// +build minimal

package widgets

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "widgets-minimal.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtWidgets_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtWidgets_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return []byte(gs)
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QAbstractButton struct {
	QWidget
}

type QAbstractButton_ITF interface {
	QWidget_ITF
	QAbstractButton_PTR() *QAbstractButton
}

func (ptr *QAbstractButton) QAbstractButton_PTR() *QAbstractButton {
	return ptr
}

func (ptr *QAbstractButton) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractButton) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractButton(ptr QAbstractButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractButton_PTR().Pointer()
	}
	return nil
}

func NewQAbstractButtonFromPointer(ptr unsafe.Pointer) (n *QAbstractButton) {
	n = new(QAbstractButton)
	n.SetPointer(ptr)
	return
}
func NewQAbstractButton(parent QWidget_ITF) *QAbstractButton {
	tmpValue := NewQAbstractButtonFromPointer(C.QAbstractButton_NewQAbstractButton(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQAbstractButton_Click
func callbackQAbstractButton_Click(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "click"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractButtonFromPointer(ptr).ClickDefault()
	}
}

func (ptr *QAbstractButton) ConnectClick(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "click"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "click", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "click", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractButton) DisconnectClick() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "click")
	}
}

func (ptr *QAbstractButton) Click() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_Click(ptr.Pointer())
	}
}

func (ptr *QAbstractButton) ClickDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_ClickDefault(ptr.Pointer())
	}
}

//export callbackQAbstractButton_Clicked
func callbackQAbstractButton_Clicked(ptr unsafe.Pointer, checked C.char) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		(*(*func(bool))(signal))(int8(checked) != 0)
	}

}

func (ptr *QAbstractButton) ConnectClicked(f func(checked bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.QAbstractButton_ConnectClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			f := func(checked bool) {
				(*(*func(bool))(signal))(checked)
				f(checked)
			}
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractButton) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *QAbstractButton) Clicked(checked bool) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_Clicked(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(checked))))
	}
}

//export callbackQAbstractButton_PaintEvent
func callbackQAbstractButton_PaintEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*gui.QPaintEvent))(signal))(gui.NewQPaintEventFromPointer(e))
	}

}

func (ptr *QAbstractButton) ConnectPaintEvent(f func(e *gui.QPaintEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "paintEvent"); signal != nil {
			f := func(e *gui.QPaintEvent) {
				(*(*func(*gui.QPaintEvent))(signal))(e)
				f(e)
			}
			qt.ConnectSignal(ptr.Pointer(), "paintEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "paintEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractButton) DisconnectPaintEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "paintEvent")
	}
}

func (ptr *QAbstractButton) PaintEvent(e gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QAbstractButton) SetText(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QAbstractButton_SetText(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QAbstractButton) DestroyQAbstractButton() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractButton_DestroyQAbstractButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractButton) Group() *QButtonGroup {
	if ptr.Pointer() != nil {
		tmpValue := NewQButtonGroupFromPointer(C.QAbstractButton_Group(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractButton) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAbstractButton_Text(ptr.Pointer()))
	}
	return ""
}

//go:generate stringer -type=QAbstractItemDelegate__EndEditHint
//QAbstractItemDelegate::EndEditHint
type QAbstractItemDelegate__EndEditHint int64

const (
	QAbstractItemDelegate__NoHint           QAbstractItemDelegate__EndEditHint = QAbstractItemDelegate__EndEditHint(0)
	QAbstractItemDelegate__EditNextItem     QAbstractItemDelegate__EndEditHint = QAbstractItemDelegate__EndEditHint(1)
	QAbstractItemDelegate__EditPreviousItem QAbstractItemDelegate__EndEditHint = QAbstractItemDelegate__EndEditHint(2)
	QAbstractItemDelegate__SubmitModelCache QAbstractItemDelegate__EndEditHint = QAbstractItemDelegate__EndEditHint(3)
	QAbstractItemDelegate__RevertModelCache QAbstractItemDelegate__EndEditHint = QAbstractItemDelegate__EndEditHint(4)
)

//go:generate stringer -type=QAbstractItemView__CursorAction
//QAbstractItemView::CursorAction
type QAbstractItemView__CursorAction int64

const (
	QAbstractItemView__MoveUp       QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(0)
	QAbstractItemView__MoveDown     QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(1)
	QAbstractItemView__MoveLeft     QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(2)
	QAbstractItemView__MoveRight    QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(3)
	QAbstractItemView__MoveHome     QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(4)
	QAbstractItemView__MoveEnd      QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(5)
	QAbstractItemView__MovePageUp   QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(6)
	QAbstractItemView__MovePageDown QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(7)
	QAbstractItemView__MoveNext     QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(8)
	QAbstractItemView__MovePrevious QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(9)
)

//go:generate stringer -type=QAbstractItemView__DropIndicatorPosition
//QAbstractItemView::DropIndicatorPosition
type QAbstractItemView__DropIndicatorPosition int64

const (
	QAbstractItemView__OnItem     QAbstractItemView__DropIndicatorPosition = QAbstractItemView__DropIndicatorPosition(0)
	QAbstractItemView__AboveItem  QAbstractItemView__DropIndicatorPosition = QAbstractItemView__DropIndicatorPosition(1)
	QAbstractItemView__BelowItem  QAbstractItemView__DropIndicatorPosition = QAbstractItemView__DropIndicatorPosition(2)
	QAbstractItemView__OnViewport QAbstractItemView__DropIndicatorPosition = QAbstractItemView__DropIndicatorPosition(3)
)

//go:generate stringer -type=QAbstractItemView__SelectionMode
//QAbstractItemView::SelectionMode
type QAbstractItemView__SelectionMode int64

const (
	QAbstractItemView__NoSelection         QAbstractItemView__SelectionMode = QAbstractItemView__SelectionMode(0)
	QAbstractItemView__SingleSelection     QAbstractItemView__SelectionMode = QAbstractItemView__SelectionMode(1)
	QAbstractItemView__MultiSelection      QAbstractItemView__SelectionMode = QAbstractItemView__SelectionMode(2)
	QAbstractItemView__ExtendedSelection   QAbstractItemView__SelectionMode = QAbstractItemView__SelectionMode(3)
	QAbstractItemView__ContiguousSelection QAbstractItemView__SelectionMode = QAbstractItemView__SelectionMode(4)
)

//go:generate stringer -type=QAbstractItemView__DragDropMode
//QAbstractItemView::DragDropMode
type QAbstractItemView__DragDropMode int64

const (
	QAbstractItemView__NoDragDrop   QAbstractItemView__DragDropMode = QAbstractItemView__DragDropMode(0)
	QAbstractItemView__DragOnly     QAbstractItemView__DragDropMode = QAbstractItemView__DragDropMode(1)
	QAbstractItemView__DropOnly     QAbstractItemView__DragDropMode = QAbstractItemView__DragDropMode(2)
	QAbstractItemView__DragDrop     QAbstractItemView__DragDropMode = QAbstractItemView__DragDropMode(3)
	QAbstractItemView__InternalMove QAbstractItemView__DragDropMode = QAbstractItemView__DragDropMode(4)
)

//go:generate stringer -type=QAbstractItemView__EditTrigger
//QAbstractItemView::EditTrigger
type QAbstractItemView__EditTrigger int64

const (
	QAbstractItemView__NoEditTriggers  QAbstractItemView__EditTrigger = QAbstractItemView__EditTrigger(0)
	QAbstractItemView__CurrentChanged  QAbstractItemView__EditTrigger = QAbstractItemView__EditTrigger(1)
	QAbstractItemView__DoubleClicked   QAbstractItemView__EditTrigger = QAbstractItemView__EditTrigger(2)
	QAbstractItemView__SelectedClicked QAbstractItemView__EditTrigger = QAbstractItemView__EditTrigger(4)
	QAbstractItemView__EditKeyPressed  QAbstractItemView__EditTrigger = QAbstractItemView__EditTrigger(8)
	QAbstractItemView__AnyKeyPressed   QAbstractItemView__EditTrigger = QAbstractItemView__EditTrigger(16)
	QAbstractItemView__AllEditTriggers QAbstractItemView__EditTrigger = QAbstractItemView__EditTrigger(31)
)

//go:generate stringer -type=QAbstractItemView__ScrollHint
//QAbstractItemView::ScrollHint
type QAbstractItemView__ScrollHint int64

const (
	QAbstractItemView__EnsureVisible    QAbstractItemView__ScrollHint = QAbstractItemView__ScrollHint(0)
	QAbstractItemView__PositionAtTop    QAbstractItemView__ScrollHint = QAbstractItemView__ScrollHint(1)
	QAbstractItemView__PositionAtBottom QAbstractItemView__ScrollHint = QAbstractItemView__ScrollHint(2)
	QAbstractItemView__PositionAtCenter QAbstractItemView__ScrollHint = QAbstractItemView__ScrollHint(3)
)

//go:generate stringer -type=QAbstractItemView__ScrollMode
//QAbstractItemView::ScrollMode
type QAbstractItemView__ScrollMode int64

const (
	QAbstractItemView__ScrollPerItem  QAbstractItemView__ScrollMode = QAbstractItemView__ScrollMode(0)
	QAbstractItemView__ScrollPerPixel QAbstractItemView__ScrollMode = QAbstractItemView__ScrollMode(1)
)

//go:generate stringer -type=QAbstractItemView__SelectionBehavior
//QAbstractItemView::SelectionBehavior
type QAbstractItemView__SelectionBehavior int64

const (
	QAbstractItemView__SelectItems   QAbstractItemView__SelectionBehavior = QAbstractItemView__SelectionBehavior(0)
	QAbstractItemView__SelectRows    QAbstractItemView__SelectionBehavior = QAbstractItemView__SelectionBehavior(1)
	QAbstractItemView__SelectColumns QAbstractItemView__SelectionBehavior = QAbstractItemView__SelectionBehavior(2)
)

//go:generate stringer -type=QAbstractItemView__State
//QAbstractItemView::State
type QAbstractItemView__State int64

const (
	QAbstractItemView__NoState            QAbstractItemView__State = QAbstractItemView__State(0)
	QAbstractItemView__DraggingState      QAbstractItemView__State = QAbstractItemView__State(1)
	QAbstractItemView__DragSelectingState QAbstractItemView__State = QAbstractItemView__State(2)
	QAbstractItemView__EditingState       QAbstractItemView__State = QAbstractItemView__State(3)
	QAbstractItemView__ExpandingState     QAbstractItemView__State = QAbstractItemView__State(4)
	QAbstractItemView__CollapsingState    QAbstractItemView__State = QAbstractItemView__State(5)
	QAbstractItemView__AnimatingState     QAbstractItemView__State = QAbstractItemView__State(6)
)

type QAbstractScrollArea struct {
	QFrame
}

type QAbstractScrollArea_ITF interface {
	QFrame_ITF
	QAbstractScrollArea_PTR() *QAbstractScrollArea
}

func (ptr *QAbstractScrollArea) QAbstractScrollArea_PTR() *QAbstractScrollArea {
	return ptr
}

func (ptr *QAbstractScrollArea) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QFrame_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractScrollArea) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QFrame_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractScrollArea(ptr QAbstractScrollArea_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractScrollArea_PTR().Pointer()
	}
	return nil
}

func NewQAbstractScrollAreaFromPointer(ptr unsafe.Pointer) (n *QAbstractScrollArea) {
	n = new(QAbstractScrollArea)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QAbstractScrollArea__SizeAdjustPolicy
//QAbstractScrollArea::SizeAdjustPolicy
type QAbstractScrollArea__SizeAdjustPolicy int64

const (
	QAbstractScrollArea__AdjustIgnored               QAbstractScrollArea__SizeAdjustPolicy = QAbstractScrollArea__SizeAdjustPolicy(0)
	QAbstractScrollArea__AdjustToContentsOnFirstShow QAbstractScrollArea__SizeAdjustPolicy = QAbstractScrollArea__SizeAdjustPolicy(1)
	QAbstractScrollArea__AdjustToContents            QAbstractScrollArea__SizeAdjustPolicy = QAbstractScrollArea__SizeAdjustPolicy(2)
)

func NewQAbstractScrollArea(parent QWidget_ITF) *QAbstractScrollArea {
	tmpValue := NewQAbstractScrollAreaFromPointer(C.QAbstractScrollArea_NewQAbstractScrollArea(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QAbstractScrollArea) DestroyQAbstractScrollArea() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractScrollArea_DestroyQAbstractScrollArea(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractScrollArea) __scrollBarWidgets_atList(i int) *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QAbstractScrollArea___scrollBarWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractScrollArea) __scrollBarWidgets_setList(i QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea___scrollBarWidgets_setList(ptr.Pointer(), PointerFromQWidget(i))
	}
}

func (ptr *QAbstractScrollArea) __scrollBarWidgets_newList() unsafe.Pointer {
	return C.QAbstractScrollArea___scrollBarWidgets_newList(ptr.Pointer())
}

//go:generate stringer -type=QAbstractSlider__SliderAction
//QAbstractSlider::SliderAction
type QAbstractSlider__SliderAction int64

const (
	QAbstractSlider__SliderNoAction      QAbstractSlider__SliderAction = QAbstractSlider__SliderAction(0)
	QAbstractSlider__SliderSingleStepAdd QAbstractSlider__SliderAction = QAbstractSlider__SliderAction(1)
	QAbstractSlider__SliderSingleStepSub QAbstractSlider__SliderAction = QAbstractSlider__SliderAction(2)
	QAbstractSlider__SliderPageStepAdd   QAbstractSlider__SliderAction = QAbstractSlider__SliderAction(3)
	QAbstractSlider__SliderPageStepSub   QAbstractSlider__SliderAction = QAbstractSlider__SliderAction(4)
	QAbstractSlider__SliderToMinimum     QAbstractSlider__SliderAction = QAbstractSlider__SliderAction(5)
	QAbstractSlider__SliderToMaximum     QAbstractSlider__SliderAction = QAbstractSlider__SliderAction(6)
	QAbstractSlider__SliderMove          QAbstractSlider__SliderAction = QAbstractSlider__SliderAction(7)
)

//go:generate stringer -type=QAbstractSlider__SliderChange
//QAbstractSlider::SliderChange
type QAbstractSlider__SliderChange int64

const (
	QAbstractSlider__SliderRangeChange       QAbstractSlider__SliderChange = QAbstractSlider__SliderChange(0)
	QAbstractSlider__SliderOrientationChange QAbstractSlider__SliderChange = QAbstractSlider__SliderChange(1)
	QAbstractSlider__SliderStepsChange       QAbstractSlider__SliderChange = QAbstractSlider__SliderChange(2)
	QAbstractSlider__SliderValueChange       QAbstractSlider__SliderChange = QAbstractSlider__SliderChange(3)
)

//go:generate stringer -type=QAbstractSpinBox__ButtonSymbols
//QAbstractSpinBox::ButtonSymbols
type QAbstractSpinBox__ButtonSymbols int64

const (
	QAbstractSpinBox__UpDownArrows QAbstractSpinBox__ButtonSymbols = QAbstractSpinBox__ButtonSymbols(0)
	QAbstractSpinBox__PlusMinus    QAbstractSpinBox__ButtonSymbols = QAbstractSpinBox__ButtonSymbols(1)
	QAbstractSpinBox__NoButtons    QAbstractSpinBox__ButtonSymbols = QAbstractSpinBox__ButtonSymbols(2)
)

//go:generate stringer -type=QAbstractSpinBox__CorrectionMode
//QAbstractSpinBox::CorrectionMode
type QAbstractSpinBox__CorrectionMode int64

const (
	QAbstractSpinBox__CorrectToPreviousValue QAbstractSpinBox__CorrectionMode = QAbstractSpinBox__CorrectionMode(0)
	QAbstractSpinBox__CorrectToNearestValue  QAbstractSpinBox__CorrectionMode = QAbstractSpinBox__CorrectionMode(1)
)

//go:generate stringer -type=QAbstractSpinBox__StepEnabledFlag
//QAbstractSpinBox::StepEnabledFlag
type QAbstractSpinBox__StepEnabledFlag int64

const (
	QAbstractSpinBox__StepNone        QAbstractSpinBox__StepEnabledFlag = QAbstractSpinBox__StepEnabledFlag(0x00)
	QAbstractSpinBox__StepUpEnabled   QAbstractSpinBox__StepEnabledFlag = QAbstractSpinBox__StepEnabledFlag(0x01)
	QAbstractSpinBox__StepDownEnabled QAbstractSpinBox__StepEnabledFlag = QAbstractSpinBox__StepEnabledFlag(0x02)
)

type QAction struct {
	core.QObject
}

type QAction_ITF interface {
	core.QObject_ITF
	QAction_PTR() *QAction
}

func (ptr *QAction) QAction_PTR() *QAction {
	return ptr
}

func (ptr *QAction) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QAction) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQAction(ptr QAction_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAction_PTR().Pointer()
	}
	return nil
}

func NewQActionFromPointer(ptr unsafe.Pointer) (n *QAction) {
	n = new(QAction)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QAction__ActionEvent
//QAction::ActionEvent
type QAction__ActionEvent int64

const (
	QAction__Trigger QAction__ActionEvent = QAction__ActionEvent(0)
	QAction__Hover   QAction__ActionEvent = QAction__ActionEvent(1)
)

//go:generate stringer -type=QAction__MenuRole
//QAction::MenuRole
type QAction__MenuRole int64

const (
	QAction__NoRole                  QAction__MenuRole = QAction__MenuRole(0)
	QAction__TextHeuristicRole       QAction__MenuRole = QAction__MenuRole(1)
	QAction__ApplicationSpecificRole QAction__MenuRole = QAction__MenuRole(2)
	QAction__AboutQtRole             QAction__MenuRole = QAction__MenuRole(3)
	QAction__AboutRole               QAction__MenuRole = QAction__MenuRole(4)
	QAction__PreferencesRole         QAction__MenuRole = QAction__MenuRole(5)
	QAction__QuitRole                QAction__MenuRole = QAction__MenuRole(6)
)

//go:generate stringer -type=QAction__Priority
//QAction::Priority
type QAction__Priority int64

const (
	QAction__LowPriority    QAction__Priority = QAction__Priority(0)
	QAction__NormalPriority QAction__Priority = QAction__Priority(128)
	QAction__HighPriority   QAction__Priority = QAction__Priority(256)
)

func NewQAction(parent core.QObject_ITF) *QAction {
	tmpValue := NewQActionFromPointer(C.QAction_NewQAction(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQAction3(icon gui.QIcon_ITF, text string, parent core.QObject_ITF) *QAction {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewQActionFromPointer(C.QAction_NewQAction3(gui.PointerFromQIcon(icon), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQAction2(text string, parent core.QObject_ITF) *QAction {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewQActionFromPointer(C.QAction_NewQAction2(C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQAction_SetEnabled
func callbackQAction_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewQActionFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QAction) ConnectSetEnabled(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setEnabled"); signal != nil {
			f := func(vbo bool) {
				(*(*func(bool))(signal))(vbo)
				f(vbo)
			}
			qt.ConnectSignal(ptr.Pointer(), "setEnabled", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setEnabled", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAction) DisconnectSetEnabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setEnabled")
	}
}

func (ptr *QAction) SetEnabled(vbo bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QAction) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QAction) SetText(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QAction_SetText(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QAction) DestroyQAction() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAction_DestroyQAction(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAction) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAction_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAction) Data() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QAction_Data(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QAction) __setShortcuts_shortcuts_atList(i int) *gui.QKeySequence {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQKeySequenceFromPointer(C.QAction___setShortcuts_shortcuts_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*gui.QKeySequence).DestroyQKeySequence)
		return tmpValue
	}
	return nil
}

func (ptr *QAction) __setShortcuts_shortcuts_setList(i gui.QKeySequence_ITF) {
	if ptr.Pointer() != nil {
		C.QAction___setShortcuts_shortcuts_setList(ptr.Pointer(), gui.PointerFromQKeySequence(i))
	}
}

func (ptr *QAction) __setShortcuts_shortcuts_newList() unsafe.Pointer {
	return C.QAction___setShortcuts_shortcuts_newList(ptr.Pointer())
}

func (ptr *QAction) __associatedGraphicsWidgets_atList(i int) *QGraphicsWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQGraphicsWidgetFromPointer(C.QAction___associatedGraphicsWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAction) __associatedGraphicsWidgets_setList(i QGraphicsWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QAction___associatedGraphicsWidgets_setList(ptr.Pointer(), PointerFromQGraphicsWidget(i))
	}
}

func (ptr *QAction) __associatedGraphicsWidgets_newList() unsafe.Pointer {
	return C.QAction___associatedGraphicsWidgets_newList(ptr.Pointer())
}

func (ptr *QAction) __shortcuts_atList(i int) *gui.QKeySequence {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQKeySequenceFromPointer(C.QAction___shortcuts_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*gui.QKeySequence).DestroyQKeySequence)
		return tmpValue
	}
	return nil
}

func (ptr *QAction) __shortcuts_setList(i gui.QKeySequence_ITF) {
	if ptr.Pointer() != nil {
		C.QAction___shortcuts_setList(ptr.Pointer(), gui.PointerFromQKeySequence(i))
	}
}

func (ptr *QAction) __shortcuts_newList() unsafe.Pointer {
	return C.QAction___shortcuts_newList(ptr.Pointer())
}

func (ptr *QAction) __associatedWidgets_atList(i int) *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QAction___associatedWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAction) __associatedWidgets_setList(i QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QAction___associatedWidgets_setList(ptr.Pointer(), PointerFromQWidget(i))
	}
}

func (ptr *QAction) __associatedWidgets_newList() unsafe.Pointer {
	return C.QAction___associatedWidgets_newList(ptr.Pointer())
}

func (ptr *QAction) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QAction___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QAction) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QAction___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QAction) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QAction___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QAction) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAction___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAction) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAction___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAction) __findChildren_newList2() unsafe.Pointer {
	return C.QAction___findChildren_newList2(ptr.Pointer())
}

func (ptr *QAction) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAction___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAction) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAction___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAction) __findChildren_newList3() unsafe.Pointer {
	return C.QAction___findChildren_newList3(ptr.Pointer())
}

func (ptr *QAction) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAction___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAction) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAction___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAction) __findChildren_newList() unsafe.Pointer {
	return C.QAction___findChildren_newList(ptr.Pointer())
}

func (ptr *QAction) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAction___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAction) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAction___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAction) __children_newList() unsafe.Pointer {
	return C.QAction___children_newList(ptr.Pointer())
}

//export callbackQAction_ObjectNameChanged
func callbackQAction_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

type QApplication struct {
	gui.QGuiApplication
}

type QApplication_ITF interface {
	gui.QGuiApplication_ITF
	QApplication_PTR() *QApplication
}

func (ptr *QApplication) QApplication_PTR() *QApplication {
	return ptr
}

func (ptr *QApplication) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGuiApplication_PTR().Pointer()
	}
	return nil
}

func (ptr *QApplication) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGuiApplication_PTR().SetPointer(p)
	}
}

func PointerFromQApplication(ptr QApplication_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QApplication_PTR().Pointer()
	}
	return nil
}

func NewQApplicationFromPointer(ptr unsafe.Pointer) (n *QApplication) {
	n = new(QApplication)
	n.SetPointer(ptr)
	return
}
func NewQApplication(argc int, argv []string) *QApplication {
	argvC := C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	tmpValue := NewQApplicationFromPointer(C.QApplication_NewQApplication(C.int(int32(argc)), argvC))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QApplication_Exec() int {
	return int(int32(C.QApplication_QApplication_Exec()))
}

func (ptr *QApplication) Exec() int {
	return int(int32(C.QApplication_QApplication_Exec()))
}

//export callbackQApplication_DestroyQApplication
func callbackQApplication_DestroyQApplication(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QApplication"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQApplicationFromPointer(ptr).DestroyQApplicationDefault()
	}
}

func (ptr *QApplication) ConnectDestroyQApplication(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QApplication"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QApplication", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QApplication", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QApplication) DisconnectDestroyQApplication() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QApplication")
	}
}

func (ptr *QApplication) DestroyQApplication() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QApplication_DestroyQApplication(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QApplication) DestroyQApplicationDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QApplication_DestroyQApplicationDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QApplication) __allWidgets_atList(i int) *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QApplication___allWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QApplication) __allWidgets_setList(i QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication___allWidgets_setList(ptr.Pointer(), PointerFromQWidget(i))
	}
}

func (ptr *QApplication) __allWidgets_newList() unsafe.Pointer {
	return C.QApplication___allWidgets_newList(ptr.Pointer())
}

func (ptr *QApplication) __topLevelWidgets_atList(i int) *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QApplication___topLevelWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QApplication) __topLevelWidgets_setList(i QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication___topLevelWidgets_setList(ptr.Pointer(), PointerFromQWidget(i))
	}
}

func (ptr *QApplication) __topLevelWidgets_newList() unsafe.Pointer {
	return C.QApplication___topLevelWidgets_newList(ptr.Pointer())
}

func (ptr *QApplication) __screens_atList(i int) *gui.QScreen {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQScreenFromPointer(C.QApplication___screens_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QApplication) __screens_setList(i gui.QScreen_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication___screens_setList(ptr.Pointer(), gui.PointerFromQScreen(i))
	}
}

func (ptr *QApplication) __screens_newList() unsafe.Pointer {
	return C.QApplication___screens_newList(ptr.Pointer())
}

func (ptr *QApplication) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QApplication___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QApplication) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QApplication) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QApplication___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QApplication) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QApplication___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QApplication) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QApplication) __findChildren_newList2() unsafe.Pointer {
	return C.QApplication___findChildren_newList2(ptr.Pointer())
}

func (ptr *QApplication) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QApplication___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QApplication) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QApplication) __findChildren_newList3() unsafe.Pointer {
	return C.QApplication___findChildren_newList3(ptr.Pointer())
}

func (ptr *QApplication) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QApplication___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QApplication) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QApplication) __findChildren_newList() unsafe.Pointer {
	return C.QApplication___findChildren_newList(ptr.Pointer())
}

func (ptr *QApplication) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QApplication___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QApplication) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QApplication) __children_newList() unsafe.Pointer {
	return C.QApplication___children_newList(ptr.Pointer())
}

//export callbackQApplication_ObjectNameChanged
func callbackQApplication_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

type QBoxLayout struct {
	QLayout
}

type QBoxLayout_ITF interface {
	QLayout_ITF
	QBoxLayout_PTR() *QBoxLayout
}

func (ptr *QBoxLayout) QBoxLayout_PTR() *QBoxLayout {
	return ptr
}

func (ptr *QBoxLayout) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QLayout_PTR().Pointer()
	}
	return nil
}

func (ptr *QBoxLayout) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QLayout_PTR().SetPointer(p)
	}
}

func PointerFromQBoxLayout(ptr QBoxLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBoxLayout_PTR().Pointer()
	}
	return nil
}

func NewQBoxLayoutFromPointer(ptr unsafe.Pointer) (n *QBoxLayout) {
	n = new(QBoxLayout)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QBoxLayout__Direction
//QBoxLayout::Direction
type QBoxLayout__Direction int64

const (
	QBoxLayout__LeftToRight QBoxLayout__Direction = QBoxLayout__Direction(0)
	QBoxLayout__RightToLeft QBoxLayout__Direction = QBoxLayout__Direction(1)
	QBoxLayout__TopToBottom QBoxLayout__Direction = QBoxLayout__Direction(2)
	QBoxLayout__BottomToTop QBoxLayout__Direction = QBoxLayout__Direction(3)
	QBoxLayout__Down        QBoxLayout__Direction = QBoxLayout__Direction(QBoxLayout__TopToBottom)
	QBoxLayout__Up          QBoxLayout__Direction = QBoxLayout__Direction(QBoxLayout__BottomToTop)
)

func NewQBoxLayout(dir QBoxLayout__Direction, parent QWidget_ITF) *QBoxLayout {
	tmpValue := NewQBoxLayoutFromPointer(C.QBoxLayout_NewQBoxLayout(C.longlong(dir), PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QBoxLayout) AddWidget(widget QWidget_ITF, stretch int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(int32(stretch)), C.longlong(alignment))
	}
}

func (ptr *QBoxLayout) DestroyQBoxLayout() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBoxLayout_DestroyQBoxLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQBoxLayout_TakeAt
func callbackQBoxLayout_TakeAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "takeAt"); signal != nil {
		return PointerFromQLayoutItem((*(*func(int) *QLayoutItem)(signal))(int(int32(index))))
	}

	return PointerFromQLayoutItem(NewQBoxLayoutFromPointer(ptr).TakeAtDefault(int(int32(index))))
}

func (ptr *QBoxLayout) TakeAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QBoxLayout_TakeAt(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

func (ptr *QBoxLayout) TakeAtDefault(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QBoxLayout_TakeAtDefault(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

//export callbackQBoxLayout_AddItem
func callbackQBoxLayout_AddItem(ptr unsafe.Pointer, item unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addItem"); signal != nil {
		(*(*func(*QLayoutItem))(signal))(NewQLayoutItemFromPointer(item))
	} else {
		NewQBoxLayoutFromPointer(ptr).AddItemDefault(NewQLayoutItemFromPointer(item))
	}
}

func (ptr *QBoxLayout) AddItem(item QLayoutItem_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddItem(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QBoxLayout) AddItemDefault(item QLayoutItem_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddItemDefault(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

//export callbackQBoxLayout_ItemAt
func callbackQBoxLayout_ItemAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemAt"); signal != nil {
		return PointerFromQLayoutItem((*(*func(int) *QLayoutItem)(signal))(int(int32(index))))
	}

	return PointerFromQLayoutItem(NewQBoxLayoutFromPointer(ptr).ItemAtDefault(int(int32(index))))
}

func (ptr *QBoxLayout) ItemAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QBoxLayout_ItemAt(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

func (ptr *QBoxLayout) ItemAtDefault(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QBoxLayout_ItemAtDefault(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

//export callbackQBoxLayout_Count
func callbackQBoxLayout_Count(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "count"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQBoxLayoutFromPointer(ptr).CountDefault()))
}

func (ptr *QBoxLayout) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBoxLayout_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBoxLayout) CountDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBoxLayout_CountDefault(ptr.Pointer())))
	}
	return 0
}

type QButtonGroup struct {
	core.QObject
}

type QButtonGroup_ITF interface {
	core.QObject_ITF
	QButtonGroup_PTR() *QButtonGroup
}

func (ptr *QButtonGroup) QButtonGroup_PTR() *QButtonGroup {
	return ptr
}

func (ptr *QButtonGroup) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QButtonGroup) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQButtonGroup(ptr QButtonGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QButtonGroup_PTR().Pointer()
	}
	return nil
}

func NewQButtonGroupFromPointer(ptr unsafe.Pointer) (n *QButtonGroup) {
	n = new(QButtonGroup)
	n.SetPointer(ptr)
	return
}
func NewQButtonGroup(parent core.QObject_ITF) *QButtonGroup {
	tmpValue := NewQButtonGroupFromPointer(C.QButtonGroup_NewQButtonGroup(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QButtonGroup) DestroyQButtonGroup() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QButtonGroup_DestroyQButtonGroup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QButtonGroup) Button(id int) *QAbstractButton {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractButtonFromPointer(C.QButtonGroup_Button(ptr.Pointer(), C.int(int32(id))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QButtonGroup) __buttons_atList(i int) *QAbstractButton {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractButtonFromPointer(C.QButtonGroup___buttons_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QButtonGroup) __buttons_setList(i QAbstractButton_ITF) {
	if ptr.Pointer() != nil {
		C.QButtonGroup___buttons_setList(ptr.Pointer(), PointerFromQAbstractButton(i))
	}
}

func (ptr *QButtonGroup) __buttons_newList() unsafe.Pointer {
	return C.QButtonGroup___buttons_newList(ptr.Pointer())
}

func (ptr *QButtonGroup) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QButtonGroup___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QButtonGroup) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QButtonGroup___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QButtonGroup) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QButtonGroup___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QButtonGroup) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QButtonGroup___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QButtonGroup) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QButtonGroup___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QButtonGroup) __findChildren_newList2() unsafe.Pointer {
	return C.QButtonGroup___findChildren_newList2(ptr.Pointer())
}

func (ptr *QButtonGroup) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QButtonGroup___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QButtonGroup) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QButtonGroup___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QButtonGroup) __findChildren_newList3() unsafe.Pointer {
	return C.QButtonGroup___findChildren_newList3(ptr.Pointer())
}

func (ptr *QButtonGroup) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QButtonGroup___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QButtonGroup) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QButtonGroup___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QButtonGroup) __findChildren_newList() unsafe.Pointer {
	return C.QButtonGroup___findChildren_newList(ptr.Pointer())
}

func (ptr *QButtonGroup) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QButtonGroup___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QButtonGroup) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QButtonGroup___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QButtonGroup) __children_newList() unsafe.Pointer {
	return C.QButtonGroup___children_newList(ptr.Pointer())
}

//export callbackQButtonGroup_ObjectNameChanged
func callbackQButtonGroup_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//go:generate stringer -type=QCalendarWidget__HorizontalHeaderFormat
//QCalendarWidget::HorizontalHeaderFormat
type QCalendarWidget__HorizontalHeaderFormat int64

const (
	QCalendarWidget__NoHorizontalHeader   QCalendarWidget__HorizontalHeaderFormat = QCalendarWidget__HorizontalHeaderFormat(0)
	QCalendarWidget__SingleLetterDayNames QCalendarWidget__HorizontalHeaderFormat = QCalendarWidget__HorizontalHeaderFormat(1)
	QCalendarWidget__ShortDayNames        QCalendarWidget__HorizontalHeaderFormat = QCalendarWidget__HorizontalHeaderFormat(2)
	QCalendarWidget__LongDayNames         QCalendarWidget__HorizontalHeaderFormat = QCalendarWidget__HorizontalHeaderFormat(3)
)

//go:generate stringer -type=QCalendarWidget__SelectionMode
//QCalendarWidget::SelectionMode
type QCalendarWidget__SelectionMode int64

const (
	QCalendarWidget__NoSelection     QCalendarWidget__SelectionMode = QCalendarWidget__SelectionMode(0)
	QCalendarWidget__SingleSelection QCalendarWidget__SelectionMode = QCalendarWidget__SelectionMode(1)
)

