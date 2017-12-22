BASE_PROJECT := github.com/lucas59356/gedis

travis: build_api build_benchmark core_test

build_api: install_deps
	go build -v -i $(BASE_PROJECT)/api
build_benchmark: install_deps
	go build -v -i $(BASE_PROJECT)/cli/gedis-benchmark
core_test: install_deps
	go test -i -v $(BASE_PROJECT)/core 

install_deps:
	go get github.com/cloudfoundry/gosigar
	go get github.com/gorilla/mux