// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a plan. A plan includes the list of resources to be created (when
// provisioning a new product) or modified (when updating a provisioned product)
// when the plan is executed. You can create one plan per provisioned product. To
// create a plan for an existing provisioned product, the product status must be
// AVAILBLE or TAINTED. To view the resource changes in the change set, use
// DescribeProvisionedProductPlan (). To create or modify the provisioned product,
// use ExecuteProvisionedProductPlan ().
func (c *Client) CreateProvisionedProductPlan(ctx context.Context, params *CreateProvisionedProductPlanInput, optFns ...func(*Options)) (*CreateProvisionedProductPlanOutput, error) {
	stack := middleware.NewStack("CreateProvisionedProductPlan", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateProvisionedProductPlanMiddlewares(stack)
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
	addIdempotencyToken_opCreateProvisionedProductPlanMiddleware(stack, options)
	addOpCreateProvisionedProductPlanValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProvisionedProductPlan(options.Region), middleware.Before)
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
			OperationName: "CreateProvisionedProductPlan",
			Err:           err,
		}
	}
	out := result.(*CreateProvisionedProductPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProvisionedProductPlanInput struct {
	// A unique identifier that you provide to ensure idempotency. If multiple requests
	// differ only by the idempotency token, the same response is returned for each
	// repeated request.
	IdempotencyToken *string
	// A user-friendly name for the provisioned product. This value must be unique for
	// the AWS account and cannot be updated after the product is provisioned.
	ProvisionedProductName *string
	// One or more tags. If the plan is for an existing provisioned product, the
	// product must have a RESOURCE_UPDATE constraint with
	// TagUpdatesOnProvisionedProduct set to ALLOWED to allow tag updates.
	Tags []*types.Tag
	// The path identifier of the product. This value is optional if the product has a
	// default path, and required if the product has more than one path. To list the
	// paths for a product, use ListLaunchPaths ().
	PathId *string
	// Parameters specified by the administrator that are required for provisioning the
	// product.
	ProvisioningParameters []*types.UpdateProvisioningParameter
	// The product identifier.
	ProductId *string
	// The name of the plan.
	PlanName *string
	// The plan type.
	PlanType types.ProvisionedProductPlanType
	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string
	// The identifier of the provisioning artifact.
	ProvisioningArtifactId *string
	// Passed to CloudFormation. The SNS topic ARNs to which to publish stack-related
	// events.
	NotificationArns []*string
}

type CreateProvisionedProductPlanOutput struct {
	// The name of the plan.
	PlanName *string
	// The user-friendly name of the provisioned product.
	ProvisionedProductName *string
	// The identifier of the provisioning artifact.
	ProvisioningArtifactId *string
	// The plan identifier.
	PlanId *string
	// The product identifier.
	ProvisionProductId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateProvisionedProductPlanMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProvisionedProductPlan{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProvisionedProductPlan{}, middleware.After)
}

type idempotencyToken_initializeOpCreateProvisionedProductPlan struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateProvisionedProductPlan) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateProvisionedProductPlan) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateProvisionedProductPlanInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateProvisionedProductPlanInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateProvisionedProductPlanMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateProvisionedProductPlan{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateProvisionedProductPlan(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "CreateProvisionedProductPlan",
	}
}
