package zdpgo_k8s

import (
	"errors"
	"fmt"
)

// 安装StorageClass
func (k *K8S) InstallStorageClass() (string, error) {

	if !k.IsHealth() {
		k.log.Panic("无法成功连接到指定的服务器，请检查后重试。")
		return "", errors.New("服务器无法连接")
	}

	// 上传安装脚本
	localFilePath := fmt.Sprintf("%s/assets/yaml/nfs-storage.yaml", currentPath)
	remoteDirPath := fmt.Sprintf("/home/%s", k.config.Username)
	k.ssh.UploadFile(localFilePath, remoteDirPath)

	// 执行部署命令
	result, err := k.ssh.Sudo(fmt.Sprintf("kubectl apply -f nfs-storage.yaml"))
	return result, err
}

// 安装zookeeper
func (k *K8S) InstallZookeeper() (string, error) {

	if !k.IsHealth() {
		k.log.Panic("无法成功连接到指定的服务器，请检查后重试。")
		return "", errors.New("服务器无法连接")
	}

	// 上传安装脚本
	localFilePath := fmt.Sprintf("%s/assets/yaml/zookeeper.yaml", currentPath)
	remoteDirPath := fmt.Sprintf("/home/%s", k.config.Username)
	k.ssh.UploadFile(localFilePath, remoteDirPath)

	// 执行部署命令
	result, err := k.ssh.Sudo(fmt.Sprintf("kubectl apply -f zookeeper.yaml -n kafka"))
	return result, err
}

// 安装kafka
func (k *K8S) InstallKafka() (string, error) {

	if !k.IsHealth() {
		k.log.Panic("无法成功连接到指定的服务器，请检查后重试。")
		return "", errors.New("服务器无法连接")
	}

	// 上传安装脚本
	localFilePath := fmt.Sprintf("%s/assets/yaml/kafka.yaml", currentPath)
	remoteDirPath := fmt.Sprintf("/home/%s", k.config.Username)
	k.ssh.UploadFile(localFilePath, remoteDirPath)

	// 执行部署命令
	result, err := k.ssh.Sudo(fmt.Sprintf("kubectl apply -f kafka.yaml -n kafka"))
	return result, err
}

// 安装kafka manager
func (k *K8S) InstallKafkaManager() (string, error) {

	if !k.IsHealth() {
		k.log.Panic("无法成功连接到指定的服务器，请检查后重试。")
		return "", errors.New("服务器无法连接")
	}

	// 上传安装脚本
	localFilePath := fmt.Sprintf("%s/assets/yaml/kafka-manager.yaml", currentPath)
	remoteDirPath := fmt.Sprintf("/home/%s", k.config.Username)
	k.ssh.UploadFile(localFilePath, remoteDirPath)

	// 执行部署命令
	result, err := k.ssh.Sudo(fmt.Sprintf("kubectl apply -f kafka-manager.yaml -n kafka"))
	return result, err
}
