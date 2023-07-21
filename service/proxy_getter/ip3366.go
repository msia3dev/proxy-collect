package proxy_getter

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tongsq/go-lib/logger"
	"proxy-collect/consts"
	"proxy-collect/global"
	"proxy-collect/service/common"
)

func NewGetProxyIp3366() *getProxyIp3366 {
	return &getProxyIp3366{}
}

type getProxyIp3366 struct {
}

func (s *getProxyIp3366) GetUrlList() []string {
	list := []string{
		"http://www.ip3366.net/free/?stype=1",
		"http://www.ip3366.net/free/?stype=2",
	}
	for i := 2; i < 4; i++ {
		list = append(list, fmt.Sprintf("http://www.ip3366.net/free/?stype=1&page=%d", i))
		list = append(list, fmt.Sprintf("http://www.ip3366.net/free/?stype=2&page=%d", i))
	}
	return list
}
func (s *getProxyIp3366) GetContentHtml(requestUrl string) string {
	logger.Info("get proxy from ip3366", logger.Fields{"url": requestUrl})
	data, err := global.SimpleGet(requestUrl)
	if err != nil || data == nil {
		logger.Error("get proxy from ip3366 fail", logger.Fields{"err": err, "data": data})
		return ""
	}
	return data.Body
}

func (s *getProxyIp3366) ParseHtml(body string) [][]string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error("read fail", logger.Fields{"err": err})
		return nil
	}
	var proxyList [][]string
	doc.Find("tbody > tr").Each(func(i int, selection *goquery.Selection) {
		td := selection.ChildrenFiltered("td").First()
		host := strings.TrimSpace(td.Text())
		td2 := selection.ChildrenFiltered("td").Eq(1)
		port := strings.TrimSpace(td2.Text())
		td3 := selection.ChildrenFiltered("td").Eq(3)
        proxyType := strings.ToLower(strings.TrimSpace(td3.Text()))

		if !common.CheckProxyFormat(host, port) {
			logger.Error(consts.PROXY_FORMAT_ERROR, logger.Fields{"host": host, "port": port})
			return
		}
		proxyArr := []string{host, port, proxyType}
		proxyList = append(proxyList, proxyArr)
	})
	return proxyList
}
