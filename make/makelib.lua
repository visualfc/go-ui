--[[ 
 Copyright 2012 visualfc <visualfc@gmail.com>. All rights reserved.

 This library is free software; you can redistribute it and/or
 modify it under the terms of the GNU Lesser General Public
 License as published by the Free Software Foundation; either
 version 2.1 of the License, or (at your option) any later version.

 This library is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 Lesser General Public License for more details.
--]]

-- tag +,-,@,*
function get_func(s)
	local tag,name,input,output = string.match(s,"([%+%-%@%*]?)%s*([_%w%*%[%]]+)%s*(%b())%s*(%b())")
	if name == nil then
		tag,name,input = string.match(s,"([%+%-%@%*]?)%s*([_%w%*%[%]]+)%s*(%b())")
	end
	if name == nil then
		tag,name = string.match(s,"([%+%-%@%*]?)%s*([_%w%*%[%]]+)")
	end

	if name == nil then
		return nil
	end
	input = input or "()"
	output = output or "()"

	input = string.sub(input,2,-2)..","
	output = string.sub(output,2,-2)..","

	local func = {}
	func.tag = tag or ""
	func.name = name
	func.input = {}
	func.output = {}

	for v,t in string.gmatch(input,"([_%w]+)%s+([_%w%.%*%[%]%(%),]+),") do
		func.input[#func.input+1] = {var=v,type=t}
	end

	for v,t in string.gmatch(output,"([_%w]+)%s+([_%w%.%*%[%]%(%),]+),") do
		func.output[#func.output+1] = {var=v,type=t}
	end

	return func
end

function get_lines(s)
	local i = 0
	local j = 0
	local lines = {}
	while true do
		j = string.find(s.."\n","\n",i)
		if j == nil then
			break
		end
		local line = string.sub(s,i,j-1)
		if line ~= "" then
			lines[#lines+1] = line
		end
		i = j+1
	end
	return lines
end

function get_funcs(s)
	local lines = get_lines(s)
	local funcs = {}
	for k,v in ipairs(lines) do
		local func = get_func(v)
		if func ~= nil then
			funcs[#funcs+1] = func
		end
	end
	return funcs
end

function get_comment(s)
	local lines = get_lines(s)
	local comment = {}
	for k,v in ipairs(lines) do
		comment[#comment+1] = "// "..v
	end
	return table.concat(comment,"\n//\n")
end

function make_params(t)
	local o = {}
	for k,v in ipairs(t) do
		o[k] = v.var .." ".. v.type
	end
	return "("..table.concat(o,",")..")"
end

function print_funcs(funcs)
	for k,v in ipairs(funcs) do
		print(k,v.tag,v.name,make_params(v.input),make_params(v.output))
	end
end

function lower_title(s)
	return string.lower(string.sub(s,1,1))..string.sub(s,2)
end

function upper_title(s)
	return string.upper(string.sub(s,1,1))..string.sub(s,2)
end

--cdrv_type = "qtdrv"

go_def = {}
c_def = {}

cdrv_func_go2c = function(s) return s end

function begin_def()
	go_def.heads = {}
	go_def.drvenum = {}
	go_def.enums = {}
	go_def.defs = {}
	go_def.funcs = {}
	go_def.newobj = {}

	c_def.heads = {}
	c_def.include = {}
	c_def.drvenum = {}
	c_def.drvfunc = {}
	c_def.enums = {}
	c_def.funcs = {}
	
	go_def.heads[#go_def.heads+1] = [[
// Copyright 2012 visualfc <visualfc@gmail.com>. All rights reserved.
//
// This library is free software; you can redistribute it and/or
// modify it under the terms of the GNU Lesser General Public
// License as published by the Free Software Foundation; either
// version 2.1 of the License, or (at your option) any later version.
//
// This library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// Lesser General Public License for more details.
]]
	c_def.heads[#c_def.heads+1] = [[
// Copyright 2012 visualfc <visualfc@gmail.com>. All rights reserved.
//
// This library is free software; you can redistribute it and/or
// modify it under the terms of the GNU Lesser General Public
// License as published by the Free Software Foundation; either
// version 2.1 of the License, or (at your option) any later version.
//
// This library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// Lesser General Public License for more details.
]]

	go_def.heads[#go_def.heads+1] = [[
package ui

import (
	"runtime"
	"unsafe"
	"image/color"
)
]]
	go_def.newobj[#go_def.newobj+1] = [[
func registerAllClass() {
	]]

	go_def.drvenum[#go_def.drvenum+1] = "// drvclass enums"
	go_def.drvenum[#go_def.drvenum+1] = "const ("
	go_def.drvenum[#go_def.drvenum+1] = "\tCLASSID_NONE = iota"
	
	c_def.drvenum[#c_def.drvenum+1] = "// drvclass enums"
	c_def.drvenum[#c_def.drvenum+1] = "enum DRVCLASS_ENUM {"
	c_def.drvenum[#c_def.drvenum+1] = "    DRVCLASS_NONE = 0,"

	c_def.include[#c_def.include+1] = [[#include "cdrv.h"]]

	c_def.drvfunc[#c_def.drvfunc+1] = [[
typedef int (*DRV_CALLBACK)(void* fn, void *a1,void* a2,void* a3,void* a4);
typedef int (*DRV_RESULT)(void* ch,int r);
typedef int (*DRV_APPMAIN)();
typedef void (*UTF8_INFO_COPY)(void *info,const char *data, int size);

static DRV_CALLBACK pfn_drv_callback;
static DRV_RESULT pfn_drv_result;
static DRV_APPMAIN pfn_drv_appmain;
static UTF8_INFO_COPY pfn_utf8_info_copy;

int drv_callback(void* fn, void *a1,void* a2,void* a3,void* a4)
{
    return pfn_drv_callback(fn,a1,a2,a3,a4);
}

int drv_result(void* ch, int r)
{
    return pfn_drv_result(ch,r);
}

int drv_appmain()
{
	return pfn_drv_appmain();
}

void utf8_info_copy(void *info, const char *data, int size)
{
    pfn_utf8_info_copy(info,data,size);
}

int drv(int drvcls, int drvid, void *exp, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    switch(drvcls) {
    case -1:
        pfn_drv_callback = (DRV_CALLBACK)exp;
        return 1;
    case -2:
        pfn_drv_result = (DRV_RESULT)exp;
        return 1;
    case -3:
        pfn_utf8_info_copy = (UTF8_INFO_COPY)exp;
        return 1;
	case -4:
		pfn_drv_appmain = (DRV_APPMAIN)exp;
		return 1;
    case _CLASSID_APP:
        return drv_app(drvid,a0,a1,a2,a3,a4,a5,a6);
    default:
        QMetaObject::invokeMethod(theApp,"drv",
                                  Q_ARG(int,drvcls),
                                  Q_ARG(int,drvid),
                                  Q_ARG(void*,exp),							  
                                  Q_ARG(void*,a0),
                                  Q_ARG(void*,a1),
                                  Q_ARG(void*,a2),
                                  Q_ARG(void*,a3),
                                  Q_ARG(void*,a4),
                                  Q_ARG(void*,a5),
                                  Q_ARG(void*,a6));
        return 0;
    }
    return 1;
}

int _drv(int drvcls, int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    switch(drvcls) {]]

end

function end_def()
	go_def.drvenum[#go_def.drvenum+1] = ")"
	go_def.newobj[#go_def.newobj+1] = "}"
	c_def.drvenum[#c_def.drvenum+1] = "    DRVCLASS_LAST"
	c_def.drvenum[#c_def.drvenum+1] = "};"
	c_def.drvfunc[#c_def.drvfunc+1] = [[
    default:
        return 0;
    }
    return 1;
}
]]
end

function def(item)
	local go = {}
	local c = {}
	local name = item.name
	local base = item.base or ""
	local name_tag = string.upper(name)
	local name_cls
	local export = false
	if string.sub(name,1,1) == string.upper(string.sub(name,1,1)) then
		export = true
	end
	
	if export then
		name_cls = "CLASSID_"..string.upper(name)
	else
		name_cls = "_CLASSID_"..string.upper(name)
	end

	assert(item.funcs ~= nil,string.format("%s funcs not defined!",name))

	local funcs = get_funcs(item.funcs)
	local cdrv = item[cdrv_type]
	assert(cdrv ~= nil,string.format("%s not defined!",cdrv_type))
	local cself = "self."
	local cdrv_name = string.match(cdrv.name,"([%w_]*)%s*%*")
	if cdrv_name ~= nil then
		cdrv.name = cdrv_name
		cself = "self->"
	end
	
	go.enums = {}
	go.defs = {}
	go.funcs = {}
	go.setattr = {}
	go.getattr = {}

	c.enums = {}
	c.funcs = {}

	go_def.drvenum[#go_def.drvenum+1] = "\t"..name_cls	
	c_def.drvenum[#c_def.drvenum+1] = "    "..name_cls..","

	go.enums[#go.enums+1] = string.format("// %s drvid enums",name_cls)
	go.enums[#go.enums+1] = "const ("
	go.enums[#go.enums+1] = string.format("\t_ID_%s = iota",name_tag.."_NONE")

	c.enums[#c.enums+1] = string.format("// %s drvid enums",name_cls)
	c.enums[#c.enums+1] = string.format("enum %s {","DRVID_"..name_tag.."_ENUM")
	c.enums[#c.enums+1] = string.format("    _ID_%s = 0,",name_tag.."_NONE")

	c.funcs[#c.funcs+1] = string.format([[
int drv_%s(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)]],string.lower(name))
	c.funcs[#c.funcs+1] = "{"
	if cdrv.self ~= nil then
		c.funcs[#c.funcs+1] = cdrv.self
	elseif cself == "self->" then
		c.funcs[#c.funcs+1] = string.format([[
    %s *self = (%s*)drvGetNative(a0);]],cdrv.name,cdrv.name)
	else
		c.funcs[#c.funcs+1] = string.format([[
    %s self = drvGet%s(a0);]],cdrv.name,name)
	end

	c.funcs[#c.funcs+1] = "    switch (drvid) {"
	
	local comment = string.format([[
// struct %s
//]],name)	
	if item.comment ~= nil then
		comment = comment .. "\n" .. get_comment(item.comment)
	end
	-- go struct def
	go.defs[#go.defs+1] = string.format([[
%s
type %s struct {
	%s
}
]],comment,name,base)
	go.defs[#go.defs+1] = string.format([[
func (p *%s) Name() string {
	return "%s"
}
]],name,name)
	go.defs[#go.defs+1] = string.format([[
func (p *%s) String() string {
	return DumpObject(p)
}]],name)

	go.setattr[#go.setattr+1] = string.format([[
func (p *%s) SetAttr(attr string, value interface{}) bool {
	switch attr {]],name)
	go.getattr[#go.getattr+1] = string.format([[
func (p *%s) Attr(attr string) interface{} {
	switch attr {]],name)
	
	if item.defs ~= nil then
		go.defs[#go.defs+1] = item.defs
	end
	-- funcs enums
	for k,v in ipairs(funcs) do
		local func = v.name
		local tag = v.tag
		local enum = "_ID_"..name_tag.."_"..string.upper(v.name)
		go.enums[#go.enums+1] = "\t"..enum
		c.enums[#c.enums+1] = "    "..enum..","
		local go_in = {}
		local go_out = {}
		local go_var = {}
		local go_var_ex = {}
		local go_body1 = {}
		local go_body2 = {}
		local c_in = {}
		local c_out = {}
		local go_in_type = ""
		local go_out_type = ""
		local index = 0
		for k,v in pairs(v.input) do
			index = index+1
			local var = v.var
			local type = v.type
			local _c_type = type
			local ptr,typ = string.match(type,"(%*)%s*(%w+)")
			if ptr ~= nil then
				_c_type = typ
			end
			local ar,typ = string.match(type,"(%[%])%s*(%w+)")
			if ar ~= nil then
				_c_type = typ .. "Array"
			end
			go_in[#go_in+1] = var.." "..type
			go_var[#go_var+1] = var
			go_in_type = v.type
			if type == "string" then
				go_var_ex[#go_var_ex+1] = "unsafe.Pointer((*string_info)(unsafe.Pointer(&"..var..")))"
			elseif type == "IWidget" or type == "ILayout" or type == "iobj" then
				go_var_ex[#go_var_ex+1] = "unsafe.Pointer("..var..".(iobj).info())"
			elseif type == "color.Color" then
				go_body1[#go_body1+1] = string.format("\tsh_%s := make_rgba(%s)",var,var)
				var = "sh_"..var
				go_var_ex[#go_var_ex+1] = "unsafe.Pointer(&"..var..")"
			elseif ar ~= nil then
				go_var_ex[#go_var_ex+1] = "unsafe.Pointer((*slice_info)(unsafe.Pointer(&"..var..")))"
			elseif ptr ~= nil then
				go_var_ex[#go_var_ex+1] = "unsafe.Pointer("..var..")"
			else
				go_var_ex[#go_var_ex+1] = "unsafe.Pointer(&"..var..")"
			end
			
			local c_type = "drvGet"..string.upper(string.sub(_c_type,1,1))..string.sub(_c_type,2,-1)
			if type == "iobj" then
				c_type = "drvGetNative"
			elseif type == "IWidget" then
				c_type = "drvGetWidget"
			elseif type == "ILayout" then
				c_type = "drvGetLayout"
			elseif type == "color.Color" then
				c_type = "drvGetColor"
			end
			c_in[#c_in+1] = c_type
		end				
		for k,v in pairs(v.output) do
			index = index+1
			local var = v.var
			local type = v.type
			local _c_type = type
			local ptr,typ = string.match(type,"(%*)%s*(%w+)")
			if ptr ~= nil then
				_c_type = typ
			end			
			local c_type = "drvSet"..string.upper(string.sub(_c_type,1,1))..string.sub(_c_type,2,-1)			
			go_out[#go_out+1] = var.." "..type
			go_out_type = type
			if type == "string" then
				var = "sh_"..var
				go_body1[#go_body1+1] = string.format("\tvar %s utf8_info",var)
				go_body2[#go_body2+1] = string.format("\t%s = %s.String()",v.var,var)
			elseif type == "bool" then
				var = "b_"..var
				go_body1[#go_body1+1] = string.format("\tvar %s int",var)
				go_body2[#go_body2+1] = string.format("\t%s = %s != 0",v.var,var)
			elseif type == "color.Color" then
				var = "sh_"..var
				go_body1[#go_body1+1] = string.format("\tvar %s rgba",var)
				go_body2[#go_body2+1] = string.format("\t%s = %s",v.var,var)				
				c_type = "drvSetColor"
			elseif type == "IObject" or 
				   type == "IWidget" or 
				   type == "ILayout" then
				var = "oi_"..var
				go_body1[#go_body1+1] = string.format("\tvar %s obj_info",var)
				go_body2[#go_body2+1] = string.format([[
	if %s.native != 0 {
		item := FindObject(%s.native)
		if item != nil {
			%s = item.(%s)
		}
	}]],var,var,v.var,v.type)
				c_type = "drvSetHandle"
			elseif ptr then
				local id = "CLASSID_"..string.upper(_c_type)
				var = "oi_"..var
				go_body1[#go_body1+1] = string.format("\tvar %s obj_info",var)
				go_body2[#go_body2+1] = string.format([[
	if %s.native != 0 {
		v := FindObject(%s.native)
		if v == nil {
			v = NewObjectWithNative(%s,%s.native)
		}
		if v != nil {
			%s = v.(%s)
		} 
	}]],var,var,id,var,v.var,v.type)
			end
			go_var[#go_var+1] = "&"..var
			if string.match(var,"oi_") then
				go_var_ex[#go_var_ex+1] = "unsafe.Pointer(&"..var..")"
			elseif type == "string" then
				go_var_ex[#go_var_ex+1] = "unsafe.Pointer((*string_info)(unsafe.Pointer(&"..var..")))"
			elseif type == "iobj" then
				go_var_ex[#go_var_ex+1] = "unsafe.Pointer("..var..".info())"
			elseif _c_type == "iobj" and ptr ~= nil then
				go_var_ex[#go_var_ex+1] = "unsafe.Pointer((*"..var..").info())"
			elseif ar ~= nil then
				go_var_ex[#go_var_ex+1] = "unsafe.Pointer((*slice_info)(unsafe.Pointer(&"..var..")))"
			elseif ptr ~= nil then
				go_var_ex[#go_var_ex+1] = "unsafe.Pointer("..var..")"
			else
				go_var_ex[#go_var_ex+1] = "unsafe.Pointer(&"..var..")"
			end			
			
			c_out[#c_out+1] = c_type
		end

		local go_head = string.format("func (p *%s) %s(%s)",name,func,table.concat(go_in,","))
		if #go_out > 0 then
			go_head = go_head .. string.format("(%s)",table.concat(go_out,","))
		end
		local go_new_head = nil
		local body = ""
		if tag == "+" then
			--"init and new"
			go_body1[#go_body1+1] = string.format("\tp.classid = %s",name_cls)
			if item.unmap then
				go_body2[#go_body2+1] = string.format("\truntime.SetFinalizer(p,(*%s).Close)",name)
			else
				go_body2[#go_body2+1] = string.format("\tInsertObject(p)")
			end
			go_head = go_head .. " *"..name
			local new = "new_"
			if export then
				new = "New"
			end
			go_new_head = "func "..string.gsub(func,"%w[%l]+",new..name,1) .. "("..table.concat(go_in,",") .. ") *"..name
			if export and func == "Init" and #v.input == 0 then
				go_def.newobj[#go_def.newobj+1] = string.format([[
	RegisterClass(%q,%s,func() IObject {
		return New%s()
	})]],name,name_cls,name)	
				go_def.newobj[#go_def.newobj+1] = string.format([[
	RegisterClassNative(%s,func(native uintptr) IObject {
		obj := new(%s)
		obj.classid = %s
		obj.native = native
		InsertObject(obj)
		return obj
	})]],name_cls,name,name_cls,name)		
			end
		elseif tag == "-" then
			-- "delete"
			go_head = go_head .. "(err error)"
			if item.unmap then
				go_body1[#go_body1+1] = [[
	if p == nil || p.native == 0 {
		return
	}]]		
				go_body2[#go_body2+1] = [[
	p.native = 0
	runtime.SetFinalizer(p,nil)]]					
			else
			go_body1[#go_body1+1] = [[
	if p == nil || p.native == 0 {
		return
	}]]
			go_body2[#go_body2+1] = [[
	RemoveObject(p.native)
	p.native = 0]]
			end
		elseif tag == "@" then
			-- "attribute"
			local set = string.match(func,"^Set(%u%w+)")
			if set ~= nil and #go_in == 1 then
				local attr = string.lower(set)
				go.setattr[#go.setattr+1] = string.format([[
	case %q:
		if v, ok := value.(%s); ok {
			p.%s(v)
			return true
		}
		return false]],attr,go_in_type,func)
			elseif #go_out == 1 then
				local attr = string.lower(func)
				local is = string.match(func,"^Is(%u%w+)")
				if is ~= nil then
					attr = string.lower(is)
				else
					local has = string.match(func,"^Has(%u%w+)")
					if has ~= nil then
						attr = string.lower(has)
					end
				end
				go.getattr[#go.getattr+1] = string.format([[
	case %q:
		return p.%s()]],attr,func)
			end
		elseif tag == "*" then
			-- "event"
		else
			-- "func"
		end
		
		if tag == "*" then
			if name == "app" then
				go_body = string.format("\t_drv_event(%s,%s,p",name_cls,enum)
			else
				go_body = string.format("\t_drv_event_ch(%s,%s,p",name_cls,enum)
			end
		elseif name == "app" then
			go_body = string.format("\t_drv(%s,%s,unsafe.Pointer(p.info())",name_cls,enum)
		else	
			go_body = string.format("\t_drv_ch(%s,%s,unsafe.Pointer(p.info())",name_cls,enum)
		end

		
		while #go_var_ex < 9 do
			go_var_ex[#go_var_ex+1] = "nil"
		end
		
		if tag == "*" then
			go_body = go_body..","..table.concat(go_var,",")
		else					
			go_body = go_body..","..table.concat(go_var_ex,",")
		end
								
		go_body = go_body..")"

		local go_func = {}
		go_func[#go_func+1] = go_head .. " {"
		if #go_body1 > 0 then
			go_func[#go_func+1] = table.concat(go_body1,"\n")
		end
		go_func[#go_func+1] = go_body
		if #go_body2 > 0 then
			go_func[#go_func+1] = table.concat(go_body2,"\n")
		end
		if tag == "+" then
			go_func[#go_func+1] = "\treturn p"
		else
			go_func[#go_func+1] = "\treturn"
		end
		go_func[#go_func+1] = "}"
		if go_new_head ~= nil then
			local go_new = {}
			go_new[#go_new+1] = go_new_head .. "{"
			go_new[#go_new+1] = string.format("\treturn new(%s).%s(%s)",name,func,table.concat(go_var,","))
			go_new[#go_new+1] = "}"
			go.funcs[#go.funcs+1] = table.concat(go_new,"\n").."\n"
		end
		go.funcs[#go.funcs+1] = table.concat(go_func,"\n").."\n"
		--auto make cdrv[func]		
		if cdrv[func] == nil then
			cdrv[func] = cdrv_func_go2c(func)
		end
		--assert(cdrv[func] ~= nil,string.format("cdrv <%s> class %s : %s not defined!",cdrv_type,name,func))
		local cbody = cdrv[func]
		if string.find(cbody,"[%;%(%)]") == nil then		
			if #c_out == 0 then
				local c_type = {}
				for k,v in ipairs(c_in) do
					c_type[#c_type+1] = v .. "(a"..k..")"
				end
				cbody = cself..cbody.."("..table.concat(c_type,",")..");"
			elseif #c_out == 1 then
				local c_type = {}
				for k,v in ipairs(c_in) do
					c_type[#c_type+1] = v .. "(a"..k..")"
				end
				cbody = c_out[1].."(a"..(#c_in+1)..","..cself..cbody.."("..table.concat(c_type,",").."));"
			end			
		end
		c.funcs[#c.funcs+1] = string.format([[
    case %s: {]],enum)
		--c.funcs[#c.funcs+1] = table.concat(c_var,"\n")
		local c_lines = get_lines(cbody)
		for k, v in ipairs(c_lines) do
			c.funcs[#c.funcs+1] = "        "..v
		end
		c.funcs[#c.funcs+1] = [[
        break;
    }]]
	end

	go.enums[#go.enums+1] = ")"
	go.setattr[#go.setattr+1] = string.format([[
	default:
		return p.%s.SetAttr(attr,value)
	}
	return false
}]],base)
	go.getattr[#go.getattr+1] = string.format([[
	default:
		return p.%s.Attr(attr)
	}
	return nil
}]],base)
	
	c.enums[#c.enums+1] = "    _ID_"..name_tag.."_LAST"
	c.enums[#c.enums+1] = "};"	

	if item.expands ~= nil then
		go.funcs[#go.funcs+1] = item.expands
	end
	
	c_def.enums[#c_def.enums+1] = table.concat(c.enums,"\n")
	if cdrv.inc ~= nil then
		c_def.include[#c_def.include+1] = "#include "..cdrv.inc
	end

	go_def.enums[#go_def.enums+1] = table.concat(go.enums,"\n")

	go_def.defs[#go_def.defs+1] = table.concat(go.defs,"\n")
	go_def.funcs[#go_def.funcs+1] = table.concat(go.defs,"\n")
	go_def.funcs[#go_def.funcs+1] = table.concat(go.setattr,"\n")
	go_def.funcs[#go_def.funcs+1] = table.concat(go.getattr,"\n")
	go_def.funcs[#go_def.funcs+1] = table.concat(go.funcs,"\n")
	
	c.funcs[#c.funcs+1] = [[
    default:
        return 0;
    }
    return 1;
}
]]
	c_def.funcs[#c_def.funcs+1] = table.concat(c.funcs,"\n")
	c_def.drvfunc[#c_def.drvfunc+1] = string.format([[
    case %s:
        return drv_%s(drvid,a0,a1,a2,a3,a4,a5,a6);]],name_cls,string.lower(name))
end
