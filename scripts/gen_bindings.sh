#!/bin/bash

# Generate go contract bindings.
# ref: https://geth.ethereum.org/docs/dapp/native-bindings

set -eou pipefail

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"

echo ""
echo "MXC_MONO_DIR: ${MXC_MONO_DIR}"
echo "MXC_GETH_DIR: ${MXC_GETH_DIR}"
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

cat ${MXC_MONO_DIR}/packages/protocol/artifacts/contracts/L1/MXCL1.sol/MXCL1.json |
	jq .abi |
	${ABIGEN_BIN} --abi - --type MXCL1Client --pkg bindings --out $DIR/../bindings/gen_mxc_l1.go

cat ${MXC_MONO_DIR}/packages/protocol/artifacts/contracts/L2/MXCL2.sol/MXCL2.json |
	jq .abi |
	${ABIGEN_BIN} --abi - --type MXCL2Client --pkg bindings --out $DIR/../bindings/gen_mxc_l2.go

git -C ${MXC_MONO_DIR} log --format="%H" -n 1 >./bindings/.githead

echo "🍻 Go contract bindings generated!"
