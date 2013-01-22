module("sizepolicy")

name = "SizePolicy"
base = "object"

unmap = true

funcs = [[
+ Init()
+ InitWithPolicy(horizontal SizePolicyPolicy, vertical SizePolicyPolicy, control SizePolicyControlType)
- Close()
@ HorizontalPolicy() (policy SizePolicyPolicy)
@ SetHorizontalPolicy(policy SizePolicyPolicy)
@ VerticalPolicy() (policy SizePolicyPolicy)
@ SetVerticalPolicy(policy SizePolicyPolicy)
@ HasHeightForWidth() (b bool)
Transpose()
]]


qtdrv = {
inc = "<QSizePolicy>",
name = "QSizePolicy *",

Init = [[
drvNewObj(a0, new QSizePolicy());
]],
InitWithPolicy = [[
drvNewObj(a0, new QSizePolicy(drvGetSizePolicyPolicy(a1),drvGetSizePolicyPolicy(a2),drvGetSizePolicyControlType(a3)));
]],
Close = [[
drvDelObj(a0,self);
]],
HorizontalPolicy = "horizontalPolicy",
SetHorizontalPolicy = "setHorizontalPolicy",
VerticalPolicy = "verticalPolicy",
SetVerticalPolicy = "setVerticalPolicy",
HasHeightForWidth = "hasHeightForWidth",
Transpose = "transpose",
}
