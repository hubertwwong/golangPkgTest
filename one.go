package main

import (
  "github.com/hubertwwong/golangPkgTest/one"
  "fmt"
)

func main() {
  fmt.Println(one.GetOne())
  fmt.Println(one.GetTwo())
}

/*
Test of having multiple files in a package
Seems like go just pulls every file in the package and you can access it all at once.
*/