package main

import (
	"fmt"
	"github.com/visualfc/go-ui/ui"
	"io/ioutil"
	"runtime"
)

type MainWindow struct {
	ui.Widget
	tab  *ui.TabWidget
	sbar *ui.StatusBar
}

func (p *MainWindow) Init() *MainWindow {
	if p.Widget.Init() == nil {
		return nil
	}
	p.SetWindowTitle("MainWindow")

	p.tab = ui.NewTabWidget()

	p.tab.AddTab(p.createStdTab(), "Standard", nil)
	p.tab.AddTab(p.createMyTab(), "Custom", nil)
	p.tab.AddTab(p.createToolBox(), "ToolBox", nil)

	p.sbar = ui.NewStatusBar()

	menubar := ui.NewMenuBar()
	menu := ui.NewMenuWithTitle("&File")
	//menu.SetTitle("&File")
	menubar.AddMenu(menu)

	act := ui.NewAction()
	act.SetText("&Quit")
	act.OnTriggered(func(bool) {
		p.Close()
	})
	ic := ui.NewIconWithFile("images/close.png")
	//defer ic.Close()
	act.SetIcon(ic)
	menu.AddAction(act)

	toolBar := ui.NewToolBar()
	toolBar.AddAction(act)
	toolBar.AddSeparator()
	cmb := ui.NewComboBox()
	cmb.AddItem("test1")
	cmb.AddItem("test2")
	cmb.SetToolTip("ComboBox")
	cmbAct := toolBar.AddWidget(cmb)
	fmt.Println(cmbAct)

	vbox := ui.NewVBoxLayout()
	vbox.SetMargin(0)
	vbox.SetSpacing(0)
	vbox.SetMenuBar(menubar)
	vbox.AddWidget(toolBar)
	vbox.AddWidget(p.tab)
	vbox.AddWidget(p.sbar)

	p.SetLayout(vbox)

	p.tab.OnCurrentChanged(func(index int) {
		p.sbar.ShowMessage("current: "+p.tab.TabText(index), 0)
	})

	systray := ui.NewSystemTray()
	systray.SetContextMenu(menu)
	systray.SetIcon(ic)
	systray.SetVisible(true)
	systray.ShowMessage("hello", "this is a test", ui.Information, 1000)
	ic2 := systray.Icon()
	fmt.Println(ic2)

	p.SetWindowIcon(ic2)

	return p
}

func (p *MainWindow) createStdTab() *ui.Widget {
	w := ui.NewWidget()
	vbox := ui.NewVBoxLayout()
	w.SetLayout(vbox)

	ed := ui.NewLineEdit()
	ed.SetInputMask("0000-00-00")
	ed.SetText("2012-01-12")

	lbl := ui.NewLabel()
	lbl.SetText("Label")
	btn := ui.NewButton()
	btn.SetText("Button")
	chk := ui.NewCheckBox()
	chk.SetText("CheckBox")
	radio := ui.NewRadio()
	radio.SetText("Radio")
	cmb := ui.NewComboBox()
	cmb.AddItem("001")
	cmb.AddItem("002")
	cmb.AddItem("003")
	cmb.SetCurrentIndex(2)
	fmt.Println(cmb.CurrentIndex())
	cmb.OnCurrentIndexChanged(func(v int) {
		fmt.Println(cmb.ItemText(v))
	})

	slider := ui.NewSlider()
	slider.SetTickInterval(50)
	slider.SetTickPosition(ui.TicksBothSides)
	slider.SetSingleStep(1)

	scl := ui.NewScrollBar()
	fmt.Println(slider.Range())

	dial := ui.NewDial()

	dial.SetNotchesVisible(true)
	dial.SetNotchTarget(10)
	fmt.Println(dial.NotchSize())

	vbox.AddWidget(ed)
	vbox.AddWidget(lbl)
	vbox.AddWidget(btn)
	vbox.AddWidget(chk)
	vbox.AddWidget(radio)
	vbox.AddWidget(cmb)
	vbox.AddWidget(slider)
	vbox.AddWidget(scl)
	vbox.AddWidget(dial)
	vbox.AddStretch(0)
	return w
}

