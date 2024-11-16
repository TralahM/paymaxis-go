// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paymaxis

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/TralahM/paymaxis-go/internal/apijson"
	"github.com/TralahM/paymaxis-go/internal/param"
	"github.com/TralahM/paymaxis-go/internal/requestconfig"
	"github.com/TralahM/paymaxis-go/option"
)

// SubscriptionService contains methods and other services that help with
// interacting with the paymaxis API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionService] method instead.
type SubscriptionService struct {
	Options []option.RequestOption
}

// NewSubscriptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubscriptionService(opts ...option.RequestOption) (r *SubscriptionService) {
	r = &SubscriptionService{}
	r.Options = opts
	return
}

// Find Subscription by Id
func (r *SubscriptionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/subscriptions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Patch Subscription
func (r *SubscriptionService) Update(ctx context.Context, id string, body SubscriptionUpdateParams, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/subscriptions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type Subscription struct {
	Result SubscriptionResult `json:"result"`
	// HTTP status code
	Status    int64            `json:"status"`
	Timestamp string           `json:"timestamp" format:"ISO 8601"`
	JSON      subscriptionJSON `json:"-"`
}

// subscriptionJSON contains the JSON metadata for the struct [Subscription]
type subscriptionJSON struct {
	Result      apijson.Field
	Status      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Subscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionJSON) RawJSON() string {
	return r.raw
}

type SubscriptionResult struct {
	// Subscription Id
	ID string `json:"id"`
	// The amount to be used for subsequent payments
	Amount float64 `json:"amount"`
	// Date and time the subscription was created
	CreateTime string `json:"createTime" format:"ISO 8601 (YYYY-MM-DD'T'HH24:MI:SS)"`
	// Payment currency
	Currency string `json:"currency" format:"ISO 4217 code for FIAT currencies or cryptocurrency symbol"`
	// Id of the customer from initial payment
	CustomerReferenceID string `json:"customerReferenceId"`
	// List of payments automatically generated for this subscription
	Cycles []SubscriptionResultCycle `json:"cycles"`
	// Description for subsequent recurring payments
	Description string `json:"description"`
	// The number of intervals after which a subscriber is billed. For example, if the
	// frequencyUnit is DAY with an frequency of 2, the subscription is billed once
	// every two days.
	Frequency int64 `json:"frequency"`
	// The interval at which the subscription is billed. Use 'MINUTE' for testing
	// purposes only.
	FrequencyUnit SubscriptionResultFrequencyUnit `json:"frequencyUnit"`
	// Token that is used to continue the recurring chain
	RecurringToken string `json:"recurringToken"`
	// Required number of subsequent recurring payments. Unlimited if value is not
	// specified.
	RequestedNumberOfCycles int64 `json:"requestedNumberOfCycles"`
	// Retry strategy for subscription. If not specified, the subscription is canceled
	// after the first failed payment attempt.
	RetryStrategy SubscriptionResultRetryStrategy `json:"retryStrategy"`
	// Date and time of the 1st cycle
	StartTime string `json:"startTime" format:"ISO 8601 (YYYY-MM-DD'T'HH24:MI:SS)"`
	// Subscription state
	State SubscriptionResultState `json:"state"`
	JSON  subscriptionResultJSON  `json:"-"`
}

// subscriptionResultJSON contains the JSON metadata for the struct
// [SubscriptionResult]
type subscriptionResultJSON struct {
	ID                      apijson.Field
	Amount                  apijson.Field
	CreateTime              apijson.Field
	Currency                apijson.Field
	CustomerReferenceID     apijson.Field
	Cycles                  apijson.Field
	Description             apijson.Field
	Frequency               apijson.Field
	FrequencyUnit           apijson.Field
	RecurringToken          apijson.Field
	RequestedNumberOfCycles apijson.Field
	RetryStrategy           apijson.Field
	StartTime               apijson.Field
	State                   apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SubscriptionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionResultJSON) RawJSON() string {
	return r.raw
}

type SubscriptionResultCycle struct {
	// Payment amount
	Amount float64 `json:"amount"`
	// Payment Id
	PaymentID string `json:"paymentId"`
	// Payment State
	PaymentState SubscriptionResultCyclesPaymentState `json:"paymentState"`
	// Sequence number of the cycle
	Sequence int64 `json:"sequence"`
	// Date and time when this cycle was supposed to be created according to the
	// schedule
	StartTime string `json:"startTime" format:"ISO 8601 (YYYY-MM-DD'T'HH24:MI:SS)"`
	// Cycle type
	Type SubscriptionResultCyclesType `json:"type"`
	JSON subscriptionResultCycleJSON  `json:"-"`
}

// subscriptionResultCycleJSON contains the JSON metadata for the struct
// [SubscriptionResultCycle]
type subscriptionResultCycleJSON struct {
	Amount       apijson.Field
	PaymentID    apijson.Field
	PaymentState apijson.Field
	Sequence     apijson.Field
	StartTime    apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionResultCycle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionResultCycleJSON) RawJSON() string {
	return r.raw
}

// Payment State
type SubscriptionResultCyclesPaymentState string

const (
	SubscriptionResultCyclesPaymentStateCheckout  SubscriptionResultCyclesPaymentState = "CHECKOUT"
	SubscriptionResultCyclesPaymentStatePending   SubscriptionResultCyclesPaymentState = "PENDING"
	SubscriptionResultCyclesPaymentStateCancelled SubscriptionResultCyclesPaymentState = "CANCELLED"
	SubscriptionResultCyclesPaymentStateDeclined  SubscriptionResultCyclesPaymentState = "DECLINED"
	SubscriptionResultCyclesPaymentStateCompleted SubscriptionResultCyclesPaymentState = "COMPLETED"
)

func (r SubscriptionResultCyclesPaymentState) IsKnown() bool {
	switch r {
	case SubscriptionResultCyclesPaymentStateCheckout, SubscriptionResultCyclesPaymentStatePending, SubscriptionResultCyclesPaymentStateCancelled, SubscriptionResultCyclesPaymentStateDeclined, SubscriptionResultCyclesPaymentStateCompleted:
		return true
	}
	return false
}

// Cycle type
type SubscriptionResultCyclesType string

const (
	SubscriptionResultCyclesTypeRegular SubscriptionResultCyclesType = "REGULAR"
	SubscriptionResultCyclesTypeRetry   SubscriptionResultCyclesType = "RETRY"
)

func (r SubscriptionResultCyclesType) IsKnown() bool {
	switch r {
	case SubscriptionResultCyclesTypeRegular, SubscriptionResultCyclesTypeRetry:
		return true
	}
	return false
}

// The interval at which the subscription is billed. Use 'MINUTE' for testing
// purposes only.
type SubscriptionResultFrequencyUnit string

const (
	SubscriptionResultFrequencyUnitMinute SubscriptionResultFrequencyUnit = "MINUTE"
	SubscriptionResultFrequencyUnitDay    SubscriptionResultFrequencyUnit = "DAY"
	SubscriptionResultFrequencyUnitWeek   SubscriptionResultFrequencyUnit = "WEEK"
	SubscriptionResultFrequencyUnitMonth  SubscriptionResultFrequencyUnit = "MONTH"
)

func (r SubscriptionResultFrequencyUnit) IsKnown() bool {
	switch r {
	case SubscriptionResultFrequencyUnitMinute, SubscriptionResultFrequencyUnitDay, SubscriptionResultFrequencyUnitWeek, SubscriptionResultFrequencyUnitMonth:
		return true
	}
	return false
}

// Retry strategy for subscription. If not specified, the subscription is canceled
// after the first failed payment attempt.
type SubscriptionResultRetryStrategy struct {
	// The number of intervals after which the system will retry the payment after an
	// unsuccessful attempt
	Frequency int64 `json:"frequency,required"`
	// Required number of retries
	NumberOfCycles int64 `json:"numberOfCycles,required"`
	// If specified, the nth element contains the percentage of the initial amount that
	// will be charged for the nth retry
	AmountAdjustments []int64 `json:"amountAdjustments"`
	// The interval at which the subscription is retried. Use 'MINUTE' for testing
	// purposes only.
	FrequencyUnit SubscriptionResultRetryStrategyFrequencyUnit `json:"frequencyUnit"`
	JSON          subscriptionResultRetryStrategyJSON          `json:"-"`
}

// subscriptionResultRetryStrategyJSON contains the JSON metadata for the struct
// [SubscriptionResultRetryStrategy]
type subscriptionResultRetryStrategyJSON struct {
	Frequency         apijson.Field
	NumberOfCycles    apijson.Field
	AmountAdjustments apijson.Field
	FrequencyUnit     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionResultRetryStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionResultRetryStrategyJSON) RawJSON() string {
	return r.raw
}

// The interval at which the subscription is retried. Use 'MINUTE' for testing
// purposes only.
type SubscriptionResultRetryStrategyFrequencyUnit string

const (
	SubscriptionResultRetryStrategyFrequencyUnitMinute SubscriptionResultRetryStrategyFrequencyUnit = "MINUTE"
	SubscriptionResultRetryStrategyFrequencyUnitDay    SubscriptionResultRetryStrategyFrequencyUnit = "DAY"
	SubscriptionResultRetryStrategyFrequencyUnitWeek   SubscriptionResultRetryStrategyFrequencyUnit = "WEEK"
	SubscriptionResultRetryStrategyFrequencyUnitMonth  SubscriptionResultRetryStrategyFrequencyUnit = "MONTH"
)

func (r SubscriptionResultRetryStrategyFrequencyUnit) IsKnown() bool {
	switch r {
	case SubscriptionResultRetryStrategyFrequencyUnitMinute, SubscriptionResultRetryStrategyFrequencyUnitDay, SubscriptionResultRetryStrategyFrequencyUnitWeek, SubscriptionResultRetryStrategyFrequencyUnitMonth:
		return true
	}
	return false
}

// Subscription state
type SubscriptionResultState string

const (
	SubscriptionResultStateActive    SubscriptionResultState = "ACTIVE"
	SubscriptionResultStateCancelled SubscriptionResultState = "CANCELLED"
	SubscriptionResultStateCompleted SubscriptionResultState = "COMPLETED"
)

func (r SubscriptionResultState) IsKnown() bool {
	switch r {
	case SubscriptionResultStateActive, SubscriptionResultStateCancelled, SubscriptionResultStateCompleted:
		return true
	}
	return false
}

type SubscriptionUpdateParams struct {
	// New subscription state
	State param.Field[SubscriptionUpdateParamsState] `json:"state"`
}

func (r SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// New subscription state
type SubscriptionUpdateParamsState string

const (
	SubscriptionUpdateParamsStateCancelled SubscriptionUpdateParamsState = "CANCELLED"
)

func (r SubscriptionUpdateParamsState) IsKnown() bool {
	switch r {
	case SubscriptionUpdateParamsStateCancelled:
		return true
	}
	return false
}
