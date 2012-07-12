module("dialog")

name = "Dialog"
base = "Widget"

funcs = [[
+ Init()

@ SetModal(b bool)
@ IsModal()(b bool)
@ SetResult(r int)
@ Result() (r int)
Open()
Exec()(r int)
Done(r int)
Accept()
Reject()

* OnAccepted(fn func())
* OnRejected(fn func())
]]

qtdrv = {
inc = "<QDialog>",
name = "QDialog *",

Init = [[
drvNewObj(a0,new QDialog);
]],

SetModal = "setModal",
IsModal = "isModal",
SetResult = "setResult",
Result = "result",
Open = "open",
Exec = "exec",
Done = "done",
Accept = "accept",
Reject = "reject",

OnAccepted = [[
QObject::connect(self,SIGNAL(acceped()),drvNewSignal(self,a1,a2),SLOT(call()));
]],

OnRejected = [[
QObject::connect(self,SIGNAL(rejected()),drvNewSignal(self,a1,a2),SLOT(call()));
]],

}
