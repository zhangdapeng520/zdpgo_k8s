package zdpgo_k8s

import (
	"fmt"
	"testing"
)

// 测试：检查健康状态
func TestK8S_IsHealth(t *testing.T) {
	config := K8SConfig{
		Host:        "192.168.33.10",
		Port:        22,
		Username:    "zhangdapeng",
		Password:    "zhangdapeng",
		LogFilePath: "",
		Debug:       true,
	}
	k := New(config)
	fmt.Println(k.IsHealth())
}
