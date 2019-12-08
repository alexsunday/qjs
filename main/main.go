package main

import (
	"fmt"
	"github.com/alexsunday/quickjs"
)

func test() {
	rt := quickjs.NewJsRuntime()
	defer rt.Close()

	ctx := rt.NewContext()
	fmt.Println(ctx)

	ctx.EvalCode("1+1===1")

	fmt.Println(rt)
}

func main() {
	test()
}
