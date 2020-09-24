// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Grants entitlements to an existing flow.
func (c *Client) GrantFlowEntitlements(ctx context.Context, params *GrantFlowEntitlementsInput, optFns ...func(*Options)) (*GrantFlowEntitlementsOutput, error) {
	stack := middleware.NewStack("GrantFlowEntitlements", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGrantFlowEntitlementsMiddlewares(stack)
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
	addOpGrantFlowEntitlementsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGrantFlowEntitlements(options.Region), middleware.Before)
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
			OperationName: "GrantFlowEntitlements",
			Err:           err,
		}
	}
	out := result.(*GrantFlowEntitlementsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to grant entitlements on a flow.
type GrantFlowEntitlementsInput struct {
	// The list of entitlements that you want to grant.
	Entitlements []*types.GrantEntitlementRequest
	// The flow that you want to grant entitlements on.
	FlowArn *string
}

type GrantFlowEntitlementsOutput struct {
	// The entitlements that were just granted.
	Entitlements []*types.Entitlement
	// The ARN of the flow that these entitlements were granted to.
	FlowArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGrantFlowEntitlementsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGrantFlowEntitlements{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGrantFlowEntitlements{}, middleware.After)
}

func newServiceMetadataMiddleware_opGrantFlowEntitlements(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "GrantFlowEntitlements",
	}
}
