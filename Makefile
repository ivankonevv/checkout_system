build:
	go build -o ./.bin/binary ./main.go

run: build
	./.bin/binary