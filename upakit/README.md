# upakit to do list
> * 对外暴露获取数据接口做限流
> * 对刷新认证接口做熔断

## example
```
upaUtil := upakit.NewUPAUtil(upaAuthUrl, upaRepoUrlFmt, developmentId, authSignature, key, proxy)

data, err := upa.GetUPAAuthRecognizeServer(cardNo, name, idCard, mobile, mode, merName, authCode, authFlag)
data, err = upa.GetUPACommonPersonAuthServer(resourceId, cardNo, name, idCard, mobile, authCode, authFlag)
```