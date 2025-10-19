BINARY_NAME=pianotiles

build:
	echo "building game"
	go build -o dist/$(BINARY_NAME) .

run: build
	./dist/$(BINARY_NAME)
