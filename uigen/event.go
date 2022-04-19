package uigen

import (
	"cipher/logger"
	"cipher/passutil"
	"github.com/therecipe/qt/widgets"
	"strings"
)

func (form *UIWindowsForm) BindEvent() {

	form.Encrypt.ConnectClicked(form.EncryptClick)
	form.Gen.ConnectClicked(form.GenClick)
	form.Decrypt.ConnectClicked(form.DecryptClick)
	form.ClearPubKey.ConnectClicked(form.ClearPubKeyClick)
	form.SelectPubKey.ConnectClicked(form.SelectPubKeyClick)
}
func (form *UIWindowsForm) GenClick(checked bool) {

	form.Label2.SetText("正在生成...")
	passutil.GenRsaKey()
	form.Content.SetText("您的公钥和私钥已经生成")
	form.Label2.SetText("完成")
}

func (form *UIWindowsForm) EncryptClick(checked bool) {
	if len(PubKeys) == 0 {
		logger.Logger.Info("没有选择公钥")
		return
	}
	logger.Logger.Debug("加密")
	form.Label2.SetText("加密中...")
	text := form.Content.ToPlainText()
	// 加密，多公钥
	text = passutil.Encrypt(text,PubKeys)
	form.Content.SetText(text)
	// 密文回显
	form.Label2.SetText("加密完成")
	logger.Logger.Debug("加密完成，密文: ", text)
}

func (form *UIWindowsForm) DecryptClick(checked bool) {
	text := form.Content.ToPlainText()
	logger.Logger.Debug("密文是: %s\n", text)
	text ,err := passutil.Decrypt(text)
	if err != nil {
		form.Label2.SetText(err.Error())
		return
	}
	form.Content.SetText(text)

}

func (form *UIWindowsForm) ClearPubKeyClick(checked bool) {
	PubKeys = nil
	form.PublicKeyList.SetText("")
}

func (form *UIWindowsForm) SelectPubKeyClick(checked bool) {
	// 选择文件
	keys := widgets.QFileDialog_GetOpenFileNames(nil, "", "./", "*.pem", "", widgets.QFileDialog__ShowDirsOnly)
	for _, key := range keys {
		if !isInclude(PubKeys, key) {
			PubKeys = append(PubKeys, key)
		}
	}
	logger.Logger.Debug("选择了公钥: %v\n", PubKeys)
	// 将文件回显
	form.PublicKeyList.SetText(strings.Join(PubKeys, "\n"))
}

func isInclude(ss []string, s string) bool {
	size := len(ss)
	for i := 0; i < size; i++ {
		if ss[i] == s {
			return true
		}
	}
	return false

}
