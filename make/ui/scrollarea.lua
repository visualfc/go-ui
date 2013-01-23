module("scrollarea")

name = "ScrollArea"
base = "baseScrollArea"

funcs = [[
+ Init()
@ SetAlignment(a Alignment)
@ Alignment() (a Alignment)
@ SetWidget(w IWidget)
@ Widget() (w IWidget)
@ SetWidgetResizable(b bool)
@ WidgetResizable() (b bool)
@ TakeWidget() (w IWidget)
EnsureVisible(x int, y int, xmargin int, ymargin int)
EnsureWidgetVisible(w IWidget, xmargin int, ymargin int)
]]

qtdrv = {
inc = "<QScrollArea>",	
name = "QScrollArea *",

Init = [[
drvNewObj(a0,new QScrollArea());
]],

SetAlignment = "setAlignment",
Alignment = "alignment",
SetWidget = "setWidget",
Widget = "widget",
SetWidgetResizable = "setWidgetResizable",
WidgetResizable = "widgetResizable",
TakeWidget = "takeWidget",
EnsureVisible = "ensureVisible",
EnsureWidgetVisible = "ensureWidgetVisible",
}
