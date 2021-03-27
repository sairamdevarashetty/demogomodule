package main

import (
	model "demoproject/model"
	model2 "demoproject/src/srcmodel"
	"fmt"
)

func main() {
	fmt.Println("%+v", model.Employee{})
	fmt.Printf("%+v\n", model.Employee{})
	fmt.Printf("%+v \n", model2.Employee2{12, "sai", "ram", 21})
}
