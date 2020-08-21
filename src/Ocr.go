package main

import (
	"fmt"
	ocr "github.com/ranghetto/go_ocr_space"
)

func main() {
	//this is a demo api key
	apiKey := "helloworld"

	//setting up the configuration
	config := ocr.InitConfig(apiKey, "chs")
	//actual converting the image from the url (struct content)
	result, err := config.ParseFromLocal("/Users/lgc/Desktop/a.pdf")
	if err != nil {
		fmt.Println(err)
	}
	//printing the just the parsed text
	fmt.Println(result.JustText())
}
