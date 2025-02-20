# prepare and docker

```shell
# 更新系统及依赖
sudo yum update -y
sudo yum install -y curl wget net-tools
sudo yum install -y yum-utils device-mapper-persistent-data lvm2
sudo vi /etc/hostname # 改主机名
```

```shell
# 更换为阿里云或其他镜像源，阿里云提供了 Docker 镜像仓库，可以使用以下命令
# 添加 Docker 官方存储库
#  安装 Docker CE
sudo yum-config-manager --add-repo https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
#sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
#sudo yum install -y docker-ce docker-ce-cli containerd.io
# 指定版本
sudo yum install -y docker-ce-20.10.24-3.el7 docker-ce-cli-20.10.24-3.el7 containerd.io 


```

```bash
# 启动 Docker 服务
sudo systemctl enable docker
sudo systemctl start docker

#运行以下命令验证 Docker 是否正常工作：

docker --version
sudo docker run hello-world
```

```shell
# fix docker issue
# 这个命令将一个Docker配置文件 daemon.json 写入 /etc/docker/ 目录中。
# 配置内容包括使用 systemd 作为 cgroup 驱动、
# 设置日志驱动为 json-file、
# 日志最大大小为100MB、
# 存储驱动为 overlay2。
# 更换镜像来源
cat <<EOF | sudo tee /etc/docker/daemon.json
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2",
  "registry-mirrors": ["https://ed9tlwpf.mirror.aliyuncs.com"]
}
EOF

# 这部分命令启用Docker服务，重新加载systemd配置，并重新启动Docker服务，确保配置生效。
sudo systemctl enable docker
sudo systemctl daemon-reload
sudo systemctl restart docker


# iptables
# 将 br_netfilter 写入 /etc/modules-load.d/k8s.conf 文件，
# 确保该内核模块在系统启动时加载。
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF

# 将一些网络相关的内核参数写入 /etc/sysctl.d/k8s.conf 文件，
# 启用了iptables的bridge调用规则和IPv4转发。
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward=1 # better than modify /etc/sysctl.conf
EOF

# 这个命令重新加载sysctl配置，确保新的内核参数生效。
sudo sysctl --system

# Disable Swap
# 这部分命令关闭所有交换分区，
# 并通过sed命令注释掉 /etc/fstab 文件中与交换分区相关的行，
# 确保在系统重新启动时交换分区不再被激活。
sudo swapoff -a
sudo sed -i '/\sswap\s/s/^#?/#/' /etc/fstab

# check

echo "please check these files:"
echo "/etc/docker/daemon.json"
echo "/etc/modules-load.d/k8s.conf"
echo "/etc/sysctl.d/k8s.conf"
echo "cat cat /etc/fstab"

```

# Kubernetes 安装

```shell
#sudo apt update
# 这两个命令用于更新系统的包列表和安装一些基本的软件包，包括支持 HTTPS 传输的软件包、证书、curl 工具以及 NFS 共享的支持。
sudo apt install -y apt-transport-https ca-certificates curl nfs-common

#sudo curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg
# 添加Kubernetes的APT密钥和源：
# 从阿里云的镜像站下载 Kubernetes 的 APT 密钥，并添加到系统的密钥环中。
curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | sudo apt-key add -

#echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
# 将 Kubernetes 的 APT 源添加到系统的软件源列表中，以便后续通过这个源安装 Kubernetes 软件包。
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main
EOF

# 更新系统的软件源
sudo apt update
# 安装 Kubernetes 工具，包括 kubeadm、kubelet 和 kubectl。版本号是指定安装的具体版本
sudo apt install -y kubeadm=1.23.3-00 kubelet=1.23.3-00 kubectl=1.23.3-00
# 已安装的 kubeadm、kubelet 和 kubectl 标记为"hold"状态，
# 防止它们在系统更新时被升级到新的版本。
sudo apt-mark hold kubeadm kubelet kubectl

# check
kubeadm version
kubectl version 
kubectl version --client

```

# Kubernetes 镜像

```shell
# use ali registry to speed up
# 使用阿里云容器镜像服务（Aliyun Container Registry）
# 来加速 Kubernetes 集群所需的容器镜像的下载
repo=registry.aliyuncs.com/google_containers
# kubeadm config images list 获取指定 Kubernetes 版本所需的官方容器镜像列表
# 拉取并修改 Kubernetes 官方镜像标签
# 去除镜像名称中的前缀，然后使用 docker pull 从阿里云容器镜像服务拉取镜像。
for name in `kubeadm config images list --kubernetes-version v1.23.3`;
do
    # 移除镜像名称中的前缀
    src_name=${name#k8s.gcr.io/}
    src_name=${src_name#coredns/}

    # 从阿里云容器镜像服务拉取镜像
    docker pull $repo/$src_name

    # 修改标签以符合 Kubernetes 需求
    docker tag $repo/$src_name $name

    # 删除原始的阿里云容器镜像
    docker rmi $repo/$src_name
done

## 拉取 Flannel 容器镜像
#for name in `grep image flannel.yml |grep -v '#image' | sed 's/image://g' -`;
#do
#    docker pull $name
#done

#(在能够访问外网的机器下载)
docker pull --platform linux/amd64 rancher/mirrored-flannelcni-flannel-cni-plugin:v1.0.1
docker pull --platform linux/amd64 rancher/mirrored-flannelcni-flannel:v0.17.0
# 这两个镜像需要 在自己本地下载后。导入到服务器中，离线安装
# 检查拉取的容器镜像
docker images
```

