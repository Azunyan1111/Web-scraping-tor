setup:
	go get ./...

start-tor:
	go run StartTor.go
	echo 'Tor service setup. please wait and check conection'
	echo 'curl --socks5 localhost:9001 http://ipinfo.io/'

stop-tor:
	killall tor

start:
	go run main.go