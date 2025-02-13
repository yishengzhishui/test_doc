## 安装外来chart

### **方法一：直接使用 .tgz 文件安装（无需解压）**

#### 1. 将 Chart 文件传输到目标机器

将打包好的 `.tgz` 文件（如 `mychart-0.1.0.tgz`）复制到目标机器的任意目录，例如 `/home/user/charts`。

#### 2. 直接通过 .tgz 文件安装

在目标机器上执行以下命令：

```bash
# 语法
helm install <RELEASE_NAME> <CHART_PATH.tgz>

# 示例（假设 .tgz 文件在当前目录）
helm install my-app ./mychart-0.1.0.tgz

# 或者指定绝对路径
helm install my-app /home/user/charts/mychart-0.1.0.tgz
```

- **无需解压**，Helm 会自动解析 `.tgz` 文件内容。
- `<RELEASE_NAME>` 是自定义的部署实例名称（如 `my-app`）。

---

### **方法二：解压后安装（可选）**

若需要手动修改 Chart 内容（如调整 `values.yaml`），可先解压：

#### 1. 解压 .tgz 文件

```bash
tar -xzvf mychart-0.1.0.tgz
```

解压后会生成一个目录（如 `mychart/`），目录结构如下：

```
mychart/
├── Chart.yaml
├── values.yaml
├── charts/       # 子 Chart（如果有依赖项）
└── templates/    # Kubernetes 资源模板
```

#### 2. 修改配置（可选）

编辑 `values.yaml` 或其他模板文件：

```bash
vim mychart/values.yaml
```

#### 3. 从解压目录安装

```bash
helm install my-app ./mychart
```

---

### **关键注意事项**

