package zdpgo_k8s

import "github.com/zhangdapeng520/zdpgo_ssh"

/*
@Time : 2022/6/26 12:03
@Author : 张大鹏
@File : config
@Software: Goland2021.3.1
@Description:
*/

// Config 配置信息
type Config struct {
	Ssh *zdpgo_ssh.Config `yaml:"ssh" json:"ssh"`
}
