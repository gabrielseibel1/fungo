all: verify

verify: vet check test

.PHONY: vet
vet:
	go vet ./...

.PHONY: check
check:
	staticcheck ./...

.PHONY: test
test:
	go test --count 1 --cover --coverprofile=./cover.out ./...

coverage: test
	go tool cover -html ./cover.out -o ./cover.html
	xdg-open ./cover.html
