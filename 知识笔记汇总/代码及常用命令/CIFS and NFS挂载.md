## 区别

CIFS（Common Internet File System）和 NFS（Network File System）是两种常见的网络文件系统协议，主要用于在网络上共享文件。以下是它们之间的一些主要区别：

### 1. 协议类型

- **CIFS**：
    - CIFS 是一种基于 SMB（Server Message Block）协议的文件共享协议，主要用于 Windows 系统。它允许应用程序在网络上读取和写入文件，并请求其他计算机的服务。

- **NFS**：
    - NFS 是一种主要用于 Unix 和 Linux 系统的文件共享协议。它允许在网络上不同主机之间共享文件和目录，使客户端能够像访问本地文件一样访问远程文件。

### 2. 平台支持

- **CIFS**：
    - 最初设计用于 Windows 系统，但现在也有 Linux 的支持（通过 Samba）。

- **NFS**：
    - 原生支持 Unix 和 Linux 系统，广泛用于这类系统的网络文件共享。

### 3. 性能

- **CIFS**：
    - 由于 CIFS 需要更复杂的封装和更高的协议开销，它的性能可能比 NFS 差，尤其是在高延迟的网络环境下。

- **NFS**：
    - 通常提供更好的性能，尤其是在大型文件或大文件夹的共享场景中。NFS 更加高效，适合大规模数据传输。

### 4. 安全性

- **CIFS**：
    - CIFS 提供了更强的安全功能，如 Kerberos 认证和加密，但这也使得配置更复杂。

- **NFS**：
    - NFS 也可以通过 Kerberos 实现安全认证，但默认情况下，它通常依赖于主机和用户的信任关系，可能更容易受到攻击。

### 5. 文件锁定和缓存

- **CIFS**：
    - CIFS 支持文件锁定，允许多个客户端安全地访问同一文件。

- **NFS**：
    - NFS 也支持文件锁定，但其实现可能在不同版本之间有所不同。NFS 的缓存机制通常更加高效。

### 6. 配置和管理

- **CIFS**：
    - 配置 CIFS 共享通常依赖于 Samba。在 Linux 上配置时，可能需要对 Samba 进行较多配置。

- **NFS**：
    - NFS 的配置相对简单，通常只需要编辑 `/etc/exports` 文件并使用 `exportfs` 命令。

### 7. 适用场景

- **CIFS**：
    - 适用于 Windows 客户端与服务器之间的文件共享，或者在 Linux 上与 Windows 系统共享文件时。

- **NFS**：
    - 适用于 Unix/Linux 环境中的文件共享，尤其在大规模服务器集群或数据存储环境中更常用。

  要在一台 Linux 服务器（10.65.233.31）上访问另一台 Linux 服务器（10.65.233.37）上的共享目录 `/home/ray/input`，可以使用 **NFS** 或 **Samba（CIFS）**
  来设置文件共享。以下是两种方式的详细设置步骤：

---

## 挂载步骤

**10.65.233.37 是linux服务器，我希望可以在另外一个linux的服务器10.65.233.31访问到37上的文件目录**

### 方法 1: 使用 NFS 进行共享

NFS（Network File System）适合 Unix/Linux 系统间的文件共享。以下步骤展示了如何在 10.65.233.37 上设置 NFS 共享，并在 10.65.233.31 上进行挂载。

#### 在 10.65.233.37 上配置 NFS 共享

1. **安装 NFS 服务**

- **Ubuntu/Debian**：
  ```bash
  sudo apt update
  sudo apt install nfs-kernel-server
  ```

- **CentOS/RHEL**：
  ```bash
  sudo yum install nfs-utils
  ```

2. **配置共享目录**

   编辑 `/etc/exports` 文件，添加需要共享的目录配置：

   ```bash
   sudo nano /etc/exports
   ```

   添加以下行，设置 `/home/ray/input` 为共享目录，并允许 10.65.233.31 访问：

   ```bash
   /home/ray/input 10.65.233.31(rw,sync,no_subtree_check)
   ```

- **rw**：允许读写权限。
- **sync**：确保数据在请求写入时立即写入到磁盘。
- **no_subtree_check**：禁用子目录检查，提高性能。

3. **重启 NFS 服务**

   每次修改 `/etc/exports` 文件后，重启 NFS 服务：

- **Ubuntu/Debian**：
  ```bash
  sudo systemctl restart nfs-kernel-server
  ```

