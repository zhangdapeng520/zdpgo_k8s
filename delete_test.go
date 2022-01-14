package zdpgo_k8s

import (
	"fmt"
	"testing"
)

// TestDeleteMysql 测试删除MySQL
func TestDeleteMysql(t *testing.T) {
	k := testCreateK8s()
	result, err := k.DeleteMysql(8, false)
	fmt.Println(result, err)
}
