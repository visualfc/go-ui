// Copyright 2012 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

/*
#include <stdlib.h>
extern int drv(int id, int action, void *exp, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6);

static void init_callback(int classid, int drvid)
{
	extern int drv_callback(void*,void*,void*,void*,void*);
	drv(classid,drvid,&drv_callback,0,0,0,0,0,0,0);
}

static void init_appmain(int classid, int drvid)
{
	extern void drv_appmain();
	drv(classid,drvid,&drv_appmain,0,0,0,0,0,0,0);
}

static void init_result(int classid, int drvid)
{
	extern int drv_result(void*,int);
	drv(classid,drvid,&drv_result,0,0,0,0,0,0,0);
}

static void init_utf8_info(int classid, int drvid)
{
	extern void utf8_info_copy(void*,void*,int);
	drv(classid,drvid,&utf8_info_copy,0,0,0,0,0,0,0);
}

*/
// #cgo LDFLAGS: -L../lib -lqtdrv
import "C"
import "unsafe"
import "fmt"
import "os"
import "reflect"
import "image/color"

type string_info struct {
	Data uintptr
	Len  int
}

type rgba uint32

func (c rgba) RGBA() (r, g, b, a uint32) {
	return uint32((c >> 16) & 0xff), uint32((c >> 8) & 0xff), uint32(c & 0xff), uint32(c >> 24)
}

func make_rgba(c color.Color) rgba {
	if c == nil {
		return 0
	}
	r, g, b, a := c.RGBA()
	return rgba(((a & 0xff) << 24) | ((r & 0xff) << 16) | ((g & 0xff) << 8) | (b & 0xff))
}

type slice_info struct {
	Data uintptr
	Len  int
	Cap  int
}

type utf8_info struct {
	data []byte
}

func (d utf8_info) String() string {
	return string(d.data)
}

//export utf8_info_copy
func utf8_info_copy(p unsafe.Pointer, data unsafe.Pointer, size C.int) {
	((*utf8_info)(p)).data = C.GoBytes(data, size)
}

