package rsautil

import (
	"crypto"
	"encoding/base64"
	"testing"
)

var (
	pub = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyyFbwkFShFRn/I4dZ2fi
bBRiZ4oV3UUpvDF7IryhDll440iGkFPzDsIdMdPtcMPkgxfE5BF+o34kPIdmlt1k
tKcOtG7d6xgE8kut8+oPSl1OInzG5DiQg/fRlZc6ZvBSwOS5e6BwM3XztdnldDK7
XUlZTOz3bmyqE8+egtDHKlrmz/jHvF8omGzOIOBYfLN3gKIIuZ2M8WWPfOJe0yQh
XVnyMl7hXG9KjbI7GzU8MGGii9P5JIQVMLTOy52l7bjuDNkKPj6sOJIEqE2m7Ihu
vh+2oj6A1BMFv2ax7FySA/XQpMtJ/M6+jFPdPBBozaoVXnxexGc8gVVey0pqn3k/
EQIDAQAB
-----END PUBLIC KEY-----`
	pri = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAyyFbwkFShFRn/I4dZ2fibBRiZ4oV3UUpvDF7IryhDll440iG
kFPzDsIdMdPtcMPkgxfE5BF+o34kPIdmlt1ktKcOtG7d6xgE8kut8+oPSl1OInzG
5DiQg/fRlZc6ZvBSwOS5e6BwM3XztdnldDK7XUlZTOz3bmyqE8+egtDHKlrmz/jH
vF8omGzOIOBYfLN3gKIIuZ2M8WWPfOJe0yQhXVnyMl7hXG9KjbI7GzU8MGGii9P5
JIQVMLTOy52l7bjuDNkKPj6sOJIEqE2m7Ihuvh+2oj6A1BMFv2ax7FySA/XQpMtJ
/M6+jFPdPBBozaoVXnxexGc8gVVey0pqn3k/EQIDAQABAoIBAG2z7Fxy7t+sviQ1
lEe+YRhBwgttFfXUXn/WsUvHV6vqQlFtX88ep4v25dF9RSS7hvQNqDYMBLrDa0qN
Tah4lOTDvDtSDOPkqvc4TBAI/o0I6yPRA5FJwzKiajxB2jax399xJ4NO2InST/aM
YVFp/Kqa7HGRIOgwI4JjhJXdr1J9oe8L0S+H4qLFo0WUq9gnMr2IbN8nMvDdJQGR
nlT9Is7Qcp/U43I+3HU8k7ud3GgcnHi1KBIlSLSabLZNcqCDLfBW/THzbPZCJe7b
FK2lBxIbWtgtMA++BI2OyWxapgDBXXHA6tZ2AU2Q/o07scHBlOujXFYmfTgz4gDp
EGnpluECgYEA+tn7/50IaRb2tQnDMBzgHHreFh7og/Nz7k9MNumCvvZezRhBt91S
rVKDevmDHFzwQO6HyAmQz4JJHfBuZoaJsKlAS2/8weKFBX/2szn5ysJHxBGqhHlI
NRDqX/j7gMMsOCo9qY/MXRWUpwyGr5nAVfkBKZo7LBWv/YBZXUgk/SUCgYEAz0yj
keUmY0+U3cWg/PE2jhiBK8Lvk8aVOTUtMXwKuayjjxrXObu8RHCY2YDL4zqtbKC2
142RIWqj+2Dcd9HXfMgGh3P47d45HaTh2vD+T5FqhZkKUE6Y2TDEzQCdYwvvLRdu
cuCzOdhnJLlSgCdnrm8KhROzbK2LPS44K6d71H0CgYEAvRyfBRpen9NHFD6S3u+1
6OKcETMl+WwNByjS/UbXYZ2c5KOXz8RTswTUyF3YgQZzvY/V33GOsVG4S5DZugNN
RFikdvqrI4Pg4r+QvZdEgJ4sulzTH2HLlO32s3miKXV6HbGCoRUebUJ6ueEQnMud
m3LIdJOobli/P66GMHPWJt0CgYBNOHJSHbdgFTwSJNVkhAJbiltLzvDp7naV+e4c
2eUw51OCMnBsLDfkksENfMH2olwJ9BBIWY7vkMcHFDzsUXnhHK359USMb1R9a3dK
1K0XPMcefzTtV2nuthEJgKogREjTVkApgPSinq9FadeGr6cavnh/vCgBWuBcaQQ5
lsk1DQKBgQDQut4dYnpB/j7aZ2K1dQzjOK9Q1CflrCRgXauulnyr/UPXWSLH83s1
jX/X83UcUksMx8Tjn2+6Mw+1yN6r+P0KQ6p2C4X7oudpbtdvpNM4mhadBdTbWOTk
4FLOUHIwn88wZBkWcwZj7jUJkjaxKLm4yufcL5IPwN0rbUPUlzm6kQ==
-----END RSA PRIVATE KEY-----`
)

var origin = "origin data"

func TestEncrypt(t *testing.T) {
	cfg := &Config{
		ParsePublicKey:     ParsePKIXPublicKey,
		ParsePrivateKey:    ParsePKCS1PrivateKey,
		EncryptWithPublic:  EncryptPKCS1v15,
		EncryptWithPrivate: nil,
		DecryptWithPub:     nil,
		DecryptWithPrivate: DecryptPKCS1v15,
		//是否分段加密
		CryptSub:           true,
		//分段加密步长
		SubStep:            256,
	}

	//加密
	encrypted, err := Encrypt([]byte(origin), []byte(pub), cfg)
	if err != nil {
		t.Error("rsa encrypt fail:", err)
	}
	t.Log("encrypted base64:", base64.StdEncoding.EncodeToString(encrypted))

	//解密
	plain, err := Decrypt(encrypted, []byte(pri), cfg)
	if err != nil {
		t.Error("rsa decrypt fail:", err)
	}
	t.Log("plain:", string(plain))

	//解析公私钥
	priKey, err := ParsePKCS1PrivateKey([]byte(pri))
	if err != nil {
		t.Error("parse pkcs1 private key fail:", err)
	}

	pubKey, err := ParsePKIXPublicKey([]byte(pub))
	if err != nil {
		t.Error("parse pkix public key fail:", err)
	}

	//私钥加签
	sign, err := SignByPKCS1v15([]byte(origin), priKey, crypto.MD5)
	if err != nil {
		t.Error("use priKey add sign fail:", err)
	}
	//公钥验签
	success := VerifySignByPKCS1v15([]byte(origin), sign, pubKey, crypto.MD5)
	if !success {
		t.Error("use pubKey check sign fail")
	} else {
		t.Log("use pubKey check sign success")
	}
}
