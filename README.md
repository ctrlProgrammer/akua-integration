# Payment Flows Documentation

This directory contains comprehensive payment flow implementations for the Akua payment integration. Each flow demonstrates different payment processing scenarios including authorization, capture, refund, and reversal operations.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Environment Setup](#environment-setup)
- [Flow Documentation](#flow-documentation)
  - [1. Authorization with Auto Capture](#1-authorization-with-auto-capture)
  - [2. Authorization with Manual Capture](#2-authorization-with-manual-capture)
  - [3. Authorization with Manual Reversal](#3-authorization-with-manual-reversal)
  - [4. Authorization with Auto Refund](#4-authorization-with-auto-refund)
- [Flow Diagrams](#flow-diagrams)
- [Architecture Overview](#architecture-overview)

## Prerequisites

Before running the tests, ensure you have:

- **Go 1.24.6** or higher installed
- **Akua API credentials** with appropriate permissions
- Access to the Akua API environment (sandbox or production)
- Network connectivity to the Akua API endpoints

## Environment Setup

Create a `.env` file in the project root with the following environment variables:

```env
AKUA_CLIENT_ID=your_client_id
AKUA_CLIENT_SECRET=your_client_secret
AKUA_AUDIENCE=https://api.akua.com
AKUA_ORGANIZATION_ID=your_organization_id
AKUA_COMMERCE_ID=your_commerce_id
```

These credentials are required for:

- OAuth token authentication
- API request authorization
- Merchant context identification

## Flow Documentation

### 1. Authorization with Auto Capture

**Flow Path:** `internal/flows/authorize_auto_capture/`

**Description:**
This flow demonstrates a payment authorization with automatic capture. When a payment is authorized with `CAPTURE_MODE_AUTOMATIC`, the system automatically captures the authorized amount in the same transaction. This is the most common flow for immediate payment processing.

**Key Characteristics:**

- **Capture Mode:** `AUTOMATIC`
- **Authorization Status:** `APPROVED`
- **Payment Status:** Authorized and captured automatically
- **Use Case:** Standard e-commerce transactions where immediate payment is required

**Flow Steps:**

1. **Initialization**

   - Load environment variables from `.env`
   - Create Akua client instance
   - Load JWT token via OAuth `/oauth/token` endpoint

2. **Authorization Request**

   - Prepare `AuthorizeRequest` with:
     - Amount: `{Currency: "USD", Value: 100}`
     - Intent: `"authorization"`
     - Capture Mode: `"AUTOMATIC"`
     - Merchant ID from client configuration
     - Card instrument details
   - Send POST request to `/v1/authorizations`
   - Process authorization response

3. **Verification**
   - Assert authorization response is not nil
   - Verify payment ID is present
   - Validate transaction status is `"APPROVED"`

**Expected Result:**

- Authorization is immediately approved and captured
- Transaction status: `APPROVED`
- Payment is ready for fulfillment

**Test Function:**

```go
Test_Authorize_AutoCapture_Success
```

---

### 2. Authorization with Manual Capture

**Flow Path:** `internal/flows/authorize_manual_cature/`

**Description:**
This flow demonstrates a payment authorization with manual capture. The payment is authorized first, and the capture is performed separately using the `Capture()` method. This pattern is useful for scenarios where you need to hold funds before finalizing a transaction (e.g., order confirmation, inventory verification).

**Key Characteristics:**

- **Capture Mode:** `MANUAL`
- **Initial Status:** `AUTHORIZED` (not yet captured)
- **After Capture:** `CAPTURED`
- **Use Case:** Pre-authorization scenarios, order processing with delayed capture

**Flow Steps:**

1. **Initialization**

   - Load environment variables
   - Create Akua client and providers (AuthorizationProvider, PaymentsProvider)
   - Load JWT token

2. **Authorization Request**

   - Prepare `AuthorizeRequest` with `Capture.Mode: "MANUAL"`
   - Send POST request to `/v1/authorizations`
   - Receive authorization response with payment ID

3. **Verify Payment State (Before Capture)**

   - Retrieve payment by ID using `GetPaymentById()`
   - Verify payment status is `"AUTHORIZED"`
   - Verify capture mode is `"MANUAL"`
   - Confirm only AUTHORIZATION transaction exists

4. **Manual Capture**

   - Prepare `CaptureRequest` with payment ID
   - Send POST request to `/v1/payments/{id}/captures`
   - Process capture response

5. **Verify Final Payment State (After Capture)**
   - Retrieve payment again to verify state
   - Confirm payment status changed to `"CAPTURED"`
   - Verify two transactions exist:
     - AUTHORIZATION transaction
     - CAPTURE transaction

**Expected Result:**

- Payment initially in `AUTHORIZED` state
- After manual capture, payment transitions to `CAPTURED`
- Payment contains both AUTHORIZATION and CAPTURE transactions

**Test Function:**

```go
Test_Authorize_ManualCapture_Success
```

---

### 3. Authorization with Manual Reversal

**Flow Path:** `internal/flows/authorize_manual_reversal/`

**Description:**
This flow demonstrates canceling an authorized but not yet captured payment using reversal. Reversal is used to cancel an authorization that was created with manual capture mode. Unlike refund (which requires a captured payment), reversal cancels the authorization before capture.

**Key Characteristics:**

- **Capture Mode:** `MANUAL`
- **Initial Status:** `AUTHORIZED`
- **After Reversal:** `CANCELLED`
- **Use Case:** Canceling orders before capture, freeing authorized funds

**Flow Steps:**

1. **Initialization**

   - Load environment variables
   - Create Akua client and providers
   - Load JWT token

2. **Authorization Request**

   - Prepare `AuthorizeRequest` with `Capture.Mode: "MANUAL"`
   - Send POST request to `/v1/authorizations`
   - Payment is authorized but not captured

3. **Verify Payment State (Before Reversal)**

   - Retrieve payment by ID
   - Verify payment status is `"AUTHORIZED"`
   - Verify capture mode is `"MANUAL"`

4. **Reversal Request**

   - Call `Reversal()` method with payment ID
   - Send POST request to `/v1/payments/{id}/reversals`
   - Process reversal response

5. **Verify Final Payment State (After Reversal)**
   - Retrieve payment to verify cancellation
   - Confirm payment status changed to `"CANCELLED"`
   - Verify two transactions exist:
     - AUTHORIZATION transaction
     - REVERSAL transaction

**Expected Result:**

- Payment initially in `AUTHORIZED` state
- After reversal, payment status is `"CANCELLED"`
- Payment contains AUTHORIZATION and REVERSAL transactions
- Funds are released back to the customer

**Test Function:**

```go
Test_Authorize_ManualCapture_Success
```

**Important Notes:**

- Reversal can only be performed on authorized but not captured payments
- Once a payment is captured, you must use refund instead of reversal
- Reversal cancels the authorization and releases the hold on funds

---

### 4. Authorization with Auto Refund

**Flow Path:** `internal/flows/authorize_auto_refund/`

**Description:**
This flow demonstrates refunding a payment that was authorized and automatically captured. Refund is used to return money to a customer for a payment that has already been captured. This is different from reversal, which cancels an authorization before capture.

**Key Characteristics:**

- **Capture Mode:** `AUTOMATIC`
- **Initial Status:** `AUTHORIZED` (auto-captured)
- **After Refund:** `REFUNDED`
- **Use Case:** Customer returns, order cancellations after fulfillment, refund requests

**Flow Steps:**

1. **Initialization**

   - Load environment variables
   - Create Akua client and providers
   - Load JWT token

2. **Authorization Request (Auto Capture)**

   - Prepare `AuthorizeRequest` with `Capture.Mode: "AUTOMATIC"`
   - Send POST request to `/v1/authorizations`
   - Payment is automatically authorized and captured

3. **Verify Payment State (Before Refund)**

   - Retrieve payment by ID
   - Verify payment status is `"AUTHORIZED"`
   - Verify capture mode is `"AUTOMATIC"`
   - Note: Payment may show as captured depending on timing

4. **Refund Request**

   - Call `Refund()` method with payment ID
   - Send POST request to `/v1/payments/{id}/refund`
   - Process refund response

5. **Verify Final Payment State (After Refund)**
   - Retrieve payment to verify refund
   - Confirm payment status changed to `"REFUNDED"`
   - Verify multiple transactions exist:
     - AUTHORIZATION transaction
     - CAPTURE transaction (automatic)
     - REFUND transaction

**Expected Result:**

- Payment initially authorized and automatically captured
- After refund, payment status is `"REFUNDED"`
- Payment contains AUTHORIZATION, CAPTURE, and REFUND transactions
- Funds are returned to the customer

**Test Function:**

```go
Test_Authorize_AutomaticCapture_Refund_Success
```

**Important Notes:**

- Refund requires a captured payment (unlike reversal which requires an authorized but uncaptured payment)
- Refund returns the captured amount to the customer
- Multiple refunds may be possible for partial refund scenarios

---

## Flow Diagrams

Each flow has a detailed PlantUML sequence diagram that visualizes the complete interaction between components. The diagrams are located in each flow's directory:

### Available Diagrams

1. **Auto Capture Flow Diagram**

   - Path: `internal/flows/authorize_auto_capture/flow_diagram.puml`

2. **Manual Capture Flow Diagram**

   - Path: `internal/flows/authorize_manual_cature/flow_diagram.puml`

3. **Manual Reversal Flow Diagram**

   - Path: `internal/flows/authorize_manual_reversal/flow_diagram.puml`

4. **Auto Refund Flow Diagram**
   - Path: `internal/flows/authorize_auto_refund/flow_diagram.puml`

### Viewing Flow Diagrams

#### Option 1: VS Code Extension (Recommended)

1. Install the **PlantUML** extension for VS Code
2. Open any `.puml` file
3. Press `Alt+D` or right-click → "Preview PlantUML"

#### Option 2: PlantUML CLI

1. Install PlantUML CLI:

   ```bash
   # Using Homebrew (macOS)
   brew install plantuml

   # Using package manager (Linux)
   sudo apt-get install plantuml
   ```

2. Generate diagram:

   ```bash
   plantuml internal/flows/authorize_auto_capture/flow_diagram.puml
   ```

3. This generates a PNG image file in the same directory

#### Option 3: Online Viewer

1. Copy the contents of any `.puml` file
2. Visit http://www.plantuml.com/plantuml/uml/
3. Paste the content to view the diagram

#### Option 4: GitHub Integration

If you use a GitHub Actions workflow or CI/CD pipeline, you can automatically generate images from PlantUML files. Some repositories use tools like [PlantUML Action](https://github.com/marketplace/actions/plantuml-action) to render diagrams.

---

## Architecture Overview

### Flow Structure

All flows follow a consistent structure:

```
internal/flows/
├── authorize_auto_capture/
│   ├── flow_test.go          # Test implementation
│   └── flow_diagram.puml     # Sequence diagram
├── authorize_manual_cature/
│   ├── flow_test.go
│   └── flow_diagram.puml
├── authorize_manual_reversal/
│   ├── flow_test.go
│   └── flow_diagram.puml
└── authorize_auto_refund/
    ├── flow_test.go
    └── flow_diagram.puml
```

### Common Components

All flows utilize:

1. **Akua Client** (`internal/adapters/akua`)

   - HTTP client wrapper
   - JWT token management
   - OAuth authentication

2. **Authorization Provider** (`internal/adapters/akua/authorization`)

   - `Authorize()` - Create payment authorization
   - `Capture()` - Capture an authorized payment
   - `Reversal()` - Reverse an authorized payment
   - `Refund()` - Refund a captured payment

3. **Payments Provider** (`internal/adapters/akua/payments`)
   - `GetPaymentById()` - Retrieve payment details
   - Payment state verification

### Flow Initialization Pattern

All flows use a common initialization pattern:

```go
func InitializePaymentFlow(ctx context.Context) (*Client, *AuthorizationProvider, *PaymentsProvider, error) {
    // 1. Load environment variables
    InitializeEnvVariables()

    // 2. Create Akua client
    akuaClient, err := adapters_akua.NewClient()

    // 3. Load JWT token
    err = akuaClient.LoadJwtToken()

    // 4. Create providers
    return akuaClient, NewAuthorizationProvider(), NewPaymentsProvider(), nil
}
```

### Error Handling

All flows implement comprehensive error handling:

- Environment variable validation
- JWT token validation
- HTTP request error handling
- Response parsing validation
- Status code verification

---

## Flow Comparison Matrix

| Flow                | Capture Mode | Initial Status          | Final Action | Final Status | Use Case                     |
| ------------------- | ------------ | ----------------------- | ------------ | ------------ | ---------------------------- |
| **Auto Capture**    | `AUTOMATIC`  | `APPROVED`              | N/A          | `APPROVED`   | Immediate payment processing |
| **Manual Capture**  | `MANUAL`     | `AUTHORIZED`            | Capture      | `CAPTURED`   | Delayed capture scenarios    |
| **Manual Reversal** | `MANUAL`     | `AUTHORIZED`            | Reversal     | `CANCELLED`  | Cancel before capture        |
| **Auto Refund**     | `AUTOMATIC`  | `AUTHORIZED`/`CAPTURED` | Refund       | `REFUNDED`   | Return funds after capture   |

---

## Transaction Types Reference

Throughout the flows, you'll encounter these transaction types:

- **AUTHORIZATION** - Initial payment authorization
- **CAPTURE** - Capture of authorized funds
- **REVERSAL** - Cancellation of authorization (before capture)
- **REFUND** - Return of funds (after capture)

---

## Best Practices

1. **Environment Management**

   - Never commit `.env` files to version control
   - Use different credentials for sandbox and production
   - Rotate credentials regularly

2. **Error Handling**

   - Always check error responses
   - Implement retry logic for transient failures
   - Log errors for debugging

3. **Idempotency**

   - All API requests use idempotency keys
   - Replay requests safely without duplicate charges
   - Keys are generated using transaction details and timestamp

4. **Testing**

   - Run tests in sandbox environment first
   - Verify test results match expected states
   - Check transaction logs in Akua dashboard

5. **Monitoring**
   - Monitor payment status transitions
   - Track transaction counts per payment
   - Verify final payment states

---

## Troubleshooting

### Common Issues

1. **JWT Token Errors**

   - Verify `AKUA_CLIENT_ID` and `AKUA_CLIENT_SECRET` are correct
   - Check `AKUA_AUDIENCE` matches your environment
   - Ensure credentials have not expired

2. **Authorization Failures**

   - Verify card details are valid test cards
   - Check merchant ID matches your account
   - Confirm sufficient permissions in API credentials

3. **Capture/Refund/Reversal Errors**

   - Ensure payment is in correct state (authorized for capture, captured for refund)
   - Verify payment ID is correct
   - Check payment has not already been processed

4. **Test Failures**
   - Verify `.env` file is in project root
   - Check network connectivity to Akua API
   - Review verbose output (`-v` flag) for detailed error messages

---

## Additional Resources

- [Akua API Documentation](https://docs.akua.com) - Official API reference
- [Integration Architecture](./../adapters/akua/integration_structure.puml) - System architecture diagram
- [Authorization Cases](./../adapters/akua/authorization/cases.puml) - Error case scenarios

---

## Support

For issues or questions:

1. Review the flow diagrams for visual flow representation
2. Check test output with `-v` flag for detailed logs
3. Verify environment configuration
4. Consult Akua API documentation

---

**Last Updated:** Generated from codebase analysis
**Go Version:** 1.24.6
**Test Framework:** Go testing package with testify assertions
