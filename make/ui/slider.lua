module("slider")

name = "Slider"
base = "baseSlider"

funcs = [[
+ Init()
@ SetTickInterval(value int)
@ TickInterval() (value int)
@ SetTickPosition(value TickPosition)
@ TickPosition() (value TickPosition)
]]

qtdrv = {
inc = "<QSlider>",	
name = "QSlider *",

Init = [[
drvNewObj(a0,new QSlider(Qt::Horizontal));
]],

SetTickInterval = "setTickInterval",
TickInterval = "tickInterval",
SetTickPosition = "setTickPosition",
TickPosition = "tickPosition",
}