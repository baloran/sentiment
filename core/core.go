package core

import (
  "fmt"
  
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
)

func init () {
  fmt.Printf("Sentiment \n")
}

func Sentiment () {
  m := martini.Classic()

  m.Use(render.Renderer())

  m.Get("/", index)

  m.Get("/u/", process)

  m.Run()
}