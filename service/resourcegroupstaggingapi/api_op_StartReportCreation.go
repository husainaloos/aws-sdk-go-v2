// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Generates a report that lists all tagged resources in accounts across your
// organization and tells whether each resource is compliant with the effective tag
// policy. Compliance data is refreshed daily. The generated report is saved to the
// following location:
// s3://example-bucket/AwsTagPolicies/o-exampleorgid/YYYY-MM-ddTHH:mm:ssZ/report.csv
// You can call this operation only from the organization's master account and from
// the us-east-1 Region.
func (c *Client) StartReportCreation(ctx context.Context, params *StartReportCreationInput, optFns ...func(*Options)) (*StartReportCreationOutput, error) {
	stack := middleware.NewStack("StartReportCreation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartReportCreationMiddlewares(stack)
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
	addOpStartReportCreationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartReportCreation(options.Region), middleware.Before)
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
			OperationName: "StartReportCreation",
			Err:           err,
		}
	}
	out := result.(*StartReportCreationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartReportCreationInput struct {
	// The name of the Amazon S3 bucket where the report will be stored; for example:
	// awsexamplebucket For more information on S3 bucket requirements, including an
	// example bucket policy, see the example S3 bucket policy on this page.
	S3Bucket *string
}

type StartReportCreationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartReportCreationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartReportCreation{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartReportCreation{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartReportCreation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tagging",
		OperationName: "StartReportCreation",
	}
}
