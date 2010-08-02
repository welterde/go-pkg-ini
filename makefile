
all:
	make -C ini install

test:
	make -C ini test

clean:
	make -C ini clean

format:
	gofmt -w .

