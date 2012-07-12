module("radio")

name = "Radio"
base = "baseButton"

funcs = [[
+ Init()
+ InitWithText(text string)
* OnClicked(fn func())
]]


qtdrv = {
inc = "<QRadioButton>",
name = "QRadioButton *",

Init = [[
drvNewObj(a0,new QRadioButton);
]],

InitWithText = [[
drvNewObj(a0, new QRadioButton(drvGetString(a1)));
]],

OnClicked = [[
QObject::connect(self,SIGNAL(clicked()),drvNewSignal(self,a1,a2),SLOT(call()));
]],

}
