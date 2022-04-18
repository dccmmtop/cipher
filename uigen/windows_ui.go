// WARNING! All changes made in this file will be lost!
package uigen

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type UIWindowsForm struct {
	HorizontalLayout  *widgets.QHBoxLayout
	GroupBox5         *widgets.QGroupBox
	HorizontalLayout3 *widgets.QHBoxLayout
	GroupBox          *widgets.QGroupBox
	VerticalLayout2   *widgets.QVBoxLayout
	Label2            *widgets.QLabel
	Content           *widgets.QTextEdit
	GroupBox2         *widgets.QGroupBox
	VerticalLayout    *widgets.QVBoxLayout
	Decrypt           *widgets.QPushButton
	Gen               *widgets.QPushButton
	GroupBox3         *widgets.QGroupBox
	VerticalLayout3   *widgets.QVBoxLayout
	Label             *widgets.QLabel
	PublicKeyList     *widgets.QTextEdit
	GroupBox4         *widgets.QGroupBox
	HorizontalLayout2 *widgets.QHBoxLayout
	SelectPubKey      *widgets.QPushButton
	ClearPubKey       *widgets.QPushButton
	Encrypt           *widgets.QPushButton
}

func (this *UIWindowsForm) SetupUI(Form *widgets.QWidget) {
	Form.SetObjectName("Form")
	Form.SetEnabled(true)
	this.HorizontalLayout = widgets.NewQHBoxLayout2(Form)
	this.HorizontalLayout.SetObjectName("horizontalLayout")
	this.HorizontalLayout.SetContentsMargins(0, 0, 0, 0)
	this.HorizontalLayout.SetSpacing(0)
	this.GroupBox5 = widgets.NewQGroupBox(Form)
	this.GroupBox5.SetObjectName("GroupBox5")
	this.HorizontalLayout3 = widgets.NewQHBoxLayout2(this.GroupBox5)
	this.HorizontalLayout3.SetObjectName("horizontalLayout_3")
	this.HorizontalLayout3.SetContentsMargins(0, 0, 0, 0)
	this.HorizontalLayout3.SetSpacing(0)
	this.GroupBox = widgets.NewQGroupBox(this.GroupBox5)
	this.GroupBox.SetObjectName("GroupBox")
	this.GroupBox.SetMinimumSize(core.NewQSize2(300, 0))
	this.VerticalLayout2 = widgets.NewQVBoxLayout2(this.GroupBox)
	this.VerticalLayout2.SetObjectName("verticalLayout_2")
	this.VerticalLayout2.SetContentsMargins(0, 0, 0, 0)
	this.VerticalLayout2.SetSpacing(0)
	this.Label2 = widgets.NewQLabel(this.GroupBox, core.Qt__Widget)
	this.Label2.SetObjectName("Label2")
	this.VerticalLayout2.AddWidget(this.Label2, 0, 0)
	this.Content = widgets.NewQTextEdit(this.GroupBox)
	this.Content.SetObjectName("Content")
	this.Content.SetEnabled(true)
	this.Content.SetMinimumSize(core.NewQSize2(0, 500))
	this.Content.SetReadOnly(false)
	this.VerticalLayout2.AddWidget(this.Content, 0, 0)
	this.HorizontalLayout3.AddWidget(this.GroupBox, 0, 0)
	this.GroupBox2 = widgets.NewQGroupBox(this.GroupBox5)
	this.GroupBox2.SetObjectName("GroupBox2")
	this.VerticalLayout = widgets.NewQVBoxLayout2(this.GroupBox2)
	this.VerticalLayout.SetObjectName("verticalLayout")
	this.VerticalLayout.SetContentsMargins(0, 0, 0, 0)
	this.VerticalLayout.SetSpacing(0)
	this.Decrypt = widgets.NewQPushButton(this.GroupBox2)
	this.Decrypt.SetObjectName("Decrypt")
	this.VerticalLayout.AddWidget(this.Decrypt, 0, 0)
	this.Gen = widgets.NewQPushButton(this.GroupBox2)
	this.Gen.SetObjectName("Gen")
	this.VerticalLayout.AddWidget(this.Gen, 0, 0)
	this.HorizontalLayout3.AddWidget(this.GroupBox2, 0, 0)
	this.GroupBox3 = widgets.NewQGroupBox(this.GroupBox5)
	this.GroupBox3.SetObjectName("GroupBox3")
	this.GroupBox3.SetMinimumSize(core.NewQSize2(300, 0))
	this.GroupBox3.SetMaximumSize(core.NewQSize2(300, 16777215))
	this.VerticalLayout3 = widgets.NewQVBoxLayout2(this.GroupBox3)
	this.VerticalLayout3.SetObjectName("verticalLayout_3")
	this.VerticalLayout3.SetContentsMargins(0, 0, 0, 0)
	this.VerticalLayout3.SetSpacing(0)
	this.Label = widgets.NewQLabel(this.GroupBox3, core.Qt__Widget)
	this.Label.SetObjectName("Label")
	this.VerticalLayout3.AddWidget(this.Label, 0, 0)
	this.PublicKeyList = widgets.NewQTextEdit(this.GroupBox3)
	this.PublicKeyList.SetObjectName("PublicKeyList")
	this.PublicKeyList.SetEnabled(true)
	this.PublicKeyList.SetMinimumSize(core.NewQSize2(0, 400))
	this.PublicKeyList.SetReadOnly(true)
	this.VerticalLayout3.AddWidget(this.PublicKeyList, 0, 0)
	this.GroupBox4 = widgets.NewQGroupBox(this.GroupBox3)
	this.GroupBox4.SetObjectName("GroupBox4")
	this.GroupBox4.SetMinimumSize(core.NewQSize2(60, 0))
	this.HorizontalLayout2 = widgets.NewQHBoxLayout2(this.GroupBox4)
	this.HorizontalLayout2.SetObjectName("horizontalLayout_2")
	this.HorizontalLayout2.SetContentsMargins(0, 0, 0, 0)
	this.HorizontalLayout2.SetSpacing(0)
	this.SelectPubKey = widgets.NewQPushButton(this.GroupBox4)
	this.SelectPubKey.SetObjectName("SelectPubKey")
	this.HorizontalLayout2.AddWidget(this.SelectPubKey, 0, 0)
	this.ClearPubKey = widgets.NewQPushButton(this.GroupBox4)
	this.ClearPubKey.SetObjectName("ClearPubKey")
	this.HorizontalLayout2.AddWidget(this.ClearPubKey, 0, 0)
	this.Encrypt = widgets.NewQPushButton(this.GroupBox4)
	this.Encrypt.SetObjectName("Encrypt")
	this.HorizontalLayout2.AddWidget(this.Encrypt, 0, 0)
	this.VerticalLayout3.AddWidget(this.GroupBox4, 0, 0)
	this.HorizontalLayout3.AddWidget(this.GroupBox3, 0, 0)
	this.HorizontalLayout.AddWidget(this.GroupBox5, 0, 0)

	this.RetranslateUi(Form)

}

func (this *UIWindowsForm) RetranslateUi(Form *widgets.QWidget) {
	_translate := core.QCoreApplication_Translate
	Form.SetWindowTitle(_translate("Form", "Form", "", -1))
	this.GroupBox5.SetTitle(_translate("Form", "", "", -1))
	this.GroupBox.SetTitle(_translate("Form", "", "", -1))
	this.Label2.SetText(_translate("Form", "明文", "", -1))
	this.GroupBox2.SetTitle(_translate("Form", "", "", -1))
	this.Decrypt.SetText(_translate("Form", "解密", "", -1))
	this.Gen.SetText(_translate("Form", "生成密钥", "", -1))
	this.GroupBox3.SetTitle(_translate("Form", "", "", -1))
	this.Label.SetText(_translate("Form", "能解密的用户", "", -1))
	this.GroupBox4.SetTitle(_translate("Form", "", "", -1))
	this.SelectPubKey.SetText(_translate("Form", "选择用户", "", -1))
	this.ClearPubKey.SetText(_translate("Form", "清空", "", -1))
	this.Encrypt.SetText(_translate("Form", "加密", "", -1))
}
