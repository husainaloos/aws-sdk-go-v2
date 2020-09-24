// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new IPSet, which is called a trusted IP list in the console user
// interface. An IPSet is a list of IP addresses that are trusted for secure
// communication with AWS infrastructure and applications. GuardDuty doesn't
// generate findings for IP addresses that are included in IPSets. Only users from
// the master account can use this operation.
func (c *Client) CreateIPSet(ctx context.Context, params *CreateIPSetInput, optFns ...func(*Options)) (*CreateIPSetOutput, error) {
	stack := middleware.NewStack("CreateIPSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateIPSetMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addIdempotencyToken_opCreateIPSetMiddleware(stack, options)
	addOpCreateIPSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIPSet(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "CreateIPSet",
			Err:           err,
		}
	}
	out := result.(*CreateIPSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIPSetInput struct {
	// The format of the file that contains the IPSet.
	Format types.IpSetFormat
	// A Boolean value that indicates whether GuardDuty is to start using the uploaded
	// IPSet.
	Activate *bool
	// The URI of the file that contains the IPSet. For example:
	// https://s3.us-west-2.amazonaws.com/my-bucket/my-object-key.
	Location *string
	// The unique ID of the detector of the GuardDuty account that you want to create
	// an IPSet for.
	DetectorId *string
	// The user-friendly name to identify the IPSet. Allowed characters are
	// alphanumerics, spaces, hyphens (-), and underscores (_).
	Name *string
	// The idempotency token for the create request.
	ClientToken *string
	// The tags to be added to a new IP set resource.
	Tags map[string]*string
}

type CreateIPSetOutput struct {
	// The ID of the IPSet resource.
	IpSetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateIPSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateIPSet{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateIPSet{}, middleware.After)
}

type idempotencyToken_initializeOpCreateIPSet struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateIPSet) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateIPSet) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateIPSetInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateIPSetInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateIPSetMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateIPSet{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateIPSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "CreateIPSet",
	}
}