# init

```shell

# init k8s
# --apiserver-advertise-address=192.168.10.210
# --pod-network-cidr=10.10.0.0/16 指定了 Pod 网络的 CIDR 地址，这是用于容器之间通信的 IP 地址范围。
# --kubernetes-version=v1.23.3 指定了 Kubernetes 版本。
# --v=5 设置详细的日志级别，有助于调试。
sudo kubeadm init \
    --pod-network-cidr=10.10.0.0/16 \
    --kubernetes-version=v1.23.3 \
    --v=5

# enable kubectl
# 创建存放 kubectl 配置的目录
mkdir -p $HOME/.kube
# 将集群的管理员配置复制到用户的 ~/.kube/config 文件中。
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
# 设置 ~/.kube/config 文件的所有者，确保用户有权访问该文件。
sudo chown $(id -u):$(id -g) $HOME/.kube/config

#在单节点场景中，您需要移除控制平面节点上的 NoSchedule 污点：

kubectl taint nodes --all node-role.kubernetes.io/control-plane-
#或（根据版本可能使用 master）：
kubectl taint nodes --all node-role.kubernetes.io/master-
# check
kubectl version
kubectl get node
```

# 网络通信

```shell
kubectl apply -f kube-flannel.yml
```

# 安装常用工具

## helm

```shell
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
helm version
```

### 常用国内仓库

```shell
helm repo add bitnami "https://helm-charts.itboon.top/bitnami" --force-update
helm repo add grafana "https://helm-charts.itboon.top/grafana" --force-update
helm repo add prometheus-community "https://helm-charts.itboon.top/prometheus-community" --force-update
helm repo add ingress-nginx "https://helm-charts.itboon.top/ingress-nginx" --force-update
helm repo add aliyunstable "https://kubernetes.oss-cn-hangzhou.aliyuncs.com/charts"
helm repo add kubernetes-dashboard	"https://kubernetes.github.io/dashboard/"

helm repo update

```

## kubernetes-dashboard

```shell
helm install kubernetes-dashboard kubernetes-dashboard/kubernetes-dashboard --namespace kubernetes-dashboard --create-namespace --version 5.4.1
```

可能出现镜像拉取失败的

```shell
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/dashboard:v2.5.1
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-webhook-certgen:v1.5.1
```

## kube-prometheus-stack

```shell
helm install kube-prometheus-stack prometheus-community/kube-prometheus-stack --namespace prometheus --create-namespace
```

### 在能够访问外网的服务器下载，然后导入 docker load -i

```shell
docker pull --platform linux/amd64 docker.io/grafana/grafana:11.5.1
docker pull --platform linux/amd64 registry.k8s.io/kube-state-metrics/kube-state-metrics:v2.14.0
```

## 访问Dashboard

```shell
Get the Kubernetes Dashboard URL by running:
export POD_NAME=$(kubectl get pods -n kubernetes-dashboard -l "app.kubernetes.io/name=kubernetes-dashboard,app.kubernetes.io/instance=kubernetes-dashboard" -o jsonpath="{.items[0].metadata.name}")
echo https://127.0.0.1:8443/
kubectl -n kubernetes-dashboard port-forward $POD_NAME 8443:8443 
```

在本地访问部署在 Kubernetes 集群中的 Kubernetes Dashboard。以下是对每个命令的解释：

1. **获取 Kubernetes Dashboard Pod 的名称**：
   ```bash
   export POD_NAME=$(kubectl get pods -n kubernetes-dashboard -l "app.kubernetes.io/name=kubernetes-dashboard,app.kubernetes.io/instance=kubernetes-dashboard" -o jsonpath="{.items[0].metadata.name}")
   ```
    - **作用**：通过指定命名空间（`kubernetes-dashboard`）和标签筛选器，获取 Kubernetes Dashboard Pod 的名称，并将其存储在环境变量 `POD_NAME` 中。

2. **显示本地访问的 URL**：
   ```bash
   echo https://127.0.0.1:8443/
   ```
    - **作用**：输出本地访问 Dashboard 的 URL。

3. **设置端口转发**：
   ```bash
   kubectl -n kubernetes-dashboard port-forward $POD_NAME 8443:8443
   ```
    - **作用**：使用 `kubectl port-forward` 命令，将本地机器的 `8443` 端口请求转发到 Kubernetes 集群中指定 Pod 的 `8443` 端口。
    - **详细说明**：
        - `-n kubernetes-dashboard`：指定操作的命名空间。
        - `port-forward`：`kubectl` 的子命令，用于建立本地与 Pod 之间的端口转发。
        - `$POD_NAME`：之前获取的 Kubernetes Dashboard Pod 的名称。
        - `8443:8443`：将本地的 `8443` 端口映射到 Pod 的 `8443` 端口。

