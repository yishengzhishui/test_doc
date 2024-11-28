## 云服务器通过ssh登录不了（安全组设置没有问题）

问题描述：云服务在重启之后（未完成续费，隔了一天续费的），登录异常，发现是eth0 是关闭状态，服务器了前两天有关机 开机后NetworkManager没启动 ； 这句话的意思是：

1. **eth0 是关闭状态**：`eth0` 是一个网络接口的名称，通常代表物理网络卡或者虚拟网络接口。如果说 `eth0` 是关闭状态，意味着该网络接口没有启用或没有处于活动状态，因此网络通信无法通过该接口进行。

2. **服务器了前两天有关机 开机后 NetworkManager 没启动**：这部分表明，服务器在前两天关机了，之后重新开机时，服务器的网络管理工具 `NetworkManager` 没有自动启动。`NetworkManager` 是
   Linux 系统中用来管理网络连接的服务，它可以管理有线、无线、VPN 等各种类型的网络接口。

    - **NetworkManager 没启动**：由于 `NetworkManager` 没有启动，导致系统无法自动配置和管理网络接口，可能就无法正常连接网络。

    - **eth0 无法启动**：如果 `NetworkManager` 没有启动，可能就没有有效的网络配置来启用 `eth0` 接口，导致该网络接口无法启用。

### 可能的原因和解决方案：

1. **NetworkManager 没有启动**： 由于关机后 `NetworkManager` 没有自动启动，可以手动启动该服务：

   ```bash
   sudo systemctl start NetworkManager
   ```

   如果您希望在每次开机时都自动启动 `NetworkManager`，可以启用它：

   ```bash
   sudo systemctl enable NetworkManager
   ```

2. **eth0 接口没有启用**： 如果 `NetworkManager` 启动后仍然无法启用 `eth0`，可以尝试手动启动 `eth0` 接口：

   ```bash
   sudo ifup eth0
   ```

   或者使用 `ip` 命令来启用网络接口：

   ```bash
   sudo ip link set eth0 up
   ```

3. **检查接口配置**： 如果 `eth0` 仍然无法启动，检查 `/etc/network/interfaces` 或者 `/etc/sysconfig/network-scripts/ifcfg-eth0` 文件（取决于您使用的
   Linux 发行版）是否配置正确。确保 `eth0` 接口的配置文件没有被破坏或清空。

4. **检查 NetworkManager 配置**： 如果 `NetworkManager` 还是没有正常工作，可能需要检查其配置文件是否正确，或者查看日志以找出启动失败的原因：

   ```bash
   journalctl -u NetworkManager
   ```

## 云服务器重启后 需要检查docker是否正常

```shell
docker --version
systemctl status docker

```

## mysql root忘记密码

1. 在/etc/my.cnf 加上

```shell
skip-grant-tables
```

2. 重启服务 `systemctl restart mysqld.service`

在 MySQL 中，`user` 表存储了用户的账户信息。如果您需要修改用户的密码，可以直接更新 `mysql.user` 表中的密码字段。具体方法取决于您使用的 MySQL 版本。

### 对于 MySQL 5.7 及之前版本：

在 MySQL 5.7 及之前的版本中，密码存储在 `password` 字段中。要修改 `root` 用户的密码，您可以按照以下步骤操作：

1. **登录 MySQL**（确保您已使用 `--skip-grant-tables` 启动了 MySQL 服务，或已经可以登录）

```bash
mysql -u root
```

2. **选择 `mysql` 数据库**

```sql
USE
mysql;
```

3. **更新密码**

```sql
UPDATE user
SET password=PASSWORD('newpassword')
WHERE user = 'root';
```

4. **刷新权限**

```sql
FLUSH
PRIVILEGES;
```

5. **退出 MySQL**

```sql
exit;
```

6. **重新启动 MySQL**

```bash
sudo systemctl restart mysqld
```

---

### 对于 MySQL 8.0 及之后版本：

在 MySQL 8.0 中，密码存储在 `authentication_string` 字段中，并且使用 `ALTER USER` 语句来更新密码，而不是直接修改 `user` 表。

1. **登录 MySQL**

```bash
mysql -u root
```

2. **更新密码**

```sql
ALTER
USER 'root'@'localhost' IDENTIFIED BY 'newpassword';
```

3. **刷新权限**

```sql
FLUSH
PRIVILEGES;
```

4. **退出 MySQL**

```sql
exit;
```

5. **重新启动 MySQL**

```bash
sudo systemctl restart mysqld
```

### 重要提示：

- 请确保您将 `'newpassword'` 替换为您实际想要设置的新密码。
- 在执行 `UPDATE` 或 `ALTER USER` 命令前后，记得执行 `FLUSH PRIVILEGES;` 以使更改生效。
- 确保您已经正确停止和启动 MySQL 服务，并且重新登录时使用了新的密码。

## ssh 访问外部服务器的配置

```ssh
# 针对主机 106.12.52.42 的 SSH 配置
Host 106.12.52.42
  HostName 106.12.52.42        # 远程主机的实际地址（IP 或主机名）
  User root                    # 用于连接的用户是 root
  Port 22                      # SSH 连接的端口是 22，通常是默认端口
  PreferredAuthentications publickey,password  # 优先使用公钥认证，其次使用密码认证

# 针对主机 106.12.52.42 的另一个 SSH 配置，指定端口为 3000
Host 106.12.52.42
  HostName 106.12.52.42        # 远程主机的实际地址
  Port 3000                    # 使用 3000 端口进行连接（通常不是默认的 SSH 端口）
  User git                     # 使用 git 用户进行连接（通常用于 Git 服务）
  PreferredAuthentications publickey  # 仅使用公钥认证（忽略密码认证）
  IdentitiesOnly yes           # 只使用配置文件中指定的密钥（忽略其他密钥代理）
  AddKeysToAgent yes           # 将密钥添加到 SSH agent 中，以便后续使用
  UseKeychain yes              # 在 macOS 上使用 SSH 密钥链（Keychain）管理密钥
  IdentityFile ~/.ssh/id_ed25519  # 指定使用的私钥文件（id_ed25519 是常见的 SSH 密钥类型）

# 针对 GitHub 的 SSH 配置
Host github.com
  HostName github.com          # 远程主机是 github.com
  User git                     # 使用 git 用户进行连接（用于 GitHub）
  PreferredAuthentications publickey  # 优先使用公钥认证
  IdentitiesOnly yes           # 只使用配置文件中指定的密钥
  AddKeysToAgent yes           # 将密钥添加到 SSH agent 中
  UseKeychain yes              # 在 macOS 上使用密钥链管理密钥
  IdentityFile ~/.ssh/id_ed25519  # 使用指定的私钥文件进行身份验证

# 针对所有主机的通用 SSH 配置
Host *
  ServerAliveInterval 60       # 每 60 秒发送一个 "keep-alive" 数据包，确保连接持续活跃
  ServerAliveCountMax 120      # 如果连续 120 次 "keep-alive" 信号未得到响应，则断开连接

```