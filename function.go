// Copyright 2021 Roger Chapman and the v8go contributors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package v8go

// #include "v8go.h"
import "C"
import (
	"errors"
	"runtime"
	"unsafe"
)

// Function is a JavaScript function.
type Function struct {
	*Value
}

// NewFunction creates a Function for a given callback.
func NewFunction(iso *Isolate, callback FunctionCallback) (*Function, error) {
	if iso == nil {
		return nil, errors.New("v8go: failed to create new FunctionTemplate: Isolate cannot be <nil>")
	}
	if callback == nil {
		return nil, errors.New("v8go: failed to create new FunctionTemplate: FunctionCallback cannot be <nil>")
	}

	cbref := iso.registerCallback(callback)

	tmpl := &template{
		ptr: C.NewFunctionTemplate(iso.ptr, C.int(cbref)),
		iso: iso,
	}
	runtime.SetFinalizer(tmpl, (*template).finalizer)
	return &FunctionTemplate{tmpl}, nil
}

// Call this JavaScript function with the given arguments.
func (fn *Function) Call(args ...Valuer) (*Value, error) {
	var argptr *C.ValuePtr
	if len(args) > 0 {
		var cArgs = make([]C.ValuePtr, len(args))
		for i, arg := range args {
			cArgs[i] = arg.value().ptr
		}
		argptr = (*C.ValuePtr)(unsafe.Pointer(&cArgs[0]))
	}
	rtn := C.FunctionCall(fn.ptr, C.int(len(args)), argptr)
	return getValue(fn.ctx, rtn), getError(rtn)
}
