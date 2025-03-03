package auth

import (
	"fmt"

	"github.com/dontpanicdao/caigo"
	"github.com/dontpanicdao/caigo/types"
	"github.com/tradeparadex/api/config"
)

func SignSNTypedData(messageType, dexAccountAddress, dexPrivateKey string) string {
	dexAccountAddressBN := types.HexToBN(dexAccountAddress)

	sc := caigo.StarkCurve{}
	message := typedMessage(messageType)
	typedData := verificationTypedData(messageType)
	domEnc, _ := typedData.GetTypedMessageHash("StarkNetDomain", typedData.Domain, sc)
	messageHash, _ := GnarkGetMessageHash(typedData, domEnc, dexAccountAddressBN, message, sc)
	r, s, _ := GnarkSign(messageHash, dexPrivateKey)

	return GetSignatureStr(r, s)
}

func typedMessage(messageType string) caigo.TypedMessage {
	switch messageType {
	case "onboarding":
		return &OnboardingPayload{Action: "Onboarding"}
	default:
		return nil
	}
}

func verificationTypedData(messageType string) *caigo.TypedData {
	var typedData *caigo.TypedData
	switch messageType {
	case "onboarding":
		fmt.Println("chainIDName", config.App.GetChainIDName())
		typedData, _ = NewVerificationTypedData(VerificationTypeOnboarding, config.App.GetChainIDName())
	}
	return typedData
}
