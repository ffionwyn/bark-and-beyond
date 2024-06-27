build:
	go build -o server .

start:
	./server

.PHONY: build start