//go:generate stringer -type=QCalendarWidget__VerticalHeaderFormat
//QCalendarWidget::VerticalHeaderFormat
type QCalendarWidget__VerticalHeaderFormat int64

const (
	QCalendarWidget__NoVerticalHeader QCalendarWidget__VerticalHeaderFormat = QCalendarWidget__VerticalHeaderFormat(0)
	QCalendarWidget__ISOWeekNumbers   QCalendarWidget__VerticalHeaderFormat = QCalendarWidget__VerticalHeaderFormat(1)
)

//go:generate stringer -type=QColorDialog__ColorDialogOption
//QColorDialog::ColorDialogOption
type QColorDialog__ColorDialogOption int64

const (
	QColorDialog__ShowAlphaChannel    QColorDialog__ColorDialogOption = QColorDialog__ColorDialogOption(0x00000001)
	QColorDialog__NoButtons           QColorDialog__ColorDialogOption = QColorDialog__ColorDialogOption(0x00000002)
	QColorDialog__DontUseNativeDialog QColorDialog__ColorDialogOption = QColorDialog__ColorDialogOption(0x00000004)
)

//go:generate stringer -type=QColormap__Mode
//QColormap::Mode
type QColormap__Mode int64

const (
	QColormap__Direct  QColormap__Mode = QColormap__Mode(0)
	QColormap__Indexed QColormap__Mode = QColormap__Mode(1)
	QColormap__Gray    QColormap__Mode = QColormap__Mode(2)
)

//go:generate stringer -type=QComboBox__InsertPolicy
//QComboBox::InsertPolicy
type QComboBox__InsertPolicy int64

const (
	QComboBox__NoInsert             QComboBox__InsertPolicy = QComboBox__InsertPolicy(0)
	QComboBox__InsertAtTop          QComboBox__InsertPolicy = QComboBox__InsertPolicy(1)
	QComboBox__InsertAtCurrent      QComboBox__InsertPolicy = QComboBox__InsertPolicy(2)
	QComboBox__InsertAtBottom       QComboBox__InsertPolicy = QComboBox__InsertPolicy(3)
	QComboBox__InsertAfterCurrent   QComboBox__InsertPolicy = QComboBox__InsertPolicy(4)
	QComboBox__InsertBeforeCurrent  QComboBox__InsertPolicy = QComboBox__InsertPolicy(5)
	QComboBox__InsertAlphabetically QComboBox__InsertPolicy = QComboBox__InsertPolicy(6)
)

//go:generate stringer -type=QComboBox__SizeAdjustPolicy
//QComboBox::SizeAdjustPolicy
type QComboBox__SizeAdjustPolicy int64

const (
	QComboBox__AdjustToContents                      QComboBox__SizeAdjustPolicy = QComboBox__SizeAdjustPolicy(0)
	QComboBox__AdjustToContentsOnFirstShow           QComboBox__SizeAdjustPolicy = QComboBox__SizeAdjustPolicy(1)
	QComboBox__AdjustToMinimumContentsLength         QComboBox__SizeAdjustPolicy = QComboBox__SizeAdjustPolicy(2)
	QComboBox__AdjustToMinimumContentsLengthWithIcon QComboBox__SizeAdjustPolicy = QComboBox__SizeAdjustPolicy(3)
)

//go:generate stringer -type=QCompleter__CompletionMode
//QCompleter::CompletionMode
type QCompleter__CompletionMode int64

const (
	QCompleter__PopupCompletion           QCompleter__CompletionMode = QCompleter__CompletionMode(0)
	QCompleter__UnfilteredPopupCompletion QCompleter__CompletionMode = QCompleter__CompletionMode(1)
	QCompleter__InlineCompletion          QCompleter__CompletionMode = QCompleter__CompletionMode(2)
)

//go:generate stringer -type=QCompleter__ModelSorting
//QCompleter::ModelSorting
type QCompleter__ModelSorting int64

const (
	QCompleter__UnsortedModel                QCompleter__ModelSorting = QCompleter__ModelSorting(0)
	QCompleter__CaseSensitivelySortedModel   QCompleter__ModelSorting = QCompleter__ModelSorting(1)
	QCompleter__CaseInsensitivelySortedModel QCompleter__ModelSorting = QCompleter__ModelSorting(2)
)

//go:generate stringer -type=QDataWidgetMapper__SubmitPolicy
//QDataWidgetMapper::SubmitPolicy
type QDataWidgetMapper__SubmitPolicy int64

const (
	QDataWidgetMapper__AutoSubmit   QDataWidgetMapper__SubmitPolicy = QDataWidgetMapper__SubmitPolicy(0)
	QDataWidgetMapper__ManualSubmit QDataWidgetMapper__SubmitPolicy = QDataWidgetMapper__SubmitPolicy(1)
)

//go:generate stringer -type=QDateTimeEdit__Section
//QDateTimeEdit::Section
type QDateTimeEdit__Section int64

const (
	QDateTimeEdit__NoSection     QDateTimeEdit__Section = QDateTimeEdit__Section(0x0000)
	QDateTimeEdit__AmPmSection   QDateTimeEdit__Section = QDateTimeEdit__Section(0x0001)
	QDateTimeEdit__MSecSection   QDateTimeEdit__Section = QDateTimeEdit__Section(0x0002)
	QDateTimeEdit__SecondSection QDateTimeEdit__Section = QDateTimeEdit__Section(0x0004)
	QDateTimeEdit__MinuteSection QDateTimeEdit__Section = QDateTimeEdit__Section(0x0008)
	QDateTimeEdit__HourSection   QDateTimeEdit__Section = QDateTimeEdit__Section(0x0010)
	QDateTimeEdit__DaySection    QDateTimeEdit__Section = QDateTimeEdit__Section(0x0100)
	QDateTimeEdit__MonthSection  QDateTimeEdit__Section = QDateTimeEdit__Section(0x0200)
	QDateTimeEdit__YearSection   QDateTimeEdit__Section = QDateTimeEdit__Section(0x0400)
)

//go:generate stringer -type=QDialog__DialogCode
//QDialog::DialogCode
type QDialog__DialogCode int64

const (
	QDialog__Rejected QDialog__DialogCode = QDialog__DialogCode(0)
	QDialog__Accepted QDialog__DialogCode = QDialog__DialogCode(1)
)

//go:generate stringer -type=QDialogButtonBox__ButtonLayout
//QDialogButtonBox::ButtonLayout
type QDialogButtonBox__ButtonLayout int64

const (
	QDialogButtonBox__WinLayout   QDialogButtonBox__ButtonLayout = QDialogButtonBox__ButtonLayout(0)
	QDialogButtonBox__MacLayout   QDialogButtonBox__ButtonLayout = QDialogButtonBox__ButtonLayout(1)
	QDialogButtonBox__KdeLayout   QDialogButtonBox__ButtonLayout = QDialogButtonBox__ButtonLayout(2)
	QDialogButtonBox__GnomeLayout QDialogButtonBox__ButtonLayout = QDialogButtonBox__ButtonLayout(3)
)

//go:generate stringer -type=QDialogButtonBox__ButtonRole
//QDialogButtonBox::ButtonRole
type QDialogButtonBox__ButtonRole int64

const (
	QDialogButtonBox__InvalidRole     QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(-1)
	QDialogButtonBox__AcceptRole      QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(0)
	QDialogButtonBox__RejectRole      QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(1)
	QDialogButtonBox__DestructiveRole QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(2)
	QDialogButtonBox__ActionRole      QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(3)
	QDialogButtonBox__HelpRole        QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(4)
	QDialogButtonBox__YesRole         QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(5)
	QDialogButtonBox__NoRole          QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(6)
	QDialogButtonBox__ResetRole       QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(7)
	QDialogButtonBox__ApplyRole       QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(8)
	QDialogButtonBox__NRoles          QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(9)
)

//go:generate stringer -type=QDialogButtonBox__StandardButton
//QDialogButtonBox::StandardButton
type QDialogButtonBox__StandardButton int64

const (
	QDialogButtonBox__NoButton        QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00000000)
	QDialogButtonBox__Ok              QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00000400)
	QDialogButtonBox__Save            QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00000800)
	QDialogButtonBox__SaveAll         QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00001000)
	QDialogButtonBox__Open            QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00002000)
	QDialogButtonBox__Yes             QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00004000)
	QDialogButtonBox__YesToAll        QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00008000)
	QDialogButtonBox__No              QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00010000)
	QDialogButtonBox__NoToAll         QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00020000)
	QDialogButtonBox__Abort           QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00040000)
	QDialogButtonBox__Retry           QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00080000)
	QDialogButtonBox__Ignore          QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00100000)
	QDialogButtonBox__Close           QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00200000)
	QDialogButtonBox__Cancel          QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00400000)
	QDialogButtonBox__Discard         QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00800000)
	QDialogButtonBox__Help            QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x01000000)
	QDialogButtonBox__Apply           QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x02000000)
	QDialogButtonBox__Reset           QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x04000000)
	QDialogButtonBox__RestoreDefaults QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x08000000)
	QDialogButtonBox__FirstButton     QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(QDialogButtonBox__Ok)
	QDialogButtonBox__LastButton      QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(QDialogButtonBox__RestoreDefaults)
)

//go:generate stringer -type=QDirModel__Roles
//QDirModel::Roles
type QDirModel__Roles int64

var (
	QDirModel__FileIconRole QDirModel__Roles = QDirModel__Roles(core.Qt__DecorationRole)
	QDirModel__FilePathRole QDirModel__Roles = QDirModel__Roles(C.QDirModel_FilePathRole_Type())
	QDirModel__FileNameRole QDirModel__Roles = QDirModel__Roles(C.QDirModel_FileNameRole_Type())
)

//go:generate stringer -type=QDockWidget__DockWidgetFeature
//QDockWidget::DockWidgetFeature
type QDockWidget__DockWidgetFeature int64

const (
	QDockWidget__DockWidgetClosable         QDockWidget__DockWidgetFeature = QDockWidget__DockWidgetFeature(0x01)
	QDockWidget__DockWidgetMovable          QDockWidget__DockWidgetFeature = QDockWidget__DockWidgetFeature(0x02)
	QDockWidget__DockWidgetFloatable        QDockWidget__DockWidgetFeature = QDockWidget__DockWidgetFeature(0x04)
	QDockWidget__DockWidgetVerticalTitleBar QDockWidget__DockWidgetFeature = QDockWidget__DockWidgetFeature(0x08)
	QDockWidget__DockWidgetFeatureMask      QDockWidget__DockWidgetFeature = QDockWidget__DockWidgetFeature(0x0f)
	QDockWidget__NoDockWidgetFeatures       QDockWidget__DockWidgetFeature = QDockWidget__DockWidgetFeature(0x00)
	QDockWidget__Reserved                   QDockWidget__DockWidgetFeature = QDockWidget__DockWidgetFeature(0xff)
)

//go:generate stringer -type=QFileDialog__AcceptMode
//QFileDialog::AcceptMode
type QFileDialog__AcceptMode int64

const (
	QFileDialog__AcceptOpen QFileDialog__AcceptMode = QFileDialog__AcceptMode(0)
	QFileDialog__AcceptSave QFileDialog__AcceptMode = QFileDialog__AcceptMode(1)
)

//go:generate stringer -type=QFileDialog__DialogLabel
//QFileDialog::DialogLabel
type QFileDialog__DialogLabel int64

const (
	QFileDialog__LookIn   QFileDialog__DialogLabel = QFileDialog__DialogLabel(0)
	QFileDialog__FileName QFileDialog__DialogLabel = QFileDialog__DialogLabel(1)
	QFileDialog__FileType QFileDialog__DialogLabel = QFileDialog__DialogLabel(2)
	QFileDialog__Accept   QFileDialog__DialogLabel = QFileDialog__DialogLabel(3)
	QFileDialog__Reject   QFileDialog__DialogLabel = QFileDialog__DialogLabel(4)
)

//go:generate stringer -type=QFileDialog__FileMode
//QFileDialog::FileMode
type QFileDialog__FileMode int64

const (
	QFileDialog__AnyFile       QFileDialog__FileMode = QFileDialog__FileMode(0)
	QFileDialog__ExistingFile  QFileDialog__FileMode = QFileDialog__FileMode(1)
	QFileDialog__Directory     QFileDialog__FileMode = QFileDialog__FileMode(2)
	QFileDialog__ExistingFiles QFileDialog__FileMode = QFileDialog__FileMode(3)
	QFileDialog__DirectoryOnly QFileDialog__FileMode = QFileDialog__FileMode(4)
)

//go:generate stringer -type=QFileDialog__Option
//QFileDialog::Option
type QFileDialog__Option int64

const (
	QFileDialog__ShowDirsOnly                QFileDialog__Option = QFileDialog__Option(0x00000001)
	QFileDialog__DontResolveSymlinks         QFileDialog__Option = QFileDialog__Option(0x00000002)
	QFileDialog__DontConfirmOverwrite        QFileDialog__Option = QFileDialog__Option(0x00000004)
	QFileDialog__DontUseSheet                QFileDialog__Option = QFileDialog__Option(0x00000008)
	QFileDialog__DontUseNativeDialog         QFileDialog__Option = QFileDialog__Option(0x00000010)
	QFileDialog__ReadOnly                    QFileDialog__Option = QFileDialog__Option(0x00000020)
	QFileDialog__HideNameFilterDetails       QFileDialog__Option = QFileDialog__Option(0x00000040)
	QFileDialog__DontUseCustomDirectoryIcons QFileDialog__Option = QFileDialog__Option(0x00000080)
)

//go:generate stringer -type=QFileDialog__ViewMode
//QFileDialog::ViewMode
type QFileDialog__ViewMode int64

const (
	QFileDialog__Detail QFileDialog__ViewMode = QFileDialog__ViewMode(0)
	QFileDialog__List   QFileDialog__ViewMode = QFileDialog__ViewMode(1)
)

//go:generate stringer -type=QFileIconProvider__IconType
//QFileIconProvider::IconType
type QFileIconProvider__IconType int64

const (
	QFileIconProvider__Computer QFileIconProvider__IconType = QFileIconProvider__IconType(0)
	QFileIconProvider__Desktop  QFileIconProvider__IconType = QFileIconProvider__IconType(1)
	QFileIconProvider__Trashcan QFileIconProvider__IconType = QFileIconProvider__IconType(2)
	QFileIconProvider__Network  QFileIconProvider__IconType = QFileIconProvider__IconType(3)
	QFileIconProvider__Drive    QFileIconProvider__IconType = QFileIconProvider__IconType(4)
	QFileIconProvider__Folder   QFileIconProvider__IconType = QFileIconProvider__IconType(5)
	QFileIconProvider__File     QFileIconProvider__IconType = QFileIconProvider__IconType(6)
)

//go:generate stringer -type=QFileIconProvider__Option
//QFileIconProvider::Option
type QFileIconProvider__Option int64

const (
	QFileIconProvider__DontUseCustomDirectoryIcons QFileIconProvider__Option = QFileIconProvider__Option(0x00000001)
)

//go:generate stringer -type=QFileSystemModel__Roles
//QFileSystemModel::Roles
type QFileSystemModel__Roles int64

var (
	QFileSystemModel__FileIconRole    QFileSystemModel__Roles = QFileSystemModel__Roles(core.Qt__DecorationRole)
	QFileSystemModel__FilePathRole    QFileSystemModel__Roles = QFileSystemModel__Roles(C.QFileSystemModel_FilePathRole_Type())
	QFileSystemModel__FileNameRole    QFileSystemModel__Roles = QFileSystemModel__Roles(C.QFileSystemModel_FileNameRole_Type())
	QFileSystemModel__FilePermissions QFileSystemModel__Roles = QFileSystemModel__Roles(C.QFileSystemModel_FilePermissions_Type())
)

//go:generate stringer -type=QFontComboBox__FontFilter
//QFontComboBox::FontFilter
type QFontComboBox__FontFilter int64

const (
	QFontComboBox__AllFonts          QFontComboBox__FontFilter = QFontComboBox__FontFilter(0)
	QFontComboBox__ScalableFonts     QFontComboBox__FontFilter = QFontComboBox__FontFilter(0x1)
	QFontComboBox__NonScalableFonts  QFontComboBox__FontFilter = QFontComboBox__FontFilter(0x2)
	QFontComboBox__MonospacedFonts   QFontComboBox__FontFilter = QFontComboBox__FontFilter(0x4)
	QFontComboBox__ProportionalFonts QFontComboBox__FontFilter = QFontComboBox__FontFilter(0x8)
)

//go:generate stringer -type=QFontDialog__FontDialogOption
//QFontDialog::FontDialogOption
type QFontDialog__FontDialogOption int64

const (
	QFontDialog__NoButtons           QFontDialog__FontDialogOption = QFontDialog__FontDialogOption(0x00000001)
	QFontDialog__DontUseNativeDialog QFontDialog__FontDialogOption = QFontDialog__FontDialogOption(0x00000002)
	QFontDialog__ScalableFonts       QFontDialog__FontDialogOption = QFontDialog__FontDialogOption(0x00000004)
	QFontDialog__NonScalableFonts    QFontDialog__FontDialogOption = QFontDialog__FontDialogOption(0x00000008)
	QFontDialog__MonospacedFonts     QFontDialog__FontDialogOption = QFontDialog__FontDialogOption(0x00000010)
	QFontDialog__ProportionalFonts   QFontDialog__FontDialogOption = QFontDialog__FontDialogOption(0x00000020)
)

//go:generate stringer -type=QFormLayout__FieldGrowthPolicy
//QFormLayout::FieldGrowthPolicy
type QFormLayout__FieldGrowthPolicy int64

const (
	QFormLayout__FieldsStayAtSizeHint  QFormLayout__FieldGrowthPolicy = QFormLayout__FieldGrowthPolicy(0)
	QFormLayout__ExpandingFieldsGrow   QFormLayout__FieldGrowthPolicy = QFormLayout__FieldGrowthPolicy(1)
	QFormLayout__AllNonFixedFieldsGrow QFormLayout__FieldGrowthPolicy = QFormLayout__FieldGrowthPolicy(2)
)

//go:generate stringer -type=QFormLayout__ItemRole
//QFormLayout::ItemRole
type QFormLayout__ItemRole int64

const (
	QFormLayout__LabelRole    QFormLayout__ItemRole = QFormLayout__ItemRole(0)
	QFormLayout__FieldRole    QFormLayout__ItemRole = QFormLayout__ItemRole(1)
	QFormLayout__SpanningRole QFormLayout__ItemRole = QFormLayout__ItemRole(2)
)

//go:generate stringer -type=QFormLayout__RowWrapPolicy
//QFormLayout::RowWrapPolicy
type QFormLayout__RowWrapPolicy int64

const (
	QFormLayout__DontWrapRows QFormLayout__RowWrapPolicy = QFormLayout__RowWrapPolicy(0)
	QFormLayout__WrapLongRows QFormLayout__RowWrapPolicy = QFormLayout__RowWrapPolicy(1)
	QFormLayout__WrapAllRows  QFormLayout__RowWrapPolicy = QFormLayout__RowWrapPolicy(2)
)

type QFrame struct {
	QWidget
}

type QFrame_ITF interface {
	QWidget_ITF
	QFrame_PTR() *QFrame
}

func (ptr *QFrame) QFrame_PTR() *QFrame {
	return ptr
}

func (ptr *QFrame) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QFrame) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQFrame(ptr QFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFrame_PTR().Pointer()
	}
	return nil
}

func NewQFrameFromPointer(ptr unsafe.Pointer) (n *QFrame) {
	n = new(QFrame)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QFrame__Shadow
//QFrame::Shadow
type QFrame__Shadow int64

const (
	QFrame__Plain  QFrame__Shadow = QFrame__Shadow(0x0010)
	QFrame__Raised QFrame__Shadow = QFrame__Shadow(0x0020)
	QFrame__Sunken QFrame__Shadow = QFrame__Shadow(0x0030)
)

//go:generate stringer -type=QFrame__Shape
//QFrame::Shape
type QFrame__Shape int64

const (
	QFrame__NoFrame     QFrame__Shape = QFrame__Shape(0)
	QFrame__Box         QFrame__Shape = QFrame__Shape(0x0001)
	QFrame__Panel       QFrame__Shape = QFrame__Shape(0x0002)
	QFrame__WinPanel    QFrame__Shape = QFrame__Shape(0x0003)
	QFrame__HLine       QFrame__Shape = QFrame__Shape(0x0004)
	QFrame__VLine       QFrame__Shape = QFrame__Shape(0x0005)
	QFrame__StyledPanel QFrame__Shape = QFrame__Shape(0x0006)
)

//go:generate stringer -type=QFrame__StyleMask
//QFrame::StyleMask
type QFrame__StyleMask int64

var (
	QFrame__Shadow_Mask QFrame__StyleMask = QFrame__StyleMask(0x00f0)
	QFrame__Shape_Mask  QFrame__StyleMask = QFrame__StyleMask(0x000f)
)

func NewQFrame(parent QWidget_ITF, ff core.Qt__WindowType) *QFrame {
	tmpValue := NewQFrameFromPointer(C.QFrame_NewQFrame(PointerFromQWidget(parent), C.longlong(ff)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QFrame) DestroyQFrame() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QFrame_DestroyQFrame(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QGesture__GestureCancelPolicy
//QGesture::GestureCancelPolicy
type QGesture__GestureCancelPolicy int64

const (
	QGesture__CancelNone         QGesture__GestureCancelPolicy = QGesture__GestureCancelPolicy(0)
	QGesture__CancelAllInContext QGesture__GestureCancelPolicy = QGesture__GestureCancelPolicy(1)
)

//go:generate stringer -type=QGestureRecognizer__ResultFlag
//QGestureRecognizer::ResultFlag
type QGestureRecognizer__ResultFlag int64

const (
	QGestureRecognizer__Ignore           QGestureRecognizer__ResultFlag = QGestureRecognizer__ResultFlag(0x0001)
	QGestureRecognizer__MayBeGesture     QGestureRecognizer__ResultFlag = QGestureRecognizer__ResultFlag(0x0002)
	QGestureRecognizer__TriggerGesture   QGestureRecognizer__ResultFlag = QGestureRecognizer__ResultFlag(0x0004)
	QGestureRecognizer__FinishGesture    QGestureRecognizer__ResultFlag = QGestureRecognizer__ResultFlag(0x0008)
	QGestureRecognizer__CancelGesture    QGestureRecognizer__ResultFlag = QGestureRecognizer__ResultFlag(0x0010)
	QGestureRecognizer__ResultState_Mask QGestureRecognizer__ResultFlag = QGestureRecognizer__ResultFlag(0x00ff)
	QGestureRecognizer__ConsumeEventHint QGestureRecognizer__ResultFlag = QGestureRecognizer__ResultFlag(0x0100)
	QGestureRecognizer__ResultHint_Mask  QGestureRecognizer__ResultFlag = QGestureRecognizer__ResultFlag(0xff00)
)

//go:generate stringer -type=QGraphicsBlurEffect__BlurHint
//QGraphicsBlurEffect::BlurHint
type QGraphicsBlurEffect__BlurHint int64

const (
	QGraphicsBlurEffect__PerformanceHint QGraphicsBlurEffect__BlurHint = QGraphicsBlurEffect__BlurHint(0x00)
	QGraphicsBlurEffect__QualityHint     QGraphicsBlurEffect__BlurHint = QGraphicsBlurEffect__BlurHint(0x01)
	QGraphicsBlurEffect__AnimationHint   QGraphicsBlurEffect__BlurHint = QGraphicsBlurEffect__BlurHint(0x02)
)

//go:generate stringer -type=QGraphicsEffect__ChangeFlag
//QGraphicsEffect::ChangeFlag
type QGraphicsEffect__ChangeFlag int64

const (
	QGraphicsEffect__SourceAttached            QGraphicsEffect__ChangeFlag = QGraphicsEffect__ChangeFlag(0x1)
	QGraphicsEffect__SourceDetached            QGraphicsEffect__ChangeFlag = QGraphicsEffect__ChangeFlag(0x2)
	QGraphicsEffect__SourceBoundingRectChanged QGraphicsEffect__ChangeFlag = QGraphicsEffect__ChangeFlag(0x4)
	QGraphicsEffect__SourceInvalidated         QGraphicsEffect__ChangeFlag = QGraphicsEffect__ChangeFlag(0x8)
)

//go:generate stringer -type=QGraphicsEffect__PixmapPadMode
//QGraphicsEffect::PixmapPadMode
type QGraphicsEffect__PixmapPadMode int64

const (
	QGraphicsEffect__NoPad                      QGraphicsEffect__PixmapPadMode = QGraphicsEffect__PixmapPadMode(0)
	QGraphicsEffect__PadToTransparentBorder     QGraphicsEffect__PixmapPadMode = QGraphicsEffect__PixmapPadMode(1)
	QGraphicsEffect__PadToEffectiveBoundingRect QGraphicsEffect__PixmapPadMode = QGraphicsEffect__PixmapPadMode(2)
)

type QGraphicsItem struct {
	ptr unsafe.Pointer
}

type QGraphicsItem_ITF interface {
	QGraphicsItem_PTR() *QGraphicsItem
}

func (ptr *QGraphicsItem) QGraphicsItem_PTR() *QGraphicsItem {
	return ptr
}

func (ptr *QGraphicsItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGraphicsItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGraphicsItem(ptr QGraphicsItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsItemFromPointer(ptr unsafe.Pointer) (n *QGraphicsItem) {
	n = new(QGraphicsItem)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QGraphicsItem__CacheMode
//QGraphicsItem::CacheMode
type QGraphicsItem__CacheMode int64

const (
	QGraphicsItem__NoCache               QGraphicsItem__CacheMode = QGraphicsItem__CacheMode(0)
	QGraphicsItem__ItemCoordinateCache   QGraphicsItem__CacheMode = QGraphicsItem__CacheMode(1)
	QGraphicsItem__DeviceCoordinateCache QGraphicsItem__CacheMode = QGraphicsItem__CacheMode(2)
)

//go:generate stringer -type=QGraphicsItem__GraphicsItemChange
//QGraphicsItem::GraphicsItemChange
type QGraphicsItem__GraphicsItemChange int64

const (
	QGraphicsItem__ItemPositionChange                 QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(0)
	QGraphicsItem__ItemMatrixChange                   QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(1)
	QGraphicsItem__ItemVisibleChange                  QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(2)
	QGraphicsItem__ItemEnabledChange                  QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(3)
	QGraphicsItem__ItemSelectedChange                 QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(4)
	QGraphicsItem__ItemParentChange                   QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(5)
	QGraphicsItem__ItemChildAddedChange               QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(6)
	QGraphicsItem__ItemChildRemovedChange             QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(7)
	QGraphicsItem__ItemTransformChange                QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(8)
	QGraphicsItem__ItemPositionHasChanged             QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(9)
	QGraphicsItem__ItemTransformHasChanged            QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(10)
	QGraphicsItem__ItemSceneChange                    QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(11)
	QGraphicsItem__ItemVisibleHasChanged              QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(12)
	QGraphicsItem__ItemEnabledHasChanged              QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(13)
	QGraphicsItem__ItemSelectedHasChanged             QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(14)
	QGraphicsItem__ItemParentHasChanged               QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(15)
	QGraphicsItem__ItemSceneHasChanged                QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(16)
	QGraphicsItem__ItemCursorChange                   QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(17)
	QGraphicsItem__ItemCursorHasChanged               QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(18)
	QGraphicsItem__ItemToolTipChange                  QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(19)
	QGraphicsItem__ItemToolTipHasChanged              QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(20)
	QGraphicsItem__ItemFlagsChange                    QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(21)
	QGraphicsItem__ItemFlagsHaveChanged               QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(22)
	QGraphicsItem__ItemZValueChange                   QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(23)
	QGraphicsItem__ItemZValueHasChanged               QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(24)
	QGraphicsItem__ItemOpacityChange                  QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(25)
	QGraphicsItem__ItemOpacityHasChanged              QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(26)
	QGraphicsItem__ItemScenePositionHasChanged        QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(27)
	QGraphicsItem__ItemRotationChange                 QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(28)
	QGraphicsItem__ItemRotationHasChanged             QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(29)
	QGraphicsItem__ItemScaleChange                    QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(30)
	QGraphicsItem__ItemScaleHasChanged                QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(31)
	QGraphicsItem__ItemTransformOriginPointChange     QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(32)
	QGraphicsItem__ItemTransformOriginPointHasChanged QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(33)
)

//go:generate stringer -type=QGraphicsItem__GraphicsItemFlag
//QGraphicsItem::GraphicsItemFlag
type QGraphicsItem__GraphicsItemFlag int64

const (
	QGraphicsItem__ItemIsMovable                        QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x1)
	QGraphicsItem__ItemIsSelectable                     QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x2)
	QGraphicsItem__ItemIsFocusable                      QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x4)
	QGraphicsItem__ItemClipsToShape                     QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x8)
	QGraphicsItem__ItemClipsChildrenToShape             QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x10)
	QGraphicsItem__ItemIgnoresTransformations           QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x20)
	QGraphicsItem__ItemIgnoresParentOpacity             QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x40)
	QGraphicsItem__ItemDoesntPropagateOpacityToChildren QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x80)
	QGraphicsItem__ItemStacksBehindParent               QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x100)
	QGraphicsItem__ItemUsesExtendedStyleOption          QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x200)
	QGraphicsItem__ItemHasNoContents                    QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x400)
	QGraphicsItem__ItemSendsGeometryChanges             QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x800)
	QGraphicsItem__ItemAcceptsInputMethod               QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x1000)
	QGraphicsItem__ItemNegativeZStacksBehindParent      QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x2000)
	QGraphicsItem__ItemIsPanel                          QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x4000)
	QGraphicsItem__ItemIsFocusScope                     QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x8000)
	QGraphicsItem__ItemSendsScenePositionChanges        QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x10000)
	QGraphicsItem__ItemStopsClickFocusPropagation       QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x20000)
	QGraphicsItem__ItemStopsFocusHandling               QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x40000)
	QGraphicsItem__ItemContainsChildrenInShape          QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x80000)
)

//go:generate stringer -type=QGraphicsItem__PanelModality
//QGraphicsItem::PanelModality
type QGraphicsItem__PanelModality int64

const (
	QGraphicsItem__NonModal   QGraphicsItem__PanelModality = QGraphicsItem__PanelModality(0)
	QGraphicsItem__PanelModal QGraphicsItem__PanelModality = QGraphicsItem__PanelModality(1)
	QGraphicsItem__SceneModal QGraphicsItem__PanelModality = QGraphicsItem__PanelModality(2)
)

func NewQGraphicsItem(parent QGraphicsItem_ITF) *QGraphicsItem {
	return NewQGraphicsItemFromPointer(C.QGraphicsItem_NewQGraphicsItem(PointerFromQGraphicsItem(parent)))
}

//export callbackQGraphicsItem_Paint
func callbackQGraphicsItem_Paint(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paint"); signal != nil {
		(*(*func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(signal))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	}

}

func (ptr *QGraphicsItem) ConnectPaint(f func(painter *gui.QPainter, option *QStyleOptionGraphicsItem, widget *QWidget)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "paint"); signal != nil {
			f := func(painter *gui.QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
				(*(*func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(signal))(painter, option, widget)
				f(painter, option, widget)
			}
			qt.ConnectSignal(ptr.Pointer(), "paint", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "paint", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsItem) DisconnectPaint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "paint")
	}
}

func (ptr *QGraphicsItem) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsItem) SetEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QGraphicsItem) Show() {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_Show(ptr.Pointer())
	}
}

//export callbackQGraphicsItem_DestroyQGraphicsItem
func callbackQGraphicsItem_DestroyQGraphicsItem(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGraphicsItem"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGraphicsItemFromPointer(ptr).DestroyQGraphicsItemDefault()
	}
}

func (ptr *QGraphicsItem) ConnectDestroyQGraphicsItem(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGraphicsItem"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsItem", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsItem", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsItem) DisconnectDestroyQGraphicsItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGraphicsItem")
	}
}

func (ptr *QGraphicsItem) DestroyQGraphicsItem() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsItem_DestroyQGraphicsItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsItem) DestroyQGraphicsItemDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsItem_DestroyQGraphicsItemDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsItem) Group() *QGraphicsItemGroup {
	if ptr.Pointer() != nil {
		return NewQGraphicsItemGroupFromPointer(C.QGraphicsItem_Group(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsItem) Window() *QGraphicsWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQGraphicsWidgetFromPointer(C.QGraphicsItem_Window(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsItem) Pos() *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QGraphicsItem_Pos(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsItem_BoundingRect
func callbackQGraphicsItem_BoundingRect(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "boundingRect"); signal != nil {
		return core.PointerFromQRectF((*(*func() *core.QRectF)(signal))())
	}

	return core.PointerFromQRectF(core.NewQRectF())
}

func (ptr *QGraphicsItem) ConnectBoundingRect(f func() *core.QRectF) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "boundingRect"); signal != nil {
			f := func() *core.QRectF {
				(*(*func() *core.QRectF)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "boundingRect", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "boundingRect", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsItem) DisconnectBoundingRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "boundingRect")
	}
}

func (ptr *QGraphicsItem) BoundingRect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QGraphicsItem_BoundingRect(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsItem) Data(key int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QGraphicsItem_Data(ptr.Pointer(), C.int(int32(key))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsItem_Contains
func callbackQGraphicsItem_Contains(ptr unsafe.Pointer, point unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "contains"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QPointF) bool)(signal))(core.NewQPointFFromPointer(point)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsItemFromPointer(ptr).ContainsDefault(core.NewQPointFFromPointer(point)))))
}

func (ptr *QGraphicsItem) ConnectContains(f func(point *core.QPointF) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "contains"); signal != nil {
			f := func(point *core.QPointF) bool {
				(*(*func(*core.QPointF) bool)(signal))(point)
				return f(point)
			}
			qt.ConnectSignal(ptr.Pointer(), "contains", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contains", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsItem) DisconnectContains() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "contains")
	}
}

func (ptr *QGraphicsItem) Contains(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point))) != 0
	}
	return false
}

func (ptr *QGraphicsItem) ContainsDefault(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsItem_ContainsDefault(ptr.Pointer(), core.PointerFromQPointF(point))) != 0
	}
	return false
}

