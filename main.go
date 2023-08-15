package main

import "github.com/ringsaturn/hertz-template-test-panic/app"

func main() {
	h := app.NewHertz()
	_ = h.Run()
}
