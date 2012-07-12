module("baseslider")

name = "baseSlider"
base = "Widget"

funcs = [[
@ SetTracking(b bool)
@ HasTracking() (b bool)
@ SetMaximum(value int)
@ Maximum() (value int)
@ SetMinimum(value int)
@ Minimum() (value int)
@ SetOrientation(value Orientation)
@ Orientation() (value Orientation)
@ SetPageStep(value int)
@ PageStep() (value int)
@ SetSingleStep(value int)
@ SingleStep() (value int)
@ SetSliderDown(b bool)
@ IsSliderDown() (b bool)
@ SetSliderPosition(value int)
@ SliderPosition() (value int)
@ SetValue(value int)
@ Value() (value int)
SetRange(min int, max int)
Range() (min int, max int)

* OnValueChanged(fn func(int))
* OnSliderPressed(fn func())
* OnSliderReleased(fn func())
* OnSliderMoved(fn func(int))
]]


qtdrv = {
inc = "<QAbstractSlider>",
name = "QAbstractSlider *",

SetTracking = "setTracking",
HasTracking = "hasTracking",
SetMaximum = "setMaximum",
Maximum = "maximum",
SetMinimum = "setMinimum",
Minimum = "minimum",
SetOrientation = "setOrientation",
Orientation = "orientation",
SetPageStep = "setPageStep",
PageStep = "pageStep",
SetSingleStep = "setSingleStep",
SingleStep = "singleStep",
SetSliderDown = "setSliderDown",
IsSliderDown = "isSliderDown",
SetSliderPosition = "setSliderPosition",
SliderPosition = "sliderPosition",
SetValue = "setValue",
Value = "value",
SetRange = "setRange",
Range = [[
drvSetInt(a1,self->minimum());
drvSetInt(a2,self->maximum());
]],

OnValueChanged = [[
QObject::connect(self,SIGNAL(valueChanged(int)),drvNewSignal(self,a1,a2),SLOT(call(int)));
]],

OnSliderPressed = [[
QObject::connect(self,SIGNAL(sliderPressed()),drvNewSignal(self,a1,a2),SLOT(call()));
]],
OnSliderReleased = [[
QObject::connect(self,SIGNAL(sliderReleased()),drvNewSignal(self,a1,a2),SLOT(call()));
]],

OnSliderMoved = [[
QObject::connect(self,SIGNAL(sliderMoved()),drvNewSignal(self,a1,a2),SLOT(call(int)));
]],

}
