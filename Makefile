CURRENT_OS ?= windows
CGO ?= 0
EXEC_FILE_NAME ?= main
TEMPLATE_PATH ?= .

#PREVENTS STRANGE BEHAVIOR IF IN THE PROJECT APEEAR FILE WITH NAME FROM PHONY LIST 
.PHONY: build build-opt run clean test

run:
	go run cmd/main.go

build: 
	go build -o $(EXEC_FILE_NAME) cmd/main.go

build-opt:
ifeq ($(CURRENT_OS),windows)
	set GOOS=windows
	set CGO_ENABLED=$(CGO)
	go build -trimpath -ldflags="-s -w -extldflags=-static" -o ./$(EXEC_FILE_NAME)_opt.exe cmd/main.go
else
	export GOOS=linux
	CGO_ENABLED=$(CGO) go build -trimpath -ldflags="-s -w -extldflags=-static" -o ./$(EXEC_FILE_NAME)_opt cmd/main.go
endif

test:
	set TEMPLATE_PATH=$(TEMPLATE_PATH)&& go test ./...

clean: clean-$(CURRENT_OS)

clean-windows:
	go clean
	del *.exe
	@echo "Windows cleanup done!"


clean-unix:
	go clean
	rm -f main
	@echo "Linux cleanup done!"