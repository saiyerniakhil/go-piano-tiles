BINARY_NAME=pianotiles

build:
	echo "building game"
	go build -o dist/$(BINARY_NAME) .

run: build
	./dist/$(BINARY_NAME)

release:
	@if [ "$(PUSH)" = "true" ]; then \
		npm run release; \
		git push --follow-tags origin HEAD:main; \
	else \
		npm run release; \
		echo "Release completed locally. No changes pushed to origin."; \
		goreleaser release --snapshot --clean; \
	fi

clean:
	rm -rf dist/*