//export callbackQGraphicsItem_Type
func callbackQGraphicsItem_Type(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQGraphicsItemFromPointer(ptr).TypeDefault()))
}

func (ptr *QGraphicsItem) ConnectType(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsItem) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QGraphicsItem) Type() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGraphicsItem_Type(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsItem) TypeDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGraphicsItem_TypeDefault(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsItem) X() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsItem_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsItem) Y() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsItem_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsItem) __setTransformations_transformations_atList(i int) *QGraphicsTransform {
	if ptr.Pointer() != nil {
		tmpValue := NewQGraphicsTransformFromPointer(C.QGraphicsItem___setTransformations_transformations_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsItem) __setTransformations_transformations_setList(i QGraphicsTransform_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem___setTransformations_transformations_setList(ptr.Pointer(), PointerFromQGraphicsTransform(i))
	}
}

func (ptr *QGraphicsItem) __setTransformations_transformations_newList() unsafe.Pointer {
	return C.QGraphicsItem___setTransformations_transformations_newList(ptr.Pointer())
}

func (ptr *QGraphicsItem) __childItems_atList(i int) *QGraphicsItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsItem___childItems_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QGraphicsItem) __childItems_setList(i QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem___childItems_setList(ptr.Pointer(), PointerFromQGraphicsItem(i))
	}
}

func (ptr *QGraphicsItem) __childItems_newList() unsafe.Pointer {
	return C.QGraphicsItem___childItems_newList(ptr.Pointer())
}

func (ptr *QGraphicsItem) __children_atList(i int) *QGraphicsItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsItem___children_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QGraphicsItem) __children_setList(i QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem___children_setList(ptr.Pointer(), PointerFromQGraphicsItem(i))
	}
}

func (ptr *QGraphicsItem) __children_newList() unsafe.Pointer {
	return C.QGraphicsItem___children_newList(ptr.Pointer())
}

func (ptr *QGraphicsItem) __collidingItems_atList(i int) *QGraphicsItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsItem___collidingItems_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QGraphicsItem) __collidingItems_setList(i QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem___collidingItems_setList(ptr.Pointer(), PointerFromQGraphicsItem(i))
	}
}

func (ptr *QGraphicsItem) __collidingItems_newList() unsafe.Pointer {
	return C.QGraphicsItem___collidingItems_newList(ptr.Pointer())
}

func (ptr *QGraphicsItem) __transformations_atList(i int) *QGraphicsTransform {
	if ptr.Pointer() != nil {
		tmpValue := NewQGraphicsTransformFromPointer(C.QGraphicsItem___transformations_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsItem) __transformations_setList(i QGraphicsTransform_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem___transformations_setList(ptr.Pointer(), PointerFromQGraphicsTransform(i))
	}
}

func (ptr *QGraphicsItem) __transformations_newList() unsafe.Pointer {
	return C.QGraphicsItem___transformations_newList(ptr.Pointer())
}

type QGraphicsItemGroup struct {
	QGraphicsItem
}

type QGraphicsItemGroup_ITF interface {
	QGraphicsItem_ITF
	QGraphicsItemGroup_PTR() *QGraphicsItemGroup
}

func (ptr *QGraphicsItemGroup) QGraphicsItemGroup_PTR() *QGraphicsItemGroup {
	return ptr
}

func (ptr *QGraphicsItemGroup) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsItem_PTR().Pointer()
	}
	return nil
}

func (ptr *QGraphicsItemGroup) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGraphicsItem_PTR().SetPointer(p)
	}
}

func PointerFromQGraphicsItemGroup(ptr QGraphicsItemGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsItemGroup_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsItemGroupFromPointer(ptr unsafe.Pointer) (n *QGraphicsItemGroup) {
	n = new(QGraphicsItemGroup)
	n.SetPointer(ptr)
	return
}
func NewQGraphicsItemGroup(parent QGraphicsItem_ITF) *QGraphicsItemGroup {
	return NewQGraphicsItemGroupFromPointer(C.QGraphicsItemGroup_NewQGraphicsItemGroup(PointerFromQGraphicsItem(parent)))
}

func (ptr *QGraphicsItemGroup) DestroyQGraphicsItemGroup() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsItemGroup_DestroyQGraphicsItemGroup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQGraphicsItemGroup_Paint
func callbackQGraphicsItemGroup_Paint(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paint"); signal != nil {
		(*(*func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(signal))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	} else {
		NewQGraphicsItemGroupFromPointer(ptr).PaintDefault(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	}
}

func (ptr *QGraphicsItemGroup) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsItemGroup) PaintDefault(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

//export callbackQGraphicsItemGroup_BoundingRect
func callbackQGraphicsItemGroup_BoundingRect(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "boundingRect"); signal != nil {
		return core.PointerFromQRectF((*(*func() *core.QRectF)(signal))())
	}

	return core.PointerFromQRectF(NewQGraphicsItemGroupFromPointer(ptr).BoundingRectDefault())
}

func (ptr *QGraphicsItemGroup) BoundingRect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QGraphicsItemGroup_BoundingRect(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsItemGroup) BoundingRectDefault() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QGraphicsItemGroup_BoundingRectDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

type QGraphicsLayout struct {
	QGraphicsLayoutItem
}

type QGraphicsLayout_ITF interface {
	QGraphicsLayoutItem_ITF
	QGraphicsLayout_PTR() *QGraphicsLayout
}

func (ptr *QGraphicsLayout) QGraphicsLayout_PTR() *QGraphicsLayout {
	return ptr
}

func (ptr *QGraphicsLayout) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsLayoutItem_PTR().Pointer()
	}
	return nil
}

func (ptr *QGraphicsLayout) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGraphicsLayoutItem_PTR().SetPointer(p)
	}
}

func PointerFromQGraphicsLayout(ptr QGraphicsLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsLayout_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsLayoutFromPointer(ptr unsafe.Pointer) (n *QGraphicsLayout) {
	n = new(QGraphicsLayout)
	n.SetPointer(ptr)
	return
}
func NewQGraphicsLayout(parent QGraphicsLayoutItem_ITF) *QGraphicsLayout {
	return NewQGraphicsLayoutFromPointer(C.QGraphicsLayout_NewQGraphicsLayout(PointerFromQGraphicsLayoutItem(parent)))
}

//export callbackQGraphicsLayout_RemoveAt
func callbackQGraphicsLayout_RemoveAt(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "removeAt"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *QGraphicsLayout) ConnectRemoveAt(f func(index int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "removeAt"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "removeAt", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removeAt", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsLayout) DisconnectRemoveAt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "removeAt")
	}
}

func (ptr *QGraphicsLayout) RemoveAt(index int) {
	if ptr.Pointer() != nil {
		C.QGraphicsLayout_RemoveAt(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QGraphicsLayout) SetContentsMargins(left float64, top float64, right float64, bottom float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsLayout_SetContentsMargins(ptr.Pointer(), C.double(left), C.double(top), C.double(right), C.double(bottom))
	}
}

func (ptr *QGraphicsLayout) DestroyQGraphicsLayout() {
	if ptr.Pointer() != nil {
		C.QGraphicsLayout_DestroyQGraphicsLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQGraphicsLayout_ItemAt
func callbackQGraphicsLayout_ItemAt(ptr unsafe.Pointer, i C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemAt"); signal != nil {
		return PointerFromQGraphicsLayoutItem((*(*func(int) *QGraphicsLayoutItem)(signal))(int(int32(i))))
	}

	return PointerFromQGraphicsLayoutItem(nil)
}

func (ptr *QGraphicsLayout) ConnectItemAt(f func(i int) *QGraphicsLayoutItem) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "itemAt"); signal != nil {
			f := func(i int) *QGraphicsLayoutItem {
				(*(*func(int) *QGraphicsLayoutItem)(signal))(i)
				return f(i)
			}
			qt.ConnectSignal(ptr.Pointer(), "itemAt", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "itemAt", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsLayout) DisconnectItemAt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "itemAt")
	}
}

func (ptr *QGraphicsLayout) ItemAt(i int) *QGraphicsLayoutItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsLayoutItemFromPointer(C.QGraphicsLayout_ItemAt(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

//export callbackQGraphicsLayout_Count
func callbackQGraphicsLayout_Count(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "count"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(0))
}

func (ptr *QGraphicsLayout) ConnectCount(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "count"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "count", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "count", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsLayout) DisconnectCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "count")
	}
}

func (ptr *QGraphicsLayout) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGraphicsLayout_Count(ptr.Pointer())))
	}
	return 0
}

//export callbackQGraphicsLayout_SizeHint
func callbackQGraphicsLayout_SizeHint(ptr unsafe.Pointer, which C.longlong, constraint unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSizeF((*(*func(core.Qt__SizeHint, *core.QSizeF) *core.QSizeF)(signal))(core.Qt__SizeHint(which), core.NewQSizeFFromPointer(constraint)))
	}

	return core.PointerFromQSizeF(NewQGraphicsLayoutFromPointer(ptr).SizeHintDefault(core.Qt__SizeHint(which), core.NewQSizeFFromPointer(constraint)))
}

func (ptr *QGraphicsLayout) SizeHint(which core.Qt__SizeHint, constraint core.QSizeF_ITF) *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QGraphicsLayout_SizeHint(ptr.Pointer(), C.longlong(which), core.PointerFromQSizeF(constraint)))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsLayout) SizeHintDefault(which core.Qt__SizeHint, constraint core.QSizeF_ITF) *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QGraphicsLayout_SizeHintDefault(ptr.Pointer(), C.longlong(which), core.PointerFromQSizeF(constraint)))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

type QGraphicsLayoutItem struct {
	ptr unsafe.Pointer
}

type QGraphicsLayoutItem_ITF interface {
	QGraphicsLayoutItem_PTR() *QGraphicsLayoutItem
}

func (ptr *QGraphicsLayoutItem) QGraphicsLayoutItem_PTR() *QGraphicsLayoutItem {
	return ptr
}

func (ptr *QGraphicsLayoutItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGraphicsLayoutItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGraphicsLayoutItem(ptr QGraphicsLayoutItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsLayoutItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsLayoutItemFromPointer(ptr unsafe.Pointer) (n *QGraphicsLayoutItem) {
	n = new(QGraphicsLayoutItem)
	n.SetPointer(ptr)
	return
}
func NewQGraphicsLayoutItem(parent QGraphicsLayoutItem_ITF, isLayout bool) *QGraphicsLayoutItem {
	return NewQGraphicsLayoutItemFromPointer(C.QGraphicsLayoutItem_NewQGraphicsLayoutItem(PointerFromQGraphicsLayoutItem(parent), C.char(int8(qt.GoBoolToInt(isLayout)))))
}

func (ptr *QGraphicsLayoutItem) SetMinimumSize(size core.QSizeF_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsLayoutItem_SetMinimumSize(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

func (ptr *QGraphicsLayoutItem) SetMinimumSize2(w float64, h float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsLayoutItem_SetMinimumSize2(ptr.Pointer(), C.double(w), C.double(h))
	}
}

//export callbackQGraphicsLayoutItem_DestroyQGraphicsLayoutItem
func callbackQGraphicsLayoutItem_DestroyQGraphicsLayoutItem(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGraphicsLayoutItem"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGraphicsLayoutItemFromPointer(ptr).DestroyQGraphicsLayoutItemDefault()
	}
}

func (ptr *QGraphicsLayoutItem) ConnectDestroyQGraphicsLayoutItem(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGraphicsLayoutItem"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsLayoutItem", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsLayoutItem", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsLayoutItem) DisconnectDestroyQGraphicsLayoutItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGraphicsLayoutItem")
	}
}

func (ptr *QGraphicsLayoutItem) DestroyQGraphicsLayoutItem() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsLayoutItem_DestroyQGraphicsLayoutItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsLayoutItem) DestroyQGraphicsLayoutItemDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsLayoutItem_DestroyQGraphicsLayoutItemDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsLayoutItem) MinimumSize() *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QGraphicsLayoutItem_MinimumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsLayoutItem_SizeHint
func callbackQGraphicsLayoutItem_SizeHint(ptr unsafe.Pointer, which C.longlong, constraint unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSizeF((*(*func(core.Qt__SizeHint, *core.QSizeF) *core.QSizeF)(signal))(core.Qt__SizeHint(which), core.NewQSizeFFromPointer(constraint)))
	}

	return core.PointerFromQSizeF(core.NewQSizeF())
}

func (ptr *QGraphicsLayoutItem) ConnectSizeHint(f func(which core.Qt__SizeHint, constraint *core.QSizeF) *core.QSizeF) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sizeHint"); signal != nil {
			f := func(which core.Qt__SizeHint, constraint *core.QSizeF) *core.QSizeF {
				(*(*func(core.Qt__SizeHint, *core.QSizeF) *core.QSizeF)(signal))(which, constraint)
				return f(which, constraint)
			}
			qt.ConnectSignal(ptr.Pointer(), "sizeHint", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sizeHint", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsLayoutItem) DisconnectSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sizeHint")
	}
}

func (ptr *QGraphicsLayoutItem) SizeHint(which core.Qt__SizeHint, constraint core.QSizeF_ITF) *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QGraphicsLayoutItem_SizeHint(ptr.Pointer(), C.longlong(which), core.PointerFromQSizeF(constraint)))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

type QGraphicsObject struct {
	core.QObject
	QGraphicsItem
}

type QGraphicsObject_ITF interface {
	core.QObject_ITF
	QGraphicsItem_ITF
	QGraphicsObject_PTR() *QGraphicsObject
}

func (ptr *QGraphicsObject) QGraphicsObject_PTR() *QGraphicsObject {
	return ptr
}

func (ptr *QGraphicsObject) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGraphicsObject) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QGraphicsItem_PTR().SetPointer(p)
	}
}

func PointerFromQGraphicsObject(ptr QGraphicsObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsObject_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsObjectFromPointer(ptr unsafe.Pointer) (n *QGraphicsObject) {
	n = new(QGraphicsObject)
	n.SetPointer(ptr)
	return
}
func NewQGraphicsObject(parent QGraphicsItem_ITF) *QGraphicsObject {
	tmpValue := NewQGraphicsObjectFromPointer(C.QGraphicsObject_NewQGraphicsObject(PointerFromQGraphicsItem(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGraphicsObject) DestroyQGraphicsObject() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsObject_DestroyQGraphicsObject(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsObject) SetEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QGraphicsObject) Pos() *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QGraphicsObject_Pos(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsObject) X() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsObject_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsObject) Y() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsObject_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsObject) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QGraphicsObject___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsObject) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QGraphicsObject) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QGraphicsObject___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QGraphicsObject) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGraphicsObject___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsObject) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsObject) __findChildren_newList2() unsafe.Pointer {
	return C.QGraphicsObject___findChildren_newList2(ptr.Pointer())
}

func (ptr *QGraphicsObject) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGraphicsObject___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsObject) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsObject) __findChildren_newList3() unsafe.Pointer {
	return C.QGraphicsObject___findChildren_newList3(ptr.Pointer())
}

func (ptr *QGraphicsObject) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGraphicsObject___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsObject) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsObject) __findChildren_newList() unsafe.Pointer {
	return C.QGraphicsObject___findChildren_newList(ptr.Pointer())
}

func (ptr *QGraphicsObject) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGraphicsObject___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsObject) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsObject) __children_newList() unsafe.Pointer {
	return C.QGraphicsObject___children_newList(ptr.Pointer())
}

//export callbackQGraphicsObject_ObjectNameChanged
func callbackQGraphicsObject_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQGraphicsObject_Paint
func callbackQGraphicsObject_Paint(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paint"); signal != nil {
		(*(*func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(signal))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	} else {
		NewQGraphicsObjectFromPointer(ptr).PaintDefault(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	}
}

func (ptr *QGraphicsObject) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsObject) PaintDefault(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

//export callbackQGraphicsObject_BoundingRect
func callbackQGraphicsObject_BoundingRect(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "boundingRect"); signal != nil {
		return core.PointerFromQRectF((*(*func() *core.QRectF)(signal))())
	}

	return core.PointerFromQRectF(NewQGraphicsObjectFromPointer(ptr).BoundingRectDefault())
}

func (ptr *QGraphicsObject) BoundingRect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QGraphicsObject_BoundingRect(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsObject) BoundingRectDefault() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QGraphicsObject_BoundingRectDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

//go:generate stringer -type=QGraphicsPixmapItem__ShapeMode
//QGraphicsPixmapItem::ShapeMode
type QGraphicsPixmapItem__ShapeMode int64

const (
	QGraphicsPixmapItem__MaskShape          QGraphicsPixmapItem__ShapeMode = QGraphicsPixmapItem__ShapeMode(0)
	QGraphicsPixmapItem__BoundingRectShape  QGraphicsPixmapItem__ShapeMode = QGraphicsPixmapItem__ShapeMode(1)
	QGraphicsPixmapItem__HeuristicMaskShape QGraphicsPixmapItem__ShapeMode = QGraphicsPixmapItem__ShapeMode(2)
)

//go:generate stringer -type=QGraphicsScene__ItemIndexMethod
//QGraphicsScene::ItemIndexMethod
type QGraphicsScene__ItemIndexMethod int64

const (
	QGraphicsScene__BspTreeIndex QGraphicsScene__ItemIndexMethod = QGraphicsScene__ItemIndexMethod(0)
	QGraphicsScene__NoIndex      QGraphicsScene__ItemIndexMethod = QGraphicsScene__ItemIndexMethod(-1)
)

//go:generate stringer -type=QGraphicsScene__SceneLayer
//QGraphicsScene::SceneLayer
type QGraphicsScene__SceneLayer int64

const (
	QGraphicsScene__ItemLayer       QGraphicsScene__SceneLayer = QGraphicsScene__SceneLayer(0x1)
	QGraphicsScene__BackgroundLayer QGraphicsScene__SceneLayer = QGraphicsScene__SceneLayer(0x2)
	QGraphicsScene__ForegroundLayer QGraphicsScene__SceneLayer = QGraphicsScene__SceneLayer(0x4)
	QGraphicsScene__AllLayers       QGraphicsScene__SceneLayer = QGraphicsScene__SceneLayer(0xffff)
)

//go:generate stringer -type=QGraphicsSceneContextMenuEvent__Reason
//QGraphicsSceneContextMenuEvent::Reason
type QGraphicsSceneContextMenuEvent__Reason int64

const (
	QGraphicsSceneContextMenuEvent__Mouse    QGraphicsSceneContextMenuEvent__Reason = QGraphicsSceneContextMenuEvent__Reason(0)
	QGraphicsSceneContextMenuEvent__Keyboard QGraphicsSceneContextMenuEvent__Reason = QGraphicsSceneContextMenuEvent__Reason(1)
	QGraphicsSceneContextMenuEvent__Other    QGraphicsSceneContextMenuEvent__Reason = QGraphicsSceneContextMenuEvent__Reason(2)
)

type QGraphicsTransform struct {
	core.QObject
}

type QGraphicsTransform_ITF interface {
	core.QObject_ITF
	QGraphicsTransform_PTR() *QGraphicsTransform
}

func (ptr *QGraphicsTransform) QGraphicsTransform_PTR() *QGraphicsTransform {
	return ptr
}

func (ptr *QGraphicsTransform) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGraphicsTransform) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQGraphicsTransform(ptr QGraphicsTransform_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsTransform_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsTransformFromPointer(ptr unsafe.Pointer) (n *QGraphicsTransform) {
	n = new(QGraphicsTransform)
	n.SetPointer(ptr)
	return
}
func NewQGraphicsTransform(parent core.QObject_ITF) *QGraphicsTransform {
	tmpValue := NewQGraphicsTransformFromPointer(C.QGraphicsTransform_NewQGraphicsTransform(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGraphicsTransform) DestroyQGraphicsTransform() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsTransform_DestroyQGraphicsTransform(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQGraphicsTransform_ApplyTo
func callbackQGraphicsTransform_ApplyTo(ptr unsafe.Pointer, matrix unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "applyTo"); signal != nil {
		(*(*func(*gui.QMatrix4x4))(signal))(gui.NewQMatrix4x4FromPointer(matrix))
	}

}

func (ptr *QGraphicsTransform) ConnectApplyTo(f func(matrix *gui.QMatrix4x4)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "applyTo"); signal != nil {
			f := func(matrix *gui.QMatrix4x4) {
				(*(*func(*gui.QMatrix4x4))(signal))(matrix)
				f(matrix)
			}
			qt.ConnectSignal(ptr.Pointer(), "applyTo", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "applyTo", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsTransform) DisconnectApplyTo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "applyTo")
	}
}

func (ptr *QGraphicsTransform) ApplyTo(matrix gui.QMatrix4x4_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTransform_ApplyTo(ptr.Pointer(), gui.PointerFromQMatrix4x4(matrix))
	}
}

func (ptr *QGraphicsTransform) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QGraphicsTransform___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsTransform) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTransform___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QGraphicsTransform) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QGraphicsTransform___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QGraphicsTransform) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGraphicsTransform___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsTransform) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTransform___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsTransform) __findChildren_newList2() unsafe.Pointer {
	return C.QGraphicsTransform___findChildren_newList2(ptr.Pointer())
}

func (ptr *QGraphicsTransform) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGraphicsTransform___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsTransform) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTransform___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsTransform) __findChildren_newList3() unsafe.Pointer {
	return C.QGraphicsTransform___findChildren_newList3(ptr.Pointer())
}

func (ptr *QGraphicsTransform) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGraphicsTransform___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsTransform) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTransform___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsTransform) __findChildren_newList() unsafe.Pointer {
	return C.QGraphicsTransform___findChildren_newList(ptr.Pointer())
}

func (ptr *QGraphicsTransform) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGraphicsTransform___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsTransform) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTransform___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsTransform) __children_newList() unsafe.Pointer {
	return C.QGraphicsTransform___children_newList(ptr.Pointer())
}

//export callbackQGraphicsTransform_ObjectNameChanged
func callbackQGraphicsTransform_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//go:generate stringer -type=QGraphicsView__CacheModeFlag
//QGraphicsView::CacheModeFlag
type QGraphicsView__CacheModeFlag int64

const (
	QGraphicsView__CacheNone       QGraphicsView__CacheModeFlag = QGraphicsView__CacheModeFlag(0x0)
	QGraphicsView__CacheBackground QGraphicsView__CacheModeFlag = QGraphicsView__CacheModeFlag(0x1)
)

//go:generate stringer -type=QGraphicsView__DragMode
//QGraphicsView::DragMode
type QGraphicsView__DragMode int64

const (
	QGraphicsView__NoDrag         QGraphicsView__DragMode = QGraphicsView__DragMode(0)
	QGraphicsView__ScrollHandDrag QGraphicsView__DragMode = QGraphicsView__DragMode(1)
	QGraphicsView__RubberBandDrag QGraphicsView__DragMode = QGraphicsView__DragMode(2)
)

//go:generate stringer -type=QGraphicsView__OptimizationFlag
//QGraphicsView::OptimizationFlag
type QGraphicsView__OptimizationFlag int64

const (
	QGraphicsView__DontClipPainter           QGraphicsView__OptimizationFlag = QGraphicsView__OptimizationFlag(0x1)
	QGraphicsView__DontSavePainterState      QGraphicsView__OptimizationFlag = QGraphicsView__OptimizationFlag(0x2)
	QGraphicsView__DontAdjustForAntialiasing QGraphicsView__OptimizationFlag = QGraphicsView__OptimizationFlag(0x4)
	QGraphicsView__IndirectPainting          QGraphicsView__OptimizationFlag = QGraphicsView__OptimizationFlag(0x8)
)

//go:generate stringer -type=QGraphicsView__ViewportAnchor
//QGraphicsView::ViewportAnchor
type QGraphicsView__ViewportAnchor int64

const (
	QGraphicsView__NoAnchor         QGraphicsView__ViewportAnchor = QGraphicsView__ViewportAnchor(0)
	QGraphicsView__AnchorViewCenter QGraphicsView__ViewportAnchor = QGraphicsView__ViewportAnchor(1)
	QGraphicsView__AnchorUnderMouse QGraphicsView__ViewportAnchor = QGraphicsView__ViewportAnchor(2)
)

//go:generate stringer -type=QGraphicsView__ViewportUpdateMode
//QGraphicsView::ViewportUpdateMode
type QGraphicsView__ViewportUpdateMode int64

const (
	QGraphicsView__FullViewportUpdate         QGraphicsView__ViewportUpdateMode = QGraphicsView__ViewportUpdateMode(0)
	QGraphicsView__MinimalViewportUpdate      QGraphicsView__ViewportUpdateMode = QGraphicsView__ViewportUpdateMode(1)
	QGraphicsView__SmartViewportUpdate        QGraphicsView__ViewportUpdateMode = QGraphicsView__ViewportUpdateMode(2)
	QGraphicsView__NoViewportUpdate           QGraphicsView__ViewportUpdateMode = QGraphicsView__ViewportUpdateMode(3)
	QGraphicsView__BoundingRectViewportUpdate QGraphicsView__ViewportUpdateMode = QGraphicsView__ViewportUpdateMode(4)
)

type QGraphicsWidget struct {
	QGraphicsObject
	QGraphicsLayoutItem
}

type QGraphicsWidget_ITF interface {
	QGraphicsObject_ITF
	QGraphicsLayoutItem_ITF
	QGraphicsWidget_PTR() *QGraphicsWidget
}

func (ptr *QGraphicsWidget) QGraphicsWidget_PTR() *QGraphicsWidget {
	return ptr
}

func (ptr *QGraphicsWidget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGraphicsWidget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGraphicsObject_PTR().SetPointer(p)
		ptr.QGraphicsLayoutItem_PTR().SetPointer(p)
	}
}

func PointerFromQGraphicsWidget(ptr QGraphicsWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsWidget_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsWidgetFromPointer(ptr unsafe.Pointer) (n *QGraphicsWidget) {
	n = new(QGraphicsWidget)
	n.SetPointer(ptr)
	return
}
func NewQGraphicsWidget(parent QGraphicsItem_ITF, wFlags core.Qt__WindowType) *QGraphicsWidget {
	tmpValue := NewQGraphicsWidgetFromPointer(C.QGraphicsWidget_NewQGraphicsWidget(PointerFromQGraphicsItem(parent), C.longlong(wFlags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGraphicsWidget_Close
func callbackQGraphicsWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QGraphicsWidget) ConnectClose(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "close"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "close", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "close", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsWidget) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "close")
	}
}

func (ptr *QGraphicsWidget) Close() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsWidget_Close(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGraphicsWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsWidget_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGraphicsWidget) SetContentsMargins(left float64, top float64, right float64, bottom float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetContentsMargins(ptr.Pointer(), C.double(left), C.double(top), C.double(right), C.double(bottom))
	}
}

func (ptr *QGraphicsWidget) SetWindowTitle(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.QGraphicsWidget_SetWindowTitle(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))})
	}
}

func (ptr *QGraphicsWidget) DestroyQGraphicsWidget() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsWidget_DestroyQGraphicsWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsWidget) Layout() *QGraphicsLayout {
	if ptr.Pointer() != nil {
		return NewQGraphicsLayoutFromPointer(C.QGraphicsWidget_Layout(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsWidget) Size() *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QGraphicsWidget_Size(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWidget) WindowTitle() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGraphicsWidget_WindowTitle(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsWidget) MinimumSize() *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QGraphicsWidget_MinimumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWidget) SetMinimumSize(minimumSize core.QSizeF_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetMinimumSize(ptr.Pointer(), core.PointerFromQSizeF(minimumSize))
	}
}

func (ptr *QGraphicsWidget) __addActions_actions_atList(i int) *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QGraphicsWidget___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWidget) __addActions_actions_setList(i QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget___addActions_actions_setList(ptr.Pointer(), PointerFromQAction(i))
	}
}

func (ptr *QGraphicsWidget) __addActions_actions_newList() unsafe.Pointer {
	return C.QGraphicsWidget___addActions_actions_newList(ptr.Pointer())
}

func (ptr *QGraphicsWidget) __insertActions_actions_atList(i int) *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QGraphicsWidget___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWidget) __insertActions_actions_setList(i QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget___insertActions_actions_setList(ptr.Pointer(), PointerFromQAction(i))
	}
}

func (ptr *QGraphicsWidget) __insertActions_actions_newList() unsafe.Pointer {
	return C.QGraphicsWidget___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *QGraphicsWidget) __actions_atList(i int) *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QGraphicsWidget___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWidget) __actions_setList(i QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget___actions_setList(ptr.Pointer(), PointerFromQAction(i))
	}
}

func (ptr *QGraphicsWidget) __actions_newList() unsafe.Pointer {
	return C.QGraphicsWidget___actions_newList(ptr.Pointer())
}

//export callbackQGraphicsWidget_SizeHint
func callbackQGraphicsWidget_SizeHint(ptr unsafe.Pointer, which C.longlong, constraint unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSizeF((*(*func(core.Qt__SizeHint, *core.QSizeF) *core.QSizeF)(signal))(core.Qt__SizeHint(which), core.NewQSizeFFromPointer(constraint)))
	}

	return core.PointerFromQSizeF(NewQGraphicsWidgetFromPointer(ptr).SizeHintDefault(core.Qt__SizeHint(which), core.NewQSizeFFromPointer(constraint)))
}

func (ptr *QGraphicsWidget) SizeHint(which core.Qt__SizeHint, constraint core.QSizeF_ITF) *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QGraphicsWidget_SizeHint(ptr.Pointer(), C.longlong(which), core.PointerFromQSizeF(constraint)))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWidget) SizeHintDefault(which core.Qt__SizeHint, constraint core.QSizeF_ITF) *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QGraphicsWidget_SizeHintDefault(ptr.Pointer(), C.longlong(which), core.PointerFromQSizeF(constraint)))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

type QGroupBox struct {
	QWidget
}

type QGroupBox_ITF interface {
	QWidget_ITF
	QGroupBox_PTR() *QGroupBox
}

func (ptr *QGroupBox) QGroupBox_PTR() *QGroupBox {
	return ptr
}

func (ptr *QGroupBox) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QGroupBox) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQGroupBox(ptr QGroupBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGroupBox_PTR().Pointer()
	}
	return nil
}

func NewQGroupBoxFromPointer(ptr unsafe.Pointer) (n *QGroupBox) {
	n = new(QGroupBox)
	n.SetPointer(ptr)
	return
}
func NewQGroupBox(parent QWidget_ITF) *QGroupBox {
	tmpValue := NewQGroupBoxFromPointer(C.QGroupBox_NewQGroupBox(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQGroupBox2(title string, parent QWidget_ITF) *QGroupBox {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	tmpValue := NewQGroupBoxFromPointer(C.QGroupBox_NewQGroupBox2(C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGroupBox_Clicked
func callbackQGroupBox_Clicked(ptr unsafe.Pointer, checked C.char) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		(*(*func(bool))(signal))(int8(checked) != 0)
	}

}

func (ptr *QGroupBox) ConnectClicked(f func(checked bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.QGroupBox_ConnectClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			f := func(checked bool) {
				(*(*func(bool))(signal))(checked)
				f(checked)
			}
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGroupBox) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QGroupBox_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *QGroupBox) Clicked(checked bool) {
	if ptr.Pointer() != nil {
		C.QGroupBox_Clicked(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(checked))))
	}
}

func (ptr *QGroupBox) SetTitle(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.QGroupBox_SetTitle(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))})
	}
}

func (ptr *QGroupBox) DestroyQGroupBox() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGroupBox_DestroyQGroupBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGroupBox) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGroupBox_Title(ptr.Pointer()))
	}
	return ""
}

type QHBoxLayout struct {
	QBoxLayout
}

type QHBoxLayout_ITF interface {
	QBoxLayout_ITF
	QHBoxLayout_PTR() *QHBoxLayout
}

func (ptr *QHBoxLayout) QHBoxLayout_PTR() *QHBoxLayout {
	return ptr
}

func (ptr *QHBoxLayout) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QBoxLayout_PTR().Pointer()
	}
	return nil
}

func (ptr *QHBoxLayout) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QBoxLayout_PTR().SetPointer(p)
	}
}

func PointerFromQHBoxLayout(ptr QHBoxLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHBoxLayout_PTR().Pointer()
	}
	return nil
}

