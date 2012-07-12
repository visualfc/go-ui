module("menubar")

name = "MenuBar"
base = "Widget"

funcs = [[
+ Init()
@ SetActiveAction(act *Action)
@ ActiveAction() (act *Action)
@ SetNativeMenuBar(b bool)
@ IsNativeMenuBar() (b bool)
AddMenu(menu *Menu) (act *Action)
InsertMenu(before *Action, menu *Menu)
AddSeparator() (act *Action)
InsertSeparator(before *Action) (act *Action)
Clear()

* OnHovered(fn func(*Action))
* OnTriggered(fn func(*Action))
]]


qtdrv = {
inc = "<QMenuBar>",
name = "QMenuBar *",

Init = [[
drvNewObj(a0,new QMenuBar);
]],

SetActiveAction = "setActiveAction",
ActiveAction = "activeAction",
SetNativeMenuBar = "setNativeMenuBar",
IsNativeMenuBar = "isNativeMenuBar",
AddMenu = "addMenu",
InsertMenu = "insertMenu",
AddSeparator = "addSeparator",
InsertSeparator = "insertSeparator",
Clear = "clear",

OnHovered = [[
QObject::connect(self,SIGNAL(hovered(QAction*)),drvNewSignal(self,a1,a2),SLOT(call(QAction*)));
]],

OnTriggered = [[
QObject::connect(self,SIGNAL(triggered(QAction*)),drvNewSignal(self,a1,a2),SLOT(call(QAction*)));
]],
}
