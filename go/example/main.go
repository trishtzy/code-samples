package main

import (
	"fmt"
	"log"

	"github.com/go-openapi/runtime"
	"github.com/tradeparadex/api/auth"
	"github.com/tradeparadex/api/client"
	"github.com/tradeparadex/api/client/system"
	"github.com/tradeparadex/api/config"
)

func main() {
	config.LoadEnv()

	// Create a client configuration
	clientConfig := client.DefaultTransportConfig().
		WithHost("api.nightly.paradex.trade").
		WithBasePath("/v1").
		WithSchemes([]string{"https"})

	// Create the main API client
	api := client.NewHTTPClientWithConfig(nil, clientConfig)

	// Step 1: Get system config
	configParams := system.NewGetSystemConfigParams()
	var configResp *system.GetSystemConfigOK
	handleAPICall(func() error {
		var err error
		configResp, err = api.System.GetSystemConfig(configParams)
		return err
	}, "Failed to get system config", false)

	// Step 1.1: Generate ethereum account
	_, ethereumPublicKey := auth.GetEthereumAccount()
	fmt.Printf("Ethereum public key: %s\n", ethereumPublicKey)

	// Step 1.2: Generate paradex L2 account
	dexPrivateKey, dexPublicKey, dexAccountAddress := auth.GenerateParadexAccountOld(
		*configResp.GetPayload(), config.App.EthereumPrivateKey)
	fmt.Printf("Paradex L2 account: %s\n", dexAccountAddress)
	fmt.Printf("Paradex L2 private key: %s\n", dexPrivateKey)
	fmt.Printf("Paradex L2 public key: %s\n", dexPublicKey)

	// // Step 2: Onboarding
	// onboardingParams := authentication.NewOnboardingParams()
	// starknetSignature := signer.SignTypedData("onboarding", dexAccountAddress)
	// onboardingParams.SetPARADEXETHEREUMACCOUNT(ethereumPublicKey)
	// onboardingParams.SetPARADEXSTARKNETACCOUNT(dexAccountAddress)
	// onboardingParams.SetPARADEXSTARKNETSIGNATURE(starknetSignature)

	// // Set necessary parameters for onboarding
	// var onboardingResp *authentication.OnboardingOK
	// handleAPICall(func() error {
	// 	var err error
	// 	onboardingResp, err = api.Authentication.Onboarding(onboardingParams)
	// 	return err
	// }, "Onboarding failed", true)
	// fmt.Printf("Onboarding successful: %v\n", onboardingResp.Code())

	// // Step 3: Authenticate to get JWT token
	// authParams := authentication.NewAuthParams()
	// // Set necessary headers or body parameters for authentication
	// var authResp *authentication.AuthOK
	// handleAPICall(func() error {
	// 	var err error
	// 	authResp, err = api.Authentication.Auth(authParams)
	// 	return err
	// }, "Authentication failed", true)
	// jwt := authResp.Payload.JwtToken

	// // Create bearer token authentication for subsequent calls
	// bearerAuth := httptransport.BearerToken(jwt)

	// // Example 1: Get account balance
	// balanceParams := account.NewGetBalanceParams()
	// var balance *account.GetBalanceOK
	// handleAPICall(func() error {
	// 	var err error
	// 	balance, err = api.Account.GetBalance(balanceParams, bearerAuth)
	// 	return err
	// }, "Failed to get balance", false)
	// fmt.Println("Account Balance:")
	// fmt.Printf("  Balance details: %v\n", balance.Payload)
	// fmt.Println()

	// // Example 2: Get markets information
	// marketsParams := markets.NewGetMarketsParams()
	// var marketsResp *markets.GetMarketsOK
	// handleAPICall(func() error {
	// 	var err error
	// 	marketsResp, err = api.Markets.GetMarkets(marketsParams)
	// 	return err
	// }, "Failed to get markets", false)
	// fmt.Println("Available Markets:")
	// for _, market := range marketsResp.GetPayload().Results {
	// 	fmt.Printf("  Market details: %v\n", market)
	// }
	// fmt.Println()
}

// handleAPICall wraps API calls with consistent error handling
func handleAPICall(fn func() error, errMsg string, fatal bool) {
	if err := fn(); err != nil {
		switch e := err.(type) {
		case *runtime.APIError:
			fmt.Printf("API Error: %v (code: %d)\n", e.Error(), e.Code)
		default:
			fmt.Printf("Error: %v\n", err)
		}
		if fatal {
			log.Fatal(errMsg)
		}
		return
	}
}