func NewQHBoxLayoutFromPointer(ptr unsafe.Pointer) (n *QHBoxLayout) {
	n = new(QHBoxLayout)
	n.SetPointer(ptr)
	return
}
func NewQHBoxLayout() *QHBoxLayout {
	tmpValue := NewQHBoxLayoutFromPointer(C.QHBoxLayout_NewQHBoxLayout())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQHBoxLayout2(parent QWidget_ITF) *QHBoxLayout {
	tmpValue := NewQHBoxLayoutFromPointer(C.QHBoxLayout_NewQHBoxLayout2(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QHBoxLayout) DestroyQHBoxLayout() {
	if ptr.Pointer() != nil {
		C.QHBoxLayout_DestroyQHBoxLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QHeaderView__ResizeMode
//QHeaderView::ResizeMode
type QHeaderView__ResizeMode int64

const (
	QHeaderView__Interactive      QHeaderView__ResizeMode = QHeaderView__ResizeMode(0)
	QHeaderView__Stretch          QHeaderView__ResizeMode = QHeaderView__ResizeMode(1)
	QHeaderView__Fixed            QHeaderView__ResizeMode = QHeaderView__ResizeMode(2)
	QHeaderView__ResizeToContents QHeaderView__ResizeMode = QHeaderView__ResizeMode(3)
	QHeaderView__Custom           QHeaderView__ResizeMode = QHeaderView__ResizeMode(QHeaderView__Fixed)
)

//go:generate stringer -type=QInputDialog__InputDialogOption
//QInputDialog::InputDialogOption
type QInputDialog__InputDialogOption int64

const (
	QInputDialog__NoButtons                    QInputDialog__InputDialogOption = QInputDialog__InputDialogOption(0x00000001)
	QInputDialog__UseListViewForComboBoxItems  QInputDialog__InputDialogOption = QInputDialog__InputDialogOption(0x00000002)
	QInputDialog__UsePlainTextEditForTextInput QInputDialog__InputDialogOption = QInputDialog__InputDialogOption(0x00000004)
)

//go:generate stringer -type=QInputDialog__InputMode
//QInputDialog::InputMode
type QInputDialog__InputMode int64

const (
	QInputDialog__TextInput   QInputDialog__InputMode = QInputDialog__InputMode(0)
	QInputDialog__IntInput    QInputDialog__InputMode = QInputDialog__InputMode(1)
	QInputDialog__DoubleInput QInputDialog__InputMode = QInputDialog__InputMode(2)
)

//go:generate stringer -type=QLCDNumber__Mode
//QLCDNumber::Mode
type QLCDNumber__Mode int64

const (
	QLCDNumber__Hex QLCDNumber__Mode = QLCDNumber__Mode(0)
	QLCDNumber__Dec QLCDNumber__Mode = QLCDNumber__Mode(1)
	QLCDNumber__Oct QLCDNumber__Mode = QLCDNumber__Mode(2)
	QLCDNumber__Bin QLCDNumber__Mode = QLCDNumber__Mode(3)
)

//go:generate stringer -type=QLCDNumber__SegmentStyle
//QLCDNumber::SegmentStyle
type QLCDNumber__SegmentStyle int64

var (
	QLCDNumber__Outline QLCDNumber__SegmentStyle = QLCDNumber__SegmentStyle(0)
	QLCDNumber__Filled  QLCDNumber__SegmentStyle = QLCDNumber__SegmentStyle(1)
	QLCDNumber__Flat    QLCDNumber__SegmentStyle = QLCDNumber__SegmentStyle(2)
)

type QLabel struct {
	QFrame
}

type QLabel_ITF interface {
	QFrame_ITF
	QLabel_PTR() *QLabel
}

func (ptr *QLabel) QLabel_PTR() *QLabel {
	return ptr
}

func (ptr *QLabel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QFrame_PTR().Pointer()
	}
	return nil
}

func (ptr *QLabel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QFrame_PTR().SetPointer(p)
	}
}

func PointerFromQLabel(ptr QLabel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLabel_PTR().Pointer()
	}
	return nil
}

func NewQLabelFromPointer(ptr unsafe.Pointer) (n *QLabel) {
	n = new(QLabel)
	n.SetPointer(ptr)
	return
}
func NewQLabel(parent QWidget_ITF, ff core.Qt__WindowType) *QLabel {
	tmpValue := NewQLabelFromPointer(C.QLabel_NewQLabel(PointerFromQWidget(parent), C.longlong(ff)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQLabel2(text string, parent QWidget_ITF, ff core.Qt__WindowType) *QLabel {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewQLabelFromPointer(C.QLabel_NewQLabel2(C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, PointerFromQWidget(parent), C.longlong(ff)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLabel) SetIndent(vin int) {
	if ptr.Pointer() != nil {
		C.QLabel_SetIndent(ptr.Pointer(), C.int(int32(vin)))
	}
}

//export callbackQLabel_SetText
func callbackQLabel_SetText(ptr unsafe.Pointer, vqs C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "setText"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	} else {
		NewQLabelFromPointer(ptr).SetTextDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QLabel) ConnectSetText(f func(vqs string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setText"); signal != nil {
			f := func(vqs string) {
				(*(*func(string))(signal))(vqs)
				f(vqs)
			}
			qt.ConnectSignal(ptr.Pointer(), "setText", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setText", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLabel) DisconnectSetText() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setText")
	}
}

func (ptr *QLabel) SetText(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QLabel_SetText(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

func (ptr *QLabel) SetTextDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QLabel_SetTextDefault(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

func (ptr *QLabel) DestroyQLabel() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLabel_DestroyQLabel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLabel) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLabel_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLabel) TextFormat() core.Qt__TextFormat {
	if ptr.Pointer() != nil {
		return core.Qt__TextFormat(C.QLabel_TextFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLabel) Indent() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLabel_Indent(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLabel) Margin() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLabel_Margin(ptr.Pointer())))
	}
	return 0
}

type QLayout struct {
	core.QObject
	QLayoutItem
}

type QLayout_ITF interface {
	core.QObject_ITF
	QLayoutItem_ITF
	QLayout_PTR() *QLayout
}

func (ptr *QLayout) QLayout_PTR() *QLayout {
	return ptr
}

func (ptr *QLayout) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QLayout) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QLayoutItem_PTR().SetPointer(p)
	}
}

func PointerFromQLayout(ptr QLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLayout_PTR().Pointer()
	}
	return nil
}

func NewQLayoutFromPointer(ptr unsafe.Pointer) (n *QLayout) {
	n = new(QLayout)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QLayout__SizeConstraint
//QLayout::SizeConstraint
type QLayout__SizeConstraint int64

const (
	QLayout__SetDefaultConstraint QLayout__SizeConstraint = QLayout__SizeConstraint(0)
	QLayout__SetNoConstraint      QLayout__SizeConstraint = QLayout__SizeConstraint(1)
	QLayout__SetMinimumSize       QLayout__SizeConstraint = QLayout__SizeConstraint(2)
	QLayout__SetFixedSize         QLayout__SizeConstraint = QLayout__SizeConstraint(3)
	QLayout__SetMaximumSize       QLayout__SizeConstraint = QLayout__SizeConstraint(4)
	QLayout__SetMinAndMaxSize     QLayout__SizeConstraint = QLayout__SizeConstraint(5)
)

func NewQLayout2() *QLayout {
	tmpValue := NewQLayoutFromPointer(C.QLayout_NewQLayout2())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQLayout(parent QWidget_ITF) *QLayout {
	tmpValue := NewQLayoutFromPointer(C.QLayout_NewQLayout(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQLayout_TakeAt
func callbackQLayout_TakeAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "takeAt"); signal != nil {
		return PointerFromQLayoutItem((*(*func(int) *QLayoutItem)(signal))(int(int32(index))))
	}

	return PointerFromQLayoutItem(NewQLayoutItem(0))
}

func (ptr *QLayout) ConnectTakeAt(f func(index int) *QLayoutItem) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "takeAt"); signal != nil {
			f := func(index int) *QLayoutItem {
				(*(*func(int) *QLayoutItem)(signal))(index)
				return f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "takeAt", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "takeAt", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayout) DisconnectTakeAt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "takeAt")
	}
}

func (ptr *QLayout) TakeAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QLayout_TakeAt(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

//export callbackQLayout_AddItem
func callbackQLayout_AddItem(ptr unsafe.Pointer, item unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addItem"); signal != nil {
		(*(*func(*QLayoutItem))(signal))(NewQLayoutItemFromPointer(item))
	}

}

func (ptr *QLayout) ConnectAddItem(f func(item *QLayoutItem)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addItem"); signal != nil {
			f := func(item *QLayoutItem) {
				(*(*func(*QLayoutItem))(signal))(item)
				f(item)
			}
			qt.ConnectSignal(ptr.Pointer(), "addItem", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addItem", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayout) DisconnectAddItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addItem")
	}
}

func (ptr *QLayout) AddItem(item QLayoutItem_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout_AddItem(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QLayout) AddWidget(w QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout_AddWidget(ptr.Pointer(), PointerFromQWidget(w))
	}
}

func (ptr *QLayout) SetContentsMargins2(margins core.QMargins_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout_SetContentsMargins2(ptr.Pointer(), core.PointerFromQMargins(margins))
	}
}

func (ptr *QLayout) SetContentsMargins(left int, top int, right int, bottom int) {
	if ptr.Pointer() != nil {
		C.QLayout_SetContentsMargins(ptr.Pointer(), C.int(int32(left)), C.int(int32(top)), C.int(int32(right)), C.int(int32(bottom)))
	}
}

func (ptr *QLayout) SetEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QLayout_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QLayout) SetSpacing(vin int) {
	if ptr.Pointer() != nil {
		C.QLayout_SetSpacing(ptr.Pointer(), C.int(int32(vin)))
	}
}

//export callbackQLayout_ItemAt
func callbackQLayout_ItemAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemAt"); signal != nil {
		return PointerFromQLayoutItem((*(*func(int) *QLayoutItem)(signal))(int(int32(index))))
	}

	return PointerFromQLayoutItem(NewQLayoutItem(0))
}

func (ptr *QLayout) ConnectItemAt(f func(index int) *QLayoutItem) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "itemAt"); signal != nil {
			f := func(index int) *QLayoutItem {
				(*(*func(int) *QLayoutItem)(signal))(index)
				return f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "itemAt", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "itemAt", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayout) DisconnectItemAt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "itemAt")
	}
}

func (ptr *QLayout) ItemAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QLayout_ItemAt(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

func (ptr *QLayout) ContentsMargins() *core.QMargins {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQMarginsFromPointer(C.QLayout_ContentsMargins(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QMargins).DestroyQMargins)
		return tmpValue
	}
	return nil
}

//export callbackQLayout_MinimumSize
func callbackQLayout_MinimumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSize"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQLayoutFromPointer(ptr).MinimumSizeDefault())
}

func (ptr *QLayout) ConnectMinimumSize(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "minimumSize"); signal != nil {
			f := func() *core.QSize {
				(*(*func() *core.QSize)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "minimumSize", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minimumSize", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayout) DisconnectMinimumSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "minimumSize")
	}
}

func (ptr *QLayout) MinimumSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QLayout_MinimumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QLayout) MinimumSizeDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QLayout_MinimumSizeDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQLayout_Count
func callbackQLayout_Count(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "count"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(0))
}

func (ptr *QLayout) ConnectCount(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "count"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "count", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "count", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayout) DisconnectCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "count")
	}
}

func (ptr *QLayout) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLayout_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLayout) Spacing() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLayout_Spacing(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLayout) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QLayout___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLayout) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QLayout) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QLayout___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QLayout) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLayout___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLayout) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLayout) __findChildren_newList2() unsafe.Pointer {
	return C.QLayout___findChildren_newList2(ptr.Pointer())
}

func (ptr *QLayout) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLayout___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLayout) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLayout) __findChildren_newList3() unsafe.Pointer {
	return C.QLayout___findChildren_newList3(ptr.Pointer())
}

func (ptr *QLayout) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLayout___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLayout) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLayout) __findChildren_newList() unsafe.Pointer {
	return C.QLayout___findChildren_newList(ptr.Pointer())
}

func (ptr *QLayout) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLayout___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLayout) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLayout) __children_newList() unsafe.Pointer {
	return C.QLayout___children_newList(ptr.Pointer())
}

//export callbackQLayout_ObjectNameChanged
func callbackQLayout_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQLayout_SetGeometry
func callbackQLayout_SetGeometry(ptr unsafe.Pointer, r unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGeometry"); signal != nil {
		(*(*func(*core.QRect))(signal))(core.NewQRectFromPointer(r))
	} else {
		NewQLayoutFromPointer(ptr).SetGeometryDefault(core.NewQRectFromPointer(r))
	}
}

func (ptr *QLayout) SetGeometry(r core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout_SetGeometry(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

func (ptr *QLayout) SetGeometryDefault(r core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout_SetGeometryDefault(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

//export callbackQLayout_Geometry
func callbackQLayout_Geometry(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "geometry"); signal != nil {
		return core.PointerFromQRect((*(*func() *core.QRect)(signal))())
	}

	return core.PointerFromQRect(NewQLayoutFromPointer(ptr).GeometryDefault())
}

func (ptr *QLayout) Geometry() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QLayout_Geometry(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QLayout) GeometryDefault() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QLayout_GeometryDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQLayout_MaximumSize
func callbackQLayout_MaximumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "maximumSize"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQLayoutFromPointer(ptr).MaximumSizeDefault())
}

func (ptr *QLayout) MaximumSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QLayout_MaximumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QLayout) MaximumSizeDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QLayout_MaximumSizeDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQLayout_SizeHint
func callbackQLayout_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQLayoutFromPointer(ptr).SizeHintDefault())
}

func (ptr *QLayout) SizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QLayout_SizeHint(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QLayout) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QLayout_SizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQLayout_ExpandingDirections
func callbackQLayout_ExpandingDirections(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "expandingDirections"); signal != nil {
		return C.longlong((*(*func() core.Qt__Orientation)(signal))())
	}

	return C.longlong(NewQLayoutFromPointer(ptr).ExpandingDirectionsDefault())
}

func (ptr *QLayout) ExpandingDirections() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QLayout_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayout) ExpandingDirectionsDefault() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QLayout_ExpandingDirectionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQLayout_IsEmpty
func callbackQLayout_IsEmpty(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isEmpty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLayoutFromPointer(ptr).IsEmptyDefault())))
}

func (ptr *QLayout) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLayout_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLayout) IsEmptyDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLayout_IsEmptyDefault(ptr.Pointer())) != 0
	}
	return false
}

type QLayoutItem struct {
	ptr unsafe.Pointer
}

type QLayoutItem_ITF interface {
	QLayoutItem_PTR() *QLayoutItem
}

func (ptr *QLayoutItem) QLayoutItem_PTR() *QLayoutItem {
	return ptr
}

func (ptr *QLayoutItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLayoutItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQLayoutItem(ptr QLayoutItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLayoutItem_PTR().Pointer()
	}
	return nil
}

func NewQLayoutItemFromPointer(ptr unsafe.Pointer) (n *QLayoutItem) {
	n = new(QLayoutItem)
	n.SetPointer(ptr)
	return
}

//export callbackQLayoutItem_Layout
func callbackQLayoutItem_Layout(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "layout"); signal != nil {
		return PointerFromQLayout((*(*func() *QLayout)(signal))())
	}

	return PointerFromQLayout(NewQLayoutItemFromPointer(ptr).LayoutDefault())
}

func (ptr *QLayoutItem) ConnectLayout(f func() *QLayout) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "layout"); signal != nil {
			f := func() *QLayout {
				(*(*func() *QLayout)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "layout", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "layout", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectLayout() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "layout")
	}
}

func (ptr *QLayoutItem) Layout() *QLayout {
	if ptr.Pointer() != nil {
		tmpValue := NewQLayoutFromPointer(C.QLayoutItem_Layout(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLayoutItem) LayoutDefault() *QLayout {
	if ptr.Pointer() != nil {
		tmpValue := NewQLayoutFromPointer(C.QLayoutItem_LayoutDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQLayoutItem(alignment core.Qt__AlignmentFlag) *QLayoutItem {
	return NewQLayoutItemFromPointer(C.QLayoutItem_NewQLayoutItem(C.longlong(alignment)))
}

//export callbackQLayoutItem_Widget
func callbackQLayoutItem_Widget(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "widget"); signal != nil {
		return PointerFromQWidget((*(*func() *QWidget)(signal))())
	}

	return PointerFromQWidget(NewQLayoutItemFromPointer(ptr).WidgetDefault())
}

func (ptr *QLayoutItem) ConnectWidget(f func() *QWidget) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "widget"); signal != nil {
			f := func() *QWidget {
				(*(*func() *QWidget)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "widget", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "widget", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "widget")
	}
}

func (ptr *QLayoutItem) Widget() *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QLayoutItem_Widget(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLayoutItem) WidgetDefault() *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QLayoutItem_WidgetDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQLayoutItem_SetGeometry
func callbackQLayoutItem_SetGeometry(ptr unsafe.Pointer, r unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGeometry"); signal != nil {
		(*(*func(*core.QRect))(signal))(core.NewQRectFromPointer(r))
	}

}

func (ptr *QLayoutItem) ConnectSetGeometry(f func(r *core.QRect)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setGeometry"); signal != nil {
			f := func(r *core.QRect) {
				(*(*func(*core.QRect))(signal))(r)
				f(r)
			}
			qt.ConnectSignal(ptr.Pointer(), "setGeometry", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setGeometry", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectSetGeometry() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setGeometry")
	}
}

func (ptr *QLayoutItem) SetGeometry(r core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QLayoutItem_SetGeometry(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

//export callbackQLayoutItem_DestroyQLayoutItem
func callbackQLayoutItem_DestroyQLayoutItem(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QLayoutItem"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLayoutItemFromPointer(ptr).DestroyQLayoutItemDefault()
	}
}

func (ptr *QLayoutItem) ConnectDestroyQLayoutItem(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QLayoutItem"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QLayoutItem", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QLayoutItem", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectDestroyQLayoutItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QLayoutItem")
	}
}

func (ptr *QLayoutItem) DestroyQLayoutItem() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLayoutItem_DestroyQLayoutItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLayoutItem) DestroyQLayoutItemDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLayoutItem_DestroyQLayoutItemDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQLayoutItem_Geometry
func callbackQLayoutItem_Geometry(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "geometry"); signal != nil {
		return core.PointerFromQRect((*(*func() *core.QRect)(signal))())
	}

	return core.PointerFromQRect(core.NewQRect())
}

func (ptr *QLayoutItem) ConnectGeometry(f func() *core.QRect) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "geometry"); signal != nil {
			f := func() *core.QRect {
				(*(*func() *core.QRect)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "geometry", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "geometry", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectGeometry() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "geometry")
	}
}

func (ptr *QLayoutItem) Geometry() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QLayoutItem_Geometry(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQLayoutItem_MaximumSize
func callbackQLayoutItem_MaximumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "maximumSize"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(core.NewQSize())
}

func (ptr *QLayoutItem) ConnectMaximumSize(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "maximumSize"); signal != nil {
			f := func() *core.QSize {
				(*(*func() *core.QSize)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "maximumSize", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "maximumSize", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectMaximumSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "maximumSize")
	}
}

func (ptr *QLayoutItem) MaximumSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QLayoutItem_MaximumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQLayoutItem_MinimumSize
func callbackQLayoutItem_MinimumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSize"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(core.NewQSize())
}

func (ptr *QLayoutItem) ConnectMinimumSize(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "minimumSize"); signal != nil {
			f := func() *core.QSize {
				(*(*func() *core.QSize)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "minimumSize", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minimumSize", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectMinimumSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "minimumSize")
	}
}

func (ptr *QLayoutItem) MinimumSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QLayoutItem_MinimumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQLayoutItem_SizeHint
func callbackQLayoutItem_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(core.NewQSize())
}

func (ptr *QLayoutItem) ConnectSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sizeHint"); signal != nil {
			f := func() *core.QSize {
				(*(*func() *core.QSize)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "sizeHint", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sizeHint", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sizeHint")
	}
}

func (ptr *QLayoutItem) SizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QLayoutItem_SizeHint(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQLayoutItem_ExpandingDirections
func callbackQLayoutItem_ExpandingDirections(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "expandingDirections"); signal != nil {
		return C.longlong((*(*func() core.Qt__Orientation)(signal))())
	}

	return C.longlong(0)
}

func (ptr *QLayoutItem) ConnectExpandingDirections(f func() core.Qt__Orientation) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "expandingDirections"); signal != nil {
			f := func() core.Qt__Orientation {
				(*(*func() core.Qt__Orientation)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "expandingDirections", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "expandingDirections", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectExpandingDirections() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "expandingDirections")
	}
}

func (ptr *QLayoutItem) ExpandingDirections() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QLayoutItem_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

//export callbackQLayoutItem_IsEmpty
func callbackQLayoutItem_IsEmpty(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isEmpty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QLayoutItem) ConnectIsEmpty(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isEmpty"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "isEmpty", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isEmpty", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectIsEmpty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isEmpty")
	}
}

func (ptr *QLayoutItem) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLayoutItem_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

//go:generate stringer -type=QLineEdit__ActionPosition
//QLineEdit::ActionPosition
type QLineEdit__ActionPosition int64

const (
	QLineEdit__LeadingPosition  QLineEdit__ActionPosition = QLineEdit__ActionPosition(0)
	QLineEdit__TrailingPosition QLineEdit__ActionPosition = QLineEdit__ActionPosition(1)
)

//go:generate stringer -type=QLineEdit__EchoMode
//QLineEdit::EchoMode
type QLineEdit__EchoMode int64

const (
	QLineEdit__Normal             QLineEdit__EchoMode = QLineEdit__EchoMode(0)
	QLineEdit__NoEcho             QLineEdit__EchoMode = QLineEdit__EchoMode(1)
	QLineEdit__Password           QLineEdit__EchoMode = QLineEdit__EchoMode(2)
	QLineEdit__PasswordEchoOnEdit QLineEdit__EchoMode = QLineEdit__EchoMode(3)
)

//go:generate stringer -type=QListView__Flow
//QListView::Flow
type QListView__Flow int64

const (
	QListView__LeftToRight QListView__Flow = QListView__Flow(0)
	QListView__TopToBottom QListView__Flow = QListView__Flow(1)
)

//go:generate stringer -type=QListView__LayoutMode
//QListView::LayoutMode
type QListView__LayoutMode int64

const (
	QListView__SinglePass QListView__LayoutMode = QListView__LayoutMode(0)
	QListView__Batched    QListView__LayoutMode = QListView__LayoutMode(1)
)

//go:generate stringer -type=QListView__Movement
//QListView::Movement
type QListView__Movement int64

const (
	QListView__Static QListView__Movement = QListView__Movement(0)
	QListView__Free   QListView__Movement = QListView__Movement(1)
	QListView__Snap   QListView__Movement = QListView__Movement(2)
)

//go:generate stringer -type=QListView__ResizeMode
//QListView::ResizeMode
type QListView__ResizeMode int64

const (
	QListView__Fixed  QListView__ResizeMode = QListView__ResizeMode(0)
	QListView__Adjust QListView__ResizeMode = QListView__ResizeMode(1)
)

//go:generate stringer -type=QListView__ViewMode
//QListView::ViewMode
type QListView__ViewMode int64

const (
	QListView__ListMode QListView__ViewMode = QListView__ViewMode(0)
	QListView__IconMode QListView__ViewMode = QListView__ViewMode(1)
)

//go:generate stringer -type=QListWidgetItem__ItemType
//QListWidgetItem::ItemType
type QListWidgetItem__ItemType int64

const (
	QListWidgetItem__Type     QListWidgetItem__ItemType = QListWidgetItem__ItemType(0)
	QListWidgetItem__UserType QListWidgetItem__ItemType = QListWidgetItem__ItemType(1000)
)

//go:generate stringer -type=QMainWindow__DockOption
//QMainWindow::DockOption
type QMainWindow__DockOption int64

const (
	QMainWindow__AnimatedDocks    QMainWindow__DockOption = QMainWindow__DockOption(0x01)
	QMainWindow__AllowNestedDocks QMainWindow__DockOption = QMainWindow__DockOption(0x02)
	QMainWindow__AllowTabbedDocks QMainWindow__DockOption = QMainWindow__DockOption(0x04)
	QMainWindow__ForceTabbedDocks QMainWindow__DockOption = QMainWindow__DockOption(0x08)
	QMainWindow__VerticalTabs     QMainWindow__DockOption = QMainWindow__DockOption(0x10)
	QMainWindow__GroupedDragging  QMainWindow__DockOption = QMainWindow__DockOption(0x20)
)

//go:generate stringer -type=QMdiArea__AreaOption
//QMdiArea::AreaOption
type QMdiArea__AreaOption int64

const (
	QMdiArea__DontMaximizeSubWindowOnActivation QMdiArea__AreaOption = QMdiArea__AreaOption(0x1)
)

//go:generate stringer -type=QMdiArea__ViewMode
//QMdiArea::ViewMode
type QMdiArea__ViewMode int64

const (
	QMdiArea__SubWindowView QMdiArea__ViewMode = QMdiArea__ViewMode(0)
	QMdiArea__TabbedView    QMdiArea__ViewMode = QMdiArea__ViewMode(1)
)

//go:generate stringer -type=QMdiArea__WindowOrder
//QMdiArea::WindowOrder
type QMdiArea__WindowOrder int64

const (
	QMdiArea__CreationOrder          QMdiArea__WindowOrder = QMdiArea__WindowOrder(0)
	QMdiArea__StackingOrder          QMdiArea__WindowOrder = QMdiArea__WindowOrder(1)
	QMdiArea__ActivationHistoryOrder QMdiArea__WindowOrder = QMdiArea__WindowOrder(2)
)

//go:generate stringer -type=QMdiSubWindow__SubWindowOption
//QMdiSubWindow::SubWindowOption
type QMdiSubWindow__SubWindowOption int64

const (
	QMdiSubWindow__AllowOutsideAreaHorizontally QMdiSubWindow__SubWindowOption = QMdiSubWindow__SubWindowOption(0x1)
	QMdiSubWindow__AllowOutsideAreaVertically   QMdiSubWindow__SubWindowOption = QMdiSubWindow__SubWindowOption(0x2)
	QMdiSubWindow__RubberBandResize             QMdiSubWindow__SubWindowOption = QMdiSubWindow__SubWindowOption(0x4)
	QMdiSubWindow__RubberBandMove               QMdiSubWindow__SubWindowOption = QMdiSubWindow__SubWindowOption(0x8)
)

//go:generate stringer -type=QMessageBox__ButtonRole
//QMessageBox::ButtonRole
type QMessageBox__ButtonRole int64

const (
	QMessageBox__InvalidRole     QMessageBox__ButtonRole = QMessageBox__ButtonRole(-1)
	QMessageBox__AcceptRole      QMessageBox__ButtonRole = QMessageBox__ButtonRole(0)
	QMessageBox__RejectRole      QMessageBox__ButtonRole = QMessageBox__ButtonRole(1)
	QMessageBox__DestructiveRole QMessageBox__ButtonRole = QMessageBox__ButtonRole(2)
	QMessageBox__ActionRole      QMessageBox__ButtonRole = QMessageBox__ButtonRole(3)
	QMessageBox__HelpRole        QMessageBox__ButtonRole = QMessageBox__ButtonRole(4)
	QMessageBox__YesRole         QMessageBox__ButtonRole = QMessageBox__ButtonRole(5)
	QMessageBox__NoRole          QMessageBox__ButtonRole = QMessageBox__ButtonRole(6)
	QMessageBox__ResetRole       QMessageBox__ButtonRole = QMessageBox__ButtonRole(7)
	QMessageBox__ApplyRole       QMessageBox__ButtonRole = QMessageBox__ButtonRole(8)
	QMessageBox__NRoles          QMessageBox__ButtonRole = QMessageBox__ButtonRole(9)
)

//go:generate stringer -type=QMessageBox__Icon
//QMessageBox::Icon
type QMessageBox__Icon int64

const (
	QMessageBox__NoIcon      QMessageBox__Icon = QMessageBox__Icon(0)
	QMessageBox__Information QMessageBox__Icon = QMessageBox__Icon(1)
	QMessageBox__Warning     QMessageBox__Icon = QMessageBox__Icon(2)
	QMessageBox__Critical    QMessageBox__Icon = QMessageBox__Icon(3)
	QMessageBox__Question    QMessageBox__Icon = QMessageBox__Icon(4)
)

//go:generate stringer -type=QMessageBox__StandardButton
//QMessageBox::StandardButton
type QMessageBox__StandardButton int64

var (
	QMessageBox__NoButton        QMessageBox__StandardButton = QMessageBox__StandardButton(0x00000000)
	QMessageBox__Ok              QMessageBox__StandardButton = QMessageBox__StandardButton(0x00000400)
	QMessageBox__Save            QMessageBox__StandardButton = QMessageBox__StandardButton(0x00000800)
	QMessageBox__SaveAll         QMessageBox__StandardButton = QMessageBox__StandardButton(0x00001000)
	QMessageBox__Open            QMessageBox__StandardButton = QMessageBox__StandardButton(0x00002000)
	QMessageBox__Yes             QMessageBox__StandardButton = QMessageBox__StandardButton(0x00004000)
	QMessageBox__YesToAll        QMessageBox__StandardButton = QMessageBox__StandardButton(0x00008000)
	QMessageBox__No              QMessageBox__StandardButton = QMessageBox__StandardButton(0x00010000)
	QMessageBox__NoToAll         QMessageBox__StandardButton = QMessageBox__StandardButton(0x00020000)
	QMessageBox__Abort           QMessageBox__StandardButton = QMessageBox__StandardButton(0x00040000)
	QMessageBox__Retry           QMessageBox__StandardButton = QMessageBox__StandardButton(0x00080000)
	QMessageBox__Ignore          QMessageBox__StandardButton = QMessageBox__StandardButton(0x00100000)
	QMessageBox__Close           QMessageBox__StandardButton = QMessageBox__StandardButton(0x00200000)
	QMessageBox__Cancel          QMessageBox__StandardButton = QMessageBox__StandardButton(0x00400000)
	QMessageBox__Discard         QMessageBox__StandardButton = QMessageBox__StandardButton(0x00800000)
	QMessageBox__Help            QMessageBox__StandardButton = QMessageBox__StandardButton(0x01000000)
	QMessageBox__Apply           QMessageBox__StandardButton = QMessageBox__StandardButton(0x02000000)
	QMessageBox__Reset           QMessageBox__StandardButton = QMessageBox__StandardButton(0x04000000)
	QMessageBox__RestoreDefaults QMessageBox__StandardButton = QMessageBox__StandardButton(0x08000000)
	QMessageBox__FirstButton     QMessageBox__StandardButton = QMessageBox__StandardButton(QMessageBox__Ok)
	QMessageBox__LastButton      QMessageBox__StandardButton = QMessageBox__StandardButton(QMessageBox__RestoreDefaults)
	QMessageBox__YesAll          QMessageBox__StandardButton = QMessageBox__StandardButton(QMessageBox__YesToAll)
	QMessageBox__NoAll           QMessageBox__StandardButton = QMessageBox__StandardButton(QMessageBox__NoToAll)
	QMessageBox__Default         QMessageBox__StandardButton = QMessageBox__StandardButton(0x00000100)
	QMessageBox__Escape          QMessageBox__StandardButton = QMessageBox__StandardButton(0x00000200)
	QMessageBox__FlagMask        QMessageBox__StandardButton = QMessageBox__StandardButton(0x00000300)
	QMessageBox__ButtonMask      QMessageBox__StandardButton = QMessageBox__StandardButton(C.QMessageBox_ButtonMask_Type())
)

//go:generate stringer -type=QOpenGLWidget__UpdateBehavior
//QOpenGLWidget::UpdateBehavior
type QOpenGLWidget__UpdateBehavior int64

const (
	QOpenGLWidget__NoPartialUpdate QOpenGLWidget__UpdateBehavior = QOpenGLWidget__UpdateBehavior(0)
	QOpenGLWidget__PartialUpdate   QOpenGLWidget__UpdateBehavior = QOpenGLWidget__UpdateBehavior(1)
)

//go:generate stringer -type=QPinchGesture__ChangeFlag
//QPinchGesture::ChangeFlag
type QPinchGesture__ChangeFlag int64

const (
	QPinchGesture__ScaleFactorChanged   QPinchGesture__ChangeFlag = QPinchGesture__ChangeFlag(0x1)
	QPinchGesture__RotationAngleChanged QPinchGesture__ChangeFlag = QPinchGesture__ChangeFlag(0x2)
	QPinchGesture__CenterPointChanged   QPinchGesture__ChangeFlag = QPinchGesture__ChangeFlag(0x4)
)

//go:generate stringer -type=QPlainTextEdit__LineWrapMode
//QPlainTextEdit::LineWrapMode
type QPlainTextEdit__LineWrapMode int64

const (
	QPlainTextEdit__NoWrap      QPlainTextEdit__LineWrapMode = QPlainTextEdit__LineWrapMode(0)
	QPlainTextEdit__WidgetWidth QPlainTextEdit__LineWrapMode = QPlainTextEdit__LineWrapMode(1)
)

//go:generate stringer -type=QProgressBar__Direction
//QProgressBar::Direction
type QProgressBar__Direction int64

const (
	QProgressBar__TopToBottom QProgressBar__Direction = QProgressBar__Direction(0)
	QProgressBar__BottomToTop QProgressBar__Direction = QProgressBar__Direction(1)
)

type QPushButton struct {
	QAbstractButton
}

type QPushButton_ITF interface {
	QAbstractButton_ITF
	QPushButton_PTR() *QPushButton
}

func (ptr *QPushButton) QPushButton_PTR() *QPushButton {
	return ptr
}

func (ptr *QPushButton) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractButton_PTR().Pointer()
	}
	return nil
}

func (ptr *QPushButton) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractButton_PTR().SetPointer(p)
	}
}

func PointerFromQPushButton(ptr QPushButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPushButton_PTR().Pointer()
	}
	return nil
}

