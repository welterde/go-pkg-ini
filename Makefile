
all:
	cd ini && make clean && make

clean:
	cd ini && make clean

test:
	cd ini && make && make test

format:
	gofmt -w .

