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

#include "qtevent.h"
#include "cdrv.h"
#include <QEvent>
#include <QResizeEvent>
#include <QMoveEvent>
#include <QTimerEvent>
#include <QMouseEvent>
#include <QHoverEvent>
#include <QPointF>
#include <QKeyEvent>
#include <QPaintEvent>
#include <QFocusEvent>
#include <QDebug>

struct base_event {
    int accept;
};

struct timer_event {
    int accept;
    int timerid;
};

struct move_event {
    int accept;
    int x;
    int y;
    int ox;
    int oy;
};

struct hover_event {
    int accept;
    int x;
    int y;
    int ox;
    int oy;
};

struct resize_event {
    int accept;
    int w;
    int h;
    int ow;
    int oh;
};

struct mouse_event {
    int accept;
    int modify;
    int button;
    int buttons;
    int gx;
    int gy;
    int x;
    int y;
};

struct key_event {
    int accept;
    int modify;
    int count;
    int autorepeat;
    int key;
    quint32 nativeModifiers;
    quint32 nativeScanCode;
    quint32 nativeVirtualKey;
    string_head *sh;
};

struct paint_event {
    int accept;
    int x;
    int y;
    int w;
    int h;
};

struct focus_event {
    int accept;
    int reason;
};

QtEvent::QtEvent(QObject *parent) :
    QObject(parent)
{
}

void QtEvent::setEvent(int type, void *func)
{
    m_hash.insert(type,func);
}

bool QtEvent::eventFilter(QObject *target, QEvent *event)
{
    int type = event->type();
    void *func = m_hash.value(type);
    if (func != 0) {
        int accept = event->isAccepted() ? 1 : 0;
        switch(type) {
        case QEvent::Create:
        case QEvent::Close:
        case QEvent::Show:
        case QEvent::Hide:
        case QEvent::Enter:
        case QEvent::Leave: {
            base_event ev = {accept};
            drv_callback(func,&ev,0,0,0);
            event->setAccepted(ev.accept != 0);
            break;
        }
        case QEvent::FocusIn:
        case QEvent::FocusOut: {
            QFocusEvent *e = (QFocusEvent*)event;
            focus_event ev = {accept,e->reason()};
            drv_callback(func,&ev,0,0,0);
            event->setAccepted(ev.accept != 0);
            break;
        }
        case QEvent::Timer: {
            QTimerEvent *e = (QTimerEvent*)event;
            timer_event ev = {accept,e->timerId()};
            drv_callback(func,&ev,0,0,0);
            event->setAccepted(ev.accept != 0);
            break;
        }
        case QEvent::HoverEnter:
        case QEvent::HoverLeave:
        case QEvent::HoverMove: {
            QHoverEvent *e = (QHoverEvent*)event;
            const QPoint &pt = e->pos();
            const QPoint &opt = e->oldPos();
            hover_event ev = {accept,pt.x(),pt.y(),opt.x(),opt.y()};
            drv_callback(func,&ev,0,0,0);
            event->setAccepted(ev.accept != 0);
            break;
        }
        case QEvent::KeyPress:
        case QEvent::KeyRelease: {
            QKeyEvent *e = (QKeyEvent*)event;
            string_head sh;
            drvSetString(&sh,e->text());
            key_event ev = {accept,e->modifiers(),e->count(),e->isAutoRepeat()?1:0,e->key(),e->nativeModifiers(),e->nativeScanCode(),e->nativeVirtualKey(),&sh};
            drv_callback(func,&ev,0,0,0);
            event->setAccepted(ev.accept != 0);
            break;
        }
        case QEvent::MouseButtonPress:
        case QEvent::MouseButtonRelease:
        case QEvent::MouseButtonDblClick:
        case QEvent::MouseMove: {
            QMouseEvent *e = (QMouseEvent*)event;
            const QPoint &gpt = e->globalPos();
            const QPoint &pt = e->pos();
            mouse_event ev = {accept,e->modifiers(),e->button(),e->buttons(),gpt.x(),gpt.y(),pt.x(),pt.y()};
            drv_callback(func,&ev,0,0,0);
            event->setAccepted(ev.accept != 0);
            break;
        }
        case QEvent::Move: {
            QMoveEvent *e = (QMoveEvent*)event;
            const QPoint &pt = e->pos();
            const QPoint &opt = e->oldPos();
            move_event ev = {accept,pt.x(),pt.y(),opt.x(),opt.y()};
            drv_callback(func,&ev,0,0,0);
            event->setAccepted(ev.accept != 0);
            break;
        }
        case QEvent::Resize: {
            QResizeEvent *e = (QResizeEvent*)event;
            const QSize &sz = e->size();
            const QSize &osz = e->oldSize();
            resize_event ev = {accept,sz.width(),sz.height(),osz.width(),osz.height()};
            drv_callback(func,&ev,0,0,0);
            event->setAccepted(ev.accept != 0);
            break;
        }
        case QEvent::Paint: {
            QPaintEvent *e = (QPaintEvent*)event;
            const QRect &rc = e->rect();
            paint_event ev = {accept,rc.x(),rc.y(),rc.width(),rc.height()};
            drv_callback(func,&ev,0,0,0);
            event->setAccepted(ev.accept != 0);
            break;
        }
        default: {
            return QObject::eventFilter(target,event);
        }
        }
        return true;
    }
    return QObject::eventFilter(target,event);
}
