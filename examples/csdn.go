package main

import (
	"github.com/anjude/bean-sdk-go/beansdk"
	"github.com/anjude/bean-sdk-go/services/csdn_service"
	"log"
)

func main() {
	client := beansdk.NewClient("", "")

	csdnService := csdn_service.NewCsdnService(client, "xxx", "xxx")
	resp, err := csdnService.HotArticleComment()
	if err != nil {
		panic(err)
	}
	log.Printf("done: %v", resp)
}
