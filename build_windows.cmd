rem Make the build for windows, it not depends of GNU Make

set BASE_PROJECT=github.com/lucas59356/gedis

go get github.com/cloudfoundry/gosigar
go get github.com/gorilla/mux

go test -v %BASE_PROJECT%/core
go install %BASE_PROJECT%/cli/gedis-api
go install %BASE_PROJECT%/cli/gedis-benchmark