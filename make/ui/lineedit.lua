module("lineedit")

name = "LineEdit"
base = "Widget"

comment = [[

InputMask:

A	ASCII alphabetic character required. A-Z, a-z.
a	ASCII alphabetic character permitted but not required.
N	ASCII alphanumeric character required. A-Z, a-z, 0-9.
n	ASCII alphanumeric character permitted but not required.
X	Any character required.
x	Any character permitted but not required.
9	ASCII digit required. 0-9.
0	ASCII digit permitted but not required.
D	ASCII digit required. 1-9.
d	ASCII digit permitted but not required (1-9).
#	ASCII digit or plus/minus sign permitted but not required.
H	Hexadecimal character required. A-F, a-f, 0-9.
h	Hexadecimal character permitted but not required.
B	Binary character required. 0-1.
b	Binary character permitted but not required.
>	All following alphabetic characters are uppercased.
<	All following alphabetic characters are lowercased.
!	Switch off case conversion.
\	Use \ to escape the special characters listed above to use them as separators.

InputMask Example:
000.000.000.000;_	IP address; blanks are _.
HH:HH:HH:HH:HH:HH;_	MAC address
0000-00-00	ISO Date; blanks are space
>AAAAA-AAAAA-AAAAA-AAAAA-AAAAA;#	License number; blanks are - and all (alphabetic) characters are converted to uppercase.

]]

funcs = [[
+ Init()
+ InitWithText(text string)
@ SetText(text string)
@ Text() (text string)
@ SetInputMask(text string)
@ InputMask() (text string)
@ SetAlignment(value Alignment)
@ Alignment() (value Alignment)
@ SetCursorPos(pos int)
@ CursorPos() (pos int)
@ SetDragEnabled(b bool)
@ DragEnabled() (b bool)
@ SetReadOnly(b bool)
@ IsReadOnly() (b bool)
@ SetFrame(b bool)
@ HasFrame() (b bool)
@ IsRedoAvailable() (b bool)
@ HasSelected() (b bool)
@ SelectedText() (text string)
@ SelStart() (start int)
SetSel(start int,length int)
CancelSel()
SelectAll()
Copy()
Cut()
Paste()
Redo()
Undo()
Clear()

* OnTextChanged(fn func(string))
* OnEditingFinished(fn func())
* OnReturnPressed(fn func())
]]


qtdrv = {
inc = "<QLineEdit>",
name = "QLineEdit *",

Init = [[
drvNewObj(a0,new QLineEdit);
]],
InitWithText = [[
drvNewObj(a0, new QLineEdit(drvGetString(a1)));
]],

SetText = "setText",
Text = "text",
SetInputMask = "setInputMask",
InputMask = "inputMask",
SetAlignment = "setAlignment",
Alignment = "alignment",
SetCursorPos = "setCursorPosition",
CursorPos = "cursorPosition",
SetDragEnabled = "setDragEnabled",
DragEnabled = "dragEnabled",
SetReadOnly = "setReadOnly",
IsReadOnly = "isReadOnly",
SetFrame = "setFrame",
HasFrame = "hasFrame",
IsRedoAvailable = "isRedoAvailable",
HasSelected = "hasSelectedText",
SelectedText = "selectedText",
SelStart = "selectionStart",
SetSel = "setSelection",
CancelSel = "deselect",
SelectAll = "selectAll",
Copy = "copy",
Cut = "cut",
Paste = "paste",
Redo = "redo",
Undo = "undo",
Clear = "clear",

OnTextChanged = [[
QObject::connect(self,SIGNAL(textChanged(QString)),drvNewSignal(self,a1,a2),SLOT(call(QString)));
]],

OnEditingFinished = [[
QObject::connect(self,SIGNAL(editingFinished(QString)),drvNewSignal(self,a1,a2),SLOT(call(QString)));
]],

OnReturnPressed = [[
QObject::connect(self,SIGNAL(returnPressed(QString)),drvNewSignal(self,a1,a2),SLOT(call(QString)));
]],

}
