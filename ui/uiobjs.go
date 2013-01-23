// Copyright 2012 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"image/color"
	"runtime"
	"unsafe"
)

// drvclass enums
const (
	CLASSID_NONE = iota
	_CLASSID_APP
	CLASSID_TIMER
	CLASSID_FONT
	CLASSID_PIXMAP
	CLASSID_ICON
	CLASSID_IMAGE
	CLASSID_WIDGET
	CLASSID_ACTION
	CLASSID_ACTIONGROUP
	CLASSID_MENU
	CLASSID_MENUBAR
	CLASSID_TOOLBAR
	CLASSID_STATUSBAR
	CLASSID_DOCKWIDGET
	CLASSID_SYSTEMTRAY
	CLASSID_TABWIDGET
	CLASSID_TOOLBOX
	_CLASSID_BASELAYOUT
	CLASSID_BOXLAYOUT
	CLASSID_HBOXLAYOUT
	CLASSID_VBOXLAYOUT
	CLASSID_STACKEDLAYOUT
	_CLASSID_BASEBUTTON
	CLASSID_BUTTON
	CLASSID_CHECKBOX
	CLASSID_RADIO
	CLASSID_TOOLBUTTON
	CLASSID_FRAME
	CLASSID_LABEL
	CLASSID_GROUPBOX
	CLASSID_DIALOG
	CLASSID_COMBOBOX
	CLASSID_LINEEDIT
	_CLASSID_BASESLIDER
	CLASSID_SLIDER
	CLASSID_SCROLLBAR
	CLASSID_DIAL
	CLASSID_BRUSH
	CLASSID_PEN
	CLASSID_PAINTER
	CLASSID_LISTWIDGETITEM
	CLASSID_LISTWIDGET
	CLASSID_MAINWINDOW
	CLASSID_GLWIDGET
	CLASSID_SIZEPOLICY
	_CLASSID_BASESCROLLAREA
	CLASSID_SCROLLAREA
)

// _CLASSID_APP drvid enums
const (
	_ID_APP_NONE = iota
	_ID_APP_APPMAIN
	_ID_APP_RUN
	_ID_APP_EXIT
	_ID_APP_SETFONT
	_ID_APP_FONT
	_ID_APP_SETACTIVEWINDOW
	_ID_APP_ACTIVEWINDOW
	_ID_APP_SETSTYLESHEET
	_ID_APP_STYLESHEET
	_ID_APP_CLOSEALLWINDOWS
	_ID_APP_ONREMOVEOBJECT
)

// CLASSID_TIMER drvid enums
const (
	_ID_TIMER_NONE = iota
	_ID_TIMER_INIT
	_ID_TIMER_CLOSE
	_ID_TIMER_SETINTERVAL
	_ID_TIMER_INTERVAL
	_ID_TIMER_SETSINGLESHOT
	_ID_TIMER_ISSINGLESHOT
	_ID_TIMER_ISACTIVE
	_ID_TIMER_TIMERID
	_ID_TIMER_START
	_ID_TIMER_STARTWITH
	_ID_TIMER_STOP
	_ID_TIMER_ONTIMEOUT
)

// CLASSID_FONT drvid enums
const (
	_ID_FONT_NONE = iota
	_ID_FONT_INIT
	_ID_FONT_INITWITH
	_ID_FONT_CLOSE
	_ID_FONT_SETFAMILY
	_ID_FONT_FAMILY
	_ID_FONT_SETPOINTSIZE
	_ID_FONT_POINTSIZE
	_ID_FONT_SETPOINTSIZEF
	_ID_FONT_POINTSIZEF
	_ID_FONT_SETBOLD
	_ID_FONT_BOLD
	_ID_FONT_SETITALIC
	_ID_FONT_ITALIC
	_ID_FONT_SETKERNING
	_ID_FONT_KERNING
	_ID_FONT_SETLETTERSPACING
	_ID_FONT_LETTERSPACING
	_ID_FONT_SETOVERLINE
	_ID_FONT_OVERLINE
	_ID_FONT_SETPIXELSIZE
	_ID_FONT_PIXELSIZE
	_ID_FONT_SETRAWMODE
	_ID_FONT_RAWMODE
	_ID_FONT_SETRAWNAME
	_ID_FONT_RAWNAME
	_ID_FONT_SETWEIGHT
	_ID_FONT_WEIGHT
	_ID_FONT_SETFIXEDPITCH
	_ID_FONT_FIXEDPITCH
	_ID_FONT_SETSTRETCH
	_ID_FONT_STRETCH
	_ID_FONT_SETSTRIKEOUT
	_ID_FONT_STRIKEOUT
	_ID_FONT_SETUNDERLINE
	_ID_FONT_UNDERLINE
)

// CLASSID_PIXMAP drvid enums
const (
	_ID_PIXMAP_NONE = iota
	_ID_PIXMAP_INIT
	_ID_PIXMAP_INITWITHFILE
	_ID_PIXMAP_INITWITHDATA
	_ID_PIXMAP_INITWITHSIZE
	_ID_PIXMAP_CLOSE
	_ID_PIXMAP_DEPTH
	_ID_PIXMAP_SIZE
	_ID_PIXMAP_RECT
	_ID_PIXMAP_HASALPHA
	_ID_PIXMAP_HASALPHACHANNEL
	_ID_PIXMAP_ISNULL
	_ID_PIXMAP_WIDTH
	_ID_PIXMAP_HEIGHT
	_ID_PIXMAP_SCALED
	_ID_PIXMAP_SCALEDTOHEIGHT
	_ID_PIXMAP_SCALEDTOWIDTH
	_ID_PIXMAP_TOIMAGE
	_ID_PIXMAP_LOAD
	_ID_PIXMAP_SAVE
	_ID_PIXMAP_FILL
)

// CLASSID_ICON drvid enums
const (
	_ID_ICON_NONE = iota
	_ID_ICON_INIT
	_ID_ICON_INITWITHFILE
	_ID_ICON_INITWITHPIXMAP
	_ID_ICON_CLOSE
)

// CLASSID_IMAGE drvid enums
const (
	_ID_IMAGE_NONE = iota
	_ID_IMAGE_INIT
	_ID_IMAGE_INITWITHFILE
	_ID_IMAGE_INITWITHSIZE
	_ID_IMAGE_CLOSE
	_ID_IMAGE_DEPTH
	_ID_IMAGE_SIZE
	_ID_IMAGE_RECT
	_ID_IMAGE_FILL
	_ID_IMAGE_SCALED
)

// CLASSID_WIDGET drvid enums
const (
	_ID_WIDGET_NONE = iota
	_ID_WIDGET_INIT
	_ID_WIDGET_CLOSE
	_ID_WIDGET_SETAUTOCLOSE
	_ID_WIDGET_AUTOCLOSE
	_ID_WIDGET_SETLAYOUT
	_ID_WIDGET_LAYOUT
	_ID_WIDGET_SETPARENT
	_ID_WIDGET_PARENT
	_ID_WIDGET_SETVISIBLE
	_ID_WIDGET_ISVISIBLE
	_ID_WIDGET_SETWINDOWTITLE
	_ID_WIDGET_WINDOWTITLE
	_ID_WIDGET_SETWINDOWICON
	_ID_WIDGET_WINDOWICON
	_ID_WIDGET_SETPOS
	_ID_WIDGET_POS
	_ID_WIDGET_SETSIZE
	_ID_WIDGET_SIZE
	_ID_WIDGET_SETMINIMUMSIZE
	_ID_WIDGET_MINIMUMSIZE
	_ID_WIDGET_SETMAXIMUMSIZE
	_ID_WIDGET_MAXIMUMSIZE
	_ID_WIDGET_SETFIXEDSIZE
	_ID_WIDGET_SETGEOMETRY
	_ID_WIDGET_GEOMETRY
	_ID_WIDGET_SETFONT
	_ID_WIDGET_FONT
	_ID_WIDGET_SETTOOLTIP
	_ID_WIDGET_TOOLTIP
	_ID_WIDGET_SETUPDATESENABLED
	_ID_WIDGET_UPDATESENABLED
	_ID_WIDGET_ISACTIVATEWINDOW
	_ID_WIDGET_ACTIVATEWINDOW
	_ID_WIDGET_SETSIZEPOLICY
	_ID_WIDGET_SIZEPOLICY
	_ID_WIDGET_DONE
	_ID_WIDGET_UPDATE
	_ID_WIDGET_REPAINT
	_ID_WIDGET_STARTTIMER
	_ID_WIDGET_KILLTIMER
	_ID_WIDGET_ADDACTION
	_ID_WIDGET_INSERTACTION
	_ID_WIDGET_REMOVEACTION
	_ID_WIDGET_ONSHOWEVENT
	_ID_WIDGET_ONHIDEEVENT
	_ID_WIDGET_ONCLOSEEVENT
	_ID_WIDGET_ONKEYPRESSEVENT
	_ID_WIDGET_ONKEYRELEASEEVENT
	_ID_WIDGET_ONMOUSEPRESSEVENT
	_ID_WIDGET_ONMOUSERELEASEEVENT
	_ID_WIDGET_ONMOUSEMOVEEVENT
	_ID_WIDGET_ONMOUSEDOUBLECLICKEVENT
	_ID_WIDGET_ONMOVEEVENT
	_ID_WIDGET_ONPAINTEVENT
	_ID_WIDGET_ONRESIZEEVENT
	_ID_WIDGET_ONENTEREVENT
	_ID_WIDGET_ONLEAVEEVENT
	_ID_WIDGET_ONFOCUSINEVENT
	_ID_WIDGET_ONFOCUSOUTEVENT
	_ID_WIDGET_ONTIMEREVENT
)

// CLASSID_ACTION drvid enums
const (
	_ID_ACTION_NONE = iota
	_ID_ACTION_INIT
	_ID_ACTION_INITWITHTEXT
	_ID_ACTION_CLOSE
	_ID_ACTION_SETTEXT
	_ID_ACTION_TEXT
	_ID_ACTION_SETICON
	_ID_ACTION_ICON
	_ID_ACTION_SETICONTEXT
	_ID_ACTION_ICONTEXT
	_ID_ACTION_SETTOOLTIP
	_ID_ACTION_TOOLTIP
	_ID_ACTION_SETFONT
	_ID_ACTION_FONT
	_ID_ACTION_SETCHECKABLE
	_ID_ACTION_ISCHECKABLE
	_ID_ACTION_SETCHECKED
	_ID_ACTION_ISCHECKED
	_ID_ACTION_SETVISIBLE
	_ID_ACTION_ISVISIBLE
	_ID_ACTION_ONTRIGGERED
)

// CLASSID_ACTIONGROUP drvid enums
const (
	_ID_ACTIONGROUP_NONE = iota
	_ID_ACTIONGROUP_INIT
	_ID_ACTIONGROUP_CLOSE
	_ID_ACTIONGROUP_SETVISIBLE
	_ID_ACTIONGROUP_ISVISIBLE
	_ID_ACTIONGROUP_SETENABLED
	_ID_ACTIONGROUP_ISENABLED
	_ID_ACTIONGROUP_SETEXCLUSIVE
	_ID_ACTIONGROUP_ISEXCLUSIVE
	_ID_ACTIONGROUP_ADDACTION
	_ID_ACTIONGROUP_REMOVEACTION
	_ID_ACTIONGROUP_ONHOVERED
	_ID_ACTIONGROUP_ONTRIGGERED
)

// CLASSID_MENU drvid enums
const (
	_ID_MENU_NONE = iota
	_ID_MENU_INIT
	_ID_MENU_INITWITHTITLE
	_ID_MENU_SETTITLE
	_ID_MENU_TITLE
	_ID_MENU_SETICON
	_ID_MENU_ICON
	_ID_MENU_SETDEFAULTACTION
	_ID_MENU_DEFAULTACTION
	_ID_MENU_SETACTIVEACTION
	_ID_MENU_ACTIVEACTION
	_ID_MENU_MENUACTION
	_ID_MENU_ISEMPTY
	_ID_MENU_ADDMENU
	_ID_MENU_INSERTMENU
	_ID_MENU_ADDSEPARATOR
	_ID_MENU_INSERTSEPARATOR
	_ID_MENU_CLEAR
	_ID_MENU_POPUP
	_ID_MENU_ONHOVERED
	_ID_MENU_ONTRIGGERED
)

// CLASSID_MENUBAR drvid enums
const (
	_ID_MENUBAR_NONE = iota
	_ID_MENUBAR_INIT
	_ID_MENUBAR_SETACTIVEACTION
	_ID_MENUBAR_ACTIVEACTION
	_ID_MENUBAR_SETNATIVEMENUBAR
	_ID_MENUBAR_ISNATIVEMENUBAR
	_ID_MENUBAR_ADDMENU
	_ID_MENUBAR_INSERTMENU
	_ID_MENUBAR_ADDSEPARATOR
	_ID_MENUBAR_INSERTSEPARATOR
	_ID_MENUBAR_CLEAR
	_ID_MENUBAR_ONHOVERED
	_ID_MENUBAR_ONTRIGGERED
)

// CLASSID_TOOLBAR drvid enums
const (
	_ID_TOOLBAR_NONE = iota
	_ID_TOOLBAR_INIT
	_ID_TOOLBAR_INITWITHTITLE
	_ID_TOOLBAR_SETTITLE
	_ID_TOOLBAR_TITLE
	_ID_TOOLBAR_SETICONSIZE
	_ID_TOOLBAR_ICONSIZE
	_ID_TOOLBAR_SETFLOATABLE
	_ID_TOOLBAR_ISFLOATABLE
	_ID_TOOLBAR_ISFLOATING
	_ID_TOOLBAR_SETTOOLBUTTONSTYLE
	_ID_TOOLBAR_TOOLBUTTONSTYLE
	_ID_TOOLBAR_ADDACTION
	_ID_TOOLBAR_ADDSEPARATOR
	_ID_TOOLBAR_ADDWIDGET
	_ID_TOOLBAR_CLEAR
)

// CLASSID_STATUSBAR drvid enums
const (
	_ID_STATUSBAR_NONE = iota
	_ID_STATUSBAR_INIT
	_ID_STATUSBAR_SETSIZEGRIP
	_ID_STATUSBAR_ISSIZEGRIP
	_ID_STATUSBAR_ADDWIDGET
	_ID_STATUSBAR_ADDPERMANENTWIDGET
	_ID_STATUSBAR_INSERTWIDGET
	_ID_STATUSBAR_INSERTPERMANENTWIDGET
	_ID_STATUSBAR_REMOVEWIDGET
	_ID_STATUSBAR_SHOWMESSAGE
	_ID_STATUSBAR_CURRENTMESSAGE
	_ID_STATUSBAR_CLEARMESSAGE
)

// CLASSID_DOCKWIDGET drvid enums
const (
	_ID_DOCKWIDGET_NONE = iota
	_ID_DOCKWIDGET_INIT
	_ID_DOCKWIDGET_INITWITHTITLE
	_ID_DOCKWIDGET_SETDOCK
	_ID_DOCKWIDGET_DOCK
	_ID_DOCKWIDGET_SETTITLEBAR
	_ID_DOCKWIDGET_TITLEBAR
	_ID_DOCKWIDGET_SETFLOATING
	_ID_DOCKWIDGET_ISFLOATING
	_ID_DOCKWIDGET_ONVISIBILITYCHANGED
)

// CLASSID_SYSTEMTRAY drvid enums
const (
	_ID_SYSTEMTRAY_NONE = iota
	_ID_SYSTEMTRAY_INIT
	_ID_SYSTEMTRAY_CLOSE
	_ID_SYSTEMTRAY_SETCONTEXTMENU
	_ID_SYSTEMTRAY_CONTEXTMENU
	_ID_SYSTEMTRAY_SETICON
	_ID_SYSTEMTRAY_ICON
	_ID_SYSTEMTRAY_SETTOOLTIP
	_ID_SYSTEMTRAY_TOOLTIP
	_ID_SYSTEMTRAY_SETVISIBLE
	_ID_SYSTEMTRAY_ISVISIBLE
	_ID_SYSTEMTRAY_SHOWMESSAGE
	_ID_SYSTEMTRAY_ONACTIVATED
	_ID_SYSTEMTRAY_ONMESSAGECLICKED
)

// CLASSID_TABWIDGET drvid enums
const (
	_ID_TABWIDGET_NONE = iota
	_ID_TABWIDGET_INIT
	_ID_TABWIDGET_COUNT
	_ID_TABWIDGET_CURRENTINDEX
	_ID_TABWIDGET_CURRENTWIDGET
	_ID_TABWIDGET_SETCURRENTINDEX
	_ID_TABWIDGET_SETCURRENTWIDGET
	_ID_TABWIDGET_ADDTAB
	_ID_TABWIDGET_INSERTTAB
	_ID_TABWIDGET_REMOVETAB
	_ID_TABWIDGET_CLEAR
	_ID_TABWIDGET_INDEXOF
	_ID_TABWIDGET_WIDGETOF
	_ID_TABWIDGET_SETTABTEXT
	_ID_TABWIDGET_TABTEXT
	_ID_TABWIDGET_SETTABICON
	_ID_TABWIDGET_TABICON
	_ID_TABWIDGET_SETTABTOOLTIP
	_ID_TABWIDGET_TABTOOLTIP
	_ID_TABWIDGET_ONCURRENTCHANGED
)

// CLASSID_TOOLBOX drvid enums
const (
	_ID_TOOLBOX_NONE = iota
	_ID_TOOLBOX_INIT
	_ID_TOOLBOX_SETCURRENTINDEX
	_ID_TOOLBOX_SETCURRENTWIDGET
	_ID_TOOLBOX_CURRENTINDEX
	_ID_TOOLBOX_CURRENTWIDGET
	_ID_TOOLBOX_COUNT
	_ID_TOOLBOX_ADDITEM
	_ID_TOOLBOX_INSERTITEM
	_ID_TOOLBOX_REMOVEITEM
	_ID_TOOLBOX_INDEXOF
	_ID_TOOLBOX_WIDGETOF
	_ID_TOOLBOX_SETITEMTEXT
	_ID_TOOLBOX_ITEMTEXT
	_ID_TOOLBOX_SETITEMICON
	_ID_TOOLBOX_ITEMICON
	_ID_TOOLBOX_SETITEMTOOLTIP
	_ID_TOOLBOX_ITEMTOOLTIP
	_ID_TOOLBOX_ISITEMENABLED
	_ID_TOOLBOX_ONCURRENTCHANGED
)

// _CLASSID_BASELAYOUT drvid enums
const (
	_ID_BASELAYOUT_NONE = iota
	_ID_BASELAYOUT_SETSPACING
	_ID_BASELAYOUT_SPACING
	_ID_BASELAYOUT_SETMARGINS
	_ID_BASELAYOUT_MARGINS
	_ID_BASELAYOUT_SETMENUBAR
	_ID_BASELAYOUT_MENUBAR
	_ID_BASELAYOUT_COUNT
	_ID_BASELAYOUT_ADDLAYOUT
	_ID_BASELAYOUT_ADDWIDGET
	_ID_BASELAYOUT_REMOVEWIDGET
	_ID_BASELAYOUT_INDEXOF
)

// CLASSID_BOXLAYOUT drvid enums
const (
	_ID_BOXLAYOUT_NONE = iota
	_ID_BOXLAYOUT_INIT
	_ID_BOXLAYOUT_CLOSE
	_ID_BOXLAYOUT_SETORIENTATION
	_ID_BOXLAYOUT_ORIENTATION
	_ID_BOXLAYOUT_ADDLAYOUTWITH
	_ID_BOXLAYOUT_ADDWIDGETWITH
	_ID_BOXLAYOUT_ADDSPACING
	_ID_BOXLAYOUT_ADDSTRETCH
)

// CLASSID_HBOXLAYOUT drvid enums
const (
	_ID_HBOXLAYOUT_NONE = iota
	_ID_HBOXLAYOUT_INIT
)

// CLASSID_VBOXLAYOUT drvid enums
const (
	_ID_VBOXLAYOUT_NONE = iota
	_ID_VBOXLAYOUT_INIT
)

// CLASSID_STACKEDLAYOUT drvid enums
const (
	_ID_STACKEDLAYOUT_NONE = iota
	_ID_STACKEDLAYOUT_INIT
	_ID_STACKEDLAYOUT_SETCURRENTINDEX
	_ID_STACKEDLAYOUT_CURRENTINDEX
	_ID_STACKEDLAYOUT_SETCURRENTWIDGET
	_ID_STACKEDLAYOUT_CURRENTWIDGET
	_ID_STACKEDLAYOUT_ADDWIDGET
	_ID_STACKEDLAYOUT_INSERTWIDGET
	_ID_STACKEDLAYOUT_WIDGET
	_ID_STACKEDLAYOUT_ONCURRENTCHANGED
)

// _CLASSID_BASEBUTTON drvid enums
const (
	_ID_BASEBUTTON_NONE = iota
	_ID_BASEBUTTON_SETTEXT
	_ID_BASEBUTTON_TEXT
	_ID_BASEBUTTON_SETICON
	_ID_BASEBUTTON_ICON
	_ID_BASEBUTTON_SETICONSIZE
	_ID_BASEBUTTON_ICONSIZE
	_ID_BASEBUTTON_SETCHECKABLE
	_ID_BASEBUTTON_ISCHECKABLE
	_ID_BASEBUTTON_SETDOWN
	_ID_BASEBUTTON_ISDOWN
)

// CLASSID_BUTTON drvid enums
const (
	_ID_BUTTON_NONE = iota
	_ID_BUTTON_INIT
	_ID_BUTTON_INITWITHTEXT
	_ID_BUTTON_SETFLAT
	_ID_BUTTON_ISFLAT
	_ID_BUTTON_SETDEFAULT
	_ID_BUTTON_ISDEFAULT
	_ID_BUTTON_SETMENU
	_ID_BUTTON_MENU
	_ID_BUTTON_ONCLICKED
)

// CLASSID_CHECKBOX drvid enums
const (
	_ID_CHECKBOX_NONE = iota
	_ID_CHECKBOX_INIT
	_ID_CHECKBOX_INITWITHTEXT
	_ID_CHECKBOX_SETCHECK
	_ID_CHECKBOX_CHECK
	_ID_CHECKBOX_SETTRISTATE
	_ID_CHECKBOX_ISTRISTATE
	_ID_CHECKBOX_ONSTATECHANGED
)

// CLASSID_RADIO drvid enums
const (
	_ID_RADIO_NONE = iota
	_ID_RADIO_INIT
	_ID_RADIO_INITWITHTEXT
	_ID_RADIO_ONCLICKED
)

// CLASSID_TOOLBUTTON drvid enums
const (
	_ID_TOOLBUTTON_NONE = iota
	_ID_TOOLBUTTON_INIT
	_ID_TOOLBUTTON_INITWITHTEXT
	_ID_TOOLBUTTON_SETMENU
	_ID_TOOLBUTTON_MENU
	_ID_TOOLBUTTON_SETAUTORAISE
	_ID_TOOLBUTTON_AUTORAISE
	_ID_TOOLBUTTON_SETTOOLBUTTONSTYLE
	_ID_TOOLBUTTON_TOOLBUTTONSTYLE
	_ID_TOOLBUTTON_SETTOOLBUTTONPOPUPMODE
	_ID_TOOLBUTTON_TOOLBUTTONPOPUPMODE
	_ID_TOOLBUTTON_ONCLICKED
)

// CLASSID_FRAME drvid enums
const (
	_ID_FRAME_NONE = iota
	_ID_FRAME_INIT
	_ID_FRAME_SETFRAMESTYLE
	_ID_FRAME_FRAMESTYLE
	_ID_FRAME_SETFRAMERECT
	_ID_FRAME_FRAMERECT
)

// CLASSID_LABEL drvid enums
const (
	_ID_LABEL_NONE = iota
	_ID_LABEL_INIT
	_ID_LABEL_INITWITHTEXT
	_ID_LABEL_INITWITHPIXMAP
	_ID_LABEL_SETTEXT
	_ID_LABEL_TEXT
	_ID_LABEL_SETWORDWRAP
	_ID_LABEL_WORDWRAP
	_ID_LABEL_SETTEXTFORMAT
	_ID_LABEL_TEXTFORMAT
	_ID_LABEL_SETPIXMAP
	_ID_LABEL_PIXMAP
	_ID_LABEL_SETSCALEDCONTENTS
	_ID_LABEL_HASSCALEDCONTENTS
	_ID_LABEL_SETOPENEXTERNALLINKS
	_ID_LABEL_OPENEXTERNALLINKS
	_ID_LABEL_SETALIGNMENT
	_ID_LABEL_ALIGNMENT
	_ID_LABEL_SETINDENT
	_ID_LABEL_INDENT
	_ID_LABEL_SETMARGIN
	_ID_LABEL_MARGIN
	_ID_LABEL_ONLINKACTIVATED
)

// CLASSID_GROUPBOX drvid enums
const (
	_ID_GROUPBOX_NONE = iota
	_ID_GROUPBOX_INIT
	_ID_GROUPBOX_INITWITHTITLE
	_ID_GROUPBOX_SETTITLE
	_ID_GROUPBOX_TITLE
)

// CLASSID_DIALOG drvid enums
const (
	_ID_DIALOG_NONE = iota
	_ID_DIALOG_INIT
	_ID_DIALOG_SETMODAL
	_ID_DIALOG_ISMODAL
	_ID_DIALOG_SETRESULT
	_ID_DIALOG_RESULT
	_ID_DIALOG_OPEN
	_ID_DIALOG_EXEC
	_ID_DIALOG_DONE
	_ID_DIALOG_ACCEPT
	_ID_DIALOG_REJECT
	_ID_DIALOG_ONACCEPTED
	_ID_DIALOG_ONREJECTED
)

// CLASSID_COMBOBOX drvid enums
const (
	_ID_COMBOBOX_NONE = iota
	_ID_COMBOBOX_INIT
	_ID_COMBOBOX_COUNT
	_ID_COMBOBOX_SETCURRENTINDEX
	_ID_COMBOBOX_CURRENTINDEX
	_ID_COMBOBOX_CURRENTTEXT
	_ID_COMBOBOX_SETEDITABLE
	_ID_COMBOBOX_ISEDITABLE
	_ID_COMBOBOX_SETMAXCOUNT
	_ID_COMBOBOX_MAXCOUNT
	_ID_COMBOBOX_SETMAXVISIBLEITEMS
	_ID_COMBOBOX_MAXVISIBLEITEMS
	_ID_COMBOBOX_SETMINIMUMCONTENTSLENGTH
	_ID_COMBOBOX_MINIMUNCONTENTSLENGHT
	_ID_COMBOBOX_ADDITEM
	_ID_COMBOBOX_INSERTITEM
	_ID_COMBOBOX_REMOVEITEM
	_ID_COMBOBOX_ITEMTEXT
	_ID_COMBOBOX_ONCURRENTINDEXCHANGED
)

// CLASSID_LINEEDIT drvid enums
const (
	_ID_LINEEDIT_NONE = iota
	_ID_LINEEDIT_INIT
	_ID_LINEEDIT_INITWITHTEXT
	_ID_LINEEDIT_SETTEXT
	_ID_LINEEDIT_TEXT
	_ID_LINEEDIT_SETINPUTMASK
	_ID_LINEEDIT_INPUTMASK
	_ID_LINEEDIT_SETALIGNMENT
	_ID_LINEEDIT_ALIGNMENT
	_ID_LINEEDIT_SETCURSORPOS
	_ID_LINEEDIT_CURSORPOS
	_ID_LINEEDIT_SETDRAGENABLED
	_ID_LINEEDIT_DRAGENABLED
	_ID_LINEEDIT_SETREADONLY
	_ID_LINEEDIT_ISREADONLY
	_ID_LINEEDIT_SETFRAME
	_ID_LINEEDIT_HASFRAME
	_ID_LINEEDIT_ISREDOAVAILABLE
	_ID_LINEEDIT_HASSELECTED
	_ID_LINEEDIT_SELECTEDTEXT
	_ID_LINEEDIT_SELSTART
	_ID_LINEEDIT_SETSEL
	_ID_LINEEDIT_CANCELSEL
	_ID_LINEEDIT_SELECTALL
	_ID_LINEEDIT_COPY
	_ID_LINEEDIT_CUT
	_ID_LINEEDIT_PASTE
	_ID_LINEEDIT_REDO
	_ID_LINEEDIT_UNDO
	_ID_LINEEDIT_CLEAR
	_ID_LINEEDIT_ONTEXTCHANGED
	_ID_LINEEDIT_ONEDITINGFINISHED
	_ID_LINEEDIT_ONRETURNPRESSED
)

// _CLASSID_BASESLIDER drvid enums
const (
	_ID_BASESLIDER_NONE = iota
	_ID_BASESLIDER_SETTRACKING
	_ID_BASESLIDER_HASTRACKING
	_ID_BASESLIDER_SETMAXIMUM
	_ID_BASESLIDER_MAXIMUM
	_ID_BASESLIDER_SETMINIMUM
	_ID_BASESLIDER_MINIMUM
	_ID_BASESLIDER_SETORIENTATION
	_ID_BASESLIDER_ORIENTATION
	_ID_BASESLIDER_SETPAGESTEP
	_ID_BASESLIDER_PAGESTEP
	_ID_BASESLIDER_SETSINGLESTEP
	_ID_BASESLIDER_SINGLESTEP
	_ID_BASESLIDER_SETSLIDERDOWN
	_ID_BASESLIDER_ISSLIDERDOWN
	_ID_BASESLIDER_SETSLIDERPOSITION
	_ID_BASESLIDER_SLIDERPOSITION
	_ID_BASESLIDER_SETVALUE
	_ID_BASESLIDER_VALUE
	_ID_BASESLIDER_SETRANGE
	_ID_BASESLIDER_RANGE
	_ID_BASESLIDER_ONVALUECHANGED
	_ID_BASESLIDER_ONSLIDERPRESSED
	_ID_BASESLIDER_ONSLIDERRELEASED
	_ID_BASESLIDER_ONSLIDERMOVED
)

