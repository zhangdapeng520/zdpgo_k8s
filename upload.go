package zdpgo_k8s

import (
	"fmt"
	"path"
	"runtime"
)

var (
	currentPath string // 当前路径
)

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		currentPath = path.Dir(filename)
	}
}

// 在centos8上安装docker
// @param host: 主机地址
// @param username: 用户名
// @param password: 密码
// @param port: 端口号，默认是200
func (k *K8S) UploadInstallDockerOnCentos() {

	if !k.IsHealth() {
		k.Log.Panic("无法成功连接到指定的服务器，请检查后重试。")
		return
	}

	// 上传安装脚本
	localFilePath := fmt.Sprintf("%s/assets/shell/centos/install_docker.sh", currentPath)
	remoteDirPath := fmt.Sprintf("/home/%s", k.Config.Ssh.Username)
	k.Ssh.UploadFile(localFilePath, remoteDirPath)
}

// 上传k8s Flannel配置文件
func (k *K8S) UploadFlannelConfig() {

	if !k.IsHealth() {
		k.Log.Panic("无法成功连接到指定的服务器，请检查后重试。")
		return
	}

	// 上传安装脚本
	localFilePath := fmt.Sprintf("%s/assets/yaml/kube-flannel.yml", currentPath)
	remoteDirPath := fmt.Sprintf("/home/%s", k.Config.Ssh.Username)
	k.Ssh.UploadFile(localFilePath, remoteDirPath)
}
