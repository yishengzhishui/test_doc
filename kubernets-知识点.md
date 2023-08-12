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


### 1.被隔离的进程

```shell
docker pull alpine

docker run -it alpine sh

```

`-it`：这是两个选项的组合，用于在容器内启动一个交互式的终端。`-i`表示标准输入保持打开，`-t`表示分配一个伪终端（TTY）。`sh`：这是在Alpine容器内要运行的命令，它启动了一个Shell终端。

这个命令的含义是：在一个新的Docker容器中，使用Alpine Linux镜像启动一个交互式的Shell会话，您可以在这个Shell终端中执行命令和操作容器的文件系统。

**容器，就是一个特殊的隔离环境，它能够让进程只看到这个环境里的有限信息，不能对外界环境施加影响**

1. 容器就是操作系统里一个特殊的“沙盒”环境，里面运行的进程只能看到受限的信息，与外部系统实现了隔离。
2. 容器隔离的目的是为了系统安全，限制了进程能够访问的各种资源。
3. 相比虚拟机技术，容器更加轻巧、更加高效，消耗的系统资源非常少，在云计算时代极具优势。
4. 容器的基本实现技术是 Linux 系统里的 namespace、cgroup、chroot。

![image.png](./assets/image.png)

### 镜像操作

镜像的完整名字由两个部分组成，名字和标签，中间用 : 连接起来。

```shell
docker rmi redis  
docker rmi d4c //移除镜像
```

### 容器操作

基本的格式是“docker run 设置参数”，再跟上“镜像名或 ID”，后面可能还会有附加的“运行命令”。

```shell
docker run -h srv alpine hostname
```

这里的 -h srv 就是容器的运行参数，alpine 是镜像名，它后面的 hostname 表示要在容器里运行的“hostname”这个程序，输出主机名。

docker run 常用参数

> -it 表示开启一个交互式操作的 Shell，这样可以直接进入容器内部，就好像是登录虚拟机一样。（它实际上是“-i”和“-t”两个参数的组合形式）
>
> -d 表示让容器在后台运行，这在我们启动 Nginx、Redis 等服务器程序的时候非常有用。
>
> --name 可以为容器起一个名字，方便我们查看，不过它不是必须的，如果不用这个参数，Docker 会分配一个随机的名字。



对于正在运行中的容器，我们可以使用 docker exec 命令在里面执行另一个程序

容器被停止后使用 docker ps 命令就看不到了，不过容器并没有被彻底销毁，我们可以使用 docker ps -a 命令查看系统里所有的容器，当然也包括已经停止运行的容器：

```shell
docker ps
docker ps -a //查看系统里所有的容器 
docker rm 容器的id //只是删除容器，不是删除镜像
docker run -d --rm nginx:alpine // --rm参数 是不保存容器，用完了就删除
```

![image.png](./assets/1691830013727-image.png)


### Dockerfile


```shell

// 第一条指令必须是 FROM
FROM alpine:3.15                # 选择Alpine镜像 镜像的安全和大小
FROM ubuntu:bionic              # 选择Ubuntu镜像 应用的运行稳定性

// copy 需要专门的文件路径，不可以随意指定
COPY ./a.txt  /tmp/a.txt    # 把构建上下文里的a.txt拷贝到镜像的/tmp目录
COPY /etc/hosts  /tmp       # 错误！不能使用构建上下文之外的文件 

//Dockerfile 里一条指令只能是一行， RUN 指令会在每行的末尾使用续行符 \，命令之间也会用 && 来连接
RUN apt-get update \
    && apt-get install -y \
        build-essential \
        curl \
        make \
        unzip \
    && cd /tmp \
    && curl -fSL xxx.tar.gz -o xxx.tar.gz\
    && tar xzf xxx.tar.gz \
    && cd xxx \
    && ./config \
    && make \
    && make clean




```

run 太长容易写错

一种变通的技巧：把这些 Shell 命令集中到一个脚本文件里，用 COPY 命令拷贝进去再用 RUN 来执行：

```shell
COPY setup.sh  /tmp/                # 拷贝脚本到/tmp目录

RUN cd /tmp && chmod +x setup.sh \  # 添加执行权限
    && ./setup.sh && rm setup.sh    # 运行脚本然后再删除
```

RUN 指令实际上就是 Shell 编程，如果你对它有所了解，就应该知道它有变量的概念，可以实现参数化运行，这在 Dockerfile 里也可以做到，需要使用两个指令 ARG 和 ENV。

**ARG 创建的变量只在镜像构建过程中可见，容器运行时不可见，**

