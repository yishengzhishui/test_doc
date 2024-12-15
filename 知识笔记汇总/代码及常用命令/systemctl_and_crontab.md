# systemd

## systemd 启用流程

`systemd` 是现代 Linux 系统的初始化系统和服务管理器，用于启动和管理系统的服务。以下是 `systemd` 启用服务的完整流程，包括服务的创建、配置、启用、启动以及管理。

---

### **1. 服务配置文件的位置**

`systemd` 的服务配置文件（单元文件）通常位于以下路径之一：

- **系统服务**：
    - `/usr/lib/systemd/system/`：系统默认的服务文件位置。
    - `/etc/systemd/system/`：管理员自定义服务文件的位置，优先级高于 `/usr/lib/systemd/system/`。
- **用户服务**：
    - `~/.config/systemd/user/`：用户自定义服务文件位置（针对用户会话）。

服务文件的命名格式通常为：`<service_name>.service`。

---

### **2. 创建或检查服务文件**

#### 创建自定义服务文件

假设要创建一个服务名为 `my-custom.service` 的自定义服务，执行以下步骤：

1. 编辑服务文件：

   ```bash
   sudo vi /etc/systemd/system/my-custom.service
   ```
2. 添加服务配置内容：

   ```ini
   [Unit]
   Description=My Custom Service
   After=network.target

   [Service]
   ExecStart=/path/to/your/script.sh
   Restart=always
   User=root
   Group=root

   [Install]
   WantedBy=multi-user.target
   ```

    - **[Unit]** 部分：
        - `Description`：服务的描述。
        - `After=network.target`：指定在网络服务启动后再启动此服务。
    - **[Service]** 部分：
        - `ExecStart`：指定启动服务的命令或脚本。
        - `Restart=always`：如果服务失败，自动重启。
        - `User` 和 `Group`：指定服务运行的用户和组。
    - **[Install]** 部分：
        - `WantedBy=multi-user.target`：定义服务在哪个目标（类似于运行级别）下启用。

---

### **3. 重新加载 `systemd` 配置**

在创建或修改服务文件后，重新加载 `systemd` 配置以使其生效：

```bash
sudo systemctl daemon-reload
```

---

### **4. 启用服务**

启用服务将其设置为 **开机自动启动**：

```bash
sudo systemctl enable my-custom.service
```

- 这会创建一个符号链接，将服务关联到默认的目标（如 `multi-user.target`）。

---

### **5. 启动服务**

立即启动服务：

```bash
sudo systemctl start my-custom.service
```

---

### **6. 验证服务状态**

查看服务运行状态，检查是否启动成功：

```bash
sudo systemctl status my-custom.service
```

- 输出示例：
  ```plaintext
  ● my-custom.service - My Custom Service
     Loaded: loaded (/etc/systemd/system/my-custom.service; enabled)
     Active: active (running) since Mon 2024-11-12 12:00:00 UTC; 1min ago
   Main PID: 12345 (script.sh)
      Tasks: 1
     Memory: 1.1M
     CGroup: /system.slice/my-custom.service
             └─12345 /bin/bash /path/to/your/script.sh
  ```

---

### **7. 禁用或停止服务**

- **禁用服务（取消开机启动）**：

  ```bash
  sudo systemctl disable my-custom.service
  ```
- **停止服务**：

  ```bash
  sudo systemctl stop my-custom.service
  ```

---

### **8. 日志管理**

查看服务的日志输出：

```bash
sudo journalctl -u my-custom.service
```

- 显示最新的日志：

  ```bash
  sudo journalctl -u my-custom.service -n 50
  ```
- 实时跟踪日志：

  ```bash
  sudo journalctl -u my-custom.service -f
  ```

---

### **9. 测试服务配置文件**

如果服务未按预期运行，可以测试服务文件的语法：

```bash
sudo systemd-analyze verify /etc/systemd/system/my-custom.service
```

---

### **完整流程总结**

1. **创建服务文件**：编辑 `/etc/systemd/system/<service_name>.service`。
2. **重新加载配置**：运行 `sudo systemctl daemon-reload`。
3. **启用服务**：运行 `sudo systemctl enable <service_name>`。
4. **启动服务**：运行 `sudo systemctl start <service_name>`。
5. **检查状态和日志**：使用 `sudo systemctl status <service_name>` 和 `sudo journalctl -u <service_name>`。
6. **管理服务**：可以通过 `start`, `stop`, `enable`, `disable`, `restart` 等命令管理服务。

