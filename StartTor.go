package main

import (
	"os"
	"os/exec"
	"strconv"
	"log"
)

func main() {
	i,_ := strconv.Atoi(os.Getenv("MAX_DRIVER"))
	for i > 0{
		i--
		// ポート番号生成
		s := strconv.Itoa(9000 + i)
		// ディレクトリ名
		d := "DataDirectory/" + s
		// データディレクトリ作成（ダブったらTorが起動できない）
		if err := os.MkdirAll(d, 0777); err != nil {
			log.Println(err)
		}
		err := exec.Command("tor", "-socksport", s, "DataDirectory",d).Start()
		if err !=nil {
			log.Println(err)
			return
		}
	}
}
