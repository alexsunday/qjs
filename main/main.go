package main

import (
	"fmt"
	"quickjs"
)

func test() {
	rt := quickjs.NewJsRuntime()
	defer rt.Close()

	ctx := rt.NewContext()
	fmt.Println(ctx)

	ctx.EvalCode("console.log('1+1=', 1+1)")

	fmt.Println(rt)
}

func main() {
	test()
}
