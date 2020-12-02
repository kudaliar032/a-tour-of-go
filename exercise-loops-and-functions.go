package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  z := 1.0
  var t float64
  for {
    z, t = z - ((z * z - x) / (2 * z)), z
    if math.Abs(t - z) < 1e-6 {
      break
    }
  }
  return z
}

func main() {
  value := float64(643783)
  guess := Sqrt(value)
  expected := math.Sqrt(value)
  fmt.Printf("Guest: %v, Expected: %v, Error: %v\n", guess, expected, math.Abs(guess - expected))
}
