// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Move a dedicated IP address to an existing dedicated IP pool. The dedicated IP
// address that you specify must already exist, and must be associated with your
// Amazon Pinpoint account.  </p> <p>The dedicated IP pool you specify must already
// exist. You can create a new pool by using the <code>CreateDedicatedIpPool</code>
// operation.</p> </note>
func (c *Client) PutDedicatedIpInPool(ctx context.Context, params *PutDedicatedIpInPoolInput, optFns ...func(*Options)) (*PutDedicatedIpInPoolOutput, error) {
	stack := middleware.NewStack("PutDedicatedIpInPool", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutDedicatedIpInPoolMiddlewares(stack)
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
	addOpPutDedicatedIpInPoolValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutDedicatedIpInPool(options.Region), middleware.Before)
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
			OperationName: "PutDedicatedIpInPool",
			Err:           err,
		}
	}
	out := result.(*PutDedicatedIpInPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to move a dedicated IP address to a dedicated IP pool.
type PutDedicatedIpInPoolInput struct {
	// The IP address that you want to move to the dedicated IP pool. The value you
	// specify has to be a dedicated IP address that's associated with your Amazon
	// Pinpoint account.
	Ip *string
	// The name of the IP pool that you want to add the dedicated IP address to. You
	// have to specify an IP pool that already exists.
	DestinationPoolName *string
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type PutDedicatedIpInPoolOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutDedicatedIpInPoolMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutDedicatedIpInPool{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutDedicatedIpInPool{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutDedicatedIpInPool(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutDedicatedIpInPool",
	}
}
