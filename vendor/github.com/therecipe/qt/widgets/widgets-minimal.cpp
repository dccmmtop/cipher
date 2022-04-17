// +build minimal

#define protected public
#define private public

#include "widgets-minimal.h"
#include "_cgo_export.h"

#include <QAbstractButton>
#include <QAbstractScrollArea>
#include <QAction>
#include <QApplication>
#include <QBoxLayout>
#include <QButtonGroup>
#include <QByteArray>
#include <QDirModel>
#include <QFileSystemModel>
#include <QFrame>
#include <QGraphicsItem>
#include <QGraphicsItemGroup>
#include <QGraphicsLayout>
#include <QGraphicsLayoutItem>
#include <QGraphicsObject>
#include <QGraphicsTransform>
#include <QGraphicsWidget>
#include <QGroupBox>
#include <QHBoxLayout>
#include <QIcon>
#include <QKeySequence>
#include <QLabel>
#include <QLayout>
#include <QLayoutItem>
#include <QList>
#include <QMargins>
#include <QMatrix4x4>
#include <QMessageBox>
#include <QMetaObject>
#include <QObject>
#include <QPagedPaintDevice>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPoint>
#include <QPointF>
#include <QPushButton>
#include <QRect>
#include <QRectF>
#include <QScreen>
#include <QSize>
#include <QSizeF>
#include <QString>
#include <QStyle>
#include <QStyleHintReturn>
#include <QStyleOption>
#include <QStyleOptionGraphicsItem>
#include <QTextEdit>
#include <QVBoxLayout>
#include <QVariant>
#include <QWidget>
#include <QStringList>

class MyQAbstractButton: public QAbstractButton
{
public:
	MyQAbstractButton(QWidget *parent = Q_NULLPTR) : QAbstractButton(parent) {QAbstractButton_QAbstractButton_QRegisterMetaType();};
	void click() { callbackQAbstractButton_Click(this); };
	void Signal_Clicked(bool checked) { callbackQAbstractButton_Clicked(this, checked); };
	void paintEvent(QPaintEvent * e) { callbackQAbstractButton_PaintEvent(this, e); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void lower() { callbackQWidget_Lower(this); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QAbstractButton*)
Q_DECLARE_METATYPE(MyQAbstractButton*)

int QAbstractButton_QAbstractButton_QRegisterMetaType(){qRegisterMetaType<QAbstractButton*>(); return qRegisterMetaType<MyQAbstractButton*>();}

void* QAbstractButton_NewQAbstractButton(void* parent)
{
		return new MyQAbstractButton(static_cast<QWidget*>(parent));
}

void QAbstractButton_Click(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractButton*>(ptr), "click");
}

void QAbstractButton_ClickDefault(void* ptr)
{
	if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::click();
	} else {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::click();
	}
}

void QAbstractButton_ConnectClicked(void* ptr, long long t)
{
	QObject::connect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)(bool)>(&QAbstractButton::clicked), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)(bool)>(&MyQAbstractButton::Signal_Clicked), static_cast<Qt::ConnectionType>(t));
}

void QAbstractButton_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)(bool)>(&QAbstractButton::clicked), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)(bool)>(&MyQAbstractButton::Signal_Clicked));
}

void QAbstractButton_Clicked(void* ptr, char checked)
{
	static_cast<QAbstractButton*>(ptr)->clicked(checked != 0);
}

void QAbstractButton_PaintEvent(void* ptr, void* e)
{
	static_cast<QAbstractButton*>(ptr)->paintEvent(static_cast<QPaintEvent*>(e));
}

void QAbstractButton_SetText(void* ptr, struct QtWidgets_PackedString text)
{
	static_cast<QAbstractButton*>(ptr)->setText(QString::fromUtf8(text.data, text.len));
}

void QAbstractButton_DestroyQAbstractButton(void* ptr)
{
	static_cast<QAbstractButton*>(ptr)->~QAbstractButton();
}

void* QAbstractButton_Group(void* ptr)
{
	return static_cast<QAbstractButton*>(ptr)->group();
}

struct QtWidgets_PackedString QAbstractButton_Text(void* ptr)
{
	return ({ QByteArray* t8ae7db = new QByteArray(static_cast<QAbstractButton*>(ptr)->text().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t8ae7db->prepend("WHITESPACE").constData()+10), t8ae7db->size()-10, t8ae7db }; });
}

