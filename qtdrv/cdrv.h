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
#include "qtdrv_global.h"
#include "qtsignal.h"
#include "qtevent.h"
#include "qtapp.h"
#include "qtshell.h"

#include <QApplication>
#include <QWidget>
#include <QLayout>
#include <QMetaType>
#include <QVariant>
#include <QEvent>
#include <QSlider>
#include <QSystemTrayIcon>
#include <QToolButton>
#include <QIcon>
#include <QPixmap>
#include <QPen>
#include <QBrush>
#include <QDebug>
#include <QSizePolicy>

class QDockWidget;
class QToolBar;
class QMenu;
class QMenuBar;
class QAction;
class QLayout;
class QStatusBar;
class QListWidgetItem;


#define drvInvalid (void*)-1

void utf8_info_copy(void*,const char*,int);

struct handle_head {
    void *native;
    int   classid;
};

struct string_head {
    char *data;
    int len;
};

struct slice_head {
    const char *data;
    int len;
    int cap;
};

struct point_head {
    int x;
    int y;
};

struct pointf_head {
    double x;
    double y;
};

struct size_head {
    int width;
    int hegiht;
};

struct sizef_head {
    double width;
    double height;
};

struct rect_head {
    int x;
    int y;
    int width;
    int height;
};

struct rectf_head {
    double x;
    double y;
    double width;
    double height;
};

struct margins_head {
    int left;
    int top;
    int right;
    int bottom;
};

struct color_head {
    quint32 r;
    quint32 g;
    quint32 b;
    quint32 a;
};

typedef unsigned char byte;

inline QString drvGetString(void *param)
{
    if (param == 0) {
        return QString();
    }
    string_head *h = (string_head*)param;
    return QString::fromUtf8(h->data,h->len);
}

inline void drvSetString(void *param, const QString &s)
{
    if (param == 0) {
        return;
    }    
    const QByteArray & ar = s.toUtf8();
    utf8_info_copy(param,ar.constData(),ar.length());
    /*
    string_head *sh = (string_head*)param;
    sh->data = new char[ar.length()];
    memcpy(sh->data,ar.constData(),ar.length());
    sh->len = ar.length();
    */
}

inline QColor drvGetColor(void *param)
{
    if (param == 0) {
        return QColor();
    }
    return QColor(*(QRgb*)param);
}

inline void drvSetColor(void *param, const QColor &clr)
{
    if (param == 0) {
        return;
    }
    *(QRgb*)param = clr.rgba();
}

inline QPoint drvGetPoint(void *param)
{
    if (param == 0) {
        return QPoint();
    }
    point_head *h = (point_head*)param;
    return QPoint(h->x,h->y);
}

inline QVector<QPoint> drvGetPointArray(void *param)
{
    if (param == 0) {
        return QVector<QPoint> ();
    }
    slice_head *sh = (slice_head*)param;
    QVector<QPoint> vec(sh->len);
    point_head *h = (point_head*)sh->data;
    for (int i=0; i < sh->len; i++) {
        vec[i] = QPoint(h[i].x,h[i].y);
    }
    return vec;
}

inline QVector<QRect> drvGetRectArray(void *param)
{
    if (param == 0) {
        return QVector<QRect>();
    }
    slice_head *sh = (slice_head*)param;
    QVector<QRect> vec(sh->len);
    rect_head *h = (rect_head*)sh->data;
    for (int i=0; i < sh->len; i++) {
        vec[i] = QRect(h[i].x,h[i].y,h[i].width,h[i].height);
    }
    return vec;
}

inline QByteArray drvGetByteArray(void *param)
{
    if (param == 0) {
        return QByteArray();
    }
    slice_head *sh = (slice_head*)param;
    return QByteArray(sh->data,sh->len);
}

inline void drvSetPoint(void *param, const QPoint &pt)
{
    if (param == 0) {
        return;
    }
    point_head *sh = (point_head*)param;
    sh->x = pt.x();
    sh->y = pt.y();
}

inline QSize drvGetSize(void *param)
{
    if (param == 0) {
        return QSize();
    }
    size_head *h = (size_head*)param;
    return QSize(h->width,h->hegiht);
}

inline void drvSetSize(void *param, const QSize &sz)
{
    if (param == 0) {
        return;
    }
    size_head *h = (size_head*)param;
    h->width = sz.width();
    h->hegiht = sz.height();
}

