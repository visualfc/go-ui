module("systemtray")

name = "SystemTray"
base = "object"

funcs = [[
+ Init()
- Close()
@ SetContextMenu(menu *Menu)
@ ContextMenu() (menu *Menu)
@ SetIcon(icon *Icon)
@ Icon() (icon *Icon)
@ SetToolTip(text string)
@ ToolTip() (text string)
@ SetVisible(b bool)
@ IsVisible() (b bool)
ShowMessage(title string, message string, icontype MessageIconType, mstimeouthint int)
* OnActivated(fn func(int))
* OnMessageClicked(fn func())
]]

qtdrv = {
inc = "<QSystemTrayIcon>",
name = "QSystemTrayIcon*",

Init = [[
drvNewObj(a0, new QSystemTrayIcon);
]],

Close = [[
drvDelObj(a0,self);
]],

SetContextMenu = "setContextMenu",
ContextMenu = "contextMenu",
SetIcon = "setIcon",
Icon = "icon",
SetToolTip = "setToolTip",
ToolTip = "toolTip",
SetVisible = "setVisible",
IsVisible = "isVisible",
ShowMessage = "showMessage",

OnActivated = [[
QObject::connect(self,SIGNAL(activated(int)),drvNewSignal(self,a1,a2),SLOT(call(int)));
]],
OnMessageClicked = [[
QObject::connect(self,SIGNAL(messageClicked(int)),drvNewSignal(self,a1,a2),SLOT(call()));
]],

}
