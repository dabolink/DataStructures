package LinkedList

import (
	"testing"
)

func Test_linkedList_Length(t *testing.T) {
	ll := New()
	length := ll.Length()
	if length != 0 {
		t.Errorf("Length() should be '0' but is %v", length)
	}
	ll.Add("0")
	length = ll.Length()
	if ll.Length() != 1 {
		t.Errorf("Length() should be '1' but is %v", length)
	}
	length = ll.Length()
	ll.Add("1")
	if ll.Length() != 2 {
		t.Errorf("Length() should be '2' but is %v", length)
	}
	length = ll.Length()
	ll.Add("2")
	if ll.Length() != 3 {
		t.Errorf("Length() should be '3' but is %v", length)
	}
	length = ll.Length()
	ll.Add("3")
	if ll.Length() != 4 {
		t.Errorf("Length() should be '4' but is %v", length)
	}
}

func Test_linkedList_Get(t *testing.T) {
	ll := New()
	_, err := ll.Get(0)
	if err != ErrOutOfBounds {
		t.Errorf("err should 'ErrOutOfBounds' is %v", err)
	}
	// add root
	ll.Add("0")
	val, err := ll.Get(0)
	if err != nil {
		t.Errorf("err should be 'nil' is %v", err)
	}
	if val != "0" {
		t.Errorf("val should be '\"0\"' is %v", val)
	}
	// add node
	ll.Add("1")
	val, err = ll.Get(1)
	if err != nil {
		t.Errorf("err should be 'nil' is %v", err)
	}
	if val != "1" {
		t.Errorf("val should be '\"1\"' is %v", val)
	}
}

func Test_linkedList_Remove(t *testing.T) {
	ll := New()
	err := ll.Remove(0)
	if err == nil {
		t.Errorf("err should be ErrOutOfBounds is %v", err)
	}
	if ll.size != 0 {
		t.Errorf("size should be 0 is %v", ll.size)
	}

	ll.Add("0")
	ll.Add("1")
	ll.Add("2")

	err = ll.Remove(1)
	if err != nil {
		t.Errorf("err should be 'nil' is %v", err)
	}
	if ll.size != 2 {
		t.Errorf("size should be 2 is %v", ll.size)
	}
	err = ll.Remove(1)
	if err != nil {
		t.Errorf("err should be 'nil' is %v", err)
	}
	if ll.size != 1 {
		t.Errorf("size should be 1 is %v", ll.size)
	}
	err = ll.Remove(0)
	if err != nil {
		t.Errorf("err should be 'nil' is %v", err)
	}
	if ll.size != 0 {
		t.Errorf("size should be 0 is %v", ll.size)
	}
}
func TestNew(t *testing.T) {
	ll := New()
	if ll.size != 0 {
		t.Errorf("size should be '0' is %v", ll.size)
	}
	if ll.root != nil {
		t.Errorf("root should be 'nil' is %v", ll.root)
	}
}

func Test_linkedList_Add(t *testing.T) {
	ll := New()
	// add root
	ll.Add("0")
	if ll.root == nil {
		t.Errorf("root should not be nil :: %v", ll.root)
	}
	if ll.size != 1 {
		t.Errorf("size should be '1' is %v", ll.size)
	}
	// add node
	ll.Add("1")
	if ll.root == nil {
		t.Errorf("root should not be nil :: %v", ll.root)
	}
	if ll.size != 2 {
		t.Errorf("size should be '2' is %v", ll.size)
	}
}
