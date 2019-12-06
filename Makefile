GO111MODULE = off

ci_test:
	go test -p 1 -timeout 30s -v ./...

TMP_PATH = "$(abspath ./tmp)"

tmp:
	@mkdir -p $(TMP_PATH)

tmp/gotest: tmp
	@export GO111MODULE=off
	@cd $(TMP_PATH) && go get github.com/rakyll/gotest && go build github.com/rakyll/gotest

test: tmp/gotest
	./tmp/gotest -p 1 -timeout 30s -v -run "${PATTERN}" ./...

test_bench:
	go test -p 1 -timeout 30s -bench="${PATTERN}" -benchmem -v -run "Benchmark" ./...

test_watch:
	watchexec -e go "clear && make test"

wasm/build/main.wasm: *.go
	mkdir -p wasm/build
	cp wasm/src/index.html wasm/build
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" wasm/build/
	GOOS=js GOARCH=wasm go build -o wasm/build/main.wasm

run_wasm_http: wasm/build/main.wasm
	cd wasm && go run ./webserver

clean:
	rm -rf ./tmp ./wasm/build
