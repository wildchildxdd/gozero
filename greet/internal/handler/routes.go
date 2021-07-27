// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"GOZero/greet/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/albums/:albums",
				Handler: AlbumHandler(serverCtx),
			},
		},
	)
}
