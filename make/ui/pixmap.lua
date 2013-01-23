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
@ Width() (width int)
@ Height() (height int)
Scaled(width int, height int, aspectRatioMode AspectRatioMode, transformMode TransformationMode) (pixmap *Pixmap)
ScaledToHeight(height int, transformMode TransformationMode) (pixmap *Pixmap)
ScaledToWidth(width int, transformMode TransformationMode) (pixmap *Pixmap)
ToImage() (image *Image)
Load(fileName string) (b bool)
Save(fileName string, format quality int) (b bool)
Fill(clr color.Color)
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
Width = "width",
Height = "height",
Scaled = "scaled",
ScaledToHeight = "scaledToHeight",
ScaledToWidth = "scaledToWidth",
Scroll = "scroll",
ToImage = "toImage",
Load = [[
	drvSetBool(a1, self->load(drvGetString(a0)));
]],
Save = [[
	drvSetBool(a3, self->save(drvGetString(a1), 0, drvGetInt(a2)));
]],
Fill = "fill",
}
