package set

import "testing"

func TestSetHoldsItem(t *testing.T) {
    list := New()
    list.Add(100)
    if list.Size() != 1 {
        t.Error("Should have 1 element.")
    }

    if !list.Contains(100) {
        t.Error("Should've had 100 as the element")
    }
}

func TestDoesntHoldDuplicates(t *testing.T) {
    list := New()
    list.Add(100)
    if list.Size() != 1 {
        t.Error("Should have 1 element.")
    }
    if !list.Contains(100) {
        t.Error("Should've had 100 as the element")
    }
    if list.Add(100) {
        t.Error("Shouldn't have added 100")
    }
    if list.Size() != 1 {
        t.Error("Should still have 1 element.")
    }
    if !list.Contains(100) {
        t.Error("Should've had 100 as the element")
    }
}

func TestRemovesItem(t *testing.T) {
    list := New()
    list.Add(100)
    if list.Size() != 1 {
        t.Error("Should have 1 element.")
    }

    if !list.Contains(100) {
        t.Error("Should've had 100 as the element")
    }

    list.Remove(100)
    if list.Size() != 0 {
        t.Error("Should have 0 element.")
    }

    if list.Contains(100) {
        t.Error("Shouldn't had 100 as the element")
    }
}

func TestGetsAsArray(t *testing.T) {
    list := New()
    for i:=0; i<100; i++ {
        list.Add(i)
    }

    if list.Size() != 100 {
        t.Error("Should have 100 elements.")
    }

    arr := list.AsArray()
    if len(arr) != 100 {
        t.Error("arr should have 100 elements")
    }
}

func TestContainsReturnsFalse(t *testing.T) {
    list := New()

    if list.Contains(100) {
        t.Error("Shouldn't had 100 as the element")
    }
}

func TestClear(t *testing.T) {
    list := New()
    for i:=0; i<100; i++ {
        list.Add(i)
    }

    if list.Size() != 100 {
        t.Error("Should have 100 elements.")
    }

    list.Clear()
    if list.Size() != 0 {
        t.Error("Should have 0 elements.")
    }


}
