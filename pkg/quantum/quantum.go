// Package quantum provides functionality similar to Perl's Quantum::Superpositions module.
// It allows creating superpositions of values and performing operations on them.
package quantum

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"time"
)

// SuperpositionType represents the type of superposition: Disjunctive (any) or Conjunctive (all).
type SuperpositionType int

const (
	Disjunctive SuperpositionType = iota // Represents a disjunctive superposition (any state can be true).
	Conjunctive                          // Represents a conjunctive superposition (all states must be true).
)

// Superposition represents a quantum superposition of values.
type Superposition struct {
	values []interface{}
	typ    SuperpositionType
}

// Any creates a disjunctive superposition of the provided values.
func Any(values ...interface{}) *Superposition {
	return &Superposition{
		values: values,
		typ:    Disjunctive,
	}
}

// All creates a conjunctive superposition of the provided values.
func All(values ...interface{}) *Superposition {
	return &Superposition{
		values: values,
		typ:    Conjunctive,
	}
}

// Add performs addition between two operands, which can be scalars or superpositions.
func Add(a, b interface{}) *Superposition {
	resultValues, resultTyp := combineValues(a, b, addValues)
	return &Superposition{
		values: resultValues,
		typ:    resultTyp,
	}
}

// Subtract performs subtraction between two operands, which can be scalars or superpositions.
func Subtract(a, b interface{}) *Superposition {
	resultValues, resultTyp := combineValues(a, b, subValues)
	return &Superposition{
		values: resultValues,
		typ:    resultTyp,
	}
}

// Multiply performs multiplication between two operands, which can be scalars or superpositions.
func Multiply(a, b interface{}) *Superposition {
	resultValues, resultTyp := combineValues(a, b, mulValues)
	return &Superposition{
		values: resultValues,
		typ:    resultTyp,
	}
}

// Divide performs division between two operands, which can be scalars or superpositions.
func Divide(a, b interface{}) *Superposition {
	resultValues, resultTyp := combineValues(a, b, divValues)
	return &Superposition{
		values: resultValues,
		typ:    resultTyp,
	}
}

// Modulo performs modulus operation between two operands, which can be scalars or superpositions.
func Modulo(a, b interface{}) *Superposition {
	resultValues, resultTyp := combineValues(a, b, modValues)
	return &Superposition{
		values: resultValues,
		typ:    resultTyp,
	}
}

// LessThan performs a comparison between two operands and returns a superposition of values from the left operand that satisfy the comparison.
func LessThan(a, b interface{}) *Superposition {
	resultValues, resultTyp := compareValues(a, b, lessThanValues)
	return &Superposition{
		values: resultValues,
		typ:    resultTyp,
	}
}

// GreaterThan performs a comparison between two operands and returns a superposition of values from the left operand that satisfy the comparison.
func GreaterThan(a, b interface{}) *Superposition {
	resultValues, resultTyp := compareValues(a, b, greaterThanValues)
	return &Superposition{
		values: resultValues,
		typ:    resultTyp,
	}
}

// EqualTo performs a comparison between two operands and returns a superposition of values from the left operand that satisfy the comparison.
func EqualTo(a, b interface{}) *Superposition {
	resultValues, resultTyp := compareValues(a, b, equalToValues)
	return &Superposition{
		values: resultValues,
		typ:    resultTyp,
	}
}

// IsTrue evaluates the truth value of the superposition based on its type and values.
func (s *Superposition) IsTrue() bool {
	if len(s.values) == 0 {
		return false
	}
	return true
}

// Eigenstates returns the eigenstates (possible values) of the superposition.
func (s *Superposition) Eigenstates() []interface{} {
	return s.values
}

// String returns a string representation of the superposition.
func (s *Superposition) String() string {
	if len(s.values) == 1 {
		return fmt.Sprintf("%v", s.values[0])
	}
	return fmt.Sprintf("%s(%v)", s.typ.String(), s.values)
}

// String returns a string representation of the SuperpositionType.
func (st SuperpositionType) String() string {
	switch st {
	case Disjunctive:
		return "any"
	case Conjunctive:
		return "all"
	default:
		return "unknown"
	}
}

// Helper functions to perform arithmetic operations on values.

func addValues(a, b interface{}) (interface{}, error) {
	return performArithmetic(a, b, '+')
}

func subValues(a, b interface{}) (interface{}, error) {
	return performArithmetic(a, b, '-')
}

func mulValues(a, b interface{}) (interface{}, error) {
	return performArithmetic(a, b, '*')
}

