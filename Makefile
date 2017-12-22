BASE_PROJECT := github.com/lucas59356/gedis

travis: build_api build_benchmark core_test

build_api:
	go build -v -i $(BASE_PROJECT)/api
build_benchmark:
	go build -v -i $(BASE_PROJECT)/cli/gedis-benchmark
core_test:
	go test -i -v $(BASE_PROJECT)/core 