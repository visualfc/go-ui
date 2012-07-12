module("vboxlayout")

name = "VBoxLayout"
base = "BoxLayout"

funcs = [[
+Init()
]]

qtdrv = {
inc = "<QVBoxLayout>",	
name = "QVBoxLayout *",

Init = [[
drvNewObj(a0,new QVBoxLayout);
]],
}