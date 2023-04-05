GIT_COMMIT := $(shell git rev-parse HEAD)
GIT_DATE := $(shell git show -s --format='%ct')

LD_FLAGS_ARGS +=-X github.com/MXCzkEVM/mxc-client/version.GitCommit=$(GIT_COMMIT)
LD_FLAGS_ARGS +=-X github.com/MXCzkEVM/mxc-client/version.GitDate=$(GIT_DATE)

LD_FLAGS := -ldflags "$(LD_FLAGS_ARGS)"

build:
	GO111MODULE=on go build -v $(LD_FLAGS) -o bin/mxc-client cmd/main.go

clean:
	@rm -rf bin/*

lint:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2 \
	&& golangci-lint run

test:
	@MXC_MONO_DIR=${MXC_MONO_DIR} \
	COMPILE_PROTOCOL=${COMPILE_PROTOCOL} \
	RUN_TESTS=true \
		./integration_test/entrypoint.sh

dev_net:
	@MXC_MONO_DIR=${MXC_MONO_DIR} \
	COMPILE_PROTOCOL=${COMPILE_PROTOCOL} \
		./integration_test/entrypoint.sh

gen_bindings:
	@MXC_MONO_DIR=${MXC_MONO_DIR} \
	MXC_GETH_DIR=${MXC_GETH_DIR} \
		./scripts/gen_bindings.sh

.PHONY: build \
				clean \
				lint \
				test \
				dev_net \
				gen_bindings
