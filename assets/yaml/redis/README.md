```shell
# 创建config
kubectl apply -f redis-config.yaml

# 创建pod
kubectl apply -f redis-pod.yaml

# 检查创建的对象
kubectl get pod/redis configmap/redis-config

# 进入Redis容器
kubectl exec -it redis -- redis-cli

# 删除创建的资源
kubectl delete pod/redis configmap/redis-config 
```