class MyQAbstractScrollArea: public QAbstractScrollArea
{
public:
	MyQAbstractScrollArea(QWidget *parent = Q_NULLPTR) : QAbstractScrollArea(parent) {QAbstractScrollArea_QAbstractScrollArea_QRegisterMetaType();};
	bool close() { return callbackQWidget_Close(this) != 0; };
	void lower() { callbackQWidget_Lower(this); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QAbstractScrollArea*)
Q_DECLARE_METATYPE(MyQAbstractScrollArea*)

int QAbstractScrollArea_QAbstractScrollArea_QRegisterMetaType(){qRegisterMetaType<QAbstractScrollArea*>(); return qRegisterMetaType<MyQAbstractScrollArea*>();}

void* QAbstractScrollArea_NewQAbstractScrollArea(void* parent)
{
		return new MyQAbstractScrollArea(static_cast<QWidget*>(parent));
}

void QAbstractScrollArea_DestroyQAbstractScrollArea(void* ptr)
{
	static_cast<QAbstractScrollArea*>(ptr)->~QAbstractScrollArea();
}

void* QAbstractScrollArea___scrollBarWidgets_atList(void* ptr, int i)
{
	return ({QWidget * tmp = static_cast<QList<QWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QWidget *>*>(ptr)->size()-1) { static_cast<QList<QWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractScrollArea___scrollBarWidgets_setList(void* ptr, void* i)
{
		static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* QAbstractScrollArea___scrollBarWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWidget *>();
}

class MyQAction: public QAction
{
public:
	MyQAction(QObject *parent = Q_NULLPTR) : QAction(parent) {QAction_QAction_QRegisterMetaType();};
	MyQAction(const QIcon &icon, const QString &text, QObject *parent = Q_NULLPTR) : QAction(icon, text, parent) {QAction_QAction_QRegisterMetaType();};
	MyQAction(const QString &text, QObject *parent = Q_NULLPTR) : QAction(text, parent) {QAction_QAction_QRegisterMetaType();};
	void setEnabled(bool vbo) { callbackQAction_SetEnabled(this, vbo); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQAction_ObjectNameChanged(this, objectNamePacked); };
};

Q_DECLARE_METATYPE(QAction*)
Q_DECLARE_METATYPE(MyQAction*)

int QAction_QAction_QRegisterMetaType(){qRegisterMetaType<QAction*>(); return qRegisterMetaType<MyQAction*>();}

void* QAction_NewQAction(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQAction(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAction(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQAction(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAction(static_cast<QWidget*>(parent));
	} else {
		return new MyQAction(static_cast<QObject*>(parent));
	}
}

void* QAction_NewQAction3(void* icon, struct QtWidgets_PackedString text, void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(parent));
	} else {
		return new MyQAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QObject*>(parent));
	}
}

void* QAction_NewQAction2(struct QtWidgets_PackedString text, void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQAction(QString::fromUtf8(text.data, text.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAction(QString::fromUtf8(text.data, text.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQAction(QString::fromUtf8(text.data, text.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAction(QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(parent));
	} else {
		return new MyQAction(QString::fromUtf8(text.data, text.len), static_cast<QObject*>(parent));
	}
}

void QAction_SetEnabled(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QAction_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QAction*>(ptr)->QAction::setEnabled(vbo != 0);
}

void QAction_SetText(void* ptr, struct QtWidgets_PackedString text)
{
	static_cast<QAction*>(ptr)->setText(QString::fromUtf8(text.data, text.len));
}

void QAction_DestroyQAction(void* ptr)
{
	static_cast<QAction*>(ptr)->~QAction();
}

struct QtWidgets_PackedString QAction_Text(void* ptr)
{
	return ({ QByteArray* t0f57fa = new QByteArray(static_cast<QAction*>(ptr)->text().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t0f57fa->prepend("WHITESPACE").constData()+10), t0f57fa->size()-10, t0f57fa }; });
}

void* QAction_Data(void* ptr)
{
	return new QVariant(static_cast<QAction*>(ptr)->data());
}

void* QAction___setShortcuts_shortcuts_atList(void* ptr, int i)
{
	return new QKeySequence(({QKeySequence tmp = static_cast<QList<QKeySequence>*>(ptr)->at(i); if (i == static_cast<QList<QKeySequence>*>(ptr)->size()-1) { static_cast<QList<QKeySequence>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAction___setShortcuts_shortcuts_setList(void* ptr, void* i)
{
	static_cast<QList<QKeySequence>*>(ptr)->append(*static_cast<QKeySequence*>(i));
}

void* QAction___setShortcuts_shortcuts_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QKeySequence>();
}

void* QAction___associatedGraphicsWidgets_atList(void* ptr, int i)
{
	return ({QGraphicsWidget * tmp = static_cast<QList<QGraphicsWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsWidget *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAction___associatedGraphicsWidgets_setList(void* ptr, void* i)
{
		static_cast<QList<QGraphicsWidget *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
}

void* QAction___associatedGraphicsWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsWidget *>();
}

void* QAction___shortcuts_atList(void* ptr, int i)
{
	return new QKeySequence(({QKeySequence tmp = static_cast<QList<QKeySequence>*>(ptr)->at(i); if (i == static_cast<QList<QKeySequence>*>(ptr)->size()-1) { static_cast<QList<QKeySequence>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAction___shortcuts_setList(void* ptr, void* i)
{
	static_cast<QList<QKeySequence>*>(ptr)->append(*static_cast<QKeySequence*>(i));
}

void* QAction___shortcuts_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QKeySequence>();
}

void* QAction___associatedWidgets_atList(void* ptr, int i)
{
	return ({QWidget * tmp = static_cast<QList<QWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QWidget *>*>(ptr)->size()-1) { static_cast<QList<QWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAction___associatedWidgets_setList(void* ptr, void* i)
{
		static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* QAction___associatedWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWidget *>();
}

void* QAction___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAction___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QAction___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QAction___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAction___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QAction___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAction___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAction___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QAction___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAction___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAction___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QAction___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAction___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAction___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QAction___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

class MyQApplication: public QApplication
{
public:
	MyQApplication(int &argc, char **argv) : QApplication(argc, argv) {QApplication_QApplication_QRegisterMetaType();};
	 ~MyQApplication() { callbackQApplication_DestroyQApplication(this); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQApplication_ObjectNameChanged(this, objectNamePacked); };
};

Q_DECLARE_METATYPE(QApplication*)
Q_DECLARE_METATYPE(MyQApplication*)

int QApplication_QApplication_QRegisterMetaType(){qRegisterMetaType<QApplication*>(); return qRegisterMetaType<MyQApplication*>();}

void* QApplication_NewQApplication(int argc, char* argv)
{
	static int argcs = argc;
	static char** argvs = static_cast<char**>(malloc(argcs * sizeof(char*)));

	QList<QByteArray> aList = QByteArray(argv).split('|');
	for (int i = 0; i < argcs; i++)
		argvs[i] = (new QByteArray(aList.at(i)))->data();

	return new MyQApplication(argcs, argvs);
}

int QApplication_QApplication_Exec()
{
	return QApplication::exec();
}

void QApplication_DestroyQApplication(void* ptr)
{
	static_cast<QApplication*>(ptr)->~QApplication();
}

void QApplication_DestroyQApplicationDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QApplication___allWidgets_atList(void* ptr, int i)
{
	return ({QWidget * tmp = static_cast<QList<QWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QWidget *>*>(ptr)->size()-1) { static_cast<QList<QWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QApplication___allWidgets_setList(void* ptr, void* i)
{
		static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* QApplication___allWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWidget *>();
}

void* QApplication___topLevelWidgets_atList(void* ptr, int i)
{
	return ({QWidget * tmp = static_cast<QList<QWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QWidget *>*>(ptr)->size()-1) { static_cast<QList<QWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QApplication___topLevelWidgets_setList(void* ptr, void* i)
{
		static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* QApplication___topLevelWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWidget *>();
}

void* QApplication___screens_atList(void* ptr, int i)
{
	return ({QScreen * tmp = static_cast<QList<QScreen *>*>(ptr)->at(i); if (i == static_cast<QList<QScreen *>*>(ptr)->size()-1) { static_cast<QList<QScreen *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QApplication___screens_setList(void* ptr, void* i)
{
	static_cast<QList<QScreen *>*>(ptr)->append(static_cast<QScreen*>(i));
}

void* QApplication___screens_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QScreen *>();
}

void* QApplication___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QApplication___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QApplication___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QApplication___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QApplication___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QApplication___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QApplication___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QApplication___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QApplication___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QApplication___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QApplication___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QApplication___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QApplication___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QApplication___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QApplication___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

class MyQBoxLayout: public QBoxLayout
{
public:
	MyQBoxLayout(Direction dir, QWidget *parent = Q_NULLPTR) : QBoxLayout(dir, parent) {QBoxLayout_QBoxLayout_QRegisterMetaType();};
	QSize minimumSize() const { return *static_cast<QSize*>(callbackQLayout_MinimumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	QLayout * layout() { return static_cast<QLayout*>(callbackQLayoutItem_Layout(this)); };
	QLayoutItem * takeAt(int index) { return static_cast<QLayoutItem*>(callbackQBoxLayout_TakeAt(this, index)); };
	void addItem(QLayoutItem * item) { callbackQBoxLayout_AddItem(this, item); };
	QLayoutItem * itemAt(int index) const { return static_cast<QLayoutItem*>(callbackQBoxLayout_ItemAt(const_cast<void*>(static_cast<const void*>(this)), index)); };
	int count() const { return callbackQBoxLayout_Count(const_cast<void*>(static_cast<const void*>(this))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQLayout_ObjectNameChanged(this, objectNamePacked); };
	QWidget * widget() { return static_cast<QWidget*>(callbackQLayoutItem_Widget(this)); };
	void setGeometry(const QRect & r) { callbackQLayout_SetGeometry(this, const_cast<QRect*>(&r)); };
	QRect geometry() const { return *static_cast<QRect*>(callbackQLayout_Geometry(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize maximumSize() const { return *static_cast<QSize*>(callbackQLayout_MaximumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQLayout_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::Orientations expandingDirections() const { return static_cast<Qt::Orientation>(callbackQLayout_ExpandingDirections(const_cast<void*>(static_cast<const void*>(this)))); };
	bool isEmpty() const { return callbackQLayout_IsEmpty(const_cast<void*>(static_cast<const void*>(this))) != 0; };
};

Q_DECLARE_METATYPE(QBoxLayout*)
Q_DECLARE_METATYPE(MyQBoxLayout*)

int QBoxLayout_QBoxLayout_QRegisterMetaType(){qRegisterMetaType<QBoxLayout*>(); return qRegisterMetaType<MyQBoxLayout*>();}

void* QBoxLayout_NewQBoxLayout(long long dir, void* parent)
{
		return new MyQBoxLayout(static_cast<QBoxLayout::Direction>(dir), static_cast<QWidget*>(parent));
}

void QBoxLayout_AddWidget(void* ptr, void* widget, int stretch, long long alignment)
{
		static_cast<QBoxLayout*>(ptr)->addWidget(static_cast<QWidget*>(widget), stretch, static_cast<Qt::AlignmentFlag>(alignment));
}

void QBoxLayout_DestroyQBoxLayout(void* ptr)
{
	static_cast<QBoxLayout*>(ptr)->~QBoxLayout();
}

void* QBoxLayout_TakeAt(void* ptr, int index)
{
	return static_cast<QBoxLayout*>(ptr)->takeAt(index);
}

void* QBoxLayout_TakeAtDefault(void* ptr, int index)
{
	if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::takeAt(index);
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::takeAt(index);
	} else {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::takeAt(index);
	}
}

void QBoxLayout_AddItem(void* ptr, void* item)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(item))) {
		static_cast<QBoxLayout*>(ptr)->addItem(static_cast<QLayout*>(item));
	} else {
		static_cast<QBoxLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item));
	}
}

void QBoxLayout_AddItemDefault(void* ptr, void* item)
{
	if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QLayout*>(static_cast<QObject*>(item))) {
			static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::addItem(static_cast<QLayout*>(item));
		} else {
			static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::addItem(static_cast<QLayoutItem*>(item));
		}
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QLayout*>(static_cast<QObject*>(item))) {
			static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::addItem(static_cast<QLayout*>(item));
		} else {
			static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::addItem(static_cast<QLayoutItem*>(item));
		}
	} else {
		if (dynamic_cast<QLayout*>(static_cast<QObject*>(item))) {
			static_cast<QBoxLayout*>(ptr)->QBoxLayout::addItem(static_cast<QLayout*>(item));
		} else {
			static_cast<QBoxLayout*>(ptr)->QBoxLayout::addItem(static_cast<QLayoutItem*>(item));
		}
	}
}

void* QBoxLayout_ItemAt(void* ptr, int index)
{
	return static_cast<QBoxLayout*>(ptr)->itemAt(index);
}

void* QBoxLayout_ItemAtDefault(void* ptr, int index)
{
	if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::itemAt(index);
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::itemAt(index);
	} else {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::itemAt(index);
	}
}

int QBoxLayout_Count(void* ptr)
{
	return static_cast<QBoxLayout*>(ptr)->count();
}

int QBoxLayout_CountDefault(void* ptr)
{
	if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::count();
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::count();
	} else {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::count();
	}
}

class MyQButtonGroup: public QButtonGroup
{
public:
	MyQButtonGroup(QObject *parent = Q_NULLPTR) : QButtonGroup(parent) {QButtonGroup_QButtonGroup_QRegisterMetaType();};
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQButtonGroup_ObjectNameChanged(this, objectNamePacked); };
};

Q_DECLARE_METATYPE(QButtonGroup*)
Q_DECLARE_METATYPE(MyQButtonGroup*)

int QButtonGroup_QButtonGroup_QRegisterMetaType(){qRegisterMetaType<QButtonGroup*>(); return qRegisterMetaType<MyQButtonGroup*>();}

void* QButtonGroup_NewQButtonGroup(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQButtonGroup(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQButtonGroup(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQButtonGroup(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQButtonGroup(static_cast<QWidget*>(parent));
	} else {
		return new MyQButtonGroup(static_cast<QObject*>(parent));
	}
}

void QButtonGroup_DestroyQButtonGroup(void* ptr)
{
	static_cast<QButtonGroup*>(ptr)->~QButtonGroup();
}

void* QButtonGroup_Button(void* ptr, int id)
{
	return static_cast<QButtonGroup*>(ptr)->button(id);
}

void* QButtonGroup___buttons_atList(void* ptr, int i)
{
	return ({QAbstractButton * tmp = static_cast<QList<QAbstractButton *>*>(ptr)->at(i); if (i == static_cast<QList<QAbstractButton *>*>(ptr)->size()-1) { static_cast<QList<QAbstractButton *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QButtonGroup___buttons_setList(void* ptr, void* i)
{
	static_cast<QList<QAbstractButton *>*>(ptr)->append(static_cast<QAbstractButton*>(i));
}

void* QButtonGroup___buttons_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAbstractButton *>();
}

void* QButtonGroup___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QButtonGroup___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QButtonGroup___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QButtonGroup___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QButtonGroup___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QButtonGroup___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QButtonGroup___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QButtonGroup___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QButtonGroup___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QButtonGroup___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QButtonGroup___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QButtonGroup___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QButtonGroup___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QButtonGroup___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QButtonGroup___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

int QDirModel_FilePathRole_Type()
{
	return QDirModel::FilePathRole;
}

int QDirModel_FileNameRole_Type()
{
	return QDirModel::FileNameRole;
}

int QFileSystemModel_FilePathRole_Type()
{
	return QFileSystemModel::FilePathRole;
}

int QFileSystemModel_FileNameRole_Type()
{
	return QFileSystemModel::FileNameRole;
}

int QFileSystemModel_FilePermissions_Type()
{
	return QFileSystemModel::FilePermissions;
}

class MyQFrame: public QFrame
{
public:
	MyQFrame(QWidget *parent = Q_NULLPTR, Qt::WindowFlags ff = Qt::WindowFlags()) : QFrame(parent, ff) {QFrame_QFrame_QRegisterMetaType();};
	bool close() { return callbackQWidget_Close(this) != 0; };
	void lower() { callbackQWidget_Lower(this); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QFrame*)
Q_DECLARE_METATYPE(MyQFrame*)

int QFrame_QFrame_QRegisterMetaType(){qRegisterMetaType<QFrame*>(); return qRegisterMetaType<MyQFrame*>();}

void* QFrame_NewQFrame(void* parent, long long ff)
{
		return new MyQFrame(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(ff));
}

void QFrame_DestroyQFrame(void* ptr)
{
	static_cast<QFrame*>(ptr)->~QFrame();
}

class MyQGraphicsItem: public QGraphicsItem
{
public:
	MyQGraphicsItem(QGraphicsItem *parent = Q_NULLPTR) : QGraphicsItem(parent) {QGraphicsItem_QGraphicsItem_QRegisterMetaType();};
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQGraphicsItem_Paint(this, painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	 ~MyQGraphicsItem() { callbackQGraphicsItem_DestroyQGraphicsItem(this); };
	QRectF boundingRect() const { return *static_cast<QRectF*>(callbackQGraphicsItem_BoundingRect(const_cast<void*>(static_cast<const void*>(this)))); };
	bool contains(const QPointF & point) const { return callbackQGraphicsItem_Contains(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPointF*>(&point)) != 0; };
	int type() const { return callbackQGraphicsItem_Type(const_cast<void*>(static_cast<const void*>(this))); };
};

Q_DECLARE_METATYPE(MyQGraphicsItem*)

int QGraphicsItem_QGraphicsItem_QRegisterMetaType(){qRegisterMetaType<QGraphicsItem*>(); return qRegisterMetaType<MyQGraphicsItem*>();}

void* QGraphicsItem_NewQGraphicsItem(void* parent)
{
	return new MyQGraphicsItem(static_cast<QGraphicsItem*>(parent));
}

void QGraphicsItem_Paint(void* ptr, void* painter, void* option, void* widget)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
	} else {
		static_cast<QGraphicsItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
	}
}

void QGraphicsItem_SetEnabled(void* ptr, char enabled)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->setEnabled(enabled != 0);
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setEnabled(enabled != 0);
	} else {
		static_cast<QGraphicsItem*>(ptr)->setEnabled(enabled != 0);
	}
}

void QGraphicsItem_Show(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->show();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->show();
	} else {
		static_cast<QGraphicsItem*>(ptr)->show();
	}
}

void QGraphicsItem_DestroyQGraphicsItem(void* ptr)
{
	static_cast<QGraphicsItem*>(ptr)->~QGraphicsItem();
}

void QGraphicsItem_DestroyQGraphicsItemDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QGraphicsItem_Group(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsObject*>(ptr)->group();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->group();
	} else {
		return static_cast<QGraphicsItem*>(ptr)->group();
	}
}

void* QGraphicsItem_Window(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsObject*>(ptr)->window();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->window();
	} else {
		return static_cast<QGraphicsItem*>(ptr)->window();
	}
}

void* QGraphicsItem_Pos(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return ({ QPointF tmpValue = static_cast<QGraphicsObject*>(ptr)->pos(); new QPointF(tmpValue.x(), tmpValue.y()); });
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({ QPointF tmpValue = static_cast<QGraphicsWidget*>(ptr)->pos(); new QPointF(tmpValue.x(), tmpValue.y()); });
	} else {
		return ({ QPointF tmpValue = static_cast<QGraphicsItem*>(ptr)->pos(); new QPointF(tmpValue.x(), tmpValue.y()); });
	}
}

void* QGraphicsItem_BoundingRect(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return ({ QRectF tmpValue = static_cast<QGraphicsObject*>(ptr)->boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({ QRectF tmpValue = static_cast<QGraphicsWidget*>(ptr)->boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QRectF tmpValue = static_cast<QGraphicsItem*>(ptr)->boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	}
}

void* QGraphicsItem_Data(void* ptr, int key)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QGraphicsObject*>(ptr)->data(key));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QGraphicsWidget*>(ptr)->data(key));
	} else {
		return new QVariant(static_cast<QGraphicsItem*>(ptr)->data(key));
	}
}

char QGraphicsItem_Contains(void* ptr, void* point)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsObject*>(ptr)->contains(*static_cast<QPointF*>(point));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->contains(*static_cast<QPointF*>(point));
	} else {
		return static_cast<QGraphicsItem*>(ptr)->contains(*static_cast<QPointF*>(point));
	}
}

char QGraphicsItem_ContainsDefault(void* ptr, void* point)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::contains(*static_cast<QPointF*>(point));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::contains(*static_cast<QPointF*>(point));
	} else if (dynamic_cast<QGraphicsItemGroup*>(static_cast<QGraphicsItem*>(ptr))) {
		return static_cast<QGraphicsItemGroup*>(ptr)->QGraphicsItemGroup::contains(*static_cast<QPointF*>(point));
	} else {
		return static_cast<QGraphicsItem*>(ptr)->QGraphicsItem::contains(*static_cast<QPointF*>(point));
	}
}

int QGraphicsItem_Type(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsObject*>(ptr)->type();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->type();
	} else {
		return static_cast<QGraphicsItem*>(ptr)->type();
	}
}

int QGraphicsItem_TypeDefault(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::type();
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::type();
	} else if (dynamic_cast<QGraphicsItemGroup*>(static_cast<QGraphicsItem*>(ptr))) {
		return static_cast<QGraphicsItemGroup*>(ptr)->QGraphicsItemGroup::type();
	} else {
		return static_cast<QGraphicsItem*>(ptr)->QGraphicsItem::type();
	}
}

double QGraphicsItem_X(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsObject*>(ptr)->x();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->x();
	} else {
		return static_cast<QGraphicsItem*>(ptr)->x();
	}
}

double QGraphicsItem_Y(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsObject*>(ptr)->y();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->y();
	} else {
		return static_cast<QGraphicsItem*>(ptr)->y();
	}
}

void* QGraphicsItem___setTransformations_transformations_atList(void* ptr, int i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsTransform * tmp = static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsTransform *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsTransform *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsTransform * tmp = static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsTransform *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsTransform *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else {
		return ({QGraphicsTransform * tmp = static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsTransform *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsTransform *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	}
}

void QGraphicsItem___setTransformations_transformations_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
	} else {
		static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
	}
}

void* QGraphicsItem___setTransformations_transformations_newList(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsTransform *>();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsTransform *>();
	} else {
		return new QList<QGraphicsTransform *>();
	}
}

void* QGraphicsItem___childItems_atList(void* ptr, int i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else {
		return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	}
}

void QGraphicsItem___childItems_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
		}
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
		}
	} else {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
		}
	}
}

void* QGraphicsItem___childItems_newList(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsItem *>();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsItem *>();
	} else {
		return new QList<QGraphicsItem *>();
	}
}

void* QGraphicsItem___children_atList(void* ptr, int i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else {
		return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	}
}

void QGraphicsItem___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
		}
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
		}
	} else {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
		}
	}
}

void* QGraphicsItem___children_newList(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsItem *>();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsItem *>();
	} else {
		return new QList<QGraphicsItem *>();
	}
}

void* QGraphicsItem___collidingItems_atList(void* ptr, int i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else {
		return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	}
}

void QGraphicsItem___collidingItems_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
		}
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
		}
	} else {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
		}
	}
}

void* QGraphicsItem___collidingItems_newList(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsItem *>();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsItem *>();
	} else {
		return new QList<QGraphicsItem *>();
	}
}

