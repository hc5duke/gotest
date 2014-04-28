package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Post("/tea", func() (int, string) {
    return 418, "i'm a teapot" // HTTP 418 : "i'm a teapot"
  })
  m.Run()
}