func divValues(a, b interface{}) (interface{}, error) {
	return performArithmetic(a, b, '/')
}

func modValues(a, b interface{}) (interface{}, error) {
	return performArithmetic(a, b, '%')
}

// Helper functions to perform comparison operations on values.

func lessThanValues(a, b interface{}) (bool, error) {
	return performComparison(a, b, '<')
}

func greaterThanValues(a, b interface{}) (bool, error) {
	return performComparison(a, b, '>')
}

func equalToValues(a, b interface{}) (bool, error) {
	return performComparison(a, b, '=')
}

// combineValues combines two operands using the provided operation function.
func combineValues(a, b interface{}, op func(a, b interface{}) (interface{}, error)) ([]interface{}, SuperpositionType) {
	aValues, aTyp := extractValues(a)
	bValues, bTyp := extractValues(b)

	resultValues := []interface{}{}

	// Perform operation on combinations of values
	for _, av := range aValues {
		for _, bv := range bValues {
			res, err := op(av, bv)
			if err != nil {
				continue
			}
			resultValues = append(resultValues, res)
		}
	}

	// Determine the result Superposition type
	resultTyp := Disjunctive
	if aTyp == Conjunctive || bTyp == Conjunctive {
		resultTyp = Conjunctive
	}

	return resultValues, resultTyp
}

// compareValues compares two operands using the provided comparison function.
func compareValues(a, b interface{}, comp func(a, b interface{}) (bool, error)) ([]interface{}, SuperpositionType) {
	aValues, aTyp := extractValues(a)
	bValues, bTyp := extractValues(b)

	resultValues := []interface{}{}

	// For each combination, check the comparison
	for _, av := range aValues {
		match := false
		for _, bv := range bValues {
			res, err := comp(av, bv)
			if err != nil {
				continue
			}
			if res {
				match = true
				break // For disjunctive b, one match suffices
			} else if bTyp == Conjunctive {
				// If b is conjunctive and one comparison fails, then the overall fails
				match = false
				break
			}
		}
		if match {
			resultValues = append(resultValues, av)
		}
	}

	resultTyp := aTyp
	return resultValues, resultTyp
}

// extractValues extracts values and type from an operand, which can be a scalar or a superposition.
func extractValues(val interface{}) ([]interface{}, SuperpositionType) {
	if s, ok := val.(*Superposition); ok {
		return s.values, s.typ
	}
	return []interface{}{val}, Disjunctive
}

// performArithmetic performs arithmetic operations between two values.
func performArithmetic(a, b interface{}, op rune) (interface{}, error) {
	aVal, aKind := getNumericValue(a)
	bVal, bKind := getNumericValue(b)

	if aKind == reflect.Invalid || bKind == reflect.Invalid {
		return nil, errors.New("unsupported type in arithmetic operation")
	}

	switch op {
	case '+':
		return aVal + bVal, nil
	case '-':
		return aVal - bVal, nil
	case '*':
		return aVal * bVal, nil
	case '/':
		if bVal == 0 {
			return nil, errors.New("division by zero")
		}
		return aVal / bVal, nil
	case '%':
		if bVal == 0 {
			return nil, errors.New("modulo by zero")
		}
		return math.Mod(aVal, bVal), nil
	default:
		return nil, errors.New("unsupported operation")
	}
}

// performComparison performs comparison operations between two values.
func performComparison(a, b interface{}, op rune) (bool, error) {
	aVal, aKind := getNumericValue(a)
	bVal, bKind := getNumericValue(b)

	if aKind == reflect.Invalid || bKind == reflect.Invalid {
		return false, errors.New("unsupported type in comparison")
	}

	switch op {
	case '<':
		return aVal < bVal, nil
	case '>':
		return aVal > bVal, nil
	case '=':
		return aVal == bVal, nil
	default:
		return false, errors.New("unsupported comparison")
	}
}

// getNumericValue converts an interface{} to a float64 for arithmetic operations.
func getNumericValue(val interface{}) (float64, reflect.Kind) {
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int()), v.Kind()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(v.Uint()), v.Kind()
	case reflect.Float32, reflect.Float64:
		return v.Float(), v.Kind()
	default:
		return 0, reflect.Invalid
	}
}

// RandomValue returns a random eigenstate from the superposition.
func (s *Superposition) RandomValue() interface{} {
	rand.Seed(time.Now().UnixNano())
	if len(s.values) == 0 {
		return nil
	}
	return s.values[rand.Intn(len(s.values))]
}
