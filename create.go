package zdpgo_k8s

import (
	"errors"
	"fmt"
	"path/filepath"
)

// CreateRC 执行kubectl create -f xxx.yaml命令
// @param fileName 配置文件的名称
func (k *K8S) create(fileName string) (string, error) {

	if !k.IsHealth() {
		k.log.Panic("无法成功连接到指定的服务器，请检查后重试。")
		return "", errors.New("服务器无法连接")
	}

	// 上传安装脚本
	localFilePath := fmt.Sprintf("%s/assets/yaml/%s", currentPath, fileName)
	k.log.Info("本地文件路径：", localFilePath)
	_, fileName = filepath.Split(fileName)
	remoteDirPath := fmt.Sprintf("/home/%s/%s", k.config.Username, fileName)
	k.log.Info("远程文件路径：", remoteDirPath)
	k.ssh.UploadFile(localFilePath, remoteDirPath)

	// 执行部署命令
	command := fmt.Sprintf("kubectl create -f %s", remoteDirPath)
	k.log.Info("要执行的命令是：", command)
	result, err := k.ssh.Sudo(command)
	if err != nil {
		k.log.Error("创建RC失败：", err)
	}
	return result, err
}

// CreateRC 创建RC
// @param fileName 配置文件的名称
func (k *K8S) CreateRC(fileName string) (string, error) {

	result, err := k.create(fileName)
	if err != nil {
		k.log.Error("创建RC失败：", err)
		return result, err
	}

	// 执行查询RC命令
	result, err = k.ssh.Sudo("kubectl get rc")
	return result, err
}

// CreateMysql 创建MySQL
// @param version 版本号，可选8或者5
// @param in 是否只在k8s集群内部访问
func (k *K8S) CreateMysql(version uint8, in bool) (string, error) {
	var (
		fileName string
		err      error
		result   string
	)
	// 创建pv
	fileName = fmt.Sprintf("mysql/%d/pv.yaml", version)
	result, err = k.create(fileName)
	if err != nil {
		k.log.Error("创建PV失败：", err)
		return result, err
	}

	// 创建pvc
	fileName = fmt.Sprintf("mysql/%d/pvc.yaml", version)
	result, err = k.create(fileName)
	if err != nil {
		k.log.Error("创建PVC失败：", err)
		return result, err
	}

	// 确认pv及pvc的状态
	result, err = k.ssh.Sudo("kubectl get pv")
	k.log.Info(result, err)
	result, err = k.ssh.Sudo("kubectl get pvc")
	k.log.Info(result, err)

	// 创建mysql configMap
	fileName = fmt.Sprintf("mysql/%d/config_map.yaml", version)
	result, err = k.create(fileName)
	if err != nil {
		k.log.Error("创建configMap失败：", err)
		return result, err
	}

	// 部署(Deployment)文件
	fileName = fmt.Sprintf("mysql/%d/deployment.yaml", version)
	result, err = k.create(fileName)
	if err != nil {
		k.log.Error("创建Deployment失败：", err)
		return result, err
	}

	// 创建svc（service）
	if in {
		fileName = fmt.Sprintf("mysql/%d/svc_in.yaml", version)
	} else {
		fileName = fmt.Sprintf("mysql/%d/svc_out.yaml", version)
	}
	result, err = k.create(fileName)
	if err != nil {
		k.log.Error("创建svc（service）失败：", err)
		return result, err
	}

	// 执行查询RC命令
	result, err = k.ssh.Sudo("kubectl get pods")
	return result, err
}

// CreateRC 创建Service
// @param fileName 配置文件的名称
func (k *K8S) CreateService(fileName string) (string, error) {

	result, err := k.create(fileName)
	if err != nil {
		k.log.Error("创建RC失败：", err)
		return result, err
	}

	// 执行查询Service命令
	result, err = k.ssh.Sudo("kubectl get svc")
	return result, err
}
