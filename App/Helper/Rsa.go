package Helper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"strings"
)

//"log"

func EncriptRSA(public []byte, message []byte) (string, error) {

	b1, err := RsaEncrypt(message, public)
	s := Base64Enc(b1)

	return s, err
}

func DecriptRSA(private []byte, message string) ([]byte, error) {
	res, err := Base64Dec(message)
	if err != nil {
		return nil, err
	}
	res2, err := RsaDecrypt(res, private)

	return res2, err
}

//func decript(public []byte, message []byte) (string, error) {}
//
//func encript(private []byte, message string) ([]byte, error) {}

//func main() {
//
//	//var SECRET_MESSAGE = []byte(`{"Command": "login", "Data": {"name":"vasya", "password":"qwerty"}}`)
//	var SECRET_MESSAGE = []byte(`{"Command": "login", "Data": {"name":"vasya", "password":"qwerty"}}`)
//
//	secret, _ := encriptRSA(PublicKey2048, []byte(SECRET_MESSAGE))
//
//	fmt.Println(secret) //wTTO3esTgfJnpXW1dgO9CIL1VR4EF/JpA+zdzn94FfbGkZLlqXURgxg+uMT4dYL+dIrAVRYY9IbtLL2foYAGvw==
//	return
//}
//
func Base64Enc(b1 []byte) string {
	s1 := base64.StdEncoding.EncodeToString(b1)
	s2 := ""
	var LEN int = 76
	for len(s1) > 76 {
		s2 = s2 + s1[:LEN] + "\n"
		s1 = s1[LEN:]
	}
	s2 = s2 + s1
	return s2
}

func Base64Dec(s1 string) ([]byte, error) {
	s1 = strings.Replace(s1, "\n", "", -1)
	s1 = strings.Replace(s1, "\r", "", -1)
	s1 = strings.Replace(s1, " ", "", -1)
	return base64.StdEncoding.DecodeString(s1)
}

func RsaDecrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func RsaEncrypt(origData []byte, key []byte) ([]byte, error) {
	block, _ := pem.Decode(key)

	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	pub := pubInterface.(*rsa.PublicKey)

	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

//2048
var PublicKey2048 = []byte(`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1Qu8IXdxAw/x0vB07t30
KqcPXWYU3DMXtUD0Rb4+Fxw3/vWovqiRXIrFzgqULE6+nZzliLg949Dy8u9wsT3s
wFqWsov7yRLDt2O2bhE4AWHTa9K7H/ImVhFqiPKJ4EvVst7k/BEEeb5zj2oiIoVq
vQV1cDPSqHDQgPHvB4ZfwjlenGBPYbxkQv46zy/crbPkewfvPb6ACdJufXd/2WrT
tKbvkWAf+LO6AzuRnJ9aolSQiT7KJhVix3vOy/rBp2o9+d2zrxLc8hc0URsa6VNU
02Afl/QGwC0NiZYFdZfzP+/WXD25HyDc6ppyIUZ4+uuP7Qmr14Wttpc7CaXU9kOn
DQIDAQAB
-----END PUBLIC KEY-----
`)

var PrivateKey2048 = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEA1Qu8IXdxAw/x0vB07t30KqcPXWYU3DMXtUD0Rb4+Fxw3/vWo
vqiRXIrFzgqULE6+nZzliLg949Dy8u9wsT3swFqWsov7yRLDt2O2bhE4AWHTa9K7
H/ImVhFqiPKJ4EvVst7k/BEEeb5zj2oiIoVqvQV1cDPSqHDQgPHvB4ZfwjlenGBP
YbxkQv46zy/crbPkewfvPb6ACdJufXd/2WrTtKbvkWAf+LO6AzuRnJ9aolSQiT7K
JhVix3vOy/rBp2o9+d2zrxLc8hc0URsa6VNU02Afl/QGwC0NiZYFdZfzP+/WXD25
HyDc6ppyIUZ4+uuP7Qmr14Wttpc7CaXU9kOnDQIDAQABAoIBAAWzrLNQmQu174jv
upFyC0jg875SLxi9nVLSjDcZolvH+4+tT3ja1hkd9bFQAys0aFGbk2EXXUCtpPLv
iZqUx7NWOVZJ/NEi2W4dExLbDx6qWZg3KZ7vQitnh+xmYV5JaUzCPmqoofAIMtuR
wZwe6f5pGP2JxYeQjWQGFZmj9QgjqG5hiYBaV0lDXYV/OCczluYQ8YhpWPznhltd
nXspvFN5Pfkmg5IGbICS9LE0Sz6fEZJ+W2l9ycBGgxSXvwfbqY8N+jvGTeiF8IZW
I0q7wlGwL0s5V5+NGIc677iNMiySd8RydCHhfYb+dW9OPEYVbamOQtoRaXSXDOYP
SFdBEFkCgYEA6aoYCId3N841v6SO43hVX/mL/e0nJ5YvvohmEoxu4IGc6ehBdlIW
2KJUSOz3B7MlCCRshfvC4hotWRf5PDQS3GDtyieojlayu5Rxt0ua6owapipAKq1C
lDbz+YWU45dsb8mKHmDBG2ORDqkSs4eZkxuEzU0DaN9TwaNE7rYt9S8CgYEA6WkX
i8IQrv7CPJK7AQdJMae/I53ROu6+hqs4ayTQ0p8bUz/2kQ1dKxWr17IkBp0daA5t
krWtZTqMurkjuqKHVrE7nou3H8gTdhHNDkC9zsYtGlqizjlXom8GnWFtys19zhDF
njKbZAwlpAPH0//FCO9N7ow4x1Jfkld3BpLH0IMCgYBHoroTg1Rta27W/iBAZJfo
AJ2Gj7o9TLH+e9qvPRaRBauRmn7XQ3t1lu64HM3vMmDcCi6rNPAeWWAlvE/QwTY3
EhBUtavbV4EcOPpT833fAmz5HGLnso6C2gYaLXXkAHxiGSua/Ja3SuMh13vARoYy
r3Ebb8znze/joaZ0eK3GSQKBgAGObtbCu4O/NSJcRBz0pHtNSAv/wGZpMyIIwD6O
p0nQP8llUsqS0T05TsUIbg8ROyH3fqmMxpQ3OLsCAwf6j3Z3hhM/kUAIAIE4Cyr+
KYXYcnKLiixf24K8xMmF3cqNZjpaiEbOxZR8NEAMumdifDjcB6QAkVxa40JBjQyh
K+3rAoGAZszMzYrKRdVwwv8a6EBRxvDZb7UW9cgorxcb0x3xrzOv1ANaBvzMpFZV
b9OT9JcZ45k4x0A1cKngoL25QKEqG5pHeSGs09YdEpZSVp8XFcIQnv3aqhwkgok1
nmVVq/n0oNxbvwIJZKCbW3RXPrKATGNpdn4zQc4cCdVq5fDOBsM=
-----END RSA PRIVATE KEY-----
`)

////////////////////////////////////////////////////////////////////////

//512
var PublicKey512 = []byte(`
-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBANjZbWtYWNKKgPiMhpjy91oiJ0OzrEEH
20CG361JK18MasmvBodAyoydUfhbLFmmFQt+Sg/M1VDlYplo6rldJvMCAwEAAQ==
-----END PUBLIC KEY-----
`)

var PrivateKey512 = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIBPAIBAAJBANjZbWtYWNKKgPiMhpjy91oiJ0OzrEEH20CG361JK18MasmvBodA
yoydUfhbLFmmFQt+Sg/M1VDlYplo6rldJvMCAwEAAQJBAMsv8DcFMTUWUoKSKgxm
nR73oZLuaBkHI4ny1uOoC9uiK+uBHUX1UQt2aCWlsVFiuLciEE3x27DE1NTh9+3B
nAECIQD/mKrc0HOCuMuhr+pysz9ISoxgz4V9cgNII12DJZK7kQIhANkxGGdREkmO
88IKwCM1dZC/f9IbDN9lni2QXP5j7RBDAiBw5+miPVqpRiR9ug3guRmdP4EfSsx3
C6Qze5vVUQWuYQIhAK5NM0t/CZEU81TvccEP3xjaRodkhBEoqSfh6m+R3CenAiEA
uaRIxHRRrvxsP/hagj8jNgfjef44uAXa5MA5mBtSl3Y=
-----END RSA PRIVATE KEY-----
`)
