package main

import (
	"github.com/sclevine/agouti"
	"log"
	"sync"
	"os"
	"strconv"
	"os/exec"
	"time"
)

func main() {
	sy := sync.WaitGroup{}

	i,_ := strconv.Atoi(os.Getenv("MAX_DRIVER"))
	for i > 0{
		i --
		sy.Add(1)
		go func(i int){
			defer sy.Done()
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
			defer page.CloseWindow()

			// ここからブラウザにやらせたい処理を書く
			c := 2
			for c > 0{
				c--
				// とりあえずIPを確認するだけの処理
				if err := page.Navigate("http://ipinfo.io/"); err != nil {
					log.Fatalf("Failed to navigate:%v", err)
				}
				// pageはJqueryのセレクターとか使えるよ。
				log.Println(page.Title())


				// IPアドレスを変更する
				if err := exec.Command("printf",`"AUTHENTICATE \"password\"\r\nSIGNAL NEWNYM\r\n"`, "|",
					"nc localhost " + s).Start(); err != nil{
						log.Println(err)
				}
				time.Sleep(100 * time.Millisecond) // change ip address now.
			}
		}(i)
	}
	sy.Wait()
}