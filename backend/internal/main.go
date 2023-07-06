// Code generated by hertz generator.

package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
	"github.com/wegoteam/weflow/internal/biz/conf"
	"github.com/wegoteam/weflow/internal/biz/router"
	"github.com/wegoteam/weflow/internal/docs"
	"gopkg.in/natefinch/lumberjack.v2"
)

// @title weflow
// @version 1.0
// @description weflow swagger api documention.

// @contact.name hertz-contrib
// @contact.url https://github.com/hertz-contrib

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:18080
// @BasePath /weflow
// @schemes http
func main() {
	//hertz
	hertz := conf.GetConf().Hertz
	address := hertz.Address
	basePath := hertz.BasePath
	h := server.New(server.WithHostPorts(address), server.WithBasePath(basePath))
	//swagger 文档
	swag := conf.GetConf().Swagger
	if swag.Enable {
		docs.SwaggerInfo.Title = swag.Title
		docs.SwaggerInfo.Description = swag.Description
		docs.SwaggerInfo.Version = swag.Version
		docs.SwaggerInfo.Host = swag.Host
		docs.SwaggerInfo.BasePath = swag.BasePath
		docs.SwaggerInfo.Schemes = swag.Schemes
		swagURL := fmt.Sprintf("http://%s/%s/swagger/doc.json", swag.Host, swag.BasePath)
		swaggerURL := swagger.URL(swagURL)
		h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, swaggerURL, swagger.DefaultModelsExpandDepth(-1)))
	}
	router.GeneratedRegister(h)
	registerMiddleware(h)
	h.Spin()
}

// registerMiddleware
// @Description: 注册中间件
// @param: h
func registerMiddleware(h *server.Hertz) {
	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}
	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// log
	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	hlog.SetOutput(&lumberjack.Logger{
		Filename:   conf.GetConf().Hertz.LogFileName,
		MaxSize:    conf.GetConf().Hertz.LogMaxSize,
		MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
		MaxAge:     conf.GetConf().Hertz.LogMaxAge,
	})

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())
}
