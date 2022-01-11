package zdpgo_k8s

import (
	"fmt"
	"testing"
)

// 测试：安装StorageClass
func TestK8S_InstallStorageClass(t *testing.T) {
	config := K8SConfig{
		Host:        "192.168.33.10",
		Port:        22,
		Username:    "zhangdapeng",
		Password:    "zhangdapeng",
		LogFilePath: "",
		Debug:       true,
	}
	k := New(config)
	result, err := k.InstallStorageClass()
	fmt.Println(result, err)
}

// 测试：安装Zookeeper
func TestK8S_InstallZookeeper(t *testing.T) {
	config := K8SConfig{
		Host:        "192.168.33.10",
		Port:        22,
		Username:    "zhangdapeng",
		Password:    "zhangdapeng",
		LogFilePath: "",
		Debug:       true,
	}
	k := New(config)
	result, err := k.InstallZookeeper()
	fmt.Println(result, err)
}

// 测试：安装Kafka
func TestK8S_InstallKafka(t *testing.T) {
	config := K8SConfig{
		Host:        "192.168.33.10",
		Port:        22,
		Username:    "zhangdapeng",
		Password:    "zhangdapeng",
		LogFilePath: "",
		Debug:       true,
	}
	k := New(config)
	result, err := k.InstallKafka()
	fmt.Println(result, err)
}

// 测试：安装KafkaManager
func TestK8S_InstallKafkaManager(t *testing.T) {
	config := K8SConfig{
		Host:        "192.168.33.10",
		Port:        22,
		Username:    "zhangdapeng",
		Password:    "zhangdapeng",
		LogFilePath: "",
		Debug:       true,
	}
	k := New(config)
	result, err := k.InstallKafkaManager()
	fmt.Println(result, err)
}
