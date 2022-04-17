// +build minimal

#define protected public
#define private public

#include "core-minimal.h"
#include "_cgo_export.h"

#ifndef QT_CORE_LIB
	#error ------------------------------------------------------------------
	#error please run: '$(go env GOPATH)/bin/qtsetup'
	#error more info here: https://github.com/therecipe/qt/wiki/Installation
	#error ------------------------------------------------------------------
#endif
#include <QBitArray>
#include <QByteArray>
#include <QChar>
#include <QChildEvent>
#include <QCoreApplication>
#include <QDataStream>
#include <QDate>
#include <QDateTime>
#include <QEasingCurve>
#include <QEvent>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHash>
#include <QIODevice>
#include <QJsonArray>
#include <QJsonDocument>
#include <QJsonObject>
#include <QJsonValue>
#include <QLatin1Char>
#include <QLatin1String>
#include <QLayout>
#include <QLine>
#include <QLineF>
#include <QList>
#include <QLocale>
#include <QMap>
#include <QMargins>
#include <QMetaMethod>
#include <QMetaObject>
#include <QModelIndex>
#include <QObject>
#include <QPersistentModelIndex>
#include <QPoint>
#include <QPointF>
#include <QRect>
#include <QRectF>
#include <QRegExp>
#include <QRegularExpression>
#include <QRegularExpressionMatch>
#include <QSize>
#include <QSizeF>
#include <QString>
#include <QStringRef>
#include <QSysInfo>
#include <QTime>
#include <QTimeZone>
#include <QTimerEvent>
#include <QUrl>
#include <QUuid>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QTextDocument>
#include <QObject>
#include <QStringList>

Q_DECLARE_METATYPE(QBitArray)
Q_DECLARE_METATYPE(QBitArray*)
void* QBitArray_NewQBitArray()
{
	return new QBitArray();
}

void* QBitArray_NewQBitArray4(void* other)
{
	return new QBitArray(*static_cast<QBitArray*>(other));
}

void* QBitArray_NewQBitArray3(void* other)
{
	return new QBitArray(*static_cast<QBitArray*>(other));
}

void* QBitArray_NewQBitArray2(int size, char value)
{
	return new QBitArray(size, value != 0);
}

int QBitArray_Size(void* ptr)
{
	return static_cast<QBitArray*>(ptr)->size();
}

Q_DECLARE_METATYPE(QByteArray)
Q_DECLARE_METATYPE(QByteArray*)
void* QByteArray_Remove(void* ptr, int pos, int l)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->remove(pos, l));
}

void* QByteArray_Replace11(void* ptr, char* before, char* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(*before, *after));
}

void* QByteArray_Replace5(void* ptr, char* before, void* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(*before, *static_cast<QByteArray*>(after)));
}

void* QByteArray_Replace13(void* ptr, char* before, struct QtCore_PackedString after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(*before, QString::fromUtf8(after.data, after.len)));
}

void* QByteArray_Replace4(void* ptr, char* before, char* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(*before, const_cast<const char*>(after)));
}

void* QByteArray_Replace8(void* ptr, void* before, void* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(*static_cast<QByteArray*>(before), *static_cast<QByteArray*>(after)));
}

void* QByteArray_Replace9(void* ptr, void* before, char* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(*static_cast<QByteArray*>(before), const_cast<const char*>(after)));
}

void* QByteArray_Replace14(void* ptr, struct QtCore_PackedString before, void* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(QString::fromUtf8(before.data, before.len), *static_cast<QByteArray*>(after)));
}

void* QByteArray_Replace12(void* ptr, struct QtCore_PackedString before, char* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(QString::fromUtf8(before.data, before.len), const_cast<const char*>(after)));
}

void* QByteArray_Replace10(void* ptr, char* before, void* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(const_cast<const char*>(before), *static_cast<QByteArray*>(after)));
}

void* QByteArray_Replace6(void* ptr, char* before, char* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(const_cast<const char*>(before), const_cast<const char*>(after)));
}

void* QByteArray_Replace7(void* ptr, char* before, int bsize, char* after, int asize)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(const_cast<const char*>(before), bsize, const_cast<const char*>(after), asize));
}

void* QByteArray_Replace(void* ptr, int pos, int l, void* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(pos, l, *static_cast<QByteArray*>(after)));
}

void* QByteArray_Replace3(void* ptr, int pos, int l, char* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(pos, l, const_cast<const char*>(after)));
}

void* QByteArray_Replace2(void* ptr, int pos, int l, char* after, int alen)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(pos, l, const_cast<const char*>(after), alen));
}

void* QByteArray_NewQByteArray()
{
	return new QByteArray();
}

void* QByteArray_NewQByteArray6(void* other)
{
	return new QByteArray(*static_cast<QByteArray*>(other));
}

void* QByteArray_NewQByteArray5(void* other)
{
	return new QByteArray(*static_cast<QByteArray*>(other));
}

void* QByteArray_NewQByteArray2(char* data, int size)
{
	return new QByteArray(const_cast<const char*>(data), size);
}

void* QByteArray_NewQByteArray3(int size, char* ch)
{
	return new QByteArray(size, *ch);
}

struct QtCore_PackedString QByteArray_Data(void* ptr)
{
	return QtCore_PackedString { static_cast<QByteArray*>(ptr)->data(), static_cast<QByteArray*>(ptr)->size(), NULL };
}

void QByteArray_DestroyQByteArray(void* ptr)
{
	static_cast<QByteArray*>(ptr)->~QByteArray();
}

void* QByteArray_ToLower(void* ptr)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->toLower());
}

void* QByteArray_ToUpper(void* ptr)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->toUpper());
}

struct QtCore_PackedList QByteArray_Split(void* ptr, char* sep)
{
	return ({ QList<QByteArray>* tmpValue17cac8 = new QList<QByteArray>(static_cast<QByteArray*>(ptr)->split(*sep)); QtCore_PackedList { tmpValue17cac8, tmpValue17cac8->size() }; });
}

char QByteArray_Contains3(void* ptr, char* ch)
{
	return static_cast<QByteArray*>(ptr)->contains(*ch);
}

char QByteArray_Contains(void* ptr, void* ba)
{
	return static_cast<QByteArray*>(ptr)->contains(*static_cast<QByteArray*>(ba));
}

char QByteArray_Contains2(void* ptr, char* str)
{
	return static_cast<QByteArray*>(ptr)->contains(const_cast<const char*>(str));
}

struct QtCore_PackedString QByteArray_Data2(void* ptr)
{
	return QtCore_PackedString { const_cast<char*>(static_cast<QByteArray*>(ptr)->data()), static_cast<QByteArray*>(ptr)->size(), NULL };
}

int QByteArray_Length(void* ptr)
{
	return static_cast<QByteArray*>(ptr)->length();
}

int QByteArray_Size(void* ptr)
{
	return static_cast<QByteArray*>(ptr)->size();
}

void* QByteArray___split_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QByteArray___split_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QByteArray___split_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

Q_DECLARE_METATYPE(QChar)
Q_DECLARE_METATYPE(QChar*)
void* QChar_NewQChar()
{
	return new QChar();
}

void* QChar_NewQChar8(void* ch)
{
	return new QChar(*static_cast<QLatin1Char*>(ch));
}

void* QChar_NewQChar7(long long ch)
{
	return new QChar(static_cast<QChar::SpecialCharacter>(ch));
}

void* QChar_NewQChar9(char* ch)
{
	return new QChar(*ch);
}

void* QChar_NewQChar6(int code)
{
	return new QChar(code);
}

void* QChar_NewQChar4(short code)
{
	return new QChar(code);
}

void* QChar_NewQChar3(char* cell, char* row)
{
	return new QChar(*static_cast<uchar*>(static_cast<void*>(cell)), *static_cast<uchar*>(static_cast<void*>(row)));
}

void* QChar_NewQChar10(char* ch)
{
	return new QChar(*static_cast<uchar*>(static_cast<void*>(ch)));
}

void* QChar_NewQChar5(unsigned int code)
{
	return new QChar(code);
}

void* QChar_NewQChar2(unsigned short code)
{
	return new QChar(code);
}

unsigned int QChar_QChar_ToLower2(unsigned int ucs4)
{
	return QChar::toLower(ucs4);
}

unsigned int QChar_QChar_ToUpper2(unsigned int ucs4)
{
	return QChar::toUpper(ucs4);
}

void* QChar_ToLower(void* ptr)
{
	return new QChar(static_cast<QChar*>(ptr)->toLower());
}

void* QChar_ToUpper(void* ptr)
{
	return new QChar(static_cast<QChar*>(ptr)->toUpper());
}

class MyQChildEvent: public QChildEvent
{
public:
	MyQChildEvent(Type ty, QObject *child) : QChildEvent(ty, child) {QChildEvent_QChildEvent_QRegisterMetaType();};
};

Q_DECLARE_METATYPE(QChildEvent*)
Q_DECLARE_METATYPE(MyQChildEvent*)

int QChildEvent_QChildEvent_QRegisterMetaType(){qRegisterMetaType<QChildEvent*>(); return qRegisterMetaType<MyQChildEvent*>();}

void* QChildEvent_NewQChildEvent(long long ty, void* child)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(child))) {
		return new MyQChildEvent(static_cast<QEvent::Type>(ty), static_cast<QGraphicsObject*>(child));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(child))) {
		return new MyQChildEvent(static_cast<QEvent::Type>(ty), static_cast<QGraphicsWidget*>(child));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(child))) {
		return new MyQChildEvent(static_cast<QEvent::Type>(ty), static_cast<QLayout*>(child));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(child))) {
		return new MyQChildEvent(static_cast<QEvent::Type>(ty), static_cast<QWidget*>(child));
	} else {
		return new MyQChildEvent(static_cast<QEvent::Type>(ty), static_cast<QObject*>(child));
	}
}

void* QChildEvent_C(void* ptr)
{
	return static_cast<QChildEvent*>(ptr)->c;
}

void QChildEvent_SetC(void* ptr, void* vqo)
{
	static_cast<QChildEvent*>(ptr)->c = static_cast<QObject*>(vqo);
}

class MyQCoreApplication: public QCoreApplication
{
public:
	MyQCoreApplication(int &argc, char **argv) : QCoreApplication(argc, argv) {QCoreApplication_QCoreApplication_QRegisterMetaType();};
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
};

Q_DECLARE_METATYPE(QCoreApplication*)
Q_DECLARE_METATYPE(MyQCoreApplication*)

int QCoreApplication_QCoreApplication_QRegisterMetaType(){qRegisterMetaType<QCoreApplication*>(); return qRegisterMetaType<MyQCoreApplication*>();}

