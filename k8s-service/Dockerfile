# 若运行环境是 Linux 则需把 alpine 换成 debian
# 使用最新版 alpine 作为基础镜像
FROM alpine:latest


# 将微服务的服务端运行文件拷贝到 /app 下
ADD k8s-service /k8s-service

ENTRYPOINT [ "/k8s-service" ]