func NewQPushButtonFromPointer(ptr unsafe.Pointer) (n *QPushButton) {
	n = new(QPushButton)
	n.SetPointer(ptr)
	return
}
func NewQPushButton(parent QWidget_ITF) *QPushButton {
	tmpValue := NewQPushButtonFromPointer(C.QPushButton_NewQPushButton(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQPushButton3(icon gui.QIcon_ITF, text string, parent QWidget_ITF) *QPushButton {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewQPushButtonFromPointer(C.QPushButton_NewQPushButton3(gui.PointerFromQIcon(icon), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQPushButton2(text string, parent QWidget_ITF) *QPushButton {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewQPushButtonFromPointer(C.QPushButton_NewQPushButton2(C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QPushButton) DestroyQPushButton() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPushButton_DestroyQPushButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQPushButton_PaintEvent
func callbackQPushButton_PaintEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*gui.QPaintEvent))(signal))(gui.NewQPaintEventFromPointer(e))
	} else {
		NewQPushButtonFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *QPushButton) PaintEvent(e gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPushButton_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QPushButton) PaintEventDefault(e gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPushButton_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

//go:generate stringer -type=QRubberBand__Shape
//QRubberBand::Shape
type QRubberBand__Shape int64

const (
	QRubberBand__Line      QRubberBand__Shape = QRubberBand__Shape(0)
	QRubberBand__Rectangle QRubberBand__Shape = QRubberBand__Shape(1)
)

//go:generate stringer -type=QScroller__Input
//QScroller::Input
type QScroller__Input int64

const (
	QScroller__InputPress   QScroller__Input = QScroller__Input(1)
	QScroller__InputMove    QScroller__Input = QScroller__Input(2)
	QScroller__InputRelease QScroller__Input = QScroller__Input(3)
)

//go:generate stringer -type=QScroller__ScrollerGestureType
//QScroller::ScrollerGestureType
type QScroller__ScrollerGestureType int64

const (
	QScroller__TouchGesture             QScroller__ScrollerGestureType = QScroller__ScrollerGestureType(0)
	QScroller__LeftMouseButtonGesture   QScroller__ScrollerGestureType = QScroller__ScrollerGestureType(1)
	QScroller__RightMouseButtonGesture  QScroller__ScrollerGestureType = QScroller__ScrollerGestureType(2)
	QScroller__MiddleMouseButtonGesture QScroller__ScrollerGestureType = QScroller__ScrollerGestureType(3)
)

//go:generate stringer -type=QScroller__State
//QScroller::State
type QScroller__State int64

const (
	QScroller__Inactive  QScroller__State = QScroller__State(0)
	QScroller__Pressed   QScroller__State = QScroller__State(1)
	QScroller__Dragging  QScroller__State = QScroller__State(2)
	QScroller__Scrolling QScroller__State = QScroller__State(3)
)

//go:generate stringer -type=QScrollerProperties__FrameRates
//QScrollerProperties::FrameRates
type QScrollerProperties__FrameRates int64

const (
	QScrollerProperties__Standard QScrollerProperties__FrameRates = QScrollerProperties__FrameRates(0)
	QScrollerProperties__Fps60    QScrollerProperties__FrameRates = QScrollerProperties__FrameRates(1)
	QScrollerProperties__Fps30    QScrollerProperties__FrameRates = QScrollerProperties__FrameRates(2)
	QScrollerProperties__Fps20    QScrollerProperties__FrameRates = QScrollerProperties__FrameRates(3)
)

//go:generate stringer -type=QScrollerProperties__OvershootPolicy
//QScrollerProperties::OvershootPolicy
type QScrollerProperties__OvershootPolicy int64

const (
	QScrollerProperties__OvershootWhenScrollable QScrollerProperties__OvershootPolicy = QScrollerProperties__OvershootPolicy(0)
	QScrollerProperties__OvershootAlwaysOff      QScrollerProperties__OvershootPolicy = QScrollerProperties__OvershootPolicy(1)
	QScrollerProperties__OvershootAlwaysOn       QScrollerProperties__OvershootPolicy = QScrollerProperties__OvershootPolicy(2)
)

//go:generate stringer -type=QScrollerProperties__ScrollMetric
//QScrollerProperties::ScrollMetric
type QScrollerProperties__ScrollMetric int64

const (
	QScrollerProperties__MousePressEventDelay           QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(0)
	QScrollerProperties__DragStartDistance              QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(1)
	QScrollerProperties__DragVelocitySmoothingFactor    QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(2)
	QScrollerProperties__AxisLockThreshold              QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(3)
	QScrollerProperties__ScrollingCurve                 QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(4)
	QScrollerProperties__DecelerationFactor             QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(5)
	QScrollerProperties__MinimumVelocity                QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(6)
	QScrollerProperties__MaximumVelocity                QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(7)
	QScrollerProperties__MaximumClickThroughVelocity    QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(8)
	QScrollerProperties__AcceleratingFlickMaximumTime   QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(9)
	QScrollerProperties__AcceleratingFlickSpeedupFactor QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(10)
	QScrollerProperties__SnapPositionRatio              QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(11)
	QScrollerProperties__SnapTime                       QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(12)
	QScrollerProperties__OvershootDragResistanceFactor  QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(13)
	QScrollerProperties__OvershootDragDistanceFactor    QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(14)
	QScrollerProperties__OvershootScrollDistanceFactor  QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(15)
	QScrollerProperties__OvershootScrollTime            QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(16)
	QScrollerProperties__HorizontalOvershootPolicy      QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(17)
	QScrollerProperties__VerticalOvershootPolicy        QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(18)
	QScrollerProperties__FrameRate                      QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(19)
	QScrollerProperties__ScrollMetricCount              QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(20)
)

//go:generate stringer -type=QSizePolicy__ControlType
//QSizePolicy::ControlType
type QSizePolicy__ControlType int64

const (
	QSizePolicy__DefaultType QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000001)
	QSizePolicy__ButtonBox   QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000002)
	QSizePolicy__CheckBox    QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000004)
	QSizePolicy__ComboBox    QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000008)
	QSizePolicy__Frame       QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000010)
	QSizePolicy__GroupBox    QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000020)
	QSizePolicy__Label       QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000040)
	QSizePolicy__Line        QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000080)
	QSizePolicy__LineEdit    QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000100)
	QSizePolicy__PushButton  QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000200)
	QSizePolicy__RadioButton QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000400)
	QSizePolicy__Slider      QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000800)
	QSizePolicy__SpinBox     QSizePolicy__ControlType = QSizePolicy__ControlType(0x00001000)
	QSizePolicy__TabWidget   QSizePolicy__ControlType = QSizePolicy__ControlType(0x00002000)
	QSizePolicy__ToolButton  QSizePolicy__ControlType = QSizePolicy__ControlType(0x00004000)
)

//go:generate stringer -type=QSizePolicy__Policy
//QSizePolicy::Policy
type QSizePolicy__Policy int64

const (
	QSizePolicy__Fixed            QSizePolicy__Policy = QSizePolicy__Policy(0)
	QSizePolicy__Minimum          QSizePolicy__Policy = QSizePolicy__Policy(QSizePolicy__GrowFlag)
	QSizePolicy__Maximum          QSizePolicy__Policy = QSizePolicy__Policy(QSizePolicy__ShrinkFlag)
	QSizePolicy__Preferred        QSizePolicy__Policy = QSizePolicy__Policy(QSizePolicy__GrowFlag | QSizePolicy__ShrinkFlag)
	QSizePolicy__MinimumExpanding QSizePolicy__Policy = QSizePolicy__Policy(QSizePolicy__GrowFlag | QSizePolicy__ExpandFlag)
	QSizePolicy__Expanding        QSizePolicy__Policy = QSizePolicy__Policy(QSizePolicy__GrowFlag | QSizePolicy__ShrinkFlag | QSizePolicy__ExpandFlag)
	QSizePolicy__Ignored          QSizePolicy__Policy = QSizePolicy__Policy(QSizePolicy__ShrinkFlag | QSizePolicy__GrowFlag | QSizePolicy__IgnoreFlag)
)

//go:generate stringer -type=QSizePolicy__PolicyFlag
//QSizePolicy::PolicyFlag
type QSizePolicy__PolicyFlag int64

const (
	QSizePolicy__GrowFlag   QSizePolicy__PolicyFlag = QSizePolicy__PolicyFlag(1)
	QSizePolicy__ExpandFlag QSizePolicy__PolicyFlag = QSizePolicy__PolicyFlag(2)
	QSizePolicy__ShrinkFlag QSizePolicy__PolicyFlag = QSizePolicy__PolicyFlag(4)
	QSizePolicy__IgnoreFlag QSizePolicy__PolicyFlag = QSizePolicy__PolicyFlag(8)
)

//go:generate stringer -type=QSlider__TickPosition
//QSlider::TickPosition
type QSlider__TickPosition int64

const (
	QSlider__NoTicks        QSlider__TickPosition = QSlider__TickPosition(0)
	QSlider__TicksAbove     QSlider__TickPosition = QSlider__TickPosition(1)
	QSlider__TicksLeft      QSlider__TickPosition = QSlider__TickPosition(QSlider__TicksAbove)
	QSlider__TicksBelow     QSlider__TickPosition = QSlider__TickPosition(2)
	QSlider__TicksRight     QSlider__TickPosition = QSlider__TickPosition(QSlider__TicksBelow)
	QSlider__TicksBothSides QSlider__TickPosition = QSlider__TickPosition(3)
)

//go:generate stringer -type=QStackedLayout__StackingMode
//QStackedLayout::StackingMode
type QStackedLayout__StackingMode int64

const (
	QStackedLayout__StackOne QStackedLayout__StackingMode = QStackedLayout__StackingMode(0)
	QStackedLayout__StackAll QStackedLayout__StackingMode = QStackedLayout__StackingMode(1)
)

//go:generate stringer -type=QStyle__ComplexControl
//QStyle::ComplexControl
type QStyle__ComplexControl int64

const (
	QStyle__CC_SpinBox     QStyle__ComplexControl = QStyle__ComplexControl(0)
	QStyle__CC_ComboBox    QStyle__ComplexControl = QStyle__ComplexControl(1)
	QStyle__CC_ScrollBar   QStyle__ComplexControl = QStyle__ComplexControl(2)
	QStyle__CC_Slider      QStyle__ComplexControl = QStyle__ComplexControl(3)
	QStyle__CC_ToolButton  QStyle__ComplexControl = QStyle__ComplexControl(4)
	QStyle__CC_TitleBar    QStyle__ComplexControl = QStyle__ComplexControl(5)
	QStyle__CC_Dial        QStyle__ComplexControl = QStyle__ComplexControl(6)
	QStyle__CC_GroupBox    QStyle__ComplexControl = QStyle__ComplexControl(7)
	QStyle__CC_MdiControls QStyle__ComplexControl = QStyle__ComplexControl(8)
	QStyle__CC_CustomBase  QStyle__ComplexControl = QStyle__ComplexControl(0xf0000000)
)

//go:generate stringer -type=QStyle__ContentsType
//QStyle::ContentsType
type QStyle__ContentsType int64

const (
	QStyle__CT_PushButton    QStyle__ContentsType = QStyle__ContentsType(0)
	QStyle__CT_CheckBox      QStyle__ContentsType = QStyle__ContentsType(1)
	QStyle__CT_RadioButton   QStyle__ContentsType = QStyle__ContentsType(2)
	QStyle__CT_ToolButton    QStyle__ContentsType = QStyle__ContentsType(3)
	QStyle__CT_ComboBox      QStyle__ContentsType = QStyle__ContentsType(4)
	QStyle__CT_Splitter      QStyle__ContentsType = QStyle__ContentsType(5)
	QStyle__CT_ProgressBar   QStyle__ContentsType = QStyle__ContentsType(6)
	QStyle__CT_MenuItem      QStyle__ContentsType = QStyle__ContentsType(7)
	QStyle__CT_MenuBarItem   QStyle__ContentsType = QStyle__ContentsType(8)
	QStyle__CT_MenuBar       QStyle__ContentsType = QStyle__ContentsType(9)
	QStyle__CT_Menu          QStyle__ContentsType = QStyle__ContentsType(10)
	QStyle__CT_TabBarTab     QStyle__ContentsType = QStyle__ContentsType(11)
	QStyle__CT_Slider        QStyle__ContentsType = QStyle__ContentsType(12)
	QStyle__CT_ScrollBar     QStyle__ContentsType = QStyle__ContentsType(13)
	QStyle__CT_LineEdit      QStyle__ContentsType = QStyle__ContentsType(14)
	QStyle__CT_SpinBox       QStyle__ContentsType = QStyle__ContentsType(15)
	QStyle__CT_SizeGrip      QStyle__ContentsType = QStyle__ContentsType(16)
	QStyle__CT_TabWidget     QStyle__ContentsType = QStyle__ContentsType(17)
	QStyle__CT_DialogButtons QStyle__ContentsType = QStyle__ContentsType(18)
	QStyle__CT_HeaderSection QStyle__ContentsType = QStyle__ContentsType(19)
	QStyle__CT_GroupBox      QStyle__ContentsType = QStyle__ContentsType(20)
	QStyle__CT_MdiControls   QStyle__ContentsType = QStyle__ContentsType(21)
	QStyle__CT_ItemViewItem  QStyle__ContentsType = QStyle__ContentsType(22)
	QStyle__CT_CustomBase    QStyle__ContentsType = QStyle__ContentsType(0xf0000000)
)

//go:generate stringer -type=QStyle__ControlElement
//QStyle::ControlElement
type QStyle__ControlElement int64

const (
	QStyle__CE_PushButton          QStyle__ControlElement = QStyle__ControlElement(0)
	QStyle__CE_PushButtonBevel     QStyle__ControlElement = QStyle__ControlElement(1)
	QStyle__CE_PushButtonLabel     QStyle__ControlElement = QStyle__ControlElement(2)
	QStyle__CE_CheckBox            QStyle__ControlElement = QStyle__ControlElement(3)
	QStyle__CE_CheckBoxLabel       QStyle__ControlElement = QStyle__ControlElement(4)
	QStyle__CE_RadioButton         QStyle__ControlElement = QStyle__ControlElement(5)
	QStyle__CE_RadioButtonLabel    QStyle__ControlElement = QStyle__ControlElement(6)
	QStyle__CE_TabBarTab           QStyle__ControlElement = QStyle__ControlElement(7)
	QStyle__CE_TabBarTabShape      QStyle__ControlElement = QStyle__ControlElement(8)
	QStyle__CE_TabBarTabLabel      QStyle__ControlElement = QStyle__ControlElement(9)
	QStyle__CE_ProgressBar         QStyle__ControlElement = QStyle__ControlElement(10)
	QStyle__CE_ProgressBarGroove   QStyle__ControlElement = QStyle__ControlElement(11)
	QStyle__CE_ProgressBarContents QStyle__ControlElement = QStyle__ControlElement(12)
	QStyle__CE_ProgressBarLabel    QStyle__ControlElement = QStyle__ControlElement(13)
	QStyle__CE_MenuItem            QStyle__ControlElement = QStyle__ControlElement(14)
	QStyle__CE_MenuScroller        QStyle__ControlElement = QStyle__ControlElement(15)
	QStyle__CE_MenuVMargin         QStyle__ControlElement = QStyle__ControlElement(16)
	QStyle__CE_MenuHMargin         QStyle__ControlElement = QStyle__ControlElement(17)
	QStyle__CE_MenuTearoff         QStyle__ControlElement = QStyle__ControlElement(18)
	QStyle__CE_MenuEmptyArea       QStyle__ControlElement = QStyle__ControlElement(19)
	QStyle__CE_MenuBarItem         QStyle__ControlElement = QStyle__ControlElement(20)
	QStyle__CE_MenuBarEmptyArea    QStyle__ControlElement = QStyle__ControlElement(21)
	QStyle__CE_ToolButtonLabel     QStyle__ControlElement = QStyle__ControlElement(22)
	QStyle__CE_Header              QStyle__ControlElement = QStyle__ControlElement(23)
	QStyle__CE_HeaderSection       QStyle__ControlElement = QStyle__ControlElement(24)
	QStyle__CE_HeaderLabel         QStyle__ControlElement = QStyle__ControlElement(25)
	QStyle__CE_ToolBoxTab          QStyle__ControlElement = QStyle__ControlElement(26)
	QStyle__CE_SizeGrip            QStyle__ControlElement = QStyle__ControlElement(27)
	QStyle__CE_Splitter            QStyle__ControlElement = QStyle__ControlElement(28)
	QStyle__CE_RubberBand          QStyle__ControlElement = QStyle__ControlElement(29)
	QStyle__CE_DockWidgetTitle     QStyle__ControlElement = QStyle__ControlElement(30)
	QStyle__CE_ScrollBarAddLine    QStyle__ControlElement = QStyle__ControlElement(31)
	QStyle__CE_ScrollBarSubLine    QStyle__ControlElement = QStyle__ControlElement(32)
	QStyle__CE_ScrollBarAddPage    QStyle__ControlElement = QStyle__ControlElement(33)
	QStyle__CE_ScrollBarSubPage    QStyle__ControlElement = QStyle__ControlElement(34)
	QStyle__CE_ScrollBarSlider     QStyle__ControlElement = QStyle__ControlElement(35)
	QStyle__CE_ScrollBarFirst      QStyle__ControlElement = QStyle__ControlElement(36)
	QStyle__CE_ScrollBarLast       QStyle__ControlElement = QStyle__ControlElement(37)
	QStyle__CE_FocusFrame          QStyle__ControlElement = QStyle__ControlElement(38)
	QStyle__CE_ComboBoxLabel       QStyle__ControlElement = QStyle__ControlElement(39)
	QStyle__CE_ToolBar             QStyle__ControlElement = QStyle__ControlElement(40)
	QStyle__CE_ToolBoxTabShape     QStyle__ControlElement = QStyle__ControlElement(41)
	QStyle__CE_ToolBoxTabLabel     QStyle__ControlElement = QStyle__ControlElement(42)
	QStyle__CE_HeaderEmptyArea     QStyle__ControlElement = QStyle__ControlElement(43)
	QStyle__CE_ColumnViewGrip      QStyle__ControlElement = QStyle__ControlElement(44)
	QStyle__CE_ItemViewItem        QStyle__ControlElement = QStyle__ControlElement(45)
	QStyle__CE_ShapedFrame         QStyle__ControlElement = QStyle__ControlElement(46)
	QStyle__CE_CustomBase          QStyle__ControlElement = QStyle__ControlElement(0xf0000000)
)

//go:generate stringer -type=QStyle__PixelMetric
//QStyle::PixelMetric
type QStyle__PixelMetric int64

var (
	QStyle__PM_ButtonMargin                       QStyle__PixelMetric = QStyle__PixelMetric(0)
	QStyle__PM_ButtonDefaultIndicator             QStyle__PixelMetric = QStyle__PixelMetric(1)
	QStyle__PM_MenuButtonIndicator                QStyle__PixelMetric = QStyle__PixelMetric(2)
	QStyle__PM_ButtonShiftHorizontal              QStyle__PixelMetric = QStyle__PixelMetric(3)
	QStyle__PM_ButtonShiftVertical                QStyle__PixelMetric = QStyle__PixelMetric(4)
	QStyle__PM_DefaultFrameWidth                  QStyle__PixelMetric = QStyle__PixelMetric(5)
	QStyle__PM_SpinBoxFrameWidth                  QStyle__PixelMetric = QStyle__PixelMetric(6)
	QStyle__PM_ComboBoxFrameWidth                 QStyle__PixelMetric = QStyle__PixelMetric(7)
	QStyle__PM_MaximumDragDistance                QStyle__PixelMetric = QStyle__PixelMetric(8)
	QStyle__PM_ScrollBarExtent                    QStyle__PixelMetric = QStyle__PixelMetric(9)
	QStyle__PM_ScrollBarSliderMin                 QStyle__PixelMetric = QStyle__PixelMetric(10)
	QStyle__PM_SliderThickness                    QStyle__PixelMetric = QStyle__PixelMetric(11)
	QStyle__PM_SliderControlThickness             QStyle__PixelMetric = QStyle__PixelMetric(12)
	QStyle__PM_SliderLength                       QStyle__PixelMetric = QStyle__PixelMetric(13)
	QStyle__PM_SliderTickmarkOffset               QStyle__PixelMetric = QStyle__PixelMetric(14)
	QStyle__PM_SliderSpaceAvailable               QStyle__PixelMetric = QStyle__PixelMetric(15)
	QStyle__PM_DockWidgetSeparatorExtent          QStyle__PixelMetric = QStyle__PixelMetric(16)
	QStyle__PM_DockWidgetHandleExtent             QStyle__PixelMetric = QStyle__PixelMetric(17)
	QStyle__PM_DockWidgetFrameWidth               QStyle__PixelMetric = QStyle__PixelMetric(18)
	QStyle__PM_TabBarTabOverlap                   QStyle__PixelMetric = QStyle__PixelMetric(19)
	QStyle__PM_TabBarTabHSpace                    QStyle__PixelMetric = QStyle__PixelMetric(20)
	QStyle__PM_TabBarTabVSpace                    QStyle__PixelMetric = QStyle__PixelMetric(21)
	QStyle__PM_TabBarBaseHeight                   QStyle__PixelMetric = QStyle__PixelMetric(22)
	QStyle__PM_TabBarBaseOverlap                  QStyle__PixelMetric = QStyle__PixelMetric(23)
	QStyle__PM_ProgressBarChunkWidth              QStyle__PixelMetric = QStyle__PixelMetric(24)
	QStyle__PM_SplitterWidth                      QStyle__PixelMetric = QStyle__PixelMetric(25)
	QStyle__PM_TitleBarHeight                     QStyle__PixelMetric = QStyle__PixelMetric(26)
	QStyle__PM_MenuScrollerHeight                 QStyle__PixelMetric = QStyle__PixelMetric(27)
	QStyle__PM_MenuHMargin                        QStyle__PixelMetric = QStyle__PixelMetric(28)
	QStyle__PM_MenuVMargin                        QStyle__PixelMetric = QStyle__PixelMetric(29)
	QStyle__PM_MenuPanelWidth                     QStyle__PixelMetric = QStyle__PixelMetric(30)
	QStyle__PM_MenuTearoffHeight                  QStyle__PixelMetric = QStyle__PixelMetric(31)
	QStyle__PM_MenuDesktopFrameWidth              QStyle__PixelMetric = QStyle__PixelMetric(32)
	QStyle__PM_MenuBarPanelWidth                  QStyle__PixelMetric = QStyle__PixelMetric(33)
	QStyle__PM_MenuBarItemSpacing                 QStyle__PixelMetric = QStyle__PixelMetric(34)
	QStyle__PM_MenuBarVMargin                     QStyle__PixelMetric = QStyle__PixelMetric(35)
	QStyle__PM_MenuBarHMargin                     QStyle__PixelMetric = QStyle__PixelMetric(36)
	QStyle__PM_IndicatorWidth                     QStyle__PixelMetric = QStyle__PixelMetric(37)
	QStyle__PM_IndicatorHeight                    QStyle__PixelMetric = QStyle__PixelMetric(38)
	QStyle__PM_ExclusiveIndicatorWidth            QStyle__PixelMetric = QStyle__PixelMetric(39)
	QStyle__PM_ExclusiveIndicatorHeight           QStyle__PixelMetric = QStyle__PixelMetric(40)
	QStyle__PM_DialogButtonsSeparator             QStyle__PixelMetric = QStyle__PixelMetric(41)
	QStyle__PM_DialogButtonsButtonWidth           QStyle__PixelMetric = QStyle__PixelMetric(42)
	QStyle__PM_DialogButtonsButtonHeight          QStyle__PixelMetric = QStyle__PixelMetric(43)
	QStyle__PM_MdiSubWindowFrameWidth             QStyle__PixelMetric = QStyle__PixelMetric(44)
	QStyle__PM_MDIFrameWidth                      QStyle__PixelMetric = QStyle__PixelMetric(QStyle__PM_MdiSubWindowFrameWidth)
	QStyle__PM_MdiSubWindowMinimizedWidth         QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_MdiSubWindowMinimizedWidth_Type())
	QStyle__PM_MDIMinimizedWidth                  QStyle__PixelMetric = QStyle__PixelMetric(QStyle__PM_MdiSubWindowMinimizedWidth)
	QStyle__PM_HeaderMargin                       QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_HeaderMargin_Type())
	QStyle__PM_HeaderMarkSize                     QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_HeaderMarkSize_Type())
	QStyle__PM_HeaderGripMargin                   QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_HeaderGripMargin_Type())
	QStyle__PM_TabBarTabShiftHorizontal           QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_TabBarTabShiftHorizontal_Type())
	QStyle__PM_TabBarTabShiftVertical             QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_TabBarTabShiftVertical_Type())
	QStyle__PM_TabBarScrollButtonWidth            QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_TabBarScrollButtonWidth_Type())
	QStyle__PM_ToolBarFrameWidth                  QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_ToolBarFrameWidth_Type())
	QStyle__PM_ToolBarHandleExtent                QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_ToolBarHandleExtent_Type())
	QStyle__PM_ToolBarItemSpacing                 QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_ToolBarItemSpacing_Type())
	QStyle__PM_ToolBarItemMargin                  QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_ToolBarItemMargin_Type())
	QStyle__PM_ToolBarSeparatorExtent             QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_ToolBarSeparatorExtent_Type())
	QStyle__PM_ToolBarExtensionExtent             QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_ToolBarExtensionExtent_Type())
	QStyle__PM_SpinBoxSliderHeight                QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_SpinBoxSliderHeight_Type())
	QStyle__PM_DefaultTopLevelMargin              QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_DefaultTopLevelMargin_Type())
	QStyle__PM_DefaultChildMargin                 QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_DefaultChildMargin_Type())
	QStyle__PM_DefaultLayoutSpacing               QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_DefaultLayoutSpacing_Type())
	QStyle__PM_ToolBarIconSize                    QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_ToolBarIconSize_Type())
	QStyle__PM_ListViewIconSize                   QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_ListViewIconSize_Type())
	QStyle__PM_IconViewIconSize                   QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_IconViewIconSize_Type())
	QStyle__PM_SmallIconSize                      QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_SmallIconSize_Type())
	QStyle__PM_LargeIconSize                      QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_LargeIconSize_Type())
	QStyle__PM_FocusFrameVMargin                  QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_FocusFrameVMargin_Type())
	QStyle__PM_FocusFrameHMargin                  QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_FocusFrameHMargin_Type())
	QStyle__PM_ToolTipLabelFrameWidth             QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_ToolTipLabelFrameWidth_Type())
	QStyle__PM_CheckBoxLabelSpacing               QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_CheckBoxLabelSpacing_Type())
	QStyle__PM_TabBarIconSize                     QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_TabBarIconSize_Type())
	QStyle__PM_SizeGripSize                       QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_SizeGripSize_Type())
	QStyle__PM_DockWidgetTitleMargin              QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_DockWidgetTitleMargin_Type())
	QStyle__PM_MessageBoxIconSize                 QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_MessageBoxIconSize_Type())
	QStyle__PM_ButtonIconSize                     QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_ButtonIconSize_Type())
	QStyle__PM_DockWidgetTitleBarButtonMargin     QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_DockWidgetTitleBarButtonMargin_Type())
	QStyle__PM_RadioButtonLabelSpacing            QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_RadioButtonLabelSpacing_Type())
	QStyle__PM_LayoutLeftMargin                   QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_LayoutLeftMargin_Type())
	QStyle__PM_LayoutTopMargin                    QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_LayoutTopMargin_Type())
	QStyle__PM_LayoutRightMargin                  QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_LayoutRightMargin_Type())
	QStyle__PM_LayoutBottomMargin                 QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_LayoutBottomMargin_Type())
	QStyle__PM_LayoutHorizontalSpacing            QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_LayoutHorizontalSpacing_Type())
	QStyle__PM_LayoutVerticalSpacing              QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_LayoutVerticalSpacing_Type())
	QStyle__PM_TabBar_ScrollButtonOverlap         QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_TabBar_ScrollButtonOverlap_Type())
	QStyle__PM_TextCursorWidth                    QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_TextCursorWidth_Type())
	QStyle__PM_TabCloseIndicatorWidth             QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_TabCloseIndicatorWidth_Type())
	QStyle__PM_TabCloseIndicatorHeight            QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_TabCloseIndicatorHeight_Type())
	QStyle__PM_ScrollView_ScrollBarSpacing        QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_ScrollView_ScrollBarSpacing_Type())
	QStyle__PM_ScrollView_ScrollBarOverlap        QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_ScrollView_ScrollBarOverlap_Type())
	QStyle__PM_SubMenuOverlap                     QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_SubMenuOverlap_Type())
	QStyle__PM_TreeViewIndentation                QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_TreeViewIndentation_Type())
	QStyle__PM_HeaderDefaultSectionSizeHorizontal QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_HeaderDefaultSectionSizeHorizontal_Type())
	QStyle__PM_HeaderDefaultSectionSizeVertical   QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_HeaderDefaultSectionSizeVertical_Type())
	QStyle__PM_TitleBarButtonIconSize             QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_TitleBarButtonIconSize_Type())
	QStyle__PM_TitleBarButtonSize                 QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_TitleBarButtonSize_Type())
	QStyle__PM_CustomBase                         QStyle__PixelMetric = QStyle__PixelMetric(0xf0000000)
)

//go:generate stringer -type=QStyle__PrimitiveElement
//QStyle::PrimitiveElement
type QStyle__PrimitiveElement int64

var (
	QStyle__PE_Frame                           QStyle__PrimitiveElement = QStyle__PrimitiveElement(0)
	QStyle__PE_FrameDefaultButton              QStyle__PrimitiveElement = QStyle__PrimitiveElement(1)
	QStyle__PE_FrameDockWidget                 QStyle__PrimitiveElement = QStyle__PrimitiveElement(2)
	QStyle__PE_FrameFocusRect                  QStyle__PrimitiveElement = QStyle__PrimitiveElement(3)
	QStyle__PE_FrameGroupBox                   QStyle__PrimitiveElement = QStyle__PrimitiveElement(4)
	QStyle__PE_FrameLineEdit                   QStyle__PrimitiveElement = QStyle__PrimitiveElement(5)
	QStyle__PE_FrameMenu                       QStyle__PrimitiveElement = QStyle__PrimitiveElement(6)
	QStyle__PE_FrameStatusBar                  QStyle__PrimitiveElement = QStyle__PrimitiveElement(7)
	QStyle__PE_FrameStatusBarItem              QStyle__PrimitiveElement = QStyle__PrimitiveElement(QStyle__PE_FrameStatusBar)
	QStyle__PE_FrameTabWidget                  QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_FrameTabWidget_Type())
	QStyle__PE_FrameWindow                     QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_FrameWindow_Type())
	QStyle__PE_FrameButtonBevel                QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_FrameButtonBevel_Type())
	QStyle__PE_FrameButtonTool                 QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_FrameButtonTool_Type())
	QStyle__PE_FrameTabBarBase                 QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_FrameTabBarBase_Type())
	QStyle__PE_PanelButtonCommand              QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_PanelButtonCommand_Type())
	QStyle__PE_PanelButtonBevel                QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_PanelButtonBevel_Type())
	QStyle__PE_PanelButtonTool                 QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_PanelButtonTool_Type())
	QStyle__PE_PanelMenuBar                    QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_PanelMenuBar_Type())
	QStyle__PE_PanelToolBar                    QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_PanelToolBar_Type())
	QStyle__PE_PanelLineEdit                   QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_PanelLineEdit_Type())
	QStyle__PE_IndicatorArrowDown              QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorArrowDown_Type())
	QStyle__PE_IndicatorArrowLeft              QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorArrowLeft_Type())
	QStyle__PE_IndicatorArrowRight             QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorArrowRight_Type())
	QStyle__PE_IndicatorArrowUp                QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorArrowUp_Type())
	QStyle__PE_IndicatorBranch                 QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorBranch_Type())
	QStyle__PE_IndicatorButtonDropDown         QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorButtonDropDown_Type())
	QStyle__PE_IndicatorViewItemCheck          QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorViewItemCheck_Type())
	QStyle__PE_IndicatorItemViewItemCheck      QStyle__PrimitiveElement = QStyle__PrimitiveElement(QStyle__PE_IndicatorViewItemCheck)
	QStyle__PE_IndicatorCheckBox               QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorCheckBox_Type())
	QStyle__PE_IndicatorDockWidgetResizeHandle QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorDockWidgetResizeHandle_Type())
	QStyle__PE_IndicatorHeaderArrow            QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorHeaderArrow_Type())
	QStyle__PE_IndicatorMenuCheckMark          QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorMenuCheckMark_Type())
	QStyle__PE_IndicatorProgressChunk          QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorProgressChunk_Type())
	QStyle__PE_IndicatorRadioButton            QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorRadioButton_Type())
	QStyle__PE_IndicatorSpinDown               QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorSpinDown_Type())
	QStyle__PE_IndicatorSpinMinus              QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorSpinMinus_Type())
	QStyle__PE_IndicatorSpinPlus               QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorSpinPlus_Type())
	QStyle__PE_IndicatorSpinUp                 QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorSpinUp_Type())
	QStyle__PE_IndicatorToolBarHandle          QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorToolBarHandle_Type())
	QStyle__PE_IndicatorToolBarSeparator       QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorToolBarSeparator_Type())
	QStyle__PE_PanelTipLabel                   QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_PanelTipLabel_Type())
	QStyle__PE_IndicatorTabTear                QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorTabTear_Type())
	QStyle__PE_IndicatorTabTearLeft            QStyle__PrimitiveElement = QStyle__PrimitiveElement(QStyle__PE_IndicatorTabTear)
	QStyle__PE_PanelScrollAreaCorner           QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_PanelScrollAreaCorner_Type())
	QStyle__PE_Widget                          QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_Widget_Type())
	QStyle__PE_IndicatorColumnViewArrow        QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorColumnViewArrow_Type())
	QStyle__PE_IndicatorItemViewItemDrop       QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorItemViewItemDrop_Type())
	QStyle__PE_PanelItemViewItem               QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_PanelItemViewItem_Type())
	QStyle__PE_PanelItemViewRow                QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_PanelItemViewRow_Type())
	QStyle__PE_PanelStatusBar                  QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_PanelStatusBar_Type())
	QStyle__PE_IndicatorTabClose               QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorTabClose_Type())
	QStyle__PE_PanelMenu                       QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_PanelMenu_Type())
	QStyle__PE_IndicatorTabTearRight           QStyle__PrimitiveElement = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorTabTearRight_Type())
	QStyle__PE_CustomBase                      QStyle__PrimitiveElement = QStyle__PrimitiveElement(0xf000000)
)

//go:generate stringer -type=QStyle__RequestSoftwareInputPanel
//QStyle::RequestSoftwareInputPanel
type QStyle__RequestSoftwareInputPanel int64

const (
	QStyle__RSIP_OnMouseClickAndAlreadyFocused QStyle__RequestSoftwareInputPanel = QStyle__RequestSoftwareInputPanel(0)
	QStyle__RSIP_OnMouseClick                  QStyle__RequestSoftwareInputPanel = QStyle__RequestSoftwareInputPanel(1)
)

//go:generate stringer -type=QStyle__StandardPixmap
//QStyle::StandardPixmap
type QStyle__StandardPixmap int64