void* QGraphicsItem___transformations_atList(void* ptr, int i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsTransform * tmp = static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsTransform *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsTransform *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsTransform * tmp = static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsTransform *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsTransform *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else {
		return ({QGraphicsTransform * tmp = static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsTransform *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsTransform *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	}
}

void QGraphicsItem___transformations_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
	} else {
		static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
	}
}

void* QGraphicsItem___transformations_newList(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsTransform *>();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsTransform *>();
	} else {
		return new QList<QGraphicsTransform *>();
	}
}

class MyQGraphicsItemGroup: public QGraphicsItemGroup
{
public:
	MyQGraphicsItemGroup(QGraphicsItem *parent = Q_NULLPTR) : QGraphicsItemGroup(parent) {QGraphicsItemGroup_QGraphicsItemGroup_QRegisterMetaType();};
	int type() const { return callbackQGraphicsItem_Type(const_cast<void*>(static_cast<const void*>(this))); };
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQGraphicsItemGroup_Paint(this, painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	QRectF boundingRect() const { return *static_cast<QRectF*>(callbackQGraphicsItemGroup_BoundingRect(const_cast<void*>(static_cast<const void*>(this)))); };
	bool contains(const QPointF & point) const { return callbackQGraphicsItem_Contains(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPointF*>(&point)) != 0; };
};

Q_DECLARE_METATYPE(QGraphicsItemGroup*)
Q_DECLARE_METATYPE(MyQGraphicsItemGroup*)

int QGraphicsItemGroup_QGraphicsItemGroup_QRegisterMetaType(){qRegisterMetaType<QGraphicsItemGroup*>(); return qRegisterMetaType<MyQGraphicsItemGroup*>();}

void* QGraphicsItemGroup_NewQGraphicsItemGroup(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsItemGroup(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsItemGroup(static_cast<QGraphicsWidget*>(parent));
	} else {
		return new MyQGraphicsItemGroup(static_cast<QGraphicsItem*>(parent));
	}
}

void QGraphicsItemGroup_DestroyQGraphicsItemGroup(void* ptr)
{
	static_cast<QGraphicsItemGroup*>(ptr)->~QGraphicsItemGroup();
}

void QGraphicsItemGroup_Paint(void* ptr, void* painter, void* option, void* widget)
{
		static_cast<QGraphicsItemGroup*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsItemGroup_PaintDefault(void* ptr, void* painter, void* option, void* widget)
{
		static_cast<QGraphicsItemGroup*>(ptr)->QGraphicsItemGroup::paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void* QGraphicsItemGroup_BoundingRect(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QGraphicsItemGroup*>(ptr)->boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QGraphicsItemGroup_BoundingRectDefault(void* ptr)
{
		return ({ QRectF tmpValue = static_cast<QGraphicsItemGroup*>(ptr)->QGraphicsItemGroup::boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

class MyQGraphicsLayout: public QGraphicsLayout
{
public:
	MyQGraphicsLayout(QGraphicsLayoutItem *parent = Q_NULLPTR) : QGraphicsLayout(parent) {QGraphicsLayout_QGraphicsLayout_QRegisterMetaType();};
	void removeAt(int index) { callbackQGraphicsLayout_RemoveAt(this, index); };
	QGraphicsLayoutItem * itemAt(int i) const { return static_cast<QGraphicsLayoutItem*>(callbackQGraphicsLayout_ItemAt(const_cast<void*>(static_cast<const void*>(this)), i)); };
	int count() const { return callbackQGraphicsLayout_Count(const_cast<void*>(static_cast<const void*>(this))); };
	QSizeF sizeHint(Qt::SizeHint which, const QSizeF & constraint) const { return *static_cast<QSizeF*>(callbackQGraphicsLayout_SizeHint(const_cast<void*>(static_cast<const void*>(this)), which, const_cast<QSizeF*>(&constraint))); };
};

Q_DECLARE_METATYPE(QGraphicsLayout*)
Q_DECLARE_METATYPE(MyQGraphicsLayout*)

int QGraphicsLayout_QGraphicsLayout_QRegisterMetaType(){qRegisterMetaType<QGraphicsLayout*>(); return qRegisterMetaType<MyQGraphicsLayout*>();}

void* QGraphicsLayout_NewQGraphicsLayout(void* parent)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsLayout(static_cast<QGraphicsWidget*>(parent));
	} else {
		return new MyQGraphicsLayout(static_cast<QGraphicsLayoutItem*>(parent));
	}
}

void QGraphicsLayout_RemoveAt(void* ptr, int index)
{
	static_cast<QGraphicsLayout*>(ptr)->removeAt(index);
}

void QGraphicsLayout_SetContentsMargins(void* ptr, double left, double top, double right, double bottom)
{
	static_cast<QGraphicsLayout*>(ptr)->setContentsMargins(left, top, right, bottom);
}

void QGraphicsLayout_DestroyQGraphicsLayout(void* ptr)
{
	static_cast<QGraphicsLayout*>(ptr)->~QGraphicsLayout();
}

void* QGraphicsLayout_ItemAt(void* ptr, int i)
{
	return static_cast<QGraphicsLayout*>(ptr)->itemAt(i);
}

int QGraphicsLayout_Count(void* ptr)
{
	return static_cast<QGraphicsLayout*>(ptr)->count();
}

void* QGraphicsLayout_SizeHint(void* ptr, long long which, void* constraint)
{
	return ({ QSizeF tmpValue = static_cast<QGraphicsLayout*>(ptr)->sizeHint(static_cast<Qt::SizeHint>(which), *static_cast<QSizeF*>(constraint)); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

void* QGraphicsLayout_SizeHintDefault(void* ptr, long long which, void* constraint)
{
	Q_UNUSED(ptr);
	Q_UNUSED(which);
	Q_UNUSED(constraint);
	
}

class MyQGraphicsLayoutItem: public QGraphicsLayoutItem
{
public:
	MyQGraphicsLayoutItem(QGraphicsLayoutItem *parent = Q_NULLPTR, bool isLayout = false) : QGraphicsLayoutItem(parent, isLayout) {QGraphicsLayoutItem_QGraphicsLayoutItem_QRegisterMetaType();};
	 ~MyQGraphicsLayoutItem() { callbackQGraphicsLayoutItem_DestroyQGraphicsLayoutItem(this); };
	QSizeF sizeHint(Qt::SizeHint which, const QSizeF & constraint) const { return *static_cast<QSizeF*>(callbackQGraphicsLayoutItem_SizeHint(const_cast<void*>(static_cast<const void*>(this)), which, const_cast<QSizeF*>(&constraint))); };
};

Q_DECLARE_METATYPE(QGraphicsLayoutItem*)
Q_DECLARE_METATYPE(MyQGraphicsLayoutItem*)

int QGraphicsLayoutItem_QGraphicsLayoutItem_QRegisterMetaType(){qRegisterMetaType<QGraphicsLayoutItem*>(); return qRegisterMetaType<MyQGraphicsLayoutItem*>();}

void* QGraphicsLayoutItem_NewQGraphicsLayoutItem(void* parent, char isLayout)
{
	return new MyQGraphicsLayoutItem(static_cast<QGraphicsLayoutItem*>(parent), isLayout != 0);
}

void QGraphicsLayoutItem_SetMinimumSize(void* ptr, void* size)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setMinimumSize(*static_cast<QSizeF*>(size));
	} else {
		static_cast<QGraphicsLayoutItem*>(ptr)->setMinimumSize(*static_cast<QSizeF*>(size));
	}
}

void QGraphicsLayoutItem_SetMinimumSize2(void* ptr, double w, double h)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setMinimumSize(w, h);
	} else {
		static_cast<QGraphicsLayoutItem*>(ptr)->setMinimumSize(w, h);
	}
}

void QGraphicsLayoutItem_DestroyQGraphicsLayoutItem(void* ptr)
{
	static_cast<QGraphicsLayoutItem*>(ptr)->~QGraphicsLayoutItem();
}

void QGraphicsLayoutItem_DestroyQGraphicsLayoutItemDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QGraphicsLayoutItem_MinimumSize(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({ QSizeF tmpValue = static_cast<QGraphicsWidget*>(ptr)->minimumSize(); new QSizeF(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSizeF tmpValue = static_cast<QGraphicsLayoutItem*>(ptr)->minimumSize(); new QSizeF(tmpValue.width(), tmpValue.height()); });
	}
}

void* QGraphicsLayoutItem_SizeHint(void* ptr, long long which, void* constraint)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({ QSizeF tmpValue = static_cast<QGraphicsWidget*>(ptr)->sizeHint(static_cast<Qt::SizeHint>(which), *static_cast<QSizeF*>(constraint)); new QSizeF(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSizeF tmpValue = static_cast<QGraphicsLayoutItem*>(ptr)->sizeHint(static_cast<Qt::SizeHint>(which), *static_cast<QSizeF*>(constraint)); new QSizeF(tmpValue.width(), tmpValue.height()); });
	}
}

class MyQGraphicsObject: public QGraphicsObject
{
public:
	MyQGraphicsObject(QGraphicsItem *parent = Q_NULLPTR) : QGraphicsObject(parent) {QGraphicsObject_QGraphicsObject_QRegisterMetaType();};
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQGraphicsObject_ObjectNameChanged(this, objectNamePacked); };
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQGraphicsObject_Paint(this, painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	QRectF boundingRect() const { return *static_cast<QRectF*>(callbackQGraphicsObject_BoundingRect(const_cast<void*>(static_cast<const void*>(this)))); };
	bool contains(const QPointF & point) const { return callbackQGraphicsItem_Contains(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPointF*>(&point)) != 0; };
	int type() const { return callbackQGraphicsItem_Type(const_cast<void*>(static_cast<const void*>(this))); };
};

Q_DECLARE_METATYPE(QGraphicsObject*)
Q_DECLARE_METATYPE(MyQGraphicsObject*)

int QGraphicsObject_QGraphicsObject_QRegisterMetaType(){qRegisterMetaType<QGraphicsObject*>(); return qRegisterMetaType<MyQGraphicsObject*>();}

void* QGraphicsObject_NewQGraphicsObject(void* parent)
{
	return new MyQGraphicsObject(static_cast<QGraphicsItem*>(parent));
}

void QGraphicsObject_DestroyQGraphicsObject(void* ptr)
{
	static_cast<QGraphicsObject*>(ptr)->~QGraphicsObject();
}

void QGraphicsObject_SetEnabled(void* ptr, char enabled)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setEnabled(enabled != 0);
	} else {
		static_cast<QGraphicsObject*>(ptr)->setEnabled(enabled != 0);
	}
}

void* QGraphicsObject_Pos(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({ QPointF tmpValue = static_cast<QGraphicsWidget*>(ptr)->pos(); new QPointF(tmpValue.x(), tmpValue.y()); });
	} else {
		return ({ QPointF tmpValue = static_cast<QGraphicsObject*>(ptr)->pos(); new QPointF(tmpValue.x(), tmpValue.y()); });
	}
}

double QGraphicsObject_X(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->x();
	} else {
		return static_cast<QGraphicsObject*>(ptr)->x();
	}
}

double QGraphicsObject_Y(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->y();
	} else {
		return static_cast<QGraphicsObject*>(ptr)->y();
	}
}

void* QGraphicsObject___dynamicPropertyNames_atList(void* ptr, int i)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
	} else {
		return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
	}
}

void QGraphicsObject___dynamicPropertyNames_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
	} else {
		static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
	}
}

void* QGraphicsObject___dynamicPropertyNames_newList(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QList<QByteArray>();
	} else {
		return new QList<QByteArray>();
	}
}

void* QGraphicsObject___findChildren_atList2(void* ptr, int i)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else {
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
	}
}

void QGraphicsObject___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
		} else {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
		}
	} else {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
		} else {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
		}
	}
}

void* QGraphicsObject___findChildren_newList2(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QList<QObject*>();
	} else {
		return new QList<QObject*>();
	}
}

void* QGraphicsObject___findChildren_atList3(void* ptr, int i)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else {
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
	}
}

void QGraphicsObject___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
		} else {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
		}
	} else {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
		} else {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
		}
	}
}

void* QGraphicsObject___findChildren_newList3(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QList<QObject*>();
	} else {
		return new QList<QObject*>();
	}
}

void* QGraphicsObject___findChildren_atList(void* ptr, int i)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else {
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
	}
}

void QGraphicsObject___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
		} else {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
		}
	} else {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
		} else {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
		}
	}
}

void* QGraphicsObject___findChildren_newList(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QList<QObject*>();
	} else {
		return new QList<QObject*>();
	}
}

void* QGraphicsObject___children_atList(void* ptr, int i)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else {
		return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	}
}

void QGraphicsObject___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QLayout*>(i));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWidget*>(i));
		} else {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
		}
	} else {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QLayout*>(i));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWidget*>(i));
		} else {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
		}
	}
}

void* QGraphicsObject___children_newList(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QList<QObject *>();
	} else {
		return new QList<QObject *>();
	}
}

void QGraphicsObject_Paint(void* ptr, void* painter, void* option, void* widget)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
	} else {
	
	}
}

void QGraphicsObject_PaintDefault(void* ptr, void* painter, void* option, void* widget)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
	} else {
	
	}
}

