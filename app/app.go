package app

import (
	"context"
	"embed"
	"html/template"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

//go:embed template/*
var f embed.FS

func NewHertz() *server.Hertz {
	h := server.Default()
	h.SetHTMLTemplate(template.Must(template.New("").ParseFS(f, "template/*.html")))

	h.GET("/hello", func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(http.StatusOK, "hello.html", nil)
	})
	return h
}