- **CentOS/RHEL**：
  ```bash
  sudo systemctl restart nfs
  ```

4. **验证共享目录**

   使用以下命令查看共享是否已正确配置：

   ```bash
   sudo exportfs -v
   ```

#### 在 10.65.233.31 上挂载 NFS 共享

1. **安装 NFS 客户端**

- **Ubuntu/Debian**：
  ```bash
  sudo apt install nfs-common
  ```

- **CentOS/RHEL**：
  ```bash
  sudo yum install nfs-utils
  ```

2. **创建挂载点并挂载**

   创建本地挂载点 `/mnt/input` 并挂载远程共享目录：

   ```bash
   sudo mkdir -p /mnt/input
   sudo mount -t nfs 10.65.233.37:/home/ray/input /mnt/input
   ```

3. **设置自动挂载**

   如果需要在重启后自动挂载，可以将挂载信息添加到 `/etc/fstab`：

   ```bash
   echo "10.65.233.37:/home/ray/input /mnt/input nfs defaults 0 0" | sudo tee -a /etc/fstab
   ```

---

### 方法 2: 使用 Samba（CIFS）进行共享

如果你希望使用 Samba（CIFS）共享目录，也可以按照以下步骤进行设置。

#### 在 10.65.233.37 上配置 Samba 共享

1. **安装 Samba**

- **Ubuntu/Debian**：
  ```bash
  sudo apt update
  sudo apt install samba
  ```

- **CentOS/RHEL**：
  ```bash
  sudo yum install samba
  ```

2. **配置共享目录**

在文件末尾添加以下配置，设置 `/home/ray/input` 目录为共享目录：

```ini
[input]
   path = /home/ray/input
   browsable = yes
   writable = yes
   read only = no
   guest ok = yes
```

- **[input]**：共享名称，可以改成任意名称（在客户端挂载时将使用该名称）。
- **path**：指定要共享的目录路径，即 `/home/ray/input`。
- **browsable = yes**：允许该共享在网络中被浏览。
- **writable = yes**：允许写入权限。
- **read only = no**：关闭只读模式，允许用户读写。
- **guest ok = yes**：允许匿名用户访问（如果你希望只允许授权用户访问，可以将其设置为 `no`）。


3. **创建 Samba 用户（如果不使用匿名访问）**

   如果不希望使用匿名访问，可以创建 Samba 用户并设置密码。假设要创建 `root` 用户作为 Samba 用户：

    ```bash
    sudo smbpasswd -a root
    ```

系统会提示设置密码。这个用户名和密码将在挂载时使用。

4. **设置共享目录权限**

   为 Samba 设置共享目录的权限：

   ```bash
   sudo chmod -R 0777 /home/ray/input
   ```

5. **重启 Samba 服务**

   重启 Samba 服务以应用配置更改：

- **Ubuntu/Debian**：
  ```bash
  sudo systemctl restart smbd
  ```

- **CentOS/RHEL**：
  ```bash
  sudo systemctl restart smb
  `

#### 在 10.65.233.31 上挂载 Samba 共享

1. **安装 CIFS 工具**

- **Ubuntu/Debian**：
  ```bash
  sudo apt install cifs-utils
  ```

- **CentOS/RHEL**：
  ```bash
  sudo yum install cifs-utils
  ```

2. **挂载共享目录**

   使用以下命令挂载共享：

   ```bash
   sudo mkdir -p /mnt/input
   # 匿名访问
   # sudo mount -t cifs -o username=guest,password=,vers=2.1 //10.65.233.37/input /mnt/input
   # 指定root及密码访问
   sudo mount -t cifs -o username=root,password=Baidu@123,vers=2.1 //10.65.233.37/input /home/ray/test

   ```

- **username=guest** 和 **password=** 是匿名访问方式。如果设置了用户名和密码，则在挂载命令中使用它们。

3. **设置自动挂载**

   将以下信息添加到 `/etc/fstab`，确保重启后自动挂载：

   ```bash
   echo "//10.65.233.37/input /mnt/input cifs username=guest,password=,vers=2.1 0 0" | sudo tee -a /etc/fstab
   ```

---

### 总结

- **使用 NFS**：更适合 Linux/Unix 系统之间的共享，性能高，配置简单。
- **使用 Samba（CIFS）**：更适合跨平台环境，尤其是需要 Linux 和 Windows 间共享文件时。
