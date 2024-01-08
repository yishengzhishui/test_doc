#!/bin/bash

# chrono @ 2022-04

# kubeadm config images list --kubernetes-version v1.23.3
# k8s.gcr.io/kube-apiserver:v1.23.3
# k8s.gcr.io/kube-controller-manager:v1.23.3
# k8s.gcr.io/kube-scheduler:v1.23.3
# k8s.gcr.io/kube-proxy:v1.23.3
# k8s.gcr.io/pause:3.6
# k8s.gcr.io/etcd:3.5.1-0
# k8s.gcr.io/coredns/coredns:v1.8.6

# use ali registry to speed up
# 使用阿里云容器镜像服务（Aliyun Container Registry）
# 来加速 Kubernetes 集群所需的容器镜像的下载
repo=registry.aliyuncs.com/google_containers
# kubeadm config images list 获取指定 Kubernetes 版本所需的官方容器镜像列表
# 拉取并修改 Kubernetes 官方镜像标签
# 去除镜像名称中的前缀，然后使用 docker pull 从阿里云容器镜像服务拉取镜像。
for name in `kubeadm config images list --kubernetes-version v1.23.3`;
do
    # 移除镜像名称中的前缀
    src_name=${name#k8s.gcr.io/}
    src_name=${src_name#coredns/}

    # 从阿里云容器镜像服务拉取镜像
    docker pull $repo/$src_name

    # 修改标签以符合 Kubernetes 需求
    docker tag $repo/$src_name $name

    # 删除原始的阿里云容器镜像
    docker rmi $repo/$src_name
done

# 拉取 Flannel 容器镜像
for name in `grep image flannel.yml |grep -v '#image' | sed 's/image://g' -`;
do
    docker pull $name
done

# 检查拉取的容器镜像
docker images

