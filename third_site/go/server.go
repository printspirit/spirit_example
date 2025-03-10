package main

import (
	"embed"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"net/http"
)

//go:embed statics/*
var statics embed.FS

func file_svr(file, minitype string) func(c echo.Context) error {
	return func(c echo.Context) error {
		content, err := statics.ReadFile("statics/" + file)
		if err != nil {
			return err // 处理读取文件错误
		}
		c.Response().Header().Set("Content-Type", minitype)
		return c.String(http.StatusOK, string(content))
	}
}

// TemplateRenderer defines the interface for rendering HTML templates.
type TemplateRenderer struct {
	templates *template.Template
}

// Render implements the Renderer interface for Echo.
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func list(c echo.Context) error {
	lst, err := site.GetList("user1")
	if err != nil {
		return c.Render(http.StatusOK, "err.html", err.Error())
	}
	return c.Render(http.StatusOK, "list.html", map[string]any{"host": site.Host, "list": lst})
}

func edit(c echo.Context) error {
	subclass := c.QueryParam("subclass")
	tpid := c.QueryParam("tpid")
	url, err := site.GetEditUrl(subclass, tpid)
	if err != nil {
		return c.Render(http.StatusOK, "err.html", err.Error())
	}
	if c.QueryParam("target") == "new" {
		return c.Redirect(http.StatusFound, url)
	} else {
		return c.Render(http.StatusOK, "edit.html", url)
	}
}

func new(c echo.Context) error {
	subclass := c.QueryParam("subclass")
	tpid, err := site.NewLabel("测试标签", 500, 800, 203, subclass, "")
	if err != nil {
		return c.Render(http.StatusOK, "err.html", err.Error())
	}
	if c.QueryParam("target") == "new" {
		url, err := site.GetEditUrl(subclass, tpid)
		if err != nil {
			return c.Render(http.StatusOK, "err.html", err.Error())
		}
		return c.Redirect(http.StatusFound, url)
	} else {
		return c.Redirect(http.StatusTemporaryRedirect, "edit?tpid="+tpid)
	}
}

func del(c echo.Context) error {
	tpid := c.QueryParam("tpid")
	err := site.DelLabel(tpid)
	if err != nil {
		return c.Render(http.StatusOK, "err.html", err.Error())
	}
	return c.Redirect(http.StatusTemporaryRedirect, "list")
}

func Start(port int) {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define a renderer for templates
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseFS(statics, "statics/*.html")),
	}
	e.Renderer = renderer

	e.GET("/style.css", file_svr("style.css", "text/css"))
	e.GET("/", list)
	e.GET("/list", list)
	e.GET("/edit", edit)
	e.GET("/new", new)
	e.GET("/del", del)

	err := e.Start(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("服务器启动失败\n请检查端口 %d 是否被占用， 并使用 -p 参数指定新端口\n", port)
	}
}
