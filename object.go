package main

import "fmt"

type SimpleObject struct {
}

type ExampleObject struct {
	num int
}

func main() {
	var obj SimpleObject
	obj.Object = new(ExampleObject)
	eobj := obj.(*ExampleObject)
	fmt.Printf("%d", eobj.num)
	fmt.Println("vim-go")
}
