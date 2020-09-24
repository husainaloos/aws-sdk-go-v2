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

// Gets information about one or more application revisions. The maximum number of
// application revisions that can be returned is 25.
func (c *Client) BatchGetApplicationRevisions(ctx context.Context, params *BatchGetApplicationRevisionsInput, optFns ...func(*Options)) (*BatchGetApplicationRevisionsOutput, error) {
	stack := middleware.NewStack("BatchGetApplicationRevisions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchGetApplicationRevisionsMiddlewares(stack)
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
	addOpBatchGetApplicationRevisionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetApplicationRevisions(options.Region), middleware.Before)
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
			OperationName: "BatchGetApplicationRevisions",
			Err:           err,
		}
	}
	out := result.(*BatchGetApplicationRevisionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a BatchGetApplicationRevisions operation.
type BatchGetApplicationRevisionsInput struct {
	// An array of RevisionLocation objects that specify information to get about the
	// application revisions, including type and location. The maximum number of
	// RevisionLocation objects you can specify is 25.
	Revisions []*types.RevisionLocation
	// The name of an AWS CodeDeploy application about which to get revision
	// information.
	ApplicationName *string
}

// Represents the output of a BatchGetApplicationRevisions operation.
type BatchGetApplicationRevisionsOutput struct {
	// Information about errors that might have occurred during the API call.
	ErrorMessage *string
	// Additional information about the revisions, including the type and location.
	Revisions []*types.RevisionInfo
	// The name of the application that corresponds to the revisions.
	ApplicationName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchGetApplicationRevisionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetApplicationRevisions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetApplicationRevisions{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchGetApplicationRevisions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "BatchGetApplicationRevisions",
	}
}
