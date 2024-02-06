package main

import (
	"github.com/labstack/echo/v4"

	_ "net/http"
)

func main() {
	// Membuat instance Echo
	e := echo.New()

	//mengatur folder statis untuk menyajikan file css
	e.Static("/static", "asset")

	// // Menentukan template
	// templates := template.Must(template.ParseFiles("templates/index.html"))

	// Menambahkan route
	e.GET("/", func(c echo.Context) error {
		// return templates.Execute(c.Response(), nil)
		return c.File("template/index.html")
	})

	// Menjalankan server
	e.Start(":8090")
}
