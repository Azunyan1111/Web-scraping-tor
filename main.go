package main

import (
	"github.com/sclevine/agouti"
	"log"
)

func main() {
	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu",
		"--proxy-server=socks5://localhost:9050"}),
	)
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	if err := page.Navigate("http://ipinfo.io/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	log.Println(page.Title())
}