void* QGraphicsObject_BoundingRect(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({ QRectF tmpValue = static_cast<QGraphicsWidget*>(ptr)->boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else {
	
	}
}

void* QGraphicsObject_BoundingRectDefault(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({ QRectF tmpValue = static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else {
	
	}
}

class MyQGraphicsTransform: public QGraphicsTransform
{
public:
	MyQGraphicsTransform(QObject *parent = Q_NULLPTR) : QGraphicsTransform(parent) {QGraphicsTransform_QGraphicsTransform_QRegisterMetaType();};
	void applyTo(QMatrix4x4 * matrix) const { callbackQGraphicsTransform_ApplyTo(const_cast<void*>(static_cast<const void*>(this)), matrix); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQGraphicsTransform_ObjectNameChanged(this, objectNamePacked); };
};

Q_DECLARE_METATYPE(QGraphicsTransform*)
Q_DECLARE_METATYPE(MyQGraphicsTransform*)

int QGraphicsTransform_QGraphicsTransform_QRegisterMetaType(){qRegisterMetaType<QGraphicsTransform*>(); return qRegisterMetaType<MyQGraphicsTransform*>();}

void* QGraphicsTransform_NewQGraphicsTransform(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsTransform(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsTransform(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsTransform(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsTransform(static_cast<QWidget*>(parent));
	} else {
		return new MyQGraphicsTransform(static_cast<QObject*>(parent));
	}
}

void QGraphicsTransform_DestroyQGraphicsTransform(void* ptr)
{
	static_cast<QGraphicsTransform*>(ptr)->~QGraphicsTransform();
}

void QGraphicsTransform_ApplyTo(void* ptr, void* matrix)
{
	static_cast<QGraphicsTransform*>(ptr)->applyTo(static_cast<QMatrix4x4*>(matrix));
}

void* QGraphicsTransform___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGraphicsTransform___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGraphicsTransform___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QGraphicsTransform___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGraphicsTransform___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QGraphicsTransform___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGraphicsTransform___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGraphicsTransform___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QGraphicsTransform___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGraphicsTransform___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGraphicsTransform___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QGraphicsTransform___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGraphicsTransform___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGraphicsTransform___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QGraphicsTransform___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

class MyQGraphicsWidget: public QGraphicsWidget
{
public:
	MyQGraphicsWidget(QGraphicsItem *parent = Q_NULLPTR, Qt::WindowFlags wFlags = Qt::WindowFlags()) : QGraphicsWidget(parent, wFlags) {QGraphicsWidget_QGraphicsWidget_QRegisterMetaType();};
	bool close() { return callbackQGraphicsWidget_Close(this) != 0; };
	int type() const { return callbackQGraphicsItem_Type(const_cast<void*>(static_cast<const void*>(this))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQGraphicsObject_ObjectNameChanged(this, objectNamePacked); };
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQGraphicsObject_Paint(this, painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	QRectF boundingRect() const { return *static_cast<QRectF*>(callbackQGraphicsObject_BoundingRect(const_cast<void*>(static_cast<const void*>(this)))); };
	bool contains(const QPointF & point) const { return callbackQGraphicsItem_Contains(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPointF*>(&point)) != 0; };
	QSizeF sizeHint(Qt::SizeHint which, const QSizeF & constraint) const { return *static_cast<QSizeF*>(callbackQGraphicsWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)), which, const_cast<QSizeF*>(&constraint))); };
};

Q_DECLARE_METATYPE(QGraphicsWidget*)
Q_DECLARE_METATYPE(MyQGraphicsWidget*)

int QGraphicsWidget_QGraphicsWidget_QRegisterMetaType(){qRegisterMetaType<QGraphicsWidget*>(); return qRegisterMetaType<MyQGraphicsWidget*>();}

void* QGraphicsWidget_NewQGraphicsWidget(void* parent, long long wFlags)
{
	return new MyQGraphicsWidget(static_cast<QGraphicsItem*>(parent), static_cast<Qt::WindowType>(wFlags));
}

char QGraphicsWidget_Close(void* ptr)
{
		bool returnArg;
	QMetaObject::invokeMethod(static_cast<QGraphicsWidget*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QGraphicsWidget_CloseDefault(void* ptr)
{
		return static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::close();
}

void QGraphicsWidget_SetContentsMargins(void* ptr, double left, double top, double right, double bottom)
{
		static_cast<QGraphicsWidget*>(ptr)->setContentsMargins(left, top, right, bottom);
}

void QGraphicsWidget_SetWindowTitle(void* ptr, struct QtWidgets_PackedString title)
{
		static_cast<QGraphicsWidget*>(ptr)->setWindowTitle(QString::fromUtf8(title.data, title.len));
}

void QGraphicsWidget_DestroyQGraphicsWidget(void* ptr)
{
	static_cast<QGraphicsWidget*>(ptr)->~QGraphicsWidget();
}

void* QGraphicsWidget_Layout(void* ptr)
{
		return static_cast<QGraphicsWidget*>(ptr)->layout();
}

void* QGraphicsWidget_Size(void* ptr)
{
		return ({ QSizeF tmpValue = static_cast<QGraphicsWidget*>(ptr)->size(); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

struct QtWidgets_PackedString QGraphicsWidget_WindowTitle(void* ptr)
{
		return ({ QByteArray* t8cb41d = new QByteArray(static_cast<QGraphicsWidget*>(ptr)->windowTitle().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t8cb41d->prepend("WHITESPACE").constData()+10), t8cb41d->size()-10, t8cb41d }; });
}

void* QGraphicsWidget_MinimumSize(void* ptr)
{
		return ({ QSizeF tmpValue = static_cast<QGraphicsWidget*>(ptr)->minimumSize(); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

void QGraphicsWidget_SetMinimumSize(void* ptr, void* minimumSize)
{
		static_cast<QGraphicsWidget*>(ptr)->setMinimumSize(*static_cast<QSizeF*>(minimumSize));
}

void* QGraphicsWidget___addActions_actions_atList(void* ptr, int i)
{
		return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGraphicsWidget___addActions_actions_setList(void* ptr, void* i)
{
		static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QGraphicsWidget___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QAction *>();
}

void* QGraphicsWidget___insertActions_actions_atList(void* ptr, int i)
{
		return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGraphicsWidget___insertActions_actions_setList(void* ptr, void* i)
{
		static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QGraphicsWidget___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QAction *>();
}

void* QGraphicsWidget___actions_atList(void* ptr, int i)
{
		return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGraphicsWidget___actions_setList(void* ptr, void* i)
{
		static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QGraphicsWidget___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QAction *>();
}

void* QGraphicsWidget_SizeHint(void* ptr, long long which, void* constraint)
{
		return ({ QSizeF tmpValue = static_cast<QGraphicsWidget*>(ptr)->sizeHint(static_cast<Qt::SizeHint>(which), *static_cast<QSizeF*>(constraint)); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

void* QGraphicsWidget_SizeHintDefault(void* ptr, long long which, void* constraint)
{
		return ({ QSizeF tmpValue = static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::sizeHint(static_cast<Qt::SizeHint>(which), *static_cast<QSizeF*>(constraint)); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

class MyQGroupBox: public QGroupBox
{
public:
	MyQGroupBox(QWidget *parent = Q_NULLPTR) : QGroupBox(parent) {QGroupBox_QGroupBox_QRegisterMetaType();};
	MyQGroupBox(const QString &title, QWidget *parent = Q_NULLPTR) : QGroupBox(title, parent) {QGroupBox_QGroupBox_QRegisterMetaType();};
	void Signal_Clicked(bool checked) { callbackQGroupBox_Clicked(this, checked); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void lower() { callbackQWidget_Lower(this); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QGroupBox*)
Q_DECLARE_METATYPE(MyQGroupBox*)

int QGroupBox_QGroupBox_QRegisterMetaType(){qRegisterMetaType<QGroupBox*>(); return qRegisterMetaType<MyQGroupBox*>();}

void* QGroupBox_NewQGroupBox(void* parent)
{
		return new MyQGroupBox(static_cast<QWidget*>(parent));
}

void* QGroupBox_NewQGroupBox2(struct QtWidgets_PackedString title, void* parent)
{
		return new MyQGroupBox(QString::fromUtf8(title.data, title.len), static_cast<QWidget*>(parent));
}

void QGroupBox_ConnectClicked(void* ptr, long long t)
{
	QObject::connect(static_cast<QGroupBox*>(ptr), static_cast<void (QGroupBox::*)(bool)>(&QGroupBox::clicked), static_cast<MyQGroupBox*>(ptr), static_cast<void (MyQGroupBox::*)(bool)>(&MyQGroupBox::Signal_Clicked), static_cast<Qt::ConnectionType>(t));
}

void QGroupBox_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<QGroupBox*>(ptr), static_cast<void (QGroupBox::*)(bool)>(&QGroupBox::clicked), static_cast<MyQGroupBox*>(ptr), static_cast<void (MyQGroupBox::*)(bool)>(&MyQGroupBox::Signal_Clicked));
}

void QGroupBox_Clicked(void* ptr, char checked)
{
	static_cast<QGroupBox*>(ptr)->clicked(checked != 0);
}

void QGroupBox_SetTitle(void* ptr, struct QtWidgets_PackedString title)
{
	static_cast<QGroupBox*>(ptr)->setTitle(QString::fromUtf8(title.data, title.len));
}

void QGroupBox_DestroyQGroupBox(void* ptr)
{
	static_cast<QGroupBox*>(ptr)->~QGroupBox();
}

struct QtWidgets_PackedString QGroupBox_Title(void* ptr)
{
	return ({ QByteArray* t937a1b = new QByteArray(static_cast<QGroupBox*>(ptr)->title().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t937a1b->prepend("WHITESPACE").constData()+10), t937a1b->size()-10, t937a1b }; });
}

class MyQHBoxLayout: public QHBoxLayout
{
public:
	MyQHBoxLayout() : QHBoxLayout() {QHBoxLayout_QHBoxLayout_QRegisterMetaType();};
	MyQHBoxLayout(QWidget *parent) : QHBoxLayout(parent) {QHBoxLayout_QHBoxLayout_QRegisterMetaType();};
	QSize minimumSize() const { return *static_cast<QSize*>(callbackQLayout_MinimumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	QLayout * layout() { return static_cast<QLayout*>(callbackQLayoutItem_Layout(this)); };
	QLayoutItem * takeAt(int index) { return static_cast<QLayoutItem*>(callbackQBoxLayout_TakeAt(this, index)); };
	void addItem(QLayoutItem * item) { callbackQBoxLayout_AddItem(this, item); };
	QLayoutItem * itemAt(int index) const { return static_cast<QLayoutItem*>(callbackQBoxLayout_ItemAt(const_cast<void*>(static_cast<const void*>(this)), index)); };
	int count() const { return callbackQBoxLayout_Count(const_cast<void*>(static_cast<const void*>(this))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQLayout_ObjectNameChanged(this, objectNamePacked); };
	QWidget * widget() { return static_cast<QWidget*>(callbackQLayoutItem_Widget(this)); };
	void setGeometry(const QRect & r) { callbackQLayout_SetGeometry(this, const_cast<QRect*>(&r)); };
	QRect geometry() const { return *static_cast<QRect*>(callbackQLayout_Geometry(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize maximumSize() const { return *static_cast<QSize*>(callbackQLayout_MaximumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQLayout_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::Orientations expandingDirections() const { return static_cast<Qt::Orientation>(callbackQLayout_ExpandingDirections(const_cast<void*>(static_cast<const void*>(this)))); };
	bool isEmpty() const { return callbackQLayout_IsEmpty(const_cast<void*>(static_cast<const void*>(this))) != 0; };
};

Q_DECLARE_METATYPE(QHBoxLayout*)
Q_DECLARE_METATYPE(MyQHBoxLayout*)

int QHBoxLayout_QHBoxLayout_QRegisterMetaType(){qRegisterMetaType<QHBoxLayout*>(); return qRegisterMetaType<MyQHBoxLayout*>();}

void* QHBoxLayout_NewQHBoxLayout()
{
	return new MyQHBoxLayout();
}

void* QHBoxLayout_NewQHBoxLayout2(void* parent)
{
		return new MyQHBoxLayout(static_cast<QWidget*>(parent));
}

void QHBoxLayout_DestroyQHBoxLayout(void* ptr)
{
	static_cast<QHBoxLayout*>(ptr)->~QHBoxLayout();
}

class MyQLabel: public QLabel
{
public:
	MyQLabel(QWidget *parent = Q_NULLPTR, Qt::WindowFlags ff = Qt::WindowFlags()) : QLabel(parent, ff) {QLabel_QLabel_QRegisterMetaType();};
	MyQLabel(const QString &text, QWidget *parent = Q_NULLPTR, Qt::WindowFlags ff = Qt::WindowFlags()) : QLabel(text, parent, ff) {QLabel_QLabel_QRegisterMetaType();};
	void setText(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQLabel_SetText(this, vqsPacked); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void lower() { callbackQWidget_Lower(this); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QLabel*)
Q_DECLARE_METATYPE(MyQLabel*)

int QLabel_QLabel_QRegisterMetaType(){qRegisterMetaType<QLabel*>(); return qRegisterMetaType<MyQLabel*>();}

void* QLabel_NewQLabel(void* parent, long long ff)
{
		return new MyQLabel(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(ff));
}

void* QLabel_NewQLabel2(struct QtWidgets_PackedString text, void* parent, long long ff)
{
		return new MyQLabel(QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(ff));
}

void QLabel_SetIndent(void* ptr, int vin)
{
	static_cast<QLabel*>(ptr)->setIndent(vin);
}

void QLabel_SetText(void* ptr, struct QtWidgets_PackedString vqs)
{
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "setText", Q_ARG(const QString, QString::fromUtf8(vqs.data, vqs.len)));
}

void QLabel_SetTextDefault(void* ptr, struct QtWidgets_PackedString vqs)
{
		static_cast<QLabel*>(ptr)->QLabel::setText(QString::fromUtf8(vqs.data, vqs.len));
}

void QLabel_DestroyQLabel(void* ptr)
{
	static_cast<QLabel*>(ptr)->~QLabel();
}

struct QtWidgets_PackedString QLabel_Text(void* ptr)
{
	return ({ QByteArray* t0d779d = new QByteArray(static_cast<QLabel*>(ptr)->text().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t0d779d->prepend("WHITESPACE").constData()+10), t0d779d->size()-10, t0d779d }; });
}

long long QLabel_TextFormat(void* ptr)
{
	return static_cast<QLabel*>(ptr)->textFormat();
}

int QLabel_Indent(void* ptr)
{
	return static_cast<QLabel*>(ptr)->indent();
}

int QLabel_Margin(void* ptr)
{
	return static_cast<QLabel*>(ptr)->margin();
}

class MyQLayout: public QLayout
{
public:
	MyQLayout() : QLayout() {QLayout_QLayout_QRegisterMetaType();};
	MyQLayout(QWidget *parent) : QLayout(parent) {QLayout_QLayout_QRegisterMetaType();};
	QLayout * layout() { return static_cast<QLayout*>(callbackQLayoutItem_Layout(this)); };
	QLayoutItem * takeAt(int index) { return static_cast<QLayoutItem*>(callbackQLayout_TakeAt(this, index)); };
	void addItem(QLayoutItem * item) { callbackQLayout_AddItem(this, item); };
	QLayoutItem * itemAt(int index) const { return static_cast<QLayoutItem*>(callbackQLayout_ItemAt(const_cast<void*>(static_cast<const void*>(this)), index)); };
	QSize minimumSize() const { return *static_cast<QSize*>(callbackQLayout_MinimumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	int count() const { return callbackQLayout_Count(const_cast<void*>(static_cast<const void*>(this))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQLayout_ObjectNameChanged(this, objectNamePacked); };
	QWidget * widget() { return static_cast<QWidget*>(callbackQLayoutItem_Widget(this)); };
	void setGeometry(const QRect & r) { callbackQLayout_SetGeometry(this, const_cast<QRect*>(&r)); };
	QRect geometry() const { return *static_cast<QRect*>(callbackQLayout_Geometry(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize maximumSize() const { return *static_cast<QSize*>(callbackQLayout_MaximumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQLayout_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::Orientations expandingDirections() const { return static_cast<Qt::Orientation>(callbackQLayout_ExpandingDirections(const_cast<void*>(static_cast<const void*>(this)))); };
	bool isEmpty() const { return callbackQLayout_IsEmpty(const_cast<void*>(static_cast<const void*>(this))) != 0; };
};

Q_DECLARE_METATYPE(QLayout*)
Q_DECLARE_METATYPE(MyQLayout*)

int QLayout_QLayout_QRegisterMetaType(){qRegisterMetaType<QLayout*>(); return qRegisterMetaType<MyQLayout*>();}

void* QLayout_NewQLayout2()
{
	return new MyQLayout();
}

void* QLayout_NewQLayout(void* parent)
{
	return new MyQLayout(static_cast<QWidget*>(parent));
}

void* QLayout_TakeAt(void* ptr, int index)
{
		return static_cast<QLayout*>(ptr)->takeAt(index);
}

void QLayout_AddItem(void* ptr, void* item)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(item))) {
		static_cast<QLayout*>(ptr)->addItem(static_cast<QLayout*>(item));
	} else {
		static_cast<QLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item));
	}
}

void QLayout_AddWidget(void* ptr, void* w)
{
		static_cast<QLayout*>(ptr)->addWidget(static_cast<QWidget*>(w));
}

void QLayout_SetContentsMargins2(void* ptr, void* margins)
{
		static_cast<QLayout*>(ptr)->setContentsMargins(*static_cast<QMargins*>(margins));
}

void QLayout_SetContentsMargins(void* ptr, int left, int top, int right, int bottom)
{
		static_cast<QLayout*>(ptr)->setContentsMargins(left, top, right, bottom);
}

void QLayout_SetEnabled(void* ptr, char enable)
{
		static_cast<QLayout*>(ptr)->setEnabled(enable != 0);
}

void QLayout_SetSpacing(void* ptr, int vin)
{
		static_cast<QLayout*>(ptr)->setSpacing(vin);
}

void* QLayout_ItemAt(void* ptr, int index)
{
		return static_cast<QLayout*>(ptr)->itemAt(index);
}

void* QLayout_ContentsMargins(void* ptr)
{
		return ({ QMargins tmpValue = static_cast<QLayout*>(ptr)->contentsMargins(); new QMargins(tmpValue.left(), tmpValue.top(), tmpValue.right(), tmpValue.bottom()); });
}

void* QLayout_MinimumSize(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QLayout*>(ptr)->minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QLayout_MinimumSizeDefault(void* ptr)
{
	if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QBoxLayout*>(ptr)->QBoxLayout::minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSize tmpValue = static_cast<QLayout*>(ptr)->QLayout::minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	}
}

int QLayout_Count(void* ptr)
{
		return static_cast<QLayout*>(ptr)->count();
}

int QLayout_Spacing(void* ptr)
{
		return static_cast<QLayout*>(ptr)->spacing();
}

void* QLayout___dynamicPropertyNames_atList(void* ptr, int i)
{
		return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QLayout___dynamicPropertyNames_setList(void* ptr, void* i)
{
		static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QLayout___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QByteArray>();
}

void* QLayout___findChildren_atList2(void* ptr, int i)
{
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLayout___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QLayout___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>();
}

void* QLayout___findChildren_atList3(void* ptr, int i)
{
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLayout___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QLayout___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>();
}

void* QLayout___findChildren_atList(void* ptr, int i)
{
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLayout___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QLayout___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>();
}

void* QLayout___children_atList(void* ptr, int i)
{
		return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLayout___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QLayout___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject *>();
}

void QLayout_SetGeometry(void* ptr, void* r)
{
		static_cast<QLayout*>(ptr)->setGeometry(*static_cast<QRect*>(r));
}

void QLayout_SetGeometryDefault(void* ptr, void* r)
{
	if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::setGeometry(*static_cast<QRect*>(r));
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::setGeometry(*static_cast<QRect*>(r));
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxLayout*>(ptr)->QBoxLayout::setGeometry(*static_cast<QRect*>(r));
	} else {
		static_cast<QLayout*>(ptr)->QLayout::setGeometry(*static_cast<QRect*>(r));
	}
}

void* QLayout_Geometry(void* ptr)
{
		return ({ QRect tmpValue = static_cast<QLayout*>(ptr)->geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QLayout_GeometryDefault(void* ptr)
{
	if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QRect tmpValue = static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QRect tmpValue = static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QRect tmpValue = static_cast<QBoxLayout*>(ptr)->QBoxLayout::geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QRect tmpValue = static_cast<QLayout*>(ptr)->QLayout::geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	}
}

void* QLayout_MaximumSize(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QLayout*>(ptr)->maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QLayout_MaximumSizeDefault(void* ptr)
{
	if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QBoxLayout*>(ptr)->QBoxLayout::maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSize tmpValue = static_cast<QLayout*>(ptr)->QLayout::maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	}
}

void* QLayout_SizeHint(void* ptr)
{
	Q_UNUSED(ptr);
	
}

void* QLayout_SizeHintDefault(void* ptr)
{
	if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QBoxLayout*>(ptr)->QBoxLayout::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
	
	}
}

long long QLayout_ExpandingDirections(void* ptr)
{
		return static_cast<QLayout*>(ptr)->expandingDirections();
}

long long QLayout_ExpandingDirectionsDefault(void* ptr)
{
	if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::expandingDirections();
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::expandingDirections();
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::expandingDirections();
	} else {
		return static_cast<QLayout*>(ptr)->QLayout::expandingDirections();
	}
}

char QLayout_IsEmpty(void* ptr)
{
		return static_cast<QLayout*>(ptr)->isEmpty();
}

char QLayout_IsEmptyDefault(void* ptr)
{
	if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::isEmpty();
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::isEmpty();
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::isEmpty();
	} else {
		return static_cast<QLayout*>(ptr)->QLayout::isEmpty();
	}
}

class MyQLayoutItem: public QLayoutItem
{
public:
	MyQLayoutItem(Qt::Alignment alignment = Qt::Alignment()) : QLayoutItem(alignment) {QLayoutItem_QLayoutItem_QRegisterMetaType();};
	QLayout * layout() { return static_cast<QLayout*>(callbackQLayoutItem_Layout(this)); };
	QWidget * widget() { return static_cast<QWidget*>(callbackQLayoutItem_Widget(this)); };
	void setGeometry(const QRect & r) { callbackQLayoutItem_SetGeometry(this, const_cast<QRect*>(&r)); };
	 ~MyQLayoutItem() { callbackQLayoutItem_DestroyQLayoutItem(this); };
	QRect geometry() const { return *static_cast<QRect*>(callbackQLayoutItem_Geometry(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize maximumSize() const { return *static_cast<QSize*>(callbackQLayoutItem_MaximumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSize() const { return *static_cast<QSize*>(callbackQLayoutItem_MinimumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQLayoutItem_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::Orientations expandingDirections() const { return static_cast<Qt::Orientation>(callbackQLayoutItem_ExpandingDirections(const_cast<void*>(static_cast<const void*>(this)))); };
	bool isEmpty() const { return callbackQLayoutItem_IsEmpty(const_cast<void*>(static_cast<const void*>(this))) != 0; };
};

Q_DECLARE_METATYPE(QLayoutItem*)
Q_DECLARE_METATYPE(MyQLayoutItem*)

int QLayoutItem_QLayoutItem_QRegisterMetaType(){qRegisterMetaType<QLayoutItem*>(); return qRegisterMetaType<MyQLayoutItem*>();}

void* QLayoutItem_Layout(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->layout();
	} else {
		return static_cast<QLayoutItem*>(ptr)->layout();
	}
}

void* QLayoutItem_LayoutDefault(void* ptr)
{
	if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::layout();
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::layout();
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::layout();
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->QLayout::layout();
	} else {
		return static_cast<QLayoutItem*>(ptr)->QLayoutItem::layout();
	}
}

void* QLayoutItem_NewQLayoutItem(long long alignment)
{
	return new MyQLayoutItem(static_cast<Qt::AlignmentFlag>(alignment));
}

void* QLayoutItem_Widget(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->widget();
	} else {
		return static_cast<QLayoutItem*>(ptr)->widget();
	}
}

void* QLayoutItem_WidgetDefault(void* ptr)
{
	if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::widget();
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::widget();
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::widget();
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->QLayout::widget();
	} else {
		return static_cast<QLayoutItem*>(ptr)->QLayoutItem::widget();
	}
}

void QLayoutItem_SetGeometry(void* ptr, void* r)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QLayout*>(ptr)->setGeometry(*static_cast<QRect*>(r));
	} else {
		static_cast<QLayoutItem*>(ptr)->setGeometry(*static_cast<QRect*>(r));
	}
}

void QLayoutItem_DestroyQLayoutItem(void* ptr)
{
	static_cast<QLayoutItem*>(ptr)->~QLayoutItem();
}

void QLayoutItem_DestroyQLayoutItemDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QLayoutItem_Geometry(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QRect tmpValue = static_cast<QLayout*>(ptr)->geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QRect tmpValue = static_cast<QLayoutItem*>(ptr)->geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	}
}

void* QLayoutItem_MaximumSize(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QLayout*>(ptr)->maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSize tmpValue = static_cast<QLayoutItem*>(ptr)->maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	}
}

void* QLayoutItem_MinimumSize(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QLayout*>(ptr)->minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSize tmpValue = static_cast<QLayoutItem*>(ptr)->minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	}
}

void* QLayoutItem_SizeHint(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QLayout*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSize tmpValue = static_cast<QLayoutItem*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	}
}

long long QLayoutItem_ExpandingDirections(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->expandingDirections();
	} else {
		return static_cast<QLayoutItem*>(ptr)->expandingDirections();
	}
}

char QLayoutItem_IsEmpty(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->isEmpty();
	} else {
		return static_cast<QLayoutItem*>(ptr)->isEmpty();
	}
}

int QMessageBox_ButtonMask_Type()
{
	return QMessageBox::ButtonMask;
}

class MyQPushButton: public QPushButton
{
public:
	MyQPushButton(QWidget *parent = Q_NULLPTR) : QPushButton(parent) {QPushButton_QPushButton_QRegisterMetaType();};
	MyQPushButton(const QIcon &icon, const QString &text, QWidget *parent = Q_NULLPTR) : QPushButton(icon, text, parent) {QPushButton_QPushButton_QRegisterMetaType();};
	MyQPushButton(const QString &text, QWidget *parent = Q_NULLPTR) : QPushButton(text, parent) {QPushButton_QPushButton_QRegisterMetaType();};
	void click() { callbackQAbstractButton_Click(this); };
	void Signal_Clicked(bool checked) { callbackQAbstractButton_Clicked(this, checked); };
	void paintEvent(QPaintEvent * e) { callbackQPushButton_PaintEvent(this, e); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void lower() { callbackQWidget_Lower(this); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QPushButton*)
Q_DECLARE_METATYPE(MyQPushButton*)

int QPushButton_QPushButton_QRegisterMetaType(){qRegisterMetaType<QPushButton*>(); return qRegisterMetaType<MyQPushButton*>();}

void* QPushButton_NewQPushButton(void* parent)
{
		return new MyQPushButton(static_cast<QWidget*>(parent));
}

void* QPushButton_NewQPushButton3(void* icon, struct QtWidgets_PackedString text, void* parent)
{
		return new MyQPushButton(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(parent));
}

void* QPushButton_NewQPushButton2(struct QtWidgets_PackedString text, void* parent)
{
		return new MyQPushButton(QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(parent));
}

void QPushButton_DestroyQPushButton(void* ptr)
{
	static_cast<QPushButton*>(ptr)->~QPushButton();
}

void QPushButton_PaintEvent(void* ptr, void* e)
{
	static_cast<QPushButton*>(ptr)->paintEvent(static_cast<QPaintEvent*>(e));
}

void QPushButton_PaintEventDefault(void* ptr, void* e)
{
		static_cast<QPushButton*>(ptr)->QPushButton::paintEvent(static_cast<QPaintEvent*>(e));
}

int QStyle_PM_MdiSubWindowMinimizedWidth_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_MdiSubWindowMinimizedWidth;
	#else
		return 0;
	#endif
}

int QStyle_PM_HeaderMargin_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_HeaderMargin;
	#else
		return 0;
	#endif
}

int QStyle_PM_HeaderMarkSize_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_HeaderMarkSize;
	#else
		return 0;
	#endif
}

int QStyle_PM_HeaderGripMargin_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_HeaderGripMargin;
	#else
		return 0;
	#endif
}

int QStyle_PM_TabBarTabShiftHorizontal_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_TabBarTabShiftHorizontal;
	#else
		return 0;
	#endif
}

int QStyle_PM_TabBarTabShiftVertical_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_TabBarTabShiftVertical;
	#else
		return 0;
	#endif
}

int QStyle_PM_TabBarScrollButtonWidth_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_TabBarScrollButtonWidth;
	#else
		return 0;
	#endif
}

int QStyle_PM_ToolBarFrameWidth_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_ToolBarFrameWidth;
	#else
		return 0;
	#endif
}

int QStyle_PM_ToolBarHandleExtent_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_ToolBarHandleExtent;
	#else
		return 0;
	#endif
}

int QStyle_PM_ToolBarItemSpacing_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_ToolBarItemSpacing;
	#else
		return 0;
	#endif
}

int QStyle_PM_ToolBarItemMargin_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_ToolBarItemMargin;
	#else
		return 0;
	#endif
}

int QStyle_PM_ToolBarSeparatorExtent_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_ToolBarSeparatorExtent;
	#else
		return 0;
	#endif
}

int QStyle_PM_ToolBarExtensionExtent_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_ToolBarExtensionExtent;
	#else
		return 0;
	#endif
}

int QStyle_PM_SpinBoxSliderHeight_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_SpinBoxSliderHeight;
	#else
		return 0;
	#endif
}

int QStyle_PM_DefaultTopLevelMargin_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_DefaultTopLevelMargin;
	#else
		return 0;
	#endif
}

int QStyle_PM_DefaultChildMargin_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_DefaultChildMargin;
	#else
		return 0;
	#endif
}

int QStyle_PM_DefaultLayoutSpacing_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_DefaultLayoutSpacing;
	#else
		return 0;
	#endif
}

int QStyle_PM_ToolBarIconSize_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_ToolBarIconSize;
	#else
		return 0;
	#endif
}

int QStyle_PM_ListViewIconSize_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_ListViewIconSize;
	#else
		return 0;
	#endif
}

int QStyle_PM_IconViewIconSize_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_IconViewIconSize;
	#else
		return 0;
	#endif
}

int QStyle_PM_SmallIconSize_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_SmallIconSize;
	#else
		return 0;
	#endif
}

int QStyle_PM_LargeIconSize_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_LargeIconSize;
	#else
		return 0;
	#endif
}

int QStyle_PM_FocusFrameVMargin_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_FocusFrameVMargin;
	#else
		return 0;
	#endif
}

int QStyle_PM_FocusFrameHMargin_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_FocusFrameHMargin;
	#else
		return 0;
	#endif
}

int QStyle_PM_ToolTipLabelFrameWidth_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_ToolTipLabelFrameWidth;
	#else
		return 0;
	#endif
}

int QStyle_PM_CheckBoxLabelSpacing_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_CheckBoxLabelSpacing;
	#else
		return 0;
	#endif
}

int QStyle_PM_TabBarIconSize_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_TabBarIconSize;
	#else
		return 0;
	#endif
}