void* QCoreApplication_NewQCoreApplication(int argc, char* argv)
{
	static int argcs = argc;
	static char** argvs = static_cast<char**>(malloc(argcs * sizeof(char*)));

	QList<QByteArray> aList = QByteArray(argv).split('|');
	for (int i = 0; i < argcs; i++)
		argvs[i] = (new QByteArray(aList.at(i)))->data();

	return new MyQCoreApplication(argcs, argvs);
}

struct QtCore_PackedString QCoreApplication_QCoreApplication_Translate(char* context, char* sourceText, char* disambiguation, int n)
{
	return ({ QByteArray* t901eb0 = new QByteArray(QCoreApplication::translate(const_cast<const char*>(context), const_cast<const char*>(sourceText), const_cast<const char*>(disambiguation), n).toUtf8()); QtCore_PackedString { const_cast<char*>(t901eb0->prepend("WHITESPACE").constData()+10), t901eb0->size()-10, t901eb0 }; });
}

int QCoreApplication_QCoreApplication_Exec()
{
	return QCoreApplication::exec();
}

void QCoreApplication_QCoreApplication_Exit(int returnCode)
{
	QCoreApplication::exit(returnCode);
}

void QCoreApplication_DestroyQCoreApplication(void* ptr)
{
	static_cast<QCoreApplication*>(ptr)->~QCoreApplication();
}

Q_DECLARE_METATYPE(QDataStream*)
void* QDataStream_NewQDataStream()
{
	return new QDataStream();
}

void* QDataStream_NewQDataStream3(void* a, long long mode)
{
	return new QDataStream(static_cast<QByteArray*>(a), static_cast<QIODevice::OpenModeFlag>(mode));
}

void* QDataStream_NewQDataStream2(void* d)
{
	return new QDataStream(static_cast<QIODevice*>(d));
}

void* QDataStream_NewQDataStream4(void* a)
{
	return new QDataStream(*static_cast<QByteArray*>(a));
}

void QDataStream_DestroyQDataStream(void* ptr)
{
	static_cast<QDataStream*>(ptr)->~QDataStream();
}

Q_DECLARE_METATYPE(QDate)
Q_DECLARE_METATYPE(QDate*)
void* QDate_NewQDate()
{
	return new QDate();
}

void* QDate_NewQDate3(int y, int m, int d)
{
	return new QDate(y, m, d);
}

struct QtCore_PackedString QDate_ToString2(void* ptr, long long format)
{
	return ({ QByteArray* tfa9ec6 = new QByteArray(static_cast<QDate*>(ptr)->toString(static_cast<Qt::DateFormat>(format)).toUtf8()); QtCore_PackedString { const_cast<char*>(tfa9ec6->prepend("WHITESPACE").constData()+10), tfa9ec6->size()-10, tfa9ec6 }; });
}

struct QtCore_PackedString QDate_ToString(void* ptr, struct QtCore_PackedString format)
{
	return ({ QByteArray* te68bf6 = new QByteArray(static_cast<QDate*>(ptr)->toString(QString::fromUtf8(format.data, format.len)).toUtf8()); QtCore_PackedString { const_cast<char*>(te68bf6->prepend("WHITESPACE").constData()+10), te68bf6->size()-10, te68bf6 }; });
}

Q_DECLARE_METATYPE(QDateTime)
Q_DECLARE_METATYPE(QDateTime*)
void* QDateTime_NewQDateTime()
{
	return new QDateTime();
}

void* QDateTime_NewQDateTime7(void* other)
{
	return new QDateTime(*static_cast<QDateTime*>(other));
}

void* QDateTime_NewQDateTime2(void* date)
{
	return new QDateTime(*static_cast<QDate*>(date));
}

void* QDateTime_NewQDateTime3(void* date, void* ti, long long spec)
{
	return new QDateTime(*static_cast<QDate*>(date), *static_cast<QTime*>(ti), static_cast<Qt::TimeSpec>(spec));
}

void* QDateTime_NewQDateTime4(void* date, void* ti, long long spec, int offsetSeconds)
{
	return new QDateTime(*static_cast<QDate*>(date), *static_cast<QTime*>(ti), static_cast<Qt::TimeSpec>(spec), offsetSeconds);
}

void* QDateTime_NewQDateTime5(void* date, void* ti, void* timeZone)
{
	return new QDateTime(*static_cast<QDate*>(date), *static_cast<QTime*>(ti), *static_cast<QTimeZone*>(timeZone));
}

void* QDateTime_NewQDateTime6(void* other)
{
	return new QDateTime(*static_cast<QDateTime*>(other));
}

void QDateTime_DestroyQDateTime(void* ptr)
{
	static_cast<QDateTime*>(ptr)->~QDateTime();
}

struct QtCore_PackedString QDateTime_ToString2(void* ptr, long long format)
{
	return ({ QByteArray* tbd5547 = new QByteArray(static_cast<QDateTime*>(ptr)->toString(static_cast<Qt::DateFormat>(format)).toUtf8()); QtCore_PackedString { const_cast<char*>(tbd5547->prepend("WHITESPACE").constData()+10), tbd5547->size()-10, tbd5547 }; });
}

struct QtCore_PackedString QDateTime_ToString(void* ptr, struct QtCore_PackedString format)
{
	return ({ QByteArray* t78bcae = new QByteArray(static_cast<QDateTime*>(ptr)->toString(QString::fromUtf8(format.data, format.len)).toUtf8()); QtCore_PackedString { const_cast<char*>(t78bcae->prepend("WHITESPACE").constData()+10), t78bcae->size()-10, t78bcae }; });
}

void* QDateTime_Time(void* ptr)
{
	return new QTime(static_cast<QDateTime*>(ptr)->time());
}

Q_DECLARE_METATYPE(QEasingCurve*)
void* QEasingCurve_NewQEasingCurve3(void* other)
{
	return new QEasingCurve(*static_cast<QEasingCurve*>(other));
}

void* QEasingCurve_NewQEasingCurve(long long ty)
{
	return new QEasingCurve(static_cast<QEasingCurve::Type>(ty));
}

void* QEasingCurve_NewQEasingCurve2(void* other)
{
	return new QEasingCurve(*static_cast<QEasingCurve*>(other));
}

void QEasingCurve_DestroyQEasingCurve(void* ptr)
{
	static_cast<QEasingCurve*>(ptr)->~QEasingCurve();
}

long long QEasingCurve_Type(void* ptr)
{
	return static_cast<QEasingCurve*>(ptr)->type();
}

double QEasingCurve_Period(void* ptr)
{
	return static_cast<QEasingCurve*>(ptr)->period();
}

