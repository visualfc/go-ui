#go-ui
=====

###Introduction
go-ui is a golang ui, base on qt, cross-platform ui tool kit.

###System
Windows / Linux x86 / Linux x86_64 / MacOS X 10.6

###Examples

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


