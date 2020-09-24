// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The GetAssignment operation retrieves the details of the specified Assignment.
func (c *Client) GetAssignment(ctx context.Context, params *GetAssignmentInput, optFns ...func(*Options)) (*GetAssignmentOutput, error) {
	stack := middleware.NewStack("GetAssignment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetAssignmentMiddlewares(stack)
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
	addOpGetAssignmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAssignment(options.Region), middleware.Before)
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
			OperationName: "GetAssignment",
			Err:           err,
		}
	}
	out := result.(*GetAssignmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssignmentInput struct {
	// The ID of the Assignment to be retrieved.
	AssignmentId *string
}

type GetAssignmentOutput struct {
	// The HIT associated with this assignment. The response includes one HIT element.
	HIT *types.HIT
	// The assignment. The response includes one Assignment element.
	Assignment *types.Assignment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetAssignmentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetAssignment{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAssignment{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAssignment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "GetAssignment",
	}
}
