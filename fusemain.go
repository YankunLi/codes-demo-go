package main

import (
	"fmt"
	"plugin"
)

type MD interface {
	GetName() string
	SetName(name string)
}

func main() {
	var fc func()
	var m MD
	p, err := plugin.Open("fuseplugin.so")
	if err != nil {
		panic(err)
	}

	f, err := p.Lookup("ShowModule")
	if err != nil {
		panic(err)
	}
	fc = f.(func())

	fc()

	module, err := p.Lookup("Fuse")
	if err != nil {
		panic(err)
	}
	m = module.(MD)

	fmt.Println(m.GetName())

	fmt.Println("laod again\n")

	p, err = plugin.Open("fuseplugin_v1.so")
	if err != nil {
		panic(err)
	}

	fmt.Println(m.GetName())

	module, err = p.Lookup("Fuse")
	if err != nil {
		panic(err)
	}
	m = module.(MD)

	fmt.Println(m.GetName())

}
