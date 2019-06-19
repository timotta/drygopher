

local: build test

prebuild:
	@echo Preparing build tooling...
	@go get -u github.com/vektra/mockery/.../
	@go get -u github.com/golang/dep/cmd/dep
.PHONY: prebuild

build:
	@echo Updating dependencies...
	@dep ensure
	@cd drygopher && CGO_ENABLED=0 GOARCH=amd64 go build -a -tags netgo -ldflags '-w -s' -o ../bin/drygopher main.go
.PHONY: build

test:
	@echo Purging old mocks...
	@rm -drf drygopher/mocks
	@echo Building mocks...
	@mockery -output drygopher/mocks -dir drygopher/coverage -all 
	@echo Ready to test.
	@drygopher -d -e "/mocks,/interfaces,/cmd,/host,'iface$$','drygopher$$','types$$'"
.PHONY: test
