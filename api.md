# Payments

Response Types:

- <a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go">paymaxis</a>.<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go#Payment">Payment</a>
- <a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go">paymaxis</a>.<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go#PaymentListResponse">PaymentListResponse</a>

Methods:

- <code title="post /api/v1/payments">client.Payments.<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go#PaymentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go">paymaxis</a>.<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go#PaymentNewParams">PaymentNewParams</a>) (<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go">paymaxis</a>.<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/payments/{id}">client.Payments.<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go#PaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go">paymaxis</a>.<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/payments">client.Payments.<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go#PaymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go">paymaxis</a>.<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go#PaymentListParams">PaymentListParams</a>) (<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go">paymaxis</a>.<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go#PaymentListResponse">PaymentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Operations

Response Types:

- <a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go">paymaxis</a>.<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go#Operation">Operation</a>

Methods:

- <code title="get /api/v1/payments/{id}/operations">client.Payments.Operations.<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go#PaymentOperationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go">paymaxis</a>.<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Subscriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go">paymaxis</a>.<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go#Subscription">Subscription</a>

Methods:

- <code title="get /api/v1/subscriptions/{id}">client.Subscriptions.<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go#SubscriptionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go">paymaxis</a>.<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go#Subscription">Subscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/v1/subscriptions/{id}">client.Subscriptions.<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go#SubscriptionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go">paymaxis</a>.<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go#SubscriptionUpdateParams">SubscriptionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go">paymaxis</a>.<a href="https://pkg.go.dev/github.com/TralahM/paymaxis-go#Subscription">Subscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