通过上述流程，你可以高效地管理 Linux 系统中的服务并确保其稳定运行。

## systemctl list-units

要查看当前系统中有多少个 `systemd` 服务（`systemctl` 管理的服务），你可以使用以下命令：

### 1. 列出所有 `systemctl` 服务

使用 `systemctl list-units` 命令可以列出当前所有加载的服务单元：

```bash
systemctl list-units --type=service
```

- `--type=service`：只显示服务单元（`.service` 文件）。

### 2. 统计 `systemctl` 服务数量

如果你只想知道有多少个 `systemctl` 服务单元，可以结合 `wc -l` 统计行数：

```bash
systemctl list-units --type=service --all | wc -l
```

- `--all`：显示所有服务单元，包括未加载、未激活和失败的服务。
- `wc -l`：统计行数，行数即为服务数量。

### 3. 示例输出说明

假设你运行以下命令：

```bash
systemctl list-units --type=service --all | wc -l
```

输出可能会是一个数字，比如：

```
150
```

这表示当前系统中有 150 个 `systemd` 服务单元（包括激活、未激活和失败的服务）。

### 4. 过滤特定状态的服务

如果你只想统计某种状态的服务数量，比如只统计正在运行的服务，可以这样做：

```bash
systemctl list-units --type=service --state=running | wc -l
```

- `--state=running`：仅列出正在运行的服务。

### 总结

- `systemctl list-units --type=service` 列出所有服务。
- `systemctl list-units --type=service --all | wc -l` 统计所有服务单元的数量。
- 通过结合 `wc -l`，你可以快速统计系统中有多少 `systemd` 服务。

## .service 文件的服务配置

`.service` 文件是 `systemd` 用来定义服务行为的配置文件，描述了如何启动、停止、重新加载一个服务以及服务的相关依赖关系。以下是 `.service` 文件的基本结构、常用字段解释及使用方法。

---

### **1. `.service` 文件基本结构**

`.service` 文件通常由以下部分组成：

#### **[Unit]** 部分

定义服务的元信息和依赖关系。

#### **[Service]** 部分

定义服务的行为，包括启动、停止的命令和运行环境。

#### **[Install]** 部分

定义服务与系统目标（如运行级别）的关系，决定服务是否可以启用和启动。

---

### **2. 示例 `.service` 文件**

假设我们创建一个服务来运行脚本 `/home/user/script.sh`，该服务的 `.service` 文件内容如下：

```ini
[Unit]
Description=My Custom Service
After=network.target

[Service]
Type=simple
ExecStart=/usr/bin/bash /home/user/script.sh
Restart=on-failure
User=user
Group=user

[Install]
WantedBy=multi-user.target
```

---

### **3. `.service` 文件字段详解**

#### **[Unit] 部分**

- **`Description`**：
    - 简短描述服务的功能。
    - 示例：
      ```ini
      Description=My Custom Service
      ```

- **`After`**：
    - 定义服务启动的顺序。`network.target` 表示此服务在网络启动后启动。
    - 示例：
      ```ini
      After=network.target
      ```

- **`Requires` 和 `Wants`**：
    - **`Requires`**：定义服务运行时的必需依赖项。如果依赖的服务失败，当前服务也会失败。
    - **`Wants`**：定义非强制依赖项，依赖失败不会影响当前服务。
    - 示例：
      ```ini
      Requires=network.target
      Wants=network-online.target
      ```

---

#### **[Service] 部分**

- **`Type`**：
    - 指定服务启动类型，常见选项：
        - **`simple`**：默认值，`ExecStart` 启动后，服务被视为启动成功。
        - **`forking`**：服务通过 fork 启动，父进程退出后，服务被视为启动成功。
        - **`oneshot`**：服务一次性执行完任务后退出，适合执行初始化脚本。
        - **`notify`**：服务会向 `systemd` 发送启动状态通知。
    - 示例：
      ```ini
      Type=simple
      ```

- **`ExecStart`**：
    - 指定服务启动时执行的命令。
    - 示例：
      ```ini
      ExecStart=/usr/bin/bash /home/user/script.sh
      ```

- **`ExecStop`**：
    - 指定服务停止时执行的命令（可选）。
    - 示例：
      ```ini
      ExecStop=/usr/bin/bash /home/user/stop_script.sh
      ```

- **`ExecReload`**：
    - 指定服务重新加载时执行的命令（可选）。
    - 示例：
      ```ini
      ExecReload=/usr/bin/bash /home/user/reload_script.sh
      ```

