module("hboxlayout")

name = "HBoxLayout"
base = "BoxLayout"


funcs = [[
+Init()
]]

qtdrv = {
inc = "<QHBoxLayout>",	
name = "QHBoxLayout *",

Init = [[
drvNewObj(a0,new QHBoxLayout);
]],

}