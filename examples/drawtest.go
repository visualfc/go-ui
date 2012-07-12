package main

import (
	"fmt"
	"github.com/visualfc/go-ui/ui"
	"image/color"
)

type rgba uint32

func (c rgba) RGBA() (r, g, b, a uint32) {
	return uint32((c >> 16) & 0xff), uint32((c >> 8) & 0xff), uint32(c & 0xff), uint32(c >> 24)
}

func make_rgba(c color.Color) rgba {
	if c == nil {
		return 0
	}
	r, g, b, a := c.RGBA()
	return rgba(((a & 0xff) << 24) | ((r & 0xff) << 16) | ((g & 0xff) << 8) | (b & 0xff))
}

func main() {
	ui.Main(func() {
		w := ui.NewWidget()

		w.OnPaintEvent(func(e *ui.PaintEvent) {
			paint := ui.NewPainter()
			defer paint.Close()
			paint.Begin(w)
			pen := ui.NewPen()
			pen.SetColor(color.RGBA{255, 128, 0, 0})
			pen.SetWidth(2)
			fmt.Println(pen, pen.Color(), pen.Width())
			paint.SetPen(pen)
			brush := ui.NewBrush()
			brush.SetStyle(ui.SolidPattern)
			brush.SetColor(color.RGBA{128, 128, 0, 255})
			paint.SetBrush(brush)
			paint.DrawRect(ui.Rect{10, 10, 100, 100})
		})

		w.SetSize(ui.Size{400, 400})
		w.Show()

		ui.Run()
	})
}
