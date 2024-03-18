package main

import (
	"log"
	"time"

	"github.com/tebeka/selenium"
)

func main() {
	// 创建 Chrome WebDriver 实例
	chrome, err := selenium.NewChromeDriverService()
	if err != nil {
		log.Fatalf("无法创建 Chrome WebDriver 实例: %v", err)
	}
	defer chrome.Quit()

	// 打开 Chrome 浏览器
	err = chrome.Start()
	if err != nil {
		log.Fatalf("无法打开 Chrome 浏览器: %v", err)
	}

	// 访问指定的 URL
	url := "https://www.example.com"
	err = chrome.Get(url)
	if err != nil {
		log.Fatalf("无法访问指定的 URL: %v", err)
	}

	// 等待一段时间，让页面加载完成
	time.Sleep(5 * time.Second)

	// 关闭 Chrome 浏览器
	err = chrome.Stop()
	if err != nil {
		log.Fatalf("无法关闭 Chrome 浏览器: %v", err)
	}
}