int QStyle_PM_SizeGripSize_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_SizeGripSize;
	#else
		return 0;
	#endif
}

int QStyle_PM_DockWidgetTitleMargin_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_DockWidgetTitleMargin;
	#else
		return 0;
	#endif
}

int QStyle_PM_MessageBoxIconSize_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_MessageBoxIconSize;
	#else
		return 0;
	#endif
}

int QStyle_PM_ButtonIconSize_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_ButtonIconSize;
	#else
		return 0;
	#endif
}

int QStyle_PM_DockWidgetTitleBarButtonMargin_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_DockWidgetTitleBarButtonMargin;
	#else
		return 0;
	#endif
}

int QStyle_PM_RadioButtonLabelSpacing_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_RadioButtonLabelSpacing;
	#else
		return 0;
	#endif
}

int QStyle_PM_LayoutLeftMargin_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_LayoutLeftMargin;
	#else
		return 0;
	#endif
}

int QStyle_PM_LayoutTopMargin_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_LayoutTopMargin;
	#else
		return 0;
	#endif
}

int QStyle_PM_LayoutRightMargin_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_LayoutRightMargin;
	#else
		return 0;
	#endif
}

int QStyle_PM_LayoutBottomMargin_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_LayoutBottomMargin;
	#else
		return 0;
	#endif
}

