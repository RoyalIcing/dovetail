GO111MODULE = off

ifeq ($(OS),Windows_NT)
	GO := C:\Go\bin\go
else
	GO := go
endif

ci_test:
	$(GO) test -p 1 -timeout 30s -v ./...

TMP_PATH = "$(abspath ./tmp)"

tmp:
	@mkdir -p $(TMP_PATH)

tmp/gotest: tmp
	@export GO111MODULE=off
	@cd $(TMP_PATH) && $(GO) get github.com/rakyll/gotest && $(GO) build github.com/rakyll/gotest

test: tmp/gotest
	./tmp/gotest -p 1 -timeout 30s -v -run "${PATTERN}" ./...

test_bench:
	$(GO) test -p 1 -timeout 30s -bench="${PATTERN}" -benchmem -v -run "Benchmark" ./...

test_watch_receiver:
	-make test

test_watch:
	watchexec -e go "clear && make test_watch_receiver"

wasm/build/main.wasm: *.go
	mkdir -p wasm/build
	cp wasm/src/index.html wasm/build
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" wasm/build/
	GOOS=js GOARCH=wasm go build -o wasm/build/main.wasm

run_wasm_http: wasm/build/main.wasm
	cd wasm && go run ./webserver

clean:
	rm -rf ./tmp ./wasm/build
