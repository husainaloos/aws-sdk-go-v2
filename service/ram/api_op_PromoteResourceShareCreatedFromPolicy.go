// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Resource shares that were created by attaching a policy to a resource are
// visible only to the resource share owner, and the resource share cannot be
// modified in AWS RAM.  <p>Use this API action to promote the resource share. When
// you promote the resource share, it becomes:</p> <ul> <li> <p>Visible to all
// principals that it is shared with.</p> </li> <li> <p>Modifiable in AWS RAM.</p>
// </li> </ul>
func (c *Client) PromoteResourceShareCreatedFromPolicy(ctx context.Context, params *PromoteResourceShareCreatedFromPolicyInput, optFns ...func(*Options)) (*PromoteResourceShareCreatedFromPolicyOutput, error) {
	stack := middleware.NewStack("PromoteResourceShareCreatedFromPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPromoteResourceShareCreatedFromPolicyMiddlewares(stack)
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
	addOpPromoteResourceShareCreatedFromPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPromoteResourceShareCreatedFromPolicy(options.Region), middleware.Before)
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
			OperationName: "PromoteResourceShareCreatedFromPolicy",
			Err:           err,
		}
	}
	out := result.(*PromoteResourceShareCreatedFromPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PromoteResourceShareCreatedFromPolicyInput struct {
	// The ARN of the resource share to promote.
	ResourceShareArn *string
}

type PromoteResourceShareCreatedFromPolicyOutput struct {
	// Indicates whether the request succeeded.
	ReturnValue *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPromoteResourceShareCreatedFromPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPromoteResourceShareCreatedFromPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPromoteResourceShareCreatedFromPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opPromoteResourceShareCreatedFromPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "PromoteResourceShareCreatedFromPolicy",
	}
}
