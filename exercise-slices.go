package main

import "github.com/golang-id/tour/pic"

func Pic(dx, dy int) [][]uint8 {
  pixel := make([][]uint8, dy)

  for x := range pixel {
    data := make([]uint8, dx)
    for y := range data {
      data[y] = uint8((x + y) / 2)
    }
    pixel[x] = data
  }
  return pixel
}

func main() {
  pic.Show(Pic)
}
