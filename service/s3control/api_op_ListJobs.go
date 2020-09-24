// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists current Amazon S3 Batch Operations jobs and jobs that have ended within
// the last 30 days for the AWS account making the request. For more information,
// see Amazon S3 Batch Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-basics.html) in the
// Amazon Simple Storage Service Developer Guide. Related actions include:
//
//     *
// CreateJob ()
//
//     * DescribeJob ()
//
//     * UpdateJobPriority ()
//
//     *
// UpdateJobStatus ()
func (c *Client) ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) {
	stack := middleware.NewStack("ListJobs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpListJobsMiddlewares(stack)
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
	addOpListJobsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListJobs(options.Region), middleware.Before)
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
			OperationName: "ListJobs",
			Err:           err,
		}
	}
	out := result.(*ListJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJobsInput struct {
	// The List Jobs request returns jobs that match the statuses listed in this
	// element.
	JobStatuses []types.JobStatus
	// A pagination token to request the next page of results. Use the token that
	// Amazon S3 returned in the NextToken element of the ListJobsResult from the
	// previous List Jobs request.
	NextToken *string
	// The maximum number of jobs that Amazon S3 will include in the List Jobs
	// response. If there are more jobs than this number, the response will include a
	// pagination token in the NextToken field to enable you to retrieve the next page
	// of results.
	MaxResults *int32
	//
	AccountId *string
}

type ListJobsOutput struct {
	// The list of current jobs and jobs that have ended within the last 30 days.
	Jobs []*types.JobListDescriptor
	// If the List Jobs request produced more than the maximum number of results, you
	// can pass this value into a subsequent List Jobs request in order to retrieve the
	// next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpListJobsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpListJobs{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpListJobs{}, middleware.After)
}

func newServiceMetadataMiddleware_opListJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListJobs",
	}
}
