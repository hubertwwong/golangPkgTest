package main

import (
  "github.com/hubertwwong/golangPkgTest/threeandfour/three"
  "github.com/hubertwwong/golangPkgTest/threeandfour/four"

  "fmt"
)

func main() {
  //xs := []float64{1,2,3,4}
  three := three.Three()
  fmt.Println(three)
  four := four.Four()
  fmt.Println(four)
}

/*
simple test of using 3 and 4 as seperate packages.
remember 1 directory 1 package.
*/