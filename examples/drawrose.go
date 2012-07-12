/*

http://www.romancortes.com/blog/1k-rose/

with(m = Math) C = cos,
S = sin,
P = pow,
R = random;
c.width = c.height = f = 500;
h = -250;
function p(a, b, c) {
    if (c > 60) return [S(a * 7) * (13 + 5 / (.2 + P(b * 4, 4))) - S(b) * 50, b * f + 50, 625 + C(a * 7) * (13 + 5 / (.2 + P(b * 4, 4))) + b * 400, a * 1 - b / 2, a];
    A = a * 2 - 1;
    B = b * 2 - 1;
    if (A * A + B * B < 1) {
        if (c > 37) {
            n = (j = c & 1) ? 6 : 4;
            o = .5 / (a + .01) + C(b * 125) * 3 - a * 300;
            w = b * h;
            return [o * C(n) + w * S(n) + j * 610 - 390, o * S(n) - w * C(n) + 550 - j * 350, 1180 + C(B + A) * 99 - j * 300, .4 - a * .1 + P(1 - B * B, -h * 6) * .15 - a * b * .4 + C(a + b) / 5 + P(C((o * (a + 1) + (B > 0 ? w: -w)) / 25), 30) * .1 * (1 - B * B), o / 1e3 + .7 - o * w * 3e-6]
        }
        if (c > 32) {
            c = c * 1.16 - .15;
            o = a * 45 - 20;
            w = b * b * h;
            z = o * S(c) + w * C(c) + 620;
            return [o * C(c) - w * S(c), 28 + C(B * .5) * 99 - b * b * b * 60 - z / 2 - h, z, (b * b * .3 + P((1 - (A * A)), 7) * .15 + .3) * b, b * .7]
        }
        o = A * (2 - b) * (80 - c * 2);
        w = 99 - C(A) * 120 - C(b) * ( - h - c * 4.9) + C(P(1 - b, 7)) * 50 + c * 2;
        z = o * S(c) + w * C(c) + 700;
        return [o * C(c) - w * S(c), B * 99 - C(P(b, 7)) * 50 - c / 3 - z / 1.35 + 450, z, (1 - b / 1.2) * .9 + a * .1, P((1 - b), 20) / 4 + .05]
    }
}
setInterval('for(i=0;i<1e4;i++)if(s=p(R(),R(),i%46/.74)){z=s[2];x=~~(s[0]*f/z-h);y=~~(s[1]*f/z-h);if(!m[q=y*f+x]|m[q]>z)m[q]=z,a.fillStyle="rgb("+~(s[3]*h)+","+~(s[4]*h)+","+~(s[3]*s[3]*-80)+")",a.fillRect(x,y,1,1)}',0)
*/
package main

import (
	"fmt"
	"github.com/visualfc/go-ui/ui"
	"math"
	"math/rand"
)

type rect struct {
	width, height int
}

var (
	S           = math.Sin
	C           = math.Cos
	P           = math.Pow
	R           = rand.Float64
	f   float64 = 500
	h   float64 = -250
	m           = make(map[float64]float64)
	max         = 1024
)

func main() {
	ui.Main(ui_main)
}

