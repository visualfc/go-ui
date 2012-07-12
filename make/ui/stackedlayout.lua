module("stackedlayout")

name = "StackedLayout"
base = "baseLayout"


funcs = [[
+Init()
@ SetCurrentIndex(index int)
@ CurrentIndex() (index int)
@ SetCurrentWidget(widget IWidget)
@ CurrentWidget() (widget IWidget)
AddWidget(widget IWidget)
InsertWidget(index int, widget IWidget)
Widget(index int) (widget IWidget)
* OnCurrentChanged(fn func(int))
]]

qtdrv = {
inc = "<QStackedLayout>",	
name = "QStackedLayout *",

Init = [[
drvNewObj(a0,new QStackedLayout);
]],

OnCurrentChanged = [[
QObject::connect(self,SIGNAL(currentChanged(int)),drvNewSignal(self,a1,a2),SLOT(call(int)));
]],
}