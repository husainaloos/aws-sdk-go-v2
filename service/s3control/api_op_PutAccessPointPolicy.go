// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates an access policy with the specified access point. Each access point
// can have only one policy, so a request made to this API replaces any existing
// policy associated with the specified access point.
func (c *Client) PutAccessPointPolicy(ctx context.Context, params *PutAccessPointPolicyInput, optFns ...func(*Options)) (*PutAccessPointPolicyOutput, error) {
	stack := middleware.NewStack("PutAccessPointPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpPutAccessPointPolicyMiddlewares(stack)
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
	addOpPutAccessPointPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutAccessPointPolicy(options.Region), middleware.Before)
	addResponseErrorMiddleware(stack)
	addMetadataRetrieverMiddleware(stack)

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
			OperationName: "PutAccessPointPolicy",
			Err:           err,
		}
	}
	out := result.(*PutAccessPointPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAccessPointPolicyInput struct {
	// The policy that you want to apply to the specified access point. For more
	// information about access point policies, see Managing Data Access with Amazon S3
	// Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-points.html) in the
	// Amazon Simple Storage Service Developer Guide.
	Policy *string
	// The AWS account ID for owner of the bucket associated with the specified access
	// point.
	AccountId *string
	// The name of the access point that you want to associate with the specified
	// policy.
	Name *string
}

type PutAccessPointPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpPutAccessPointPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpPutAccessPointPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpPutAccessPointPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutAccessPointPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutAccessPointPolicy",
	}
}
