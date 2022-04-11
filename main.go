package main

import (
	"html/template"
	"log"

	"github.com/SowinskiBraeden/gfbmb/messageBox"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

// Example
func main() {
	// Example engine
	engine := html.New("./exampleViews", ".html")

	// Take not of this AddFunc code,
	// This is from GoFiber enabling us to render html
	engine.AddFunc(
		// add unescape function
		"unescape", func(s string) template.HTML {
			return template.HTML(s)
		},
	)

	// After you created your engine, you can pass it to Fiber's Views Engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Index Example
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"msgBox": messageBox.EmptyMessageBox(),
		})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		msgType := c.FormValue("messageboxtype")
		// msgType is a string value of
		// 'success', 'warning' or 'danger'
		message := c.FormValue("message")

		// You can use...
		newMsgBox, err := messageBox.NewMessageBox(message, msgType)
		// msgType is optional, providing none will default to 'success'
		if err != nil {
			return c.Render("index", fiber.Map{
				"msgBox": messageBox.NewDangerBox(err.Error()),
			})
		}
		// To make a new message box

		// Or you can use any of the following:
		// newMsgBox := messageBox.NewSuccessBox(message)
		// newMsgBox := messageBox.NewWarningBox(message)
		// newMsgBox := messageBox.NewErrorBox(message)

		return c.Render("index", fiber.Map{
			"msgBox": newMsgBox,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