1. **Helm 客户端必须安装**  
   目标机器需安装 Helm CLI，验证命令：
   ```bash
   helm version
   ```
   若未安装，参考 [Helm 官方安装文档](https://helm.sh/docs/intro/install/)。

2. **依赖项处理**
    - 如果 Chart 包含依赖（如 `charts/` 子目录或 `Chart.yaml` 中的 `dependencies`），直接使用 `.tgz` 文件时依赖已打包在内，无需额外操作。
    - 若解压后手动安装，需确保依赖已存在或运行更新命令：
      ```bash
      helm dependency update ./mychart
      ```

3. **配置文件覆盖**  
   安装时可通过 `--values` 或 `--set` 覆盖默认配置：
   ```bash
   # 使用自定义 values 文件
   helm install my-app ./mychart-0.1.0.tgz --values my-custom-values.yaml

   # 动态设置参数
   helm install my-app ./mychart-0.1.0.tgz --set replicaCount=3
   ```

4. **跨机器一致性**
    - 确保目标机器的 Kubernetes 集群版本与 Chart 兼容。
    - 若 Chart 使用特定存储类或插件，需预先在目标集群中配置。

---

### **总结**

- **无需解压**：直接通过 `.tgz` 文件安装更快捷。
- **需要修改配置时解压**：解压后调整内容再安装。
- **常用命令速查**：
  ```bash
  # 查看 Chart 内容（不解压）
  helm show all ./mychart-0.1.0.tgz

  # 检查 Chart 语法
  helm lint ./mychart-0.1.0.tgz

  # 卸载应用
  helm uninstall my-app
  ```

通过上述步骤，您可以轻松将打包好的 Helm Chart 部署到其他机器！

## 更新chart后重新安装 不覆盖旧版本

以下是分步操作指南，用于修改 `.tgz` 格式的 Helm Chart 内容后重新打包并安装：

---

### **完整步骤**

#### 1. **解压原始 Chart 文件**

```bash
# 创建临时目录并解压
mkdir temp-chart && cd temp-chart
tar -xzvf /path/to/original-chart.tgz
```

解压后会得到一个目录（例如 `mychart/`），目录结构如下：

```
mychart/
├── Chart.yaml
├── values.yaml
├── charts/          # 子 Chart（如果有依赖项）
└── templates/       # Kubernetes 资源模板
```

---

#### 2. **修改配置和模板**

进入解压后的 Chart 目录：

```bash
cd mychart
```

- **修改 `values.yaml`**  
  编辑默认配置：
  ```bash
  vim values.yaml  # 修改参数（如 replicaCount、image.tag 等）
  ```

- **修改模板文件**  
  进入 `templates/` 目录，编辑任意模板（如 `deployment.yaml`）：
  ```bash
  vim templates/deployment.yaml  # 调整资源定义
  ```

---

#### 3. **更新 Chart 版本（重要！）**

修改 `Chart.yaml` 中的 `version` 字段，避免与原始 Chart 冲突：

```yaml
apiVersion: v2
name: mychart
version: 0.1.1  # 递增版本号
```

---

#### 4. **验证 Chart 合法性**

```bash
# 检查语法和配置
helm lint .
```

若输出 `[INFO] Chart.yaml: icon is recommended` 等提示，表示校验通过。

---

#### 5. **重新打包 Chart**

回到 Chart 的上级目录，重新生成 `.tgz` 文件：

```bash
cd ..  # 确保在 mychart/ 的父目录中
helm package mychart
```

输出文件为 `mychart-0.1.1.tgz`（名称格式：`<Chart名称>-<版本>.tgz`）。

---

#### 6. **安装修改后的 Chart**

使用新生成的 `.tgz` 文件部署：

```bash
# 安装（可覆盖默认 values.yaml）
helm install my-release ./mychart-0.1.1.tgz

# 或指定自定义 values 文件
helm install my-release ./mychart-0.1.1.tgz -f custom-values.yaml
```

---

### **关键注意事项**

1. **依赖项处理**  
   如果 Chart 包含子 Chart（`charts/` 目录）或依赖项（`Chart.yaml` 中的 `dependencies`），需确保依赖已更新：
   ```bash
   helm dependency update mychart/  # 更新依赖到 charts/ 目录
   ```

2. **模拟安装测试**  
   在正式安装前，使用 `--dry-run` 验证渲染结果：
   ```bash
   helm install my-release ./mychart-0.1.1.tgz --dry-run --debug
   ```

3. **版本管理**
    - 每次修改后必须更新 `Chart.yaml` 中的 `version`，否则重新打包会覆盖历史版本。
    - 建议遵循 [SemVer 版本规范](https://semver.org/)。

4. **清理临时文件**  
   操作完成后删除临时目录：
   ```bash
   rm -rf temp-chart/
   ```

---

### **操作示例**

假设原始 Chart 为 `myapp-1.0.0.tgz`，目标修改镜像标签并调整副本数：

1. 解压并进入目录：
   ```bash
   tar -xzvf myapp-1.0.0.tgz && cd myapp
   ```
2. 修改 `values.yaml`：
   ```yaml
   replicaCount: 3
   image:
     tag: "v2.0.0"
   ```
3. 更新 `Chart.yaml`：
   ```yaml
   version: 1.0.1
   ```
4. 重新打包并安装：
   ```bash
   helm package . && helm install myapp-v2 ./myapp-1.0.1.tgz
   ```

---

### **总结**

通过解压 → 修改 → 打包 → 安装的流程，可灵活自定义已有的 Helm Chart。此方法适用于二次开发或适配不同环境的需求。

# QA

## 1. `--set replicaCount=3` 会自动修改values.yaml吗

**不会**，`helm upgrade my-release ./my-chart -n my-namespace --set replicaCount=3` **不会自动修改 `values.yaml` 文件**，它只是**临时覆盖**
Chart 的默认值，并将这些更改存储在 Kubernetes 释放（Release）历史记录中。

---

### **1. 如何验证 `--set` 覆盖的值？**

可以使用以下命令查看 Helm Release 的当前 `values`：

```bash
helm get values my-release -n my-namespace
```

如果 `replicaCount` 被修改，它会出现在输出中。

如果要查看完整的 `values.yaml`（包括默认值 + 覆盖的值）：

```bash
helm get values my-release -n my-namespace --all
```

---

### **2. `--set` 修改是否持久化？**

- **持久化到 Release 记录**：Helm 会在 Release 历史记录中保存 `--set` 覆盖的值，但不会修改 Chart 源文件（`values.yaml`）。
- **不会修改 `values.yaml` 文件**：如果你删除 Release 并重新安装，`values.yaml` 仍然保持原来的配置。

---

### **3. 如何让 `values.yaml` 也更新？**

如果你希望 `values.yaml` 也被永久修改，而不是只在 Helm 记录中生效，需要**手动编辑 `values.yaml` 文件**：

```yaml
replicaCount: 3
```

然后运行：

```bash
helm upgrade my-release ./my-chart -n my-namespace -f values.yaml
```

---

### **4. 如果既使用 `--set` 又提供 `values.yaml`？**

如果你同时提供 `--set` 和 `-f values.yaml`：

```bash
helm upgrade my-release ./my-chart -n my-namespace -f values.yaml --set replicaCount=5
```

Helm 的行为：

1. **先加载 `values.yaml`** 的内容。
2. **然后应用 `--set` 参数**，`--set` 的值**优先级更高**，会覆盖 `values.yaml` 里的相同字段。

---

### **5. 如何恢复 Helm Release 到 `values.yaml` 的默认状态？**

如果你想取消 `--set` 造成的覆盖，强制使用 `values.yaml`，可以执行：

```bash
helm upgrade my-release ./my-chart -n my-namespace -f values.yaml
```

这样 Helm 会完全使用 `values.yaml` 的内容，而忽略 `--set` 的历史覆盖。

---

### **总结**

✅ `helm upgrade --set` **不会修改 `values.yaml` 文件**，但会存储到 Helm Release 记录中。  
✅ 如果需要持久化修改，应该**手动编辑 `values.yaml` 并重新升级**。  
✅ **`--set` 的优先级高于 `values.yaml`**，如果两者冲突，`--set` 会覆盖 `values.yaml`。  
✅ **如果要恢复 `values.yaml` 的默认值，可以运行 `helm upgrade` 仅使用 `values.yaml`**。

## 2. 在回滚之前，查看之前版本的values

### **1. 查看当前 Helm Release 的 `values.yaml`**

```bash
helm get values my-release -n my-namespace
```

这将显示 **当前生效的 values 配置**。

---

### **2. 查看指定版本的 `values.yaml`**

如果你想查看**过去某个版本的 `values.yaml`**（例如 `revision 2`），可以使用：

```bash
helm get values my-release -n my-namespace --revision 2
```

这样可以在回滚之前，确认该版本的 `values.yaml` 是否符合你的需求。

---

### **3. 如何查看所有历史版本**

你可以使用以下命令列出所有历史版本：

```bash
helm history my-release -n my-namespace
```

示例输出：

```
REVISION    UPDATED                     STATUS     CHART           APP VERSION
1           2024-02-12 10:00:00         DEPLOYED   my-chart-1.0.0   1.19
2           2024-02-13 14:00:00         DEPLOYED   my-chart-1.0.1   1.21
3           2024-02-14 16:00:00         DEPLOYED   my-chart-1.0.2   1.22
```

然后你可以选择某个 `revision` 来查看它的 `values.yaml`：

```bash
helm get values my-release -n my-namespace --revision 2
```

---

### **4. 只有 `--set` 修改的值会在 `get values` 显示**

⚠️ **注意**：

- 如果 `helm upgrade` 时使用了 `-f values.yaml`，那么 `helm get values` **不会显示完整的 `values.yaml` 内容**，而只会显示**覆盖的部分**。
- 如果你想查看完整的 `values.yaml`，可以尝试：
  ```bash
  helm get manifest my-release -n my-namespace --revision 2
  ```
  这会展示该版本的完整 Kubernetes 资源清单，包括 ConfigMap、Deployment 等。

---

### **5. 如何回滚到某个版本**

如果确认 `revision 2` 适合回滚，可以执行：

```bash
helm rollback my-release 2 -n my-namespace
```

---

### **总结**

✅ `helm get values my-release -n my-namespace --revision <revision>` **可以在回滚前查看历史 `values.yaml`**。  
✅ `helm history` **列出所有版本编号**，用于选择正确的 `revision`。  
✅ `helm get manifest` **可以查看完整的 Kubernetes 配置，包括 Deployment、Service 等**。  
✅ **回滚前先检查 `values.yaml`，避免回滚到错误的配置**。
