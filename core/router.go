package core

import (
  "fmt"
  "log"

  "github.com/PuerkitoBio/goquery"
  
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
)

func index (r render.Render) {
  r.HTML(200, "hello", nil)
}

func process (r render.Render, params martini.Params) {

  doc, err := goquery.NewDocument("http://designspartan.com/actualite/lancement-de-la-campagne-de-crowdfunding-de-digitalpainting-school/") 

  if err != nil {
    log.Fatal(err)
  }

  doc.Find("article").Each(func(i int, s *goquery.Selection) {
    if i == 0 {
      fmt.Printf("%+v\n", s.First().Text()) 
    }
  })
}