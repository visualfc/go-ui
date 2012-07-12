module("toolbox")

name = "ToolBox"
base = "Widget"

funcs = [[
+ Init()
@ SetCurrentIndex(index int)
@ SetCurrentWidget(widget IWidget)
@ CurrentIndex() (n int)
@ CurrentWidget() (widget IWidget)
@ Count() (n int)
AddItem(widget IWidget, text string, icon *Icon)
InsertItem(index int, widget IWidget, text string, icon *Icon)
RemoveItem(index int)
IndexOf(widget IWidget) (index int)
WidgetOf(index int) (widget IWidget)
SetItemText(index int, text string)
ItemText(index int) (text string)
SetItemIcon(index int, icon *Icon)
ItemIcon(index int) (icon *Icon)
SetItemToolTip(index int, text string)
ItemToolTip(index int) (text string)
IsItemEnabled(index int) (b bool)

* OnCurrentChanged(fn func(int))
]]

qtdrv = {
inc = "<QToolBox>",
name = "QToolBox*",

Init = [[
drvNewObj(a0, new QToolBox);
]],
SetCurrentIndex = "setCurrentIndex",
SetCurrentWidget = "setCurrentWidget",
CurrentIndex = "currentIndex",
CurrentWidget = "currentWidget",
Count = "count",
AddItem = [[
self->addItem(drvGetWidget(a1),drvGetIcon(a3),drvGetString(a2));
]],
InsertItem = [[
self->insertItem(drvGetInt(a1),drvGetWidget(a2),drvGetIcon(a4),drvGetString(a3));
]],
RemoveItem = "removeItem",
IndexOf = "indexOf",
WidgetOf = "widget",
SetItemText = "setItemText",
ItemText = "itemText",
SetItemIcon = "setItemIcon",
ItemIcon = "itemIcon",
SetItemToolTip = "setItemToolTip",
ItemToolTip = "itemToolTip",
IsItemEnabled = "isItemEnabled",

OnCurrentChanged = [[
QObject::connect(self,SIGNAL(currentChanged(int)),drvNewSignal(self,a1,a2),SLOT(call(int)));
]]

}
