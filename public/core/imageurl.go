package core

import (
	"strings"
	"github.com/hunterhug/GoSpider/util"
	"errors"
	"os"
	"regexp"
)

func replaceRegString(re, repl, filename string) string{
	// re 	: regular expression
	// repl : replaced string
	// filename : target
    regex, err := regexp.Compile(re)
    if err != nil {
		return filename // there was a problem with the regular expression.
    }

	result := regex.ReplaceAllString(filename, repl)
	return result
}

func replaceImageUrlWithFormat(imageUrlformat, url string) string {
	// replace the image format
	return replaceRegString("[.][_](.*)[_][.]", imageUrlformat, url)
}

// get image with the URL
func GetImageUrl(ip string, url string) ([]byte, error) {
	// Filename should be Asin + image_rank + .jpg
	// TODO 2017.12.11 modify the filename layer
	filename := strings.Split(url, "/I/")

	if len(filename) != 2{
	}
	keepdirtemp := MyConfig.Datadir + "/asin/" + Today + "/" + filename[1] + ".jpeg"
	if MyConfig.Asinlocalkeep {
		if util.FileExist(keepdirtemp){
			AmazonAsinLog.Debugf("FileExists:%s", keepdirtemp)
			return util.ReadfromFile(keepdirtemp)
		}
		if util.FileExist(keepdirtemp + "sql") {
			AmazonAsinLog.Debugf("FileExist: %s", keepdirtemp)
			return util.ReadfromFile(keepdirtemp + "sql")
		}
	}
	// TODO change to normal later
	content, err := Download("*", replaceImageUrlWithFormat(MyConfig.ImageUrlFormat, url))

	if err != nil {
		return nil, err
	}
	// Judge if it is robot.
	if IsRobot(content){
		return nil, errors.New("robot")
	}
	if Is404(content){
		return nil, errors.New("404")
	}
	if MyConfig.Asinlocalkeep {
		util.SaveToFile(keepdirtemp, content)
	}
	// Save the image and insert into the asin table.
	img, err := os.Create("/Users/kcyu/mount/plots/"  + filename[1] + ".jpg")
	defer img.Close()
	if err != nil {
		return nil, err
	}
	_, err = img.Write(content)

	return content, nil
}

func GetImageUrls() error {
	AmazonAsinLog.Log("Atart Get Image from Url")
	ip := GetIP()

	// before use, send to hash pool
	ipbegintimes := util.GetSecend2DateTimes(util.GetSecendTimes())
	RedisClient.Hset(MyConfig.Proxyhashpool, ip, ipbegintimes)

	// do a lot urls still can't pop url. ??
	for {
		// take url, if block
		url, err := RedisClient.Brpoplpush(MyConfig.Asinpool, MyConfig.Asindealpool, 0)
	}
}