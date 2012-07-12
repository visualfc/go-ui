module("toolbutton")

name = "ToolButton"
base = "baseButton"

funcs = [[
+ Init()
+ InitWithText(text string)
@ SetMenu(menu *Menu)
@ Menu()(menu *Menu)
@ SetAutoRaise(b bool)
@ AutoRaise() (b bool)
@ SetToolButtonStyle(style ToolButtonStyle)
@ ToolButtonStyle() (style ToolButtonStyle)
@ SetToolButtonPopupMode(mode ToolButtonPopupMode)
@ ToolButtonPopupMode() (mode ToolButtonPopupMode)
* OnClicked(fn func())
]]


qtdrv = {
inc = "<QToolButton>",
name = "QToolButton *",

Init = [[
drvNewObj(a0,new QToolButton);
]],

InitWithText = [[
drvNewObj(a0,new QPushButton(drvGetString(a1)));
]],

SetMenu = "setMenu",
Menu = "menu",
SetAutoRaise = "setAutoRaise",
AutoRaise = "autoRaise",
SetToolButtonStyle = "setToolButtonStyle",
ToolButtonStyle = "toolButtonStyle",
SetToolButtonPopupMode = "setPopupMode",
ToolButtonPopupMode = "popupMode",

OnClicked = [[
QObject::connect(self,SIGNAL(clicked()),drvNewSignal(self,a1,a2),SLOT(call()));
]],

}
