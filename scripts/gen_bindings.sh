#!/bin/bash

# Generate go contract bindings.
# ref: https://geth.ethereum.org/docs/dapp/native-bindings

set -eou pipefail

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"

echo ""
echo "MXC_MONO_DIR: ${MXC_MONO_DIR}"
echo "MXC_GETH_DIR: ${MXC_GETH_DIR}"
echo "MXC_CORE_DIR: ${MXC_CORE_DIR}"
echo ""

cd ${MXC_GETH_DIR} &&
  make all &&
  cd -

cd ${MXC_MONO_DIR}/packages/protocol &&
  pnpm clean &&
  pnpm compile &&
  cd -

ABIGEN_BIN=$MXC_GETH_DIR/build/bin/abigen

echo ""
echo "Start generating go contract bindings..."
echo ""

cat ${MXC_MONO_DIR}/packages/protocol/out/MxcL1.sol/MxcL1.json |
	jq .abi |
	${ABIGEN_BIN} --abi - --type MxcL1Client --pkg bindings --out $DIR/../bindings/gen_mxc_l1.go

cat ${MXC_MONO_DIR}/packages/protocol/out/MxcL2.sol/MxcL2.json |
	jq .abi |
	${ABIGEN_BIN} --abi - --type MxcL2Client --pkg bindings --out $DIR/../bindings/gen_mxc_l2.go

cat ${MXC_MONO_DIR}/packages/protocol/out/MxcToken.sol/MxcToken.json |
  jq .abi |
  ${ABIGEN_BIN} --abi - --type MxcTokenClient --pkg bindings --out $DIR/../bindings/gen_mxc_token.go

cat ${MXC_CORE_DIR}/artifacts/contracts/token/LPWAN.sol/LPWAN.json |
	jq .abi |
	${ABIGEN_BIN} --abi - --type LPWANClient --pkg bindings --out $DIR/../bindings/gen_lpwan.go

git -C ${MXC_MONO_DIR} log --format="%H" -n 1 >./bindings/.githead

echo "🍻 Go contract bindings generated!"
