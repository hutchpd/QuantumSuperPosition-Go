package quantum

import "testing"

func TestAny(t *testing.T) {
    sp := Any(1, 2, 3)
    if len(sp.Eigenstates()) != 3 {
        t.Errorf("Expected 3 eigenstates, got %d", len(sp.Eigenstates()))
    }
}

func TestAdd(t *testing.T) {
    sp := Add(1, 2)
    if len(sp.Eigenstates()) != 1 {
        t.Errorf("Expected 1 eigenstate, got %d", len(sp.Eigenstates()))
    }
    if sp.Eigenstates()[0] != 3 {
        t.Errorf("Expected result 3, got %v", sp.Eigenstates()[0])
    }
}
