// Copyright 2019 Roger Chapman and the v8go contributors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package v8go

//go:generate clang-format -i --verbose -style=Chromium v8go.h v8go.cc

// #cgo CXXFLAGS: -fno-rtti -fpic -std=c++14 -DV8_COMPRESS_POINTERS -DV8_31BIT_SMIS_ON_64BIT_ARCH
// #cgo darwin linux CXXFLAGS: -I${SRCDIR}/deps/include
// #cgo LDFLAGS: -pthread -lv8
// #cgo windows LDFLAGS: -lv8_libplatform
// #cgo darwin LDFLAGS: -L${SRCDIR}/deps/darwin_x86_64
// #cgo linux LDFLAGS: -L${SRCDIR}/deps/linux_x86_64
import "C"
