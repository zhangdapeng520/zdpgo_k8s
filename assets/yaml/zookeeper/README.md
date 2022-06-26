```shell
# 安装zookeeper
kubectl apply -f zookeeper.yml

# 查看节点状态
kubectl get all -l app=zk

# 查看Pod
kubectl get pods -w -l app=zk

# 删除所有deployment
kubectl delete deployment -all

# 删除一个
kubectl delete deployment DEPLOYMENT_NAME -n NAMESPACE_NAME

# 删除所有pod
kubectl delete pod --all

# 删除一个
kubectl delete pod zk-0 
```