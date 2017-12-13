package core

import (
	"strings"
	"github.com/hunterhug/GoSpider/util"
	"errors"
	"os"
	"regexp"
	"fmt"
)

func replaceRegString(re string, repl string, filename string) string{
	// re 	: regular expression
	// repl : replaced string
	// filename : target
    regex, err := regexp.Compile(re)
    if err != nil {
		panic(err)
		return filename // there was a problem with the regular expression.
    }

	result := regex.ReplaceAllString(filename, repl)
	return result
}

func replaceImageUrlWithFormat(imageUrlformat string, url string) string {
	// replace the image format
	return replaceRegString("[.][_](.*)[_][.]", imageUrlformat, url)
}

func saveImageFromBinaryWebContent(content []byte, filename string) error {
	/* Save the corresponding image based on the web content */
	img, err := os.Create(filename)
	defer img.Close()
	if err != nil {
		return err
	}
	_, err = img.Write(content)
	return err
}

// get image with the URL
func GetImageUrl(ip string, url string) ([]byte, error) {
	// TODO 2017.12.11 modify the filename layer
	filename := strings.Split(url, "/")

	keepdirtemp := MyConfig.Datadir + "/tmp_image/" + Today + "/" + filename[len(filename)-1]
	if MyConfig.Asinlocalkeep {
		if util.FileExist(keepdirtemp){
			AmazonImageLog.Debugf("FileExists:%s", keepdirtemp)
			return util.ReadfromFile(keepdirtemp)
		}
		if util.FileExist(keepdirtemp + "sql") {
			AmazonImageLog.Debugf("FileExist: %s", keepdirtemp)
			return util.ReadfromFile(keepdirtemp + "sql")
		}
	}
	// TODO change to normal later
	url = strings.Replace(url, "0jpg", "0.jpg", -1)
	content, err := Download(ip, replaceImageUrlWithFormat(MyConfig.ImageUrlFormat, url))

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
		err = util.SaveToFile(keepdirtemp, content)
		if err != nil {
			AmazonImageLog.Errorf( "save tmp file %s error %s", keepdirtemp, err.Error())
			panic(err)
		}
	}
	return content, nil
}
//
func GetImageUrls() error {
	AmazonImageLog.Log("Start Get Image from Url")
	ip := GetIP()

	// before use, send to hash pool
	ipbegintimes := util.GetSecond2DateTimes(util.GetSecondTimes())
	RedisClient.Hset(MyConfig.Proxyhashpool, ip, ipbegintimes)

	// Keep looping until there is no url left in the Redis.ImageUrlpool.
	for {
		// take url, if block
		url_id, err := RedisClient.Brpoplpush(MyConfig.ImageUrlpool, MyConfig.ImageUrldealpool, 0)
		if err != nil {
			// No result in ImageUrlpool
			return err
		}
		//fmt.Println(url)
		decode := strings.Split(url_id, "|")
		if len(decode) != 2 {
			AmazonImageLog.Errorf("Url ID error %s", url_id)
			continue
		}
		url := decode[0]
		imageId := decode[1]

		exist, _ := RedisClient.Hexists(MyConfig.ImageUrlhashpool, url)
		if exist {
			AmazonImageLog.Errorf("exists %s", url)
			continue
		}

		urlbegintime := util.GetSecond2DateTimes(util.GetSecondTimes())

		// Initialize the content and err
		content := []byte("")
		err = nil
		for {
			// Grab the content with the same logic
			content, err = GetImageUrl(ip, url)
			// Save the image and insert into the asin table.
			spider, ok := Spiders.Get(ip)
			// Handle the no error logic.
			if err == nil {
				break
			} else {
				if strings.Contains(err.Error(), "404") {
					// If 404, (due to the invalid address) break
					break
				}
				if strings.Contains(err.Error(), "robot") {
					if ok {
						spider.Errortimes = spider.Errortimes + 1
					}
				}
				if ok {
					AmazonImageLog.Errorf("get %s fail(%d),total(%d) error:%s,ip:%s", url, spider.Errortimes, spider.Fetchtimes, err.Error(), ip)
				}
			}
			if ok && spider.Errortimes > MyConfig.Proxymaxtrytimes {
				// die
				ipendtimes := util.GetSecond2DateTimes(util.GetSecondTimes())
				insidetemp := ipbegintimes + "|" + ipendtimes + "|" + util.IS(spider.Fetchtimes - spider.Errortimes) + "|" + util.IS(spider.Errortimes)
				RedisClient.Hset(MyConfig.Proxyhashpool, ip, insidetemp)
				// delete the ip from spider
				Spiders.Delete(ip)
				// get new proxy again
				ip = GetIP()
				ipbegintimes = util.GetSecond2DateTimes(util.GetSecondTimes())
				RedisClient.Hset(MyConfig.Proxyhashpool, ip, ipbegintimes)
			}
		} // end IP grabbing error
		if err != nil && strings.Contains(err.Error(), "404") {
			// handle 404 error
			err = SetImageInvalid(imageId)
			if err != nil {
				AmazonImageLog.Errorf("%s set invalid error: %s", imageId, err.Error())
			}
		} else {
			// save the actual image here and insert into the mysql database
			// TODO find a better logic to save the directory
			image_filename := "/image/" + Today + "/" + imageId + ".jpeg"
			AmazonImageLog.Logf("%s save image", imageId)
			err := saveImageFromBinaryWebContent(content, MyConfig.Datadir + image_filename)
			if err != nil {
				// save failure, should never happen
				AmazonImageLog.Errorf("%s save image error %s", imageId, err.Error())
				panic(err)
			}
			err = InsertImageMysql(imageId, url, image_filename)
			if err != nil {
				AmazonImageLog.Errorf("%s mysql insert error", imageId, err.Error())
				panic(err.Error())
			}
		}
		// done, remove redis deal pool

		RedisClient.Lrem(MyConfig.ImageUrldealpool, 0, url)
		urlendtime := util.GetSecond2DateTimes(util.GetSecondTimes())
		// Put to the hash pool, for duplication checking
		RedisClient.Hset(MyConfig.ImageUrlhashpool, url, urlbegintime + "|" + urlendtime)
		numm, _ := RedisClient.Hlen(MyConfig.ImageUrlhashpool)
		//if e != nil {
		//	panic(e)
		//}
		AmazonImageLog.Logf("Reduce ImageUrldealpool and add to hash %s %d", imageId, numm)
	}
	return nil
}