func _b(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func _p(a interface{}) uintptr {
	if a == nil {
		return 0
	}
	switch v := a.(type) {
	case bool:
		return uintptr(unsafe.Pointer(&v))
	case uintptr:
		return uintptr(unsafe.Pointer(&v))
	case *uintptr:
		return uintptr(unsafe.Pointer(v))
	case int:
		return uintptr(unsafe.Pointer(&v))
	case *int:
		return uintptr(unsafe.Pointer(v))
	case uint:
		return uintptr(unsafe.Pointer(&v))
	case *uint:
		return uintptr(unsafe.Pointer(v))
	case float32:
		return uintptr(unsafe.Pointer(&v))
	case *float32:
		return uintptr(unsafe.Pointer(v))
	case float64:
		return uintptr(unsafe.Pointer(&v))
	case *float64:
		return uintptr(unsafe.Pointer(v))
	case string:
		return uintptr(unsafe.Pointer((*string_info)(unsafe.Pointer(&v))))
	case rgba:
		return uintptr(unsafe.Pointer(&v))
	case *rgba:
		return uintptr(unsafe.Pointer(v))
	case color.Color:
		clr := make_rgba(v)
		return uintptr(unsafe.Pointer(&clr))
	case *utf8_info:
		return uintptr(unsafe.Pointer(v))
	case *obj_info:
		return uintptr(unsafe.Pointer(v))
	case iobj:
		return uintptr(unsafe.Pointer(v.info()))
	case *iobj:
		return uintptr(unsafe.Pointer((*v).info()))
	case Point:
		return uintptr(unsafe.Pointer(&v))
	case *Point:
		return uintptr(unsafe.Pointer(v))
	case PointF:
		return uintptr(unsafe.Pointer(&v))
	case *PointF:
		return uintptr(unsafe.Pointer(v))
	case Size:
		return uintptr(unsafe.Pointer(&v))
	case *Size:
		return uintptr(unsafe.Pointer(v))
	case SizeF:
		return uintptr(unsafe.Pointer(&v))
	case *SizeF:
		return uintptr(unsafe.Pointer(v))
	case Rect:
		return uintptr(unsafe.Pointer(&v))
	case *Rect:
		return uintptr(unsafe.Pointer(v))
	case RectF:
		return uintptr(unsafe.Pointer(&v))
	case *RectF:
		return uintptr(unsafe.Pointer(v))
	case Margins:
		return uintptr(unsafe.Pointer(&v))
	case *Margins:
		return uintptr(unsafe.Pointer(v))
	case []Point:
		return uintptr(unsafe.Pointer((*slice_info)(unsafe.Pointer(&v))))
	case []Rect:
		return uintptr(unsafe.Pointer((*slice_info)(unsafe.Pointer(&v))))
	case []byte:
		return uintptr(unsafe.Pointer((*slice_info)(unsafe.Pointer(&v))))
	case Font:
		return uintptr(unsafe.Pointer(&v))
	case *Font:
		return uintptr(unsafe.Pointer(v))
	case *Icon:
		return uintptr(unsafe.Pointer(v))
	case *Pixmap:
		return uintptr(unsafe.Pointer(v))
	case *Image:
		return uintptr(unsafe.Pointer(v))
	case *Pen:
		return uintptr(unsafe.Pointer(v))
	case *Brush:
		return uintptr(unsafe.Pointer(v))
	case MessageIconType:
		return uintptr(unsafe.Pointer(&v))
	case Alignment:
		return uintptr(unsafe.Pointer(&v))
	case *Alignment:
		return uintptr(unsafe.Pointer(v))
	case Orientation:
		return uintptr(unsafe.Pointer(&v))
	case *Orientation:
		return uintptr(unsafe.Pointer(v))
	case TickPosition:
		return uintptr(unsafe.Pointer(&v))
	case *TickPosition:
		return uintptr(unsafe.Pointer(v))
	case ToolButtonPopupMode:
		return uintptr(unsafe.Pointer(&v))
	case *ToolButtonPopupMode:
		return uintptr(unsafe.Pointer(v))
	case ToolButtonStyle:
		return uintptr(unsafe.Pointer(&v))
	case *ToolButtonStyle:
		return uintptr(unsafe.Pointer(v))
	case DockWidgetArea:
		return uintptr(unsafe.Pointer(&v))
	case *DockWidgetArea:
		return uintptr(unsafe.Pointer(v))
	case PenStyle:
		return uintptr(unsafe.Pointer(&v))
	case *PenStyle:
		return uintptr(unsafe.Pointer(v))
	case BrushStyle:
		return uintptr(unsafe.Pointer(&v))
	case *BrushStyle:
		return uintptr(unsafe.Pointer(v))
	case AspectRatioMode:
		return uintptr(unsafe.Pointer(&v))
	case *AspectRatioMode:
		return uintptr(unsafe.Pointer(v))
	case TransformationMode:
		return uintptr(unsafe.Pointer(&v))
	case *TransformationMode:
		return uintptr(unsafe.Pointer(v))
	default:
		warning("Warning drv, param type \"%s\" not match!", reflect.TypeOf(v))
	}
	return 0
}

//func drv10(id int, act int, a0, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) int {
//	if id <= _CLASSID_APP {
//		return int(C.drv(C.int(id), C.int(act), unsafe.Pointer(a0), unsafe.Pointer(a1), unsafe.Pointer(a2), unsafe.Pointer(a3), unsafe.Pointer(a4), unsafe.Pointer(a5), unsafe.Pointer(a6), unsafe.Pointer(a7), unsafe.Pointer(a8), unsafe.Pointer(a9)))
//	}
//	ch := make(chan int)
//	C.drv(C.int(id), C.int(act), unsafe.Pointer(a0), unsafe.Pointer(a1), unsafe.Pointer(a2), unsafe.Pointer(a3), unsafe.Pointer(a4), unsafe.Pointer(a5), unsafe.Pointer(a6), unsafe.Pointer(&ch), unsafe.Pointer(a8), unsafe.Pointer(a9))
//	<-ch
//	return 0
//}

//func drv(id int, act int, a ...interface{}) int {
//	switch len(a) {
//	case 0:
//		return drv10(id, act, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
//	case 1:
//		return drv10(id, act, _p(a[0]), 0, 0, 0, 0, 0, 0, 0, 0, 0)
//	case 2:
//		return drv10(id, act, _p(a[0]), _p(a[1]), 0, 0, 0, 0, 0, 0, 0, 0)
//	case 3:
//		return drv10(id, act, _p(a[0]), _p(a[1]), _p(a[2]), 0, 0, 0, 0, 0, 0, 0)
//	case 4:
//		return drv10(id, act, _p(a[0]), _p(a[1]), _p(a[2]), _p(a[3]), 0, 0, 0, 0, 0, 0)
//	case 5:
//		return drv10(id, act, _p(a[0]), _p(a[1]), _p(a[2]), _p(a[3]), _p(a[4]), 0, 0, 0, 0, 0)
//	case 6:
//		return drv10(id, act, _p(a[0]), _p(a[1]), _p(a[2]), _p(a[3]), _p(a[4]), _p(a[5]), 0, 0, 0, 0)
//	case 7:
//		return drv10(id, act, _p(a[0]), _p(a[1]), _p(a[2]), _p(a[3]), _p(a[4]), _p(a[5]), _p(a[6]), 0, 0, 0)
//	case 8:
//		return drv10(id, act, _p(a[0]), _p(a[1]), _p(a[2]), _p(a[3]), _p(a[4]), _p(a[5]), _p(a[6]), _p(a[7]), 0, 0)
//	case 9:
//		return drv10(id, act, _p(a[0]), _p(a[1]), _p(a[2]), _p(a[3]), _p(a[4]), _p(a[5]), _p(a[6]), _p(a[7]), _p(a[9]), 0)
//	case 10:
//		return drv10(id, act, _p(a[0]), _p(a[1]), _p(a[2]), _p(a[3]), _p(a[4]), _p(a[5]), _p(a[6]), _p(a[7]), _p(a[9]), _p(a[10]))
//	default:
//		panic("drv param count must <= 10")
//	}
//	return 0
//}

func init() {
	C.init_callback(-1, 0)
	C.init_result(-2, 0)
	C.init_utf8_info(-3, 0)
	C.init_appmain(-4, 0)
}

var func_map = make(map[unsafe.Pointer]interface{})

func _drv(id int, act int, a0, a1, a2, a3, a4, a5, a6, a7, a8, a9 unsafe.Pointer) int {
	return int(C.drv(C.int(id), C.int(act), nil, a0, a1, a2, a3, a4, a5, a6))
}

func _drv_ch(id int, act int, a0, a1, a2, a3, a4, a5, a6, a7, a8, a9 unsafe.Pointer) int {
	ch := make(chan int)
	C.drv(C.int(id), C.int(act), unsafe.Pointer(&ch), a0, a1, a2, a3, a4, a5, a6)
	<-ch
	return 0
}

func _drv_event(id int, event int, obj iobj, fn interface{}) {
	var pfn unsafe.Pointer = unsafe.Pointer(&fn)
	func_map[pfn] = fn
	_drv(id, event, unsafe.Pointer(obj.info()), pfn, nil, nil, nil, nil, nil, nil, nil, nil)
}

func _drv_event_ch(id int, event int, obj iobj, fn interface{}) {
	var pfn unsafe.Pointer = unsafe.Pointer(&fn)
	func_map[pfn] = fn
	_drv_ch(id, event, unsafe.Pointer(obj.info()), pfn, nil, nil, nil, nil, nil, nil, nil, nil)
}

func nativeToObject(native uintptr, classid int) interface{} {
	if native == 0 {
		return nil
	}
	obj := FindObject(native)
	if obj == nil {
		obj = NewObjectWithNative(classid, native)
	}
	return obj
}

//export drv_result
func drv_result(ch unsafe.Pointer, r int) int {
	go func() {
		*(*chan int)(ch) <- r
	}()
	return 0
}

var theApp app
var fnAppMain func()

//export drv_appmain
func drv_appmain() {
	theApp.onRemoveObject(drvRemoveObject)
	registerAllClass()
	defer clearAllObject()
	if fnAppMain != nil {
		fnAppMain()
	}
}

//export drv_callback
func drv_callback(pfn unsafe.Pointer, a1, a2, a3, a4 unsafe.Pointer) int {
	fn, ok := func_map[pfn]
	if !ok {
		return 0
	}
	switch v := (fn).(type) {
	case func():
		v()
	case func(int):
		v(*(*int)(a1))
	case func(uint):
		v(*(*uint)(a1))
	case func(bool):
		v(*(*int)(a1) != 0)
	case func(uintptr):
		v(uintptr(a1))
	case func(string):
		v(((*utf8_info)(a1)).String())
	case func(*Action):
		obj := nativeToObject(uintptr(a1), CLASSID_ACTION)
		var act *Action
		if v1, ok := obj.(*Action); ok {
			act = v1
		}
		v(act)
	case func(*ListWidgetItem):
		obj := nativeToObject(uintptr(a1), CLASSID_LISTWIDGETITEM)
		var item *ListWidgetItem
		if v1, ok := obj.(*ListWidgetItem); ok {
			item = v1
		}
		v(item)
	case func(*ListWidgetItem, *ListWidgetItem):
		obj1 := nativeToObject(uintptr(a1), CLASSID_LISTWIDGETITEM)
		obj2 := nativeToObject(uintptr(a2), CLASSID_LISTWIDGETITEM)
		var item *ListWidgetItem
		var oldItem *ListWidgetItem
		if v1, ok := obj1.(*ListWidgetItem); ok {
			item = v1
		}
		if v2, ok := obj2.(*ListWidgetItem); ok {
			oldItem = v2
		}
		v(item, oldItem)
	case func(*ShowEvent):
		v((*ShowEvent)(a1))
	case func(*HideEvent):
		v((*HideEvent)(a1))
	case func(*CloseEvent):
		v((*CloseEvent)(a1))
	case func(*KeyEvent):
		v((*KeyEvent)(a1))
	case func(*MouseEvent):
		v((*MouseEvent)(a1))
	case func(*MoveEvent):
		v((*MoveEvent)(a1))
	case func(*ResizeEvent):
		v((*ResizeEvent)(a1))
	case func(*EnterEvent):
		v((*EnterEvent)(a1))
	case func(*LeaveEvent):
		v((*LeaveEvent)(a1))
	case func(*FocusEvent):
		v((*FocusEvent)(a1))
	case func(*PaintEvent):
		v((*PaintEvent)(a1))
	case func(*TimerEvent):
		v((*TimerEvent)(a1))
	case func(int,int):
		v(*(*int)(a1),*(*int)(a2))

	default:
		warning("Warning drv_callback, func type \"%s\" not match!", reflect.TypeOf(v))
	}
	return 1
}

func warning(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
