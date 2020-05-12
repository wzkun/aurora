package decode

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
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
func EncodeRSAs(pubkey *rsa.PublicKey, datas ...interface{}) ([]string, error) {
	var resp []string
	for _, data := range datas {
		signature, err := rsa.EncryptPKCS1v15(rand.Reader, pubkey, data.([]byte))
		if err != nil {
			return nil, err
		}

		encoded := base64.StdEncoding.EncodeToString(signature)
		resp = append(resp, encoded)
	}

	return resp, nil
}

// EncodeRSA RSA+BASE64 encode
func EncodeRSA(pubkey *rsa.PublicKey, data interface{}) (string, error) {
	signature, err := rsa.EncryptPKCS1v15(rand.Reader, pubkey, data.([]byte))
	if err != nil {
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(signature)

	return encoded, nil
}

// DecodeRSA  RSA+BASE64 decode
func DecodeRSA(prikey *rsa.PrivateKey, data interface{}) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(data.(string))
	if err != nil {
		return nil, err
	}

	dsignature, err := rsa.DecryptPKCS1v15(rand.Reader, prikey, []byte(decoded))
	if err != nil {
		return nil, err
	}

	return dsignature, nil
}
