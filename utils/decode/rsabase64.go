package decode

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
)

// 编码
func StdEncodeToString(src []byte) string {
	dst := base64.StdEncoding.EncodeToString(src)
	return dst
}

// 解码
func StdDecodeString(src string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(src)
	return decoded, err
}

// EncodeRSAs RSA+BASE64 encode
func EncodeRSAs(pubkey *rsa.PublicKey, datas [][]byte) ([]string, error) {
	var resp []string
	for _, data := range datas {
		signature, err := rsa.EncryptPKCS1v15(rand.Reader, pubkey, data)
		if err != nil {
			return nil, err
		}

		encoded := base64.StdEncoding.EncodeToString(signature)
		resp = append(resp, encoded)
	}

	return resp, nil
}

// EncodeRSA RSA+BASE64 encode
func EncodeRSA(pubkey *rsa.PublicKey, data []byte) (string, error) {
	signature, err := rsa.EncryptPKCS1v15(rand.Reader, pubkey, data)
	if err != nil {
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(signature)

	return encoded, nil
}

// DecodeRSA  RSA+BASE64 decode
func DecodeRSA(prikey *rsa.PrivateKey, data string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	dsignature, err := rsa.DecryptPKCS1v15(rand.Reader, prikey, []byte(decoded))
	if err != nil {
		return nil, err
	}

	return dsignature, nil
}

// GetPubKey function
func GetPubKey(filename string) (*rsa.PublicKey, error) {
	PubKey, err := ioutil.ReadFile(filename)
	if err != nil {
		os.Exit(-1)
	}

	if PubKey == nil {
		return nil, errors.New("input arguments error")
	}

	block, _ := pem.Decode(PubKey)
	if block == nil {
		return nil, errors.New("public rsaKey error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)

	return pub, nil
}

// GetPriKey function.
func GetPriKey(filename string) (*rsa.PrivateKey, error) {
	PriKey, err := ioutil.ReadFile(filename)
	if err != nil {
		os.Exit(-1)
	}
	if PriKey == nil {
		return nil, errors.New("input arguments error")
	}

	block, _ := pem.Decode(PriKey)
	if block == nil {
		return nil, errors.New("private rsaKey error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

// GenRsaKey function.
func GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("config/pri.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("config/pub.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}
