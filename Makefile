CURRENT_OS ?= windows

#PREVENTS STRANGE BEHAVIOR IF IN THE PROJECT APEEAR FILE WITH NAME FROM PHONY LIST 
.PHONY: build build-opt run clean 

run:
	go run cmd/main.go

build: 
	go build cmd/main.go

build-opt:
ifeq ($(CURRENT_OS),windows)
	go build -trimpath -ldflags="-s -w -extldflags=-static" -o ./main_opt.exe cmd/main.go
else
	go build -trimpath -ldflags="-s -w -extldflags=-static" -o ./main_opt cmd/main.go
endif


clean: clean-$(CURRENT_OS)

clean-windows:
	go clean
	del *.exe
	@echo "Windows cleanup done!"


clean-unix:
	go clean
	rm -f main
	@echo "Linux cleanup done!"