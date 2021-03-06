// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get a list of key phrase detection jobs that you have submitted.
func (c *Client) ListKeyPhrasesDetectionJobs(ctx context.Context, params *ListKeyPhrasesDetectionJobsInput, optFns ...func(*Options)) (*ListKeyPhrasesDetectionJobsOutput, error) {
	if params == nil {
		params = &ListKeyPhrasesDetectionJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListKeyPhrasesDetectionJobs", params, optFns, addOperationListKeyPhrasesDetectionJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListKeyPhrasesDetectionJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListKeyPhrasesDetectionJobsInput struct {

	// Filters the jobs that are returned. You can filter jobs on their name, status,
	// or the date and time that they were submitted. You can only set one filter at a
	// time.
	Filter *types.KeyPhrasesDetectionJobFilter

	// The maximum number of results to return in each page. The default is 100.
	MaxResults *int32

	// Identifies the next page of results to return.
	NextToken *string
}

type ListKeyPhrasesDetectionJobsOutput struct {

	// A list containing the properties of each job that is returned.
	KeyPhrasesDetectionJobPropertiesList []*types.KeyPhrasesDetectionJobProperties

	// Identifies the next page of results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListKeyPhrasesDetectionJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListKeyPhrasesDetectionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListKeyPhrasesDetectionJobs{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListKeyPhrasesDetectionJobs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListKeyPhrasesDetectionJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "ListKeyPhrasesDetectionJobs",
	}
}
