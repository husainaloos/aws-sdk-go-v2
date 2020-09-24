// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves a service last accessed report that was created using the
// GenerateServiceLastAccessedDetails operation. You can use the JobId parameter in
// GetServiceLastAccessedDetails to retrieve the status of your report job. When
// the report is complete, you can retrieve the generated report. The report
// includes a list of AWS services that the resource (user, group, role, or managed
// policy) can access. Service last accessed data does not use other policy types
// when determining whether a resource could access a service. These other policy
// types include resource-based policies, access control lists, AWS Organizations
// policies, IAM permissions boundaries, and AWS STS assume role policies. It only
// applies permissions policy logic. For more about the evaluation of policy types,
// see Evaluating Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-basics)
// in the IAM User Guide. For each service that the resource could access using
// permissions policies, the operation returns details about the most recent access
// attempt. If there was no attempt, the service is listed without details about
// the most recent attempt to access the service. If the operation fails, the
// GetServiceLastAccessedDetails operation returns the reason that it failed. The
// GetServiceLastAccessedDetails operation returns a list of services. This list
// includes the number of entities that have attempted to access the service and
// the date and time of the last attempt. It also returns the ARN of the following
// entity, depending on the resource ARN that you used to generate the report:
//
//
// * User – Returns the user ARN that you used to generate the report
//
//     * Group
// – Returns the ARN of the group member (user) that last attempted to access the
// service
//
//     * Role – Returns the role ARN that you used to generate the
// report
//
//     * Policy – Returns the ARN of the user or role that last used the
// policy to attempt to access the service
//
// By default, the list is sorted by
// service namespace. If you specified ACTION_LEVEL granularity when you generated
// the report, this operation returns service and action last accessed data. This
// includes the most recent access attempt for each tracked action within a
// service. Otherwise, this operation returns only service data. For more
// information about service and action last accessed data, see Reducing
// Permissions Using Service Last Accessed Data
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html)
// in the IAM User Guide.
func (c *Client) GetServiceLastAccessedDetails(ctx context.Context, params *GetServiceLastAccessedDetailsInput, optFns ...func(*Options)) (*GetServiceLastAccessedDetailsOutput, error) {
	stack := middleware.NewStack("GetServiceLastAccessedDetails", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetServiceLastAccessedDetailsMiddlewares(stack)
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
	addOpGetServiceLastAccessedDetailsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetServiceLastAccessedDetails(options.Region), middleware.Before)
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
			OperationName: "GetServiceLastAccessedDetails",
			Err:           err,
		}
	}
	out := result.(*GetServiceLastAccessedDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceLastAccessedDetailsInput struct {
	// The ID of the request generated by the GenerateServiceLastAccessedDetails ()
	// operation. The JobId returned by GenerateServiceLastAccessedDetail must be used
	// by the same role within a session, or by the same user when used to call
	// GetServiceLastAccessedDetail.
	JobId *string
	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true. If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true, and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	MaxItems *int32
	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string
}

type GetServiceLastAccessedDetailsOutput struct {
	// The date and time, in ISO 8601 date-time format
	// (http://www.iso.org/iso/iso8601), when the report job was created.
	JobCreationDate *time.Time
	// A ServiceLastAccessed object that contains details about the most recent attempt
	// to access the service.
	ServicesLastAccessed []*types.ServiceLastAccessed
	// The status of the job.
	JobStatus types.JobStatusType
	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you receive
	// all your results.
	IsTruncated *bool
	// The type of job. Service jobs return information about when each service was
	// last accessed. Action jobs also include information about when tracked actions
	// within the service were last accessed.
	JobType types.AccessAdvisorUsageGranularityType
	// When IsTruncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string
	// The date and time, in ISO 8601 date-time format
	// (http://www.iso.org/iso/iso8601), when the generated report job was completed or
	// failed. This field is null if the job is still in progress, as indicated by a
	// job status value of IN_PROGRESS.
	JobCompletionDate *time.Time
	// An object that contains details about the reason the operation failed.
	Error *types.ErrorDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetServiceLastAccessedDetailsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetServiceLastAccessedDetails{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetServiceLastAccessedDetails{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetServiceLastAccessedDetails(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GetServiceLastAccessedDetails",
	}
}
