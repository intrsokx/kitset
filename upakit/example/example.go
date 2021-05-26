package main

import (
	"github.com/intrsokx/kitset/upakit"
	"net/url"
)

var upa *upakit.UPAUtil

func init() {
	upaAuthUrl := "xxx"
	upaRepoUrlFmt := "xxx"
	developmentId := "xxx"
	authSignature := "xxxx"
	baseKey := "xxxx"
	proxy, _ := url.Parse("http://xxx")

	upa = upakit.NewUPAUtil(upaAuthUrl, upaRepoUrlFmt, developmentId, authSignature, baseKey, proxy)
}

func main() {
	//upa.GetUPACommonPersonAuthServer()
	//upa.GetUPAAuthRecognizeServer()
}
