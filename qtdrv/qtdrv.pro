# Copyright 2012 visualfc <visualfc@gmail.com>. All rights reserved.
#
# This library is free software; you can redistribute it and/or
# modify it under the terms of the GNU Lesser General Public
# License as published by the Free Software Foundation; either
# version 2.1 of the License, or (at your option) any later version.
#
# This library is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
# Lesser General Public License for more details.

TARGET = qtdrv
TEMPLATE = lib

DEFINES += QTDRV_LIBRARY

SOURCES += cdrv.cpp \
    qtsignal.cpp \
    qtevent.cpp \
    qtapp.cpp

HEADERS += cdrv.h\
        qtdrv_global.h \
    qtsignal.h \
    qtevent.h \
    qtapp.h \
    qtshell.h \
	gldrv.h

IDE_BUILD_TREE = $$PWD
IDE_LIB_PATH = $$IDE_BUILD_TREE/../lib

DESTDIR = $$IDE_LIB_PATH

symbian {
    MMP_RULES += EXPORTUNFROZEN
    TARGET.UID3 = 0xE1B1CE15
    TARGET.CAPABILITY = 
    TARGET.EPOCALLOWDLLDATA = 1
    addFiles.sources = qtdrv.dll
    addFiles.path = !:/sys/bin
    DEPLOYMENT += addFiles
}

unix:!symbian {
    maemo5 {
        target.path = /opt/usr/lib
    } else {
        target.path = /usr/lib
    }
    INSTALLS += target
}

QT += opengl
