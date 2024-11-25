
# QuantumSuperPosition

QuantumSuperPosition is a Go module that replicates the C# module `QuantumSuperposition`. It allows you to create superpositions of values and perform operations on them as if they were in multiple states simultaneously.

## Features

- Create disjunctive (`Any`) and conjunctive (`All`) superpositions of values.
- Perform arithmetic operations (`Add`, `Subtract`, `Multiply`, `Divide`, `Modulo`) on superpositions.
- Perform comparison operations (`LessThan`, `GreaterThan`, `EqualTo`) on superpositions.
- Evaluate truthiness of superpositions with `IsTrue`.
- Extract eigenstates (possible values) with `Eigenstates`.
- String representation of superpositions.

## Usage Examples

### Example 1: Basic Arithmetic Operations

```go
package main

import (
    "fmt"
    "quantum"
)

func main() {
    a := quantum.Any(1, 2, 3)
    b := quantum.All(4, 5)
    
    sum := quantum.Add(a, b)
    fmt.Println("Sum:", sum) // Output: Sum: all([5 6 6 7 7 8])
    
    product := quantum.Multiply(a, b)
    fmt.Println("Product:", product) // Output: Product: all([4 5 8 10 12 15])
}
```

### Example 2: Comparison Operations

```go
package main

import (
    "fmt"
    "quantum"
)

func main() {
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
}
```

### Example 3: Primality Testing

```go
package main

import (
    "fmt"
    "math"
    "quantum"
)

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

func main() {
    number := 29
    if isPrime(number) {
        fmt.Printf("%d is prime\n", number) // This will print
    } else {
        fmt.Printf("%d is not prime\n", number)
    }
}
```

### Example 4: Set Intersection

```go
package main

import (
    "fmt"
    "quantum"
)

func main() {
    set1 := quantum.Any(1, 2, 3, 4)
    set2 := quantum.Any(3, 4, 5, 6)
    intersection := quantum.EqualTo(set1, set2)
    fmt.Println("Intersection:", intersection.Eigenstates()) // Output: Intersection: [3 4]
}
```

## Notes

- Arithmetic and comparison functions are used as Go does not support operator overloading.
- The `Superposition` type holds a slice of values (`values`) and a type (`typ`) indicating whether it's disjunctive or conjunctive.

## Limitations

- Currently supports numerical types (integers and floats) for operations.
- Additional handling for other types like strings or custom structs can be added.
- Error handling can be improved based on specific needs. (seriously error handling in go is weird, or is it just me?)
