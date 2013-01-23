module("basescrollarea")

name = "baseScrollArea"
base = "Widget"

funcs = [[
* CornerWidget() (w IWidget)
* HorizontalScrollBar() (s *ScrollBar)
* VerticalScrollBar() (s *ScrollBar)
* Viewport() (w IWidget)
]]


qtdrv = {
inc = "<QAbstractScrollArea>",
name = "QAbstractScrollArea *",

CornerWidget = "cornerWidget",
HorizontalScrollBar = "horizontalScrollBar",
VerticalScrollBar = "verticalScrollBar",
Viewport = "viewport",
}
