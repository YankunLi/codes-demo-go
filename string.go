package main

import (
	"fmt"
	"strings"
)

type S struct {
	id string
}

func (p *S) String() string {
	return fmt.Sprintf("%s--", p.id)
}

func main() {
	s := S{id: "hello world"}
	fmt.Printf("s format #v: %#v\n", &s)
	fmt.Printf("s format v: %v\n", &s)
	fmt.Printf("s format s: %s\n", &s)
	strs := strings.Split("/sdf/", "/")
	for i, v := range strs {
		fmt.Println(i, v)
	}

}
