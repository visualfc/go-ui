/********************************************************************
** Copyright 2012 visualfc <visualfc@gmail.com>. All rights reserved.
**
** This library is free software; you can redistribute it and/or
** modify it under the terms of the GNU Lesser General Public
** License as published by the Free Software Foundation; either
** version 2.1 of the License, or (at your option) any later version.
**
** This library is distributed in the hope that it will be useful,
** but WITHOUT ANY WARRANTY; without even the implied warranty of
** MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
** Lesser General Public License for more details.
*********************************************************************/

#include "cdrv.h"
#include "qtapp.h"
#include <QTimer>
#include <QFont>
#include <QPixmap>
#include <QIcon>
#include <QImage>
#include <QWidget>
#include <QAction>
#include <QActionGroup>
#include <QMenu>
#include <QMenuBar>
#include <QToolBar>
#include <QStatusBar>
#include <QDockWidget>
#include <QSystemTrayIcon>
#include <QTabWidget>
#include <QToolBox>
#include <QLayout>
#include <QBoxLayout>
#include <QHBoxLayout>
#include <QVBoxLayout>
#include <QStackedLayout>
#include <QAbstractButton>
#include <QPushButton>
#include <QCheckBox>
#include <QRadioButton>
#include <QToolButton>
#include <QFrame>
#include <QLabel>
#include <QGroupBox>
#include <QDialog>
#include <QComboBox>
#include <QLineEdit>
#include <QAbstractSlider>
#include <QSlider>
#include <QScrollBar>
#include <QDial>
#include <QBrush>
#include <QPen>
#include <QPainter>
#include <QListWidget>
#include <QMainWindow>
// drvclass enums
enum DRVCLASS_ENUM {
    DRVCLASS_NONE = 0,
    _CLASSID_APP,
    CLASSID_TIMER,
    CLASSID_FONT,
    CLASSID_PIXMAP,
    CLASSID_ICON,
    CLASSID_IMAGE,
    CLASSID_WIDGET,
    CLASSID_ACTION,
    CLASSID_ACTIONGROUP,
    CLASSID_MENU,
    CLASSID_MENUBAR,
    CLASSID_TOOLBAR,
    CLASSID_STATUSBAR,
    CLASSID_DOCKWIDGET,
    CLASSID_SYSTEMTRAY,
    CLASSID_TABWIDGET,
    CLASSID_TOOLBOX,
    _CLASSID_BASELAYOUT,
    CLASSID_BOXLAYOUT,
    CLASSID_HBOXLAYOUT,
    CLASSID_VBOXLAYOUT,
    CLASSID_STACKEDLAYOUT,
    _CLASSID_BASEBUTTON,
    CLASSID_BUTTON,
    CLASSID_CHECKBOX,
    CLASSID_RADIO,
    CLASSID_TOOLBUTTON,
    CLASSID_FRAME,
    CLASSID_LABEL,
    CLASSID_GROUPBOX,
    CLASSID_DIALOG,
    CLASSID_COMBOBOX,
    CLASSID_LINEEDIT,
    _CLASSID_BASESLIDER,
    CLASSID_SLIDER,
    CLASSID_SCROLLBAR,
    CLASSID_DIAL,
    CLASSID_BRUSH,
    CLASSID_PEN,
    CLASSID_PAINTER,
    CLASSID_LISTWIDGETITEM,
    CLASSID_LISTWIDGET,
    CLASSID_MAINWINDOW,
    DRVCLASS_LAST
};
// _CLASSID_APP drvid enums
enum DRVID_APP_ENUM {
    _ID_APP_NONE = 0,
    _ID_APP_APPMAIN,
    _ID_APP_RUN,
    _ID_APP_EXIT,
    _ID_APP_SETFONT,
    _ID_APP_FONT,
    _ID_APP_SETACTIVEWINDOW,
    _ID_APP_ACTIVEWINDOW,
    _ID_APP_SETSTYLESHEET,
    _ID_APP_STYLESHEET,
    _ID_APP_CLOSEALLWINDOWS,
    _ID_APP_ONREMOVEOBJECT,
    _ID_APP_LAST
};
// CLASSID_TIMER drvid enums
enum DRVID_TIMER_ENUM {
    _ID_TIMER_NONE = 0,
    _ID_TIMER_INIT,
    _ID_TIMER_CLOSE,
    _ID_TIMER_SETINTERVAL,
    _ID_TIMER_INTERVAL,
    _ID_TIMER_SETSINGLESHOT,
    _ID_TIMER_ISSINGLESHOT,
    _ID_TIMER_ISACTIVE,
    _ID_TIMER_TIMERID,
    _ID_TIMER_START,
    _ID_TIMER_STARTWITH,
    _ID_TIMER_STOP,
    _ID_TIMER_ONTIMEOUT,
    _ID_TIMER_LAST
};
// CLASSID_FONT drvid enums
enum DRVID_FONT_ENUM {
    _ID_FONT_NONE = 0,
    _ID_FONT_INIT,
    _ID_FONT_INITWITH,
    _ID_FONT_CLOSE,
    _ID_FONT_SETFAMILY,
    _ID_FONT_FAMILY,
    _ID_FONT_SETPOINTSIZE,
    _ID_FONT_POINTSIZE,
    _ID_FONT_SETPOINTSIZEF,
    _ID_FONT_POINTSIZEF,
    _ID_FONT_SETBOLD,
    _ID_FONT_BOLD,
    _ID_FONT_SETITALIC,
    _ID_FONT_ITALIC,
    _ID_FONT_SETKERNING,
    _ID_FONT_KERNING,
    _ID_FONT_SETLETTERSPACING,
    _ID_FONT_LETTERSPACING,
    _ID_FONT_SETOVERLINE,
    _ID_FONT_OVERLINE,
    _ID_FONT_SETPIXELSIZE,
    _ID_FONT_PIXELSIZE,
    _ID_FONT_SETRAWMODE,
    _ID_FONT_RAWMODE,
    _ID_FONT_SETRAWNAME,
    _ID_FONT_RAWNAME,
    _ID_FONT_SETWEIGHT,
    _ID_FONT_WEIGHT,
    _ID_FONT_SETFIXEDPITCH,
    _ID_FONT_FIXEDPITCH,
    _ID_FONT_SETSTRETCH,
    _ID_FONT_STRETCH,
    _ID_FONT_SETSTRIKEOUT,
    _ID_FONT_STRIKEOUT,
    _ID_FONT_SETUNDERLINE,
    _ID_FONT_UNDERLINE,
    _ID_FONT_LAST
};
// CLASSID_PIXMAP drvid enums
enum DRVID_PIXMAP_ENUM {
    _ID_PIXMAP_NONE = 0,
    _ID_PIXMAP_INIT,
    _ID_PIXMAP_INITWITHFILE,
    _ID_PIXMAP_INITWITHDATA,
    _ID_PIXMAP_INITWITHSIZE,
    _ID_PIXMAP_CLOSE,
    _ID_PIXMAP_DEPTH,
    _ID_PIXMAP_SIZE,
    _ID_PIXMAP_RECT,
    _ID_PIXMAP_HASALPHA,
    _ID_PIXMAP_HASALPHACHANNEL,
    _ID_PIXMAP_ISNULL,
    _ID_PIXMAP_WIDTH,
    _ID_PIXMAP_HEIGHT,
    _ID_PIXMAP_SCALED,
    _ID_PIXMAP_SCALEDTOHEIGHT,
    _ID_PIXMAP_SCALEDTOWIDTH,
    _ID_PIXMAP_TOIMAGE,
    _ID_PIXMAP_LOAD,
    _ID_PIXMAP_SAVE,
    _ID_PIXMAP_FILL,
    _ID_PIXMAP_LAST
};
// CLASSID_ICON drvid enums
enum DRVID_ICON_ENUM {
    _ID_ICON_NONE = 0,
    _ID_ICON_INIT,
    _ID_ICON_INITWITHFILE,
    _ID_ICON_INITWITHPIXMAP,
    _ID_ICON_CLOSE,
    _ID_ICON_LAST
};
// CLASSID_IMAGE drvid enums
enum DRVID_IMAGE_ENUM {
    _ID_IMAGE_NONE = 0,
    _ID_IMAGE_INIT,
    _ID_IMAGE_INITWITHFILE,
    _ID_IMAGE_INITWITHSIZE,
    _ID_IMAGE_CLOSE,
    _ID_IMAGE_DEPTH,
    _ID_IMAGE_SIZE,
    _ID_IMAGE_RECT,
    _ID_IMAGE_FILL,
    _ID_IMAGE_SCALED,
    _ID_IMAGE_LAST
};
// CLASSID_WIDGET drvid enums
enum DRVID_WIDGET_ENUM {
    _ID_WIDGET_NONE = 0,
    _ID_WIDGET_INIT,
    _ID_WIDGET_CLOSE,
    _ID_WIDGET_SETAUTOCLOSE,
    _ID_WIDGET_AUTOCLOSE,
    _ID_WIDGET_SETLAYOUT,
    _ID_WIDGET_LAYOUT,
    _ID_WIDGET_SETPARENT,
    _ID_WIDGET_PARENT,
    _ID_WIDGET_SETVISIBLE,
    _ID_WIDGET_ISVISIBLE,
    _ID_WIDGET_SETWINDOWTITLE,
    _ID_WIDGET_WINDOWTITLE,
    _ID_WIDGET_SETWINDOWICON,
    _ID_WIDGET_WINDOWICON,
    _ID_WIDGET_SETPOS,
    _ID_WIDGET_POS,
    _ID_WIDGET_SETSIZE,
    _ID_WIDGET_SIZE,
    _ID_WIDGET_SETMINIMUMSIZE,
    _ID_WIDGET_MINIMUMSIZE,
    _ID_WIDGET_SETMAXIMUMSIZE,
    _ID_WIDGET_MAXIMUMSIZE,
    _ID_WIDGET_SETFIXEDSIZE,
    _ID_WIDGET_SETGEOMETRY,
    _ID_WIDGET_GEOMETRY,
    _ID_WIDGET_SETFONT,
    _ID_WIDGET_FONT,
    _ID_WIDGET_SETTOOLTIP,
    _ID_WIDGET_TOOLTIP,
    _ID_WIDGET_SETUPDATESENABLED,
    _ID_WIDGET_UPDATESENABLED,
    _ID_WIDGET_ISACTIVATEWINDOW,
    _ID_WIDGET_ACTIVATEWINDOW,
    _ID_WIDGET_DONE,
    _ID_WIDGET_UPDATE,
    _ID_WIDGET_REPAINT,
    _ID_WIDGET_STARTTIMER,
    _ID_WIDGET_KILLTIMER,
    _ID_WIDGET_ADDACTION,
    _ID_WIDGET_INSERTACTION,
    _ID_WIDGET_REMOVEACTION,
    _ID_WIDGET_ONSHOWEVENT,
    _ID_WIDGET_ONHIDEEVENT,
    _ID_WIDGET_ONCLOSEEVENT,
    _ID_WIDGET_ONKEYPRESSEVENT,
    _ID_WIDGET_ONKEYRELEASEEVENT,
    _ID_WIDGET_ONMOUSEPRESSEVENT,
    _ID_WIDGET_ONMOUSERELEASEEVENT,
    _ID_WIDGET_ONMOUSEMOVEEVENT,
    _ID_WIDGET_ONMOUSEDOUBLECLICKEVENT,
    _ID_WIDGET_ONMOVEEVENT,
    _ID_WIDGET_ONPAINTEVENT,
    _ID_WIDGET_ONRESIZEEVENT,
    _ID_WIDGET_ONENTEREVENT,
    _ID_WIDGET_ONLEAVEEVENT,
    _ID_WIDGET_ONFOCUSINEVENT,
    _ID_WIDGET_ONFOCUSOUTEVENT,
    _ID_WIDGET_ONTIMEREVENT,
    _ID_WIDGET_LAST
};
// CLASSID_ACTION drvid enums
enum DRVID_ACTION_ENUM {
    _ID_ACTION_NONE = 0,
    _ID_ACTION_INIT,
    _ID_ACTION_INITWITHTEXT,
    _ID_ACTION_CLOSE,
    _ID_ACTION_SETTEXT,
    _ID_ACTION_TEXT,
    _ID_ACTION_SETICON,
    _ID_ACTION_ICON,
    _ID_ACTION_SETICONTEXT,
    _ID_ACTION_ICONTEXT,
    _ID_ACTION_SETTOOLTIP,
    _ID_ACTION_TOOLTIP,
    _ID_ACTION_SETFONT,
    _ID_ACTION_FONT,
    _ID_ACTION_SETCHECKABLE,
    _ID_ACTION_ISCHECKABLE,
    _ID_ACTION_SETCHECKED,
    _ID_ACTION_ISCHECKED,
    _ID_ACTION_SETVISIBLE,
    _ID_ACTION_ISVISIBLE,
    _ID_ACTION_ONTRIGGERED,
    _ID_ACTION_LAST
};
// CLASSID_ACTIONGROUP drvid enums
enum DRVID_ACTIONGROUP_ENUM {
    _ID_ACTIONGROUP_NONE = 0,
    _ID_ACTIONGROUP_INIT,
    _ID_ACTIONGROUP_CLOSE,
    _ID_ACTIONGROUP_SETVISIBLE,
    _ID_ACTIONGROUP_ISVISIBLE,
    _ID_ACTIONGROUP_SETENABLED,
    _ID_ACTIONGROUP_ISENABLED,
    _ID_ACTIONGROUP_SETEXCLUSIVE,
    _ID_ACTIONGROUP_ISEXCLUSIVE,
    _ID_ACTIONGROUP_ADDACTION,
    _ID_ACTIONGROUP_REMOVEACTION,
    _ID_ACTIONGROUP_ONHOVERED,
    _ID_ACTIONGROUP_ONTRIGGERED,
    _ID_ACTIONGROUP_LAST
};
// CLASSID_MENU drvid enums
enum DRVID_MENU_ENUM {
    _ID_MENU_NONE = 0,
    _ID_MENU_INIT,
    _ID_MENU_INITWITHTITLE,
    _ID_MENU_SETTITLE,
    _ID_MENU_TITLE,
    _ID_MENU_SETICON,
    _ID_MENU_ICON,
    _ID_MENU_SETDEFAULTACTION,
    _ID_MENU_DEFAULTACTION,
    _ID_MENU_SETACTIVEACTION,
    _ID_MENU_ACTIVEACTION,
    _ID_MENU_MENUACTION,
    _ID_MENU_ISEMPTY,
    _ID_MENU_ADDMENU,
    _ID_MENU_INSERTMENU,
    _ID_MENU_ADDSEPARATOR,
    _ID_MENU_INSERTSEPARATOR,
    _ID_MENU_CLEAR,
    _ID_MENU_POPUP,
    _ID_MENU_ONHOVERED,
    _ID_MENU_ONTRIGGERED,
    _ID_MENU_LAST
};
// CLASSID_MENUBAR drvid enums
enum DRVID_MENUBAR_ENUM {
    _ID_MENUBAR_NONE = 0,
    _ID_MENUBAR_INIT,
    _ID_MENUBAR_SETACTIVEACTION,
    _ID_MENUBAR_ACTIVEACTION,
    _ID_MENUBAR_SETNATIVEMENUBAR,
    _ID_MENUBAR_ISNATIVEMENUBAR,
    _ID_MENUBAR_ADDMENU,
    _ID_MENUBAR_INSERTMENU,
    _ID_MENUBAR_ADDSEPARATOR,
    _ID_MENUBAR_INSERTSEPARATOR,
    _ID_MENUBAR_CLEAR,
    _ID_MENUBAR_ONHOVERED,
    _ID_MENUBAR_ONTRIGGERED,
    _ID_MENUBAR_LAST
};
// CLASSID_TOOLBAR drvid enums
enum DRVID_TOOLBAR_ENUM {
    _ID_TOOLBAR_NONE = 0,
    _ID_TOOLBAR_INIT,
    _ID_TOOLBAR_INITWITHTITLE,
    _ID_TOOLBAR_SETTITLE,
    _ID_TOOLBAR_TITLE,
    _ID_TOOLBAR_SETICONSIZE,
    _ID_TOOLBAR_ICONSIZE,
    _ID_TOOLBAR_SETFLOATABLE,
    _ID_TOOLBAR_ISFLOATABLE,
    _ID_TOOLBAR_ISFLOATING,
    _ID_TOOLBAR_SETTOOLBUTTONSTYLE,
    _ID_TOOLBAR_TOOLBUTTONSTYLE,
    _ID_TOOLBAR_ADDACTION,
    _ID_TOOLBAR_ADDSEPARATOR,
    _ID_TOOLBAR_ADDWIDGET,
    _ID_TOOLBAR_CLEAR,
    _ID_TOOLBAR_LAST
};
// CLASSID_STATUSBAR drvid enums
enum DRVID_STATUSBAR_ENUM {
    _ID_STATUSBAR_NONE = 0,
    _ID_STATUSBAR_INIT,
    _ID_STATUSBAR_SETSIZEGRIP,
    _ID_STATUSBAR_ISSIZEGRIP,
    _ID_STATUSBAR_ADDWIDGET,
    _ID_STATUSBAR_ADDPERMANENTWIDGET,
    _ID_STATUSBAR_INSERTWIDGET,
    _ID_STATUSBAR_INSERTPERMANENTWIDGET,
    _ID_STATUSBAR_REMOVEWIDGET,
    _ID_STATUSBAR_SHOWMESSAGE,
    _ID_STATUSBAR_CURRENTMESSAGE,
    _ID_STATUSBAR_CLEARMESSAGE,
    _ID_STATUSBAR_LAST
};
// CLASSID_DOCKWIDGET drvid enums
enum DRVID_DOCKWIDGET_ENUM {
    _ID_DOCKWIDGET_NONE = 0,
    _ID_DOCKWIDGET_INIT,
    _ID_DOCKWIDGET_INITWITHTITLE,
    _ID_DOCKWIDGET_SETDOCK,
    _ID_DOCKWIDGET_DOCK,
    _ID_DOCKWIDGET_SETTITLEBAR,
    _ID_DOCKWIDGET_TITLEBAR,
    _ID_DOCKWIDGET_SETFLOATING,
    _ID_DOCKWIDGET_ISFLOATING,
    _ID_DOCKWIDGET_ONVISIBILITYCHANGED,
    _ID_DOCKWIDGET_LAST
};
// CLASSID_SYSTEMTRAY drvid enums
enum DRVID_SYSTEMTRAY_ENUM {
    _ID_SYSTEMTRAY_NONE = 0,
    _ID_SYSTEMTRAY_INIT,
    _ID_SYSTEMTRAY_CLOSE,
    _ID_SYSTEMTRAY_SETCONTEXTMENU,
    _ID_SYSTEMTRAY_CONTEXTMENU,
    _ID_SYSTEMTRAY_SETICON,
    _ID_SYSTEMTRAY_ICON,
    _ID_SYSTEMTRAY_SETTOOLTIP,
    _ID_SYSTEMTRAY_TOOLTIP,
    _ID_SYSTEMTRAY_SETVISIBLE,
    _ID_SYSTEMTRAY_ISVISIBLE,
    _ID_SYSTEMTRAY_SHOWMESSAGE,
    _ID_SYSTEMTRAY_ONACTIVATED,
    _ID_SYSTEMTRAY_ONMESSAGECLICKED,
    _ID_SYSTEMTRAY_LAST
};
// CLASSID_TABWIDGET drvid enums
enum DRVID_TABWIDGET_ENUM {
    _ID_TABWIDGET_NONE = 0,
    _ID_TABWIDGET_INIT,
    _ID_TABWIDGET_COUNT,
    _ID_TABWIDGET_CURRENTINDEX,
    _ID_TABWIDGET_CURRENTWIDGET,
    _ID_TABWIDGET_SETCURRENTINDEX,
    _ID_TABWIDGET_SETCURRENTWIDGET,
    _ID_TABWIDGET_ADDTAB,
    _ID_TABWIDGET_INSERTTAB,
    _ID_TABWIDGET_REMOVETAB,
    _ID_TABWIDGET_CLEAR,
    _ID_TABWIDGET_INDEXOF,
    _ID_TABWIDGET_WIDGETOF,
    _ID_TABWIDGET_SETTABTEXT,
    _ID_TABWIDGET_TABTEXT,
    _ID_TABWIDGET_SETTABICON,
    _ID_TABWIDGET_TABICON,
    _ID_TABWIDGET_SETTABTOOLTIP,
    _ID_TABWIDGET_TABTOOLTIP,
    _ID_TABWIDGET_ONCURRENTCHANGED,
    _ID_TABWIDGET_LAST
};
// CLASSID_TOOLBOX drvid enums
enum DRVID_TOOLBOX_ENUM {
    _ID_TOOLBOX_NONE = 0,
    _ID_TOOLBOX_INIT,
    _ID_TOOLBOX_SETCURRENTINDEX,
    _ID_TOOLBOX_SETCURRENTWIDGET,
    _ID_TOOLBOX_CURRENTINDEX,
    _ID_TOOLBOX_CURRENTWIDGET,
    _ID_TOOLBOX_COUNT,
    _ID_TOOLBOX_ADDITEM,
    _ID_TOOLBOX_INSERTITEM,
    _ID_TOOLBOX_REMOVEITEM,
    _ID_TOOLBOX_INDEXOF,
    _ID_TOOLBOX_WIDGETOF,
    _ID_TOOLBOX_SETITEMTEXT,
    _ID_TOOLBOX_ITEMTEXT,
    _ID_TOOLBOX_SETITEMICON,
    _ID_TOOLBOX_ITEMICON,
    _ID_TOOLBOX_SETITEMTOOLTIP,
    _ID_TOOLBOX_ITEMTOOLTIP,
    _ID_TOOLBOX_ISITEMENABLED,
    _ID_TOOLBOX_ONCURRENTCHANGED,
    _ID_TOOLBOX_LAST
};
// _CLASSID_BASELAYOUT drvid enums
enum DRVID_BASELAYOUT_ENUM {
    _ID_BASELAYOUT_NONE = 0,
    _ID_BASELAYOUT_SETSPACING,
    _ID_BASELAYOUT_SPACING,
    _ID_BASELAYOUT_SETMARGINS,
    _ID_BASELAYOUT_MARGINS,
    _ID_BASELAYOUT_SETMENUBAR,
    _ID_BASELAYOUT_MENUBAR,
    _ID_BASELAYOUT_COUNT,
    _ID_BASELAYOUT_ADDLAYOUT,
    _ID_BASELAYOUT_ADDWIDGET,
    _ID_BASELAYOUT_REMOVEWIDGET,
    _ID_BASELAYOUT_INDEXOF,
    _ID_BASELAYOUT_LAST
};
// CLASSID_BOXLAYOUT drvid enums
enum DRVID_BOXLAYOUT_ENUM {
    _ID_BOXLAYOUT_NONE = 0,
    _ID_BOXLAYOUT_INIT,
    _ID_BOXLAYOUT_CLOSE,
    _ID_BOXLAYOUT_SETORIENTATION,
    _ID_BOXLAYOUT_ORIENTATION,
    _ID_BOXLAYOUT_ADDLAYOUTWITH,
    _ID_BOXLAYOUT_ADDWIDGETWITH,
    _ID_BOXLAYOUT_ADDSPACING,
    _ID_BOXLAYOUT_ADDSTRETCH,
    _ID_BOXLAYOUT_LAST
};
// CLASSID_HBOXLAYOUT drvid enums
enum DRVID_HBOXLAYOUT_ENUM {
    _ID_HBOXLAYOUT_NONE = 0,
    _ID_HBOXLAYOUT_INIT,
    _ID_HBOXLAYOUT_LAST
};
// CLASSID_VBOXLAYOUT drvid enums
enum DRVID_VBOXLAYOUT_ENUM {
    _ID_VBOXLAYOUT_NONE = 0,
    _ID_VBOXLAYOUT_INIT,
    _ID_VBOXLAYOUT_LAST
};
// CLASSID_STACKEDLAYOUT drvid enums
enum DRVID_STACKEDLAYOUT_ENUM {
    _ID_STACKEDLAYOUT_NONE = 0,
    _ID_STACKEDLAYOUT_INIT,
    _ID_STACKEDLAYOUT_SETCURRENTINDEX,
    _ID_STACKEDLAYOUT_CURRENTINDEX,
    _ID_STACKEDLAYOUT_SETCURRENTWIDGET,
    _ID_STACKEDLAYOUT_CURRENTWIDGET,
    _ID_STACKEDLAYOUT_ADDWIDGET,
    _ID_STACKEDLAYOUT_INSERTWIDGET,
    _ID_STACKEDLAYOUT_WIDGET,
    _ID_STACKEDLAYOUT_ONCURRENTCHANGED,
    _ID_STACKEDLAYOUT_LAST
};
// _CLASSID_BASEBUTTON drvid enums
enum DRVID_BASEBUTTON_ENUM {
    _ID_BASEBUTTON_NONE = 0,
    _ID_BASEBUTTON_SETTEXT,
    _ID_BASEBUTTON_TEXT,
    _ID_BASEBUTTON_SETICON,
    _ID_BASEBUTTON_ICON,
    _ID_BASEBUTTON_SETICONSIZE,
    _ID_BASEBUTTON_ICONSIZE,
    _ID_BASEBUTTON_SETCHECKABLE,
    _ID_BASEBUTTON_ISCHECKABLE,
    _ID_BASEBUTTON_SETDOWN,
    _ID_BASEBUTTON_ISDOWN,
    _ID_BASEBUTTON_LAST
};
// CLASSID_BUTTON drvid enums
enum DRVID_BUTTON_ENUM {
    _ID_BUTTON_NONE = 0,
    _ID_BUTTON_INIT,
    _ID_BUTTON_INITWITHTEXT,
    _ID_BUTTON_SETFLAT,
    _ID_BUTTON_ISFLAT,
    _ID_BUTTON_SETDEFAULT,
    _ID_BUTTON_ISDEFAULT,
    _ID_BUTTON_SETMENU,
    _ID_BUTTON_MENU,
    _ID_BUTTON_ONCLICKED,
    _ID_BUTTON_LAST
};
// CLASSID_CHECKBOX drvid enums
enum DRVID_CHECKBOX_ENUM {
    _ID_CHECKBOX_NONE = 0,
    _ID_CHECKBOX_INIT,
    _ID_CHECKBOX_INITWITHTEXT,
    _ID_CHECKBOX_SETCHECK,
    _ID_CHECKBOX_CHECK,
    _ID_CHECKBOX_SETTRISTATE,
    _ID_CHECKBOX_ISTRISTATE,
    _ID_CHECKBOX_ONSTATECHANGED,
    _ID_CHECKBOX_LAST
};
// CLASSID_RADIO drvid enums
enum DRVID_RADIO_ENUM {
    _ID_RADIO_NONE = 0,
    _ID_RADIO_INIT,
    _ID_RADIO_INITWITHTEXT,
    _ID_RADIO_ONCLICKED,
    _ID_RADIO_LAST
};
// CLASSID_TOOLBUTTON drvid enums
enum DRVID_TOOLBUTTON_ENUM {
    _ID_TOOLBUTTON_NONE = 0,
    _ID_TOOLBUTTON_INIT,
    _ID_TOOLBUTTON_INITWITHTEXT,
    _ID_TOOLBUTTON_SETMENU,
    _ID_TOOLBUTTON_MENU,
    _ID_TOOLBUTTON_SETAUTORAISE,
    _ID_TOOLBUTTON_AUTORAISE,
    _ID_TOOLBUTTON_SETTOOLBUTTONSTYLE,
    _ID_TOOLBUTTON_TOOLBUTTONSTYLE,
    _ID_TOOLBUTTON_SETTOOLBUTTONPOPUPMODE,
    _ID_TOOLBUTTON_TOOLBUTTONPOPUPMODE,
    _ID_TOOLBUTTON_ONCLICKED,
    _ID_TOOLBUTTON_LAST
};
// CLASSID_FRAME drvid enums
enum DRVID_FRAME_ENUM {
    _ID_FRAME_NONE = 0,
    _ID_FRAME_INIT,
    _ID_FRAME_SETFRAMESTYLE,
    _ID_FRAME_FRAMESTYLE,
    _ID_FRAME_SETFRAMERECT,
    _ID_FRAME_FRAMERECT,
    _ID_FRAME_LAST
};
// CLASSID_LABEL drvid enums
enum DRVID_LABEL_ENUM {
    _ID_LABEL_NONE = 0,
    _ID_LABEL_INIT,
    _ID_LABEL_INITWITHTEXT,
    _ID_LABEL_INITWITHPIXMAP,
    _ID_LABEL_SETTEXT,
    _ID_LABEL_TEXT,
    _ID_LABEL_SETWORDWRAP,
    _ID_LABEL_WORDWRAP,
    _ID_LABEL_SETTEXTFORMAT,
    _ID_LABEL_TEXTFORMAT,
    _ID_LABEL_SETPIXMAP,
    _ID_LABEL_PIXMAP,
    _ID_LABEL_ONLINKACTIVATED,
    _ID_LABEL_LAST
};
// CLASSID_GROUPBOX drvid enums
enum DRVID_GROUPBOX_ENUM {
    _ID_GROUPBOX_NONE = 0,
    _ID_GROUPBOX_INIT,
    _ID_GROUPBOX_INITWITHTITLE,
    _ID_GROUPBOX_SETTITLE,
    _ID_GROUPBOX_TITLE,
    _ID_GROUPBOX_LAST
};
// CLASSID_DIALOG drvid enums
enum DRVID_DIALOG_ENUM {
    _ID_DIALOG_NONE = 0,
    _ID_DIALOG_INIT,
    _ID_DIALOG_SETMODAL,
    _ID_DIALOG_ISMODAL,
    _ID_DIALOG_SETRESULT,
    _ID_DIALOG_RESULT,
    _ID_DIALOG_OPEN,
    _ID_DIALOG_EXEC,
    _ID_DIALOG_DONE,
    _ID_DIALOG_ACCEPT,
    _ID_DIALOG_REJECT,
    _ID_DIALOG_ONACCEPTED,
    _ID_DIALOG_ONREJECTED,
    _ID_DIALOG_LAST
};
// CLASSID_COMBOBOX drvid enums
enum DRVID_COMBOBOX_ENUM {
    _ID_COMBOBOX_NONE = 0,
    _ID_COMBOBOX_INIT,
    _ID_COMBOBOX_COUNT,
    _ID_COMBOBOX_SETCURRENTINDEX,
    _ID_COMBOBOX_CURRENTINDEX,
    _ID_COMBOBOX_CURRENTTEXT,
    _ID_COMBOBOX_SETEDITABLE,
    _ID_COMBOBOX_ISEDITABLE,
    _ID_COMBOBOX_SETMAXCOUNT,
    _ID_COMBOBOX_MAXCOUNT,
    _ID_COMBOBOX_SETMAXVISIBLEITEMS,
    _ID_COMBOBOX_MAXVISIBLEITEMS,
    _ID_COMBOBOX_SETMINIMUMCONTENTSLENGTH,
    _ID_COMBOBOX_MINIMUNCONTENTSLENGHT,
    _ID_COMBOBOX_ADDITEM,
    _ID_COMBOBOX_INSERTITEM,
    _ID_COMBOBOX_REMOVEITEM,
    _ID_COMBOBOX_ITEMTEXT,
    _ID_COMBOBOX_ONCURRENTINDEXCHANGED,
    _ID_COMBOBOX_LAST
};
// CLASSID_LINEEDIT drvid enums
enum DRVID_LINEEDIT_ENUM {
    _ID_LINEEDIT_NONE = 0,
    _ID_LINEEDIT_INIT,
    _ID_LINEEDIT_INITWITHTEXT,
    _ID_LINEEDIT_SETTEXT,
    _ID_LINEEDIT_TEXT,
    _ID_LINEEDIT_SETINPUTMASK,
    _ID_LINEEDIT_INPUTMASK,
    _ID_LINEEDIT_SETALIGNMENT,
    _ID_LINEEDIT_ALIGNMENT,
    _ID_LINEEDIT_SETCURSORPOS,
    _ID_LINEEDIT_CURSORPOS,
    _ID_LINEEDIT_SETDRAGENABLED,
    _ID_LINEEDIT_DRAGENABLED,
    _ID_LINEEDIT_SETREADONLY,
    _ID_LINEEDIT_ISREADONLY,
    _ID_LINEEDIT_SETFRAME,
    _ID_LINEEDIT_HASFRAME,
    _ID_LINEEDIT_ISREDOAVAILABLE,
    _ID_LINEEDIT_HASSELECTED,
    _ID_LINEEDIT_SELECTEDTEXT,
    _ID_LINEEDIT_SELSTART,
    _ID_LINEEDIT_SETSEL,
    _ID_LINEEDIT_CANCELSEL,
    _ID_LINEEDIT_SELECTALL,
    _ID_LINEEDIT_COPY,
    _ID_LINEEDIT_CUT,
    _ID_LINEEDIT_PASTE,
    _ID_LINEEDIT_REDO,
    _ID_LINEEDIT_UNDO,
    _ID_LINEEDIT_CLEAR,
    _ID_LINEEDIT_ONTEXTCHANGED,
    _ID_LINEEDIT_ONEDITINGFINISHED,
    _ID_LINEEDIT_ONRETURNPRESSED,
    _ID_LINEEDIT_LAST
};
// _CLASSID_BASESLIDER drvid enums
enum DRVID_BASESLIDER_ENUM {
    _ID_BASESLIDER_NONE = 0,
    _ID_BASESLIDER_SETTRACKING,
    _ID_BASESLIDER_HASTRACKING,
    _ID_BASESLIDER_SETMAXIMUM,
    _ID_BASESLIDER_MAXIMUM,
    _ID_BASESLIDER_SETMINIMUM,
    _ID_BASESLIDER_MINIMUM,
    _ID_BASESLIDER_SETORIENTATION,
    _ID_BASESLIDER_ORIENTATION,
    _ID_BASESLIDER_SETPAGESTEP,
    _ID_BASESLIDER_PAGESTEP,
    _ID_BASESLIDER_SETSINGLESTEP,
    _ID_BASESLIDER_SINGLESTEP,
    _ID_BASESLIDER_SETSLIDERDOWN,
    _ID_BASESLIDER_ISSLIDERDOWN,
    _ID_BASESLIDER_SETSLIDERPOSITION,
    _ID_BASESLIDER_SLIDERPOSITION,
    _ID_BASESLIDER_SETVALUE,
    _ID_BASESLIDER_VALUE,
    _ID_BASESLIDER_SETRANGE,
    _ID_BASESLIDER_RANGE,
    _ID_BASESLIDER_ONVALUECHANGED,
    _ID_BASESLIDER_ONSLIDERPRESSED,
    _ID_BASESLIDER_ONSLIDERRELEASED,
    _ID_BASESLIDER_ONSLIDERMOVED,
    _ID_BASESLIDER_LAST
};
// CLASSID_SLIDER drvid enums
enum DRVID_SLIDER_ENUM {
    _ID_SLIDER_NONE = 0,
    _ID_SLIDER_INIT,
    _ID_SLIDER_SETTICKINTERVAL,
    _ID_SLIDER_TICKINTERVAL,
    _ID_SLIDER_SETTICKPOSITION,
    _ID_SLIDER_TICKPOSITION,
    _ID_SLIDER_LAST
};
// CLASSID_SCROLLBAR drvid enums
enum DRVID_SCROLLBAR_ENUM {
    _ID_SCROLLBAR_NONE = 0,
    _ID_SCROLLBAR_INIT,
    _ID_SCROLLBAR_LAST
};
// CLASSID_DIAL drvid enums
enum DRVID_DIAL_ENUM {
    _ID_DIAL_NONE = 0,
    _ID_DIAL_INIT,
    _ID_DIAL_NOTCHSIZE,
    _ID_DIAL_SETNOTCHTARGET,
    _ID_DIAL_NOTCHTARGET,
    _ID_DIAL_SETNOTCHESVISIBLE,
    _ID_DIAL_NOTCHESVISIBLE,
    _ID_DIAL_SETWRAPPING,
    _ID_DIAL_WRAPPING,
    _ID_DIAL_LAST
};
// CLASSID_BRUSH drvid enums
enum DRVID_BRUSH_ENUM {
    _ID_BRUSH_NONE = 0,
    _ID_BRUSH_INIT,
    _ID_BRUSH_CLOSE,
    _ID_BRUSH_SETCOLOR,
    _ID_BRUSH_COLOR,
    _ID_BRUSH_SETSTYLE,
    _ID_BRUSH_STYLE,
    _ID_BRUSH_LAST
};
// CLASSID_PEN drvid enums
enum DRVID_PEN_ENUM {
    _ID_PEN_NONE = 0,
    _ID_PEN_INIT,
    _ID_PEN_CLOSE,
    _ID_PEN_SETCOLOR,
    _ID_PEN_COLOR,
    _ID_PEN_SETWIDTH,
    _ID_PEN_WIDTH,
    _ID_PEN_SETSTYLE,
    _ID_PEN_STYLE,
    _ID_PEN_LAST
};
// CLASSID_PAINTER drvid enums
enum DRVID_PAINTER_ENUM {
    _ID_PAINTER_NONE = 0,
    _ID_PAINTER_INIT,
    _ID_PAINTER_INITWITHIMAGE,
    _ID_PAINTER_CLOSE,
    _ID_PAINTER_INITFROM,
    _ID_PAINTER_BEGIN,
    _ID_PAINTER_END,
    _ID_PAINTER_SETFONT,
    _ID_PAINTER_FONT,
    _ID_PAINTER_SETPEN,
    _ID_PAINTER_PEN,
    _ID_PAINTER_SETBRUSH,
    _ID_PAINTER_BRUSH,
    _ID_PAINTER_DRAWPOINT,
    _ID_PAINTER_DRAWPOINTS,
    _ID_PAINTER_DRAWLINE,
    _ID_PAINTER_DRAWLINES,
    _ID_PAINTER_DRAWPOLYLINE,
    _ID_PAINTER_DRAWPOLYGON,
    _ID_PAINTER_DRAWRECT,
    _ID_PAINTER_DRAWRECTS,
    _ID_PAINTER_DRAWROUNDEDRECT,
    _ID_PAINTER_DRAWELLIPSE,
    _ID_PAINTER_DRAWARC,
    _ID_PAINTER_DRAWCHORD,
    _ID_PAINTER_DRAWPIE,
    _ID_PAINTER_DRAWTEXT,
    _ID_PAINTER_DRAWTEXTRECT,
    _ID_PAINTER_DRAWPIXMAP,
    _ID_PAINTER_DRAWPIXMAPEX,
    _ID_PAINTER_DRAWPIXMAPRECT,
    _ID_PAINTER_DRAWPIXMAPRECTEX,
    _ID_PAINTER_DRAWIMAGE,
    _ID_PAINTER_DRAWIMAGEEX,
    _ID_PAINTER_DRAWIMAGERECT,
    _ID_PAINTER_DRAWIMAGERECTEX,
    _ID_PAINTER_FILLRECT,
    _ID_PAINTER_FILLRECTF,
    _ID_PAINTER_LAST
};
// CLASSID_LISTWIDGETITEM drvid enums
enum DRVID_LISTWIDGETITEM_ENUM {
    _ID_LISTWIDGETITEM_NONE = 0,
    _ID_LISTWIDGETITEM_INIT,
    _ID_LISTWIDGETITEM_INITWITHTEXT,
    _ID_LISTWIDGETITEM_CLOSE,
    _ID_LISTWIDGETITEM_SETTEXT,
    _ID_LISTWIDGETITEM_TEXT,
    _ID_LISTWIDGETITEM_SETICON,
    _ID_LISTWIDGETITEM_ICON,
    _ID_LISTWIDGETITEM_SETSELECTED,
    _ID_LISTWIDGETITEM_ISSELECTED,
    _ID_LISTWIDGETITEM_SETHIDDEN,
    _ID_LISTWIDGETITEM_ISHIDDEN,
    _ID_LISTWIDGETITEM_SETFONT,
    _ID_LISTWIDGETITEM_FONT,
    _ID_LISTWIDGETITEM_SETFOREGROUND,
    _ID_LISTWIDGETITEM_FOREGROUND,
    _ID_LISTWIDGETITEM_SETBACKGROUND,
    _ID_LISTWIDGETITEM_BACKGROUND,
    _ID_LISTWIDGETITEM_SETTOOLTIP,
    _ID_LISTWIDGETITEM_TOOLTIP,
    _ID_LISTWIDGETITEM_SETTEXTALIGNMENT,
    _ID_LISTWIDGETITEM_TEXTALIGNMENT,
    _ID_LISTWIDGETITEM_SETFLAGS,
    _ID_LISTWIDGETITEM_FLAGS,
    _ID_LISTWIDGETITEM_LAST
};
// CLASSID_LISTWIDGET drvid enums
enum DRVID_LISTWIDGET_ENUM {
    _ID_LISTWIDGET_NONE = 0,
    _ID_LISTWIDGET_INIT,
    _ID_LISTWIDGET_COUNT,
    _ID_LISTWIDGET_SETCURRENTITEM,
    _ID_LISTWIDGET_CURRENTITEM,
    _ID_LISTWIDGET_SETCURRENTROW,
    _ID_LISTWIDGET_CURRENTROW,
    _ID_LISTWIDGET_ADDITEM,
    _ID_LISTWIDGET_INSERTITEM,
    _ID_LISTWIDGET_EDITITEM,
    _ID_LISTWIDGET_TAKEITEM,
    _ID_LISTWIDGET_ITEM,
    _ID_LISTWIDGET_CLEAR,
    _ID_LISTWIDGET_ONCURRENTITEMCHANGED,
    _ID_LISTWIDGET_ONCURRENTROWCHANGED,
    _ID_LISTWIDGET_ONITEMACTIVATED,
    _ID_LISTWIDGET_ONITEMCHANGED,
    _ID_LISTWIDGET_ONITEMCLICKED,
    _ID_LISTWIDGET_ONITEMDOUBLECLICKED,
    _ID_LISTWIDGET_ONITEMENTERED,
    _ID_LISTWIDGET_ONITEMPRESSED,
    _ID_LISTWIDGET_ONITEMSELECTIONCHANGED,
    _ID_LISTWIDGET_LAST
};
// CLASSID_MAINWINDOW drvid enums
enum DRVID_MAINWINDOW_ENUM {
    _ID_MAINWINDOW_NONE = 0,
    _ID_MAINWINDOW_INIT,
    _ID_MAINWINDOW_SETCENTRALWIDGET,
    _ID_MAINWINDOW_CENTRALWIDGET,
    _ID_MAINWINDOW_SETMENUBAR,
    _ID_MAINWINDOW_MENUBAR,
    _ID_MAINWINDOW_SETSTATUSBAR,
    _ID_MAINWINDOW_STATUSBAR,
    _ID_MAINWINDOW_ADDTOOLBAR,
    _ID_MAINWINDOW_INSERTTOOLBAR,
    _ID_MAINWINDOW_REMOVETOOLBAR,
    _ID_MAINWINDOW_ADDDOCKWIDGET,
    _ID_MAINWINDOW_REMOVEDOCKWIDGET,
    _ID_MAINWINDOW_LAST
};
int drv_app(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QtApp *self = (QtApp*)drvGetNative(a0);
    switch (drvid) {
    case _ID_APP_APPMAIN: {
        QtApp app;
        handle_head *head =(handle_head*)a0;
        head->native = &app;
        drvSetInt(a1,drv_appmain());
        break;
    }
    case _ID_APP_RUN: {
        drvSetInt(a1,self->exec());
        break;
    }
    case _ID_APP_EXIT: {
        self->exit(drvGetInt(a1));
        break;
    }
    case _ID_APP_SETFONT: {
        self->setFont(drvGetFont(a1));
        break;
    }
    case _ID_APP_FONT: {
        drvSetFont(a1,self->font());
        break;
    }
    case _ID_APP_SETACTIVEWINDOW: {
        self->setActiveWindow(drvGetWidget(a1));
        break;
    }
    case _ID_APP_ACTIVEWINDOW: {
        drvSetHandle(a1,self->activeWindow());
        break;
    }
    case _ID_APP_SETSTYLESHEET: {
        self->setStyleSheet(drvGetString(a1));
        break;
    }
    case _ID_APP_STYLESHEET: {
        drvSetString(a1,self->styleSheet());
        break;
    }
    case _ID_APP_CLOSEALLWINDOWS: {
        self->closeAllWindows();
        break;
    }
    case _ID_APP_ONREMOVEOBJECT: {
        self->pfnDeleteObj = a1;
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_timer(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QTimer *self = (QTimer*)drvGetNative(a0);
    switch (drvid) {
    case _ID_TIMER_INIT: {
        drvNewObj(a0,new QTimer(theApp));
        break;
    }
    case _ID_TIMER_CLOSE: {
        drvDelObj(a0,self);
        break;
    }
    case _ID_TIMER_SETINTERVAL: {
        self->setInterval(drvGetInt(a1));
        break;
    }
    case _ID_TIMER_INTERVAL: {
        drvSetInt(a1,self->interval());
        break;
    }
    case _ID_TIMER_SETSINGLESHOT: {
        self->setSingleShot(drvGetBool(a1));
        break;
    }
    case _ID_TIMER_ISSINGLESHOT: {
        drvSetBool(a1,self->isSingleShot());
        break;
    }
    case _ID_TIMER_ISACTIVE: {
        drvSetBool(a1,self->isActive());
        break;
    }
    case _ID_TIMER_TIMERID: {
        drvSetInt(a1,self->timerId());
        break;
    }
    case _ID_TIMER_START: {
        self->start();
        break;
    }
    case _ID_TIMER_STARTWITH: {
        self->start(drvGetInt(a1));
        break;
    }
    case _ID_TIMER_STOP: {
        self->stop();
        break;
    }
    case _ID_TIMER_ONTIMEOUT: {
        QObject::connect(self,SIGNAL(timeout()),drvNewSignal(self,a1,a2),SLOT(call()));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_font(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QFont *self = (QFont*)drvGetNative(a0);
    switch (drvid) {
    case _ID_FONT_INIT: {
        drvSetFont(a0,QFont());
        break;
    }
    case _ID_FONT_INITWITH: {
        drvSetFont(a0,QFont(drvGetString(a1),drvGetInt(a2),drvGetInt(a3)));
        break;
    }
    case _ID_FONT_CLOSE: {
        drvDelFont(a0,self);
        break;
    }
    case _ID_FONT_SETFAMILY: {
        self->setFamily(drvGetString(a1));
        break;
    }
    case _ID_FONT_FAMILY: {
        drvSetString(a1,self->family());
        break;
    }
    case _ID_FONT_SETPOINTSIZE: {
        self->setPointSize(drvGetInt(a1));
        break;
    }
    case _ID_FONT_POINTSIZE: {
        drvSetInt(a1,self->pointSize());
        break;
    }
    case _ID_FONT_SETPOINTSIZEF: {
        self->setPointSizeF(drvGetFloat64(a1));
        break;
    }
    case _ID_FONT_POINTSIZEF: {
        drvSetFloat64(a1,self->pointSizeF());
        break;
    }
    case _ID_FONT_SETBOLD: {
        self->setBold(drvGetBool(a1));
        break;
    }
    case _ID_FONT_BOLD: {
        drvSetBool(a1,self->bold());
        break;
    }
    case _ID_FONT_SETITALIC: {
        self->setItalic(drvGetBool(a1));
        break;
    }
    case _ID_FONT_ITALIC: {
        drvSetBool(a1,self->italic());
        break;
    }
    case _ID_FONT_SETKERNING: {
        self->setKerning(drvGetBool(a1));
        break;
    }
    case _ID_FONT_KERNING: {
        drvSetBool(a1,self->kerning());
        break;
    }
    case _ID_FONT_SETLETTERSPACING: {
        self->setLetterSpacing((QFont::SpacingType)drvGetInt(a1),drvGetFloat64(a2));
        break;
    }
    case _ID_FONT_LETTERSPACING: {
        drvSetInt(a1,self->letterSpacing());
        drvSetInt(a2,self->letterSpacingType());
        break;
    }
    case _ID_FONT_SETOVERLINE: {
        self->setOverline(drvGetBool(a1));
        break;
    }
    case _ID_FONT_OVERLINE: {
        drvSetBool(a1,self->overline());
        break;
    }
    case _ID_FONT_SETPIXELSIZE: {
        self->setPixelSize(drvGetInt(a1));
        break;
    }
    case _ID_FONT_PIXELSIZE: {
        drvSetInt(a1,self->pixelSize());
        break;
    }
    case _ID_FONT_SETRAWMODE: {
        self->setRawMode(drvGetBool(a1));
        break;
    }
    case _ID_FONT_RAWMODE: {
        drvSetBool(a1,self->rawMode());
        break;
    }
    case _ID_FONT_SETRAWNAME: {
        self->setRawName(drvGetString(a1));
        break;
    }
    case _ID_FONT_RAWNAME: {
        drvSetString(a1,self->rawName());
        break;
    }
    case _ID_FONT_SETWEIGHT: {
        self->setWeight(drvGetInt(a1));
        break;
    }
    case _ID_FONT_WEIGHT: {
        drvSetInt(a1,self->weight());
        break;
    }
    case _ID_FONT_SETFIXEDPITCH: {
        self->setFixedPitch(drvGetBool(a1));
        break;
    }
    case _ID_FONT_FIXEDPITCH: {
        drvSetBool(a1,self->fixedPitch());
        break;
    }
    case _ID_FONT_SETSTRETCH: {
        self->setStretch(drvGetInt(a1));
        break;
    }
    case _ID_FONT_STRETCH: {
        drvSetInt(a1,self->stretch());
        break;
    }
    case _ID_FONT_SETSTRIKEOUT: {
        self->setStrikeOut(drvGetBool(a1));
        break;
    }
    case _ID_FONT_STRIKEOUT: {
        drvSetBool(a1,self->strikeOut());
        break;
    }
    case _ID_FONT_SETUNDERLINE: {
        self->setUnderline(drvGetBool(a1));
        break;
    }
    case _ID_FONT_UNDERLINE: {
        drvSetBool(a1,self->underline());
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_pixmap(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QPixmap *self = (QPixmap*)drvGetNative(a0);
    switch (drvid) {
    case _ID_PIXMAP_INIT: {
        drvSetPixmap(a0, QPixmap());
        break;
    }
    case _ID_PIXMAP_INITWITHFILE: {
        drvSetPixmap(a0, QPixmap(drvGetString(a1)));
        break;
    }
    case _ID_PIXMAP_INITWITHDATA: {
        QPixmap pixmap;
        if (pixmap.loadFromData(drvGetByteArray(a1))) {
        	drvSetPixmap(a0, pixmap);
        }
        break;
    }
    case _ID_PIXMAP_INITWITHSIZE: {
        drvSetPixmap(a0, QPixmap(drvGetInt(a1),drvGetInt(a2)));
        break;
    }
    case _ID_PIXMAP_CLOSE: {
        drvDelPixmap(a0,self);
        break;
    }
    case _ID_PIXMAP_DEPTH: {
        drvSetInt(a1,self->depth());
        break;
    }
    case _ID_PIXMAP_SIZE: {
        drvSetSize(a1,self->size());
        break;
    }
    case _ID_PIXMAP_RECT: {
        drvSetRect(a1,self->rect());
        break;
    }
    case _ID_PIXMAP_HASALPHA: {
        drvSetBool(a1,self->hasAlpha());
        break;
    }
    case _ID_PIXMAP_HASALPHACHANNEL: {
        drvSetBool(a1,self->hasAlphaChannel());
        break;
    }
    case _ID_PIXMAP_ISNULL: {
        drvSetBool(a1,self->isNull());
        break;
    }
    case _ID_PIXMAP_WIDTH: {
        drvSetInt(a1,self->width());
        break;
    }
    case _ID_PIXMAP_HEIGHT: {
        drvSetInt(a1,self->height());
        break;
    }
    case _ID_PIXMAP_SCALED: {
        drvSetPixmap(a5,self->scaled(drvGetInt(a1),drvGetInt(a2),drvGetAspectRatioMode(a3),drvGetTransformationMode(a4)));
        break;
    }
    case _ID_PIXMAP_SCALEDTOHEIGHT: {
        drvSetPixmap(a3,self->scaledToHeight(drvGetInt(a1),drvGetTransformationMode(a2)));
        break;
    }
    case _ID_PIXMAP_SCALEDTOWIDTH: {
        drvSetPixmap(a3,self->scaledToWidth(drvGetInt(a1),drvGetTransformationMode(a2)));
        break;
    }
    case _ID_PIXMAP_TOIMAGE: {
        drvSetImage(a1,self->toImage());
        break;
    }
    case _ID_PIXMAP_LOAD: {
        	drvSetBool(a1, self->load(drvGetString(a0)));
        break;
    }
    case _ID_PIXMAP_SAVE: {
        	drvSetBool(a3, self->save(drvGetString(a1), 0, drvGetInt(a2)));
        break;
    }
    case _ID_PIXMAP_FILL: {
        self->fill(drvGetColor(a1));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_icon(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QIcon *self = (QIcon*)drvGetNative(a0);
    switch (drvid) {
    case _ID_ICON_INIT: {
        drvSetIcon(a0, QIcon());
        break;
    }
    case _ID_ICON_INITWITHFILE: {
        drvSetIcon(a0, QIcon(drvGetString(a1)));
        break;
    }
    case _ID_ICON_INITWITHPIXMAP: {
        drvSetIcon(a0, QIcon(drvGetPixmap(a1)));
        break;
    }
    case _ID_ICON_CLOSE: {
        drvDelIcon(a0,self);
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_image(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QImage *self = (QImage*)drvGetNative(a0);
    switch (drvid) {
    case _ID_IMAGE_INIT: {
        drvSetImage(a0, QImage());
        break;
    }
    case _ID_IMAGE_INITWITHFILE: {
        drvSetImage(a0, QImage(drvGetString(a1)));
        break;
    }
    case _ID_IMAGE_INITWITHSIZE: {
        drvSetImage(a0, QImage(drvGetInt(a1),drvGetInt(a2),QImage::Format_ARGB32_Premultiplied));
        break;
    }
    case _ID_IMAGE_CLOSE: {
        drvDelImage(a0,self);
        break;
    }
    case _ID_IMAGE_DEPTH: {
        drvSetInt(a1,self->depth());
        break;
    }
    case _ID_IMAGE_SIZE: {
        drvSetSize(a1,self->size());
        break;
    }
    case _ID_IMAGE_RECT: {
        drvSetRect(a1,self->rect());
        break;
    }
    case _ID_IMAGE_FILL: {
        self->fill(drvGetUint(a1));
        break;
    }
    case _ID_IMAGE_SCALED: {
        drvSetImage(a5,self->scaled(drvGetInt(a1),drvGetInt(a2),drvGetAspectRatioMode(a3),drvGetTransformationMode(a4)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_widget(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QWidget *self = (QWidget*)drvGetNative(a0);
    switch (drvid) {
    case _ID_WIDGET_INIT: {
        drvNewObj(a0,new QWidget);
        break;
    }
    case _ID_WIDGET_CLOSE: {
        if (self->testAttribute(Qt::WA_DeleteOnClose)) {
        	self->close();
        } else {
        	self->close();
        	drvDelObj(a0,self);
        }
        break;
    }
    case _ID_WIDGET_SETAUTOCLOSE: {
        self->setAttribute(Qt::WA_DeleteOnClose,drvGetBool(a1));
        break;
    }
    case _ID_WIDGET_AUTOCLOSE: {
        drvSetBool(a1,self->testAttribute(Qt::WA_DeleteOnClose));
        break;
    }
    case _ID_WIDGET_SETLAYOUT: {
        self->setLayout(drvGetLayout(a1));
        break;
    }
    case _ID_WIDGET_LAYOUT: {
        drvSetHandle(a1,self->layout());
        break;
    }
    case _ID_WIDGET_SETPARENT: {
        self->setParent(drvGetWidget(a1));
        break;
    }
    case _ID_WIDGET_PARENT: {
        drvSetHandle(a1,self->parent());
        break;
    }
    case _ID_WIDGET_SETVISIBLE: {
        self->setVisible(drvGetBool(a1));
        break;
    }
    case _ID_WIDGET_ISVISIBLE: {
        drvSetBool(a1,self->isVisible());
        break;
    }
    case _ID_WIDGET_SETWINDOWTITLE: {
        self->setWindowTitle(drvGetString(a1));
        break;
    }
    case _ID_WIDGET_WINDOWTITLE: {
        drvSetString(a1,self->windowTitle());
        break;
    }
    case _ID_WIDGET_SETWINDOWICON: {
        self->setWindowIcon(drvGetIcon(a1));
        break;
    }
    case _ID_WIDGET_WINDOWICON: {
        drvSetIcon(a1,self->windowIcon());
        break;
    }
    case _ID_WIDGET_SETPOS: {
        self->move(drvGetPoint(a1));
        break;
    }
    case _ID_WIDGET_POS: {
        drvSetPoint(a1,self->pos());
        break;
    }
    case _ID_WIDGET_SETSIZE: {
        self->resize(drvGetSize(a1));
        break;
    }
    case _ID_WIDGET_SIZE: {
        drvSetSize(a1,self->size());
        break;
    }
    case _ID_WIDGET_SETMINIMUMSIZE: {
        self->setMinimumSize(drvGetSize(a1));
        break;
    }
    case _ID_WIDGET_MINIMUMSIZE: {
        drvSetSize(a1,self->minimumSize());
        break;
    }
    case _ID_WIDGET_SETMAXIMUMSIZE: {
        self->setMaximumSize(drvGetSize(a1));
        break;
    }
    case _ID_WIDGET_MAXIMUMSIZE: {
        drvSetSize(a1,self->maximumSize());
        break;
    }
    case _ID_WIDGET_SETFIXEDSIZE: {
        self->setFixedSize(drvGetSize(a1));
        break;
    }
    case _ID_WIDGET_SETGEOMETRY: {
        self->setGeometry(drvGetRect(a1));
        break;
    }
    case _ID_WIDGET_GEOMETRY: {
        drvSetRect(a1,self->geometry());
        break;
    }
    case _ID_WIDGET_SETFONT: {
        self->setFont(drvGetFont(a1));
        break;
    }
    case _ID_WIDGET_FONT: {
        drvSetFont(a1,self->font());
        break;
    }
    case _ID_WIDGET_SETTOOLTIP: {
        self->setToolTip(drvGetString(a1));
        break;
    }
    case _ID_WIDGET_TOOLTIP: {
        drvSetString(a1,self->toolTip());
        break;
    }
    case _ID_WIDGET_SETUPDATESENABLED: {
        self->setUpdatesEnabled(drvGetBool(a1));
        break;
    }
    case _ID_WIDGET_UPDATESENABLED: {
        drvSetBool(a1,self->updatesEnabled());
        break;
    }
    case _ID_WIDGET_ISACTIVATEWINDOW: {
        drvSetBool(a1,self->isActiveWindow());
        break;
    }
    case _ID_WIDGET_ACTIVATEWINDOW: {
        self->activateWindow();
        break;
    }
    case _ID_WIDGET_DONE: {
        self->close();
        break;
    }
    case _ID_WIDGET_UPDATE: {
        self->update();
        break;
    }
    case _ID_WIDGET_REPAINT: {
        self->repaint();
        break;
    }
    case _ID_WIDGET_STARTTIMER: {
        drvSetInt(a2,self->startTimer(drvGetInt(a1)));
        break;
    }
    case _ID_WIDGET_KILLTIMER: {
        self->killTimer(drvGetInt(a1));
        break;
    }
    case _ID_WIDGET_ADDACTION: {
        self->addAction(drvGetAction(a1));
        break;
    }
    case _ID_WIDGET_INSERTACTION: {
        self->insertAction(drvGetAction(a1),drvGetAction(a2));
        break;
    }
    case _ID_WIDGET_REMOVEACTION: {
        self->removeAction(drvGetAction(a1));
        break;
    }
    case _ID_WIDGET_ONSHOWEVENT: {
        drvNewEvent(QEvent::Show,a0,a1,a2);
        break;
    }
    case _ID_WIDGET_ONHIDEEVENT: {
        drvNewEvent(QEvent::Hide,a0,a1,a2);
        break;
    }
    case _ID_WIDGET_ONCLOSEEVENT: {
        drvNewEvent(QEvent::Close,a0,a1,a2);
        break;
    }
    case _ID_WIDGET_ONKEYPRESSEVENT: {
        drvNewEvent(QEvent::KeyPress,a0,a1,a2);
        break;
    }
    case _ID_WIDGET_ONKEYRELEASEEVENT: {
        drvNewEvent(QEvent::KeyRelease,a0,a1,a2);
        break;
    }
    case _ID_WIDGET_ONMOUSEPRESSEVENT: {
        drvNewEvent(QEvent::MouseButtonPress,a0,a1,a2);
        break;
    }
    case _ID_WIDGET_ONMOUSERELEASEEVENT: {
        drvNewEvent(QEvent::MouseButtonRelease,a0,a1,a2);
        break;
    }
    case _ID_WIDGET_ONMOUSEMOVEEVENT: {
        drvNewEvent(QEvent::MouseMove,a0,a1,a2);
        break;
    }
    case _ID_WIDGET_ONMOUSEDOUBLECLICKEVENT: {
        drvNewEvent(QEvent::MouseButtonDblClick,a0,a1,a2);
        break;
    }
    case _ID_WIDGET_ONMOVEEVENT: {
        drvNewEvent(QEvent::Move,a0,a1,a2);
        break;
    }
    case _ID_WIDGET_ONPAINTEVENT: {
        drvNewEvent(QEvent::Paint,a0,a1,a2);
        break;
    }
    case _ID_WIDGET_ONRESIZEEVENT: {
        drvNewEvent(QEvent::Resize,a0,a1,a2);
        break;
    }
    case _ID_WIDGET_ONENTEREVENT: {
        drvNewEvent(QEvent::Enter,a0,a1,a2);
        break;
    }
    case _ID_WIDGET_ONLEAVEEVENT: {
        drvNewEvent(QEvent::Leave,a0,a1,a2);
        break;
    }
    case _ID_WIDGET_ONFOCUSINEVENT: {
        drvNewEvent(QEvent::FocusIn,a0,a1,a2);
        break;
    }
    case _ID_WIDGET_ONFOCUSOUTEVENT: {
        drvNewEvent(QEvent::FocusOut,a0,a1,a2);
        break;
    }
    case _ID_WIDGET_ONTIMEREVENT: {
        drvNewEvent(QEvent::Timer,a0,a1,a2);
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_action(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QAction *self = (QAction*)drvGetNative(a0);
    switch (drvid) {
    case _ID_ACTION_INIT: {
        drvNewObj(a0, new QAction(qApp));
        break;
    }
    case _ID_ACTION_INITWITHTEXT: {
        drvNewObj(a0, new QAction(drvGetString(a1),qApp));
        break;
    }
    case _ID_ACTION_CLOSE: {
        drvDelObj(a0,self);
        break;
    }
    case _ID_ACTION_SETTEXT: {
        self->setText(drvGetString(a1));
        break;
    }
    case _ID_ACTION_TEXT: {
        drvSetString(a1,self->text());
        break;
    }
    case _ID_ACTION_SETICON: {
        self->setIcon(drvGetIcon(a1));
        break;
    }
    case _ID_ACTION_ICON: {
        drvSetIcon(a1,self->icon());
        break;
    }
    case _ID_ACTION_SETICONTEXT: {
        self->setIconText(drvGetString(a1));
        break;
    }
    case _ID_ACTION_ICONTEXT: {
        drvSetString(a1,self->iconText());
        break;
    }
    case _ID_ACTION_SETTOOLTIP: {
        self->setToolTip(drvGetString(a1));
        break;
    }
    case _ID_ACTION_TOOLTIP: {
        drvSetString(a1,self->toolTip());
        break;
    }
    case _ID_ACTION_SETFONT: {
        self->setFont(drvGetFont(a1));
        break;
    }
    case _ID_ACTION_FONT: {
        drvSetFont(a1,self->font());
        break;
    }
    case _ID_ACTION_SETCHECKABLE: {
        self->setCheckable(drvGetBool(a1));
        break;
    }
    case _ID_ACTION_ISCHECKABLE: {
        drvSetBool(a1,self->isCheckable());
        break;
    }
    case _ID_ACTION_SETCHECKED: {
        self->setChecked(drvGetBool(a1));
        break;
    }
    case _ID_ACTION_ISCHECKED: {
        drvSetBool(a1,self->isChecked());
        break;
    }
    case _ID_ACTION_SETVISIBLE: {
        self->setVisible(drvGetBool(a1));
        break;
    }
    case _ID_ACTION_ISVISIBLE: {
        drvSetBool(a1,self->isVisible());
        break;
    }
    case _ID_ACTION_ONTRIGGERED: {
        QObject::connect(self,SIGNAL(triggered(bool)),drvNewSignal(self,a1,a2),SLOT(call(bool)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_actiongroup(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QActionGroup *self = (QActionGroup*)drvGetNative(a0);
    switch (drvid) {
    case _ID_ACTIONGROUP_INIT: {
        drvNewObj(a0, new QActionGroup(qApp));
        break;
    }
    case _ID_ACTIONGROUP_CLOSE: {
        drvDelObj(a0,self);
        break;
    }
    case _ID_ACTIONGROUP_SETVISIBLE: {
        self->setVisible(drvGetBool(a1));
        break;
    }
    case _ID_ACTIONGROUP_ISVISIBLE: {
        drvSetBool(a1,self->isVisible());
        break;
    }
    case _ID_ACTIONGROUP_SETENABLED: {
        self->setEnabled(drvGetBool(a1));
        break;
    }
    case _ID_ACTIONGROUP_ISENABLED: {
        drvSetBool(a1,self->isEnabled());
        break;
    }
    case _ID_ACTIONGROUP_SETEXCLUSIVE: {
        self->setExclusive(drvGetBool(a1));
        break;
    }
    case _ID_ACTIONGROUP_ISEXCLUSIVE: {
        drvSetBool(a1,self->isExclusive());
        break;
    }
    case _ID_ACTIONGROUP_ADDACTION: {
        self->addAction(drvGetAction(a1));
        break;
    }
    case _ID_ACTIONGROUP_REMOVEACTION: {
        self->removeAction(drvGetAction(a1));
        break;
    }
    case _ID_ACTIONGROUP_ONHOVERED: {
        QObject::connect(self,SIGNAL(hovered(QAction*)),drvNewSignal(self,a1,a2),SLOT(call(QAction*)));
        break;
    }
    case _ID_ACTIONGROUP_ONTRIGGERED: {
        QObject::connect(self,SIGNAL(triggered(QAction*)),drvNewSignal(self,a1,a2),SLOT(call(QAction*)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_menu(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QMenu *self = (QMenu*)drvGetNative(a0);
    switch (drvid) {
    case _ID_MENU_INIT: {
        drvNewObj(a0,new QMenu);
        break;
    }
    case _ID_MENU_INITWITHTITLE: {
        drvNewObj(a0, new QMenu(drvGetString(a1)));
        break;
    }
    case _ID_MENU_SETTITLE: {
        self->setTitle(drvGetString(a1));
        break;
    }
    case _ID_MENU_TITLE: {
        drvSetString(a1,self->title());
        break;
    }
    case _ID_MENU_SETICON: {
        self->setIcon(drvGetIcon(a1));
        break;
    }
    case _ID_MENU_ICON: {
        drvSetIcon(a1,self->icon());
        break;
    }
    case _ID_MENU_SETDEFAULTACTION: {
        self->setDefaultAction(drvGetAction(a1));
        break;
    }
    case _ID_MENU_DEFAULTACTION: {
        drvSetAction(a1,self->defaultAction());
        break;
    }
    case _ID_MENU_SETACTIVEACTION: {
        self->setActiveAction(drvGetAction(a1));
        break;
    }
    case _ID_MENU_ACTIVEACTION: {
        drvSetAction(a1,self->activeAction());
        break;
    }
    case _ID_MENU_MENUACTION: {
        drvSetAction(a1,self->menuAction());
        break;
    }
    case _ID_MENU_ISEMPTY: {
        drvSetBool(a1,self->isEmpty());
        break;
    }
    case _ID_MENU_ADDMENU: {
        drvSetAction(a2,self->addMenu(drvGetMenu(a1)));
        break;
    }
    case _ID_MENU_INSERTMENU: {
        self->insertMenu(drvGetAction(a1),drvGetMenu(a2));
        break;
    }
    case _ID_MENU_ADDSEPARATOR: {
        drvSetAction(a1,self->addSeparator());
        break;
    }
    case _ID_MENU_INSERTSEPARATOR: {
        drvSetAction(a2,self->insertSeparator(drvGetAction(a1)));
        break;
    }
    case _ID_MENU_CLEAR: {
        self->clear();
        break;
    }
    case _ID_MENU_POPUP: {
        self->popup(drvGetPoint(a1),drvGetAction(a2));
        break;
    }
    case _ID_MENU_ONHOVERED: {
        QObject::connect(self,SIGNAL(hovered(QAction*)),drvNewSignal(self,a1,a2),SLOT(call(QAction*)));
        break;
    }
    case _ID_MENU_ONTRIGGERED: {
        QObject::connect(self,SIGNAL(triggered(QAction*)),drvNewSignal(self,a1,a2),SLOT(call(QAction*)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_menubar(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QMenuBar *self = (QMenuBar*)drvGetNative(a0);
    switch (drvid) {
    case _ID_MENUBAR_INIT: {
        drvNewObj(a0,new QMenuBar);
        break;
    }
    case _ID_MENUBAR_SETACTIVEACTION: {
        self->setActiveAction(drvGetAction(a1));
        break;
    }
    case _ID_MENUBAR_ACTIVEACTION: {
        drvSetAction(a1,self->activeAction());
        break;
    }
    case _ID_MENUBAR_SETNATIVEMENUBAR: {
        self->setNativeMenuBar(drvGetBool(a1));
        break;
    }
    case _ID_MENUBAR_ISNATIVEMENUBAR: {
        drvSetBool(a1,self->isNativeMenuBar());
        break;
    }
    case _ID_MENUBAR_ADDMENU: {
        drvSetAction(a2,self->addMenu(drvGetMenu(a1)));
        break;
    }
    case _ID_MENUBAR_INSERTMENU: {
        self->insertMenu(drvGetAction(a1),drvGetMenu(a2));
        break;
    }
    case _ID_MENUBAR_ADDSEPARATOR: {
        drvSetAction(a1,self->addSeparator());
        break;
    }
    case _ID_MENUBAR_INSERTSEPARATOR: {
        drvSetAction(a2,self->insertSeparator(drvGetAction(a1)));
        break;
    }
    case _ID_MENUBAR_CLEAR: {
        self->clear();
        break;
    }
    case _ID_MENUBAR_ONHOVERED: {
        QObject::connect(self,SIGNAL(hovered(QAction*)),drvNewSignal(self,a1,a2),SLOT(call(QAction*)));
        break;
    }
    case _ID_MENUBAR_ONTRIGGERED: {
        QObject::connect(self,SIGNAL(triggered(QAction*)),drvNewSignal(self,a1,a2),SLOT(call(QAction*)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_toolbar(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QToolBar *self = (QToolBar*)drvGetNative(a0);
    switch (drvid) {
    case _ID_TOOLBAR_INIT: {
        drvNewObj(a0, new QToolBar,true,false);
        break;
    }
    case _ID_TOOLBAR_INITWITHTITLE: {
        drvNewObj(a0, new QToolBar(drvGetString(a1)),true,false);
        break;
    }
    case _ID_TOOLBAR_SETTITLE: {
        self->setWindowTitle(drvGetString(a1));
        break;
    }
    case _ID_TOOLBAR_TITLE: {
        drvSetString(a1,self->windowTitle());
        break;
    }
    case _ID_TOOLBAR_SETICONSIZE: {
        self->setIconSize(drvGetSize(a1));
        break;
    }
    case _ID_TOOLBAR_ICONSIZE: {
        drvSetSize(a1,self->iconSize());
        break;
    }
    case _ID_TOOLBAR_SETFLOATABLE: {
        self->setFloatable(drvGetBool(a1));
        break;
    }
    case _ID_TOOLBAR_ISFLOATABLE: {
        drvSetBool(a1,self->isFloatable());
        break;
    }
    case _ID_TOOLBAR_ISFLOATING: {
        drvSetBool(a1,self->isFloating());
        break;
    }
    case _ID_TOOLBAR_SETTOOLBUTTONSTYLE: {
        self->setToolButtonStyle(drvGetToolButtonStyle(a1));
        break;
    }
    case _ID_TOOLBAR_TOOLBUTTONSTYLE: {
        drvSetToolButtonStyle(a1,self->toolButtonStyle());
        break;
    }
    case _ID_TOOLBAR_ADDACTION: {
        self->addAction(drvGetAction(a1));
        break;
    }
    case _ID_TOOLBAR_ADDSEPARATOR: {
        drvNewObj(a1,self->addSeparator());
        break;
    }
    case _ID_TOOLBAR_ADDWIDGET: {
        drvNewObj(a2,self->addWidget(drvGetWidget(a1)));
        break;
    }
    case _ID_TOOLBAR_CLEAR: {
        self->clear();
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_statusbar(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QStatusBar *self = (QStatusBar*)drvGetNative(a0);
    switch (drvid) {
    case _ID_STATUSBAR_INIT: {
        drvNewObj(a0, new QStatusBar());
        break;
    }
    case _ID_STATUSBAR_SETSIZEGRIP: {
        self->setSizeGripEnabled(drvGetBool(a1));
        break;
    }
    case _ID_STATUSBAR_ISSIZEGRIP: {
        drvSetBool(a1,self->isSizeGripEnabled());
        break;
    }
    case _ID_STATUSBAR_ADDWIDGET: {
        self->addWidget(drvGetWidget(a1),drvGetInt(a2));
        break;
    }
    case _ID_STATUSBAR_ADDPERMANENTWIDGET: {
        self->addPermanentWidget(drvGetWidget(a1),drvGetInt(a2));
        break;
    }
    case _ID_STATUSBAR_INSERTWIDGET: {
        self->insertWidget(drvGetInt(a1),drvGetWidget(a2),drvGetInt(a3));
        break;
    }
    case _ID_STATUSBAR_INSERTPERMANENTWIDGET: {
        self->insertPermanentWidget(drvGetInt(a1),drvGetWidget(a2),drvGetInt(a3));
        break;
    }
    case _ID_STATUSBAR_REMOVEWIDGET: {
        self->removeWidget(drvGetWidget(a1));
        break;
    }
    case _ID_STATUSBAR_SHOWMESSAGE: {
        self->showMessage(drvGetString(a1),drvGetInt(a2));
        break;
    }
    case _ID_STATUSBAR_CURRENTMESSAGE: {
        drvSetString(a1,self->currentMessage());
        break;
    }
    case _ID_STATUSBAR_CLEARMESSAGE: {
        self->clearMessage();
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_dockwidget(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QDockWidget *self = (QDockWidget*)drvGetNative(a0);
    switch (drvid) {
    case _ID_DOCKWIDGET_INIT: {
        drvNewObj(a0,new QDockWidget,true,false);
        break;
    }
    case _ID_DOCKWIDGET_INITWITHTITLE: {
        drvNewObj(a0,new QDockWidget(drvGetString(a1)),true,false);
        break;
    }
    case _ID_DOCKWIDGET_SETDOCK: {
        self->setWidget(drvGetWidget(a1));
        break;
    }
    case _ID_DOCKWIDGET_DOCK: {
        drvSetHandle(a1,self->widget());
        break;
    }
    case _ID_DOCKWIDGET_SETTITLEBAR: {
        self->setTitleBarWidget(drvGetWidget(a1));
        break;
    }
    case _ID_DOCKWIDGET_TITLEBAR: {
        drvSetHandle(a1,self->titleBarWidget());
        break;
    }
    case _ID_DOCKWIDGET_SETFLOATING: {
        self->setFloating(drvGetBool(a1));
        break;
    }
    case _ID_DOCKWIDGET_ISFLOATING: {
        drvSetBool(a1,self->isFloating());
        break;
    }
    case _ID_DOCKWIDGET_ONVISIBILITYCHANGED: {
        QObject::connect(self,SIGNAL(visibilityChanged(bool)),drvNewSignal(self,a1,a2),SLOT(call(bool)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_systemtray(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QSystemTrayIcon *self = (QSystemTrayIcon*)drvGetNative(a0);
    switch (drvid) {
    case _ID_SYSTEMTRAY_INIT: {
        drvNewObj(a0, new QSystemTrayIcon);
        break;
    }
    case _ID_SYSTEMTRAY_CLOSE: {
        drvDelObj(a0,self);
        break;
    }
    case _ID_SYSTEMTRAY_SETCONTEXTMENU: {
        self->setContextMenu(drvGetMenu(a1));
        break;
    }
    case _ID_SYSTEMTRAY_CONTEXTMENU: {
        drvSetMenu(a1,self->contextMenu());
        break;
    }
    case _ID_SYSTEMTRAY_SETICON: {
        self->setIcon(drvGetIcon(a1));
        break;
    }
    case _ID_SYSTEMTRAY_ICON: {
        drvSetIcon(a1,self->icon());
        break;
    }
    case _ID_SYSTEMTRAY_SETTOOLTIP: {
        self->setToolTip(drvGetString(a1));
        break;
    }
    case _ID_SYSTEMTRAY_TOOLTIP: {
        drvSetString(a1,self->toolTip());
        break;
    }
    case _ID_SYSTEMTRAY_SETVISIBLE: {
        self->setVisible(drvGetBool(a1));
        break;
    }
    case _ID_SYSTEMTRAY_ISVISIBLE: {
        drvSetBool(a1,self->isVisible());
        break;
    }
    case _ID_SYSTEMTRAY_SHOWMESSAGE: {
        self->showMessage(drvGetString(a1),drvGetString(a2),drvGetMessageIconType(a3),drvGetInt(a4));
        break;
    }
    case _ID_SYSTEMTRAY_ONACTIVATED: {
        QObject::connect(self,SIGNAL(activated(int)),drvNewSignal(self,a1,a2),SLOT(call(int)));
        break;
    }
    case _ID_SYSTEMTRAY_ONMESSAGECLICKED: {
        QObject::connect(self,SIGNAL(messageClicked(int)),drvNewSignal(self,a1,a2),SLOT(call()));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_tabwidget(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QTabWidget *self = (QTabWidget*)drvGetNative(a0);
    switch (drvid) {
    case _ID_TABWIDGET_INIT: {
        drvNewObj(a0,new QTabWidget);
        break;
    }
    case _ID_TABWIDGET_COUNT: {
        drvSetInt(a1,self->count());
        break;
    }
    case _ID_TABWIDGET_CURRENTINDEX: {
        drvSetInt(a1,self->currentIndex());
        break;
    }
    case _ID_TABWIDGET_CURRENTWIDGET: {
        drvSetHandle(a1,self->currentWidget());
        break;
    }
    case _ID_TABWIDGET_SETCURRENTINDEX: {
        self->setCurrentIndex(drvGetInt(a1));
        break;
    }
    case _ID_TABWIDGET_SETCURRENTWIDGET: {
        self->setCurrentWidget(drvGetWidget(a1));
        break;
    }
    case _ID_TABWIDGET_ADDTAB: {
        self->addTab(drvGetWidget(a1),drvGetIcon(a3),drvGetString(a2));
        break;
    }
    case _ID_TABWIDGET_INSERTTAB: {
        self->insertTab(drvGetInt(a1),drvGetWidget(a2),drvGetIcon(a4),drvGetString(a3));
        break;
    }
    case _ID_TABWIDGET_REMOVETAB: {
        self->removeTab(drvGetInt(a1));
        break;
    }
    case _ID_TABWIDGET_CLEAR: {
        self->clear();
        break;
    }
    case _ID_TABWIDGET_INDEXOF: {
        drvSetInt(a2,self->indexOf(drvGetWidget(a1)));
        break;
    }
    case _ID_TABWIDGET_WIDGETOF: {
        drvSetHandle(a2,self->widget(drvGetInt(a1)));
        break;
    }
    case _ID_TABWIDGET_SETTABTEXT: {
        self->setTabText(drvGetInt(a1),drvGetString(a2));
        break;
    }
    case _ID_TABWIDGET_TABTEXT: {
        drvSetString(a2,self->tabText(drvGetInt(a1)));
        break;
    }
    case _ID_TABWIDGET_SETTABICON: {
        self->setTabIcon(drvGetInt(a1),drvGetIcon(a2));
        break;
    }
    case _ID_TABWIDGET_TABICON: {
        drvSetIcon(a2,self->tabIcon(drvGetInt(a1)));
        break;
    }
    case _ID_TABWIDGET_SETTABTOOLTIP: {
        self->setTabToolTip(drvGetInt(a1),drvGetString(a2));
        break;
    }
    case _ID_TABWIDGET_TABTOOLTIP: {
        drvSetString(a2,self->tabToolTip(drvGetInt(a1)));
        break;
    }
    case _ID_TABWIDGET_ONCURRENTCHANGED: {
        QObject::connect(self,SIGNAL(currentChanged(int)),drvNewSignal(self,a1,a2),SLOT(call(int)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_toolbox(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QToolBox *self = (QToolBox*)drvGetNative(a0);
    switch (drvid) {
    case _ID_TOOLBOX_INIT: {
        drvNewObj(a0, new QToolBox);
        break;
    }
    case _ID_TOOLBOX_SETCURRENTINDEX: {
        self->setCurrentIndex(drvGetInt(a1));
        break;
    }
    case _ID_TOOLBOX_SETCURRENTWIDGET: {
        self->setCurrentWidget(drvGetWidget(a1));
        break;
    }
    case _ID_TOOLBOX_CURRENTINDEX: {
        drvSetInt(a1,self->currentIndex());
        break;
    }
    case _ID_TOOLBOX_CURRENTWIDGET: {
        drvSetHandle(a1,self->currentWidget());
        break;
    }
    case _ID_TOOLBOX_COUNT: {
        drvSetInt(a1,self->count());
        break;
    }
    case _ID_TOOLBOX_ADDITEM: {
        self->addItem(drvGetWidget(a1),drvGetIcon(a3),drvGetString(a2));
        break;
    }
    case _ID_TOOLBOX_INSERTITEM: {
        self->insertItem(drvGetInt(a1),drvGetWidget(a2),drvGetIcon(a4),drvGetString(a3));
        break;
    }
    case _ID_TOOLBOX_REMOVEITEM: {
        self->removeItem(drvGetInt(a1));
        break;
    }
    case _ID_TOOLBOX_INDEXOF: {
        drvSetInt(a2,self->indexOf(drvGetWidget(a1)));
        break;
    }
    case _ID_TOOLBOX_WIDGETOF: {
        drvSetHandle(a2,self->widget(drvGetInt(a1)));
        break;
    }
    case _ID_TOOLBOX_SETITEMTEXT: {
        self->setItemText(drvGetInt(a1),drvGetString(a2));
        break;
    }
    case _ID_TOOLBOX_ITEMTEXT: {
        drvSetString(a2,self->itemText(drvGetInt(a1)));
        break;
    }
    case _ID_TOOLBOX_SETITEMICON: {
        self->setItemIcon(drvGetInt(a1),drvGetIcon(a2));
        break;
    }
    case _ID_TOOLBOX_ITEMICON: {
        drvSetIcon(a2,self->itemIcon(drvGetInt(a1)));
        break;
    }
    case _ID_TOOLBOX_SETITEMTOOLTIP: {
        self->setItemToolTip(drvGetInt(a1),drvGetString(a2));
        break;
    }
    case _ID_TOOLBOX_ITEMTOOLTIP: {
        drvSetString(a2,self->itemToolTip(drvGetInt(a1)));
        break;
    }
    case _ID_TOOLBOX_ISITEMENABLED: {
        drvSetBool(a2,self->isItemEnabled(drvGetInt(a1)));
        break;
    }
    case _ID_TOOLBOX_ONCURRENTCHANGED: {
        QObject::connect(self,SIGNAL(currentChanged(int)),drvNewSignal(self,a1,a2),SLOT(call(int)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_baselayout(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QLayout *self = (QLayout*)drvGetNative(a0);
    switch (drvid) {
    case _ID_BASELAYOUT_SETSPACING: {
        self->setSpacing(drvGetInt(a1));
        break;
    }
    case _ID_BASELAYOUT_SPACING: {
        drvSetInt(a1,self->spacing());
        break;
    }
    case _ID_BASELAYOUT_SETMARGINS: {
        self->setContentsMargins(drvGetMargins(a1));
        break;
    }
    case _ID_BASELAYOUT_MARGINS: {
        drvSetMargins(a1,self->contentsMargins());
        break;
    }
    case _ID_BASELAYOUT_SETMENUBAR: {
        self->setMenuBar(drvGetWidget(a1));
        break;
    }
    case _ID_BASELAYOUT_MENUBAR: {
        drvSetHandle(a1,self->menuBar());
        break;
    }
    case _ID_BASELAYOUT_COUNT: {
        drvSetInt(a1,self->count());
        break;
    }
    case _ID_BASELAYOUT_ADDLAYOUT: {
        self->addItem(drvGetLayout(a1));
        break;
    }
    case _ID_BASELAYOUT_ADDWIDGET: {
        self->addWidget(drvGetWidget(a1));
        break;
    }
    case _ID_BASELAYOUT_REMOVEWIDGET: {
        self->removeWidget(drvGetWidget(a1));
        break;
    }
    case _ID_BASELAYOUT_INDEXOF: {
        drvSetInt(a2,self->indexOf(drvGetWidget(a1)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_boxlayout(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QBoxLayout *self = (QBoxLayout*)drvGetNative(a0);
    switch (drvid) {
    case _ID_BOXLAYOUT_INIT: {
        drvNewObj(a0,new QBoxLayout(QBoxLayout::TopToBottom));
        break;
    }
    case _ID_BOXLAYOUT_CLOSE: {
        drvDelObj(a0,self);
        break;
    }
    case _ID_BOXLAYOUT_SETORIENTATION: {
        switch (drvGetInt(a1)) {
        	case Qt::Horizontal: {
        		self->setDirection(QBoxLayout::TopToBottom);
        		break;		
        	}
        	case Qt::Vertical: {
        		self->setDirection(QBoxLayout::LeftToRight);
        		break;
        	}
        }
        break;
    }
    case _ID_BOXLAYOUT_ORIENTATION: {
        switch((int)self->direction()) {
        	case QBoxLayout::TopToBottom: {
        		drvSetInt(a1,Qt::Horizontal);
        		break;
        	}
        	case QBoxLayout::LeftToRight: {
        		drvSetInt(a1,Qt::Vertical);
        		break;
        	}
        }
        break;
    }
    case _ID_BOXLAYOUT_ADDLAYOUTWITH: {
        self->addLayout((QLayout*)drvGetNative(a1),drvGetInt(a2));
        break;
    }
    case _ID_BOXLAYOUT_ADDWIDGETWITH: {
        self->addWidget((QWidget*)drvGetNative(a1),drvGetInt(a2),Qt::Alignment(drvGetInt(a3)));
        break;
    }
    case _ID_BOXLAYOUT_ADDSPACING: {
        self->addSpacing(drvGetInt(a1));
        break;
    }
    case _ID_BOXLAYOUT_ADDSTRETCH: {
        self->addStretch(drvGetInt(a1));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_hboxlayout(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QHBoxLayout *self = (QHBoxLayout*)drvGetNative(a0);
    switch (drvid) {
    case _ID_HBOXLAYOUT_INIT: {
        drvNewObj(a0,new QHBoxLayout);
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_vboxlayout(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QVBoxLayout *self = (QVBoxLayout*)drvGetNative(a0);
    switch (drvid) {
    case _ID_VBOXLAYOUT_INIT: {
        drvNewObj(a0,new QVBoxLayout);
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_stackedlayout(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QStackedLayout *self = (QStackedLayout*)drvGetNative(a0);
    switch (drvid) {
    case _ID_STACKEDLAYOUT_INIT: {
        drvNewObj(a0,new QStackedLayout);
        break;
    }
    case _ID_STACKEDLAYOUT_SETCURRENTINDEX: {
        self->setCurrentIndex(drvGetInt(a1));
        break;
    }
    case _ID_STACKEDLAYOUT_CURRENTINDEX: {
        drvSetInt(a1,self->currentIndex());
        break;
    }
    case _ID_STACKEDLAYOUT_SETCURRENTWIDGET: {
        self->setCurrentWidget(drvGetWidget(a1));
        break;
    }
    case _ID_STACKEDLAYOUT_CURRENTWIDGET: {
        drvSetHandle(a1,self->currentWidget());
        break;
    }
    case _ID_STACKEDLAYOUT_ADDWIDGET: {
        self->addWidget(drvGetWidget(a1));
        break;
    }
    case _ID_STACKEDLAYOUT_INSERTWIDGET: {
        self->insertWidget(drvGetInt(a1),drvGetWidget(a2));
        break;
    }
    case _ID_STACKEDLAYOUT_WIDGET: {
        drvSetHandle(a2,self->widget(drvGetInt(a1)));
        break;
    }
    case _ID_STACKEDLAYOUT_ONCURRENTCHANGED: {
        QObject::connect(self,SIGNAL(currentChanged(int)),drvNewSignal(self,a1,a2),SLOT(call(int)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_basebutton(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QAbstractButton *self = (QAbstractButton*)drvGetNative(a0);
    switch (drvid) {
    case _ID_BASEBUTTON_SETTEXT: {
        self->setText(drvGetString(a1));
        break;
    }
    case _ID_BASEBUTTON_TEXT: {
        drvSetString(a1,self->text());
        break;
    }
    case _ID_BASEBUTTON_SETICON: {
        self->setIcon(drvGetIcon(a1));
        break;
    }
    case _ID_BASEBUTTON_ICON: {
        drvSetIcon(a1,self->icon());
        break;
    }
    case _ID_BASEBUTTON_SETICONSIZE: {
        self->setIconSize(drvGetSize(a1));
        break;
    }
    case _ID_BASEBUTTON_ICONSIZE: {
        drvSetSize(a1,self->iconSize());
        break;
    }
    case _ID_BASEBUTTON_SETCHECKABLE: {
        self->setCheckable(drvGetBool(a1));
        break;
    }
    case _ID_BASEBUTTON_ISCHECKABLE: {
        drvSetBool(a1,self->isCheckable());
        break;
    }
    case _ID_BASEBUTTON_SETDOWN: {
        self->setDown(drvGetBool(a1));
        break;
    }
    case _ID_BASEBUTTON_ISDOWN: {
        drvSetBool(a1,self->isDown());
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_button(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QPushButton *self = (QPushButton*)drvGetNative(a0);
    switch (drvid) {
    case _ID_BUTTON_INIT: {
        drvNewObj(a0,new QPushButton);
        break;
    }
    case _ID_BUTTON_INITWITHTEXT: {
        drvNewObj(a0,new QPushButton(drvGetString(a1)));
        break;
    }
    case _ID_BUTTON_SETFLAT: {
        self->setFlat(drvGetBool(a1));
        break;
    }
    case _ID_BUTTON_ISFLAT: {
        drvSetBool(a1,self->isFlat());
        break;
    }
    case _ID_BUTTON_SETDEFAULT: {
        self->setDefault(drvGetBool(a1));
        break;
    }
    case _ID_BUTTON_ISDEFAULT: {
        drvSetBool(a1,self->isDefault());
        break;
    }
    case _ID_BUTTON_SETMENU: {
        self->setMenu(drvGetMenu(a1));
        break;
    }
    case _ID_BUTTON_MENU: {
        drvSetMenu(a1,self->menu());
        break;
    }
    case _ID_BUTTON_ONCLICKED: {
        QObject::connect(self,SIGNAL(clicked()),drvNewSignal(self,a1,a2),SLOT(call()));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_checkbox(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QCheckBox *self = (QCheckBox*)drvGetNative(a0);
    switch (drvid) {
    case _ID_CHECKBOX_INIT: {
        drvNewObj(a0,new QCheckBox);
        break;
    }
    case _ID_CHECKBOX_INITWITHTEXT: {
        drvNewObj(a0,new QCheckBox(drvGetString(a1)));
        break;
    }
    case _ID_CHECKBOX_SETCHECK: {
        self->setCheckState((Qt::CheckState)drvGetInt(a1));
        break;
    }
    case _ID_CHECKBOX_CHECK: {
        drvSetInt(a1,self->checkState());
        break;
    }
    case _ID_CHECKBOX_SETTRISTATE: {
        self->setTristate(drvGetBool(a1));
        break;
    }
    case _ID_CHECKBOX_ISTRISTATE: {
        drvSetBool(a1,self->isTristate());
        break;
    }
    case _ID_CHECKBOX_ONSTATECHANGED: {
        QObject::connect(self,SIGNAL(stateChanged(int)),drvNewSignal(self,a1,a2),SLOT(call(int)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_radio(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QRadioButton *self = (QRadioButton*)drvGetNative(a0);
    switch (drvid) {
    case _ID_RADIO_INIT: {
        drvNewObj(a0,new QRadioButton);
        break;
    }
    case _ID_RADIO_INITWITHTEXT: {
        drvNewObj(a0, new QRadioButton(drvGetString(a1)));
        break;
    }
    case _ID_RADIO_ONCLICKED: {
        QObject::connect(self,SIGNAL(clicked()),drvNewSignal(self,a1,a2),SLOT(call()));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_toolbutton(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QToolButton *self = (QToolButton*)drvGetNative(a0);
    switch (drvid) {
    case _ID_TOOLBUTTON_INIT: {
        drvNewObj(a0,new QToolButton);
        break;
    }
    case _ID_TOOLBUTTON_INITWITHTEXT: {
        drvNewObj(a0,new QPushButton(drvGetString(a1)));
        break;
    }
    case _ID_TOOLBUTTON_SETMENU: {
        self->setMenu(drvGetMenu(a1));
        break;
    }
    case _ID_TOOLBUTTON_MENU: {
        drvSetMenu(a1,self->menu());
        break;
    }
    case _ID_TOOLBUTTON_SETAUTORAISE: {
        self->setAutoRaise(drvGetBool(a1));
        break;
    }
    case _ID_TOOLBUTTON_AUTORAISE: {
        drvSetBool(a1,self->autoRaise());
        break;
    }
    case _ID_TOOLBUTTON_SETTOOLBUTTONSTYLE: {
        self->setToolButtonStyle(drvGetToolButtonStyle(a1));
        break;
    }
    case _ID_TOOLBUTTON_TOOLBUTTONSTYLE: {
        drvSetToolButtonStyle(a1,self->toolButtonStyle());
        break;
    }
    case _ID_TOOLBUTTON_SETTOOLBUTTONPOPUPMODE: {
        self->setPopupMode(drvGetToolButtonPopupMode(a1));
        break;
    }
    case _ID_TOOLBUTTON_TOOLBUTTONPOPUPMODE: {
        drvSetToolButtonPopupMode(a1,self->popupMode());
        break;
    }
    case _ID_TOOLBUTTON_ONCLICKED: {
        QObject::connect(self,SIGNAL(clicked()),drvNewSignal(self,a1,a2),SLOT(call()));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_frame(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QFrame *self = (QFrame*)drvGetNative(a0);
    switch (drvid) {
    case _ID_FRAME_INIT: {
        drvNewObj(a0,new QFrame);
        break;
    }
    case _ID_FRAME_SETFRAMESTYLE: {
        self->setFrameStyle(drvGetInt(a1));
        break;
    }
    case _ID_FRAME_FRAMESTYLE: {
        drvSetInt(a1,self->frameStyle());
        break;
    }
    case _ID_FRAME_SETFRAMERECT: {
        self->setFrameRect(drvGetRect(a1));
        break;
    }
    case _ID_FRAME_FRAMERECT: {
        drvSetRect(a1,self->frameRect());
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_label(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QLabel *self = (QLabel*)drvGetNative(a0);
    switch (drvid) {
    case _ID_LABEL_INIT: {
        drvNewObj(a0,new QLabel);
        break;
    }
    case _ID_LABEL_INITWITHTEXT: {
        drvNewObj(a0, new QLabel(drvGetString(a1)));
        break;
    }
    case _ID_LABEL_INITWITHPIXMAP: {
        QLabel *label = new QLabel;
        label->setPixmap(drvGetPixmap(a1));
        drvNewObj(a0,label);
        break;
    }
    case _ID_LABEL_SETTEXT: {
        self->setText(drvGetString(a1));
        break;
    }
    case _ID_LABEL_TEXT: {
        drvSetString(a1,self->text());
        break;
    }
    case _ID_LABEL_SETWORDWRAP: {
        self->setWordWrap(drvGetBool(a1));
        break;
    }
    case _ID_LABEL_WORDWRAP: {
        drvSetBool(a1,self->wordWrap());
        break;
    }
    case _ID_LABEL_SETTEXTFORMAT: {
        self->setTextFormat((Qt::TextFormat)drvGetInt(a1));
        break;
    }
    case _ID_LABEL_TEXTFORMAT: {
        drvSetInt(a1,self->textFormat());
        break;
    }
    case _ID_LABEL_SETPIXMAP: {
        self->setPixmap(drvGetPixmap(a1));
        break;
    }
    case _ID_LABEL_PIXMAP: {
        drvSetPixmap(a1,self->pixmap());
        break;
    }
    case _ID_LABEL_ONLINKACTIVATED: {
        QObject::connect(self,SIGNAL(linkActivated(QString)),drvNewSignal(self,a1,a2),SLOT(call(QString)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_groupbox(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QGroupBox *self = (QGroupBox*)drvGetNative(a0);
    switch (drvid) {
    case _ID_GROUPBOX_INIT: {
        drvNewObj(a0,new QGroupBox);
        break;
    }
    case _ID_GROUPBOX_INITWITHTITLE: {
        drvNewObj(a0, new QGroupBox(drvGetString(a1)));
        break;
    }
    case _ID_GROUPBOX_SETTITLE: {
        self->setTitle(drvGetString(a1));
        break;
    }
    case _ID_GROUPBOX_TITLE: {
        drvSetString(a1,self->title());
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_dialog(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QDialog *self = (QDialog*)drvGetNative(a0);
    switch (drvid) {
    case _ID_DIALOG_INIT: {
        drvNewObj(a0,new QDialog);
        break;
    }
    case _ID_DIALOG_SETMODAL: {
        self->setModal(drvGetBool(a1));
        break;
    }
    case _ID_DIALOG_ISMODAL: {
        drvSetBool(a1,self->isModal());
        break;
    }
    case _ID_DIALOG_SETRESULT: {
        self->setResult(drvGetInt(a1));
        break;
    }
    case _ID_DIALOG_RESULT: {
        drvSetInt(a1,self->result());
        break;
    }
    case _ID_DIALOG_OPEN: {
        self->open();
        break;
    }
    case _ID_DIALOG_EXEC: {
        drvSetInt(a1,self->exec());
        break;
    }
    case _ID_DIALOG_DONE: {
        self->done(drvGetInt(a1));
        break;
    }
    case _ID_DIALOG_ACCEPT: {
        self->accept();
        break;
    }
    case _ID_DIALOG_REJECT: {
        self->reject();
        break;
    }
    case _ID_DIALOG_ONACCEPTED: {
        QObject::connect(self,SIGNAL(acceped()),drvNewSignal(self,a1,a2),SLOT(call()));
        break;
    }
    case _ID_DIALOG_ONREJECTED: {
        QObject::connect(self,SIGNAL(rejected()),drvNewSignal(self,a1,a2),SLOT(call()));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_combobox(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QComboBox *self = (QComboBox*)drvGetNative(a0);
    switch (drvid) {
    case _ID_COMBOBOX_INIT: {
        drvNewObj(a0,new QComboBox);
        break;
    }
    case _ID_COMBOBOX_COUNT: {
        drvSetInt(a1,self->count());
        break;
    }
    case _ID_COMBOBOX_SETCURRENTINDEX: {
        self->setCurrentIndex(drvGetInt(a1));
        break;
    }
    case _ID_COMBOBOX_CURRENTINDEX: {
        drvSetInt(a1,self->currentIndex());
        break;
    }
    case _ID_COMBOBOX_CURRENTTEXT: {
        drvSetString(a1,self->currentText());
        break;
    }
    case _ID_COMBOBOX_SETEDITABLE: {
        self->setEditable(drvGetBool(a1));
        break;
    }
    case _ID_COMBOBOX_ISEDITABLE: {
        drvSetBool(a1,self->isEditable());
        break;
    }
    case _ID_COMBOBOX_SETMAXCOUNT: {
        self->setMaxCount(drvGetInt(a1));
        break;
    }
    case _ID_COMBOBOX_MAXCOUNT: {
        drvSetInt(a1,self->maxCount());
        break;
    }
    case _ID_COMBOBOX_SETMAXVISIBLEITEMS: {
        self->setMaxVisibleItems(drvGetInt(a1));
        break;
    }
    case _ID_COMBOBOX_MAXVISIBLEITEMS: {
        drvSetInt(a1,self->maxVisibleItems());
        break;
    }
    case _ID_COMBOBOX_SETMINIMUMCONTENTSLENGTH: {
        self->setMinimumContentsLength(drvGetInt(a1));
        break;
    }
    case _ID_COMBOBOX_MINIMUNCONTENTSLENGHT: {
        drvSetInt(a1,self->minimumContentsLength());
        break;
    }
    case _ID_COMBOBOX_ADDITEM: {
        self->addItem(drvGetString(a1));
        break;
    }
    case _ID_COMBOBOX_INSERTITEM: {
        self->insertItem(drvGetInt(a1),drvGetString(a2));
        break;
    }
    case _ID_COMBOBOX_REMOVEITEM: {
        self->removeItem(drvGetInt(a1));
        break;
    }
    case _ID_COMBOBOX_ITEMTEXT: {
        drvSetString(a2,self->itemText(drvGetInt(a1)));
        break;
    }
    case _ID_COMBOBOX_ONCURRENTINDEXCHANGED: {
        QObject::connect(self,SIGNAL(currentIndexChanged(int)),drvNewSignal(self,a1,a2),SLOT(call(int)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_lineedit(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QLineEdit *self = (QLineEdit*)drvGetNative(a0);
    switch (drvid) {
    case _ID_LINEEDIT_INIT: {
        drvNewObj(a0,new QLineEdit);
        break;
    }
    case _ID_LINEEDIT_INITWITHTEXT: {
        drvNewObj(a0, new QLineEdit(drvGetString(a1)));
        break;
    }
    case _ID_LINEEDIT_SETTEXT: {
        self->setText(drvGetString(a1));
        break;
    }
    case _ID_LINEEDIT_TEXT: {
        drvSetString(a1,self->text());
        break;
    }
    case _ID_LINEEDIT_SETINPUTMASK: {
        self->setInputMask(drvGetString(a1));
        break;
    }
    case _ID_LINEEDIT_INPUTMASK: {
        drvSetString(a1,self->inputMask());
        break;
    }
    case _ID_LINEEDIT_SETALIGNMENT: {
        self->setAlignment(drvGetAlignment(a1));
        break;
    }
    case _ID_LINEEDIT_ALIGNMENT: {
        drvSetAlignment(a1,self->alignment());
        break;
    }
    case _ID_LINEEDIT_SETCURSORPOS: {
        self->setCursorPosition(drvGetInt(a1));
        break;
    }
    case _ID_LINEEDIT_CURSORPOS: {
        drvSetInt(a1,self->cursorPosition());
        break;
    }
    case _ID_LINEEDIT_SETDRAGENABLED: {
        self->setDragEnabled(drvGetBool(a1));
        break;
    }
    case _ID_LINEEDIT_DRAGENABLED: {
        drvSetBool(a1,self->dragEnabled());
        break;
    }
    case _ID_LINEEDIT_SETREADONLY: {
        self->setReadOnly(drvGetBool(a1));
        break;
    }
    case _ID_LINEEDIT_ISREADONLY: {
        drvSetBool(a1,self->isReadOnly());
        break;
    }
    case _ID_LINEEDIT_SETFRAME: {
        self->setFrame(drvGetBool(a1));
        break;
    }
    case _ID_LINEEDIT_HASFRAME: {
        drvSetBool(a1,self->hasFrame());
        break;
    }
    case _ID_LINEEDIT_ISREDOAVAILABLE: {
        drvSetBool(a1,self->isRedoAvailable());
        break;
    }
    case _ID_LINEEDIT_HASSELECTED: {
        drvSetBool(a1,self->hasSelectedText());
        break;
    }
    case _ID_LINEEDIT_SELECTEDTEXT: {
        drvSetString(a1,self->selectedText());
        break;
    }
    case _ID_LINEEDIT_SELSTART: {
        drvSetInt(a1,self->selectionStart());
        break;
    }
    case _ID_LINEEDIT_SETSEL: {
        self->setSelection(drvGetInt(a1),drvGetInt(a2));
        break;
    }
    case _ID_LINEEDIT_CANCELSEL: {
        self->deselect();
        break;
    }
    case _ID_LINEEDIT_SELECTALL: {
        self->selectAll();
        break;
    }
    case _ID_LINEEDIT_COPY: {
        self->copy();
        break;
    }
    case _ID_LINEEDIT_CUT: {
        self->cut();
        break;
    }
    case _ID_LINEEDIT_PASTE: {
        self->paste();
        break;
    }
    case _ID_LINEEDIT_REDO: {
        self->redo();
        break;
    }
    case _ID_LINEEDIT_UNDO: {
        self->undo();
        break;
    }
    case _ID_LINEEDIT_CLEAR: {
        self->clear();
        break;
    }
    case _ID_LINEEDIT_ONTEXTCHANGED: {
        QObject::connect(self,SIGNAL(textChanged(QString)),drvNewSignal(self,a1,a2),SLOT(call(QString)));
        break;
    }
    case _ID_LINEEDIT_ONEDITINGFINISHED: {
        QObject::connect(self,SIGNAL(editingFinished(QString)),drvNewSignal(self,a1,a2),SLOT(call(QString)));
        break;
    }
    case _ID_LINEEDIT_ONRETURNPRESSED: {
        QObject::connect(self,SIGNAL(returnPressed(QString)),drvNewSignal(self,a1,a2),SLOT(call(QString)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_baseslider(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QAbstractSlider *self = (QAbstractSlider*)drvGetNative(a0);
    switch (drvid) {
    case _ID_BASESLIDER_SETTRACKING: {
        self->setTracking(drvGetBool(a1));
        break;
    }
    case _ID_BASESLIDER_HASTRACKING: {
        drvSetBool(a1,self->hasTracking());
        break;
    }
    case _ID_BASESLIDER_SETMAXIMUM: {
        self->setMaximum(drvGetInt(a1));
        break;
    }
    case _ID_BASESLIDER_MAXIMUM: {
        drvSetInt(a1,self->maximum());
        break;
    }
    case _ID_BASESLIDER_SETMINIMUM: {
        self->setMinimum(drvGetInt(a1));
        break;
    }
    case _ID_BASESLIDER_MINIMUM: {
        drvSetInt(a1,self->minimum());
        break;
    }
    case _ID_BASESLIDER_SETORIENTATION: {
        self->setOrientation(drvGetOrientation(a1));
        break;
    }
    case _ID_BASESLIDER_ORIENTATION: {
        drvSetOrientation(a1,self->orientation());
        break;
    }
    case _ID_BASESLIDER_SETPAGESTEP: {
        self->setPageStep(drvGetInt(a1));
        break;
    }
    case _ID_BASESLIDER_PAGESTEP: {
        drvSetInt(a1,self->pageStep());
        break;
    }
    case _ID_BASESLIDER_SETSINGLESTEP: {
        self->setSingleStep(drvGetInt(a1));
        break;
    }
    case _ID_BASESLIDER_SINGLESTEP: {
        drvSetInt(a1,self->singleStep());
        break;
    }
    case _ID_BASESLIDER_SETSLIDERDOWN: {
        self->setSliderDown(drvGetBool(a1));
        break;
    }
    case _ID_BASESLIDER_ISSLIDERDOWN: {
        drvSetBool(a1,self->isSliderDown());
        break;
    }
    case _ID_BASESLIDER_SETSLIDERPOSITION: {
        self->setSliderPosition(drvGetInt(a1));
        break;
    }
    case _ID_BASESLIDER_SLIDERPOSITION: {
        drvSetInt(a1,self->sliderPosition());
        break;
    }
    case _ID_BASESLIDER_SETVALUE: {
        self->setValue(drvGetInt(a1));
        break;
    }
    case _ID_BASESLIDER_VALUE: {
        drvSetInt(a1,self->value());
        break;
    }
    case _ID_BASESLIDER_SETRANGE: {
        self->setRange(drvGetInt(a1),drvGetInt(a2));
        break;
    }
    case _ID_BASESLIDER_RANGE: {
        drvSetInt(a1,self->minimum());
        drvSetInt(a2,self->maximum());
        break;
    }
    case _ID_BASESLIDER_ONVALUECHANGED: {
        QObject::connect(self,SIGNAL(valueChanged(int)),drvNewSignal(self,a1,a2),SLOT(call(int)));
        break;
    }
    case _ID_BASESLIDER_ONSLIDERPRESSED: {
        QObject::connect(self,SIGNAL(sliderPressed()),drvNewSignal(self,a1,a2),SLOT(call()));
        break;
    }
    case _ID_BASESLIDER_ONSLIDERRELEASED: {
        QObject::connect(self,SIGNAL(sliderReleased()),drvNewSignal(self,a1,a2),SLOT(call()));
        break;
    }
    case _ID_BASESLIDER_ONSLIDERMOVED: {
        QObject::connect(self,SIGNAL(sliderMoved()),drvNewSignal(self,a1,a2),SLOT(call(int)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_slider(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QSlider *self = (QSlider*)drvGetNative(a0);
    switch (drvid) {
    case _ID_SLIDER_INIT: {
        drvNewObj(a0,new QSlider(Qt::Horizontal));
        break;
    }
    case _ID_SLIDER_SETTICKINTERVAL: {
        self->setTickInterval(drvGetInt(a1));
        break;
    }
    case _ID_SLIDER_TICKINTERVAL: {
        drvSetInt(a1,self->tickInterval());
        break;
    }
    case _ID_SLIDER_SETTICKPOSITION: {
        self->setTickPosition(drvGetTickPosition(a1));
        break;
    }
    case _ID_SLIDER_TICKPOSITION: {
        drvSetTickPosition(a1,self->tickPosition());
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_scrollbar(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QScrollBar *self = (QScrollBar*)drvGetNative(a0);
    switch (drvid) {
    case _ID_SCROLLBAR_INIT: {
        drvNewObj(a0,new QScrollBar(Qt::Horizontal));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_dial(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QDial *self = (QDial*)drvGetNative(a0);
    switch (drvid) {
    case _ID_DIAL_INIT: {
        drvNewObj(a0,new QDial);
        break;
    }
    case _ID_DIAL_NOTCHSIZE: {
        drvSetInt(a1,self->notchSize());
        break;
    }
    case _ID_DIAL_SETNOTCHTARGET: {
        self->setNotchTarget(drvGetFloat64(a1));
        break;
    }
    case _ID_DIAL_NOTCHTARGET: {
        drvSetFloat64(a1,self->notchTarget());
        break;
    }
    case _ID_DIAL_SETNOTCHESVISIBLE: {
        self->setNotchesVisible(drvGetBool(a1));
        break;
    }
    case _ID_DIAL_NOTCHESVISIBLE: {
        drvSetBool(a1,self->notchesVisible());
        break;
    }
    case _ID_DIAL_SETWRAPPING: {
        self->setWrapping(drvGetBool(a1));
        break;
    }
    case _ID_DIAL_WRAPPING: {
        drvSetBool(a1,self->wrapping());
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_brush(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QBrush *self = (QBrush*)drvGetNative(a0);
    switch (drvid) {
    case _ID_BRUSH_INIT: {
        drvSetBrush(a0, QBrush());
        break;
    }
    case _ID_BRUSH_CLOSE: {
        drvDelBrush(a0,self);
        break;
    }
    case _ID_BRUSH_SETCOLOR: {
        self->setColor(drvGetColor(a1));
        break;
    }
    case _ID_BRUSH_COLOR: {
        drvSetColor(a1,self->color());
        break;
    }
    case _ID_BRUSH_SETSTYLE: {
        self->setStyle((Qt::BrushStyle)drvGetInt(a1));
        break;
    }
    case _ID_BRUSH_STYLE: {
        drvSetInt(a1,self->style());
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_pen(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QPen *self = (QPen*)drvGetNative(a0);
    switch (drvid) {
    case _ID_PEN_INIT: {
        drvSetPen(a0, QPen());
        break;
    }
    case _ID_PEN_CLOSE: {
        drvDelPen(a0,self);
        break;
    }
    case _ID_PEN_SETCOLOR: {
        self->setColor(drvGetColor(a1));
        break;
    }
    case _ID_PEN_COLOR: {
        drvSetColor(a1,self->color());
        break;
    }
    case _ID_PEN_SETWIDTH: {
        self->setWidth(drvGetInt(a1));
        break;
    }
    case _ID_PEN_WIDTH: {
        drvSetInt(a1,self->width());
        break;
    }
    case _ID_PEN_SETSTYLE: {
        self->setStyle((Qt::PenStyle)drvGetInt(a1));
        break;
    }
    case _ID_PEN_STYLE: {
        drvSetInt(a1,self->style());
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_painter(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QPainter *self = (QPainter*)drvGetNative(a0);
    switch (drvid) {
    case _ID_PAINTER_INIT: {
        drvNewObj(a0,new QPainter);
        break;
    }
    case _ID_PAINTER_INITWITHIMAGE: {
        drvNewObj(a0,new QPainter((QImage*)drvGetNative(a1)));
        break;
    }
    case _ID_PAINTER_CLOSE: {
        drvDelObj(a0,self);
        break;
    }
    case _ID_PAINTER_INITFROM: {
        self->initFrom(drvGetWidget(a1));
        break;
    }
    case _ID_PAINTER_BEGIN: {
        self->begin(drvGetWidget(a1));
        break;
    }
    case _ID_PAINTER_END: {
        self->end();
        break;
    }
    case _ID_PAINTER_SETFONT: {
        self->setFont(drvGetFont(a1));
        break;
    }
    case _ID_PAINTER_FONT: {
        drvSetFont(a1,self->font());
        break;
    }
    case _ID_PAINTER_SETPEN: {
        self->setPen(drvGetPen(a1));
        break;
    }
    case _ID_PAINTER_PEN: {
        drvSetPen(a1,self->pen());
        break;
    }
    case _ID_PAINTER_SETBRUSH: {
        self->setBrush(drvGetBrush(a1));
        break;
    }
    case _ID_PAINTER_BRUSH: {
        drvSetBrush(a1,self->brush());
        break;
    }
    case _ID_PAINTER_DRAWPOINT: {
        self->drawPoint(drvGetPoint(a1));
        break;
    }
    case _ID_PAINTER_DRAWPOINTS: {
        self->drawPoints(drvGetPointArray(a1));
        break;
    }
    case _ID_PAINTER_DRAWLINE: {
        self->drawLine(drvGetPoint(a1),drvGetPoint(a2));
        break;
    }
    case _ID_PAINTER_DRAWLINES: {
        self->drawLines(drvGetPointArray(a1));
        break;
    }
    case _ID_PAINTER_DRAWPOLYLINE: {
        self->drawPolyline(drvGetPointArray(a1));
        break;
    }
    case _ID_PAINTER_DRAWPOLYGON: {
        self->drawPolygon(drvGetPointArray(a1));
        break;
    }
    case _ID_PAINTER_DRAWRECT: {
        self->drawRect(drvGetRect(a1));
        break;
    }
    case _ID_PAINTER_DRAWRECTS: {
        self->drawRects(drvGetRectArray(a1));
        break;
    }
    case _ID_PAINTER_DRAWROUNDEDRECT: {
        self->drawRoundedRect(drvGetRect(a1),drvGetFloat64(a2),drvGetFloat64(a3),(Qt::SizeMode)drvGetInt(a4));
        break;
    }
    case _ID_PAINTER_DRAWELLIPSE: {
        self->drawEllipse(drvGetRect(a1));
        break;
    }
    case _ID_PAINTER_DRAWARC: {
        self->drawArc(drvGetRect(a1),drvGetInt(a2),drvGetInt(a3));
        break;
    }
    case _ID_PAINTER_DRAWCHORD: {
        self->drawChord(drvGetRect(a1),drvGetInt(a2),drvGetInt(a3));
        break;
    }
    case _ID_PAINTER_DRAWPIE: {
        self->drawPie(drvGetRect(a1),drvGetInt(a2),drvGetInt(a3));
        break;
    }
    case _ID_PAINTER_DRAWTEXT: {
        self->drawText(drvGetPoint(a1),drvGetString(a2));
        break;
    }
    case _ID_PAINTER_DRAWTEXTRECT: {
        QRect bound;
        self->drawText(drvGetRect(a1),drvGetInt(a2),drvGetString(a3),&bound);
        drvSetRect(a4,bound);
        break;
    }
    case _ID_PAINTER_DRAWPIXMAP: {
        self->drawPixmap(drvGetPoint(a1),drvGetPixmap(a2));
        break;
    }
    case _ID_PAINTER_DRAWPIXMAPEX: {
        self->drawPixmap(drvGetPoint(a1),drvGetPixmap(a2),drvGetRect(a3));
        break;
    }
    case _ID_PAINTER_DRAWPIXMAPRECT: {
        self->drawPixmap(drvGetRect(a1),drvGetPixmap(a2));
        break;
    }
    case _ID_PAINTER_DRAWPIXMAPRECTEX: {
        self->drawPixmap(drvGetRect(a1),drvGetPixmap(a2),drvGetRect(a3));
        break;
    }
    case _ID_PAINTER_DRAWIMAGE: {
        self->drawImage(drvGetPoint(a1),drvGetImage(a2));
        break;
    }
    case _ID_PAINTER_DRAWIMAGEEX: {
        self->drawImage(drvGetPoint(a1),drvGetImage(a2),drvGetRect(a3));
        break;
    }
    case _ID_PAINTER_DRAWIMAGERECT: {
        self->drawImage(drvGetRect(a1),drvGetImage(a2));
        break;
    }
    case _ID_PAINTER_DRAWIMAGERECTEX: {
        self->drawImage(drvGetRect(a1),drvGetImage(a2),drvGetRect(a3));
        break;
    }
    case _ID_PAINTER_FILLRECT: {
        self->fillRect(drvGetRect(a1),drvGetUint(a2));
        break;
    }
    case _ID_PAINTER_FILLRECTF: {
        self->fillRect(drvGetRectF(a1),drvGetUint(a2));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_listwidgetitem(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    ShellListWidgetItem *self = (ShellListWidgetItem*)drvGetNative(a0);
    switch (drvid) {
    case _ID_LISTWIDGETITEM_INIT: {
        drvNewObj(a0,new ShellListWidgetItem);
        break;
    }
    case _ID_LISTWIDGETITEM_INITWITHTEXT: {
        drvNewObj(a0, new ShellListWidgetItem(drvGetString(a1)));
        break;
    }
    case _ID_LISTWIDGETITEM_CLOSE: {
        drvDelObj(a0,self);
        break;
    }
    case _ID_LISTWIDGETITEM_SETTEXT: {
        self->setText(drvGetString(a1));
        break;
    }
    case _ID_LISTWIDGETITEM_TEXT: {
        drvSetString(a1,self->text());
        break;
    }
    case _ID_LISTWIDGETITEM_SETICON: {
        self->setIcon(drvGetIcon(a1));
        break;
    }
    case _ID_LISTWIDGETITEM_ICON: {
        drvSetIcon(a1,self->icon());
        break;
    }
    case _ID_LISTWIDGETITEM_SETSELECTED: {
        self->setSelected(drvGetBool(a1));
        break;
    }
    case _ID_LISTWIDGETITEM_ISSELECTED: {
        drvSetBool(a1,self->isSelected());
        break;
    }
    case _ID_LISTWIDGETITEM_SETHIDDEN: {
        self->setHidden(drvGetBool(a1));
        break;
    }
    case _ID_LISTWIDGETITEM_ISHIDDEN: {
        drvSetBool(a1,self->isHidden());
        break;
    }
    case _ID_LISTWIDGETITEM_SETFONT: {
        self->setFont(drvGetFont(a1));
        break;
    }
    case _ID_LISTWIDGETITEM_FONT: {
        self->font();
        break;
    }
    case _ID_LISTWIDGETITEM_SETFOREGROUND: {
        self->setForeground(drvGetBrush(a1));
        break;
    }
    case _ID_LISTWIDGETITEM_FOREGROUND: {
        drvSetBrush(a1,self->foreground());
        break;
    }
    case _ID_LISTWIDGETITEM_SETBACKGROUND: {
        self->setBackground(drvGetBrush(a1));
        break;
    }
    case _ID_LISTWIDGETITEM_BACKGROUND: {
        drvSetBrush(a1,self->background());
        break;
    }
    case _ID_LISTWIDGETITEM_SETTOOLTIP: {
        self->setToolTip(drvGetString(a1));
        break;
    }
    case _ID_LISTWIDGETITEM_TOOLTIP: {
        drvSetString(a1,self->toolTip());
        break;
    }
    case _ID_LISTWIDGETITEM_SETTEXTALIGNMENT: {
        self->setTextAlignment((Qt::Alignment)drvGetInt(a1));
        break;
    }
    case _ID_LISTWIDGETITEM_TEXTALIGNMENT: {
        drvSetInt(a1,self->textAlignment());
        break;
    }
    case _ID_LISTWIDGETITEM_SETFLAGS: {
        self->setFlags((Qt::ItemFlag)drvGetInt(a1));
        break;
    }
    case _ID_LISTWIDGETITEM_FLAGS: {
        drvSetInt(a1,self->flags());
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_listwidget(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QListWidget *self = (QListWidget*)drvGetNative(a0);
    switch (drvid) {
    case _ID_LISTWIDGET_INIT: {
        drvNewObj(a0,new QListWidget);
        break;
    }
    case _ID_LISTWIDGET_COUNT: {
        drvSetInt(a1,self->count());
        break;
    }
    case _ID_LISTWIDGET_SETCURRENTITEM: {
        self->setCurrentItem(drvGetListWidgetItem(a1));
        break;
    }
    case _ID_LISTWIDGET_CURRENTITEM: {
        drvSetListWidgetItem(a1,self->currentItem());
        break;
    }
    case _ID_LISTWIDGET_SETCURRENTROW: {
        self->setCurrentRow(drvGetInt(a1));
        break;
    }
    case _ID_LISTWIDGET_CURRENTROW: {
        drvSetInt(a1,self->currentRow());
        break;
    }
    case _ID_LISTWIDGET_ADDITEM: {
        self->addItem(drvGetListWidgetItem(a1));
        break;
    }
    case _ID_LISTWIDGET_INSERTITEM: {
        self->insertItem(drvGetInt(a1),drvGetListWidgetItem(a2));
        break;
    }
    case _ID_LISTWIDGET_EDITITEM: {
        self->editItem(drvGetListWidgetItem(a1));
        break;
    }
    case _ID_LISTWIDGET_TAKEITEM: {
        drvSetListWidgetItem(a2,self->takeItem(drvGetInt(a1)));
        break;
    }
    case _ID_LISTWIDGET_ITEM: {
        drvSetListWidgetItem(a2,self->item(drvGetInt(a1)));
        break;
    }
    case _ID_LISTWIDGET_CLEAR: {
        self->clear();
        break;
    }
    case _ID_LISTWIDGET_ONCURRENTITEMCHANGED: {
        QObject::connect(self,SIGNAL(currentItemChanged(QListWidgetItem*,QListWidgetItem*)),drvNewSignal(self,a1,a2),SLOT(call(QListWidgetItem*,QListWidgetItem*)));
        break;
    }
    case _ID_LISTWIDGET_ONCURRENTROWCHANGED: {
        QObject::connect(self,SIGNAL(currentRowChanged(int)),drvNewSignal(self,a1,a2),SLOT(call(int)));
        break;
    }
    case _ID_LISTWIDGET_ONITEMACTIVATED: {
        QObject::connect(self,SIGNAL(itemActivated(QListWidgetItem*)),drvNewSignal(self,a1,a2),SLOT(call(QListWidgetItem*)));
        break;
    }
    case _ID_LISTWIDGET_ONITEMCHANGED: {
        QObject::connect(self,SIGNAL(itemChanged(QListWidgetItem*)),drvNewSignal(self,a1,a2),SLOT(call(QListWidgetItem*)));
        break;
    }
    case _ID_LISTWIDGET_ONITEMCLICKED: {
        QObject::connect(self,SIGNAL(itemClicked(QListWidgetItem*)),drvNewSignal(self,a1,a2),SLOT(call(QListWidgetItem*)));
        break;
    }
    case _ID_LISTWIDGET_ONITEMDOUBLECLICKED: {
        QObject::connect(self,SIGNAL(itemDoubleClicked(QListWidgetItem*)),drvNewSignal(self,a1,a2),SLOT(call(QListWidgetItem*)));
        break;
    }
    case _ID_LISTWIDGET_ONITEMENTERED: {
        QObject::connect(self,SIGNAL(itemEntered(QListWidgetItem*)),drvNewSignal(self,a1,a2),SLOT(call(QListWidgetItem*)));
        break;
    }
    case _ID_LISTWIDGET_ONITEMPRESSED: {
        QObject::connect(self,SIGNAL(itemPressed(QListWidgetItem*)),drvNewSignal(self,a1,a2),SLOT(call(QListWidgetItem*)));
        break;
    }
    case _ID_LISTWIDGET_ONITEMSELECTIONCHANGED: {
        QObject::connect(self,SIGNAL(itemSelectionChanged(QListWidgetItem*)),drvNewSignal(self,a1,a2),SLOT(call(QListWidgetItem*)));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

int drv_mainwindow(int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6)
{
    QMainWindow *self = (QMainWindow*)drvGetNative(a0);
    switch (drvid) {
    case _ID_MAINWINDOW_INIT: {
        drvNewObj(a0,new QMainWindow);
        break;
    }
    case _ID_MAINWINDOW_SETCENTRALWIDGET: {
        self->setCentralWidget(drvGetWidget(a1));
        break;
    }
    case _ID_MAINWINDOW_CENTRALWIDGET: {
        drvSetHandle(a1,self->centralWidget());
        break;
    }
    case _ID_MAINWINDOW_SETMENUBAR: {
        self->setMenuBar(drvGetMenuBar(a1));
        break;
    }
    case _ID_MAINWINDOW_MENUBAR: {
        drvSetMenuBar(a1,self->menuBar());
        break;
    }
    case _ID_MAINWINDOW_SETSTATUSBAR: {
        self->setStatusBar(drvGetStatusBar(a1));
        break;
    }
    case _ID_MAINWINDOW_STATUSBAR: {
        drvSetStatusBar(a1,self->statusBar());
        break;
    }
    case _ID_MAINWINDOW_ADDTOOLBAR: {
        self->addToolBar(drvGetToolBar(a1));
        break;
    }
    case _ID_MAINWINDOW_INSERTTOOLBAR: {
        self->insertToolBar(drvGetToolBar(a1),drvGetToolBar(a2));
        break;
    }
    case _ID_MAINWINDOW_REMOVETOOLBAR: {
        self->removeToolBar(drvGetToolBar(a1));
        break;
    }
    case _ID_MAINWINDOW_ADDDOCKWIDGET: {
        self->addDockWidget((Qt::DockWidgetArea)drvGetInt(a1),drvGetDockWidget(a2));
        break;
    }
    case _ID_MAINWINDOW_REMOVEDOCKWIDGET: {
        self->removeDockWidget(drvGetDockWidget(a1));
        break;
    }
    default:
        return 0;
    }
    return 1;
}

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
    switch(drvcls) {
    case _CLASSID_APP:
        return drv_app(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_TIMER:
        return drv_timer(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_FONT:
        return drv_font(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_PIXMAP:
        return drv_pixmap(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_ICON:
        return drv_icon(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_IMAGE:
        return drv_image(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_WIDGET:
        return drv_widget(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_ACTION:
        return drv_action(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_ACTIONGROUP:
        return drv_actiongroup(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_MENU:
        return drv_menu(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_MENUBAR:
        return drv_menubar(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_TOOLBAR:
        return drv_toolbar(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_STATUSBAR:
        return drv_statusbar(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_DOCKWIDGET:
        return drv_dockwidget(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_SYSTEMTRAY:
        return drv_systemtray(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_TABWIDGET:
        return drv_tabwidget(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_TOOLBOX:
        return drv_toolbox(drvid,a0,a1,a2,a3,a4,a5,a6);
    case _CLASSID_BASELAYOUT:
        return drv_baselayout(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_BOXLAYOUT:
        return drv_boxlayout(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_HBOXLAYOUT:
        return drv_hboxlayout(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_VBOXLAYOUT:
        return drv_vboxlayout(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_STACKEDLAYOUT:
        return drv_stackedlayout(drvid,a0,a1,a2,a3,a4,a5,a6);
    case _CLASSID_BASEBUTTON:
        return drv_basebutton(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_BUTTON:
        return drv_button(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_CHECKBOX:
        return drv_checkbox(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_RADIO:
        return drv_radio(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_TOOLBUTTON:
        return drv_toolbutton(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_FRAME:
        return drv_frame(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_LABEL:
        return drv_label(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_GROUPBOX:
        return drv_groupbox(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_DIALOG:
        return drv_dialog(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_COMBOBOX:
        return drv_combobox(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_LINEEDIT:
        return drv_lineedit(drvid,a0,a1,a2,a3,a4,a5,a6);
    case _CLASSID_BASESLIDER:
        return drv_baseslider(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_SLIDER:
        return drv_slider(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_SCROLLBAR:
        return drv_scrollbar(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_DIAL:
        return drv_dial(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_BRUSH:
        return drv_brush(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_PEN:
        return drv_pen(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_PAINTER:
        return drv_painter(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_LISTWIDGETITEM:
        return drv_listwidgetitem(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_LISTWIDGET:
        return drv_listwidget(drvid,a0,a1,a2,a3,a4,a5,a6);
    case CLASSID_MAINWINDOW:
        return drv_mainwindow(drvid,a0,a1,a2,a3,a4,a5,a6);
    default:
        return 0;
    }
    return 1;
}

