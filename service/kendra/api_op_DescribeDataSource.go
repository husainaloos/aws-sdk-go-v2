// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets information about a Amazon Kendra data source.
func (c *Client) DescribeDataSource(ctx context.Context, params *DescribeDataSourceInput, optFns ...func(*Options)) (*DescribeDataSourceOutput, error) {
	stack := middleware.NewStack("DescribeDataSource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeDataSourceMiddlewares(stack)
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
	addOpDescribeDataSourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDataSource(options.Region), middleware.Before)
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
			OperationName: "DescribeDataSource",
			Err:           err,
		}
	}
	out := result.(*DescribeDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDataSourceInput struct {
	// The unique identifier of the data source to describe.
	Id *string
	// The identifier of the index that contains the data source.
	IndexId *string
}

type DescribeDataSourceOutput struct {
	// The identifier of the index that contains the data source.
	IndexId *string
	// The current status of the data source. When the status is ACTIVE the data source
	// is ready to use. When the status is FAILED, the ErrorMessage field contains the
	// reason that the data source failed.
	Status types.DataSourceStatus
	// The description of the data source.
	Description *string
	// The identifier of the data source.
	Id *string
	// The schedule that Amazon Kendra will update the data source.
	Schedule *string
	// The Unix timestamp of when the data source was created.
	CreatedAt *time.Time
	// Information that describes where the data source is located and how the data
	// source is configured. The specific information in the description depends on the
	// data source provider.
	Configuration *types.DataSourceConfiguration
	// When the Status field value is FAILED, the ErrorMessage field contains a
	// description of the error that caused the data source to fail.
	ErrorMessage *string
	// The Unix timestamp of when the data source was last updated.
	UpdatedAt *time.Time
	// The name that you gave the data source when it was created.
	Name *string
	// The type of the data source.
	Type types.DataSourceType
	// The Amazon Resource Name (ARN) of the role that enables the data source to
	// access its resources.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeDataSourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDataSource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDataSource{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDataSource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "DescribeDataSource",
	}
}
