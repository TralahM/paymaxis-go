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

func TestSubscriptionGet(t *testing.T) {
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
	_, err := client.Subscriptions.Get(context.TODO(), "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	if err != nil {
		var apierr *paymaxis.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.Update(
		context.TODO(),
		"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		paymaxis.SubscriptionUpdateParams{
			State: paymaxis.F(paymaxis.SubscriptionUpdateParamsStateCancelled),
		},
	)
	if err != nil {
		var apierr *paymaxis.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
