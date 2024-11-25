package main

import (
    "fmt"
    "github.com/hutchpd/QuantumSuperPosition-Go/pkg/quantum"
)

func main() {
    sp := quantum.All("cat", "dog", "bird")
    fmt.Printf("Conjunctive superposition eigenstates: %v\n", sp.Eigenstates())

    lt := quantum.LessThan(10, sp)
    fmt.Printf("Values less than 10: %v\n", lt.Eigenstates())
}
