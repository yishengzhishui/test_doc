#!/bin/bash

# chrono @ 2022-04

# https://kubernetes.io/zh/docs/reference/setup-tools/kubeadm/kubeadm-init/

# init k8s
# --apiserver-advertise-address=192.168.10.210
# --pod-network-cidr=10.10.0.0/16 指定了 Pod 网络的 CIDR 地址，这是用于容器之间通信的 IP 地址范围。
# --kubernetes-version=v1.23.3 指定了 Kubernetes 版本。
# --v=5 设置详细的日志级别，有助于调试。
sudo kubeadm init \
    --pod-network-cidr=10.10.0.0/16 \
    --kubernetes-version=v1.23.3 \
    --v=5

# enable kubectl
# 创建存放 kubectl 配置的目录
mkdir -p $HOME/.kube
# 将集群的管理员配置复制到用户的 ~/.kube/config 文件中。
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
# 设置 ~/.kube/config 文件的所有者，确保用户有权访问该文件。
sudo chown $(id -u):$(id -g) $HOME/.kube/config

# check
kubectl version
kubectl get node

# kubeadm join 192.168.10.210:6443 --token tlx8h6.nqq9ae0x6n311ur2 \
#   --discovery-token-ca-cert-hash sha256:3ad1e8a51484ec125e2394f03eb3c0429f467a88b432a9408faef6d00f197e87

# get join token
# kubeadm token list
# kubeadm token create --print-join-command
# openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'


