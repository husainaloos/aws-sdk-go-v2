// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieve a list of application versions.
func (c *Client) DescribeApplicationVersions(ctx context.Context, params *DescribeApplicationVersionsInput, optFns ...func(*Options)) (*DescribeApplicationVersionsOutput, error) {
	stack := middleware.NewStack("DescribeApplicationVersions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeApplicationVersionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeApplicationVersions(options.Region), middleware.Before)
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
			OperationName: "DescribeApplicationVersions",
			Err:           err,
		}
	}
	out := result.(*DescribeApplicationVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to describe application versions.
type DescribeApplicationVersionsInput struct {
	// Specify an application name to show only application versions for that
	// application.
	ApplicationName *string
	// For a paginated request. Specify a token from a previous response page to
	// retrieve the next response page. All other parameter values must be identical to
	// the ones specified in the initial request. If no NextToken is specified, the
	// first page is retrieved.
	NextToken *string
	// For a paginated request. Specify a maximum number of application versions to
	// include in each response. If no MaxRecords is specified, all available
	// application versions are retrieved in a single response.
	MaxRecords *int32
	// Specify a version label to show a specific application version.
	VersionLabels []*string
}

// Result message wrapping a list of application version descriptions.
type DescribeApplicationVersionsOutput struct {
	// List of ApplicationVersionDescription objects sorted in order of creation.
	ApplicationVersions []*types.ApplicationVersionDescription
	// In a paginated request, the token that you can pass in a subsequent request to
	// get the next response page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeApplicationVersionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeApplicationVersions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeApplicationVersions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeApplicationVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DescribeApplicationVersions",
	}
}
