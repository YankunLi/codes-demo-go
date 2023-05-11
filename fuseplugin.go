package main

import "fmt"

type FUSE struct {
	name     string
	version  int
	describe string
}

func (f FUSE) GetName() string {
	return f.name
}

func (f FUSE) SetName(name string) {
	f.name = name
}

var (
	Fuse = FUSE{name: "fuse_v1", version: 1.0, describe: "fuse module"}
)

func init() {
	fmt.Printf("Load %s plugin version %d desc: %s\n", Fuse.name, Fuse.version, Fuse.describe)
}

func ShowModule() {
	fmt.Printf("%v", Fuse)
}
