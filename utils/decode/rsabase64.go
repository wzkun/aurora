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

	"golang.org/x/crypto/bcrypt"
)

// StdEncodeToString base64编码
func StdEncodeToString(src []byte) string {
	dst := base64.StdEncoding.EncodeToString(src)
	return dst
}

// StdDecodeString base64解码
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

		encoded := StdEncodeToString(signature)
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

	encoded := StdEncodeToString(signature)

	return encoded, nil
}

// EncodeRSA RSA+BASE64 encode
func EncodeRSAWithKeyFile(filename string, data []byte) (string, error) {
	pubkey, err := GetPubKey(filename)
	if err != nil {
		return "", err
	}

	signature, err := rsa.EncryptPKCS1v15(rand.Reader, pubkey, data)
	if err != nil {
		return "", err
	}

	encoded := StdEncodeToString(signature)

	return encoded, nil
}

// DecodeRSA  RSA+BASE64 decode
func DecodeRSA(prikey *rsa.PrivateKey, data string) ([]byte, error) {
	decoded, err := StdDecodeString(data)
	if err != nil {
		return nil, err
	}

	dsignature, err := rsa.DecryptPKCS1v15(rand.Reader, prikey, []byte(decoded))
	if err != nil {
		return nil, err
	}

	return dsignature, nil
}

// DecodeRSA  RSA+BASE64 decode
func DecodeRSAWithKeyFile(filename, data string) ([]byte, error) {
	prikey, err := GetPriKey(filename)
	if err != nil {
		return nil, err
	}

	decoded, err := StdDecodeString(data)
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

//GetPubKeyLen 获取RSA公钥
func GetPubKeyLen(filename string) (int, error) {
	pub, err := GetPubKey(filename)
	if err != nil {
		return 0, err
	}
	return pub.N.BitLen(), nil
}

//GetPriKeyLen 获取RSA私钥长度
func GetPriKeyLen(filename string) (int, error) {
	prikey, err := GetPriKey(filename)
	if err != nil {
		return 0, err
	}

	return prikey.N.BitLen(), nil
}

// SaltedValue return salted hash value.
func SaltedValue(value string) (string, error) {
	salt, err := bcrypt.GenerateFromPassword([]byte(value), 10)
	if err != nil {
		return "", err
	}

	encodePW := string(salt)
	return encodePW, nil
}

// Verfy function.
func Verify(password, hpassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hpassword))
	if err != nil {
		return false, err
	}
	return true, nil
}
