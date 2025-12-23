BINARY_NAME=pianotiles

build:
	echo "building game"
	go build -o dist/$(BINARY_NAME) .

run: build
	./dist/$(BINARY_NAME)

release:
	@if [ "$(PUSH)" = "true" ]; then \
		npm run release; \
		git push --follow-tags origin $$(git rev-parse --abbrev-ref HEAD); \
	else \
		npm run release; \
		echo "Release completed locally. No changes pushed to origin."; \
	fi
