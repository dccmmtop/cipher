package main

import (
	"cipher/uigen"
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	widget := widgets.NewQApplication(len(os.Args), os.Args)

	mainwindows := widgets.NewQWidget(nil, 0) //建立空白的 widget
	loginUI := uigen.UIWindowsForm{}
	loginUI.SetupUI(mainwindows) // 将UI初始化给创建的mainwindows
	mainwindows.Show()           // 显示

	widget.Exec()

}
