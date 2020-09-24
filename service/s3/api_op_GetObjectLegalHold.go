// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets an object's current Legal Hold status. For more information, see Locking
// Objects (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html).
func (c *Client) GetObjectLegalHold(ctx context.Context, params *GetObjectLegalHoldInput, optFns ...func(*Options)) (*GetObjectLegalHoldOutput, error) {
	stack := middleware.NewStack("GetObjectLegalHold", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetObjectLegalHoldMiddlewares(stack)
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
	addOpGetObjectLegalHoldValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetObjectLegalHold(options.Region), middleware.Before)
	addUpdateEndpointMiddleware(stack, options)
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
			OperationName: "GetObjectLegalHold",
			Err:           err,
		}
	}
	out := result.(*GetObjectLegalHoldOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetObjectLegalHoldInput struct {
	// The key name for the object whose Legal Hold status you want to retrieve.
	Key *string
	// The bucket name containing the object whose Legal Hold status you want to
	// retrieve. When using this API with an access point, you must direct requests to
	// the access point hostname. The access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// operation using an access point through the AWS SDKs, you provide the access
	// point ARN in place of the bucket name. For more information about access point
	// ARNs, see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) in
	// the Amazon Simple Storage Service Developer Guide.
	Bucket *string
	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer
	// The version ID of the object whose Legal Hold status you want to retrieve.
	VersionId *string
}

type GetObjectLegalHoldOutput struct {
	// The current Legal Hold status for the specified object.
	LegalHold *types.ObjectLockLegalHold

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetObjectLegalHoldMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetObjectLegalHold{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetObjectLegalHold{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetObjectLegalHold(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetObjectLegalHold",
	}
}
