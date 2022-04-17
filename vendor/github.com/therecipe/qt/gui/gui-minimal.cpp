// +build minimal

#define protected public
#define private public

#include "gui-minimal.h"
#include "_cgo_export.h"

#include <QAccessible>
#include <QBitmap>
#include <QByteArray>
#include <QDataStream>
#include <QFont>
#include <QFontDatabase>
#include <QGuiApplication>
#include <QIcon>
#include <QIconEngine>
#include <QKeySequence>
#include <QLine>
#include <QLineF>
#include <QList>
#include <QMatrix>
#include <QMatrix4x4>
#include <QObject>
#include <QPagedPaintDevice>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPaintEngineState>
#include <QPaintEvent>
#include <QPainter>
#include <QPainterPath>
#include <QPalette>
#include <QPixmap>
#include <QPoint>
#include <QPointF>
#include <QPolygon>
#include <QPolygonF>
#include <QRect>
#include <QRectF>
#include <QRegion>
#include <QScreen>
#include <QSize>
#include <QString>
#include <QTransform>
#include <QVariant>
#include <QVector>
#include <QVector2D>
#include <QVector3D>
#include <QVector4D>
#include <QWidget>
#include <QStringList>

int QAccessible_InvalidEvent_Type()
{
	return QAccessible::InvalidEvent;
}

class MyQBitmap: public QBitmap
{
public:
	MyQBitmap() : QBitmap() {QBitmap_QBitmap_QRegisterMetaType();};
	MyQBitmap(int width, int height) : QBitmap(width, height) {QBitmap_QBitmap_QRegisterMetaType();};
	MyQBitmap(const QBitmap &other) : QBitmap(other) {QBitmap_QBitmap_QRegisterMetaType();};
	MyQBitmap(const QPixmap &pixmap) : QBitmap(pixmap) {QBitmap_QBitmap_QRegisterMetaType();};
	MyQBitmap(const QSize &size) : QBitmap(size) {QBitmap_QBitmap_QRegisterMetaType();};
	MyQBitmap(const QString &fileName, const char *format = Q_NULLPTR) : QBitmap(fileName, format) {QBitmap_QBitmap_QRegisterMetaType();};
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPixmap_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QBitmap*)
Q_DECLARE_METATYPE(MyQBitmap*)

int QBitmap_QBitmap_QRegisterMetaType(){qRegisterMetaType<QBitmap*>(); return qRegisterMetaType<MyQBitmap*>();}

void* QBitmap_NewQBitmap()
{
	return new MyQBitmap();
}

void* QBitmap_NewQBitmap3(int width, int height)
{
	return new MyQBitmap(width, height);
}

void* QBitmap_NewQBitmap6(void* other)
{
	return new MyQBitmap(*static_cast<QBitmap*>(other));
}

void* QBitmap_NewQBitmap2(void* pixmap)
{
	return new MyQBitmap(*static_cast<QPixmap*>(pixmap));
}

void* QBitmap_NewQBitmap4(void* size)
{
	return new MyQBitmap(*static_cast<QSize*>(size));
}

void* QBitmap_NewQBitmap5(struct QtGui_PackedString fileName, char* format)
{
	return new MyQBitmap(QString::fromUtf8(fileName.data, fileName.len), const_cast<const char*>(format));
}

void QBitmap_DestroyQBitmap(void* ptr)
{
	static_cast<QBitmap*>(ptr)->~QBitmap();
}

int QFont_Times_Type()
{
	return QFont::Times;
}

int QFont_Courier_Type()
{
	return QFont::Courier;
}

int QFont_OldEnglish_Type()
{
	return QFont::OldEnglish;
}

int QFont_System_Type()
{
	return QFont::System;
}

int QFont_AnyStyle_Type()
{
	return QFont::AnyStyle;
}

int QFont_Cursive_Type()
{
	return QFont::Cursive;
}

int QFont_Monospace_Type()
{
	return QFont::Monospace;
}

int QFont_Fantasy_Type()
{
	return QFont::Fantasy;
}

int QFontDatabase_Ogham_Type()
{
	return QFontDatabase::Ogham;
}

int QFontDatabase_Runic_Type()
{
	return QFontDatabase::Runic;
}

int QFontDatabase_Nko_Type()
{
	return QFontDatabase::Nko;
}

int QFontDatabase_WritingSystemsCount_Type()
{
	return QFontDatabase::WritingSystemsCount;
}

class MyQGuiApplication: public QGuiApplication
{
public:
	MyQGuiApplication(int &argc, char **argv) : QGuiApplication(argc, argv) {QGuiApplication_QGuiApplication_QRegisterMetaType();};
	 ~MyQGuiApplication() { callbackQGuiApplication_DestroyQGuiApplication(this); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQGuiApplication_ObjectNameChanged(this, objectNamePacked); };
};

Q_DECLARE_METATYPE(QGuiApplication*)
Q_DECLARE_METATYPE(MyQGuiApplication*)

int QGuiApplication_QGuiApplication_QRegisterMetaType(){qRegisterMetaType<QGuiApplication*>(); return qRegisterMetaType<MyQGuiApplication*>();}

void* QGuiApplication_NewQGuiApplication(int argc, char* argv)
{
	static int argcs = argc;
	static char** argvs = static_cast<char**>(malloc(argcs * sizeof(char*)));

	QList<QByteArray> aList = QByteArray(argv).split('|');
	for (int i = 0; i < argcs; i++)
		argvs[i] = (new QByteArray(aList.at(i)))->data();

	return new MyQGuiApplication(argcs, argvs);
}

int QGuiApplication_QGuiApplication_Exec()
{
	return QGuiApplication::exec();
}

void QGuiApplication_DestroyQGuiApplication(void* ptr)
{
	static_cast<QGuiApplication*>(ptr)->~QGuiApplication();
}

