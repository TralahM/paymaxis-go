// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paymaxis

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/TralahM/paymaxis-go/internal/apijson"
	"github.com/TralahM/paymaxis-go/internal/apiquery"
	"github.com/TralahM/paymaxis-go/internal/param"
	"github.com/TralahM/paymaxis-go/internal/requestconfig"
	"github.com/TralahM/paymaxis-go/option"
)

// PaymentService contains methods and other services that help with interacting
// with the paymaxis API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentService] method instead.
type PaymentService struct {
	Options    []option.RequestOption
	Operations *PaymentOperationService
}

// NewPaymentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaymentService(opts ...option.RequestOption) (r *PaymentService) {
	r = &PaymentService{}
	r.Options = opts
	r.Operations = NewPaymentOperationService(opts...)
	return
}

// Payment request, used for DEPOSITS, WITHDRAWALS and REFUNDS
func (r *PaymentService) New(ctx context.Context, body PaymentNewParams, opts ...option.RequestOption) (res *Payment, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Find Payment by Id
func (r *PaymentService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Payment, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/payments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of payments sorted by creation date (most recent first)
func (r *PaymentService) List(ctx context.Context, query PaymentListParams, opts ...option.RequestOption) (res *PaymentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v1/payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Payment struct {
	Result PaymentResult `json:"result"`
	// HTTP status code
	Status    int64       `json:"status"`
	Timestamp string      `json:"timestamp" format:"ISO 8601"`
	JSON      paymentJSON `json:"-"`
}

// paymentJSON contains the JSON metadata for the struct [Payment]
type paymentJSON struct {
	Result      apijson.Field
	Status      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Payment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentJSON) RawJSON() string {
	return r.raw
}

type PaymentResult struct {
	// Payment Id
	ID string `json:"id"`
	// Amount sent to the payment provider
	Amount float64 `json:"amount"`
	// Customer's billing address
	BillingAddress PaymentResultBillingAddress `json:"billingAddress"`
	// Currency sent to the payment provider
	Currency string                `json:"currency" format:"ISO 4217 code for FIAT currencies or cryptocurrency symbol"`
	Customer PaymentResultCustomer `json:"customer"`
	// Amount from payment request. Filled only if the request currency differs from
	// the currency sent to the payment provider.
	CustomerAmount float64 `json:"customerAmount"`
	// Currency from payment request. Filled only if it differs from the currency sent
	// to the payment provider.
	CustomerCurrency string `json:"customerCurrency" format:"ISO 4217 code for FIAT currencies or cryptocurrency symbol"`
	// Description of the transaction
	Description string `json:"description"`
	// Check 'Error Codes' section for details
	ErrorCode string `json:"errorCode"`
	// Provider fee. Filled only if supported by the provider.
	ExternalFeeAmount float64 `json:"externalFeeAmount"`
	// Provider fee currency. Filled only if supported by the provider.
	ExternalFeeCurrency string `json:"externalFeeCurrency" format:"ISO 4217 code for FIAT currencies or cryptocurrency symbol"`
	// Result code from external provider
	ExternalResultCode string `json:"externalResultCode"`
	// Initial transaction Id from payment request
	ParentPaymentID string `json:"parentPaymentId"`
	// Payment Method
	PaymentMethod        PaymentResultPaymentMethod        `json:"paymentMethod"`
	PaymentMethodDetails PaymentResultPaymentMethodDetails `json:"paymentMethodDetails"`
	// Payment Type
	PaymentType PaymentResultPaymentType `json:"paymentType"`
	// Token that can be used to continue the recurring chain
	RecurringToken string `json:"recurringToken"`
	// URL to redirect the customer
	RedirectURL string `json:"redirectUrl"`
	// referenceId from payment request
	ReferenceID string `json:"referenceId"`
	// Indicates whether this payment has started a recurring chain
	StartRecurring bool `json:"startRecurring"`
	// Payment State
	State PaymentResultState `json:"state"`
	// The name of the provider that was used to process this payment
	TerminalName string            `json:"terminalName"`
	JSON         paymentResultJSON `json:"-"`
}

// paymentResultJSON contains the JSON metadata for the struct [PaymentResult]
type paymentResultJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	BillingAddress       apijson.Field
	Currency             apijson.Field
	Customer             apijson.Field
	CustomerAmount       apijson.Field
	CustomerCurrency     apijson.Field
	Description          apijson.Field
	ErrorCode            apijson.Field
	ExternalFeeAmount    apijson.Field
	ExternalFeeCurrency  apijson.Field
	ExternalResultCode   apijson.Field
	ParentPaymentID      apijson.Field
	PaymentMethod        apijson.Field
	PaymentMethodDetails apijson.Field
	PaymentType          apijson.Field
	RecurringToken       apijson.Field
	RedirectURL          apijson.Field
	ReferenceID          apijson.Field
	StartRecurring       apijson.Field
	State                apijson.Field
	TerminalName         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PaymentResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentResultJSON) RawJSON() string {
	return r.raw
}

// Customer's billing address
type PaymentResultBillingAddress struct {
	// Line 1 of the address (e.g., Number, street, etc)
	AddressLine1 string `json:"addressLine1"`
	// Line 2 of the address (e.g., Suite, apt)
	AddressLine2 string `json:"addressLine2"`
	// City name
	City string `json:"city"`
	// 2-character IS0-3166-1 country code
	CountryCode string `json:"countryCode"`
	// Postal code
	PostalCode string `json:"postalCode"`
	// State code
	State string                          `json:"state"`
	JSON  paymentResultBillingAddressJSON `json:"-"`
}

// paymentResultBillingAddressJSON contains the JSON metadata for the struct
// [PaymentResultBillingAddress]
type paymentResultBillingAddressJSON struct {
	AddressLine1 apijson.Field
	AddressLine2 apijson.Field
	City         apijson.Field
	CountryCode  apijson.Field
	PostalCode   apijson.Field
	State        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PaymentResultBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentResultBillingAddressJSON) RawJSON() string {
	return r.raw
}

type PaymentResultCustomer struct {
	// Customer account name in the provider's system. Used for some types of
	// withdrawals.
	AccountName string `json:"accountName"`
	// Customer account number in the provider's system. Used for some types of
	// withdrawals.
	AccountNumber string `json:"accountNumber"`
	// Customer bank. Used for some types of withdrawals.
	Bank string `json:"bank"`
	// Customer bank branch. Used for some types of withdrawals.
	BankBranch string `json:"bankBranch"`
	// Customer country of citizenship
	CitizenshipCountryCode string `json:"citizenshipCountryCode"`
	DateOfBirth            string `json:"dateOfBirth" format:"ISO 8601 (YYYY-MM-DD)"`
	// Date of the first deposit from the customer
	DateOfFirstDeposit string `json:"dateOfFirstDeposit" format:"ISO 8601 (YYYY-MM-DD)"`
	// How much the customer has deposited, in base currency
	DepositsAmount int64 `json:"depositsAmount"`
	// How many times the customer made a deposit
	DepositsCnt int64 `json:"depositsCnt"`
	// An identifier for the customer assigned by a government authority
	DocumentNumber string `json:"documentNumber"`
	// Document Type
	DocumentType PaymentResultCustomerDocumentType `json:"documentType"`
	// Email address of the customer
	Email     string `json:"email" format:"email"`
	FirstName string `json:"firstName"`
	// Indicates whether the customer has passed KYC verification
	KYCStatus bool   `json:"kycStatus"`
	LastName  string `json:"lastName"`
	// Customer preferred display language
	Locale string `json:"locale"`
	// Indicates whether the payment instrument (usually the card number) has passed
	// KYC verification
	PaymentInstrumentKYCStatus bool `json:"paymentInstrumentKycStatus"`
	// International phone number of the customer, without the '+'. Use a space as a
	// separator between the dialing country code and local phone number.
	Phone string `json:"phone"`
	// Id of the customer assigned by Merchant
	ReferenceID string `json:"referenceId"`
	// Identify the customer as belonging to a specific group that is used for routing
	RoutingGroup string `json:"routingGroup"`
	// How much the customer has withdrawn, in base currency
	WithdrawalsAmount int64 `json:"withdrawalsAmount"`
	// How many times the customer made a withdrawal
	WithdrawalsCnt int64                     `json:"withdrawalsCnt"`
	JSON           paymentResultCustomerJSON `json:"-"`
}

// paymentResultCustomerJSON contains the JSON metadata for the struct
// [PaymentResultCustomer]
type paymentResultCustomerJSON struct {
	AccountName                apijson.Field
	AccountNumber              apijson.Field
	Bank                       apijson.Field
	BankBranch                 apijson.Field
	CitizenshipCountryCode     apijson.Field
	DateOfBirth                apijson.Field
	DateOfFirstDeposit         apijson.Field
	DepositsAmount             apijson.Field
	DepositsCnt                apijson.Field
	DocumentNumber             apijson.Field
	DocumentType               apijson.Field
	Email                      apijson.Field
	FirstName                  apijson.Field
	KYCStatus                  apijson.Field
	LastName                   apijson.Field
	Locale                     apijson.Field
	PaymentInstrumentKYCStatus apijson.Field
	Phone                      apijson.Field
	ReferenceID                apijson.Field
	RoutingGroup               apijson.Field
	WithdrawalsAmount          apijson.Field
	WithdrawalsCnt             apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *PaymentResultCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentResultCustomerJSON) RawJSON() string {
	return r.raw
}

// Document Type
type PaymentResultCustomerDocumentType string

const (
	PaymentResultCustomerDocumentTypeArCdi  PaymentResultCustomerDocumentType = "AR_CDI"
	PaymentResultCustomerDocumentTypeArCuil PaymentResultCustomerDocumentType = "AR_CUIL"
	PaymentResultCustomerDocumentTypeArCuit PaymentResultCustomerDocumentType = "AR_CUIT"
	PaymentResultCustomerDocumentTypeArDni  PaymentResultCustomerDocumentType = "AR_DNI"
	PaymentResultCustomerDocumentTypeArOtro PaymentResultCustomerDocumentType = "AR_OTRO"
	PaymentResultCustomerDocumentTypeBrCnpj PaymentResultCustomerDocumentType = "BR_CNPJ"
	PaymentResultCustomerDocumentTypeBrCpf  PaymentResultCustomerDocumentType = "BR_CPF"
	PaymentResultCustomerDocumentTypeClOtro PaymentResultCustomerDocumentType = "CL_OTRO"
	PaymentResultCustomerDocumentTypeClRun  PaymentResultCustomerDocumentType = "CL_RUN"
	PaymentResultCustomerDocumentTypeClRut  PaymentResultCustomerDocumentType = "CL_RUT"
	PaymentResultCustomerDocumentTypeCoCc   PaymentResultCustomerDocumentType = "CO_CC"
	PaymentResultCustomerDocumentTypeCoCe   PaymentResultCustomerDocumentType = "CO_CE"
	PaymentResultCustomerDocumentTypeCoDl   PaymentResultCustomerDocumentType = "CO_DL"
	PaymentResultCustomerDocumentTypeCoDni  PaymentResultCustomerDocumentType = "CO_DNI"
	PaymentResultCustomerDocumentTypeCoNe   PaymentResultCustomerDocumentType = "CO_NE"
	PaymentResultCustomerDocumentTypeCoNit  PaymentResultCustomerDocumentType = "CO_NIT"
	PaymentResultCustomerDocumentTypeCoPp   PaymentResultCustomerDocumentType = "CO_PP"
	PaymentResultCustomerDocumentTypeCoSS   PaymentResultCustomerDocumentType = "CO_SS"
	PaymentResultCustomerDocumentTypeCoTi   PaymentResultCustomerDocumentType = "CO_TI"
	PaymentResultCustomerDocumentTypeCrCdi  PaymentResultCustomerDocumentType = "CR_CDI"
	PaymentResultCustomerDocumentTypeEcDni  PaymentResultCustomerDocumentType = "EC_DNI"
	PaymentResultCustomerDocumentTypeEcPp   PaymentResultCustomerDocumentType = "EC_PP"
	PaymentResultCustomerDocumentTypeEcRuc  PaymentResultCustomerDocumentType = "EC_RUC"
	PaymentResultCustomerDocumentTypeGtCui  PaymentResultCustomerDocumentType = "GT_CUI"
	PaymentResultCustomerDocumentTypeGtDpi  PaymentResultCustomerDocumentType = "GT_DPI"
	PaymentResultCustomerDocumentTypeGtNit  PaymentResultCustomerDocumentType = "GT_NIT"
	PaymentResultCustomerDocumentTypeMxCurp PaymentResultCustomerDocumentType = "MX_CURP"
	PaymentResultCustomerDocumentTypeMxIfe  PaymentResultCustomerDocumentType = "MX_IFE"
	PaymentResultCustomerDocumentTypeMxPp   PaymentResultCustomerDocumentType = "MX_PP"
	PaymentResultCustomerDocumentTypeMxRfc  PaymentResultCustomerDocumentType = "MX_RFC"
	PaymentResultCustomerDocumentTypePaCip  PaymentResultCustomerDocumentType = "PA_CIP"
	PaymentResultCustomerDocumentTypePeCe   PaymentResultCustomerDocumentType = "PE_CE"
	PaymentResultCustomerDocumentTypePeDni  PaymentResultCustomerDocumentType = "PE_DNI"
	PaymentResultCustomerDocumentTypePeOtro PaymentResultCustomerDocumentType = "PE_OTRO"
	PaymentResultCustomerDocumentTypePePp   PaymentResultCustomerDocumentType = "PE_PP"
	PaymentResultCustomerDocumentTypePeRuc  PaymentResultCustomerDocumentType = "PE_RUC"
)

func (r PaymentResultCustomerDocumentType) IsKnown() bool {
	switch r {
	case PaymentResultCustomerDocumentTypeArCdi, PaymentResultCustomerDocumentTypeArCuil, PaymentResultCustomerDocumentTypeArCuit, PaymentResultCustomerDocumentTypeArDni, PaymentResultCustomerDocumentTypeArOtro, PaymentResultCustomerDocumentTypeBrCnpj, PaymentResultCustomerDocumentTypeBrCpf, PaymentResultCustomerDocumentTypeClOtro, PaymentResultCustomerDocumentTypeClRun, PaymentResultCustomerDocumentTypeClRut, PaymentResultCustomerDocumentTypeCoCc, PaymentResultCustomerDocumentTypeCoCe, PaymentResultCustomerDocumentTypeCoDl, PaymentResultCustomerDocumentTypeCoDni, PaymentResultCustomerDocumentTypeCoNe, PaymentResultCustomerDocumentTypeCoNit, PaymentResultCustomerDocumentTypeCoPp, PaymentResultCustomerDocumentTypeCoSS, PaymentResultCustomerDocumentTypeCoTi, PaymentResultCustomerDocumentTypeCrCdi, PaymentResultCustomerDocumentTypeEcDni, PaymentResultCustomerDocumentTypeEcPp, PaymentResultCustomerDocumentTypeEcRuc, PaymentResultCustomerDocumentTypeGtCui, PaymentResultCustomerDocumentTypeGtDpi, PaymentResultCustomerDocumentTypeGtNit, PaymentResultCustomerDocumentTypeMxCurp, PaymentResultCustomerDocumentTypeMxIfe, PaymentResultCustomerDocumentTypeMxPp, PaymentResultCustomerDocumentTypeMxRfc, PaymentResultCustomerDocumentTypePaCip, PaymentResultCustomerDocumentTypePeCe, PaymentResultCustomerDocumentTypePeDni, PaymentResultCustomerDocumentTypePeOtro, PaymentResultCustomerDocumentTypePePp, PaymentResultCustomerDocumentTypePeRuc:
		return true
	}
	return false
}

// Payment Method
type PaymentResultPaymentMethod string

const (
	PaymentResultPaymentMethodBasicCard            PaymentResultPaymentMethod = "BASIC_CARD"
	PaymentResultPaymentMethodCrypto               PaymentResultPaymentMethod = "CRYPTO"
	PaymentResultPaymentMethodFlexepin             PaymentResultPaymentMethod = "FLEXEPIN"
	PaymentResultPaymentMethodMacropay             PaymentResultPaymentMethod = "MACROPAY"
	PaymentResultPaymentMethodSkrill               PaymentResultPaymentMethod = "SKRILL"
	PaymentResultPaymentMethodPayretailers         PaymentResultPaymentMethod = "PAYRETAILERS"
	PaymentResultPaymentMethodLocalpayment         PaymentResultPaymentMethod = "LOCALPAYMENT"
	PaymentResultPaymentMethodMonnet               PaymentResultPaymentMethod = "MONNET"
	PaymentResultPaymentMethodPaypal               PaymentResultPaymentMethod = "PAYPAL"
	PaymentResultPaymentMethodNeteller             PaymentResultPaymentMethod = "NETELLER"
	PaymentResultPaymentMethodTrustpayments        PaymentResultPaymentMethod = "TRUSTPAYMENTS"
	PaymentResultPaymentMethodPaymaxis             PaymentResultPaymentMethod = "PAYMAXIS"
	PaymentResultPaymentMethodGate8Transact        PaymentResultPaymentMethod = "GATE8TRANSACT"
	PaymentResultPaymentMethodTink                 PaymentResultPaymentMethod = "TINK"
	PaymentResultPaymentMethodB2Binpay             PaymentResultPaymentMethod = "B2BINPAY"
	PaymentResultPaymentMethodClick                PaymentResultPaymentMethod = "CLICK"
	PaymentResultPaymentMethodMonetix              PaymentResultPaymentMethod = "MONETIX"
	PaymentResultPaymentMethodPerfectmoney         PaymentResultPaymentMethod = "PERFECTMONEY"
	PaymentResultPaymentMethodVolt                 PaymentResultPaymentMethod = "VOLT"
	PaymentResultPaymentMethodKesspay              PaymentResultPaymentMethod = "KESSPAY"
	PaymentResultPaymentMethodBillline             PaymentResultPaymentMethod = "BILLLINE"
	PaymentResultPaymentMethodNgenius              PaymentResultPaymentMethod = "NGENIUS"
	PaymentResultPaymentMethodAstropay             PaymentResultPaymentMethod = "ASTROPAY"
	PaymentResultPaymentMethodAlycepay             PaymentResultPaymentMethod = "ALYCEPAY"
	PaymentResultPaymentMethodPix                  PaymentResultPaymentMethod = "PIX"
	PaymentResultPaymentMethodBoleto               PaymentResultPaymentMethod = "BOLETO"
	PaymentResultPaymentMethodUpi                  PaymentResultPaymentMethod = "UPI"
	PaymentResultPaymentMethodPaytm                PaymentResultPaymentMethod = "PAYTM"
	PaymentResultPaymentMethodNetbanking           PaymentResultPaymentMethod = "NETBANKING"
	PaymentResultPaymentMethodFinrax               PaymentResultPaymentMethod = "FINRAX"
	PaymentResultPaymentMethodSpoynt               PaymentResultPaymentMethod = "SPOYNT"
	PaymentResultPaymentMethodXinpay               PaymentResultPaymentMethod = "XINPAY"
	PaymentResultPaymentMethodOmnimatrix           PaymentResultPaymentMethod = "OMNIMATRIX"
	PaymentResultPaymentMethodDpopay               PaymentResultPaymentMethod = "DPOPAY"
	PaymentResultPaymentMethodExternalHpp          PaymentResultPaymentMethod = "EXTERNAL_HPP"
	PaymentResultPaymentMethodXanpay               PaymentResultPaymentMethod = "XANPAY"
	PaymentResultPaymentMethodInrpay               PaymentResultPaymentMethod = "INRPAY"
	PaymentResultPaymentMethodAri10                PaymentResultPaymentMethod = "ARI10"
	PaymentResultPaymentMethodSofort               PaymentResultPaymentMethod = "SOFORT"
	PaymentResultPaymentMethodGiropay              PaymentResultPaymentMethod = "GIROPAY"
	PaymentResultPaymentMethodPaysafecard          PaymentResultPaymentMethod = "PAYSAFECARD"
	PaymentResultPaymentMethodPaysafecash          PaymentResultPaymentMethod = "PAYSAFECASH"
	PaymentResultPaymentMethodOpenBanking          PaymentResultPaymentMethod = "OPEN_BANKING"
	PaymentResultPaymentMethodKlarna               PaymentResultPaymentMethod = "KLARNA"
	PaymentResultPaymentMethodSpei                 PaymentResultPaymentMethod = "SPEI"
	PaymentResultPaymentMethodPaycash              PaymentResultPaymentMethod = "PAYCASH"
	PaymentResultPaymentMethodRapipago             PaymentResultPaymentMethod = "RAPIPAGO"
	PaymentResultPaymentMethodPagofacil            PaymentResultPaymentMethod = "PAGOFACIL"
	PaymentResultPaymentMethodRapidtransfer        PaymentResultPaymentMethod = "RAPIDTRANSFER"
	PaymentResultPaymentMethodMobileMoney          PaymentResultPaymentMethod = "MOBILE_MONEY"
	PaymentResultPaymentMethodInterac              PaymentResultPaymentMethod = "INTERAC"
	PaymentResultPaymentMethodInteracEto           PaymentResultPaymentMethod = "INTERAC_ETO"
	PaymentResultPaymentMethodInteracRto           PaymentResultPaymentMethod = "INTERAC_RTO"
	PaymentResultPaymentMethodInteracACH           PaymentResultPaymentMethod = "INTERAC_ACH"
	PaymentResultPaymentMethodPicpay               PaymentResultPaymentMethod = "PICPAY"
	PaymentResultPaymentMethodMollie               PaymentResultPaymentMethod = "MOLLIE"
	PaymentResultPaymentMethodTed                  PaymentResultPaymentMethod = "TED"
	PaymentResultPaymentMethodZipay                PaymentResultPaymentMethod = "ZIPAY"
	PaymentResultPaymentMethodPse                  PaymentResultPaymentMethod = "PSE"
	PaymentResultPaymentMethodEfecty               PaymentResultPaymentMethod = "EFECTY"
	PaymentResultPaymentMethodBanktransfer         PaymentResultPaymentMethod = "BANKTRANSFER"
	PaymentResultPaymentMethodPec                  PaymentResultPaymentMethod = "PEC"
	PaymentResultPaymentMethodOxxo                 PaymentResultPaymentMethod = "OXXO"
	PaymentResultPaymentMethodWebpay               PaymentResultPaymentMethod = "WEBPAY"
	PaymentResultPaymentMethodPagoefectivo         PaymentResultPaymentMethod = "PAGOEFECTIVO"
	PaymentResultPaymentMethodMifinity             PaymentResultPaymentMethod = "MIFINITY"
	PaymentResultPaymentMethodPayport              PaymentResultPaymentMethod = "PAYPORT"
	PaymentResultPaymentMethodJetoncash            PaymentResultPaymentMethod = "JETONCASH"
	PaymentResultPaymentMethodJetonwallet          PaymentResultPaymentMethod = "JETONWALLET"
	PaymentResultPaymentMethodNoda                 PaymentResultPaymentMethod = "NODA"
	PaymentResultPaymentMethodNodaRevolut          PaymentResultPaymentMethod = "NODA_REVOLUT"
	PaymentResultPaymentMethodAlfakit              PaymentResultPaymentMethod = "ALFAKIT"
	PaymentResultPaymentMethodPayfun               PaymentResultPaymentMethod = "PAYFUN"
	PaymentResultPaymentMethodEmanat               PaymentResultPaymentMethod = "EMANAT"
	PaymentResultPaymentMethodM10                  PaymentResultPaymentMethod = "M10"
	PaymentResultPaymentMethodRubpay               PaymentResultPaymentMethod = "RUBPAY"
	PaymentResultPaymentMethodMonerchy             PaymentResultPaymentMethod = "MONERCHY"
	PaymentResultPaymentMethodMuchbetter           PaymentResultPaymentMethod = "MUCHBETTER"
	PaymentResultPaymentMethodYapily               PaymentResultPaymentMethod = "YAPILY"
	PaymentResultPaymentMethodInai                 PaymentResultPaymentMethod = "INAI"
	PaymentResultPaymentMethodImps                 PaymentResultPaymentMethod = "IMPS"
	PaymentResultPaymentMethodRtgs                 PaymentResultPaymentMethod = "RTGS"
	PaymentResultPaymentMethodPayid                PaymentResultPaymentMethod = "PAYID"
	PaymentResultPaymentMethodZotapay              PaymentResultPaymentMethod = "ZOTAPAY"
	PaymentResultPaymentMethodSbp                  PaymentResultPaymentMethod = "SBP"
	PaymentResultPaymentMethodP2PCard              PaymentResultPaymentMethod = "P2P_CARD"
	PaymentResultPaymentMethodP2PIban              PaymentResultPaymentMethod = "P2P_IBAN"
	PaymentResultPaymentMethodP2PSbp               PaymentResultPaymentMethod = "P2P_SBP"
	PaymentResultPaymentMethodP2PMobile            PaymentResultPaymentMethod = "P2P_MOBILE"
	PaymentResultPaymentMethodPush                 PaymentResultPaymentMethod = "PUSH"
	PaymentResultPaymentMethodGateiq               PaymentResultPaymentMethod = "GATEIQ"
	PaymentResultPaymentMethodViettel              PaymentResultPaymentMethod = "VIETTEL"
	PaymentResultPaymentMethodZalo                 PaymentResultPaymentMethod = "ZALO"
	PaymentResultPaymentMethodQr                   PaymentResultPaymentMethod = "QR"
	PaymentResultPaymentMethodCup                  PaymentResultPaymentMethod = "CUP"
	PaymentResultPaymentMethodCodi                 PaymentResultPaymentMethod = "CODI"
	PaymentResultPaymentMethodPay2Play             PaymentResultPaymentMethod = "PAY2PLAY"
	PaymentResultPaymentMethodBkash                PaymentResultPaymentMethod = "BKASH"
	PaymentResultPaymentMethodNagad                PaymentResultPaymentMethod = "NAGAD"
	PaymentResultPaymentMethodRocket               PaymentResultPaymentMethod = "ROCKET"
	PaymentResultPaymentMethodVirtualAccount       PaymentResultPaymentMethod = "VIRTUAL_ACCOUNT"
	PaymentResultPaymentMethodMultibanco           PaymentResultPaymentMethod = "MULTIBANCO"
	PaymentResultPaymentMethodBlik                 PaymentResultPaymentMethod = "BLIK"
	PaymentResultPaymentMethodMbway                PaymentResultPaymentMethod = "MBWAY"
	PaymentResultPaymentMethodP24                  PaymentResultPaymentMethod = "P24"
	PaymentResultPaymentMethodMistercash           PaymentResultPaymentMethod = "MISTERCASH"
	PaymentResultPaymentMethodMach                 PaymentResultPaymentMethod = "MACH"
	PaymentResultPaymentMethodKhipu                PaymentResultPaymentMethod = "KHIPU"
	PaymentResultPaymentMethodNeft                 PaymentResultPaymentMethod = "NEFT"
	PaymentResultPaymentMethodSticpay              PaymentResultPaymentMethod = "STICPAY"
	PaymentResultPaymentMethodSberpay              PaymentResultPaymentMethod = "SBERPAY"
	PaymentResultPaymentMethodMobileCommerce       PaymentResultPaymentMethod = "MOBILE_COMMERCE"
	PaymentResultPaymentMethodBinancePay           PaymentResultPaymentMethod = "BINANCE_PAY"
	PaymentResultPaymentMethodMpay                 PaymentResultPaymentMethod = "MPAY"
	PaymentResultPaymentMethodChek                 PaymentResultPaymentMethod = "CHEK"
	PaymentResultPaymentMethodKlapEfectivo         PaymentResultPaymentMethod = "KLAP_EFECTIVO"
	PaymentResultPaymentMethodKlapTransferencia    PaymentResultPaymentMethod = "KLAP_TRANSFERENCIA"
	PaymentResultPaymentMethodPago46               PaymentResultPaymentMethod = "PAGO46"
	PaymentResultPaymentMethodHites                PaymentResultPaymentMethod = "HITES"
	PaymentResultPaymentMethodServifacil           PaymentResultPaymentMethod = "SERVIFACIL"
	PaymentResultPaymentMethodOpenpayd             PaymentResultPaymentMethod = "OPENPAYD"
	PaymentResultPaymentMethodFawry                PaymentResultPaymentMethod = "FAWRY"
	PaymentResultPaymentMethodEps                  PaymentResultPaymentMethod = "EPS"
	PaymentResultPaymentMethodIdeal                PaymentResultPaymentMethod = "IDEAL"
	PaymentResultPaymentMethodTrustly              PaymentResultPaymentMethod = "TRUSTLY"
	PaymentResultPaymentMethodUssd                 PaymentResultPaymentMethod = "USSD"
	PaymentResultPaymentMethodMpesa                PaymentResultPaymentMethod = "MPESA"
	PaymentResultPaymentMethodEnaira               PaymentResultPaymentMethod = "ENAIRA"
	PaymentResultPaymentMethodOnevoucher           PaymentResultPaymentMethod = "ONEVOUCHER"
	PaymentResultPaymentMethodBancontact           PaymentResultPaymentMethod = "BANCONTACT"
	PaymentResultPaymentMethodSwish                PaymentResultPaymentMethod = "SWISH"
	PaymentResultPaymentMethodEft                  PaymentResultPaymentMethod = "EFT"
	PaymentResultPaymentMethodGcash                PaymentResultPaymentMethod = "GCASH"
	PaymentResultPaymentMethodPaymaya              PaymentResultPaymentMethod = "PAYMAYA"
	PaymentResultPaymentMethodPagoMovil            PaymentResultPaymentMethod = "PAGO_MOVIL"
	PaymentResultPaymentMethodPagoMovilInst        PaymentResultPaymentMethod = "PAGO_MOVIL_INST"
	PaymentResultPaymentMethodBiopago              PaymentResultPaymentMethod = "BIOPAGO"
	PaymentResultPaymentMethodCash                 PaymentResultPaymentMethod = "CASH"
	PaymentResultPaymentMethodVoucherry            PaymentResultPaymentMethod = "VOUCHERRY"
	PaymentResultPaymentMethodApplepay             PaymentResultPaymentMethod = "APPLEPAY"
	PaymentResultPaymentMethodGooglepay            PaymentResultPaymentMethod = "GOOGLEPAY"
	PaymentResultPaymentMethodBrite                PaymentResultPaymentMethod = "BRITE"
	PaymentResultPaymentMethodVouchstar            PaymentResultPaymentMethod = "VOUCHSTAR"
	PaymentResultPaymentMethodRevolut              PaymentResultPaymentMethod = "REVOLUT"
	PaymentResultPaymentMethodOnlineBanking        PaymentResultPaymentMethod = "ONLINE_BANKING"
	PaymentResultPaymentMethodPromptpay            PaymentResultPaymentMethod = "PROMPTPAY"
	PaymentResultPaymentMethodTruemoney            PaymentResultPaymentMethod = "TRUEMONEY"
	PaymentResultPaymentMethodMomopayVn            PaymentResultPaymentMethod = "MOMOPAY_VN"
	PaymentResultPaymentMethodMomopayRw            PaymentResultPaymentMethod = "MOMOPAY_RW"
	PaymentResultPaymentMethodVnpayQr              PaymentResultPaymentMethod = "VNPAY_QR"
	PaymentResultPaymentMethodN26                  PaymentResultPaymentMethod = "N26"
	PaymentResultPaymentMethodWise                 PaymentResultPaymentMethod = "WISE"
	PaymentResultPaymentMethodPaydoWallet          PaymentResultPaymentMethod = "PAYDO_WALLET"
	PaymentResultPaymentMethodPapara               PaymentResultPaymentMethod = "PAPARA"
	PaymentResultPaymentMethodPayoutSepaBatch      PaymentResultPaymentMethod = "PAYOUT_SEPA_BATCH"
	PaymentResultPaymentMethodPayoutNonsepaRequest PaymentResultPaymentMethod = "PAYOUT_NONSEPA_REQUEST"
)

func (r PaymentResultPaymentMethod) IsKnown() bool {
	switch r {
	case PaymentResultPaymentMethodBasicCard, PaymentResultPaymentMethodCrypto, PaymentResultPaymentMethodFlexepin, PaymentResultPaymentMethodMacropay, PaymentResultPaymentMethodSkrill, PaymentResultPaymentMethodPayretailers, PaymentResultPaymentMethodLocalpayment, PaymentResultPaymentMethodMonnet, PaymentResultPaymentMethodPaypal, PaymentResultPaymentMethodNeteller, PaymentResultPaymentMethodTrustpayments, PaymentResultPaymentMethodPaymaxis, PaymentResultPaymentMethodGate8Transact, PaymentResultPaymentMethodTink, PaymentResultPaymentMethodB2Binpay, PaymentResultPaymentMethodClick, PaymentResultPaymentMethodMonetix, PaymentResultPaymentMethodPerfectmoney, PaymentResultPaymentMethodVolt, PaymentResultPaymentMethodKesspay, PaymentResultPaymentMethodBillline, PaymentResultPaymentMethodNgenius, PaymentResultPaymentMethodAstropay, PaymentResultPaymentMethodAlycepay, PaymentResultPaymentMethodPix, PaymentResultPaymentMethodBoleto, PaymentResultPaymentMethodUpi, PaymentResultPaymentMethodPaytm, PaymentResultPaymentMethodNetbanking, PaymentResultPaymentMethodFinrax, PaymentResultPaymentMethodSpoynt, PaymentResultPaymentMethodXinpay, PaymentResultPaymentMethodOmnimatrix, PaymentResultPaymentMethodDpopay, PaymentResultPaymentMethodExternalHpp, PaymentResultPaymentMethodXanpay, PaymentResultPaymentMethodInrpay, PaymentResultPaymentMethodAri10, PaymentResultPaymentMethodSofort, PaymentResultPaymentMethodGiropay, PaymentResultPaymentMethodPaysafecard, PaymentResultPaymentMethodPaysafecash, PaymentResultPaymentMethodOpenBanking, PaymentResultPaymentMethodKlarna, PaymentResultPaymentMethodSpei, PaymentResultPaymentMethodPaycash, PaymentResultPaymentMethodRapipago, PaymentResultPaymentMethodPagofacil, PaymentResultPaymentMethodRapidtransfer, PaymentResultPaymentMethodMobileMoney, PaymentResultPaymentMethodInterac, PaymentResultPaymentMethodInteracEto, PaymentResultPaymentMethodInteracRto, PaymentResultPaymentMethodInteracACH, PaymentResultPaymentMethodPicpay, PaymentResultPaymentMethodMollie, PaymentResultPaymentMethodTed, PaymentResultPaymentMethodZipay, PaymentResultPaymentMethodPse, PaymentResultPaymentMethodEfecty, PaymentResultPaymentMethodBanktransfer, PaymentResultPaymentMethodPec, PaymentResultPaymentMethodOxxo, PaymentResultPaymentMethodWebpay, PaymentResultPaymentMethodPagoefectivo, PaymentResultPaymentMethodMifinity, PaymentResultPaymentMethodPayport, PaymentResultPaymentMethodJetoncash, PaymentResultPaymentMethodJetonwallet, PaymentResultPaymentMethodNoda, PaymentResultPaymentMethodNodaRevolut, PaymentResultPaymentMethodAlfakit, PaymentResultPaymentMethodPayfun, PaymentResultPaymentMethodEmanat, PaymentResultPaymentMethodM10, PaymentResultPaymentMethodRubpay, PaymentResultPaymentMethodMonerchy, PaymentResultPaymentMethodMuchbetter, PaymentResultPaymentMethodYapily, PaymentResultPaymentMethodInai, PaymentResultPaymentMethodImps, PaymentResultPaymentMethodRtgs, PaymentResultPaymentMethodPayid, PaymentResultPaymentMethodZotapay, PaymentResultPaymentMethodSbp, PaymentResultPaymentMethodP2PCard, PaymentResultPaymentMethodP2PIban, PaymentResultPaymentMethodP2PSbp, PaymentResultPaymentMethodP2PMobile, PaymentResultPaymentMethodPush, PaymentResultPaymentMethodGateiq, PaymentResultPaymentMethodViettel, PaymentResultPaymentMethodZalo, PaymentResultPaymentMethodQr, PaymentResultPaymentMethodCup, PaymentResultPaymentMethodCodi, PaymentResultPaymentMethodPay2Play, PaymentResultPaymentMethodBkash, PaymentResultPaymentMethodNagad, PaymentResultPaymentMethodRocket, PaymentResultPaymentMethodVirtualAccount, PaymentResultPaymentMethodMultibanco, PaymentResultPaymentMethodBlik, PaymentResultPaymentMethodMbway, PaymentResultPaymentMethodP24, PaymentResultPaymentMethodMistercash, PaymentResultPaymentMethodMach, PaymentResultPaymentMethodKhipu, PaymentResultPaymentMethodNeft, PaymentResultPaymentMethodSticpay, PaymentResultPaymentMethodSberpay, PaymentResultPaymentMethodMobileCommerce, PaymentResultPaymentMethodBinancePay, PaymentResultPaymentMethodMpay, PaymentResultPaymentMethodChek, PaymentResultPaymentMethodKlapEfectivo, PaymentResultPaymentMethodKlapTransferencia, PaymentResultPaymentMethodPago46, PaymentResultPaymentMethodHites, PaymentResultPaymentMethodServifacil, PaymentResultPaymentMethodOpenpayd, PaymentResultPaymentMethodFawry, PaymentResultPaymentMethodEps, PaymentResultPaymentMethodIdeal, PaymentResultPaymentMethodTrustly, PaymentResultPaymentMethodUssd, PaymentResultPaymentMethodMpesa, PaymentResultPaymentMethodEnaira, PaymentResultPaymentMethodOnevoucher, PaymentResultPaymentMethodBancontact, PaymentResultPaymentMethodSwish, PaymentResultPaymentMethodEft, PaymentResultPaymentMethodGcash, PaymentResultPaymentMethodPaymaya, PaymentResultPaymentMethodPagoMovil, PaymentResultPaymentMethodPagoMovilInst, PaymentResultPaymentMethodBiopago, PaymentResultPaymentMethodCash, PaymentResultPaymentMethodVoucherry, PaymentResultPaymentMethodApplepay, PaymentResultPaymentMethodGooglepay, PaymentResultPaymentMethodBrite, PaymentResultPaymentMethodVouchstar, PaymentResultPaymentMethodRevolut, PaymentResultPaymentMethodOnlineBanking, PaymentResultPaymentMethodPromptpay, PaymentResultPaymentMethodTruemoney, PaymentResultPaymentMethodMomopayVn, PaymentResultPaymentMethodMomopayRw, PaymentResultPaymentMethodVnpayQr, PaymentResultPaymentMethodN26, PaymentResultPaymentMethodWise, PaymentResultPaymentMethodPaydoWallet, PaymentResultPaymentMethodPapara, PaymentResultPaymentMethodPayoutSepaBatch, PaymentResultPaymentMethodPayoutNonsepaRequest:
		return true
	}
	return false
}

type PaymentResultPaymentMethodDetails struct {
	// Card expiration month (for BASIC_CARD payment method only)
	CardExpiryMonth string `json:"cardExpiryMonth"`
	// Card expiration year (for BASIC_CARD payment method only)
	CardExpiryYear string `json:"cardExpiryYear"`
	// Cardholder name (for BASIC_CARD payment method only)
	CardholderName string `json:"cardholderName"`
	// Card issuing country code (for BASIC_CARD payment method only)
	CardIssuingCountryCode string `json:"cardIssuingCountryCode"`
	// Customer account Id in external system or masked card PAN
	CustomerAccountNumber string                                `json:"customerAccountNumber"`
	JSON                  paymentResultPaymentMethodDetailsJSON `json:"-"`
}

// paymentResultPaymentMethodDetailsJSON contains the JSON metadata for the struct
// [PaymentResultPaymentMethodDetails]
type paymentResultPaymentMethodDetailsJSON struct {
	CardExpiryMonth        apijson.Field
	CardExpiryYear         apijson.Field
	CardholderName         apijson.Field
	CardIssuingCountryCode apijson.Field
	CustomerAccountNumber  apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PaymentResultPaymentMethodDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentResultPaymentMethodDetailsJSON) RawJSON() string {
	return r.raw
}

// Payment Type
type PaymentResultPaymentType string

const (
	PaymentResultPaymentTypeDeposit    PaymentResultPaymentType = "DEPOSIT"
	PaymentResultPaymentTypeWithdrawal PaymentResultPaymentType = "WITHDRAWAL"
	PaymentResultPaymentTypeRefund     PaymentResultPaymentType = "REFUND"
)

func (r PaymentResultPaymentType) IsKnown() bool {
	switch r {
	case PaymentResultPaymentTypeDeposit, PaymentResultPaymentTypeWithdrawal, PaymentResultPaymentTypeRefund:
		return true
	}
	return false
}

// Payment State
type PaymentResultState string

const (
	PaymentResultStateCheckout  PaymentResultState = "CHECKOUT"
	PaymentResultStatePending   PaymentResultState = "PENDING"
	PaymentResultStateCancelled PaymentResultState = "CANCELLED"
	PaymentResultStateDeclined  PaymentResultState = "DECLINED"
	PaymentResultStateCompleted PaymentResultState = "COMPLETED"
)

func (r PaymentResultState) IsKnown() bool {
	switch r {
	case PaymentResultStateCheckout, PaymentResultStatePending, PaymentResultStateCancelled, PaymentResultStateDeclined, PaymentResultStateCompleted:
		return true
	}
	return false
}

type PaymentListResponse struct {
	// Indicates if there are more pages to return
	HasMore bool                        `json:"hasMore"`
	Result  []PaymentListResponseResult `json:"result"`
	// HTTP status code
	Status    int64                   `json:"status"`
	Timestamp string                  `json:"timestamp" format:"ISO 8601"`
	JSON      paymentListResponseJSON `json:"-"`
}

// paymentListResponseJSON contains the JSON metadata for the struct
// [PaymentListResponse]
type paymentListResponseJSON struct {
	HasMore     apijson.Field
	Result      apijson.Field
	Status      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentListResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentListResponseResult struct {
	// Payment Id
	ID string `json:"id"`
	// Amount sent to the payment provider
	Amount float64 `json:"amount"`
	// Customer's billing address
	BillingAddress PaymentListResponseResultBillingAddress `json:"billingAddress"`
	// Currency sent to the payment provider
	Currency string                            `json:"currency" format:"ISO 4217 code for FIAT currencies or cryptocurrency symbol"`
	Customer PaymentListResponseResultCustomer `json:"customer"`
	// Amount from payment request. Filled only if the request currency differs from
	// the currency sent to the payment provider.
	CustomerAmount float64 `json:"customerAmount"`
	// Currency from payment request. Filled only if it differs from the currency sent
	// to the payment provider.
	CustomerCurrency string `json:"customerCurrency" format:"ISO 4217 code for FIAT currencies or cryptocurrency symbol"`
	// Description of the transaction
	Description string `json:"description"`
	// Check 'Error Codes' section for details
	ErrorCode string `json:"errorCode"`
	// Provider fee. Filled only if supported by the provider.
	ExternalFeeAmount float64 `json:"externalFeeAmount"`
	// Provider fee currency. Filled only if supported by the provider.
	ExternalFeeCurrency string `json:"externalFeeCurrency" format:"ISO 4217 code for FIAT currencies or cryptocurrency symbol"`
	// Result code from external provider
	ExternalResultCode string `json:"externalResultCode"`
	// Initial transaction Id from payment request
	ParentPaymentID string `json:"parentPaymentId"`
	// Payment Method
	PaymentMethod        PaymentListResponseResultPaymentMethod        `json:"paymentMethod"`
	PaymentMethodDetails PaymentListResponseResultPaymentMethodDetails `json:"paymentMethodDetails"`
	// Payment Type
	PaymentType PaymentListResponseResultPaymentType `json:"paymentType"`
	// Token that can be used to continue the recurring chain
	RecurringToken string `json:"recurringToken"`
	// URL to redirect the customer
	RedirectURL string `json:"redirectUrl"`
	// referenceId from payment request
	ReferenceID string `json:"referenceId"`
	// Indicates whether this payment has started a recurring chain
	StartRecurring bool `json:"startRecurring"`
	// Payment State
	State PaymentListResponseResultState `json:"state"`
	// The name of the provider that was used to process this payment
	TerminalName string                        `json:"terminalName"`
	JSON         paymentListResponseResultJSON `json:"-"`
}

// paymentListResponseResultJSON contains the JSON metadata for the struct
// [PaymentListResponseResult]
type paymentListResponseResultJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	BillingAddress       apijson.Field
	Currency             apijson.Field
	Customer             apijson.Field
	CustomerAmount       apijson.Field
	CustomerCurrency     apijson.Field
	Description          apijson.Field
	ErrorCode            apijson.Field
	ExternalFeeAmount    apijson.Field
	ExternalFeeCurrency  apijson.Field
	ExternalResultCode   apijson.Field
	ParentPaymentID      apijson.Field
	PaymentMethod        apijson.Field
	PaymentMethodDetails apijson.Field
	PaymentType          apijson.Field
	RecurringToken       apijson.Field
	RedirectURL          apijson.Field
	ReferenceID          apijson.Field
	StartRecurring       apijson.Field
	State                apijson.Field
	TerminalName         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PaymentListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentListResponseResultJSON) RawJSON() string {
	return r.raw
}

// Customer's billing address
type PaymentListResponseResultBillingAddress struct {
	// Line 1 of the address (e.g., Number, street, etc)
	AddressLine1 string `json:"addressLine1"`
	// Line 2 of the address (e.g., Suite, apt)
	AddressLine2 string `json:"addressLine2"`
	// City name
	City string `json:"city"`
	// 2-character IS0-3166-1 country code
	CountryCode string `json:"countryCode"`
	// Postal code
	PostalCode string `json:"postalCode"`
	// State code
	State string                                      `json:"state"`
	JSON  paymentListResponseResultBillingAddressJSON `json:"-"`
}

// paymentListResponseResultBillingAddressJSON contains the JSON metadata for the
// struct [PaymentListResponseResultBillingAddress]
type paymentListResponseResultBillingAddressJSON struct {
	AddressLine1 apijson.Field
	AddressLine2 apijson.Field
	City         apijson.Field
	CountryCode  apijson.Field
	PostalCode   apijson.Field
	State        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PaymentListResponseResultBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentListResponseResultBillingAddressJSON) RawJSON() string {
	return r.raw
}

type PaymentListResponseResultCustomer struct {
	// Customer account name in the provider's system. Used for some types of
	// withdrawals.
	AccountName string `json:"accountName"`
	// Customer account number in the provider's system. Used for some types of
	// withdrawals.
	AccountNumber string `json:"accountNumber"`
	// Customer bank. Used for some types of withdrawals.
	Bank string `json:"bank"`
	// Customer bank branch. Used for some types of withdrawals.
	BankBranch string `json:"bankBranch"`
	// Customer country of citizenship
	CitizenshipCountryCode string `json:"citizenshipCountryCode"`
	DateOfBirth            string `json:"dateOfBirth" format:"ISO 8601 (YYYY-MM-DD)"`
	// Date of the first deposit from the customer
	DateOfFirstDeposit string `json:"dateOfFirstDeposit" format:"ISO 8601 (YYYY-MM-DD)"`
	// How much the customer has deposited, in base currency
	DepositsAmount int64 `json:"depositsAmount"`
	// How many times the customer made a deposit
	DepositsCnt int64 `json:"depositsCnt"`
	// An identifier for the customer assigned by a government authority
	DocumentNumber string `json:"documentNumber"`
	// Document Type
	DocumentType PaymentListResponseResultCustomerDocumentType `json:"documentType"`
	// Email address of the customer
	Email     string `json:"email" format:"email"`
	FirstName string `json:"firstName"`
	// Indicates whether the customer has passed KYC verification
	KYCStatus bool   `json:"kycStatus"`
	LastName  string `json:"lastName"`
	// Customer preferred display language
	Locale string `json:"locale"`
	// Indicates whether the payment instrument (usually the card number) has passed
	// KYC verification
	PaymentInstrumentKYCStatus bool `json:"paymentInstrumentKycStatus"`
	// International phone number of the customer, without the '+'. Use a space as a
	// separator between the dialing country code and local phone number.
	Phone string `json:"phone"`
	// Id of the customer assigned by Merchant
	ReferenceID string `json:"referenceId"`
	// Identify the customer as belonging to a specific group that is used for routing
	RoutingGroup string `json:"routingGroup"`
	// How much the customer has withdrawn, in base currency
	WithdrawalsAmount int64 `json:"withdrawalsAmount"`
	// How many times the customer made a withdrawal
	WithdrawalsCnt int64                                 `json:"withdrawalsCnt"`
	JSON           paymentListResponseResultCustomerJSON `json:"-"`
}

// paymentListResponseResultCustomerJSON contains the JSON metadata for the struct
// [PaymentListResponseResultCustomer]
type paymentListResponseResultCustomerJSON struct {
	AccountName                apijson.Field
	AccountNumber              apijson.Field
	Bank                       apijson.Field
	BankBranch                 apijson.Field
	CitizenshipCountryCode     apijson.Field
	DateOfBirth                apijson.Field
	DateOfFirstDeposit         apijson.Field
	DepositsAmount             apijson.Field
	DepositsCnt                apijson.Field
	DocumentNumber             apijson.Field
	DocumentType               apijson.Field
	Email                      apijson.Field
	FirstName                  apijson.Field
	KYCStatus                  apijson.Field
	LastName                   apijson.Field
	Locale                     apijson.Field
	PaymentInstrumentKYCStatus apijson.Field
	Phone                      apijson.Field
	ReferenceID                apijson.Field
	RoutingGroup               apijson.Field
	WithdrawalsAmount          apijson.Field
	WithdrawalsCnt             apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *PaymentListResponseResultCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentListResponseResultCustomerJSON) RawJSON() string {
	return r.raw
}

// Document Type
type PaymentListResponseResultCustomerDocumentType string

const (
	PaymentListResponseResultCustomerDocumentTypeArCdi  PaymentListResponseResultCustomerDocumentType = "AR_CDI"
	PaymentListResponseResultCustomerDocumentTypeArCuil PaymentListResponseResultCustomerDocumentType = "AR_CUIL"
	PaymentListResponseResultCustomerDocumentTypeArCuit PaymentListResponseResultCustomerDocumentType = "AR_CUIT"
	PaymentListResponseResultCustomerDocumentTypeArDni  PaymentListResponseResultCustomerDocumentType = "AR_DNI"
	PaymentListResponseResultCustomerDocumentTypeArOtro PaymentListResponseResultCustomerDocumentType = "AR_OTRO"
	PaymentListResponseResultCustomerDocumentTypeBrCnpj PaymentListResponseResultCustomerDocumentType = "BR_CNPJ"
	PaymentListResponseResultCustomerDocumentTypeBrCpf  PaymentListResponseResultCustomerDocumentType = "BR_CPF"
	PaymentListResponseResultCustomerDocumentTypeClOtro PaymentListResponseResultCustomerDocumentType = "CL_OTRO"
	PaymentListResponseResultCustomerDocumentTypeClRun  PaymentListResponseResultCustomerDocumentType = "CL_RUN"
	PaymentListResponseResultCustomerDocumentTypeClRut  PaymentListResponseResultCustomerDocumentType = "CL_RUT"
	PaymentListResponseResultCustomerDocumentTypeCoCc   PaymentListResponseResultCustomerDocumentType = "CO_CC"
	PaymentListResponseResultCustomerDocumentTypeCoCe   PaymentListResponseResultCustomerDocumentType = "CO_CE"
	PaymentListResponseResultCustomerDocumentTypeCoDl   PaymentListResponseResultCustomerDocumentType = "CO_DL"
	PaymentListResponseResultCustomerDocumentTypeCoDni  PaymentListResponseResultCustomerDocumentType = "CO_DNI"
	PaymentListResponseResultCustomerDocumentTypeCoNe   PaymentListResponseResultCustomerDocumentType = "CO_NE"
	PaymentListResponseResultCustomerDocumentTypeCoNit  PaymentListResponseResultCustomerDocumentType = "CO_NIT"
	PaymentListResponseResultCustomerDocumentTypeCoPp   PaymentListResponseResultCustomerDocumentType = "CO_PP"
	PaymentListResponseResultCustomerDocumentTypeCoSS   PaymentListResponseResultCustomerDocumentType = "CO_SS"
	PaymentListResponseResultCustomerDocumentTypeCoTi   PaymentListResponseResultCustomerDocumentType = "CO_TI"
	PaymentListResponseResultCustomerDocumentTypeCrCdi  PaymentListResponseResultCustomerDocumentType = "CR_CDI"
	PaymentListResponseResultCustomerDocumentTypeEcDni  PaymentListResponseResultCustomerDocumentType = "EC_DNI"
	PaymentListResponseResultCustomerDocumentTypeEcPp   PaymentListResponseResultCustomerDocumentType = "EC_PP"
	PaymentListResponseResultCustomerDocumentTypeEcRuc  PaymentListResponseResultCustomerDocumentType = "EC_RUC"
	PaymentListResponseResultCustomerDocumentTypeGtCui  PaymentListResponseResultCustomerDocumentType = "GT_CUI"
	PaymentListResponseResultCustomerDocumentTypeGtDpi  PaymentListResponseResultCustomerDocumentType = "GT_DPI"
	PaymentListResponseResultCustomerDocumentTypeGtNit  PaymentListResponseResultCustomerDocumentType = "GT_NIT"
	PaymentListResponseResultCustomerDocumentTypeMxCurp PaymentListResponseResultCustomerDocumentType = "MX_CURP"
	PaymentListResponseResultCustomerDocumentTypeMxIfe  PaymentListResponseResultCustomerDocumentType = "MX_IFE"
	PaymentListResponseResultCustomerDocumentTypeMxPp   PaymentListResponseResultCustomerDocumentType = "MX_PP"
	PaymentListResponseResultCustomerDocumentTypeMxRfc  PaymentListResponseResultCustomerDocumentType = "MX_RFC"
	PaymentListResponseResultCustomerDocumentTypePaCip  PaymentListResponseResultCustomerDocumentType = "PA_CIP"
	PaymentListResponseResultCustomerDocumentTypePeCe   PaymentListResponseResultCustomerDocumentType = "PE_CE"
	PaymentListResponseResultCustomerDocumentTypePeDni  PaymentListResponseResultCustomerDocumentType = "PE_DNI"
	PaymentListResponseResultCustomerDocumentTypePeOtro PaymentListResponseResultCustomerDocumentType = "PE_OTRO"
	PaymentListResponseResultCustomerDocumentTypePePp   PaymentListResponseResultCustomerDocumentType = "PE_PP"
	PaymentListResponseResultCustomerDocumentTypePeRuc  PaymentListResponseResultCustomerDocumentType = "PE_RUC"
)

func (r PaymentListResponseResultCustomerDocumentType) IsKnown() bool {
	switch r {
	case PaymentListResponseResultCustomerDocumentTypeArCdi, PaymentListResponseResultCustomerDocumentTypeArCuil, PaymentListResponseResultCustomerDocumentTypeArCuit, PaymentListResponseResultCustomerDocumentTypeArDni, PaymentListResponseResultCustomerDocumentTypeArOtro, PaymentListResponseResultCustomerDocumentTypeBrCnpj, PaymentListResponseResultCustomerDocumentTypeBrCpf, PaymentListResponseResultCustomerDocumentTypeClOtro, PaymentListResponseResultCustomerDocumentTypeClRun, PaymentListResponseResultCustomerDocumentTypeClRut, PaymentListResponseResultCustomerDocumentTypeCoCc, PaymentListResponseResultCustomerDocumentTypeCoCe, PaymentListResponseResultCustomerDocumentTypeCoDl, PaymentListResponseResultCustomerDocumentTypeCoDni, PaymentListResponseResultCustomerDocumentTypeCoNe, PaymentListResponseResultCustomerDocumentTypeCoNit, PaymentListResponseResultCustomerDocumentTypeCoPp, PaymentListResponseResultCustomerDocumentTypeCoSS, PaymentListResponseResultCustomerDocumentTypeCoTi, PaymentListResponseResultCustomerDocumentTypeCrCdi, PaymentListResponseResultCustomerDocumentTypeEcDni, PaymentListResponseResultCustomerDocumentTypeEcPp, PaymentListResponseResultCustomerDocumentTypeEcRuc, PaymentListResponseResultCustomerDocumentTypeGtCui, PaymentListResponseResultCustomerDocumentTypeGtDpi, PaymentListResponseResultCustomerDocumentTypeGtNit, PaymentListResponseResultCustomerDocumentTypeMxCurp, PaymentListResponseResultCustomerDocumentTypeMxIfe, PaymentListResponseResultCustomerDocumentTypeMxPp, PaymentListResponseResultCustomerDocumentTypeMxRfc, PaymentListResponseResultCustomerDocumentTypePaCip, PaymentListResponseResultCustomerDocumentTypePeCe, PaymentListResponseResultCustomerDocumentTypePeDni, PaymentListResponseResultCustomerDocumentTypePeOtro, PaymentListResponseResultCustomerDocumentTypePePp, PaymentListResponseResultCustomerDocumentTypePeRuc:
		return true
	}
	return false
}

// Payment Method
type PaymentListResponseResultPaymentMethod string

const (
	PaymentListResponseResultPaymentMethodBasicCard            PaymentListResponseResultPaymentMethod = "BASIC_CARD"
	PaymentListResponseResultPaymentMethodCrypto               PaymentListResponseResultPaymentMethod = "CRYPTO"
	PaymentListResponseResultPaymentMethodFlexepin             PaymentListResponseResultPaymentMethod = "FLEXEPIN"
	PaymentListResponseResultPaymentMethodMacropay             PaymentListResponseResultPaymentMethod = "MACROPAY"
	PaymentListResponseResultPaymentMethodSkrill               PaymentListResponseResultPaymentMethod = "SKRILL"
	PaymentListResponseResultPaymentMethodPayretailers         PaymentListResponseResultPaymentMethod = "PAYRETAILERS"
	PaymentListResponseResultPaymentMethodLocalpayment         PaymentListResponseResultPaymentMethod = "LOCALPAYMENT"
	PaymentListResponseResultPaymentMethodMonnet               PaymentListResponseResultPaymentMethod = "MONNET"
	PaymentListResponseResultPaymentMethodPaypal               PaymentListResponseResultPaymentMethod = "PAYPAL"
	PaymentListResponseResultPaymentMethodNeteller             PaymentListResponseResultPaymentMethod = "NETELLER"
	PaymentListResponseResultPaymentMethodTrustpayments        PaymentListResponseResultPaymentMethod = "TRUSTPAYMENTS"
	PaymentListResponseResultPaymentMethodPaymaxis             PaymentListResponseResultPaymentMethod = "PAYMAXIS"
	PaymentListResponseResultPaymentMethodGate8Transact        PaymentListResponseResultPaymentMethod = "GATE8TRANSACT"
	PaymentListResponseResultPaymentMethodTink                 PaymentListResponseResultPaymentMethod = "TINK"
	PaymentListResponseResultPaymentMethodB2Binpay             PaymentListResponseResultPaymentMethod = "B2BINPAY"
	PaymentListResponseResultPaymentMethodClick                PaymentListResponseResultPaymentMethod = "CLICK"
	PaymentListResponseResultPaymentMethodMonetix              PaymentListResponseResultPaymentMethod = "MONETIX"
	PaymentListResponseResultPaymentMethodPerfectmoney         PaymentListResponseResultPaymentMethod = "PERFECTMONEY"
	PaymentListResponseResultPaymentMethodVolt                 PaymentListResponseResultPaymentMethod = "VOLT"
	PaymentListResponseResultPaymentMethodKesspay              PaymentListResponseResultPaymentMethod = "KESSPAY"
	PaymentListResponseResultPaymentMethodBillline             PaymentListResponseResultPaymentMethod = "BILLLINE"
	PaymentListResponseResultPaymentMethodNgenius              PaymentListResponseResultPaymentMethod = "NGENIUS"
	PaymentListResponseResultPaymentMethodAstropay             PaymentListResponseResultPaymentMethod = "ASTROPAY"
	PaymentListResponseResultPaymentMethodAlycepay             PaymentListResponseResultPaymentMethod = "ALYCEPAY"
	PaymentListResponseResultPaymentMethodPix                  PaymentListResponseResultPaymentMethod = "PIX"
	PaymentListResponseResultPaymentMethodBoleto               PaymentListResponseResultPaymentMethod = "BOLETO"
	PaymentListResponseResultPaymentMethodUpi                  PaymentListResponseResultPaymentMethod = "UPI"
	PaymentListResponseResultPaymentMethodPaytm                PaymentListResponseResultPaymentMethod = "PAYTM"
	PaymentListResponseResultPaymentMethodNetbanking           PaymentListResponseResultPaymentMethod = "NETBANKING"
	PaymentListResponseResultPaymentMethodFinrax               PaymentListResponseResultPaymentMethod = "FINRAX"
	PaymentListResponseResultPaymentMethodSpoynt               PaymentListResponseResultPaymentMethod = "SPOYNT"
	PaymentListResponseResultPaymentMethodXinpay               PaymentListResponseResultPaymentMethod = "XINPAY"
	PaymentListResponseResultPaymentMethodOmnimatrix           PaymentListResponseResultPaymentMethod = "OMNIMATRIX"
	PaymentListResponseResultPaymentMethodDpopay               PaymentListResponseResultPaymentMethod = "DPOPAY"
	PaymentListResponseResultPaymentMethodExternalHpp          PaymentListResponseResultPaymentMethod = "EXTERNAL_HPP"
	PaymentListResponseResultPaymentMethodXanpay               PaymentListResponseResultPaymentMethod = "XANPAY"
	PaymentListResponseResultPaymentMethodInrpay               PaymentListResponseResultPaymentMethod = "INRPAY"
	PaymentListResponseResultPaymentMethodAri10                PaymentListResponseResultPaymentMethod = "ARI10"
	PaymentListResponseResultPaymentMethodSofort               PaymentListResponseResultPaymentMethod = "SOFORT"
	PaymentListResponseResultPaymentMethodGiropay              PaymentListResponseResultPaymentMethod = "GIROPAY"
	PaymentListResponseResultPaymentMethodPaysafecard          PaymentListResponseResultPaymentMethod = "PAYSAFECARD"
	PaymentListResponseResultPaymentMethodPaysafecash          PaymentListResponseResultPaymentMethod = "PAYSAFECASH"
	PaymentListResponseResultPaymentMethodOpenBanking          PaymentListResponseResultPaymentMethod = "OPEN_BANKING"
	PaymentListResponseResultPaymentMethodKlarna               PaymentListResponseResultPaymentMethod = "KLARNA"
	PaymentListResponseResultPaymentMethodSpei                 PaymentListResponseResultPaymentMethod = "SPEI"
	PaymentListResponseResultPaymentMethodPaycash              PaymentListResponseResultPaymentMethod = "PAYCASH"
	PaymentListResponseResultPaymentMethodRapipago             PaymentListResponseResultPaymentMethod = "RAPIPAGO"
	PaymentListResponseResultPaymentMethodPagofacil            PaymentListResponseResultPaymentMethod = "PAGOFACIL"
	PaymentListResponseResultPaymentMethodRapidtransfer        PaymentListResponseResultPaymentMethod = "RAPIDTRANSFER"
	PaymentListResponseResultPaymentMethodMobileMoney          PaymentListResponseResultPaymentMethod = "MOBILE_MONEY"
	PaymentListResponseResultPaymentMethodInterac              PaymentListResponseResultPaymentMethod = "INTERAC"
	PaymentListResponseResultPaymentMethodInteracEto           PaymentListResponseResultPaymentMethod = "INTERAC_ETO"
	PaymentListResponseResultPaymentMethodInteracRto           PaymentListResponseResultPaymentMethod = "INTERAC_RTO"
	PaymentListResponseResultPaymentMethodInteracACH           PaymentListResponseResultPaymentMethod = "INTERAC_ACH"
	PaymentListResponseResultPaymentMethodPicpay               PaymentListResponseResultPaymentMethod = "PICPAY"
	PaymentListResponseResultPaymentMethodMollie               PaymentListResponseResultPaymentMethod = "MOLLIE"
	PaymentListResponseResultPaymentMethodTed                  PaymentListResponseResultPaymentMethod = "TED"
	PaymentListResponseResultPaymentMethodZipay                PaymentListResponseResultPaymentMethod = "ZIPAY"
	PaymentListResponseResultPaymentMethodPse                  PaymentListResponseResultPaymentMethod = "PSE"
	PaymentListResponseResultPaymentMethodEfecty               PaymentListResponseResultPaymentMethod = "EFECTY"
	PaymentListResponseResultPaymentMethodBanktransfer         PaymentListResponseResultPaymentMethod = "BANKTRANSFER"
	PaymentListResponseResultPaymentMethodPec                  PaymentListResponseResultPaymentMethod = "PEC"
	PaymentListResponseResultPaymentMethodOxxo                 PaymentListResponseResultPaymentMethod = "OXXO"
	PaymentListResponseResultPaymentMethodWebpay               PaymentListResponseResultPaymentMethod = "WEBPAY"
	PaymentListResponseResultPaymentMethodPagoefectivo         PaymentListResponseResultPaymentMethod = "PAGOEFECTIVO"
	PaymentListResponseResultPaymentMethodMifinity             PaymentListResponseResultPaymentMethod = "MIFINITY"
	PaymentListResponseResultPaymentMethodPayport              PaymentListResponseResultPaymentMethod = "PAYPORT"
	PaymentListResponseResultPaymentMethodJetoncash            PaymentListResponseResultPaymentMethod = "JETONCASH"
	PaymentListResponseResultPaymentMethodJetonwallet          PaymentListResponseResultPaymentMethod = "JETONWALLET"
	PaymentListResponseResultPaymentMethodNoda                 PaymentListResponseResultPaymentMethod = "NODA"
	PaymentListResponseResultPaymentMethodNodaRevolut          PaymentListResponseResultPaymentMethod = "NODA_REVOLUT"
	PaymentListResponseResultPaymentMethodAlfakit              PaymentListResponseResultPaymentMethod = "ALFAKIT"
	PaymentListResponseResultPaymentMethodPayfun               PaymentListResponseResultPaymentMethod = "PAYFUN"
	PaymentListResponseResultPaymentMethodEmanat               PaymentListResponseResultPaymentMethod = "EMANAT"
	PaymentListResponseResultPaymentMethodM10                  PaymentListResponseResultPaymentMethod = "M10"
	PaymentListResponseResultPaymentMethodRubpay               PaymentListResponseResultPaymentMethod = "RUBPAY"
	PaymentListResponseResultPaymentMethodMonerchy             PaymentListResponseResultPaymentMethod = "MONERCHY"
	PaymentListResponseResultPaymentMethodMuchbetter           PaymentListResponseResultPaymentMethod = "MUCHBETTER"
	PaymentListResponseResultPaymentMethodYapily               PaymentListResponseResultPaymentMethod = "YAPILY"
	PaymentListResponseResultPaymentMethodInai                 PaymentListResponseResultPaymentMethod = "INAI"
	PaymentListResponseResultPaymentMethodImps                 PaymentListResponseResultPaymentMethod = "IMPS"
	PaymentListResponseResultPaymentMethodRtgs                 PaymentListResponseResultPaymentMethod = "RTGS"
	PaymentListResponseResultPaymentMethodPayid                PaymentListResponseResultPaymentMethod = "PAYID"
	PaymentListResponseResultPaymentMethodZotapay              PaymentListResponseResultPaymentMethod = "ZOTAPAY"
	PaymentListResponseResultPaymentMethodSbp                  PaymentListResponseResultPaymentMethod = "SBP"
	PaymentListResponseResultPaymentMethodP2PCard              PaymentListResponseResultPaymentMethod = "P2P_CARD"
	PaymentListResponseResultPaymentMethodP2PIban              PaymentListResponseResultPaymentMethod = "P2P_IBAN"
	PaymentListResponseResultPaymentMethodP2PSbp               PaymentListResponseResultPaymentMethod = "P2P_SBP"
	PaymentListResponseResultPaymentMethodP2PMobile            PaymentListResponseResultPaymentMethod = "P2P_MOBILE"
	PaymentListResponseResultPaymentMethodPush                 PaymentListResponseResultPaymentMethod = "PUSH"
	PaymentListResponseResultPaymentMethodGateiq               PaymentListResponseResultPaymentMethod = "GATEIQ"
	PaymentListResponseResultPaymentMethodViettel              PaymentListResponseResultPaymentMethod = "VIETTEL"
	PaymentListResponseResultPaymentMethodZalo                 PaymentListResponseResultPaymentMethod = "ZALO"
	PaymentListResponseResultPaymentMethodQr                   PaymentListResponseResultPaymentMethod = "QR"
	PaymentListResponseResultPaymentMethodCup                  PaymentListResponseResultPaymentMethod = "CUP"
	PaymentListResponseResultPaymentMethodCodi                 PaymentListResponseResultPaymentMethod = "CODI"
	PaymentListResponseResultPaymentMethodPay2Play             PaymentListResponseResultPaymentMethod = "PAY2PLAY"
	PaymentListResponseResultPaymentMethodBkash                PaymentListResponseResultPaymentMethod = "BKASH"
	PaymentListResponseResultPaymentMethodNagad                PaymentListResponseResultPaymentMethod = "NAGAD"
	PaymentListResponseResultPaymentMethodRocket               PaymentListResponseResultPaymentMethod = "ROCKET"
	PaymentListResponseResultPaymentMethodVirtualAccount       PaymentListResponseResultPaymentMethod = "VIRTUAL_ACCOUNT"
	PaymentListResponseResultPaymentMethodMultibanco           PaymentListResponseResultPaymentMethod = "MULTIBANCO"
	PaymentListResponseResultPaymentMethodBlik                 PaymentListResponseResultPaymentMethod = "BLIK"
	PaymentListResponseResultPaymentMethodMbway                PaymentListResponseResultPaymentMethod = "MBWAY"
	PaymentListResponseResultPaymentMethodP24                  PaymentListResponseResultPaymentMethod = "P24"
	PaymentListResponseResultPaymentMethodMistercash           PaymentListResponseResultPaymentMethod = "MISTERCASH"
	PaymentListResponseResultPaymentMethodMach                 PaymentListResponseResultPaymentMethod = "MACH"
	PaymentListResponseResultPaymentMethodKhipu                PaymentListResponseResultPaymentMethod = "KHIPU"
	PaymentListResponseResultPaymentMethodNeft                 PaymentListResponseResultPaymentMethod = "NEFT"
	PaymentListResponseResultPaymentMethodSticpay              PaymentListResponseResultPaymentMethod = "STICPAY"
	PaymentListResponseResultPaymentMethodSberpay              PaymentListResponseResultPaymentMethod = "SBERPAY"
	PaymentListResponseResultPaymentMethodMobileCommerce       PaymentListResponseResultPaymentMethod = "MOBILE_COMMERCE"
	PaymentListResponseResultPaymentMethodBinancePay           PaymentListResponseResultPaymentMethod = "BINANCE_PAY"
	PaymentListResponseResultPaymentMethodMpay                 PaymentListResponseResultPaymentMethod = "MPAY"
	PaymentListResponseResultPaymentMethodChek                 PaymentListResponseResultPaymentMethod = "CHEK"
	PaymentListResponseResultPaymentMethodKlapEfectivo         PaymentListResponseResultPaymentMethod = "KLAP_EFECTIVO"
	PaymentListResponseResultPaymentMethodKlapTransferencia    PaymentListResponseResultPaymentMethod = "KLAP_TRANSFERENCIA"
	PaymentListResponseResultPaymentMethodPago46               PaymentListResponseResultPaymentMethod = "PAGO46"
	PaymentListResponseResultPaymentMethodHites                PaymentListResponseResultPaymentMethod = "HITES"
	PaymentListResponseResultPaymentMethodServifacil           PaymentListResponseResultPaymentMethod = "SERVIFACIL"
	PaymentListResponseResultPaymentMethodOpenpayd             PaymentListResponseResultPaymentMethod = "OPENPAYD"
	PaymentListResponseResultPaymentMethodFawry                PaymentListResponseResultPaymentMethod = "FAWRY"
	PaymentListResponseResultPaymentMethodEps                  PaymentListResponseResultPaymentMethod = "EPS"
	PaymentListResponseResultPaymentMethodIdeal                PaymentListResponseResultPaymentMethod = "IDEAL"
	PaymentListResponseResultPaymentMethodTrustly              PaymentListResponseResultPaymentMethod = "TRUSTLY"
	PaymentListResponseResultPaymentMethodUssd                 PaymentListResponseResultPaymentMethod = "USSD"
	PaymentListResponseResultPaymentMethodMpesa                PaymentListResponseResultPaymentMethod = "MPESA"
	PaymentListResponseResultPaymentMethodEnaira               PaymentListResponseResultPaymentMethod = "ENAIRA"
	PaymentListResponseResultPaymentMethodOnevoucher           PaymentListResponseResultPaymentMethod = "ONEVOUCHER"
	PaymentListResponseResultPaymentMethodBancontact           PaymentListResponseResultPaymentMethod = "BANCONTACT"
	PaymentListResponseResultPaymentMethodSwish                PaymentListResponseResultPaymentMethod = "SWISH"
	PaymentListResponseResultPaymentMethodEft                  PaymentListResponseResultPaymentMethod = "EFT"
	PaymentListResponseResultPaymentMethodGcash                PaymentListResponseResultPaymentMethod = "GCASH"
	PaymentListResponseResultPaymentMethodPaymaya              PaymentListResponseResultPaymentMethod = "PAYMAYA"
	PaymentListResponseResultPaymentMethodPagoMovil            PaymentListResponseResultPaymentMethod = "PAGO_MOVIL"
	PaymentListResponseResultPaymentMethodPagoMovilInst        PaymentListResponseResultPaymentMethod = "PAGO_MOVIL_INST"
	PaymentListResponseResultPaymentMethodBiopago              PaymentListResponseResultPaymentMethod = "BIOPAGO"
	PaymentListResponseResultPaymentMethodCash                 PaymentListResponseResultPaymentMethod = "CASH"
	PaymentListResponseResultPaymentMethodVoucherry            PaymentListResponseResultPaymentMethod = "VOUCHERRY"
	PaymentListResponseResultPaymentMethodApplepay             PaymentListResponseResultPaymentMethod = "APPLEPAY"
	PaymentListResponseResultPaymentMethodGooglepay            PaymentListResponseResultPaymentMethod = "GOOGLEPAY"
	PaymentListResponseResultPaymentMethodBrite                PaymentListResponseResultPaymentMethod = "BRITE"
	PaymentListResponseResultPaymentMethodVouchstar            PaymentListResponseResultPaymentMethod = "VOUCHSTAR"
	PaymentListResponseResultPaymentMethodRevolut              PaymentListResponseResultPaymentMethod = "REVOLUT"
	PaymentListResponseResultPaymentMethodOnlineBanking        PaymentListResponseResultPaymentMethod = "ONLINE_BANKING"
	PaymentListResponseResultPaymentMethodPromptpay            PaymentListResponseResultPaymentMethod = "PROMPTPAY"
	PaymentListResponseResultPaymentMethodTruemoney            PaymentListResponseResultPaymentMethod = "TRUEMONEY"
	PaymentListResponseResultPaymentMethodMomopayVn            PaymentListResponseResultPaymentMethod = "MOMOPAY_VN"
	PaymentListResponseResultPaymentMethodMomopayRw            PaymentListResponseResultPaymentMethod = "MOMOPAY_RW"
	PaymentListResponseResultPaymentMethodVnpayQr              PaymentListResponseResultPaymentMethod = "VNPAY_QR"
	PaymentListResponseResultPaymentMethodN26                  PaymentListResponseResultPaymentMethod = "N26"
	PaymentListResponseResultPaymentMethodWise                 PaymentListResponseResultPaymentMethod = "WISE"
	PaymentListResponseResultPaymentMethodPaydoWallet          PaymentListResponseResultPaymentMethod = "PAYDO_WALLET"
	PaymentListResponseResultPaymentMethodPapara               PaymentListResponseResultPaymentMethod = "PAPARA"
	PaymentListResponseResultPaymentMethodPayoutSepaBatch      PaymentListResponseResultPaymentMethod = "PAYOUT_SEPA_BATCH"
	PaymentListResponseResultPaymentMethodPayoutNonsepaRequest PaymentListResponseResultPaymentMethod = "PAYOUT_NONSEPA_REQUEST"
)

func (r PaymentListResponseResultPaymentMethod) IsKnown() bool {
	switch r {
	case PaymentListResponseResultPaymentMethodBasicCard, PaymentListResponseResultPaymentMethodCrypto, PaymentListResponseResultPaymentMethodFlexepin, PaymentListResponseResultPaymentMethodMacropay, PaymentListResponseResultPaymentMethodSkrill, PaymentListResponseResultPaymentMethodPayretailers, PaymentListResponseResultPaymentMethodLocalpayment, PaymentListResponseResultPaymentMethodMonnet, PaymentListResponseResultPaymentMethodPaypal, PaymentListResponseResultPaymentMethodNeteller, PaymentListResponseResultPaymentMethodTrustpayments, PaymentListResponseResultPaymentMethodPaymaxis, PaymentListResponseResultPaymentMethodGate8Transact, PaymentListResponseResultPaymentMethodTink, PaymentListResponseResultPaymentMethodB2Binpay, PaymentListResponseResultPaymentMethodClick, PaymentListResponseResultPaymentMethodMonetix, PaymentListResponseResultPaymentMethodPerfectmoney, PaymentListResponseResultPaymentMethodVolt, PaymentListResponseResultPaymentMethodKesspay, PaymentListResponseResultPaymentMethodBillline, PaymentListResponseResultPaymentMethodNgenius, PaymentListResponseResultPaymentMethodAstropay, PaymentListResponseResultPaymentMethodAlycepay, PaymentListResponseResultPaymentMethodPix, PaymentListResponseResultPaymentMethodBoleto, PaymentListResponseResultPaymentMethodUpi, PaymentListResponseResultPaymentMethodPaytm, PaymentListResponseResultPaymentMethodNetbanking, PaymentListResponseResultPaymentMethodFinrax, PaymentListResponseResultPaymentMethodSpoynt, PaymentListResponseResultPaymentMethodXinpay, PaymentListResponseResultPaymentMethodOmnimatrix, PaymentListResponseResultPaymentMethodDpopay, PaymentListResponseResultPaymentMethodExternalHpp, PaymentListResponseResultPaymentMethodXanpay, PaymentListResponseResultPaymentMethodInrpay, PaymentListResponseResultPaymentMethodAri10, PaymentListResponseResultPaymentMethodSofort, PaymentListResponseResultPaymentMethodGiropay, PaymentListResponseResultPaymentMethodPaysafecard, PaymentListResponseResultPaymentMethodPaysafecash, PaymentListResponseResultPaymentMethodOpenBanking, PaymentListResponseResultPaymentMethodKlarna, PaymentListResponseResultPaymentMethodSpei, PaymentListResponseResultPaymentMethodPaycash, PaymentListResponseResultPaymentMethodRapipago, PaymentListResponseResultPaymentMethodPagofacil, PaymentListResponseResultPaymentMethodRapidtransfer, PaymentListResponseResultPaymentMethodMobileMoney, PaymentListResponseResultPaymentMethodInterac, PaymentListResponseResultPaymentMethodInteracEto, PaymentListResponseResultPaymentMethodInteracRto, PaymentListResponseResultPaymentMethodInteracACH, PaymentListResponseResultPaymentMethodPicpay, PaymentListResponseResultPaymentMethodMollie, PaymentListResponseResultPaymentMethodTed, PaymentListResponseResultPaymentMethodZipay, PaymentListResponseResultPaymentMethodPse, PaymentListResponseResultPaymentMethodEfecty, PaymentListResponseResultPaymentMethodBanktransfer, PaymentListResponseResultPaymentMethodPec, PaymentListResponseResultPaymentMethodOxxo, PaymentListResponseResultPaymentMethodWebpay, PaymentListResponseResultPaymentMethodPagoefectivo, PaymentListResponseResultPaymentMethodMifinity, PaymentListResponseResultPaymentMethodPayport, PaymentListResponseResultPaymentMethodJetoncash, PaymentListResponseResultPaymentMethodJetonwallet, PaymentListResponseResultPaymentMethodNoda, PaymentListResponseResultPaymentMethodNodaRevolut, PaymentListResponseResultPaymentMethodAlfakit, PaymentListResponseResultPaymentMethodPayfun, PaymentListResponseResultPaymentMethodEmanat, PaymentListResponseResultPaymentMethodM10, PaymentListResponseResultPaymentMethodRubpay, PaymentListResponseResultPaymentMethodMonerchy, PaymentListResponseResultPaymentMethodMuchbetter, PaymentListResponseResultPaymentMethodYapily, PaymentListResponseResultPaymentMethodInai, PaymentListResponseResultPaymentMethodImps, PaymentListResponseResultPaymentMethodRtgs, PaymentListResponseResultPaymentMethodPayid, PaymentListResponseResultPaymentMethodZotapay, PaymentListResponseResultPaymentMethodSbp, PaymentListResponseResultPaymentMethodP2PCard, PaymentListResponseResultPaymentMethodP2PIban, PaymentListResponseResultPaymentMethodP2PSbp, PaymentListResponseResultPaymentMethodP2PMobile, PaymentListResponseResultPaymentMethodPush, PaymentListResponseResultPaymentMethodGateiq, PaymentListResponseResultPaymentMethodViettel, PaymentListResponseResultPaymentMethodZalo, PaymentListResponseResultPaymentMethodQr, PaymentListResponseResultPaymentMethodCup, PaymentListResponseResultPaymentMethodCodi, PaymentListResponseResultPaymentMethodPay2Play, PaymentListResponseResultPaymentMethodBkash, PaymentListResponseResultPaymentMethodNagad, PaymentListResponseResultPaymentMethodRocket, PaymentListResponseResultPaymentMethodVirtualAccount, PaymentListResponseResultPaymentMethodMultibanco, PaymentListResponseResultPaymentMethodBlik, PaymentListResponseResultPaymentMethodMbway, PaymentListResponseResultPaymentMethodP24, PaymentListResponseResultPaymentMethodMistercash, PaymentListResponseResultPaymentMethodMach, PaymentListResponseResultPaymentMethodKhipu, PaymentListResponseResultPaymentMethodNeft, PaymentListResponseResultPaymentMethodSticpay, PaymentListResponseResultPaymentMethodSberpay, PaymentListResponseResultPaymentMethodMobileCommerce, PaymentListResponseResultPaymentMethodBinancePay, PaymentListResponseResultPaymentMethodMpay, PaymentListResponseResultPaymentMethodChek, PaymentListResponseResultPaymentMethodKlapEfectivo, PaymentListResponseResultPaymentMethodKlapTransferencia, PaymentListResponseResultPaymentMethodPago46, PaymentListResponseResultPaymentMethodHites, PaymentListResponseResultPaymentMethodServifacil, PaymentListResponseResultPaymentMethodOpenpayd, PaymentListResponseResultPaymentMethodFawry, PaymentListResponseResultPaymentMethodEps, PaymentListResponseResultPaymentMethodIdeal, PaymentListResponseResultPaymentMethodTrustly, PaymentListResponseResultPaymentMethodUssd, PaymentListResponseResultPaymentMethodMpesa, PaymentListResponseResultPaymentMethodEnaira, PaymentListResponseResultPaymentMethodOnevoucher, PaymentListResponseResultPaymentMethodBancontact, PaymentListResponseResultPaymentMethodSwish, PaymentListResponseResultPaymentMethodEft, PaymentListResponseResultPaymentMethodGcash, PaymentListResponseResultPaymentMethodPaymaya, PaymentListResponseResultPaymentMethodPagoMovil, PaymentListResponseResultPaymentMethodPagoMovilInst, PaymentListResponseResultPaymentMethodBiopago, PaymentListResponseResultPaymentMethodCash, PaymentListResponseResultPaymentMethodVoucherry, PaymentListResponseResultPaymentMethodApplepay, PaymentListResponseResultPaymentMethodGooglepay, PaymentListResponseResultPaymentMethodBrite, PaymentListResponseResultPaymentMethodVouchstar, PaymentListResponseResultPaymentMethodRevolut, PaymentListResponseResultPaymentMethodOnlineBanking, PaymentListResponseResultPaymentMethodPromptpay, PaymentListResponseResultPaymentMethodTruemoney, PaymentListResponseResultPaymentMethodMomopayVn, PaymentListResponseResultPaymentMethodMomopayRw, PaymentListResponseResultPaymentMethodVnpayQr, PaymentListResponseResultPaymentMethodN26, PaymentListResponseResultPaymentMethodWise, PaymentListResponseResultPaymentMethodPaydoWallet, PaymentListResponseResultPaymentMethodPapara, PaymentListResponseResultPaymentMethodPayoutSepaBatch, PaymentListResponseResultPaymentMethodPayoutNonsepaRequest:
		return true
	}
	return false
}

type PaymentListResponseResultPaymentMethodDetails struct {
	// Card expiration month (for BASIC_CARD payment method only)
	CardExpiryMonth string `json:"cardExpiryMonth"`
	// Card expiration year (for BASIC_CARD payment method only)
	CardExpiryYear string `json:"cardExpiryYear"`
	// Cardholder name (for BASIC_CARD payment method only)
	CardholderName string `json:"cardholderName"`
	// Card issuing country code (for BASIC_CARD payment method only)
	CardIssuingCountryCode string `json:"cardIssuingCountryCode"`
	// Customer account Id in external system or masked card PAN
	CustomerAccountNumber string                                            `json:"customerAccountNumber"`
	JSON                  paymentListResponseResultPaymentMethodDetailsJSON `json:"-"`
}

// paymentListResponseResultPaymentMethodDetailsJSON contains the JSON metadata for
// the struct [PaymentListResponseResultPaymentMethodDetails]
type paymentListResponseResultPaymentMethodDetailsJSON struct {
	CardExpiryMonth        apijson.Field
	CardExpiryYear         apijson.Field
	CardholderName         apijson.Field
	CardIssuingCountryCode apijson.Field
	CustomerAccountNumber  apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PaymentListResponseResultPaymentMethodDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentListResponseResultPaymentMethodDetailsJSON) RawJSON() string {
	return r.raw
}

// Payment Type
type PaymentListResponseResultPaymentType string

const (
	PaymentListResponseResultPaymentTypeDeposit    PaymentListResponseResultPaymentType = "DEPOSIT"
	PaymentListResponseResultPaymentTypeWithdrawal PaymentListResponseResultPaymentType = "WITHDRAWAL"
	PaymentListResponseResultPaymentTypeRefund     PaymentListResponseResultPaymentType = "REFUND"
)

func (r PaymentListResponseResultPaymentType) IsKnown() bool {
	switch r {
	case PaymentListResponseResultPaymentTypeDeposit, PaymentListResponseResultPaymentTypeWithdrawal, PaymentListResponseResultPaymentTypeRefund:
		return true
	}
	return false
}

// Payment State
type PaymentListResponseResultState string

const (
	PaymentListResponseResultStateCheckout  PaymentListResponseResultState = "CHECKOUT"
	PaymentListResponseResultStatePending   PaymentListResponseResultState = "PENDING"
	PaymentListResponseResultStateCancelled PaymentListResponseResultState = "CANCELLED"
	PaymentListResponseResultStateDeclined  PaymentListResponseResultState = "DECLINED"
	PaymentListResponseResultStateCompleted PaymentListResponseResultState = "COMPLETED"
)

func (r PaymentListResponseResultState) IsKnown() bool {
	switch r {
	case PaymentListResponseResultStateCheckout, PaymentListResponseResultStatePending, PaymentListResponseResultStateCancelled, PaymentListResponseResultStateDeclined, PaymentListResponseResultStateCompleted:
		return true
	}
	return false
}

type PaymentNewParams struct {
	// Payment currency
	Currency param.Field[string] `json:"currency,required" format:"ISO 4217 code for FIAT currencies or cryptocurrency symbol"`
	// Payment Type
	PaymentType param.Field[PaymentNewParamsPaymentType] `json:"paymentType,required"`
	// Additional parameters required by some payment providers. Contact support for
	// more information.
	AdditionalParameters param.Field[map[string]string] `json:"additionalParameters"`
	// Payment amount
	Amount param.Field[float64] `json:"amount"`
	// Customer's billing address
	BillingAddress param.Field[PaymentNewParamsBillingAddress] `json:"billingAddress"`
	// You must be PCI DSS compliant to collect card data on your side. If you are not
	// certified, do not add this field to your request and we will collect the data on
	// our page.
	Card     param.Field[PaymentNewParamsCard]     `json:"card"`
	Customer param.Field[PaymentNewParamsCustomer] `json:"customer"`
	// Description of the transaction shown to the Customer. Can be sent outside the
	// system.
	Description param.Field[string] `json:"description"`
	// Id of initial deposit for refunds, Id of initial recurring payment for
	// subsequent payments
	ParentPaymentID param.Field[string] `json:"parentPaymentId"`
	// Payment Method
	PaymentMethod param.Field[PaymentNewParamsPaymentMethod] `json:"paymentMethod"`
	// To continue recurring chain, send a token from a previously initiated recurring
	// payment.
	RecurringToken param.Field[string] `json:"recurringToken"`
	// Reference assigned by Merchant. Will not go outside the system. Will be sent
	// unchanged in the PaymentResponse.
	ReferenceID param.Field[string] `json:"referenceId"`
	// URL to redirect Customer after processing
	ReturnURL param.Field[string] `json:"returnUrl"`
	// Send 'true' if you want this payment to initiate recurring chain. Default is
	// 'false'.
	StartRecurring param.Field[bool] `json:"startRecurring"`
	// Subscription to bill customers at regular intervals. Used only with
	// 'startRecurring=true'.
	Subscription param.Field[PaymentNewParamsSubscription] `json:"subscription"`
	// Url to receive payment status notifications
	WebhookURL param.Field[string] `json:"webhookUrl"`
}

func (r PaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Payment Type
type PaymentNewParamsPaymentType string

const (
	PaymentNewParamsPaymentTypeDeposit    PaymentNewParamsPaymentType = "DEPOSIT"
	PaymentNewParamsPaymentTypeWithdrawal PaymentNewParamsPaymentType = "WITHDRAWAL"
	PaymentNewParamsPaymentTypeRefund     PaymentNewParamsPaymentType = "REFUND"
)

func (r PaymentNewParamsPaymentType) IsKnown() bool {
	switch r {
	case PaymentNewParamsPaymentTypeDeposit, PaymentNewParamsPaymentTypeWithdrawal, PaymentNewParamsPaymentTypeRefund:
		return true
	}
	return false
}

// Customer's billing address
type PaymentNewParamsBillingAddress struct {
	// Line 1 of the address (e.g., Number, street, etc)
	AddressLine1 param.Field[string] `json:"addressLine1"`
	// Line 2 of the address (e.g., Suite, apt)
	AddressLine2 param.Field[string] `json:"addressLine2"`
	// City name
	City param.Field[string] `json:"city"`
	// 2-character IS0-3166-1 country code
	CountryCode param.Field[string] `json:"countryCode"`
	// Postal code
	PostalCode param.Field[string] `json:"postalCode"`
	// State code
	State param.Field[string] `json:"state"`
}

func (r PaymentNewParamsBillingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// You must be PCI DSS compliant to collect card data on your side. If you are not
// certified, do not add this field to your request and we will collect the data on
// our page.
type PaymentNewParamsCard struct {
	// Cardholder's name printed on the card
	CardholderName param.Field[string] `json:"cardholderName"`
	// Card primary account number (PAN). All non-numeric characters will be ignored.
	CardNumber param.Field[string] `json:"cardNumber"`
	// Card security code (CVV2 / CVC2 / CAV2)
	CardSecurityCode param.Field[string] `json:"cardSecurityCode"`
	// Card expiration month, 2 digits
	ExpiryMonth param.Field[string] `json:"expiryMonth"`
	// Card expiration year, 4 digits
	ExpiryYear param.Field[string] `json:"expiryYear"`
}

func (r PaymentNewParamsCard) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsCustomer struct {
	// Customer account name in the provider's system. Used for some types of
	// withdrawals.
	AccountName param.Field[string] `json:"accountName"`
	// Customer account number in the provider's system. Used for some types of
	// withdrawals.
	AccountNumber param.Field[string] `json:"accountNumber"`
	// Customer bank. Used for some types of withdrawals.
	Bank param.Field[string] `json:"bank"`
	// Customer bank branch. Used for some types of withdrawals.
	BankBranch param.Field[string] `json:"bankBranch"`
	// Customer country of citizenship
	CitizenshipCountryCode param.Field[string] `json:"citizenshipCountryCode"`
	DateOfBirth            param.Field[string] `json:"dateOfBirth" format:"ISO 8601 (YYYY-MM-DD)"`
	// Date of the first deposit from the customer
	DateOfFirstDeposit param.Field[string] `json:"dateOfFirstDeposit" format:"ISO 8601 (YYYY-MM-DD)"`
	// How much the customer has deposited, in base currency
	DepositsAmount param.Field[int64] `json:"depositsAmount"`
	// How many times the customer made a deposit
	DepositsCnt param.Field[int64] `json:"depositsCnt"`
	// An identifier for the customer assigned by a government authority
	DocumentNumber param.Field[string] `json:"documentNumber"`
	// Document Type
	DocumentType param.Field[PaymentNewParamsCustomerDocumentType] `json:"documentType"`
	// Email address of the customer
	Email     param.Field[string] `json:"email" format:"email"`
	FirstName param.Field[string] `json:"firstName"`
	// Indicates whether the customer has passed KYC verification
	KYCStatus param.Field[bool]   `json:"kycStatus"`
	LastName  param.Field[string] `json:"lastName"`
	// Customer preferred display language
	Locale param.Field[string] `json:"locale"`
	// Indicates whether the payment instrument (usually the card number) has passed
	// KYC verification
	PaymentInstrumentKYCStatus param.Field[bool] `json:"paymentInstrumentKycStatus"`
	// International phone number of the customer, without the '+'. Use a space as a
	// separator between the dialing country code and local phone number.
	Phone param.Field[string] `json:"phone"`
	// Id of the customer assigned by Merchant
	ReferenceID param.Field[string] `json:"referenceId"`
	// Identify the customer as belonging to a specific group that is used for routing
	RoutingGroup param.Field[string] `json:"routingGroup"`
	// How much the customer has withdrawn, in base currency
	WithdrawalsAmount param.Field[int64] `json:"withdrawalsAmount"`
	// How many times the customer made a withdrawal
	WithdrawalsCnt param.Field[int64] `json:"withdrawalsCnt"`
}

func (r PaymentNewParamsCustomer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Document Type
type PaymentNewParamsCustomerDocumentType string

const (
	PaymentNewParamsCustomerDocumentTypeArCdi  PaymentNewParamsCustomerDocumentType = "AR_CDI"
	PaymentNewParamsCustomerDocumentTypeArCuil PaymentNewParamsCustomerDocumentType = "AR_CUIL"
	PaymentNewParamsCustomerDocumentTypeArCuit PaymentNewParamsCustomerDocumentType = "AR_CUIT"
	PaymentNewParamsCustomerDocumentTypeArDni  PaymentNewParamsCustomerDocumentType = "AR_DNI"
	PaymentNewParamsCustomerDocumentTypeArOtro PaymentNewParamsCustomerDocumentType = "AR_OTRO"
	PaymentNewParamsCustomerDocumentTypeBrCnpj PaymentNewParamsCustomerDocumentType = "BR_CNPJ"
	PaymentNewParamsCustomerDocumentTypeBrCpf  PaymentNewParamsCustomerDocumentType = "BR_CPF"
	PaymentNewParamsCustomerDocumentTypeClOtro PaymentNewParamsCustomerDocumentType = "CL_OTRO"
	PaymentNewParamsCustomerDocumentTypeClRun  PaymentNewParamsCustomerDocumentType = "CL_RUN"
	PaymentNewParamsCustomerDocumentTypeClRut  PaymentNewParamsCustomerDocumentType = "CL_RUT"
	PaymentNewParamsCustomerDocumentTypeCoCc   PaymentNewParamsCustomerDocumentType = "CO_CC"
	PaymentNewParamsCustomerDocumentTypeCoCe   PaymentNewParamsCustomerDocumentType = "CO_CE"
	PaymentNewParamsCustomerDocumentTypeCoDl   PaymentNewParamsCustomerDocumentType = "CO_DL"
	PaymentNewParamsCustomerDocumentTypeCoDni  PaymentNewParamsCustomerDocumentType = "CO_DNI"
	PaymentNewParamsCustomerDocumentTypeCoNe   PaymentNewParamsCustomerDocumentType = "CO_NE"
	PaymentNewParamsCustomerDocumentTypeCoNit  PaymentNewParamsCustomerDocumentType = "CO_NIT"
	PaymentNewParamsCustomerDocumentTypeCoPp   PaymentNewParamsCustomerDocumentType = "CO_PP"
	PaymentNewParamsCustomerDocumentTypeCoSS   PaymentNewParamsCustomerDocumentType = "CO_SS"
	PaymentNewParamsCustomerDocumentTypeCoTi   PaymentNewParamsCustomerDocumentType = "CO_TI"
	PaymentNewParamsCustomerDocumentTypeCrCdi  PaymentNewParamsCustomerDocumentType = "CR_CDI"
	PaymentNewParamsCustomerDocumentTypeEcDni  PaymentNewParamsCustomerDocumentType = "EC_DNI"
	PaymentNewParamsCustomerDocumentTypeEcPp   PaymentNewParamsCustomerDocumentType = "EC_PP"
	PaymentNewParamsCustomerDocumentTypeEcRuc  PaymentNewParamsCustomerDocumentType = "EC_RUC"
	PaymentNewParamsCustomerDocumentTypeGtCui  PaymentNewParamsCustomerDocumentType = "GT_CUI"
	PaymentNewParamsCustomerDocumentTypeGtDpi  PaymentNewParamsCustomerDocumentType = "GT_DPI"
	PaymentNewParamsCustomerDocumentTypeGtNit  PaymentNewParamsCustomerDocumentType = "GT_NIT"
	PaymentNewParamsCustomerDocumentTypeMxCurp PaymentNewParamsCustomerDocumentType = "MX_CURP"
	PaymentNewParamsCustomerDocumentTypeMxIfe  PaymentNewParamsCustomerDocumentType = "MX_IFE"
	PaymentNewParamsCustomerDocumentTypeMxPp   PaymentNewParamsCustomerDocumentType = "MX_PP"
	PaymentNewParamsCustomerDocumentTypeMxRfc  PaymentNewParamsCustomerDocumentType = "MX_RFC"
	PaymentNewParamsCustomerDocumentTypePaCip  PaymentNewParamsCustomerDocumentType = "PA_CIP"
	PaymentNewParamsCustomerDocumentTypePeCe   PaymentNewParamsCustomerDocumentType = "PE_CE"
	PaymentNewParamsCustomerDocumentTypePeDni  PaymentNewParamsCustomerDocumentType = "PE_DNI"
	PaymentNewParamsCustomerDocumentTypePeOtro PaymentNewParamsCustomerDocumentType = "PE_OTRO"
	PaymentNewParamsCustomerDocumentTypePePp   PaymentNewParamsCustomerDocumentType = "PE_PP"
	PaymentNewParamsCustomerDocumentTypePeRuc  PaymentNewParamsCustomerDocumentType = "PE_RUC"
)

func (r PaymentNewParamsCustomerDocumentType) IsKnown() bool {
	switch r {
	case PaymentNewParamsCustomerDocumentTypeArCdi, PaymentNewParamsCustomerDocumentTypeArCuil, PaymentNewParamsCustomerDocumentTypeArCuit, PaymentNewParamsCustomerDocumentTypeArDni, PaymentNewParamsCustomerDocumentTypeArOtro, PaymentNewParamsCustomerDocumentTypeBrCnpj, PaymentNewParamsCustomerDocumentTypeBrCpf, PaymentNewParamsCustomerDocumentTypeClOtro, PaymentNewParamsCustomerDocumentTypeClRun, PaymentNewParamsCustomerDocumentTypeClRut, PaymentNewParamsCustomerDocumentTypeCoCc, PaymentNewParamsCustomerDocumentTypeCoCe, PaymentNewParamsCustomerDocumentTypeCoDl, PaymentNewParamsCustomerDocumentTypeCoDni, PaymentNewParamsCustomerDocumentTypeCoNe, PaymentNewParamsCustomerDocumentTypeCoNit, PaymentNewParamsCustomerDocumentTypeCoPp, PaymentNewParamsCustomerDocumentTypeCoSS, PaymentNewParamsCustomerDocumentTypeCoTi, PaymentNewParamsCustomerDocumentTypeCrCdi, PaymentNewParamsCustomerDocumentTypeEcDni, PaymentNewParamsCustomerDocumentTypeEcPp, PaymentNewParamsCustomerDocumentTypeEcRuc, PaymentNewParamsCustomerDocumentTypeGtCui, PaymentNewParamsCustomerDocumentTypeGtDpi, PaymentNewParamsCustomerDocumentTypeGtNit, PaymentNewParamsCustomerDocumentTypeMxCurp, PaymentNewParamsCustomerDocumentTypeMxIfe, PaymentNewParamsCustomerDocumentTypeMxPp, PaymentNewParamsCustomerDocumentTypeMxRfc, PaymentNewParamsCustomerDocumentTypePaCip, PaymentNewParamsCustomerDocumentTypePeCe, PaymentNewParamsCustomerDocumentTypePeDni, PaymentNewParamsCustomerDocumentTypePeOtro, PaymentNewParamsCustomerDocumentTypePePp, PaymentNewParamsCustomerDocumentTypePeRuc:
		return true
	}
	return false
}

// Payment Method
type PaymentNewParamsPaymentMethod string

const (
	PaymentNewParamsPaymentMethodBasicCard            PaymentNewParamsPaymentMethod = "BASIC_CARD"
	PaymentNewParamsPaymentMethodCrypto               PaymentNewParamsPaymentMethod = "CRYPTO"
	PaymentNewParamsPaymentMethodFlexepin             PaymentNewParamsPaymentMethod = "FLEXEPIN"
	PaymentNewParamsPaymentMethodMacropay             PaymentNewParamsPaymentMethod = "MACROPAY"
	PaymentNewParamsPaymentMethodSkrill               PaymentNewParamsPaymentMethod = "SKRILL"
	PaymentNewParamsPaymentMethodPayretailers         PaymentNewParamsPaymentMethod = "PAYRETAILERS"
	PaymentNewParamsPaymentMethodLocalpayment         PaymentNewParamsPaymentMethod = "LOCALPAYMENT"
	PaymentNewParamsPaymentMethodMonnet               PaymentNewParamsPaymentMethod = "MONNET"
	PaymentNewParamsPaymentMethodPaypal               PaymentNewParamsPaymentMethod = "PAYPAL"
	PaymentNewParamsPaymentMethodNeteller             PaymentNewParamsPaymentMethod = "NETELLER"
	PaymentNewParamsPaymentMethodTrustpayments        PaymentNewParamsPaymentMethod = "TRUSTPAYMENTS"
	PaymentNewParamsPaymentMethodPaymaxis             PaymentNewParamsPaymentMethod = "PAYMAXIS"
	PaymentNewParamsPaymentMethodGate8Transact        PaymentNewParamsPaymentMethod = "GATE8TRANSACT"
	PaymentNewParamsPaymentMethodTink                 PaymentNewParamsPaymentMethod = "TINK"
	PaymentNewParamsPaymentMethodB2Binpay             PaymentNewParamsPaymentMethod = "B2BINPAY"
	PaymentNewParamsPaymentMethodClick                PaymentNewParamsPaymentMethod = "CLICK"
	PaymentNewParamsPaymentMethodMonetix              PaymentNewParamsPaymentMethod = "MONETIX"
	PaymentNewParamsPaymentMethodPerfectmoney         PaymentNewParamsPaymentMethod = "PERFECTMONEY"
	PaymentNewParamsPaymentMethodVolt                 PaymentNewParamsPaymentMethod = "VOLT"
	PaymentNewParamsPaymentMethodKesspay              PaymentNewParamsPaymentMethod = "KESSPAY"
	PaymentNewParamsPaymentMethodBillline             PaymentNewParamsPaymentMethod = "BILLLINE"
	PaymentNewParamsPaymentMethodNgenius              PaymentNewParamsPaymentMethod = "NGENIUS"
	PaymentNewParamsPaymentMethodAstropay             PaymentNewParamsPaymentMethod = "ASTROPAY"
	PaymentNewParamsPaymentMethodAlycepay             PaymentNewParamsPaymentMethod = "ALYCEPAY"
	PaymentNewParamsPaymentMethodPix                  PaymentNewParamsPaymentMethod = "PIX"
	PaymentNewParamsPaymentMethodBoleto               PaymentNewParamsPaymentMethod = "BOLETO"
	PaymentNewParamsPaymentMethodUpi                  PaymentNewParamsPaymentMethod = "UPI"
	PaymentNewParamsPaymentMethodPaytm                PaymentNewParamsPaymentMethod = "PAYTM"
	PaymentNewParamsPaymentMethodNetbanking           PaymentNewParamsPaymentMethod = "NETBANKING"
	PaymentNewParamsPaymentMethodFinrax               PaymentNewParamsPaymentMethod = "FINRAX"
	PaymentNewParamsPaymentMethodSpoynt               PaymentNewParamsPaymentMethod = "SPOYNT"
	PaymentNewParamsPaymentMethodXinpay               PaymentNewParamsPaymentMethod = "XINPAY"
	PaymentNewParamsPaymentMethodOmnimatrix           PaymentNewParamsPaymentMethod = "OMNIMATRIX"
	PaymentNewParamsPaymentMethodDpopay               PaymentNewParamsPaymentMethod = "DPOPAY"
	PaymentNewParamsPaymentMethodExternalHpp          PaymentNewParamsPaymentMethod = "EXTERNAL_HPP"
	PaymentNewParamsPaymentMethodXanpay               PaymentNewParamsPaymentMethod = "XANPAY"
	PaymentNewParamsPaymentMethodInrpay               PaymentNewParamsPaymentMethod = "INRPAY"
	PaymentNewParamsPaymentMethodAri10                PaymentNewParamsPaymentMethod = "ARI10"
	PaymentNewParamsPaymentMethodSofort               PaymentNewParamsPaymentMethod = "SOFORT"
	PaymentNewParamsPaymentMethodGiropay              PaymentNewParamsPaymentMethod = "GIROPAY"
	PaymentNewParamsPaymentMethodPaysafecard          PaymentNewParamsPaymentMethod = "PAYSAFECARD"
	PaymentNewParamsPaymentMethodPaysafecash          PaymentNewParamsPaymentMethod = "PAYSAFECASH"
	PaymentNewParamsPaymentMethodOpenBanking          PaymentNewParamsPaymentMethod = "OPEN_BANKING"
	PaymentNewParamsPaymentMethodKlarna               PaymentNewParamsPaymentMethod = "KLARNA"
	PaymentNewParamsPaymentMethodSpei                 PaymentNewParamsPaymentMethod = "SPEI"
	PaymentNewParamsPaymentMethodPaycash              PaymentNewParamsPaymentMethod = "PAYCASH"
	PaymentNewParamsPaymentMethodRapipago             PaymentNewParamsPaymentMethod = "RAPIPAGO"
	PaymentNewParamsPaymentMethodPagofacil            PaymentNewParamsPaymentMethod = "PAGOFACIL"
	PaymentNewParamsPaymentMethodRapidtransfer        PaymentNewParamsPaymentMethod = "RAPIDTRANSFER"
	PaymentNewParamsPaymentMethodMobileMoney          PaymentNewParamsPaymentMethod = "MOBILE_MONEY"
	PaymentNewParamsPaymentMethodInterac              PaymentNewParamsPaymentMethod = "INTERAC"
	PaymentNewParamsPaymentMethodInteracEto           PaymentNewParamsPaymentMethod = "INTERAC_ETO"
	PaymentNewParamsPaymentMethodInteracRto           PaymentNewParamsPaymentMethod = "INTERAC_RTO"
	PaymentNewParamsPaymentMethodInteracACH           PaymentNewParamsPaymentMethod = "INTERAC_ACH"
	PaymentNewParamsPaymentMethodPicpay               PaymentNewParamsPaymentMethod = "PICPAY"
	PaymentNewParamsPaymentMethodMollie               PaymentNewParamsPaymentMethod = "MOLLIE"
	PaymentNewParamsPaymentMethodTed                  PaymentNewParamsPaymentMethod = "TED"
	PaymentNewParamsPaymentMethodZipay                PaymentNewParamsPaymentMethod = "ZIPAY"
	PaymentNewParamsPaymentMethodPse                  PaymentNewParamsPaymentMethod = "PSE"
	PaymentNewParamsPaymentMethodEfecty               PaymentNewParamsPaymentMethod = "EFECTY"
	PaymentNewParamsPaymentMethodBanktransfer         PaymentNewParamsPaymentMethod = "BANKTRANSFER"
	PaymentNewParamsPaymentMethodPec                  PaymentNewParamsPaymentMethod = "PEC"
	PaymentNewParamsPaymentMethodOxxo                 PaymentNewParamsPaymentMethod = "OXXO"
	PaymentNewParamsPaymentMethodWebpay               PaymentNewParamsPaymentMethod = "WEBPAY"
	PaymentNewParamsPaymentMethodPagoefectivo         PaymentNewParamsPaymentMethod = "PAGOEFECTIVO"
	PaymentNewParamsPaymentMethodMifinity             PaymentNewParamsPaymentMethod = "MIFINITY"
	PaymentNewParamsPaymentMethodPayport              PaymentNewParamsPaymentMethod = "PAYPORT"
	PaymentNewParamsPaymentMethodJetoncash            PaymentNewParamsPaymentMethod = "JETONCASH"
	PaymentNewParamsPaymentMethodJetonwallet          PaymentNewParamsPaymentMethod = "JETONWALLET"
	PaymentNewParamsPaymentMethodNoda                 PaymentNewParamsPaymentMethod = "NODA"
	PaymentNewParamsPaymentMethodNodaRevolut          PaymentNewParamsPaymentMethod = "NODA_REVOLUT"
	PaymentNewParamsPaymentMethodAlfakit              PaymentNewParamsPaymentMethod = "ALFAKIT"
	PaymentNewParamsPaymentMethodPayfun               PaymentNewParamsPaymentMethod = "PAYFUN"
	PaymentNewParamsPaymentMethodEmanat               PaymentNewParamsPaymentMethod = "EMANAT"
	PaymentNewParamsPaymentMethodM10                  PaymentNewParamsPaymentMethod = "M10"
	PaymentNewParamsPaymentMethodRubpay               PaymentNewParamsPaymentMethod = "RUBPAY"
	PaymentNewParamsPaymentMethodMonerchy             PaymentNewParamsPaymentMethod = "MONERCHY"
	PaymentNewParamsPaymentMethodMuchbetter           PaymentNewParamsPaymentMethod = "MUCHBETTER"
	PaymentNewParamsPaymentMethodYapily               PaymentNewParamsPaymentMethod = "YAPILY"
	PaymentNewParamsPaymentMethodInai                 PaymentNewParamsPaymentMethod = "INAI"
	PaymentNewParamsPaymentMethodImps                 PaymentNewParamsPaymentMethod = "IMPS"
	PaymentNewParamsPaymentMethodRtgs                 PaymentNewParamsPaymentMethod = "RTGS"
	PaymentNewParamsPaymentMethodPayid                PaymentNewParamsPaymentMethod = "PAYID"
	PaymentNewParamsPaymentMethodZotapay              PaymentNewParamsPaymentMethod = "ZOTAPAY"
	PaymentNewParamsPaymentMethodSbp                  PaymentNewParamsPaymentMethod = "SBP"
	PaymentNewParamsPaymentMethodP2PCard              PaymentNewParamsPaymentMethod = "P2P_CARD"
	PaymentNewParamsPaymentMethodP2PIban              PaymentNewParamsPaymentMethod = "P2P_IBAN"
	PaymentNewParamsPaymentMethodP2PSbp               PaymentNewParamsPaymentMethod = "P2P_SBP"
	PaymentNewParamsPaymentMethodP2PMobile            PaymentNewParamsPaymentMethod = "P2P_MOBILE"
	PaymentNewParamsPaymentMethodPush                 PaymentNewParamsPaymentMethod = "PUSH"
	PaymentNewParamsPaymentMethodGateiq               PaymentNewParamsPaymentMethod = "GATEIQ"
	PaymentNewParamsPaymentMethodViettel              PaymentNewParamsPaymentMethod = "VIETTEL"
	PaymentNewParamsPaymentMethodZalo                 PaymentNewParamsPaymentMethod = "ZALO"
	PaymentNewParamsPaymentMethodQr                   PaymentNewParamsPaymentMethod = "QR"
	PaymentNewParamsPaymentMethodCup                  PaymentNewParamsPaymentMethod = "CUP"
	PaymentNewParamsPaymentMethodCodi                 PaymentNewParamsPaymentMethod = "CODI"
	PaymentNewParamsPaymentMethodPay2Play             PaymentNewParamsPaymentMethod = "PAY2PLAY"
	PaymentNewParamsPaymentMethodBkash                PaymentNewParamsPaymentMethod = "BKASH"
	PaymentNewParamsPaymentMethodNagad                PaymentNewParamsPaymentMethod = "NAGAD"
	PaymentNewParamsPaymentMethodRocket               PaymentNewParamsPaymentMethod = "ROCKET"
	PaymentNewParamsPaymentMethodVirtualAccount       PaymentNewParamsPaymentMethod = "VIRTUAL_ACCOUNT"
	PaymentNewParamsPaymentMethodMultibanco           PaymentNewParamsPaymentMethod = "MULTIBANCO"
	PaymentNewParamsPaymentMethodBlik                 PaymentNewParamsPaymentMethod = "BLIK"
	PaymentNewParamsPaymentMethodMbway                PaymentNewParamsPaymentMethod = "MBWAY"
	PaymentNewParamsPaymentMethodP24                  PaymentNewParamsPaymentMethod = "P24"
	PaymentNewParamsPaymentMethodMistercash           PaymentNewParamsPaymentMethod = "MISTERCASH"
	PaymentNewParamsPaymentMethodMach                 PaymentNewParamsPaymentMethod = "MACH"
	PaymentNewParamsPaymentMethodKhipu                PaymentNewParamsPaymentMethod = "KHIPU"
	PaymentNewParamsPaymentMethodNeft                 PaymentNewParamsPaymentMethod = "NEFT"
	PaymentNewParamsPaymentMethodSticpay              PaymentNewParamsPaymentMethod = "STICPAY"
	PaymentNewParamsPaymentMethodSberpay              PaymentNewParamsPaymentMethod = "SBERPAY"
	PaymentNewParamsPaymentMethodMobileCommerce       PaymentNewParamsPaymentMethod = "MOBILE_COMMERCE"
	PaymentNewParamsPaymentMethodBinancePay           PaymentNewParamsPaymentMethod = "BINANCE_PAY"
	PaymentNewParamsPaymentMethodMpay                 PaymentNewParamsPaymentMethod = "MPAY"
	PaymentNewParamsPaymentMethodChek                 PaymentNewParamsPaymentMethod = "CHEK"
	PaymentNewParamsPaymentMethodKlapEfectivo         PaymentNewParamsPaymentMethod = "KLAP_EFECTIVO"
	PaymentNewParamsPaymentMethodKlapTransferencia    PaymentNewParamsPaymentMethod = "KLAP_TRANSFERENCIA"
	PaymentNewParamsPaymentMethodPago46               PaymentNewParamsPaymentMethod = "PAGO46"
	PaymentNewParamsPaymentMethodHites                PaymentNewParamsPaymentMethod = "HITES"
	PaymentNewParamsPaymentMethodServifacil           PaymentNewParamsPaymentMethod = "SERVIFACIL"
	PaymentNewParamsPaymentMethodOpenpayd             PaymentNewParamsPaymentMethod = "OPENPAYD"
	PaymentNewParamsPaymentMethodFawry                PaymentNewParamsPaymentMethod = "FAWRY"
	PaymentNewParamsPaymentMethodEps                  PaymentNewParamsPaymentMethod = "EPS"
	PaymentNewParamsPaymentMethodIdeal                PaymentNewParamsPaymentMethod = "IDEAL"
	PaymentNewParamsPaymentMethodTrustly              PaymentNewParamsPaymentMethod = "TRUSTLY"
	PaymentNewParamsPaymentMethodUssd                 PaymentNewParamsPaymentMethod = "USSD"
	PaymentNewParamsPaymentMethodMpesa                PaymentNewParamsPaymentMethod = "MPESA"
	PaymentNewParamsPaymentMethodEnaira               PaymentNewParamsPaymentMethod = "ENAIRA"
	PaymentNewParamsPaymentMethodOnevoucher           PaymentNewParamsPaymentMethod = "ONEVOUCHER"
	PaymentNewParamsPaymentMethodBancontact           PaymentNewParamsPaymentMethod = "BANCONTACT"
	PaymentNewParamsPaymentMethodSwish                PaymentNewParamsPaymentMethod = "SWISH"
	PaymentNewParamsPaymentMethodEft                  PaymentNewParamsPaymentMethod = "EFT"
	PaymentNewParamsPaymentMethodGcash                PaymentNewParamsPaymentMethod = "GCASH"
	PaymentNewParamsPaymentMethodPaymaya              PaymentNewParamsPaymentMethod = "PAYMAYA"
	PaymentNewParamsPaymentMethodPagoMovil            PaymentNewParamsPaymentMethod = "PAGO_MOVIL"
	PaymentNewParamsPaymentMethodPagoMovilInst        PaymentNewParamsPaymentMethod = "PAGO_MOVIL_INST"
	PaymentNewParamsPaymentMethodBiopago              PaymentNewParamsPaymentMethod = "BIOPAGO"
	PaymentNewParamsPaymentMethodCash                 PaymentNewParamsPaymentMethod = "CASH"
	PaymentNewParamsPaymentMethodVoucherry            PaymentNewParamsPaymentMethod = "VOUCHERRY"
	PaymentNewParamsPaymentMethodApplepay             PaymentNewParamsPaymentMethod = "APPLEPAY"
	PaymentNewParamsPaymentMethodGooglepay            PaymentNewParamsPaymentMethod = "GOOGLEPAY"
	PaymentNewParamsPaymentMethodBrite                PaymentNewParamsPaymentMethod = "BRITE"
	PaymentNewParamsPaymentMethodVouchstar            PaymentNewParamsPaymentMethod = "VOUCHSTAR"
	PaymentNewParamsPaymentMethodRevolut              PaymentNewParamsPaymentMethod = "REVOLUT"
	PaymentNewParamsPaymentMethodOnlineBanking        PaymentNewParamsPaymentMethod = "ONLINE_BANKING"
	PaymentNewParamsPaymentMethodPromptpay            PaymentNewParamsPaymentMethod = "PROMPTPAY"
	PaymentNewParamsPaymentMethodTruemoney            PaymentNewParamsPaymentMethod = "TRUEMONEY"
	PaymentNewParamsPaymentMethodMomopayVn            PaymentNewParamsPaymentMethod = "MOMOPAY_VN"
	PaymentNewParamsPaymentMethodMomopayRw            PaymentNewParamsPaymentMethod = "MOMOPAY_RW"
	PaymentNewParamsPaymentMethodVnpayQr              PaymentNewParamsPaymentMethod = "VNPAY_QR"
	PaymentNewParamsPaymentMethodN26                  PaymentNewParamsPaymentMethod = "N26"
	PaymentNewParamsPaymentMethodWise                 PaymentNewParamsPaymentMethod = "WISE"
	PaymentNewParamsPaymentMethodPaydoWallet          PaymentNewParamsPaymentMethod = "PAYDO_WALLET"
	PaymentNewParamsPaymentMethodPapara               PaymentNewParamsPaymentMethod = "PAPARA"
	PaymentNewParamsPaymentMethodPayoutSepaBatch      PaymentNewParamsPaymentMethod = "PAYOUT_SEPA_BATCH"
	PaymentNewParamsPaymentMethodPayoutNonsepaRequest PaymentNewParamsPaymentMethod = "PAYOUT_NONSEPA_REQUEST"
)

func (r PaymentNewParamsPaymentMethod) IsKnown() bool {
	switch r {
	case PaymentNewParamsPaymentMethodBasicCard, PaymentNewParamsPaymentMethodCrypto, PaymentNewParamsPaymentMethodFlexepin, PaymentNewParamsPaymentMethodMacropay, PaymentNewParamsPaymentMethodSkrill, PaymentNewParamsPaymentMethodPayretailers, PaymentNewParamsPaymentMethodLocalpayment, PaymentNewParamsPaymentMethodMonnet, PaymentNewParamsPaymentMethodPaypal, PaymentNewParamsPaymentMethodNeteller, PaymentNewParamsPaymentMethodTrustpayments, PaymentNewParamsPaymentMethodPaymaxis, PaymentNewParamsPaymentMethodGate8Transact, PaymentNewParamsPaymentMethodTink, PaymentNewParamsPaymentMethodB2Binpay, PaymentNewParamsPaymentMethodClick, PaymentNewParamsPaymentMethodMonetix, PaymentNewParamsPaymentMethodPerfectmoney, PaymentNewParamsPaymentMethodVolt, PaymentNewParamsPaymentMethodKesspay, PaymentNewParamsPaymentMethodBillline, PaymentNewParamsPaymentMethodNgenius, PaymentNewParamsPaymentMethodAstropay, PaymentNewParamsPaymentMethodAlycepay, PaymentNewParamsPaymentMethodPix, PaymentNewParamsPaymentMethodBoleto, PaymentNewParamsPaymentMethodUpi, PaymentNewParamsPaymentMethodPaytm, PaymentNewParamsPaymentMethodNetbanking, PaymentNewParamsPaymentMethodFinrax, PaymentNewParamsPaymentMethodSpoynt, PaymentNewParamsPaymentMethodXinpay, PaymentNewParamsPaymentMethodOmnimatrix, PaymentNewParamsPaymentMethodDpopay, PaymentNewParamsPaymentMethodExternalHpp, PaymentNewParamsPaymentMethodXanpay, PaymentNewParamsPaymentMethodInrpay, PaymentNewParamsPaymentMethodAri10, PaymentNewParamsPaymentMethodSofort, PaymentNewParamsPaymentMethodGiropay, PaymentNewParamsPaymentMethodPaysafecard, PaymentNewParamsPaymentMethodPaysafecash, PaymentNewParamsPaymentMethodOpenBanking, PaymentNewParamsPaymentMethodKlarna, PaymentNewParamsPaymentMethodSpei, PaymentNewParamsPaymentMethodPaycash, PaymentNewParamsPaymentMethodRapipago, PaymentNewParamsPaymentMethodPagofacil, PaymentNewParamsPaymentMethodRapidtransfer, PaymentNewParamsPaymentMethodMobileMoney, PaymentNewParamsPaymentMethodInterac, PaymentNewParamsPaymentMethodInteracEto, PaymentNewParamsPaymentMethodInteracRto, PaymentNewParamsPaymentMethodInteracACH, PaymentNewParamsPaymentMethodPicpay, PaymentNewParamsPaymentMethodMollie, PaymentNewParamsPaymentMethodTed, PaymentNewParamsPaymentMethodZipay, PaymentNewParamsPaymentMethodPse, PaymentNewParamsPaymentMethodEfecty, PaymentNewParamsPaymentMethodBanktransfer, PaymentNewParamsPaymentMethodPec, PaymentNewParamsPaymentMethodOxxo, PaymentNewParamsPaymentMethodWebpay, PaymentNewParamsPaymentMethodPagoefectivo, PaymentNewParamsPaymentMethodMifinity, PaymentNewParamsPaymentMethodPayport, PaymentNewParamsPaymentMethodJetoncash, PaymentNewParamsPaymentMethodJetonwallet, PaymentNewParamsPaymentMethodNoda, PaymentNewParamsPaymentMethodNodaRevolut, PaymentNewParamsPaymentMethodAlfakit, PaymentNewParamsPaymentMethodPayfun, PaymentNewParamsPaymentMethodEmanat, PaymentNewParamsPaymentMethodM10, PaymentNewParamsPaymentMethodRubpay, PaymentNewParamsPaymentMethodMonerchy, PaymentNewParamsPaymentMethodMuchbetter, PaymentNewParamsPaymentMethodYapily, PaymentNewParamsPaymentMethodInai, PaymentNewParamsPaymentMethodImps, PaymentNewParamsPaymentMethodRtgs, PaymentNewParamsPaymentMethodPayid, PaymentNewParamsPaymentMethodZotapay, PaymentNewParamsPaymentMethodSbp, PaymentNewParamsPaymentMethodP2PCard, PaymentNewParamsPaymentMethodP2PIban, PaymentNewParamsPaymentMethodP2PSbp, PaymentNewParamsPaymentMethodP2PMobile, PaymentNewParamsPaymentMethodPush, PaymentNewParamsPaymentMethodGateiq, PaymentNewParamsPaymentMethodViettel, PaymentNewParamsPaymentMethodZalo, PaymentNewParamsPaymentMethodQr, PaymentNewParamsPaymentMethodCup, PaymentNewParamsPaymentMethodCodi, PaymentNewParamsPaymentMethodPay2Play, PaymentNewParamsPaymentMethodBkash, PaymentNewParamsPaymentMethodNagad, PaymentNewParamsPaymentMethodRocket, PaymentNewParamsPaymentMethodVirtualAccount, PaymentNewParamsPaymentMethodMultibanco, PaymentNewParamsPaymentMethodBlik, PaymentNewParamsPaymentMethodMbway, PaymentNewParamsPaymentMethodP24, PaymentNewParamsPaymentMethodMistercash, PaymentNewParamsPaymentMethodMach, PaymentNewParamsPaymentMethodKhipu, PaymentNewParamsPaymentMethodNeft, PaymentNewParamsPaymentMethodSticpay, PaymentNewParamsPaymentMethodSberpay, PaymentNewParamsPaymentMethodMobileCommerce, PaymentNewParamsPaymentMethodBinancePay, PaymentNewParamsPaymentMethodMpay, PaymentNewParamsPaymentMethodChek, PaymentNewParamsPaymentMethodKlapEfectivo, PaymentNewParamsPaymentMethodKlapTransferencia, PaymentNewParamsPaymentMethodPago46, PaymentNewParamsPaymentMethodHites, PaymentNewParamsPaymentMethodServifacil, PaymentNewParamsPaymentMethodOpenpayd, PaymentNewParamsPaymentMethodFawry, PaymentNewParamsPaymentMethodEps, PaymentNewParamsPaymentMethodIdeal, PaymentNewParamsPaymentMethodTrustly, PaymentNewParamsPaymentMethodUssd, PaymentNewParamsPaymentMethodMpesa, PaymentNewParamsPaymentMethodEnaira, PaymentNewParamsPaymentMethodOnevoucher, PaymentNewParamsPaymentMethodBancontact, PaymentNewParamsPaymentMethodSwish, PaymentNewParamsPaymentMethodEft, PaymentNewParamsPaymentMethodGcash, PaymentNewParamsPaymentMethodPaymaya, PaymentNewParamsPaymentMethodPagoMovil, PaymentNewParamsPaymentMethodPagoMovilInst, PaymentNewParamsPaymentMethodBiopago, PaymentNewParamsPaymentMethodCash, PaymentNewParamsPaymentMethodVoucherry, PaymentNewParamsPaymentMethodApplepay, PaymentNewParamsPaymentMethodGooglepay, PaymentNewParamsPaymentMethodBrite, PaymentNewParamsPaymentMethodVouchstar, PaymentNewParamsPaymentMethodRevolut, PaymentNewParamsPaymentMethodOnlineBanking, PaymentNewParamsPaymentMethodPromptpay, PaymentNewParamsPaymentMethodTruemoney, PaymentNewParamsPaymentMethodMomopayVn, PaymentNewParamsPaymentMethodMomopayRw, PaymentNewParamsPaymentMethodVnpayQr, PaymentNewParamsPaymentMethodN26, PaymentNewParamsPaymentMethodWise, PaymentNewParamsPaymentMethodPaydoWallet, PaymentNewParamsPaymentMethodPapara, PaymentNewParamsPaymentMethodPayoutSepaBatch, PaymentNewParamsPaymentMethodPayoutNonsepaRequest:
		return true
	}
	return false
}

// Subscription to bill customers at regular intervals. Used only with
// 'startRecurring=true'.
type PaymentNewParamsSubscription struct {
	// The number of intervals after which a subscriber is billed. For example, if the
	// frequencyUnit is DAY with an frequency of 2, the subscription is billed once
	// every two days.
	Frequency param.Field[int64] `json:"frequency,required"`
	// The amount to be used for subsequent payments. If not specified, the amount of
	// the original payment is used.
	Amount param.Field[float64] `json:"amount"`
	// Description for subsequent recurring payments
	Description param.Field[string] `json:"description"`
	// The interval at which the subscription is billed. Use 'MINUTE' for testing
	// purposes only.
	FrequencyUnit param.Field[PaymentNewParamsSubscriptionFrequencyUnit] `json:"frequencyUnit"`
	// Required number of subsequent recurring payments. Unlimited if value is not
	// specified.
	NumberOfCycles param.Field[int64] `json:"numberOfCycles"`
	// Retry strategy for subscription. If not specified, the subscription is canceled
	// after the first failed payment attempt.
	RetryStrategy param.Field[PaymentNewParamsSubscriptionRetryStrategy] `json:"retryStrategy"`
	// Date and time of the 1st cycle. if not specified, then calculated as
	// (initialDeposit.createTime + frequency\*frequencyUnit).
	StartTime param.Field[string] `json:"startTime" format:"ISO 8601 (YYYY-MM-DD'T'HH24:MI:SS)"`
}

func (r PaymentNewParamsSubscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The interval at which the subscription is billed. Use 'MINUTE' for testing
// purposes only.
type PaymentNewParamsSubscriptionFrequencyUnit string

const (
	PaymentNewParamsSubscriptionFrequencyUnitMinute PaymentNewParamsSubscriptionFrequencyUnit = "MINUTE"
	PaymentNewParamsSubscriptionFrequencyUnitDay    PaymentNewParamsSubscriptionFrequencyUnit = "DAY"
	PaymentNewParamsSubscriptionFrequencyUnitWeek   PaymentNewParamsSubscriptionFrequencyUnit = "WEEK"
	PaymentNewParamsSubscriptionFrequencyUnitMonth  PaymentNewParamsSubscriptionFrequencyUnit = "MONTH"
)

func (r PaymentNewParamsSubscriptionFrequencyUnit) IsKnown() bool {
	switch r {
	case PaymentNewParamsSubscriptionFrequencyUnitMinute, PaymentNewParamsSubscriptionFrequencyUnitDay, PaymentNewParamsSubscriptionFrequencyUnitWeek, PaymentNewParamsSubscriptionFrequencyUnitMonth:
		return true
	}
	return false
}

// Retry strategy for subscription. If not specified, the subscription is canceled
// after the first failed payment attempt.
type PaymentNewParamsSubscriptionRetryStrategy struct {
	// The number of intervals after which the system will retry the payment after an
	// unsuccessful attempt
	Frequency param.Field[int64] `json:"frequency,required"`
	// Required number of retries
	NumberOfCycles param.Field[int64] `json:"numberOfCycles,required"`
	// If specified, the nth element contains the percentage of the initial amount that
	// will be charged for the nth retry
	AmountAdjustments param.Field[[]int64] `json:"amountAdjustments"`
	// The interval at which the subscription is retried. Use 'MINUTE' for testing
	// purposes only.
	FrequencyUnit param.Field[PaymentNewParamsSubscriptionRetryStrategyFrequencyUnit] `json:"frequencyUnit"`
}

func (r PaymentNewParamsSubscriptionRetryStrategy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The interval at which the subscription is retried. Use 'MINUTE' for testing
// purposes only.
type PaymentNewParamsSubscriptionRetryStrategyFrequencyUnit string

const (
	PaymentNewParamsSubscriptionRetryStrategyFrequencyUnitMinute PaymentNewParamsSubscriptionRetryStrategyFrequencyUnit = "MINUTE"
	PaymentNewParamsSubscriptionRetryStrategyFrequencyUnitDay    PaymentNewParamsSubscriptionRetryStrategyFrequencyUnit = "DAY"
	PaymentNewParamsSubscriptionRetryStrategyFrequencyUnitWeek   PaymentNewParamsSubscriptionRetryStrategyFrequencyUnit = "WEEK"
	PaymentNewParamsSubscriptionRetryStrategyFrequencyUnitMonth  PaymentNewParamsSubscriptionRetryStrategyFrequencyUnit = "MONTH"
)

func (r PaymentNewParamsSubscriptionRetryStrategyFrequencyUnit) IsKnown() bool {
	switch r {
	case PaymentNewParamsSubscriptionRetryStrategyFrequencyUnitMinute, PaymentNewParamsSubscriptionRetryStrategyFrequencyUnitDay, PaymentNewParamsSubscriptionRetryStrategyFrequencyUnitWeek, PaymentNewParamsSubscriptionRetryStrategyFrequencyUnitMonth:
		return true
	}
	return false
}

type PaymentListParams struct {
	Created param.Field[PaymentListParamsCreated] `query:"created"`
	// The numbers of items to return. Default is 50.
	Limit param.Field[int64] `query:"limit"`
	// The number of items to skip before starting to collect the result set. Default
	// is 0.
	Offset  param.Field[int64]                    `query:"offset"`
	Updated param.Field[PaymentListParamsUpdated] `query:"updated"`
}

// URLQuery serializes [PaymentListParams]'s query parameters as `url.Values`.
func (r PaymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaymentListParamsCreated struct {
	// If passed, return only payments created at or after the specified time
	Gte param.Field[string] `query:"gte" format:"ISO 8601 (YYYY-MM-DD'T'HH24:MI:SS)"`
	// If passed, return only payments created strictly before the specified time
	Lt param.Field[string] `query:"lt" format:"ISO 8601 (YYYY-MM-DD'T'HH24:MI:SS)"`
}

// URLQuery serializes [PaymentListParamsCreated]'s query parameters as
// `url.Values`.
func (r PaymentListParamsCreated) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaymentListParamsUpdated struct {
	// If passed, return only payments updated at or after the specified time
	Gte param.Field[string] `query:"gte" format:"ISO 8601 (YYYY-MM-DD'T'HH24:MI:SS)"`
	// If passed, return only payments updated strictly before the specified time
	Lt param.Field[string] `query:"lt" format:"ISO 8601 (YYYY-MM-DD'T'HH24:MI:SS)"`
}

// URLQuery serializes [PaymentListParamsUpdated]'s query parameters as
// `url.Values`.
func (r PaymentListParamsUpdated) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
