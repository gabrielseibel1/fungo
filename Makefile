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
