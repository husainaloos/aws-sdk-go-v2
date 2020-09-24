// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an access policy that grants the specified AWS Single Sign-On user or
// group access to the specified AWS IoT SiteWise Monitor portal or project
// resource.
func (c *Client) CreateAccessPolicy(ctx context.Context, params *CreateAccessPolicyInput, optFns ...func(*Options)) (*CreateAccessPolicyOutput, error) {
	stack := middleware.NewStack("CreateAccessPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateAccessPolicyMiddlewares(stack)
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
	addIdempotencyToken_opCreateAccessPolicyMiddleware(stack, options)
	addOpCreateAccessPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccessPolicy(options.Region), middleware.Before)
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
			OperationName: "CreateAccessPolicy",
			Err:           err,
		}
	}
	out := result.(*CreateAccessPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccessPolicyInput struct {
	// The identity for this access policy. Choose either a user or a group but not
	// both.
	AccessPolicyIdentity *types.Identity
	// A unique case-sensitive identifier that you can provide to ensure the
	// idempotency of the request. Don't reuse this client token if a new idempotent
	// request is required.
	ClientToken *string
	// The AWS IoT SiteWise Monitor resource for this access policy. Choose either
	// portal or project but not both.
	AccessPolicyResource *types.Resource
	// The permission level for this access policy. Note that a project ADMINISTRATOR
	// is also known as a project owner.
	AccessPolicyPermission types.Permission
	// A list of key-value pairs that contain metadata for the access policy. For more
	// information, see Tagging your AWS IoT SiteWise resources
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html)
	// in the AWS IoT SiteWise User Guide.
	Tags map[string]*string
}

type CreateAccessPolicyOutput struct {
	// The ID of the access policy.
	AccessPolicyId *string
	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the access policy, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:access-policy/${AccessPolicyId}
	AccessPolicyArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateAccessPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateAccessPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAccessPolicy{}, middleware.After)
}

type idempotencyToken_initializeOpCreateAccessPolicy struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAccessPolicy) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAccessPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAccessPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAccessPolicyInput ")
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
func addIdempotencyToken_opCreateAccessPolicyMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateAccessPolicy{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAccessPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "CreateAccessPolicy",
	}
}
