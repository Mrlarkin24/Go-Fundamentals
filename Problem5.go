package main

import (
    "fmt"
    "math/rand"
)

// range specification, note that min <= max
type IntRange struct {
    min, max int
}

// get next random value within the interval including min and max
func (ir *IntRange) NextRandom(r* rand.Rand) int {
    return r.Intn(ir.max - ir.min +1) + ir.min
}

func main() {
    r := rand.New(rand.NewSource(55))
    ir := IntRange{-1,1}
    for i := 0; i<10; i++ {
        fmt.Println(ir.NextRandom(r))
    }
}

//was trying to find a number generator when i ran out of time
//Resource: https://stackoverflow.com/questions/23577091/generating-random-numbers-over-a-range-in-go