const (
	QStyle__SP_TitleBarMenuButton               QStyle__StandardPixmap = QStyle__StandardPixmap(0)
	QStyle__SP_TitleBarMinButton                QStyle__StandardPixmap = QStyle__StandardPixmap(1)
	QStyle__SP_TitleBarMaxButton                QStyle__StandardPixmap = QStyle__StandardPixmap(2)
	QStyle__SP_TitleBarCloseButton              QStyle__StandardPixmap = QStyle__StandardPixmap(3)
	QStyle__SP_TitleBarNormalButton             QStyle__StandardPixmap = QStyle__StandardPixmap(4)
	QStyle__SP_TitleBarShadeButton              QStyle__StandardPixmap = QStyle__StandardPixmap(5)
	QStyle__SP_TitleBarUnshadeButton            QStyle__StandardPixmap = QStyle__StandardPixmap(6)
	QStyle__SP_TitleBarContextHelpButton        QStyle__StandardPixmap = QStyle__StandardPixmap(7)
	QStyle__SP_DockWidgetCloseButton            QStyle__StandardPixmap = QStyle__StandardPixmap(8)
	QStyle__SP_MessageBoxInformation            QStyle__StandardPixmap = QStyle__StandardPixmap(9)
	QStyle__SP_MessageBoxWarning                QStyle__StandardPixmap = QStyle__StandardPixmap(10)
	QStyle__SP_MessageBoxCritical               QStyle__StandardPixmap = QStyle__StandardPixmap(11)
	QStyle__SP_MessageBoxQuestion               QStyle__StandardPixmap = QStyle__StandardPixmap(12)
	QStyle__SP_DesktopIcon                      QStyle__StandardPixmap = QStyle__StandardPixmap(13)
	QStyle__SP_TrashIcon                        QStyle__StandardPixmap = QStyle__StandardPixmap(14)
	QStyle__SP_ComputerIcon                     QStyle__StandardPixmap = QStyle__StandardPixmap(15)
	QStyle__SP_DriveFDIcon                      QStyle__StandardPixmap = QStyle__StandardPixmap(16)
	QStyle__SP_DriveHDIcon                      QStyle__StandardPixmap = QStyle__StandardPixmap(17)
	QStyle__SP_DriveCDIcon                      QStyle__StandardPixmap = QStyle__StandardPixmap(18)
	QStyle__SP_DriveDVDIcon                     QStyle__StandardPixmap = QStyle__StandardPixmap(19)
	QStyle__SP_DriveNetIcon                     QStyle__StandardPixmap = QStyle__StandardPixmap(20)
	QStyle__SP_DirOpenIcon                      QStyle__StandardPixmap = QStyle__StandardPixmap(21)
	QStyle__SP_DirClosedIcon                    QStyle__StandardPixmap = QStyle__StandardPixmap(22)
	QStyle__SP_DirLinkIcon                      QStyle__StandardPixmap = QStyle__StandardPixmap(23)
	QStyle__SP_DirLinkOpenIcon                  QStyle__StandardPixmap = QStyle__StandardPixmap(24)
	QStyle__SP_FileIcon                         QStyle__StandardPixmap = QStyle__StandardPixmap(25)
	QStyle__SP_FileLinkIcon                     QStyle__StandardPixmap = QStyle__StandardPixmap(26)
	QStyle__SP_ToolBarHorizontalExtensionButton QStyle__StandardPixmap = QStyle__StandardPixmap(27)
	QStyle__SP_ToolBarVerticalExtensionButton   QStyle__StandardPixmap = QStyle__StandardPixmap(28)
	QStyle__SP_FileDialogStart                  QStyle__StandardPixmap = QStyle__StandardPixmap(29)
	QStyle__SP_FileDialogEnd                    QStyle__StandardPixmap = QStyle__StandardPixmap(30)
	QStyle__SP_FileDialogToParent               QStyle__StandardPixmap = QStyle__StandardPixmap(31)
	QStyle__SP_FileDialogNewFolder              QStyle__StandardPixmap = QStyle__StandardPixmap(32)
	QStyle__SP_FileDialogDetailedView           QStyle__StandardPixmap = QStyle__StandardPixmap(33)
	QStyle__SP_FileDialogInfoView               QStyle__StandardPixmap = QStyle__StandardPixmap(34)
	QStyle__SP_FileDialogContentsView           QStyle__StandardPixmap = QStyle__StandardPixmap(35)
	QStyle__SP_FileDialogListView               QStyle__StandardPixmap = QStyle__StandardPixmap(36)
	QStyle__SP_FileDialogBack                   QStyle__StandardPixmap = QStyle__StandardPixmap(37)
	QStyle__SP_DirIcon                          QStyle__StandardPixmap = QStyle__StandardPixmap(38)
	QStyle__SP_DialogOkButton                   QStyle__StandardPixmap = QStyle__StandardPixmap(39)
	QStyle__SP_DialogCancelButton               QStyle__StandardPixmap = QStyle__StandardPixmap(40)
	QStyle__SP_DialogHelpButton                 QStyle__StandardPixmap = QStyle__StandardPixmap(41)
	QStyle__SP_DialogOpenButton                 QStyle__StandardPixmap = QStyle__StandardPixmap(42)
	QStyle__SP_DialogSaveButton                 QStyle__StandardPixmap = QStyle__StandardPixmap(43)
	QStyle__SP_DialogCloseButton                QStyle__StandardPixmap = QStyle__StandardPixmap(44)
	QStyle__SP_DialogApplyButton                QStyle__StandardPixmap = QStyle__StandardPixmap(45)
	QStyle__SP_DialogResetButton                QStyle__StandardPixmap = QStyle__StandardPixmap(46)
	QStyle__SP_DialogDiscardButton              QStyle__StandardPixmap = QStyle__StandardPixmap(47)
	QStyle__SP_DialogYesButton                  QStyle__StandardPixmap = QStyle__StandardPixmap(48)
	QStyle__SP_DialogNoButton                   QStyle__StandardPixmap = QStyle__StandardPixmap(49)
	QStyle__SP_ArrowUp                          QStyle__StandardPixmap = QStyle__StandardPixmap(50)
	QStyle__SP_ArrowDown                        QStyle__StandardPixmap = QStyle__StandardPixmap(51)
	QStyle__SP_ArrowLeft                        QStyle__StandardPixmap = QStyle__StandardPixmap(52)
	QStyle__SP_ArrowRight                       QStyle__StandardPixmap = QStyle__StandardPixmap(53)
	QStyle__SP_ArrowBack                        QStyle__StandardPixmap = QStyle__StandardPixmap(54)
	QStyle__SP_ArrowForward                     QStyle__StandardPixmap = QStyle__StandardPixmap(55)
	QStyle__SP_DirHomeIcon                      QStyle__StandardPixmap = QStyle__StandardPixmap(56)
	QStyle__SP_CommandLink                      QStyle__StandardPixmap = QStyle__StandardPixmap(57)
	QStyle__SP_VistaShield                      QStyle__StandardPixmap = QStyle__StandardPixmap(58)
	QStyle__SP_BrowserReload                    QStyle__StandardPixmap = QStyle__StandardPixmap(59)
	QStyle__SP_BrowserStop                      QStyle__StandardPixmap = QStyle__StandardPixmap(60)
	QStyle__SP_MediaPlay                        QStyle__StandardPixmap = QStyle__StandardPixmap(61)
	QStyle__SP_MediaStop                        QStyle__StandardPixmap = QStyle__StandardPixmap(62)
	QStyle__SP_MediaPause                       QStyle__StandardPixmap = QStyle__StandardPixmap(63)
	QStyle__SP_MediaSkipForward                 QStyle__StandardPixmap = QStyle__StandardPixmap(64)
	QStyle__SP_MediaSkipBackward                QStyle__StandardPixmap = QStyle__StandardPixmap(65)
	QStyle__SP_MediaSeekForward                 QStyle__StandardPixmap = QStyle__StandardPixmap(66)
	QStyle__SP_MediaSeekBackward                QStyle__StandardPixmap = QStyle__StandardPixmap(67)
	QStyle__SP_MediaVolume                      QStyle__StandardPixmap = QStyle__StandardPixmap(68)
	QStyle__SP_MediaVolumeMuted                 QStyle__StandardPixmap = QStyle__StandardPixmap(69)
	QStyle__SP_LineEditClearButton              QStyle__StandardPixmap = QStyle__StandardPixmap(70)
	QStyle__SP_CustomBase                       QStyle__StandardPixmap = QStyle__StandardPixmap(0xf0000000)
)

//go:generate stringer -type=QStyle__StateFlag
//QStyle::StateFlag
type QStyle__StateFlag int64

const (
	QStyle__State_None                QStyle__StateFlag = QStyle__StateFlag(0x00000000)
	QStyle__State_Enabled             QStyle__StateFlag = QStyle__StateFlag(0x00000001)
	QStyle__State_Raised              QStyle__StateFlag = QStyle__StateFlag(0x00000002)
	QStyle__State_Sunken              QStyle__StateFlag = QStyle__StateFlag(0x00000004)
	QStyle__State_Off                 QStyle__StateFlag = QStyle__StateFlag(0x00000008)
	QStyle__State_NoChange            QStyle__StateFlag = QStyle__StateFlag(0x00000010)
	QStyle__State_On                  QStyle__StateFlag = QStyle__StateFlag(0x00000020)
	QStyle__State_DownArrow           QStyle__StateFlag = QStyle__StateFlag(0x00000040)
	QStyle__State_Horizontal          QStyle__StateFlag = QStyle__StateFlag(0x00000080)
	QStyle__State_HasFocus            QStyle__StateFlag = QStyle__StateFlag(0x00000100)
	QStyle__State_Top                 QStyle__StateFlag = QStyle__StateFlag(0x00000200)
	QStyle__State_Bottom              QStyle__StateFlag = QStyle__StateFlag(0x00000400)
	QStyle__State_FocusAtBorder       QStyle__StateFlag = QStyle__StateFlag(0x00000800)
	QStyle__State_AutoRaise           QStyle__StateFlag = QStyle__StateFlag(0x00001000)
	QStyle__State_MouseOver           QStyle__StateFlag = QStyle__StateFlag(0x00002000)
	QStyle__State_UpArrow             QStyle__StateFlag = QStyle__StateFlag(0x00004000)
	QStyle__State_Selected            QStyle__StateFlag = QStyle__StateFlag(0x00008000)
	QStyle__State_Active              QStyle__StateFlag = QStyle__StateFlag(0x00010000)
	QStyle__State_Window              QStyle__StateFlag = QStyle__StateFlag(0x00020000)
	QStyle__State_Open                QStyle__StateFlag = QStyle__StateFlag(0x00040000)
	QStyle__State_Children            QStyle__StateFlag = QStyle__StateFlag(0x00080000)
	QStyle__State_Item                QStyle__StateFlag = QStyle__StateFlag(0x00100000)
	QStyle__State_Sibling             QStyle__StateFlag = QStyle__StateFlag(0x00200000)
	QStyle__State_Editing             QStyle__StateFlag = QStyle__StateFlag(0x00400000)
	QStyle__State_KeyboardFocusChange QStyle__StateFlag = QStyle__StateFlag(0x00800000)
	QStyle__State_HasEditFocus        QStyle__StateFlag = QStyle__StateFlag(0x01000000)
	QStyle__State_ReadOnly            QStyle__StateFlag = QStyle__StateFlag(0x02000000)
	QStyle__State_Small               QStyle__StateFlag = QStyle__StateFlag(0x04000000)
	QStyle__State_Mini                QStyle__StateFlag = QStyle__StateFlag(0x08000000)
)

//go:generate stringer -type=QStyle__StyleHint
//QStyle::StyleHint
type QStyle__StyleHint int64

var (
	QStyle__SH_EtchDisabledText                               QStyle__StyleHint = QStyle__StyleHint(0)
	QStyle__SH_DitherDisabledText                             QStyle__StyleHint = QStyle__StyleHint(1)
	QStyle__SH_ScrollBar_MiddleClickAbsolutePosition          QStyle__StyleHint = QStyle__StyleHint(2)
	QStyle__SH_ScrollBar_ScrollWhenPointerLeavesControl       QStyle__StyleHint = QStyle__StyleHint(3)
	QStyle__SH_TabBar_SelectMouseType                         QStyle__StyleHint = QStyle__StyleHint(4)
	QStyle__SH_TabBar_Alignment                               QStyle__StyleHint = QStyle__StyleHint(5)
	QStyle__SH_Header_ArrowAlignment                          QStyle__StyleHint = QStyle__StyleHint(6)
	QStyle__SH_Slider_SnapToValue                             QStyle__StyleHint = QStyle__StyleHint(7)
	QStyle__SH_Slider_SloppyKeyEvents                         QStyle__StyleHint = QStyle__StyleHint(8)
	QStyle__SH_ProgressDialog_CenterCancelButton              QStyle__StyleHint = QStyle__StyleHint(9)
	QStyle__SH_ProgressDialog_TextLabelAlignment              QStyle__StyleHint = QStyle__StyleHint(10)
	QStyle__SH_PrintDialog_RightAlignButtons                  QStyle__StyleHint = QStyle__StyleHint(11)
	QStyle__SH_MainWindow_SpaceBelowMenuBar                   QStyle__StyleHint = QStyle__StyleHint(12)
	QStyle__SH_FontDialog_SelectAssociatedText                QStyle__StyleHint = QStyle__StyleHint(13)
	QStyle__SH_Menu_AllowActiveAndDisabled                    QStyle__StyleHint = QStyle__StyleHint(14)
	QStyle__SH_Menu_SpaceActivatesItem                        QStyle__StyleHint = QStyle__StyleHint(15)
	QStyle__SH_Menu_SubMenuPopupDelay                         QStyle__StyleHint = QStyle__StyleHint(16)
	QStyle__SH_ScrollView_FrameOnlyAroundContents             QStyle__StyleHint = QStyle__StyleHint(17)
	QStyle__SH_MenuBar_AltKeyNavigation                       QStyle__StyleHint = QStyle__StyleHint(18)
	QStyle__SH_ComboBox_ListMouseTracking                     QStyle__StyleHint = QStyle__StyleHint(19)
	QStyle__SH_Menu_MouseTracking                             QStyle__StyleHint = QStyle__StyleHint(20)
	QStyle__SH_MenuBar_MouseTracking                          QStyle__StyleHint = QStyle__StyleHint(21)
	QStyle__SH_ItemView_ChangeHighlightOnFocus                QStyle__StyleHint = QStyle__StyleHint(22)
	QStyle__SH_Widget_ShareActivation                         QStyle__StyleHint = QStyle__StyleHint(23)
	QStyle__SH_Workspace_FillSpaceOnMaximize                  QStyle__StyleHint = QStyle__StyleHint(24)
	QStyle__SH_ComboBox_Popup                                 QStyle__StyleHint = QStyle__StyleHint(25)
	QStyle__SH_TitleBar_NoBorder                              QStyle__StyleHint = QStyle__StyleHint(26)
	QStyle__SH_Slider_StopMouseOverSlider                     QStyle__StyleHint = QStyle__StyleHint(27)
	QStyle__SH_ScrollBar_StopMouseOverSlider                  QStyle__StyleHint = QStyle__StyleHint(QStyle__SH_Slider_StopMouseOverSlider)
	QStyle__SH_BlinkCursorWhenTextSelected                    QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_BlinkCursorWhenTextSelected_Type())
	QStyle__SH_RichText_FullWidthSelection                    QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_RichText_FullWidthSelection_Type())
	QStyle__SH_Menu_Scrollable                                QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Menu_Scrollable_Type())
	QStyle__SH_GroupBox_TextLabelVerticalAlignment            QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_GroupBox_TextLabelVerticalAlignment_Type())
	QStyle__SH_GroupBox_TextLabelColor                        QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_GroupBox_TextLabelColor_Type())
	QStyle__SH_Menu_SloppySubMenus                            QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Menu_SloppySubMenus_Type())
	QStyle__SH_Table_GridLineColor                            QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Table_GridLineColor_Type())
	QStyle__SH_LineEdit_PasswordCharacter                     QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_LineEdit_PasswordCharacter_Type())
	QStyle__SH_DialogButtons_DefaultButton                    QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_DialogButtons_DefaultButton_Type())
	QStyle__SH_ToolBox_SelectedPageTitleBold                  QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ToolBox_SelectedPageTitleBold_Type())
	QStyle__SH_TabBar_PreferNoArrows                          QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_TabBar_PreferNoArrows_Type())
	QStyle__SH_ScrollBar_LeftClickAbsolutePosition            QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ScrollBar_LeftClickAbsolutePosition_Type())
	QStyle__SH_ListViewExpand_SelectMouseType                 QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ListViewExpand_SelectMouseType_Type())
	QStyle__SH_UnderlineShortcut                              QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_UnderlineShortcut_Type())
	QStyle__SH_SpinBox_AnimateButton                          QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_SpinBox_AnimateButton_Type())
	QStyle__SH_SpinBox_KeyPressAutoRepeatRate                 QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_SpinBox_KeyPressAutoRepeatRate_Type())
	QStyle__SH_SpinBox_ClickAutoRepeatRate                    QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_SpinBox_ClickAutoRepeatRate_Type())
	QStyle__SH_Menu_FillScreenWithScroll                      QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Menu_FillScreenWithScroll_Type())
	QStyle__SH_ToolTipLabel_Opacity                           QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ToolTipLabel_Opacity_Type())
	QStyle__SH_DrawMenuBarSeparator                           QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_DrawMenuBarSeparator_Type())
	QStyle__SH_TitleBar_ModifyNotification                    QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_TitleBar_ModifyNotification_Type())
	QStyle__SH_Button_FocusPolicy                             QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Button_FocusPolicy_Type())
	QStyle__SH_MessageBox_UseBorderForButtonSpacing           QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_MessageBox_UseBorderForButtonSpacing_Type())
	QStyle__SH_TitleBar_AutoRaise                             QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_TitleBar_AutoRaise_Type())
	QStyle__SH_ToolButton_PopupDelay                          QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ToolButton_PopupDelay_Type())
	QStyle__SH_FocusFrame_Mask                                QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_FocusFrame_Mask_Type())
	QStyle__SH_RubberBand_Mask                                QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_RubberBand_Mask_Type())
	QStyle__SH_WindowFrame_Mask                               QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_WindowFrame_Mask_Type())
	QStyle__SH_SpinControls_DisableOnBounds                   QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_SpinControls_DisableOnBounds_Type())
	QStyle__SH_Dial_BackgroundRole                            QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Dial_BackgroundRole_Type())
	QStyle__SH_ComboBox_LayoutDirection                       QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ComboBox_LayoutDirection_Type())
	QStyle__SH_ItemView_EllipsisLocation                      QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ItemView_EllipsisLocation_Type())
	QStyle__SH_ItemView_ShowDecorationSelected                QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ItemView_ShowDecorationSelected_Type())
	QStyle__SH_ItemView_ActivateItemOnSingleClick             QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ItemView_ActivateItemOnSingleClick_Type())
	QStyle__SH_ScrollBar_ContextMenu                          QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ScrollBar_ContextMenu_Type())
	QStyle__SH_ScrollBar_RollBetweenButtons                   QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ScrollBar_RollBetweenButtons_Type())
	QStyle__SH_Slider_AbsoluteSetButtons                      QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Slider_AbsoluteSetButtons_Type())
	QStyle__SH_Slider_PageSetButtons                          QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Slider_PageSetButtons_Type())
	QStyle__SH_Menu_KeyboardSearch                            QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Menu_KeyboardSearch_Type())
	QStyle__SH_TabBar_ElideMode                               QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_TabBar_ElideMode_Type())
	QStyle__SH_DialogButtonLayout                             QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_DialogButtonLayout_Type())
	QStyle__SH_ComboBox_PopupFrameStyle                       QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ComboBox_PopupFrameStyle_Type())
	QStyle__SH_MessageBox_TextInteractionFlags                QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_MessageBox_TextInteractionFlags_Type())
	QStyle__SH_DialogButtonBox_ButtonsHaveIcons               QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_DialogButtonBox_ButtonsHaveIcons_Type())
	QStyle__SH_SpellCheckUnderlineStyle                       QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_SpellCheckUnderlineStyle_Type())
	QStyle__SH_MessageBox_CenterButtons                       QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_MessageBox_CenterButtons_Type())
	QStyle__SH_Menu_SelectionWrap                             QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Menu_SelectionWrap_Type())
	QStyle__SH_ItemView_MovementWithoutUpdatingSelection      QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ItemView_MovementWithoutUpdatingSelection_Type())
	QStyle__SH_ToolTip_Mask                                   QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ToolTip_Mask_Type())
	QStyle__SH_FocusFrame_AboveWidget                         QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_FocusFrame_AboveWidget_Type())
	QStyle__SH_TextControl_FocusIndicatorTextCharFormat       QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_TextControl_FocusIndicatorTextCharFormat_Type())
	QStyle__SH_WizardStyle                                    QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_WizardStyle_Type())
	QStyle__SH_ItemView_ArrowKeysNavigateIntoChildren         QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ItemView_ArrowKeysNavigateIntoChildren_Type())
	QStyle__SH_Menu_Mask                                      QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Menu_Mask_Type())
	QStyle__SH_Menu_FlashTriggeredItem                        QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Menu_FlashTriggeredItem_Type())
	QStyle__SH_Menu_FadeOutOnHide                             QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Menu_FadeOutOnHide_Type())
	QStyle__SH_SpinBox_ClickAutoRepeatThreshold               QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_SpinBox_ClickAutoRepeatThreshold_Type())
	QStyle__SH_ItemView_PaintAlternatingRowColorsForEmptyArea QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ItemView_PaintAlternatingRowColorsForEmptyArea_Type())
	QStyle__SH_FormLayoutWrapPolicy                           QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_FormLayoutWrapPolicy_Type())
	QStyle__SH_TabWidget_DefaultTabPosition                   QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_TabWidget_DefaultTabPosition_Type())
	QStyle__SH_ToolBar_Movable                                QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ToolBar_Movable_Type())
	QStyle__SH_FormLayoutFieldGrowthPolicy                    QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_FormLayoutFieldGrowthPolicy_Type())
	QStyle__SH_FormLayoutFormAlignment                        QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_FormLayoutFormAlignment_Type())
	QStyle__SH_FormLayoutLabelAlignment                       QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_FormLayoutLabelAlignment_Type())
	QStyle__SH_ItemView_DrawDelegateFrame                     QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ItemView_DrawDelegateFrame_Type())
	QStyle__SH_TabBar_CloseButtonPosition                     QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_TabBar_CloseButtonPosition_Type())
	QStyle__SH_DockWidget_ButtonsHaveFrame                    QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_DockWidget_ButtonsHaveFrame_Type())
	QStyle__SH_ToolButtonStyle                                QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ToolButtonStyle_Type())
	QStyle__SH_RequestSoftwareInputPanel                      QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_RequestSoftwareInputPanel_Type())
	QStyle__SH_ScrollBar_Transient                            QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ScrollBar_Transient_Type())
	QStyle__SH_Menu_SupportsSections                          QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Menu_SupportsSections_Type())
	QStyle__SH_ToolTip_WakeUpDelay                            QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ToolTip_WakeUpDelay_Type())
	QStyle__SH_ToolTip_FallAsleepDelay                        QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ToolTip_FallAsleepDelay_Type())
	QStyle__SH_Widget_Animate                                 QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Widget_Animate_Type())
	QStyle__SH_Splitter_OpaqueResize                          QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Splitter_OpaqueResize_Type())
	QStyle__SH_ComboBox_UseNativePopup                        QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ComboBox_UseNativePopup_Type())
	QStyle__SH_LineEdit_PasswordMaskDelay                     QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_LineEdit_PasswordMaskDelay_Type())
	QStyle__SH_TabBar_ChangeCurrentDelay                      QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_TabBar_ChangeCurrentDelay_Type())
	QStyle__SH_Menu_SubMenuUniDirection                       QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Menu_SubMenuUniDirection_Type())
	QStyle__SH_Menu_SubMenuUniDirectionFailCount              QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Menu_SubMenuUniDirectionFailCount_Type())
	QStyle__SH_Menu_SubMenuSloppySelectOtherActions           QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Menu_SubMenuSloppySelectOtherActions_Type())
	QStyle__SH_Menu_SubMenuSloppyCloseTimeout                 QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Menu_SubMenuSloppyCloseTimeout_Type())
	QStyle__SH_Menu_SubMenuResetWhenReenteringParent          QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Menu_SubMenuResetWhenReenteringParent_Type())
	QStyle__SH_Menu_SubMenuDontStartSloppyOnLeave             QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_Menu_SubMenuDontStartSloppyOnLeave_Type())
	QStyle__SH_ItemView_ScrollMode                            QStyle__StyleHint = QStyle__StyleHint(C.QStyle_SH_ItemView_ScrollMode_Type())
	QStyle__SH_CustomBase                                     QStyle__StyleHint = QStyle__StyleHint(0xf0000000)
)

//go:generate stringer -type=QStyle__SubControl
//QStyle::SubControl
type QStyle__SubControl int64

const (
	QStyle__SC_None                      QStyle__SubControl = QStyle__SubControl(0x00000000)
	QStyle__SC_ScrollBarAddLine          QStyle__SubControl = QStyle__SubControl(0x00000001)
	QStyle__SC_ScrollBarSubLine          QStyle__SubControl = QStyle__SubControl(0x00000002)
	QStyle__SC_ScrollBarAddPage          QStyle__SubControl = QStyle__SubControl(0x00000004)
	QStyle__SC_ScrollBarSubPage          QStyle__SubControl = QStyle__SubControl(0x00000008)
	QStyle__SC_ScrollBarFirst            QStyle__SubControl = QStyle__SubControl(0x00000010)
	QStyle__SC_ScrollBarLast             QStyle__SubControl = QStyle__SubControl(0x00000020)
	QStyle__SC_ScrollBarSlider           QStyle__SubControl = QStyle__SubControl(0x00000040)
	QStyle__SC_ScrollBarGroove           QStyle__SubControl = QStyle__SubControl(0x00000080)
	QStyle__SC_SpinBoxUp                 QStyle__SubControl = QStyle__SubControl(0x00000001)
	QStyle__SC_SpinBoxDown               QStyle__SubControl = QStyle__SubControl(0x00000002)
	QStyle__SC_SpinBoxFrame              QStyle__SubControl = QStyle__SubControl(0x00000004)
	QStyle__SC_SpinBoxEditField          QStyle__SubControl = QStyle__SubControl(0x00000008)
	QStyle__SC_ComboBoxFrame             QStyle__SubControl = QStyle__SubControl(0x00000001)
	QStyle__SC_ComboBoxEditField         QStyle__SubControl = QStyle__SubControl(0x00000002)
	QStyle__SC_ComboBoxArrow             QStyle__SubControl = QStyle__SubControl(0x00000004)
	QStyle__SC_ComboBoxListBoxPopup      QStyle__SubControl = QStyle__SubControl(0x00000008)
	QStyle__SC_SliderGroove              QStyle__SubControl = QStyle__SubControl(0x00000001)
	QStyle__SC_SliderHandle              QStyle__SubControl = QStyle__SubControl(0x00000002)
	QStyle__SC_SliderTickmarks           QStyle__SubControl = QStyle__SubControl(0x00000004)
	QStyle__SC_ToolButton                QStyle__SubControl = QStyle__SubControl(0x00000001)
	QStyle__SC_ToolButtonMenu            QStyle__SubControl = QStyle__SubControl(0x00000002)
	QStyle__SC_TitleBarSysMenu           QStyle__SubControl = QStyle__SubControl(0x00000001)
	QStyle__SC_TitleBarMinButton         QStyle__SubControl = QStyle__SubControl(0x00000002)
	QStyle__SC_TitleBarMaxButton         QStyle__SubControl = QStyle__SubControl(0x00000004)
	QStyle__SC_TitleBarCloseButton       QStyle__SubControl = QStyle__SubControl(0x00000008)
	QStyle__SC_TitleBarNormalButton      QStyle__SubControl = QStyle__SubControl(0x00000010)
	QStyle__SC_TitleBarShadeButton       QStyle__SubControl = QStyle__SubControl(0x00000020)
	QStyle__SC_TitleBarUnshadeButton     QStyle__SubControl = QStyle__SubControl(0x00000040)
	QStyle__SC_TitleBarContextHelpButton QStyle__SubControl = QStyle__SubControl(0x00000080)
	QStyle__SC_TitleBarLabel             QStyle__SubControl = QStyle__SubControl(0x00000100)
	QStyle__SC_DialGroove                QStyle__SubControl = QStyle__SubControl(0x00000001)
	QStyle__SC_DialHandle                QStyle__SubControl = QStyle__SubControl(0x00000002)
	QStyle__SC_DialTickmarks             QStyle__SubControl = QStyle__SubControl(0x00000004)
	QStyle__SC_GroupBoxCheckBox          QStyle__SubControl = QStyle__SubControl(0x00000001)
	QStyle__SC_GroupBoxLabel             QStyle__SubControl = QStyle__SubControl(0x00000002)
	QStyle__SC_GroupBoxContents          QStyle__SubControl = QStyle__SubControl(0x00000004)
	QStyle__SC_GroupBoxFrame             QStyle__SubControl = QStyle__SubControl(0x00000008)
	QStyle__SC_MdiMinButton              QStyle__SubControl = QStyle__SubControl(0x00000001)
	QStyle__SC_MdiNormalButton           QStyle__SubControl = QStyle__SubControl(0x00000002)
	QStyle__SC_MdiCloseButton            QStyle__SubControl = QStyle__SubControl(0x00000004)
	QStyle__SC_CustomBase                QStyle__SubControl = QStyle__SubControl(0xf0000000)
	QStyle__SC_All                       QStyle__SubControl = QStyle__SubControl(0xffffffff)
)

//go:generate stringer -type=QStyle__SubElement
//QStyle::SubElement
type QStyle__SubElement int64

var (
	QStyle__SE_PushButtonContents         QStyle__SubElement = QStyle__SubElement(0)
	QStyle__SE_PushButtonFocusRect        QStyle__SubElement = QStyle__SubElement(1)
	QStyle__SE_CheckBoxIndicator          QStyle__SubElement = QStyle__SubElement(2)
	QStyle__SE_CheckBoxContents           QStyle__SubElement = QStyle__SubElement(3)
	QStyle__SE_CheckBoxFocusRect          QStyle__SubElement = QStyle__SubElement(4)
	QStyle__SE_CheckBoxClickRect          QStyle__SubElement = QStyle__SubElement(5)
	QStyle__SE_RadioButtonIndicator       QStyle__SubElement = QStyle__SubElement(6)
	QStyle__SE_RadioButtonContents        QStyle__SubElement = QStyle__SubElement(7)
	QStyle__SE_RadioButtonFocusRect       QStyle__SubElement = QStyle__SubElement(8)
	QStyle__SE_RadioButtonClickRect       QStyle__SubElement = QStyle__SubElement(9)
	QStyle__SE_ComboBoxFocusRect          QStyle__SubElement = QStyle__SubElement(10)
	QStyle__SE_SliderFocusRect            QStyle__SubElement = QStyle__SubElement(11)
	QStyle__SE_ProgressBarGroove          QStyle__SubElement = QStyle__SubElement(12)
	QStyle__SE_ProgressBarContents        QStyle__SubElement = QStyle__SubElement(13)
	QStyle__SE_ProgressBarLabel           QStyle__SubElement = QStyle__SubElement(14)
	QStyle__SE_ToolBoxTabContents         QStyle__SubElement = QStyle__SubElement(15)
	QStyle__SE_HeaderLabel                QStyle__SubElement = QStyle__SubElement(16)
	QStyle__SE_HeaderArrow                QStyle__SubElement = QStyle__SubElement(17)
	QStyle__SE_TabWidgetTabBar            QStyle__SubElement = QStyle__SubElement(18)
	QStyle__SE_TabWidgetTabPane           QStyle__SubElement = QStyle__SubElement(19)
	QStyle__SE_TabWidgetTabContents       QStyle__SubElement = QStyle__SubElement(20)
	QStyle__SE_TabWidgetLeftCorner        QStyle__SubElement = QStyle__SubElement(21)
	QStyle__SE_TabWidgetRightCorner       QStyle__SubElement = QStyle__SubElement(22)
	QStyle__SE_ViewItemCheckIndicator     QStyle__SubElement = QStyle__SubElement(23)
	QStyle__SE_ItemViewItemCheckIndicator QStyle__SubElement = QStyle__SubElement(QStyle__SE_ViewItemCheckIndicator)
	QStyle__SE_TabBarTearIndicator        QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_TabBarTearIndicator_Type())
	QStyle__SE_TabBarTearIndicatorLeft    QStyle__SubElement = QStyle__SubElement(QStyle__SE_TabBarTearIndicator)
	QStyle__SE_TreeViewDisclosureItem     QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_TreeViewDisclosureItem_Type())
	QStyle__SE_LineEditContents           QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_LineEditContents_Type())
	QStyle__SE_FrameContents              QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_FrameContents_Type())
	QStyle__SE_DockWidgetCloseButton      QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_DockWidgetCloseButton_Type())
	QStyle__SE_DockWidgetFloatButton      QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_DockWidgetFloatButton_Type())
	QStyle__SE_DockWidgetTitleBarText     QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_DockWidgetTitleBarText_Type())
	QStyle__SE_DockWidgetIcon             QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_DockWidgetIcon_Type())
	QStyle__SE_CheckBoxLayoutItem         QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_CheckBoxLayoutItem_Type())
	QStyle__SE_ComboBoxLayoutItem         QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_ComboBoxLayoutItem_Type())
	QStyle__SE_DateTimeEditLayoutItem     QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_DateTimeEditLayoutItem_Type())
	QStyle__SE_DialogButtonBoxLayoutItem  QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_DialogButtonBoxLayoutItem_Type())
	QStyle__SE_LabelLayoutItem            QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_LabelLayoutItem_Type())
	QStyle__SE_ProgressBarLayoutItem      QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_ProgressBarLayoutItem_Type())
	QStyle__SE_PushButtonLayoutItem       QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_PushButtonLayoutItem_Type())
	QStyle__SE_RadioButtonLayoutItem      QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_RadioButtonLayoutItem_Type())
	QStyle__SE_SliderLayoutItem           QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_SliderLayoutItem_Type())
	QStyle__SE_SpinBoxLayoutItem          QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_SpinBoxLayoutItem_Type())
	QStyle__SE_ToolButtonLayoutItem       QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_ToolButtonLayoutItem_Type())
	QStyle__SE_FrameLayoutItem            QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_FrameLayoutItem_Type())
	QStyle__SE_GroupBoxLayoutItem         QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_GroupBoxLayoutItem_Type())
	QStyle__SE_TabWidgetLayoutItem        QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_TabWidgetLayoutItem_Type())
	QStyle__SE_ItemViewItemDecoration     QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_ItemViewItemDecoration_Type())
	QStyle__SE_ItemViewItemText           QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_ItemViewItemText_Type())
	QStyle__SE_ItemViewItemFocusRect      QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_ItemViewItemFocusRect_Type())
	QStyle__SE_TabBarTabLeftButton        QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_TabBarTabLeftButton_Type())
	QStyle__SE_TabBarTabRightButton       QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_TabBarTabRightButton_Type())
	QStyle__SE_TabBarTabText              QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_TabBarTabText_Type())
	QStyle__SE_ShapedFrameContents        QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_ShapedFrameContents_Type())
	QStyle__SE_ToolBarHandle              QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_ToolBarHandle_Type())
	QStyle__SE_TabBarScrollLeftButton     QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_TabBarScrollLeftButton_Type())
	QStyle__SE_TabBarScrollRightButton    QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_TabBarScrollRightButton_Type())
	QStyle__SE_TabBarTearIndicatorRight   QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_TabBarTearIndicatorRight_Type())
	QStyle__SE_CustomBase                 QStyle__SubElement = QStyle__SubElement(0xf0000000)
)