// CLASSID_SLIDER drvid enums
const (
	_ID_SLIDER_NONE = iota
	_ID_SLIDER_INIT
	_ID_SLIDER_SETTICKINTERVAL
	_ID_SLIDER_TICKINTERVAL
	_ID_SLIDER_SETTICKPOSITION
	_ID_SLIDER_TICKPOSITION
)

// CLASSID_SCROLLBAR drvid enums
const (
	_ID_SCROLLBAR_NONE = iota
	_ID_SCROLLBAR_INIT
)

// CLASSID_DIAL drvid enums
const (
	_ID_DIAL_NONE = iota
	_ID_DIAL_INIT
	_ID_DIAL_NOTCHSIZE
	_ID_DIAL_SETNOTCHTARGET
	_ID_DIAL_NOTCHTARGET
	_ID_DIAL_SETNOTCHESVISIBLE
	_ID_DIAL_NOTCHESVISIBLE
	_ID_DIAL_SETWRAPPING
	_ID_DIAL_WRAPPING
)

// CLASSID_BRUSH drvid enums
const (
	_ID_BRUSH_NONE = iota
	_ID_BRUSH_INIT
	_ID_BRUSH_CLOSE
	_ID_BRUSH_SETCOLOR
	_ID_BRUSH_COLOR
	_ID_BRUSH_SETSTYLE
	_ID_BRUSH_STYLE
)

// CLASSID_PEN drvid enums
const (
	_ID_PEN_NONE = iota
	_ID_PEN_INIT
	_ID_PEN_CLOSE
	_ID_PEN_SETCOLOR
	_ID_PEN_COLOR
	_ID_PEN_SETWIDTH
	_ID_PEN_WIDTH
	_ID_PEN_SETSTYLE
	_ID_PEN_STYLE
)

// CLASSID_PAINTER drvid enums
const (
	_ID_PAINTER_NONE = iota
	_ID_PAINTER_INIT
	_ID_PAINTER_INITWITHIMAGE
	_ID_PAINTER_CLOSE
	_ID_PAINTER_INITFROM
	_ID_PAINTER_BEGIN
	_ID_PAINTER_END
	_ID_PAINTER_SETFONT
	_ID_PAINTER_FONT
	_ID_PAINTER_SETPEN
	_ID_PAINTER_PEN
	_ID_PAINTER_SETBRUSH
	_ID_PAINTER_BRUSH
	_ID_PAINTER_DRAWPOINT
	_ID_PAINTER_DRAWPOINTS
	_ID_PAINTER_DRAWLINE
	_ID_PAINTER_DRAWLINES
	_ID_PAINTER_DRAWPOLYLINE
	_ID_PAINTER_DRAWPOLYGON
	_ID_PAINTER_DRAWRECT
	_ID_PAINTER_DRAWRECTS
	_ID_PAINTER_DRAWROUNDEDRECT
	_ID_PAINTER_DRAWELLIPSE
	_ID_PAINTER_DRAWARC
	_ID_PAINTER_DRAWCHORD
	_ID_PAINTER_DRAWPIE
	_ID_PAINTER_DRAWTEXT
	_ID_PAINTER_DRAWTEXTRECT
	_ID_PAINTER_DRAWPIXMAP
	_ID_PAINTER_DRAWPIXMAPEX
	_ID_PAINTER_DRAWPIXMAPRECT
	_ID_PAINTER_DRAWPIXMAPRECTEX
	_ID_PAINTER_DRAWIMAGE
	_ID_PAINTER_DRAWIMAGEEX
	_ID_PAINTER_DRAWIMAGERECT
	_ID_PAINTER_DRAWIMAGERECTEX
	_ID_PAINTER_FILLRECT
	_ID_PAINTER_FILLRECTF
)

// CLASSID_LISTWIDGETITEM drvid enums
const (
	_ID_LISTWIDGETITEM_NONE = iota
	_ID_LISTWIDGETITEM_INIT
	_ID_LISTWIDGETITEM_INITWITHTEXT
	_ID_LISTWIDGETITEM_CLOSE
	_ID_LISTWIDGETITEM_SETTEXT
	_ID_LISTWIDGETITEM_TEXT
	_ID_LISTWIDGETITEM_SETICON
	_ID_LISTWIDGETITEM_ICON
	_ID_LISTWIDGETITEM_SETSELECTED
	_ID_LISTWIDGETITEM_ISSELECTED
	_ID_LISTWIDGETITEM_SETHIDDEN
	_ID_LISTWIDGETITEM_ISHIDDEN
	_ID_LISTWIDGETITEM_SETFONT
	_ID_LISTWIDGETITEM_FONT
	_ID_LISTWIDGETITEM_SETFOREGROUND
	_ID_LISTWIDGETITEM_FOREGROUND
	_ID_LISTWIDGETITEM_SETBACKGROUND
	_ID_LISTWIDGETITEM_BACKGROUND
	_ID_LISTWIDGETITEM_SETTOOLTIP
	_ID_LISTWIDGETITEM_TOOLTIP
	_ID_LISTWIDGETITEM_SETTEXTALIGNMENT
	_ID_LISTWIDGETITEM_TEXTALIGNMENT
	_ID_LISTWIDGETITEM_SETFLAGS
	_ID_LISTWIDGETITEM_FLAGS
)

// CLASSID_LISTWIDGET drvid enums
const (
	_ID_LISTWIDGET_NONE = iota
	_ID_LISTWIDGET_INIT
	_ID_LISTWIDGET_COUNT
	_ID_LISTWIDGET_SETCURRENTITEM
	_ID_LISTWIDGET_CURRENTITEM
	_ID_LISTWIDGET_SETCURRENTROW
	_ID_LISTWIDGET_CURRENTROW
	_ID_LISTWIDGET_ADDITEM
	_ID_LISTWIDGET_INSERTITEM
	_ID_LISTWIDGET_EDITITEM
	_ID_LISTWIDGET_TAKEITEM
	_ID_LISTWIDGET_ITEM
	_ID_LISTWIDGET_CLEAR
	_ID_LISTWIDGET_ONCURRENTITEMCHANGED
	_ID_LISTWIDGET_ONCURRENTROWCHANGED
	_ID_LISTWIDGET_ONITEMACTIVATED
	_ID_LISTWIDGET_ONITEMCHANGED
	_ID_LISTWIDGET_ONITEMCLICKED
	_ID_LISTWIDGET_ONITEMDOUBLECLICKED
	_ID_LISTWIDGET_ONITEMENTERED
	_ID_LISTWIDGET_ONITEMPRESSED
	_ID_LISTWIDGET_ONITEMSELECTIONCHANGED
)

// CLASSID_MAINWINDOW drvid enums
const (
	_ID_MAINWINDOW_NONE = iota
	_ID_MAINWINDOW_INIT
	_ID_MAINWINDOW_SETCENTRALWIDGET
	_ID_MAINWINDOW_CENTRALWIDGET
	_ID_MAINWINDOW_SETMENUBAR
	_ID_MAINWINDOW_MENUBAR
	_ID_MAINWINDOW_SETSTATUSBAR
	_ID_MAINWINDOW_STATUSBAR
	_ID_MAINWINDOW_ADDTOOLBAR
	_ID_MAINWINDOW_INSERTTOOLBAR
	_ID_MAINWINDOW_REMOVETOOLBAR
	_ID_MAINWINDOW_ADDDOCKWIDGET
	_ID_MAINWINDOW_REMOVEDOCKWIDGET
)

// CLASSID_GLWIDGET drvid enums
const (
	_ID_GLWIDGET_NONE = iota
	_ID_GLWIDGET_INIT
	_ID_GLWIDGET_DELETETEXTURE
	_ID_GLWIDGET_DONECURRENT
	_ID_GLWIDGET_DOUBLEBUFFER
	_ID_GLWIDGET_CONVERTTOGLFORMAT
	_ID_GLWIDGET_SETMOUSETRACKING
	_ID_GLWIDGET_RENDERTEXT
	_ID_GLWIDGET_UPDATEGL
	_ID_GLWIDGET_UPDATEOVERLAYGL
	_ID_GLWIDGET_ONINITIALIZEGL
	_ID_GLWIDGET_ONINITIALIZEOVERLAYGL
	_ID_GLWIDGET_ONPAINTGL
	_ID_GLWIDGET_ONPAINTOVERLAYGL
	_ID_GLWIDGET_ONRESIZEGL
	_ID_GLWIDGET_ONRESIZEOVERLAYGL
)

// CLASSID_SIZEPOLICY drvid enums
const (
	_ID_SIZEPOLICY_NONE = iota
	_ID_SIZEPOLICY_INIT
	_ID_SIZEPOLICY_INITWITHPOLICY
	_ID_SIZEPOLICY_CLOSE
	_ID_SIZEPOLICY_HORIZONTALPOLICY
	_ID_SIZEPOLICY_SETHORIZONTALPOLICY
	_ID_SIZEPOLICY_VERTICALPOLICY
	_ID_SIZEPOLICY_SETVERTICALPOLICY
	_ID_SIZEPOLICY_HASHEIGHTFORWIDTH
	_ID_SIZEPOLICY_TRANSPOSE
)

// _CLASSID_BASESCROLLAREA drvid enums
const (
	_ID_BASESCROLLAREA_NONE = iota
	_ID_BASESCROLLAREA_CORNERWIDGET
	_ID_BASESCROLLAREA_HORIZONTALSCROLLBAR
	_ID_BASESCROLLAREA_VERTICALSCROLLBAR
	_ID_BASESCROLLAREA_VIEWPORT
)

// CLASSID_SCROLLAREA drvid enums
const (
	_ID_SCROLLAREA_NONE = iota
	_ID_SCROLLAREA_INIT
	_ID_SCROLLAREA_SETALIGNMENT
	_ID_SCROLLAREA_ALIGNMENT
	_ID_SCROLLAREA_SETWIDGET
	_ID_SCROLLAREA_WIDGET
	_ID_SCROLLAREA_SETWIDGETRESIZABLE
	_ID_SCROLLAREA_WIDGETRESIZABLE
	_ID_SCROLLAREA_TAKEWIDGET
	_ID_SCROLLAREA_ENSUREVISIBLE
	_ID_SCROLLAREA_ENSUREWIDGETVISIBLE
)