**ENV 创建的变量不仅能够在构建镜像的过程中使用，在容器运行时也能够以环境变量的形式被应用程序使用**

```shell
ARG IMAGE_BASE="node"
ARG IMAGE_TAG="alpine"

ENV PATH=$PATH:/tmp
ENV DEBUG=OFF

//EXPOSE，它用来声明容器对外服务的端口号
EXPOSE 443           # 默认是tcp协议
EXPOSE 53/udp        # 可以指定udp协议
```

如果目录里有的文件（例如 readme/.git/.svn 等）不需要拷贝进镜像，docker 也会一股脑地打包上传，效率很低。

为了避免这种问题，你可以在“构建上下文”目录里再建立一个 .dockerignore 文件，语法与 .gitignore 类似，排除那些不需要的文件。

```shell
# docker ignore
*.swp
*.sh
```

#### 小结

```shell
docker build -f Dockerfile文件名 .  // .是当前路径 代表构建上下文 就是 context
```

好了，今天我们一起学习了容器镜像的内部结构，重点理解容器镜像是由多个只读的 Layer 构成的，同一个 Layer 可以被不同的镜像共享，减少了存储和传输的成本。

**只有 RUN, COPY, ADD 会生成新的镜像层，其它指令只会产生临时层**

1. 创建镜像需要编写 Dockerfile，写清楚创建镜像的步骤，每个指令都会生成一个 Layer。

2. Dockerfile 里，第一个指令必须是 FROM，用来选择基础镜像，常用的有 Alpine、Ubuntu 等。

3. 其他常用的指令有：COPY、RUN、EXPOSE，分别是拷贝文件，运行 Shell 命令，声明服务端口号。docker build 需要用 -f 来指定 Dockerfile，如果不指定就使用当前目录下名字是“Dockerfile”的文件。

4. docker build 需要指定“构建上下文”，其中的文件会打包上传到 Docker daemon，所以尽量不要在“构建上下文”中存放多余的文件。

5. 创建镜像的时候应当尽量使用 -t 参数，为镜像起一个有意义的名字，方便管理。


### 数据交换

![image.png](./assets/1691908922910-image.png)

```shell
docker run -d --rm redis
docker ps // dedao container id
docker cp a.txt {container_id}:/tmp //数据拷贝 到容器中
docker cp {container_id}:/tmp/a.txt ./b.txt //从容器中拷贝数据
```

#### 共享主机文件

docker run 命令启动容器的时候使用 -v 参数就行，具体的格式是“宿主机路径: 容器内路径”。

我还是以 Redis 为例，启动容器，使用 -v 参数把本机的“/tmp”目录挂载到容器里的“/tmp”目录，也就是说让容器共享宿主机的“/tmp”目录：

`docker run -d --rm -v /tmp:/tmp redis`

然后我们再用 docker exec 进入容器，查看一下容器内的“/tmp”目录，应该就可以看到文件与宿主机是完全一致的。

```shell
docker ps //得到容器ID
docker exec -it {容器ID} sh 
```

我们可以在不变动本机环境的前提下，使用镜像安装任意的应用，然后直接以容器来运行我们本地的源码

比如我本机上只有 Python 2.7，但我想用 Python 3 开发，如果同时安装 Python 2 和 Python 3 很容易就会把系统搞乱，所以我就可以这么做：

* 先使用 docker pull 拉取一个 Python 3 的镜像，因为它打包了完整的运行环境，运行时有隔离，所以不会对现有系统的 Python 2.7 产生任何影响。
* 在本地的某个目录编写 Python 代码，然后用 -v 参数让容器共享这个目录。
* 现在就可以在容器里以 Python 3 来安装各种包，再运行脚本做开发了。

```shell
docker pull python:alpine
docker run -it --rm -v `pwd`:/tmp python:alpine sh
```

#### 网络互通

Docker 提供了三种网络模式，分别是 null、host 和 bridge。

1. null 是最简单的模式，也就是没有网络，但允许其他的网络插件来自定义网络连接，这里就不多做介绍了。
2. host 的意思是直接使用宿主机网络，相当于去掉了容器的网络隔离（其他隔离依然保留），所有的容器会共享宿主机的 IP 地址和网卡。这种模式没有中间层，自然通信效率高，但缺少了隔离，运行太多的容器也容易导致端口冲突。

host 模式需要在 docker run 时使用 --net=host 参数，下面我就用这个参数启动 Nginx：

```shell
docker run -d --rm --net=host nginx:alpine
```
