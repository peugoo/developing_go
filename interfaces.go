package main

import (
	"fmt"
)

type InterOcean interface{
	Addfish()
	Deletefish()
}

type Ocean struct {
	Creatures []string
	fish_c int
}

func (o *Ocean) Addfish() {
	o.fish_c ++
}

func (o *Ocean) Deletefish() {
	o.fish_c --
}

func log(a string, d &InterOcean){
	fmt.Print(a, d)
}

func main() {
	o := Ocean{
		[]string{
			"sea urchin",
			"lobster",
			"shark",
		},
		123,
	}
	o.Addfish()
	log("asd", &o)
}