func registerAllClass() {

	RegisterClass("Timer", CLASSID_TIMER, func() IObject {
		return NewTimer()
	})
	RegisterClassNative(CLASSID_TIMER, func(native uintptr) IObject {
		obj := new(Timer)
		obj.classid = CLASSID_TIMER
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("Font", CLASSID_FONT, func() IObject {
		return NewFont()
	})
	RegisterClassNative(CLASSID_FONT, func(native uintptr) IObject {
		obj := new(Font)
		obj.classid = CLASSID_FONT
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("Pixmap", CLASSID_PIXMAP, func() IObject {
		return NewPixmap()
	})
	RegisterClassNative(CLASSID_PIXMAP, func(native uintptr) IObject {
		obj := new(Pixmap)
		obj.classid = CLASSID_PIXMAP
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("Icon", CLASSID_ICON, func() IObject {
		return NewIcon()
	})
	RegisterClassNative(CLASSID_ICON, func(native uintptr) IObject {
		obj := new(Icon)
		obj.classid = CLASSID_ICON
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("Image", CLASSID_IMAGE, func() IObject {
		return NewImage()
	})
	RegisterClassNative(CLASSID_IMAGE, func(native uintptr) IObject {
		obj := new(Image)
		obj.classid = CLASSID_IMAGE
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("Widget", CLASSID_WIDGET, func() IObject {
		return NewWidget()
	})
	RegisterClassNative(CLASSID_WIDGET, func(native uintptr) IObject {
		obj := new(Widget)
		obj.classid = CLASSID_WIDGET
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("Action", CLASSID_ACTION, func() IObject {
		return NewAction()
	})
	RegisterClassNative(CLASSID_ACTION, func(native uintptr) IObject {
		obj := new(Action)
		obj.classid = CLASSID_ACTION
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("ActionGroup", CLASSID_ACTIONGROUP, func() IObject {
		return NewActionGroup()
	})
	RegisterClassNative(CLASSID_ACTIONGROUP, func(native uintptr) IObject {
		obj := new(ActionGroup)
		obj.classid = CLASSID_ACTIONGROUP
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("Menu", CLASSID_MENU, func() IObject {
		return NewMenu()
	})
	RegisterClassNative(CLASSID_MENU, func(native uintptr) IObject {
		obj := new(Menu)
		obj.classid = CLASSID_MENU
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("MenuBar", CLASSID_MENUBAR, func() IObject {
		return NewMenuBar()
	})
	RegisterClassNative(CLASSID_MENUBAR, func(native uintptr) IObject {
		obj := new(MenuBar)
		obj.classid = CLASSID_MENUBAR
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("ToolBar", CLASSID_TOOLBAR, func() IObject {
		return NewToolBar()
	})
	RegisterClassNative(CLASSID_TOOLBAR, func(native uintptr) IObject {
		obj := new(ToolBar)
		obj.classid = CLASSID_TOOLBAR
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("StatusBar", CLASSID_STATUSBAR, func() IObject {
		return NewStatusBar()
	})
	RegisterClassNative(CLASSID_STATUSBAR, func(native uintptr) IObject {
		obj := new(StatusBar)
		obj.classid = CLASSID_STATUSBAR
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("DockWidget", CLASSID_DOCKWIDGET, func() IObject {
		return NewDockWidget()
	})
	RegisterClassNative(CLASSID_DOCKWIDGET, func(native uintptr) IObject {
		obj := new(DockWidget)
		obj.classid = CLASSID_DOCKWIDGET
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("SystemTray", CLASSID_SYSTEMTRAY, func() IObject {
		return NewSystemTray()
	})
	RegisterClassNative(CLASSID_SYSTEMTRAY, func(native uintptr) IObject {
		obj := new(SystemTray)
		obj.classid = CLASSID_SYSTEMTRAY
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("TabWidget", CLASSID_TABWIDGET, func() IObject {
		return NewTabWidget()
	})
	RegisterClassNative(CLASSID_TABWIDGET, func(native uintptr) IObject {
		obj := new(TabWidget)
		obj.classid = CLASSID_TABWIDGET
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("ToolBox", CLASSID_TOOLBOX, func() IObject {
		return NewToolBox()
	})
	RegisterClassNative(CLASSID_TOOLBOX, func(native uintptr) IObject {
		obj := new(ToolBox)
		obj.classid = CLASSID_TOOLBOX
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("BoxLayout", CLASSID_BOXLAYOUT, func() IObject {
		return NewBoxLayout()
	})
	RegisterClassNative(CLASSID_BOXLAYOUT, func(native uintptr) IObject {
		obj := new(BoxLayout)
		obj.classid = CLASSID_BOXLAYOUT
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("HBoxLayout", CLASSID_HBOXLAYOUT, func() IObject {
		return NewHBoxLayout()
	})
	RegisterClassNative(CLASSID_HBOXLAYOUT, func(native uintptr) IObject {
		obj := new(HBoxLayout)
		obj.classid = CLASSID_HBOXLAYOUT
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("VBoxLayout", CLASSID_VBOXLAYOUT, func() IObject {
		return NewVBoxLayout()
	})
	RegisterClassNative(CLASSID_VBOXLAYOUT, func(native uintptr) IObject {
		obj := new(VBoxLayout)
		obj.classid = CLASSID_VBOXLAYOUT
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("StackedLayout", CLASSID_STACKEDLAYOUT, func() IObject {
		return NewStackedLayout()
	})
	RegisterClassNative(CLASSID_STACKEDLAYOUT, func(native uintptr) IObject {
		obj := new(StackedLayout)
		obj.classid = CLASSID_STACKEDLAYOUT
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("Button", CLASSID_BUTTON, func() IObject {
		return NewButton()
	})
	RegisterClassNative(CLASSID_BUTTON, func(native uintptr) IObject {
		obj := new(Button)
		obj.classid = CLASSID_BUTTON
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("CheckBox", CLASSID_CHECKBOX, func() IObject {
		return NewCheckBox()
	})
	RegisterClassNative(CLASSID_CHECKBOX, func(native uintptr) IObject {
		obj := new(CheckBox)
		obj.classid = CLASSID_CHECKBOX
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("Radio", CLASSID_RADIO, func() IObject {
		return NewRadio()
	})
	RegisterClassNative(CLASSID_RADIO, func(native uintptr) IObject {
		obj := new(Radio)
		obj.classid = CLASSID_RADIO
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("ToolButton", CLASSID_TOOLBUTTON, func() IObject {
		return NewToolButton()
	})
	RegisterClassNative(CLASSID_TOOLBUTTON, func(native uintptr) IObject {
		obj := new(ToolButton)
		obj.classid = CLASSID_TOOLBUTTON
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("Frame", CLASSID_FRAME, func() IObject {
		return NewFrame()
	})
	RegisterClassNative(CLASSID_FRAME, func(native uintptr) IObject {
		obj := new(Frame)
		obj.classid = CLASSID_FRAME
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("Label", CLASSID_LABEL, func() IObject {
		return NewLabel()
	})
	RegisterClassNative(CLASSID_LABEL, func(native uintptr) IObject {
		obj := new(Label)
		obj.classid = CLASSID_LABEL
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("GroupBox", CLASSID_GROUPBOX, func() IObject {
		return NewGroupBox()
	})
	RegisterClassNative(CLASSID_GROUPBOX, func(native uintptr) IObject {
		obj := new(GroupBox)
		obj.classid = CLASSID_GROUPBOX
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("Dialog", CLASSID_DIALOG, func() IObject {
		return NewDialog()
	})
	RegisterClassNative(CLASSID_DIALOG, func(native uintptr) IObject {
		obj := new(Dialog)
		obj.classid = CLASSID_DIALOG
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("ComboBox", CLASSID_COMBOBOX, func() IObject {
		return NewComboBox()
	})
	RegisterClassNative(CLASSID_COMBOBOX, func(native uintptr) IObject {
		obj := new(ComboBox)
		obj.classid = CLASSID_COMBOBOX
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("LineEdit", CLASSID_LINEEDIT, func() IObject {
		return NewLineEdit()
	})
	RegisterClassNative(CLASSID_LINEEDIT, func(native uintptr) IObject {
		obj := new(LineEdit)
		obj.classid = CLASSID_LINEEDIT
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("Slider", CLASSID_SLIDER, func() IObject {
		return NewSlider()
	})
	RegisterClassNative(CLASSID_SLIDER, func(native uintptr) IObject {
		obj := new(Slider)
		obj.classid = CLASSID_SLIDER
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("ScrollBar", CLASSID_SCROLLBAR, func() IObject {
		return NewScrollBar()
	})
	RegisterClassNative(CLASSID_SCROLLBAR, func(native uintptr) IObject {
		obj := new(ScrollBar)
		obj.classid = CLASSID_SCROLLBAR
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("Dial", CLASSID_DIAL, func() IObject {
		return NewDial()
	})
	RegisterClassNative(CLASSID_DIAL, func(native uintptr) IObject {
		obj := new(Dial)
		obj.classid = CLASSID_DIAL
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("Brush", CLASSID_BRUSH, func() IObject {
		return NewBrush()
	})
	RegisterClassNative(CLASSID_BRUSH, func(native uintptr) IObject {
		obj := new(Brush)
		obj.classid = CLASSID_BRUSH
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("Pen", CLASSID_PEN, func() IObject {
		return NewPen()
	})
	RegisterClassNative(CLASSID_PEN, func(native uintptr) IObject {
		obj := new(Pen)
		obj.classid = CLASSID_PEN
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("Painter", CLASSID_PAINTER, func() IObject {
		return NewPainter()
	})
	RegisterClassNative(CLASSID_PAINTER, func(native uintptr) IObject {
		obj := new(Painter)
		obj.classid = CLASSID_PAINTER
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("ListWidgetItem", CLASSID_LISTWIDGETITEM, func() IObject {
		return NewListWidgetItem()
	})
	RegisterClassNative(CLASSID_LISTWIDGETITEM, func(native uintptr) IObject {
		obj := new(ListWidgetItem)
		obj.classid = CLASSID_LISTWIDGETITEM
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("ListWidget", CLASSID_LISTWIDGET, func() IObject {
		return NewListWidget()
	})
	RegisterClassNative(CLASSID_LISTWIDGET, func(native uintptr) IObject {
		obj := new(ListWidget)
		obj.classid = CLASSID_LISTWIDGET
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("MainWindow", CLASSID_MAINWINDOW, func() IObject {
		return NewMainWindow()
	})
	RegisterClassNative(CLASSID_MAINWINDOW, func(native uintptr) IObject {
		obj := new(MainWindow)
		obj.classid = CLASSID_MAINWINDOW
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("GLWidget", CLASSID_GLWIDGET, func() IObject {
		return NewGLWidget()
	})
	RegisterClassNative(CLASSID_GLWIDGET, func(native uintptr) IObject {
		obj := new(GLWidget)
		obj.classid = CLASSID_GLWIDGET
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("SizePolicy", CLASSID_SIZEPOLICY, func() IObject {
		return NewSizePolicy()
	})
	RegisterClassNative(CLASSID_SIZEPOLICY, func(native uintptr) IObject {
		obj := new(SizePolicy)
		obj.classid = CLASSID_SIZEPOLICY
		obj.native = native
		InsertObject(obj)
		return obj
	})
	RegisterClass("ScrollArea", CLASSID_SCROLLAREA, func() IObject {
		return NewScrollArea()
	})
	RegisterClassNative(CLASSID_SCROLLAREA, func(native uintptr) IObject {
		obj := new(ScrollArea)
		obj.classid = CLASSID_SCROLLAREA
		obj.native = native
		InsertObject(obj)
		return obj
	})
}

// struct app
//
type app struct {
	object
}

func (p *app) Name() string {
	return "app"
}

func (p *app) String() string {
	return DumpObject(p)
}

func (p *app) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "font":
		if v, ok := value.(*Font); ok {
			p.SetFont(v)
			return true
		}
		return false
	case "activewindow":
		if v, ok := value.(IWidget); ok {
			p.SetActiveWindow(v)
			return true
		}
		return false
	case "stylesheet":
		if v, ok := value.(string); ok {
			p.SetStyleSheet(v)
			return true
		}
		return false
	default:
		return p.object.SetAttr(attr, value)
	}
	return false
}
func (p *app) Attr(attr string) interface{} {
	switch attr {
	case "font":
		return p.Font()
	case "activewindow":
		return p.ActiveWindow()
	case "stylesheet":
		return p.StyleSheet()
	default:
		return p.object.Attr(attr)
	}
	return nil
}
func (p *app) AppMain() (code int) {
	_drv(_CLASSID_APP, _ID_APP_APPMAIN, unsafe.Pointer(p.info()), unsafe.Pointer(&code), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *app) Run() (code int) {
	_drv(_CLASSID_APP, _ID_APP_RUN, unsafe.Pointer(p.info()), unsafe.Pointer(&code), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *app) Exit(code int) {
	_drv(_CLASSID_APP, _ID_APP_EXIT, unsafe.Pointer(p.info()), unsafe.Pointer(&code), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *app) SetFont(font *Font) {
	_drv(_CLASSID_APP, _ID_APP_SETFONT, unsafe.Pointer(p.info()), unsafe.Pointer(font), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *app) Font() (font *Font) {
	var oi_font obj_info
	_drv(_CLASSID_APP, _ID_APP_FONT, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_font), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_font.native != 0 {
		v := FindObject(oi_font.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_FONT, oi_font.native)
		}
		if v != nil {
			font = v.(*Font)
		}
	}
	return
}

func (p *app) SetActiveWindow(widget IWidget) {
	_drv(_CLASSID_APP, _ID_APP_SETACTIVEWINDOW, unsafe.Pointer(p.info()), unsafe.Pointer(widget.(iobj).info()), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *app) ActiveWindow() (widget IWidget) {
	var oi_widget obj_info
	_drv(_CLASSID_APP, _ID_APP_ACTIVEWINDOW, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_widget), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_widget.native != 0 {
		item := FindObject(oi_widget.native)
		if item != nil {
			widget = item.(IWidget)
		}
	}
	return
}

func (p *app) SetStyleSheet(style string) {
	_drv(_CLASSID_APP, _ID_APP_SETSTYLESHEET, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&style))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *app) StyleSheet() (style string) {
	var sh_style utf8_info
	_drv(_CLASSID_APP, _ID_APP_STYLESHEET, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_style))), nil, nil, nil, nil, nil, nil, nil, nil)
	style = sh_style.String()
	return
}

func (p *app) CloseAllWindows() {
	_drv(_CLASSID_APP, _ID_APP_CLOSEALLWINDOWS, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *app) onRemoveObject(fn func(uintptr)) {
	_drv_event(_CLASSID_APP, _ID_APP_ONREMOVEOBJECT, p, fn)
	return
}

// struct Timer
//
type Timer struct {
	object
}

func (p *Timer) Name() string {
	return "Timer"
}

func (p *Timer) String() string {
	return DumpObject(p)
}
func (p *Timer) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "interval":
		if v, ok := value.(int); ok {
			p.SetInterval(v)
			return true
		}
		return false
	case "singleshot":
		if v, ok := value.(bool); ok {
			p.SetSingleShot(v)
			return true
		}
		return false
	default:
		return p.object.SetAttr(attr, value)
	}
	return false
}
func (p *Timer) Attr(attr string) interface{} {
	switch attr {
	case "interval":
		return p.Interval()
	case "singleshot":
		return p.IsSingleShot()
	case "active":
		return p.IsActive()
	case "timerid":
		return p.TimerId()
	default:
		return p.object.Attr(attr)
	}
	return nil
}
func NewTimer() *Timer {
	return new(Timer).Init()
}

func (p *Timer) Init() *Timer {
	p.classid = CLASSID_TIMER
	_drv_ch(CLASSID_TIMER, _ID_TIMER_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	runtime.SetFinalizer(p, (*Timer).Close)
	return p
}

func (p *Timer) Close() (err error) {
	if p == nil || p.native == 0 {
		return
	}
	_drv_ch(CLASSID_TIMER, _ID_TIMER_CLOSE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	p.native = 0
	runtime.SetFinalizer(p, nil)
	return
}

func (p *Timer) SetInterval(msec int) {
	_drv_ch(CLASSID_TIMER, _ID_TIMER_SETINTERVAL, unsafe.Pointer(p.info()), unsafe.Pointer(&msec), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Timer) Interval() (msec int) {
	_drv_ch(CLASSID_TIMER, _ID_TIMER_INTERVAL, unsafe.Pointer(p.info()), unsafe.Pointer(&msec), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Timer) SetSingleShot(b bool) {
	_drv_ch(CLASSID_TIMER, _ID_TIMER_SETSINGLESHOT, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Timer) IsSingleShot() (b bool) {
	var b_b int
	_drv_ch(CLASSID_TIMER, _ID_TIMER_ISSINGLESHOT, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Timer) IsActive() (b bool) {
	var b_b int
	_drv_ch(CLASSID_TIMER, _ID_TIMER_ISACTIVE, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Timer) TimerId() (id int) {
	_drv_ch(CLASSID_TIMER, _ID_TIMER_TIMERID, unsafe.Pointer(p.info()), unsafe.Pointer(&id), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Timer) Start() {
	_drv_ch(CLASSID_TIMER, _ID_TIMER_START, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Timer) StartWith(msec int) {
	_drv_ch(CLASSID_TIMER, _ID_TIMER_STARTWITH, unsafe.Pointer(p.info()), unsafe.Pointer(&msec), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Timer) Stop() {
	_drv_ch(CLASSID_TIMER, _ID_TIMER_STOP, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Timer) OnTimeout(fn func()) {
	_drv_event_ch(CLASSID_TIMER, _ID_TIMER_ONTIMEOUT, p, fn)
	return
}

func AfterFunc(ms int, fn func()) *Timer {
	timer := NewTimer()
	timer.SetSingleShot(true)
	timer.StartWith(ms)
	timer.OnTimeout(func() {
		timer.Close()
		fn()
	})
	return timer
}

// struct Font
//
type Font struct {
	object
}

func (p *Font) Name() string {
	return "Font"
}

func (p *Font) String() string {
	return DumpObject(p)
}
func (p *Font) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "family":
		if v, ok := value.(string); ok {
			p.SetFamily(v)
			return true
		}
		return false
	case "pointsize":
		if v, ok := value.(int); ok {
			p.SetPointSize(v)
			return true
		}
		return false
	case "pointsizef":
		if v, ok := value.(float64); ok {
			p.SetPointSizeF(v)
			return true
		}
		return false
	case "bold":
		if v, ok := value.(bool); ok {
			p.SetBold(v)
			return true
		}
		return false
	case "italic":
		if v, ok := value.(bool); ok {
			p.SetItalic(v)
			return true
		}
		return false
	case "kerning":
		if v, ok := value.(bool); ok {
			p.SetKerning(v)
			return true
		}
		return false
	case "overline":
		if v, ok := value.(bool); ok {
			p.SetOverline(v)
			return true
		}
		return false
	case "pixelsize":
		if v, ok := value.(int); ok {
			p.SetPixelSize(v)
			return true
		}
		return false
	case "rawmode":
		if v, ok := value.(bool); ok {
			p.SetRawMode(v)
			return true
		}
		return false
	case "rawname":
		if v, ok := value.(string); ok {
			p.SetRawName(v)
			return true
		}
		return false
	case "weight":
		if v, ok := value.(int); ok {
			p.SetWeight(v)
			return true
		}
		return false
	case "fixedpitch":
		if v, ok := value.(bool); ok {
			p.SetFixedPitch(v)
			return true
		}
		return false
	case "stretch":
		if v, ok := value.(int); ok {
			p.SetStretch(v)
			return true
		}
		return false
	case "strikeout":
		if v, ok := value.(bool); ok {
			p.SetStrikeOut(v)
			return true
		}
		return false
	case "underline":
		if v, ok := value.(bool); ok {
			p.SetUnderline(v)
			return true
		}
		return false
	default:
		return p.object.SetAttr(attr, value)
	}
	return false
}
func (p *Font) Attr(attr string) interface{} {
	switch attr {
	case "family":
		return p.Family()
	case "pointsize":
		return p.PointSize()
	case "pointsizef":
		return p.PointSizeF()
	case "bold":
		return p.Bold()
	case "italic":
		return p.Italic()
	case "kerning":
		return p.Kerning()
	case "overline":
		return p.Overline()
	case "pixelsize":
		return p.PixelSize()
	case "rawmode":
		return p.RawMode()
	case "rawname":
		return p.RawName()
	case "weight":
		return p.Weight()
	case "fixedpitch":
		return p.FixedPitch()
	case "stretch":
		return p.Stretch()
	case "strikeout":
		return p.StrikeOut()
	case "underline":
		return p.Underline()
	default:
		return p.object.Attr(attr)
	}
	return nil
}
func NewFont() *Font {
	return new(Font).Init()
}

func (p *Font) Init() *Font {
	p.classid = CLASSID_FONT
	_drv_ch(CLASSID_FONT, _ID_FONT_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewFontWith(family string, size int, weight int) *Font {
	return new(Font).InitWith(family, size, weight)
}

func (p *Font) InitWith(family string, size int, weight int) *Font {
	p.classid = CLASSID_FONT
	_drv_ch(CLASSID_FONT, _ID_FONT_INITWITH, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&family))), unsafe.Pointer(&size), unsafe.Pointer(&weight), nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *Font) Close() (err error) {
	if p == nil || p.native == 0 {
		return
	}
	_drv_ch(CLASSID_FONT, _ID_FONT_CLOSE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	RemoveObject(p.native)
	p.native = 0
	return
}

func (p *Font) SetFamily(text string) {
	_drv_ch(CLASSID_FONT, _ID_FONT_SETFAMILY, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) Family() (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_FONT, _ID_FONT_FAMILY, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *Font) SetPointSize(size int) {
	_drv_ch(CLASSID_FONT, _ID_FONT_SETPOINTSIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&size), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) PointSize() (size int) {
	_drv_ch(CLASSID_FONT, _ID_FONT_POINTSIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&size), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) SetPointSizeF(size float64) {
	_drv_ch(CLASSID_FONT, _ID_FONT_SETPOINTSIZEF, unsafe.Pointer(p.info()), unsafe.Pointer(&size), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) PointSizeF() (size float64) {
	_drv_ch(CLASSID_FONT, _ID_FONT_POINTSIZEF, unsafe.Pointer(p.info()), unsafe.Pointer(&size), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) SetBold(b bool) {
	_drv_ch(CLASSID_FONT, _ID_FONT_SETBOLD, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) Bold() (b bool) {
	var b_b int
	_drv_ch(CLASSID_FONT, _ID_FONT_BOLD, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Font) SetItalic(b bool) {
	_drv_ch(CLASSID_FONT, _ID_FONT_SETITALIC, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) Italic() (b bool) {
	var b_b int
	_drv_ch(CLASSID_FONT, _ID_FONT_ITALIC, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Font) SetKerning(b bool) {
	_drv_ch(CLASSID_FONT, _ID_FONT_SETKERNING, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) Kerning() (b bool) {
	var b_b int
	_drv_ch(CLASSID_FONT, _ID_FONT_KERNING, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Font) SetLetterSpacing(typ int, spacing float64) {
	_drv_ch(CLASSID_FONT, _ID_FONT_SETLETTERSPACING, unsafe.Pointer(p.info()), unsafe.Pointer(&typ), unsafe.Pointer(&spacing), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) LetterSpacing() (typ int, spacing float64) {
	_drv_ch(CLASSID_FONT, _ID_FONT_LETTERSPACING, unsafe.Pointer(p.info()), unsafe.Pointer(&typ), unsafe.Pointer(&spacing), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) SetOverline(b bool) {
	_drv_ch(CLASSID_FONT, _ID_FONT_SETOVERLINE, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) Overline() (b bool) {
	var b_b int
	_drv_ch(CLASSID_FONT, _ID_FONT_OVERLINE, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Font) SetPixelSize(size int) {
	_drv_ch(CLASSID_FONT, _ID_FONT_SETPIXELSIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&size), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) PixelSize() (size int) {
	_drv_ch(CLASSID_FONT, _ID_FONT_PIXELSIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&size), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) SetRawMode(b bool) {
	_drv_ch(CLASSID_FONT, _ID_FONT_SETRAWMODE, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) RawMode() (b bool) {
	var b_b int
	_drv_ch(CLASSID_FONT, _ID_FONT_RAWMODE, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Font) SetRawName(name string) {
	_drv_ch(CLASSID_FONT, _ID_FONT_SETRAWNAME, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&name))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) RawName() (name string) {
	var sh_name utf8_info
	_drv_ch(CLASSID_FONT, _ID_FONT_RAWNAME, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_name))), nil, nil, nil, nil, nil, nil, nil, nil)
	name = sh_name.String()
	return
}

func (p *Font) SetWeight(weight int) {
	_drv_ch(CLASSID_FONT, _ID_FONT_SETWEIGHT, unsafe.Pointer(p.info()), unsafe.Pointer(&weight), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) Weight() (weight int) {
	_drv_ch(CLASSID_FONT, _ID_FONT_WEIGHT, unsafe.Pointer(p.info()), unsafe.Pointer(&weight), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) SetFixedPitch(b bool) {
	_drv_ch(CLASSID_FONT, _ID_FONT_SETFIXEDPITCH, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) FixedPitch() (b bool) {
	var b_b int
	_drv_ch(CLASSID_FONT, _ID_FONT_FIXEDPITCH, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Font) SetStretch(factor int) {
	_drv_ch(CLASSID_FONT, _ID_FONT_SETSTRETCH, unsafe.Pointer(p.info()), unsafe.Pointer(&factor), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) Stretch() (factor int) {
	_drv_ch(CLASSID_FONT, _ID_FONT_STRETCH, unsafe.Pointer(p.info()), unsafe.Pointer(&factor), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) SetStrikeOut(b bool) {
	_drv_ch(CLASSID_FONT, _ID_FONT_SETSTRIKEOUT, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) StrikeOut() (b bool) {
	var b_b int
	_drv_ch(CLASSID_FONT, _ID_FONT_STRIKEOUT, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Font) SetUnderline(b bool) {
	_drv_ch(CLASSID_FONT, _ID_FONT_SETUNDERLINE, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Font) Underline() (b bool) {
	var b_b int
	_drv_ch(CLASSID_FONT, _ID_FONT_UNDERLINE, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

// struct Pixmap
//
type Pixmap struct {
	object
}

func (p *Pixmap) Name() string {
	return "Pixmap"
}

func (p *Pixmap) String() string {
	return DumpObject(p)
}
func (p *Pixmap) SetAttr(attr string, value interface{}) bool {
	switch attr {
	default:
		return p.object.SetAttr(attr, value)
	}
	return false
}
func (p *Pixmap) Attr(attr string) interface{} {
	switch attr {
	case "depth":
		return p.Depth()
	case "size":
		return p.Size()
	case "rect":
		return p.Rect()
	case "alpha":
		return p.HasAlpha()
	case "alphachannel":
		return p.HasAlphaChannel()
	case "null":
		return p.IsNull()
	case "width":
		return p.Width()
	case "height":
		return p.Height()
	default:
		return p.object.Attr(attr)
	}
	return nil
}
func NewPixmap() *Pixmap {
	return new(Pixmap).Init()
}

func (p *Pixmap) Init() *Pixmap {
	p.classid = CLASSID_PIXMAP
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewPixmapWithFile(name string) *Pixmap {
	return new(Pixmap).InitWithFile(name)
}

func (p *Pixmap) InitWithFile(name string) *Pixmap {
	p.classid = CLASSID_PIXMAP
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_INITWITHFILE, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&name))), nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewPixmapWithData(data []byte) *Pixmap {
	return new(Pixmap).InitWithData(data)
}

func (p *Pixmap) InitWithData(data []byte) *Pixmap {
	p.classid = CLASSID_PIXMAP
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_INITWITHDATA, unsafe.Pointer(p.info()), unsafe.Pointer((*slice_info)(unsafe.Pointer(&data))), nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewPixmapWithSize(width int, height int) *Pixmap {
	return new(Pixmap).InitWithSize(width, height)
}

func (p *Pixmap) InitWithSize(width int, height int) *Pixmap {
	p.classid = CLASSID_PIXMAP
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_INITWITHSIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&width), unsafe.Pointer(&height), nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *Pixmap) Close() (err error) {
	if p == nil || p.native == 0 {
		return
	}
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_CLOSE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	RemoveObject(p.native)
	p.native = 0
	return
}

func (p *Pixmap) Depth() (n int) {
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_DEPTH, unsafe.Pointer(p.info()), unsafe.Pointer(&n), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Pixmap) Size() (sz Size) {
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_SIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&sz), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Pixmap) Rect() (rc Rect) {
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_RECT, unsafe.Pointer(p.info()), unsafe.Pointer(&rc), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Pixmap) HasAlpha() (b bool) {
	var b_b int
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_HASALPHA, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Pixmap) HasAlphaChannel() (b bool) {
	var b_b int
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_HASALPHACHANNEL, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Pixmap) IsNull() (b bool) {
	var b_b int
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_ISNULL, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Pixmap) Width() (width int) {
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_WIDTH, unsafe.Pointer(p.info()), unsafe.Pointer(&width), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Pixmap) Height() (height int) {
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_HEIGHT, unsafe.Pointer(p.info()), unsafe.Pointer(&height), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Pixmap) Scaled(width int, height int, aspectRatioMode AspectRatioMode, transformMode TransformationMode) (pixmap *Pixmap) {
	var oi_pixmap obj_info
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_SCALED, unsafe.Pointer(p.info()), unsafe.Pointer(&width), unsafe.Pointer(&height), unsafe.Pointer(&aspectRatioMode), unsafe.Pointer(&transformMode), unsafe.Pointer(&oi_pixmap), nil, nil, nil, nil)
	if oi_pixmap.native != 0 {
		v := FindObject(oi_pixmap.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_PIXMAP, oi_pixmap.native)
		}
		if v != nil {
			pixmap = v.(*Pixmap)
		}
	}
	return
}

func (p *Pixmap) ScaledToHeight(height int, transformMode TransformationMode) (pixmap *Pixmap) {
	var oi_pixmap obj_info
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_SCALEDTOHEIGHT, unsafe.Pointer(p.info()), unsafe.Pointer(&height), unsafe.Pointer(&transformMode), unsafe.Pointer(&oi_pixmap), nil, nil, nil, nil, nil, nil)
	if oi_pixmap.native != 0 {
		v := FindObject(oi_pixmap.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_PIXMAP, oi_pixmap.native)
		}
		if v != nil {
			pixmap = v.(*Pixmap)
		}
	}
	return
}

func (p *Pixmap) ScaledToWidth(width int, transformMode TransformationMode) (pixmap *Pixmap) {
	var oi_pixmap obj_info
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_SCALEDTOWIDTH, unsafe.Pointer(p.info()), unsafe.Pointer(&width), unsafe.Pointer(&transformMode), unsafe.Pointer(&oi_pixmap), nil, nil, nil, nil, nil, nil)
	if oi_pixmap.native != 0 {
		v := FindObject(oi_pixmap.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_PIXMAP, oi_pixmap.native)
		}
		if v != nil {
			pixmap = v.(*Pixmap)
		}
	}
	return
}

func (p *Pixmap) ToImage() (image *Image) {
	var oi_image obj_info
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_TOIMAGE, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_image), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_image.native != 0 {
		v := FindObject(oi_image.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_IMAGE, oi_image.native)
		}
		if v != nil {
			image = v.(*Image)
		}
	}
	return
}

func (p *Pixmap) Load(fileName string) (b bool) {
	var b_b int
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_LOAD, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&fileName))), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Pixmap) Save(fileName string, quality int) (b bool) {
	var b_b int
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_SAVE, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&fileName))), unsafe.Pointer(&quality), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Pixmap) Fill(clr color.Color) {
	sh_clr := make_rgba(clr)
	_drv_ch(CLASSID_PIXMAP, _ID_PIXMAP_FILL, unsafe.Pointer(p.info()), unsafe.Pointer(&sh_clr), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

// struct Icon
//
type Icon struct {
	object
}

func (p *Icon) Name() string {
	return "Icon"
}

func (p *Icon) String() string {
	return DumpObject(p)
}
func (p *Icon) SetAttr(attr string, value interface{}) bool {
	switch attr {
	default:
		return p.object.SetAttr(attr, value)
	}
	return false
}
func (p *Icon) Attr(attr string) interface{} {
	switch attr {
	default:
		return p.object.Attr(attr)
	}
	return nil
}
func NewIcon() *Icon {
	return new(Icon).Init()
}

func (p *Icon) Init() *Icon {
	p.classid = CLASSID_ICON
	_drv_ch(CLASSID_ICON, _ID_ICON_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewIconWithFile(name string) *Icon {
	return new(Icon).InitWithFile(name)
}

func (p *Icon) InitWithFile(name string) *Icon {
	p.classid = CLASSID_ICON
	_drv_ch(CLASSID_ICON, _ID_ICON_INITWITHFILE, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&name))), nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewIconWithPixmap(pixmap *Pixmap) *Icon {
	return new(Icon).InitWithPixmap(pixmap)
}

func (p *Icon) InitWithPixmap(pixmap *Pixmap) *Icon {
	p.classid = CLASSID_ICON
	_drv_ch(CLASSID_ICON, _ID_ICON_INITWITHPIXMAP, unsafe.Pointer(p.info()), unsafe.Pointer(pixmap), nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *Icon) Close() (err error) {
	if p == nil || p.native == 0 {
		return
	}
	_drv_ch(CLASSID_ICON, _ID_ICON_CLOSE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	RemoveObject(p.native)
	p.native = 0
	return
}

// struct Image
//
type Image struct {
	object
}

func (p *Image) Name() string {
	return "Image"
}

func (p *Image) String() string {
	return DumpObject(p)
}
func (p *Image) SetAttr(attr string, value interface{}) bool {
	switch attr {
	default:
		return p.object.SetAttr(attr, value)
	}
	return false
}
func (p *Image) Attr(attr string) interface{} {
	switch attr {
	case "depth":
		return p.Depth()
	case "size":
		return p.Size()
	case "rect":
		return p.Rect()
	default:
		return p.object.Attr(attr)
	}
	return nil
}
func NewImage() *Image {
	return new(Image).Init()
}

func (p *Image) Init() *Image {
	p.classid = CLASSID_IMAGE
	_drv_ch(CLASSID_IMAGE, _ID_IMAGE_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewImageWithFile(name string) *Image {
	return new(Image).InitWithFile(name)
}

func (p *Image) InitWithFile(name string) *Image {
	p.classid = CLASSID_IMAGE
	_drv_ch(CLASSID_IMAGE, _ID_IMAGE_INITWITHFILE, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&name))), nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewImageWithSize(width int, height int) *Image {
	return new(Image).InitWithSize(width, height)
}

func (p *Image) InitWithSize(width int, height int) *Image {
	p.classid = CLASSID_IMAGE
	_drv_ch(CLASSID_IMAGE, _ID_IMAGE_INITWITHSIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&width), unsafe.Pointer(&height), nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *Image) Close() (err error) {
	if p == nil || p.native == 0 {
		return
	}
	_drv_ch(CLASSID_IMAGE, _ID_IMAGE_CLOSE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	RemoveObject(p.native)
	p.native = 0
	return
}

func (p *Image) Depth() (n int) {
	_drv_ch(CLASSID_IMAGE, _ID_IMAGE_DEPTH, unsafe.Pointer(p.info()), unsafe.Pointer(&n), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Image) Size() (sz Size) {
	_drv_ch(CLASSID_IMAGE, _ID_IMAGE_SIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&sz), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Image) Rect() (rc Rect) {
	_drv_ch(CLASSID_IMAGE, _ID_IMAGE_RECT, unsafe.Pointer(p.info()), unsafe.Pointer(&rc), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Image) Fill(color uint) {
	_drv_ch(CLASSID_IMAGE, _ID_IMAGE_FILL, unsafe.Pointer(p.info()), unsafe.Pointer(&color), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Image) Scaled(width int, height int, aspectRatioMode AspectRatioMode, transformMode TransformationMode) (image *Image) {
	var oi_image obj_info
	_drv_ch(CLASSID_IMAGE, _ID_IMAGE_SCALED, unsafe.Pointer(p.info()), unsafe.Pointer(&width), unsafe.Pointer(&height), unsafe.Pointer(&aspectRatioMode), unsafe.Pointer(&transformMode), unsafe.Pointer(&oi_image), nil, nil, nil, nil)
	if oi_image.native != 0 {
		v := FindObject(oi_image.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_IMAGE, oi_image.native)
		}
		if v != nil {
			image = v.(*Image)
		}
	}
	return
}

// struct Widget
//
type Widget struct {
	object
}

func (p *Widget) Name() string {
	return "Widget"
}

func (p *Widget) String() string {
	return DumpObject(p)
}

func (p *Widget) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "autoclose":
		if v, ok := value.(bool); ok {
			p.SetAutoClose(v)
			return true
		}
		return false
	case "layout":
		if v, ok := value.(ILayout); ok {
			p.SetLayout(v)
			return true
		}
		return false
	case "parent":
		if v, ok := value.(IWidget); ok {
			p.SetParent(v)
			return true
		}
		return false
	case "visible":
		if v, ok := value.(bool); ok {
			p.SetVisible(v)
			return true
		}
		return false
	case "windowtitle":
		if v, ok := value.(string); ok {
			p.SetWindowTitle(v)
			return true
		}
		return false
	case "windowicon":
		if v, ok := value.(*Icon); ok {
			p.SetWindowIcon(v)
			return true
		}
		return false
	case "pos":
		if v, ok := value.(Point); ok {
			p.SetPos(v)
			return true
		}
		return false
	case "size":
		if v, ok := value.(Size); ok {
			p.SetSize(v)
			return true
		}
		return false
	case "minimumsize":
		if v, ok := value.(Size); ok {
			p.SetMinimumSize(v)
			return true
		}
		return false
	case "maximumsize":
		if v, ok := value.(Size); ok {
			p.SetMaximumSize(v)
			return true
		}
		return false
	case "fixedsize":
		if v, ok := value.(Size); ok {
			p.SetFixedSize(v)
			return true
		}
		return false
	case "geometry":
		if v, ok := value.(Rect); ok {
			p.SetGeometry(v)
			return true
		}
		return false
	case "font":
		if v, ok := value.(*Font); ok {
			p.SetFont(v)
			return true
		}
		return false
	case "tooltip":
		if v, ok := value.(string); ok {
			p.SetToolTip(v)
			return true
		}
		return false
	case "updatesenabled":
		if v, ok := value.(bool); ok {
			p.SetUpdatesEnabled(v)
			return true
		}
		return false
	case "sizepolicy":
		if v, ok := value.(*SizePolicy); ok {
			p.SetSizePolicy(v)
			return true
		}
		return false
	default:
		return p.object.SetAttr(attr, value)
	}
	return false
}
func (p *Widget) Attr(attr string) interface{} {
	switch attr {
	case "autoclose":
		return p.AutoClose()
	case "layout":
		return p.Layout()
	case "parent":
		return p.Parent()
	case "visible":
		return p.IsVisible()
	case "windowtitle":
		return p.WindowTitle()
	case "windowicon":
		return p.WindowIcon()
	case "pos":
		return p.Pos()
	case "size":
		return p.Size()
	case "minimumsize":
		return p.MinimumSize()
	case "maximumsize":
		return p.MaximumSize()
	case "geometry":
		return p.Geometry()
	case "font":
		return p.Font()
	case "tooltip":
		return p.ToolTip()
	case "updatesenabled":
		return p.UpdatesEnabled()
	case "activatewindow":
		return p.IsActivateWindow()
	case "sizepolicy":
		return p.SizePolicy()
	default:
		return p.object.Attr(attr)
	}
	return nil
}
func NewWidget() *Widget {
	return new(Widget).Init()
}

func (p *Widget) Init() *Widget {
	p.classid = CLASSID_WIDGET
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *Widget) Close() (err error) {
	if p == nil || p.native == 0 {
		return
	}
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_CLOSE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	RemoveObject(p.native)
	p.native = 0
	return
}

func (p *Widget) SetAutoClose(b bool) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_SETAUTOCLOSE, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) AutoClose() (b bool) {
	var b_b int
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_AUTOCLOSE, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Widget) SetLayout(layout ILayout) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_SETLAYOUT, unsafe.Pointer(p.info()), unsafe.Pointer(layout.(iobj).info()), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) Layout() (layout ILayout) {
	var oi_layout obj_info
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_LAYOUT, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_layout), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_layout.native != 0 {
		item := FindObject(oi_layout.native)
		if item != nil {
			layout = item.(ILayout)
		}
	}
	return
}

func (p *Widget) SetParent(parent IWidget) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_SETPARENT, unsafe.Pointer(p.info()), unsafe.Pointer(parent.(iobj).info()), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) Parent() (parent IWidget) {
	var oi_parent obj_info
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_PARENT, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_parent), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_parent.native != 0 {
		item := FindObject(oi_parent.native)
		if item != nil {
			parent = item.(IWidget)
		}
	}
	return
}

func (p *Widget) SetVisible(b bool) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_SETVISIBLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) IsVisible() (b bool) {
	var b_b int
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_ISVISIBLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Widget) SetWindowTitle(text string) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_SETWINDOWTITLE, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) WindowTitle() (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_WINDOWTITLE, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *Widget) SetWindowIcon(icon *Icon) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_SETWINDOWICON, unsafe.Pointer(p.info()), unsafe.Pointer(icon), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) WindowIcon() (icon *Icon) {
	var oi_icon obj_info
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_WINDOWICON, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_icon), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_icon.native != 0 {
		v := FindObject(oi_icon.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ICON, oi_icon.native)
		}
		if v != nil {
			icon = v.(*Icon)
		}
	}
	return
}

func (p *Widget) SetPos(pt Point) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_SETPOS, unsafe.Pointer(p.info()), unsafe.Pointer(&pt), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) Pos() (pt Point) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_POS, unsafe.Pointer(p.info()), unsafe.Pointer(&pt), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) SetSize(sz Size) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_SETSIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&sz), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) Size() (sz Size) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_SIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&sz), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) SetMinimumSize(sz Size) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_SETMINIMUMSIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&sz), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) MinimumSize() (sz Size) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_MINIMUMSIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&sz), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) SetMaximumSize(sz Size) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_SETMAXIMUMSIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&sz), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) MaximumSize() (sz Size) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_MAXIMUMSIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&sz), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) SetFixedSize(sz Size) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_SETFIXEDSIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&sz), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) SetGeometry(rc Rect) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_SETGEOMETRY, unsafe.Pointer(p.info()), unsafe.Pointer(&rc), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) Geometry() (rc Rect) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_GEOMETRY, unsafe.Pointer(p.info()), unsafe.Pointer(&rc), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) SetFont(font *Font) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_SETFONT, unsafe.Pointer(p.info()), unsafe.Pointer(font), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) Font() (font *Font) {
	var oi_font obj_info
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_FONT, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_font), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_font.native != 0 {
		v := FindObject(oi_font.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_FONT, oi_font.native)
		}
		if v != nil {
			font = v.(*Font)
		}
	}
	return
}

func (p *Widget) SetToolTip(text string) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_SETTOOLTIP, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) ToolTip() (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_TOOLTIP, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *Widget) SetUpdatesEnabled(b bool) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_SETUPDATESENABLED, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) UpdatesEnabled() (b bool) {
	var b_b int
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_UPDATESENABLED, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Widget) IsActivateWindow() (b bool) {
	var b_b int
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_ISACTIVATEWINDOW, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Widget) ActivateWindow() {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_ACTIVATEWINDOW, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) SetSizePolicy(policy *SizePolicy) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_SETSIZEPOLICY, unsafe.Pointer(p.info()), unsafe.Pointer(policy), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) SizePolicy() (policy *SizePolicy) {
	var oi_policy obj_info
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_SIZEPOLICY, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_policy), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_policy.native != 0 {
		v := FindObject(oi_policy.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_SIZEPOLICY, oi_policy.native)
		}
		if v != nil {
			policy = v.(*SizePolicy)
		}
	}
	return
}

func (p *Widget) Done() {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_DONE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) Update() {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_UPDATE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) Repaint() {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_REPAINT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) StartTimer(millisecond int) (id int) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_STARTTIMER, unsafe.Pointer(p.info()), unsafe.Pointer(&millisecond), unsafe.Pointer(&id), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) KillTimer(id int) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_KILLTIMER, unsafe.Pointer(p.info()), unsafe.Pointer(&id), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) AddAction(act *Action) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_ADDACTION, unsafe.Pointer(p.info()), unsafe.Pointer(act), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) InsertAction(before *Action, act *Action) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_INSERTACTION, unsafe.Pointer(p.info()), unsafe.Pointer(before), unsafe.Pointer(act), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) RemoveAction(act *Action) {
	_drv_ch(CLASSID_WIDGET, _ID_WIDGET_REMOVEACTION, unsafe.Pointer(p.info()), unsafe.Pointer(act), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Widget) OnShowEvent(fn func(*ShowEvent)) {
	_drv_event_ch(CLASSID_WIDGET, _ID_WIDGET_ONSHOWEVENT, p, fn)
	return
}

func (p *Widget) OnHideEvent(fn func(*HideEvent)) {
	_drv_event_ch(CLASSID_WIDGET, _ID_WIDGET_ONHIDEEVENT, p, fn)
	return
}

func (p *Widget) OnCloseEvent(fn func(*CloseEvent)) {
	_drv_event_ch(CLASSID_WIDGET, _ID_WIDGET_ONCLOSEEVENT, p, fn)
	return
}

func (p *Widget) OnKeyPressEvent(fn func(*KeyEvent)) {
	_drv_event_ch(CLASSID_WIDGET, _ID_WIDGET_ONKEYPRESSEVENT, p, fn)
	return
}

func (p *Widget) OnKeyReleaseEvent(fn func(*KeyEvent)) {
	_drv_event_ch(CLASSID_WIDGET, _ID_WIDGET_ONKEYRELEASEEVENT, p, fn)
	return
}

func (p *Widget) OnMousePressEvent(fn func(*MouseEvent)) {
	_drv_event_ch(CLASSID_WIDGET, _ID_WIDGET_ONMOUSEPRESSEVENT, p, fn)
	return
}

func (p *Widget) OnMouseReleaseEvent(fn func(*MouseEvent)) {
	_drv_event_ch(CLASSID_WIDGET, _ID_WIDGET_ONMOUSERELEASEEVENT, p, fn)
	return
}

func (p *Widget) OnMouseMoveEvent(fn func(*MouseEvent)) {
	_drv_event_ch(CLASSID_WIDGET, _ID_WIDGET_ONMOUSEMOVEEVENT, p, fn)
	return
}

func (p *Widget) OnMouseDoubleClickEvent(fn func(*MouseEvent)) {
	_drv_event_ch(CLASSID_WIDGET, _ID_WIDGET_ONMOUSEDOUBLECLICKEVENT, p, fn)
	return
}

func (p *Widget) OnMoveEvent(fn func(*MoveEvent)) {
	_drv_event_ch(CLASSID_WIDGET, _ID_WIDGET_ONMOVEEVENT, p, fn)
	return
}

func (p *Widget) OnPaintEvent(fn func(*PaintEvent)) {
	_drv_event_ch(CLASSID_WIDGET, _ID_WIDGET_ONPAINTEVENT, p, fn)
	return
}

func (p *Widget) OnResizeEvent(fn func(*ResizeEvent)) {
	_drv_event_ch(CLASSID_WIDGET, _ID_WIDGET_ONRESIZEEVENT, p, fn)
	return
}

func (p *Widget) OnEnterEvent(fn func(*EnterEvent)) {
	_drv_event_ch(CLASSID_WIDGET, _ID_WIDGET_ONENTEREVENT, p, fn)
	return
}

func (p *Widget) OnLeaveEvent(fn func(*LeaveEvent)) {
	_drv_event_ch(CLASSID_WIDGET, _ID_WIDGET_ONLEAVEEVENT, p, fn)
	return
}

func (p *Widget) OnFocusInEvent(fn func(*FocusEvent)) {
	_drv_event_ch(CLASSID_WIDGET, _ID_WIDGET_ONFOCUSINEVENT, p, fn)
	return
}

func (p *Widget) OnFocusOutEvent(fn func(*FocusEvent)) {
	_drv_event_ch(CLASSID_WIDGET, _ID_WIDGET_ONFOCUSOUTEVENT, p, fn)
	return
}

func (p *Widget) OnTimerEvent(fn func(*TimerEvent)) {
	_drv_event_ch(CLASSID_WIDGET, _ID_WIDGET_ONTIMEREVENT, p, fn)
	return
}

func (p *Widget) Show() {
	p.SetVisible(true)
}

func (p *Widget) Hide() {
	p.SetVisible(false)
}

func (p *Widget) SetPosv(x, y int) {
	p.SetPos(Pt(x, y))
}

func (p *Widget) Posv() (int, int) {
	return p.Pos().Unpack()
}

func (p *Widget) SetSizev(width, height int) {
	p.SetSize(Sz(width, height))
}

func (p *Widget) Sizev() (int, int) {
	return p.Size().Unpack()
}

func (p *Widget) SetGeometryv(x, y, width, height int) {
	p.SetGeometry(Rc(x, y, width, height))
}

func (p *Widget) Geometryv() (int, int, int, int) {
	return p.Geometry().Unpack()
}

// struct Action
//
type Action struct {
	object
}

func (p *Action) Name() string {
	return "Action"
}

func (p *Action) String() string {
	return DumpObject(p)
}
func (p *Action) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "text":
		if v, ok := value.(string); ok {
			p.SetText(v)
			return true
		}
		return false
	case "icon":
		if v, ok := value.(*Icon); ok {
			p.SetIcon(v)
			return true
		}
		return false
	case "icontext":
		if v, ok := value.(string); ok {
			p.SetIconText(v)
			return true
		}
		return false
	case "tooltip":
		if v, ok := value.(string); ok {
			p.SetToolTip(v)
			return true
		}
		return false
	case "font":
		if v, ok := value.(*Font); ok {
			p.SetFont(v)
			return true
		}
		return false
	case "checkable":
		if v, ok := value.(bool); ok {
			p.SetCheckable(v)
			return true
		}
		return false
	case "checked":
		if v, ok := value.(bool); ok {
			p.SetChecked(v)
			return true
		}
		return false
	case "visible":
		if v, ok := value.(bool); ok {
			p.SetVisible(v)
			return true
		}
		return false
	default:
		return p.object.SetAttr(attr, value)
	}
	return false
}
func (p *Action) Attr(attr string) interface{} {
	switch attr {
	case "text":
		return p.Text()
	case "icon":
		return p.Icon()
	case "icontext":
		return p.IconText()
	case "tooltip":
		return p.ToolTip()
	case "font":
		return p.Font()
	case "checkable":
		return p.IsCheckable()
	case "checked":
		return p.IsChecked()
	case "visible":
		return p.IsVisible()
	default:
		return p.object.Attr(attr)
	}
	return nil
}
func NewAction() *Action {
	return new(Action).Init()
}

func (p *Action) Init() *Action {
	p.classid = CLASSID_ACTION
	_drv_ch(CLASSID_ACTION, _ID_ACTION_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewActionWithText(text string) *Action {
	return new(Action).InitWithText(text)
}

func (p *Action) InitWithText(text string) *Action {
	p.classid = CLASSID_ACTION
	_drv_ch(CLASSID_ACTION, _ID_ACTION_INITWITHTEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *Action) Close() (err error) {
	if p == nil || p.native == 0 {
		return
	}
	_drv_ch(CLASSID_ACTION, _ID_ACTION_CLOSE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	RemoveObject(p.native)
	p.native = 0
	return
}

func (p *Action) SetText(text string) {
	_drv_ch(CLASSID_ACTION, _ID_ACTION_SETTEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Action) Text() (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_ACTION, _ID_ACTION_TEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *Action) SetIcon(icon *Icon) {
	_drv_ch(CLASSID_ACTION, _ID_ACTION_SETICON, unsafe.Pointer(p.info()), unsafe.Pointer(icon), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Action) Icon() (icon *Icon) {
	var oi_icon obj_info
	_drv_ch(CLASSID_ACTION, _ID_ACTION_ICON, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_icon), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_icon.native != 0 {
		v := FindObject(oi_icon.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ICON, oi_icon.native)
		}
		if v != nil {
			icon = v.(*Icon)
		}
	}
	return
}

func (p *Action) SetIconText(text string) {
	_drv_ch(CLASSID_ACTION, _ID_ACTION_SETICONTEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Action) IconText() (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_ACTION, _ID_ACTION_ICONTEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *Action) SetToolTip(text string) {
	_drv_ch(CLASSID_ACTION, _ID_ACTION_SETTOOLTIP, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Action) ToolTip() (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_ACTION, _ID_ACTION_TOOLTIP, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *Action) SetFont(font *Font) {
	_drv_ch(CLASSID_ACTION, _ID_ACTION_SETFONT, unsafe.Pointer(p.info()), unsafe.Pointer(font), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Action) Font() (font *Font) {
	var oi_font obj_info
	_drv_ch(CLASSID_ACTION, _ID_ACTION_FONT, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_font), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_font.native != 0 {
		v := FindObject(oi_font.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_FONT, oi_font.native)
		}
		if v != nil {
			font = v.(*Font)
		}
	}
	return
}

func (p *Action) SetCheckable(b bool) {
	_drv_ch(CLASSID_ACTION, _ID_ACTION_SETCHECKABLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Action) IsCheckable() (b bool) {
	var b_b int
	_drv_ch(CLASSID_ACTION, _ID_ACTION_ISCHECKABLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Action) SetChecked(b bool) {
	_drv_ch(CLASSID_ACTION, _ID_ACTION_SETCHECKED, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Action) IsChecked() (b bool) {
	var b_b int
	_drv_ch(CLASSID_ACTION, _ID_ACTION_ISCHECKED, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Action) SetVisible(b bool) {
	_drv_ch(CLASSID_ACTION, _ID_ACTION_SETVISIBLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Action) IsVisible() (b bool) {
	var b_b int
	_drv_ch(CLASSID_ACTION, _ID_ACTION_ISVISIBLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Action) OnTriggered(fn func(bool)) {
	_drv_event_ch(CLASSID_ACTION, _ID_ACTION_ONTRIGGERED, p, fn)
	return
}

// struct ActionGroup
//
type ActionGroup struct {
	object
}

func (p *ActionGroup) Name() string {
	return "ActionGroup"
}

func (p *ActionGroup) String() string {
	return DumpObject(p)
}
func (p *ActionGroup) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "visible":
		if v, ok := value.(bool); ok {
			p.SetVisible(v)
			return true
		}
		return false
	case "enabled":
		if v, ok := value.(bool); ok {
			p.SetEnabled(v)
			return true
		}
		return false
	case "exclusive":
		if v, ok := value.(bool); ok {
			p.SetExclusive(v)
			return true
		}
		return false
	default:
		return p.object.SetAttr(attr, value)
	}
	return false
}
func (p *ActionGroup) Attr(attr string) interface{} {
	switch attr {
	case "visible":
		return p.IsVisible()
	case "enabled":
		return p.IsEnabled()
	case "exclusive":
		return p.IsExclusive()
	default:
		return p.object.Attr(attr)
	}
	return nil
}
func NewActionGroup() *ActionGroup {
	return new(ActionGroup).Init()
}

func (p *ActionGroup) Init() *ActionGroup {
	p.classid = CLASSID_ACTIONGROUP
	_drv_ch(CLASSID_ACTIONGROUP, _ID_ACTIONGROUP_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *ActionGroup) Close() (err error) {
	if p == nil || p.native == 0 {
		return
	}
	_drv_ch(CLASSID_ACTIONGROUP, _ID_ACTIONGROUP_CLOSE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	RemoveObject(p.native)
	p.native = 0
	return
}

func (p *ActionGroup) SetVisible(b bool) {
	_drv_ch(CLASSID_ACTIONGROUP, _ID_ACTIONGROUP_SETVISIBLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ActionGroup) IsVisible() (b bool) {
	var b_b int
	_drv_ch(CLASSID_ACTIONGROUP, _ID_ACTIONGROUP_ISVISIBLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *ActionGroup) SetEnabled(b bool) {
	_drv_ch(CLASSID_ACTIONGROUP, _ID_ACTIONGROUP_SETENABLED, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ActionGroup) IsEnabled() (b bool) {
	var b_b int
	_drv_ch(CLASSID_ACTIONGROUP, _ID_ACTIONGROUP_ISENABLED, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *ActionGroup) SetExclusive(b bool) {
	_drv_ch(CLASSID_ACTIONGROUP, _ID_ACTIONGROUP_SETEXCLUSIVE, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ActionGroup) IsExclusive() (b bool) {
	var b_b int
	_drv_ch(CLASSID_ACTIONGROUP, _ID_ACTIONGROUP_ISEXCLUSIVE, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *ActionGroup) AddAction(act *Action) {
	_drv_ch(CLASSID_ACTIONGROUP, _ID_ACTIONGROUP_ADDACTION, unsafe.Pointer(p.info()), unsafe.Pointer(act), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ActionGroup) RemoveAction(act *Action) {
	_drv_ch(CLASSID_ACTIONGROUP, _ID_ACTIONGROUP_REMOVEACTION, unsafe.Pointer(p.info()), unsafe.Pointer(act), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ActionGroup) OnHovered(fn func(*Action)) {
	_drv_event_ch(CLASSID_ACTIONGROUP, _ID_ACTIONGROUP_ONHOVERED, p, fn)
	return
}

func (p *ActionGroup) OnTriggered(fn func(*Action)) {
	_drv_event_ch(CLASSID_ACTIONGROUP, _ID_ACTIONGROUP_ONTRIGGERED, p, fn)
	return
}

// struct Menu
//
type Menu struct {
	Widget
}

func (p *Menu) Name() string {
	return "Menu"
}

func (p *Menu) String() string {
	return DumpObject(p)
}
func (p *Menu) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "title":
		if v, ok := value.(string); ok {
			p.SetTitle(v)
			return true
		}
		return false
	case "icon":
		if v, ok := value.(*Icon); ok {
			p.SetIcon(v)
			return true
		}
		return false
	case "defaultaction":
		if v, ok := value.(*Action); ok {
			p.SetDefaultAction(v)
			return true
		}
		return false
	case "activeaction":
		if v, ok := value.(*Action); ok {
			p.SetActiveAction(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr, value)
	}
	return false
}
func (p *Menu) Attr(attr string) interface{} {
	switch attr {
	case "title":
		return p.Title()
	case "icon":
		return p.Icon()
	case "defaultaction":
		return p.DefaultAction()
	case "activeaction":
		return p.ActiveAction()
	case "menuaction":
		return p.MenuAction()
	case "empty":
		return p.IsEmpty()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewMenu() *Menu {
	return new(Menu).Init()
}

func (p *Menu) Init() *Menu {
	p.classid = CLASSID_MENU
	_drv_ch(CLASSID_MENU, _ID_MENU_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewMenuWithTitle(text string) *Menu {
	return new(Menu).InitWithTitle(text)
}

func (p *Menu) InitWithTitle(text string) *Menu {
	p.classid = CLASSID_MENU
	_drv_ch(CLASSID_MENU, _ID_MENU_INITWITHTITLE, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *Menu) SetTitle(text string) {
	_drv_ch(CLASSID_MENU, _ID_MENU_SETTITLE, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Menu) Title() (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_MENU, _ID_MENU_TITLE, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *Menu) SetIcon(icon *Icon) {
	_drv_ch(CLASSID_MENU, _ID_MENU_SETICON, unsafe.Pointer(p.info()), unsafe.Pointer(icon), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Menu) Icon() (icon *Icon) {
	var oi_icon obj_info
	_drv_ch(CLASSID_MENU, _ID_MENU_ICON, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_icon), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_icon.native != 0 {
		v := FindObject(oi_icon.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ICON, oi_icon.native)
		}
		if v != nil {
			icon = v.(*Icon)
		}
	}
	return
}

func (p *Menu) SetDefaultAction(act *Action) {
	_drv_ch(CLASSID_MENU, _ID_MENU_SETDEFAULTACTION, unsafe.Pointer(p.info()), unsafe.Pointer(act), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Menu) DefaultAction() (act *Action) {
	var oi_act obj_info
	_drv_ch(CLASSID_MENU, _ID_MENU_DEFAULTACTION, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_act), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_act.native != 0 {
		v := FindObject(oi_act.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ACTION, oi_act.native)
		}
		if v != nil {
			act = v.(*Action)
		}
	}
	return
}

func (p *Menu) SetActiveAction(act *Action) {
	_drv_ch(CLASSID_MENU, _ID_MENU_SETACTIVEACTION, unsafe.Pointer(p.info()), unsafe.Pointer(act), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Menu) ActiveAction() (act *Action) {
	var oi_act obj_info
	_drv_ch(CLASSID_MENU, _ID_MENU_ACTIVEACTION, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_act), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_act.native != 0 {
		v := FindObject(oi_act.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ACTION, oi_act.native)
		}
		if v != nil {
			act = v.(*Action)
		}
	}
	return
}

func (p *Menu) MenuAction() (act *Action) {
	var oi_act obj_info
	_drv_ch(CLASSID_MENU, _ID_MENU_MENUACTION, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_act), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_act.native != 0 {
		v := FindObject(oi_act.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ACTION, oi_act.native)
		}
		if v != nil {
			act = v.(*Action)
		}
	}
	return
}

func (p *Menu) IsEmpty() (b bool) {
	var b_b int
	_drv_ch(CLASSID_MENU, _ID_MENU_ISEMPTY, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Menu) AddMenu(menu *Menu) (act *Action) {
	var oi_act obj_info
	_drv_ch(CLASSID_MENU, _ID_MENU_ADDMENU, unsafe.Pointer(p.info()), unsafe.Pointer(menu), unsafe.Pointer(&oi_act), nil, nil, nil, nil, nil, nil, nil)
	if oi_act.native != 0 {
		v := FindObject(oi_act.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ACTION, oi_act.native)
		}
		if v != nil {
			act = v.(*Action)
		}
	}
	return
}

func (p *Menu) InsertMenu(before *Action, menu *Menu) {
	_drv_ch(CLASSID_MENU, _ID_MENU_INSERTMENU, unsafe.Pointer(p.info()), unsafe.Pointer(before), unsafe.Pointer(menu), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Menu) AddSeparator() (act *Action) {
	var oi_act obj_info
	_drv_ch(CLASSID_MENU, _ID_MENU_ADDSEPARATOR, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_act), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_act.native != 0 {
		v := FindObject(oi_act.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ACTION, oi_act.native)
		}
		if v != nil {
			act = v.(*Action)
		}
	}
	return
}

func (p *Menu) InsertSeparator(before *Action) (act *Action) {
	var oi_act obj_info
	_drv_ch(CLASSID_MENU, _ID_MENU_INSERTSEPARATOR, unsafe.Pointer(p.info()), unsafe.Pointer(before), unsafe.Pointer(&oi_act), nil, nil, nil, nil, nil, nil, nil)
	if oi_act.native != 0 {
		v := FindObject(oi_act.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ACTION, oi_act.native)
		}
		if v != nil {
			act = v.(*Action)
		}
	}
	return
}

func (p *Menu) Clear() {
	_drv_ch(CLASSID_MENU, _ID_MENU_CLEAR, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Menu) Popup(pt Point, act *Action) {
	_drv_ch(CLASSID_MENU, _ID_MENU_POPUP, unsafe.Pointer(p.info()), unsafe.Pointer(&pt), unsafe.Pointer(act), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Menu) OnHovered(fn func(*Action)) {
	_drv_event_ch(CLASSID_MENU, _ID_MENU_ONHOVERED, p, fn)
	return
}

func (p *Menu) OnTriggered(fn func(*Action)) {
	_drv_event_ch(CLASSID_MENU, _ID_MENU_ONTRIGGERED, p, fn)
	return
}

// struct MenuBar
//
type MenuBar struct {
	Widget
}

func (p *MenuBar) Name() string {
	return "MenuBar"
}

func (p *MenuBar) String() string {
	return DumpObject(p)
}
func (p *MenuBar) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "activeaction":
		if v, ok := value.(*Action); ok {
			p.SetActiveAction(v)
			return true
		}
		return false
	case "nativemenubar":
		if v, ok := value.(bool); ok {
			p.SetNativeMenuBar(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr, value)
	}
	return false
}
func (p *MenuBar) Attr(attr string) interface{} {
	switch attr {
	case "activeaction":
		return p.ActiveAction()
	case "nativemenubar":
		return p.IsNativeMenuBar()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewMenuBar() *MenuBar {
	return new(MenuBar).Init()
}

func (p *MenuBar) Init() *MenuBar {
	p.classid = CLASSID_MENUBAR
	_drv_ch(CLASSID_MENUBAR, _ID_MENUBAR_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *MenuBar) SetActiveAction(act *Action) {
	_drv_ch(CLASSID_MENUBAR, _ID_MENUBAR_SETACTIVEACTION, unsafe.Pointer(p.info()), unsafe.Pointer(act), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *MenuBar) ActiveAction() (act *Action) {
	var oi_act obj_info
	_drv_ch(CLASSID_MENUBAR, _ID_MENUBAR_ACTIVEACTION, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_act), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_act.native != 0 {
		v := FindObject(oi_act.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ACTION, oi_act.native)
		}
		if v != nil {
			act = v.(*Action)
		}
	}
	return
}

func (p *MenuBar) SetNativeMenuBar(b bool) {
	_drv_ch(CLASSID_MENUBAR, _ID_MENUBAR_SETNATIVEMENUBAR, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *MenuBar) IsNativeMenuBar() (b bool) {
	var b_b int
	_drv_ch(CLASSID_MENUBAR, _ID_MENUBAR_ISNATIVEMENUBAR, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *MenuBar) AddMenu(menu *Menu) (act *Action) {
	var oi_act obj_info
	_drv_ch(CLASSID_MENUBAR, _ID_MENUBAR_ADDMENU, unsafe.Pointer(p.info()), unsafe.Pointer(menu), unsafe.Pointer(&oi_act), nil, nil, nil, nil, nil, nil, nil)
	if oi_act.native != 0 {
		v := FindObject(oi_act.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ACTION, oi_act.native)
		}
		if v != nil {
			act = v.(*Action)
		}
	}
	return
}

func (p *MenuBar) InsertMenu(before *Action, menu *Menu) {
	_drv_ch(CLASSID_MENUBAR, _ID_MENUBAR_INSERTMENU, unsafe.Pointer(p.info()), unsafe.Pointer(before), unsafe.Pointer(menu), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *MenuBar) AddSeparator() (act *Action) {
	var oi_act obj_info
	_drv_ch(CLASSID_MENUBAR, _ID_MENUBAR_ADDSEPARATOR, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_act), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_act.native != 0 {
		v := FindObject(oi_act.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ACTION, oi_act.native)
		}
		if v != nil {
			act = v.(*Action)
		}
	}
	return
}

func (p *MenuBar) InsertSeparator(before *Action) (act *Action) {
	var oi_act obj_info
	_drv_ch(CLASSID_MENUBAR, _ID_MENUBAR_INSERTSEPARATOR, unsafe.Pointer(p.info()), unsafe.Pointer(before), unsafe.Pointer(&oi_act), nil, nil, nil, nil, nil, nil, nil)
	if oi_act.native != 0 {
		v := FindObject(oi_act.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ACTION, oi_act.native)
		}
		if v != nil {
			act = v.(*Action)
		}
	}
	return
}

func (p *MenuBar) Clear() {
	_drv_ch(CLASSID_MENUBAR, _ID_MENUBAR_CLEAR, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *MenuBar) OnHovered(fn func(*Action)) {
	_drv_event_ch(CLASSID_MENUBAR, _ID_MENUBAR_ONHOVERED, p, fn)
	return
}

func (p *MenuBar) OnTriggered(fn func(*Action)) {
	_drv_event_ch(CLASSID_MENUBAR, _ID_MENUBAR_ONTRIGGERED, p, fn)
	return
}

// struct ToolBar
//
type ToolBar struct {
	Widget
}

func (p *ToolBar) Name() string {
	return "ToolBar"
}

func (p *ToolBar) String() string {
	return DumpObject(p)
}
func (p *ToolBar) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "title":
		if v, ok := value.(string); ok {
			p.SetTitle(v)
			return true
		}
		return false
	case "iconsize":
		if v, ok := value.(Size); ok {
			p.SetIconSize(v)
			return true
		}
		return false
	case "floatable":
		if v, ok := value.(bool); ok {
			p.SetFloatable(v)
			return true
		}
		return false
	case "toolbuttonstyle":
		if v, ok := value.(ToolButtonStyle); ok {
			p.SetToolButtonStyle(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr, value)
	}
	return false
}
func (p *ToolBar) Attr(attr string) interface{} {
	switch attr {
	case "title":
		return p.Title()
	case "iconsize":
		return p.IconSize()
	case "floatable":
		return p.IsFloatable()
	case "floating":
		return p.IsFloating()
	case "toolbuttonstyle":
		return p.ToolButtonStyle()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewToolBar() *ToolBar {
	return new(ToolBar).Init()
}

func (p *ToolBar) Init() *ToolBar {
	p.classid = CLASSID_TOOLBAR
	_drv_ch(CLASSID_TOOLBAR, _ID_TOOLBAR_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewToolBarWithTitle(text string) *ToolBar {
	return new(ToolBar).InitWithTitle(text)
}

func (p *ToolBar) InitWithTitle(text string) *ToolBar {
	p.classid = CLASSID_TOOLBAR
	_drv_ch(CLASSID_TOOLBAR, _ID_TOOLBAR_INITWITHTITLE, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *ToolBar) SetTitle(text string) {
	_drv_ch(CLASSID_TOOLBAR, _ID_TOOLBAR_SETTITLE, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolBar) Title() (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_TOOLBAR, _ID_TOOLBAR_TITLE, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *ToolBar) SetIconSize(sz Size) {
	_drv_ch(CLASSID_TOOLBAR, _ID_TOOLBAR_SETICONSIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&sz), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolBar) IconSize() (sz Size) {
	_drv_ch(CLASSID_TOOLBAR, _ID_TOOLBAR_ICONSIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&sz), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolBar) SetFloatable(b bool) {
	_drv_ch(CLASSID_TOOLBAR, _ID_TOOLBAR_SETFLOATABLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolBar) IsFloatable() (b bool) {
	var b_b int
	_drv_ch(CLASSID_TOOLBAR, _ID_TOOLBAR_ISFLOATABLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *ToolBar) IsFloating() (b bool) {
	var b_b int
	_drv_ch(CLASSID_TOOLBAR, _ID_TOOLBAR_ISFLOATING, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *ToolBar) SetToolButtonStyle(style ToolButtonStyle) {
	_drv_ch(CLASSID_TOOLBAR, _ID_TOOLBAR_SETTOOLBUTTONSTYLE, unsafe.Pointer(p.info()), unsafe.Pointer(&style), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolBar) ToolButtonStyle() (style ToolButtonStyle) {
	_drv_ch(CLASSID_TOOLBAR, _ID_TOOLBAR_TOOLBUTTONSTYLE, unsafe.Pointer(p.info()), unsafe.Pointer(&style), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolBar) AddAction(act *Action) {
	_drv_ch(CLASSID_TOOLBAR, _ID_TOOLBAR_ADDACTION, unsafe.Pointer(p.info()), unsafe.Pointer(act), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolBar) AddSeparator() (act *Action) {
	var oi_act obj_info
	_drv_ch(CLASSID_TOOLBAR, _ID_TOOLBAR_ADDSEPARATOR, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_act), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_act.native != 0 {
		v := FindObject(oi_act.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ACTION, oi_act.native)
		}
		if v != nil {
			act = v.(*Action)
		}
	}
	return
}

func (p *ToolBar) AddWidget(widget IWidget) (act *Action) {
	var oi_act obj_info
	_drv_ch(CLASSID_TOOLBAR, _ID_TOOLBAR_ADDWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(widget.(iobj).info()), unsafe.Pointer(&oi_act), nil, nil, nil, nil, nil, nil, nil)
	if oi_act.native != 0 {
		v := FindObject(oi_act.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ACTION, oi_act.native)
		}
		if v != nil {
			act = v.(*Action)
		}
	}
	return
}

func (p *ToolBar) Clear() {
	_drv_ch(CLASSID_TOOLBAR, _ID_TOOLBAR_CLEAR, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

// struct StatusBar
//
type StatusBar struct {
	Widget
}

func (p *StatusBar) Name() string {
	return "StatusBar"
}

func (p *StatusBar) String() string {
	return DumpObject(p)
}
func (p *StatusBar) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "sizegrip":
		if v, ok := value.(bool); ok {
			p.SetSizeGrip(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr, value)
	}
	return false
}
func (p *StatusBar) Attr(attr string) interface{} {
	switch attr {
	case "sizegrip":
		return p.IsSizeGrip()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewStatusBar() *StatusBar {
	return new(StatusBar).Init()
}

func (p *StatusBar) Init() *StatusBar {
	p.classid = CLASSID_STATUSBAR
	_drv_ch(CLASSID_STATUSBAR, _ID_STATUSBAR_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *StatusBar) SetSizeGrip(b bool) {
	_drv_ch(CLASSID_STATUSBAR, _ID_STATUSBAR_SETSIZEGRIP, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *StatusBar) IsSizeGrip() (b bool) {
	var b_b int
	_drv_ch(CLASSID_STATUSBAR, _ID_STATUSBAR_ISSIZEGRIP, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *StatusBar) AddWidget(widget IWidget, stretch int) {
	_drv_ch(CLASSID_STATUSBAR, _ID_STATUSBAR_ADDWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(widget.(iobj).info()), unsafe.Pointer(&stretch), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *StatusBar) AddPermanentWidget(widget IWidget, stretch int) {
	_drv_ch(CLASSID_STATUSBAR, _ID_STATUSBAR_ADDPERMANENTWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(widget.(iobj).info()), unsafe.Pointer(&stretch), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *StatusBar) InsertWidget(index int, widget IWidget, stretch int) {
	_drv_ch(CLASSID_STATUSBAR, _ID_STATUSBAR_INSERTWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer(widget.(iobj).info()), unsafe.Pointer(&stretch), nil, nil, nil, nil, nil, nil)
	return
}

func (p *StatusBar) InsertPermanentWidget(index int, widget IWidget, stretch int) {
	_drv_ch(CLASSID_STATUSBAR, _ID_STATUSBAR_INSERTPERMANENTWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer(widget.(iobj).info()), unsafe.Pointer(&stretch), nil, nil, nil, nil, nil, nil)
	return
}

func (p *StatusBar) RemoveWidget(widget IWidget) {
	_drv_ch(CLASSID_STATUSBAR, _ID_STATUSBAR_REMOVEWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(widget.(iobj).info()), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *StatusBar) ShowMessage(text string, timeout int) {
	_drv_ch(CLASSID_STATUSBAR, _ID_STATUSBAR_SHOWMESSAGE, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), unsafe.Pointer(&timeout), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *StatusBar) CurrentMessage() (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_STATUSBAR, _ID_STATUSBAR_CURRENTMESSAGE, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *StatusBar) ClearMessage() {
	_drv_ch(CLASSID_STATUSBAR, _ID_STATUSBAR_CLEARMESSAGE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

// struct DockWidget
//
type DockWidget struct {
	Widget
}

func (p *DockWidget) Name() string {
	return "DockWidget"
}

func (p *DockWidget) String() string {
	return DumpObject(p)
}
func (p *DockWidget) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "dock":
		if v, ok := value.(IWidget); ok {
			p.SetDock(v)
			return true
		}
		return false
	case "titlebar":
		if v, ok := value.(IWidget); ok {
			p.SetTitleBar(v)
			return true
		}
		return false
	case "floating":
		if v, ok := value.(bool); ok {
			p.SetFloating(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr, value)
	}
	return false
}
func (p *DockWidget) Attr(attr string) interface{} {
	switch attr {
	case "dock":
		return p.Dock()
	case "titlebar":
		return p.TitleBar()
	case "floating":
		return p.IsFloating()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewDockWidget() *DockWidget {
	return new(DockWidget).Init()
}

func (p *DockWidget) Init() *DockWidget {
	p.classid = CLASSID_DOCKWIDGET
	_drv_ch(CLASSID_DOCKWIDGET, _ID_DOCKWIDGET_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewDockWidgetWithTitle(title string) *DockWidget {
	return new(DockWidget).InitWithTitle(title)
}

func (p *DockWidget) InitWithTitle(title string) *DockWidget {
	p.classid = CLASSID_DOCKWIDGET
	_drv_ch(CLASSID_DOCKWIDGET, _ID_DOCKWIDGET_INITWITHTITLE, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&title))), nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *DockWidget) SetDock(widget IWidget) {
	_drv_ch(CLASSID_DOCKWIDGET, _ID_DOCKWIDGET_SETDOCK, unsafe.Pointer(p.info()), unsafe.Pointer(widget.(iobj).info()), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *DockWidget) Dock() (widget IWidget) {
	var oi_widget obj_info
	_drv_ch(CLASSID_DOCKWIDGET, _ID_DOCKWIDGET_DOCK, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_widget), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_widget.native != 0 {
		item := FindObject(oi_widget.native)
		if item != nil {
			widget = item.(IWidget)
		}
	}
	return
}

func (p *DockWidget) SetTitleBar(widget IWidget) {
	_drv_ch(CLASSID_DOCKWIDGET, _ID_DOCKWIDGET_SETTITLEBAR, unsafe.Pointer(p.info()), unsafe.Pointer(widget.(iobj).info()), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *DockWidget) TitleBar() (widget IWidget) {
	var oi_widget obj_info
	_drv_ch(CLASSID_DOCKWIDGET, _ID_DOCKWIDGET_TITLEBAR, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_widget), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_widget.native != 0 {
		item := FindObject(oi_widget.native)
		if item != nil {
			widget = item.(IWidget)
		}
	}
	return
}

func (p *DockWidget) SetFloating(b bool) {
	_drv_ch(CLASSID_DOCKWIDGET, _ID_DOCKWIDGET_SETFLOATING, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *DockWidget) IsFloating() (b bool) {
	var b_b int
	_drv_ch(CLASSID_DOCKWIDGET, _ID_DOCKWIDGET_ISFLOATING, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *DockWidget) OnVisibilityChanged(fn func(bool)) {
	_drv_event_ch(CLASSID_DOCKWIDGET, _ID_DOCKWIDGET_ONVISIBILITYCHANGED, p, fn)
	return
}

// struct SystemTray
//
type SystemTray struct {
	object
}

func (p *SystemTray) Name() string {
	return "SystemTray"
}

func (p *SystemTray) String() string {
	return DumpObject(p)
}
func (p *SystemTray) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "contextmenu":
		if v, ok := value.(*Menu); ok {
			p.SetContextMenu(v)
			return true
		}
		return false
	case "icon":
		if v, ok := value.(*Icon); ok {
			p.SetIcon(v)
			return true
		}
		return false
	case "tooltip":
		if v, ok := value.(string); ok {
			p.SetToolTip(v)
			return true
		}
		return false
	case "visible":
		if v, ok := value.(bool); ok {
			p.SetVisible(v)
			return true
		}
		return false
	default:
		return p.object.SetAttr(attr, value)
	}
	return false
}
func (p *SystemTray) Attr(attr string) interface{} {
	switch attr {
	case "contextmenu":
		return p.ContextMenu()
	case "icon":
		return p.Icon()
	case "tooltip":
		return p.ToolTip()
	case "visible":
		return p.IsVisible()
	default:
		return p.object.Attr(attr)
	}
	return nil
}
func NewSystemTray() *SystemTray {
	return new(SystemTray).Init()
}

func (p *SystemTray) Init() *SystemTray {
	p.classid = CLASSID_SYSTEMTRAY
	_drv_ch(CLASSID_SYSTEMTRAY, _ID_SYSTEMTRAY_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *SystemTray) Close() (err error) {
	if p == nil || p.native == 0 {
		return
	}
	_drv_ch(CLASSID_SYSTEMTRAY, _ID_SYSTEMTRAY_CLOSE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	RemoveObject(p.native)
	p.native = 0
	return
}

func (p *SystemTray) SetContextMenu(menu *Menu) {
	_drv_ch(CLASSID_SYSTEMTRAY, _ID_SYSTEMTRAY_SETCONTEXTMENU, unsafe.Pointer(p.info()), unsafe.Pointer(menu), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *SystemTray) ContextMenu() (menu *Menu) {
	var oi_menu obj_info
	_drv_ch(CLASSID_SYSTEMTRAY, _ID_SYSTEMTRAY_CONTEXTMENU, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_menu), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_menu.native != 0 {
		v := FindObject(oi_menu.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_MENU, oi_menu.native)
		}
		if v != nil {
			menu = v.(*Menu)
		}
	}
	return
}

func (p *SystemTray) SetIcon(icon *Icon) {
	_drv_ch(CLASSID_SYSTEMTRAY, _ID_SYSTEMTRAY_SETICON, unsafe.Pointer(p.info()), unsafe.Pointer(icon), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *SystemTray) Icon() (icon *Icon) {
	var oi_icon obj_info
	_drv_ch(CLASSID_SYSTEMTRAY, _ID_SYSTEMTRAY_ICON, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_icon), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_icon.native != 0 {
		v := FindObject(oi_icon.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ICON, oi_icon.native)
		}
		if v != nil {
			icon = v.(*Icon)
		}
	}
	return
}

func (p *SystemTray) SetToolTip(text string) {
	_drv_ch(CLASSID_SYSTEMTRAY, _ID_SYSTEMTRAY_SETTOOLTIP, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *SystemTray) ToolTip() (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_SYSTEMTRAY, _ID_SYSTEMTRAY_TOOLTIP, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *SystemTray) SetVisible(b bool) {
	_drv_ch(CLASSID_SYSTEMTRAY, _ID_SYSTEMTRAY_SETVISIBLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *SystemTray) IsVisible() (b bool) {
	var b_b int
	_drv_ch(CLASSID_SYSTEMTRAY, _ID_SYSTEMTRAY_ISVISIBLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *SystemTray) ShowMessage(title string, message string, icontype MessageIconType, mstimeouthint int) {
	_drv_ch(CLASSID_SYSTEMTRAY, _ID_SYSTEMTRAY_SHOWMESSAGE, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&title))), unsafe.Pointer((*string_info)(unsafe.Pointer(&message))), unsafe.Pointer(&icontype), unsafe.Pointer(&mstimeouthint), nil, nil, nil, nil, nil)
	return
}

func (p *SystemTray) OnActivated(fn func(int)) {
	_drv_event_ch(CLASSID_SYSTEMTRAY, _ID_SYSTEMTRAY_ONACTIVATED, p, fn)
	return
}

func (p *SystemTray) OnMessageClicked(fn func()) {
	_drv_event_ch(CLASSID_SYSTEMTRAY, _ID_SYSTEMTRAY_ONMESSAGECLICKED, p, fn)
	return
}

// struct TabWidget
//
type TabWidget struct {
	Widget
}

func (p *TabWidget) Name() string {
	return "TabWidget"
}

func (p *TabWidget) String() string {
	return DumpObject(p)
}
func (p *TabWidget) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "currentindex":
		if v, ok := value.(int); ok {
			p.SetCurrentIndex(v)
			return true
		}
		return false
	case "currentwidget":
		if v, ok := value.(IWidget); ok {
			p.SetCurrentWidget(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr, value)
	}
	return false
}
func (p *TabWidget) Attr(attr string) interface{} {
	switch attr {
	case "count":
		return p.Count()
	case "currentindex":
		return p.CurrentIndex()
	case "currentwidget":
		return p.CurrentWidget()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewTabWidget() *TabWidget {
	return new(TabWidget).Init()
}

func (p *TabWidget) Init() *TabWidget {
	p.classid = CLASSID_TABWIDGET
	_drv_ch(CLASSID_TABWIDGET, _ID_TABWIDGET_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *TabWidget) Count() (n int) {
	_drv_ch(CLASSID_TABWIDGET, _ID_TABWIDGET_COUNT, unsafe.Pointer(p.info()), unsafe.Pointer(&n), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *TabWidget) CurrentIndex() (index int) {
	_drv_ch(CLASSID_TABWIDGET, _ID_TABWIDGET_CURRENTINDEX, unsafe.Pointer(p.info()), unsafe.Pointer(&index), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *TabWidget) CurrentWidget() (w IWidget) {
	var oi_w obj_info
	_drv_ch(CLASSID_TABWIDGET, _ID_TABWIDGET_CURRENTWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_w), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_w.native != 0 {
		item := FindObject(oi_w.native)
		if item != nil {
			w = item.(IWidget)
		}
	}
	return
}

func (p *TabWidget) SetCurrentIndex(index int) {
	_drv_ch(CLASSID_TABWIDGET, _ID_TABWIDGET_SETCURRENTINDEX, unsafe.Pointer(p.info()), unsafe.Pointer(&index), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *TabWidget) SetCurrentWidget(w IWidget) {
	_drv_ch(CLASSID_TABWIDGET, _ID_TABWIDGET_SETCURRENTWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(w.(iobj).info()), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *TabWidget) AddTab(w IWidget, label string, icon *Icon) {
	_drv_ch(CLASSID_TABWIDGET, _ID_TABWIDGET_ADDTAB, unsafe.Pointer(p.info()), unsafe.Pointer(w.(iobj).info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&label))), unsafe.Pointer(icon), nil, nil, nil, nil, nil, nil)
	return
}

func (p *TabWidget) InsertTab(index int, w IWidget, label string, icon *Icon) {
	_drv_ch(CLASSID_TABWIDGET, _ID_TABWIDGET_INSERTTAB, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer(w.(iobj).info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&label))), unsafe.Pointer(icon), nil, nil, nil, nil, nil)
	return
}

func (p *TabWidget) RemoveTab(index int) {
	_drv_ch(CLASSID_TABWIDGET, _ID_TABWIDGET_REMOVETAB, unsafe.Pointer(p.info()), unsafe.Pointer(&index), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *TabWidget) Clear() {
	_drv_ch(CLASSID_TABWIDGET, _ID_TABWIDGET_CLEAR, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *TabWidget) IndexOf(w IWidget) (index int) {
	_drv_ch(CLASSID_TABWIDGET, _ID_TABWIDGET_INDEXOF, unsafe.Pointer(p.info()), unsafe.Pointer(w.(iobj).info()), unsafe.Pointer(&index), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *TabWidget) WidgetOf(index int) (w IWidget) {
	var oi_w obj_info
	_drv_ch(CLASSID_TABWIDGET, _ID_TABWIDGET_WIDGETOF, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer(&oi_w), nil, nil, nil, nil, nil, nil, nil)
	if oi_w.native != 0 {
		item := FindObject(oi_w.native)
		if item != nil {
			w = item.(IWidget)
		}
	}
	return
}

func (p *TabWidget) SetTabText(index int, label string) {
	_drv_ch(CLASSID_TABWIDGET, _ID_TABWIDGET_SETTABTEXT, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer((*string_info)(unsafe.Pointer(&label))), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *TabWidget) TabText(index int) (label string) {
	var sh_label utf8_info
	_drv_ch(CLASSID_TABWIDGET, _ID_TABWIDGET_TABTEXT, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_label))), nil, nil, nil, nil, nil, nil, nil)
	label = sh_label.String()
	return
}

func (p *TabWidget) SetTabIcon(index int, icon *Icon) {
	_drv_ch(CLASSID_TABWIDGET, _ID_TABWIDGET_SETTABICON, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer(icon), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *TabWidget) TabIcon(index int) (icon *Icon) {
	var oi_icon obj_info
	_drv_ch(CLASSID_TABWIDGET, _ID_TABWIDGET_TABICON, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer(&oi_icon), nil, nil, nil, nil, nil, nil, nil)
	if oi_icon.native != 0 {
		v := FindObject(oi_icon.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ICON, oi_icon.native)
		}
		if v != nil {
			icon = v.(*Icon)
		}
	}
	return
}

func (p *TabWidget) SetTabToolTip(index int, tip string) {
	_drv_ch(CLASSID_TABWIDGET, _ID_TABWIDGET_SETTABTOOLTIP, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer((*string_info)(unsafe.Pointer(&tip))), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *TabWidget) TabToolTip(index int) (tip string) {
	var sh_tip utf8_info
	_drv_ch(CLASSID_TABWIDGET, _ID_TABWIDGET_TABTOOLTIP, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_tip))), nil, nil, nil, nil, nil, nil, nil)
	tip = sh_tip.String()
	return
}

func (p *TabWidget) OnCurrentChanged(fn func(int)) {
	_drv_event_ch(CLASSID_TABWIDGET, _ID_TABWIDGET_ONCURRENTCHANGED, p, fn)
	return
}

// struct ToolBox
//
type ToolBox struct {
	Widget
}

func (p *ToolBox) Name() string {
	return "ToolBox"
}

func (p *ToolBox) String() string {
	return DumpObject(p)
}
func (p *ToolBox) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "currentindex":
		if v, ok := value.(int); ok {
			p.SetCurrentIndex(v)
			return true
		}
		return false
	case "currentwidget":
		if v, ok := value.(IWidget); ok {
			p.SetCurrentWidget(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr, value)
	}
	return false
}
func (p *ToolBox) Attr(attr string) interface{} {
	switch attr {
	case "currentindex":
		return p.CurrentIndex()
	case "currentwidget":
		return p.CurrentWidget()
	case "count":
		return p.Count()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewToolBox() *ToolBox {
	return new(ToolBox).Init()
}

func (p *ToolBox) Init() *ToolBox {
	p.classid = CLASSID_TOOLBOX
	_drv_ch(CLASSID_TOOLBOX, _ID_TOOLBOX_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *ToolBox) SetCurrentIndex(index int) {
	_drv_ch(CLASSID_TOOLBOX, _ID_TOOLBOX_SETCURRENTINDEX, unsafe.Pointer(p.info()), unsafe.Pointer(&index), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolBox) SetCurrentWidget(widget IWidget) {
	_drv_ch(CLASSID_TOOLBOX, _ID_TOOLBOX_SETCURRENTWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(widget.(iobj).info()), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolBox) CurrentIndex() (n int) {
	_drv_ch(CLASSID_TOOLBOX, _ID_TOOLBOX_CURRENTINDEX, unsafe.Pointer(p.info()), unsafe.Pointer(&n), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolBox) CurrentWidget() (widget IWidget) {
	var oi_widget obj_info
	_drv_ch(CLASSID_TOOLBOX, _ID_TOOLBOX_CURRENTWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_widget), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_widget.native != 0 {
		item := FindObject(oi_widget.native)
		if item != nil {
			widget = item.(IWidget)
		}
	}
	return
}

func (p *ToolBox) Count() (n int) {
	_drv_ch(CLASSID_TOOLBOX, _ID_TOOLBOX_COUNT, unsafe.Pointer(p.info()), unsafe.Pointer(&n), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolBox) AddItem(widget IWidget, text string, icon *Icon) {
	_drv_ch(CLASSID_TOOLBOX, _ID_TOOLBOX_ADDITEM, unsafe.Pointer(p.info()), unsafe.Pointer(widget.(iobj).info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), unsafe.Pointer(icon), nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolBox) InsertItem(index int, widget IWidget, text string, icon *Icon) {
	_drv_ch(CLASSID_TOOLBOX, _ID_TOOLBOX_INSERTITEM, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer(widget.(iobj).info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), unsafe.Pointer(icon), nil, nil, nil, nil, nil)
	return
}

func (p *ToolBox) RemoveItem(index int) {
	_drv_ch(CLASSID_TOOLBOX, _ID_TOOLBOX_REMOVEITEM, unsafe.Pointer(p.info()), unsafe.Pointer(&index), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolBox) IndexOf(widget IWidget) (index int) {
	_drv_ch(CLASSID_TOOLBOX, _ID_TOOLBOX_INDEXOF, unsafe.Pointer(p.info()), unsafe.Pointer(widget.(iobj).info()), unsafe.Pointer(&index), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolBox) WidgetOf(index int) (widget IWidget) {
	var oi_widget obj_info
	_drv_ch(CLASSID_TOOLBOX, _ID_TOOLBOX_WIDGETOF, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer(&oi_widget), nil, nil, nil, nil, nil, nil, nil)
	if oi_widget.native != 0 {
		item := FindObject(oi_widget.native)
		if item != nil {
			widget = item.(IWidget)
		}
	}
	return
}

func (p *ToolBox) SetItemText(index int, text string) {
	_drv_ch(CLASSID_TOOLBOX, _ID_TOOLBOX_SETITEMTEXT, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolBox) ItemText(index int) (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_TOOLBOX, _ID_TOOLBOX_ITEMTEXT, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *ToolBox) SetItemIcon(index int, icon *Icon) {
	_drv_ch(CLASSID_TOOLBOX, _ID_TOOLBOX_SETITEMICON, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer(icon), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolBox) ItemIcon(index int) (icon *Icon) {
	var oi_icon obj_info
	_drv_ch(CLASSID_TOOLBOX, _ID_TOOLBOX_ITEMICON, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer(&oi_icon), nil, nil, nil, nil, nil, nil, nil)
	if oi_icon.native != 0 {
		v := FindObject(oi_icon.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ICON, oi_icon.native)
		}
		if v != nil {
			icon = v.(*Icon)
		}
	}
	return
}

func (p *ToolBox) SetItemToolTip(index int, text string) {
	_drv_ch(CLASSID_TOOLBOX, _ID_TOOLBOX_SETITEMTOOLTIP, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolBox) ItemToolTip(index int) (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_TOOLBOX, _ID_TOOLBOX_ITEMTOOLTIP, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *ToolBox) IsItemEnabled(index int) (b bool) {
	var b_b int
	_drv_ch(CLASSID_TOOLBOX, _ID_TOOLBOX_ISITEMENABLED, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *ToolBox) OnCurrentChanged(fn func(int)) {
	_drv_event_ch(CLASSID_TOOLBOX, _ID_TOOLBOX_ONCURRENTCHANGED, p, fn)
	return
}

// struct baseLayout
//
type baseLayout struct {
	object
}

func (p *baseLayout) Name() string {
	return "baseLayout"
}

func (p *baseLayout) String() string {
	return DumpObject(p)
}
func (p *baseLayout) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "spacing":
		if v, ok := value.(int); ok {
			p.SetSpacing(v)
			return true
		}
		return false
	case "margins":
		if v, ok := value.(Margins); ok {
			p.SetMargins(v)
			return true
		}
		return false
	case "menubar":
		if v, ok := value.(IWidget); ok {
			p.SetMenuBar(v)
			return true
		}
		return false
	default:
		return p.object.SetAttr(attr, value)
	}
	return false
}
func (p *baseLayout) Attr(attr string) interface{} {
	switch attr {
	case "spacing":
		return p.Spacing()
	case "margins":
		return p.Margins()
	case "menubar":
		return p.MenuBar()
	case "count":
		return p.Count()
	default:
		return p.object.Attr(attr)
	}
	return nil
}
func (p *baseLayout) SetSpacing(spac int) {
	_drv_ch(_CLASSID_BASELAYOUT, _ID_BASELAYOUT_SETSPACING, unsafe.Pointer(p.info()), unsafe.Pointer(&spac), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseLayout) Spacing() (spac int) {
	_drv_ch(_CLASSID_BASELAYOUT, _ID_BASELAYOUT_SPACING, unsafe.Pointer(p.info()), unsafe.Pointer(&spac), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseLayout) SetMargins(m Margins) {
	_drv_ch(_CLASSID_BASELAYOUT, _ID_BASELAYOUT_SETMARGINS, unsafe.Pointer(p.info()), unsafe.Pointer(&m), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseLayout) Margins() (m Margins) {
	_drv_ch(_CLASSID_BASELAYOUT, _ID_BASELAYOUT_MARGINS, unsafe.Pointer(p.info()), unsafe.Pointer(&m), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseLayout) SetMenuBar(widget IWidget) {
	_drv_ch(_CLASSID_BASELAYOUT, _ID_BASELAYOUT_SETMENUBAR, unsafe.Pointer(p.info()), unsafe.Pointer(widget.(iobj).info()), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseLayout) MenuBar() (widget IWidget) {
	var oi_widget obj_info
	_drv_ch(_CLASSID_BASELAYOUT, _ID_BASELAYOUT_MENUBAR, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_widget), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_widget.native != 0 {
		item := FindObject(oi_widget.native)
		if item != nil {
			widget = item.(IWidget)
		}
	}
	return
}

func (p *baseLayout) Count() (count int) {
	_drv_ch(_CLASSID_BASELAYOUT, _ID_BASELAYOUT_COUNT, unsafe.Pointer(p.info()), unsafe.Pointer(&count), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseLayout) AddLayout(layout ILayout) {
	_drv_ch(_CLASSID_BASELAYOUT, _ID_BASELAYOUT_ADDLAYOUT, unsafe.Pointer(p.info()), unsafe.Pointer(layout.(iobj).info()), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseLayout) AddWidget(widget IWidget) {
	_drv_ch(_CLASSID_BASELAYOUT, _ID_BASELAYOUT_ADDWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(widget.(iobj).info()), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseLayout) RemoveWidget(widget IWidget) {
	_drv_ch(_CLASSID_BASELAYOUT, _ID_BASELAYOUT_REMOVEWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(widget.(iobj).info()), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseLayout) IndexOf(widget IWidget) (index int) {
	_drv_ch(_CLASSID_BASELAYOUT, _ID_BASELAYOUT_INDEXOF, unsafe.Pointer(p.info()), unsafe.Pointer(widget.(iobj).info()), unsafe.Pointer(&index), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseLayout) SetMargin(v int) {
	p.SetMargins(Mr(v, v, v, v))
}
func (p *baseLayout) SetMarginsv(left, top, right, bottom int) {
	p.SetMargins(Mr(left, top, right, bottom))
}

func (p *baseLayout) Marginsv() (int, int, int, int) {
	return p.Margins().Unpack()
}

// struct BoxLayout
//
type BoxLayout struct {
	baseLayout
}

func (p *BoxLayout) Name() string {
	return "BoxLayout"
}

func (p *BoxLayout) String() string {
	return DumpObject(p)
}
func (p *BoxLayout) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "orientation":
		if v, ok := value.(Orientation); ok {
			p.SetOrientation(v)
			return true
		}
		return false
	default:
		return p.baseLayout.SetAttr(attr, value)
	}
	return false
}
func (p *BoxLayout) Attr(attr string) interface{} {
	switch attr {
	case "orientation":
		return p.Orientation()
	default:
		return p.baseLayout.Attr(attr)
	}
	return nil
}
func NewBoxLayout() *BoxLayout {
	return new(BoxLayout).Init()
}

func (p *BoxLayout) Init() *BoxLayout {
	p.classid = CLASSID_BOXLAYOUT
	_drv_ch(CLASSID_BOXLAYOUT, _ID_BOXLAYOUT_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *BoxLayout) Close() (err error) {
	if p == nil || p.native == 0 {
		return
	}
	_drv_ch(CLASSID_BOXLAYOUT, _ID_BOXLAYOUT_CLOSE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	RemoveObject(p.native)
	p.native = 0
	return
}

func (p *BoxLayout) SetOrientation(value Orientation) {
	_drv_ch(CLASSID_BOXLAYOUT, _ID_BOXLAYOUT_SETORIENTATION, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *BoxLayout) Orientation() (value Orientation) {
	_drv_ch(CLASSID_BOXLAYOUT, _ID_BOXLAYOUT_ORIENTATION, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *BoxLayout) AddLayoutWith(layout ILayout, stetch int) {
	_drv_ch(CLASSID_BOXLAYOUT, _ID_BOXLAYOUT_ADDLAYOUTWITH, unsafe.Pointer(p.info()), unsafe.Pointer(layout.(iobj).info()), unsafe.Pointer(&stetch), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *BoxLayout) AddWidgetWith(widget IWidget, stretch int, alignment Alignment) {
	_drv_ch(CLASSID_BOXLAYOUT, _ID_BOXLAYOUT_ADDWIDGETWITH, unsafe.Pointer(p.info()), unsafe.Pointer(widget.(iobj).info()), unsafe.Pointer(&stretch), unsafe.Pointer(&alignment), nil, nil, nil, nil, nil, nil)
	return
}

func (p *BoxLayout) AddSpacing(size int) {
	_drv_ch(CLASSID_BOXLAYOUT, _ID_BOXLAYOUT_ADDSPACING, unsafe.Pointer(p.info()), unsafe.Pointer(&size), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *BoxLayout) AddStretch(size int) {
	_drv_ch(CLASSID_BOXLAYOUT, _ID_BOXLAYOUT_ADDSTRETCH, unsafe.Pointer(p.info()), unsafe.Pointer(&size), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

// struct HBoxLayout
//
type HBoxLayout struct {
	BoxLayout
}

func (p *HBoxLayout) Name() string {
	return "HBoxLayout"
}

func (p *HBoxLayout) String() string {
	return DumpObject(p)
}
func (p *HBoxLayout) SetAttr(attr string, value interface{}) bool {
	switch attr {
	default:
		return p.BoxLayout.SetAttr(attr, value)
	}
	return false
}
func (p *HBoxLayout) Attr(attr string) interface{} {
	switch attr {
	default:
		return p.BoxLayout.Attr(attr)
	}
	return nil
}
func NewHBoxLayout() *HBoxLayout {
	return new(HBoxLayout).Init()
}

func (p *HBoxLayout) Init() *HBoxLayout {
	p.classid = CLASSID_HBOXLAYOUT
	_drv_ch(CLASSID_HBOXLAYOUT, _ID_HBOXLAYOUT_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

// struct VBoxLayout
//
type VBoxLayout struct {
	BoxLayout
}

func (p *VBoxLayout) Name() string {
	return "VBoxLayout"
}

func (p *VBoxLayout) String() string {
	return DumpObject(p)
}
func (p *VBoxLayout) SetAttr(attr string, value interface{}) bool {
	switch attr {
	default:
		return p.BoxLayout.SetAttr(attr, value)
	}
	return false
}
func (p *VBoxLayout) Attr(attr string) interface{} {
	switch attr {
	default:
		return p.BoxLayout.Attr(attr)
	}
	return nil
}
func NewVBoxLayout() *VBoxLayout {
	return new(VBoxLayout).Init()
}

func (p *VBoxLayout) Init() *VBoxLayout {
	p.classid = CLASSID_VBOXLAYOUT
	_drv_ch(CLASSID_VBOXLAYOUT, _ID_VBOXLAYOUT_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

// struct StackedLayout
//
type StackedLayout struct {
	baseLayout
}

func (p *StackedLayout) Name() string {
	return "StackedLayout"
}

func (p *StackedLayout) String() string {
	return DumpObject(p)
}
func (p *StackedLayout) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "currentindex":
		if v, ok := value.(int); ok {
			p.SetCurrentIndex(v)
			return true
		}
		return false
	case "currentwidget":
		if v, ok := value.(IWidget); ok {
			p.SetCurrentWidget(v)
			return true
		}
		return false
	default:
		return p.baseLayout.SetAttr(attr, value)
	}
	return false
}
func (p *StackedLayout) Attr(attr string) interface{} {
	switch attr {
	case "currentindex":
		return p.CurrentIndex()
	case "currentwidget":
		return p.CurrentWidget()
	default:
		return p.baseLayout.Attr(attr)
	}
	return nil
}
func NewStackedLayout() *StackedLayout {
	return new(StackedLayout).Init()
}

func (p *StackedLayout) Init() *StackedLayout {
	p.classid = CLASSID_STACKEDLAYOUT
	_drv_ch(CLASSID_STACKEDLAYOUT, _ID_STACKEDLAYOUT_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *StackedLayout) SetCurrentIndex(index int) {
	_drv_ch(CLASSID_STACKEDLAYOUT, _ID_STACKEDLAYOUT_SETCURRENTINDEX, unsafe.Pointer(p.info()), unsafe.Pointer(&index), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *StackedLayout) CurrentIndex() (index int) {
	_drv_ch(CLASSID_STACKEDLAYOUT, _ID_STACKEDLAYOUT_CURRENTINDEX, unsafe.Pointer(p.info()), unsafe.Pointer(&index), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *StackedLayout) SetCurrentWidget(widget IWidget) {
	_drv_ch(CLASSID_STACKEDLAYOUT, _ID_STACKEDLAYOUT_SETCURRENTWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(widget.(iobj).info()), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *StackedLayout) CurrentWidget() (widget IWidget) {
	var oi_widget obj_info
	_drv_ch(CLASSID_STACKEDLAYOUT, _ID_STACKEDLAYOUT_CURRENTWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_widget), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_widget.native != 0 {
		item := FindObject(oi_widget.native)
		if item != nil {
			widget = item.(IWidget)
		}
	}
	return
}

func (p *StackedLayout) AddWidget(widget IWidget) {
	_drv_ch(CLASSID_STACKEDLAYOUT, _ID_STACKEDLAYOUT_ADDWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(widget.(iobj).info()), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *StackedLayout) InsertWidget(index int, widget IWidget) {
	_drv_ch(CLASSID_STACKEDLAYOUT, _ID_STACKEDLAYOUT_INSERTWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer(widget.(iobj).info()), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *StackedLayout) Widget(index int) (widget IWidget) {
	var oi_widget obj_info
	_drv_ch(CLASSID_STACKEDLAYOUT, _ID_STACKEDLAYOUT_WIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer(&oi_widget), nil, nil, nil, nil, nil, nil, nil)
	if oi_widget.native != 0 {
		item := FindObject(oi_widget.native)
		if item != nil {
			widget = item.(IWidget)
		}
	}
	return
}

func (p *StackedLayout) OnCurrentChanged(fn func(int)) {
	_drv_event_ch(CLASSID_STACKEDLAYOUT, _ID_STACKEDLAYOUT_ONCURRENTCHANGED, p, fn)
	return
}

// struct baseButton
//
type baseButton struct {
	Widget
}

func (p *baseButton) Name() string {
	return "baseButton"
}

func (p *baseButton) String() string {
	return DumpObject(p)
}
func (p *baseButton) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "text":
		if v, ok := value.(string); ok {
			p.SetText(v)
			return true
		}
		return false
	case "icon":
		if v, ok := value.(*Icon); ok {
			p.SetIcon(v)
			return true
		}
		return false
	case "iconsize":
		if v, ok := value.(Size); ok {
			p.SetIconSize(v)
			return true
		}
		return false
	case "checkable":
		if v, ok := value.(bool); ok {
			p.SetCheckable(v)
			return true
		}
		return false
	case "down":
		if v, ok := value.(bool); ok {
			p.SetDown(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr, value)
	}
	return false
}
func (p *baseButton) Attr(attr string) interface{} {
	switch attr {
	case "text":
		return p.Text()
	case "icon":
		return p.Icon()
	case "iconsize":
		return p.IconSize()
	case "checkable":
		return p.IsCheckable()
	case "down":
		return p.IsDown()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func (p *baseButton) SetText(text string) {
	_drv_ch(_CLASSID_BASEBUTTON, _ID_BASEBUTTON_SETTEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseButton) Text() (text string) {
	var sh_text utf8_info
	_drv_ch(_CLASSID_BASEBUTTON, _ID_BASEBUTTON_TEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *baseButton) SetIcon(icon *Icon) {
	_drv_ch(_CLASSID_BASEBUTTON, _ID_BASEBUTTON_SETICON, unsafe.Pointer(p.info()), unsafe.Pointer(icon), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseButton) Icon() (icon *Icon) {
	var oi_icon obj_info
	_drv_ch(_CLASSID_BASEBUTTON, _ID_BASEBUTTON_ICON, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_icon), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_icon.native != 0 {
		v := FindObject(oi_icon.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ICON, oi_icon.native)
		}
		if v != nil {
			icon = v.(*Icon)
		}
	}
	return
}

func (p *baseButton) SetIconSize(sz Size) {
	_drv_ch(_CLASSID_BASEBUTTON, _ID_BASEBUTTON_SETICONSIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&sz), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseButton) IconSize() (sz Size) {
	_drv_ch(_CLASSID_BASEBUTTON, _ID_BASEBUTTON_ICONSIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&sz), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseButton) SetCheckable(b bool) {
	_drv_ch(_CLASSID_BASEBUTTON, _ID_BASEBUTTON_SETCHECKABLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseButton) IsCheckable() (b bool) {
	var b_b int
	_drv_ch(_CLASSID_BASEBUTTON, _ID_BASEBUTTON_ISCHECKABLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *baseButton) SetDown(b bool) {
	_drv_ch(_CLASSID_BASEBUTTON, _ID_BASEBUTTON_SETDOWN, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseButton) IsDown() (b bool) {
	var b_b int
	_drv_ch(_CLASSID_BASEBUTTON, _ID_BASEBUTTON_ISDOWN, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

// struct Button
//
type Button struct {
	baseButton
}

func (p *Button) Name() string {
	return "Button"
}

func (p *Button) String() string {
	return DumpObject(p)
}
func (p *Button) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "flat":
		if v, ok := value.(bool); ok {
			p.SetFlat(v)
			return true
		}
		return false
	case "default":
		if v, ok := value.(bool); ok {
			p.SetDefault(v)
			return true
		}
		return false
	case "menu":
		if v, ok := value.(*Menu); ok {
			p.SetMenu(v)
			return true
		}
		return false
	default:
		return p.baseButton.SetAttr(attr, value)
	}
	return false
}
func (p *Button) Attr(attr string) interface{} {
	switch attr {
	case "flat":
		return p.IsFlat()
	case "default":
		return p.IsDefault()
	case "menu":
		return p.Menu()
	default:
		return p.baseButton.Attr(attr)
	}
	return nil
}
func NewButton() *Button {
	return new(Button).Init()
}

func (p *Button) Init() *Button {
	p.classid = CLASSID_BUTTON
	_drv_ch(CLASSID_BUTTON, _ID_BUTTON_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewButtonWithText(text string) *Button {
	return new(Button).InitWithText(text)
}

func (p *Button) InitWithText(text string) *Button {
	p.classid = CLASSID_BUTTON
	_drv_ch(CLASSID_BUTTON, _ID_BUTTON_INITWITHTEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *Button) SetFlat(b bool) {
	_drv_ch(CLASSID_BUTTON, _ID_BUTTON_SETFLAT, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Button) IsFlat() (b bool) {
	var b_b int
	_drv_ch(CLASSID_BUTTON, _ID_BUTTON_ISFLAT, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Button) SetDefault(b bool) {
	_drv_ch(CLASSID_BUTTON, _ID_BUTTON_SETDEFAULT, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Button) IsDefault() (b bool) {
	var b_b int
	_drv_ch(CLASSID_BUTTON, _ID_BUTTON_ISDEFAULT, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Button) SetMenu(menu *Menu) {
	_drv_ch(CLASSID_BUTTON, _ID_BUTTON_SETMENU, unsafe.Pointer(p.info()), unsafe.Pointer(menu), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Button) Menu() (menu *Menu) {
	var oi_menu obj_info
	_drv_ch(CLASSID_BUTTON, _ID_BUTTON_MENU, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_menu), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_menu.native != 0 {
		v := FindObject(oi_menu.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_MENU, oi_menu.native)
		}
		if v != nil {
			menu = v.(*Menu)
		}
	}
	return
}

func (p *Button) OnClicked(fn func()) {
	_drv_event_ch(CLASSID_BUTTON, _ID_BUTTON_ONCLICKED, p, fn)
	return
}

// struct CheckBox
//
type CheckBox struct {
	baseButton
}

func (p *CheckBox) Name() string {
	return "CheckBox"
}

func (p *CheckBox) String() string {
	return DumpObject(p)
}
func (p *CheckBox) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "check":
		if v, ok := value.(int); ok {
			p.SetCheck(v)
			return true
		}
		return false
	case "tristate":
		if v, ok := value.(bool); ok {
			p.SetTristate(v)
			return true
		}
		return false
	default:
		return p.baseButton.SetAttr(attr, value)
	}
	return false
}
func (p *CheckBox) Attr(attr string) interface{} {
	switch attr {
	case "check":
		return p.Check()
	case "tristate":
		return p.IsTristate()
	default:
		return p.baseButton.Attr(attr)
	}
	return nil
}
func NewCheckBox() *CheckBox {
	return new(CheckBox).Init()
}

func (p *CheckBox) Init() *CheckBox {
	p.classid = CLASSID_CHECKBOX
	_drv_ch(CLASSID_CHECKBOX, _ID_CHECKBOX_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewCheckBoxWithText(text string) *CheckBox {
	return new(CheckBox).InitWithText(text)
}

func (p *CheckBox) InitWithText(text string) *CheckBox {
	p.classid = CLASSID_CHECKBOX
	_drv_ch(CLASSID_CHECKBOX, _ID_CHECKBOX_INITWITHTEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *CheckBox) SetCheck(state int) {
	_drv_ch(CLASSID_CHECKBOX, _ID_CHECKBOX_SETCHECK, unsafe.Pointer(p.info()), unsafe.Pointer(&state), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *CheckBox) Check() (state int) {
	_drv_ch(CLASSID_CHECKBOX, _ID_CHECKBOX_CHECK, unsafe.Pointer(p.info()), unsafe.Pointer(&state), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *CheckBox) SetTristate(b bool) {
	_drv_ch(CLASSID_CHECKBOX, _ID_CHECKBOX_SETTRISTATE, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *CheckBox) IsTristate() (b bool) {
	var b_b int
	_drv_ch(CLASSID_CHECKBOX, _ID_CHECKBOX_ISTRISTATE, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *CheckBox) OnStateChanged(fn func(int)) {
	_drv_event_ch(CLASSID_CHECKBOX, _ID_CHECKBOX_ONSTATECHANGED, p, fn)
	return
}

// struct Radio
//
type Radio struct {
	baseButton
}

func (p *Radio) Name() string {
	return "Radio"
}

func (p *Radio) String() string {
	return DumpObject(p)
}
func (p *Radio) SetAttr(attr string, value interface{}) bool {
	switch attr {
	default:
		return p.baseButton.SetAttr(attr, value)
	}
	return false
}
func (p *Radio) Attr(attr string) interface{} {
	switch attr {
	default:
		return p.baseButton.Attr(attr)
	}
	return nil
}
func NewRadio() *Radio {
	return new(Radio).Init()
}

func (p *Radio) Init() *Radio {
	p.classid = CLASSID_RADIO
	_drv_ch(CLASSID_RADIO, _ID_RADIO_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewRadioWithText(text string) *Radio {
	return new(Radio).InitWithText(text)
}

func (p *Radio) InitWithText(text string) *Radio {
	p.classid = CLASSID_RADIO
	_drv_ch(CLASSID_RADIO, _ID_RADIO_INITWITHTEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *Radio) OnClicked(fn func()) {
	_drv_event_ch(CLASSID_RADIO, _ID_RADIO_ONCLICKED, p, fn)
	return
}

// struct ToolButton
//
type ToolButton struct {
	baseButton
}

func (p *ToolButton) Name() string {
	return "ToolButton"
}

func (p *ToolButton) String() string {
	return DumpObject(p)
}
func (p *ToolButton) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "menu":
		if v, ok := value.(*Menu); ok {
			p.SetMenu(v)
			return true
		}
		return false
	case "autoraise":
		if v, ok := value.(bool); ok {
			p.SetAutoRaise(v)
			return true
		}
		return false
	case "toolbuttonstyle":
		if v, ok := value.(ToolButtonStyle); ok {
			p.SetToolButtonStyle(v)
			return true
		}
		return false
	case "toolbuttonpopupmode":
		if v, ok := value.(ToolButtonPopupMode); ok {
			p.SetToolButtonPopupMode(v)
			return true
		}
		return false
	default:
		return p.baseButton.SetAttr(attr, value)
	}
	return false
}
func (p *ToolButton) Attr(attr string) interface{} {
	switch attr {
	case "menu":
		return p.Menu()
	case "autoraise":
		return p.AutoRaise()
	case "toolbuttonstyle":
		return p.ToolButtonStyle()
	case "toolbuttonpopupmode":
		return p.ToolButtonPopupMode()
	default:
		return p.baseButton.Attr(attr)
	}
	return nil
}
func NewToolButton() *ToolButton {
	return new(ToolButton).Init()
}

func (p *ToolButton) Init() *ToolButton {
	p.classid = CLASSID_TOOLBUTTON
	_drv_ch(CLASSID_TOOLBUTTON, _ID_TOOLBUTTON_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewToolButtonWithText(text string) *ToolButton {
	return new(ToolButton).InitWithText(text)
}

func (p *ToolButton) InitWithText(text string) *ToolButton {
	p.classid = CLASSID_TOOLBUTTON
	_drv_ch(CLASSID_TOOLBUTTON, _ID_TOOLBUTTON_INITWITHTEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *ToolButton) SetMenu(menu *Menu) {
	_drv_ch(CLASSID_TOOLBUTTON, _ID_TOOLBUTTON_SETMENU, unsafe.Pointer(p.info()), unsafe.Pointer(menu), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolButton) Menu() (menu *Menu) {
	var oi_menu obj_info
	_drv_ch(CLASSID_TOOLBUTTON, _ID_TOOLBUTTON_MENU, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_menu), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_menu.native != 0 {
		v := FindObject(oi_menu.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_MENU, oi_menu.native)
		}
		if v != nil {
			menu = v.(*Menu)
		}
	}
	return
}

func (p *ToolButton) SetAutoRaise(b bool) {
	_drv_ch(CLASSID_TOOLBUTTON, _ID_TOOLBUTTON_SETAUTORAISE, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolButton) AutoRaise() (b bool) {
	var b_b int
	_drv_ch(CLASSID_TOOLBUTTON, _ID_TOOLBUTTON_AUTORAISE, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *ToolButton) SetToolButtonStyle(style ToolButtonStyle) {
	_drv_ch(CLASSID_TOOLBUTTON, _ID_TOOLBUTTON_SETTOOLBUTTONSTYLE, unsafe.Pointer(p.info()), unsafe.Pointer(&style), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolButton) ToolButtonStyle() (style ToolButtonStyle) {
	_drv_ch(CLASSID_TOOLBUTTON, _ID_TOOLBUTTON_TOOLBUTTONSTYLE, unsafe.Pointer(p.info()), unsafe.Pointer(&style), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolButton) SetToolButtonPopupMode(mode ToolButtonPopupMode) {
	_drv_ch(CLASSID_TOOLBUTTON, _ID_TOOLBUTTON_SETTOOLBUTTONPOPUPMODE, unsafe.Pointer(p.info()), unsafe.Pointer(&mode), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolButton) ToolButtonPopupMode() (mode ToolButtonPopupMode) {
	_drv_ch(CLASSID_TOOLBUTTON, _ID_TOOLBUTTON_TOOLBUTTONPOPUPMODE, unsafe.Pointer(p.info()), unsafe.Pointer(&mode), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ToolButton) OnClicked(fn func()) {
	_drv_event_ch(CLASSID_TOOLBUTTON, _ID_TOOLBUTTON_ONCLICKED, p, fn)
	return
}

// struct Frame
//
type Frame struct {
	Widget
}

func (p *Frame) Name() string {
	return "Frame"
}

func (p *Frame) String() string {
	return DumpObject(p)
}
func (p *Frame) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "framestyle":
		if v, ok := value.(int); ok {
			p.SetFrameStyle(v)
			return true
		}
		return false
	case "framerect":
		if v, ok := value.(Rect); ok {
			p.SetFrameRect(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr, value)
	}
	return false
}
func (p *Frame) Attr(attr string) interface{} {
	switch attr {
	case "framestyle":
		return p.FrameStyle()
	case "framerect":
		return p.FrameRect()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewFrame() *Frame {
	return new(Frame).Init()
}

func (p *Frame) Init() *Frame {
	p.classid = CLASSID_FRAME
	_drv_ch(CLASSID_FRAME, _ID_FRAME_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *Frame) SetFrameStyle(style int) {
	_drv_ch(CLASSID_FRAME, _ID_FRAME_SETFRAMESTYLE, unsafe.Pointer(p.info()), unsafe.Pointer(&style), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Frame) FrameStyle() (style int) {
	_drv_ch(CLASSID_FRAME, _ID_FRAME_FRAMESTYLE, unsafe.Pointer(p.info()), unsafe.Pointer(&style), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Frame) SetFrameRect(rc Rect) {
	_drv_ch(CLASSID_FRAME, _ID_FRAME_SETFRAMERECT, unsafe.Pointer(p.info()), unsafe.Pointer(&rc), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Frame) FrameRect() (rc Rect) {
	_drv_ch(CLASSID_FRAME, _ID_FRAME_FRAMERECT, unsafe.Pointer(p.info()), unsafe.Pointer(&rc), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Frame) SetFrameRectv(x, y, width, height int) {
	p.SetFrameRect(Rc(x, y, width, height))
}

func (p *Frame) FrameRectv() (int, int, int, int) {
	return p.FrameRect().Unpack()
}

// struct Label
//
type Label struct {
	Frame
}

func (p *Label) Name() string {
	return "Label"
}

func (p *Label) String() string {
	return DumpObject(p)
}
func (p *Label) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "text":
		if v, ok := value.(string); ok {
			p.SetText(v)
			return true
		}
		return false
	case "wordwrap":
		if v, ok := value.(bool); ok {
			p.SetWordWrap(v)
			return true
		}
		return false
	case "textformat":
		if v, ok := value.(int); ok {
			p.SetTextFormat(v)
			return true
		}
		return false
	case "pixmap":
		if v, ok := value.(*Pixmap); ok {
			p.SetPixmap(v)
			return true
		}
		return false
	case "scaledcontents":
		if v, ok := value.(bool); ok {
			p.SetScaledContents(v)
			return true
		}
		return false
	case "openexternallinks":
		if v, ok := value.(bool); ok {
			p.SetOpenExternalLinks(v)
			return true
		}
		return false
	case "alignment":
		if v, ok := value.(Alignment); ok {
			p.SetAlignment(v)
			return true
		}
		return false
	case "indent":
		if v, ok := value.(int); ok {
			p.SetIndent(v)
			return true
		}
		return false
	case "margin":
		if v, ok := value.(int); ok {
			p.SetMargin(v)
			return true
		}
		return false
	default:
		return p.Frame.SetAttr(attr, value)
	}
	return false
}
func (p *Label) Attr(attr string) interface{} {
	switch attr {
	case "text":
		return p.Text()
	case "wordwrap":
		return p.WordWrap()
	case "textformat":
		return p.TextFormat()
	case "pixmap":
		return p.Pixmap()
	case "scaledcontents":
		return p.HasScaledContents()
	case "openexternallinks":
		return p.OpenExternalLinks()
	case "alignment":
		return p.Alignment()
	case "indent":
		return p.Indent()
	case "margin":
		return p.Margin()
	default:
		return p.Frame.Attr(attr)
	}
	return nil
}
func NewLabel() *Label {
	return new(Label).Init()
}

func (p *Label) Init() *Label {
	p.classid = CLASSID_LABEL
	_drv_ch(CLASSID_LABEL, _ID_LABEL_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewLabelWithText(text string) *Label {
	return new(Label).InitWithText(text)
}

func (p *Label) InitWithText(text string) *Label {
	p.classid = CLASSID_LABEL
	_drv_ch(CLASSID_LABEL, _ID_LABEL_INITWITHTEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewLabelWithPixmap(pixmap *Pixmap) *Label {
	return new(Label).InitWithPixmap(pixmap)
}

func (p *Label) InitWithPixmap(pixmap *Pixmap) *Label {
	p.classid = CLASSID_LABEL
	_drv_ch(CLASSID_LABEL, _ID_LABEL_INITWITHPIXMAP, unsafe.Pointer(p.info()), unsafe.Pointer(pixmap), nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *Label) SetText(text string) {
	_drv_ch(CLASSID_LABEL, _ID_LABEL_SETTEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Label) Text() (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_LABEL, _ID_LABEL_TEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *Label) SetWordWrap(b bool) {
	_drv_ch(CLASSID_LABEL, _ID_LABEL_SETWORDWRAP, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Label) WordWrap() (b bool) {
	var b_b int
	_drv_ch(CLASSID_LABEL, _ID_LABEL_WORDWRAP, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Label) SetTextFormat(format int) {
	_drv_ch(CLASSID_LABEL, _ID_LABEL_SETTEXTFORMAT, unsafe.Pointer(p.info()), unsafe.Pointer(&format), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Label) TextFormat() (format int) {
	_drv_ch(CLASSID_LABEL, _ID_LABEL_TEXTFORMAT, unsafe.Pointer(p.info()), unsafe.Pointer(&format), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Label) SetPixmap(pixmap *Pixmap) {
	_drv_ch(CLASSID_LABEL, _ID_LABEL_SETPIXMAP, unsafe.Pointer(p.info()), unsafe.Pointer(pixmap), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Label) Pixmap() (pixmap *Pixmap) {
	var oi_pixmap obj_info
	_drv_ch(CLASSID_LABEL, _ID_LABEL_PIXMAP, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_pixmap), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_pixmap.native != 0 {
		v := FindObject(oi_pixmap.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_PIXMAP, oi_pixmap.native)
		}
		if v != nil {
			pixmap = v.(*Pixmap)
		}
	}
	return
}

func (p *Label) SetScaledContents(b bool) {
	_drv_ch(CLASSID_LABEL, _ID_LABEL_SETSCALEDCONTENTS, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Label) HasScaledContents() (b bool) {
	var b_b int
	_drv_ch(CLASSID_LABEL, _ID_LABEL_HASSCALEDCONTENTS, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Label) SetOpenExternalLinks(b bool) {
	_drv_ch(CLASSID_LABEL, _ID_LABEL_SETOPENEXTERNALLINKS, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Label) OpenExternalLinks() (b bool) {
	var b_b int
	_drv_ch(CLASSID_LABEL, _ID_LABEL_OPENEXTERNALLINKS, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Label) SetAlignment(a Alignment) {
	_drv_ch(CLASSID_LABEL, _ID_LABEL_SETALIGNMENT, unsafe.Pointer(p.info()), unsafe.Pointer(&a), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Label) Alignment() (a Alignment) {
	_drv_ch(CLASSID_LABEL, _ID_LABEL_ALIGNMENT, unsafe.Pointer(p.info()), unsafe.Pointer(&a), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Label) SetIndent(i int) {
	_drv_ch(CLASSID_LABEL, _ID_LABEL_SETINDENT, unsafe.Pointer(p.info()), unsafe.Pointer(&i), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Label) Indent() (i int) {
	_drv_ch(CLASSID_LABEL, _ID_LABEL_INDENT, unsafe.Pointer(p.info()), unsafe.Pointer(&i), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Label) SetMargin(i int) {
	_drv_ch(CLASSID_LABEL, _ID_LABEL_SETMARGIN, unsafe.Pointer(p.info()), unsafe.Pointer(&i), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Label) Margin() (i int) {
	_drv_ch(CLASSID_LABEL, _ID_LABEL_MARGIN, unsafe.Pointer(p.info()), unsafe.Pointer(&i), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Label) OnLinkActivated(fn func(string)) {
	_drv_event_ch(CLASSID_LABEL, _ID_LABEL_ONLINKACTIVATED, p, fn)
	return
}

// struct GroupBox
//
type GroupBox struct {
	Widget
}

func (p *GroupBox) Name() string {
	return "GroupBox"
}

func (p *GroupBox) String() string {
	return DumpObject(p)
}
func (p *GroupBox) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "title":
		if v, ok := value.(string); ok {
			p.SetTitle(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr, value)
	}
	return false
}
func (p *GroupBox) Attr(attr string) interface{} {
	switch attr {
	case "title":
		return p.Title()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewGroupBox() *GroupBox {
	return new(GroupBox).Init()
}

func (p *GroupBox) Init() *GroupBox {
	p.classid = CLASSID_GROUPBOX
	_drv_ch(CLASSID_GROUPBOX, _ID_GROUPBOX_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewGroupBoxWithTitle(text string) *GroupBox {
	return new(GroupBox).InitWithTitle(text)
}

func (p *GroupBox) InitWithTitle(text string) *GroupBox {
	p.classid = CLASSID_GROUPBOX
	_drv_ch(CLASSID_GROUPBOX, _ID_GROUPBOX_INITWITHTITLE, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *GroupBox) SetTitle(text string) {
	_drv_ch(CLASSID_GROUPBOX, _ID_GROUPBOX_SETTITLE, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *GroupBox) Title() (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_GROUPBOX, _ID_GROUPBOX_TITLE, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

// struct Dialog
//
type Dialog struct {
	Widget
}

func (p *Dialog) Name() string {
	return "Dialog"
}

func (p *Dialog) String() string {
	return DumpObject(p)
}
func (p *Dialog) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "modal":
		if v, ok := value.(bool); ok {
			p.SetModal(v)
			return true
		}
		return false
	case "result":
		if v, ok := value.(int); ok {
			p.SetResult(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr, value)
	}
	return false
}
func (p *Dialog) Attr(attr string) interface{} {
	switch attr {
	case "modal":
		return p.IsModal()
	case "result":
		return p.Result()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewDialog() *Dialog {
	return new(Dialog).Init()
}

func (p *Dialog) Init() *Dialog {
	p.classid = CLASSID_DIALOG
	_drv_ch(CLASSID_DIALOG, _ID_DIALOG_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *Dialog) SetModal(b bool) {
	_drv_ch(CLASSID_DIALOG, _ID_DIALOG_SETMODAL, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Dialog) IsModal() (b bool) {
	var b_b int
	_drv_ch(CLASSID_DIALOG, _ID_DIALOG_ISMODAL, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Dialog) SetResult(r int) {
	_drv_ch(CLASSID_DIALOG, _ID_DIALOG_SETRESULT, unsafe.Pointer(p.info()), unsafe.Pointer(&r), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Dialog) Result() (r int) {
	_drv_ch(CLASSID_DIALOG, _ID_DIALOG_RESULT, unsafe.Pointer(p.info()), unsafe.Pointer(&r), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Dialog) Open() {
	_drv_ch(CLASSID_DIALOG, _ID_DIALOG_OPEN, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Dialog) Exec() (r int) {
	_drv_ch(CLASSID_DIALOG, _ID_DIALOG_EXEC, unsafe.Pointer(p.info()), unsafe.Pointer(&r), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Dialog) Done(r int) {
	_drv_ch(CLASSID_DIALOG, _ID_DIALOG_DONE, unsafe.Pointer(p.info()), unsafe.Pointer(&r), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Dialog) Accept() {
	_drv_ch(CLASSID_DIALOG, _ID_DIALOG_ACCEPT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Dialog) Reject() {
	_drv_ch(CLASSID_DIALOG, _ID_DIALOG_REJECT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Dialog) OnAccepted(fn func()) {
	_drv_event_ch(CLASSID_DIALOG, _ID_DIALOG_ONACCEPTED, p, fn)
	return
}

func (p *Dialog) OnRejected(fn func()) {
	_drv_event_ch(CLASSID_DIALOG, _ID_DIALOG_ONREJECTED, p, fn)
	return
}

// struct ComboBox
//
type ComboBox struct {
	Widget
}

func (p *ComboBox) Name() string {
	return "ComboBox"
}

func (p *ComboBox) String() string {
	return DumpObject(p)
}
func (p *ComboBox) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "currentindex":
		if v, ok := value.(int); ok {
			p.SetCurrentIndex(v)
			return true
		}
		return false
	case "editable":
		if v, ok := value.(bool); ok {
			p.SetEditable(v)
			return true
		}
		return false
	case "maxcount":
		if v, ok := value.(int); ok {
			p.SetMaxCount(v)
			return true
		}
		return false
	case "maxvisibleitems":
		if v, ok := value.(int); ok {
			p.SetMaxVisibleItems(v)
			return true
		}
		return false
	case "minimumcontentslength":
		if v, ok := value.(int); ok {
			p.SetMinimumContentsLength(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr, value)
	}
	return false
}
func (p *ComboBox) Attr(attr string) interface{} {
	switch attr {
	case "count":
		return p.Count()
	case "currentindex":
		return p.CurrentIndex()
	case "currenttext":
		return p.CurrentText()
	case "editable":
		return p.IsEditable()
	case "maxcount":
		return p.MaxCount()
	case "maxvisibleitems":
		return p.MaxVisibleItems()
	case "minimuncontentslenght":
		return p.MinimunContentsLenght()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewComboBox() *ComboBox {
	return new(ComboBox).Init()
}

func (p *ComboBox) Init() *ComboBox {
	p.classid = CLASSID_COMBOBOX
	_drv_ch(CLASSID_COMBOBOX, _ID_COMBOBOX_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *ComboBox) Count() (count int) {
	_drv_ch(CLASSID_COMBOBOX, _ID_COMBOBOX_COUNT, unsafe.Pointer(p.info()), unsafe.Pointer(&count), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ComboBox) SetCurrentIndex(index int) {
	_drv_ch(CLASSID_COMBOBOX, _ID_COMBOBOX_SETCURRENTINDEX, unsafe.Pointer(p.info()), unsafe.Pointer(&index), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ComboBox) CurrentIndex() (index int) {
	_drv_ch(CLASSID_COMBOBOX, _ID_COMBOBOX_CURRENTINDEX, unsafe.Pointer(p.info()), unsafe.Pointer(&index), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ComboBox) CurrentText() (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_COMBOBOX, _ID_COMBOBOX_CURRENTTEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *ComboBox) SetEditable(b bool) {
	_drv_ch(CLASSID_COMBOBOX, _ID_COMBOBOX_SETEDITABLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ComboBox) IsEditable() (b bool) {
	var b_b int
	_drv_ch(CLASSID_COMBOBOX, _ID_COMBOBOX_ISEDITABLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *ComboBox) SetMaxCount(count int) {
	_drv_ch(CLASSID_COMBOBOX, _ID_COMBOBOX_SETMAXCOUNT, unsafe.Pointer(p.info()), unsafe.Pointer(&count), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ComboBox) MaxCount() (count int) {
	_drv_ch(CLASSID_COMBOBOX, _ID_COMBOBOX_MAXCOUNT, unsafe.Pointer(p.info()), unsafe.Pointer(&count), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ComboBox) SetMaxVisibleItems(max int) {
	_drv_ch(CLASSID_COMBOBOX, _ID_COMBOBOX_SETMAXVISIBLEITEMS, unsafe.Pointer(p.info()), unsafe.Pointer(&max), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ComboBox) MaxVisibleItems() (max int) {
	_drv_ch(CLASSID_COMBOBOX, _ID_COMBOBOX_MAXVISIBLEITEMS, unsafe.Pointer(p.info()), unsafe.Pointer(&max), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ComboBox) SetMinimumContentsLength(characters int) {
	_drv_ch(CLASSID_COMBOBOX, _ID_COMBOBOX_SETMINIMUMCONTENTSLENGTH, unsafe.Pointer(p.info()), unsafe.Pointer(&characters), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ComboBox) MinimunContentsLenght() (characters int) {
	_drv_ch(CLASSID_COMBOBOX, _ID_COMBOBOX_MINIMUNCONTENTSLENGHT, unsafe.Pointer(p.info()), unsafe.Pointer(&characters), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ComboBox) AddItem(text string) {
	_drv_ch(CLASSID_COMBOBOX, _ID_COMBOBOX_ADDITEM, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ComboBox) InsertItem(index int, text string) {
	_drv_ch(CLASSID_COMBOBOX, _ID_COMBOBOX_INSERTITEM, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ComboBox) RemoveItem(index int) {
	_drv_ch(CLASSID_COMBOBOX, _ID_COMBOBOX_REMOVEITEM, unsafe.Pointer(p.info()), unsafe.Pointer(&index), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ComboBox) ItemText(index int) (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_COMBOBOX, _ID_COMBOBOX_ITEMTEXT, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *ComboBox) OnCurrentIndexChanged(fn func(int)) {
	_drv_event_ch(CLASSID_COMBOBOX, _ID_COMBOBOX_ONCURRENTINDEXCHANGED, p, fn)
	return
}

// struct LineEdit
//
// InputMask:
//
// A	ASCII alphabetic character required. A-Z, a-z.
//
// a	ASCII alphabetic character permitted but not required.
//
// N	ASCII alphanumeric character required. A-Z, a-z, 0-9.
//
// n	ASCII alphanumeric character permitted but not required.
//
// X	Any character required.
//
// x	Any character permitted but not required.
//
// 9	ASCII digit required. 0-9.
//
// 0	ASCII digit permitted but not required.
//
// D	ASCII digit required. 1-9.
//
// d	ASCII digit permitted but not required (1-9).
//
// #	ASCII digit or plus/minus sign permitted but not required.
//
// H	Hexadecimal character required. A-F, a-f, 0-9.
//
// h	Hexadecimal character permitted but not required.
//
// B	Binary character required. 0-1.
//
// b	Binary character permitted but not required.
//
// >	All following alphabetic characters are uppercased.
//
// <	All following alphabetic characters are lowercased.
//
// !	Switch off case conversion.
//
// \	Use \ to escape the special characters listed above to use them as separators.
//
// InputMask Example:
//
// 000.000.000.000;_	IP address; blanks are _.
//
// HH:HH:HH:HH:HH:HH;_	MAC address
//
// 0000-00-00	ISO Date; blanks are space
//
// >AAAAA-AAAAA-AAAAA-AAAAA-AAAAA;#	License number; blanks are - and all (alphabetic) characters are converted to uppercase.
type LineEdit struct {
	Widget
}

func (p *LineEdit) Name() string {
	return "LineEdit"
}

func (p *LineEdit) String() string {
	return DumpObject(p)
}
func (p *LineEdit) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "text":
		if v, ok := value.(string); ok {
			p.SetText(v)
			return true
		}
		return false
	case "inputmask":
		if v, ok := value.(string); ok {
			p.SetInputMask(v)
			return true
		}
		return false
	case "alignment":
		if v, ok := value.(Alignment); ok {
			p.SetAlignment(v)
			return true
		}
		return false
	case "cursorpos":
		if v, ok := value.(int); ok {
			p.SetCursorPos(v)
			return true
		}
		return false
	case "dragenabled":
		if v, ok := value.(bool); ok {
			p.SetDragEnabled(v)
			return true
		}
		return false
	case "readonly":
		if v, ok := value.(bool); ok {
			p.SetReadOnly(v)
			return true
		}
		return false
	case "frame":
		if v, ok := value.(bool); ok {
			p.SetFrame(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr, value)
	}
	return false
}
func (p *LineEdit) Attr(attr string) interface{} {
	switch attr {
	case "text":
		return p.Text()
	case "inputmask":
		return p.InputMask()
	case "alignment":
		return p.Alignment()
	case "cursorpos":
		return p.CursorPos()
	case "dragenabled":
		return p.DragEnabled()
	case "readonly":
		return p.IsReadOnly()
	case "frame":
		return p.HasFrame()
	case "redoavailable":
		return p.IsRedoAvailable()
	case "selected":
		return p.HasSelected()
	case "selectedtext":
		return p.SelectedText()
	case "selstart":
		return p.SelStart()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewLineEdit() *LineEdit {
	return new(LineEdit).Init()
}

func (p *LineEdit) Init() *LineEdit {
	p.classid = CLASSID_LINEEDIT
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewLineEditWithText(text string) *LineEdit {
	return new(LineEdit).InitWithText(text)
}

func (p *LineEdit) InitWithText(text string) *LineEdit {
	p.classid = CLASSID_LINEEDIT
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_INITWITHTEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *LineEdit) SetText(text string) {
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_SETTEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *LineEdit) Text() (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_TEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *LineEdit) SetInputMask(text string) {
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_SETINPUTMASK, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *LineEdit) InputMask() (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_INPUTMASK, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *LineEdit) SetAlignment(value Alignment) {
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_SETALIGNMENT, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *LineEdit) Alignment() (value Alignment) {
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_ALIGNMENT, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *LineEdit) SetCursorPos(pos int) {
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_SETCURSORPOS, unsafe.Pointer(p.info()), unsafe.Pointer(&pos), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *LineEdit) CursorPos() (pos int) {
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_CURSORPOS, unsafe.Pointer(p.info()), unsafe.Pointer(&pos), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *LineEdit) SetDragEnabled(b bool) {
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_SETDRAGENABLED, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *LineEdit) DragEnabled() (b bool) {
	var b_b int
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_DRAGENABLED, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *LineEdit) SetReadOnly(b bool) {
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_SETREADONLY, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *LineEdit) IsReadOnly() (b bool) {
	var b_b int
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_ISREADONLY, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *LineEdit) SetFrame(b bool) {
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_SETFRAME, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *LineEdit) HasFrame() (b bool) {
	var b_b int
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_HASFRAME, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *LineEdit) IsRedoAvailable() (b bool) {
	var b_b int
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_ISREDOAVAILABLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *LineEdit) HasSelected() (b bool) {
	var b_b int
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_HASSELECTED, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *LineEdit) SelectedText() (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_SELECTEDTEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *LineEdit) SelStart() (start int) {
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_SELSTART, unsafe.Pointer(p.info()), unsafe.Pointer(&start), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *LineEdit) SetSel(start int, length int) {
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_SETSEL, unsafe.Pointer(p.info()), unsafe.Pointer(&start), unsafe.Pointer(&length), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *LineEdit) CancelSel() {
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_CANCELSEL, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *LineEdit) SelectAll() {
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_SELECTALL, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *LineEdit) Copy() {
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_COPY, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *LineEdit) Cut() {
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_CUT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *LineEdit) Paste() {
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_PASTE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *LineEdit) Redo() {
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_REDO, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *LineEdit) Undo() {
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_UNDO, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *LineEdit) Clear() {
	_drv_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_CLEAR, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *LineEdit) OnTextChanged(fn func(string)) {
	_drv_event_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_ONTEXTCHANGED, p, fn)
	return
}

func (p *LineEdit) OnEditingFinished(fn func()) {
	_drv_event_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_ONEDITINGFINISHED, p, fn)
	return
}

func (p *LineEdit) OnReturnPressed(fn func()) {
	_drv_event_ch(CLASSID_LINEEDIT, _ID_LINEEDIT_ONRETURNPRESSED, p, fn)
	return
}

// struct baseSlider
//
type baseSlider struct {
	Widget
}

func (p *baseSlider) Name() string {
	return "baseSlider"
}

func (p *baseSlider) String() string {
	return DumpObject(p)
}
func (p *baseSlider) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "tracking":
		if v, ok := value.(bool); ok {
			p.SetTracking(v)
			return true
		}
		return false
	case "maximum":
		if v, ok := value.(int); ok {
			p.SetMaximum(v)
			return true
		}
		return false
	case "minimum":
		if v, ok := value.(int); ok {
			p.SetMinimum(v)
			return true
		}
		return false
	case "orientation":
		if v, ok := value.(Orientation); ok {
			p.SetOrientation(v)
			return true
		}
		return false
	case "pagestep":
		if v, ok := value.(int); ok {
			p.SetPageStep(v)
			return true
		}
		return false
	case "singlestep":
		if v, ok := value.(int); ok {
			p.SetSingleStep(v)
			return true
		}
		return false
	case "sliderdown":
		if v, ok := value.(bool); ok {
			p.SetSliderDown(v)
			return true
		}
		return false
	case "sliderposition":
		if v, ok := value.(int); ok {
			p.SetSliderPosition(v)
			return true
		}
		return false
	case "value":
		if v, ok := value.(int); ok {
			p.SetValue(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr, value)
	}
	return false
}
func (p *baseSlider) Attr(attr string) interface{} {
	switch attr {
	case "tracking":
		return p.HasTracking()
	case "maximum":
		return p.Maximum()
	case "minimum":
		return p.Minimum()
	case "orientation":
		return p.Orientation()
	case "pagestep":
		return p.PageStep()
	case "singlestep":
		return p.SingleStep()
	case "sliderdown":
		return p.IsSliderDown()
	case "sliderposition":
		return p.SliderPosition()
	case "value":
		return p.Value()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func (p *baseSlider) SetTracking(b bool) {
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_SETTRACKING, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseSlider) HasTracking() (b bool) {
	var b_b int
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_HASTRACKING, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *baseSlider) SetMaximum(value int) {
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_SETMAXIMUM, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseSlider) Maximum() (value int) {
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_MAXIMUM, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseSlider) SetMinimum(value int) {
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_SETMINIMUM, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseSlider) Minimum() (value int) {
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_MINIMUM, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseSlider) SetOrientation(value Orientation) {
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_SETORIENTATION, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseSlider) Orientation() (value Orientation) {
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_ORIENTATION, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseSlider) SetPageStep(value int) {
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_SETPAGESTEP, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseSlider) PageStep() (value int) {
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_PAGESTEP, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseSlider) SetSingleStep(value int) {
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_SETSINGLESTEP, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseSlider) SingleStep() (value int) {
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_SINGLESTEP, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseSlider) SetSliderDown(b bool) {
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_SETSLIDERDOWN, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseSlider) IsSliderDown() (b bool) {
	var b_b int
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_ISSLIDERDOWN, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *baseSlider) SetSliderPosition(value int) {
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_SETSLIDERPOSITION, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseSlider) SliderPosition() (value int) {
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_SLIDERPOSITION, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseSlider) SetValue(value int) {
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_SETVALUE, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseSlider) Value() (value int) {
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_VALUE, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseSlider) SetRange(min int, max int) {
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_SETRANGE, unsafe.Pointer(p.info()), unsafe.Pointer(&min), unsafe.Pointer(&max), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseSlider) Range() (min int, max int) {
	_drv_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_RANGE, unsafe.Pointer(p.info()), unsafe.Pointer(&min), unsafe.Pointer(&max), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *baseSlider) OnValueChanged(fn func(int)) {
	_drv_event_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_ONVALUECHANGED, p, fn)
	return
}

func (p *baseSlider) OnSliderPressed(fn func()) {
	_drv_event_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_ONSLIDERPRESSED, p, fn)
	return
}

func (p *baseSlider) OnSliderReleased(fn func()) {
	_drv_event_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_ONSLIDERRELEASED, p, fn)
	return
}

func (p *baseSlider) OnSliderMoved(fn func(int)) {
	_drv_event_ch(_CLASSID_BASESLIDER, _ID_BASESLIDER_ONSLIDERMOVED, p, fn)
	return
}

// struct Slider
//
type Slider struct {
	baseSlider
}

func (p *Slider) Name() string {
	return "Slider"
}

func (p *Slider) String() string {
	return DumpObject(p)
}
func (p *Slider) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "tickinterval":
		if v, ok := value.(int); ok {
			p.SetTickInterval(v)
			return true
		}
		return false
	case "tickposition":
		if v, ok := value.(TickPosition); ok {
			p.SetTickPosition(v)
			return true
		}
		return false
	default:
		return p.baseSlider.SetAttr(attr, value)
	}
	return false
}
func (p *Slider) Attr(attr string) interface{} {
	switch attr {
	case "tickinterval":
		return p.TickInterval()
	case "tickposition":
		return p.TickPosition()
	default:
		return p.baseSlider.Attr(attr)
	}
	return nil
}
func NewSlider() *Slider {
	return new(Slider).Init()
}

func (p *Slider) Init() *Slider {
	p.classid = CLASSID_SLIDER
	_drv_ch(CLASSID_SLIDER, _ID_SLIDER_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *Slider) SetTickInterval(value int) {
	_drv_ch(CLASSID_SLIDER, _ID_SLIDER_SETTICKINTERVAL, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Slider) TickInterval() (value int) {
	_drv_ch(CLASSID_SLIDER, _ID_SLIDER_TICKINTERVAL, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Slider) SetTickPosition(value TickPosition) {
	_drv_ch(CLASSID_SLIDER, _ID_SLIDER_SETTICKPOSITION, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Slider) TickPosition() (value TickPosition) {
	_drv_ch(CLASSID_SLIDER, _ID_SLIDER_TICKPOSITION, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

// struct ScrollBar
//
type ScrollBar struct {
	baseSlider
}

func (p *ScrollBar) Name() string {
	return "ScrollBar"
}

func (p *ScrollBar) String() string {
	return DumpObject(p)
}
func (p *ScrollBar) SetAttr(attr string, value interface{}) bool {
	switch attr {
	default:
		return p.baseSlider.SetAttr(attr, value)
	}
	return false
}
func (p *ScrollBar) Attr(attr string) interface{} {
	switch attr {
	default:
		return p.baseSlider.Attr(attr)
	}
	return nil
}
func NewScrollBar() *ScrollBar {
	return new(ScrollBar).Init()
}

func (p *ScrollBar) Init() *ScrollBar {
	p.classid = CLASSID_SCROLLBAR
	_drv_ch(CLASSID_SCROLLBAR, _ID_SCROLLBAR_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

// struct Dial
//
type Dial struct {
	baseSlider
}

func (p *Dial) Name() string {
	return "Dial"
}

func (p *Dial) String() string {
	return DumpObject(p)
}
func (p *Dial) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "notchtarget":
		if v, ok := value.(float64); ok {
			p.SetNotchTarget(v)
			return true
		}
		return false
	case "notchesvisible":
		if v, ok := value.(bool); ok {
			p.SetNotchesVisible(v)
			return true
		}
		return false
	case "wrapping":
		if v, ok := value.(bool); ok {
			p.SetWrapping(v)
			return true
		}
		return false
	default:
		return p.baseSlider.SetAttr(attr, value)
	}
	return false
}
func (p *Dial) Attr(attr string) interface{} {
	switch attr {
	case "notchsize":
		return p.NotchSize()
	case "notchtarget":
		return p.NotchTarget()
	case "notchesvisible":
		return p.NotchesVisible()
	case "wrapping":
		return p.Wrapping()
	default:
		return p.baseSlider.Attr(attr)
	}
	return nil
}
func NewDial() *Dial {
	return new(Dial).Init()
}

func (p *Dial) Init() *Dial {
	p.classid = CLASSID_DIAL
	_drv_ch(CLASSID_DIAL, _ID_DIAL_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *Dial) NotchSize() (size int) {
	_drv_ch(CLASSID_DIAL, _ID_DIAL_NOTCHSIZE, unsafe.Pointer(p.info()), unsafe.Pointer(&size), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Dial) SetNotchTarget(value float64) {
	_drv_ch(CLASSID_DIAL, _ID_DIAL_SETNOTCHTARGET, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Dial) NotchTarget() (value float64) {
	_drv_ch(CLASSID_DIAL, _ID_DIAL_NOTCHTARGET, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Dial) SetNotchesVisible(b bool) {
	_drv_ch(CLASSID_DIAL, _ID_DIAL_SETNOTCHESVISIBLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Dial) NotchesVisible() (b bool) {
	var b_b int
	_drv_ch(CLASSID_DIAL, _ID_DIAL_NOTCHESVISIBLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *Dial) SetWrapping(b bool) {
	_drv_ch(CLASSID_DIAL, _ID_DIAL_SETWRAPPING, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Dial) Wrapping() (b bool) {
	var b_b int
	_drv_ch(CLASSID_DIAL, _ID_DIAL_WRAPPING, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

// struct Brush
//
type Brush struct {
	object
}

func (p *Brush) Name() string {
	return "Brush"
}

func (p *Brush) String() string {
	return DumpObject(p)
}
func (p *Brush) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "color":
		if v, ok := value.(color.Color); ok {
			p.SetColor(v)
			return true
		}
		return false
	case "style":
		if v, ok := value.(BrushStyle); ok {
			p.SetStyle(v)
			return true
		}
		return false
	default:
		return p.object.SetAttr(attr, value)
	}
	return false
}
func (p *Brush) Attr(attr string) interface{} {
	switch attr {
	case "color":
		return p.Color()
	case "style":
		return p.Style()
	default:
		return p.object.Attr(attr)
	}
	return nil
}
func NewBrush() *Brush {
	return new(Brush).Init()
}

func (p *Brush) Init() *Brush {
	p.classid = CLASSID_BRUSH
	_drv_ch(CLASSID_BRUSH, _ID_BRUSH_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *Brush) Close() (err error) {
	if p == nil || p.native == 0 {
		return
	}
	_drv_ch(CLASSID_BRUSH, _ID_BRUSH_CLOSE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	RemoveObject(p.native)
	p.native = 0
	return
}

func (p *Brush) SetColor(clr color.Color) {
	sh_clr := make_rgba(clr)
	_drv_ch(CLASSID_BRUSH, _ID_BRUSH_SETCOLOR, unsafe.Pointer(p.info()), unsafe.Pointer(&sh_clr), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Brush) Color() (clr color.Color) {
	var sh_clr rgba
	_drv_ch(CLASSID_BRUSH, _ID_BRUSH_COLOR, unsafe.Pointer(p.info()), unsafe.Pointer(&sh_clr), nil, nil, nil, nil, nil, nil, nil, nil)
	clr = sh_clr
	return
}

func (p *Brush) SetStyle(style BrushStyle) {
	_drv_ch(CLASSID_BRUSH, _ID_BRUSH_SETSTYLE, unsafe.Pointer(p.info()), unsafe.Pointer(&style), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Brush) Style() (style BrushStyle) {
	_drv_ch(CLASSID_BRUSH, _ID_BRUSH_STYLE, unsafe.Pointer(p.info()), unsafe.Pointer(&style), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

// struct Pen
//
type Pen struct {
	object
}

func (p *Pen) Name() string {
	return "Pen"
}

func (p *Pen) String() string {
	return DumpObject(p)
}
func (p *Pen) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "color":
		if v, ok := value.(color.Color); ok {
			p.SetColor(v)
			return true
		}
		return false
	case "width":
		if v, ok := value.(int); ok {
			p.SetWidth(v)
			return true
		}
		return false
	case "style":
		if v, ok := value.(PenStyle); ok {
			p.SetStyle(v)
			return true
		}
		return false
	default:
		return p.object.SetAttr(attr, value)
	}
	return false
}
func (p *Pen) Attr(attr string) interface{} {
	switch attr {
	case "color":
		return p.Color()
	case "width":
		return p.Width()
	case "style":
		return p.Style()
	default:
		return p.object.Attr(attr)
	}
	return nil
}
func NewPen() *Pen {
	return new(Pen).Init()
}

func (p *Pen) Init() *Pen {
	p.classid = CLASSID_PEN
	_drv_ch(CLASSID_PEN, _ID_PEN_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *Pen) Close() (err error) {
	if p == nil || p.native == 0 {
		return
	}
	_drv_ch(CLASSID_PEN, _ID_PEN_CLOSE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	RemoveObject(p.native)
	p.native = 0
	return
}

func (p *Pen) SetColor(clr color.Color) {
	sh_clr := make_rgba(clr)
	_drv_ch(CLASSID_PEN, _ID_PEN_SETCOLOR, unsafe.Pointer(p.info()), unsafe.Pointer(&sh_clr), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Pen) Color() (clr color.Color) {
	var sh_clr rgba
	_drv_ch(CLASSID_PEN, _ID_PEN_COLOR, unsafe.Pointer(p.info()), unsafe.Pointer(&sh_clr), nil, nil, nil, nil, nil, nil, nil, nil)
	clr = sh_clr
	return
}

func (p *Pen) SetWidth(width int) {
	_drv_ch(CLASSID_PEN, _ID_PEN_SETWIDTH, unsafe.Pointer(p.info()), unsafe.Pointer(&width), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Pen) Width() (width int) {
	_drv_ch(CLASSID_PEN, _ID_PEN_WIDTH, unsafe.Pointer(p.info()), unsafe.Pointer(&width), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Pen) SetStyle(style PenStyle) {
	_drv_ch(CLASSID_PEN, _ID_PEN_SETSTYLE, unsafe.Pointer(p.info()), unsafe.Pointer(&style), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Pen) Style() (style PenStyle) {
	_drv_ch(CLASSID_PEN, _ID_PEN_STYLE, unsafe.Pointer(p.info()), unsafe.Pointer(&style), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

// struct Painter
//
type Painter struct {
	object
}

func (p *Painter) Name() string {
	return "Painter"
}

func (p *Painter) String() string {
	return DumpObject(p)
}
func (p *Painter) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "font":
		if v, ok := value.(*Font); ok {
			p.SetFont(v)
			return true
		}
		return false
	case "pen":
		if v, ok := value.(*Pen); ok {
			p.SetPen(v)
			return true
		}
		return false
	case "brush":
		if v, ok := value.(*Brush); ok {
			p.SetBrush(v)
			return true
		}
		return false
	default:
		return p.object.SetAttr(attr, value)
	}
	return false
}
func (p *Painter) Attr(attr string) interface{} {
	switch attr {
	case "font":
		return p.Font()
	case "pen":
		return p.Pen()
	case "brush":
		return p.Brush()
	default:
		return p.object.Attr(attr)
	}
	return nil
}
func NewPainter() *Painter {
	return new(Painter).Init()
}

func (p *Painter) Init() *Painter {
	p.classid = CLASSID_PAINTER
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	runtime.SetFinalizer(p, (*Painter).Close)
	return p
}

func NewPainterWithImage(image *Image) *Painter {
	return new(Painter).InitWithImage(image)
}

func (p *Painter) InitWithImage(image *Image) *Painter {
	p.classid = CLASSID_PAINTER
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_INITWITHIMAGE, unsafe.Pointer(p.info()), unsafe.Pointer(image), nil, nil, nil, nil, nil, nil, nil, nil)
	runtime.SetFinalizer(p, (*Painter).Close)
	return p
}

func (p *Painter) Close() (err error) {
	if p == nil || p.native == 0 {
		return
	}
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_CLOSE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	p.native = 0
	runtime.SetFinalizer(p, nil)
	return
}

func (p *Painter) InitFrom(w IWidget) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_INITFROM, unsafe.Pointer(p.info()), unsafe.Pointer(w.(iobj).info()), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) Begin(w IWidget) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_BEGIN, unsafe.Pointer(p.info()), unsafe.Pointer(w.(iobj).info()), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) End() {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_END, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) SetFont(font *Font) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_SETFONT, unsafe.Pointer(p.info()), unsafe.Pointer(font), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) Font() (font *Font) {
	var oi_font obj_info
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_FONT, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_font), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_font.native != 0 {
		v := FindObject(oi_font.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_FONT, oi_font.native)
		}
		if v != nil {
			font = v.(*Font)
		}
	}
	return
}

func (p *Painter) SetPen(pen *Pen) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_SETPEN, unsafe.Pointer(p.info()), unsafe.Pointer(pen), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) Pen() (pen *Pen) {
	var oi_pen obj_info
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_PEN, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_pen), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_pen.native != 0 {
		v := FindObject(oi_pen.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_PEN, oi_pen.native)
		}
		if v != nil {
			pen = v.(*Pen)
		}
	}
	return
}

func (p *Painter) SetBrush(brush *Brush) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_SETBRUSH, unsafe.Pointer(p.info()), unsafe.Pointer(brush), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) Brush() (brush *Brush) {
	var oi_brush obj_info
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_BRUSH, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_brush), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_brush.native != 0 {
		v := FindObject(oi_brush.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_BRUSH, oi_brush.native)
		}
		if v != nil {
			brush = v.(*Brush)
		}
	}
	return
}

func (p *Painter) DrawPoint(pt Point) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWPOINT, unsafe.Pointer(p.info()), unsafe.Pointer(&pt), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawPoints(pts []Point) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWPOINTS, unsafe.Pointer(p.info()), unsafe.Pointer((*slice_info)(unsafe.Pointer(&pts))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawLine(pt1 Point, pt2 Point) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWLINE, unsafe.Pointer(p.info()), unsafe.Pointer(&pt1), unsafe.Pointer(&pt2), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawLines(pts []Point) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWLINES, unsafe.Pointer(p.info()), unsafe.Pointer((*slice_info)(unsafe.Pointer(&pts))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawPolyline(pts []Point) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWPOLYLINE, unsafe.Pointer(p.info()), unsafe.Pointer((*slice_info)(unsafe.Pointer(&pts))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawPolygon(pts []Point) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWPOLYGON, unsafe.Pointer(p.info()), unsafe.Pointer((*slice_info)(unsafe.Pointer(&pts))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawRect(rc Rect) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWRECT, unsafe.Pointer(p.info()), unsafe.Pointer(&rc), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawRects(rcs []Rect) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWRECTS, unsafe.Pointer(p.info()), unsafe.Pointer((*slice_info)(unsafe.Pointer(&rcs))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawRoundedRect(rc Rect, xRadius float64, yRadius float64, sizeMode int) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWROUNDEDRECT, unsafe.Pointer(p.info()), unsafe.Pointer(&rc), unsafe.Pointer(&xRadius), unsafe.Pointer(&yRadius), unsafe.Pointer(&sizeMode), nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawEllipse(rc Rect) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWELLIPSE, unsafe.Pointer(p.info()), unsafe.Pointer(&rc), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawArc(rc Rect, startAngle int, spanAngle int) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWARC, unsafe.Pointer(p.info()), unsafe.Pointer(&rc), unsafe.Pointer(&startAngle), unsafe.Pointer(&spanAngle), nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawChord(rc Rect, startAngle int, spanAngle int) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWCHORD, unsafe.Pointer(p.info()), unsafe.Pointer(&rc), unsafe.Pointer(&startAngle), unsafe.Pointer(&spanAngle), nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawPie(rc Rect, startAngle int, spanAngle int) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWPIE, unsafe.Pointer(p.info()), unsafe.Pointer(&rc), unsafe.Pointer(&startAngle), unsafe.Pointer(&spanAngle), nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawText(pt Point, text string) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWTEXT, unsafe.Pointer(p.info()), unsafe.Pointer(&pt), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawTextRect(rc Rect, flags int, text string) (bound Rect) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWTEXTRECT, unsafe.Pointer(p.info()), unsafe.Pointer(&rc), unsafe.Pointer(&flags), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), unsafe.Pointer(&bound), nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawPixmap(pt Point, pixmap *Pixmap) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWPIXMAP, unsafe.Pointer(p.info()), unsafe.Pointer(&pt), unsafe.Pointer(pixmap), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawPixmapEx(pt Point, pixmap *Pixmap, source Rect) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWPIXMAPEX, unsafe.Pointer(p.info()), unsafe.Pointer(&pt), unsafe.Pointer(pixmap), unsafe.Pointer(&source), nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawPixmapRect(rc Rect, pixmap *Pixmap) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWPIXMAPRECT, unsafe.Pointer(p.info()), unsafe.Pointer(&rc), unsafe.Pointer(pixmap), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawPixmapRectEx(rc Rect, pixmap *Pixmap, source Rect) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWPIXMAPRECTEX, unsafe.Pointer(p.info()), unsafe.Pointer(&rc), unsafe.Pointer(pixmap), unsafe.Pointer(&source), nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawImage(pt Point, image *Image) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWIMAGE, unsafe.Pointer(p.info()), unsafe.Pointer(&pt), unsafe.Pointer(image), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawImageEx(pt Point, image *Image, source Rect) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWIMAGEEX, unsafe.Pointer(p.info()), unsafe.Pointer(&pt), unsafe.Pointer(image), unsafe.Pointer(&source), nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawImageRect(rc Rect, image *Image) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWIMAGERECT, unsafe.Pointer(p.info()), unsafe.Pointer(&rc), unsafe.Pointer(image), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) DrawImageRectEx(rc Rect, image *Image, source Rect) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_DRAWIMAGERECTEX, unsafe.Pointer(p.info()), unsafe.Pointer(&rc), unsafe.Pointer(image), unsafe.Pointer(&source), nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) FillRect(rc Rect, color uint) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_FILLRECT, unsafe.Pointer(p.info()), unsafe.Pointer(&rc), unsafe.Pointer(&color), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *Painter) FillRectF(rc RectF, color uint) {
	_drv_ch(CLASSID_PAINTER, _ID_PAINTER_FILLRECTF, unsafe.Pointer(p.info()), unsafe.Pointer(&rc), unsafe.Pointer(&color), nil, nil, nil, nil, nil, nil, nil)
	return
}

// struct ListWidgetItem
//
type ListWidgetItem struct {
	object
}

func (p *ListWidgetItem) Name() string {
	return "ListWidgetItem"
}

func (p *ListWidgetItem) String() string {
	return DumpObject(p)
}
func (p *ListWidgetItem) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "text":
		if v, ok := value.(string); ok {
			p.SetText(v)
			return true
		}
		return false
	case "icon":
		if v, ok := value.(*Icon); ok {
			p.SetIcon(v)
			return true
		}
		return false
	case "selected":
		if v, ok := value.(bool); ok {
			p.SetSelected(v)
			return true
		}
		return false
	case "hidden":
		if v, ok := value.(bool); ok {
			p.SetHidden(v)
			return true
		}
		return false
	case "font":
		if v, ok := value.(*Font); ok {
			p.SetFont(v)
			return true
		}
		return false
	case "foreground":
		if v, ok := value.(*Brush); ok {
			p.SetForeground(v)
			return true
		}
		return false
	case "background":
		if v, ok := value.(*Brush); ok {
			p.SetBackground(v)
			return true
		}
		return false
	case "tooltip":
		if v, ok := value.(string); ok {
			p.SetToolTip(v)
			return true
		}
		return false
	case "textalignment":
		if v, ok := value.(Alignment); ok {
			p.SetTextAlignment(v)
			return true
		}
		return false
	case "flags":
		if v, ok := value.(ItemFlag); ok {
			p.SetFlags(v)
			return true
		}
		return false
	default:
		return p.object.SetAttr(attr, value)
	}
	return false
}
func (p *ListWidgetItem) Attr(attr string) interface{} {
	switch attr {
	case "text":
		return p.Text()
	case "icon":
		return p.Icon()
	case "selected":
		return p.IsSelected()
	case "hidden":
		return p.IsHidden()
	case "foreground":
		return p.Foreground()
	case "background":
		return p.Background()
	case "tooltip":
		return p.ToolTip()
	case "textalignment":
		return p.TextAlignment()
	case "flags":
		return p.Flags()
	default:
		return p.object.Attr(attr)
	}
	return nil
}
func NewListWidgetItem() *ListWidgetItem {
	return new(ListWidgetItem).Init()
}

func (p *ListWidgetItem) Init() *ListWidgetItem {
	p.classid = CLASSID_LISTWIDGETITEM
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func NewListWidgetItemWithText(text string) *ListWidgetItem {
	return new(ListWidgetItem).InitWithText(text)
}

func (p *ListWidgetItem) InitWithText(text string) *ListWidgetItem {
	p.classid = CLASSID_LISTWIDGETITEM
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_INITWITHTEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *ListWidgetItem) Close() (err error) {
	if p == nil || p.native == 0 {
		return
	}
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_CLOSE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	RemoveObject(p.native)
	p.native = 0
	return
}

func (p *ListWidgetItem) SetText(text string) {
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_SETTEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&text))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidgetItem) Text() (text string) {
	var sh_text utf8_info
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_TEXT, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_text))), nil, nil, nil, nil, nil, nil, nil, nil)
	text = sh_text.String()
	return
}

func (p *ListWidgetItem) SetIcon(icon *Icon) {
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_SETICON, unsafe.Pointer(p.info()), unsafe.Pointer(icon), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidgetItem) Icon() (icon *Icon) {
	var oi_icon obj_info
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_ICON, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_icon), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_icon.native != 0 {
		v := FindObject(oi_icon.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_ICON, oi_icon.native)
		}
		if v != nil {
			icon = v.(*Icon)
		}
	}
	return
}

func (p *ListWidgetItem) SetSelected(b bool) {
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_SETSELECTED, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidgetItem) IsSelected() (b bool) {
	var b_b int
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_ISSELECTED, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *ListWidgetItem) SetHidden(b bool) {
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_SETHIDDEN, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidgetItem) IsHidden() (b bool) {
	var b_b int
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_ISHIDDEN, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *ListWidgetItem) SetFont(font *Font) {
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_SETFONT, unsafe.Pointer(p.info()), unsafe.Pointer(font), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidgetItem) Font() {
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_FONT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidgetItem) SetForeground(brush *Brush) {
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_SETFOREGROUND, unsafe.Pointer(p.info()), unsafe.Pointer(brush), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidgetItem) Foreground() (brush *Brush) {
	var oi_brush obj_info
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_FOREGROUND, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_brush), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_brush.native != 0 {
		v := FindObject(oi_brush.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_BRUSH, oi_brush.native)
		}
		if v != nil {
			brush = v.(*Brush)
		}
	}
	return
}

func (p *ListWidgetItem) SetBackground(brush *Brush) {
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_SETBACKGROUND, unsafe.Pointer(p.info()), unsafe.Pointer(brush), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidgetItem) Background() (brush *Brush) {
	var oi_brush obj_info
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_BACKGROUND, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_brush), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_brush.native != 0 {
		v := FindObject(oi_brush.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_BRUSH, oi_brush.native)
		}
		if v != nil {
			brush = v.(*Brush)
		}
	}
	return
}

func (p *ListWidgetItem) SetToolTip(tip string) {
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_SETTOOLTIP, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&tip))), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidgetItem) ToolTip() (tip string) {
	var sh_tip utf8_info
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_TOOLTIP, unsafe.Pointer(p.info()), unsafe.Pointer((*string_info)(unsafe.Pointer(&sh_tip))), nil, nil, nil, nil, nil, nil, nil, nil)
	tip = sh_tip.String()
	return
}

func (p *ListWidgetItem) SetTextAlignment(value Alignment) {
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_SETTEXTALIGNMENT, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidgetItem) TextAlignment() (value Alignment) {
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_TEXTALIGNMENT, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidgetItem) SetFlags(value ItemFlag) {
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_SETFLAGS, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidgetItem) Flags() (value ItemFlag) {
	_drv_ch(CLASSID_LISTWIDGETITEM, _ID_LISTWIDGETITEM_FLAGS, unsafe.Pointer(p.info()), unsafe.Pointer(&value), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

// struct ListWidget
//
type ListWidget struct {
	Widget
}

func (p *ListWidget) Name() string {
	return "ListWidget"
}

func (p *ListWidget) String() string {
	return DumpObject(p)
}
func (p *ListWidget) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "currentitem":
		if v, ok := value.(*ListWidgetItem); ok {
			p.SetCurrentItem(v)
			return true
		}
		return false
	case "currentrow":
		if v, ok := value.(int); ok {
			p.SetCurrentRow(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr, value)
	}
	return false
}
func (p *ListWidget) Attr(attr string) interface{} {
	switch attr {
	case "count":
		return p.Count()
	case "currentitem":
		return p.CurrentItem()
	case "currentrow":
		return p.CurrentRow()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewListWidget() *ListWidget {
	return new(ListWidget).Init()
}

func (p *ListWidget) Init() *ListWidget {
	p.classid = CLASSID_LISTWIDGET
	_drv_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *ListWidget) Count() (count int) {
	_drv_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_COUNT, unsafe.Pointer(p.info()), unsafe.Pointer(&count), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidget) SetCurrentItem(item *ListWidgetItem) {
	_drv_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_SETCURRENTITEM, unsafe.Pointer(p.info()), unsafe.Pointer(item), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidget) CurrentItem() (item *ListWidgetItem) {
	var oi_item obj_info
	_drv_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_CURRENTITEM, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_item), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_item.native != 0 {
		v := FindObject(oi_item.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_LISTWIDGETITEM, oi_item.native)
		}
		if v != nil {
			item = v.(*ListWidgetItem)
		}
	}
	return
}

func (p *ListWidget) SetCurrentRow(row int) {
	_drv_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_SETCURRENTROW, unsafe.Pointer(p.info()), unsafe.Pointer(&row), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidget) CurrentRow() (row int) {
	_drv_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_CURRENTROW, unsafe.Pointer(p.info()), unsafe.Pointer(&row), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidget) AddItem(item *ListWidgetItem) {
	_drv_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_ADDITEM, unsafe.Pointer(p.info()), unsafe.Pointer(item), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidget) InsertItem(index int, item *ListWidgetItem) {
	_drv_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_INSERTITEM, unsafe.Pointer(p.info()), unsafe.Pointer(&index), unsafe.Pointer(item), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidget) EditItem(item *ListWidgetItem) {
	_drv_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_EDITITEM, unsafe.Pointer(p.info()), unsafe.Pointer(item), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidget) TakeItem(row int) (item *ListWidgetItem) {
	var oi_item obj_info
	_drv_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_TAKEITEM, unsafe.Pointer(p.info()), unsafe.Pointer(&row), unsafe.Pointer(&oi_item), nil, nil, nil, nil, nil, nil, nil)
	if oi_item.native != 0 {
		v := FindObject(oi_item.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_LISTWIDGETITEM, oi_item.native)
		}
		if v != nil {
			item = v.(*ListWidgetItem)
		}
	}
	return
}

func (p *ListWidget) Item(row int) (item *ListWidgetItem) {
	var oi_item obj_info
	_drv_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_ITEM, unsafe.Pointer(p.info()), unsafe.Pointer(&row), unsafe.Pointer(&oi_item), nil, nil, nil, nil, nil, nil, nil)
	if oi_item.native != 0 {
		v := FindObject(oi_item.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_LISTWIDGETITEM, oi_item.native)
		}
		if v != nil {
			item = v.(*ListWidgetItem)
		}
	}
	return
}

func (p *ListWidget) Clear() {
	_drv_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_CLEAR, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ListWidget) OnCurrentItemChanged(fn func(*ListWidgetItem, *ListWidgetItem)) {
	_drv_event_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_ONCURRENTITEMCHANGED, p, fn)
	return
}

func (p *ListWidget) OnCurrentRowChanged(fn func(int)) {
	_drv_event_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_ONCURRENTROWCHANGED, p, fn)
	return
}

func (p *ListWidget) OnItemActivated(fn func(*ListWidgetItem)) {
	_drv_event_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_ONITEMACTIVATED, p, fn)
	return
}

func (p *ListWidget) OnItemChanged(fn func(*ListWidgetItem)) {
	_drv_event_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_ONITEMCHANGED, p, fn)
	return
}

func (p *ListWidget) OnItemClicked(fn func(*ListWidgetItem)) {
	_drv_event_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_ONITEMCLICKED, p, fn)
	return
}

func (p *ListWidget) OnItemDoubleClicked(fn func(*ListWidgetItem)) {
	_drv_event_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_ONITEMDOUBLECLICKED, p, fn)
	return
}

func (p *ListWidget) OnItemEntered(fn func(*ListWidgetItem)) {
	_drv_event_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_ONITEMENTERED, p, fn)
	return
}

func (p *ListWidget) OnItemPressed(fn func(*ListWidgetItem)) {
	_drv_event_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_ONITEMPRESSED, p, fn)
	return
}

func (p *ListWidget) OnItemSelectionChanged(fn func()) {
	_drv_event_ch(CLASSID_LISTWIDGET, _ID_LISTWIDGET_ONITEMSELECTIONCHANGED, p, fn)
	return
}

// struct MainWindow
//
type MainWindow struct {
	Widget
}

func (p *MainWindow) Name() string {
	return "MainWindow"
}

func (p *MainWindow) String() string {
	return DumpObject(p)
}
func (p *MainWindow) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "centralwidget":
		if v, ok := value.(IWidget); ok {
			p.SetCentralWidget(v)
			return true
		}
		return false
	case "menubar":
		if v, ok := value.(*MenuBar); ok {
			p.SetMenuBar(v)
			return true
		}
		return false
	case "statusbar":
		if v, ok := value.(*StatusBar); ok {
			p.SetStatusBar(v)
			return true
		}
		return false
	default:
		return p.Widget.SetAttr(attr, value)
	}
	return false
}
func (p *MainWindow) Attr(attr string) interface{} {
	switch attr {
	case "centralwidget":
		return p.CentralWidget()
	case "menubar":
		return p.MenuBar()
	case "statusbar":
		return p.StatusBar()
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewMainWindow() *MainWindow {
	return new(MainWindow).Init()
}

func (p *MainWindow) Init() *MainWindow {
	p.classid = CLASSID_MAINWINDOW
	_drv_ch(CLASSID_MAINWINDOW, _ID_MAINWINDOW_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *MainWindow) SetCentralWidget(widget IWidget) {
	_drv_ch(CLASSID_MAINWINDOW, _ID_MAINWINDOW_SETCENTRALWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(widget.(iobj).info()), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *MainWindow) CentralWidget() (widget IWidget) {
	var oi_widget obj_info
	_drv_ch(CLASSID_MAINWINDOW, _ID_MAINWINDOW_CENTRALWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_widget), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_widget.native != 0 {
		item := FindObject(oi_widget.native)
		if item != nil {
			widget = item.(IWidget)
		}
	}
	return
}

func (p *MainWindow) SetMenuBar(bar *MenuBar) {
	_drv_ch(CLASSID_MAINWINDOW, _ID_MAINWINDOW_SETMENUBAR, unsafe.Pointer(p.info()), unsafe.Pointer(bar), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *MainWindow) MenuBar() (bar *MenuBar) {
	var oi_bar obj_info
	_drv_ch(CLASSID_MAINWINDOW, _ID_MAINWINDOW_MENUBAR, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_bar), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_bar.native != 0 {
		v := FindObject(oi_bar.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_MENUBAR, oi_bar.native)
		}
		if v != nil {
			bar = v.(*MenuBar)
		}
	}
	return
}

func (p *MainWindow) SetStatusBar(bar *StatusBar) {
	_drv_ch(CLASSID_MAINWINDOW, _ID_MAINWINDOW_SETSTATUSBAR, unsafe.Pointer(p.info()), unsafe.Pointer(bar), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *MainWindow) StatusBar() (bar *StatusBar) {
	var oi_bar obj_info
	_drv_ch(CLASSID_MAINWINDOW, _ID_MAINWINDOW_STATUSBAR, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_bar), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_bar.native != 0 {
		v := FindObject(oi_bar.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_STATUSBAR, oi_bar.native)
		}
		if v != nil {
			bar = v.(*StatusBar)
		}
	}
	return
}

func (p *MainWindow) AddToolBar(bar *ToolBar) {
	_drv_ch(CLASSID_MAINWINDOW, _ID_MAINWINDOW_ADDTOOLBAR, unsafe.Pointer(p.info()), unsafe.Pointer(bar), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *MainWindow) InsertToolBar(before *ToolBar, bar *ToolBar) {
	_drv_ch(CLASSID_MAINWINDOW, _ID_MAINWINDOW_INSERTTOOLBAR, unsafe.Pointer(p.info()), unsafe.Pointer(before), unsafe.Pointer(bar), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *MainWindow) RemoveToolBar(bar *ToolBar) {
	_drv_ch(CLASSID_MAINWINDOW, _ID_MAINWINDOW_REMOVETOOLBAR, unsafe.Pointer(p.info()), unsafe.Pointer(bar), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *MainWindow) AddDockWidget(area DockWidgetArea, dock *DockWidget) {
	_drv_ch(CLASSID_MAINWINDOW, _ID_MAINWINDOW_ADDDOCKWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(&area), unsafe.Pointer(dock), nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *MainWindow) RemoveDockWidget(dock *DockWidget) {
	_drv_ch(CLASSID_MAINWINDOW, _ID_MAINWINDOW_REMOVEDOCKWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(dock), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

// struct GLWidget
//
type GLWidget struct {
	Widget
}

func (p *GLWidget) Name() string {
	return "GLWidget"
}

func (p *GLWidget) String() string {
	return DumpObject(p)
}
func (p *GLWidget) SetAttr(attr string, value interface{}) bool {
	switch attr {
	default:
		return p.Widget.SetAttr(attr, value)
	}
	return false
}
func (p *GLWidget) Attr(attr string) interface{} {
	switch attr {
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func NewGLWidget() *GLWidget {
	return new(GLWidget).Init()
}

func (p *GLWidget) Init() *GLWidget {
	p.classid = CLASSID_GLWIDGET
	_drv_ch(CLASSID_GLWIDGET, _ID_GLWIDGET_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *GLWidget) DeleteTexture(id uint) {
	_drv_ch(CLASSID_GLWIDGET, _ID_GLWIDGET_DELETETEXTURE, unsafe.Pointer(p.info()), unsafe.Pointer(&id), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *GLWidget) DoneCurrent() {
	_drv_ch(CLASSID_GLWIDGET, _ID_GLWIDGET_DONECURRENT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *GLWidget) DoubleBuffer() (b bool) {
	var b_b int
	_drv_ch(CLASSID_GLWIDGET, _ID_GLWIDGET_DOUBLEBUFFER, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *GLWidget) ConvertToGLFormat(image *Image) (glImage *Image) {
	var oi_glImage obj_info
	_drv_ch(CLASSID_GLWIDGET, _ID_GLWIDGET_CONVERTTOGLFORMAT, unsafe.Pointer(p.info()), unsafe.Pointer(image), unsafe.Pointer(&oi_glImage), nil, nil, nil, nil, nil, nil, nil)
	if oi_glImage.native != 0 {
		v := FindObject(oi_glImage.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_IMAGE, oi_glImage.native)
		}
		if v != nil {
			glImage = v.(*Image)
		}
	}
	return
}

func (p *GLWidget) SetMouseTracking(enable bool) {
	_drv_ch(CLASSID_GLWIDGET, _ID_GLWIDGET_SETMOUSETRACKING, unsafe.Pointer(p.info()), unsafe.Pointer(&enable), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *GLWidget) RenderText(x int, y int, z int, str string, font *Font) {
	_drv_ch(CLASSID_GLWIDGET, _ID_GLWIDGET_RENDERTEXT, unsafe.Pointer(p.info()), unsafe.Pointer(&x), unsafe.Pointer(&y), unsafe.Pointer(&z), unsafe.Pointer((*string_info)(unsafe.Pointer(&str))), unsafe.Pointer(font), nil, nil, nil, nil)
	return
}

func (p *GLWidget) UpdateGL() {
	_drv_ch(CLASSID_GLWIDGET, _ID_GLWIDGET_UPDATEGL, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *GLWidget) UpdateOverlayGL() {
	_drv_ch(CLASSID_GLWIDGET, _ID_GLWIDGET_UPDATEOVERLAYGL, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *GLWidget) OnInitializeGL(fn func()) {
	_drv_event_ch(CLASSID_GLWIDGET, _ID_GLWIDGET_ONINITIALIZEGL, p, fn)
	return
}

func (p *GLWidget) OnInitializeOverlayGL(fn func()) {
	_drv_event_ch(CLASSID_GLWIDGET, _ID_GLWIDGET_ONINITIALIZEOVERLAYGL, p, fn)
	return
}

func (p *GLWidget) OnPaintGL(fn func()) {
	_drv_event_ch(CLASSID_GLWIDGET, _ID_GLWIDGET_ONPAINTGL, p, fn)
	return
}

func (p *GLWidget) OnPaintOverlayGL(fn func()) {
	_drv_event_ch(CLASSID_GLWIDGET, _ID_GLWIDGET_ONPAINTOVERLAYGL, p, fn)
	return
}

func (p *GLWidget) OnResizeGL(fn func(int, int)) {
	_drv_event_ch(CLASSID_GLWIDGET, _ID_GLWIDGET_ONRESIZEGL, p, fn)
	return
}

func (p *GLWidget) OnResizeOverlayGL(fn func(int, int)) {
	_drv_event_ch(CLASSID_GLWIDGET, _ID_GLWIDGET_ONRESIZEOVERLAYGL, p, fn)
	return
}

// struct SizePolicy
//
type SizePolicy struct {
	object
}

func (p *SizePolicy) Name() string {
	return "SizePolicy"
}

func (p *SizePolicy) String() string {
	return DumpObject(p)
}
func (p *SizePolicy) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "horizontalpolicy":
		if v, ok := value.(SizePolicyPolicy); ok {
			p.SetHorizontalPolicy(v)
			return true
		}
		return false
	case "verticalpolicy":
		if v, ok := value.(SizePolicyPolicy); ok {
			p.SetVerticalPolicy(v)
			return true
		}
		return false
	default:
		return p.object.SetAttr(attr, value)
	}
	return false
}
func (p *SizePolicy) Attr(attr string) interface{} {
	switch attr {
	case "horizontalpolicy":
		return p.HorizontalPolicy()
	case "verticalpolicy":
		return p.VerticalPolicy()
	case "heightforwidth":
		return p.HasHeightForWidth()
	default:
		return p.object.Attr(attr)
	}
	return nil
}
func NewSizePolicy() *SizePolicy {
	return new(SizePolicy).Init()
}

func (p *SizePolicy) Init() *SizePolicy {
	p.classid = CLASSID_SIZEPOLICY
	_drv_ch(CLASSID_SIZEPOLICY, _ID_SIZEPOLICY_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	runtime.SetFinalizer(p, (*SizePolicy).Close)
	return p
}

func NewSizePolicyWithPolicy(horizontal SizePolicyPolicy, vertical SizePolicyPolicy, control SizePolicyControlType) *SizePolicy {
	return new(SizePolicy).InitWithPolicy(horizontal, vertical, control)
}

func (p *SizePolicy) InitWithPolicy(horizontal SizePolicyPolicy, vertical SizePolicyPolicy, control SizePolicyControlType) *SizePolicy {
	p.classid = CLASSID_SIZEPOLICY
	_drv_ch(CLASSID_SIZEPOLICY, _ID_SIZEPOLICY_INITWITHPOLICY, unsafe.Pointer(p.info()), unsafe.Pointer(&horizontal), unsafe.Pointer(&vertical), unsafe.Pointer(&control), nil, nil, nil, nil, nil, nil)
	runtime.SetFinalizer(p, (*SizePolicy).Close)
	return p
}

func (p *SizePolicy) Close() (err error) {
	if p == nil || p.native == 0 {
		return
	}
	_drv_ch(CLASSID_SIZEPOLICY, _ID_SIZEPOLICY_CLOSE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	p.native = 0
	runtime.SetFinalizer(p, nil)
	return
}

func (p *SizePolicy) HorizontalPolicy() (policy SizePolicyPolicy) {
	_drv_ch(CLASSID_SIZEPOLICY, _ID_SIZEPOLICY_HORIZONTALPOLICY, unsafe.Pointer(p.info()), unsafe.Pointer(&policy), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *SizePolicy) SetHorizontalPolicy(policy SizePolicyPolicy) {
	_drv_ch(CLASSID_SIZEPOLICY, _ID_SIZEPOLICY_SETHORIZONTALPOLICY, unsafe.Pointer(p.info()), unsafe.Pointer(&policy), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *SizePolicy) VerticalPolicy() (policy SizePolicyPolicy) {
	_drv_ch(CLASSID_SIZEPOLICY, _ID_SIZEPOLICY_VERTICALPOLICY, unsafe.Pointer(p.info()), unsafe.Pointer(&policy), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *SizePolicy) SetVerticalPolicy(policy SizePolicyPolicy) {
	_drv_ch(CLASSID_SIZEPOLICY, _ID_SIZEPOLICY_SETVERTICALPOLICY, unsafe.Pointer(p.info()), unsafe.Pointer(&policy), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *SizePolicy) HasHeightForWidth() (b bool) {
	var b_b int
	_drv_ch(CLASSID_SIZEPOLICY, _ID_SIZEPOLICY_HASHEIGHTFORWIDTH, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *SizePolicy) Transpose() {
	_drv_ch(CLASSID_SIZEPOLICY, _ID_SIZEPOLICY_TRANSPOSE, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

// struct baseScrollArea
//
type baseScrollArea struct {
	Widget
}

func (p *baseScrollArea) Name() string {
	return "baseScrollArea"
}

func (p *baseScrollArea) String() string {
	return DumpObject(p)
}
func (p *baseScrollArea) SetAttr(attr string, value interface{}) bool {
	switch attr {
	default:
		return p.Widget.SetAttr(attr, value)
	}
	return false
}
func (p *baseScrollArea) Attr(attr string) interface{} {
	switch attr {
	default:
		return p.Widget.Attr(attr)
	}
	return nil
}
func (p *baseScrollArea) CornerWidget() (w IWidget) {
	var oi_w obj_info
	_drv_event_ch(_CLASSID_BASESCROLLAREA, _ID_BASESCROLLAREA_CORNERWIDGET, p, &oi_w)
	if oi_w.native != 0 {
		item := FindObject(oi_w.native)
		if item != nil {
			w = item.(IWidget)
		}
	}
	return
}

func (p *baseScrollArea) HorizontalScrollBar() (s *ScrollBar) {
	var oi_s obj_info
	_drv_event_ch(_CLASSID_BASESCROLLAREA, _ID_BASESCROLLAREA_HORIZONTALSCROLLBAR, p, &oi_s)
	if oi_s.native != 0 {
		v := FindObject(oi_s.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_SCROLLBAR, oi_s.native)
		}
		if v != nil {
			s = v.(*ScrollBar)
		}
	}
	return
}

func (p *baseScrollArea) VerticalScrollBar() (s *ScrollBar) {
	var oi_s obj_info
	_drv_event_ch(_CLASSID_BASESCROLLAREA, _ID_BASESCROLLAREA_VERTICALSCROLLBAR, p, &oi_s)
	if oi_s.native != 0 {
		v := FindObject(oi_s.native)
		if v == nil {
			v = NewObjectWithNative(CLASSID_SCROLLBAR, oi_s.native)
		}
		if v != nil {
			s = v.(*ScrollBar)
		}
	}
	return
}

func (p *baseScrollArea) Viewport() (w IWidget) {
	var oi_w obj_info
	_drv_event_ch(_CLASSID_BASESCROLLAREA, _ID_BASESCROLLAREA_VIEWPORT, p, &oi_w)
	if oi_w.native != 0 {
		item := FindObject(oi_w.native)
		if item != nil {
			w = item.(IWidget)
		}
	}
	return
}

// struct ScrollArea
//
type ScrollArea struct {
	baseScrollArea
}

func (p *ScrollArea) Name() string {
	return "ScrollArea"
}

func (p *ScrollArea) String() string {
	return DumpObject(p)
}
func (p *ScrollArea) SetAttr(attr string, value interface{}) bool {
	switch attr {
	case "alignment":
		if v, ok := value.(Alignment); ok {
			p.SetAlignment(v)
			return true
		}
		return false
	case "widget":
		if v, ok := value.(IWidget); ok {
			p.SetWidget(v)
			return true
		}
		return false
	case "widgetresizable":
		if v, ok := value.(bool); ok {
			p.SetWidgetResizable(v)
			return true
		}
		return false
	default:
		return p.baseScrollArea.SetAttr(attr, value)
	}
	return false
}
func (p *ScrollArea) Attr(attr string) interface{} {
	switch attr {
	case "alignment":
		return p.Alignment()
	case "widget":
		return p.Widget()
	case "widgetresizable":
		return p.WidgetResizable()
	case "takewidget":
		return p.TakeWidget()
	default:
		return p.baseScrollArea.Attr(attr)
	}
	return nil
}
func NewScrollArea() *ScrollArea {
	return new(ScrollArea).Init()
}

func (p *ScrollArea) Init() *ScrollArea {
	p.classid = CLASSID_SCROLLAREA
	_drv_ch(CLASSID_SCROLLAREA, _ID_SCROLLAREA_INIT, unsafe.Pointer(p.info()), nil, nil, nil, nil, nil, nil, nil, nil, nil)
	InsertObject(p)
	return p
}

func (p *ScrollArea) SetAlignment(a Alignment) {
	_drv_ch(CLASSID_SCROLLAREA, _ID_SCROLLAREA_SETALIGNMENT, unsafe.Pointer(p.info()), unsafe.Pointer(&a), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ScrollArea) Alignment() (a Alignment) {
	_drv_ch(CLASSID_SCROLLAREA, _ID_SCROLLAREA_ALIGNMENT, unsafe.Pointer(p.info()), unsafe.Pointer(&a), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ScrollArea) SetWidget(w IWidget) {
	_drv_ch(CLASSID_SCROLLAREA, _ID_SCROLLAREA_SETWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(w.(iobj).info()), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ScrollArea) Widget() (w IWidget) {
	var oi_w obj_info
	_drv_ch(CLASSID_SCROLLAREA, _ID_SCROLLAREA_WIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_w), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_w.native != 0 {
		item := FindObject(oi_w.native)
		if item != nil {
			w = item.(IWidget)
		}
	}
	return
}

func (p *ScrollArea) SetWidgetResizable(b bool) {
	_drv_ch(CLASSID_SCROLLAREA, _ID_SCROLLAREA_SETWIDGETRESIZABLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b), nil, nil, nil, nil, nil, nil, nil, nil)
	return
}

func (p *ScrollArea) WidgetResizable() (b bool) {
	var b_b int
	_drv_ch(CLASSID_SCROLLAREA, _ID_SCROLLAREA_WIDGETRESIZABLE, unsafe.Pointer(p.info()), unsafe.Pointer(&b_b), nil, nil, nil, nil, nil, nil, nil, nil)
	b = b_b != 0
	return
}

func (p *ScrollArea) TakeWidget() (w IWidget) {
	var oi_w obj_info
	_drv_ch(CLASSID_SCROLLAREA, _ID_SCROLLAREA_TAKEWIDGET, unsafe.Pointer(p.info()), unsafe.Pointer(&oi_w), nil, nil, nil, nil, nil, nil, nil, nil)
	if oi_w.native != 0 {
		item := FindObject(oi_w.native)
		if item != nil {
			w = item.(IWidget)
		}
	}
	return
}

func (p *ScrollArea) EnsureVisible(x int, y int, xmargin int, ymargin int) {
	_drv_ch(CLASSID_SCROLLAREA, _ID_SCROLLAREA_ENSUREVISIBLE, unsafe.Pointer(p.info()), unsafe.Pointer(&x), unsafe.Pointer(&y), unsafe.Pointer(&xmargin), unsafe.Pointer(&ymargin), nil, nil, nil, nil, nil)
	return
}

func (p *ScrollArea) EnsureWidgetVisible(w IWidget, xmargin int, ymargin int) {
	_drv_ch(CLASSID_SCROLLAREA, _ID_SCROLLAREA_ENSUREWIDGETVISIBLE, unsafe.Pointer(p.info()), unsafe.Pointer(w.(iobj).info()), unsafe.Pointer(&xmargin), unsafe.Pointer(&ymargin), nil, nil, nil, nil, nil, nil)
	return
}
