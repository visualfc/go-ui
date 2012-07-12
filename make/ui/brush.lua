module("brush")

name = "Brush"
base = "object"

funcs = [[
+ Init()
- Close()
@ SetColor(clr color.Color)
@ Color() (clr color.Color)
@ SetStyle(style BrushStyle)
@ Style() (style BrushStyle)
]]

qtdrv = {
inc = "<QBrush>",
name = "QBrush*",

Init = [[
drvSetBrush(a0, QBrush());
]],
Close = [[
drvDelBrush(a0,self);
]],
SetColor = "setColor",
Color = "color",
SetStyle = [[
self->setStyle((Qt::BrushStyle)drvGetInt(a1));
]],
Style = [[
drvSetInt(a1,self->style());
]],

}