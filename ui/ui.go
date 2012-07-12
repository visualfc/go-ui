// Copyright 2012 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

func About() string {
	return "go-ui 0.1.1 <visualfc@gmail.com>"
}

func Version() string {
	return "go-ui 0.1.1"
}

func Main(fn func()) int {
	fnAppMain = fn
	return theApp.AppMain()
}

func Run() int {
	return theApp.Run()
}

func Exit(code int) {
	theApp.Exit(code)
}

func CloseAllWindows() {
	theApp.CloseAllWindows()
}

func App() *app {
	return &theApp
}
