package zdpgo_k8s

import (
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_ssh"
	"testing"
)

// 测试：检查健康状态
func TestK8S_IsHealth(t *testing.T) {
	k := NewWithConfig(&Config{Ssh: &zdpgo_ssh.Config{
		Host:     "192.168.33.10",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	}}, zdpgo_log.Tmp)
	if k == nil {
		panic("new error")
	}
}

// 测试获取RC列表
func TestK8S_GetRc(t *testing.T) {
	k := NewWithConfig(&Config{Ssh: &zdpgo_ssh.Config{
		Host:     "192.168.33.10",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	}}, zdpgo_log.Tmp)
	rc, err := k.GetRc()
	if err != nil {
		panic(err)
	}
	if rc == "" {
		panic("rc结果为空")
	}
}

// 测试获取Service列表
func TestK8S_GetSvc(t *testing.T) {
	k := NewWithConfig(&Config{Ssh: &zdpgo_ssh.Config{
		Host:     "192.168.33.10",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	}}, zdpgo_log.Tmp)
	rc, err := k.GetSvc()
	if err != nil {
		panic(err)
	}
	if rc == "" {
		panic("svc结果为空")
	}
}
