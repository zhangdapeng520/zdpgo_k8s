package main

import (
	"github.com/zhangdapeng520/zdpgo_k8s"
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_ssh"
)

/*
@Time : 2022/6/26 14:43
@Author : 张大鹏
@File : main
@Software: Goland2021.3.1
@Description: 部署Redis实例
*/

func main() {
	// 创建k8s对象
	k := zdpgo_k8s.NewWithConfig(&zdpgo_k8s.Config{Ssh: &zdpgo_ssh.Config{
		Host:     "192.168.33.10",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	}}, zdpgo_log.Tmp)

	// 部署Redis
	err := k.DeployRedis()
	if err != nil {
		k.Log.Error("部署Redis失败", "error", err)
	}
}
