// main.go
package phantomjs

import (
	"github.com/sclevine/agouti"
	"log"
)

func main() {
	command := []string{"phantomjs", "--webdriver={{.Address}}","--proxy=localhost:9050","--proxy-type=socks5"}
	driver := agouti.NewWebDriver("http://{{.Address}}", command)

	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	if err := page.Navigate("http://ipinfo.io/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	log.Println(page.Title())
}
