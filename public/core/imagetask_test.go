package core

import (
	"github.com/hunterhug/GoSpider/spider"
	"testing"
)

func TestAsinImageGo(t *testing.T) {
	// TODO Finish this test module later
	spider.SetLogLevel("debug")
	ip := "104.128.124.122:808"
	url := "https://www.amazon.com/dp/B01M4R0M9V"
	_, err := GetAsinUrl(ip, url)
	if err != nil {
		t.Error(err)
	}
}