inline QRect drvGetRect(void *param)
{
    if (param == 0) {
        return QRect();
    }
    rect_head *h = (rect_head*)param;
    return QRect(h->x,h->y,h->width,h->height);
}

inline QRectF drvGetRectF(void *param)
{
    if (param == 0) {
        return QRectF();
    }
    rectf_head *h = (rectf_head*)param;
    return QRectF(h->x,h->y,h->width,h->height);
}


inline void drvSetRect(void *param, const QRect &rc)
{
    if (param == 0) {
        return;
    }
    rect_head *h = (rect_head*)param;
    h->x = rc.x();
    h->y = rc.y();
    h->width = rc.width();
    h->height = rc.height();
}

inline QMargins drvGetMargins(void *param)
{
    if (param == 0) {
        return QMargins();
    }
    margins_head *h = (margins_head*)param;
    return QMargins(h->left,h->top,h->right,h->bottom);
}

inline void drvSetMargins(void *param, const QMargins &mr)
{
    if (param == 0) {
        return;
    }
    margins_head *h = (margins_head*)param;
    h->left = mr.left();
    h->top = mr.top();
    h->right = mr.right();
    h->bottom = mr.bottom();
}

inline bool drvGetBool(void *param)
{
    return *(int*)param != 0;
}

inline void drvSetBool(void *param, bool b)
{
    *(int*)param = b?1:0;
}

inline int drvGetInt(void *param)
{
    return *(int*)param;
}

inline int drvGetUint(void *param)
{
    return *(uint*)param;
}

inline void drvSetInt(void *param, int value)
{
    *(int*)param = value;
}

inline QSystemTrayIcon::MessageIcon drvGetMessageIconType(void *param)
{
    return QSystemTrayIcon::MessageIcon(*(int*)param);
}

inline Qt::Alignment drvGetAlignment(void *param)
{
    return (Qt::Alignment)(*(int*)param);
}

inline void drvSetAlignment(void *param,Qt::Alignment value)
{
    *(int*)param = value;
}

inline Qt::Orientation drvGetOrientation(void *param)
{
    return (Qt::Orientation)(*(int*)param);
}

inline void drvSetOrientation(void *param, Qt::Orientation value)
{
    *(int*)param = value;
}

inline QSlider::TickPosition drvGetTickPosition(void *param)
{
    return (QSlider::TickPosition)(*(int*)param);
}

inline void drvSetTickPosition(void *param, QSlider::TickPosition value)
{
    *(int*)param = value;
}

inline Qt::ToolButtonStyle drvGetToolButtonStyle(void *param)
{
    return (Qt::ToolButtonStyle)(*(int*)param);
}

inline void drvSetToolButtonStyle(void *param, Qt::ToolButtonStyle value)
{
    *(int*)param = value;
}

inline QToolButton::ToolButtonPopupMode drvGetToolButtonPopupMode(void *param)
{
    return (QToolButton::ToolButtonPopupMode)(*(int*)param);
}

inline void drvSetToolButtonPopupMode(void *param, QToolButton::ToolButtonPopupMode value)
{
    *(int*)param = value;
}


inline double drvGetFloat64(void *param)
{
    return *(double*)param;
}

inline void drvSetFloat64(void *param, double value)
{
    *(double*)param = value;
}

inline void drvSetHandle(void *param, void *obj)
{
    if (obj && param) {
         handle_head *head = (handle_head*)param;
         head->native = obj;
    }
}

inline void* drvGetNative(void *param)
{
    if (param == 0) {
        return 0;
    }    
    return ((handle_head*)param)->native;
}

inline QWidget* drvGetWidget(void *param)
{
    if (param == 0) {
        return 0;
    }
    return (QWidget*)((handle_head*)param)->native;
}

inline QLayout* drvGetLayout(void *param)
{
    if (param == 0) {
        return 0;
    }
    return (QLayout*)((handle_head*)param)->native;
}

inline QMenu* drvGetMenu(void *param)
{
    if (param == 0) {
        return 0;
    }
    return (QMenu*)((handle_head*)param)->native;
}

inline QMenuBar* drvGetMenuBar(void *param)
{
    return (QMenuBar*)drvGetNative(param);
}

inline QAction* drvGetAction(void *param)
{
    if (param == 0) {
        return 0;
    }
    return (QAction*)((handle_head*)param)->native;
}