void QGuiApplication_DestroyQGuiApplicationDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QGuiApplication___screens_atList(void* ptr, int i)
{
	return ({QScreen * tmp = static_cast<QList<QScreen *>*>(ptr)->at(i); if (i == static_cast<QList<QScreen *>*>(ptr)->size()-1) { static_cast<QList<QScreen *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGuiApplication___screens_setList(void* ptr, void* i)
{
	static_cast<QList<QScreen *>*>(ptr)->append(static_cast<QScreen*>(i));
}

void* QGuiApplication___screens_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QScreen *>();
}

void* QGuiApplication___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGuiApplication___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGuiApplication___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QGuiApplication___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGuiApplication___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGuiApplication___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGuiApplication___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGuiApplication___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGuiApplication___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGuiApplication___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGuiApplication___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGuiApplication___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGuiApplication___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGuiApplication___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGuiApplication___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

Q_DECLARE_METATYPE(QIcon)
Q_DECLARE_METATYPE(QIcon*)
void* QIcon_NewQIcon()
{
	return new QIcon();
}

void* QIcon_NewQIcon4(void* other)
{
	return new QIcon(*static_cast<QIcon*>(other));
}

void* QIcon_NewQIcon6(void* engine)
{
	return new QIcon(static_cast<QIconEngine*>(engine));
}

void* QIcon_NewQIcon3(void* other)
{
	return new QIcon(*static_cast<QIcon*>(other));
}

void* QIcon_NewQIcon2(void* pixmap)
{
	return new QIcon(*static_cast<QPixmap*>(pixmap));
}

void* QIcon_NewQIcon5(struct QtGui_PackedString fileName)
{
	return new QIcon(QString::fromUtf8(fileName.data, fileName.len));
}

void QIcon_DestroyQIcon(void* ptr)
{
	static_cast<QIcon*>(ptr)->~QIcon();
}

struct QtGui_PackedString QIcon_Name(void* ptr)
{
	return ({ QByteArray* t03700a = new QByteArray(static_cast<QIcon*>(ptr)->name().toUtf8()); QtGui_PackedString { const_cast<char*>(t03700a->prepend("WHITESPACE").constData()+10), t03700a->size()-10, t03700a }; });
}

void* QIcon_ToVariant(void* ptr)
{
	return new QVariant(*static_cast<QIcon*>(ptr));
}

void* QIcon___availableSizes_atList(void* ptr, int i)
{
	return ({ QSize tmpValue = ({QSize tmp = static_cast<QList<QSize>*>(ptr)->at(i); if (i == static_cast<QList<QSize>*>(ptr)->size()-1) { static_cast<QList<QSize>*>(ptr)->~QList(); free(ptr); }; tmp; }); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QIcon___availableSizes_setList(void* ptr, void* i)
{
	static_cast<QList<QSize>*>(ptr)->append(*static_cast<QSize*>(i));
}

void* QIcon___availableSizes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSize>();
}

class MyQIconEngine: public QIconEngine
{
public:
	MyQIconEngine() : QIconEngine() {QIconEngine_QIconEngine_QRegisterMetaType();};
	bool read(QDataStream & in) { return callbackQIconEngine_Read(this, static_cast<QDataStream*>(&in)) != 0; };
	void paint(QPainter * painter, const QRect & rect, QIcon::Mode mode, QIcon::State state) { callbackQIconEngine_Paint(this, painter, const_cast<QRect*>(&rect), mode, state); };
	 ~MyQIconEngine() { callbackQIconEngine_DestroyQIconEngine(this); };
	QIconEngine * clone() const { return static_cast<QIconEngine*>(callbackQIconEngine_Clone(const_cast<void*>(static_cast<const void*>(this)))); };
	QString key() const { return ({ QtGui_PackedString tempVal = callbackQIconEngine_Key(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool write(QDataStream & out) const { return callbackQIconEngine_Write(const_cast<void*>(static_cast<const void*>(this)), static_cast<QDataStream*>(&out)) != 0; };
};

Q_DECLARE_METATYPE(QIconEngine*)
Q_DECLARE_METATYPE(MyQIconEngine*)

int QIconEngine_QIconEngine_QRegisterMetaType(){qRegisterMetaType<QIconEngine*>(); return qRegisterMetaType<MyQIconEngine*>();}

void* QIconEngine_NewQIconEngine()
{
	return new MyQIconEngine();
}

char QIconEngine_Read(void* ptr, void* in)
{
	return static_cast<QIconEngine*>(ptr)->read(*static_cast<QDataStream*>(in));
}

char QIconEngine_ReadDefault(void* ptr, void* in)
{
		return static_cast<QIconEngine*>(ptr)->QIconEngine::read(*static_cast<QDataStream*>(in));
}

void QIconEngine_Paint(void* ptr, void* painter, void* rect, long long mode, long long state)
{
	static_cast<QIconEngine*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QRect*>(rect), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void QIconEngine_DestroyQIconEngine(void* ptr)
{
	static_cast<QIconEngine*>(ptr)->~QIconEngine();
}

void QIconEngine_DestroyQIconEngineDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QIconEngine_Clone(void* ptr)
{
	return static_cast<QIconEngine*>(ptr)->clone();
}

struct QtGui_PackedString QIconEngine_Key(void* ptr)
{
	return ({ QByteArray* tfa2543 = new QByteArray(static_cast<QIconEngine*>(ptr)->key().toUtf8()); QtGui_PackedString { const_cast<char*>(tfa2543->prepend("WHITESPACE").constData()+10), tfa2543->size()-10, tfa2543 }; });
}

struct QtGui_PackedString QIconEngine_KeyDefault(void* ptr)
{
		return ({ QByteArray* t9979b6 = new QByteArray(static_cast<QIconEngine*>(ptr)->QIconEngine::key().toUtf8()); QtGui_PackedString { const_cast<char*>(t9979b6->prepend("WHITESPACE").constData()+10), t9979b6->size()-10, t9979b6 }; });
}

char QIconEngine_Write(void* ptr, void* out)
{
	return static_cast<QIconEngine*>(ptr)->write(*static_cast<QDataStream*>(out));
}

char QIconEngine_WriteDefault(void* ptr, void* out)
{
		return static_cast<QIconEngine*>(ptr)->QIconEngine::write(*static_cast<QDataStream*>(out));
}

void* QIconEngine___availableSizes_atList(void* ptr, int i)
{
	return ({ QSize tmpValue = ({QSize tmp = static_cast<QList<QSize>*>(ptr)->at(i); if (i == static_cast<QList<QSize>*>(ptr)->size()-1) { static_cast<QList<QSize>*>(ptr)->~QList(); free(ptr); }; tmp; }); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QIconEngine___availableSizes_setList(void* ptr, void* i)
{
	static_cast<QList<QSize>*>(ptr)->append(*static_cast<QSize*>(i));
}

void* QIconEngine___availableSizes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSize>();
}

Q_DECLARE_METATYPE(QKeySequence)
Q_DECLARE_METATYPE(QKeySequence*)
void* QKeySequence_NewQKeySequence()
{
	return new QKeySequence();
}

void* QKeySequence_NewQKeySequence5(long long key)
{
	return new QKeySequence(static_cast<QKeySequence::StandardKey>(key));
}

void* QKeySequence_NewQKeySequence4(void* keysequence)
{
	return new QKeySequence(*static_cast<QKeySequence*>(keysequence));
}

void* QKeySequence_NewQKeySequence2(struct QtGui_PackedString key, long long format)
{
	return new QKeySequence(QString::fromUtf8(key.data, key.len), static_cast<QKeySequence::SequenceFormat>(format));
}

void* QKeySequence_NewQKeySequence3(int k1, int k2, int k3, int k4)
{
	return new QKeySequence(k1, k2, k3, k4);
}

void QKeySequence_DestroyQKeySequence(void* ptr)
{
	static_cast<QKeySequence*>(ptr)->~QKeySequence();
}

struct QtGui_PackedString QKeySequence_ToString(void* ptr, long long format)
{
	return ({ QByteArray* t3bf8a8 = new QByteArray(static_cast<QKeySequence*>(ptr)->toString(static_cast<QKeySequence::SequenceFormat>(format)).toUtf8()); QtGui_PackedString { const_cast<char*>(t3bf8a8->prepend("WHITESPACE").constData()+10), t3bf8a8->size()-10, t3bf8a8 }; });
}

void* QKeySequence___keyBindings_atList(void* ptr, int i)
{
	return new QKeySequence(({QKeySequence tmp = static_cast<QList<QKeySequence>*>(ptr)->at(i); if (i == static_cast<QList<QKeySequence>*>(ptr)->size()-1) { static_cast<QList<QKeySequence>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QKeySequence___keyBindings_setList(void* ptr, void* i)
{
	static_cast<QList<QKeySequence>*>(ptr)->append(*static_cast<QKeySequence*>(i));
}

void* QKeySequence___keyBindings_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QKeySequence>();
}

void* QKeySequence___listFromString_atList(void* ptr, int i)
{
	return new QKeySequence(({QKeySequence tmp = static_cast<QList<QKeySequence>*>(ptr)->at(i); if (i == static_cast<QList<QKeySequence>*>(ptr)->size()-1) { static_cast<QList<QKeySequence>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QKeySequence___listFromString_setList(void* ptr, void* i)
{
	static_cast<QList<QKeySequence>*>(ptr)->append(*static_cast<QKeySequence*>(i));
}

void* QKeySequence___listFromString_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QKeySequence>();
}

void* QKeySequence___listToString_list_atList(void* ptr, int i)
{
	return new QKeySequence(({QKeySequence tmp = static_cast<QList<QKeySequence>*>(ptr)->at(i); if (i == static_cast<QList<QKeySequence>*>(ptr)->size()-1) { static_cast<QList<QKeySequence>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QKeySequence___listToString_list_setList(void* ptr, void* i)
{
	static_cast<QList<QKeySequence>*>(ptr)->append(*static_cast<QKeySequence*>(i));
}

void* QKeySequence___listToString_list_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QKeySequence>();
}

Q_DECLARE_METATYPE(QMatrix)
Q_DECLARE_METATYPE(QMatrix*)
void* QMatrix_Translate(void* ptr, double dx, double dy)
{
	return new QMatrix(static_cast<QMatrix*>(ptr)->translate(dx, dy));
}

void* QMatrix_NewQMatrix()
{
	return new QMatrix();
}

void* QMatrix_NewQMatrix4(void* other)
{
	return new QMatrix(*static_cast<QMatrix*>(other));
}

void* QMatrix_NewQMatrix5(void* matrix)
{
	return new QMatrix(*static_cast<QMatrix*>(matrix));
}

void* QMatrix_NewQMatrix3(double m11, double m12, double m21, double m22, double dx, double dy)
{
	return new QMatrix(m11, m12, m21, m22, dx, dy);
}

void QMatrix_Reset(void* ptr)
{
	static_cast<QMatrix*>(ptr)->reset();
}

void* QMatrix_Map5(void* ptr, void* line)
{
	return ({ QLine tmpValue = static_cast<QMatrix*>(ptr)->map(*static_cast<QLine*>(line)); new QLine(tmpValue.p1(), tmpValue.p2()); });
}

void* QMatrix_Map6(void* ptr, void* line)
{
	return ({ QLineF tmpValue = static_cast<QMatrix*>(ptr)->map(*static_cast<QLineF*>(line)); new QLineF(tmpValue.p1(), tmpValue.p2()); });
}

void* QMatrix_Map10(void* ptr, void* path)
{
	return new QPainterPath(static_cast<QMatrix*>(ptr)->map(*static_cast<QPainterPath*>(path)));
}

void* QMatrix_Map3(void* ptr, void* point)
{
	return ({ QPoint tmpValue = static_cast<QMatrix*>(ptr)->map(*static_cast<QPoint*>(point)); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void* QMatrix_Map4(void* ptr, void* point)
{
	return ({ QPointF tmpValue = static_cast<QMatrix*>(ptr)->map(*static_cast<QPointF*>(point)); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void* QMatrix_Map8(void* ptr, void* polygon)
{
	return new QPolygon(static_cast<QMatrix*>(ptr)->map(*static_cast<QPolygon*>(polygon)));
}

void* QMatrix_Map7(void* ptr, void* polygon)
{
	return new QPolygonF(static_cast<QMatrix*>(ptr)->map(*static_cast<QPolygonF*>(polygon)));
}

void* QMatrix_Map9(void* ptr, void* region)
{
	return new QRegion(static_cast<QMatrix*>(ptr)->map(*static_cast<QRegion*>(region)));
}

void QMatrix_Map2(void* ptr, int x, int y, int tx, int ty)
{
	static_cast<QMatrix*>(ptr)->map(x, y, &tx, &ty);
}

void QMatrix_Map(void* ptr, double x, double y, double tx, double ty)
{
	static_cast<QMatrix*>(ptr)->map(x, y, &tx, &ty);
}

Q_DECLARE_METATYPE(QMatrix4x4)
Q_DECLARE_METATYPE(QMatrix4x4*)
void* QMatrix4x4_NewQMatrix4x4()
{
	return new QMatrix4x4();
}

void* QMatrix4x4_NewQMatrix4x48(void* matrix)
{
	return new QMatrix4x4(*static_cast<QMatrix*>(matrix));
}

void* QMatrix4x4_NewQMatrix4x47(void* transform)
{
	return new QMatrix4x4(*static_cast<QTransform*>(transform));
}

void* QMatrix4x4_NewQMatrix4x43(float values)
{
	return new QMatrix4x4(const_cast<const float*>(&values));
}

void* QMatrix4x4_NewQMatrix4x44(float m11, float m12, float m13, float m14, float m21, float m22, float m23, float m24, float m31, float m32, float m33, float m34, float m41, float m42, float m43, float m44)
{
	return new QMatrix4x4(m11, m12, m13, m14, m21, m22, m23, m24, m31, m32, m33, m34, m41, m42, m43, m44);
}

float QMatrix4x4_Data(void* ptr)
{
	return *static_cast<QMatrix4x4*>(ptr)->data();
}

void QMatrix4x4_Translate(void* ptr, void* vector)
{
	static_cast<QMatrix4x4*>(ptr)->translate(*static_cast<QVector3D*>(vector));
}

void QMatrix4x4_Translate2(void* ptr, float x, float y)
{
	static_cast<QMatrix4x4*>(ptr)->translate(x, y);
}

void QMatrix4x4_Translate3(void* ptr, float x, float y, float z)
{
	static_cast<QMatrix4x4*>(ptr)->translate(x, y, z);
}

void* QMatrix4x4_Map(void* ptr, void* point)
{
	return ({ QPoint tmpValue = static_cast<QMatrix4x4*>(ptr)->map(*static_cast<QPoint*>(point)); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void* QMatrix4x4_Map2(void* ptr, void* point)
{
	return ({ QPointF tmpValue = static_cast<QMatrix4x4*>(ptr)->map(*static_cast<QPointF*>(point)); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void* QMatrix4x4_Map3(void* ptr, void* point)
{
	return new QVector3D(static_cast<QMatrix4x4*>(ptr)->map(*static_cast<QVector3D*>(point)));
}

void* QMatrix4x4_Map4(void* ptr, void* point)
{
	return new QVector4D(static_cast<QMatrix4x4*>(ptr)->map(*static_cast<QVector4D*>(point)));
}

float QMatrix4x4_Data2(void* ptr)
{
	return *static_cast<QMatrix4x4*>(ptr)->data();
}

class MyQPagedPaintDevice: public QPagedPaintDevice
{
public:
	MyQPagedPaintDevice() : QPagedPaintDevice() {QPagedPaintDevice_QPagedPaintDevice_QRegisterMetaType();};
	bool newPage() { return callbackQPagedPaintDevice_NewPage(this) != 0; };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPagedPaintDevice_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QPagedPaintDevice*)
Q_DECLARE_METATYPE(MyQPagedPaintDevice*)

int QPagedPaintDevice_QPagedPaintDevice_QRegisterMetaType(){qRegisterMetaType<QPagedPaintDevice*>(); return qRegisterMetaType<MyQPagedPaintDevice*>();}

void* QPagedPaintDevice_NewQPagedPaintDevice()
{
	return new MyQPagedPaintDevice();
}

char QPagedPaintDevice_NewPage(void* ptr)
{
	return static_cast<QPagedPaintDevice*>(ptr)->newPage();
}

void QPagedPaintDevice_DestroyQPagedPaintDevice(void* ptr)
{
	static_cast<QPagedPaintDevice*>(ptr)->~QPagedPaintDevice();
}

void* QPagedPaintDevice_PaintEngine(void* ptr)
{
	return static_cast<QPagedPaintDevice*>(ptr)->paintEngine();
}

void* QPagedPaintDevice_PaintEngineDefault(void* ptr)
{
	Q_UNUSED(ptr);
	
}

class MyQPaintDevice: public QPaintDevice
{
public:
	MyQPaintDevice() : QPaintDevice() {QPaintDevice_QPaintDevice_QRegisterMetaType();};
	 ~MyQPaintDevice() { callbackQPaintDevice_DestroyQPaintDevice(this); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPaintDevice_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QPaintDevice*)
Q_DECLARE_METATYPE(MyQPaintDevice*)

int QPaintDevice_QPaintDevice_QRegisterMetaType(){qRegisterMetaType<QPaintDevice*>(); return qRegisterMetaType<MyQPaintDevice*>();}

void* QPaintDevice_NewQPaintDevice()
{
	return new MyQPaintDevice();
}

void QPaintDevice_DestroyQPaintDevice(void* ptr)
{
	static_cast<QPaintDevice*>(ptr)->~QPaintDevice();
}

void QPaintDevice_DestroyQPaintDeviceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QPaintDevice_PaintEngine(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->paintEngine();
}

int QPaintDevice_Depth(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->depth();
}

class MyQPaintEngine: public QPaintEngine
{
public:
	MyQPaintEngine(PaintEngineFeatures caps = PaintEngineFeatures()) : QPaintEngine(caps) {QPaintEngine_QPaintEngine_QRegisterMetaType();};
	bool end() { return callbackQPaintEngine_End(this) != 0; };
	bool begin(QPaintDevice * pdev) { return callbackQPaintEngine_Begin(this, pdev) != 0; };
	void drawPixmap(const QRectF & r, const QPixmap & pm, const QRectF & sr) { callbackQPaintEngine_DrawPixmap(this, const_cast<QRectF*>(&r), const_cast<QPixmap*>(&pm), const_cast<QRectF*>(&sr)); };
	void updateState(const QPaintEngineState & state) { callbackQPaintEngine_UpdateState(this, const_cast<QPaintEngineState*>(&state)); };
	 ~MyQPaintEngine() { callbackQPaintEngine_DestroyQPaintEngine(this); };
	Type type() const { return static_cast<QPaintEngine::Type>(callbackQPaintEngine_Type(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QPaintEngine*)
Q_DECLARE_METATYPE(MyQPaintEngine*)

int QPaintEngine_QPaintEngine_QRegisterMetaType(){qRegisterMetaType<QPaintEngine*>(); return qRegisterMetaType<MyQPaintEngine*>();}

char QPaintEngine_End(void* ptr)
{
	return static_cast<QPaintEngine*>(ptr)->end();
}

void* QPaintEngine_NewQPaintEngine(long long caps)
{
	return new MyQPaintEngine(static_cast<QPaintEngine::PaintEngineFeature>(caps));
}

char QPaintEngine_Begin(void* ptr, void* pdev)
{
	return static_cast<QPaintEngine*>(ptr)->begin(static_cast<QPaintDevice*>(pdev));
}

void QPaintEngine_DrawPixmap(void* ptr, void* r, void* pm, void* sr)
{
	static_cast<QPaintEngine*>(ptr)->drawPixmap(*static_cast<QRectF*>(r), *static_cast<QPixmap*>(pm), *static_cast<QRectF*>(sr));
}

void QPaintEngine_UpdateState(void* ptr, void* state)
{
	static_cast<QPaintEngine*>(ptr)->updateState(*static_cast<QPaintEngineState*>(state));
}

void QPaintEngine_DestroyQPaintEngine(void* ptr)
{
	static_cast<QPaintEngine*>(ptr)->~QPaintEngine();
}

void QPaintEngine_DestroyQPaintEngineDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QPaintEngine_Type(void* ptr)
{
	return static_cast<QPaintEngine*>(ptr)->type();
}

class MyQPaintEvent: public QPaintEvent
{
public:
	MyQPaintEvent(const QRect &paintRect) : QPaintEvent(paintRect) {QPaintEvent_QPaintEvent_QRegisterMetaType();};
	MyQPaintEvent(const QRegion &paintRegion) : QPaintEvent(paintRegion) {QPaintEvent_QPaintEvent_QRegisterMetaType();};
};

Q_DECLARE_METATYPE(QPaintEvent*)
Q_DECLARE_METATYPE(MyQPaintEvent*)

int QPaintEvent_QPaintEvent_QRegisterMetaType(){qRegisterMetaType<QPaintEvent*>(); return qRegisterMetaType<MyQPaintEvent*>();}

void* QPaintEvent_NewQPaintEvent2(void* paintRect)
{
	return new MyQPaintEvent(*static_cast<QRect*>(paintRect));
}

void* QPaintEvent_NewQPaintEvent(void* paintRegion)
{
	return new MyQPaintEvent(*static_cast<QRegion*>(paintRegion));
}

Q_DECLARE_METATYPE(QPainter*)
void* QPainter_NewQPainter2(void* device)
{
	if (dynamic_cast<QWidget*>(static_cast<QObject*>(device))) {
		return new QPainter(static_cast<QWidget*>(device));
	} else {
		return new QPainter(static_cast<QPaintDevice*>(device));
	}
}

void* QPainter_NewQPainter()
{
	return new QPainter();
}

void QPainter_SetWindow(void* ptr, void* rectangle)
{
	static_cast<QPainter*>(ptr)->setWindow(*static_cast<QRect*>(rectangle));
}

void QPainter_SetWindow2(void* ptr, int x, int y, int width, int height)
{
	static_cast<QPainter*>(ptr)->setWindow(x, y, width, height);
}

void QPainter_Translate2(void* ptr, void* offset)
{
	static_cast<QPainter*>(ptr)->translate(*static_cast<QPoint*>(offset));
}

void QPainter_Translate(void* ptr, void* offset)
{
	static_cast<QPainter*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

void QPainter_Translate3(void* ptr, double dx, double dy)
{
	static_cast<QPainter*>(ptr)->translate(dx, dy);
}

void QPainter_DestroyQPainter(void* ptr)
{
	static_cast<QPainter*>(ptr)->~QPainter();
}

void* QPainter_Window(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QPainter*>(ptr)->window(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QPainter___drawLines_lines_atList6(void* ptr, int i)
{
	return ({ QLine tmpValue = ({QLine tmp = static_cast<QVector<QLine>*>(ptr)->at(i); if (i == static_cast<QVector<QLine>*>(ptr)->size()-1) { static_cast<QVector<QLine>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QLine(tmpValue.p1(), tmpValue.p2()); });
}

void QPainter___drawLines_lines_setList6(void* ptr, void* i)
{
	static_cast<QVector<QLine>*>(ptr)->append(*static_cast<QLine*>(i));
}

void* QPainter___drawLines_lines_newList6(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QLine>();
}

void* QPainter___drawLines_lines_atList2(void* ptr, int i)
{
	return ({ QLineF tmpValue = ({QLineF tmp = static_cast<QVector<QLineF>*>(ptr)->at(i); if (i == static_cast<QVector<QLineF>*>(ptr)->size()-1) { static_cast<QVector<QLineF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QLineF(tmpValue.p1(), tmpValue.p2()); });
}

void QPainter___drawLines_lines_setList2(void* ptr, void* i)
{
	static_cast<QVector<QLineF>*>(ptr)->append(*static_cast<QLineF*>(i));
}

void* QPainter___drawLines_lines_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QLineF>();
}

void* QPainter___drawLines_pointPairs_atList8(void* ptr, int i)
{
	return ({ QPoint tmpValue = ({QPoint tmp = static_cast<QVector<QPoint>*>(ptr)->at(i); if (i == static_cast<QVector<QPoint>*>(ptr)->size()-1) { static_cast<QVector<QPoint>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void QPainter___drawLines_pointPairs_setList8(void* ptr, void* i)
{
	static_cast<QVector<QPoint>*>(ptr)->append(*static_cast<QPoint*>(i));
}

void* QPainter___drawLines_pointPairs_newList8(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPoint>();
}

void* QPainter___drawLines_pointPairs_atList4(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QVector<QPointF>*>(ptr)->at(i); if (i == static_cast<QVector<QPointF>*>(ptr)->size()-1) { static_cast<QVector<QPointF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QPainter___drawLines_pointPairs_setList4(void* ptr, void* i)
{
	static_cast<QVector<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QPainter___drawLines_pointPairs_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPointF>();
}

void* QPainter___drawRects_rectangles_atList4(void* ptr, int i)
{
	return ({ QRect tmpValue = ({QRect tmp = static_cast<QVector<QRect>*>(ptr)->at(i); if (i == static_cast<QVector<QRect>*>(ptr)->size()-1) { static_cast<QVector<QRect>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QPainter___drawRects_rectangles_setList4(void* ptr, void* i)
{
	static_cast<QVector<QRect>*>(ptr)->append(*static_cast<QRect*>(i));
}

void* QPainter___drawRects_rectangles_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QRect>();
}

void* QPainter___drawRects_rectangles_atList2(void* ptr, int i)
{
	return ({ QRectF tmpValue = ({QRectF tmp = static_cast<QVector<QRectF>*>(ptr)->at(i); if (i == static_cast<QVector<QRectF>*>(ptr)->size()-1) { static_cast<QVector<QRectF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QPainter___drawRects_rectangles_setList2(void* ptr, void* i)
{
	static_cast<QVector<QRectF>*>(ptr)->append(*static_cast<QRectF*>(i));
}

void* QPainter___drawRects_rectangles_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QRectF>();
}

Q_DECLARE_METATYPE(QPainterPath*)
void* QPainterPath_NewQPainterPath3(void* path)
{
	return new QPainterPath(*static_cast<QPainterPath*>(path));
}

void* QPainterPath_NewQPainterPath()
{
	return new QPainterPath();
}

void* QPainterPath_NewQPainterPath2(void* startPoint)
{
	return new QPainterPath(*static_cast<QPointF*>(startPoint));
}

void QPainterPath_Translate2(void* ptr, void* offset)
{
	static_cast<QPainterPath*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

void QPainterPath_Translate(void* ptr, double dx, double dy)
{
	static_cast<QPainterPath*>(ptr)->translate(dx, dy);
}

void QPainterPath_DestroyQPainterPath(void* ptr)
{
	static_cast<QPainterPath*>(ptr)->~QPainterPath();
}

char QPainterPath_Contains3(void* ptr, void* p)
{
	return static_cast<QPainterPath*>(ptr)->contains(*static_cast<QPainterPath*>(p));
}

char QPainterPath_Contains(void* ptr, void* point)
{
	return static_cast<QPainterPath*>(ptr)->contains(*static_cast<QPointF*>(point));
}

char QPainterPath_Contains2(void* ptr, void* rectangle)
{
	return static_cast<QPainterPath*>(ptr)->contains(*static_cast<QRectF*>(rectangle));
}

double QPainterPath_Length(void* ptr)
{
	return static_cast<QPainterPath*>(ptr)->length();
}

void* QPainterPath___toFillPolygons_atList2(void* ptr, int i)
{
	return new QPolygonF(({QPolygonF tmp = static_cast<QList<QPolygonF>*>(ptr)->at(i); if (i == static_cast<QList<QPolygonF>*>(ptr)->size()-1) { static_cast<QList<QPolygonF>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QPainterPath___toFillPolygons_setList2(void* ptr, void* i)
{
	static_cast<QList<QPolygonF>*>(ptr)->append(*static_cast<QPolygonF*>(i));
}

void* QPainterPath___toFillPolygons_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPolygonF>();
}

void* QPainterPath___toFillPolygons_atList(void* ptr, int i)
{
	return new QPolygonF(({QPolygonF tmp = static_cast<QList<QPolygonF>*>(ptr)->at(i); if (i == static_cast<QList<QPolygonF>*>(ptr)->size()-1) { static_cast<QList<QPolygonF>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QPainterPath___toFillPolygons_setList(void* ptr, void* i)
{
	static_cast<QList<QPolygonF>*>(ptr)->append(*static_cast<QPolygonF*>(i));
}

void* QPainterPath___toFillPolygons_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPolygonF>();
}

void* QPainterPath___toSubpathPolygons_atList2(void* ptr, int i)
{
	return new QPolygonF(({QPolygonF tmp = static_cast<QList<QPolygonF>*>(ptr)->at(i); if (i == static_cast<QList<QPolygonF>*>(ptr)->size()-1) { static_cast<QList<QPolygonF>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QPainterPath___toSubpathPolygons_setList2(void* ptr, void* i)
{
	static_cast<QList<QPolygonF>*>(ptr)->append(*static_cast<QPolygonF*>(i));
}

void* QPainterPath___toSubpathPolygons_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPolygonF>();
}

void* QPainterPath___toSubpathPolygons_atList(void* ptr, int i)
{
	return new QPolygonF(({QPolygonF tmp = static_cast<QList<QPolygonF>*>(ptr)->at(i); if (i == static_cast<QList<QPolygonF>*>(ptr)->size()-1) { static_cast<QList<QPolygonF>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QPainterPath___toSubpathPolygons_setList(void* ptr, void* i)
{
	static_cast<QList<QPolygonF>*>(ptr)->append(*static_cast<QPolygonF*>(i));
}

void* QPainterPath___toSubpathPolygons_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPolygonF>();
}

int QPalette_NColorRoles_Type()
{
	return QPalette::NColorRoles;
}

class MyQPixmap: public QPixmap
{
public:
	MyQPixmap() : QPixmap() {QPixmap_QPixmap_QRegisterMetaType();};
	MyQPixmap(const QPixmap &pixmap) : QPixmap(pixmap) {QPixmap_QPixmap_QRegisterMetaType();};
	MyQPixmap(const QSize &size) : QPixmap(size) {QPixmap_QPixmap_QRegisterMetaType();};
	MyQPixmap(const QString &fileName, const char *format = Q_NULLPTR, Qt::ImageConversionFlags flags = Qt::AutoColor) : QPixmap(fileName, format, flags) {QPixmap_QPixmap_QRegisterMetaType();};
	MyQPixmap(int width, int height) : QPixmap(width, height) {QPixmap_QPixmap_QRegisterMetaType();};
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPixmap_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QPixmap*)
Q_DECLARE_METATYPE(MyQPixmap*)

int QPixmap_QPixmap_QRegisterMetaType(){qRegisterMetaType<QPixmap*>(); return qRegisterMetaType<MyQPixmap*>();}

void* QPixmap_NewQPixmap()
{
	return new MyQPixmap();
}

void* QPixmap_NewQPixmap7(void* pixmap)
{
	return new MyQPixmap(*static_cast<QPixmap*>(pixmap));
}

void* QPixmap_NewQPixmap4(void* size)
{
	return new MyQPixmap(*static_cast<QSize*>(size));
}

void* QPixmap_NewQPixmap5(struct QtGui_PackedString fileName, char* format, long long flags)
{
	return new MyQPixmap(QString::fromUtf8(fileName.data, fileName.len), const_cast<const char*>(format), static_cast<Qt::ImageConversionFlag>(flags));
}

void* QPixmap_NewQPixmap3(int width, int height)
{
	return new MyQPixmap(width, height);
}

char QPixmap_Load(void* ptr, struct QtGui_PackedString fileName, char* format, long long flags)
{
	return static_cast<QPixmap*>(ptr)->load(QString::fromUtf8(fileName.data, fileName.len), const_cast<const char*>(format), static_cast<Qt::ImageConversionFlag>(flags));
}

void QPixmap_DestroyQPixmap(void* ptr)
{
	static_cast<QPixmap*>(ptr)->~QPixmap();
}

void* QPixmap_Copy(void* ptr, void* rectangle)
{
	return new QPixmap(static_cast<QPixmap*>(ptr)->copy(*static_cast<QRect*>(rectangle)));
}

void* QPixmap_Copy2(void* ptr, int x, int y, int width, int height)
{
	return new QPixmap(static_cast<QPixmap*>(ptr)->copy(x, y, width, height));
}

void* QPixmap_Size(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QPixmap*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QPixmap_ToVariant(void* ptr)
{
	return new QVariant(*static_cast<QPixmap*>(ptr));
}

void* QPixmap_PaintEngine(void* ptr)
{
	return static_cast<QPixmap*>(ptr)->paintEngine();
}

void* QPixmap_PaintEngineDefault(void* ptr)
{
	if (dynamic_cast<QBitmap*>(static_cast<QPixmap*>(ptr))) {
		return static_cast<QBitmap*>(ptr)->QBitmap::paintEngine();
	} else {
		return static_cast<QPixmap*>(ptr)->QPixmap::paintEngine();
	}
}

Q_DECLARE_METATYPE(QPolygon)
Q_DECLARE_METATYPE(QPolygon*)
void* QPolygon_NewQPolygon()
{
	return new QPolygon();
}

void* QPolygon_NewQPolygon8(void* other)
{
	return new QPolygon(*static_cast<QPolygon*>(other));
}

void* QPolygon_NewQPolygon4(void* v)
{
	return new QPolygon(({ QVector<QPoint>* tmpP = static_cast<QVector<QPoint>*>(v); QVector<QPoint> tmpV = *tmpP; tmpP->~QVector(); free(tmpP); tmpV; }));
}

void* QPolygon_NewQPolygon7(void* polygon)
{
	return new QPolygon(*static_cast<QPolygon*>(polygon));
}

void* QPolygon_NewQPolygon5(void* rectangle, char closed)
{
	return new QPolygon(*static_cast<QRect*>(rectangle), closed != 0);
}

void* QPolygon_NewQPolygon3(void* points)
{
	return new QPolygon(*static_cast<QVector<QPoint>*>(points));
}

void* QPolygon_NewQPolygon2(int size)
{
	return new QPolygon(size);
}

void QPolygon_SetPoint2(void* ptr, int index, void* point)
{
	static_cast<QPolygon*>(ptr)->setPoint(index, *static_cast<QPoint*>(point));
}

void QPolygon_SetPoint(void* ptr, int index, int x, int y)
{
	static_cast<QPolygon*>(ptr)->setPoint(index, x, y);
}

void QPolygon_Translate2(void* ptr, void* offset)
{
	static_cast<QPolygon*>(ptr)->translate(*static_cast<QPoint*>(offset));
}

void QPolygon_Translate(void* ptr, int dx, int dy)
{
	static_cast<QPolygon*>(ptr)->translate(dx, dy);
}

void QPolygon_DestroyQPolygon(void* ptr)
{
	static_cast<QPolygon*>(ptr)->~QPolygon();
}

void* QPolygon_Point2(void* ptr, int index)
{
	return ({ QPoint tmpValue = static_cast<QPolygon*>(ptr)->point(index); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void QPolygon_Point(void* ptr, int index, int x, int y)
{
	static_cast<QPolygon*>(ptr)->point(index, &x, &y);
}

void* QPolygon___QPolygon_v_atList4(void* ptr, int i)
{
	return ({ QPoint tmpValue = ({QPoint tmp = static_cast<QVector<QPoint>*>(ptr)->at(i); if (i == static_cast<QVector<QPoint>*>(ptr)->size()-1) { static_cast<QVector<QPoint>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void QPolygon___QPolygon_v_setList4(void* ptr, void* i)
{
	static_cast<QVector<QPoint>*>(ptr)->append(*static_cast<QPoint*>(i));
}

void* QPolygon___QPolygon_v_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPoint>();
}

void* QPolygon___QPolygon_points_atList3(void* ptr, int i)
{
	return ({ QPoint tmpValue = ({QPoint tmp = static_cast<QVector<QPoint>*>(ptr)->at(i); if (i == static_cast<QVector<QPoint>*>(ptr)->size()-1) { static_cast<QVector<QPoint>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void QPolygon___QPolygon_points_setList3(void* ptr, void* i)
{
	static_cast<QVector<QPoint>*>(ptr)->append(*static_cast<QPoint*>(i));
}

void* QPolygon___QPolygon_points_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPoint>();
}

void* QPolygon___QVector_other_atList5(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___QVector_other_setList5(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygon___QVector_other_newList5(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___QVector_other_atList4(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___QVector_other_setList4(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygon___QVector_other_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___fill_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___fill_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygon___fill_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___fromList_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___fromList_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygon___fromList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___fromList_list_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___fromList_list_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygon___fromList_list_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___fromStdVector_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___fromStdVector_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygon___fromStdVector_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___append_value_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___append_value_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygon___append_value_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___swap_other_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___swap_other_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygon___swap_other_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___toList_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___toList_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygon___toList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___mid_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___mid_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygon___mid_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

Q_DECLARE_METATYPE(QPolygonF)
Q_DECLARE_METATYPE(QPolygonF*)
void* QPolygonF_NewQPolygonF()
{
	return new QPolygonF();
}

void* QPolygonF_NewQPolygonF8(void* other)
{
	return new QPolygonF(*static_cast<QPolygonF*>(other));
}

void* QPolygonF_NewQPolygonF4(void* v)
{
	return new QPolygonF(({ QVector<QPointF>* tmpP = static_cast<QVector<QPointF>*>(v); QVector<QPointF> tmpV = *tmpP; tmpP->~QVector(); free(tmpP); tmpV; }));
}

void* QPolygonF_NewQPolygonF6(void* polygon)
{
	return new QPolygonF(*static_cast<QPolygon*>(polygon));
}

void* QPolygonF_NewQPolygonF7(void* polygon)
{
	return new QPolygonF(*static_cast<QPolygonF*>(polygon));
}

void* QPolygonF_NewQPolygonF5(void* rectangle)
{
	return new QPolygonF(*static_cast<QRectF*>(rectangle));
}

void* QPolygonF_NewQPolygonF3(void* points)
{
	return new QPolygonF(*static_cast<QVector<QPointF>*>(points));
}

void* QPolygonF_NewQPolygonF2(int size)
{
	return new QPolygonF(size);
}

void QPolygonF_Translate(void* ptr, void* offset)
{
	static_cast<QPolygonF*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

void QPolygonF_Translate2(void* ptr, double dx, double dy)
{
	static_cast<QPolygonF*>(ptr)->translate(dx, dy);
}

void QPolygonF_DestroyQPolygonF(void* ptr)
{
	static_cast<QPolygonF*>(ptr)->~QPolygonF();
}

void* QPolygonF___QPolygonF_v_atList4(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QVector<QPointF>*>(ptr)->at(i); if (i == static_cast<QVector<QPointF>*>(ptr)->size()-1) { static_cast<QVector<QPointF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QPolygonF___QPolygonF_v_setList4(void* ptr, void* i)
{
	static_cast<QVector<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QPolygonF___QPolygonF_v_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPointF>();
}

void* QPolygonF___QPolygonF_points_atList3(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QVector<QPointF>*>(ptr)->at(i); if (i == static_cast<QVector<QPointF>*>(ptr)->size()-1) { static_cast<QVector<QPointF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QPolygonF___QPolygonF_points_setList3(void* ptr, void* i)
{
	static_cast<QVector<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QPolygonF___QPolygonF_points_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPointF>();
}

void* QPolygonF___QVector_other_atList5(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___QVector_other_setList5(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygonF___QVector_other_newList5(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygonF___QVector_other_atList4(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___QVector_other_setList4(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygonF___QVector_other_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygonF___fill_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___fill_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygonF___fill_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygonF___fromList_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___fromList_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygonF___fromList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygonF___fromList_list_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___fromList_list_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygonF___fromList_list_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygonF___fromStdVector_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___fromStdVector_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygonF___fromStdVector_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygonF___append_value_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___append_value_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygonF___append_value_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygonF___swap_other_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___swap_other_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygonF___swap_other_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygonF___toList_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___toList_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygonF___toList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygonF___mid_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___mid_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPolygonF___mid_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

Q_DECLARE_METATYPE(QRegion)
Q_DECLARE_METATYPE(QRegion*)
void* QRegion_NewQRegion()
{
	return new QRegion();
}

void* QRegion_NewQRegion6(void* other)
{
	return new QRegion(*static_cast<QRegion*>(other));
}

void* QRegion_NewQRegion7(void* bm)
{
	return new QRegion(*static_cast<QBitmap*>(bm));
}

void* QRegion_NewQRegion4(void* a, long long fillRule)
{
	return new QRegion(*static_cast<QPolygon*>(a), static_cast<Qt::FillRule>(fillRule));
}

void* QRegion_NewQRegion3(void* r, long long t)
{
	return new QRegion(*static_cast<QRect*>(r), static_cast<QRegion::RegionType>(t));
}

void* QRegion_NewQRegion5(void* r)
{
	return new QRegion(*static_cast<QRegion*>(r));
}

void* QRegion_NewQRegion2(int x, int y, int w, int h, long long t)
{
	return new QRegion(x, y, w, h, static_cast<QRegion::RegionType>(t));
}

void QRegion_Translate2(void* ptr, void* point)
{
	static_cast<QRegion*>(ptr)->translate(*static_cast<QPoint*>(point));
}

void QRegion_Translate(void* ptr, int dx, int dy)
{
	static_cast<QRegion*>(ptr)->translate(dx, dy);
}

char QRegion_Contains(void* ptr, void* p)
{
	return static_cast<QRegion*>(ptr)->contains(*static_cast<QPoint*>(p));
}

char QRegion_Contains2(void* ptr, void* r)
{
	return static_cast<QRegion*>(ptr)->contains(*static_cast<QRect*>(r));
}

void* QRegion___rects_atList(void* ptr, int i)
{
	return ({ QRect tmpValue = ({QRect tmp = static_cast<QVector<QRect>*>(ptr)->at(i); if (i == static_cast<QVector<QRect>*>(ptr)->size()-1) { static_cast<QVector<QRect>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QRegion___rects_setList(void* ptr, void* i)
{
	static_cast<QVector<QRect>*>(ptr)->append(*static_cast<QRect*>(i));
}

void* QRegion___rects_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QRect>();
}

class MyQScreen: public QScreen
{
public:
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQScreen_ObjectNameChanged(this, objectNamePacked); };
};

Q_DECLARE_METATYPE(QScreen*)
Q_DECLARE_METATYPE(MyQScreen*)

int QScreen_QScreen_QRegisterMetaType(){qRegisterMetaType<QScreen*>(); return qRegisterMetaType<MyQScreen*>();}

void QScreen_DestroyQScreen(void* ptr)
{
	static_cast<QScreen*>(ptr)->~QScreen();
}

void* QScreen_Size(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QScreen*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct QtGui_PackedString QScreen_Name(void* ptr)
{
	return ({ QByteArray* tc60f02 = new QByteArray(static_cast<QScreen*>(ptr)->name().toUtf8()); QtGui_PackedString { const_cast<char*>(tc60f02->prepend("WHITESPACE").constData()+10), tc60f02->size()-10, tc60f02 }; });
}

int QScreen_Depth(void* ptr)
{
	return static_cast<QScreen*>(ptr)->depth();
}

void* QScreen___virtualSiblings_atList(void* ptr, int i)
{
	return ({QScreen * tmp = static_cast<QList<QScreen *>*>(ptr)->at(i); if (i == static_cast<QList<QScreen *>*>(ptr)->size()-1) { static_cast<QList<QScreen *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QScreen___virtualSiblings_setList(void* ptr, void* i)
{
	static_cast<QList<QScreen *>*>(ptr)->append(static_cast<QScreen*>(i));
}

void* QScreen___virtualSiblings_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QScreen *>();
}

void* QScreen___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QScreen___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QScreen___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QScreen___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QScreen___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScreen___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QScreen___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QScreen___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScreen___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QScreen___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QScreen___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScreen___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QScreen___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QScreen___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QScreen___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

Q_DECLARE_METATYPE(QTransform)
Q_DECLARE_METATYPE(QTransform*)
void* QTransform_Translate(void* ptr, double dx, double dy)
{
	return new QTransform(static_cast<QTransform*>(ptr)->translate(dx, dy));
}

void* QTransform_NewQTransform()
{
	return new QTransform();
}

void* QTransform_NewQTransform6(void* other)
{
	return new QTransform(*static_cast<QTransform*>(other));
}

void* QTransform_NewQTransform5(void* matrix)
{
	return new QTransform(*static_cast<QMatrix*>(matrix));
}

void* QTransform_NewQTransform7(void* other)
{
	return new QTransform(*static_cast<QTransform*>(other));
}

void* QTransform_NewQTransform3(double m11, double m12, double m13, double m21, double m22, double m23, double m31, double m32, double m33)
{
	return new QTransform(m11, m12, m13, m21, m22, m23, m31, m32, m33);
}

void* QTransform_NewQTransform4(double m11, double m12, double m21, double m22, double dx, double dy)
{
	return new QTransform(m11, m12, m21, m22, dx, dy);
}

void QTransform_Reset(void* ptr)
{
	static_cast<QTransform*>(ptr)->reset();
}

void* QTransform_Map3(void* ptr, void* l)
{
	return ({ QLine tmpValue = static_cast<QTransform*>(ptr)->map(*static_cast<QLine*>(l)); new QLine(tmpValue.p1(), tmpValue.p2()); });
}

void* QTransform_Map4(void* ptr, void* line)
{
	return ({ QLineF tmpValue = static_cast<QTransform*>(ptr)->map(*static_cast<QLineF*>(line)); new QLineF(tmpValue.p1(), tmpValue.p2()); });
}

void* QTransform_Map8(void* ptr, void* path)
{
	return new QPainterPath(static_cast<QTransform*>(ptr)->map(*static_cast<QPainterPath*>(path)));
}

void* QTransform_Map10(void* ptr, void* point)
{
	return ({ QPoint tmpValue = static_cast<QTransform*>(ptr)->map(*static_cast<QPoint*>(point)); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void* QTransform_Map2(void* ptr, void* p)
{
	return ({ QPointF tmpValue = static_cast<QTransform*>(ptr)->map(*static_cast<QPointF*>(p)); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void* QTransform_Map6(void* ptr, void* polygon)
{
	return new QPolygon(static_cast<QTransform*>(ptr)->map(*static_cast<QPolygon*>(polygon)));
}

void* QTransform_Map5(void* ptr, void* polygon)
{
	return new QPolygonF(static_cast<QTransform*>(ptr)->map(*static_cast<QPolygonF*>(polygon)));
}

void* QTransform_Map7(void* ptr, void* region)
{
	return new QRegion(static_cast<QTransform*>(ptr)->map(*static_cast<QRegion*>(region)));
}

long long QTransform_Type(void* ptr)
{
	return static_cast<QTransform*>(ptr)->type();
}

void QTransform_Map9(void* ptr, int x, int y, int tx, int ty)
{
	static_cast<QTransform*>(ptr)->map(x, y, &tx, &ty);
}

void QTransform_Map(void* ptr, double x, double y, double tx, double ty)
{
	static_cast<QTransform*>(ptr)->map(x, y, &tx, &ty);
}

Q_DECLARE_METATYPE(QVector2D)
Q_DECLARE_METATYPE(QVector2D*)
void* QVector2D_NewQVector2D()
{
	return new QVector2D();
}

void* QVector2D_NewQVector2D4(void* point)
{
	return new QVector2D(*static_cast<QPoint*>(point));
}

void* QVector2D_NewQVector2D5(void* point)
{
	return new QVector2D(*static_cast<QPointF*>(point));
}

void* QVector2D_NewQVector2D6(void* vector)
{
	return new QVector2D(*static_cast<QVector3D*>(vector));
}

void* QVector2D_NewQVector2D7(void* vector)
{
	return new QVector2D(*static_cast<QVector4D*>(vector));
}

void* QVector2D_NewQVector2D3(float xpos, float ypos)
{
	return new QVector2D(xpos, ypos);
}

float QVector2D_Length(void* ptr)
{
	return static_cast<QVector2D*>(ptr)->length();
}

float QVector2D_X(void* ptr)
{
	return static_cast<QVector2D*>(ptr)->x();
}

float QVector2D_Y(void* ptr)
{
	return static_cast<QVector2D*>(ptr)->y();
}

Q_DECLARE_METATYPE(QVector3D)
Q_DECLARE_METATYPE(QVector3D*)
void* QVector3D_NewQVector3D()
{
	return new QVector3D();
}

void* QVector3D_NewQVector3D4(void* point)
{
	return new QVector3D(*static_cast<QPoint*>(point));
}

void* QVector3D_NewQVector3D5(void* point)
{
	return new QVector3D(*static_cast<QPointF*>(point));
}

void* QVector3D_NewQVector3D6(void* vector)
{
	return new QVector3D(*static_cast<QVector2D*>(vector));
}

void* QVector3D_NewQVector3D7(void* vector, float zpos)
{
	return new QVector3D(*static_cast<QVector2D*>(vector), zpos);
}

void* QVector3D_NewQVector3D8(void* vector)
{
	return new QVector3D(*static_cast<QVector4D*>(vector));
}

void* QVector3D_NewQVector3D3(float xpos, float ypos, float zpos)
{
	return new QVector3D(xpos, ypos, zpos);
}

float QVector3D_Length(void* ptr)
{
	return static_cast<QVector3D*>(ptr)->length();
}

float QVector3D_X(void* ptr)
{
	return static_cast<QVector3D*>(ptr)->x();
}

float QVector3D_Y(void* ptr)
{
	return static_cast<QVector3D*>(ptr)->y();
}

float QVector3D_Z(void* ptr)
{
	return static_cast<QVector3D*>(ptr)->z();
}

Q_DECLARE_METATYPE(QVector4D)
Q_DECLARE_METATYPE(QVector4D*)
void* QVector4D_NewQVector4D()
{
	return new QVector4D();
}

void* QVector4D_NewQVector4D4(void* point)
{
	return new QVector4D(*static_cast<QPoint*>(point));
}

void* QVector4D_NewQVector4D5(void* point)
{
	return new QVector4D(*static_cast<QPointF*>(point));
}

void* QVector4D_NewQVector4D6(void* vector)
{
	return new QVector4D(*static_cast<QVector2D*>(vector));
}

void* QVector4D_NewQVector4D7(void* vector, float zpos, float wpos)
{
	return new QVector4D(*static_cast<QVector2D*>(vector), zpos, wpos);
}

void* QVector4D_NewQVector4D8(void* vector)
{
	return new QVector4D(*static_cast<QVector3D*>(vector));
}

void* QVector4D_NewQVector4D9(void* vector, float wpos)
{
	return new QVector4D(*static_cast<QVector3D*>(vector), wpos);
}

void* QVector4D_NewQVector4D3(float xpos, float ypos, float zpos, float wpos)
{
	return new QVector4D(xpos, ypos, zpos, wpos);
}

void QVector4D_SetW(void* ptr, float w)
{
	static_cast<QVector4D*>(ptr)->setW(w);
}

float QVector4D_Length(void* ptr)
{
	return static_cast<QVector4D*>(ptr)->length();
}

float QVector4D_W(void* ptr)
{
	return static_cast<QVector4D*>(ptr)->w();
}

float QVector4D_X(void* ptr)
{
	return static_cast<QVector4D*>(ptr)->x();
}

float QVector4D_Y(void* ptr)
{
	return static_cast<QVector4D*>(ptr)->y();
}

float QVector4D_Z(void* ptr)
{
	return static_cast<QVector4D*>(ptr)->z();
}

