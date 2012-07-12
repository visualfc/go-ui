module("statusbar")

name = "StatusBar"
base = "Widget"

funcs = [[
+ Init()
@ SetSizeGrip(b bool)
@ IsSizeGrip() (b bool)
AddWidget(widget IWidget, stretch int)
AddPermanentWidget(widget IWidget, stretch int)
InsertWidget(index int,widget IWidget, stretch int)
InsertPermanentWidget(index int, widget IWidget, stretch int)
RemoveWidget(widget IWidget)
ShowMessage(text string, timeout int)
CurrentMessage() (text string)
ClearMessage()
]]

qtdrv = {
inc = "<QStatusBar>",
name = "QStatusBar*",

Init = [[
drvNewObj(a0, new QStatusBar());
]],

SetSizeGrip = "setSizeGripEnabled",
IsSizeGrip = "isSizeGripEnabled",
AddWidget = "addWidget",
AddPermanentWidget = "addPermanentWidget",
InsertWidget = "insertWidget",
InsertPermanentWidget = "insertPermanentWidget",
RemoveWidget = "removeWidget",
ShowMessage = "showMessage",
CurrentMessage = "currentMessage",
ClearMessage = "clearMessage",

}
