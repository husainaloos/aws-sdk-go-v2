// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	"net/http"
	"time"
)

const ServiceID = "Support"

// AWS Support The AWS Support API reference is intended for programmers who need
// detailed information about the AWS Support operations and data types. This
// service enables you to manage your AWS Support cases programmatically. It uses
// HTTP methods that return results in JSON format.
//
//     * You must have a Business
// or Enterprise support plan to use the AWS Support API.
//
//     * If you call the
// AWS Support API from an account that does not have a Business or Enterprise
// support plan, the SubscriptionRequiredException error message appears. For
// information about changing your support plan, see AWS Support
// (http://aws.amazon.com/premiumsupport/).
//
// The AWS Support service also exposes a
// set of AWS Trusted Advisor
// (http://aws.amazon.com/premiumsupport/trustedadvisor/) features. You can
// retrieve a list of checks and their descriptions, get check results, specify
// checks to refresh, and get the refresh status of checks. The following list
// describes the AWS Support case management operations:
//
//     * Service names,
// issue categories, and available severity levels. The DescribeServices () and
// DescribeSeverityLevels () operations return AWS service names, service codes,
// service categories, and problem severity levels. You use these values when you
// call the CreateCase () operation.
//
//     * Case creation, case details, and case
// resolution. The CreateCase (), DescribeCases (), DescribeAttachment (), and
// ResolveCase () operations create AWS Support cases, retrieve information about
// cases, and resolve cases.
//
//     * Case communication. The DescribeCommunications
// (), AddCommunicationToCase (), and AddAttachmentsToSet () operations retrieve
// and add communications and attachments to AWS Support cases.
//
// The following list
// describes the operations available from the AWS Support service for Trusted
// Advisor:
//
//     * DescribeTrustedAdvisorChecks () returns the list of checks that
// run against your AWS resources.
//
//     * Using the checkId for a specific check
// returned by DescribeTrustedAdvisorChecks (), you can call
// DescribeTrustedAdvisorCheckResult () to obtain the results for the check that
// you specified.
//
//     * DescribeTrustedAdvisorCheckSummaries () returns summarized
// results for one or more Trusted Advisor checks.
//
//     *
// RefreshTrustedAdvisorCheck () requests that Trusted Advisor rerun a specified
// check.
//
//     * DescribeTrustedAdvisorCheckRefreshStatuses () reports the refresh
// status of one or more checks.
//
// For authentication of requests, AWS Support uses
// Signature Version 4 Signing Process
// (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html). See
// About the AWS Support API
// (https://docs.aws.amazon.com/awssupport/latest/user/Welcome.html) in the AWS
// Support User Guide for information about how to use this service to create and
// manage your support cases, and how to call Trusted Advisor for results of checks
// on your resources.
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

	resolveHTTPSignerV4(&options)

	resolveDefaultEndpointConfiguration(&options)

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

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions ResolverOptions

	// The service endpoint resolver.
	EndpointResolver EndpointResolver

	// Signature Version 4 (SigV4) Signer
	HTTPSignerV4 HTTPSignerV4

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

func (o Options) GetEndpointOptions() ResolverOptions {
	return o.EndpointOptions
}

func (o Options) GetEndpointResolver() EndpointResolver {
	return o.EndpointResolver
}

func (o Options) GetHTTPSignerV4() HTTPSignerV4 {
	return o.HTTPSignerV4
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
	awsmiddleware.AddUserAgentKey("support")(stack)
}

func addHTTPSignerV4Middleware(stack *middleware.Stack, o Options) {
	stack.Finalize.Add(v4.NewSignHTTPRequestMiddleware(o.Credentials, o.HTTPSignerV4), middleware.After)
}

type HTTPSignerV4 interface {
	SignHTTP(ctx context.Context, credentials aws.Credentials, r *http.Request, payloadHash string, service string, region string, signingTime time.Time) error
}

func resolveHTTPSignerV4(o *Options) {
	if o.HTTPSignerV4 != nil {
		return
	}
	o.HTTPSignerV4 = v4.NewSigner()
}
