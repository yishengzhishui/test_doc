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

# Gluster的存储卷类型

在 GlusterFS 中，可以创建多种类型的卷，每种类型适用于不同的场景和需求。这些卷类型是通过 GlusterFS 的底层机制（例如分布式存储、复制存储等）实现的。以下是 GlusterFS 支持的主要卷类型及其特点：

---

### 1. **分布式卷 (Distributed Volume)**

#### 特点：

- 默认的卷类型。
- 文件被分布式存储在多个节点上。
- 不提供冗余保护（即不复制文件），如果一个节点或磁盘故障，对应的文件将丢失。
- 适合存储对可靠性要求不高的大量数据。

#### 创建命令：

```bash
gluster volume create <volume-name> <node1>:/data <node2>:/data ...
```

#### 适用场景：

- 非关键性数据。
- 对性能和存储效率要求高，但对数据冗余保护要求低的场景。

---

### 2. **复制卷 (Replicated Volume)**

#### 特点：

- 文件被复制到多个节点，提供数据冗余。
- 如果一个节点或磁盘发生故障，数据仍然可用。
- 适合需要高数据可靠性的场景。
- 可以指定副本数量（例如 2 或 3 副本）。

#### 创建命令：

```bash
gluster volume create <volume-name> replica <replica-count> <node1>:/data <node2>:/data ...
```

#### 适用场景：

- 关键性数据存储。
- 需要高可靠性和高可用性的场景。

---

### 3. **分布式复制卷 (Distributed Replicated Volume)**

#### 特点：

- 分布式和复制的结合。
- 文件被分布在多个节点上，同时每个文件在多个节点间有冗余副本。
- 提供可靠性（复制）和扩展性（分布式）。
- 副本的数量由 `replica` 参数决定。

#### 创建命令：

```bash
gluster volume create <volume-name> replica <replica-count> <node1>:/data <node2>:/data <node3>:/data ...
```

#### 适用场景：

- 既需要高可靠性又需要扩展性的场景。
- 例如，分布式日志存储或分布式数据库。

---

### 4. **条带化卷 (Striped Volume)**

#### 特点：

- 文件被切分为固定大小的数据块，并条带化存储在多个节点上。
- 提高文件的读取和写入性能。
- 不提供冗余保护（类似 RAID 0）。

#### 创建命令：

```bash
gluster volume create <volume-name> stripe <stripe-count> <node1>:/data <node2>:/data ...
```

#### 适用场景：

- 高性能要求的大文件存储，例如多媒体处理、科学计算等。
- 文件较大且需要高并发访问。

---

### 5. **分布式条带化卷 (Distributed Striped Volume)**

#### 特点：

- 分布式和条带化的结合。
- 文件被分布在多个节点上，并在每个文件上采用条带化存储。
- 提供扩展性和高性能，但不提供冗余保护。

#### 创建命令：

```bash
gluster volume create <volume-name> stripe <stripe-count> <node1>:/data <node2>:/data <node3>:/data ...
```

#### 适用场景：

- 高性能要求的大文件存储。
- 数据对冗余保护需求较低，但对扩展性和性能需求高的场景。

---

### 6. **条带化复制卷 (Striped Replicated Volume)**

#### 特点：

- 文件数据被切分为条带化块，每个块在多个节点间复制。
- 提供条带化的高性能和复制的可靠性。

#### 创建命令：

```bash
gluster volume create <volume-name> stripe <stripe-count> replica <replica-count> <node1>:/data <node2>:/data ...
```

#### 适用场景：

- 高性能与高可靠性同时要求的场景。
- 存储大型关键数据，例如视频流或数据库。

---

### 7. **分布式条带化复制卷 (Distributed Striped Replicated Volume)**

#### 特点：

- 分布式、条带化和复制的结合。
- 文件被分布在多个节点上，并条带化存储，同时每个块有多个副本。
- 提供可靠性、高性能和扩展性。

#### 创建命令：

```bash
gluster volume create <volume-name> stripe <stripe-count> replica <replica-count> <node1>:/data <node2>:/data ...
```

#### 适用场景：

- 高可靠性、高性能和扩展性需求并存的场景。

---

### 8. **Arbiter Volume (仲裁卷)**

#### 特点：

- 是复制卷的一种优化形式，用于减少存储空间消耗。
- 第三个副本只存储文件的元数据，而不是完整的文件内容。
- 提供可靠性，同时减少存储开销。

#### 创建命令：

```bash
gluster volume create <volume-name> replica 3 arbiter 1 <node1>:/data <node2>:/data <node3>:/data ...
```

#### 适用场景：

- 数据可靠性要求高，同时存储成本敏感的场景。

---

### 9. **Geo-replication Volume (地理复制卷)**