void* QEasingCurve___cubicBezierSpline_atList(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QList<QPointF>*>(ptr)->at(i); if (i == static_cast<QList<QPointF>*>(ptr)->size()-1) { static_cast<QList<QPointF>*>(ptr)->~QList(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QEasingCurve___cubicBezierSpline_setList(void* ptr, void* i)
{
	static_cast<QList<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QEasingCurve___cubicBezierSpline_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPointF>();
}

void* QEasingCurve___toCubicSpline_atList(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QVector<QPointF>*>(ptr)->at(i); if (i == static_cast<QVector<QPointF>*>(ptr)->size()-1) { static_cast<QVector<QPointF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QEasingCurve___toCubicSpline_setList(void* ptr, void* i)
{
	static_cast<QVector<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QEasingCurve___toCubicSpline_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPointF>();
}

class MyQEvent: public QEvent
{
public:
	MyQEvent(Type ty) : QEvent(ty) {QEvent_QEvent_QRegisterMetaType();};
	 ~MyQEvent() { callbackQEvent_DestroyQEvent(this); };
};

Q_DECLARE_METATYPE(QEvent*)
Q_DECLARE_METATYPE(MyQEvent*)

int QEvent_QEvent_QRegisterMetaType(){qRegisterMetaType<QEvent*>(); return qRegisterMetaType<MyQEvent*>();}

void* QEvent_NewQEvent(long long ty)
{
	return new MyQEvent(static_cast<QEvent::Type>(ty));
}

void QEvent_DestroyQEvent(void* ptr)
{
	static_cast<QEvent*>(ptr)->~QEvent();
}

void QEvent_DestroyQEventDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QEvent_Type(void* ptr)
{
	return static_cast<QEvent*>(ptr)->type();
}

unsigned short QEvent_T(void* ptr)
{
	return static_cast<QEvent*>(ptr)->t;
}

void QEvent_SetT(void* ptr, unsigned short vus)
{
	static_cast<QEvent*>(ptr)->t = vus;
}

class MyQIODevice: public QIODevice
{
public:
	MyQIODevice(QObject *parent) : QIODevice(parent) {QIODevice_QIODevice_QRegisterMetaType();};
	MyQIODevice() : QIODevice() {QIODevice_QIODevice_QRegisterMetaType();};
	bool open(QIODevice::OpenMode mode) { return callbackQIODevice_Open(this, mode) != 0; };
	bool reset() { return callbackQIODevice_Reset(this) != 0; };
	qint64 readData(char * data, qint64 maxSize) { QtCore_PackedString dataPacked = { data, maxSize, NULL };return callbackQIODevice_ReadData(this, dataPacked, maxSize); };
	qint64 writeData(const char * data, qint64 maxSize) { QtCore_PackedString dataPacked = { const_cast<char*>(data), maxSize, NULL };return callbackQIODevice_WriteData(this, dataPacked, maxSize); };
	void close() { callbackQIODevice_Close(this); };
	 ~MyQIODevice() { callbackQIODevice_DestroyQIODevice(this); };
	qint64 pos() const { return callbackQIODevice_Pos(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 size() const { return callbackQIODevice_Size(const_cast<void*>(static_cast<const void*>(this))); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
};

Q_DECLARE_METATYPE(QIODevice*)
Q_DECLARE_METATYPE(MyQIODevice*)

int QIODevice_QIODevice_QRegisterMetaType(){qRegisterMetaType<QIODevice*>(); return qRegisterMetaType<MyQIODevice*>();}

void* QIODevice_Read2(void* ptr, long long maxSize)
{
	return new QByteArray(static_cast<QIODevice*>(ptr)->read(maxSize));
}

void* QIODevice_NewQIODevice2(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQIODevice(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQIODevice(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQIODevice(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQIODevice(static_cast<QWidget*>(parent));
	} else {
		return new MyQIODevice(static_cast<QObject*>(parent));
	}
}

void* QIODevice_NewQIODevice()
{
	return new MyQIODevice();
}

char QIODevice_Open(void* ptr, long long mode)
{
	return static_cast<QIODevice*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

char QIODevice_OpenDefault(void* ptr, long long mode)
{
		return static_cast<QIODevice*>(ptr)->QIODevice::open(static_cast<QIODevice::OpenModeFlag>(mode));
}

char QIODevice_Reset(void* ptr)
{
	return static_cast<QIODevice*>(ptr)->reset();
}

char QIODevice_ResetDefault(void* ptr)
{
		return static_cast<QIODevice*>(ptr)->QIODevice::reset();
}

long long QIODevice_Read(void* ptr, char* data, long long maxSize)
{
	return static_cast<QIODevice*>(ptr)->read(data, maxSize);
}

long long QIODevice_ReadData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QIODevice*>(ptr)->readData(data, maxSize);
}

long long QIODevice_Write3(void* ptr, void* byteArray)
{
	return static_cast<QIODevice*>(ptr)->write(*static_cast<QByteArray*>(byteArray));
}

long long QIODevice_Write2(void* ptr, char* data)
{
	return static_cast<QIODevice*>(ptr)->write(const_cast<const char*>(data));
}

long long QIODevice_Write(void* ptr, char* data, long long maxSize)
{
	return static_cast<QIODevice*>(ptr)->write(const_cast<const char*>(data), maxSize);
}

long long QIODevice_WriteData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QIODevice*>(ptr)->writeData(const_cast<const char*>(data), maxSize);
}

void QIODevice_Close(void* ptr)
{
	static_cast<QIODevice*>(ptr)->close();
}

void QIODevice_CloseDefault(void* ptr)
{
		static_cast<QIODevice*>(ptr)->QIODevice::close();
}

void QIODevice_DestroyQIODevice(void* ptr)
{
	static_cast<QIODevice*>(ptr)->~QIODevice();
}

void QIODevice_DestroyQIODeviceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QIODevice_Pos(void* ptr)
{
	return static_cast<QIODevice*>(ptr)->pos();
}

long long QIODevice_PosDefault(void* ptr)
{
		return static_cast<QIODevice*>(ptr)->QIODevice::pos();
}

long long QIODevice_Size(void* ptr)
{
	return static_cast<QIODevice*>(ptr)->size();
}

long long QIODevice_SizeDefault(void* ptr)
{
		return static_cast<QIODevice*>(ptr)->QIODevice::size();
}

Q_DECLARE_METATYPE(QJsonArray)
Q_DECLARE_METATYPE(QJsonArray*)
void* QJsonArray_NewQJsonArray()
{
	return new QJsonArray();
}

void* QJsonArray_NewQJsonArray3(void* other)
{
	return new QJsonArray(*static_cast<QJsonArray*>(other));
}

void QJsonArray_Replace(void* ptr, int i, void* value)
{
	static_cast<QJsonArray*>(ptr)->replace(i, *static_cast<QJsonValue*>(value));
}

void QJsonArray_DestroyQJsonArray(void* ptr)
{
	static_cast<QJsonArray*>(ptr)->~QJsonArray();
}

void* QJsonArray_Last(void* ptr)
{
	return new QJsonValue(static_cast<QJsonArray*>(ptr)->last());
}

char QJsonArray_Contains(void* ptr, void* value)
{
	return static_cast<QJsonArray*>(ptr)->contains(*static_cast<QJsonValue*>(value));
}

char QJsonArray_Empty(void* ptr)
{
	return static_cast<QJsonArray*>(ptr)->empty();
}

int QJsonArray_Size(void* ptr)
{
	return static_cast<QJsonArray*>(ptr)->size();
}

void* QJsonArray___fromVariantList_list_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QJsonArray___fromVariantList_list_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QJsonArray___fromVariantList_list_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QJsonArray___toVariantList_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QJsonArray___toVariantList_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QJsonArray___toVariantList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

Q_DECLARE_METATYPE(QJsonDocument)
Q_DECLARE_METATYPE(QJsonDocument*)
void* QJsonDocument_NewQJsonDocument()
{
	return new QJsonDocument();
}

void* QJsonDocument_NewQJsonDocument3(void* array)
{
	return new QJsonDocument(*static_cast<QJsonArray*>(array));
}

void* QJsonDocument_NewQJsonDocument4(void* other)
{
	return new QJsonDocument(*static_cast<QJsonDocument*>(other));
}

void* QJsonDocument_NewQJsonDocument2(void* object)
{
	return new QJsonDocument(*static_cast<QJsonObject*>(object));
}

void QJsonDocument_SetObject(void* ptr, void* object)
{
	static_cast<QJsonDocument*>(ptr)->setObject(*static_cast<QJsonObject*>(object));
}

void QJsonDocument_DestroyQJsonDocument(void* ptr)
{
	static_cast<QJsonDocument*>(ptr)->~QJsonDocument();
}

void* QJsonDocument_ToJson(void* ptr, long long format)
{
	return new QByteArray(static_cast<QJsonDocument*>(ptr)->toJson(static_cast<QJsonDocument::JsonFormat>(format)));
}

void* QJsonDocument_Object(void* ptr)
{
	return new QJsonObject(static_cast<QJsonDocument*>(ptr)->object());
}

void* QJsonDocument_ToVariant(void* ptr)
{
	return new QVariant(static_cast<QJsonDocument*>(ptr)->toVariant());
}

Q_DECLARE_METATYPE(QJsonObject)
Q_DECLARE_METATYPE(QJsonObject*)
void* QJsonObject_NewQJsonObject()
{
	return new QJsonObject();
}

void* QJsonObject_NewQJsonObject3(void* other)
{
	return new QJsonObject(*static_cast<QJsonObject*>(other));
}

void QJsonObject_Remove(void* ptr, struct QtCore_PackedString key)
{
	static_cast<QJsonObject*>(ptr)->remove(QString::fromUtf8(key.data, key.len));
}

void QJsonObject_DestroyQJsonObject(void* ptr)
{
	static_cast<QJsonObject*>(ptr)->~QJsonObject();
}

void* QJsonObject_Value2(void* ptr, void* key)
{
	return new QJsonValue(static_cast<QJsonObject*>(ptr)->value(*static_cast<QLatin1String*>(key)));
}

void* QJsonObject_Value(void* ptr, struct QtCore_PackedString key)
{
	return new QJsonValue(static_cast<QJsonObject*>(ptr)->value(QString::fromUtf8(key.data, key.len)));
}

struct QtCore_PackedString QJsonObject_Keys(void* ptr)
{
	return ({ QByteArray* t6c6b45 = new QByteArray(static_cast<QJsonObject*>(ptr)->keys().join("¡¦!").toUtf8()); QtCore_PackedString { const_cast<char*>(t6c6b45->prepend("WHITESPACE").constData()+10), t6c6b45->size()-10, t6c6b45 }; });
}

char QJsonObject_Contains2(void* ptr, void* key)
{
	return static_cast<QJsonObject*>(ptr)->contains(*static_cast<QLatin1String*>(key));
}

char QJsonObject_Contains(void* ptr, struct QtCore_PackedString key)
{
	return static_cast<QJsonObject*>(ptr)->contains(QString::fromUtf8(key.data, key.len));
}

char QJsonObject_Empty(void* ptr)
{
	return static_cast<QJsonObject*>(ptr)->empty();
}

int QJsonObject_Length(void* ptr)
{
	return static_cast<QJsonObject*>(ptr)->length();
}

int QJsonObject_Size(void* ptr)
{
	return static_cast<QJsonObject*>(ptr)->size();
}

void* QJsonObject___fromVariantHash_hash_atList(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QHash<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QHash<QString, QVariant>*>(ptr)->size()-1) { static_cast<QHash<QString, QVariant>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QJsonObject___fromVariantHash_hash_setList(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QHash<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QJsonObject___fromVariantHash_hash_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<QString, QVariant>();
}

struct QtCore_PackedList QJsonObject___fromVariantHash_hash_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValuef43bc5 = new QList<QString>(static_cast<QHash<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValuef43bc5, tmpValuef43bc5->size() }; });
}

void* QJsonObject___toVariantHash_atList(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QHash<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QHash<QString, QVariant>*>(ptr)->size()-1) { static_cast<QHash<QString, QVariant>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QJsonObject___toVariantHash_setList(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QHash<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QJsonObject___toVariantHash_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<QString, QVariant>();
}

struct QtCore_PackedList QJsonObject___toVariantHash_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValuef43bc5 = new QList<QString>(static_cast<QHash<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValuef43bc5, tmpValuef43bc5->size() }; });
}

void* QJsonObject___toVariantMap_atList(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QJsonObject___toVariantMap_setList(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QJsonObject___toVariantMap_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtCore_PackedList QJsonObject___toVariantMap_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct QtCore_PackedString QJsonObject_____fromVariantHash_hash_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtCore_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QJsonObject_____fromVariantHash_hash_keyList_setList(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QJsonObject_____fromVariantHash_hash_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtCore_PackedString QJsonObject_____fromVariantMap_map_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtCore_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QJsonObject_____fromVariantMap_map_keyList_setList(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QJsonObject_____fromVariantMap_map_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtCore_PackedString QJsonObject_____toVariantHash_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtCore_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QJsonObject_____toVariantHash_keyList_setList(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QJsonObject_____toVariantHash_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtCore_PackedString QJsonObject_____toVariantMap_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtCore_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QJsonObject_____toVariantMap_keyList_setList(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QJsonObject_____toVariantMap_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

Q_DECLARE_METATYPE(QJsonValue*)
void* QJsonValue_NewQJsonValue7(void* s)
{
	return new QJsonValue(*static_cast<QLatin1String*>(s));
}

void* QJsonValue_NewQJsonValue(long long ty)
{
	return new QJsonValue(static_cast<QJsonValue::Type>(ty));
}

void* QJsonValue_NewQJsonValue2(char b)
{
	return new QJsonValue(b != 0);
}

void* QJsonValue_NewQJsonValue9(void* a)
{
	return new QJsonValue(*static_cast<QJsonArray*>(a));
}

void* QJsonValue_NewQJsonValue10(void* o)
{
	return new QJsonValue(*static_cast<QJsonObject*>(o));
}

void* QJsonValue_NewQJsonValue11(void* other)
{
	return new QJsonValue(*static_cast<QJsonValue*>(other));
}

void* QJsonValue_NewQJsonValue6(struct QtCore_PackedString s)
{
	return new QJsonValue(QString::fromUtf8(s.data, s.len));
}

void* QJsonValue_NewQJsonValue8(char* s)
{
	return new QJsonValue(const_cast<const char*>(s));
}

void* QJsonValue_NewQJsonValue3(double n)
{
	return new QJsonValue(n);
}

void* QJsonValue_NewQJsonValue4(int n)
{
	return new QJsonValue(n);
}

void* QJsonValue_NewQJsonValue5(long long n)
{
	return new QJsonValue(n);
}

void QJsonValue_DestroyQJsonValue(void* ptr)
{
	static_cast<QJsonValue*>(ptr)->~QJsonValue();
}

struct QtCore_PackedString QJsonValue_ToString(void* ptr)
{
	return ({ QByteArray* t54cc8a = new QByteArray(static_cast<QJsonValue*>(ptr)->toString().toUtf8()); QtCore_PackedString { const_cast<char*>(t54cc8a->prepend("WHITESPACE").constData()+10), t54cc8a->size()-10, t54cc8a }; });
}

struct QtCore_PackedString QJsonValue_ToString2(void* ptr, struct QtCore_PackedString defaultValue)
{
	return ({ QByteArray* tb25a8c = new QByteArray(static_cast<QJsonValue*>(ptr)->toString(QString::fromUtf8(defaultValue.data, defaultValue.len)).toUtf8()); QtCore_PackedString { const_cast<char*>(tb25a8c->prepend("WHITESPACE").constData()+10), tb25a8c->size()-10, tb25a8c }; });
}

void* QJsonValue_ToVariant(void* ptr)
{
	return new QVariant(static_cast<QJsonValue*>(ptr)->toVariant());
}

long long QJsonValue_Type(void* ptr)
{
	return static_cast<QJsonValue*>(ptr)->type();
}

Q_DECLARE_METATYPE(QLatin1Char*)
void* QLatin1Char_NewQLatin1Char(char* c)
{
	return new QLatin1Char(*c);
}

Q_DECLARE_METATYPE(QLatin1String)
Q_DECLARE_METATYPE(QLatin1String*)
void* QLatin1String_NewQLatin1String()
{
	return new QLatin1String();
}

void* QLatin1String_NewQLatin1String4(void* str)
{
	return new QLatin1String(*static_cast<QByteArray*>(str));
}

void* QLatin1String_NewQLatin1String2(char* str)
{
	return new QLatin1String(const_cast<const char*>(str));
}

void* QLatin1String_NewQLatin1String3(char* str, int size)
{
	return new QLatin1String(const_cast<const char*>(str), size);
}

struct QtCore_PackedString QLatin1String_Data(void* ptr)
{
	return QtCore_PackedString { const_cast<char*>(static_cast<QLatin1String*>(ptr)->data()), static_cast<QLatin1String*>(ptr)->size(), NULL };
}

int QLatin1String_Size(void* ptr)
{
	return static_cast<QLatin1String*>(ptr)->size();
}

Q_DECLARE_METATYPE(QLine)
Q_DECLARE_METATYPE(QLine*)
void* QLine_NewQLine()
{
	return new QLine();
}

void* QLine_NewQLine2(void* p1, void* p2)
{
	return new QLine(*static_cast<QPoint*>(p1), *static_cast<QPoint*>(p2));
}

void* QLine_NewQLine3(int x1, int y1, int x2, int y2)
{
	return new QLine(x1, y1, x2, y2);
}

void QLine_Translate(void* ptr, void* offset)
{
	static_cast<QLine*>(ptr)->translate(*static_cast<QPoint*>(offset));
}

void QLine_Translate2(void* ptr, int dx, int dy)
{
	static_cast<QLine*>(ptr)->translate(dx, dy);
}

Q_DECLARE_METATYPE(QLineF)
Q_DECLARE_METATYPE(QLineF*)
void* QLineF_NewQLineF()
{
	return new QLineF();
}

void* QLineF_NewQLineF4(void* line)
{
	return new QLineF(*static_cast<QLine*>(line));
}

void* QLineF_NewQLineF2(void* p1, void* p2)
{
	return new QLineF(*static_cast<QPointF*>(p1), *static_cast<QPointF*>(p2));
}

void* QLineF_NewQLineF3(double x1, double y1, double x2, double y2)
{
	return new QLineF(x1, y1, x2, y2);
}

void QLineF_Translate(void* ptr, void* offset)
{
	static_cast<QLineF*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

void QLineF_Translate2(void* ptr, double dx, double dy)
{
	static_cast<QLineF*>(ptr)->translate(dx, dy);
}

double QLineF_Length(void* ptr)
{
	return static_cast<QLineF*>(ptr)->length();
}

Q_DECLARE_METATYPE(QLocale)
Q_DECLARE_METATYPE(QLocale*)
void* QLocale_QLocale_C()
{
	return new QLocale(QLocale::c());
}

void* QLocale_NewQLocale()
{
	return new QLocale();
}

void* QLocale_NewQLocale3(long long language, long long country)
{
	return new QLocale(static_cast<QLocale::Language>(language), static_cast<QLocale::Country>(country));
}

void* QLocale_NewQLocale4(long long language, long long scri, long long country)
{
	return new QLocale(static_cast<QLocale::Language>(language), static_cast<QLocale::Script>(scri), static_cast<QLocale::Country>(country));
}

void* QLocale_NewQLocale5(void* other)
{
	return new QLocale(*static_cast<QLocale*>(other));
}

void* QLocale_NewQLocale2(struct QtCore_PackedString name)
{
	return new QLocale(QString::fromUtf8(name.data, name.len));
}

void QLocale_DestroyQLocale(void* ptr)
{
	static_cast<QLocale*>(ptr)->~QLocale();
}

struct QtCore_PackedString QLocale_Name(void* ptr)
{
	return ({ QByteArray* tee867e = new QByteArray(static_cast<QLocale*>(ptr)->name().toUtf8()); QtCore_PackedString { const_cast<char*>(tee867e->prepend("WHITESPACE").constData()+10), tee867e->size()-10, tee867e }; });
}

struct QtCore_PackedString QLocale_ToLower(void* ptr, struct QtCore_PackedString str)
{
	return ({ QByteArray* t112fbe = new QByteArray(static_cast<QLocale*>(ptr)->toLower(QString::fromUtf8(str.data, str.len)).toUtf8()); QtCore_PackedString { const_cast<char*>(t112fbe->prepend("WHITESPACE").constData()+10), t112fbe->size()-10, t112fbe }; });
}

struct QtCore_PackedString QLocale_ToString10(void* ptr, void* date, long long format)
{
	return ({ QByteArray* t1fa262 = new QByteArray(static_cast<QLocale*>(ptr)->toString(*static_cast<QDate*>(date), static_cast<QLocale::FormatType>(format)).toUtf8()); QtCore_PackedString { const_cast<char*>(t1fa262->prepend("WHITESPACE").constData()+10), t1fa262->size()-10, t1fa262 }; });
}

struct QtCore_PackedString QLocale_ToString9(void* ptr, void* date, struct QtCore_PackedString format)
{
	return ({ QByteArray* t198629 = new QByteArray(static_cast<QLocale*>(ptr)->toString(*static_cast<QDate*>(date), QString::fromUtf8(format.data, format.len)).toUtf8()); QtCore_PackedString { const_cast<char*>(t198629->prepend("WHITESPACE").constData()+10), t198629->size()-10, t198629 }; });
}

struct QtCore_PackedString QLocale_ToString13(void* ptr, void* dateTime, long long format)
{
	return ({ QByteArray* t96745f = new QByteArray(static_cast<QLocale*>(ptr)->toString(*static_cast<QDateTime*>(dateTime), static_cast<QLocale::FormatType>(format)).toUtf8()); QtCore_PackedString { const_cast<char*>(t96745f->prepend("WHITESPACE").constData()+10), t96745f->size()-10, t96745f }; });
}

struct QtCore_PackedString QLocale_ToString14(void* ptr, void* dateTime, struct QtCore_PackedString format)
{
	return ({ QByteArray* t336e5f = new QByteArray(static_cast<QLocale*>(ptr)->toString(*static_cast<QDateTime*>(dateTime), QString::fromUtf8(format.data, format.len)).toUtf8()); QtCore_PackedString { const_cast<char*>(t336e5f->prepend("WHITESPACE").constData()+10), t336e5f->size()-10, t336e5f }; });
}

struct QtCore_PackedString QLocale_ToString12(void* ptr, void* ti, long long format)
{
	return ({ QByteArray* tb6230e = new QByteArray(static_cast<QLocale*>(ptr)->toString(*static_cast<QTime*>(ti), static_cast<QLocale::FormatType>(format)).toUtf8()); QtCore_PackedString { const_cast<char*>(tb6230e->prepend("WHITESPACE").constData()+10), tb6230e->size()-10, tb6230e }; });
}

struct QtCore_PackedString QLocale_ToString11(void* ptr, void* ti, struct QtCore_PackedString format)
{
	return ({ QByteArray* t607af9 = new QByteArray(static_cast<QLocale*>(ptr)->toString(*static_cast<QTime*>(ti), QString::fromUtf8(format.data, format.len)).toUtf8()); QtCore_PackedString { const_cast<char*>(t607af9->prepend("WHITESPACE").constData()+10), t607af9->size()-10, t607af9 }; });
}

struct QtCore_PackedString QLocale_ToString7(void* ptr, double i, char* ff, int prec)
{
	return ({ QByteArray* t07d146 = new QByteArray(static_cast<QLocale*>(ptr)->toString(i, *ff, prec).toUtf8()); QtCore_PackedString { const_cast<char*>(t07d146->prepend("WHITESPACE").constData()+10), t07d146->size()-10, t07d146 }; });
}

struct QtCore_PackedString QLocale_ToString8(void* ptr, float i, char* ff, int prec)
{
	return ({ QByteArray* t07d146 = new QByteArray(static_cast<QLocale*>(ptr)->toString(i, *ff, prec).toUtf8()); QtCore_PackedString { const_cast<char*>(t07d146->prepend("WHITESPACE").constData()+10), t07d146->size()-10, t07d146 }; });
}

struct QtCore_PackedString QLocale_ToString5(void* ptr, int i)
{
	return ({ QByteArray* teea9c2 = new QByteArray(static_cast<QLocale*>(ptr)->toString(i).toUtf8()); QtCore_PackedString { const_cast<char*>(teea9c2->prepend("WHITESPACE").constData()+10), teea9c2->size()-10, teea9c2 }; });
}

struct QtCore_PackedString QLocale_ToString(void* ptr, long long i)
{
	return ({ QByteArray* teea9c2 = new QByteArray(static_cast<QLocale*>(ptr)->toString(i).toUtf8()); QtCore_PackedString { const_cast<char*>(teea9c2->prepend("WHITESPACE").constData()+10), teea9c2->size()-10, teea9c2 }; });
}

struct QtCore_PackedString QLocale_ToString2(void* ptr, unsigned long long i)
{
	return ({ QByteArray* teea9c2 = new QByteArray(static_cast<QLocale*>(ptr)->toString(i).toUtf8()); QtCore_PackedString { const_cast<char*>(teea9c2->prepend("WHITESPACE").constData()+10), teea9c2->size()-10, teea9c2 }; });
}

struct QtCore_PackedString QLocale_ToString3(void* ptr, short i)
{
	return ({ QByteArray* teea9c2 = new QByteArray(static_cast<QLocale*>(ptr)->toString(i).toUtf8()); QtCore_PackedString { const_cast<char*>(teea9c2->prepend("WHITESPACE").constData()+10), teea9c2->size()-10, teea9c2 }; });
}

struct QtCore_PackedString QLocale_ToString6(void* ptr, unsigned int i)
{
	return ({ QByteArray* teea9c2 = new QByteArray(static_cast<QLocale*>(ptr)->toString(i).toUtf8()); QtCore_PackedString { const_cast<char*>(teea9c2->prepend("WHITESPACE").constData()+10), teea9c2->size()-10, teea9c2 }; });
}

struct QtCore_PackedString QLocale_ToString4(void* ptr, unsigned short i)
{
	return ({ QByteArray* teea9c2 = new QByteArray(static_cast<QLocale*>(ptr)->toString(i).toUtf8()); QtCore_PackedString { const_cast<char*>(teea9c2->prepend("WHITESPACE").constData()+10), teea9c2->size()-10, teea9c2 }; });
}

struct QtCore_PackedString QLocale_ToUpper(void* ptr, struct QtCore_PackedString str)
{
	return ({ QByteArray* t6969d0 = new QByteArray(static_cast<QLocale*>(ptr)->toUpper(QString::fromUtf8(str.data, str.len)).toUtf8()); QtCore_PackedString { const_cast<char*>(t6969d0->prepend("WHITESPACE").constData()+10), t6969d0->size()-10, t6969d0 }; });
}

void* QLocale___matchingLocales_atList(void* ptr, int i)
{
	return new QLocale(({QLocale tmp = static_cast<QList<QLocale>*>(ptr)->at(i); if (i == static_cast<QList<QLocale>*>(ptr)->size()-1) { static_cast<QList<QLocale>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QLocale___matchingLocales_setList(void* ptr, void* i)
{
	static_cast<QList<QLocale>*>(ptr)->append(*static_cast<QLocale*>(i));
}

void* QLocale___matchingLocales_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QLocale>();
}

long long QLocale___weekdays_atList(void* ptr, int i)
{
	return ({Qt::DayOfWeek tmp = static_cast<QList<Qt::DayOfWeek>*>(ptr)->at(i); if (i == static_cast<QList<Qt::DayOfWeek>*>(ptr)->size()-1) { static_cast<QList<Qt::DayOfWeek>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLocale___weekdays_setList(void* ptr, long long i)
{
	static_cast<QList<Qt::DayOfWeek>*>(ptr)->append(static_cast<Qt::DayOfWeek>(i));
}

void* QLocale___weekdays_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Qt::DayOfWeek>();
}

Q_DECLARE_METATYPE(QMargins)
Q_DECLARE_METATYPE(QMargins*)
void* QMargins_NewQMargins()
{
	return new QMargins();
}

void* QMargins_NewQMargins2(int left, int top, int right, int bottom)
{
	return new QMargins(left, top, right, bottom);
}

void* QMetaMethod_Name(void* ptr)
{
	return new QByteArray(static_cast<QMetaMethod*>(ptr)->name());
}

void* QMetaMethod___parameterNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QMetaMethod___parameterNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QMetaMethod___parameterNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QMetaMethod___parameterTypes_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QMetaMethod___parameterTypes_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QMetaMethod___parameterTypes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QMetaObject_Constructor(void* ptr, int index)
{
	return new QMetaMethod(static_cast<QMetaObject*>(ptr)->constructor(index));
}

void* QMetaObject_Method(void* ptr, int index)
{
	return new QMetaMethod(static_cast<QMetaObject*>(ptr)->method(index));
}

struct QtCore_PackedString QMetaObject_ClassName(void* ptr)
{
	return QtCore_PackedString { const_cast<char*>(static_cast<QMetaObject*>(ptr)->className()), -1, NULL };
}

Q_DECLARE_METATYPE(QModelIndex)
Q_DECLARE_METATYPE(QModelIndex*)
void* QModelIndex_NewQModelIndex()
{
	return new QModelIndex();
}

void* QModelIndex_Data(void* ptr, int role)
{
	return new QVariant(static_cast<QModelIndex*>(ptr)->data(role));
}

class MyQObject: public QObject
{
public:
	MyQObject(QObject *parent = Q_NULLPTR) : QObject(parent) {QObject_QObject_QRegisterMetaType();};
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
	 ~MyQObject() { callbackQObject_DestroyQObject(this); };
};

Q_DECLARE_METATYPE(MyQObject*)

int QObject_QObject_QRegisterMetaType(){qRegisterMetaType<QObject*>(); return qRegisterMetaType<MyQObject*>();}

void* QObject_NewQObject(void* parent)
{
	return new MyQObject(static_cast<QObject*>(parent));
}

struct QtCore_PackedString QObject_QObject_Tr(char* sourceText, char* disambiguation, int n)
{
	return ({ QByteArray* te5b32b = new QByteArray(QObject::tr(const_cast<const char*>(sourceText), const_cast<const char*>(disambiguation), n).toUtf8()); QtCore_PackedString { const_cast<char*>(te5b32b->prepend("WHITESPACE").constData()+10), te5b32b->size()-10, te5b32b }; });
}

void QObject_ConnectDestroyed(void* ptr, long long t)
{
	QObject::connect(static_cast<QObject*>(ptr), static_cast<void (QObject::*)(QObject *)>(&QObject::destroyed), static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(QObject *)>(&MyQObject::Signal_Destroyed), static_cast<Qt::ConnectionType>(t));
}

void QObject_DisconnectDestroyed(void* ptr)
{
	QObject::disconnect(static_cast<QObject*>(ptr), static_cast<void (QObject::*)(QObject *)>(&QObject::destroyed), static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(QObject *)>(&MyQObject::Signal_Destroyed));
}

void QObject_Destroyed(void* ptr, void* obj)
{
	static_cast<QObject*>(ptr)->destroyed(static_cast<QObject*>(obj));
}

void QObject_ConnectObjectNameChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QObject*>(ptr), &QObject::objectNameChanged, static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(const QString &)>(&MyQObject::Signal_ObjectNameChanged), static_cast<Qt::ConnectionType>(t));
}

void QObject_DisconnectObjectNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QObject*>(ptr), &QObject::objectNameChanged, static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(const QString &)>(&MyQObject::Signal_ObjectNameChanged));
}

void QObject_SetObjectName(void* ptr, struct QtCore_PackedString name)
{
	static_cast<QObject*>(ptr)->setObjectName(QString::fromUtf8(name.data, name.len));
}

void QObject_DestroyQObject(void* ptr)
{
	static_cast<QObject*>(ptr)->~QObject();
}

void QObject_DestroyQObjectDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtCore_PackedString QObject_ObjectName(void* ptr)
{
	return ({ QByteArray* te77be1 = new QByteArray(static_cast<QObject*>(ptr)->objectName().toUtf8()); QtCore_PackedString { const_cast<char*>(te77be1->prepend("WHITESPACE").constData()+10), te77be1->size()-10, te77be1 }; });
}

void* QObject_ToVariant(void* ptr)
{
	return new QVariant(QVariant::fromValue(static_cast<QObject*>(ptr)));
}

void* QObject___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QObject___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QObject___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QObject___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QObject___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QObject___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QObject___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QObject___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QObject___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QObject___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QObject___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QObject___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QObject___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QObject___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QObject___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

Q_DECLARE_METATYPE(QPersistentModelIndex*)
void* QPersistentModelIndex_NewQPersistentModelIndex4(void* other)
{
	return new QPersistentModelIndex(*static_cast<QPersistentModelIndex*>(other));
}

void* QPersistentModelIndex_NewQPersistentModelIndex(void* index)
{
	return new QPersistentModelIndex(*static_cast<QModelIndex*>(index));
}

void* QPersistentModelIndex_NewQPersistentModelIndex3(void* other)
{
	return new QPersistentModelIndex(*static_cast<QPersistentModelIndex*>(other));
}

void* QPersistentModelIndex_Data(void* ptr, int role)
{
	return new QVariant(static_cast<QPersistentModelIndex*>(ptr)->data(role));
}

Q_DECLARE_METATYPE(QPoint)
Q_DECLARE_METATYPE(QPoint*)
void* QPoint_NewQPoint()
{
	return new QPoint();
}

void* QPoint_NewQPoint2(int xpos, int ypos)
{
	return new QPoint(xpos, ypos);
}

int QPoint_X(void* ptr)
{
	return static_cast<QPoint*>(ptr)->x();
}

int QPoint_Y(void* ptr)
{
	return static_cast<QPoint*>(ptr)->y();
}

Q_DECLARE_METATYPE(QPointF)
Q_DECLARE_METATYPE(QPointF*)
void* QPointF_NewQPointF()
{
	return new QPointF();
}

void* QPointF_NewQPointF2(void* point)
{
	return new QPointF(*static_cast<QPoint*>(point));
}

void* QPointF_NewQPointF3(double xpos, double ypos)
{
	return new QPointF(xpos, ypos);
}

double QPointF_X(void* ptr)
{
	return static_cast<QPointF*>(ptr)->x();
}

double QPointF_Y(void* ptr)
{
	return static_cast<QPointF*>(ptr)->y();
}

Q_DECLARE_METATYPE(QRect)
Q_DECLARE_METATYPE(QRect*)
void* QRect_NewQRect()
{
	return new QRect();
}

void* QRect_NewQRect2(void* topLeft, void* bottomRight)
{
	return new QRect(*static_cast<QPoint*>(topLeft), *static_cast<QPoint*>(bottomRight));
}

void* QRect_NewQRect3(void* topLeft, void* size)
{
	return new QRect(*static_cast<QPoint*>(topLeft), *static_cast<QSize*>(size));
}

void* QRect_NewQRect4(int x, int y, int width, int height)
{
	return new QRect(x, y, width, height);
}

void QRect_Translate2(void* ptr, void* offset)
{
	static_cast<QRect*>(ptr)->translate(*static_cast<QPoint*>(offset));
}

void QRect_Translate(void* ptr, int dx, int dy)
{
	static_cast<QRect*>(ptr)->translate(dx, dy);
}

void* QRect_Size(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QRect*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

char QRect_Contains(void* ptr, void* point, char proper)
{
	return static_cast<QRect*>(ptr)->contains(*static_cast<QPoint*>(point), proper != 0);
}

char QRect_Contains2(void* ptr, void* rectangle, char proper)
{
	return static_cast<QRect*>(ptr)->contains(*static_cast<QRect*>(rectangle), proper != 0);
}

char QRect_Contains3(void* ptr, int x, int y)
{
	return static_cast<QRect*>(ptr)->contains(x, y);
}

char QRect_Contains4(void* ptr, int x, int y, char proper)
{
	return static_cast<QRect*>(ptr)->contains(x, y, proper != 0);
}

int QRect_X(void* ptr)
{
	return static_cast<QRect*>(ptr)->x();
}

int QRect_Y(void* ptr)
{
	return static_cast<QRect*>(ptr)->y();
}

Q_DECLARE_METATYPE(QRectF)
Q_DECLARE_METATYPE(QRectF*)
void* QRectF_NewQRectF()
{
	return new QRectF();
}

void* QRectF_NewQRectF3(void* topLeft, void* bottomRight)
{
	return new QRectF(*static_cast<QPointF*>(topLeft), *static_cast<QPointF*>(bottomRight));
}

void* QRectF_NewQRectF2(void* topLeft, void* size)
{
	return new QRectF(*static_cast<QPointF*>(topLeft), *static_cast<QSizeF*>(size));
}

void* QRectF_NewQRectF5(void* rectangle)
{
	return new QRectF(*static_cast<QRect*>(rectangle));
}

void* QRectF_NewQRectF4(double x, double y, double width, double height)
{
	return new QRectF(x, y, width, height);
}

void QRectF_Translate2(void* ptr, void* offset)
{
	static_cast<QRectF*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

void QRectF_Translate(void* ptr, double dx, double dy)
{
	static_cast<QRectF*>(ptr)->translate(dx, dy);
}

void* QRectF_Size(void* ptr)
{
	return ({ QSizeF tmpValue = static_cast<QRectF*>(ptr)->size(); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

char QRectF_Contains(void* ptr, void* point)
{
	return static_cast<QRectF*>(ptr)->contains(*static_cast<QPointF*>(point));
}

char QRectF_Contains2(void* ptr, void* rectangle)
{
	return static_cast<QRectF*>(ptr)->contains(*static_cast<QRectF*>(rectangle));
}

char QRectF_Contains3(void* ptr, double x, double y)
{
	return static_cast<QRectF*>(ptr)->contains(x, y);
}

double QRectF_X(void* ptr)
{
	return static_cast<QRectF*>(ptr)->x();
}

double QRectF_Y(void* ptr)
{
	return static_cast<QRectF*>(ptr)->y();
}

Q_DECLARE_METATYPE(QRegExp)
Q_DECLARE_METATYPE(QRegExp*)
void* QRegExp_NewQRegExp()
{
	return new QRegExp();
}

void* QRegExp_NewQRegExp3(void* rx)
{
	return new QRegExp(*static_cast<QRegExp*>(rx));
}

void* QRegExp_NewQRegExp2(struct QtCore_PackedString pattern, long long cs, long long syntax)
{
	return new QRegExp(QString::fromUtf8(pattern.data, pattern.len), static_cast<Qt::CaseSensitivity>(cs), static_cast<QRegExp::PatternSyntax>(syntax));
}

void QRegExp_DestroyQRegExp(void* ptr)
{
	static_cast<QRegExp*>(ptr)->~QRegExp();
}

int QRegExp_Pos(void* ptr, int nth)
{
	return static_cast<QRegExp*>(ptr)->pos(nth);
}

Q_DECLARE_METATYPE(QRegularExpression)
Q_DECLARE_METATYPE(QRegularExpression*)
void* QRegularExpression_NewQRegularExpression()
{
	return new QRegularExpression();
}

void* QRegularExpression_NewQRegularExpression3(void* re)
{
	return new QRegularExpression(*static_cast<QRegularExpression*>(re));
}

void* QRegularExpression_NewQRegularExpression2(struct QtCore_PackedString pattern, long long options)
{
	return new QRegularExpression(QString::fromUtf8(pattern.data, pattern.len), static_cast<QRegularExpression::PatternOption>(options));
}

void QRegularExpression_DestroyQRegularExpression(void* ptr)
{
	static_cast<QRegularExpression*>(ptr)->~QRegularExpression();
}

Q_DECLARE_METATYPE(QRegularExpressionMatch)
Q_DECLARE_METATYPE(QRegularExpressionMatch*)
void* QRegularExpressionMatch_NewQRegularExpressionMatch()
{
	return new QRegularExpressionMatch();
}

void* QRegularExpressionMatch_NewQRegularExpressionMatch2(void* match)
{
	return new QRegularExpressionMatch(*static_cast<QRegularExpressionMatch*>(match));
}

void QRegularExpressionMatch_DestroyQRegularExpressionMatch(void* ptr)
{
	static_cast<QRegularExpressionMatch*>(ptr)->~QRegularExpressionMatch();
}

Q_DECLARE_METATYPE(QSize)
Q_DECLARE_METATYPE(QSize*)
void* QSize_NewQSize()
{
	return new QSize();
}

void* QSize_NewQSize2(int width, int height)
{
	return new QSize(width, height);
}

Q_DECLARE_METATYPE(QSizeF)
Q_DECLARE_METATYPE(QSizeF*)
void* QSizeF_NewQSizeF()
{
	return new QSizeF();
}

void* QSizeF_NewQSizeF2(void* size)
{
	return new QSizeF(*static_cast<QSize*>(size));
}

void* QSizeF_NewQSizeF3(double width, double height)
{
	return new QSizeF(width, height);
}

Q_DECLARE_METATYPE(QStringRef)
Q_DECLARE_METATYPE(QStringRef*)
void* QStringRef_NewQStringRef()
{
	return new QStringRef();
}

void* QStringRef_NewQStringRef5(void* other)
{
	return new QStringRef(*static_cast<QStringRef*>(other));
}

void* QStringRef_NewQStringRef3(struct QtCore_PackedString stri)
{
	return new QStringRef(new QString(QString::fromUtf8(stri.data, stri.len)));
}

void* QStringRef_NewQStringRef2(struct QtCore_PackedString stri, int position, int length)
{
	return new QStringRef(new QString(QString::fromUtf8(stri.data, stri.len)), position, length);
}

void* QStringRef_NewQStringRef4(void* other)
{
	return new QStringRef(*static_cast<QStringRef*>(other));
}

void QStringRef_DestroyQStringRef(void* ptr)
{
	static_cast<QStringRef*>(ptr)->~QStringRef();
}

struct QtCore_PackedString QStringRef_ToString(void* ptr)
{
	return ({ QByteArray* t336448 = new QByteArray(static_cast<QStringRef*>(ptr)->toString().toUtf8()); QtCore_PackedString { const_cast<char*>(t336448->prepend("WHITESPACE").constData()+10), t336448->size()-10, t336448 }; });
}

struct QtCore_PackedList QStringRef_Split2(void* ptr, void* sep, long long behavior, long long cs)
{
	return ({ QVector<QStringRef>* tmpValue643751 = new QVector<QStringRef>(static_cast<QStringRef*>(ptr)->split(*static_cast<QChar*>(sep), static_cast<QString::SplitBehavior>(behavior), static_cast<Qt::CaseSensitivity>(cs))); QtCore_PackedList { tmpValue643751, tmpValue643751->size() }; });
}

struct QtCore_PackedList QStringRef_Split(void* ptr, struct QtCore_PackedString sep, long long behavior, long long cs)
{
	return ({ QVector<QStringRef>* tmpValued1bf8d = new QVector<QStringRef>(static_cast<QStringRef*>(ptr)->split(QString::fromUtf8(sep.data, sep.len), static_cast<QString::SplitBehavior>(behavior), static_cast<Qt::CaseSensitivity>(cs))); QtCore_PackedList { tmpValued1bf8d, tmpValued1bf8d->size() }; });
}

char QStringRef_Contains2(void* ptr, void* ch, long long cs)
{
	return static_cast<QStringRef*>(ptr)->contains(*static_cast<QChar*>(ch), static_cast<Qt::CaseSensitivity>(cs));
}

char QStringRef_Contains3(void* ptr, void* str, long long cs)
{
	return static_cast<QStringRef*>(ptr)->contains(*static_cast<QLatin1String*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

char QStringRef_Contains(void* ptr, struct QtCore_PackedString str, long long cs)
{
	return static_cast<QStringRef*>(ptr)->contains(QString::fromUtf8(str.data, str.len), static_cast<Qt::CaseSensitivity>(cs));
}

char QStringRef_Contains4(void* ptr, void* str, long long cs)
{
	return static_cast<QStringRef*>(ptr)->contains(*static_cast<QStringRef*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

void* QStringRef_Data(void* ptr)
{
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->data());
}

struct QtCore_PackedString QStringRef_String(void* ptr)
{
	return ({ QByteArray* t75a779 = new QByteArray(static_cast<QStringRef*>(ptr)->string()->toUtf8()); QtCore_PackedString { const_cast<char*>(t75a779->prepend("WHITESPACE").constData()+10), t75a779->size()-10, t75a779 }; });
}

int QStringRef_Length(void* ptr)
{
	return static_cast<QStringRef*>(ptr)->length();
}

int QStringRef_Position(void* ptr)
{
	return static_cast<QStringRef*>(ptr)->position();
}

int QStringRef_Size(void* ptr)
{
	return static_cast<QStringRef*>(ptr)->size();
}

void* QStringRef___split_atList2(void* ptr, int i)
{
	return new QStringRef(({QStringRef tmp = static_cast<QVector<QStringRef>*>(ptr)->at(i); if (i == static_cast<QVector<QStringRef>*>(ptr)->size()-1) { static_cast<QVector<QStringRef>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QStringRef___split_setList2(void* ptr, void* i)
{
	static_cast<QVector<QStringRef>*>(ptr)->append(*static_cast<QStringRef*>(i));
}

void* QStringRef___split_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QStringRef>();
}

void* QStringRef___split_atList(void* ptr, int i)
{
	return new QStringRef(({QStringRef tmp = static_cast<QVector<QStringRef>*>(ptr)->at(i); if (i == static_cast<QVector<QStringRef>*>(ptr)->size()-1) { static_cast<QVector<QStringRef>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QStringRef___split_setList(void* ptr, void* i)
{
	static_cast<QVector<QStringRef>*>(ptr)->append(*static_cast<QStringRef*>(i));
}

void* QStringRef___split_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QStringRef>();
}

unsigned int QStringRef___toUcs4_atList(void* ptr, int i)
{
	return ({uint tmp = static_cast<QVector<uint>*>(ptr)->at(i); if (i == static_cast<QVector<uint>*>(ptr)->size()-1) { static_cast<QVector<uint>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QStringRef___toUcs4_setList(void* ptr, unsigned int i)
{
	static_cast<QVector<uint>*>(ptr)->append(i);
}

void* QStringRef___toUcs4_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<uint>();
}

int QSysInfo_WordSize_Type()
{
	return QSysInfo::WordSize;
}

Q_DECLARE_METATYPE(QTime)
Q_DECLARE_METATYPE(QTime*)
void* QTime_NewQTime()
{
	return new QTime();
}

void* QTime_NewQTime3(int h, int m, int s, int ms)
{
	return new QTime(h, m, s, ms);
}

void QTime_Start(void* ptr)
{
	static_cast<QTime*>(ptr)->start();
}

struct QtCore_PackedString QTime_ToString2(void* ptr, long long format)
{
	return ({ QByteArray* tb222d7 = new QByteArray(static_cast<QTime*>(ptr)->toString(static_cast<Qt::DateFormat>(format)).toUtf8()); QtCore_PackedString { const_cast<char*>(tb222d7->prepend("WHITESPACE").constData()+10), tb222d7->size()-10, tb222d7 }; });
}

struct QtCore_PackedString QTime_ToString(void* ptr, struct QtCore_PackedString format)
{
	return ({ QByteArray* t864fb5 = new QByteArray(static_cast<QTime*>(ptr)->toString(QString::fromUtf8(format.data, format.len)).toUtf8()); QtCore_PackedString { const_cast<char*>(t864fb5->prepend("WHITESPACE").constData()+10), t864fb5->size()-10, t864fb5 }; });
}

int QTime_Second(void* ptr)
{
	return static_cast<QTime*>(ptr)->second();
}

Q_DECLARE_METATYPE(QTimeZone)
Q_DECLARE_METATYPE(QTimeZone*)
void* QTimeZone_NewQTimeZone()
{
	return new QTimeZone();
}

void* QTimeZone_NewQTimeZone2(void* ianaId)
{
	return new QTimeZone(*static_cast<QByteArray*>(ianaId));
}

void* QTimeZone_NewQTimeZone4(void* ianaId, int offsetSeconds, struct QtCore_PackedString name, struct QtCore_PackedString abbreviation, long long country, struct QtCore_PackedString comment)
{
	return new QTimeZone(*static_cast<QByteArray*>(ianaId), offsetSeconds, QString::fromUtf8(name.data, name.len), QString::fromUtf8(abbreviation.data, abbreviation.len), static_cast<QLocale::Country>(country), QString::fromUtf8(comment.data, comment.len));
}

void* QTimeZone_NewQTimeZone5(void* other)
{
	return new QTimeZone(*static_cast<QTimeZone*>(other));
}

void* QTimeZone_NewQTimeZone3(int offsetSeconds)
{
	return new QTimeZone(offsetSeconds);
}

void QTimeZone_DestroyQTimeZone(void* ptr)
{
	static_cast<QTimeZone*>(ptr)->~QTimeZone();
}

void* QTimeZone___availableTimeZoneIds_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTimeZone___availableTimeZoneIds_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTimeZone___availableTimeZoneIds_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QTimeZone___availableTimeZoneIds_atList2(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTimeZone___availableTimeZoneIds_setList2(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTimeZone___availableTimeZoneIds_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QTimeZone___availableTimeZoneIds_atList3(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTimeZone___availableTimeZoneIds_setList3(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTimeZone___availableTimeZoneIds_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QTimeZone___windowsIdToIanaIds_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTimeZone___windowsIdToIanaIds_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTimeZone___windowsIdToIanaIds_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QTimeZone___windowsIdToIanaIds_atList2(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTimeZone___windowsIdToIanaIds_setList2(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTimeZone___windowsIdToIanaIds_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

class MyQTimerEvent: public QTimerEvent
{
public:
	MyQTimerEvent(int timerId) : QTimerEvent(timerId) {QTimerEvent_QTimerEvent_QRegisterMetaType();};
};

Q_DECLARE_METATYPE(QTimerEvent*)
Q_DECLARE_METATYPE(MyQTimerEvent*)

int QTimerEvent_QTimerEvent_QRegisterMetaType(){qRegisterMetaType<QTimerEvent*>(); return qRegisterMetaType<MyQTimerEvent*>();}

void* QTimerEvent_NewQTimerEvent(int timerId)
{
	return new MyQTimerEvent(timerId);
}

Q_DECLARE_METATYPE(QUrl)
Q_DECLARE_METATYPE(QUrl*)
void* QUrl_NewQUrl()
{
	return new QUrl();
}

void* QUrl_NewQUrl4(void* other)
{
	return new QUrl(*static_cast<QUrl*>(other));
}

void* QUrl_NewQUrl3(struct QtCore_PackedString url, long long parsingMode)
{
	return new QUrl(QString::fromUtf8(url.data, url.len), static_cast<QUrl::ParsingMode>(parsingMode));
}

void* QUrl_NewQUrl2(void* other)
{
	return new QUrl(*static_cast<QUrl*>(other));
}

void QUrl_DestroyQUrl(void* ptr)
{
	static_cast<QUrl*>(ptr)->~QUrl();
}

struct QtCore_PackedString QUrl_Path(void* ptr, long long options)
{
	return ({ QByteArray* t70ef65 = new QByteArray(static_cast<QUrl*>(ptr)->path(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8()); QtCore_PackedString { const_cast<char*>(t70ef65->prepend("WHITESPACE").constData()+10), t70ef65->size()-10, t70ef65 }; });
}

struct QtCore_PackedString QUrl_ToString(void* ptr, long long options)
{
	return ({ QByteArray* tc0d7b7 = new QByteArray(static_cast<QUrl*>(ptr)->toString(static_cast<QUrl::UrlFormattingOption>(options)).toUtf8()); QtCore_PackedString { const_cast<char*>(tc0d7b7->prepend("WHITESPACE").constData()+10), tc0d7b7->size()-10, tc0d7b7 }; });
}

struct QtCore_PackedString QUrl_Url(void* ptr, long long options)
{
	return ({ QByteArray* t2ea726 = new QByteArray(static_cast<QUrl*>(ptr)->url(static_cast<QUrl::UrlFormattingOption>(options)).toUtf8()); QtCore_PackedString { const_cast<char*>(t2ea726->prepend("WHITESPACE").constData()+10), t2ea726->size()-10, t2ea726 }; });
}

void* QUrl___fromStringList_atList(void* ptr, int i)
{
	return new QUrl(({QUrl tmp = static_cast<QList<QUrl>*>(ptr)->at(i); if (i == static_cast<QList<QUrl>*>(ptr)->size()-1) { static_cast<QList<QUrl>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QUrl___fromStringList_setList(void* ptr, void* i)
{
	static_cast<QList<QUrl>*>(ptr)->append(*static_cast<QUrl*>(i));
}

void* QUrl___fromStringList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QUrl>();
}

void* QUrl___toStringList_urls_atList(void* ptr, int i)
{
	return new QUrl(({QUrl tmp = static_cast<QList<QUrl>*>(ptr)->at(i); if (i == static_cast<QList<QUrl>*>(ptr)->size()-1) { static_cast<QList<QUrl>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QUrl___toStringList_urls_setList(void* ptr, void* i)
{
	static_cast<QList<QUrl>*>(ptr)->append(*static_cast<QUrl*>(i));
}

void* QUrl___toStringList_urls_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QUrl>();
}

void* QUrl___allEncodedQueryItemValues_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QUrl___allEncodedQueryItemValues_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QUrl___allEncodedQueryItemValues_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

Q_DECLARE_METATYPE(QUuid)
Q_DECLARE_METATYPE(QUuid*)
void* QUuid_NewQUuid()
{
	return new QUuid();
}

void* QUuid_NewQUuid5(void* text)
{
	return new QUuid(*static_cast<QByteArray*>(text));
}

void* QUuid_NewQUuid3(struct QtCore_PackedString text)
{
	return new QUuid(QString::fromUtf8(text.data, text.len));
}

void* QUuid_NewQUuid2(unsigned int l, unsigned short w1, unsigned short w2, char* b1, char* b2, char* b3, char* b4, char* b5, char* b6, char* b7, char* b8)
{
	return new QUuid(l, w1, w2, *static_cast<uchar*>(static_cast<void*>(b1)), *static_cast<uchar*>(static_cast<void*>(b2)), *static_cast<uchar*>(static_cast<void*>(b3)), *static_cast<uchar*>(static_cast<void*>(b4)), *static_cast<uchar*>(static_cast<void*>(b5)), *static_cast<uchar*>(static_cast<void*>(b6)), *static_cast<uchar*>(static_cast<void*>(b7)), *static_cast<uchar*>(static_cast<void*>(b8)));
}

struct QtCore_PackedString QUuid_ToString(void* ptr)
{
	return ({ QByteArray* t6f28ec = new QByteArray(static_cast<QUuid*>(ptr)->toString().toUtf8()); QtCore_PackedString { const_cast<char*>(t6f28ec->prepend("WHITESPACE").constData()+10), t6f28ec->size()-10, t6f28ec }; });
}

Q_DECLARE_METATYPE(QVariant*)
void* QVariant_NewQVariant20(void* c)
{
	return new QVariant(*static_cast<QChar*>(c));
}

void* QVariant_NewQVariant()
{
	return new QVariant();
}

void* QVariant_NewQVariant6(void* s)
{
	return new QVariant(*static_cast<QDataStream*>(s));
}

void* QVariant_NewQVariant18(void* val)
{
	return new QVariant(*static_cast<QLatin1String*>(val));
}

void* QVariant_NewQVariant47(void* other)
{
	return new QVariant(*static_cast<QVariant*>(other));
}

void* QVariant_NewQVariant2(long long ty)
{
	return new QVariant(static_cast<QVariant::Type>(ty));
}

void* QVariant_NewQVariant11(char val)
{
	return new QVariant(val != 0);
}

void* QVariant_NewQVariant16(void* val)
{
	return new QVariant(*static_cast<QBitArray*>(val));
}

void* QVariant_NewQVariant15(void* val)
{
	return new QVariant(*static_cast<QByteArray*>(val));
}

void* QVariant_NewQVariant21(void* val)
{
	return new QVariant(*static_cast<QDate*>(val));
}

void* QVariant_NewQVariant23(void* val)
{
	return new QVariant(*static_cast<QDateTime*>(val));
}

void* QVariant_NewQVariant39(void* val)
{
	return new QVariant(*static_cast<QEasingCurve*>(val));
}

void* QVariant_NewQVariant26(void* val)
{
	return new QVariant(({ QHash<QString, QVariant>* tmpP = static_cast<QHash<QString, QVariant>*>(val); QHash<QString, QVariant> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }));
}

void* QVariant_NewQVariant45(void* val)
{
	return new QVariant(*static_cast<QJsonArray*>(val));
}

void* QVariant_NewQVariant46(void* val)
{
	return new QVariant(*static_cast<QJsonDocument*>(val));
}

void* QVariant_NewQVariant44(void* val)
{
	return new QVariant(*static_cast<QJsonObject*>(val));
}

void* QVariant_NewQVariant43(void* val)
{
	return new QVariant(*static_cast<QJsonValue*>(val));
}

void* QVariant_NewQVariant31(void* val)
{
	return new QVariant(*static_cast<QLine*>(val));
}

void* QVariant_NewQVariant32(void* val)
{
	return new QVariant(*static_cast<QLineF*>(val));
}

void* QVariant_NewQVariant24(void* val)
{
	return new QVariant(({ QList<QVariant>* tmpP = static_cast<QList<QVariant>*>(val); QList<QVariant> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void* QVariant_NewQVariant35(void* l)
{
	return new QVariant(*static_cast<QLocale*>(l));
}

void* QVariant_NewQVariant25(void* val)
{
	return new QVariant(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(val); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void* QVariant_NewQVariant41(void* val)
{
	return new QVariant(*static_cast<QModelIndex*>(val));
}

void* QVariant_NewQVariant42(void* val)
{
	return new QVariant(*static_cast<QPersistentModelIndex*>(val));
}

void* QVariant_NewQVariant29(void* val)
{
	return new QVariant(*static_cast<QPoint*>(val));
}

void* QVariant_NewQVariant30(void* val)
{
	return new QVariant(*static_cast<QPointF*>(val));
}

void* QVariant_NewQVariant33(void* val)
{
	return new QVariant(*static_cast<QRect*>(val));
}

void* QVariant_NewQVariant34(void* val)
{
	return new QVariant(*static_cast<QRectF*>(val));
}

void* QVariant_NewQVariant36(void* regExp)
{
	return new QVariant(*static_cast<QRegExp*>(regExp));
}

void* QVariant_NewQVariant37(void* re)
{
	return new QVariant(*static_cast<QRegularExpression*>(re));
}

void* QVariant_NewQVariant27(void* val)
{
	return new QVariant(*static_cast<QSize*>(val));
}

void* QVariant_NewQVariant28(void* val)
{
	return new QVariant(*static_cast<QSizeF*>(val));
}

void* QVariant_NewQVariant17(struct QtCore_PackedString val)
{
	return new QVariant(QString::fromUtf8(val.data, val.len));
}

void* QVariant_NewQVariant19(struct QtCore_PackedString val)
{
	return new QVariant(QString::fromUtf8(val.data, val.len).split("¡¦!", QString::SkipEmptyParts));
}

void* QVariant_NewQVariant22(void* val)
{
	return new QVariant(*static_cast<QTime*>(val));
}

void* QVariant_NewQVariant38(void* val)
{
	return new QVariant(*static_cast<QUrl*>(val));
}

void* QVariant_NewQVariant40(void* val)
{
	return new QVariant(*static_cast<QUuid*>(val));
}

void* QVariant_NewQVariant5(void* p)
{
	return new QVariant(*static_cast<QVariant*>(p));
}

void* QVariant_NewQVariant14(char* val)
{
	return new QVariant(const_cast<const char*>(val));
}

void* QVariant_NewQVariant12(double val)
{
	return new QVariant(val);
}

void* QVariant_NewQVariant13(float val)
{
	return new QVariant(val);
}

void* QVariant_NewQVariant3(int typeId, void* copy)
{
	return new QVariant(typeId, copy);
}

void* QVariant_NewQVariant7(int val)
{
	return new QVariant(val);
}

void* QVariant_NewQVariant9(long long val)
{
	return new QVariant(val);
}

void* QVariant_NewQVariant10(unsigned long long val)
{
	return new QVariant(val);
}

void* QVariant_NewQVariant8(unsigned int val)
{
	return new QVariant(val);
}

char QVariant_Convert(void* ptr, int targetTypeId)
{
	return static_cast<QVariant*>(ptr)->convert(targetTypeId);
}

void QVariant_DestroyQVariant(void* ptr)
{
	static_cast<QVariant*>(ptr)->~QVariant();
}

struct QtCore_PackedList QVariant_ToHash(void* ptr)
{
	return ({ QHash<QString, QVariant>* tmpValue00701e = new QHash<QString, QVariant>(static_cast<QVariant*>(ptr)->toHash()); QtCore_PackedList { tmpValue00701e, tmpValue00701e->size() }; });
}

struct QtCore_PackedList QVariant_ToList(void* ptr)
{
	return ({ QList<QVariant>* tmpValue8f6950 = new QList<QVariant>(static_cast<QVariant*>(ptr)->toList()); QtCore_PackedList { tmpValue8f6950, tmpValue8f6950->size() }; });
}

struct QtCore_PackedList QVariant_ToMap(void* ptr)
{
	return ({ QMap<QString, QVariant>* tmpValue1e0d76 = new QMap<QString, QVariant>(static_cast<QVariant*>(ptr)->toMap()); QtCore_PackedList { tmpValue1e0d76, tmpValue1e0d76->size() }; });
}

struct QtCore_PackedString QVariant_ToString(void* ptr)
{
	return ({ QByteArray* tf9e1e4 = new QByteArray(static_cast<QVariant*>(ptr)->toString().toUtf8()); QtCore_PackedString { const_cast<char*>(tf9e1e4->prepend("WHITESPACE").constData()+10), tf9e1e4->size()-10, tf9e1e4 }; });
}

struct QtCore_PackedString QVariant_ToStringList(void* ptr)
{
	return ({ QByteArray* tf99cb6 = new QByteArray(static_cast<QVariant*>(ptr)->toStringList().join("¡¦!").toUtf8()); QtCore_PackedString { const_cast<char*>(tf99cb6->prepend("WHITESPACE").constData()+10), tf99cb6->size()-10, tf99cb6 }; });
}

long long QVariant_Type(void* ptr)
{
	return static_cast<QVariant*>(ptr)->type();
}

char QVariant_CanConvert(void* ptr, int targetTypeId)
{
	return static_cast<QVariant*>(ptr)->canConvert(targetTypeId);
}

char QVariant_IsNull(void* ptr)
{
	return static_cast<QVariant*>(ptr)->isNull();
}

char QVariant_IsValid(void* ptr)
{
	return static_cast<QVariant*>(ptr)->isValid();
}

char QVariant_ToBool(void* ptr)
{
	return static_cast<QVariant*>(ptr)->toBool();
}

double QVariant_ToDouble(void* ptr, char* ok)
{
	return static_cast<QVariant*>(ptr)->toDouble(reinterpret_cast<bool*>(ok));
}

int QVariant_ToInt(void* ptr, char* ok)
{
	return static_cast<QVariant*>(ptr)->toInt(reinterpret_cast<bool*>(ok));
}

long long QVariant_ToLongLong(void* ptr, char* ok)
{
	return static_cast<QVariant*>(ptr)->toLongLong(reinterpret_cast<bool*>(ok));
}

unsigned long long QVariant_ToULongLong(void* ptr, char* ok)
{
	return static_cast<QVariant*>(ptr)->toULongLong(reinterpret_cast<bool*>(ok));
}

unsigned int QVariant_ToUInt(void* ptr, char* ok)
{
	return static_cast<QVariant*>(ptr)->toUInt(reinterpret_cast<bool*>(ok));
}

void* QVariant___QVariant_val_atList26(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QHash<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QHash<QString, QVariant>*>(ptr)->size()-1) { static_cast<QHash<QString, QVariant>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QVariant___QVariant_val_setList26(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QHash<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QVariant___QVariant_val_newList26(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<QString, QVariant>();
}

struct QtCore_PackedList QVariant___QVariant_val_keyList26(void* ptr)
{
	return ({ QList<QString>* tmpValuef43bc5 = new QList<QString>(static_cast<QHash<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValuef43bc5, tmpValuef43bc5->size() }; });
}

void* QVariant___QVariant_val_atList24(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVariant___QVariant_val_setList24(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QVariant___QVariant_val_newList24(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QVariant___QVariant_val_atList25(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QVariant___QVariant_val_setList25(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QVariant___QVariant_val_newList25(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtCore_PackedList QVariant___QVariant_val_keyList25(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

void* QVariant___toHash_atList(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QHash<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QHash<QString, QVariant>*>(ptr)->size()-1) { static_cast<QHash<QString, QVariant>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QVariant___toHash_setList(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QHash<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QVariant___toHash_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<QString, QVariant>();
}

struct QtCore_PackedList QVariant___toHash_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValuef43bc5 = new QList<QString>(static_cast<QHash<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValuef43bc5, tmpValuef43bc5->size() }; });
}

void* QVariant___toList_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVariant___toList_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QVariant___toList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QVariant___toMap_atList(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QVariant___toMap_setList(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QVariant___toMap_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtCore_PackedList QVariant___toMap_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct QtCore_PackedString QVariant_____QVariant_val_keyList_atList26(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtCore_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QVariant_____QVariant_val_keyList_setList26(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QVariant_____QVariant_val_keyList_newList26(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtCore_PackedString QVariant_____QVariant_val_keyList_atList25(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtCore_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QVariant_____QVariant_val_keyList_setList25(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QVariant_____QVariant_val_keyList_newList25(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtCore_PackedString QVariant_____toHash_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtCore_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QVariant_____toHash_keyList_setList(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QVariant_____toHash_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtCore_PackedString QVariant_____toMap_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtCore_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QVariant_____toMap_keyList_setList(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QVariant_____toMap_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

int Qt_LocaleDate_Type()
{
	return Qt::LocaleDate;
}

int Qt_SystemLocaleShortDate_Type()
{
	return Qt::SystemLocaleShortDate;
}

int Qt_SystemLocaleLongDate_Type()
{
	return Qt::SystemLocaleLongDate;
}

int Qt_DefaultLocaleShortDate_Type()
{
	return Qt::DefaultLocaleShortDate;
}

int Qt_DefaultLocaleLongDate_Type()
{
	return Qt::DefaultLocaleLongDate;
}

int Qt_RFC2822Date_Type()
{
	return Qt::RFC2822Date;
}

int Qt_ISODateWithMs_Type()
{
	#if QT_VERSION >= 0x056000
		return Qt::ISODateWithMs;
	#else
		return 0;
	#endif
}

int Qt_LastGestureType_Type()
{
	return Qt::LastGestureType;
}

