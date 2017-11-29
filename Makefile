setup:
	go get ./...

start-tor:
	go run StartTor.go

stop-tor:
	killall tor

start:
	go run main.go