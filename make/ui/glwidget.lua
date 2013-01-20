module("glwidget")

name = "GLWidget"
base = "Widget"

funcs = [[
+ Init()
DeleteTexture(id uint)
DoneCurrent()
DoubleBuffer() (b bool)
ConvertToGLFormat(image *Image) (glImage *Image)
SetMouseTracking(enable bool)
RenderText(x int, y int, z int, str string, font *Font)
UpdateGL()
UpdateOverlayGL()
* OnInitializeGL(fn func())
* OnInitializeOverlayGL(fn func())
* OnPaintGL(fn func())
* OnPaintOverlayGL(fn func())
* OnResizeGL(fn func(int,int))
* OnResizeOverlayGL(fn func(int,int))
]]

qtdrv = {
inc = "\"gldrv.h\"",	
name = "QGLWidgetCustom *",

Init = [[
drvNewObj(a0,new QGLWidgetCustom);
]],

DeleteTexture = "deleteTexture",
DoneCurrent = "doneCurrent",
DoubleBuffer = "doubleBuffer",
ConvertToGLFormat = "convertToGLFormat",
SetMouseTracking = "setMouseTracking",
RenderText = "renderText",
UpdateGL = "updateGL",
UpdateOverlayGL = "updateOverlayGL",

OnInitializeGL = [[
	QObject::connect(self,SIGNAL(initializeGLSignal()),drvNewSignal(self,a1,a2),SLOT(call()));
]],
OnPaintGL = [[
	QObject::connect(self,SIGNAL(paintGLSignal()),drvNewSignal(self,a1,a2),SLOT(call()));
]],
OnResizeGL = [[
	QObject::connect(self,SIGNAL(resizeGLSignal(int,int)),drvNewSignal(self,a1,a2),SLOT(call(int,int)));
]],

OnInitializeOverlayGL = [[
	QObject::connect(self,SIGNAL(initializeOverlayGLSignal()),drvNewSignal(self,a1,a2),SLOT(call()));
]],
OnPaintOverlayGL = [[
	QObject::connect(self,SIGNAL(paintOverlayGLSignal()),drvNewSignal(self,a1,a2),SLOT(call()));
]],
OnResizeOverlayGL = [[
	QObject::connect(self,SIGNAL(resizeOverlayGLSignal(int,int)),drvNewSignal(self,a1,a2),SLOT(call(int,int)));
]],
}