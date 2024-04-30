package main

import (
	"github.com/labstack/echo/v4"
	"go-todo/models"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	e := echo.New()
	renderer := &TemplateRenderer {
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/list", list)
	e.POST("/create", create)
	e.GET("/edit", edit)
	e.POST("/update", update)
	e.GET("/destroy", destroy)

	e.Logger.Fatal(e.Start(":1323"))
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t * TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func destroy(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		log.Fatalln(err)
	}
	models.DeleteTodo(id)
	return c.Redirect(http.StatusMovedPermanently, "/list")
}

func update(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	content := c.FormValue("content")
	todo, _ := models.GetTodo(id)
	todo.Content = content
	models.UpdateTodo(todo)

	return c.Redirect(http.StatusMovedPermanently, "/list")
}

func edit(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		log.Fatalln(err)
	}
	todo, _ := models.GetTodo(id)
	data := map[string]interface{} {
		"todo": todo,
	}

	return c.Render(http.StatusOK, "edit.html", data)
}

func create(c echo.Context) error {
	content := c.FormValue("content")
	models.CreateTodo(content)
	return c.Redirect(http.StatusMovedPermanently, "/list")
}

func list(c echo.Context) error {
	var todos []models.Todo
	models.Db.Find(&todos)

	data := map[string]interface{} {
		"todos": todos,
	}
	return c.Render(http.StatusOK, "list.html", data)
}
