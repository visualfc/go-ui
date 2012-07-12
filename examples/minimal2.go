package main

import (
	"github.com/visualfc/go-ui/ui"
)

var exit = make(chan bool)

func main() {
	ui.Main(func() {
		go ui_main()
		ui.Run()
		exit <- true
	})
}

func ui_main() {
	w := ui.NewWidget()
	w.SetWindowTitle(ui.Version())
	w.SetSizev(300, 200)
	defer w.Close()
	w.Show()
	<-exit
}
