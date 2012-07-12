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


#ifndef QTDRV_GLOBAL_H
#define QTDRV_GLOBAL_H

#include <QtCore/qglobal.h>

#if defined(QTDRV_LIBRARY)
#  define QTDRVSHARED_EXPORT Q_DECL_EXPORT
#else
#  define QTDRVSHARED_EXPORT Q_DECL_IMPORT
#endif

#endif // QTDRV_GLOBAL_H
