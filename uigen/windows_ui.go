// WARNING! All changes made in this file will be lost!
package uigen

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type UIWindowsForm struct {
	VerticalLayout4 *widgets.QVBoxLayout
	GroupBox *widgets.QGroupBox
	HorizontalLayout2 *widgets.QHBoxLayout
	GroupBox4 *widgets.QGroupBox
	VerticalLayout2 *widgets.QVBoxLayout
	Label2 *widgets.QLabel
	Content *widgets.QTextEdit
	GroupBox2 *widgets.QGroupBox
	VerticalLayout *widgets.QVBoxLayout
	Gen *widgets.QPushButton
	Encrypt *widgets.QPushButton
	Decrypt *widgets.QPushButton
}

func (this *UIWindowsForm) SetupUI(Form *widgets.QWidget) {
	Form.SetObjectName("Form")
	this.VerticalLayout4 = widgets.NewQVBoxLayout2(Form)
	this.VerticalLayout4.SetObjectName("verticalLayout_4")
	this.VerticalLayout4.SetContentsMargins(0, 0, 0, 0)
	this.VerticalLayout4.SetSpacing(0)
	this.GroupBox = widgets.NewQGroupBox(Form)
	this.GroupBox.SetObjectName("GroupBox")
	this.GroupBox.SetMinimumSize(core.NewQSize2(500, 500))
	this.HorizontalLayout2 = widgets.NewQHBoxLayout2(this.GroupBox)
	this.HorizontalLayout2.SetObjectName("horizontalLayout_2")
	this.HorizontalLayout2.SetContentsMargins(0, 0, 0, 0)
	this.HorizontalLayout2.SetSpacing(0)
	this.GroupBox4 = widgets.NewQGroupBox(this.GroupBox)
	this.GroupBox4.SetObjectName("GroupBox4")
	this.VerticalLayout2 = widgets.NewQVBoxLayout2(this.GroupBox4)
	this.VerticalLayout2.SetObjectName("verticalLayout_2")
	this.VerticalLayout2.SetContentsMargins(0, 0, 0, 0)
	this.VerticalLayout2.SetSpacing(0)
	this.Label2 = widgets.NewQLabel(this.GroupBox4, core.Qt__Widget)
	this.Label2.SetObjectName("Label2")
	this.VerticalLayout2.AddWidget(this.Label2, 0, 0)
	this.Content = widgets.NewQTextEdit(this.GroupBox4)
	this.Content.SetObjectName("Content")
	this.Content.SetEnabled(true)
	this.Content.SetMinimumSize(core.NewQSize2(0, 500))
	this.VerticalLayout2.AddWidget(this.Content, 0, 0)
	this.HorizontalLayout2.AddWidget(this.GroupBox4, 0, 0)
	this.GroupBox2 = widgets.NewQGroupBox(this.GroupBox)
	this.GroupBox2.SetObjectName("GroupBox2")
	this.VerticalLayout = widgets.NewQVBoxLayout2(this.GroupBox2)
	this.VerticalLayout.SetObjectName("verticalLayout")
	this.VerticalLayout.SetContentsMargins(0, 0, 0, 0)
	this.VerticalLayout.SetSpacing(0)
	this.Gen = widgets.NewQPushButton(this.GroupBox2)
	this.Gen.SetObjectName("Gen")
	this.VerticalLayout.AddWidget(this.Gen, 0, 0)
	this.Encrypt = widgets.NewQPushButton(this.GroupBox2)
	this.Encrypt.SetObjectName("Encrypt")
	this.VerticalLayout.AddWidget(this.Encrypt, 0, 0)
	this.Decrypt = widgets.NewQPushButton(this.GroupBox2)
	this.Decrypt.SetObjectName("Decrypt")
	this.VerticalLayout.AddWidget(this.Decrypt, 0, 0)
	this.HorizontalLayout2.AddWidget(this.GroupBox2, 0, 0)
	this.VerticalLayout4.AddWidget(this.GroupBox, 0, 0)


    this.RetranslateUi(Form)

	this.Gen.ConnectClicked(this.GenClick)
	this.Decrypt.ConnectClicked(this.DecryptClick)
	this.Encrypt.ConnectClicked(this.EncryptClick)
}

func (this *UIWindowsForm) RetranslateUi(Form *widgets.QWidget) {
    _translate := core.QCoreApplication_Translate
	Form.SetWindowTitle(_translate("Form", "Form", "", -1))
	this.GroupBox.SetTitle(_translate("Form", "", "", -1))
	this.GroupBox4.SetTitle(_translate("Form", "", "", -1))
	this.Label2.SetText(_translate("Form", "明文", "", -1))
	this.GroupBox2.SetTitle(_translate("Form", "", "", -1))
	this.Gen.SetText(_translate("Form", "生成密钥", "", -1))
	this.Encrypt.SetText(_translate("Form", "加密", "", -1))
	this.Decrypt.SetText(_translate("Form", "解密", "", -1))
}