func ui_main() {
	w := ui.NewWidget()
	defer w.Close()
	w.SetWindowTitle("Rose")

	rose := ui.NewWidget()
	rose.SetMinimumSize(ui.Sz(510, 510))

	img := ui.NewImageWithSize(500, 500)
	defer img.Close()

	imgPainter := ui.NewPainterWithImage(img)
	imgPainter.InitFrom(rose)
	defer imgPainter.Close()

	var timerid int = -1
	var count int
	var timerValue int = 50

	rose.OnTimerEvent(func(e *ui.TimerEvent) {
		if e.TimerId() == timerid {
			draw(imgPainter)
			rose.Update()
			count++
			if count >= max {
				w.KillTimer(timerid)
			}
			w.SetWindowTitle(fmt.Sprintf("Rose: %d*1e4", count))
		}
	})
	rose.OnPaintEvent(func(e *ui.PaintEvent) {
		painter := ui.NewPainter()
		defer painter.Close()
		painter.Begin(rose)
		painter.DrawImageEx(ui.Point{0, 0}, img, img.Rect())
		painter.End()
	})

	load := func() {
		if timerid != -1 {
			rose.KillTimer(timerid)
		}
		for k, _ := range m {
			delete(m, k)
		}
		img.Fill(rgb(255, 255, 255))
		count = 0
		timerid = rose.StartTimer(timerValue)
	}
	stop := func() {
		rose.KillTimer(timerid)
		timerid = -1
	}

	vbox := ui.NewVBoxLayout()

	hbox := ui.NewHBoxLayout()
	vbox.AddWidget(rose)
	vbox.AddLayout(hbox)

	loadBtn := ui.NewButtonWithText("Reload")
	stopBtn := ui.NewButtonWithText("Stop")
	exitBtn := ui.NewButtonWithText("Exit")
	timerLabel := ui.NewLabelWithText("Timer: 50 ")
	timerSlider := ui.NewSlider()
	timerSlider.SetRange(1, 600)
	timerSlider.SetValue(50)
	timerSlider.OnValueChanged(func(v int) {
		timerValue = v
		timerLabel.SetText(fmt.Sprintf("Timer:%4d ", v))
	})
	maxLabel := ui.NewLabelWithText("Max: 1024 ")
	maxSlider := ui.NewSlider()
	maxSlider.SetRange(10, 10240)
	maxSlider.SetValue(1024)
	maxSlider.OnValueChanged(func(v int) {
		max = v
		maxLabel.SetText(fmt.Sprintf("Max:%5d ", v))
	})
	hbox.AddWidget(maxLabel)
	hbox.AddWidget(maxSlider)
	hbox.AddSpacing(10)
	hbox.AddWidget(timerLabel)
	hbox.AddWidget(timerSlider)
	hbox.AddSpacing(10)
	hbox.AddWidget(loadBtn)
	hbox.AddWidget(stopBtn)
	hbox.AddStretch(0)
	hbox.AddWidget(exitBtn)

	loadBtn.OnClicked(load)
	stopBtn.OnClicked(stop)
	exitBtn.OnClicked(func() {
		w.Close()
	})

	w.SetLayout(vbox)
	w.Show()

	w.OnCloseEvent(func(e *ui.CloseEvent) {
		rose.KillTimer(timerid)
	})

	load()

	ui.Run()
}

func sel(b bool, v1, v2 float64) float64 {
	if b {
		return v1
	}
	return v2
}
func p(a, b, c float64) []float64 {
	if c > 60 {
		return []float64{S(a*7)*(13+5/(.2+P(b*4, 4))) - S(b)*50, b*f + 50, 625 + C(a*7)*(13+5/(.2+P(b*4, 4))) + b*400, a*1 - b/2, a}
	}
	A := a*2 - 1
	B := b*2 - 1
	if A*A+B*B < 1 {
		if c > 37 {
			var j float64 = float64(int64(c) & 1)
			var n float64 = sel(j > 0, 6, 4)
			var o float64 = .5/(a+.01) + C(b*125)*3 - a*300
			var w float64 = b * h
			return []float64{o*C(n) + w*S(n) + j*610 - 390, o*S(n) - w*C(n) + 550 - j*350, 1180 + C(B+A)*99 - j*300, .4 - a*.1 + P(1-B*B, -h*6)*.15 - a*b*.4 + C(a+b)/5 + P(C((o*(a+1)+sel(B > 0, w, -w))/25), 30)*.1*(1-B*B), o/1e3 + .7 - o*w*3e-6}
		}
		if c > 32 {
			c = c*1.16 - .15
			var o float64 = a*45 - 20
			var w float64 = b * b * h
			var z float64 = o*S(c) + w*C(c) + 620
			return []float64{o*C(c) - w*S(c), 28 + C(B*.5)*99 - b*b*b*60 - z/2 - h, z, (b*b*.3 + P((1-(A*A)), 7)*.15 + .3) * b, b * .7}
		}
		var o float64 = A * (2 - b) * (80 - c*2)
		var w float64 = 99 - C(A)*120 - C(b)*(-h-c*4.9) + C(P(1-b, 7))*50 + c*2
		var z float64 = o*S(c) + w*C(c) + 700
		return []float64{o*C(c) - w*S(c), B*99 - C(P(b, 7))*50 - c/3 - z/1.35 + 450, z, (1-b/1.2)*.9 + a*.1, P((1-b), 20)/4 + .05}
	}
	return nil
}

func rgb(r, g, b int64) uint {
	if r < 0 {
		r = -r
	}
	if g < 0 {
		g = -g
	}
	if b < 0 {
		b = -b
	}
	return uint(r%256*256*256 + g%256*256 + b)
}

func draw(painter *ui.Painter) {
	for i := 0; i < 1e4; i++ {
		s := p(R(), R(), float64(i%46)/.74)
		if s != nil {
			z := s[2]
			x := float64(^^int(s[0]*f/z - h))
			y := float64(^^int(s[1]*f/z - h))

			q := y*f + x
			if v, ok := m[q]; !ok || (v > z) {
				m[q] = z
				clr := rgb(^int64(s[3]*h), ^int64(s[4]*h), ^int64(s[3]*s[3]*-80))
				painter.FillRectF(ui.RcF(x, y, 1, 1), uint(clr))
			}
		}
	}
}
