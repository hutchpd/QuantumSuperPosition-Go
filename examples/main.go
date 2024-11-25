package main

import (
    "fmt"
	"math"
    "github.com/hutchpd/QuantumSuperPosition-Go/pkg/quantum"
)

func main() {
    // Example 1: Basic Arithmetic Operations
    fmt.Println("Example 1: Basic Arithmetic Operations")
    a := quantum.Any(1, 2, 3)
    b := quantum.All(4, 5)
    sum := quantum.Add(a, b)
    fmt.Println("Sum:", sum) // Output: Sum: all([5 6 6 7 7 8])
    product := quantum.Multiply(a, b)
    fmt.Println("Product:", product) // Output: Product: all([4 5 8 10 12 15])

    // Example 2: Comparison Operations
    fmt.Println("\nExample 2: Comparison Operations")
    x := 5
    if quantum.EqualTo(x, quantum.Any(3, 4, 5)).IsTrue() {
        fmt.Println("x is 3, 4, or 5") // This will print
    }
    thresholds := []int{10, 15, 20}
    nextVal := 12
    if quantum.LessThan(nextVal, quantum.All(thresholds)).IsTrue() {
        fmt.Println("nextVal is less than all thresholds")
    } else {
        fmt.Println("nextVal is not less than all thresholds") // This will print
    }

    // Example 3: Primality Testing
    fmt.Println("\nExample 3: Primality Testing")
    number := 29
    if isPrime(number) {
        fmt.Printf("%d is prime\n", number) // This will print
    } else {
        fmt.Printf("%d is not prime\n", number)
    }

    // Example 4: Set Intersection
    fmt.Println("\nExample 4: Set Intersection")
    set1 := quantum.Any(1, 2, 3, 4)
    set2 := quantum.Any(3, 4, 5, 6)
    intersection := quantum.EqualTo(set1, set2)
    fmt.Println("Intersection:", intersection.Eigenstates()) // Output: Intersection: [3 4]
}

// Helper function for primality testing
func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    divisors := quantum.All()
    for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
        divisors = quantum.Add(divisors, i)
    }
    modResults := quantum.Modulo(n, divisors)
    return !modResults.IsTrue()
}
