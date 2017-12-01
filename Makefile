setup:
	go get ./...

start-tor:
	go run main/StartTor.go
	export MAX_DRIVER=10
	source ~/.bashrc
	# Tor service setup now. please wait and check conection
	# curl --socks5 localhost:9001 http://ipinfo.io/

stop:
	killall tor

start:
	go run main.go