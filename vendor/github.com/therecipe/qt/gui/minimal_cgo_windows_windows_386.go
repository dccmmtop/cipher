// +build minimal

package gui

/*
#cgo CFLAGS: -fno-keep-inline-dllexport -pipe -O2 -Wextra -Wall -W -DUNICODE -DQT_NEEDS_QMAIN -DQT_NO_DEBUG -DQT_WIDGETS_LIB -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -fno-keep-inline-dllexport -pipe -O2 -std=gnu++11 -Wextra -Wall -W -fexceptions -mthreads -DUNICODE -DQT_NEEDS_QMAIN -DQT_NO_DEBUG -DQT_WIDGETS_LIB -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -IF:/code/go/Qt/cipher -I. -ID:/Qt5.9.1/5.9.1/mingw53_32/include -ID:/Qt5.9.1/5.9.1/mingw53_32/include/QtWidgets -ID:/Qt5.9.1/5.9.1/mingw53_32/include/QtGui -ID:/Qt5.9.1/5.9.1/mingw53_32/include/QtANGLE -ID:/Qt5.9.1/5.9.1/mingw53_32/include/QtCore -Irelease -ID:/Qt5.9.1/5.9.1/mingw53_32/mkspecs/win32-g++
#cgo LDFLAGS:        -Wl,-subsystem,windows -mthreads -L D:/Qt5.9.1/5.9.1/mingw53_32/lib
#cgo LDFLAGS:        -lmingw32 -LD:/Qt5.9.1/5.9.1/mingw53_32/lib -lqtmain -LC:/utils/my_sql/my_sql/lib -LC:/utils/postgresql/pgsql/lib -lshell32 -lQt5Core -lQt5Widgets -lQt5Gui -lQt5Core
#cgo LDFLAGS: -Wl,--allow-multiple-definition
#cgo CFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
#cgo CXXFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
*/
import "C"
