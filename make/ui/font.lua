module("font")

name = "Font"
base = "object"

funcs = [[
+ Init()
+ InitWith(family string, size int, weight int)
- Close()
@ SetFamily(text string)
@ Family() (text string)
@ SetPointSize(size int)
@ PointSize() (size int)
@ SetPointSizeF(size float64)
@ PointSizeF() (size float64)
@ SetBold(b bool)
@ Bold() (b bool)
@ SetItalic(b bool)
@ Italic() (b bool)
@ SetKerning(b bool)
@ Kerning() (b bool)
@ SetLetterSpacing(typ int, spacing float64)
@ LetterSpacing() (typ int, spacing float64)
@ SetOverline(b bool)
@ Overline() (b bool)
@ SetPixelSize(size int)
@ PixelSize() (size int)
@ SetRawMode(b bool)
@ RawMode() (b bool)
@ SetRawName(name string)
@ RawName() (name string)
@ SetWeight(weight int)
@ Weight() (weight int)
@ SetFixedPitch(b bool)
@ FixedPitch() (b bool)
@ SetStretch(factor int)
@ Stretch() (factor int)
@ SetStrikeOut(b bool)
@ StrikeOut() (b bool)
@ SetUnderline(b bool)
@ Underline() (b bool)

]]

qtdrv = {
inc = "<QFont>",
name = "QFont*",

Init = [[
drvSetFont(a0,QFont());
]],

InitWith = [[
drvSetFont(a0,QFont(drvGetString(a1),drvGetInt(a2),drvGetInt(a3)));
]],

Close = [[
drvDelFont(a0,self);
]],

SetFamily = "setFamily",
Family = "family",
SetPointSize = "setPointSize",
PointSize = "pointSize",
SetPointSizeF = "setPointSizeF",
PointSizeF = "pointSizeF",
SetBold = "setBold",
Bold = "bold",
SetItalic = "setItalic",
Italic = "italic",
SetKerning = "setKerning",
Kerning = "kerning",
SetLetterSpacing = [[
self->setLetterSpacing((QFont::SpacingType)drvGetInt(a1),drvGetFloat64(a2));
]],
LetterSpacing = [[
drvSetInt(a1,self->letterSpacing());
drvSetInt(a2,self->letterSpacingType());
]],
SetOverline = "setOverline",
Overline = "overline",
SetPixelSize = "setPixelSize",
PixelSize = "pixelSize",
SetRawMode = "setRawMode",
RawMode = "rawMode",
SetRawName = "setRawName",
RawName = "rawName",
SetWeight = "setWeight",
Weight = "weight",
SetFixedPitch = "setFixedPitch",
FixedPitch = "fixedPitch",
SetStretch = "setStretch",
Stretch = "stretch",
SetStrikeOut = "setStrikeOut",
StrikeOut = "strikeOut",
SetUnderline = "setUnderline",
Underline = "underline",

}