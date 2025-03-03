package auth

import (
	"github.com/dontpanicdao/caigo"
	"github.com/dontpanicdao/caigo/types"
	"github.com/tradeparadex/api/config"
)

const DEFAULT_EXPIRY_IN_SECONDS = int64(30)

type SignerParams struct {
	MessageType       string
	DexAccountAddress string
	DexPrivateKey     string
	TimestampStr      string
	ExpirationStr     string
}

func SignSNTypedData(signerParams SignerParams) string {
	dexAccountAddressBN := types.HexToBN(signerParams.DexAccountAddress)

	sc := caigo.StarkCurve{}
	message := typedMessage(signerParams)
	typedData := verificationTypedData(signerParams.MessageType)
	domEnc, _ := typedData.GetTypedMessageHash("StarkNetDomain", typedData.Domain, sc)
	messageHash, _ := GnarkGetMessageHash(typedData, domEnc, dexAccountAddressBN, message, sc)
	r, s, _ := GnarkSign(messageHash, signerParams.DexPrivateKey)

	return GetSignatureStr(r, s)
}

func typedMessage(signerParams SignerParams) caigo.TypedMessage {
	switch signerParams.MessageType {
	case "onboarding":
		return &OnboardingPayload{Action: "Onboarding"}
	case "auth":
		return &AuthPayload{
			Method:     "POST",
			Path:       "/v1/auth",
			Body:       "",
			Timestamp:  signerParams.TimestampStr,
			Expiration: signerParams.ExpirationStr,
		}
	default:
		return nil
	}
}

func verificationTypedData(messageType string) *caigo.TypedData {
	var typedData *caigo.TypedData
	var verificationType VerificationType
	switch messageType {
	case "onboarding":
		verificationType = VerificationTypeOnboarding
	case "auth":
		verificationType = VerificationTypeAuth
	}

	typedData, _ = NewVerificationTypedData(verificationType, config.App.GetChainIDName())
	return typedData
}
