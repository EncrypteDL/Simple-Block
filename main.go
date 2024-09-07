package main

import (
	"SimpleBlock/internal"
	"encoding/json"
	"fmt"
)

func main() {
	c := internal.NewChain()
	c.Write([]byte("foo"))
	c.Write([]byte("bar"))
	if !c.Verify() {
		panic("Blockchain verification failure")
	}

	data, _ := json.Marshal(c)
	fmt.Print(string(data))
}
