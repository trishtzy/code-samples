# Paradex REST API Client

A Go client for interacting with the Paradex API, generated using go-swagger.

## Installation

```bash
go get github.com/tradeparadex/api
```

## Usage


### Basic Client Setup

```go
import (
    "github.com/tradeparadex/api/client"
)

// Create default client
api := client.Default

// Or create custom client with configuration
config := client.DefaultTransportConfig().
    WithHost("api.testnet.paradex.trade").
    WithBasePath("/v1").
    WithSchemes([]string{"https"})

api := client.NewHTTPClientWithConfig(nil, config)
```


### Authentication

The client supports both Basic Auth and Bearer Token authentication:

```go
import (
    "github.com/tradeparadex/api/client/authentication"
)

// Using Bearer Token
authClient := authentication.NewClientWithBearerToken(
    "api.testnet.paradex.trade",  // host
    "/v1",                        // basePath
    "https",                      // scheme
    "your-bearer-token"           // token
)

// Using Basic Auth
authClient := authentication.NewClientWithBasicAuth(
    "api.testnet.paradex.trade",  // host
    "/v1",                        // basePath
    "https",                      // scheme
    "username",                   // user
    "password"                    // password
)
```


### Example API Calls

#### Get Account Balance

```go
import (
    "context"
    "github.com/tradeparadex/api/client/account"
)

// Create bearer token authentication for subsequent calls
jwt := "JWT_TOKEN"
bearerAuth := httptransport.BearerToken(jwt)

// Create params
params := account.NewGetBalanceParams()

// Make API call
balance, err := api.Account.GetBalance(params, bearerAuth)
if err != nil {
    log.Fatal(err)
}

// Use the response
fmt.Printf("Balance: %+v\n", balance)
```

#### Place an Order

```go
import (
    "context"
    "github.com/tradeparadex/api/client/orders"
    "github.com/tradeparadex/api/models"
)

// Create bearer token authentication for subsequent calls
jwt := "JWT_TOKEN"
bearerAuth := httptransport.BearerToken(jwt)

// Create order params with authentication
orderParams := orders.NewOrdersNewParams()
	market := "ETH-USD-PERP"
	price := "1000"
	size := "1"
	now = time.Now().UnixMilli()
	orderSignature := auth.SignSNTypedData(auth.SignerParams{
		MessageType:       "order",
		DexAccountAddress: dexAccountAddress,
		DexPrivateKey:     dexPrivateKey,
		Params: map[string]interface{}{
			"timestamp": now,
			"market":    market,
			"side":      "BUY",
			"orderType": "LIMIT",
			"size":      size,
			"price":     price,
		},
	})

// Make authenticated API call
orderParams.SetParams(&models.RequestsOrderRequest{
		ClientID:           "client-id",
		Type:               struct{ models.ResponsesOrderType }{models.ResponsesOrderType("LIMIT")},
		Side:               struct{ models.ResponsesOrderSide }{models.ResponsesOrderSide("BUY")},
		Market:             &market,
		Price:              &price,
		Size:               &size,
		Signature:          &orderSignature,
		SignatureTimestamp: &now,
	})

	orderResp, err := api.Orders.OrdersNew(orderParams, bearerAuth)
	if err != nil {
		fmt.Printf("Failed to place order: %v\n", err)
	}
	fmt.Printf("Order placed successfully: %v\n", orderResp.Code())
```

## Available Services

The client provides access to various API services:

- Account
- Authentication
- Orders
- Trades
- Markets
- Vaults
- Insurance
- Liquidations
- Transfers
- System
- Algos
- Referrals

Each service has its own set of methods for interacting with different endpoints.

## Configuration

The default configuration uses:


```30:40:client/paradex_r_e_s_t_api_client.go
const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.testnet.paradex.trade"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}
```


You can customize these values using the `TransportConfig` methods when creating a new client.

## Documentation

For detailed API documentation, visit:
- Production API: https://docs.api.prod.paradex.trade/
- Testnet API: https://docs.api.testnet.paradex.trade/

For the generated Go client documentation, see the godoc for this package.