# Centos 安装 Gluster

### 1 install

```shell
yum -y install centos-release-gluster
yum -y install glusterfs glusterfs-server glusterfs-fuse
```

### 2 启动

以下操作需要在所有需要部署 GlusterFS 的机器节点执行一遍

```shell
# 查看 GlusterFS 版本，确保服务正常工作
gluster --version

systemctl start glusterd
systemctl enable glusterd
systemctl status glusterd
```

### 3 添加 peer 节点

假设我们在一个 3 节点的 k8s 集群上进行操作，其中 node1 和 node2 是此集群的 node 节点。node1 的 IP 为 192.168.10.1，node2 的 IP 为 192.168.10.2 以下操作只需要在
k8s master 节点上执行

```shell
gluster peer probe 192.168.10.1
gluster peer probe 192.168.10.2
# 查看集群状态
gluster peer status
```

- **问题**： 如果节点之间的网络通信有问题，`gluster peer probe` 会失败。可能是防火墙或 SELinux 导致的。
- **建议**：

    - 确保节点之间的网络互通，特别是 TCP 24007 和 24008 端口。
    - 如果开启了防火墙，可以放行 GlusterFS 所需的端口：
      ```shell
      firewall-cmd --add-service=glusterfs --permanent
      firewall-cmd --reload
      ```
    - 如果使用 SELinux，确保 GlusterFS 相关权限：
      ```shell
      setsebool -P virt_sandbox_use_fusefs 1
      setsebool -P virt_use_samba 1
      ```

### 4 创建分布式卷

确保对应的目录已经建立 在 k8s master 节点上执行以下命令（下面的 192.168.10.1、192.168.10.2、192.168.10.3 分别表示3台机器 的内网IP，为示例）

```shell
gluster volume create ivip-pcb-output 192.168.10.1:/data/gluster/ivip-pcb
192.168.10.2:/data/gluster/ivip-pcb 192.168.10.3:/data/gluster/ivip-pcb
```

建议使用 replica 或 distribute 参数明确指定卷类型

```shell
gluster volume create ivip-pcb-output replica 3 192.168.10.1:/data/gluster/ivip-pcb \
192.168.10.2:/data/gluster/ivip-pcb 192.168.10.3:/data/gluster/ivip-pcb

```

### 5 启动 volume

```shell
gluster volume start ivip-pcb-output
```

### 6 验证 volume 部署成功

```shell
gluster volume info
```

### 7 查看volume状态

```shell
 gluster volume status ivip-pcb-output
```

### 8 部署 endpoint

glusterfs-cluster.json 内容如下示例，其中 subsets.addressed.ip 设置为GlusterFS集群的所有机器ip

```shell
# 部署 endpoint，在集群master执行
kubectl apply -f glusterfs-cluster.json 
```

- **验证 Endpoint 部署**： 使用以下命令验证 Endpoint 是否正常：
  ```shell
  kubectl get endpoints glusterfs-cluster -n ivip
  ```

glusterfs-cluster.json 例子：

```json
{
  "kind": "Endpoints",
  "apiVersion": "v1",
  "metadata": {
    "name": "glusterfs-cluster",
    "namespace": "ivip"
  },
  "subsets": [
    {
      "addresses": [
        {
          "ip": "192.168.9.27"
        },
        {
          "ip": "192.168.9.28"
        },
        {
          "ip": "192.168.9.29"
        }
      ],
      "ports": [
        {
          "port": 1990
        }
      ]
    }
  ]
}
```

### 9 部署service

部署 service，在集群master执行（避免ep在k8s重启时被清理）

```shell
kubectl apply -f glusterfs-cluster-service.yaml
```

- **验证 Service 部署**： 使用以下命令检查 Service 是否创建成功：
  ```shell
  kubectl get service glusterfs-cluster -n ivip
  ```

glusterfs-cluster-service.yaml 内容如下示例

```yaml
kind: Service
apiVersion: v1
metadata:
  name: glusterfs-cluster
  namespace: ivip
spec:
  ports: [ ]
  clusterIP: None
  type: ClusterIP

```
