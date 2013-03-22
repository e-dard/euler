package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := 1.0
    delta := func() float64 { return (z * z - x) / (2.0 * z) }
    
    for d := delta(); math.Abs(d) >= 0.001; d = delta() {
        z -= d
    }
    return z
}

func main() {
    fmt.Println(euler.Sqrt(2))
}