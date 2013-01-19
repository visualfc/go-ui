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

#ifndef QTSIGNAL_H
#define QTSIGNAL_H

#include <QObject>
#include <QAction>
#include <QListWidgetItem>

class QtSignal : public QObject
{
    Q_OBJECT
public:
    explicit QtSignal(QObject *parent,void *func);
    virtual ~QtSignal();
public slots:
    void call();
    void call(bool);
	void call(int);
    void call(int,int);
    void call(const QString&);
    void call(QAction*);
    void call(QListWidgetItem*);
    void call(QListWidgetItem*,QListWidgetItem*);
public:
    void *m_func;
};

#endif // QTSIGNAL_H
