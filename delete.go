package zdpgo_k8s

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

// delete 执行kubectl delete -f xxx.yaml命令
// @param fileName 配置文件的名称
func (k *K8S) delete(fileName string) (string, error) {

	if !k.IsHealth() {
		k.Log.Panic("无法成功连接到指定的服务器，请检查后重试。")
		return "", errors.New("服务器无法连接")
	}

	// 上传安装脚本
	localFilePath := fmt.Sprintf("%s/assets/yaml/%s", currentPath, fileName)
	k.Log.Info("本地文件路径：", localFilePath)
	_, fileName = filepath.Split(fileName)
	remoteDirPath := fmt.Sprintf("/home/%s/%s", k.Config.Ssh.Username, fileName)
	k.Log.Info("远程文件路径：", remoteDirPath)
	k.Ssh.UploadFile(localFilePath, remoteDirPath)

	// 执行部署命令
	command := fmt.Sprintf("kubectl delete -f %s", remoteDirPath)
	k.Log.Info("要执行的命令是：", command)
	result, err := k.Ssh.Sudo(command)
	if err != nil {
		k.Log.Error("删除资源失败：", err)
	}

	// 执行删除命令
	go func(path string) {
		if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
			command := fmt.Sprintf("rm -rf %s", remoteDirPath)
			result, err = k.Ssh.Sudo(command)
			k.Log.Info(command)
			k.Log.Info("删除远程配置文件成功：", result, err)
		}
	}(remoteDirPath)

	return result, err
}

// DeleteMysql 删除MySQL
// @param version 版本号，可选8或者5
// @param in 是否只在k8s集群内部访问
func (k *K8S) DeleteMysql(version uint8, in bool) (string, error) {
	var (
		fileName string
		err      error
		result   string
	)

	// 删除Deployment
	fileName = fmt.Sprintf("mysql/%d/deployment.yaml", version)
	result, err = k.delete(fileName)
	k.Log.Info(result, err)

	// 删除svc（service）
	if in {
		fileName = fmt.Sprintf("mysql/%d/svc_in.yaml", version)
	} else {
		fileName = fmt.Sprintf("mysql/%d/svc_out.yaml", version)
	}
	result, err = k.delete(fileName)
	k.Log.Info(result, err)

	// 删除mysql configMap
	fileName = fmt.Sprintf("mysql/%d/config_map.yaml", version)
	result, err = k.delete(fileName)
	k.Log.Info(result, err)

	// 删除pv
	fileName = fmt.Sprintf("mysql/%d/pv.yaml", version)
	result, err = k.delete(fileName)
	k.Log.Info(result, err)

	// 删除pvc
	fileName = fmt.Sprintf("mysql/%d/pvc.yaml", version)
	result, err = k.delete(fileName)
	k.Log.Info(result, err)

	// 确认pv及pvc的状态
	result, err = k.Ssh.Sudo("kubectl get pv")
	k.Log.Info(result, err)
	result, err = k.Ssh.Sudo("kubectl get pvc")
	k.Log.Info(result, err)

	// 执行查询RC命令
	result, err = k.Ssh.Sudo("kubectl get pods")
	return result, err
}
