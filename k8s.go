package zdpgo_k8s

import (
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_ssh"
	"strings"
)

// k8s核心结构体
type K8S struct {
	ssh         *zdpgo_ssh.SSH // ssh连接对象
	log         *zdpgo_log.Log // 日志对象
	logFilePath string         // 日志文件路径
	debug       bool           // 是否为debug模式
}

// k8s配置结构体
type K8SConfig struct {
	Host        string // 主机地址
	Port        int    // SSH端口号
	Username    string // 登录用户名
	Password    string // 登录密码
	LogFilePath string // 日志文件路径
	Debug       bool   // 是否为debug模式
}

func New(config K8SConfig) *K8S {
	k := K8S{}

	// 初始化日志
	if config.LogFilePath == "" {
		k.logFilePath = "zdpgo_k8s.log"
	} else {
		k.logFilePath = config.LogFilePath
	}
	logConfig := zdpgo_log.LogConfig{
		Debug:       config.Debug,
		LogFilePath: k.logFilePath,
	}
	k.log = zdpgo_log.New(logConfig)

	// 校验参数
	if config.Host == "" {
		k.log.Panic("主机地址Host不能为空")
	}
	if config.Username == "" {
		k.log.Panic("SSH登录用户名Username不能为空")
	}
	if config.Password == "" {
		k.log.Panic("SSH登录密码Password不能为空")
	}
	if config.Port == 0 {
		config.Port = 22
	}

	// 初始化ssh连接
	ssh := zdpgo_ssh.New(config.Host, config.Username, config.Password, config.Port)
	ssh.Connect()
	k.ssh = ssh

	return &k
}

// 设置debug模式
func (k *K8S) SetDebug(debug bool) {
	k.debug = debug
	logConfig := zdpgo_log.LogConfig{
		Debug:       debug,
		LogFilePath: k.logFilePath,
	}
	k.log = zdpgo_log.New(logConfig)
}

// 判断是否为debug模式
func (k *K8S) IsDebug() bool {
	return k.debug
}

// 判断是否为健康状态
func (k *K8S) IsHealth() bool {
	command := "echo 'ok'"
	result, err := k.ssh.Run(command)
	flag := strings.Trim(result, "\n") == "ok" && err == nil
	return flag
}
