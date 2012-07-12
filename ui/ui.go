// Copyright 2012 visualfc <visualfc@gmail.com>. All rights reserved.
//
// This library is free software; you can redistribute it and/or
// modify it under the terms of the GNU Lesser General Public
// License as published by the Free Software Foundation; either
// version 2.1 of the License, or (at your option) any later version.
//
// This library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// Lesser General Public License for more details.

package ui

func About() string {
	return "go-ui 0.1 <visualfc@gmail.com>"
}

func Version() string {
	return "go-ui 0.1"
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
