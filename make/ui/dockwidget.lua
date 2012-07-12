module("dockwidget")

name = "DockWidget"
base = "Widget"

funcs = [[
+ Init()
+ InitWithTitle(title string)
@ SetDock(widget IWidget)
@ Dock() (widget IWidget)
@ SetTitleBar(widget IWidget)
@ TitleBar() (widget IWidget)
@ SetFloating(b bool)
@ IsFloating() (b bool)
* OnVisibilityChanged(fn func(bool))
]]

qtdrv = {
inc = "<QDockWidget>",
name = "QDockWidget *",

Init = [[
drvNewObj(a0,new QDockWidget,true,false);
]],
InitWithTitle = [[
drvNewObj(a0,new QDockWidget(drvGetString(a1)),true,false);
]],

SetDock = "setWidget",
Dock = "widget",
SetTitleBar = "setTitleBarWidget",
TitleBar = "titleBarWidget",
SetFloating = "setFloating",
IsFloating = "isFloating",

OnVisibilityChanged = [[
QObject::connect(self,SIGNAL(visibilityChanged(bool)),drvNewSignal(self,a1,a2),SLOT(call(bool)));
]]
}
