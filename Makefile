all: build run

run:
	./marvin-go

build:
	go get
	go build

clean:
	rm marvin-go

update:
	git pull
	go get -u github.com/kevinzwang/marvin-go/errhandler