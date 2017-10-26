/*
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:459527502

	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:459527502

	2017.7 by hunterhug
*/
package AmazonBigSpider

import (
	"flag"
	"fmt"
	"github.com/hunterhug/GoSpider/util"
	// 为了依赖
	_ "github.com/hunterhug/GoSpider/query"
	_ "github.com/hunterhug/GoSpider/spider"
	_ "github.com/hunterhug/GoSpider/store/myredis"
	_ "github.com/hunterhug/GoSpider/store/mysql"
	"path/filepath"
)

var Dir = util.CurDir()
var CoreDir = filepath.Join(Dir, "public", "core")
var Local = true
var ToolStep int = 0
var ToolProxy bool = false
var User = ""

func init() {
	rootdir := flag.String("root", "", "root config")
	coredir := flag.String("core", "", "core config")
	temp := flag.Int("toolstep", 0, "which step get category url")
	temp1 := flag.Bool("toolproxy", false, "proxy get category url?")
	temp2 := flag.String("user", "", "user")
	if !flag.Parsed() {
		flag.Parse()
	}

	ToolStep = *temp
	ToolProxy = *temp1
	User = *temp2
	if *rootdir != "" {
		Dir = *rootdir
	}
	if *coredir != "" {
		CoreDir = *coredir
	}
	// 在根目录建一个远程.txt使用远程配置
	if util.FileExist(Dir + "/远程.txt") {
		Local = false
		fmt.Println("远程方式！！！")
	}

	//addrs, err := net.InterfaceAddrs()
	//
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//
	//	for _, address := range addrs {
	//
	//		// 检查ip地址判断是否回环地址
	//		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	//			if ipnet.IP.To4() != nil {
	//				fmt.Println(ipnet.IP.String())
	//			}
	//
	//		}
	//	}
	//}
}
