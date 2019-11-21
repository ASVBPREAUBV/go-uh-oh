package main

import (
	"fmt"
	"strconv"
)

func main() {
	var listOfNumberStrings StringContainerArray

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("#%s", strconv.Itoa(i))
		othername := fmt.Sprintf("#%s", strconv.Itoa(i))
		listOfNumberStrings = append(listOfNumberStrings, &StringContainer{name: name, othername: othername})
	}

	for _, n := range listOfNumberStrings {
		fmt.Printf("%s\n", n)
	}

	return
}

type StringContainerArray []*StringContainer

type StringContainer struct {
	name      string
	othername string
}

func (p StringContainer) String() string {
	return fmt.Sprintf("{name:%s, othername:%s}", p.name, p.othername)
}