**总结**：通过上述命令，您可以在本地通过 `https://127.0.0.1:8443/` 访问部署在 Kubernetes 集群中的 Kubernetes
Dashboard，而无需直接暴露服务到外部网络。这对于在本地管理和监控集群非常方便。

# glusterFs

## base

```shell
sudo apt install -y glusterfs-server glusterfs-client
sudo systemctl start glusterd
sudo systemctl enable glusterd
sudo systemctl status glusterd
# 创建volume
gluster volume create pcb-test 172.16.114.51:/data/gluster/test_1 172.16.114.51:/data/gluster/test_2 force
gluster volume start pcb-test
gluster volume info
gluster volume status pcb-test
```

k8s 相关的服务

```shell
kubectl apply -f endpint.yaml
kubectl apply -f service.yaml
kubectl apply -f pv.yaml
kubectl apply -f pvc.yaml

需要验证 endpoint service pv pvc 全部正常
```

## kuberay-operator and  ray-cluster chart

镜像

```shell
docker pull --platform linux/amd64 kuberay/operator:v0.5.2
docker pull --platform linux/amd64 fluent/fluent-bit:2.0.9
```

```shell
helm install kuberay-operator ~/charts/kuberay-operator-0.0.1.tgz
helm install ray-cluster ray-cluster-0.0.2.tgz
```

### **你的 Pod 处于 `Pending` 状态，原因：**

```
Warning  FailedScheduling  0/1 nodes are available: 1 node(s) didn't match Pod's node affinity/selector.
```

✅ **这个错误说明 `Pod` 受到了 `nodeSelector` 或 `nodeAffinity` 规则的限制，无法调度到任何节点**。

---

## **1. 确认 `nodeSelector` 限制**

你的 Pod `ray-cluster-kuberay-head-82ws5` 里面有：

```yaml
Node-Selectors: raycluster=enable
```

这意味着：

- 你的 Pod **只能调度到带有 `raycluster=enable` 标签的节点**。
- **如果当前节点没有 `raycluster=enable`，Pod 会一直 `Pending`，无法调度**。

### **检查集群节点是否有 `raycluster=enable`**

```bash
kubectl get nodes --show-labels
```

✅ **如果 `LABELS` 里没有 `raycluster=enable`，说明没有合适的节点来运行 Pod**。

---

## **2. 解决方案**

### ✅ **方法 1：给节点添加 `raycluster=enable` 标签**

如果 `kubectl get nodes --show-labels` 发现 **没有 `raycluster=enable` 标签的节点**，你可以手动给节点打标签：

```bash
kubectl label node <你的节点名称> raycluster=enable
```

然后检查：

```bash
kubectl get nodes --show-labels
```

如果 `raycluster=enable` 出现在 **你的目标节点** 上，Pod 会自动被调度。
---

## **3. 结论**

✅ **你的 Pod 由于 `nodeSelector: raycluster=enable` 限制，无法找到符合要求的节点。**  
✅ **你可以给节点手动添加 `raycluster=enable` 标签，或者删除 `nodeSelector` 让 Pod 可以调度到任意节点。**

# mysql

镜像

```shell
docker pull --platform linux/amd64 mysql:5.7.30
docker pull --platform linux/amd64 prom/mysqld-exporter:v0.12.1
```

## 创建 pv和storageClass

```shell
kubectl apply -f mysql-storage.yaml
kubectl apply -f mysql_pv.yaml
```

```shell
helm install mysql ./mysql 
#helm install mysql mysql-0.0.2.tgz

kubectl exec -it mysql-0  -- mysql -u root -p
```

# pcb-vrs

## 目录挂载

```shell
mkdir /mnt/glusterfs/pcb-test
mount -t glusterfs 172.16.114.51:/pcb-test /mnt/glusterfs/pcb-test
```

## 自动挂载

将下面内容添加到 /etc/fstab 文件中。

```shell
172.16.114.51:/pcb-test /mnt/glusterfs/pcb-test glusterfs defaults 0 0
```

需要注意数据库的配置关系

# ingress-controller 
```shell
docker pull --platform linux/amd64 k8s.gcr.io/ingress-nginx/controller:v1.0.0
docker pull --platform linux/amd64 k8s.gcr.io/ingress-nginx/kube-webhook-certgen:v1.0
docker pull --platform linux/amd64 k8s.gcr.io/defaultbackend-amd64:1.5
```


ingress file
ray-cluster ingress
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
   name: ray-cluster-ingress
   namespace: ivip
   annotations:
      kubernetes.io/ingress.class: nginx-pcb
      nginx.ingress.kubernetes.io/rewrite-target: /$2
spec:
   rules:
      - http:
           paths:
              - path: /raycluster(/|$)(.*)
                pathType: ImplementationSpecific
                backend:
                   service:
                      name: ray-cluster-kuberay-head-svc
                      port:
                         number: 8265
              - path: /static(.*)
                pathType: ImplementationSpecific
                backend:
                   service:
                      name: ray-cluster-kuberay-head-svc
                      port:
                         number: 8265

```










