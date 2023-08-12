## 安装虚拟机

基本命令

```shell
sudo apt update
sudo apt install -y git vim curl jq
sudo apt install -y openssh-server
ip addr
```

### 1.terminal打不开

语言的问题，setting中全部切换为china

### 2.sudo apt update 时 is not in the sudoers file解决

```shell
su - //切换为root 密码是root
```

### 3.通过item2 连接

ssh vobxuser@192.168.56.2 changeme

### 4.change root

```
su - // 需要密码
sudo -i //(不需要密码)

passwd // change password

```

## Docker install

```
sudo apt install -y docker.io #安装Docker Engine
sudo service docker start         #启动docker服务
sudo usermod -aG docker ${USER}   #当前用户加入docker组
docker ps #它会列出当前系统里运行的容器
docker pull busybox      #拉取busybox镜像
docker images # 它会列出当前 Docker 所存储的所有镜像：
```