- **`Restart`**：
    - 定义服务失败时的重启策略：
        - `no`：不重启（默认）。
        - `on-success`：成功退出时重启。
        - `on-failure`：失败时重启。
        - `always`：无论任何退出状态都重启。
    - 示例：
      ```ini
      Restart=on-failure
      ```

- **`User` 和 `Group`**：
    - 指定运行服务的用户和用户组。
    - 示例：
      ```ini
      User=user
      Group=user
      ```

- **`Environment`**：
    - 定义环境变量。
    - 示例：
      ```ini
      Environment="ENV_VAR=value"
      ```

---

#### **[Install] 部分**

- **`WantedBy`**：
    - 指定服务关联的目标（类似运行级别），常见值：
        - **`multi-user.target`**：类似于传统的运行级别 3，适用于非图形环境。
        - **`graphical.target`**：图形化环境。
    - 示例：
      ```ini
      WantedBy=multi-user.target
      ```

---

## servcie 自动重启

### **通过 `systemd` 自动重启服务**

在服务的 `.service` 文件中，可以配置 `Restart` 和相关参数来实现服务失败后的自动重启。以下是具体步骤：

---

#### **1. 修改 `.service` 文件**

编辑服务的 `.service` 文件，添加或修改以下内容：

```ini
[Service]
# 服务启动命令
ExecStart=/path/to/your/program

# 自动重启配置
Restart=on-failure
RestartSec=5
StartLimitInterval=60
StartLimitBurst=3
```

##### **字段解释：**

- **`Restart=on-failure`**：
    - 指定何时自动重启服务：
        - `no`：默认值，不自动重启。
        - `on-failure`：仅在非 0 退出状态或超时导致服务失败时重启。
        - `always`：无论任何退出状态都重启。
        - `on-abnormal`：仅在异常（如信号中止、核心转储）导致服务失败时重启。

- **`RestartSec=5`**：
    - 服务停止后等待 5 秒再尝试重启。

- **`StartLimitInterval=60`** 和 **`StartLimitBurst=3`**：
    - 限制在 **60 秒**内允许最多 **3 次重启**。
    - 如果超出限制，服务将被标记为失败并不会再重启。

---

#### **2. 重新加载 `systemd` 配置**

每次修改 `.service` 文件后，需要重新加载 `systemd` 配置：

```bash
sudo systemctl daemon-reload
```

### **验证服务自动重启**

1. 手动停止服务并观察重启行为：
   ```bash
   sudo systemctl stop <service_name>
   ```

2. 或者模拟服务失败（通过修改程序或发送 `kill` 信号）：
   ```bash
   sudo kill -SIGTERM <service_pid>
   ```

3. 检查服务是否重启：
   ```bash
   sudo systemctl status <service_name>
   ```

4. 查看服务的日志：
   ```bash
   sudo journalctl -u <service_name>
   ```

---

# chomd

`chmod 644` 是一个用于更改文件权限的命令，具体作用是将文件的权限设置为 `644`。以下是详细的解释：

---

### **文件权限基础**

在 Linux 中，文件权限分为三类：

1. **所有者（Owner）**：文件的拥有者。
2. **所在组（Group）**：与文件所有者在同一用户组的用户。
3. **其他人（Others）**：不属于上述两类的其他用户。

权限包括：

- **读（r）**：值为 `4`，允许读取文件内容。
- **写（w）**：值为 `2`，允许修改文件内容。
- **执行（x）**：值为 `1`，允许运行该文件（如果是脚本或可执行文件）。

权限表示法分为两种：

- **符号表示法**：如 `rw-r--r--`。
- **八进制表示法**：如 `644`。

---

### **八进制权限解释**

`chmod 644` 中的 `644` 是一个八进制数，每一位表示一类用户的权限。

| 数字 | 二进制 | 权限表示 | 含义                  |
|------|--------|----------|-----------------------|
| `6`  | `110`  | `rw-`    | 读和写权限（没有执行） |
| `4`  | `100`  | `r--`    | 只读权限              |
| `4`  | `100`  | `r--`    | 只读权限              |

#### **权限分配**

- **第一位**：`6` 表示文件的所有者权限（读和写）。
- **第二位**：`4` 表示文件所在组的权限（只读）。
- **第三位**：`4` 表示其他用户的权限（只读）。

---

### **`chmod 644` 的作用**

运行 `chmod 644 filename` 后，文件的权限变为：

