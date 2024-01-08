#!/bin/bash

# chrono @ 2022-04

# https://kubernetes.io/zh/docs/setup/production-environment/tools/kubeadm/install-kubeadm/
# https://kubernetes.io/zh/docs/setup/production-environment/container-runtimes/#docker

# sudo vi /etc/hostname

# fix docker issue
# 这个命令将一个Docker配置文件 daemon.json 写入 /etc/docker/ 目录中。
# 配置内容包括使用 systemd 作为 cgroup 驱动、
# 设置日志驱动为 json-file、
# 日志最大大小为100MB、
# 存储驱动为 overlay2。
cat <<EOF | sudo tee /etc/docker/daemon.json
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2"
}
EOF

# 这部分命令启用Docker服务，重新加载systemd配置，并重新启动Docker服务，确保配置生效。
sudo systemctl enable docker
sudo systemctl daemon-reload
sudo systemctl restart docker


# iptables
# 将 br_netfilter 写入 /etc/modules-load.d/k8s.conf 文件，
# 确保该内核模块在系统启动时加载。
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF

# 将一些网络相关的内核参数写入 /etc/sysctl.d/k8s.conf 文件，
# 启用了iptables的bridge调用规则和IPv4转发。
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward=1 # better than modify /etc/sysctl.conf
EOF

# 这个命令重新加载sysctl配置，确保新的内核参数生效。
sudo sysctl --system

# Disable Swap
# 这部分命令关闭所有交换分区，
# 并通过sed命令注释掉 /etc/fstab 文件中与交换分区相关的行，
# 确保在系统重新启动时交换分区不再被激活。
sudo swapoff -a
sudo sed -ri '/\sswap\s/s/^#?/#/' /etc/fstab

# check

echo "please check these files:"
echo "/etc/docker/daemon.json"
echo "/etc/modules-load.d/k8s.conf"
echo "/etc/sysctl.d/k8s.conf"
echo "cat cat /etc/fstab"
