module("action")

name = "Action"
base = "object"

funcs = [[
+ Init()
+ InitWithText(text string)
- Close()
@ SetText(text string)
@ Text() (text string)
@ SetIcon(icon *Icon)
@ Icon() (icon *Icon)
@ SetIconText(text string)
@ IconText() (text string)
@ SetToolTip(text string)
@ ToolTip() (text string)
@ SetFont(font *Font)
@ Font() (font *Font)
@ SetCheckable(b bool)
@ IsCheckable() (b bool)
@ SetChecked(b bool)
@ IsChecked() (b bool)
@ SetVisible(b bool)
@ IsVisible() (b bool)
* OnTriggered(fn func(bool))
]]

qtdrv = {
inc = "<QAction>",
name = "QAction*",

Init = [[
drvNewObj(a0, new QAction(qApp));
]],

InitWithText = [[
drvNewObj(a0, new QAction(drvGetString(a1),qApp));
]],

Close = [[
drvDelObj(a0,self);
]],

SetText = "setText",
Text = "text",
SetIcon = "setIcon",
Icon = "icon",
SetIconText = "setIconText",
IconText = "iconText",
SetToolTip = "setToolTip",
ToolTip = "toolTip",
SetFont = "setFont",
Font = "font",
SetCheckable = "setCheckable",
IsCheckable = "isCheckable",
SetChecked = "setChecked",
IsChecked = "isChecked",
SetVisible = "setVisible",
IsVisible = "isVisible",

OnTriggered = [[
QObject::connect(self,SIGNAL(triggered(bool)),drvNewSignal(self,a1,a2),SLOT(call(bool)));
]]

}
