module("frame")

name = "Frame"
base = "Widget"

funcs = [[
+ Init()
@ SetFrameStyle(style int)
@ FrameStyle() (style int)
@ SetFrameRect(rc Rect)
@ FrameRect()(rc Rect)
]]

expands = [[
func (p *Frame) SetFrameRectv(x,y,width,height int) {
	p.SetFrameRect(Rc(x,y,width,height))
}

func (p *Frame) FrameRectv() (int,int,int,int) {
	return p.FrameRect().Unpack()
}
]]


qtdrv = {
inc = "<QFrame>",
name = "QFrame *",

Init = [[
drvNewObj(a0,new QFrame);
]],

SetFrameStyle = "setFrameStyle",
FrameStyle = "frameStyle",
SetFrameRect = "setFrameRect",
FrameRect = "frameRect",

}
