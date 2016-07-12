package core

import (
  "fmt"
  "log"

  "github.com/PuerkitoBio/goquery"

  "github.com/kataras/iris"
)

func index (ctx *iris.Context) {
  ctx.Text(iris.StatusOK, "Plain text here")
}

func process (ctx *iris.Context) {

  doc, err := goquery.NewDocument("http://www.liberation.fr/france/2016/04/29/loi-travail-un-manifestant-eborgne-par-un-tir-de-flash-ball-a-rennes_1449427")

  if err != nil {
    log.Fatal(err)
  }

  s := doc.Find("article p")
  for i := range s.Nodes {
    single := s.Eq(i)
    fmt.Printf("%+v\n", len(single.Text()))
    if len(single.Text()) > 500 {
      fmt.Printf("%+v\n", single.Text())
    }
  }


}
