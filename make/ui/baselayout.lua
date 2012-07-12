module("baselayout")

name = "baseLayout"
base = "object"

funcs = [[
@ SetSpacing(spac int)
@ Spacing() (spac int)
@ SetMargins(m Margins)
@ Margins() (m Margins)
@ SetMenuBar(widget IWidget)
@ MenuBar() (widget IWidget)
@ Count() (count int)
AddLayout(layout ILayout)
AddWidget(widget IWidget)
RemoveWidget(widget IWidget)
IndexOf(widget IWidget) (index int)
]]

expands = [[
func (p *baseLayout) SetMargin(v int) {
	p.SetMargins(Mr(v,v,v,v))
}
func (p *baseLayout) SetMarginsv(left,top,right,bottom int) {
	p.SetMargins(Mr(left,top,right,bottom))
}

func (p *baseLayout) Marginsv() (int,int,int,int) {
	return p.Margins().Unpack()
}
]]

qtdrv = {
inc = "<QLayout>",
name = "QLayout *",

AddLayout = "addItem",
AddWidget = "addWidget",
SetSpacing = "setSpacing",
Spacing = "spacing",
SetMargins = "setContentsMargins",
Margins = "contentsMargins",
SetMenuBar = "setMenuBar",
MenuBar = "menuBar",
}
