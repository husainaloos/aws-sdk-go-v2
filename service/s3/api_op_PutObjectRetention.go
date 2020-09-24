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

// Places an Object Retention configuration on an object. Related Resources
//
//     *
// Locking Objects
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html)
func (c *Client) PutObjectRetention(ctx context.Context, params *PutObjectRetentionInput, optFns ...func(*Options)) (*PutObjectRetentionOutput, error) {
	stack := middleware.NewStack("PutObjectRetention", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpPutObjectRetentionMiddlewares(stack)
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
	addOpPutObjectRetentionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutObjectRetention(options.Region), middleware.Before)
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
			OperationName: "PutObjectRetention",
			Err:           err,
		}
	}
	out := result.(*PutObjectRetentionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutObjectRetentionInput struct {
	// The version ID for the object that you want to apply this Object Retention
	// configuration to.
	VersionId *string
	// The MD5 hash for the request body.
	ContentMD5 *string
	// The bucket name that contains the object you want to apply this Object Retention
	// configuration to. When using this API with an access point, you must direct
	// requests to the access point hostname. The access point hostname takes the form
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
	// Indicates whether this operation should bypass Governance-mode restrictions.
	BypassGovernanceRetention *bool
	// The key name for the object that you want to apply this Object Retention
	// configuration to.
	Key *string
	// The container element for the Object Retention configuration.
	Retention *types.ObjectLockRetention
}

type PutObjectRetentionOutput struct {
	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpPutObjectRetentionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpPutObjectRetention{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpPutObjectRetention{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutObjectRetention(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutObjectRetention",
	}
}
