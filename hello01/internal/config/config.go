// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import "github.com/zeromicro/go-zero/rest"

// 源码将yaml转为了json处理，可使用json tag，变量名大小写敏感。
type Config struct {
	rest.RestConf
	Address string `json:"address"`
}
