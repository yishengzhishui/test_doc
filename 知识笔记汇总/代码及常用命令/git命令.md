## 替换代码的远程仓库
---

### 方法 1：查看当前配置的远程仓库

先查看已配置的远程仓库信息：

```bash
git remote -v
```

示例输出：

```plaintext
origin  https://example.com/old-repo.git (fetch)
origin  https://example.com/old-repo.git (push)
```

---

### 方法 2：修改现有的 `origin` 远程地址

如果您想将现有的 `origin` 修改为新的地址：

```bash
git remote set-url origin http://106.12.52.42:3000/bailan/pcb-vrs.git
```

修改完成后，可以再查看确认：

```bash
git remote -v
```

---

### 方法 3：删除并重新添加 `origin`

如果您想完全重新设置 `origin`：

```bash
git remote remove origin
```

然后再添加新的地址：

```bash
git remote add origin http://106.12.52.42:3000/bailan/pcb-vrs.git
```

---

### 方法 4：添加一个新的远程名称

如果您想保留现有的 `origin`，可以添加一个不同的名称，比如 `new-origin`：

```bash
git remote add new-origin http://106.12.52.42:3000/bailan/pcb-vrs.git
```

确认后，新的远程会以 `new-origin` 作为标识。

---

通过这些方法，您可以根据实际需要调整远程仓库的配置。