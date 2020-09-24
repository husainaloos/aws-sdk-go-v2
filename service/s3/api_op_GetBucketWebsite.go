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

// Returns the website configuration for a bucket. To host website on Amazon S3,
// you can configure a bucket as website by adding a website configuration. For
// more information about hosting websites, see Hosting Websites on Amazon S3
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html). This GET
// operation requires the S3:GetBucketWebsite permission. By default, only the
// bucket owner can read the bucket website configuration. However, bucket owners
// can allow other users to read the website configuration by writing a bucket
// policy granting them the S3:GetBucketWebsite permission. The following
// operations are related to DeleteBucketWebsite:
//
//     * DeleteBucketWebsite ()
//
//
// * PutBucketWebsite ()
func (c *Client) GetBucketWebsite(ctx context.Context, params *GetBucketWebsiteInput, optFns ...func(*Options)) (*GetBucketWebsiteOutput, error) {
	stack := middleware.NewStack("GetBucketWebsite", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetBucketWebsiteMiddlewares(stack)
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
	addOpGetBucketWebsiteValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketWebsite(options.Region), middleware.Before)
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
			OperationName: "GetBucketWebsite",
			Err:           err,
		}
	}
	out := result.(*GetBucketWebsiteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketWebsiteInput struct {
	// The bucket name for which to get the website configuration.
	Bucket *string
}

type GetBucketWebsiteOutput struct {
	// Specifies the redirect behavior of all requests to a website endpoint of an
	// Amazon S3 bucket.
	RedirectAllRequestsTo *types.RedirectAllRequestsTo
	// The object key name of the website error document to use for 4XX class errors.
	ErrorDocument *types.ErrorDocument
	// The name of the index document for the website (for example index.html).
	IndexDocument *types.IndexDocument
	// Rules that define when a redirect is applied and the redirect behavior.
	RoutingRules []*types.RoutingRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetBucketWebsiteMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetBucketWebsite{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketWebsite{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBucketWebsite(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketWebsite",
	}
}
