module("mainwindow")

name = "MainWindow"
base = "Widget"

funcs = [[
+ Init()
@ SetCentralWidget(widget IWidget)
@ CentralWidget() (widget IWidget)
@ SetMenuBar(bar *MenuBar)
@ MenuBar() (bar *MenuBar)
@ SetStatusBar(bar *StatusBar)
@ StatusBar() (bar *StatusBar)
AddToolBar(bar *ToolBar)
InsertToolBar(before *ToolBar, bar *ToolBar)
RemoveToolBar(bar *ToolBar)
AddDockWidget(area DockWidgetArea, dock *DockWidget)
RemoveDockWidget(dock *DockWidget)
]]

qtdrv = {
inc = "<QMainWindow>",
name = "QMainWindow *",

Init = [[
drvNewObj(a0,new QMainWindow);
]],

SetCentralWidget = "setCentralWidget",
CentralWidget = "centralWidget",
SetMenuBar = "setMenuBar",
MenuBar = "menuBar",
SetStatusBar = "setStatusBar",
StatusBar = "statusBar",
AddToolBar = "addToolBar",
InsertToolBar = "insertToolBar",
RemoveToolBar = "removeToolBar",
AddDockWidget = [[
self->addDockWidget((Qt::DockWidgetArea)drvGetInt(a1),drvGetDockWidget(a2));
]],
RemoveDockWidget = "removeDockWidget",
}
