SET PATH=%PATH%;%GOPATH%\bin;

REM 配置QT目录和版本
SET QT_DIR=D:\Qt5.9.1
SET QT_VERSION_MAJOR=5.9.1

REM 编译32位程序
SET GOARCH=386

REM 编译具体项目时运行
qtdeploy build desktop ./main.go
