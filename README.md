使用非对称加密方法实现文字的加解密
### 编译环境:
- Qt 5.9.1
- golang 1.16

### 编译方式:
`./build.bat`

### 运行方式
执行 deploy/windows/cipher.exe


### 进展
目前已经实现基于单个固定公钥加密,下一步实现支持多选公钥加密

### 备注
ui 转 go 方法
`go get -u -v github.com/stephenlyu/goqtuic`

`goqtuic -ui-file .\windows.ui`

