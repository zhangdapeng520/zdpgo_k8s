package zdpgo_k8s

import (
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_ssh"
	"testing"
)

/*
@Time : 2022/6/26 22:13
@Author : 张大鹏
@File : mysql_test
@Software: Goland2021.3.1
@Description: k8s操作MySQL相关测试
*/

func TestK8S_DeleteMysqlRc(t *testing.T) {
	k := NewWithConfig(&Config{Ssh: &zdpgo_ssh.Config{
		Host:     "192.168.33.10",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	}}, zdpgo_log.Tmp)

	k.DeleteMysqlRc()
}

func TestK8S_DeleteMysqlSvc(t *testing.T) {
	k := NewWithConfig(&Config{Ssh: &zdpgo_ssh.Config{
		Host:     "192.168.33.10",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	}}, zdpgo_log.Tmp)

	k.DeleteMysqlSvc()
}
