package main

import (
	"github.com/visualfc/go-ui/ui"
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	ui.Main(main_ui)
}

func main_ui() {
	ui.OnInsertObject(func(v interface{}) {
		fmt.Println("add item", v)
	})
	ui.OnRemoveObject(func(v interface{}) {
		fmt.Println("remove item", v)
	})
	w := ui.NewMainWindow()
	defer w.Close()
	go func() {
		dock := ui.NewDockWidgetWithTitle("Dock")
		dock.SetDock(ui.NewButtonWithText("Hello"))
		w.AddDockWidget(ui.LeftDockWidgetArea, dock)
		btn := ui.NewButtonWithText("CloseDock测试")
		w.SetCentralWidget(btn)
		w.SetSize(ui.Sz(200, 200))

		tb := ui.NewToolBarWithTitle("Standard")
		tb.AddWidget(ui.NewButtonWithText("ok"))
		w.AddToolBar(tb)

		tb.OnCloseEvent(func(e *ui.CloseEvent) {
			fmt.Println("tb close", e)
		})
		sb := ui.NewStatusBar()
		w.SetStatusBar(sb)
		sb.OnCloseEvent(func(e *ui.CloseEvent) {
			fmt.Println("sb close", e)
		})

		btn.OnClicked(func() {
			//dock.Close()
			runtime.GC()
			btn.SetText(btn.Text())
		})
		dock.OnCloseEvent(func(e *ui.CloseEvent) {
			fmt.Println(e)
		})

		go func() {
			for {
				timer := time.NewTimer(1)
				select {
				case <-timer.C:
					btn.SetText(btn.Text())
					btn.SetText(btn.Text())
					btn.SetText(btn.Text())
					fmt.Println(">", btn.Text())
					if btn.Text() != "CloseDock测试" {
						panic("close")
					}
				}
			}
		}()

		dock.OnVisibilityChanged(func(b bool) {
			fmt.Println(b)
			if !b {
				time.AfterFunc(1e9, func() {
					if dock.IsValid() {
						dock.Show()
					}
				})
			}
		})

		w.Show()
	}()

	ui.Run()
}
