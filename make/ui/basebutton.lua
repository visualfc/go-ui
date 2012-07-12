module("basebutton")

name = "baseButton"
base = "Widget"

funcs = [[
@ SetText(text string)
@ Text() (text string)
@ SetIcon(icon *Icon)
@ Icon() (icon *Icon)
@ SetIconSize(sz Size)
@ IconSize() (sz Size)
@ SetCheckable(b bool)
@ IsCheckable() (b bool)
@ SetDown(b bool)
@ IsDown() (b bool)
]]

qtdrv = {
inc = "<QAbstractButton>",
name = "QAbstractButton *",

SetText = "setText",
Text = "text",
SetIcon = "setIcon",
Icon = "icon",
SetIconSize = "setIconSize",
IconSize = "iconSize",
SetCheckable = "setCheckable",
IsCheckable = "isCheckable",
SetDown = "setDown",
IsDown = "isDown",

}
