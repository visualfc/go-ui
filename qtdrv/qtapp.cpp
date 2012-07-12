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

#define _private private
#define private public
#include <QFont>
#include <QPixmap>
#include <QIcon>
#include <QImage>
#include <QPen>
#include <QBrush>
static void* drvGetFontKey(const QFont *font)
{
    return font->d.data();
}
void* drvGetPixmapKey(const QPixmap *pixmap)
{
    return pixmap->pixmapData();
}
void* drvGetImageKey(const QImage *image)
{
    return image->d;
}
void* drvGetIconKey(const QIcon *icon)
{
    return icon->d;
}
void* drvGetPenKey(const QPen *pen)
{
    return pen->d;
}
void* drvGetBrushKey(const QBrush *brush)
{
    return brush->d.data();
}
#undef private
#define private _private

#include "qtapp.h"
#include "cdrv.h"

QtApp *theApp = 0;

QtApp::QtApp(int argc, char **argv) :
    QApplication(argc,argv), pfnDeleteObj(0)
{
    theApp = this;
}

int QtApp::drv(int drvclass, int drvid, void *ch, void *a0, void *a1, void *a2, void *a3, void *a4, void *a5, void *a6)
{
    int r = _drv(drvclass,drvid,a0,a1,a2,a3,a4,a5,a6);
    drv_result(ch,r);
    return 0;
}

void QtApp::deleteObject(QObject *obj)
{
    drv_callback(pfnDeleteObj,obj,0,0,0);
}

void QtApp::deleteObj(void *obj)
{
    drv_callback(pfnDeleteObj,obj,0,0,0);
}

void* QtApp::insertFont(const QFont &font)
{
    void *key = drvGetFontKey(&font);
    QFont *p = fonts.value(key);
    if (p == 0) {
        p = new QFont(font);
        fonts.insert(key,p);
    }
    return p;
}

void QtApp::removeFont(QFont *font)
{
    void *key = drvGetFontKey(font);
    if (fonts.remove(key)) {
        drv_callback(pfnDeleteObj,font,0,0,0);
    }
}

void* QtApp::insertIcon(const QIcon &icon)
{
    void *key = drvGetIconKey(&icon);
    QIcon *p = icons.value(key);
    if (p == 0) {
        p = new QIcon(icon);
        icons.insert(key,p);
    }
    return p;
}

void QtApp::removeIcon(QIcon *icon)
{
    void *key = drvGetIconKey(icon);
    if (icons.remove(key)) {
        drv_callback(pfnDeleteObj,icon,0,0,0);
    }
}

void* QtApp::insertPixmap(const QPixmap &pix)
{
    void *key = drvGetPixmapKey(&pix);
    QPixmap *p = pixs.value(key);
    if (p == 0) {
        p = new QPixmap(pix);
        pixs.insert(key,p);
    }
    return p;
}

void QtApp::removePixmap(QPixmap *pix)
{
    void *key = drvGetPixmapKey(pix);
    if (pixs.remove(key)) {
        drv_callback(pfnDeleteObj,pix,0,0,0);
    }
}

void* QtApp::insertImage(const QImage &img)
{
    void *key = drvGetImageKey(&img);
    QImage *p = imgs.value(key);
    if (p == 0) {
        p = new QImage(img);
        imgs.insert(key,p);
    }
    return p;
}

void QtApp::removeImage(QImage *img)
{
    void *key = drvGetImageKey(img);
    if (imgs.remove(key)) {
        drv_callback(pfnDeleteObj,img,0,0,0);
    }
}

void* QtApp::insertPen(const QPen &v)
{
    void *key = drvGetPenKey(&v);
    QPen *p = pens.value(key);
    if (p == 0) {
        p = new QPen(v);
        pens.insert(key,p);
    }
    return p;
}

void QtApp::removePen(QPen *v)
{
    void *key = drvGetPenKey(v);
    if (pens.remove(key)) {
        drv_callback(pfnDeleteObj,v,0,0,0);
    }
}

void* QtApp::insertBrush(const QBrush &v)
{
    void *key = drvGetBrushKey(&v);
    QBrush *p = brus.value(key);
    if (p == 0) {
        p = new QBrush(v);
        brus.insert(key,p);
    }
    return p;
}

void QtApp::removeBrush(QBrush *v)
{
    void *key = drvGetBrushKey(v);
    if (brus.remove(key)) {
        drv_callback(pfnDeleteObj,v,0,0,0);
    }
}

