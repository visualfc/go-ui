module("toolbar")

name = "ToolBar"
base = "Widget"

funcs = [[
+ Init()
+ InitWithTitle(text string)
@ SetTitle(text string)
@ Title() (text string)
@ SetIconSize(sz Size)
@ IconSize() (sz Size)
@ SetFloatable(b bool)
@ IsFloatable() (b bool)
@ IsFloating() (b bool)
@ SetToolButtonStyle(style ToolButtonStyle)
@ ToolButtonStyle() (style ToolButtonStyle)
AddAction(act *Action)
AddSeparator() (act *Action)
AddWidget(widget IWidget) (act *Action)
Clear()
]]

qtdrv = {
inc = "<QToolBar>",
name = "QToolBar*",

Init = [[
drvNewObj(a0, new QToolBar,true,false);
]],
InitWithTitle = [[
drvNewObj(a0, new QToolBar(drvGetString(a1)),true,false);
]],

SetTitle = "setWindowTitle",
Title = "windowTitle",
SetIconSize = "setIconSize",
IconSize = "iconSize",
SetFloatable = "setFloatable",
IsFloatable = "isFloatable",
IsFloating = "isFloating",
AddAction = "addAction",
AddSeparator = [[
drvNewObj(a1,self->addSeparator());
]],
AddWidget = [[
drvNewObj(a2,self->addWidget(drvGetWidget(a1)));
]],
Clear = "clear",

}
