package zdpgo_k8s

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_ssh"
	"testing"
)

func testCreateK8s() *K8S {
	k := NewWithConfig(&Config{Ssh: &zdpgo_ssh.Config{
		Host:     "192.168.33.10",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	}}, zdpgo_log.Tmp)
	if k != nil {
		k.Log.Info("链接服务器成功。")
	}
	return k
}

// CreateRC 创建RC
// @param fileName 配置文件的名称
func TestCreateRC(t *testing.T) {
	k := NewWithConfig(&Config{Ssh: &zdpgo_ssh.Config{
		Host:     "192.168.33.10",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	}}, zdpgo_log.Tmp)
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
