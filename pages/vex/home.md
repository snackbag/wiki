# Vex

Vex is a simple Go UI library based on [raylib-go](https://github.com/gen2brain/raylib-go). It implements a widget
system, while not limiting you on how you want your items to be scaled. It comes with a CSS-like styling system while
not having to emulate a browser (at all. it's 100% native)

This is how to open your first window
```go
package main

import "github.com/snackbag/vex"

func main() {
	process := vex.Init("Title", 400, 400)
	process.SetAllowResize(true)
	
	process.Run()
}
```

HINT: Once you initialized your program you can also use `vex.Process` to access the process globally
