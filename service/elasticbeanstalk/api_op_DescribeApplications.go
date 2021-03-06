// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the descriptions of existing applications.
func (c *Client) DescribeApplications(ctx context.Context, params *DescribeApplicationsInput, optFns ...func(*Options)) (*DescribeApplicationsOutput, error) {
	if params == nil {
		params = &DescribeApplicationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeApplications", params, optFns, addOperationDescribeApplicationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeApplicationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to describe one or more applications.
type DescribeApplicationsInput struct {

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to only
	// include those with the specified names.
	ApplicationNames []*string
}

// Result message containing a list of application descriptions.
type DescribeApplicationsOutput struct {

	// This parameter contains a list of ApplicationDescription ().
	Applications []*types.ApplicationDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeApplicationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeApplications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeApplications{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeApplications(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeApplications(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DescribeApplications",
	}
}
