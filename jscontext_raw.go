package quickjs

//#include "quickjs.h"
import "C"

type JSContextRaw struct {
	ref *C.JSContext
}

func JSNewContextRaw(rt *JSRuntime) *JSContextRaw {
	ctx := new(JSContextRaw)
	ctx.ref = C.JS_NewContextRaw(rt.ref)
	return ctx
}

func (m *JSContextRaw) ToContext() JSContext {
	return JSContext{
		ref: m.ref,
	}
}

// no needed.
func (m *JSContextRaw) AddIntrinsicBaseObjects() {
	C.JS_AddIntrinsicBaseObjects(m.ref)
}

func (m *JSContextRaw) AddIntrinsicDate() {
	C.JS_AddIntrinsicDate(m.ref)
}

func (m *JSContextRaw) AddIntrinsicEval() {
	C.JS_AddIntrinsicEval(m.ref)
}

func (m *JSContextRaw) AddIntrinsicStringNormalize() {
	C.JS_AddIntrinsicStringNormalize(m.ref)
}

func (m *JSContextRaw) AddIntrinsicRegExpCompiler() {
	C.JS_AddIntrinsicRegExpCompiler(m.ref)
}

func (m *JSContextRaw) AddIntrinsicRegExp() {
	C.JS_AddIntrinsicRegExp(m.ref)
}

func (m *JSContextRaw) AddIntrinsicJSON() {
	C.JS_AddIntrinsicJSON(m.ref)
}

func (m *JSContextRaw) AddIntrinsicProxy() {
	C.JS_AddIntrinsicProxy(m.ref)
}

func (m *JSContextRaw) AddIntrinsicMapSet() {
	C.JS_AddIntrinsicMapSet(m.ref)
}

func (m *JSContextRaw) AddIntrinsicTypedArrays() {
	C.JS_AddIntrinsicTypedArrays(m.ref)
}

func (m *JSContextRaw) AddIntrinsicPromise() {
	C.JS_AddIntrinsicPromise(m.ref)
}