#### 特点：

- 通过跨数据中心或地域复制，实现灾备。
- 异步复制，保证数据在多个位置保持一致。
- 提供容灾能力，但性能不如同步复制。

#### 创建命令：

```bash
gluster volume geo-replication <master-vol> <remote-host>::<remote-vol> start
```

#### 适用场景：

- 跨区域的容灾备份和数据同步场景。

---

### 总结

| 卷类型                 | 特点                                   | 适用场景                   |
| ------------------------ | ---------------------------------------- | ---------------------------- |
| **分布式卷**           | 分布存储，无冗余保护                   | 非关键性数据存储           |
| **复制卷**             | 数据冗余，可靠性高                     | 关键数据存储               |
| **分布式复制卷**       | 分布式+复制，可靠性与扩展性兼具        | 既要高扩展性又要高可靠性   |
| **条带化卷**           | 高性能，但无冗余保护                   | 处理大文件的高性能场景     |
| **分布式条带化卷**     | 分布式+条带化，高性能和扩展性          | 高性能需求的大型数据存储   |
| **条带化复制卷**       | 条带化+复制，高性能和可靠性            | 高性能和高可靠性需求       |
| **分布式条带化复制卷** | 分布式+条带化+复制，可靠性和扩展性兼具 | 综合需求场景               |
| **仲裁卷**             | 节省存储空间的复制卷                   | 存储成本敏感的高可靠性场景 |
| **地理复制卷**         | 跨地域复制，提供容灾能力               | 跨数据中心的容灾备份       |

选择适合的卷类型取决于你的性能、可靠性和存储成本需求。

# 复制卷 (Replicated Volume) 和 分布式复制卷 (Distributed Replicated Volume)

**复制卷 (Replicated Volume)** 和 **分布式复制卷 (Distributed Replicated Volume)** 是 GlusterFS 中两种不同的卷类型，它们的主要区别在于文件的分布方式和适用场景：

---

### 1. **复制卷 (Replicated Volume)**

#### 特点

- **每个文件**会被完全复制到所有副本节点上。
- 数据冗余：副本数决定了数据的冗余级别（例如，`replica 2` 表示每个文件有两个副本）。
- **所有文件**存储在相同的副本节点组中，没有分布机制。
- 高可靠性：如果某个节点或磁盘发生故障，数据仍然可用。

#### 结构示意图

假设有 3 个节点，每个文件都会被复制到所有 3 个节点上。

```
节点1       节点2       节点3
 file1      file1      file1
 file2      file2      file2
 file3      file3      file3
```

#### 创建命令

```bash
gluster volume create <volume-name> replica 3 <node1>:/data <node2>:/data <node3>:/data
```

#### 适用场景

- 数据可靠性要求高的场景，例如关键性数据存储。
- 文件数量不多，但每个文件需要高可用性的场景。

#### 缺点

- 扩展性差：因为每个文件都存储在所有副本节点上，增加节点不会增加存储容量。
- 存储效率低：实际可用存储容量为单节点容量的 `1/N`，其中 `N` 是副本数。

---

### 2. **分布式复制卷 (Distributed Replicated Volume)**

#### 特点

- 结合了分布式卷和复制卷的特点。
- **文件被分布式存储**在多个副本组上，每个副本组内部使用复制机制。
- 提供数据冗余的同时增加了扩展性：副本组数量决定了存储容量的扩展能力。
- 每个副本组内的节点数决定了冗余级别。

#### 结构示意图

假设有 6 个节点，每个副本组有 3 个节点（`replica 3`），文件分布在不同的副本组中。

```
节点1       节点2       节点3       节点4       节点5       节点6
 file1      file1      file1       file2      file2      file2
 file3      file3      file3       file4      file4      file4
```

- 文件 `file1` 被复制到副本组 1（节点 1、2、3）。
- 文件 `file2` 被复制到副本组 2（节点 4、5、6）。

#### 创建命令

```bash
gluster volume create <volume-name> replica 3 <node1>:/data <node2>:/data <node3>:/data \
<node4>:/data <node5>:/data <node6>:/data
```

#### 适用场景

- 数据可靠性要求高，同时需要较大存储容量的场景。
- 文件数量多，分布式访问的性能要求较高的场景。

#### 优点

- **扩展性强**：增加副本组可以同时提高存储容量和吞吐量。
- **高可靠性**：每个副本组提供数据冗余，防止单节点或单副本组故障。

#### 缺点

- 管理复杂度高：需要维护多个副本组。
- 存储效率：每个副本组内仍然存在冗余，因此实际可用容量是单副本组容量的 `1/N`，其中 `N` 是副本数。

---

### 对比总结

