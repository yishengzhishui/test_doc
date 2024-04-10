### pod 常用命令

以下是一些常用的 `kubectl` 命令，用于操作 Pod：

1. **查看 Pod：**

   - `kubectl get pods`：显示所有 Pod。
   - `kubectl get pods -n [命名空间]`：显示特定命名空间中的所有 Pod。
2. **查看 Pod 详细信息：**

   - `kubectl describe pod [Pod名称]`：显示特定 Pod 的详细信息。
3. **查看 Pod 日志：**

   - `kubectl logs [Pod名称]`：查看 Pod 的日志。
   - `kubectl logs -f [Pod名称]`：实时跟踪 Pod 的日志。
4. **进入 Pod 内部：**

   - `kubectl exec -it [Pod名称] -- /bin/bash`：在 Pod 内部启动交互式 shell。
5. **在 Pod 中运行命令：**

   - `kubectl exec [Pod名称] -- [命令]`：在 Pod 内部执行单次命令，例如 `kubectl exec mypod -- ls /app`。
6. **删除 Pod：**

   - `kubectl delete pod [Pod名称]`：删除特定 Pod。
   - `kubectl delete pods --all`：删除所有 Pod。
7. **查看 Pod 中的环境变量：**

   - `kubectl exec [Pod名称] -- env`：查看 Pod 中运行的容器的环境变量。
8. **查看 Pod 中的资源使用情况：**

   - `kubectl top pod [Pod名称]`：查看特定 Pod 中的 CPU 和内存使用情况。
9. **复制文件到/从 Pod：**

   - `kubectl cp [本地文件路径] [Pod名称]:[目标路径]`：将本地文件复制到 Pod 内部。
   - `kubectl cp [Pod名称]:[源路径] [本地目标路径]`：从 Pod 内部复制文件到本地。
10. **设置当前上下文的默认命名空间**：

    ```bash
    kubectl config set-context --current --namespace=<命名空间>
    ```

    将 `<命名空间>` 替换为您要使用的命名空间的名称。
11. **获取 Pod**：

    ```bash
    kubectl get pods
    ```

    现在，您可以直接运行 `kubectl get pods` 命令，它将默认获取您设置的命名空间中的 Pod。

请注意，通过这种方式设置的命名空间默认值在当前会话中仍然有效，直到您通过其他命令更改它或者结束当前会话。

1. **获取资源列表**：

   - `kubectl get pods`：获取 Pod 列表。
   - `kubectl get deployments`：获取 Deployment 列表。
   - `kubectl get services`：获取 Service 列表。
   - `kubectl get nodes`：获取 Node 列表。
   - `kubectl get namespaces`：获取 Namespace 列表。
2. **查看资源详情**：

   - `kubectl describe pod <pod名称>`：查看 Pod 的详细信息。
   - `kubectl describe deployment <deployment名称>`：查看 Deployment 的详细信息。
   - `kubectl describe service <service名称>`：查看 Service 的详细信息。
3. **创建、删除和编辑资源**：

   - `kubectl create -f <yaml文件>`：根据 YAML 文件创建资源。
   - `kubectl delete <资源类型> <资源名称>`：删除资源。
   - `kubectl edit <资源类型> <资源名称>`：编辑资源。
4. **查看日志和执行命令**：

   - `kubectl logs <pod名称>`：查看 Pod 的日志。
   - `kubectl exec -it <pod名称> -- <命令>`：在 Pod 中执行命令。
5. **调试和故障排除**：

   - `kubectl get events`：查看集群事件。
   - `kubectl describe <资源类型> <资源名称>`：查看资源的详细描述以进行故障排除。
   - `kubectl logs -p <pod名称>`：查看 Pod 中上一个容器的日志。
6. **扩展和缩减资源**：

   - `kubectl scale deployment <deployment名称> --replicas=<副本数>`：调整 Deployment 的副本数。
   - `kubectl autoscale deployment <deployment名称> --min=<最小副本数> --max=<最大副本数> --cpu-percent=<CPU利用率百分比>`：根据 CPU 使用情况自动缩放 Deployment 的副本数。
7. **查看集群信息**：

   - `kubectl cluster-info`：查看集群信息，包括 API 服务器的地址和版本信息。
   - `kubectl version`：查看 Kubernetes 客户端和服务器的版本信息。
8. **管理存储**：

   - `kubectl get pv`：获取持久卷（PersistentVolume）列表。
   - `kubectl get pvc`：获取持久卷声明（PersistentVolumeClaim）列表。
9. **管理网络**：

   - `kubectl get ingress`：获取 Ingress 资源列表，用于管理集群的入口流量。
10. **管理配置**：

- `kubectl get configmaps`：获取 ConfigMap 列表。
- `kubectl get secrets`：获取 Secret 列表。

这些命令涵盖了 Kubernetes 中更多的方面，包括资源的扩展、集群信息的查看、存储和网络的管理以及配置的管理。使用这些命令，您可以更好地管理和监控 Kubernetes 集群。

label
```shell
kubectl get pod -l app=nginx
kubectl get pod -l 'app in (ngx, nginx, ngx-dep)'

```

`grep` 是一个非常强大和常用的 Linux/Unix 命令，用于在文本中搜索指定的模式。以下是一些常用的 `grep` 命令及其说明：

1. **基本搜索**：
   ```bash
   grep "pattern" file
   ```
   在文件中搜索指定的模式，并输出包含该模式的所有行。

2. **忽略大小写搜索**：
   ```bash
   grep -i "pattern" file
   ```
   在搜索时忽略大小写。

3. **显示匹配行的行号**：
   ```bash
   grep -n "pattern" file
   ```
   显示匹配模式的行号。

4. **显示不匹配的行**：
   ```bash
   grep -v "pattern" file
   ```
   显示不包含指定模式的所有行。

5. **显示匹配模式的行数**：
   ```bash
   grep -c "pattern" file
   ```
   统计包含指定模式的行数。

6. **显示匹配模式的前后行**：
   ```bash
   grep -A <num> "pattern" file  # 显示匹配行及其后 <num> 行
   grep -B <num> "pattern" file  # 显示匹配行及其前 <num> 行
   grep -C <num> "pattern" file  # 显示匹配行及其前后各 <num> 行
   ```
   显示包含匹配模式的行及其周围的行，`<num>` 是行数。

7. **递归搜索目录**：
   ```bash
   grep -r "pattern" directory
   ```
   在指定目录及其子目录中递归搜索指定模式。

这些是一些常用的 `grep` 命令，能够满足大多数搜索文本的需求。

要查询 Pod 是否有异常日志，您可以使用 `kubectl logs` 命令结合管道 (`|`) 和 `grep` 工具来筛选出包含异常信息的日志行。例如，您可以使用以下命令：

```bash
kubectl logs <pod名称> | grep -i "error\|exception"
```

这个命令将从指定的 Pod 日志中筛选出包含 "error" 或 "exception" 关键字的行。您可以根据需要修改关键字或正则表达式以适应您的特定情况。

请注意，`-i` 选项用于忽略大小写，这样无论日志中的关键字是大写还是小写，都能匹配到。