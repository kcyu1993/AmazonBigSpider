/*
	Modify the spider to create an image task.
	Create the corresponding image task.

	2017.12 by kcyu1993
*/
package core

import (
	"github.com/hunterhug/GoSpider/util"
	"math/rand"
	"os"
)

var (
	asinimagetasknum int
	asinimageendchan chan string
)

func asinimagetask(taskname string) {
	second := rand.Intn(5)
	AmazonImageLog.Debugf("%s:%d second", taskname, second)
	util.Sleep(second)
	if MyConfig.Proxyasin {
		err := GetImageUrls()
		if err != nil {
			AmazonImageLog.Errorf(taskname + "-error:" + err.Error())
		}
	} else {
		err := GetNoneProxyImageUrls(taskname)
		if err != nil {
			AmazonImageLog.Errorf(taskname + "-error:" + err.Error())
		}
	}
	asinendchan <- "done!"
}

func AsinImageGo() {
	OpenMysql()
	err := CreateAsinImageTabels()

	if err != nil {
		AmazonImageLog.Errorf("createtables:%s,error:%s", Today, err.Error())
	}

	// Create the directory accordingly.
	err = os.MkdirAll(MyConfig.Datadir + "/image/" + Today, os.ModePerm)
	if err != nil {
		AmazonImageLog.Errorf("create folder error %s", err.Error())
	}

	// Create the tmp folder accordingly.
	err = os.MkdirAll(MyConfig.Datadir + "/tmp_image/" + Today, os.ModePerm)
	if err != nil {
		AmazonImageLog.Errorf("create folder error %s", err.Error())
	}

	asinimagetasknum = MyConfig.Imagetasknum
	asinimageendchan = make(chan string, asinimagetasknum)
	for i := 0; i < asinimagetasknum; i++ {
		go asinimagetask("imagetask" + util.IS(i))
	}
	go Clean()
	for i := 0; i < asinimagetasknum; i++ {
		<-asinimageendchan
	}
	AmazonImageLog.Log("List All done")
}
