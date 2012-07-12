module("listwidget")

name = "ListWidget"
base = "Widget"

funcs = [[
+ Init()
@ Count() (count int)
@ SetCurrentItem(item *ListWidgetItem)
@ CurrentItem() (item *ListWidgetItem)
@ SetCurrentRow(row int)
@ CurrentRow() (row int)
AddItem(item *ListWidgetItem)
InsertItem(index int,item *ListWidgetItem)
EditItem(item *ListWidgetItem)
TakeItem(row int) (item *ListWidgetItem)
Item(row int) (item *ListWidgetItem)
Clear()
* OnCurrentItemChanged(fn func(*ListWidgetItem,*ListWidgetItem))
* OnCurrentRowChanged(fn func(int))
* OnItemActivated(fn func(*ListWidgetItem))
* OnItemChanged(fn func(*ListWidgetItem))
* OnItemClicked(fn func(*ListWidgetItem))
* OnItemDoubleClicked(fn func(*ListWidgetItem))
* OnItemEntered(fn func(*ListWidgetItem))
* OnItemPressed(fn func(*ListWidgetItem))
* OnItemSelectionChanged(fn func())
]]

qtdrv = {
inc = "<QListWidget>",
name = "QListWidget *",

Init = [[
drvNewObj(a0,new QListWidget);
]],
OnCurrentItemChanged = [[
QObject::connect(self,SIGNAL(currentItemChanged(QListWidgetItem*,QListWidgetItem*)),drvNewSignal(self,a1,a2),SLOT(call(QListWidgetItem*,QListWidgetItem*)));
]],
OnCurrentRowChanged = [[
QObject::connect(self,SIGNAL(currentRowChanged(int)),drvNewSignal(self,a1,a2),SLOT(call(int)));
]],
OnItemActivated = [[
QObject::connect(self,SIGNAL(itemActivated(QListWidgetItem*)),drvNewSignal(self,a1,a2),SLOT(call(QListWidgetItem*)));
]],
OnItemChanged = [[
QObject::connect(self,SIGNAL(itemChanged(QListWidgetItem*)),drvNewSignal(self,a1,a2),SLOT(call(QListWidgetItem*)));
]],
OnItemClicked = [[
QObject::connect(self,SIGNAL(itemClicked(QListWidgetItem*)),drvNewSignal(self,a1,a2),SLOT(call(QListWidgetItem*)));
]],
OnItemDoubleClicked = [[
QObject::connect(self,SIGNAL(itemDoubleClicked(QListWidgetItem*)),drvNewSignal(self,a1,a2),SLOT(call(QListWidgetItem*)));
]],
OnItemEntered = [[
QObject::connect(self,SIGNAL(itemEntered(QListWidgetItem*)),drvNewSignal(self,a1,a2),SLOT(call(QListWidgetItem*)));
]],
OnItemPressed = [[
QObject::connect(self,SIGNAL(itemPressed(QListWidgetItem*)),drvNewSignal(self,a1,a2),SLOT(call(QListWidgetItem*)));
]],
OnItemSelectionChanged = [[
QObject::connect(self,SIGNAL(itemSelectionChanged(QListWidgetItem*)),drvNewSignal(self,a1,a2),SLOT(call(QListWidgetItem*)));
]],
}
