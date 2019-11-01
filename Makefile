export GO111MODULE=on

install:
	mkdir /etc/dongl
	cp dist/configs/config.yml /etc/dongl
	mkdir /var/run/dongl/
	mkdir /var/log/dongl/
	touch /var/log/dongl/error.log

uninstall:
	rm -rf /etc/dongl
	rm -rf /var/log/dongl
	rm -rf /var/run/dongl

rebuild:
	go clean
	make build

format:
	gofmt -s -w .

test-package:
	RUN_MODE=test go test -race ./${package} -coverprofile coverprofile/${package}.out
	go tool cover -html=coverprofile/${package}.out -o coverprofile/${package}.html

test:
	cd agent && go vet
	cd api && go vet
	cd daemon && go vet
	RUN_MODE=test go test -race ./api -coverprofile=./api/coverage.txt -covermode=atomic -v
	RUN_MODE=test go test -race ./agent -coverprofile=./agent/coverage.txt -covermode=atomic -v
	RUN_MODE=test go test -race ./daemon -coverprofile=./daemon/coverage.txt -covermode=atomic -v

test-with-report:
	cd config && go vet
	cd agent && go vet
	cd api && go vet
	cd daemon && go vet
	mkdir -p coverprofile
	RUN_MODE=test go test -race ./api -coverprofile coverprofile/api.out
	go tool cover -html=coverprofile/api.out -o coverprofile/api.html
	RUN_MODE=test go test -race ./agent -coverprofile coverprofile/agent.out
	go tool cover -html=coverprofile/agent.out -o coverprofile/agent.html
	RUN_MODE=test go test -race ./daemon -coverprofile coverprofile/daemon.out
	go tool cover -html=coverprofile/daemon.out -o coverprofile/daemon.html

debug:
	go run cmd/example/main.go --pid=tests/var/run/dongl/.pid --cfg=tests/configs --debug=true