//go:generate stringer -type=QStyleHintReturn__HintReturnType
//QStyleHintReturn::HintReturnType
type QStyleHintReturn__HintReturnType int64

var (
	QStyleHintReturn__SH_Default QStyleHintReturn__HintReturnType = QStyleHintReturn__HintReturnType(0xf000)
	QStyleHintReturn__SH_Mask    QStyleHintReturn__HintReturnType = QStyleHintReturn__HintReturnType(C.QStyleHintReturn_SH_Mask_Type())
	QStyleHintReturn__SH_Variant QStyleHintReturn__HintReturnType = QStyleHintReturn__HintReturnType(C.QStyleHintReturn_SH_Variant_Type())
)

//go:generate stringer -type=QStyleHintReturn__StyleOptionType
//QStyleHintReturn::StyleOptionType
type QStyleHintReturn__StyleOptionType int64

var (
	QStyleHintReturn__Type QStyleHintReturn__StyleOptionType = QStyleHintReturn__StyleOptionType(QStyleHintReturn__SH_Default)
)

//go:generate stringer -type=QStyleHintReturn__StyleOptionVersion
//QStyleHintReturn::StyleOptionVersion
type QStyleHintReturn__StyleOptionVersion int64

var (
	QStyleHintReturn__Version QStyleHintReturn__StyleOptionVersion = QStyleHintReturn__StyleOptionVersion(1)
)

//go:generate stringer -type=QStyleHintReturnMask__StyleOptionType
//QStyleHintReturnMask::StyleOptionType
type QStyleHintReturnMask__StyleOptionType int64

var (
	QStyleHintReturnMask__Type QStyleHintReturnMask__StyleOptionType = QStyleHintReturnMask__StyleOptionType(QStyleHintReturn__SH_Mask)
)

//go:generate stringer -type=QStyleHintReturnMask__StyleOptionVersion
//QStyleHintReturnMask::StyleOptionVersion
type QStyleHintReturnMask__StyleOptionVersion int64

var (
	QStyleHintReturnMask__Version QStyleHintReturnMask__StyleOptionVersion = QStyleHintReturnMask__StyleOptionVersion(1)
)

//go:generate stringer -type=QStyleHintReturnVariant__StyleOptionType
//QStyleHintReturnVariant::StyleOptionType
type QStyleHintReturnVariant__StyleOptionType int64

var (
	QStyleHintReturnVariant__Type QStyleHintReturnVariant__StyleOptionType = QStyleHintReturnVariant__StyleOptionType(QStyleHintReturn__SH_Variant)
)

//go:generate stringer -type=QStyleHintReturnVariant__StyleOptionVersion
//QStyleHintReturnVariant::StyleOptionVersion
type QStyleHintReturnVariant__StyleOptionVersion int64

var (
	QStyleHintReturnVariant__Version QStyleHintReturnVariant__StyleOptionVersion = QStyleHintReturnVariant__StyleOptionVersion(1)
)

type QStyleOption struct {
	ptr unsafe.Pointer
}

type QStyleOption_ITF interface {
	QStyleOption_PTR() *QStyleOption
}

func (ptr *QStyleOption) QStyleOption_PTR() *QStyleOption {
	return ptr
}

func (ptr *QStyleOption) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QStyleOption) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQStyleOption(ptr QStyleOption_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOption_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionFromPointer(ptr unsafe.Pointer) (n *QStyleOption) {
	n = new(QStyleOption)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QStyleOption__OptionType
//QStyleOption::OptionType
type QStyleOption__OptionType int64

var (
	QStyleOption__SO_Default           QStyleOption__OptionType = QStyleOption__OptionType(0)
	QStyleOption__SO_FocusRect         QStyleOption__OptionType = QStyleOption__OptionType(1)
	QStyleOption__SO_Button            QStyleOption__OptionType = QStyleOption__OptionType(2)
	QStyleOption__SO_Tab               QStyleOption__OptionType = QStyleOption__OptionType(3)
	QStyleOption__SO_MenuItem          QStyleOption__OptionType = QStyleOption__OptionType(4)
	QStyleOption__SO_Frame             QStyleOption__OptionType = QStyleOption__OptionType(5)
	QStyleOption__SO_ProgressBar       QStyleOption__OptionType = QStyleOption__OptionType(6)
	QStyleOption__SO_ToolBox           QStyleOption__OptionType = QStyleOption__OptionType(7)
	QStyleOption__SO_Header            QStyleOption__OptionType = QStyleOption__OptionType(8)
	QStyleOption__SO_DockWidget        QStyleOption__OptionType = QStyleOption__OptionType(9)
	QStyleOption__SO_ViewItem          QStyleOption__OptionType = QStyleOption__OptionType(10)
	QStyleOption__SO_TabWidgetFrame    QStyleOption__OptionType = QStyleOption__OptionType(11)
	QStyleOption__SO_TabBarBase        QStyleOption__OptionType = QStyleOption__OptionType(12)
	QStyleOption__SO_RubberBand        QStyleOption__OptionType = QStyleOption__OptionType(13)
	QStyleOption__SO_ToolBar           QStyleOption__OptionType = QStyleOption__OptionType(14)
	QStyleOption__SO_GraphicsItem      QStyleOption__OptionType = QStyleOption__OptionType(15)
	QStyleOption__SO_Complex           QStyleOption__OptionType = QStyleOption__OptionType(0xf0000)
	QStyleOption__SO_Slider            QStyleOption__OptionType = QStyleOption__OptionType(C.QStyleOption_SO_Slider_Type())
	QStyleOption__SO_SpinBox           QStyleOption__OptionType = QStyleOption__OptionType(C.QStyleOption_SO_SpinBox_Type())
	QStyleOption__SO_ToolButton        QStyleOption__OptionType = QStyleOption__OptionType(C.QStyleOption_SO_ToolButton_Type())
	QStyleOption__SO_ComboBox          QStyleOption__OptionType = QStyleOption__OptionType(C.QStyleOption_SO_ComboBox_Type())
	QStyleOption__SO_TitleBar          QStyleOption__OptionType = QStyleOption__OptionType(C.QStyleOption_SO_TitleBar_Type())
	QStyleOption__SO_GroupBox          QStyleOption__OptionType = QStyleOption__OptionType(C.QStyleOption_SO_GroupBox_Type())
	QStyleOption__SO_SizeGrip          QStyleOption__OptionType = QStyleOption__OptionType(C.QStyleOption_SO_SizeGrip_Type())
	QStyleOption__SO_CustomBase        QStyleOption__OptionType = QStyleOption__OptionType(0xf00)
	QStyleOption__SO_ComplexCustomBase QStyleOption__OptionType = QStyleOption__OptionType(0xf000000)
)

//go:generate stringer -type=QStyleOption__StyleOptionType
//QStyleOption::StyleOptionType
type QStyleOption__StyleOptionType int64

var (
	QStyleOption__Type QStyleOption__StyleOptionType = QStyleOption__StyleOptionType(QStyleOption__SO_Default)
)

//go:generate stringer -type=QStyleOption__StyleOptionVersion
//QStyleOption::StyleOptionVersion
type QStyleOption__StyleOptionVersion int64

var (
	QStyleOption__Version QStyleOption__StyleOptionVersion = QStyleOption__StyleOptionVersion(1)
)

func NewQStyleOption2(other QStyleOption_ITF) *QStyleOption {
	tmpValue := NewQStyleOptionFromPointer(C.QStyleOption_NewQStyleOption2(PointerFromQStyleOption(other)))
	qt.SetFinalizer(tmpValue, (*QStyleOption).DestroyQStyleOption)
	return tmpValue
}

func NewQStyleOption(version int, ty int) *QStyleOption {
	tmpValue := NewQStyleOptionFromPointer(C.QStyleOption_NewQStyleOption(C.int(int32(version)), C.int(int32(ty))))
	qt.SetFinalizer(tmpValue, (*QStyleOption).DestroyQStyleOption)
	return tmpValue
}

func (ptr *QStyleOption) InitFrom(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QStyleOption_InitFrom(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QStyleOption) DestroyQStyleOption() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QStyleOption_DestroyQStyleOption(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QStyleOption) Type() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStyleOption_Type(ptr.Pointer())))
	}
	return 0
}

//go:generate stringer -type=QStyleOptionButton__ButtonFeature
//QStyleOptionButton::ButtonFeature
type QStyleOptionButton__ButtonFeature int64

const (
	QStyleOptionButton__None              QStyleOptionButton__ButtonFeature = QStyleOptionButton__ButtonFeature(0x00)
	QStyleOptionButton__Flat              QStyleOptionButton__ButtonFeature = QStyleOptionButton__ButtonFeature(0x01)
	QStyleOptionButton__HasMenu           QStyleOptionButton__ButtonFeature = QStyleOptionButton__ButtonFeature(0x02)
	QStyleOptionButton__DefaultButton     QStyleOptionButton__ButtonFeature = QStyleOptionButton__ButtonFeature(0x04)
	QStyleOptionButton__AutoDefaultButton QStyleOptionButton__ButtonFeature = QStyleOptionButton__ButtonFeature(0x08)
	QStyleOptionButton__CommandLinkButton QStyleOptionButton__ButtonFeature = QStyleOptionButton__ButtonFeature(0x10)
)

//go:generate stringer -type=QStyleOptionButton__StyleOptionType
//QStyleOptionButton::StyleOptionType
type QStyleOptionButton__StyleOptionType int64

var (
	QStyleOptionButton__Type QStyleOptionButton__StyleOptionType = QStyleOptionButton__StyleOptionType(QStyleOption__SO_Button)
)

//go:generate stringer -type=QStyleOptionButton__StyleOptionVersion
//QStyleOptionButton::StyleOptionVersion
type QStyleOptionButton__StyleOptionVersion int64

var (
	QStyleOptionButton__Version QStyleOptionButton__StyleOptionVersion = QStyleOptionButton__StyleOptionVersion(1)
)

//go:generate stringer -type=QStyleOptionComboBox__StyleOptionType
//QStyleOptionComboBox::StyleOptionType
type QStyleOptionComboBox__StyleOptionType int64

var (
	QStyleOptionComboBox__Type QStyleOptionComboBox__StyleOptionType = QStyleOptionComboBox__StyleOptionType(QStyleOption__SO_ComboBox)
)

//go:generate stringer -type=QStyleOptionComboBox__StyleOptionVersion
//QStyleOptionComboBox::StyleOptionVersion
type QStyleOptionComboBox__StyleOptionVersion int64

var (
	QStyleOptionComboBox__Version QStyleOptionComboBox__StyleOptionVersion = QStyleOptionComboBox__StyleOptionVersion(1)
)

//go:generate stringer -type=QStyleOptionComplex__StyleOptionType
//QStyleOptionComplex::StyleOptionType
type QStyleOptionComplex__StyleOptionType int64

var (
	QStyleOptionComplex__Type QStyleOptionComplex__StyleOptionType = QStyleOptionComplex__StyleOptionType(QStyleOption__SO_Complex)
)

//go:generate stringer -type=QStyleOptionComplex__StyleOptionVersion
//QStyleOptionComplex::StyleOptionVersion
type QStyleOptionComplex__StyleOptionVersion int64

var (
	QStyleOptionComplex__Version QStyleOptionComplex__StyleOptionVersion = QStyleOptionComplex__StyleOptionVersion(1)
)

//go:generate stringer -type=QStyleOptionDockWidget__StyleOptionType
//QStyleOptionDockWidget::StyleOptionType
type QStyleOptionDockWidget__StyleOptionType int64

var (
	QStyleOptionDockWidget__Type QStyleOptionDockWidget__StyleOptionType = QStyleOptionDockWidget__StyleOptionType(QStyleOption__SO_DockWidget)
)

//go:generate stringer -type=QStyleOptionDockWidget__StyleOptionVersion
//QStyleOptionDockWidget::StyleOptionVersion
type QStyleOptionDockWidget__StyleOptionVersion int64

var (
	QStyleOptionDockWidget__Version QStyleOptionDockWidget__StyleOptionVersion = QStyleOptionDockWidget__StyleOptionVersion(2)
)

//go:generate stringer -type=QStyleOptionFocusRect__StyleOptionType
//QStyleOptionFocusRect::StyleOptionType
type QStyleOptionFocusRect__StyleOptionType int64

var (
	QStyleOptionFocusRect__Type QStyleOptionFocusRect__StyleOptionType = QStyleOptionFocusRect__StyleOptionType(QStyleOption__SO_FocusRect)
)

//go:generate stringer -type=QStyleOptionFocusRect__StyleOptionVersion
//QStyleOptionFocusRect::StyleOptionVersion
type QStyleOptionFocusRect__StyleOptionVersion int64

var (
	QStyleOptionFocusRect__Version QStyleOptionFocusRect__StyleOptionVersion = QStyleOptionFocusRect__StyleOptionVersion(1)
)

//go:generate stringer -type=QStyleOptionFrame__FrameFeature
//QStyleOptionFrame::FrameFeature
type QStyleOptionFrame__FrameFeature int64

const (
	QStyleOptionFrame__None    QStyleOptionFrame__FrameFeature = QStyleOptionFrame__FrameFeature(0x00)
	QStyleOptionFrame__Flat    QStyleOptionFrame__FrameFeature = QStyleOptionFrame__FrameFeature(0x01)
	QStyleOptionFrame__Rounded QStyleOptionFrame__FrameFeature = QStyleOptionFrame__FrameFeature(0x02)
)

//go:generate stringer -type=QStyleOptionFrame__StyleOptionType
//QStyleOptionFrame::StyleOptionType
type QStyleOptionFrame__StyleOptionType int64

var (
	QStyleOptionFrame__Type QStyleOptionFrame__StyleOptionType = QStyleOptionFrame__StyleOptionType(QStyleOption__SO_Frame)
)

//go:generate stringer -type=QStyleOptionFrame__StyleOptionVersion
//QStyleOptionFrame::StyleOptionVersion
type QStyleOptionFrame__StyleOptionVersion int64

var (
	QStyleOptionFrame__Version QStyleOptionFrame__StyleOptionVersion = QStyleOptionFrame__StyleOptionVersion(3)
)

type QStyleOptionGraphicsItem struct {
	QStyleOption
}

type QStyleOptionGraphicsItem_ITF interface {
	QStyleOption_ITF
	QStyleOptionGraphicsItem_PTR() *QStyleOptionGraphicsItem
}

func (ptr *QStyleOptionGraphicsItem) QStyleOptionGraphicsItem_PTR() *QStyleOptionGraphicsItem {
	return ptr
}

func (ptr *QStyleOptionGraphicsItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOption_PTR().Pointer()
	}
	return nil
}

func (ptr *QStyleOptionGraphicsItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QStyleOption_PTR().SetPointer(p)
	}
}

func PointerFromQStyleOptionGraphicsItem(ptr QStyleOptionGraphicsItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionGraphicsItem_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionGraphicsItemFromPointer(ptr unsafe.Pointer) (n *QStyleOptionGraphicsItem) {
	n = new(QStyleOptionGraphicsItem)
	n.SetPointer(ptr)
	return
}
func (ptr *QStyleOptionGraphicsItem) DestroyQStyleOptionGraphicsItem() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QStyleOptionGraphicsItem__StyleOptionType
//QStyleOptionGraphicsItem::StyleOptionType
type QStyleOptionGraphicsItem__StyleOptionType int64

var (
	QStyleOptionGraphicsItem__Type QStyleOptionGraphicsItem__StyleOptionType = QStyleOptionGraphicsItem__StyleOptionType(QStyleOption__SO_GraphicsItem)
)

//go:generate stringer -type=QStyleOptionGraphicsItem__StyleOptionVersion
//QStyleOptionGraphicsItem::StyleOptionVersion
type QStyleOptionGraphicsItem__StyleOptionVersion int64

var (
	QStyleOptionGraphicsItem__Version QStyleOptionGraphicsItem__StyleOptionVersion = QStyleOptionGraphicsItem__StyleOptionVersion(1)
)

func NewQStyleOptionGraphicsItem() *QStyleOptionGraphicsItem {
	tmpValue := NewQStyleOptionGraphicsItemFromPointer(C.QStyleOptionGraphicsItem_NewQStyleOptionGraphicsItem())
	qt.SetFinalizer(tmpValue, (*QStyleOptionGraphicsItem).DestroyQStyleOptionGraphicsItem)
	return tmpValue
}

func NewQStyleOptionGraphicsItem2(other QStyleOptionGraphicsItem_ITF) *QStyleOptionGraphicsItem {
	tmpValue := NewQStyleOptionGraphicsItemFromPointer(C.QStyleOptionGraphicsItem_NewQStyleOptionGraphicsItem2(PointerFromQStyleOptionGraphicsItem(other)))
	qt.SetFinalizer(tmpValue, (*QStyleOptionGraphicsItem).DestroyQStyleOptionGraphicsItem)
	return tmpValue
}

//go:generate stringer -type=QStyleOptionGroupBox__StyleOptionType
//QStyleOptionGroupBox::StyleOptionType
type QStyleOptionGroupBox__StyleOptionType int64

var (
	QStyleOptionGroupBox__Type QStyleOptionGroupBox__StyleOptionType = QStyleOptionGroupBox__StyleOptionType(QStyleOption__SO_GroupBox)
)

//go:generate stringer -type=QStyleOptionGroupBox__StyleOptionVersion
//QStyleOptionGroupBox::StyleOptionVersion
type QStyleOptionGroupBox__StyleOptionVersion int64

var (
	QStyleOptionGroupBox__Version QStyleOptionGroupBox__StyleOptionVersion = QStyleOptionGroupBox__StyleOptionVersion(1)
)

//go:generate stringer -type=QStyleOptionHeader__SectionPosition
//QStyleOptionHeader::SectionPosition
type QStyleOptionHeader__SectionPosition int64

const (
	QStyleOptionHeader__Beginning      QStyleOptionHeader__SectionPosition = QStyleOptionHeader__SectionPosition(0)
	QStyleOptionHeader__Middle         QStyleOptionHeader__SectionPosition = QStyleOptionHeader__SectionPosition(1)
	QStyleOptionHeader__End            QStyleOptionHeader__SectionPosition = QStyleOptionHeader__SectionPosition(2)
	QStyleOptionHeader__OnlyOneSection QStyleOptionHeader__SectionPosition = QStyleOptionHeader__SectionPosition(3)
)

//go:generate stringer -type=QStyleOptionHeader__SelectedPosition
//QStyleOptionHeader::SelectedPosition
type QStyleOptionHeader__SelectedPosition int64

const (
	QStyleOptionHeader__NotAdjacent                QStyleOptionHeader__SelectedPosition = QStyleOptionHeader__SelectedPosition(0)
	QStyleOptionHeader__NextIsSelected             QStyleOptionHeader__SelectedPosition = QStyleOptionHeader__SelectedPosition(1)
	QStyleOptionHeader__PreviousIsSelected         QStyleOptionHeader__SelectedPosition = QStyleOptionHeader__SelectedPosition(2)
	QStyleOptionHeader__NextAndPreviousAreSelected QStyleOptionHeader__SelectedPosition = QStyleOptionHeader__SelectedPosition(3)
)

//go:generate stringer -type=QStyleOptionHeader__SortIndicator
//QStyleOptionHeader::SortIndicator
type QStyleOptionHeader__SortIndicator int64

const (
	QStyleOptionHeader__None     QStyleOptionHeader__SortIndicator = QStyleOptionHeader__SortIndicator(0)
	QStyleOptionHeader__SortUp   QStyleOptionHeader__SortIndicator = QStyleOptionHeader__SortIndicator(1)
	QStyleOptionHeader__SortDown QStyleOptionHeader__SortIndicator = QStyleOptionHeader__SortIndicator(2)
)

//go:generate stringer -type=QStyleOptionHeader__StyleOptionType
//QStyleOptionHeader::StyleOptionType
type QStyleOptionHeader__StyleOptionType int64

var (
	QStyleOptionHeader__Type QStyleOptionHeader__StyleOptionType = QStyleOptionHeader__StyleOptionType(QStyleOption__SO_Header)
)

//go:generate stringer -type=QStyleOptionHeader__StyleOptionVersion
//QStyleOptionHeader::StyleOptionVersion
type QStyleOptionHeader__StyleOptionVersion int64

var (
	QStyleOptionHeader__Version QStyleOptionHeader__StyleOptionVersion = QStyleOptionHeader__StyleOptionVersion(1)
)

//go:generate stringer -type=QStyleOptionMenuItem__CheckType
//QStyleOptionMenuItem::CheckType
type QStyleOptionMenuItem__CheckType int64

const (
	QStyleOptionMenuItem__NotCheckable QStyleOptionMenuItem__CheckType = QStyleOptionMenuItem__CheckType(0)
	QStyleOptionMenuItem__Exclusive    QStyleOptionMenuItem__CheckType = QStyleOptionMenuItem__CheckType(1)
	QStyleOptionMenuItem__NonExclusive QStyleOptionMenuItem__CheckType = QStyleOptionMenuItem__CheckType(2)
)

//go:generate stringer -type=QStyleOptionMenuItem__MenuItemType
//QStyleOptionMenuItem::MenuItemType
type QStyleOptionMenuItem__MenuItemType int64

const (
	QStyleOptionMenuItem__Normal      QStyleOptionMenuItem__MenuItemType = QStyleOptionMenuItem__MenuItemType(0)
	QStyleOptionMenuItem__DefaultItem QStyleOptionMenuItem__MenuItemType = QStyleOptionMenuItem__MenuItemType(1)
	QStyleOptionMenuItem__Separator   QStyleOptionMenuItem__MenuItemType = QStyleOptionMenuItem__MenuItemType(2)
	QStyleOptionMenuItem__SubMenu     QStyleOptionMenuItem__MenuItemType = QStyleOptionMenuItem__MenuItemType(3)
	QStyleOptionMenuItem__Scroller    QStyleOptionMenuItem__MenuItemType = QStyleOptionMenuItem__MenuItemType(4)
	QStyleOptionMenuItem__TearOff     QStyleOptionMenuItem__MenuItemType = QStyleOptionMenuItem__MenuItemType(5)
	QStyleOptionMenuItem__Margin      QStyleOptionMenuItem__MenuItemType = QStyleOptionMenuItem__MenuItemType(6)
	QStyleOptionMenuItem__EmptyArea   QStyleOptionMenuItem__MenuItemType = QStyleOptionMenuItem__MenuItemType(7)
)

//go:generate stringer -type=QStyleOptionMenuItem__StyleOptionType
//QStyleOptionMenuItem::StyleOptionType
type QStyleOptionMenuItem__StyleOptionType int64

var (
	QStyleOptionMenuItem__Type QStyleOptionMenuItem__StyleOptionType = QStyleOptionMenuItem__StyleOptionType(QStyleOption__SO_MenuItem)
)

//go:generate stringer -type=QStyleOptionMenuItem__StyleOptionVersion
//QStyleOptionMenuItem::StyleOptionVersion
type QStyleOptionMenuItem__StyleOptionVersion int64

var (
	QStyleOptionMenuItem__Version QStyleOptionMenuItem__StyleOptionVersion = QStyleOptionMenuItem__StyleOptionVersion(1)
)

//go:generate stringer -type=QStyleOptionProgressBar__StyleOptionType
//QStyleOptionProgressBar::StyleOptionType
type QStyleOptionProgressBar__StyleOptionType int64

var (
	QStyleOptionProgressBar__Type QStyleOptionProgressBar__StyleOptionType = QStyleOptionProgressBar__StyleOptionType(QStyleOption__SO_ProgressBar)
)

//go:generate stringer -type=QStyleOptionProgressBar__StyleOptionVersion
//QStyleOptionProgressBar::StyleOptionVersion
type QStyleOptionProgressBar__StyleOptionVersion int64

var (
	QStyleOptionProgressBar__Version QStyleOptionProgressBar__StyleOptionVersion = QStyleOptionProgressBar__StyleOptionVersion(2)
)

//go:generate stringer -type=QStyleOptionRubberBand__StyleOptionType
//QStyleOptionRubberBand::StyleOptionType
type QStyleOptionRubberBand__StyleOptionType int64

var (
	QStyleOptionRubberBand__Type QStyleOptionRubberBand__StyleOptionType = QStyleOptionRubberBand__StyleOptionType(QStyleOption__SO_RubberBand)
)

//go:generate stringer -type=QStyleOptionRubberBand__StyleOptionVersion
//QStyleOptionRubberBand::StyleOptionVersion
type QStyleOptionRubberBand__StyleOptionVersion int64

var (
	QStyleOptionRubberBand__Version QStyleOptionRubberBand__StyleOptionVersion = QStyleOptionRubberBand__StyleOptionVersion(1)
)

//go:generate stringer -type=QStyleOptionSizeGrip__StyleOptionType
//QStyleOptionSizeGrip::StyleOptionType
type QStyleOptionSizeGrip__StyleOptionType int64

var (
	QStyleOptionSizeGrip__Type QStyleOptionSizeGrip__StyleOptionType = QStyleOptionSizeGrip__StyleOptionType(QStyleOption__SO_SizeGrip)
)

//go:generate stringer -type=QStyleOptionSizeGrip__StyleOptionVersion
//QStyleOptionSizeGrip::StyleOptionVersion
type QStyleOptionSizeGrip__StyleOptionVersion int64

var (
	QStyleOptionSizeGrip__Version QStyleOptionSizeGrip__StyleOptionVersion = QStyleOptionSizeGrip__StyleOptionVersion(1)
)

//go:generate stringer -type=QStyleOptionSlider__StyleOptionType
//QStyleOptionSlider::StyleOptionType
type QStyleOptionSlider__StyleOptionType int64

var (
	QStyleOptionSlider__Type QStyleOptionSlider__StyleOptionType = QStyleOptionSlider__StyleOptionType(QStyleOption__SO_Slider)
)

//go:generate stringer -type=QStyleOptionSlider__StyleOptionVersion
//QStyleOptionSlider::StyleOptionVersion
type QStyleOptionSlider__StyleOptionVersion int64

var (
	QStyleOptionSlider__Version QStyleOptionSlider__StyleOptionVersion = QStyleOptionSlider__StyleOptionVersion(1)
)

//go:generate stringer -type=QStyleOptionSpinBox__StyleOptionType
//QStyleOptionSpinBox::StyleOptionType
type QStyleOptionSpinBox__StyleOptionType int64

var (
	QStyleOptionSpinBox__Type QStyleOptionSpinBox__StyleOptionType = QStyleOptionSpinBox__StyleOptionType(QStyleOption__SO_SpinBox)
)

//go:generate stringer -type=QStyleOptionSpinBox__StyleOptionVersion
//QStyleOptionSpinBox::StyleOptionVersion
type QStyleOptionSpinBox__StyleOptionVersion int64

var (
	QStyleOptionSpinBox__Version QStyleOptionSpinBox__StyleOptionVersion = QStyleOptionSpinBox__StyleOptionVersion(1)
)

//go:generate stringer -type=QStyleOptionTab__CornerWidget
//QStyleOptionTab::CornerWidget
type QStyleOptionTab__CornerWidget int64

const (
	QStyleOptionTab__NoCornerWidgets   QStyleOptionTab__CornerWidget = QStyleOptionTab__CornerWidget(0x00)
	QStyleOptionTab__LeftCornerWidget  QStyleOptionTab__CornerWidget = QStyleOptionTab__CornerWidget(0x01)
	QStyleOptionTab__RightCornerWidget QStyleOptionTab__CornerWidget = QStyleOptionTab__CornerWidget(0x02)
)

//go:generate stringer -type=QStyleOptionTab__SelectedPosition
//QStyleOptionTab::SelectedPosition
type QStyleOptionTab__SelectedPosition int64

const (
	QStyleOptionTab__NotAdjacent        QStyleOptionTab__SelectedPosition = QStyleOptionTab__SelectedPosition(0)
	QStyleOptionTab__NextIsSelected     QStyleOptionTab__SelectedPosition = QStyleOptionTab__SelectedPosition(1)
	QStyleOptionTab__PreviousIsSelected QStyleOptionTab__SelectedPosition = QStyleOptionTab__SelectedPosition(2)
)

//go:generate stringer -type=QStyleOptionTab__StyleOptionType
//QStyleOptionTab::StyleOptionType
type QStyleOptionTab__StyleOptionType int64

var (
	QStyleOptionTab__Type QStyleOptionTab__StyleOptionType = QStyleOptionTab__StyleOptionType(QStyleOption__SO_Tab)
)

//go:generate stringer -type=QStyleOptionTab__StyleOptionVersion
//QStyleOptionTab::StyleOptionVersion
type QStyleOptionTab__StyleOptionVersion int64

var (
	QStyleOptionTab__Version QStyleOptionTab__StyleOptionVersion = QStyleOptionTab__StyleOptionVersion(3)
)

//go:generate stringer -type=QStyleOptionTab__TabFeature
//QStyleOptionTab::TabFeature
type QStyleOptionTab__TabFeature int64

const (
	QStyleOptionTab__None     QStyleOptionTab__TabFeature = QStyleOptionTab__TabFeature(0x00)
	QStyleOptionTab__HasFrame QStyleOptionTab__TabFeature = QStyleOptionTab__TabFeature(0x01)
)

//go:generate stringer -type=QStyleOptionTab__TabPosition
//QStyleOptionTab::TabPosition
type QStyleOptionTab__TabPosition int64

const (
	QStyleOptionTab__Beginning  QStyleOptionTab__TabPosition = QStyleOptionTab__TabPosition(0)
	QStyleOptionTab__Middle     QStyleOptionTab__TabPosition = QStyleOptionTab__TabPosition(1)
	QStyleOptionTab__End        QStyleOptionTab__TabPosition = QStyleOptionTab__TabPosition(2)
	QStyleOptionTab__OnlyOneTab QStyleOptionTab__TabPosition = QStyleOptionTab__TabPosition(3)
)

//go:generate stringer -type=QStyleOptionTabBarBase__StyleOptionType
//QStyleOptionTabBarBase::StyleOptionType
type QStyleOptionTabBarBase__StyleOptionType int64

var (
	QStyleOptionTabBarBase__Type QStyleOptionTabBarBase__StyleOptionType = QStyleOptionTabBarBase__StyleOptionType(QStyleOption__SO_TabBarBase)
)

//go:generate stringer -type=QStyleOptionTabBarBase__StyleOptionVersion
//QStyleOptionTabBarBase::StyleOptionVersion
type QStyleOptionTabBarBase__StyleOptionVersion int64

var (
	QStyleOptionTabBarBase__Version QStyleOptionTabBarBase__StyleOptionVersion = QStyleOptionTabBarBase__StyleOptionVersion(2)
)

//go:generate stringer -type=QStyleOptionTabWidgetFrame__StyleOptionType
//QStyleOptionTabWidgetFrame::StyleOptionType
type QStyleOptionTabWidgetFrame__StyleOptionType int64

var (
	QStyleOptionTabWidgetFrame__Type QStyleOptionTabWidgetFrame__StyleOptionType = QStyleOptionTabWidgetFrame__StyleOptionType(QStyleOption__SO_TabWidgetFrame)
)

//go:generate stringer -type=QStyleOptionTabWidgetFrame__StyleOptionVersion
//QStyleOptionTabWidgetFrame::StyleOptionVersion
type QStyleOptionTabWidgetFrame__StyleOptionVersion int64

var (
	QStyleOptionTabWidgetFrame__Version QStyleOptionTabWidgetFrame__StyleOptionVersion = QStyleOptionTabWidgetFrame__StyleOptionVersion(2)
)

//go:generate stringer -type=QStyleOptionTitleBar__StyleOptionType
//QStyleOptionTitleBar::StyleOptionType
type QStyleOptionTitleBar__StyleOptionType int64

var (
	QStyleOptionTitleBar__Type QStyleOptionTitleBar__StyleOptionType = QStyleOptionTitleBar__StyleOptionType(QStyleOption__SO_TitleBar)
)

//go:generate stringer -type=QStyleOptionTitleBar__StyleOptionVersion
//QStyleOptionTitleBar::StyleOptionVersion
type QStyleOptionTitleBar__StyleOptionVersion int64

var (
	QStyleOptionTitleBar__Version QStyleOptionTitleBar__StyleOptionVersion = QStyleOptionTitleBar__StyleOptionVersion(1)
)

//go:generate stringer -type=QStyleOptionToolBar__StyleOptionType
//QStyleOptionToolBar::StyleOptionType
type QStyleOptionToolBar__StyleOptionType int64

var (
	QStyleOptionToolBar__Type QStyleOptionToolBar__StyleOptionType = QStyleOptionToolBar__StyleOptionType(QStyleOption__SO_ToolBar)
)

//go:generate stringer -type=QStyleOptionToolBar__StyleOptionVersion
//QStyleOptionToolBar::StyleOptionVersion
type QStyleOptionToolBar__StyleOptionVersion int64

var (
	QStyleOptionToolBar__Version QStyleOptionToolBar__StyleOptionVersion = QStyleOptionToolBar__StyleOptionVersion(1)
)

//go:generate stringer -type=QStyleOptionToolBar__ToolBarFeature
//QStyleOptionToolBar::ToolBarFeature
type QStyleOptionToolBar__ToolBarFeature int64

const (
	QStyleOptionToolBar__None    QStyleOptionToolBar__ToolBarFeature = QStyleOptionToolBar__ToolBarFeature(0x0)
	QStyleOptionToolBar__Movable QStyleOptionToolBar__ToolBarFeature = QStyleOptionToolBar__ToolBarFeature(0x1)
)

//go:generate stringer -type=QStyleOptionToolBar__ToolBarPosition
//QStyleOptionToolBar::ToolBarPosition
type QStyleOptionToolBar__ToolBarPosition int64

const (
	QStyleOptionToolBar__Beginning QStyleOptionToolBar__ToolBarPosition = QStyleOptionToolBar__ToolBarPosition(0)
	QStyleOptionToolBar__Middle    QStyleOptionToolBar__ToolBarPosition = QStyleOptionToolBar__ToolBarPosition(1)
	QStyleOptionToolBar__End       QStyleOptionToolBar__ToolBarPosition = QStyleOptionToolBar__ToolBarPosition(2)
	QStyleOptionToolBar__OnlyOne   QStyleOptionToolBar__ToolBarPosition = QStyleOptionToolBar__ToolBarPosition(3)
)

//go:generate stringer -type=QStyleOptionToolBox__SelectedPosition
//QStyleOptionToolBox::SelectedPosition
type QStyleOptionToolBox__SelectedPosition int64

const (
	QStyleOptionToolBox__NotAdjacent        QStyleOptionToolBox__SelectedPosition = QStyleOptionToolBox__SelectedPosition(0)
	QStyleOptionToolBox__NextIsSelected     QStyleOptionToolBox__SelectedPosition = QStyleOptionToolBox__SelectedPosition(1)
	QStyleOptionToolBox__PreviousIsSelected QStyleOptionToolBox__SelectedPosition = QStyleOptionToolBox__SelectedPosition(2)
)

//go:generate stringer -type=QStyleOptionToolBox__StyleOptionType
//QStyleOptionToolBox::StyleOptionType
type QStyleOptionToolBox__StyleOptionType int64

var (
	QStyleOptionToolBox__Type QStyleOptionToolBox__StyleOptionType = QStyleOptionToolBox__StyleOptionType(QStyleOption__SO_ToolBox)
)

//go:generate stringer -type=QStyleOptionToolBox__StyleOptionVersion
//QStyleOptionToolBox::StyleOptionVersion
type QStyleOptionToolBox__StyleOptionVersion int64

var (
	QStyleOptionToolBox__Version QStyleOptionToolBox__StyleOptionVersion = QStyleOptionToolBox__StyleOptionVersion(2)
)

//go:generate stringer -type=QStyleOptionToolBox__TabPosition
//QStyleOptionToolBox::TabPosition
type QStyleOptionToolBox__TabPosition int64

const (
	QStyleOptionToolBox__Beginning  QStyleOptionToolBox__TabPosition = QStyleOptionToolBox__TabPosition(0)
	QStyleOptionToolBox__Middle     QStyleOptionToolBox__TabPosition = QStyleOptionToolBox__TabPosition(1)
	QStyleOptionToolBox__End        QStyleOptionToolBox__TabPosition = QStyleOptionToolBox__TabPosition(2)
	QStyleOptionToolBox__OnlyOneTab QStyleOptionToolBox__TabPosition = QStyleOptionToolBox__TabPosition(3)
)

//go:generate stringer -type=QStyleOptionToolButton__StyleOptionType
//QStyleOptionToolButton::StyleOptionType
type QStyleOptionToolButton__StyleOptionType int64

var (
	QStyleOptionToolButton__Type QStyleOptionToolButton__StyleOptionType = QStyleOptionToolButton__StyleOptionType(QStyleOption__SO_ToolButton)
)

//go:generate stringer -type=QStyleOptionToolButton__StyleOptionVersion
//QStyleOptionToolButton::StyleOptionVersion
type QStyleOptionToolButton__StyleOptionVersion int64

var (
	QStyleOptionToolButton__Version QStyleOptionToolButton__StyleOptionVersion = QStyleOptionToolButton__StyleOptionVersion(1)
)

//go:generate stringer -type=QStyleOptionToolButton__ToolButtonFeature
//QStyleOptionToolButton::ToolButtonFeature
type QStyleOptionToolButton__ToolButtonFeature int64

const (
	QStyleOptionToolButton__None            QStyleOptionToolButton__ToolButtonFeature = QStyleOptionToolButton__ToolButtonFeature(0x00)
	QStyleOptionToolButton__Arrow           QStyleOptionToolButton__ToolButtonFeature = QStyleOptionToolButton__ToolButtonFeature(0x01)
	QStyleOptionToolButton__Menu            QStyleOptionToolButton__ToolButtonFeature = QStyleOptionToolButton__ToolButtonFeature(0x04)
	QStyleOptionToolButton__MenuButtonPopup QStyleOptionToolButton__ToolButtonFeature = QStyleOptionToolButton__ToolButtonFeature(QStyleOptionToolButton__Menu)
	QStyleOptionToolButton__PopupDelay      QStyleOptionToolButton__ToolButtonFeature = QStyleOptionToolButton__ToolButtonFeature(0x08)
	QStyleOptionToolButton__HasMenu         QStyleOptionToolButton__ToolButtonFeature = QStyleOptionToolButton__ToolButtonFeature(0x10)
)

//go:generate stringer -type=QStyleOptionViewItem__Position
//QStyleOptionViewItem::Position
type QStyleOptionViewItem__Position int64

const (
	QStyleOptionViewItem__Left   QStyleOptionViewItem__Position = QStyleOptionViewItem__Position(0)
	QStyleOptionViewItem__Right  QStyleOptionViewItem__Position = QStyleOptionViewItem__Position(1)
	QStyleOptionViewItem__Top    QStyleOptionViewItem__Position = QStyleOptionViewItem__Position(2)
	QStyleOptionViewItem__Bottom QStyleOptionViewItem__Position = QStyleOptionViewItem__Position(3)
)

//go:generate stringer -type=QStyleOptionViewItem__StyleOptionType
//QStyleOptionViewItem::StyleOptionType
type QStyleOptionViewItem__StyleOptionType int64

var (
	QStyleOptionViewItem__Type QStyleOptionViewItem__StyleOptionType = QStyleOptionViewItem__StyleOptionType(QStyleOption__SO_ViewItem)
)

//go:generate stringer -type=QStyleOptionViewItem__StyleOptionVersion
//QStyleOptionViewItem::StyleOptionVersion
type QStyleOptionViewItem__StyleOptionVersion int64

var (
	QStyleOptionViewItem__Version QStyleOptionViewItem__StyleOptionVersion = QStyleOptionViewItem__StyleOptionVersion(4)
)

//go:generate stringer -type=QStyleOptionViewItem__ViewItemFeature
//QStyleOptionViewItem::ViewItemFeature
type QStyleOptionViewItem__ViewItemFeature int64

const (
	QStyleOptionViewItem__None              QStyleOptionViewItem__ViewItemFeature = QStyleOptionViewItem__ViewItemFeature(0x00)
	QStyleOptionViewItem__WrapText          QStyleOptionViewItem__ViewItemFeature = QStyleOptionViewItem__ViewItemFeature(0x01)
	QStyleOptionViewItem__Alternate         QStyleOptionViewItem__ViewItemFeature = QStyleOptionViewItem__ViewItemFeature(0x02)
	QStyleOptionViewItem__HasCheckIndicator QStyleOptionViewItem__ViewItemFeature = QStyleOptionViewItem__ViewItemFeature(0x04)
	QStyleOptionViewItem__HasDisplay        QStyleOptionViewItem__ViewItemFeature = QStyleOptionViewItem__ViewItemFeature(0x08)
	QStyleOptionViewItem__HasDecoration     QStyleOptionViewItem__ViewItemFeature = QStyleOptionViewItem__ViewItemFeature(0x10)
)

//go:generate stringer -type=QStyleOptionViewItem__ViewItemPosition
//QStyleOptionViewItem::ViewItemPosition
type QStyleOptionViewItem__ViewItemPosition int64

const (
	QStyleOptionViewItem__Invalid   QStyleOptionViewItem__ViewItemPosition = QStyleOptionViewItem__ViewItemPosition(0)
	QStyleOptionViewItem__Beginning QStyleOptionViewItem__ViewItemPosition = QStyleOptionViewItem__ViewItemPosition(1)
	QStyleOptionViewItem__Middle    QStyleOptionViewItem__ViewItemPosition = QStyleOptionViewItem__ViewItemPosition(2)
	QStyleOptionViewItem__End       QStyleOptionViewItem__ViewItemPosition = QStyleOptionViewItem__ViewItemPosition(3)
	QStyleOptionViewItem__OnlyOne   QStyleOptionViewItem__ViewItemPosition = QStyleOptionViewItem__ViewItemPosition(4)
)

//go:generate stringer -type=QSwipeGesture__SwipeDirection
//QSwipeGesture::SwipeDirection
type QSwipeGesture__SwipeDirection int64

const (
	QSwipeGesture__NoDirection QSwipeGesture__SwipeDirection = QSwipeGesture__SwipeDirection(0)
	QSwipeGesture__Left        QSwipeGesture__SwipeDirection = QSwipeGesture__SwipeDirection(1)
	QSwipeGesture__Right       QSwipeGesture__SwipeDirection = QSwipeGesture__SwipeDirection(2)
	QSwipeGesture__Up          QSwipeGesture__SwipeDirection = QSwipeGesture__SwipeDirection(3)
	QSwipeGesture__Down        QSwipeGesture__SwipeDirection = QSwipeGesture__SwipeDirection(4)
)

//go:generate stringer -type=QSystemTrayIcon__ActivationReason
//QSystemTrayIcon::ActivationReason
type QSystemTrayIcon__ActivationReason int64

const (
	QSystemTrayIcon__Unknown     QSystemTrayIcon__ActivationReason = QSystemTrayIcon__ActivationReason(0)
	QSystemTrayIcon__Context     QSystemTrayIcon__ActivationReason = QSystemTrayIcon__ActivationReason(1)
	QSystemTrayIcon__DoubleClick QSystemTrayIcon__ActivationReason = QSystemTrayIcon__ActivationReason(2)
	QSystemTrayIcon__Trigger     QSystemTrayIcon__ActivationReason = QSystemTrayIcon__ActivationReason(3)
	QSystemTrayIcon__MiddleClick QSystemTrayIcon__ActivationReason = QSystemTrayIcon__ActivationReason(4)
)

//go:generate stringer -type=QSystemTrayIcon__MessageIcon
//QSystemTrayIcon::MessageIcon
type QSystemTrayIcon__MessageIcon int64

const (
	QSystemTrayIcon__NoIcon      QSystemTrayIcon__MessageIcon = QSystemTrayIcon__MessageIcon(0)
	QSystemTrayIcon__Information QSystemTrayIcon__MessageIcon = QSystemTrayIcon__MessageIcon(1)
	QSystemTrayIcon__Warning     QSystemTrayIcon__MessageIcon = QSystemTrayIcon__MessageIcon(2)
	QSystemTrayIcon__Critical    QSystemTrayIcon__MessageIcon = QSystemTrayIcon__MessageIcon(3)
)

//go:generate stringer -type=QTabBar__ButtonPosition
//QTabBar::ButtonPosition
type QTabBar__ButtonPosition int64

const (
	QTabBar__LeftSide  QTabBar__ButtonPosition = QTabBar__ButtonPosition(0)
	QTabBar__RightSide QTabBar__ButtonPosition = QTabBar__ButtonPosition(1)
)

//go:generate stringer -type=QTabBar__SelectionBehavior
//QTabBar::SelectionBehavior
type QTabBar__SelectionBehavior int64

const (
	QTabBar__SelectLeftTab     QTabBar__SelectionBehavior = QTabBar__SelectionBehavior(0)
	QTabBar__SelectRightTab    QTabBar__SelectionBehavior = QTabBar__SelectionBehavior(1)
	QTabBar__SelectPreviousTab QTabBar__SelectionBehavior = QTabBar__SelectionBehavior(2)
)

//go:generate stringer -type=QTabBar__Shape
//QTabBar::Shape
type QTabBar__Shape int64

const (
	QTabBar__RoundedNorth    QTabBar__Shape = QTabBar__Shape(0)
	QTabBar__RoundedSouth    QTabBar__Shape = QTabBar__Shape(1)
	QTabBar__RoundedWest     QTabBar__Shape = QTabBar__Shape(2)
	QTabBar__RoundedEast     QTabBar__Shape = QTabBar__Shape(3)
	QTabBar__TriangularNorth QTabBar__Shape = QTabBar__Shape(4)
	QTabBar__TriangularSouth QTabBar__Shape = QTabBar__Shape(5)
	QTabBar__TriangularWest  QTabBar__Shape = QTabBar__Shape(6)
	QTabBar__TriangularEast  QTabBar__Shape = QTabBar__Shape(7)
)

//go:generate stringer -type=QTabWidget__TabPosition
//QTabWidget::TabPosition
type QTabWidget__TabPosition int64

const (
	QTabWidget__North QTabWidget__TabPosition = QTabWidget__TabPosition(0)
	QTabWidget__South QTabWidget__TabPosition = QTabWidget__TabPosition(1)
	QTabWidget__West  QTabWidget__TabPosition = QTabWidget__TabPosition(2)
	QTabWidget__East  QTabWidget__TabPosition = QTabWidget__TabPosition(3)
)

//go:generate stringer -type=QTabWidget__TabShape
//QTabWidget::TabShape
type QTabWidget__TabShape int64

const (
	QTabWidget__Rounded    QTabWidget__TabShape = QTabWidget__TabShape(0)
	QTabWidget__Triangular QTabWidget__TabShape = QTabWidget__TabShape(1)
)

//go:generate stringer -type=QTableWidgetItem__ItemType
//QTableWidgetItem::ItemType
type QTableWidgetItem__ItemType int64

const (
	QTableWidgetItem__Type     QTableWidgetItem__ItemType = QTableWidgetItem__ItemType(0)
	QTableWidgetItem__UserType QTableWidgetItem__ItemType = QTableWidgetItem__ItemType(1000)
)

type QTextEdit struct {
	QAbstractScrollArea
}

type QTextEdit_ITF interface {
	QAbstractScrollArea_ITF
	QTextEdit_PTR() *QTextEdit
}

func (ptr *QTextEdit) QTextEdit_PTR() *QTextEdit {
	return ptr
}

func (ptr *QTextEdit) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractScrollArea_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextEdit) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractScrollArea_PTR().SetPointer(p)
	}
}

