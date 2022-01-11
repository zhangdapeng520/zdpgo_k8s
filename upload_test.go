package zdpgo_k8s

import (
	"testing"
)

// 测试：检查健康状态
func TestK8S_UploadInstallDockerOnCentos(t *testing.T) {
	config := K8SConfig{
		Host:        "192.168.33.10",
		Port:        22,
		Username:    "zhangdapeng",
		Password:    "zhangdapeng",
		LogFilePath: "",
		Debug:       true,
	}
	k := New(config)
	k.UploadInstallDockerOnCentos()
	k.UploadFlannelConfig()
}
