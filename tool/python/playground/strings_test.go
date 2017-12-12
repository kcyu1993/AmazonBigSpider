package main

import (
	"fmt"
	"github.com/hunterhug/AmazonBigSpider/public/misc"
)


func main(){
	REPLACE_PATTERN := "SL1000_SR640,640"
	imageUrl := "https://images-na.ssl-images-amazon.com/images/I/81fE70hd%2BjL._SL500_SR129,160_.jpg"

	newImageUrl := misc.replaceImageUrlWithFormat(REPLACE_PATTERN, imageUrl)
	fmt.Println(newImageUrl)

}