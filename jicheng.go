package main

import "fmt"

type A struct {
	c int
	b int
}

func (p *A) String() string {
	return fmt.Sprintf("a: %d b: %d", p.c, p.b)
}

type C struct {
	c int
}

type B struct {
	C
	A
	c int
}

func (p *B) String() string {
	return fmt.Sprintf("a: %d b: %d", p.c, p.b)
}

func main() {
	var a1 A
	fmt.Println(a1)
	var b1 B
	fmt.Println(b1)

	a1.c = 12
	fmt.Println(b1)

	b1.A = a1
	fmt.Println(b1)
	b1.A.c = 13
	b1.c = 14
	fmt.Println(b1)

	fmt.Println("vim-go")
}