int QStyle_PM_LayoutHorizontalSpacing_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_LayoutHorizontalSpacing;
	#else
		return 0;
	#endif
}

int QStyle_PM_LayoutVerticalSpacing_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_LayoutVerticalSpacing;
	#else
		return 0;
	#endif
}

int QStyle_PM_TabBar_ScrollButtonOverlap_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_TabBar_ScrollButtonOverlap;
	#else
		return 0;
	#endif
}

int QStyle_PM_TextCursorWidth_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_TextCursorWidth;
	#else
		return 0;
	#endif
}

int QStyle_PM_TabCloseIndicatorWidth_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_TabCloseIndicatorWidth;
	#else
		return 0;
	#endif
}

int QStyle_PM_TabCloseIndicatorHeight_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_TabCloseIndicatorHeight;
	#else
		return 0;
	#endif
}

int QStyle_PM_ScrollView_ScrollBarSpacing_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_ScrollView_ScrollBarSpacing;
	#else
		return 0;
	#endif
}

int QStyle_PM_ScrollView_ScrollBarOverlap_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_ScrollView_ScrollBarOverlap;
	#else
		return 0;
	#endif
}

int QStyle_PM_SubMenuOverlap_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_SubMenuOverlap;
	#else
		return 0;
	#endif
}

int QStyle_PM_TreeViewIndentation_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_TreeViewIndentation;
	#else
		return 0;
	#endif
}

int QStyle_PM_HeaderDefaultSectionSizeHorizontal_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_HeaderDefaultSectionSizeHorizontal;
	#else
		return 0;
	#endif
}

int QStyle_PM_HeaderDefaultSectionSizeVertical_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_HeaderDefaultSectionSizeVertical;
	#else
		return 0;
	#endif
}

int QStyle_PM_TitleBarButtonIconSize_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_TitleBarButtonIconSize;
	#else
		return 0;
	#endif
}

int QStyle_PM_TitleBarButtonSize_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_TitleBarButtonSize;
	#else
		return 0;
	#endif
}

int QStyle_PE_FrameTabWidget_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_FrameTabWidget;
	#else
		return 0;
	#endif
}

int QStyle_PE_FrameWindow_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_FrameWindow;
	#else
		return 0;
	#endif
}

int QStyle_PE_FrameButtonBevel_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_FrameButtonBevel;
	#else
		return 0;
	#endif
}

int QStyle_PE_FrameButtonTool_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_FrameButtonTool;
	#else
		return 0;
	#endif
}

int QStyle_PE_FrameTabBarBase_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_FrameTabBarBase;
	#else
		return 0;
	#endif
}

int QStyle_PE_PanelButtonCommand_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_PanelButtonCommand;
	#else
		return 0;
	#endif
}

int QStyle_PE_PanelButtonBevel_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_PanelButtonBevel;
	#else
		return 0;
	#endif
}

int QStyle_PE_PanelButtonTool_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_PanelButtonTool;
	#else
		return 0;
	#endif
}

int QStyle_PE_PanelMenuBar_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_PanelMenuBar;
	#else
		return 0;
	#endif
}

int QStyle_PE_PanelToolBar_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_PanelToolBar;
	#else
		return 0;
	#endif
}

int QStyle_PE_PanelLineEdit_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_PanelLineEdit;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorArrowDown_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorArrowDown;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorArrowLeft_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorArrowLeft;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorArrowRight_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorArrowRight;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorArrowUp_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorArrowUp;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorBranch_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorBranch;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorButtonDropDown_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorButtonDropDown;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorViewItemCheck_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorViewItemCheck;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorCheckBox_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorCheckBox;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorDockWidgetResizeHandle_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorDockWidgetResizeHandle;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorHeaderArrow_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorHeaderArrow;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorMenuCheckMark_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorMenuCheckMark;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorProgressChunk_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorProgressChunk;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorRadioButton_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorRadioButton;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorSpinDown_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorSpinDown;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorSpinMinus_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorSpinMinus;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorSpinPlus_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorSpinPlus;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorSpinUp_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorSpinUp;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorToolBarHandle_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorToolBarHandle;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorToolBarSeparator_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorToolBarSeparator;
	#else
		return 0;
	#endif
}

int QStyle_PE_PanelTipLabel_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_PanelTipLabel;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorTabTear_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorTabTear;
	#else
		return 0;
	#endif
}

int QStyle_PE_PanelScrollAreaCorner_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_PanelScrollAreaCorner;
	#else
		return 0;
	#endif
}

int QStyle_PE_Widget_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_Widget;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorColumnViewArrow_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorColumnViewArrow;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorItemViewItemDrop_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorItemViewItemDrop;
	#else
		return 0;
	#endif
}

int QStyle_PE_PanelItemViewItem_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_PanelItemViewItem;
	#else
		return 0;
	#endif
}

int QStyle_PE_PanelItemViewRow_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_PanelItemViewRow;
	#else
		return 0;
	#endif
}

int QStyle_PE_PanelStatusBar_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_PanelStatusBar;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorTabClose_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorTabClose;
	#else
		return 0;
	#endif
}

int QStyle_PE_PanelMenu_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_PanelMenu;
	#else
		return 0;
	#endif
}

int QStyle_PE_IndicatorTabTearRight_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::PE_IndicatorTabTearRight;
	#else
		return 0;
	#endif
}

int QStyle_SH_BlinkCursorWhenTextSelected_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_BlinkCursorWhenTextSelected;
	#else
		return 0;
	#endif
}

int QStyle_SH_RichText_FullWidthSelection_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_RichText_FullWidthSelection;
	#else
		return 0;
	#endif
}

int QStyle_SH_Menu_Scrollable_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Menu_Scrollable;
	#else
		return 0;
	#endif
}

int QStyle_SH_GroupBox_TextLabelVerticalAlignment_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_GroupBox_TextLabelVerticalAlignment;
	#else
		return 0;
	#endif
}

int QStyle_SH_GroupBox_TextLabelColor_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_GroupBox_TextLabelColor;
	#else
		return 0;
	#endif
}

int QStyle_SH_Menu_SloppySubMenus_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Menu_SloppySubMenus;
	#else
		return 0;
	#endif
}

int QStyle_SH_Table_GridLineColor_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Table_GridLineColor;
	#else
		return 0;
	#endif
}

int QStyle_SH_LineEdit_PasswordCharacter_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_LineEdit_PasswordCharacter;
	#else
		return 0;
	#endif
}

int QStyle_SH_DialogButtons_DefaultButton_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_DialogButtons_DefaultButton;
	#else
		return 0;
	#endif
}

int QStyle_SH_ToolBox_SelectedPageTitleBold_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ToolBox_SelectedPageTitleBold;
	#else
		return 0;
	#endif
}

int QStyle_SH_TabBar_PreferNoArrows_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_TabBar_PreferNoArrows;
	#else
		return 0;
	#endif
}

int QStyle_SH_ScrollBar_LeftClickAbsolutePosition_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ScrollBar_LeftClickAbsolutePosition;
	#else
		return 0;
	#endif
}

int QStyle_SH_ListViewExpand_SelectMouseType_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ListViewExpand_SelectMouseType;
	#else
		return 0;
	#endif
}

int QStyle_SH_UnderlineShortcut_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_UnderlineShortcut;
	#else
		return 0;
	#endif
}

int QStyle_SH_SpinBox_AnimateButton_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_SpinBox_AnimateButton;
	#else
		return 0;
	#endif
}

int QStyle_SH_SpinBox_KeyPressAutoRepeatRate_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_SpinBox_KeyPressAutoRepeatRate;
	#else
		return 0;
	#endif
}

int QStyle_SH_SpinBox_ClickAutoRepeatRate_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_SpinBox_ClickAutoRepeatRate;
	#else
		return 0;
	#endif
}

int QStyle_SH_Menu_FillScreenWithScroll_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Menu_FillScreenWithScroll;
	#else
		return 0;
	#endif
}

int QStyle_SH_ToolTipLabel_Opacity_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ToolTipLabel_Opacity;
	#else
		return 0;
	#endif
}

int QStyle_SH_DrawMenuBarSeparator_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_DrawMenuBarSeparator;
	#else
		return 0;
	#endif
}

int QStyle_SH_TitleBar_ModifyNotification_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_TitleBar_ModifyNotification;
	#else
		return 0;
	#endif
}

int QStyle_SH_Button_FocusPolicy_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Button_FocusPolicy;
	#else
		return 0;
	#endif
}

int QStyle_SH_MessageBox_UseBorderForButtonSpacing_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_MessageBox_UseBorderForButtonSpacing;
	#else
		return 0;
	#endif
}

int QStyle_SH_TitleBar_AutoRaise_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_TitleBar_AutoRaise;
	#else
		return 0;
	#endif
}

int QStyle_SH_ToolButton_PopupDelay_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ToolButton_PopupDelay;
	#else
		return 0;
	#endif
}

int QStyle_SH_FocusFrame_Mask_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_FocusFrame_Mask;
	#else
		return 0;
	#endif
}

int QStyle_SH_RubberBand_Mask_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_RubberBand_Mask;
	#else
		return 0;
	#endif
}

int QStyle_SH_WindowFrame_Mask_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_WindowFrame_Mask;
	#else
		return 0;
	#endif
}

int QStyle_SH_SpinControls_DisableOnBounds_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_SpinControls_DisableOnBounds;
	#else
		return 0;
	#endif
}

int QStyle_SH_Dial_BackgroundRole_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Dial_BackgroundRole;
	#else
		return 0;
	#endif
}

