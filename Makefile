BASE_PROJECT := github.com/lucas59356/gedis

travis: core_test install_deps
	go build -v $(BASE_PROJECT)/cli/gedis-api
	go build -v $(BASE_PROJECT)/cli/gedis-api
	go build -v $(BASE_PROJECT)/cli/gedis-benchmark

core_test: install_deps
	go test -v $(BASE_PROJECT)/core

install_deps:
	go get github.com/cloudfoundry/gosigar
	go get github.com/gorilla/mux


# Docker related builds
docker_build:
	docker build -t gedis .

docker_run: docker_build
	docker run -P gedis