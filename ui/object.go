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


package ui

import (
	"fmt"
)

const (
	drvInvalid uintptr = ^uintptr(0)
)

type IObject interface {
	Close() error
	IsValid() bool
	Name() string
	Native() uintptr
	String() string
	SetAttr(attr string, value interface{}) bool
	Attr(attr string) interface{}
	Key() (string, string)
}

type IWidget interface {
	IObject
	SetLayout(ILayout)
	Layout() ILayout
}

type ILayout interface {
	IObject
	AddWidget(IWidget)
}

type obj_info struct {
	native  uintptr
	classid int
}

type iobj interface {
	info() *obj_info
	setKey(key string, attr string)
}

type object struct {
	obj_info
	key      string
	key_attr string
}

func (p *object) info() *obj_info {
	return &p.obj_info
}

func DumpObject(obj IObject) string {
	return fmt.Sprintf("&{%s{%p}}", obj.Name(), obj)
}

func DumpObjectInfo(obj iobj) string {
	return fmt.Sprintf("&{{%d %d}}", obj.info().native, obj.info().classid)
}

func (p *object) String() string {
	return DumpObject(p)
}

func (p *object) Native() uintptr {
	return p.native
}

func (p *object) Name() string {
	return "object"
}

func (p *object) Close() error {
	p.native = 0
	return nil
}

func (p *object) IsValid() bool {
	return p != nil && p.native != 0
}

func (p *object) setKey(key string, attr string) {
	p.key = key
	p.key_attr = attr
}

func (p *object) Key() (string, string) {
	return p.key, p.key_attr
}

func FindObjectForKey(key string) IObject {
	return key_map[key]
}

func (p *object) SetAttr(attr string, value interface{}) bool {
	return false
}

func (p *object) Attr(attr string) interface{} {
	return nil
}

var obj_map = make(map[uintptr]IObject)
var key_map = make(map[string]IObject)

func SetKey(obj IObject, key string, attr string) {
	k, _ := obj.Key()
	if k != "" && k != key {
		delete(key_map, k)
	}
	key_map[key] = obj
	if v, ok := obj.(iobj); ok {
		v.setKey(key, attr)
	}
}

func Value(key string) interface{} {
	if obj, ok := key_map[key]; ok {
		_, attr := obj.Key()
		return obj.Attr(attr)
	}
	return nil
}

func SetValue(key string, value interface{}) bool {
	if obj, ok := key_map[key]; ok {
		_, attr := obj.Key()
		return obj.SetAttr(attr, value)
	}
	return false
}

var fnOnInsertObject func(interface{})
var fnOnRemoveObject func(interface{})

func OnInsertObject(fn func(interface{})) {
	fnOnInsertObject = fn
}

func OnRemoveObject(fn func(interface{})) {
	fnOnRemoveObject = fn
}

func InsertObject(obj IObject) {
	obj_map[obj.Native()] = obj
	if fnOnInsertObject != nil {
		fnOnInsertObject(obj)
	}
}

func FindObject(native uintptr) interface{} {
	return obj_map[native]
}

func RemoveObject(native uintptr) {
	v := FindObject(native)
	if v == nil {
		return
	}
	if fnOnRemoveObject != nil {
		fnOnRemoveObject(v)
	}
	delete(obj_map, native)
}

func drvRemoveObject(native uintptr) {
	v := FindObject(native)
	if v == nil {
		return
	}
	if fnOnRemoveObject != nil {
		fnOnRemoveObject(v)
	}
	if obj, ok := v.(iobj); ok {
		obj.info().native = 0
	}
	delete(obj_map, native)
}

func clearAllObject() {
	for k, v := range obj_map {
		if fnOnRemoveObject != nil {
			fnOnRemoveObject(v)
		}
		delete(obj_map, k)
		v.Close()
	}
}

var newobj_map = make(map[string](func() IObject))
var newobjid_map = make(map[int](func() IObject))
var newobjnative_map = make(map[int](func(uintptr) IObject))

func RegisterClass(classname string, classid int, fn func() IObject) {
	newobj_map[classname] = fn
	newobjid_map[classid] = fn
}

func RegisterClassNative(classid int, fn func(uintptr) IObject) {
	newobjnative_map[classid] = fn
}

func NewObject(classname string) IObject {
	if fn, ok := newobj_map[classname]; ok {
		return fn()
	}
	return nil
}

func NewObjectWithId(classid int) IObject {
	if fn, ok := newobjid_map[classid]; ok {
		return fn()
	}
	return nil
}

func NewObjectWithNative(classid int, native uintptr) IObject {
	if fn, ok := newobjnative_map[classid]; ok {
		return fn(native)
	}
	return nil
}