func (p *MainWindow) createToolBox() ui.IWidget {
	tb := ui.NewToolBox()
	tb.AddItem(ui.NewButtonWithText("button"), "btn", nil)
	tb.AddItem(ui.NewLabelWithText("Label\nInfo"), "Label", nil)
	pixmap := ui.NewPixmapWithFile("images/liteide128.png")
	//defer pixmap.Close()
	lbl := ui.NewLabel()
	lbl.SetPixmap(pixmap)
	tb.AddItem(lbl, "Lalel Pixmap", nil)
	buf, err := ioutil.ReadFile("images/liteide128.png")
	if err == nil {
		pixmap2 := ui.NewPixmapWithData(buf)
		tb.AddItem(ui.NewLabelWithPixmap(pixmap2), "Lalel Pixmap2", nil)
	}
	return tb
}

func (p *MainWindow) createMyTab() *ui.Widget {
	w := ui.NewWidget()
	vbox := ui.NewVBoxLayout()
	hbox := ui.NewHBoxLayout()
	my := new(MyWidget).Init()
	lbl := ui.NewLabel()
	lbl.SetText("this is custome widget - draw lines")
	btn := ui.NewButton()
	btn.SetText("Clear")
	btn.OnClicked(func() {
		my.Clear()
	})
	hbox.AddWidget(lbl)
	hbox.AddWidgetWith(btn, 0, ui.AlignRight)
	vbox.AddLayout(hbox)
	vbox.AddWidgetWith(my, 1, 0)
	w.SetLayout(vbox)
	return w
}

func main() {
	ui.Main(ui_main)
}

func ui_main() {
	fmt.Println("vfc/ui")
	ui.OnInsertObject(func(v interface{}) {
		fmt.Println("add item", v)
	})
	ui.OnRemoveObject(func(v interface{}) {
		fmt.Println("remove item", v)
	})
	w := new(MainWindow).Init()
	defer w.Close()

	w.SetSizev(400, 300)
	w.OnCloseEvent(func(e *ui.CloseEvent) {
		fmt.Println("close", e)
	})
	w.Show()
	ui.Run()
}

type MyWidget struct {
	ui.Widget
	lines [][]ui.Point
	line  []ui.Point
	font  *ui.Font
}

func (p *MyWidget) Name() string {
	return "MyWidget"
}

func (p *MyWidget) String() string {
	return ui.DumpObject(p)
}

func (p *MyWidget) Init() *MyWidget {
	if p.Widget.Init() == nil {
		return nil
	}
	p.font = ui.NewFontWith("Timer", 16, 87)
	p.Widget.OnPaintEvent(func(e *ui.PaintEvent) {
		p.paintEvent(e)
	})
	p.Widget.OnMousePressEvent(func(e *ui.MouseEvent) {
		p.mousePressEvent(e)
	})
	p.Widget.OnMouseMoveEvent(func(e *ui.MouseEvent) {
		p.mouseMoveEvent(e)
	})
	p.Widget.OnMouseReleaseEvent(func(e *ui.MouseEvent) {
		p.mouseReleaseEvent(e)
	})
	ui.InsertObject(p)
	return p
}

func (p *MyWidget) Clear() {
	p.lines = [][]ui.Point{}
	p.Update()
}

func (p *MyWidget) paintEvent(e *ui.PaintEvent) {
	paint := ui.NewPainter()
	defer paint.Close()

	paint.Begin(p)
	paint.SetFont(p.font)
	paint.DrawLines(p.line)
	paint.SetFont(p.font)
	paint.DrawText(ui.Pt(100, 100), "this is a test")
	for _, v := range p.lines {
		//paint.DrawLines(v)
		paint.DrawPolyline(v)
	}
	paint.End()
	runtime.GC()
}

func (p *MyWidget) mousePressEvent(e *ui.MouseEvent) {
	p.line = append(p.line, e.Pos())
	p.Update()
}

func (p *MyWidget) mouseMoveEvent(e *ui.MouseEvent) {
	p.line = append(p.line, e.Pos())
	p.Update()
}

func (p *MyWidget) mouseReleaseEvent(e *ui.MouseEvent) {
	p.line = append(p.line, e.Pos())
	p.lines = append(p.lines, p.line)
	p.line = []ui.Point{}
	p.Update()
}
