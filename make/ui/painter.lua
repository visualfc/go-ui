module("painter")

name = "Painter"
base = "object"
unmap = true

funcs = [[
+ Init()
+ InitWithImage(image *Image)
- Close()
InitFrom(w IWidget)
Begin(w IWidget)
End()
@ SetFont(font *Font)
@ Font() (font *Font)
@ SetPen(pen *Pen)
@ Pen() (pen *Pen)
@ SetBrush(brush *Brush)
@ Brush() (brush *Brush)
DrawPoint(pt Point)
DrawPoints(pts []Point)
DrawLine(pt1 Point, pt2 Point)
DrawLines(pts []Point)
DrawPolyline(pts []Point)
DrawPolygon(pts []Point)
DrawRect(rc Rect)
DrawRects(rcs []Rect)
DrawRoundedRect(rc Rect, xRadius float64, yRadius float64, sizeMode int)
DrawEllipse(rc Rect)
DrawArc(rc Rect, startAngle int, spanAngle int)
DrawChord(rc Rect, startAngle int, spanAngle int)
DrawPie(rc Rect, startAngle int, spanAngle int)
DrawText(pt Point,text string)
DrawTextRect(rc Rect, flags int, text string) (bound Rect)
DrawPixmap(pt Point,pixmap *Pixmap)
DrawPixmapEx(pt Point,pixmap *Pixmap, source Rect)
DrawPixmapRect(rc Rect,pixmap *Pixmap)
DrawPixmapRectEx(rc Rect,pixmap *Pixmap, source Rect)
DrawImage(pt Point, image *Image)
DrawImageEx(pt Point, image *Image, source Rect)
DrawImageRect(rc Rect, image *Image)
DrawImageRectEx(rc Rect, image *Image, source Rect)
FillRect(rc Rect, color uint)
FillRectF(rc RectF, color uint)
]]

qtdrv = {
inc = "<QPainter>",
name = "QPainter *",

Init = [[
drvNewObj(a0,new QPainter);
]],
InitWithImage = [[
drvNewObj(a0,new QPainter((QImage*)drvGetNative(a1)));
]],

Close = [[
drvDelObj(a0,self);
]],

InitFrom = "initFrom",
Begin = "begin",
End = "end",
SetFont = "setFont",
Font = "font",
SetPen = "setPen",
Pen = "pen",
SetBrush = "setBrush",
Brush = "brush",
DrawPoint = "drawPoint",
DrawPoints = "drawPoints",
DrawLine = "drawLine",
DrawLines = "drawLines",
DrawPolyline = "drawPolyline",
DrawPolygon = "drawPolygon",
DrawRect = "drawRect",
DrawRects = "drawRects",
DrawRoundedRect = [[
self->drawRoundedRect(drvGetRect(a1),drvGetFloat64(a2),drvGetFloat64(a3),(Qt::SizeMode)drvGetInt(a4));
]],
DrawEllipse = "drawEllipse",
DrawArc = "drawArc",
DrawChord = "drawChord",
DrawPie = "drawPie",
DrawText = "drawText",
DrawTextRect = [[
QRect bound;
self->drawText(drvGetRect(a1),drvGetInt(a2),drvGetString(a3),&bound);
drvSetRect(a4,bound);
]],
DrawPixmap = "drawPixmap",
DrawPixmapEx = "drawPixmap",
DrawPixmapRect = "drawPixmap",
DrawPixmapRectEx = "drawPixmap",
DrawImage = "drawImage",
DrawImageEx = "drawImage",
DrawImageRect = "drawImage",
DrawImageRectEx = "drawImage",
FillRect = "fillRect",
FillRectF = "fillRect",
}