package handler

import (
	"net/http"

	"iron-go/internal/logic/category"
	"iron-go/internal/svc"
	"iron-go/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func AddCategoryHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCategoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddCategoryLogic(r.Context(), ctx)
		resp, err := l.AddCategory(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
