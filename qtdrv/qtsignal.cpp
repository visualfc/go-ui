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

#include "qtsignal.h"
#include "cdrv.h"
#include <QDebug>

QtSignal::QtSignal(QObject *parent, void *func) :
    QObject(parent), m_func(func)
{
}

QtSignal::~QtSignal()
{
}

void QtSignal::call()
{
    drv_callback(m_func,0,0,0,0);
}

void QtSignal::call(bool b)
{
    int i = b ? 1:0;
    drv_callback(m_func,&i,0,0,0);
}

void QtSignal::call(int i)
{
    drv_callback(m_func,&i,0,0,0);
}

void QtSignal::call(int i, int j)
{
    drv_callback(m_func,&i,&j,0,0);
}

void QtSignal::call(const QString& s)
{
    string_head sh;
    drvSetString(&sh,s);
    drv_callback(m_func,&sh,0,0,0);
}

void QtSignal::call(QAction *act)
{
    drv_callback(m_func,act,0,0,0);
}

void QtSignal::call(QListWidgetItem* item)
{
    drv_callback(m_func,item,0,0,0);
}

void QtSignal::call(QListWidgetItem* item,QListWidgetItem* olditem)
{
    drv_callback(m_func,item,olditem,0,0);
}