int QStyle_SH_ComboBox_LayoutDirection_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ComboBox_LayoutDirection;
	#else
		return 0;
	#endif
}

int QStyle_SH_ItemView_EllipsisLocation_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ItemView_EllipsisLocation;
	#else
		return 0;
	#endif
}

int QStyle_SH_ItemView_ShowDecorationSelected_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ItemView_ShowDecorationSelected;
	#else
		return 0;
	#endif
}

int QStyle_SH_ItemView_ActivateItemOnSingleClick_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ItemView_ActivateItemOnSingleClick;
	#else
		return 0;
	#endif
}

int QStyle_SH_ScrollBar_ContextMenu_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ScrollBar_ContextMenu;
	#else
		return 0;
	#endif
}

int QStyle_SH_ScrollBar_RollBetweenButtons_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ScrollBar_RollBetweenButtons;
	#else
		return 0;
	#endif
}

int QStyle_SH_Slider_AbsoluteSetButtons_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Slider_AbsoluteSetButtons;
	#else
		return 0;
	#endif
}

int QStyle_SH_Slider_PageSetButtons_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Slider_PageSetButtons;
	#else
		return 0;
	#endif
}

int QStyle_SH_Menu_KeyboardSearch_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Menu_KeyboardSearch;
	#else
		return 0;
	#endif
}

int QStyle_SH_TabBar_ElideMode_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_TabBar_ElideMode;
	#else
		return 0;
	#endif
}

int QStyle_SH_DialogButtonLayout_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_DialogButtonLayout;
	#else
		return 0;
	#endif
}

int QStyle_SH_ComboBox_PopupFrameStyle_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ComboBox_PopupFrameStyle;
	#else
		return 0;
	#endif
}

int QStyle_SH_MessageBox_TextInteractionFlags_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_MessageBox_TextInteractionFlags;
	#else
		return 0;
	#endif
}

int QStyle_SH_DialogButtonBox_ButtonsHaveIcons_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_DialogButtonBox_ButtonsHaveIcons;
	#else
		return 0;
	#endif
}

int QStyle_SH_SpellCheckUnderlineStyle_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_SpellCheckUnderlineStyle;
	#else
		return 0;
	#endif
}

int QStyle_SH_MessageBox_CenterButtons_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_MessageBox_CenterButtons;
	#else
		return 0;
	#endif
}

int QStyle_SH_Menu_SelectionWrap_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Menu_SelectionWrap;
	#else
		return 0;
	#endif
}

int QStyle_SH_ItemView_MovementWithoutUpdatingSelection_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ItemView_MovementWithoutUpdatingSelection;
	#else
		return 0;
	#endif
}

int QStyle_SH_ToolTip_Mask_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ToolTip_Mask;
	#else
		return 0;
	#endif
}

int QStyle_SH_FocusFrame_AboveWidget_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_FocusFrame_AboveWidget;
	#else
		return 0;
	#endif
}

int QStyle_SH_TextControl_FocusIndicatorTextCharFormat_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_TextControl_FocusIndicatorTextCharFormat;
	#else
		return 0;
	#endif
}

int QStyle_SH_WizardStyle_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_WizardStyle;
	#else
		return 0;
	#endif
}

int QStyle_SH_ItemView_ArrowKeysNavigateIntoChildren_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ItemView_ArrowKeysNavigateIntoChildren;
	#else
		return 0;
	#endif
}

int QStyle_SH_Menu_Mask_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Menu_Mask;
	#else
		return 0;
	#endif
}

int QStyle_SH_Menu_FlashTriggeredItem_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Menu_FlashTriggeredItem;
	#else
		return 0;
	#endif
}

int QStyle_SH_Menu_FadeOutOnHide_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Menu_FadeOutOnHide;
	#else
		return 0;
	#endif
}

int QStyle_SH_SpinBox_ClickAutoRepeatThreshold_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_SpinBox_ClickAutoRepeatThreshold;
	#else
		return 0;
	#endif
}

int QStyle_SH_ItemView_PaintAlternatingRowColorsForEmptyArea_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ItemView_PaintAlternatingRowColorsForEmptyArea;
	#else
		return 0;
	#endif
}

int QStyle_SH_FormLayoutWrapPolicy_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_FormLayoutWrapPolicy;
	#else
		return 0;
	#endif
}

int QStyle_SH_TabWidget_DefaultTabPosition_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_TabWidget_DefaultTabPosition;
	#else
		return 0;
	#endif
}

int QStyle_SH_ToolBar_Movable_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ToolBar_Movable;
	#else
		return 0;
	#endif
}

int QStyle_SH_FormLayoutFieldGrowthPolicy_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_FormLayoutFieldGrowthPolicy;
	#else
		return 0;
	#endif
}

int QStyle_SH_FormLayoutFormAlignment_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_FormLayoutFormAlignment;
	#else
		return 0;
	#endif
}

int QStyle_SH_FormLayoutLabelAlignment_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_FormLayoutLabelAlignment;
	#else
		return 0;
	#endif
}

int QStyle_SH_ItemView_DrawDelegateFrame_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ItemView_DrawDelegateFrame;
	#else
		return 0;
	#endif
}

int QStyle_SH_TabBar_CloseButtonPosition_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_TabBar_CloseButtonPosition;
	#else
		return 0;
	#endif
}

int QStyle_SH_DockWidget_ButtonsHaveFrame_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_DockWidget_ButtonsHaveFrame;
	#else
		return 0;
	#endif
}

int QStyle_SH_ToolButtonStyle_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ToolButtonStyle;
	#else
		return 0;
	#endif
}

int QStyle_SH_RequestSoftwareInputPanel_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_RequestSoftwareInputPanel;
	#else
		return 0;
	#endif
}

int QStyle_SH_ScrollBar_Transient_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ScrollBar_Transient;
	#else
		return 0;
	#endif
}

int QStyle_SH_Menu_SupportsSections_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Menu_SupportsSections;
	#else
		return 0;
	#endif
}

int QStyle_SH_ToolTip_WakeUpDelay_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ToolTip_WakeUpDelay;
	#else
		return 0;
	#endif
}

int QStyle_SH_ToolTip_FallAsleepDelay_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ToolTip_FallAsleepDelay;
	#else
		return 0;
	#endif
}

int QStyle_SH_Widget_Animate_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Widget_Animate;
	#else
		return 0;
	#endif
}

int QStyle_SH_Splitter_OpaqueResize_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Splitter_OpaqueResize;
	#else
		return 0;
	#endif
}

int QStyle_SH_ComboBox_UseNativePopup_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ComboBox_UseNativePopup;
	#else
		return 0;
	#endif
}

int QStyle_SH_LineEdit_PasswordMaskDelay_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_LineEdit_PasswordMaskDelay;
	#else
		return 0;
	#endif
}

int QStyle_SH_TabBar_ChangeCurrentDelay_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_TabBar_ChangeCurrentDelay;
	#else
		return 0;
	#endif
}

int QStyle_SH_Menu_SubMenuUniDirection_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Menu_SubMenuUniDirection;
	#else
		return 0;
	#endif
}

int QStyle_SH_Menu_SubMenuUniDirectionFailCount_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Menu_SubMenuUniDirectionFailCount;
	#else
		return 0;
	#endif
}

int QStyle_SH_Menu_SubMenuSloppySelectOtherActions_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Menu_SubMenuSloppySelectOtherActions;
	#else
		return 0;
	#endif
}

int QStyle_SH_Menu_SubMenuSloppyCloseTimeout_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Menu_SubMenuSloppyCloseTimeout;
	#else
		return 0;
	#endif
}

int QStyle_SH_Menu_SubMenuResetWhenReenteringParent_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Menu_SubMenuResetWhenReenteringParent;
	#else
		return 0;
	#endif
}

int QStyle_SH_Menu_SubMenuDontStartSloppyOnLeave_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_Menu_SubMenuDontStartSloppyOnLeave;
	#else
		return 0;
	#endif
}

int QStyle_SH_ItemView_ScrollMode_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::SH_ItemView_ScrollMode;
	#else
		return 0;
	#endif
}

int QStyle_SE_TabBarTearIndicator_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_TabBarTearIndicator;
	#else
		return 0;
	#endif
}

int QStyle_SE_TreeViewDisclosureItem_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_TreeViewDisclosureItem;
	#else
		return 0;
	#endif
}

int QStyle_SE_LineEditContents_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_LineEditContents;
	#else
		return 0;
	#endif
}

int QStyle_SE_FrameContents_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_FrameContents;
	#else
		return 0;
	#endif
}

int QStyle_SE_DockWidgetCloseButton_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_DockWidgetCloseButton;
	#else
		return 0;
	#endif
}

int QStyle_SE_DockWidgetFloatButton_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_DockWidgetFloatButton;
	#else
		return 0;
	#endif
}

int QStyle_SE_DockWidgetTitleBarText_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_DockWidgetTitleBarText;
	#else
		return 0;
	#endif
}

int QStyle_SE_DockWidgetIcon_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_DockWidgetIcon;
	#else
		return 0;
	#endif
}

int QStyle_SE_CheckBoxLayoutItem_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_CheckBoxLayoutItem;
	#else
		return 0;
	#endif
}

int QStyle_SE_ComboBoxLayoutItem_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_ComboBoxLayoutItem;
	#else
		return 0;
	#endif
}

int QStyle_SE_DateTimeEditLayoutItem_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_DateTimeEditLayoutItem;
	#else
		return 0;
	#endif
}

int QStyle_SE_DialogButtonBoxLayoutItem_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_DialogButtonBoxLayoutItem;
	#else
		return 0;
	#endif
}

int QStyle_SE_LabelLayoutItem_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_LabelLayoutItem;
	#else
		return 0;
	#endif
}

int QStyle_SE_ProgressBarLayoutItem_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_ProgressBarLayoutItem;
	#else
		return 0;
	#endif
}

int QStyle_SE_PushButtonLayoutItem_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_PushButtonLayoutItem;
	#else
		return 0;
	#endif
}

int QStyle_SE_RadioButtonLayoutItem_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_RadioButtonLayoutItem;
	#else
		return 0;
	#endif
}

int QStyle_SE_SliderLayoutItem_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_SliderLayoutItem;
	#else
		return 0;
	#endif
}

int QStyle_SE_SpinBoxLayoutItem_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_SpinBoxLayoutItem;
	#else
		return 0;
	#endif
}

int QStyle_SE_ToolButtonLayoutItem_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_ToolButtonLayoutItem;
	#else
		return 0;
	#endif
}

int QStyle_SE_FrameLayoutItem_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_FrameLayoutItem;
	#else
		return 0;
	#endif
}

int QStyle_SE_GroupBoxLayoutItem_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_GroupBoxLayoutItem;
	#else
		return 0;
	#endif
}

int QStyle_SE_TabWidgetLayoutItem_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_TabWidgetLayoutItem;
	#else
		return 0;
	#endif
}

int QStyle_SE_ItemViewItemDecoration_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_ItemViewItemDecoration;
	#else
		return 0;
	#endif
}

int QStyle_SE_ItemViewItemText_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_ItemViewItemText;
	#else
		return 0;
	#endif
}

int QStyle_SE_ItemViewItemFocusRect_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_ItemViewItemFocusRect;
	#else
		return 0;
	#endif
}

int QStyle_SE_TabBarTabLeftButton_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_TabBarTabLeftButton;
	#else
		return 0;
	#endif
}

int QStyle_SE_TabBarTabRightButton_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_TabBarTabRightButton;
	#else
		return 0;
	#endif
}

int QStyle_SE_TabBarTabText_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_TabBarTabText;
	#else
		return 0;
	#endif
}

int QStyle_SE_ShapedFrameContents_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_ShapedFrameContents;
	#else
		return 0;
	#endif
}

int QStyle_SE_ToolBarHandle_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_ToolBarHandle;
	#else
		return 0;
	#endif
}

int QStyle_SE_TabBarScrollLeftButton_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_TabBarScrollLeftButton;
	#else
		return 0;
	#endif
}

int QStyle_SE_TabBarScrollRightButton_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_TabBarScrollRightButton;
	#else
		return 0;
	#endif
}

int QStyle_SE_TabBarTearIndicatorRight_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_TabBarTearIndicatorRight;
	#else
		return 0;
	#endif
}

int QStyleHintReturn_SH_Mask_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyleHintReturn::SH_Mask;
	#else
		return 0;
	#endif
}

int QStyleHintReturn_SH_Variant_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyleHintReturn::SH_Variant;
	#else
		return 0;
	#endif
}

Q_DECLARE_METATYPE(QStyleOption*)
int QStyleOption_SO_Slider_Type()
{
	return QStyleOption::SO_Slider;
}

int QStyleOption_SO_SpinBox_Type()
{
	return QStyleOption::SO_SpinBox;
}

int QStyleOption_SO_ToolButton_Type()
{
	return QStyleOption::SO_ToolButton;
}

int QStyleOption_SO_ComboBox_Type()
{
	return QStyleOption::SO_ComboBox;
}

int QStyleOption_SO_TitleBar_Type()
{
	return QStyleOption::SO_TitleBar;
}

int QStyleOption_SO_GroupBox_Type()
{
	return QStyleOption::SO_GroupBox;
}

int QStyleOption_SO_SizeGrip_Type()
{
	return QStyleOption::SO_SizeGrip;
}

void* QStyleOption_NewQStyleOption2(void* other)
{
	return new QStyleOption(*static_cast<QStyleOption*>(other));
}

void* QStyleOption_NewQStyleOption(int version, int ty)
{
	return new QStyleOption(version, ty);
}

void QStyleOption_InitFrom(void* ptr, void* widget)
{
		static_cast<QStyleOption*>(ptr)->initFrom(static_cast<QWidget*>(widget));
}

void QStyleOption_DestroyQStyleOption(void* ptr)
{
	static_cast<QStyleOption*>(ptr)->~QStyleOption();
}

