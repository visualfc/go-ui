module("groupbox")

name = "GroupBox"
base = "Widget"

funcs = [[
+ Init()
+ InitWithTitle(text string)
@ SetTitle(text string)
@ Title() (text string)
]]


qtdrv = {
inc = "<QGroupBox>",
name = "QGroupBox *",

Init = [[
drvNewObj(a0,new QGroupBox);
]],

InitWithTitle = [[
drvNewObj(a0, new QGroupBox(drvGetString(a1)));
]],

SetTitle = "setTitle",
Title = "title",

}
