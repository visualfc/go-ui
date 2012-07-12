package main

import (
	"fmt"
	"github.com/visualfc/go-ui/ui"
)

var exit = make(chan bool)

func main() {
	fmt.Println(ui.Version())
	ui.Main(func() {
		go main_ui()
		ui.Run()
		exit <- true
	})
}

func main_ui() {
ui.OnInsertObject(func(item interface{}) {
		fmt.Println("add item", item)
	})
	ui.OnRemoveObject(func(item interface{}) {
		fmt.Println("remove item", item)
	})

	w := ui.NewWidget()
	defer w.Close()

	w.SetWindowTitle("This is a test")
	fmt.Println(w.WindowTitle())

	vbox := ui.NewVBoxLayout()
	fmt.Println(vbox)
	w.SetLayout(vbox)

	lbl := ui.NewLabel()
	lbl.SetText("<h2><i>Hello</i> <font color=blue><a href=\"ui\">UI</a></font></h2>")
	lbl.OnLinkActivated(fnTEST)
	vbox.AddWidget(lbl)
	vbox.AddStretch(0)

	//runtime.GC()

	btn := ui.NewButton()
	btn.SetText("WbcdefgwqABCDEFQW")
	font := btn.Font()
	defer font.Close()
	font.SetPointSize(16)
	btn.SetFont(font)
	fmt.Println("f3->", btn.Font())

	btn2 := ui.NewButton()
	font.SetPointSize(18)
	btn2.SetAttr("text", "WbcdefgwqABCDEFQW")
	btn2.SetAttr("font", font)

	btn.OnClicked(func() {
		fmt.Println(btn)
		btn.Close()
	})
	btn.OnCloseEvent(func(e *ui.CloseEvent) {
		fmt.Println("Close", e)
	})
	btn3 := ui.NewButton()
	btn3.SetText("Exit")
	btn3.OnClicked(func() {
		ui.Exit(0)
	})

	l := w.Layout()
	fmt.Println("ll", l)
	l.AddWidget(btn)
	l.AddWidget(btn2)
	l.AddWidget(btn3)
	//vbox.AddWidget(btn)
	f := btn2.Attr("parent")
	fmt.Println("parent->", f, f == nil)

	fmt.Println(btn.Font())

	w.OnResizeEvent(func(e *ui.ResizeEvent) {
		fmt.Println(e)
	})

	w.OnPaintEvent(func(ev *ui.PaintEvent) {
		fnPaint(ev, w)
	})

	//w.Show()
	w.SetVisible(true)
	<-exit
}

func fnPaint(ev *ui.PaintEvent, w *ui.Widget) {
	p := ui.NewPainter()
	defer p.Close()
	p.Begin(w)
	p.DrawPoint(ui.Pt(10, 10))
	p.DrawLine(ui.Pt(10, 10), ui.Pt(100, 100))
	p.End()
}

func fnTEST(link string) {
	fmt.Println("link:", link)
}
