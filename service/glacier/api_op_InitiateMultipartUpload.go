// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation initiates a multipart upload. Amazon S3 Glacier creates a
// multipart upload resource and returns its ID in the response. The multipart
// upload ID is used in subsequent requests to upload parts of an archive (see
// UploadMultipartPart ()).  <p>When you initiate a multipart upload, you specify
// the part size in number of bytes. The part size must be a megabyte (1024 KB)
// multiplied by a power of 2-for example, 1048576 (1 MB), 2097152 (2 MB), 4194304
// (4 MB), 8388608 (8 MB), and so on. The minimum allowable part size is 1 MB, and
// the maximum is 4 GB.</p> <p>Every part you upload to this resource (see
// <a>UploadMultipartPart</a>), except the last one, must have the same size. The
// last one can be the same size or smaller. For example, suppose you want to
// upload a 16.2 MB file. If you initiate the multipart upload with a part size of
// 4 MB, you will upload four parts of 4 MB each and one part of 0.2 MB. </p>
// <note> <p>You don't need to know the size of the archive when you start a
// multipart upload because Amazon S3 Glacier does not require you to specify the
// overall archive size.</p> </note> <p>After you complete the multipart upload,
// Amazon S3 Glacier (Glacier) removes the multipart upload resource referenced by
// the ID. Glacier also removes the multipart upload resource if you cancel the
// multipart upload or it may be removed if there is no activity for a period of 24
// hours.</p> <p>An AWS account has full permission to perform all operations
// (actions). However, AWS Identity and Access Management (IAM) users don't have
// any permissions by default. You must grant them explicit permission to perform
// specific actions. For more information, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html">Access
// Control Using AWS Identity and Access Management (IAM)</a>.</p> <p>For
// conceptual information and underlying REST API, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/uploading-archive-mpu.html">Uploading
// Large Archives in Parts (Multipart Upload)</a> and <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/api-multipart-initiate-upload.html">Initiate
// Multipart Upload</a> in the <i>Amazon Glacier Developer Guide</i>.</p>
func (c *Client) InitiateMultipartUpload(ctx context.Context, params *InitiateMultipartUploadInput, optFns ...func(*Options)) (*InitiateMultipartUploadOutput, error) {
	stack := middleware.NewStack("InitiateMultipartUpload", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpInitiateMultipartUploadMiddlewares(stack)
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
	addOpInitiateMultipartUploadValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opInitiateMultipartUpload(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	glaciercust.AddTreeHashMiddleware(stack)

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
			OperationName: "InitiateMultipartUpload",
			Err:           err,
		}
	}
	out := result.(*InitiateMultipartUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options for initiating a multipart upload to an Amazon S3 Glacier
// vault.
type InitiateMultipartUploadInput struct {
	// The size of each part except the last, in bytes. The last part can be smaller
	// than this part size.
	PartSize *string
	// The name of the vault.
	VaultName *string
	// The archive description that you are uploading in parts. The part size must be a
	// megabyte (1024 KB) multiplied by a power of 2, for example 1048576 (1 MB),
	// 2097152 (2 MB), 4194304 (4 MB), 8388608 (8 MB), and so on. The minimum allowable
	// part size is 1 MB, and the maximum is 4 GB (4096 MB).
	ArchiveDescription *string
	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	AccountId *string
}

// The Amazon S3 Glacier response to your request.
type InitiateMultipartUploadOutput struct {
	// The relative URI path of the multipart upload ID Amazon S3 Glacier created.
	Location *string
	// The ID of the multipart upload. This value is also included as part of the
	// location.
	UploadId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpInitiateMultipartUploadMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpInitiateMultipartUpload{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpInitiateMultipartUpload{}, middleware.After)
}

func newServiceMetadataMiddleware_opInitiateMultipartUpload(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "InitiateMultipartUpload",
	}
}
