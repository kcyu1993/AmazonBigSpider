package main

import (
	"github.com/hunterhug/AmazonBigSpider"
	"github.com/hunterhug/AmazonBigSpider/public/core"
	"fmt"
)


func displayAsinPool() {
	num1, _ := core.RedisClient.Llen(core.MyConfig.Asinpool)
	num2, _ := core.RedisClient.Hlen(core.MyConfig.Asinpool)
	fmt.Printf("Image Asin pool Llen %d Hlen %d \n" ,num1, num2)
	num1, _ = core.RedisClient.Llen(core.MyConfig.Asinhashpool)
	num2, _ = core.RedisClient.Hlen(core.MyConfig.Asinhashpool)
	fmt.Printf("Image Asin hash pool Llen %d Hlen %d \n" ,num1, num2)
	num1, _ = core.RedisClient.Llen(core.MyConfig.Asindealpool)
	num2, _ = core.RedisClient.Hlen(core.MyConfig.Asindealpool)
	fmt.Printf("Image Asin deal pool Llen %d Hlen %d \n" ,num1, num2)
}

func displayImagePool(){
	num1, _ := core.RedisClient.Llen(core.MyConfig.ImageUrldealpool)
	num2, _ := core.RedisClient.Hlen(core.MyConfig.ImageUrldealpool)
	fmt.Printf("Image URL deal pool Llen %d Hlen %d \n" ,num1, num2)
	num1, _ = core.RedisClient.Llen(core.MyConfig.ImageUrlhashpool)
	num2, _ = core.RedisClient.Hlen(core.MyConfig.ImageUrlhashpool)
	fmt.Printf("Image URL hash pool Llen %d Hlen %d \n" ,num1, num2)
	num1, _ = core.RedisClient.Llen(core.MyConfig.ImageUrlpool)
	num2, _ = core.RedisClient.Hlen(core.MyConfig.ImageUrlpool)
	fmt.Printf("Image URL pool Llen %d Hlen %d \n" ,num1, num2)

}

func cleanPool(key string) {
	num, e := core.RedisClient.Llen(key)
	if e != nil {
		panic(e)
	}
	for i := int64(0); i < num; i ++ {
		core.RedisClient.Lpop(key)
	}
}

func cleanImagePool() {
	cleanPool(core.MyConfig.ImageUrlpool)
	cleanPool(core.MyConfig.ImageUrldealpool)
	cleanPool(core.MyConfig.ImageUrlhashpool)

}

func main(){
	if AmazonBigSpider.Local {
		core.InitConfig(AmazonBigSpider.Dir + "/config/usa_local_config.json", AmazonBigSpider.Dir + "/config/usa_local_log.json")
	} else {
		core.InitConfig(AmazonBigSpider.Dir + "/config/usa_config.json", AmazonBigSpider.Dir + "/config/usa_log.json")
	}
	//cleanImagePool()
	displayImagePool()
}


