// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of solutions that use the given dataset group. When a dataset
// group is not specified, all the solutions associated with the account are
// listed. The response provides the properties for each solution, including the
// Amazon Resource Name (ARN). For more information on solutions, see
// CreateSolution ().
func (c *Client) ListSolutions(ctx context.Context, params *ListSolutionsInput, optFns ...func(*Options)) (*ListSolutionsOutput, error) {
	stack := middleware.NewStack("ListSolutions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListSolutionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSolutions(options.Region), middleware.Before)
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
			OperationName: "ListSolutions",
			Err:           err,
		}
	}
	out := result.(*ListSolutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSolutionsInput struct {
	// The Amazon Resource Name (ARN) of the dataset group.
	DatasetGroupArn *string
	// The maximum number of solutions to return.
	MaxResults *int32
	// A token returned from the previous call to ListSolutions for getting the next
	// set of solutions (if they exist).
	NextToken *string
}

type ListSolutionsOutput struct {
	// A token for getting the next set of solutions (if they exist).
	NextToken *string
	// A list of the current solutions.
	Solutions []*types.SolutionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListSolutionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListSolutions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSolutions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListSolutions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "ListSolutions",
	}
}
