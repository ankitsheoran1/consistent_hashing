package main

import (
	"sort"
	"testing"
)

func checkNum(num, expected int, t *testing.T) {
	if num != expected {
		t.Errorf("got %d, expected %d", num, expected)
	}
}

func TestNew(t *testing.T) {
	x := New()
	if x == nil {
		t.Errorf("expected obj")
	}
	checkNum(x.NumberOfReplicas, 20, t)
}

func TestAdd(t *testing.T) {
	x := New()
	x.AddNode("abcdefg")
	checkNum(len(x.sortedHashes), 20, t)
	checkNum(len(x.sortedHashes), 20, t)
	if sort.IsSorted(x.sortedHashes) == false {
		t.Errorf("expected sorted hashes to be sorted")
	}

	x.AddNode("qwer")
	checkNum(len(x.sortedHashes), 40, t)

	if sort.IsSorted(x.sortedHashes) == false {
		t.Errorf("expected sorted hashes to be sorted")
	}
}

func TestRemove(t *testing.T) {
	x := New()
	x.AddNode("abcdefg")
	x.AddNode("qwer")
	x.RemoveNode("qwer")
	checkNum(len(x.sortedHashes), 20, t)
}

func TestRemoveNonExisting(t *testing.T) {
	x := New()
	x.AddNode("abcdefg")
	x.RemoveNode("abcdefghijk")
	checkNum(len(x.sortedHashes), 0, t)
}

func TestGetEmpty(t *testing.T) {
	x := New()
	_, err := x.Get("asdfsadfsadf")
	if err == nil {
		t.Errorf("expected error")
	}
	if err != ErrEmptyCircle {
		t.Errorf("expected empty circle error")
	}
}

func TestGetSingle(t *testing.T) {
	x := New()
	x.AddNode("abcdefg")
	x.AddNode("defgh")
	x.AddNode("mnop")

	y, err := x.Get("abcde")
	if err != nil {
		t.Errorf(err.Error())
	}
	if y.addr != "mnop" {
		t.Errorf("expectation failed" + y.addr)
	}
}
