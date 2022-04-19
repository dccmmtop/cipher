package test

import (
	"cipher/passutil"
	"testing"
)

func TestGenKey(t *testing.T){
	t.Skip("skip")
	 passutil.GenRsaKey()
}


func TestEncrypt(t *testing.T){

	text := passutil.Encrypt("你好呀!",[]string{"F:\\code\\go\\Qt\\cipher\\test\\公钥\\public_我.pem","F:\\code\\go\\Qt\\cipher\\test\\公钥\\public_张三.pem"})
	t.Log(len(text))
	t.Log(text)
}

func TestDecrypt(t *testing.T){
	text := "4934bf26adb5c497c186834aa1e6686d59fadd23912769892acebd8c93895f37a73d44333bb68e6f73426564519dc55a85c151b41baee79c28c39866bf972f38464d90d5a17341ef50ac86df3af37e464fc316ddcf4124960d94d6e93168407440e580f814cfc8687c1a75f9de8dd512ae72f2125950c3eb9d8825954d3a3d8b" +
		"|+d@c#m$m+|" +
		"6d916dde2d797a8cbaf5309f9fc9431e9424d96146bc8395397ea1a282cab54be01a9bb51f3084c93474d9e6eaec65baef817bdeb67b9141fee1fae51be7ca1f3acd674fb5e82242d5cd2bcc81e7dfd3ddeccec0436a56f692e9bea7f5f58bc42e14daa114525446b1ed0161b543b9c172223380aed7f53f9159799efe5acb09" +
		"-|=w-a-n=|-" +
		"utcgh/0fujchk3wtSwDBjQ=="
	decrypt, err := passutil.Decrypt(text)
	if err != nil {
		t.Log(err)
		t.Failed()
	}
	if decrypt != "你好呀!" {
		t.Log(decrypt)
		t.Failed()
	}
}