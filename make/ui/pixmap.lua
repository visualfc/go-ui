module("pixmap")

name = "Pixmap"
base = "object"

funcs = [[
+ Init()
+ InitWithFile(name string)
+ InitWithData(data []byte)
+ InitWithSize(width int, height int)
- Close()
@ Depth() (n int)
@ Size() (sz Size)
@ Rect() (rc Rect)
@ HasAlpha() (b bool)
@ HasAlphaChannel() (b bool)
@ IsNull() (b bool)

]]


qtdrv = {
inc = "<QPixmap>",
name = "QPixmap *",

Init = [[
drvSetPixmap(a0, QPixmap());
]],
InitWithFile = [[
drvSetPixmap(a0, QPixmap(drvGetString(a1)));
]],
InitWithSize = [[
drvSetPixmap(a0, QPixmap(drvGetInt(a1),drvGetInt(a2)));
]],
InitWithData = [[
QPixmap pixmap;
if (pixmap.loadFromData(drvGetByteArray(a1))) {
	drvSetPixmap(a0, pixmap);
}
]],
Close = [[
drvDelPixmap(a0,self);
]],

Depth = "depth",
Size = "size",
Rect = "rect",
HasAlpha = "hasAlpha",
HasAlphaChannel = "hasAlphaChannel",
IsNull = "isNull",

}
