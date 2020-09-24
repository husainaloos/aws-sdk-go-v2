// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds tags to on-premises instances.
func (c *Client) AddTagsToOnPremisesInstances(ctx context.Context, params *AddTagsToOnPremisesInstancesInput, optFns ...func(*Options)) (*AddTagsToOnPremisesInstancesOutput, error) {
	stack := middleware.NewStack("AddTagsToOnPremisesInstances", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAddTagsToOnPremisesInstancesMiddlewares(stack)
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
	addOpAddTagsToOnPremisesInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddTagsToOnPremisesInstances(options.Region), middleware.Before)
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
			OperationName: "AddTagsToOnPremisesInstances",
			Err:           err,
		}
	}
	out := result.(*AddTagsToOnPremisesInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of, and adds tags to, an on-premises instance operation.
type AddTagsToOnPremisesInstancesInput struct {
	// The names of the on-premises instances to which to add tags.
	InstanceNames []*string
	// The tag key-value pairs to add to the on-premises instances. Keys and values are
	// both required. Keys cannot be null or empty strings. Value-only tags are not
	// allowed.
	Tags []*types.Tag
}

type AddTagsToOnPremisesInstancesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAddTagsToOnPremisesInstancesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAddTagsToOnPremisesInstances{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddTagsToOnPremisesInstances{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddTagsToOnPremisesInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "AddTagsToOnPremisesInstances",
	}
}
