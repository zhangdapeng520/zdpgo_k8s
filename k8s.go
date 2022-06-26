package zdpgo_k8s

import (
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_ssh"
)

// K8S 核心结构体
type K8S struct {
	Ssh    *zdpgo_ssh.SSH // ssh连接对象
	Log    *zdpgo_log.Log // 日志对象
	Config *Config        // k8s配置
}

func New(log *zdpgo_log.Log) *K8S {
	return NewWithConfig(&Config{}, log)
}

func NewWithConfig(config *Config, log *zdpgo_log.Log) *K8S {
	k := &K8S{}

	// 初始化日志
	k.Log = log

	// 配置
	k.Config = config

	// 初始化ssh连接
	k.Ssh = zdpgo_ssh.NewWithConfig(config.Ssh, log)

	return k
}

// IsHealth 判断是否为健康状态
func (k *K8S) IsHealth() bool {
	return k.Ssh.Status()
}

// GetRc 获取RC列表
func (k *K8S) GetRc() (string, error) {
	executeResult, err := k.Ssh.Sudo("kubectl get rc")
	if err != nil {
		k.Log.Error("获取RC列表失败", "error", err)
		return "", err
	}
	k.Log.Debug("获取RC列表成功", "result", executeResult)
	return executeResult, nil
}

// GetSvc 获取Svc列表
func (k *K8S) GetSvc() (string, error) {
	executeResult, err := k.Ssh.Sudo("kubectl get svc")
	if err != nil {
		k.Log.Error("获取Service列表失败", "error", err)
		return "", err
	}
	k.Log.Debug("获取Service列表成功", "result", executeResult)
	return executeResult, nil
}
