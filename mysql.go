package zdpgo_k8s

import (
	"io/ioutil"
	"os"
)

/*
@Time : 2022/6/26 21:42
@Author : 张大鹏
@File : mysql
@Software: Goland2021.3.1
@Description: k8s操作MySQL相关
*/

func (k *K8S) CreateMysqlRc() {
	mysqlRcYaml := `apiVersion: v1
kind: ReplicationController # 副本控制器RC
metadata:
  name: mysql # RC的名称，全局唯一
spec:
  replicas: 1 # Pod副本的期待数量
  selector:
    app: mysql # 符合目标的Pod拥有此标签
  template: # 根据此模板创建Pod的副本，也就是实例
    metadata:
      labels:
        app: mysql # Pod副本拥有的标签，对应RC的selector
    spec:
      containers: # Pod内容器的定义部分
        - name: mysql
          image: mysql
          ports:
            - containerPort: 3306 # 容器监听的端口号
          env:
            - name: MYSQL_ROOT_PASSWORD
              value: "root" # 自定义MySQL的root用户密码
`
	mysqlRcFileName := "mysql-rc.yaml"

	// 创建本地文件
	err := ioutil.WriteFile(mysqlRcFileName, []byte(mysqlRcYaml), os.ModePerm)
	if err != nil {
		k.Log.Error("创建本地文件失败", "error", err)
		return
	}
	defer os.Remove(mysqlRcFileName)

	// 上传
	k.Ssh.UploadFile(mysqlRcFileName, mysqlRcFileName)
	defer k.Ssh.Sudo("rm -rf " + mysqlRcFileName)

	// 创建
	executeResult, err := k.Ssh.Sudo("kubectl create -f " + mysqlRcFileName)
	if err != nil {
		k.Log.Error(err.Error())
		return
	}
	k.Log.Debug(executeResult)
}

func (k *K8S) CreateMysqlSvc() {
	mysqlSvcYaml := `apiVersion: v1
kind: Service # 表名是Kubernetes Service
metadata:
  name: mysql # Service的全局唯一名称
spec:
  type: NodePort # 指定可以通过外网访问
  ports:
    - port: 3306 # Service提供服务的端口号
  selector:
    app: mysql # Service对应的Pod拥有这里定义的标签
`
	mysqlSvcFileName := "mysql-svc.yaml"

	// 创建本地文件
	err := ioutil.WriteFile(mysqlSvcFileName, []byte(mysqlSvcYaml), os.ModePerm)
	if err != nil {
		k.Log.Error("创建本地文件失败", "error", err)
		return
	}
	defer os.Remove(mysqlSvcFileName)

	// 上传
	k.Ssh.UploadFile(mysqlSvcFileName, mysqlSvcFileName)
	defer k.Ssh.Sudo("rm -rf " + mysqlSvcFileName)

	// 创建
	executeResult, err := k.Ssh.Sudo("kubectl create -f " + mysqlSvcFileName)
	if err != nil {
		k.Log.Error(err.Error())
		return
	}
	k.Log.Debug(executeResult)
}

func (k *K8S) DeleteMysqlRc() {
	executeResult, err := k.Ssh.Sudo("kubectl delete rc mysql")
	if err != nil {
		k.Log.Error("删除MySQL RC失败", "error", err)
		return
	}
	k.Log.Debug("删除MySQL RC成功", "result", executeResult)
}

func (k *K8S) DeleteMysqlSvc() {
	executeResult, err := k.Ssh.Sudo("kubectl delete svc mysql")
	if err != nil {
		k.Log.Error("删除MySQL SVC失败", "error", err)
		return
	}
	k.Log.Debug("删除MySQL SVC成功", "result", executeResult)
}