func GetNoneProxyImageUrls(taskname string) error {
	/* Remove the IP handling logic, keep the rest same as proxy image urls */
	AmazonImageLog.Log("Atart Get Image from Url, without proxy")
	ip := "*" + taskname

	// Keep looping until there is no url left in the Redis.ImageUrlpool.
	for {
		// take url, if block
		url_id, err := RedisClient.Brpoplpush(MyConfig.ImageUrlpool, MyConfig.ImageUrldealpool, 0)
		if err != nil {
			return err
		}
		fmt.Println(url)
		decode := strings.Split(url_id, "|")
		if len(decode) != 2 {
			AmazonImageLog.Errorf("Url ID error %s", url_id)
			continue
		}
		url := decode[0]
		imageId := decode[1]

		exist, _ := RedisClient.Hexists(MyConfig.ImageUrlhashpool, url)
		if exist {
			AmazonImageLog.Errorf("exists %s", url)
			continue
		}

		urlbegintime := util.GetSecond2DateTimes(util.GetSecondTimes())

		// Initialize the content and err
		content := []byte("")
		err = nil
		for {
			// Grab the content with the same logic
			content, err = GetImageUrl(ip, url)
			// Save the image and insert into the asin table.
			spider, ok := Spiders.Get(ip)
			// Handle the no error logic.
			if err == nil {
				break
			} else {
				if strings.Contains(err.Error(), "404") {
					// If 404, (due to the invalid address) break
					break
				}
				if strings.Contains(err.Error(), "robot") {
					if ok {
						spider.Errortimes = spider.Errortimes + 1
					}
				}
				if ok {
					AmazonAsinLog.Errorf("get %s fail(%d),total(%d) error:%s,ip:%s", url, spider.Errortimes, spider.Fetchtimes, err.Error(), ip)
				}
			}
			if ok && spider.Errortimes > MyConfig.Proxymaxtrytimes {
				// die
				Spiders.Delete(ip)
			}
		} // end IP grabbing error
		if err != nil && strings.Contains(err.Error(), "404") {
			// handle 404 error
			err = SetImageInvalid(imageId)
			if err != nil {
				AmazonImageLog.Errorf("%s set invalid error: %s", imageId, err.Error())
			}
		} else {
			// save the actual image here and insert into the mysql database
			// TODO find a better logic to save the directory
			image_filename := "/image/" + Today + "/" + imageId + ".jpeg"
			err := saveImageFromBinaryWebContent(content, MyConfig.Datadir + image_filename)
			if err != nil {
				// save failure, should never happen
				AmazonImageLog.Errorf("%s save image error %s", imageId, err.Error())
				panic(err)
			}
			err = InsertImageMysql(imageId, url, image_filename)
			if err != nil {
				AmazonImageLog.Errorf("%s mysql insert error", imageId, err.Error())
			}
		}
		// done, remove redis deal pool
		RedisClient.Lrem(MyConfig.ImageUrldealpool, 0, url)
		urlendtime := util.GetSecond2DateTimes(util.GetSecondTimes())
		// Put to the hash pool, for duplication checking
		RedisClient.Hset(MyConfig.ImageUrlhashpool, url, urlbegintime + "|" + urlendtime)
	}
	return nil
}


//func GetImageUrl(ip string, url string) ([]byte, error) {
//	// Filename should be Asin + image_rank + .jpg
//	// TODO 2017.12.11 modify the filename layer
//	filename := strings.Split(url, "/I/")
//
//	if len(filename) != 2{
//	}
//	keepdirtemp := MyConfig.Datadir + "/image/" + Today + "/" + filename[1] + ".jpeg"
//	if MyConfig.Asinlocalkeep {
//		if util.FileExist(keepdirtemp){
//			AmazonImageLog.Debugf("FileExists:%s", keepdirtemp)
//			return util.ReadfromFile(keepdirtemp)
//		}
//		if util.FileExist(keepdirtemp + "sql") {
//			AmazonImageLog.Debugf("FileExist: %s", keepdirtemp)
//			return util.ReadfromFile(keepdirtemp + "sql")
//		}
//	}
//	// TODO change to normal later
//	content, err := Download("*", replaceImageUrlWithFormat(MyConfig.ImageUrlFormat, url))
//
//	if err != nil {
//		return nil, err
//	}
//	// Judge if it is robot.
//	if IsRobot(content){
//		return nil, errors.New("robot")
//	}
//	if Is404(content){
//		return nil, errors.New("404")
//	}
//	if MyConfig.Asinlocalkeep {
//		util.SaveToFile(keepdirtemp, content)
//	}
//	// Save the image and insert into the asin table.
//	img, err := os.Create("/Users/kcyu/mount/plots/"  + filename[1] + ".jpg")
//	defer img.Close()
//	if err != nil {
//		return nil, err
//	}
//	_, err = img.Write(content)
//
//	return content, nil
//}