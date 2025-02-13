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
sudo yum install -y yum-utils device-mapper-persistent-data lvm2 nfs-utils curl yum-plugin-versionlock

# 添加Kubernetes的密钥和源：
#curl -fsSL https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg | sudo rpm --import -
sudo rpm --import https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg
sudo rpm --import https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg

#echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
# 将 Kubernetes 的 APT 源添加到系统的软件源列表中，以便后续通过这个源安装 Kubernetes 软件包。
# baseurl：指定了 Kubernetes 仓库的 URL，适用于 CentOS 7。
#enabled=1：启用该仓库。
#gpgcheck=1：启用 GPG 签名检查。
#repo_gpgcheck=1：启用仓库的 GPG 密钥验证。
#gpgkey：指定了用于验证软件包签名的 GPG 密钥 URL。
cat <<EOF | sudo tee /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-\$basearch
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg
EOF


# 安装 Kubernetes 工具，包括 kubeadm、kubelet 和 kubectl。版本号是指定安装的具体版本
sudo yum install -y kubeadm-1.23.3-0 kubelet-1.23.3-0 kubectl-1.23.3-0
# 已安装的 kubeadm、kubelet 和 kubectl 标记为"hold"状态，
# 防止它们在系统更新时被升级到新的版本。
sudo yum versionlock kubeadm kubelet kubectl

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
