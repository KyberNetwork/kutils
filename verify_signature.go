package kutils

import (
	"errors"
	"fmt"

	awstime "github.com/aws/smithy-go/time"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"regexp"
	"strings"

	"time"
)

var (
	DefaultAuthRegexp = regexp.MustCompile(`Click sign to (.+?) at Kyberswap.com without logging in.
This request won’t trigger any blockchain transaction or cost any gas fee. Expires in 7 days. 

Issued at: (.+)`)
	DefaultAuthExpiry = 168 * time.Hour
)

type EIP191 struct {
	Msg       string
	Signature string
	Address   string
}

func VerifyEIP191SignatureWithDefaults(eipChallenge *EIP191) error {
	return VerifyEIP191Signature(eipChallenge, DefaultAuthRegexp, DefaultAuthExpiry)
}

func VerifyEIP191Signature(eipChallenge *EIP191, authMessageRegexp *regexp.Regexp, authExpiry time.Duration) error {
	decodedSig, err := hexutil.Decode(eipChallenge.Signature)
	if err != nil {
		return err
	}

	if decodedSig[64] < 27 {
		if !hasValidLastByte(decodedSig) {
			return errors.New("invalid last byte")
		}
	} else {
		decodedSig[64] -= 27 // shift byte?
	}
	signHash := signEIP191(eipChallenge.Msg)
	recoveredPublicKey, err := crypto.Ecrecover(signHash.Bytes(), decodedSig)
	if err != nil {
		return err
	}

	secp256k1RecoveredPublicKey, err := crypto.UnmarshalPubkey(recoveredPublicKey)
	if err != nil {
		return err
	}

	recoveredAddress := crypto.PubkeyToAddress(*secp256k1RecoveredPublicKey).Hex()

	if !hasMatchingAddress(eipChallenge.Address, recoveredAddress) {
		return errors.New("invalid user")
	}

	if authMessageRegexp != nil {
		matched := authMessageRegexp.FindStringSubmatch(eipChallenge.Msg)
		if len(matched) < 3 {
			return errors.New("authMessageRegexp: invalid message")
		}

		timestamp, err := awstime.ParseDateTime(matched[2])
		if err != nil {
			return err
		}
		if timestamp.Add(authExpiry).Before(time.Now()) {
			return errors.New("expired message")
		}
	}

	return nil
}

func hasMatchingAddress(knownAddress string, recoveredAddress string) bool {
	return strings.EqualFold(strings.ToLower(knownAddress), strings.ToLower(recoveredAddress))
}

func hasValidLastByte(sig []byte) bool {
	return sig[64] == 0 || sig[64] == 1
}

func signEIP191(message string) common.Hash {
	msg := []byte(message)
	formattedMsg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)
	return crypto.Keccak256Hash([]byte(formattedMsg))
}
