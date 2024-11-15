package cryption

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	"github.com/pkg/errors"
)

// 兼容多商户模式
// 回调地址为同一个.请求根据语言进行渠道选择
// map[string]any : map[渠道ID对应的货币code]any

var (
	ChannelMap = map[string]*ConfInfo{
		"INR": {
			PUBLIC_KEY: []byte(`-----BEGIN RSA PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAKtt9fbGwO/YXTS7uArPxu1XbvVu0hbQ
MrjC7iflNl+5sXTmZa0VdbHXfpa6GgEBs7LxTYziHpwFNN1xCJgDu1UCAwEAAQ==
-----END RSA PUBLIC KEY-----`),
			PRIVATE_KEY: []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIBVAIBADANBgkqhkiG9w0BAQEFAASCAT4wggE6AgEAAkEAq2319sbA79hdNLu4
Cs/G7Vdu9W7SFtAyuMLuJ+U2X7mxdOZlrRV1sdd+lroaAQGzsvFNjOIenAU03XEI
mAO7VQIDAQABAkBvxDlgsySHOyFJiDntkpm3lBGboq0GgqcPHsf26tIQhgeoJurC
+H3Ecb7Xd4gnAS9wmIw3bVZ5bUgZ6SfrTssBAiEA0yJ7HLzeA9ck7vaO3qOHfa0E
x9r3I8YbnmOBZKI3uQ0CIQDP2475437PSlT61xuPinoHU9wmHtsgkJTiZ21aT6zp
aQIgUm5s6IxoldFCrzV+FDh+ZrfNXQYFZWcsU0uAMj0PSmUCIEKh3r6quA8ZhA55
gKNOP/ozXuJ1h8dzsda2Uy7dVc/pAiEAuDrRVx67j+k6fZBiHSJigyO46FaiOskJ
1y2jg1oDc2k=
-----END RSA PRIVATE KEY-----`),
		},
		"THB": {
			PUBLIC_KEY: []byte(`-----BEGIN RSA PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAItgrSos7d81/pOjEtYAixouuTOBO49V
YdG9uLIvopWaOyr13bp9hc+z3DRfKvDN3GqAXXcF6DQJeH4kx7AQNrUCAwEAAQ==
-----END RSA PUBLIC KEY-----`),
			PRIVATE_KEY: []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIBVQIBADANBgkqhkiG9w0BAQEFAASCAT8wggE7AgEAAkEAi2CtKizt3zX+k6MS
1gCLGi65M4E7j1Vh0b24si+ilZo7KvXdun2Fz7PcNF8q8M3caoBddwXoNAl4fiTH
sBA2tQIDAQABAkA3p1XeBmzBeXT7DN3DmBnzTX7kQxE+TbCuqJwsgnrgU0rCuDTO
1HGkYatBkdzy62qlmHracVkLQyOItqalHWKtAiEA06QUVAwRZYRFGGgrzu8aZlQj
79QEr3Jo13RnkFcZNTcCIQColzLa6lkQqfBjLEOm7va9f/3xn7dCv4g+PE2r/0Sp
cwIhAKmTW5JoRPwhOjevKALbhDyfiFfeeo6VTjnsFBsMHq+LAiBsh3Eghk4xsQRk
A2HFu/2w/g0vF5804/HEBrQFrj24SwIhAI+sesaW3xG7ADp42GVx+goTYoixVQrL
IBLEGCTk1rad
-----END RSA PRIVATE KEY-----`),
		},
	}
)

type ConfInfo struct {
	PUBLIC_KEY  []byte // 渠道ID的公钥
	PRIVATE_KEY []byte // 渠道ID的私钥
}

// 请求rsa签名
func (e *ConfInfo) ReqRSASign(data string) (string, error) {
	msgHash := md5.New()
	_, err := msgHash.Write([]byte(data))
	if err != nil {
		return "", err
	}
	msgHashSum := msgHash.Sum(nil)

	priPem, _ := pem.Decode(e.PRIVATE_KEY)
	if priPem == nil {
		return "", errors.New("空私钥")
	}

	priInterface, err := x509.ParsePKCS8PrivateKey(priPem.Bytes)
	if err != nil {
		return "", errors.Wrap(err, "解析私钥")
	}

	privateKey := priInterface.(*rsa.PrivateKey)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.MD5, msgHashSum)
	if err != nil {
		return "", errors.Wrap(err, "生成rsa签名失败")
	}

	sign := base64.StdEncoding.EncodeToString(signature)

	return sign, nil
}

// 响应rsa签名校验
func (e *ConfInfo) RespRSAVerify(data, signature string) (bool, error) {
	msgHash := md5.New()
	_, err := msgHash.Write([]byte(data))
	if err != nil {
		return false, err
	}
	msgHashSum := msgHash.Sum(nil)

	pubPem, _ := pem.Decode(e.PUBLIC_KEY)
	if pubPem == nil {
		return false, errors.New("空公钥")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(pubPem.Bytes)
	if err != nil {
		return false, err
	}

	publicKey := pubInterface.(*rsa.PublicKey)

	sign, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}

	err = rsa.VerifyPKCS1v15(publicKey, crypto.MD5, msgHashSum, sign)
	if err != nil {
		return false, errors.Wrap(err, "校验rsa签名失败")
	}

	return true, nil
}

// 响应rsa签名加解密数据
func (e *ConfInfo) RespRSATest(data string) error {
	pubPem, _ := pem.Decode(e.PUBLIC_KEY)
	if pubPem == nil {
		return errors.New("空公钥")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(pubPem.Bytes)
	if err != nil {
		return err
	}

	publicKey := pubInterface.(*rsa.PublicKey)

	signature, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(data))
	if err != nil {
		return err
	}
	signStr := base64.StdEncoding.EncodeToString(signature)
	fmt.Println("加密字符串", signStr)

	sign, err := base64.StdEncoding.DecodeString(signStr)
	if err != nil {
		return err
	}

	priPem, _ := pem.Decode(e.PRIVATE_KEY)
	if priPem == nil {
		return errors.New("空私钥")
	}

	priInterface, err := x509.ParsePKCS8PrivateKey(priPem.Bytes)
	if err != nil {
		return errors.Wrap(err, "解析私钥")
	}

	privateKey := priInterface.(*rsa.PrivateKey)

	result, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, sign)
	if err != nil {
		return errors.Wrap(err, "解密字符串")
	}

	fmt.Println("结果", string(result))

	return nil
}
