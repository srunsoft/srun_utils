package utils

import (
	log "github.com/sirupsen/logrus"
	"github.com/tiaguinho/gosoap"
	"net/http"
	"time"
)

type Soap struct {
	Wsdl     string
	Function string
	Params   gosoap.Params
}

func (sp *Soap) Req() (res *gosoap.Response, err error) {
	httpClient := &http.Client{Timeout: 1500 * time.Millisecond}
	// 调用wsdl定义文件
	soap, err := gosoap.SoapClient(sp.Wsdl, httpClient)
	if err != nil {
		log.Errorf("SoapClient error: %s", err.Error())
		return
	}
	// 调用ws中的服务
	res, err = soap.Call(sp.Function, sp.Params)
	if err != nil {
		log.Errorf("Call error: %s", err.Error())
		return
	}
	return
}
