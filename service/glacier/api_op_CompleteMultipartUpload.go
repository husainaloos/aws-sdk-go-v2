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

// You call this operation to inform Amazon S3 Glacier (Glacier) that all the
// archive parts have been uploaded and that Glacier can now assemble the archive
// from the uploaded parts. After assembling and saving the archive to the vault,
// Glacier returns the URI path of the newly created archive resource. Using the
// URI path, you can then access the archive. After you upload an archive, you
// should save the archive ID returned to retrieve the archive at a later point.
// You can also get the vault inventory to obtain a list of archive IDs in a vault.
// For more information, see InitiateJob ().  <p>In the request, you must include
// the computed SHA256 tree hash of the entire archive you have uploaded. For
// information about computing a SHA256 tree hash, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/checksum-calculations.html">Computing
// Checksums</a>. On the server side, Glacier also constructs the SHA256 tree hash
// of the assembled archive. If the values match, Glacier saves the archive to the
// vault; otherwise, it returns an error, and the operation fails. The
// <a>ListParts</a> operation returns a list of parts uploaded for a specific
// multipart upload. It includes checksum information for each uploaded part that
// can be used to debug a bad checksum issue.</p> <p>Additionally, Glacier also
// checks for any missing content ranges when assembling the archive, if missing
// content ranges are found, Glacier returns an error and the operation fails.</p>
// <p>Complete Multipart Upload is an idempotent operation. After your first
// successful complete multipart upload, if you call the operation again within a
// short period, the operation will succeed and return the same archive ID. This is
// useful in the event you experience a network issue that causes an aborted
// connection or receive a 500 server error, in which case you can repeat your
// Complete Multipart Upload request and get the same archive ID without creating
// duplicate archives. Note, however, that after the multipart upload completes,
// you cannot call the List Parts operation and the multipart upload will not
// appear in List Multipart Uploads response, even if idempotent complete is
// possible.</p> <p>An AWS account has full permission to perform all operations
// (actions). However, AWS Identity and Access Management (IAM) users don't have
// any permissions by default. You must grant them explicit permission to perform
// specific actions. For more information, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html">Access
// Control Using AWS Identity and Access Management (IAM)</a>.</p> <p> For
// conceptual information and underlying REST API, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/uploading-archive-mpu.html">Uploading
// Large Archives in Parts (Multipart Upload)</a> and <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/api-multipart-complete-upload.html">Complete
// Multipart Upload</a> in the <i>Amazon Glacier Developer Guide</i>. </p>
func (c *Client) CompleteMultipartUpload(ctx context.Context, params *CompleteMultipartUploadInput, optFns ...func(*Options)) (*CompleteMultipartUploadOutput, error) {
	stack := middleware.NewStack("CompleteMultipartUpload", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCompleteMultipartUploadMiddlewares(stack)
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
	addOpCompleteMultipartUploadValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCompleteMultipartUpload(options.Region), middleware.Before)
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
			OperationName: "CompleteMultipartUpload",
			Err:           err,
		}
	}
	out := result.(*CompleteMultipartUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options to complete a multipart upload operation. This informs Amazon
// Glacier that all the archive parts have been uploaded and Amazon S3 Glacier
// (Glacier) can now assemble the archive from the uploaded parts. After assembling
// and saving the archive to the vault, Glacier returns the URI path of the newly
// created archive resource.
type CompleteMultipartUploadInput struct {
	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	AccountId *string
	// The upload ID of the multipart upload.
	UploadId *string
	// The name of the vault.
	VaultName *string
	// The total size, in bytes, of the entire archive. This value should be the sum of
	// all the sizes of the individual parts that you uploaded.
	ArchiveSize *string
	// The SHA256 tree hash of the entire archive. It is the tree hash of SHA256 tree
	// hash of the individual parts. If the value you specify in the request does not
	// match the SHA256 tree hash of the final assembled archive as computed by Amazon
	// S3 Glacier (Glacier), Glacier returns an error and the request fails.
	Checksum *string
}

// Contains the Amazon S3 Glacier response to your request. For information about
// the underlying REST API, see Upload Archive
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-archive-post.html).
// For conceptual information, see Working with Archives in Amazon S3 Glacier
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-archives.html).
type CompleteMultipartUploadOutput struct {
	// The checksum of the archive computed by Amazon S3 Glacier.
	Checksum *string
	// The ID of the archive. This value is also included as part of the location.
	ArchiveId *string
	// The relative URI path of the newly added archive resource.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCompleteMultipartUploadMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCompleteMultipartUpload{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCompleteMultipartUpload{}, middleware.After)
}

func newServiceMetadataMiddleware_opCompleteMultipartUpload(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "CompleteMultipartUpload",
	}
}
