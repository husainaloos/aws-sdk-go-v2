// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns descriptions for existing environments.
func (c *Client) DescribeEnvironments(ctx context.Context, params *DescribeEnvironmentsInput, optFns ...func(*Options)) (*DescribeEnvironmentsOutput, error) {
	if params == nil {
		params = &DescribeEnvironmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEnvironments", params, optFns, addOperationDescribeEnvironmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEnvironmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to describe one or more environments.
type DescribeEnvironmentsInput struct {

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// include only those that are associated with this application.
	ApplicationName *string

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// include only those that have the specified IDs.
	EnvironmentIds []*string

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// include only those that have the specified names.
	EnvironmentNames []*string

	// Indicates whether to include deleted environments: true: Environments that have
	// been deleted after IncludedDeletedBackTo are displayed. false: Do not include
	// deleted environments.
	IncludeDeleted *bool

	// If specified when IncludeDeleted is set to true, then environments deleted after
	// this date are displayed.
	IncludedDeletedBackTo *time.Time

	// For a paginated request. Specify a maximum number of environments to include in
	// each response. If no MaxRecords is specified, all available environments are
	// retrieved in a single response.
	MaxRecords *int32

	// For a paginated request. Specify a token from a previous response page to
	// retrieve the next response page. All other parameter values must be identical to
	// the ones specified in the initial request. If no NextToken is specified, the
	// first page is retrieved.
	NextToken *string

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// include only those that are associated with this application version.
	VersionLabel *string
}

// Result message containing a list of environment descriptions.
type DescribeEnvironmentsOutput struct {

	// Returns an EnvironmentDescription () list.
	Environments []*types.EnvironmentDescription

	// In a paginated request, the token that you can pass in a subsequent request to
	// get the next response page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEnvironmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEnvironments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEnvironments{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEnvironments(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeEnvironments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DescribeEnvironments",
	}
}
