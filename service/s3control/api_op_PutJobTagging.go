// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Set the supplied tag-set on an Amazon S3 Batch Operations job. A tag is a
// key-value pair. You can associate Amazon S3 Batch Operations tags with any job
// by sending a PUT request against the tagging subresource that is associated with
// the job. To modify the existing tag set, you can either replace the existing tag
// set entirely, or make changes within the existing tag set by retrieving the
// existing tag set using GetJobTagging (), modify that tag set, and use this API
// action to replace the tag set with the one you have modified.. For more
// information, see Using Job Tags
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-managing-jobs.html#batch-ops-job-tags)
// in the Amazon Simple Storage Service Developer Guide.
//
//     * If you send this
// request with an empty tag set, Amazon S3 deletes the existing tag set on the
// Batch Operations job. If you use this method, you will be charged for a Tier 1
// Request (PUT). For more information, see Amazon S3 pricing
// (http://aws.amazon.com/s3/pricing/).
//
//     * For deleting existing tags for your
// batch operations job, DeleteJobTagging () request is preferred because it
// achieves the same result without incurring charges.
//
//     * A few things to
// consider about using tags:
//
//         * Amazon S3 limits the maximum number of
// tags to 50 tags per job.
//
//         * You can associate up to 50 tags with a job
// as long as they have unique tag keys.
//
//         * A tag key can be up to 128
// Unicode characters in length, and tag values can be up to 256 Unicode characters
// in length.
//
//         * The key and values are case sensitive.
//
//         * For
// tagging-related restrictions related to characters and encodings, see
// User-Defined Tag Restrictions
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.html).
//
// To
// use this operation, you must have permission to perform the s3:PutJobTagging
// action. Related actions include:
//
//     * CreateJob ()
//
//     * GetJobTagging ()
//
//
// * DeleteJobTagging ()
func (c *Client) PutJobTagging(ctx context.Context, params *PutJobTaggingInput, optFns ...func(*Options)) (*PutJobTaggingOutput, error) {
	stack := middleware.NewStack("PutJobTagging", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpPutJobTaggingMiddlewares(stack)
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
	addOpPutJobTaggingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutJobTagging(options.Region), middleware.Before)
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
			OperationName: "PutJobTagging",
			Err:           err,
		}
	}
	out := result.(*PutJobTaggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutJobTaggingInput struct {
	// The ID for the Amazon S3 Batch Operations job whose tags you want to replace.
	JobId *string
	// The set of tags to associate with the Amazon S3 Batch Operations job.
	Tags []*types.S3Tag
	// The AWS account ID associated with the Amazon S3 Batch Operations job.
	AccountId *string
}

type PutJobTaggingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpPutJobTaggingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpPutJobTagging{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpPutJobTagging{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutJobTagging(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutJobTagging",
	}
}
