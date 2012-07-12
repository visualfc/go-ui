module("scrollbar")

name = "ScrollBar"
base = "baseSlider"

funcs = [[
+ Init()
]]

qtdrv = {
inc = "<QScrollBar>",	
name = "QScrollBar *",

Init = [[
drvNewObj(a0,new QScrollBar(Qt::Horizontal));
]],

}