```plaintext
rw-r--r--
```

- **所有者**：拥有读 (`r`) 和写 (`w`) 权限。
- **所在组**：只有读 (`r`) 权限。
- **其他人**：只有读 (`r`) 权限。

这通常用于配置普通文件的权限，确保只有文件的所有者可以修改，而其他人只能读取。

---

### **适用场景**

- 配置文件（如 `/etc/samba/smb.conf`）通常使用 `644`，这样其他用户可以读取配置，但不能修改。
- HTML 文件、日志文件等需要被其他用户读取但不允许修改时，也常用 `644`。

---

### **命令示例**

1. **设置权限**：
   ```bash
   chmod 644 example.txt
   ```

   权限变为：
   ```plaintext
   rw-r--r--
   ```

2. **查看权限**： 使用 `ls -l` 命令查看文件权限：
   ```bash
   ls -l example.txt
   ```

   输出示例：
   ```plaintext
   -rw-r--r--  1 user group 1024 Nov 12 14:00 example.txt
   ```

---

### **对比其他常见权限**

| 权限    | 八进制 | 符号表示  | 含义                                           |
|---------|--------|-----------|------------------------------------------------|
| `777`   | `rwxrwxrwx` | 所有用户都有读、写、执行权限                     |
| `755`   | `rwxr-xr-x` | 所有者有读、写、执行权限，其他用户有读和执行权限  |
| `644`   | `rw-r--r--` | 所有者有读、写权限，其他用户只有读权限           |
| `600`   | `rw-------` | 只有所有者有读、写权限，其他用户没有任何权限      |

---

### **总结**

- `chmod 644` 设置文件权限为：**所有者可读写，所在组和其他用户只读**。
- 这是非常常见的权限设置，适合普通文件（如配置文件），确保安全性和可读性。

# crontab

`crontab -e` 是 Linux 系统中用来编辑当前用户的 **cron 表（crontab）** 的命令。  
**Cron 表** 是一个配置文件，用于定义定时任务，cron 会根据配置文件的内容在指定的时间自动执行相应的任务。

---

### **Cron 表的用途**

- 定期运行脚本或命令，例如备份文件、清理日志、发送邮件等。
- 可以精确到分钟、小时、天、月、星期来调度任务。
- 每个用户（包括系统级别的任务）都可以拥有自己的 `crontab`。

---

### **如何使用 `crontab -e`**

#### **步骤 1: 打开编辑器**

执行以下命令：

```bash
crontab -e
```

这会打开当前用户的 cron 表。如果是第一次使用，会提示选择默认的文本编辑器（如 `nano`, `vim`）。

---

#### **步骤 2: 编辑 cron 表**

在编辑器中添加任务，每一行代表一个定时任务。  
Cron 表的语法如下：

```
* * * * * command_to_execute
```

每个字段的含义如下：

| 字段         | 含义              | 值的范围           |
|--------------|-------------------|--------------------|
| `minute`     | 分钟              | 0–59              |
| `hour`       | 小时              | 0–23              |
| `day`        | 每月的某一天      | 1–31              |
| `month`      | 月份              | 1–12 或者 `jan-dec` |
| `weekday`    | 星期几            | 0–7 或者 `sun-sat` |
| `command`    | 要执行的命令或脚本 | 完整路径的命令     |

---

#### **步骤 3: 保存并退出**

编辑完成后，保存文件并退出编辑器（如在 `vim` 中按 `ESC`，然后输入 `:wq`）。

---

### **示例任务**

#### 1. 每天凌晨 2 点运行脚本

运行位于 `/home/user/backup.sh` 的备份脚本：

```bash
0 2 * * * /home/user/backup.sh
```

#### 2. 每小时清理一次临时目录

清理 `/tmp` 文件夹中的过期文件：

```bash
0 * * * * /usr/bin/find /tmp -type f -mtime +7 -delete
```

#### 3. 每 5 分钟运行一个任务

运行 `/usr/local/bin/check_status.sh` 脚本：

```bash
*/5 * * * * /usr/local/bin/check_status.sh
```

#### 4. 每周日凌晨 1 点重启服务

重启 Apache 服务：

```bash
0 1 * * 0 /usr/sbin/service apache2 restart
```

---

### **管理 `crontab`**

#### 1. 查看当前用户的定时任务

使用以下命令查看当前用户的 `crontab` 内容：

```bash
crontab -l
```

#### 2. 删除当前用户的定时任务

使用以下命令清空当前用户的 `crontab`：

