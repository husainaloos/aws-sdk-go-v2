// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This operation initiates a multipart upload and returns an upload ID. This
// upload ID is used to associate all of the parts in the specific multipart
// upload. You specify this upload ID in each of your subsequent upload part
// requests (see UploadPart ()). You also include this upload ID in the final
// request to either complete or abort the multipart upload request.  <p>For more
// information about multipart uploads, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html">Multipart
// Upload Overview</a>.</p> <p>If you have configured a lifecycle rule to abort
// incomplete multipart uploads, the upload must complete within the number of days
// specified in the bucket lifecycle configuration. Otherwise, the incomplete
// multipart upload becomes eligible for an abort operation and Amazon S3 aborts
// the multipart upload. For more information, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html#mpu-abort-incomplete-mpu-lifecycle-config">Aborting
// Incomplete Multipart Uploads Using a Bucket Lifecycle Policy</a>.</p> <p>For
// information about the permissions required to use the multipart upload API, see
// <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html">Multipart
// Upload API and Permissions</a>.</p> <p>For request signing, multipart upload is
// just a series of regular requests. You initiate a multipart upload, send one or
// more requests to upload parts, and then complete the multipart upload process.
// You sign each request individually. There is nothing special about signing
// multipart upload requests. For more information about signing, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html">Authenticating
// Requests (AWS Signature Version 4)</a>.</p> <note> <p> After you initiate a
// multipart upload and upload one or more parts, to stop being charged for storing
// the uploaded parts, you must either complete or abort the multipart upload.
// Amazon S3 frees up the space used to store the parts and stop charging you for
// storing them only after you either complete or abort a multipart upload. </p>
// </note> <p>You can optionally request server-side encryption. For server-side
// encryption, Amazon S3 encrypts your data as it writes it to disks in its data
// centers and decrypts it when you access it. You can provide your own encryption
// key, or use AWS Key Management Service (AWS KMS) customer master keys (CMKs) or
// Amazon S3-managed encryption keys. If you choose to provide your own encryption
// key, the request headers you provide in <a>UploadPart</a>) and
// <a>UploadPartCopy</a>) requests must match the headers you used in the request
// to initiate the upload by using <code>CreateMultipartUpload</code>. </p> <p>To
// perform a multipart upload with encryption using an AWS KMS CMK, the requester
// must have permission to the <code>kms:Encrypt</code>, <code>kms:Decrypt</code>,
// <code>kms:ReEncrypt*</code>, <code>kms:GenerateDataKey*</code>, and
// <code>kms:DescribeKey</code> actions on the key. These permissions are required
// because Amazon S3 must decrypt and read data from the encrypted file parts
// before it completes the multipart upload.</p> <p>If your AWS Identity and Access
// Management (IAM) user or role is in the same AWS account as the AWS KMS CMK,
// then you must have these permissions on the key policy. If your IAM user or role
// belongs to a different account than the key, then you must have the permissions
// on both the key policy and your IAM user or role.</p> <p> For more information,
// see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/serv-side-encryption.html">Protecting
// Data Using Server-Side Encryption</a>.</p> <dl> <dt>Access Permissions</dt> <dd>
// <p>When copying an object, you can optionally specify the accounts or groups
// that should be granted specific permissions on the new object. There are two
// ways to grant the permissions using the request headers:</p> <ul> <li>
// <p>Specify a canned ACL with the <code>x-amz-acl</code> request header. For more
// information, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL">Canned
// ACL</a>.</p> </li> <li> <p>Specify access permissions explicitly with the
// <code>x-amz-grant-read</code>, <code>x-amz-grant-read-acp</code>,
// <code>x-amz-grant-write-acp</code>, and <code>x-amz-grant-full-control</code>
// headers. These parameters map to the set of permissions that Amazon S3 supports
// in an ACL. For more information, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html">Access
// Control List (ACL) Overview</a>.</p> </li> </ul> <p>You can use either a canned
// ACL or specify access permissions explicitly. You cannot do both.</p> </dd>
// <dt>Server-Side- Encryption-Specific Request Headers</dt> <dd> <p>You can
// optionally tell Amazon S3 to encrypt data at rest using server-side encryption.
// Server-side encryption is for data encryption at rest. Amazon S3 encrypts your
// data as it writes it to disks in its data centers and decrypts it when you
// access it. The option you use depends on whether you want to use AWS managed
// encryption keys or provide your own encryption key. </p> <ul> <li> <p>Use
// encryption keys managed by Amazon S3 or customer master keys (CMKs) stored in
// AWS Key Management Service (AWS KMS) – If you want AWS to manage the keys used
// to encrypt data, specify the following headers in the request.</p> <ul> <li>
// <p>x-amz-server-side-encryption</p> </li> <li>
// <p>x-amz-server-side-encryption-aws-kms-key-id</p> </li> <li>
// <p>x-amz-server-side-encryption-context</p> </li> </ul> <note> <p>If you specify
// <code>x-amz-server-side-encryption:aws:kms</code>, but don't provide
// <code>x-amz-server-side-encryption-aws-kms-key-id</code>, Amazon S3 uses the AWS
// managed CMK in AWS KMS to protect the data.</p> </note> <important> <p>All GET
// and PUT requests for an object protected by AWS KMS fail if you don't make them
// with SSL or by using SigV4.</p> </important> <p>For more information about
// server-side encryption with CMKs stored in AWS KMS (SSE-KMS), see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html">Protecting
// Data Using Server-Side Encryption with CMKs stored in AWS KMS</a>.</p> </li>
// <li> <p>Use customer-provided encryption keys – If you want to manage your own
// encryption keys, provide all the following headers in the request.</p> <ul> <li>
// <p>x-amz-server-side-encryption-customer-algorithm</p> </li> <li>
// <p>x-amz-server-side-encryption-customer-key</p> </li> <li>
// <p>x-amz-server-side-encryption-customer-key-MD5</p> </li> </ul> <p>For more
// information about server-side encryption with CMKs stored in AWS KMS (SSE-KMS),
// see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html">Protecting
// Data Using Server-Side Encryption with CMKs stored in AWS KMS</a>.</p> </li>
// </ul> </dd> <dt>Access-Control-List (ACL)-Specific Request Headers</dt> <dd>
// <p>You also can use the following access control–related headers with this
// operation. By default, all objects are private. Only the owner has full access
// control. When adding a new object, you can grant permissions to individual AWS
// accounts or to predefined groups defined by Amazon S3. These permissions are
// then added to the access control list (ACL) on the object. For more information,
// see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/S3_ACLs_UsingACLs.html">Using
// ACLs</a>. With this operation, you can grant access permissions using one of the
// following two methods:</p> <ul> <li> <p>Specify a canned ACL
// (<code>x-amz-acl</code>) — Amazon S3 supports a set of predefined ACLs, known as
// <i>canned ACLs</i>. Each canned ACL has a predefined set of grantees and
// permissions. For more information, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL">Canned
// ACL</a>.</p> </li> <li> <p>Specify access permissions explicitly — To explicitly
// grant access permissions to specific AWS accounts or groups, use the following
// headers. Each header maps to specific permissions that Amazon S3 supports in an
// ACL. For more information, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html">Access
// Control List (ACL) Overview</a>. In the header, you specify a list of grantees
// who get the specific permission. To grant permissions explicitly, use:</p> <ul>
// <li> <p>x-amz-grant-read</p> </li> <li> <p>x-amz-grant-write</p> </li> <li>
// <p>x-amz-grant-read-acp</p> </li> <li> <p>x-amz-grant-write-acp</p> </li> <li>
// <p>x-amz-grant-full-control</p> </li> </ul> <p>You specify each grantee as a
// type=value pair, where the type is one of the following:</p> <ul> <li> <p>
// <code>id</code> – if the value specified is the canonical user ID of an AWS
// account</p> </li> <li> <p> <code>uri</code> – if you are granting permissions to
// a predefined group</p> </li> <li> <p> <code>emailAddress</code> – if the value
// specified is the email address of an AWS account</p> <note> <p>Using email
// addresses to specify a grantee is only supported in the following AWS Regions:
// </p> <ul> <li> <p>US East (N. Virginia)</p> </li> <li> <p>US West (N.
// California)</p> </li> <li> <p> US West (Oregon)</p> </li> <li> <p> Asia Pacific
// (Singapore)</p> </li> <li> <p>Asia Pacific (Sydney)</p> </li> <li> <p>Asia
// Pacific (Tokyo)</p> </li> <li> <p>Europe (Ireland)</p> </li> <li> <p>South
// America (São Paulo)</p> </li> </ul> <p>For a list of all the Amazon S3 supported
// Regions and endpoints, see <a
// href="https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region">Regions
// and Endpoints</a> in the AWS General Reference.</p> </note> </li> </ul> <p>For
// example, the following <code>x-amz-grant-read</code> header grants the AWS
// accounts identified by account IDs permissions to read object data and its
// metadata:</p> <p> <code>x-amz-grant-read: id="11112222333", id="444455556666"
// </code> </p> </li> </ul> </dd> </dl> <p>The following operations are related to
// <code>CreateMultipartUpload</code>:</p> <ul> <li> <p> <a>UploadPart</a> </p>
// </li> <li> <p> <a>CompleteMultipartUpload</a> </p> </li> <li> <p>
// <a>AbortMultipartUpload</a> </p> </li> <li> <p> <a>ListParts</a> </p> </li> <li>
// <p> <a>ListMultipartUploads</a> </p> </li> </ul>
func (c *Client) CreateMultipartUpload(ctx context.Context, params *CreateMultipartUploadInput, optFns ...func(*Options)) (*CreateMultipartUploadOutput, error) {
	if params == nil {
		params = &CreateMultipartUploadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMultipartUpload", params, optFns, addOperationCreateMultipartUploadMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMultipartUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMultipartUploadInput struct {

	// The name of the bucket to which to initiate the upload
	//
	// This member is required.
	Bucket *string

	// Object key for which the multipart upload is to be initiated.
	//
	// This member is required.
	Key *string

	// The canned ACL to apply to the object.
	ACL types.ObjectCannedACL

	// Specifies caching behavior along the request/reply chain.
	CacheControl *string

	// Specifies presentational information for the object.
	ContentDisposition *string

	// Specifies what content encodings have been applied to the object and thus what
	// decoding mechanisms must be applied to obtain the media-type referenced by the
	// Content-Type header field.
	ContentEncoding *string

	// The language the content is in.
	ContentLanguage *string

	// A standard MIME type describing the format of the object data.
	ContentType *string

	// The date and time at which the object is no longer cacheable.
	Expires *time.Time

	// Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.
	GrantFullControl *string

	// Allows grantee to read the object data and its metadata.
	GrantRead *string

	// Allows grantee to read the object ACL.
	GrantReadACP *string

	// Allows grantee to write the ACL for the applicable object.
	GrantWriteACP *string

	// A map of metadata to store with the object in S3.
	Metadata map[string]*string

	// Specifies whether you want to apply a Legal Hold to the uploaded object.
	ObjectLockLegalHoldStatus types.ObjectLockLegalHoldStatus

	// Specifies the Object Lock mode that you want to apply to the uploaded object.
	ObjectLockMode types.ObjectLockMode

	// Specifies the date and time when you want the Object Lock to expire.
	ObjectLockRetainUntilDate *time.Time

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer

	// Specifies the algorithm to use to when encrypting the object (for example,
	// AES256).
	SSECustomerAlgorithm *string

	// Specifies the customer-provided encryption key for Amazon S3 to use in
	// encrypting data. This value is used to store the object and then it is
	// discarded; Amazon S3 does not store the encryption key. The key must be
	// appropriate for use with the algorithm specified in the
	// x-amz-server-side-encryption-customer-algorithm header.
	SSECustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	SSECustomerKeyMD5 *string

	// Specifies the AWS KMS Encryption Context to use for object encryption. The value
	// of this header is a base64-encoded UTF-8 string holding JSON with the encryption
	// context key-value pairs.
	SSEKMSEncryptionContext *string

	// Specifies the ID of the symmetric customer managed AWS KMS CMK to use for object
	// encryption. All GET and PUT requests for an object protected by AWS KMS will
	// fail if not made via SSL or using SigV4. For information about configuring using
	// any of the officially supported AWS SDKs and AWS CLI, see Specifying the
	// Signature Version in Request Authentication
	// (https://docs.aws.amazon.com/http:/docs.aws.amazon.com/AmazonS3/latest/dev/UsingAWSSDK.html#specify-signature-version)
	// in the Amazon S3 Developer Guide.
	SSEKMSKeyId *string

	// The server-side encryption algorithm used when storing this object in Amazon S3
	// (for example, AES256, aws:kms).
	ServerSideEncryption types.ServerSideEncryption

	// The type of storage to use for the object. Defaults to 'STANDARD'.
	StorageClass types.StorageClass

	// The tag-set for the object. The tag-set must be encoded as URL Query parameters.
	Tagging *string

	// If the bucket is configured as a website, redirects requests for this object to
	// another object in the same bucket or to an external URL. Amazon S3 stores the
	// value of this header in the object metadata.
	WebsiteRedirectLocation *string
}

type CreateMultipartUploadOutput struct {

	// If the bucket has a lifecycle rule configured with an action to abort incomplete
	// multipart uploads and the prefix in the lifecycle rule matches the object name
	// in the request, the response includes this header. The header indicates when the
	// initiated multipart upload becomes eligible for an abort operation. For more
	// information, see  Aborting Incomplete Multipart Uploads Using a Bucket Lifecycle
	// Policy
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html#mpu-abort-incomplete-mpu-lifecycle-config).
	// <p>The response also includes the <code>x-amz-abort-rule-id</code> header that
	// provides the ID of the lifecycle configuration rule that defines this
	// action.</p>
	AbortDate *time.Time

	// This header is returned along with the x-amz-abort-date header. It identifies
	// the applicable lifecycle configuration rule that defines the action to abort
	// incomplete multipart uploads.
	AbortRuleId *string

	// Name of the bucket to which the multipart upload was initiated. When using this
	// API with an access point, you must direct requests to the access point hostname.
	// The access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// operation using an access point through the AWS SDKs, you provide the access
	// point ARN in place of the bucket name. For more information about access point
	// ARNs, see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) in
	// the Amazon Simple Storage Service Developer Guide.
	Bucket *string

	// Object key for which the multipart upload was initiated.
	Key *string

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm used.
	SSECustomerAlgorithm *string

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round-trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string

	// If present, specifies the AWS KMS Encryption Context to use for object
	// encryption. The value of this header is a base64-encoded UTF-8 string holding
	// JSON with the encryption context key-value pairs.
	SSEKMSEncryptionContext *string

	// If present, specifies the ID of the AWS Key Management Service (AWS KMS)
	// symmetric customer managed customer master key (CMK) that was used for the
	// object.
	SSEKMSKeyId *string

	// The server-side encryption algorithm used when storing this object in Amazon S3
	// (for example, AES256, aws:kms).
	ServerSideEncryption types.ServerSideEncryption

	// ID for the initiated multipart upload.
	UploadId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateMultipartUploadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateMultipartUpload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateMultipartUpload{}, middleware.After)
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
	addOpCreateMultipartUploadValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMultipartUpload(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateMultipartUpload(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "CreateMultipartUpload",
	}
}
