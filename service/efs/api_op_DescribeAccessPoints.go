// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the description of a specific Amazon EFS access point if the
// AccessPointId is provided. If you provide an EFS FileSystemId, it returns
// descriptions of all access points for that file system. You can provide either
// an AccessPointId or a FileSystemId in the request, but not both. This operation
// requires permissions for the elasticfilesystem:DescribeAccessPoints action.
func (c *Client) DescribeAccessPoints(ctx context.Context, params *DescribeAccessPointsInput, optFns ...func(*Options)) (*DescribeAccessPointsOutput, error) {
	stack := middleware.NewStack("DescribeAccessPoints", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeAccessPointsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccessPoints(options.Region), middleware.Before)
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
			OperationName: "DescribeAccessPoints",
			Err:           err,
		}
	}
	out := result.(*DescribeAccessPointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccessPointsInput struct {
	// (Optional) When retrieving all access points for a file system, you can
	// optionally specify the MaxItems parameter to limit the number of objects
	// returned in a response. The default value is 100.
	MaxResults *int32
	// NextToken is present if the response is paginated. You can use NextMarker in the
	// subsequent request to fetch the next page of access point descriptions.
	NextToken *string
	// (Optional) Specifies an EFS access point to describe in the response; mutually
	// exclusive with FileSystemId.
	AccessPointId *string
	// (Optional) If you provide a FileSystemId, EFS returns all access points for that
	// file system; mutually exclusive with AccessPointId.
	FileSystemId *string
}

type DescribeAccessPointsOutput struct {
	// An array of access point descriptions.
	AccessPoints []*types.AccessPointDescription
	// Present if there are more access points than returned in the response. You can
	// use the NextMarker in the subsequent request to fetch the additional
	// descriptions.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeAccessPointsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAccessPoints{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAccessPoints{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAccessPoints(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "DescribeAccessPoints",
	}
}
