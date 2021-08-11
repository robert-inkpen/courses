package main

import (
	"fmt"
)

// YOUR CODE:
func changeFoo(opts *Options) {
	opts.foo = 2
}
func changeBar(opts *Options) {
	opts.bar = "hello"
}
func changeBarDiff(opts *Options) {
	opts.bar = "there"
}
func main() {
	DoSomething(&Input{}, changeFoo, changeBar)
	DoSomething(&Input{}, changeFoo, changeBarDiff)
}

// AWS CODE:
type Options struct {
	foo int
	bar string
}
type Input struct{}

func DoSomething(input *Input, salamander ...func(*Options)) {
	opts := Options{
		foo: 1,
		bar: "sausage",
	}
	fmt.Printf("Functions %[1]T: %[1]v\n", salamander)
	fmt.Printf("Options: %+v\n", opts)
	for _, function := range salamander { // optFuncs []func(*Options)
		function(&opts)
		fmt.Printf("Options: %+v\n", opts)
	}
}
