module("pen")

name = "Pen"
base = "object"

funcs = [[
+ Init()
- Close()
@ SetColor(clr color.Color)
@ Color() (clr color.Color)
@ SetWidth(width int)
@ Width() (width int)
@ SetStyle(style PenStyle)
@ Style() (style PenStyle)
]]

qtdrv = {
inc = "<QPen>",
name = "QPen*",

Init = [[
drvSetPen(a0, QPen());
]],
Close = [[
drvDelPen(a0,self);
]],
SetColor = "setColor",
Color = "color",
SetWidth = "setWidth",
Width = "width",
SetStyle = [[
self->setStyle((Qt::PenStyle)drvGetInt(a1));
]],
Style = [[
drvSetInt(a1,self->style());
]],

}