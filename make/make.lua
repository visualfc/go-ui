--[[ 
// Copyright 2012 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
--]]

require "ui.app"
require "ui.timer"

require "ui.font"
require "ui.icon"
require "ui.pixmap"
require "ui.image"

require "ui.widget"
require "ui.tabwidget"
require "ui.toolbox"

require "ui.dialog"

require "ui.baselayout"
require "ui.boxlayout"
require "ui.hboxlayout"
require "ui.vboxlayout"
require "ui.stackedlayout"

require "ui.basebutton"
require "ui.button"
require "ui.checkbox"
require "ui.radio"
require "ui.toolbutton"

require "ui.frame"
require "ui.label"

require "ui.groupbox"

require "ui.brush"
require "ui.pen"
require "ui.painter"

require "ui.menubar"
require "ui.menu"
require "ui.action"
require "ui.actiongroup"
require "ui.toolbar"
require "ui.statusbar"
require "ui.dockwidget"

require "ui.combobox"
require "ui.lineedit"

require "ui.systemtray"

require "ui.baseslider"
require "ui.slider"
require "ui.scrollbar"
require "ui.dial"

require "ui.listwidgetitem"
require "ui.listwidget"

require "ui.mainwindow"

require "makelib"

function make()
	cdrv_type = "qtdrv"
	cdrv_func_go2c = lower_title
	begin_def()
		def(app)
		def(timer)
		
		def(font)
		def(pixmap)
		def(icon)	
		def(image)	
		
		def(widget)

		def(action)
		def(actiongroup)
		def(menu)
		def(menubar)
		def(toolbar)
		def(statusbar)
		def(dockwidget)
		
		def(systemtray)

		def(tabwidget)
		def(toolbox)

		def(baselayout)
		def(boxlayout)
		def(hboxlayout)
		def(vboxlayout)
		def(stackedlayout)

		def(basebutton)
		def(button)
		def(checkbox)
		def(radio)
		def(toolbutton)

		def(frame)
		def(label)
		def(groupbox)
		def(dialog)

		def(combobox)
		def(lineedit)
		
		def(baseslider)
		def(slider)
		def(scrollbar)
		def(dial)

		def(brush)
		def(pen)
		def(painter)
		
		def(listwidgetitem)
		def(listwidget)
		
		def(mainwindow)
		
	end_def()

	local ui = io.open("../ui/uiobjs.go","w")
	ui:write(table.concat(go_def.heads,"\n"))
	ui:write("\n")
	ui:write(table.concat(go_def.drvenum,"\n"))
	ui:write("\n")
	ui:write(table.concat(go_def.enums,"\n"))
	ui:write("\n")
	ui:write(table.concat(go_def.newobj,"\n"))
	ui:write("\n")
	ui:write(table.concat(go_def.funcs,"\n"))
	ui:write("\n")
	ui:close()

	local cdrv = io.open(string.format("../%s/cdrv.cpp",cdrv_type),"w")
	cdrv:write(table.concat(c_def.heads,"\n"))
	cdrv:write("\n")
	cdrv:write(table.concat(c_def.include,"\n"))
	cdrv:write("\n")
	cdrv:write(table.concat(c_def.drvenum,"\n"))
	cdrv:write("\n")
	cdrv:write(table.concat(c_def.enums,"\n"))
	cdrv:write("\n")
	cdrv:write(table.concat(c_def.funcs,"\n"))
	cdrv:write("\n")
	cdrv:write(table.concat(c_def.drvfunc,"\n"))
	cdrv:write("\n")
	cdrv:close()
end

make()
