// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paymaxis_test

import (
	"context"
	"os"
	"testing"

	"github.com/TralahM/paymaxis-go"
	"github.com/TralahM/paymaxis-go/internal/testutil"
	"github.com/TralahM/paymaxis-go/option"
)

func TestUsage(t *testing.T) {
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
	payment, err := client.Payments.New(context.TODO(), paymaxis.PaymentNewParams{
		Currency:       paymaxis.F("EUR"),
		PaymentType:    paymaxis.F(paymaxis.PaymentNewParamsPaymentTypeDeposit),
		Amount:         paymaxis.F(4000.000000),
		PaymentMethod:  paymaxis.F(paymaxis.PaymentNewParamsPaymentMethodBasicCard),
		StartRecurring: paymaxis.F(true),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", payment.Result)
}
