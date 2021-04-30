package kcdgin

import (
	"context"
	"github.com/alexisvisco/kcd"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() {
	kcd.Config.StringsExtractors = append(kcd.Config.StringsExtractors, GinPathExtractor{})
}

type GinPathExtractor struct{}

func (g GinPathExtractor) Extract(req *http.Request, res http.ResponseWriter, valueOfTag string) ([]string, error) {
	params := req.Context().Value("gin-params")

	if params != nil {
		p, ok := params.(gin.Params)
		if ok {
			name, ok := p.Get(valueOfTag)

			if ok {
				return []string{name}, nil
			}
		}
	}

	return nil, nil
}

func (g GinPathExtractor) Tag() string {
	return "path"
}

func Handler(h interface{}, defaultStatusCode int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), "gin-params", ctx.Params))
		kcd.Handler(h, defaultStatusCode)(ctx.Writer, req)
	}
}