func PointerFromQTextEdit(ptr QTextEdit_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextEdit_PTR().Pointer()
	}
	return nil
}

func NewQTextEditFromPointer(ptr unsafe.Pointer) (n *QTextEdit) {
	n = new(QTextEdit)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QTextEdit__AutoFormattingFlag
//QTextEdit::AutoFormattingFlag
type QTextEdit__AutoFormattingFlag int64

const (
	QTextEdit__AutoNone       QTextEdit__AutoFormattingFlag = QTextEdit__AutoFormattingFlag(0)
	QTextEdit__AutoBulletList QTextEdit__AutoFormattingFlag = QTextEdit__AutoFormattingFlag(0x00000001)
	QTextEdit__AutoAll        QTextEdit__AutoFormattingFlag = QTextEdit__AutoFormattingFlag(0xffffffff)
)

//go:generate stringer -type=QTextEdit__LineWrapMode
//QTextEdit::LineWrapMode
type QTextEdit__LineWrapMode int64

const (
	QTextEdit__NoWrap           QTextEdit__LineWrapMode = QTextEdit__LineWrapMode(0)
	QTextEdit__WidgetWidth      QTextEdit__LineWrapMode = QTextEdit__LineWrapMode(1)
	QTextEdit__FixedPixelWidth  QTextEdit__LineWrapMode = QTextEdit__LineWrapMode(2)
	QTextEdit__FixedColumnWidth QTextEdit__LineWrapMode = QTextEdit__LineWrapMode(3)
)

func NewQTextEdit(parent QWidget_ITF) *QTextEdit {
	tmpValue := NewQTextEditFromPointer(C.QTextEdit_NewQTextEdit(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQTextEdit2(text string, parent QWidget_ITF) *QTextEdit {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewQTextEditFromPointer(C.QTextEdit_NewQTextEdit2(C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQTextEdit_Copy
func callbackQTextEdit_Copy(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "copy"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTextEditFromPointer(ptr).CopyDefault()
	}
}

func (ptr *QTextEdit) ConnectCopy(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "copy"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "copy", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "copy", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextEdit) DisconnectCopy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "copy")
	}
}

func (ptr *QTextEdit) Copy() {
	if ptr.Pointer() != nil {
		C.QTextEdit_Copy(ptr.Pointer())
	}
}

func (ptr *QTextEdit) CopyDefault() {
	if ptr.Pointer() != nil {
		C.QTextEdit_CopyDefault(ptr.Pointer())
	}
}

//export callbackQTextEdit_SetText
func callbackQTextEdit_SetText(ptr unsafe.Pointer, text C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "setText"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(text))
	} else {
		NewQTextEditFromPointer(ptr).SetTextDefault(cGoUnpackString(text))
	}
}

func (ptr *QTextEdit) ConnectSetText(f func(text string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setText"); signal != nil {
			f := func(text string) {
				(*(*func(string))(signal))(text)
				f(text)
			}
			qt.ConnectSignal(ptr.Pointer(), "setText", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setText", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextEdit) DisconnectSetText() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setText")
	}
}

func (ptr *QTextEdit) SetText(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QTextEdit_SetText(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QTextEdit) SetTextDefault(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QTextEdit_SetTextDefault(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

//export callbackQTextEdit_DestroyQTextEdit
func callbackQTextEdit_DestroyQTextEdit(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QTextEdit"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTextEditFromPointer(ptr).DestroyQTextEditDefault()
	}
}

func (ptr *QTextEdit) ConnectDestroyQTextEdit(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QTextEdit"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QTextEdit", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QTextEdit", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextEdit) DisconnectDestroyQTextEdit() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QTextEdit")
	}
}

func (ptr *QTextEdit) DestroyQTextEdit() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextEdit_DestroyQTextEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextEdit) DestroyQTextEditDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextEdit_DestroyQTextEditDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextEdit) ToPlainText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTextEdit_ToPlainText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextEdit) Print(printer gui.QPagedPaintDevice_ITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_Print(ptr.Pointer(), gui.PointerFromQPagedPaintDevice(printer))
	}
}

//go:generate stringer -type=QToolButton__ToolButtonPopupMode
//QToolButton::ToolButtonPopupMode
type QToolButton__ToolButtonPopupMode int64

const (
	QToolButton__DelayedPopup    QToolButton__ToolButtonPopupMode = QToolButton__ToolButtonPopupMode(0)
	QToolButton__MenuButtonPopup QToolButton__ToolButtonPopupMode = QToolButton__ToolButtonPopupMode(1)
	QToolButton__InstantPopup    QToolButton__ToolButtonPopupMode = QToolButton__ToolButtonPopupMode(2)
)

//go:generate stringer -type=QTreeWidgetItem__ChildIndicatorPolicy
//QTreeWidgetItem::ChildIndicatorPolicy
type QTreeWidgetItem__ChildIndicatorPolicy int64

const (
	QTreeWidgetItem__ShowIndicator                  QTreeWidgetItem__ChildIndicatorPolicy = QTreeWidgetItem__ChildIndicatorPolicy(0)
	QTreeWidgetItem__DontShowIndicator              QTreeWidgetItem__ChildIndicatorPolicy = QTreeWidgetItem__ChildIndicatorPolicy(1)
	QTreeWidgetItem__DontShowIndicatorWhenChildless QTreeWidgetItem__ChildIndicatorPolicy = QTreeWidgetItem__ChildIndicatorPolicy(2)
)

//go:generate stringer -type=QTreeWidgetItem__ItemType
//QTreeWidgetItem::ItemType
type QTreeWidgetItem__ItemType int64

const (
	QTreeWidgetItem__Type     QTreeWidgetItem__ItemType = QTreeWidgetItem__ItemType(0)
	QTreeWidgetItem__UserType QTreeWidgetItem__ItemType = QTreeWidgetItem__ItemType(1000)
)

//go:generate stringer -type=QTreeWidgetItemIterator__IteratorFlag
//QTreeWidgetItemIterator::IteratorFlag
type QTreeWidgetItemIterator__IteratorFlag int64

const (
	QTreeWidgetItemIterator__All           QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000000)
	QTreeWidgetItemIterator__Hidden        QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000001)
	QTreeWidgetItemIterator__NotHidden     QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000002)
	QTreeWidgetItemIterator__Selected      QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000004)
	QTreeWidgetItemIterator__Unselected    QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000008)
	QTreeWidgetItemIterator__Selectable    QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000010)
	QTreeWidgetItemIterator__NotSelectable QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000020)
	QTreeWidgetItemIterator__DragEnabled   QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000040)
	QTreeWidgetItemIterator__DragDisabled  QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000080)
	QTreeWidgetItemIterator__DropEnabled   QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000100)
	QTreeWidgetItemIterator__DropDisabled  QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000200)
	QTreeWidgetItemIterator__HasChildren   QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000400)
	QTreeWidgetItemIterator__NoChildren    QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000800)
	QTreeWidgetItemIterator__Checked       QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00001000)
	QTreeWidgetItemIterator__NotChecked    QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00002000)
	QTreeWidgetItemIterator__Enabled       QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00004000)
	QTreeWidgetItemIterator__Disabled      QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00008000)
	QTreeWidgetItemIterator__Editable      QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00010000)
	QTreeWidgetItemIterator__NotEditable   QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00020000)
	QTreeWidgetItemIterator__UserFlag      QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x01000000)
)

type QVBoxLayout struct {
	QBoxLayout
}

type QVBoxLayout_ITF interface {
	QBoxLayout_ITF
	QVBoxLayout_PTR() *QVBoxLayout
}

func (ptr *QVBoxLayout) QVBoxLayout_PTR() *QVBoxLayout {
	return ptr
}

func (ptr *QVBoxLayout) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QBoxLayout_PTR().Pointer()
	}
	return nil
}

func (ptr *QVBoxLayout) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QBoxLayout_PTR().SetPointer(p)
	}
}

func PointerFromQVBoxLayout(ptr QVBoxLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVBoxLayout_PTR().Pointer()
	}
	return nil
}

func NewQVBoxLayoutFromPointer(ptr unsafe.Pointer) (n *QVBoxLayout) {
	n = new(QVBoxLayout)
	n.SetPointer(ptr)
	return
}
func NewQVBoxLayout() *QVBoxLayout {
	tmpValue := NewQVBoxLayoutFromPointer(C.QVBoxLayout_NewQVBoxLayout())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQVBoxLayout2(parent QWidget_ITF) *QVBoxLayout {
	tmpValue := NewQVBoxLayoutFromPointer(C.QVBoxLayout_NewQVBoxLayout2(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QVBoxLayout) DestroyQVBoxLayout() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QVBoxLayout_DestroyQVBoxLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QWidget struct {
	core.QObject
	gui.QPaintDevice
}

type QWidget_ITF interface {
	core.QObject_ITF
	gui.QPaintDevice_ITF
	QWidget_PTR() *QWidget
}

func (ptr *QWidget) QWidget_PTR() *QWidget {
	return ptr
}

func (ptr *QWidget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWidget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QPaintDevice_PTR().SetPointer(p)
	}
}

func PointerFromQWidget(ptr QWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func NewQWidgetFromPointer(ptr unsafe.Pointer) (n *QWidget) {
	n = new(QWidget)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QWidget__RenderFlag
//QWidget::RenderFlag
type QWidget__RenderFlag int64

const (
	QWidget__DrawWindowBackground QWidget__RenderFlag = QWidget__RenderFlag(0x1)
	QWidget__DrawChildren         QWidget__RenderFlag = QWidget__RenderFlag(0x2)
	QWidget__IgnoreMask           QWidget__RenderFlag = QWidget__RenderFlag(0x4)
)

func NewQWidget(parent QWidget_ITF, ff core.Qt__WindowType) *QWidget {
	tmpValue := NewQWidgetFromPointer(C.QWidget_NewQWidget(PointerFromQWidget(parent), C.longlong(ff)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWidget_Close
func callbackQWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QWidget) ConnectClose(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "close"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "close", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "close", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "close")
	}
}

func (ptr *QWidget) Close() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWidget_Close(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWidget_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) Create(window uintptr, initializeWindow bool, destroyOldWindow bool) {
	if ptr.Pointer() != nil {
		C.QWidget_Create(ptr.Pointer(), C.uintptr_t(window), C.char(int8(qt.GoBoolToInt(initializeWindow))), C.char(int8(qt.GoBoolToInt(destroyOldWindow))))
	}
}

//export callbackQWidget_Lower
func callbackQWidget_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QWidget) ConnectLower(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "lower"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "lower", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lower", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectLower() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "lower")
	}
}

func (ptr *QWidget) Lower() {
	if ptr.Pointer() != nil {
		C.QWidget_Lower(ptr.Pointer())
	}
}

func (ptr *QWidget) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QWidget_LowerDefault(ptr.Pointer())
	}
}

func (ptr *QWidget) SetContentsMargins2(margins core.QMargins_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetContentsMargins2(ptr.Pointer(), core.PointerFromQMargins(margins))
	}
}

func (ptr *QWidget) SetContentsMargins(left int, top int, right int, bottom int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetContentsMargins(ptr.Pointer(), C.int(int32(left)), C.int(int32(top)), C.int(int32(right)), C.int(int32(bottom)))
	}
}

//export callbackQWidget_SetEnabled
func callbackQWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewQWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QWidget) ConnectSetEnabled(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setEnabled"); signal != nil {
			f := func(vbo bool) {
				(*(*func(bool))(signal))(vbo)
				f(vbo)
			}
			qt.ConnectSignal(ptr.Pointer(), "setEnabled", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setEnabled", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectSetEnabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setEnabled")
	}
}

func (ptr *QWidget) SetEnabled(vbo bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QWidget) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QWidget) SetMinimumSize(vqs core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMinimumSize(ptr.Pointer(), core.PointerFromQSize(vqs))
	}
}

func (ptr *QWidget) SetMinimumSize2(minw int, minh int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMinimumSize2(ptr.Pointer(), C.int(int32(minw)), C.int(int32(minh)))
	}
}

//export callbackQWidget_SetWindowTitle
func callbackQWidget_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	} else {
		NewQWidgetFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QWidget) ConnectSetWindowTitle(f func(vqs string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setWindowTitle"); signal != nil {
			f := func(vqs string) {
				(*(*func(string))(signal))(vqs)
				f(vqs)
			}
			qt.ConnectSignal(ptr.Pointer(), "setWindowTitle", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setWindowTitle", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectSetWindowTitle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setWindowTitle")
	}
}

func (ptr *QWidget) SetWindowTitle(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QWidget_SetWindowTitle(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

func (ptr *QWidget) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QWidget_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQWidget_Show
func callbackQWidget_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QWidget) ConnectShow(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "show"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "show", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "show", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectShow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "show")
	}
}

func (ptr *QWidget) Show() {
	if ptr.Pointer() != nil {
		C.QWidget_Show(ptr.Pointer())
	}
}

func (ptr *QWidget) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QWidget_ShowDefault(ptr.Pointer())
	}
}

func (ptr *QWidget) DestroyQWidget() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWidget_DestroyQWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWidget) Layout() *QLayout {
	if ptr.Pointer() != nil {
		tmpValue := NewQLayoutFromPointer(C.QWidget_Layout(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) ContentsMargins() *core.QMargins {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQMarginsFromPointer(C.QWidget_ContentsMargins(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QMargins).DestroyQMargins)
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) MapTo(parent QWidget_ITF, pos core.QPoint_ITF) *core.QPoint {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFromPointer(C.QWidget_MapTo(ptr.Pointer(), PointerFromQWidget(parent), core.PointerFromQPoint(pos)))
		qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) Pos() *core.QPoint {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFromPointer(C.QWidget_Pos(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) MinimumSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QWidget_MinimumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) Size() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QWidget_Size(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) WindowTitle() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWidget_WindowTitle(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) Window() *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QWidget_Window(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) X() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWidget_X(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) Y() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWidget_Y(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) __addActions_actions_atList(i int) *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QWidget___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) __addActions_actions_setList(i QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget___addActions_actions_setList(ptr.Pointer(), PointerFromQAction(i))
	}
}

func (ptr *QWidget) __addActions_actions_newList() unsafe.Pointer {
	return C.QWidget___addActions_actions_newList(ptr.Pointer())
}

func (ptr *QWidget) __insertActions_actions_atList(i int) *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QWidget___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) __insertActions_actions_setList(i QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget___insertActions_actions_setList(ptr.Pointer(), PointerFromQAction(i))
	}
}

func (ptr *QWidget) __insertActions_actions_newList() unsafe.Pointer {
	return C.QWidget___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *QWidget) __actions_atList(i int) *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QWidget___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) __actions_setList(i QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget___actions_setList(ptr.Pointer(), PointerFromQAction(i))
	}
}

func (ptr *QWidget) __actions_newList() unsafe.Pointer {
	return C.QWidget___actions_newList(ptr.Pointer())
}

func (ptr *QWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWidget___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWidget) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWidget___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWidget) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWidget___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWidget) __findChildren_newList2() unsafe.Pointer {
	return C.QWidget___findChildren_newList2(ptr.Pointer())
}

func (ptr *QWidget) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWidget___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWidget) __findChildren_newList3() unsafe.Pointer {
	return C.QWidget___findChildren_newList3(ptr.Pointer())
}

func (ptr *QWidget) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWidget___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWidget) __findChildren_newList() unsafe.Pointer {
	return C.QWidget___findChildren_newList(ptr.Pointer())
}

func (ptr *QWidget) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWidget___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWidget) __children_newList() unsafe.Pointer {
	return C.QWidget___children_newList(ptr.Pointer())
}

//export callbackQWidget_ObjectNameChanged
func callbackQWidget_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQWidget_PaintEngine
func callbackQWidget_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine((*(*func() *gui.QPaintEngine)(signal))())
	}

	return gui.PointerFromQPaintEngine(NewQWidgetFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QWidget) PaintEngine() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QWidget_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QWidget_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//go:generate stringer -type=QWizard__WizardButton
//QWizard::WizardButton
type QWizard__WizardButton int64

const (
	QWizard__BackButton       QWizard__WizardButton = QWizard__WizardButton(0)
	QWizard__NextButton       QWizard__WizardButton = QWizard__WizardButton(1)
	QWizard__CommitButton     QWizard__WizardButton = QWizard__WizardButton(2)
	QWizard__FinishButton     QWizard__WizardButton = QWizard__WizardButton(3)
	QWizard__CancelButton     QWizard__WizardButton = QWizard__WizardButton(4)
	QWizard__HelpButton       QWizard__WizardButton = QWizard__WizardButton(5)
	QWizard__CustomButton1    QWizard__WizardButton = QWizard__WizardButton(6)
	QWizard__CustomButton2    QWizard__WizardButton = QWizard__WizardButton(7)
	QWizard__CustomButton3    QWizard__WizardButton = QWizard__WizardButton(8)
	QWizard__Stretch          QWizard__WizardButton = QWizard__WizardButton(9)
	QWizard__NoButton         QWizard__WizardButton = QWizard__WizardButton(-1)
	QWizard__NStandardButtons QWizard__WizardButton = QWizard__WizardButton(6)
	QWizard__NButtons         QWizard__WizardButton = QWizard__WizardButton(9)
)

//go:generate stringer -type=QWizard__WizardOption
//QWizard::WizardOption
type QWizard__WizardOption int64

const (
	QWizard__IndependentPages             QWizard__WizardOption = QWizard__WizardOption(0x00000001)
	QWizard__IgnoreSubTitles              QWizard__WizardOption = QWizard__WizardOption(0x00000002)
	QWizard__ExtendedWatermarkPixmap      QWizard__WizardOption = QWizard__WizardOption(0x00000004)
	QWizard__NoDefaultButton              QWizard__WizardOption = QWizard__WizardOption(0x00000008)
	QWizard__NoBackButtonOnStartPage      QWizard__WizardOption = QWizard__WizardOption(0x00000010)
	QWizard__NoBackButtonOnLastPage       QWizard__WizardOption = QWizard__WizardOption(0x00000020)
	QWizard__DisabledBackButtonOnLastPage QWizard__WizardOption = QWizard__WizardOption(0x00000040)
	QWizard__HaveNextButtonOnLastPage     QWizard__WizardOption = QWizard__WizardOption(0x00000080)
	QWizard__HaveFinishButtonOnEarlyPages QWizard__WizardOption = QWizard__WizardOption(0x00000100)
	QWizard__NoCancelButton               QWizard__WizardOption = QWizard__WizardOption(0x00000200)
	QWizard__CancelButtonOnLeft           QWizard__WizardOption = QWizard__WizardOption(0x00000400)
	QWizard__HaveHelpButton               QWizard__WizardOption = QWizard__WizardOption(0x00000800)
	QWizard__HelpButtonOnRight            QWizard__WizardOption = QWizard__WizardOption(0x00001000)
	QWizard__HaveCustomButton1            QWizard__WizardOption = QWizard__WizardOption(0x00002000)
	QWizard__HaveCustomButton2            QWizard__WizardOption = QWizard__WizardOption(0x00004000)
	QWizard__HaveCustomButton3            QWizard__WizardOption = QWizard__WizardOption(0x00008000)
	QWizard__NoCancelButtonOnLastPage     QWizard__WizardOption = QWizard__WizardOption(0x00010000)
)

//go:generate stringer -type=QWizard__WizardPixmap
//QWizard::WizardPixmap
type QWizard__WizardPixmap int64

const (
	QWizard__WatermarkPixmap  QWizard__WizardPixmap = QWizard__WizardPixmap(0)
	QWizard__LogoPixmap       QWizard__WizardPixmap = QWizard__WizardPixmap(1)
	QWizard__BannerPixmap     QWizard__WizardPixmap = QWizard__WizardPixmap(2)
	QWizard__BackgroundPixmap QWizard__WizardPixmap = QWizard__WizardPixmap(3)
	QWizard__NPixmaps         QWizard__WizardPixmap = QWizard__WizardPixmap(4)
)

//go:generate stringer -type=QWizard__WizardStyle
//QWizard::WizardStyle
type QWizard__WizardStyle int64

var (
	QWizard__ClassicStyle QWizard__WizardStyle = QWizard__WizardStyle(0)
	QWizard__ModernStyle  QWizard__WizardStyle = QWizard__WizardStyle(1)
	QWizard__MacStyle     QWizard__WizardStyle = QWizard__WizardStyle(2)
	QWizard__AeroStyle    QWizard__WizardStyle = QWizard__WizardStyle(3)
	QWizard__NStyles      QWizard__WizardStyle = QWizard__WizardStyle(4)
)

func init() {
}
