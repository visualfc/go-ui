module("icon")

name = "Icon"
base = "object"

funcs = [[
+ Init()
+ InitWithFile(name string)
+ InitWithPixmap(pixmap *Pixmap)
- Close()
]]

qtdrv = {
inc = "<QIcon>",
name = "QIcon*",

Init = [[
drvSetIcon(a0, QIcon());
]],
InitWithFile = [[
drvSetIcon(a0, QIcon(drvGetString(a1)));
]],
InitWithPixmap = [[
drvSetIcon(a0, QIcon(drvGetPixmap(a1)));
]],
Close = [[
drvDelIcon(a0,self);
]],

}