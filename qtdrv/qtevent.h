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

#ifndef QTEVENT_H
#define QTEVENT_H

#include <QObject>
#include <QHash>

class QtEvent : public QObject
{
    Q_OBJECT
public:
    explicit QtEvent(QObject *parent = 0);
    void setEvent(int type, void *func);
    virtual bool eventFilter(QObject *target, QEvent *event);
protected:
    QHash<int,void*> m_hash;
};

#endif // QTEVENT_H
