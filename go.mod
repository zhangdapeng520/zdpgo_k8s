module github.com/zhangdapeng520/zdpgo_k8s

go 1.17

require (
	github.com/zhangdapeng520/zdpgo_log v1.2.0
	github.com/zhangdapeng520/zdpgo_ssh v1.0.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/fs v0.1.0 // indirect
	github.com/pkg/sftp v1.13.4 // indirect
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace (
	 zdpgo_ssh v1.0.2 => ../zdpgo_ssh // 重定向到本地，稳定以后再打版本更新
)