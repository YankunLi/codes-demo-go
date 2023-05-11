package main

import "fmt"

type SS interface {
	Show()
}

type AA struct {
	Name string
}

func (p *AA) Show() {
	fmt.Println(p.Name)
}

type BB struct {
	Name string
	Age  int
}

func (p *BB) Show() {
	fmt.Println(p.Name, " ", p.Age)
}

func main() {
	var ss SS
	ss = &AA{Name: "AA"}
	ss.Show()
	ss = &BB{Name: "BB", Age: 12}
	ss.Show()
	fmt.Println("vim-go")
}
