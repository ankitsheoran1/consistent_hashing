package main

import (
	"fmt"
	"log"
)

func ExampleNew() {
	c := New()
	c.AddNode("cacheA")
	c.AddNode("cacheB")
	c.AddNode("cacheC")
	users := []string{"user_mcnulty", "user_bunk", "user_omar", "user_bunny", "user_stringer"}
	for _, u := range users {
		server, err := c.Get(u)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s => %s\n", u, server.addr)
	}
	// Output:
	// user_mcnulty => cacheC
	// user_bunk => cacheA
	// user_omar => cacheA
	// user_bunny => cacheC
	// user_stringer => cacheB
}

func ExampleAdd() {
	c := New()
	c.AddNode("cacheA")
	c.AddNode("cacheB")
	c.AddNode("cacheC")
	users := []string{"user_mcnulty", "user_bunk", "user_omar", "user_bunny", "user_stringer"}
	fmt.Println("initial state [A, B, C]")
	for _, u := range users {
		server, err := c.Get(u)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s => %s\n", u, server.addr)
	}
	c.AddNode("cacheD")
	c.AddNode("cacheE")
	fmt.Println("\nwith cacheD, cacheE [A, B, C, D, E]")
	for _, u := range users {
		server, err := c.Get(u)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s => %s\n", u, server.addr)
	}
	// Output:
	// initial state [A, B, C]
	// user_mcnulty => cacheC
	// user_bunk => cacheA
	// user_omar => cacheA
	// user_bunny => cacheC
	// user_stringer => cacheB
	//
	// with cacheD, cacheE [A, B, C, D, E]
	// user_mcnulty => cacheC
	// user_bunk => cacheE
	// user_omar => cacheE
	// user_bunny => cacheC
	// user_stringer => cacheB
}

func ExampleRemove() {
	c := New()
	c.AddNode("cacheA")
	c.AddNode("cacheB")
	c.AddNode("cacheC")
	users := []string{"user_mcnulty", "user_bunk", "user_omar", "user_bunny", "user_stringer"}
	fmt.Println("initial state [A, B, C]")
	for _, u := range users {
		server, err := c.Get(u)
		if err != nil {
			log.Fatal(err, u)
		}
		fmt.Printf("%s => %s\n", u, server.addr)
	}
	c.RemoveNode("cacheC")
	fmt.Println("\ncacheC removed [A, B]")
	for _, u := range users {
		server, err := c.Get(u)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%s => %s\n", u, server.addr)

	}
	// Output:
	// initial state [A, B, C]
	// user_mcnulty => cacheC
	// user_bunk => cacheA
	// user_omar => cacheA
	// user_bunny => cacheC
	// user_stringer => cacheB
	//
	// cacheC removed [A, B]
	// user_mcnulty => cacheA
	// user_bunk => cacheA
	// user_omar => cacheA
	// user_bunny => cacheA
	// user_stringer => cacheB
}
