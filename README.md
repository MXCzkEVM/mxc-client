# mxc-client

[![CI](https://github.com/MXCzkEVM/mxc-client/actions/workflows/test.yml/badge.svg)](https://github.com/MXCzkEVM/mxc-client/actions/workflows/test.yml)
[![Codecov](https://img.shields.io/codecov/c/github/MXCzkEVM/mxc-client?logo=codecov&token=OH6BJMVP6O)](https://codecov.io/gh/MXCzkEVM/mxc-client)

Mxc protocol's client software implementation in Go. Learn more about Mxc nodes with [the docs](https://geneva.mxc.com/docs/Designs/Supernode).

## Project structure

| Path                | Description                                                                                                                            |
| ------------------- |----------------------------------------------------------------------------------------------------------------------------------------|
| `bindings/`         | [Go contract bindings](https://geth.ethereum.org/docs/dapp/native-bindings) for mxc smart contracts, and few related utility functions |
| `cmd/`              | Main executable for this project                                                                                                       |
| `docs/`             | Documentation                                                                                                                          |
| `driver/`           | Driver sub-command                                                                                                                     |
| `integration_test/` | Scripts to do the integration testing of all client softwares                                                                          |
| `metrics/`          | Metrics related                                                                                                                        |
| `pkg/`              | Library code which used by all sub-commands                                                                                            |
| `proposer/`         | Proposer sub-command                                                                                                                   |
| `prover/`           | Prover sub-command                                                                                                                     |
| `scripts/`          | Helpful scripts                                                                                                                        |
| `testutils/`        | Test utils                                                                                                                             |
| `version/`          | Version information                                                                                                                    |

## Build the source

Building the `mxc-client` binary requires a Go compiler. Once installed, run:

```sh
make build
```

## Usage

Review all available sub-commands:

```sh
bin/mxc-client --help
```

Review each sub-command's command line flags:

```sh
bin/mxc-client <sub-command> --help
```

## Testing

Ensure you have Docker running, and pnpm installed.

Then, run the integration tests:

1. Start Docker locally
2. Perform a `pnpm install` in `mxc-mono/packages/protocol`
3. Replace `<PATH_TO_MXC_MONO_REPO>` and execute:

   ```bash
   MXC_MONO_DIR=<PATH_TO_MXC_MONO_REPO> \
   COMPILE_PROTOCOL=true \
     make test
   ```
