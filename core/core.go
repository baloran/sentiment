package core

import (
  "fmt"

  "github.com/kataras/iris"
)

func init () {
  fmt.Printf("Sentiment \n")
}

func Sentiment () {

  iris.Get("/", index)

  iris.Get("/u/", process)

  iris.Listen(":8080")
}
