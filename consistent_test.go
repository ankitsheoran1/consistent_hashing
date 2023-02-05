package main

import (
	"fmt"
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

func TestAddC(t *testing.T) {
	x := New()
	x.AddNode("abcdefg")
	checkNum(len(x.sortedHashes), 20, t)
	checkNum(len(x.sortedHashes), 20, t)
	if sort.IsSorted(x.sortedHashes) == false {
		t.Errorf("expected sorted hashes to be sorted")
	}

	fmt.Println("===============", x.sortedHashes)

	x.AddNode("qwer")
	checkNum(len(x.sortedHashes), 40, t)
	fmt.Println("===============", x.sortedHashes)

	if sort.IsSorted(x.sortedHashes) == false {
		t.Errorf("expected sorted hashes to be sorted")
	}
}

func TestRemoveC(t *testing.T) {
	x := New()
	x.AddNode("abcdefg")
	fmt.Println("===============", x.sortedHashes)
	x.AddNode("qwer")
	x.RemoveNode("qwer")
	fmt.Println("===============", x.sortedHashes)
	//checkNum(len(x.circle), 0, t)
	checkNum(len(x.sortedHashes), 20, t)
}

func TestRemoveNonExistingC(t *testing.T) {
	x := New()
	x.AddNode("abcdefg")
	x.RemoveNode("abcdefghijk")
	checkNum(len(x.sortedHashes), 0, t)
}

func TestGetEmptyC(t *testing.T) {
	x := New()
	_, err := x.Get("asdfsadfsadf")
	if err == nil {
		t.Errorf("expected error")
	}
	if err != ErrEmptyCircle {
		t.Errorf("expected empty circle error")
	}
}

func TestGetSingleC(t *testing.T) {
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
