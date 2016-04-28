package main

import (
  "fmt"
  "log"

  "github.com/PuerkitoBio/goquery"
  
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
)

func main() {
  m := martini.Classic()

  m.Use(render.Renderer())

  m.Get("/", func(r render.Render) {
    r.HTML(200, "hello", nil)
  })

  m.Get("/u/", func(r render.Render, params martini.Params) {

    doc, err := goquery.NewDocument("https://fr.wikipedia.org/wiki/New_York") 

    if err != nil {
      log.Fatal(err)
    }

    doc.Find("h1").Each(func(i int, s *goquery.Selection) {

      fmt.Printf(s.Text())
    })
  })

  m.Run()
}