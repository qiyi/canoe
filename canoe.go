package canoe

import (
	"github.com/qiyi/etc"
)

const (
	Deck   = "deck"
	Gate   = "gate"
	Driver = "clouddriver"
	Front  = "front50"
	Igor   = "igor"
	Echo   = "echo"
	Rosco  = "rosco"
	Flat   = "flat"
)

// TODO
// * 各模块如何初始化，各模块可以加载不同模块配置，也能读取共用外部依赖配置
// * 初始化时，各模块可以自行注册 HTTP 接口/ 进程内接口
// * 模块间的调用机制，可以调用进程内服务，也可以远程调用

type Canoe interface {
	Env() etc.Environment
	Module(string) Module
}

type Board interface {
	Initialize(Canoe) error
	Destroy(Canoe)
}

type Module interface {
	Board
	Name() string
	Credential() map[string]interface{}
}
