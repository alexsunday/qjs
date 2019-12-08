# QuickJS Golang binding

example:
```go
package main

import (
	"github.com/alexsunday/quickjs"
)

func main()  {
    rt := quickjs.NewJsRuntime()
    defer rt.Close()
    ctx := rt.NewContext()
    ctx.EvalCode("console.log('1+1=', 1+1)")
}
```
