## GoFiber BootStrap Message Box (gfbmb)

When using [GoFiber](https://github.com/gofiber/fiber) with the [HTML template](https://github.com/gofiber/template/tree/master/html) to serve a basic webpage, a common tool used is BootStrap. Though, there was no simple way to add a BootStrap component like a Message Box to the page through GoFiber.

I solved this issue, by writing a small package for a couple of my projects that returns the required code for a message box. Combining [GoFiber's HTML template](https://github.com/gofiber/template/tree/master/html) and my basic functions I am able to easly insert and render a message box on my html page.

<br>

### Getting Started

1. Get the module in your terminal.
```
go get github.com/SowinskiBraeden/gfbmb
```

2. Ensure your GoFiber engine has `unescape` defined.
```GoLang
engine.AddFunc(
  "unescape", func(s string) template.HTML {
    return template.HTML(s)
  },
)
```

3. Ensure your html page has the required modules.
```html
<link rel="stylesheet" href="https://bootswatch.com/4/journal/bootstrap.min.css"/>
<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/js/bootstrap.min.js"></script>
```

4. Place where you want your Message box to go in your html page.
```html
{{ unescape msgBox }}
```

5. Render your html page with the empty message box.
```GoLang
app.Get("/", func(c *fiber.Ctx) error {
  return c.Render("index", fiber.Map{
    "msgBox": messageBox.EmptyMessageBox(),
  })
})
```

6. When rendering a page with a message, use one of the following...
```GoLang
newMsgBox, err := messageBox.NewMessageBox(message, msgType)
// msgType is optional, providing none will default to 'success'
// msgType can be 'success', 'warning' or 'danger'
if err != nil {
  return c.Render("index", fiber.Map{
    "msgBox": messageBox.NewDangerBox(err.Error()),
  })
}
```
or one of these...
```GoLang
newMsgBox := messageBox.NewSuccessBox(message)
newMsgBox := messageBox.NewWarningBox(message)
newMsgBox := messageBox.NewErrorBox(message)
```

7. Render your html page with your new Message Box.
```GoLang
return c.Render("index", fiber.Map{
  "msgBox": newMsgBox,
})
```

You can see the full example using the `main.go` file.