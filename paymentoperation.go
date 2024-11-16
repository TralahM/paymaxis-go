// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paymaxis

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/TralahM/paymaxis-go/internal/apijson"
	"github.com/TralahM/paymaxis-go/internal/requestconfig"
	"github.com/TralahM/paymaxis-go/option"
)

// PaymentOperationService contains methods and other services that help with
// interacting with the paymaxis API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentOperationService] method instead.
type PaymentOperationService struct {
	Options []option.RequestOption
}

// NewPaymentOperationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPaymentOperationService(opts ...option.RequestOption) (r *PaymentOperationService) {
	r = &PaymentOperationService{}
	r.Options = opts
	return
}

// Get a list of operations performed during payment processing sorted by time
// (most recent first)
func (r *PaymentOperationService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *Operation, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/payments/%s/operations", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Operation struct {
	Result []OperationResult `json:"result"`
	// HTTP status code
	Status    int64         `json:"status"`
	Timestamp string        `json:"timestamp" format:"ISO 8601"`
	JSON      operationJSON `json:"-"`
}

// operationJSON contains the JSON metadata for the struct [Operation]
type operationJSON struct {
	Result      apijson.Field
	Status      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Operation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationJSON) RawJSON() string {
	return r.raw
}

type OperationResult struct {
	// Operation Id
	ID int64 `json:"id"`
	// Operation end time
	Completed string `json:"completed" format:"ISO 8601"`
	// List of messages received from external APIs during operation processing
	IncomingMessages string `json:"incomingMessages"`
	// Operation performed during payment processing
	Operation OperationResultOperation `json:"operation"`
	// List of messages sent to external APIs during operation processing
	OutgoingMessages string `json:"outgoingMessages"`
	// Payment State
	PaymentState OperationResultPaymentState `json:"paymentState"`
	// Operation start time
	Started string              `json:"started" format:"ISO 8601"`
	JSON    operationResultJSON `json:"-"`
}

// operationResultJSON contains the JSON metadata for the struct [OperationResult]
type operationResultJSON struct {
	ID               apijson.Field
	Completed        apijson.Field
	IncomingMessages apijson.Field
	Operation        apijson.Field
	OutgoingMessages apijson.Field
	PaymentState     apijson.Field
	Started          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationResultJSON) RawJSON() string {
	return r.raw
}

// Operation performed during payment processing
type OperationResultOperation string

const (
	OperationResultOperationCreatePayment     OperationResultOperation = "CREATE_PAYMENT"
	OperationResultOperationCheckout          OperationResultOperation = "CHECKOUT"
	OperationResultOperationCancel            OperationResultOperation = "CANCEL"
	OperationResultOperationConfirmation      OperationResultOperation = "CONFIRMATION"
	OperationResultOperationCascading         OperationResultOperation = "CASCADING"
	OperationResultOperationRedirect          OperationResultOperation = "REDIRECT"
	OperationResultOperationContinue          OperationResultOperation = "CONTINUE"
	OperationResultOperationContinueAntiFraud OperationResultOperation = "CONTINUE_ANTI_FRAUD"
	OperationResultOperationDetectFraud       OperationResultOperation = "DETECT_FRAUD"
	OperationResultOperationDeposit           OperationResultOperation = "DEPOSIT"
	OperationResultOperationWithdrawal        OperationResultOperation = "WITHDRAWAL"
	OperationResultOperationRefund            OperationResultOperation = "REFUND"
	OperationResultOperationChargeback        OperationResultOperation = "CHARGEBACK"
	OperationResultOperationCheckState        OperationResultOperation = "CHECK_STATE"
	OperationResultOperationHandleWebhook     OperationResultOperation = "HANDLE_WEBHOOK"
	OperationResultOperationManualUpdate      OperationResultOperation = "MANUAL_UPDATE"
)

func (r OperationResultOperation) IsKnown() bool {
	switch r {
	case OperationResultOperationCreatePayment, OperationResultOperationCheckout, OperationResultOperationCancel, OperationResultOperationConfirmation, OperationResultOperationCascading, OperationResultOperationRedirect, OperationResultOperationContinue, OperationResultOperationContinueAntiFraud, OperationResultOperationDetectFraud, OperationResultOperationDeposit, OperationResultOperationWithdrawal, OperationResultOperationRefund, OperationResultOperationChargeback, OperationResultOperationCheckState, OperationResultOperationHandleWebhook, OperationResultOperationManualUpdate:
		return true
	}
	return false
}

// Payment State
type OperationResultPaymentState string

const (
	OperationResultPaymentStateCheckout  OperationResultPaymentState = "CHECKOUT"
	OperationResultPaymentStatePending   OperationResultPaymentState = "PENDING"
	OperationResultPaymentStateCancelled OperationResultPaymentState = "CANCELLED"
	OperationResultPaymentStateDeclined  OperationResultPaymentState = "DECLINED"
	OperationResultPaymentStateCompleted OperationResultPaymentState = "COMPLETED"
)

func (r OperationResultPaymentState) IsKnown() bool {
	switch r {
	case OperationResultPaymentStateCheckout, OperationResultPaymentStatePending, OperationResultPaymentStateCancelled, OperationResultPaymentStateDeclined, OperationResultPaymentStateCompleted:
		return true
	}
	return false
}
