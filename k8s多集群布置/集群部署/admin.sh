#!/bin/bash

# chrono @ 2022-04

# https://kubernetes.io/zh/docs/setup/production-environment/tools/kubeadm/install-kubeadm/
# https://kubernetes.io/zh/docs/tasks/administer-cluster/kubeadm/kubeadm-upgrade/

#sudo apt update
# 这两个命令用于更新系统的包列表和安装一些基本的软件包，包括支持 HTTPS 传输的软件包、证书、curl 工具以及 NFS 共享的支持。
sudo apt install -y apt-transport-https ca-certificates curl nfs-common

#sudo curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg
# 添加Kubernetes的APT密钥和源：
# 从阿里云的镜像站下载 Kubernetes 的 APT 密钥，并添加到系统的密钥环中。
curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | sudo apt-key add -

#echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
# 将 Kubernetes 的 APT 源添加到系统的软件源列表中，以便后续通过这个源安装 Kubernetes 软件包。
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main
EOF

# 更新系统的软件源
sudo apt update
# 安装 Kubernetes 工具，包括 kubeadm、kubelet 和 kubectl。版本号是指定安装的具体版本
sudo apt install -y kubeadm=1.23.3-00 kubelet=1.23.3-00 kubectl=1.23.3-00
# 已安装的 kubeadm、kubelet 和 kubectl 标记为"hold"状态，
# 防止它们在系统更新时被升级到新的版本。
sudo apt-mark hold kubeadm kubelet kubectl

# check
kubeadm version
kubectl version --client

# add this line to '.bashrc'
#source <(kubectl completion bash)

