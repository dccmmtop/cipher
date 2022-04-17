package uigen

import (
	"cipher/logger"
	"cipher/passutil"
)

func (form *UIWindowsForm) GenClick(checked bool) {

	form.Label2.SetText("正在生成...")
	passutil.GenRsaKey()
	form.Content.SetText("您的公钥和私钥已经生成")
	form.Label2.SetText("完成")
}

func (form *UIWindowsForm) EncryptClick(checked bool) {
	logger.Logger.Debug("加密")
	form.Label2.SetText("加密中...")
	text := form.Content.ToPlainText()
	text = passutil.Encrypt(text)
	form.Content.SetText(text)
	form.Label2.SetText("加密完成")
	logger.Logger.Debug("加密完成，密文: %s\n",text)
}

func (form *UIWindowsForm) DecryptClick(checked bool) {
	text := form.Content.ToPlainText()
	logger.Logger.Debug("密文是: %s\n",text)
	text = passutil.Decrypt(text)
	form.Content.SetText(text)
}
