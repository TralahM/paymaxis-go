// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paymaxis_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/TralahM/paymaxis-go"
	"github.com/TralahM/paymaxis-go/internal/testutil"
	"github.com/TralahM/paymaxis-go/option"
)

func TestPaymentNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paymaxis.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Payments.New(context.TODO(), paymaxis.PaymentNewParams{
		Currency:    paymaxis.F("EUR"),
		PaymentType: paymaxis.F(paymaxis.PaymentNewParamsPaymentTypeDeposit),
		AdditionalParameters: paymaxis.F(map[string]string{
			"bankCode":       "ABHY0065032",
			"countryOfBirth": "CY",
		}),
		Amount: paymaxis.F(0.000000),
		BillingAddress: paymaxis.F(paymaxis.PaymentNewParamsBillingAddress{
			AddressLine1: paymaxis.F("7, Sunny street"),
			AddressLine2: paymaxis.F("Office 3"),
			City:         paymaxis.F("Limassol"),
			CountryCode:  paymaxis.F("CY"),
			PostalCode:   paymaxis.F("4141"),
			State:        paymaxis.F("CA"),
		}),
		Card: paymaxis.F(paymaxis.PaymentNewParamsCard{
			CardholderName:   paymaxis.F("John Smith"),
			CardNumber:       paymaxis.F("4000 0000 0000 0002"),
			CardSecurityCode: paymaxis.F("010"),
			ExpiryMonth:      paymaxis.F("01"),
			ExpiryYear:       paymaxis.F("2030"),
		}),
		Customer: paymaxis.F(paymaxis.PaymentNewParamsCustomer{
			AccountName:                paymaxis.F("accountName"),
			AccountNumber:              paymaxis.F("accountNumber"),
			Bank:                       paymaxis.F("bank"),
			BankBranch:                 paymaxis.F("bankBranch"),
			CitizenshipCountryCode:     paymaxis.F("AU"),
			DateOfBirth:                paymaxis.F("2001-12-03"),
			DateOfFirstDeposit:         paymaxis.F("2021-01-01"),
			DepositsAmount:             paymaxis.F(int64(5000)),
			DepositsCnt:                paymaxis.F(int64(5000)),
			DocumentNumber:             paymaxis.F("documentNumber"),
			DocumentType:               paymaxis.F(paymaxis.PaymentNewParamsCustomerDocumentTypeArCdi),
			Email:                      paymaxis.F("my@email.com"),
			FirstName:                  paymaxis.F("John"),
			KYCStatus:                  paymaxis.F(true),
			LastName:                   paymaxis.F("Smith"),
			Locale:                     paymaxis.F("ru"),
			PaymentInstrumentKYCStatus: paymaxis.F(true),
			Phone:                      paymaxis.F("357 123123123"),
			ReferenceID:                paymaxis.F("customer_123"),
			RoutingGroup:               paymaxis.F("VIP"),
			WithdrawalsAmount:          paymaxis.F(int64(1000)),
			WithdrawalsCnt:             paymaxis.F(int64(1000)),
		}),
		Description:     paymaxis.F("Deposit 123 via TEST shop"),
		ParentPaymentID: paymaxis.F("91d27876e87f4b22b3ecd53924bf973d"),
		PaymentMethod:   paymaxis.F(paymaxis.PaymentNewParamsPaymentMethodBasicCard),
		RecurringToken:  paymaxis.F("recurringToken"),
		ReferenceID:     paymaxis.F("payment_id=123;custom_ref=456"),
		ReturnURL:       paymaxis.F("https://mywebsite.com/{id}/{referenceId}/{state}/{type}"),
		StartRecurring:  paymaxis.F(true),
		Subscription: paymaxis.F(paymaxis.PaymentNewParamsSubscription{
			Frequency:      paymaxis.F(int64(2)),
			Amount:         paymaxis.F(99.990000),
			Description:    paymaxis.F("Subscription to service"),
			FrequencyUnit:  paymaxis.F(paymaxis.PaymentNewParamsSubscriptionFrequencyUnitMinute),
			NumberOfCycles: paymaxis.F(int64(12)),
			RetryStrategy: paymaxis.F(paymaxis.PaymentNewParamsSubscriptionRetryStrategy{
				Frequency:         paymaxis.F(int64(2)),
				NumberOfCycles:    paymaxis.F(int64(12)),
				AmountAdjustments: paymaxis.F([]int64{int64(1), int64(1), int64(1)}),
				FrequencyUnit:     paymaxis.F(paymaxis.PaymentNewParamsSubscriptionRetryStrategyFrequencyUnitMinute),
			}),
			StartTime: paymaxis.F("2030-12-25T10:11:12"),
		}),
		WebhookURL: paymaxis.F("https://mywebsite.com/webhooks"),
	})
	if err != nil {
		var apierr *paymaxis.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paymaxis.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Payments.Get(context.TODO(), "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	if err != nil {
		var apierr *paymaxis.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paymaxis.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Payments.List(context.TODO(), paymaxis.PaymentListParams{
		Created: paymaxis.F(paymaxis.PaymentListParamsCreated{
			Gte: paymaxis.F("2021-10-13T10:26:18"),
			Lt:  paymaxis.F("2021-10-13T10:39:34"),
		}),
		Limit:  paymaxis.F(int64(1)),
		Offset: paymaxis.F(int64(0)),
		Updated: paymaxis.F(paymaxis.PaymentListParamsUpdated{
			Gte: paymaxis.F("2021-10-13T10:26:18"),
			Lt:  paymaxis.F("2021-10-13T10:39:34"),
		}),
	})
	if err != nil {
		var apierr *paymaxis.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
