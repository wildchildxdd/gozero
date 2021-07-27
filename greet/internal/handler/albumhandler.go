package handler

import (
	"net/http"

	"GOZero/greet/internal/logic"
	"GOZero/greet/internal/svc"
	"GOZero/greet/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func AlbumHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAlbumLogic(r.Context(), ctx)
		resp, err := l.Album(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
