package zdpgo_k8s

import (
	"fmt"
	"testing"
)

func testCreateK8s() *K8S {
	config := K8SConfig{
		Host:        "192.168.33.10",
		Port:        22,
		Username:    "zhangdapeng",
		Password:    "zhangdapeng",
		LogFilePath: "",
		Debug:       true,
	}
	k := New(config)
	if k != nil {
		k.log.Info("链接服务器成功。")
	}
	return k
}

// CreateRC 创建RC
// @param fileName 配置文件的名称
func TestCreateRC(t *testing.T) {
	k := testCreateK8s()
	result, err := k.CreateRC("mysql-rc.yaml")
	fmt.Println(result, err)
}

// CreateRC 创建Service
func TestCreateService(t *testing.T) {
	k := testCreateK8s()
	result, err := k.CreateService("mysql-svc.yaml")
	fmt.Println(result, err)
}

// CreateWeb 创建tomcat web服务
func TestCreateWeb(t *testing.T) {
	k := testCreateK8s()
	result, err := k.CreateRC("myweb-rc.yaml")
	fmt.Println(result, err)
	result, err = k.CreateService("myweb-svc.yaml")
	fmt.Println(result, err)
}

// TestCreateMysql 测试创建MySQL
func TestCreateMysql(t *testing.T) {
	k := testCreateK8s()
	result, err := k.CreateMysql(8, false)
	fmt.Println(result, err)
}
