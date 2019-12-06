ci_test:
	go test -p 1 -timeout 30s -v ./...

TMP_PATH = "$(abspath ./tmp)"

tmp/gotest:
	@mkdir -p $(TMP_PATH)
	@cd $(TMP_PATH) && GO111MODULE=off go get github.com/rakyll/gotest && GO111MODULE=off go build github.com/rakyll/gotest

test: tmp/gotest
	./tmp/gotest -p 1 -timeout 30s -v -run "${PATTERN}" ./...

BENCH_PATTERN ?= Bench
test_bench:
	# ./tmp/gotest -p 1 -timeout 30s -gcflags=-m -bench=Bench -benchmem -v -run "Bench" ./...
	go test -p 1 -timeout 30s -bench="${BENCH_PATTERN}" -benchmem -v -run "${BENCH_PATTERN}" ./...

test_watch:
	watchexec -e go "clear && make test"
