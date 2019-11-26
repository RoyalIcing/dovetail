ci_test:
	go test -p 1 -timeout 30s -v ./...

TMP_PATH = "$(abspath ./tmp)"

tmp/gotest:
	@mkdir -p $(TMP_PATH)
	@cd $(TMP_PATH) && GO111MODULE=off go get github.com/rakyll/gotest && GO111MODULE=off go build github.com/rakyll/gotest

test_rich: tmp/gotest
	./tmp/gotest -p 1 -timeout 30s -v ./...

test: tmp/gotest
	./tmp/gotest -p 1 -timeout 30s -v -run "${PATTERN}" ./...

test_bench: tmp/gotest
	# ./tmp/gotest -p 1 -timeout 30s -gcflags=-m -bench=Bench -benchmem -v -run "Bench" ./...
	./tmp/gotest -p 1 -timeout 30s -bench=Bench -benchmem -v -run "Bench" ./...

test_watch:
	watchexec -e go "clear && make test"
