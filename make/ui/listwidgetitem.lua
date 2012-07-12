module("listwidgetitem")

name = "ListWidgetItem"
base = "object"

funcs = [[
+ Init()
+ InitWithText(text string)
- Close()
@ SetText(text string)
@ Text() (text string)
@ SetIcon(icon *Icon)
@ Icon() (icon *Icon)
@ SetSelected(b bool)
@ IsSelected() (b bool)
@ SetHidden(b bool)
@ IsHidden() (b bool)
@ SetFont(font *Font)
@ Font() font *Font)
@ SetForeground(brush *Brush)
@ Foreground() (brush *Brush)
@ SetBackground(brush *Brush)
@ Background() (brush *Brush)
@ SetToolTip(tip string)
@ ToolTip() (tip string)
@ SetTextAlignment(value Alignment)
@ TextAlignment() (value Alignment)
@ SetFlags(value ItemFlag)
@ Flags() (value ItemFlag)
]]

qtdrv = {
--inc = "<ListWidgetItem>",
name = "ShellListWidgetItem *",

Init = [[
drvNewObj(a0,new ShellListWidgetItem);
]],
InitWithText = [[
drvNewObj(a0, new ShellListWidgetItem(drvGetString(a1)));
]],
Close = [[
drvDelObj(a0,self);
]],

SetTextAlignment = [[
self->setTextAlignment((Qt::Alignment)drvGetInt(a1));
]],
TextAlignment = [[
drvSetInt(a1,self->textAlignment());
]],
SetFlags = [[
self->setFlags((Qt::ItemFlag)drvGetInt(a1));
]],
Flags = [[
drvSetInt(a1,self->flags());
]],

}
