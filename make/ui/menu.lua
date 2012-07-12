module("menu")

name = "Menu"
base = "Widget"

funcs = [[
+ Init()
+ InitWithTitle(text string)
@ SetTitle(text string)
@ Title() (text string)
@ SetIcon(icon *Icon)
@ Icon() (icon *Icon)
@ SetDefaultAction(act *Action)
@ DefaultAction() (act *Action)
@ SetActiveAction(act *Action)
@ ActiveAction() (act *Action)
@ MenuAction() (act *Action)
@ IsEmpty() (b bool)
AddMenu(menu *Menu) (act *Action)
InsertMenu(before *Action, menu *Menu)
AddSeparator() (act *Action)
InsertSeparator(before *Action) (act *Action)
Clear()
Popup(pt Point, act *Action)

* OnHovered(fn func(*Action))
* OnTriggered(fn func(*Action))
]]


qtdrv = {
inc = "<QMenu>",
name = "QMenu *",

Init = [[
drvNewObj(a0,new QMenu);
]],
InitWithTitle = [[
drvNewObj(a0, new QMenu(drvGetString(a1)));
]],

SetTitle = "setTitle",
Title = "title",
SetIcon = "setIcon",
Icon = "icon",
SetDefaultAction = "setDefaultAction",
DefaultAction = "defaultAction",
SetActiveAction = "setActiveAction",
ActiveAction = "activeAction",
MenuAction = "menuAction",
IsEmpty = "isEmpty",
AddMenu = "addMenu",
InsertMenu = "insertMenu",
AddSeparator = "addSeparator",
InsertSeparator = "insertSeparator",
Clear = "clear",
Popup = "popup",

OnHovered = [[
QObject::connect(self,SIGNAL(hovered(QAction*)),drvNewSignal(self,a1,a2),SLOT(call(QAction*)));
]],

OnTriggered = [[
QObject::connect(self,SIGNAL(triggered(QAction*)),drvNewSignal(self,a1,a2),SLOT(call(QAction*)));
]],
}
