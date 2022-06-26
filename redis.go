package zdpgo_k8s

import (
	"io/ioutil"
	"os"
)

/*
@Time : 2022/6/26 14:30
@Author : 张大鹏
@File : redis
@Software: Goland2021.3.1
@Description: k8s操作Redis相关
*/

func (k *K8S) DeployRedis() error {
	redisConfigYaml := `apiVersion: v1
kind: ConfigMap
metadata:
  name: redis-config
data:
  redis-config: |
    maxmemory 2mb
    maxmemory-policy allkeys-lru    `

	redisPodYaml := `apiVersion: v1
kind: Pod
metadata:
  name: redis
spec:
  containers:
    - name: redis
      image: redis:5.0.4
      command:
        - redis-server
        - "/redis-master/redis.conf"
      env:
        - name: MASTER
          value: "true"
      ports:
        - containerPort: 6379
      resources:
        limits:
          cpu: "0.1"
      volumeMounts:
        - mountPath: /redis-master-data
          name: data
        - mountPath: /redis-master
          name: config
  volumes:
    - name: data
      emptyDir: {}
    - name: config
      configMap:
        name: redis-config
        items:
          - key: redis-config
            path: redis.conf
`

	// 创建config配置文件
	configFile := "redis-config.yaml"
	err := ioutil.WriteFile(configFile, []byte(redisConfigYaml), os.ModePerm)
	if err != nil {
		k.Log.Error("创建config配置文件失败", "error", err)
		return err
	}
	defer os.Remove(configFile)

	// 创建pod配置文件
	podFile := "redis-pod.yaml"
	err = ioutil.WriteFile(podFile, []byte(redisPodYaml), os.ModePerm)
	if err != nil {
		k.Log.Error("创建pod配置文件失败", "error", err)
		return err
	}
	defer os.Remove(podFile)

	// 上传配置文件
	k.Ssh.UploadFile(configFile, configFile)
	defer k.Ssh.Run("rm -rf " + configFile)
	k.Ssh.UploadFile(podFile, podFile)
	defer k.Ssh.Run("rm -rf " + podFile)

	// 执行命令
	_, err = k.Ssh.Sudo("kubectl apply -f " + configFile)
	if err != nil {
		k.Log.Error("创建ConfigMap失败", "error", err)
		return err
	}
	_, err = k.Ssh.Sudo("kubectl apply -f " + podFile)
	if err != nil {
		k.Log.Error("创建Pod失败", "error", err)
		return err
	}

	// 查看
	commandResult, _ := k.Ssh.Sudo("kubectl get pod/redis configmap/redis-config")
	k.Log.Debug("部署Redis结束", "result", commandResult)

	return nil
}