inline QIcon drvGetIcon(void *param)
{
    QIcon *icon = (QIcon*)drvGetNative(param);
    if (!icon) {
        return QIcon();
    }
    return *icon;
}

inline void drvSetIcon(void *param, const QIcon &icon)
{
    drvSetHandle(param,theApp->insertIcon(icon));
}

inline QFont drvGetFont(void *param)
{
    QFont *font = (QFont*)drvGetNative(param);
    if (!font) {
        return QFont();
    }
    return *font;
}

inline void drvSetFont(void *param, const QFont &font)
{
    drvSetHandle(param,theApp->insertFont(font));
}

inline void drvSetAction(void *param, QAction *act)
{
    drvSetHandle(param,act);
}

inline void drvSetMenu(void *param, QMenu *menu)
{
    drvSetHandle(param,menu);
}

inline void drvSetMenuBar(void *param, QMenuBar *bar)
{
    drvSetHandle(param,bar);
}

inline void drvSetStatusBar(void *param, QStatusBar *bar)
{
    drvSetHandle(param,bar);
}

inline QStatusBar* drvGetStatusBar(void *param)
{
    return (QStatusBar*)drvGetNative(param);
}

inline QToolBar* drvGetToolBar(void *param)
{
    return (QToolBar*)drvGetNative(param);
}

inline QDockWidget* drvGetDockWidget(void *param)
{
    return (QDockWidget*)drvGetNative(param);
}

inline QListWidgetItem* drvGetListWidgetItem(void *param)
{
    return (QListWidgetItem*)drvGetNative(param);
}
inline void drvSetListWidgetItem(void *param, QListWidgetItem* item)
{
    drvSetHandle(param,item);
}

inline QPixmap drvGetPixmap(void *param)
{
    QPixmap *pixmap = (QPixmap*)drvGetNative(param);
    if (!pixmap) {
        return QPixmap();
    }
    return *pixmap;
}

inline void drvSetPixmap(void *param, const QPixmap &pixmap)
{
    drvSetHandle(param,theApp->insertPixmap(pixmap));
}

inline void drvSetPixmap(void *param, const QPixmap *pixmap)
{
    if (pixmap == 0) {
        return;
    }
    drvSetHandle(param,theApp->insertPixmap(*pixmap));
}

inline QImage drvGetImage(void *param)
{
    QImage *image = (QImage*)drvGetNative(param);
    if (!image) {
        return QImage();
    }
    return *image;
}

inline void drvSetImage(void *param, const QImage &image)
{
    drvSetHandle(param,theApp->insertImage(image));
}

inline void drvSetImage(void *param, const QImage *image)
{
    if (image == 0) {
        return;
    }
    drvSetHandle(param,theApp->insertImage(*image));
}

inline QPen drvGetPen(void *param)
{
    QPen *pen = (QPen*)drvGetNative(param);
    if (!pen) {
        return QPen();
    }
    return *pen;
}

inline void drvSetPen(void *param, const QPen &pen)
{
    drvSetHandle(param,theApp->insertPen(pen));
}

inline QBrush drvGetBrush(void *param)
{
    QBrush *brush = (QBrush*)drvGetNative(param);
    if (!brush) {
        return QBrush();
    }
    return *brush;
}

inline void drvSetBrush(void *param, const QBrush &brush)
{
    drvSetHandle(param,theApp->insertBrush(brush));
}

inline void drvNewObj(void *a0, QObject *obj, bool reg = true)
{
    handle_head *head =(handle_head*)a0;
    head->native = obj;
    if (reg) {
        QObject::connect(obj,SIGNAL(destroyed(QObject*)),theApp,SLOT(deleteObject(QObject*)));
    }
}

inline void drvNewObj(void *a0, QWidget *obj, bool reg = true, bool delete_close = true)
{
    handle_head *head =(handle_head*)a0;
    head->native = obj;
    if (delete_close) {
        obj->setAttribute(Qt::WA_DeleteOnClose);
    }
    if (reg) {
        QObject::connect(obj,SIGNAL(destroyed(QObject*)),((QtApp*)qApp),SLOT(deleteObject(QObject*)));
    }
}

inline void drvNewObj(void *a0, void *obj)
{
    handle_head *head =(handle_head*)a0;
    head->native = obj;
}

