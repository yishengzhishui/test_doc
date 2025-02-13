以下是一个完整的 Helm 使用教程，涵盖安装、基础操作和高级用法：

---

### 一、Helm 简介

Helm 是 Kubernetes 的包管理工具，通过 **Charts**（预配置的 Kubernetes 资源模板）简化应用部署。核心概念：

- **Chart**：应用模板集合
- **Release**：Chart 的运行实例
- **Repository**：Chart 存储仓库

---

### 二、安装 Helm

#### 1. 安装 CLI

```bash
# Mac (Homebrew)
brew install helm

# Linux
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

# Windows (Chocolatey)
choco install kubernetes-helm
```

#### 2. 验证安装

```bash
helm version
```

---

### 三、基础操作

#### 1. 添加仓库

```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
```

#### 2. 搜索 Chart

```bash
helm search repo nginx
```

#### 3. 安装应用

```bash
helm install my-nginx bitnami/nginx
```

#### 4. 查看 Release

```bash
helm list
```

#### 5. 卸载应用

```bash
helm uninstall my-nginx
```

---

### 四、创建自定义 Chart

#### 1. 生成 Chart 模板

```bash
helm create myapp
```

生成目录结构：

```
myapp/
├── charts/          # 子 Chart
├── Chart.yaml       # Chart 元数据
├── values.yaml      # 默认配置
└── templates/       # Kubernetes 资源模板
```

#### 2. 编辑 Chart.yaml

```yaml
apiVersion: v2
name: myapp
description: My Helm Chart
version: 0.1.0
```

#### 3. 配置 values.yaml

```yaml
replicaCount: 2
image:
  repository: nginx
  tag: "1.23.4"
```

#### 4. 编写模板 (templates/deployment.yaml)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: { { .Release.Name } }-deployment
spec:
  replicas: { { .Values.replicaCount } }
  template:
    spec:
      containers:
        - name: { { .Chart.Name } }
          image: "{{ .Values.image.repository }}:{{ .Values.image.tag }}"
```

#### 5. 本地测试

```bash
helm install --dry-run --debug myapp ./myapp
```

---

### 五、打包与发布

#### 1. 打包 Chart

```bash
helm package myapp
# 生成 myapp-0.1.0.tgz
```

#### 2. 发布到仓库

（以本地目录为例）

```bash
mkdir -p myrepo
mv myapp-0.1.0.tgz myrepo/
helm repo index myrepo/
```

---

### 六、高级操作

#### 1. 升级 Release

```bash
helm upgrade myapp ./myapp --set replicaCount=3
```

#### 2. 回滚版本

```bash
helm history myapp
helm rollback myapp 1
```

#### 3. 依赖管理

创建 `requirements.yaml`：

```yaml
dependencies:
  - name: redis
    version: "17.3.3"
    repository: "https://charts.bitnami.com/bitnami"
```

更新依赖：

```bash
helm dependency update myapp
```

#### 4. 使用 Hooks

在 templates/ 中添加注解：

```yaml
annotations:
  "helm.sh/hook": post-install
```

---

### 七、最佳实践

1. **版本控制**：Chart 与 App 版本解耦
2. **参数化配置**：通过 values.yaml 管理变量
3. **测试验证**：
   ```bash
   helm test <RELEASE_NAME>
   ```
4. **安全检查**：
   ```bash
   helm lint myapp
   ```

---

### 八、常用命令速查

| 命令 | 说明 |
|------|------|
| `helm install` | 安装 Chart |
| `helm upgrade` | 升级 Release |
| `helm uninstall` | 卸载 Release |
| `helm list` | 查看已安装 Release |
| `helm show values` | 查看 Chart 默认配置 |
| `helm get manifest` | 查看生成的 Kubernetes 资源 |

---

通过这个教程，您可以快速上手 Helm 的基础操作并了解高级用法。建议结合官方文档 (https://helm.sh/docs/) 深入实践。