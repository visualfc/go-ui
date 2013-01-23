// Copyright 2012 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"syscall"
)

var (
	EINVAL error = syscall.EINVAL
)

type Alignment int

const (
	AlignLeft            Alignment = 0x0001
	AlignRight           Alignment = 0x0002
	AlignHCenter         Alignment = 0x0004
	AlignJustify         Alignment = 0x0008
	AlignAbsolute        Alignment = 0x0010
	AlignHorizontal_Mask Alignment = AlignLeft | AlignRight | AlignHCenter | AlignJustify | AlignAbsolute

	AlignTop           Alignment = 0x0020
	AlignBottom        Alignment = 0x0040
	AlignVCenter       Alignment = 0x0080
	AlignVertical_Mask Alignment = AlignTop | AlignBottom | AlignVCenter

	AlignCenter Alignment = AlignVCenter | AlignHCenter
)

type Orientation int

const (
	Horizontal Orientation = 0x1
	Vertical   Orientation = 0x2
)

type TickPosition int

const (
	NoTicks TickPosition = iota
	TicksAbove
	TicksBelow
	TicksBothSides
)

type DockWidgetArea int

const (
	NoDockWidgetArea     DockWidgetArea = 0
	LeftDockWidgetArea   DockWidgetArea = 0x01
	RightDockWidgetArea  DockWidgetArea = 0x02
	TopDockWidgetArea    DockWidgetArea = 0x04
	BottomDockWidgetArea DockWidgetArea = 0x08
	AllDockWidgetAreas   DockWidgetArea = LeftDockWidgetArea | RightDockWidgetArea | TopDockWidgetArea | BottomDockWidgetArea
)

type ItemFlag int

const (
	NoItemFlags ItemFlag = iota
	ItemIsSelectable
	ItemIsEditable
	ItemIsDragEnabled
	ItemIsDropEnabled
	ItemIsUserCheckable
	ItemIsEnabled
	ItemIsTristate
)

const (
	Unchecked        = 0
	PartiallyChecked = 1
	Checked          = 2
)

const (
	PlainText = 0
	RichText  = 1
	AutoText  = 2
	LogText   = 3
)

const (
	LightWeight    = 25
	NormalWeight   = 50
	DemiBoldWeight = 63
	BoldWeight     = 75
	BlackWeight    = 87
)

const (
	NoModifier          = 0x00000000
	ShiftModifier       = 0x02000000
	ControlModifier     = 0x04000000
	AltModifier         = 0x08000000
	MetaModifier        = 0x10000000
	KeypadModifier      = 0x20000000
	GroupSwitchModifier = 0x40000000
	// Do not extend the mask to include 0x01000000
	KeyboardModifierMask = 0xfe000000
)

//shorter names for shortcuts
const (
	META          = MetaModifier
	SHIFT         = ShiftModifier
	CTRL          = ControlModifier
	ALT           = AltModifier
	MODIFIER_MASK = KeyboardModifierMask
	UNICODE_ACCEL = 0x00000000
)

const (
	NoButton        = 0x00000000
	LeftButton      = 0x00000001
	RightButton     = 0x00000002
	MidButton       = 0x00000004 // ### Qt 5: remove me
	MiddleButton    = MidButton
	XButton1        = 0x00000008
	XButton2        = 0x00000010
	MouseButtonMask = 0x000000ff
)

//enum MessageIcon { NoIcon, Information, Warning, Critical };
type MessageIconType int

const (
	NoIcon MessageIconType = iota
	Information
	Warning
	Critical
)

type ToolButtonStyle int

const (
	ToolButtonIconOnly ToolButtonStyle = iota
	ToolButtonTextOnly
	ToolButtonTextBesideIcon
	ToolButtonTextUnderIcon
	ToolButtonFollowStyle
)

type ToolButtonPopupMode int

const (
	ToolButtonDelayedPopup ToolButtonPopupMode = iota
	ToolButtonMenuButtonPopup
	ToolButtonInstantPopup
)

type PenStyle int

const (
	NoPen PenStyle = iota
	SolidLine
	DashLine
	DotLine
	DashDotLine
	DashDotDotLine
	CustomDashLine
)

type BrushStyle int

