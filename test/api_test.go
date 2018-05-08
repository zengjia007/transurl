package test

import (
	"testing"
	"transurl/api"
	"os"
	"log"
)

func TestLongToShort(t *testing.T) {
	path, _ := os.Getwd()
	log.Println("当前路径为：", path)
	var longUrl = "http://www.baidu.com"
	_, err := api.LongToShort(longUrl)
	if err != nil {
		t.Error("A error accurrs!")
	}
}

func TestShortToLongUrl(t *testing.T) {
	var shortUrl = "http://127.0.0.1:8088/w"
	_, err := api.ShortToLongUrl(shortUrl)
	if err != nil {
		t.Error("A error accurrs!")
	}
}