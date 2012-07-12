module("boxlayout")

name = "BoxLayout"
base = "baseLayout"

funcs = [[
+ Init()
- Close()
@ SetOrientation(value Orientation)
@ Orientation() (value Orientation)
AddLayoutWith(layout ILayout, stetch int)
AddWidgetWith(widget IWidget, stretch int, alignment Alignment)
AddSpacing(size int)
AddStretch(size int)
]]

expands = [[
]]

qtdrv = {
inc = "<QBoxLayout>",
name = "QBoxLayout *",

Init = [[
drvNewObj(a0,new QBoxLayout(QBoxLayout::TopToBottom));
]],

Close = [[
drvDelObj(a0,self);
]],

SetOrientation = [[
switch (drvGetInt(a1)) {
	case Qt::Horizontal: {
		self->setDirection(QBoxLayout::TopToBottom);
		break;		
	}
	case Qt::Vertical: {
		self->setDirection(QBoxLayout::LeftToRight);
		break;
	}
}
]],
Orientation = [[
switch((int)self->direction()) {
	case QBoxLayout::TopToBottom: {
		drvSetInt(a1,Qt::Horizontal);
		break;
	}
	case QBoxLayout::LeftToRight: {
		drvSetInt(a1,Qt::Vertical);
		break;
	}
}
]],
AddLayoutWith = [[
self->addLayout((QLayout*)drvGetNative(a1),drvGetInt(a2));
]],
AddWidgetWith = [[
self->addWidget((QWidget*)drvGetNative(a1),drvGetInt(a2),Qt::Alignment(drvGetInt(a3)));
]],
AddSpacing = "addSpacing",
AddStretch = "addStretch",
}
