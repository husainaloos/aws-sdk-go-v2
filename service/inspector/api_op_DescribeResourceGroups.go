// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the resource groups that are specified by the ARNs of the resource
// groups.
func (c *Client) DescribeResourceGroups(ctx context.Context, params *DescribeResourceGroupsInput, optFns ...func(*Options)) (*DescribeResourceGroupsOutput, error) {
	stack := middleware.NewStack("DescribeResourceGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeResourceGroupsMiddlewares(stack)
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
	addOpDescribeResourceGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeResourceGroups(options.Region), middleware.Before)

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
			OperationName: "DescribeResourceGroups",
			Err:           err,
		}
	}
	out := result.(*DescribeResourceGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeResourceGroupsInput struct {
	// The ARN that specifies the resource group that you want to describe.
	ResourceGroupArns []*string
}

type DescribeResourceGroupsOutput struct {
	// Resource group details that cannot be described. An error code is provided for
	// each failed item.
	FailedItems map[string]*types.FailedItemDetails
	// Information about a resource group.
	ResourceGroups []*types.ResourceGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeResourceGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeResourceGroups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeResourceGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeResourceGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "DescribeResourceGroups",
	}
}