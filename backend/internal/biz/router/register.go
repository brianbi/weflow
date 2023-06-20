// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	comm "github.com/wegoteam/weflow/internal/biz/router/comm"
	flow "github.com/wegoteam/weflow/internal/biz/router/flow"
	insttask "github.com/wegoteam/weflow/internal/biz/router/insttask"
	model "github.com/wegoteam/weflow/internal/biz/router/model"
	usertask "github.com/wegoteam/weflow/internal/biz/router/usertask"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(hertz *server.Hertz) {
	//注册公共路由
	comm.Register(hertz)
	//注册流程路由
	flow.Register(hertz)
	//注册模板路由
	model.Register(hertz)
	//注册用户任务路由
	usertask.Register(hertz)
	//注册实例任务路由
	insttask.Register(hertz)
}
