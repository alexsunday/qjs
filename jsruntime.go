package quickjs

//#include "quickjs.h"
import "C"

type JSRuntime struct {
	ref    *C.JSRuntime
	ctxSet []*JSContext
}

func NewJsRuntime() *JSRuntime {
	out := &JSRuntime{}
	out.ref = C.JS_NewRuntime()
	//runtime.SetFinalizer(out, (*out).Close)
	return out
}

func (m *JSRuntime) NewContext() *JSContext {
	ctx := NewJSContext(m)
	m.ctxSet = append(m.ctxSet, ctx)
	return ctx
}

func (m *JSRuntime) removeContext(ctx *JSContext) {
	var pos = -1
	for idx, item := range m.ctxSet {
		if item == ctx {
			pos = idx
		}
	}
	if pos != -1 {
		m.ctxSet = append(m.ctxSet[:pos], m.ctxSet[pos+1:]...)
	}
}

func (m *JSRuntime) Close() {
	for {
		if len(m.ctxSet) > 0 {
			m.ctxSet[0].Close()
		} else {
			break
		}
	}

	C.JS_FreeRuntime(m.ref)
}
