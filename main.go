package main

import (
	"github.com/sclevine/agouti"
	"log"
	"sync"
	"os"
	"strconv"
)

func main() {
	sy := sync.WaitGroup{}

	i,_ := strconv.Atoi(os.Getenv("MAX_DRIVER"))

	for i > 0{
		i --
		sy.Add(1)
		go func(i int){
			s := strconv.Itoa(9000 + i)
			driver := agouti.ChromeDriver(
				agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu",
					"--proxy-server=socks5://localhost:" + s}),
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

			defer sy.Done()
		}(i)
	}
	sy.Wait()
}