#go-ui
=====

##Introduction
go-ui is a cross-platform golang ui tool kit, based on qt.

##System
Windows / Linux / MacOS X

##Build go-ui

###1.build qtdrv
    $ cd go-ui/qtdrv
    $ qmake "CONFIG+=release"
    $ make
###2.build go-ui
    $ cd go-ui/ui
    $ go install

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