template <typename T>
inline void drvDelObj(void *a0, T *obj)
{
	if (obj != 0) {
		delete obj;
		obj = 0;
	}

    if (a0 == 0) {
        return;
    }
    
    handle_head *head =(handle_head*)a0;
    if (head->native == 0 || head->native == drvInvalid) {
        return;
    }
    head->native = 0;
}

inline void drvDelFont(void *a0, QFont *font)
{
    theApp->removeFont(font);
    drvDelObj(a0,font);
}

inline void drvDelIcon(void *a0, QIcon *icon)
{
    theApp->removeIcon(icon);
    drvDelObj(a0,icon);
}

inline void drvDelPixmap(void *a0, QPixmap *pixmap)
{
    theApp->removePixmap(pixmap);
    drvDelObj(a0,pixmap);
}

inline void drvDelImage(void *a0, QImage *image)
{
    theApp->removeImage(image);
    drvDelObj(a0,image);
}

inline void drvDelPen(void *a0, QPen *pen)
{
    theApp->removePen(pen);
    drvDelObj(a0,pen);
}

inline void drvDelBrush(void *a0, QBrush *brush)
{
    theApp->removeBrush(brush);
    drvDelObj(a0,brush);
}

inline QtSignal* drvNewSignal(QObject *parent, void *fn, void *param = 0)
{
    QtSignal *s = new QtSignal(0,fn);
    s->moveToThread(parent->thread());
    s->setParent(parent);
    return s;
}

inline void drvNewEvent(int type, void *a0, void *a1, void *a2 = 0)
{
    handle_head* head = (handle_head*)a0;
    QObject *obj = (QObject*)head->native;
    QtApp *app = (QtApp*)qApp;
    app->eventLock.lock();
    QtEvent *ev = ((QtApp*)qApp)->eventMap.value(obj);
    if (ev == 0) {
        ev = new QtEvent;
        ev->moveToThread(obj->thread());
        ev->setParent(obj);
        obj->installEventFilter(ev);
        ((QtApp*)qApp)->eventMap.insert(obj,ev);
    }
    ev->setEvent(type,a1);
    app->eventLock.unlock();
}


int drv_callback(void* fn, void *a1,void* a2,void* a3,void* a4);
int drv_result(void* ch, int r);
int drv_appmain();

extern "C"
int QTDRVSHARED_EXPORT drv(int drvclass, int drvid, void *exp, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6);

int _drv(int drvclass, int drvid, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6);

inline Qt::AspectRatioMode drvGetAspectRatioMode(void *param)
{
	return (Qt::AspectRatioMode)(*(int*)param);
}

inline Qt::TransformationMode drvGetTransformationMode(void *param)
{
	return (Qt::TransformationMode)(*(int*)param);
}

inline QSizePolicy::Policy drvGetSizePolicyPolicy(void *param)
{
	if (param == 0) {
		return QSizePolicy::Fixed;
	}

	return (QSizePolicy::Policy)(*(int*)param);
}

inline void drvSetSizePolicyPolicy(void *param, QSizePolicy::Policy policy)
{
	if (param == 0) {
		return;
	}

	(*(int*)param) = (int)policy;
}

inline QSizePolicy::ControlType drvGetSizePolicyControlType(void *param)
{
	if (param == 0) {
		return QSizePolicy::DefaultType;
	}

	return (QSizePolicy::ControlType)(*(int*)param);
}

inline void drvSetSizePolicyControlType(void *param, QSizePolicy::ControlType control)
{
	if (param == 0) {
		return;
	}

	(*(int*)param) = (int)control;
}

inline void drvSetSizePolicy(void *param, const QSizePolicy &p) // FIXME
{
	QSizePolicy* policy = (QSizePolicy*)drvGetNative(param);
	if (policy == 0) {
		drvSetHandle(param, new QSizePolicy(p));
		return;
	}
	*policy = p;
}

inline QSizePolicy drvGetSizePolicy(void *param) // FIXME
{
	QSizePolicy* policy = (QSizePolicy*)drvGetNative(param);
	if (policy == 0) {
		return QSizePolicy();
	}

    return QSizePolicy(*policy);
}

inline void drvSetScrollBar(void *param, const QScrollBar *scrollbar)
{
    if (scrollbar == 0) {
        return;
    }
    drvSetHandle(param,(void*)scrollbar);
}
