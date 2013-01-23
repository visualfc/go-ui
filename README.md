#go-ui 0.1.1
=====

##Introduction
go-ui is a cross-platform golang ui tool kit, based on qt.

Lua code that generates the wrappers (`uiobjs.go` and `cdrv.cpp`) can be found under `make`.

The wrapper code is by far incomplete. Adding new functionality usually consists of editing or adding files under `make/ui`, and updating `make/make.lua` script, and making relevant changes in `ui` and `qtdrv`.

##System
Windows / Linux / MacOS X

##License
    go-ui lib BSD
    qtdrv lib LGPL

##Build go-ui and examples

###1.get go-ui
    $ go get github.com/visualfc/go-ui
###2.build qtdrv, need install QtSDK
    $ cd go-ui/qtdrv
    $ qmake "CONFIG+=release"
    $ make
###3.build go-ui
    $ cd go-ui/ui
    $ go install
###4.build examples
    $ cd go-ui/examples
    $ go build -ldflags '-r ../lib' minimal.go
    $ ./minimal

##Examples

    package main

    import (
	    "github.com/visualfc/go-ui/ui"
    )
    
    func main() {
	    ui.Main(func() {
		    w := ui.NewWidget()
		    w.SetWindowTitle(ui.Version())
		    w.SetSizev(300, 200)
		    defer w.Close()
		    w.Show()
		    ui.Run()
	    })
    }


