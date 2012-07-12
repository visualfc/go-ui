// Copyright 2012 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

type Event interface {
	Accept()
	Ignore()
	IsAccept() bool
}

type BaseEvent struct {
	accept int
}

func (p *BaseEvent) Accept() {
	p.accept = 1
}

func (p *BaseEvent) Ignore() {
	p.accept = 0
}

func (p *BaseEvent) IsAccept() bool {
	return p.accept != 0
}

type ResizeEvent struct {
	BaseEvent
	w, h, ow, oh int
}

func (p *ResizeEvent) Size() Size {
	return Sz(p.w, p.h)
}

func (p *ResizeEvent) OldSize() Size {
	return Sz(p.ow, p.oh)
}

type MoveEvent struct {
	BaseEvent
	x, y   int
	ox, oy int
}

func (p *MoveEvent) Pos() Point {
	return Pt(p.x, p.y)
}

func (p *MoveEvent) OldPos() Point {
	return Pt(p.ox, p.oy)
}

type HoverEvent struct {
	BaseEvent
	x, y   int
	ox, oy int
}

func (p *HoverEvent) Pos() Point {
	return Pt(p.x, p.y)
}

func (p *HoverEvent) OldPos() Point {
	return Pt(p.ox, p.oy)
}

type ShowEvent struct {
	BaseEvent
}

type HideEvent struct {
	BaseEvent
}

type CloseEvent struct {
	BaseEvent
}

type EnterEvent struct {
	BaseEvent
}

type LeaveEvent struct {
	BaseEvent
}

type TimerEvent struct {
	BaseEvent
	timerid int
}

func (p *TimerEvent) TimerId() int {
	return p.timerid
}

type MouseEvent struct {
	BaseEvent
	modifiers int
	button    int
	buttons   int
	gx, gy    int
	x, y      int
}

func (p *MouseEvent) Modifiers() int {
	return p.modifiers
}

func (p *MouseEvent) Button() int {
	return p.button
}

func (p *MouseEvent) Buttons() int {
	return p.buttons
}

func (p *MouseEvent) GlobalPos() Point {
	return Pt(p.x, p.y)
}

func (p *MouseEvent) Pos() Point {
	return Pt(p.x, p.y)
}

type PaintEvent struct {
	BaseEvent
	x, y int
	w, h int
}

func (p *PaintEvent) Rect() Rect {
	return Rc(p.x, p.y, p.w, p.h)
}

type KeyEvent struct {
	BaseEvent
	modifiers        int
	count            int
	autorepeat       int
	key              int
	nativeModifiers  uint32
	nativeScanCode   uint32
	nativeVirtualKey uint32
	text             *utf8_info
}

func (p *KeyEvent) Count() int {
	return p.count
}

func (p *KeyEvent) Modifiers() int {
	return p.modifiers
}

func (p *KeyEvent) Key() int {
	return p.key
}

func (p *KeyEvent) IsAutoRepeat() bool {
	return p.autorepeat != 0
}

func (p *KeyEvent) NativeModifiers() uint32 {
	return p.nativeModifiers
}

func (p *KeyEvent) NativeScanCode() uint32 {
	return p.nativeScanCode
}

func (p *KeyEvent) NativeVirtualKey() uint32 {
	return p.nativeVirtualKey
}

func (p *KeyEvent) Text() string {
	return p.text.String()
}

type FocusEvent struct {
	BaseEvent
	reason int
}

func (p *FocusEvent) Reason() int {
	return p.reason
}
