### 1. Authorization with Auto Capture

**Flow Path:** `internal/flows/authorize_auto_capture/`

![Automatic capture](diagram.svg)

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
