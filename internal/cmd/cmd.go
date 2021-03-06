package cmd

import (
	"context"

	"github.com/gogf/gf/v2/protocol/goai"

	"go-ops/internal/consts"
	"go-ops/internal/controller"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(MiddlewareCORS)

				//	group.Middleware(controller.MiddlewareGetApp)
				group.Middleware(ghttp.MiddlewareHandlerResponse)

				group.Bind(
					controller.Auth,
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					// 打开这里使用认证功能
					group.Middleware(controller.AuthUser)
					group.Bind(
						controller.ScritptTask,
						controller.PeerManagaer,
						controller.DownloadFileTask,
						controller.App,
						controller.Task,
						controller.Script,
						controller.Agent,
						controller.CheckItem,
						controller.TaskPreset,
						controller.TaskCron,
						controller.Plugin,
						controller.VM,
						controller.User,
					)
				})

			})

			s.SetIndexFolder(true)
			s.AddSearchPath("public")
			s.AddSearchPath("swagger")
			s.AddStaticPath("/public", "public")
			s.AddStaticPath("/swagger", "swagger")
			enhanceOpenAPIDoc(s)
			s.Run()
			return nil
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {

	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription + consts.OpenAPISign,
		Contact: &goai.Contact{
			Name: "LUXINGWEN",
			URL:  "https://github.com/luxingwen",
		},
	}
}