| 特性             | **复制卷**                   | **分布式复制卷**               |
| ------------------ | ------------------------------ | -------------------------------- |
| **文件分布方式** | 每个文件复制到所有副本节点上 | 文件分布到不同副本组，组内复制 |
| **可靠性**       | 高（由副本数决定）           | 高（每个副本组提供冗余）       |
| **扩展性**       | 差（节点数增加不会增加容量） | 强（增加副本组可扩展容量）     |
| **存储效率**     | 单节点容量的`1/N`            | 单副本组容量的`1/N`            |
| **适用场景**     | 小型集群，关键数据存储       | 大型集群，可靠性和扩展性要求高 |
| **创建命令**     | `replica <副本数>`           | `replica <副本数> <多节点>`    |

---

### 选择建议

- **使用复制卷**：
    - 当存储容量需求较小，数据可靠性要求高时。
    - 适用于关键数据存储的小型集群。
- **使用分布式复制卷**：
    - 当需要在多个节点之间分布存储数据，同时保证可靠性时。
    - 适用于大规模集群、分布式访问的场景，例如日志存储、分布式文件系统等。

# 卷类型更改 分布式改成复制卷

在 GlusterFS 中，**已经创建的卷无法直接更改类型**（例如，从分布式卷改为复制卷）。这是因为卷的类型决定了底层文件存储的逻辑结构（例如，如何分布和复制文件），这些结构在卷创建时就固定下来，不能动态修改。

如果需要更改卷类型，通常需要以下步骤：

---

### 1. **备份数据**

更改卷类型的唯一方法是创建一个新卷，然后将数据从旧卷迁移到新卷。这意味着需要先备份当前数据。

#### 方法：

- 如果卷已挂载到某个挂载点，例如 `/mnt/glusterfs`，可以使用工具如 `rsync` 或 `cp` 将数据备份到其他存储位置：
  ```bash
  rsync -avz /mnt/glusterfs /backup/location/
  ```

---

### 2. **停止现有卷**

为了避免数据不一致，先停止旧卷：

```bash
gluster volume stop <volume-name>
```

---

### 3. **删除旧卷**

删除当前卷以释放节点资源：

```bash
gluster volume delete <volume-name>
```

---

### 4. **创建新卷**

根据需求，重新创建卷。例如：

#### 将分布式卷改为复制卷：

假设原先有一个分布式卷，现在需要将其改为 **复制卷**（`replica 2`）。

```bash
gluster volume create <new-volume-name> replica 2 <node1>:/data <node2>:/data
```

#### 启动新卷：

```bash
gluster volume start <new-volume-name>
```

---

### 5. **恢复数据**

将备份的数据复制回新创建的卷：

```bash
rsync -avz /backup/location/ /mnt/glusterfs
```

---

### 6. **验证新卷**

确认新卷类型和数据是否正常。

#### 查看卷类型：

```bash
gluster volume info <new-volume-name>
```

#### 验证挂载和数据完整性：

```bash
mount -t glusterfs <node>:/<new-volume-name> /mnt/glusterfs
ls /mnt/glusterfs
```

---

### 7. **自动化迁移（高级）**

如果涉及大量数据，可以使用 GlusterFS 的 Geo-replication 功能实现数据自动迁移。

#### 配置 Geo-replication：

1. 在新旧卷之间设置异步复制关系。
2. 完成复制后切换到新卷。

---

### 为什么不能直接更改卷类型？

1. **文件分布方式不同**：

- 分布式卷：文件分散存储在多个节点上。
- 复制卷：每个文件有多个副本存储在不同节点上。
- 分布式复制卷：文件分布在多个副本组中，每个副本组内有多个副本。 更改卷类型会导致现有文件无法适应新结构。

2. **数据一致性和完整性**： 如果动态修改卷类型，可能会破坏现有数据的冗余关系，导致数据丢失或无法恢复。

---

### 总结

- GlusterFS 不支持直接更改卷类型。
- 解决方案是：
    1. **备份数据**。
    2. **删除旧卷并创建新卷**。
    3. **将数据恢复到新卷**。
- 使用 Geo-replication 功能可以简化大规模数据迁移的过程。

## 挂载验证

在 Kubernetes 中使用 GlusterFS 时，建议先测试挂载。可以通过创建 Pod 并挂载 GlusterFS 卷验证是否工作正常。

**测试挂载 YAML 文件**：

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: gluster-test
  namespace: ivip
spec:
  containers:
    - name: test-container
      image: busybox
      command: [ "/bin/sh", "-c", "while true; do sleep 3600; done" ]
      volumeMounts:
        - name: glusterfs-volume
          mountPath: /mnt/glusterfs
  volumes:
    - name: glusterfs-volume
      glusterfs:
        endpoints: glusterfs-cluster
        path: ivip-pcb-output
        readOnly: false
