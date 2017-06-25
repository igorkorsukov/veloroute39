TEMPLATE = app

QT += qml quick quickcontrols2 location
CONFIG += c++11

TARGET = veloroute39

OBJ_PATH = $$PWD/../../obj
LIB_PATH = $$PWD/../../lib
BIN_PATH = $$PWD/../../bin

OBJECTS_DIR = $$OBJ_PATH/$$TARGET
MOC_DIR = $$OBJECTS_DIR
OBJMOC = $$OBJECTS_DIR
RCC_DIR = $$OBJECTS_DIR

DESTDIR = $$BIN_PATH


include($$PWD/qgotoc/qgotoc.pri)


#LIBS += \
#    -L$$LIB_PATH \



INCLUDEPATH += .\

#DEPENDPATH += \
#    $$PWD/../infra \
#    $$PWD/../ui \
#    $$PWD/../map \

#PRE_TARGETDEPS += \
#    $$LIB_PATH/libinfra.a \
#    $$LIB_PATH/libui.a \
#    $$LIB_PATH/libmap.a \


RESOURCES += qml.qrc

SOURCES += main.cpp






# The following define makes your compiler emit warnings if you use
# any feature of Qt which as been marked deprecated (the exact warnings
# depend on your compiler). Please consult the documentation of the
# deprecated API in order to know how to port your code away from it.
DEFINES += QT_DEPRECATED_WARNINGS


# Default rules for deployment.
qnx: target.path = /tmp/$${TARGET}/bin
else: unix:!android: target.path = /opt/$${TARGET}/bin
!isEmpty(target.path): INSTALLS += target

DISTFILES += \
    android/AndroidManifest.xml \
    android/gradle/wrapper/gradle-wrapper.jar \
    android/gradlew \
    android/res/values/libs.xml \
    android/build.gradle \
    android/gradle/wrapper/gradle-wrapper.properties \
    android/gradlew.bat \
    qtquickcontrols2.conf \
    qml/main.qml \
    qml/MapView.qml

ANDROID_PACKAGE_SOURCE_DIR = $$PWD/../android
