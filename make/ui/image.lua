module("image")

name = "Image"
base = "object"

funcs = [[
+ Init()
+ InitWithFile(name string)
+ InitWithSize(width int, height int)
- Close()
@ Depth() (n int)
@ Size() (sz Size)
@ Rect() (rc Rect)
Fill(color uint)
Scaled(width int, height int, aspectRatioMode AspectRatioMode, transformMode TransformationMode) (image *Image)
]]


qtdrv = {
inc = "<QImage>",
name = "QImage *",

Init = [[
drvSetImage(a0, QImage());
]],
InitWithFile = [[
drvSetImage(a0, QImage(drvGetString(a1)));
]],
InitWithSize = [[
drvSetImage(a0, QImage(drvGetInt(a1),drvGetInt(a2),QImage::Format_ARGB32_Premultiplied));
]],
Close = [[
drvDelImage(a0,self);
]],

Depth = "depth",
Size = "size",
Rect = "rect",
Fill = "fill",
Scaled = "scaled",
}
