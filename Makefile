setup:
	go get ./...

start-tor:
	tor -f tor/torrc1 &

stop-tor:
	killall tor

start:
	go run main.go