package hello

import (
  "net/http"
  "github.com/go-martini/martini"
)

func init() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Post("/tea", func() (int, string) {
    return 418, "I'm a teapot."
  })
  http.Handle("/", m)
}
