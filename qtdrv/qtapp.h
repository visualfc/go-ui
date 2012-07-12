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


#ifndef QTAPP_H
#define QTAPP_H

#include <QApplication>
#include <QVariant>
#include <QMap>
#include <QFont>
#include <QMutex>
#include "qtevent.h"

struct drv_param {
    int classid;
    int drvid;
    void *a0;
    void *a1;
    void *a2;
    void *a3;
    void *a4;
    void *a5;
    void *a6;
    void *a7;
    void *a8;
    void *a9;
};

Q_DECLARE_METATYPE(drv_param)
Q_DECLARE_METATYPE(drv_param*)

class QtApp : public QApplication
{
    Q_OBJECT
public:
    explicit QtApp(int argc = 0, char **argv = NULL);
public slots:
    int drv(int drvclass, int drvid, void *ch, void *a0, void* a1, void* a2, void* a3, void* a4, void* a5, void* a6);
    void deleteObject(QObject *obj);
public:
    void deleteObj(void *obj);
public:
    void* insertFont(const QFont &font);
    void removeFont(QFont *font);
    void* insertIcon(const QIcon &icon);
    void removeIcon(QIcon *icon);
    void* insertPixmap(const QPixmap &pix);
    void removePixmap(QPixmap *pix);
    void* insertImage(const QImage &img);
    void removeImage(QImage *img);
    void* insertPen(const QPen &v);
    void removePen(QPen *v);
    void* insertBrush(const QBrush &v);
    void removeBrush(QBrush *v);
protected:
    QMap<void*,QFont*> fonts;
    QMap<void*,QIcon*> icons;
    QMap<void*,QPixmap*> pixs;
    QMap<void*,QImage*> imgs;
    QMap<void*,QPen*> pens;
    QMap<void*,QBrush*> brus;
public:
    void *pfnDeleteObj;
    QMap<QObject*,QtEvent*> eventMap;
    QMutex eventLock;
};

extern QtApp *theApp;


#endif // QTAPP_H
