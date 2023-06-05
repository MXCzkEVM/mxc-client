package submitter

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/MXCzkEVM/mxc-client/bindings/encoding"
	"github.com/MXCzkEVM/mxc-client/pkg/rpc"
	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

var (
	errUnretryable = errors.New("unretryable")
)

// isSubmitProofTxErrorRetryable checks whether the error returned by a proof submission transaction
// is retryable.
func isSubmitProofTxErrorRetryable(err error, blockID *big.Int) bool {
	if strings.HasPrefix(err.Error(), "L1_NOT_SPECIAL_PROVER") ||
		!strings.HasPrefix(err.Error(), "L1_") {
		return true
	}

	log.Warn("🤷‍♂️ Unretryable proof submission error", "error", err, "blockID", blockID)
	return false
}

// getProveBlocksTxOpts creates a bind.TransactOpts instance using the given private key.
// Used for creating MxcL1.proveBlock and MxcL1.proveBlockInvalid transactions.
func getProveBlocksTxOpts(
	ctx context.Context,
	cli *ethclient.Client,
	chainID *big.Int,
	proverPrivKey *ecdsa.PrivateKey,
) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(proverPrivKey, chainID)
	if err != nil {
		return nil, err
	}
	gasTipCap, err := cli.SuggestGasTipCap(ctx)
	if err != nil {
		if rpc.IsMaxPriorityFeePerGasNotFoundError(err) {
			gasTipCap = rpc.FallbackGasTipCap
		} else {
			return nil, err
		}
	}

	opts.GasTipCap = gasTipCap

	return opts, nil
}

// sendTxWithBackoff tries to send the given proof submission transaction with a backoff policy.
func sendTxWithBackoff(
	ctx context.Context,
	cli *rpc.Client,
	blockID *big.Int,
	sendTxFunc func() (*types.Transaction, error),
) error {
	var isUnretryableError bool
	if err := backoff.Retry(func() error {
		if ctx.Err() != nil {
			return nil
		}

		tx, err := sendTxFunc()
		if err != nil {
			err = encoding.TryParsingCustomError(err)
			if isSubmitProofTxErrorRetryable(err, blockID) {
				log.Info("Retry sending MxcL1.proveBlock transaction", "blockID", blockID, "reason", err)
				return err
			}

			isUnretryableError = true
			return nil
		}

		if _, err := rpc.WaitReceipt(ctx, cli.L1, tx); err != nil {
			log.Warn("Failed to wait till transaction executed", "blockID", blockID, "txHash", tx.Hash(), "error", err)
			isUnretryableError = true
			return nil
			//return err
		}

		return nil
	}, backoff.NewExponentialBackOff()); err != nil {
		return fmt.Errorf("failed to send MxcL1.proveBlock transaction: %w", err)
	}

	if isUnretryableError {
		return errUnretryable
	}

	return nil
}
