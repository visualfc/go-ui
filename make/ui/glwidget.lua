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
* OnUpdateGL(fn func())
* OnUpdateOverlayGL(fn func())
* OnGLDraw(fn func())
* OnGLInit(fn func())
* OnInitializeGL(fn func())
* OnInitializeOverlayGL(fn func())
* OnPaintGL(fn func())
* OnPaintOverlayGL(fn func())
* OnResizeGL(fn func(int,int))
* OnResizeOverlayGL(fn func(int,int))
]]

qtdrv = {
inc = "\"gldrv.h\"",	
name = "QGLWidget *",

Init = [[
drvNewObj(a0,new QGLWidgetCustom);
]],

DeleteTexture = "deleteTexture",
DoneCurrent = "doneCurrent",
DoubleBuffer = "doubleBuffer",
ConvertToGLFormat = "convertToGLFormat",
SetMouseTracking = "setMouseTracking",
RenderText = "renderText",
OnUpdateGL = [[
	QObject::connect(self,SIGNAL(updateGL()),drvNewSignal(self,a1,a2),SLOT(call()));
]],
OnUpdateOverlayGL = [[
	QObject::connect(self,SIGNAL(updateOverlayGL()),drvNewSignal(self,a1,a2),SLOT(call()));
]],
OnGLDraw = [[
	QObject::connect(self,SIGNAL(glDraw()),drvNewSignal(self,a1,a2),SLOT(call()));
]],
OnGLInit = [[
	QObject::connect(self,SIGNAL(glInit()),drvNewSignal(self,a1,a2),SLOT(call()));
]],
OnInitializeGL = [[
	QObject::connect(self,SIGNAL(initializeGLSignal()),drvNewSignal(self,a1,a2),SLOT(call()));
]],
OnInitializeOverlayGL = [[
	QObject::connect(self,SIGNAL(initializeOverlayGL()),drvNewSignal(self,a1,a2),SLOT(call()));
]],
OnPaintGL = [[
	QObject::connect(self,SIGNAL(paintGLSignal()),drvNewSignal(self,a1,a2),SLOT(call()));
]],
OnPaintOverlayGL = [[
	QObject::connect(self,SIGNAL(paintOverlayGL()),drvNewSignal(self,a1,a2),SLOT(call()));
]],
OnResizeGL = [[
	QObject::connect(self,SIGNAL(resizeGLSignal(int,int)),drvNewSignal(self,a1,a2),SLOT(call(int,int)));
]],
OnResizeOverlayGL = [[
	QObject::connect(self,SIGNAL(resizeOverlayGLSignal(int,int)),drvNewSignal(self,a1,a2),SLOT(call(int,int)));
]],
}