package cryption

import "testing"

func Test_encrypt(t *testing.T) {
	thb := ChannelMap["THB"]
	data := "31271652780605295"
	sign, err := thb.ReqRSASign(data)
	t.Log("\n sign:", sign, "err:", err)

	result, err := thb.RespRSAVerify(data, sign)

	t.Log("\n result:", result, "err:", err)
}

func Test_encryptData(t *testing.T) {
	thb := ChannelMap["THB"]
	data := "202212141157"
	err := thb.RespRSATest(data)
	t.Log("err:", err)
}