const (
	NoBrush BrushStyle = iota
	SolidPattern
	Dense1Pattern
	Dense2Pattern
	Dense3Pattern
	Dense4Pattern
	Dense5Pattern
	Dense6Pattern
	Dense7Pattern
	HorPattern
	VerPattern
	CrossPattern
	BDiagPattern
	FDiagPattern
	DiagCrossPattern
	LinearGradientPattern
	RadialGradientPattern
	ConicalGradientPattern
	TexturePattern BrushStyle = 24
)

type TransformationMode int

const (
	FastTransformation TransformationMode = iota
	SmoothTransformation
)

type AspectRatioMode int

const (
	IgnoreAspectRatio AspectRatioMode = iota
	KeepAspectRatio
	KeepAspectRatioByExpanding
)

type SizePolicyPolicyFlag int

const (
	GrowFlag   SizePolicyPolicyFlag = 1
	ExpandFlag                      = 2
	ShrinkFlag                      = 4
	IgnoreFlag                      = 8
)

type SizePolicyPolicy int

const (
	Fixed            SizePolicyPolicy = 0
	Minimum                           = SizePolicyPolicy(GrowFlag)
	Maximum                           = SizePolicyPolicy(ShrinkFlag)
	Preferred                         = SizePolicyPolicy(GrowFlag | ShrinkFlag)
	Expanding                         = SizePolicyPolicy(GrowFlag | ShrinkFlag | ExpandFlag)
	MinimumExpanding                  = SizePolicyPolicy(GrowFlag | ExpandFlag)
	Ignored                           = SizePolicyPolicy(ShrinkFlag | GrowFlag | IgnoreFlag)
)

type SizePolicyControlType int

const (
	ControlTypeDefaultType SizePolicyControlType = 0x00000001
	ControlTypeButtonBox                         = 0x00000002
	ControlTypeCheckBox                          = 0x00000004
	ControlTypeComboBox                          = 0x00000008
	ControlTypeFrame                             = 0x00000010
	ControlTypeGroupBox                          = 0x00000020
	ControlTypeLabel                             = 0x00000040
	ControlTypeLine                              = 0x00000080
	ControlTypeLineEdit                          = 0x00000100
	ControlTypePushButton                        = 0x00000200
	ControlTypeRadioButton                       = 0x00000400
	ControlTypeSlider                            = 0x00000800
	ControlTypeSpinBox                           = 0x00001000
	ControlTypeTabWidget                         = 0x00002000
	ControlTypeToolButton                        = 0x00004000
)

type Point struct {
	X, Y int
}

func (p Point) Unpack() (int, int) {
	return p.X, p.Y
}

type PointF struct {
	X, Y float64
}

func (p PointF) Unpack() (float64, float64) {
	return p.X, p.Y
}

type Size struct {
	Width, Height int
}

func (p Size) Unpack() (int, int) {
	return p.Width, p.Height
}

type SizeF struct {
	Width, Height float64
}

func (p SizeF) Unpack() (float64, float64) {
	return p.Width, p.Height
}

type Rect struct {
	X, Y          int
	Width, Height int
}

func (p Rect) Unpack() (int, int, int, int) {
	return p.X, p.Y, p.Width, p.Height
}

func (p Rect) Point() Point {
	return Point{p.X, p.Y}
}

func (p Rect) Size() Size {
	return Size{p.Width, p.Height}
}

type RectF struct {
	X, Y          float64
	Width, Height float64
}

func (p RectF) Unpack() (float64, float64, float64, float64) {
	return p.X, p.Y, p.Width, p.Height
}

func (p RectF) PointF() PointF {
	return PointF{p.X, p.Y}
}

func (p RectF) SizeF() SizeF {
	return SizeF{p.Width, p.Height}
}

type Margins struct {
	Left, Top, Right, Bottom int
}

func (p Margins) Unpack() (int, int, int, int) {
	return p.Left, p.Top, p.Right, p.Bottom
}

func Pt(x, y int) Point {
	return Point{x, y}
}

func PtF(x, y float64) PointF {
	return PointF{x, y}
}

func Sz(w, h int) Size {
	return Size{w, h}
}

func SzF(w, h float64) SizeF {
	return SizeF{w, h}
}

func Rc(x, y, w, h int) Rect {
	return Rect{x, y, w, h}
}

func RcF(x, y, w, h float64) RectF {
	return RectF{x, y, w, h}
}

func Mr(left, right, top, bottom int) Margins {
	return Margins{left, right, top, bottom}
}