```

部署后进入容器查看挂载点：

```bash
kubectl exec -it gluster-test -n ivip -- /bin/sh
ls /mnt/glusterfs
```

## 变更卷后重新进行挂载到pod

如果是通过这个 YAML 文件挂载 GlusterFS 卷的，那么你在进行 GlusterFS 卷类型变换时，操作步骤与之前相似，关键步骤是停止 Pod，进行卷类型更改，然后再恢复 Pod。下面是针对这个 YAML 文件挂载的具体操作流程。

### 操作步骤：

#### 1. **查看和确认 Pod 状态**

首先，确认你所挂载 GlusterFS 卷的 Pod 是否正在运行。使用以下命令查看 Pod 状态：

```bash
kubectl get pods -n ivip
```

确认挂载 GlusterFS 卷的 Pod `gluster-test` 是否在运行。

#### 2. **停止 Pod**

为了安全地更改卷类型，你需要先停止相关的 Pod。可以通过以下命令删除 Pod，这样 Kubernetes 会自动重启 Pod（如果它是 Deployment 或 StatefulSet 的一部分）。

```bash
kubectl delete pod gluster-test -n ivip
```

删除 Pod 后，Pod 会停止并解除与 GlusterFS 卷的挂载。

如果是通过 Deployment 或 StatefulSet 进行管理，首先需要 `scale`（缩容）应用程序，以便将 Pod 设置为 0 副本。

- **对于 Deployment**： 如果你的应用程序是通过 `Deployment` 管理的，可以通过以下命令将副本数调整为 0，停止所有 Pod：

  ```bash
  kubectl scale deployment <deployment-name> --replicas=0
  ```

- **对于 StatefulSet**： 如果是通过 `StatefulSet` 管理的，可以将副本数调整为 0：

  ```bash
  kubectl scale statefulset <statefulset-name> --replicas=0
  ```

这样可以确保所有挂载了卷的 Pod 停止，避免数据丢失或不一致。

#### 3. **停止 GlusterFS 卷**

在更改卷类型之前，最好停止 GlusterFS 卷，以确保没有正在进行的读写操作。使用以下命令停止卷：

```bash
gluster volume stop ivip-pcb-output
```

#### 4. **更改卷类型**

此时，如果你需要将 GlusterFS 卷从分布式卷更改为复制卷或其他类型，可以执行以下命令来更改卷的类型：

```bash
gluster volume set ivip-pcb-output volume-type replica 2
```

这会将 `ivip-pcb-output` 卷的类型更改为复制卷（`replica 2`）。根据需求，你可以选择不同的配置类型。

#### 5. **重新启动 GlusterFS 卷**

卷类型更改完成后，重新启动 GlusterFS 卷：

```bash
gluster volume start ivip-pcb-output
```

#### 6. **重新部署 Pod**

卷已更新并且 GlusterFS 卷类型已更改后，重新部署之前删除的 Pod。由于该 Pod 是通过 Kubernetes YAML 文件定义的，你可以直接重新创建该 Pod：

```bash
kubectl apply -f <your-pod-yaml-file>.yaml
```

- **重新启动 Deployment**：

  如果你之前使用 `Deployment` 进行管理，可以通过以下命令恢复 Pod：

  ```bash
  kubectl scale deployment <deployment-name> --replicas=<desired-replicas>
  ```

- **重新启动 StatefulSet**：

  如果你使用 `StatefulSet` 进行管理，可以通过以下命令恢复 Pod：

  ```bash
  kubectl scale statefulset <statefulset-name> --replicas=<desired-replicas>
  ```

#### 7. **验证挂载是否正常**

确认 Pod 已经成功挂载新的 GlusterFS 卷并且可以访问数据。你可以查看 Pod 的日志，确保没有错误，并且可以访问 GlusterFS 卷。

查看 Pod 状态：

```bash
kubectl get pods -n ivip
```

查看 Pod 的日志：

```bash
kubectl logs gluster-test -n ivip
```

### 总结

通过 YAML 文件挂载的 GlusterFS 卷，在更改卷类型时，操作流程如下：

1. **删除 Pod**：停止正在挂载的 Pod。
2. **停止 GlusterFS 卷**：停止正在使用的 GlusterFS 卷。
3. **更改卷类型**：修改 GlusterFS 卷的类型。
4. **启动 GlusterFS 卷**：重新启动卷以应用更改。
5. **重新创建 Pod**：使用 `kubectl apply` 或 `kubectl scale` 恢复 Pod。
6. **验证 Pod**：确保 Pod 已成功挂载新类型的 GlusterFS 卷并能正常工作。