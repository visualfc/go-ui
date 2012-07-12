package main

import (
	"fmt"
	"github.com/visualfc/go-ui/ui"
)

func main() {
	ui.Main(ui_main)
}

func ui_main() {
	w := ui.NewWidget()
	defer w.Close()
	list := ui.NewListWidget()
	vbox := ui.NewVBoxLayout()
	vbox.AddWidget(list)
	w.SetLayout(vbox)
	go func() {
		list.OnCurrentItemChanged(func(item, old *ui.ListWidgetItem) {
			go func() {
				fmt.Println(item.Attr("text"))
			}()
		})

		item := ui.NewListWidgetItem()
		item.SetText("Item1")
		list.AddItem(item)
		list.AddItem(ui.NewListWidgetItemWithText("Item2"))
	}()

	w.Show()

	ui.Run()
}
