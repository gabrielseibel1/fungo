all: verify

verify: vet check test run

vet:
	go vet ./...

check:
	staticcheck ./...

test:
	go test --count 1 --cover ./...

run:
	go run .
