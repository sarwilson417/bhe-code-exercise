make run:
	go run cmd/bhe-code-exercise/main.go

generate:
	go generate ./...

test: generate
	go clean -testcache
	go test ./... -coverpkg=./... -coverprofile test-coverage.out
	# update README.md badge with test coverage
	test_coverage=$$(go tool cover -func=test-coverage.out | grep "total:" | awk '{print $$3}');\
	sed -i "" -e "s/Test%20Coverage-.*25-green.svg/Test%20Coverage-$${test_coverage}25-green.svg/" "README.md"
	go tool cover -html=test-coverage.out -o test-coverage.html