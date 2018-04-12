all: update build run

no-update: build run

run:
	./marvin-go

build:
	go get
	go build

clean:
	rm marvin-go

update:
	git pull