# Payment Flow - Authorization Command

## Overview

This directory contains the payment flow command that handles authorization requests to the Akua API.

## Authorization Flow Diagram

See `authorization_flow.puml` for the complete sequence diagram showing all steps from command execution to API response.

## Prerequisites

1. Set up environment variables in `.env` file:
   ```
   AKUA_CLIENT_ID=your_client_id
   AKUA_CLIENT_SECRET=your_client_secret
   AKUA_AUDIENCE=https://api.akua.com
   AKUA_ORGANIZATION_ID=your_org_id
   AKUA_COMMERCE_ID=your_commerce_id
   ```

## Commands to Run

### Option 1: Run via Main Entry Point

Create a `main.go` in the project root and add:

```go
package main

import (
    "akua-project/cmd/flows"
    "context"
    "log"
)

func main() {
    ctx := context.Background()

    // Initialize the payment flow (loads JWT token)
    err := flows.InitializePaymentFlow(ctx)
    if err != nil {
        log.Fatal(err)
    }

    // Execute payment flow
    err = flows.PaymentFlow(ctx)
    if err != nil {
        log.Fatal(err)
    }
}
```

Then run:

```bash
go run main.go
```

### Option 2: Run as a Test

You can also test the authorization flow using the test file:

```bash
go test ./internal/adapters/akua/authorization/... -v
```

### Option 3: Build and Run Binary

```bash
# Build the application
go build -o akua-payment cmd/flows/main.go

# Run the binary
./akua-payment
```

## Authorization Flow Steps

1. **Initialization**

   - Create Akua client
   - Load JWT token from `/oauth/token` endpoint

2. **Authorization Request**

   - Prepare authorization request with:
     - Amount (currency, value)
     - Intent (authorization/pre-authorization)
     - Merchant ID
     - Instrument (card details)
     - Capture mode (automatic/manual)
   - Create AuthorizationProvider
   - Call `Authorize()` method

3. **API Communication**

   - Validate JWT token
   - Marshal request to JSON
   - Set HTTP headers (Authorization, Idempotency-Key, Content-Type)
   - POST to `/v1/authorizations`
   - Handle response (201 Created or errors)

4. **Response Processing**
   - Unmarshal JSON response
   - Check authorization status (approved/declined/rejected)
   - Return result to caller

## View the Flow Diagram

To view the PlantUML diagram:

1. **Using PlantUML CLI:**

   ```bash
   plantuml cmd/flows/authorization_flow.puml
   ```

2. **Using VS Code Extension:**

   - Install "PlantUML" extension
   - Open `authorization_flow.puml`
   - Press `Alt+D` or right-click → "Preview PlantUML"

3. **Using Online Viewer:**
   - Copy contents of `authorization_flow.puml`
   - Paste at http://www.plantuml.com/plantuml/uml/

## Example Authorization Request

```go
request := AuthorizeRequest{
    Amount: instruments.AmountObject{
        Currency: "USD",
        Value:    100.00,
    },
    Intent:     INTENT_AUTHORIZE,
    MerchantId: akuaClient.GetMerchantId(),
    Capture: instruments.CaptureObject{
        Mode: instruments.CAPTURE_MODE_AUTOMATIC,
    },
    Instrument: instruments.InstrumentObject{
        Type: "CARD",
        Card: instruments.Instrument{
            Number:          "4111111111111111",
            CVV:             "123",
            ExpirationMonth: "12",
            ExpirationYear:  "2026",
            HolderName:      "John Doe",
        },
    },
}
```
