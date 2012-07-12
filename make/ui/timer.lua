module("timer")

name = "Timer"
base = "object"

unmap = true

funcs = [[
+ Init()
- Close()
@ SetInterval(msec int)
@ Interval() (msec int)
@ SetSingleShot(b bool)
@ IsSingleShot() (b bool)
@ IsActive() (b bool)
@ TimerId() (id int)
Start()
StartWith(msec int)
Stop()
* OnTimeout(fn func())
]]

expands = [[

func AfterFunc(ms int,fn func()) *Timer {
	timer := NewTimer()
	timer.SetSingleShot(true)
	timer.StartWith(ms)
	timer.OnTimeout(func() {
		timer.Close()
		fn()		
	})
	return timer
}

]]


qtdrv = {
inc = "<QTimer>",
name = "QTimer *",

Init = [[
drvNewObj(a0,new QTimer(theApp));
]],
Close = [[
drvDelObj(a0,self);
]],
SetInterval = "setInterval",
Interval = "interval",
SetSingleShot = "setSingleShot",
IsSingleShot = "isSingleShot",
IsActive = "isActive",
TimerId = "timerId",
Start = "start",
StartWith = "start",
Stop = "stop",

OnTimeout = [[
QObject::connect(self,SIGNAL(timeout()),drvNewSignal(self,a1,a2),SLOT(call()));
]],

}
