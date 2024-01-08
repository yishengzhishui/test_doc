## 问题及解答：

### 1）docker 容器启动，外部却无法访问

```shell
docker run -itd -P  registry.cn-hangzhou.aliyuncs.com/acr-toolkit/ack-cube
##
1c893f5cba84   registry.cn-hangzhou.aliyuncs.com/acr-toolkit/ack-cube   "/docker-entrypoint.…"   10 minutes ago   Up 10 minutes   0.0.0.0:32768->80/tcp, :::32768->80/tcp   practical_hofstadter
```

原因：阿里云的安全组规则需要改动，允许外部访问这个端口

![image.png](./assets/1704701378559-image.png)
