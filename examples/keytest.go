package main

import (
	"fmt"
	"github.com/visualfc/go-ui/ui"
)

func printInfo() {
	info := ui.Value("info")
	fmt.Println(info)
}

func setInfo() {
	ui.SetValue("info", "new info")
}

func main() {
	ui.Main(ui_main)
}

func ui_main() {

	w := ui.NewWidget()

	lbox := ui.NewVBoxLayout()
	lbl1 := ui.NewLabel()
	lbl1.SetText("This is a info1")
	lbl2 := ui.NewLabel()
	lbl2.SetText("This is a info2")

	ed1 := ui.NewLineEdit()

	printInfo := func() {
		info := ui.Value("info")
		ed1.SetText(fmt.Sprint(info))
	}

	lbox.AddWidget(lbl1)
	lbox.AddWidget(lbl2)
	lbox.AddWidget(ed1)

	rbox := ui.NewVBoxLayout()

	btn1 := ui.NewButton()
	btn1.SetText("Change")

	btn2 := ui.NewButton()
	btn2.SetText("Value")

	btn3 := ui.NewButton()
	btn3.SetText("SetValue")

	rbox.AddWidget(btn1)
	rbox.AddWidget(btn2)
	rbox.AddWidget(btn3)

	b := true
	btn1.OnClicked(func() {
		var text string
		if b {
			ui.SetKey(lbl1, "info", "text")
			text = "info1"
		} else {
			ui.SetKey(lbl2, "info", "text")
			text = "info2"
		}
		b = !b
		btn1.SetText(text)
	})

	btn2.OnClicked(printInfo)
	btn3.OnClicked(setInfo)

	hbox := ui.NewHBoxLayout()
	hbox.AddLayout(lbox)
	hbox.AddLayout(rbox)

	w.SetLayout(hbox)

	w.Show()

	ui.Run()
}
