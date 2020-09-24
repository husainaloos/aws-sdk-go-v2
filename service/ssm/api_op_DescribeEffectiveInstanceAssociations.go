// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// All associations for the instance(s).
func (c *Client) DescribeEffectiveInstanceAssociations(ctx context.Context, params *DescribeEffectiveInstanceAssociationsInput, optFns ...func(*Options)) (*DescribeEffectiveInstanceAssociationsOutput, error) {
	stack := middleware.NewStack("DescribeEffectiveInstanceAssociations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeEffectiveInstanceAssociationsMiddlewares(stack)
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
	addOpDescribeEffectiveInstanceAssociationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEffectiveInstanceAssociations(options.Region), middleware.Before)
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
			OperationName: "DescribeEffectiveInstanceAssociations",
			Err:           err,
		}
	}
	out := result.(*DescribeEffectiveInstanceAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEffectiveInstanceAssociationsInput struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32
	// The instance ID for which you want to view all associations.
	InstanceId *string
	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
}

type DescribeEffectiveInstanceAssociationsOutput struct {
	// The associations for the requested instance.
	Associations []*types.InstanceAssociation
	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeEffectiveInstanceAssociationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEffectiveInstanceAssociations{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEffectiveInstanceAssociations{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeEffectiveInstanceAssociations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeEffectiveInstanceAssociations",
	}
}
