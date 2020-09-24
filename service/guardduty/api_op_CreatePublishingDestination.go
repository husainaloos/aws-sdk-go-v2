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

// Creates a publishing destination to export findings to. The resource to export
// findings to must exist before you use this operation.
func (c *Client) CreatePublishingDestination(ctx context.Context, params *CreatePublishingDestinationInput, optFns ...func(*Options)) (*CreatePublishingDestinationOutput, error) {
	stack := middleware.NewStack("CreatePublishingDestination", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreatePublishingDestinationMiddlewares(stack)
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
	addIdempotencyToken_opCreatePublishingDestinationMiddleware(stack, options)
	addOpCreatePublishingDestinationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePublishingDestination(options.Region), middleware.Before)
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
			OperationName: "CreatePublishingDestination",
			Err:           err,
		}
	}
	out := result.(*CreatePublishingDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePublishingDestinationInput struct {
	// The type of resource for the publishing destination. Currently only Amazon S3
	// buckets are supported.
	DestinationType types.DestinationType
	// The ID of the GuardDuty detector associated with the publishing destination.
	DetectorId *string
	// The idempotency token for the request.
	ClientToken *string
	// The properties of the publishing destination, including the ARNs for the
	// destination and the KMS key used for encryption.
	DestinationProperties *types.DestinationProperties
}

type CreatePublishingDestinationOutput struct {
	// The ID of the publishing destination that is created.
	DestinationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreatePublishingDestinationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreatePublishingDestination{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePublishingDestination{}, middleware.After)
}

type idempotencyToken_initializeOpCreatePublishingDestination struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePublishingDestination) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePublishingDestination) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePublishingDestinationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePublishingDestinationInput ")
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
func addIdempotencyToken_opCreatePublishingDestinationMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreatePublishingDestination{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePublishingDestination(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "CreatePublishingDestination",
	}
}