int QStyleOption_Type(void* ptr)
{
	return static_cast<QStyleOption*>(ptr)->type;
}

Q_DECLARE_METATYPE(QStyleOptionGraphicsItem)
Q_DECLARE_METATYPE(QStyleOptionGraphicsItem*)
void* QStyleOptionGraphicsItem_NewQStyleOptionGraphicsItem()
{
	return new QStyleOptionGraphicsItem();
}

void* QStyleOptionGraphicsItem_NewQStyleOptionGraphicsItem2(void* other)
{
	return new QStyleOptionGraphicsItem(*static_cast<QStyleOptionGraphicsItem*>(other));
}

class MyQTextEdit: public QTextEdit
{
public:
	MyQTextEdit(QWidget *parent = Q_NULLPTR) : QTextEdit(parent) {QTextEdit_QTextEdit_QRegisterMetaType();};
	MyQTextEdit(const QString &text, QWidget *parent = Q_NULLPTR) : QTextEdit(text, parent) {QTextEdit_QTextEdit_QRegisterMetaType();};
	void copy() { callbackQTextEdit_Copy(this); };
	void setText(const QString & text) { QByteArray* t372ea0 = new QByteArray(text.toUtf8()); QtWidgets_PackedString textPacked = { const_cast<char*>(t372ea0->prepend("WHITESPACE").constData()+10), t372ea0->size()-10, t372ea0 };callbackQTextEdit_SetText(this, textPacked); };
	 ~MyQTextEdit() { callbackQTextEdit_DestroyQTextEdit(this); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void lower() { callbackQWidget_Lower(this); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QTextEdit*)
Q_DECLARE_METATYPE(MyQTextEdit*)

int QTextEdit_QTextEdit_QRegisterMetaType(){qRegisterMetaType<QTextEdit*>(); return qRegisterMetaType<MyQTextEdit*>();}

void* QTextEdit_NewQTextEdit(void* parent)
{
		return new MyQTextEdit(static_cast<QWidget*>(parent));
}

void* QTextEdit_NewQTextEdit2(struct QtWidgets_PackedString text, void* parent)
{
		return new MyQTextEdit(QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(parent));
}

void QTextEdit_Copy(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "copy");
}

void QTextEdit_CopyDefault(void* ptr)
{
		static_cast<QTextEdit*>(ptr)->QTextEdit::copy();
}

void QTextEdit_SetText(void* ptr, struct QtWidgets_PackedString text)
{
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setText", Q_ARG(const QString, QString::fromUtf8(text.data, text.len)));
}

void QTextEdit_SetTextDefault(void* ptr, struct QtWidgets_PackedString text)
{
		static_cast<QTextEdit*>(ptr)->QTextEdit::setText(QString::fromUtf8(text.data, text.len));
}

void QTextEdit_DestroyQTextEdit(void* ptr)
{
	static_cast<QTextEdit*>(ptr)->~QTextEdit();
}

void QTextEdit_DestroyQTextEditDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtWidgets_PackedString QTextEdit_ToPlainText(void* ptr)
{
	return ({ QByteArray* te62381 = new QByteArray(static_cast<QTextEdit*>(ptr)->toPlainText().toUtf8()); QtWidgets_PackedString { const_cast<char*>(te62381->prepend("WHITESPACE").constData()+10), te62381->size()-10, te62381 }; });
}

void QTextEdit_Print(void* ptr, void* printer)
{
#ifndef Q_OS_IOS
	static_cast<QTextEdit*>(ptr)->print(static_cast<QPagedPaintDevice*>(printer));
#endif
}

class MyQVBoxLayout: public QVBoxLayout
{
public:
	MyQVBoxLayout() : QVBoxLayout() {QVBoxLayout_QVBoxLayout_QRegisterMetaType();};
	MyQVBoxLayout(QWidget *parent) : QVBoxLayout(parent) {QVBoxLayout_QVBoxLayout_QRegisterMetaType();};
	QSize minimumSize() const { return *static_cast<QSize*>(callbackQLayout_MinimumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	QLayout * layout() { return static_cast<QLayout*>(callbackQLayoutItem_Layout(this)); };
	QLayoutItem * takeAt(int index) { return static_cast<QLayoutItem*>(callbackQBoxLayout_TakeAt(this, index)); };
	void addItem(QLayoutItem * item) { callbackQBoxLayout_AddItem(this, item); };
	QLayoutItem * itemAt(int index) const { return static_cast<QLayoutItem*>(callbackQBoxLayout_ItemAt(const_cast<void*>(static_cast<const void*>(this)), index)); };
	int count() const { return callbackQBoxLayout_Count(const_cast<void*>(static_cast<const void*>(this))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQLayout_ObjectNameChanged(this, objectNamePacked); };
	QWidget * widget() { return static_cast<QWidget*>(callbackQLayoutItem_Widget(this)); };
	void setGeometry(const QRect & r) { callbackQLayout_SetGeometry(this, const_cast<QRect*>(&r)); };
	QRect geometry() const { return *static_cast<QRect*>(callbackQLayout_Geometry(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize maximumSize() const { return *static_cast<QSize*>(callbackQLayout_MaximumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQLayout_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::Orientations expandingDirections() const { return static_cast<Qt::Orientation>(callbackQLayout_ExpandingDirections(const_cast<void*>(static_cast<const void*>(this)))); };
	bool isEmpty() const { return callbackQLayout_IsEmpty(const_cast<void*>(static_cast<const void*>(this))) != 0; };
};

Q_DECLARE_METATYPE(QVBoxLayout*)
Q_DECLARE_METATYPE(MyQVBoxLayout*)

int QVBoxLayout_QVBoxLayout_QRegisterMetaType(){qRegisterMetaType<QVBoxLayout*>(); return qRegisterMetaType<MyQVBoxLayout*>();}

void* QVBoxLayout_NewQVBoxLayout()
{
	return new MyQVBoxLayout();
}

void* QVBoxLayout_NewQVBoxLayout2(void* parent)
{
		return new MyQVBoxLayout(static_cast<QWidget*>(parent));
}

void QVBoxLayout_DestroyQVBoxLayout(void* ptr)
{
	static_cast<QVBoxLayout*>(ptr)->~QVBoxLayout();
}

class MyQWidget: public QWidget
{
public:
	MyQWidget(QWidget *parent = Q_NULLPTR, Qt::WindowFlags ff = Qt::WindowFlags()) : QWidget(parent, ff) {QWidget_QWidget_QRegisterMetaType();};
	bool close() { return callbackQWidget_Close(this) != 0; };
	void lower() { callbackQWidget_Lower(this); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QWidget*)
Q_DECLARE_METATYPE(MyQWidget*)

int QWidget_QWidget_QRegisterMetaType(){qRegisterMetaType<QWidget*>(); return qRegisterMetaType<MyQWidget*>();}

void* QWidget_NewQWidget(void* parent, long long ff)
{
	return new MyQWidget(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(ff));
}

char QWidget_Close(void* ptr)
{
		bool returnArg;
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QWidget_CloseDefault(void* ptr)
{
	if (dynamic_cast<QGroupBox*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGroupBox*>(ptr)->QGroupBox::close();
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLabel*>(ptr)->QLabel::close();
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTextEdit*>(ptr)->QTextEdit::close();
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::close();
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFrame*>(ptr)->QFrame::close();
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPushButton*>(ptr)->QPushButton::close();
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractButton*>(ptr)->QAbstractButton::close();
	} else {
		return static_cast<QWidget*>(ptr)->QWidget::close();
	}
}

void QWidget_Create(void* ptr, uintptr_t window, char initializeWindow, char destroyOldWindow)
{
		static_cast<QWidget*>(ptr)->create(window, initializeWindow != 0, destroyOldWindow != 0);
}

void QWidget_Lower(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "lower");
}

void QWidget_LowerDefault(void* ptr)
{
	if (dynamic_cast<QGroupBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QGroupBox*>(ptr)->QGroupBox::lower();
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::lower();
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::lower();
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::lower();
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::lower();
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::lower();
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::lower();
	} else {
		static_cast<QWidget*>(ptr)->QWidget::lower();
	}
}

void QWidget_SetContentsMargins2(void* ptr, void* margins)
{
		static_cast<QWidget*>(ptr)->setContentsMargins(*static_cast<QMargins*>(margins));
}

void QWidget_SetContentsMargins(void* ptr, int left, int top, int right, int bottom)
{
		static_cast<QWidget*>(ptr)->setContentsMargins(left, top, right, bottom);
}

void QWidget_SetEnabled(void* ptr, char vbo)
{
		QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QWidget_SetEnabledDefault(void* ptr, char vbo)
{
	if (dynamic_cast<QGroupBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QGroupBox*>(ptr)->QGroupBox::setEnabled(vbo != 0);
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::setEnabled(vbo != 0);
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::setEnabled(vbo != 0);
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::setEnabled(vbo != 0);
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::setEnabled(vbo != 0);
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::setEnabled(vbo != 0);
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::setEnabled(vbo != 0);
	} else {
		static_cast<QWidget*>(ptr)->QWidget::setEnabled(vbo != 0);
	}
}

void QWidget_SetMinimumSize(void* ptr, void* vqs)
{
		static_cast<QWidget*>(ptr)->setMinimumSize(*static_cast<QSize*>(vqs));
}

void QWidget_SetMinimumSize2(void* ptr, int minw, int minh)
{
		static_cast<QWidget*>(ptr)->setMinimumSize(minw, minh);
}

void QWidget_SetWindowTitle(void* ptr, struct QtWidgets_PackedString vqs)
{
		QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setWindowTitle", Q_ARG(const QString, QString::fromUtf8(vqs.data, vqs.len)));
}

void QWidget_SetWindowTitleDefault(void* ptr, struct QtWidgets_PackedString vqs)
{
	if (dynamic_cast<QGroupBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QGroupBox*>(ptr)->QGroupBox::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else {
		static_cast<QWidget*>(ptr)->QWidget::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	}
}

void QWidget_Show(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "show");
}

void QWidget_ShowDefault(void* ptr)
{
	if (dynamic_cast<QGroupBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QGroupBox*>(ptr)->QGroupBox::show();
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::show();
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::show();
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::show();
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::show();
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::show();
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::show();
	} else {
		static_cast<QWidget*>(ptr)->QWidget::show();
	}
}

void QWidget_DestroyQWidget(void* ptr)
{
	static_cast<QWidget*>(ptr)->~QWidget();
}

void* QWidget_Layout(void* ptr)
{
		return static_cast<QWidget*>(ptr)->layout();
}

void* QWidget_ContentsMargins(void* ptr)
{
		return ({ QMargins tmpValue = static_cast<QWidget*>(ptr)->contentsMargins(); new QMargins(tmpValue.left(), tmpValue.top(), tmpValue.right(), tmpValue.bottom()); });
}

void* QWidget_MapTo(void* ptr, void* parent, void* pos)
{
		return ({ QPoint tmpValue = static_cast<QWidget*>(ptr)->mapTo(static_cast<QWidget*>(parent), *static_cast<QPoint*>(pos)); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void* QWidget_Pos(void* ptr)
{
		return ({ QPoint tmpValue = static_cast<QWidget*>(ptr)->pos(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void* QWidget_MinimumSize(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWidget*>(ptr)->minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QWidget_Size(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWidget*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct QtWidgets_PackedString QWidget_WindowTitle(void* ptr)
{
		return ({ QByteArray* tf3cd6c = new QByteArray(static_cast<QWidget*>(ptr)->windowTitle().toUtf8()); QtWidgets_PackedString { const_cast<char*>(tf3cd6c->prepend("WHITESPACE").constData()+10), tf3cd6c->size()-10, tf3cd6c }; });
}

void* QWidget_Window(void* ptr)
{
		return static_cast<QWidget*>(ptr)->window();
}

int QWidget_X(void* ptr)
{
		return static_cast<QWidget*>(ptr)->x();
}

int QWidget_Y(void* ptr)
{
		return static_cast<QWidget*>(ptr)->y();
}

void* QWidget___addActions_actions_atList(void* ptr, int i)
{
		return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWidget___addActions_actions_setList(void* ptr, void* i)
{
		static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QWidget___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QAction *>();
}

void* QWidget___insertActions_actions_atList(void* ptr, int i)
{
		return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWidget___insertActions_actions_setList(void* ptr, void* i)
{
		static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QWidget___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QAction *>();
}

void* QWidget___actions_atList(void* ptr, int i)
{
		return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWidget___actions_setList(void* ptr, void* i)
{
		static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QWidget___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QAction *>();
}

void* QWidget___dynamicPropertyNames_atList(void* ptr, int i)
{
		return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWidget___dynamicPropertyNames_setList(void* ptr, void* i)
{
		static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWidget___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QByteArray>();
}

void* QWidget___findChildren_atList2(void* ptr, int i)
{
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWidget___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QWidget___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>();
}

void* QWidget___findChildren_atList3(void* ptr, int i)
{
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWidget___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QWidget___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>();
}

void* QWidget___findChildren_atList(void* ptr, int i)
{
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWidget___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QWidget___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>();
}

void* QWidget___children_atList(void* ptr, int i)
{
		return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWidget___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QWidget___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject *>();
}

void* QWidget_PaintEngine(void* ptr)
{
		return static_cast<QWidget*>(ptr)->paintEngine();
}

void* QWidget_PaintEngineDefault(void* ptr)
{
	if (dynamic_cast<QGroupBox*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGroupBox*>(ptr)->QGroupBox::paintEngine();
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLabel*>(ptr)->QLabel::paintEngine();
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTextEdit*>(ptr)->QTextEdit::paintEngine();
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::paintEngine();
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFrame*>(ptr)->QFrame::paintEngine();
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPushButton*>(ptr)->QPushButton::paintEngine();
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractButton*>(ptr)->QAbstractButton::paintEngine();
	} else {
		return static_cast<QWidget*>(ptr)->QWidget::paintEngine();
	}
}

