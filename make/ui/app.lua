module("app")

name = "app"
base = "object"
sync = true

defs = [[
]]

funcs = [[
AppMain()(code int)
Run() (code int)
Exit(code int)
@ SetFont(font *Font)
@ Font() (font *Font)
@ SetActiveWindow(widget IWidget)
@ ActiveWindow() (widget IWidget)
@ SetStyleSheet(style string)
@ StyleSheet() (style string)
CloseAllWindows()

* onRemoveObject(fn func(uintptr))
]]

qtdrv = {
inc = [["qtapp.h"]],
name = "QtApp*",

Init = [[
drvNewObj(a0, new QtApp);
]],
Close = [[
drvDelObj(a0,self);
]],

AppMain = [[
QtApp app;
handle_head *head =(handle_head*)a0;
head->native = &app;
drvSetInt(a1,drv_appmain());
]],

Run = "exec",
Exit = "exit",
SetFont = "setFont",
Font = "font",
SetActiveWindow = "setActiveWindow",
ActiveWindow = "activeWindow",
CloseAllWindows = "closeAllWindows",

onRemoveObject = [[
self->pfnDeleteObj = a1;
]],

}

wtldrv = {
name = "CApp",
Init = [[
_Module.Init(NULL,NULL);
]],
Close = [[
_Module.Term();
]],
}