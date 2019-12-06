ci_test:
	go test -p 1 -timeout 30s -v ./...

TMP_PATH = "$(abspath ./tmp)"

tmp/gotest:
	@mkdir -p $(TMP_PATH)
	@cd $(TMP_PATH) && GO111MODULE=off go get github.com/rakyll/gotest && GO111MODULE=off go build github.com/rakyll/gotest

test: tmp/gotest
	./tmp/gotest -p 1 -timeout 30s -v -run "${PATTERN}" ./...

test_bench:
	go test -p 1 -timeout 30s -bench="${PATTERN}" -benchmem -v -run "Benchmark" ./...

test_watch:
	watchexec -e go "clear && make test"
