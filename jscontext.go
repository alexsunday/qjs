package quickjs

//#include "quickjs.h"
import "C"

type JSContext struct {
	rt  *JSRuntime
	ref *C.JSContext
}

func NewJSContext(rt *JSRuntime) *JSContext {
	out := &JSContext{}
	out.rt = rt
	out.ref = C.JS_NewContext(rt.ref)
	return out
}

/*
Close
1, remove from JSRuntime
2, free self
*/
func (m *JSContext) Close() {
	m.rt.removeContext(m)
	m.rt = nil

	C.JS_FreeContext(m.ref)
}

func (m *JSContext) EvalCode(code string) {
	inputLen := C.size_t(len(code))
	filename := C.CString("code")
	evalFlags := C.int(0)
	C.JS_Eval(m.ref, C.CString(code), inputLen, filename, evalFlags)
}
