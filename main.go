package main

import (
	"cipher/uigen"
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	widget := widgets.NewQApplication(len(os.Args), os.Args)

	mainwindows := widgets.NewQWidget(nil, 0) //建立空白的 widget
	form := uigen.UIWindowsForm{}
	// 绘制界面
	form.SetupUI(mainwindows)
	// 事件绑定
	form.BindEvent()
	// 显示
	mainwindows.Show()
	widget.Exec()

}
