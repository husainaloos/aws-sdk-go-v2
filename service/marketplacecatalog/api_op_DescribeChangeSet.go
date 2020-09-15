// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides information about a given change set.
func (c *Client) DescribeChangeSet(ctx context.Context, params *DescribeChangeSetInput, optFns ...func(*Options)) (*DescribeChangeSetOutput, error) {
	stack := middleware.NewStack("DescribeChangeSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeChangeSetMiddlewares(stack)
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
	addOpDescribeChangeSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeChangeSet(options.Region), middleware.Before)

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
			OperationName: "DescribeChangeSet",
			Err:           err,
		}
	}
	out := result.(*DescribeChangeSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeChangeSetInput struct {
	// Required. The catalog related to the request. Fixed value: AWSMarketplace
	Catalog *string
	// Required. The unique identifier for the StartChangeSet request that you want to
	// describe the details for.
	ChangeSetId *string
}

type DescribeChangeSetOutput struct {
	// The date and time, in ISO 8601 format (2018-02-27T13:45:22Z), the request
	// transitioned to a terminal state. The change cannot transition to a different
	// state. Null if the request is not in a terminal state.
	EndTime *string
	// The optional name provided in the StartChangeSet request. If you do not provide
	// a name, one is set by default.
	ChangeSetName *string
	// The status of the change request.
	Status types.ChangeStatus
	// Required. The unique identifier for the change set referenced in this request.
	ChangeSetId *string
	// The date and time, in ISO 8601 format (2018-02-27T13:45:22Z), the request
	// started.
	StartTime *string
	// The ARN associated with the unique identifier for the change set referenced in
	// this request.
	ChangeSetArn *string
	// Returned if there is a failure on the change set, but that failure is not
	// related to any of the changes in the request.
	FailureDescription *string
	// An array of ChangeSummary objects.
	ChangeSet []*types.ChangeSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeChangeSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeChangeSet{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeChangeSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeChangeSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aws-marketplace",
		OperationName: "DescribeChangeSet",
	}
}