module("actiongroup")

name = "ActionGroup"
base = "object"

funcs = [[
+ Init()
- Close()
@ SetVisible(b bool)
@ IsVisible() (b bool)
@ SetEnabled(b bool)
@ IsEnabled() (b bool)
@ SetExclusive(b bool)
@ IsExclusive() (b bool)
AddAction(act *Action)
RemoveAction(act *Action)
* OnHovered(fn func(*Action))
* OnTriggered(fn func(*Action))
]]

qtdrv = {
inc = "<QActionGroup>",
name = "QActionGroup*",

Init = [[
drvNewObj(a0, new QActionGroup(qApp));
]],
Close = [[
drvDelObj(a0,self);
]],

OnHovered = [[
QObject::connect(self,SIGNAL(hovered(QAction*)),drvNewSignal(self,a1,a2),SLOT(call(QAction*)));
]],
OnTriggered = [[
QObject::connect(self,SIGNAL(triggered(QAction*)),drvNewSignal(self,a1,a2),SLOT(call(QAction*)));
]]

}
