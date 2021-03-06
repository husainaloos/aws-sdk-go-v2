// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of solution versions for the given solution. When a solution is
// not specified, all the solution versions associated with the account are listed.
// The response provides the properties for each solution version, including the
// Amazon Resource Name (ARN). For more information on solutions, see
// CreateSolution ().
func (c *Client) ListSolutionVersions(ctx context.Context, params *ListSolutionVersionsInput, optFns ...func(*Options)) (*ListSolutionVersionsOutput, error) {
	if params == nil {
		params = &ListSolutionVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSolutionVersions", params, optFns, addOperationListSolutionVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSolutionVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSolutionVersionsInput struct {

	// The maximum number of solution versions to return.
	MaxResults *int32

	// A token returned from the previous call to ListSolutionVersions for getting the
	// next set of solution versions (if they exist).
	NextToken *string

	// The Amazon Resource Name (ARN) of the solution.
	SolutionArn *string
}

type ListSolutionVersionsOutput struct {

	// A token for getting the next set of solution versions (if they exist).
	NextToken *string

	// A list of solution versions describing the version properties.
	SolutionVersions []*types.SolutionVersionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSolutionVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSolutionVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSolutionVersions{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSolutionVersions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListSolutionVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "ListSolutionVersions",
	}
}
