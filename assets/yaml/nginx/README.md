```shell
#创建pod
kubectl create[apply] -f xx.yaml
#创建成功后，发现报错：因为在这个pod中创建了2个容器，但是此2个容器出现了端口冲突
#查看原因：
kubectl describe pod my-app
# 查询某个容器的日志
kubectl log my-app -c test
```