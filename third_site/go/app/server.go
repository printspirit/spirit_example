package app

import (
	_ "fmt"
	"github.com/kardianos/osext"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func file_svr(root, minitype, file string) func(c echo.Context) error {
	return func(c echo.Context) error {
		filePath := filepath.Join(root, file)
		content, err := ioutil.ReadFile(filePath)
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
	lst, _ := GetList("user1")
	return c.Render(http.StatusOK, "list.html", lst)
}

func edit(c echo.Context) error {
	subclass := c.QueryParam("subclass")
	tpid := c.QueryParam("tpid")
	url, err := GetEditUrl(subclass, tpid)
	if err != nil {
		return c.Render(http.StatusOK, "err.html", err.Error())
	}
	if c.QueryParam("target") == "new" {
		return c.Redirect(http.StatusFound, url)
	} else {
		return c.Render(http.StatusOK, "edit.html", url)
	}
}

func Start() {

	root, _ := osext.ExecutableFolder()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define a renderer for templates
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = renderer

	e.GET("/style.css", file_svr(root, "text/css", "style.css"))
	e.GET("/", list)
	e.GET("/list", list)
	e.GET("/edit", edit)

	e.Logger.Fatal(e.Start(":8000"))
}
