module("tabwidget")

name = "TabWidget"
base = "Widget"

funcs = [[
+Init()
@ Count() (n int)
@ CurrentIndex() (index int)
@ CurrentWidget() (w IWidget)
@ SetCurrentIndex(index int)
@ SetCurrentWidget(w IWidget)
AddTab(w IWidget, label string, icon *Icon)
InsertTab(index int, w IWidget, label string, icon *Icon)
RemoveTab(index int)
Clear()
IndexOf(w IWidget) (index int)
WidgetOf(index int) (w IWidget)
SetTabText(index int, label string)
TabText(index int) (label string)
SetTabIcon(index int, icon *Icon)
TabIcon(index int) (icon *Icon)
SetTabToolTip(index int, tip string)
TabToolTip(index int) (tip string)
* OnCurrentChanged(fn func(int))
]]


qtdrv = {
inc = "<QTabWidget>",
name = "QTabWidget *",

Init = [[
drvNewObj(a0,new QTabWidget);
]],

Count = "count",
CurrentIndex = "currentIndex",
CurrentWidget = "currentWidget",
SetCurrentIndex = "setCurrentIndex",
SetCurrentWidget = "setCurrentWidget",
IndexOf = "indexOf",
WidgetOf = "widget",
AddTab = [[
self->addTab(drvGetWidget(a1),drvGetIcon(a3),drvGetString(a2));
]],
InsertTab = [[
self->insertTab(drvGetInt(a1),drvGetWidget(a2),drvGetIcon(a4),drvGetString(a3));
]],
RemoveTab = "removeTab",
Clear = "clear",
SetTabText = "setTabText",
TabText = "tabText",
SetTabIcon = "setTabIcon",
TabIcon = "tabIcon",
SetTabToolTip = "setTabToolTip",
TabToolTip = "tabToolTip",



OnCurrentChanged = [[
QObject::connect(self,SIGNAL(currentChanged(int)),drvNewSignal(self,a1,a2),SLOT(call(int)));
]],

}