```bash
crontab -r
```

#### 3. 修改其他用户的 `crontab`（需要 root 权限）

管理员可以为其他用户编辑 `crontab`：

```bash
sudo crontab -u username -e
```

---

### **注意事项**

1. **路径问题**：
    - `cron` 运行的环境和用户的登录环境不同，命令应使用完整路径。例如，`/usr/bin/python3` 而不是 `python3`。
    - 可以在 `crontab` 中的顶部设置环境变量，例如：
      ```bash
      PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
      ```

2. **日志记录**：
    - Cron 的默认日志在 `/var/log/cron` 或 `/var/log/syslog` 中。
    - 如果需要输出任务日志，可以重定向输出，例如：
      ```bash
      0 2 * * * /home/user/backup.sh >> /home/user/backup.log 2>&1
      ```

3. **权限问题**：
    - 确保运行的命令或脚本有执行权限（`chmod +x script.sh`）。
    - 一些系统任务需要以 `root` 用户身份运行。

---

### **总结**

- `crontab -e` 用于编辑当前用户的定时任务表。
- 按照 cron 表语法定义任务运行的时间和命令。
- 定时任务是非常灵活的工具，适用于自动化日常任务的管理。
- 在编辑和调试定时任务时，注意环境变量、路径和权限问题。

## service 文件可以放置在自己的文件假中，软链到/etc/systemd/system

---

### **操作步骤**

#### 1. **创建软链接**

假设服务文件位于 `/home/test/my-custom.service`，需要将其链接到 `/etc/systemd/system/`。

运行以下命令：

```bash
sudo ln -s /home/test/my-custom.service /etc/systemd/system/my-custom.service
```

- `ln -s`：创建软链接。
- `/home/test/my-custom.service`：原文件路径。
- `/etc/systemd/system/my-custom.service`：目标链接路径。

---

#### 2. **验证软链接是否成功**

运行以下命令，检查软链接是否正确：

```bash
ls -l /etc/systemd/system/my-custom.service
```

#### 3. **运行**

```bash
# 重新加载
sudo systemctl daemon-reload
```

```bash
#启动并启用服务
sudo systemctl enable my-custom.service
sudo systemctl start my-custom.service
```

### **注意事项**

1. **推荐使用 `/etc/systemd/system/`**：
    - `/etc/systemd/system/` 是用于管理员定义的服务，优先级高于 `/usr/lib/systemd/system/`。
    - 自定义服务一般应放置在 `/etc/systemd/system/` 中。

2. **权限问题**：
    - 确保源文件 `/home/test/my-custom.service` 的权限允许 `systemd` 读取（通常需要 `root` 权限）。
    - 检查权限：
      ```bash
      ls -l /home/test/my-custom.service
      ```
    - 如果需要，调整权限：
      ```bash
      sudo chmod 644 /home/test/my-custom.service
      ```

3. **路径问题**：
    - 如果 `/home/test` 是网络文件系统（如 NFS），可能会导致启动问题。在这种情况下，建议复制文件而非创建软链接。

## 查看服务器的systemctl的服务：

### 1. 列出所有 `systemctl` 服务

使用 `systemctl list-units` 命令可以列出当前所有加载的服务单元：

```bash
systemctl list-units --type=service
```

- `--type=service`：只显示服务单元（`.service` 文件）。

### 2. 统计 `systemctl` 服务数量

如果你只想知道有多少个 `systemctl` 服务单元，可以结合 `wc -l` 统计行数：

```bash
systemctl list-units --type=service --all | wc -l
```

- `--all`：显示所有服务单元，包括未加载、未激活和失败的服务。
- `wc -l`：统计行数，行数即为服务数量。

### 3. 示例输出说明

假设你运行以下命令：

```bash
systemctl list-units --type=service --all | wc -l
```

输出可能会是一个数字，比如：

```
150
```

这表示当前系统中有 150 个 `systemd` 服务单元（包括激活、未激活和失败的服务）。

### 4. 过滤特定状态的服务

如果你只想统计某种状态的服务数量，比如只统计正在运行的服务，可以这样做：

```bash
systemctl list-units --type=service --state=running | wc -l
```

- `--state=running`：仅列出正在运行的服务。

### 总结

- `systemctl list-units --type=service` 列出所有服务。
- `systemctl list-units --type=service --all | wc -l` 统计所有服务单元的数量。
- 通过结合 `wc -l`，你可以快速统计系统中有多少 `systemd` 服务。