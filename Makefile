build:
	go build -o bin/portfolio main.go

run:
	./bin/portfolio &

stop:
	pkill portfolio

restart:
	make stop; make run
