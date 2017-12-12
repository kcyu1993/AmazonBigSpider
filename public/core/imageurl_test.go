package core

import (
	"github.com/hunterhug/GoSpider/spider"
	"github.com/hunterhug/AmazonBigSpider/public/misc"
	"testing"
)

func TestGetImageUrl(t *testing.T) {
	spider.SetLogLevel("debug")
	ip := "10.90.37.46:30981"
	url := "https://images-na.ssl-images-amazon.com/images/I/81mO7wXAu4L._SL500_SR123,160_.jpg"
	misc.replace
	_, err := GetImageUrl(ip, url)
	if err != nil {
		t.Error(err)
	}
}
