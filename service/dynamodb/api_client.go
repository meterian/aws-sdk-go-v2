// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	cryptorand "crypto/rand"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	ddbcust "github.com/aws/aws-sdk-go-v2/service/dynamodb/internal/customizations"
	"github.com/awslabs/smithy-go/middleware"
	smithyrand "github.com/awslabs/smithy-go/rand"
	"net/http"
)

const ServiceID = "DynamoDB"

// Amazon DynamoDB  <p>Amazon DynamoDB is a fully managed NoSQL database service
// that provides fast and predictable performance with seamless scalability.
// DynamoDB lets you offload the administrative burdens of operating and scaling a
// distributed database, so that you don't have to worry about hardware
// provisioning, setup and configuration, replication, software patching, or
// cluster scaling.</p> <p>With DynamoDB, you can create database tables that can
// store and retrieve any amount of data, and serve any level of request traffic.
// You can scale up or scale down your tables' throughput capacity without downtime
// or performance degradation, and use the AWS Management Console to monitor
// resource utilization and performance metrics.</p> <p>DynamoDB automatically
// spreads the data and traffic for your tables over a sufficient number of servers
// to handle your throughput and storage requirements, while maintaining consistent
// and fast performance. All of your data is stored on solid state disks (SSDs) and
// automatically replicated across multiple Availability Zones in an AWS region,
// providing built-in high availability and data durability. </p>
type Client struct {
	options Options
}

// New returns an initialized Client based on the functional options. Provide
// additional functional options to further configure the behavior of the client,
// such as changing the client's endpoint or adding custom middleware behavior.
func New(options Options, optFns ...func(*Options)) *Client {
	options = options.Copy()

	resolveRetryer(&options)

	resolveHTTPClient(&options)

	resolveDefaultEndpointConfiguration(&options)

	resolveIdempotencyTokenProvider(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	client := &Client{
		options: options,
	}

	return client
}

type Options struct {
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []APIOptionFunc

	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// Allows you to disable the client's validation of response integrity using CRC32
	// checksum. Enabled by default.
	DisableValidateResponseChecksum bool

	// Allows you to enable the client's support for compressed gzip responses.
	// Disabled by default.
	EnableAcceptEncodingGzip bool

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions ResolverOptions

	// The service endpoint resolver.
	EndpointResolver EndpointResolver

	// Provides idempotency tokens values that will be automatically populated into
	// idempotent API operations.
	IdempotencyTokenProvider IdempotencyTokenProvider

	// An integer value representing the logging level.
	LogLevel aws.LogLevel

	// The logger writer interface to write logging messages to.
	Logger aws.Logger

	// The region to send requests to. (Required)
	Region string

	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer.
	Retryer retry.Retryer

	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient
}

func (o Options) GetCredentials() aws.CredentialsProvider {
	return o.Credentials
}

func (o Options) GetDisableValidateResponseChecksum() bool {
	return o.DisableValidateResponseChecksum
}

func (o Options) GetEnableAcceptEncodingGzip() bool {
	return o.EnableAcceptEncodingGzip
}

func (o Options) GetEndpointOptions() ResolverOptions {
	return o.EndpointOptions
}

func (o Options) GetEndpointResolver() EndpointResolver {
	return o.EndpointResolver
}

func (o Options) GetIdempotencyTokenProvider() IdempotencyTokenProvider {
	return o.IdempotencyTokenProvider
}

func (o Options) GetLogLevel() aws.LogLevel {
	return o.LogLevel
}

func (o Options) GetLogger() aws.Logger {
	return o.Logger
}

func (o Options) GetRegion() string {
	return o.Region
}

func (o Options) GetRetryer() retry.Retryer {
	return o.Retryer
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	to.APIOptions = make([]APIOptionFunc, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)
	return to
}

type APIOptionFunc func(*middleware.Stack) error

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:      cfg.Region,
		Retryer:     cfg.Retryer,
		LogLevel:    cfg.LogLevel,
		Logger:      cfg.Logger,
		HTTPClient:  cfg.HTTPClient,
		Credentials: cfg.Credentials,
	}
	return New(opts, optFns...)
}

func resolveHTTPClient(o *Options) {
	if o.HTTPClient != nil {
		return
	}
	o.HTTPClient = aws.NewBuildableHTTPClient()
}

func resolveRetryer(o *Options) {
	if o.Retryer != nil {
		return
	}
	o.Retryer = retry.NewStandard()
}

func addClientUserAgent(stack *middleware.Stack) {
	awsmiddleware.AddUserAgentKey("dynamodb")(stack)
}

func addHTTPSignerV4Middleware(stack *middleware.Stack, o Options) {
	signer := v4.Signer{}
	stack.Finalize.Add(v4.NewSignHTTPRequestMiddleware(o.Credentials, signer), middleware.After)
}

func resolveIdempotencyTokenProvider(o *Options) {
	if o.IdempotencyTokenProvider != nil {
		return
	}
	o.IdempotencyTokenProvider = smithyrand.NewUUIDIdempotencyToken(cryptorand.Reader)
}

// IdempotencyTokenProvider interface for providing idempotency token
type IdempotencyTokenProvider interface {
	GetIdempotencyToken() (string, error)
}

func addValidateResponseChecksum(stack *middleware.Stack, options Options) {
	ddbcust.AddValidateResponseChecksum(stack, ddbcust.AddValidateResponseChecksumOptions{Disable: options.DisableValidateResponseChecksum})
}

func addAcceptEncodingGzip(stack *middleware.Stack, options Options) {
	ddbcust.AddAcceptEncodingGzip(stack, ddbcust.AddAcceptEncodingGzipOptions{Enable: options.EnableAcceptEncodingGzip})
}
