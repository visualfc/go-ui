module("dial")

name = "Dial"
base = "baseSlider"

funcs = [[
+ Init()
@ NotchSize() (size int)
@ SetNotchTarget(value float64)
@ NotchTarget() (value float64)
@ SetNotchesVisible(b bool)
@ NotchesVisible() (b bool)
@ SetWrapping(b bool)
@ Wrapping() (b bool)
]]

qtdrv = {
inc = "<QDial>",	
name = "QDial *",

Init = [[
drvNewObj(a0,new QDial);
]],

NotchSize = "notchSize",
SetNotchTarget = "setNotchTarget",
NotchTarget = "notchTarget",
SetNotchesVisible = "setNotchesVisible",
NotchesVisible = "notchesVisible",
SetWrapping = "setWrapping",
Wrapping = "wrapping",

}