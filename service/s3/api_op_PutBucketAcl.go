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

// Sets the permissions on an existing bucket using access control lists (ACL). For
// more information, see Using ACLs
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3_ACLs_UsingACLs.html). To set
// the ACL of a bucket, you must have WRITE_ACP permission.  <p>You can use one of
// the following two ways to set a bucket's permissions:</p> <ul> <li> <p>Specify
// the ACL in the request body</p> </li> <li> <p>Specify permissions using request
// headers</p> </li> </ul> <note> <p>You cannot specify access permission using
// both the body and the request headers.</p> </note> <p>Depending on your
// application needs, you may choose to set the ACL on a bucket using either the
// request body or the headers. For example, if you have an existing application
// that updates a bucket ACL using the request body, then you can continue to use
// that approach.</p> <p> <b>Access Permissions</b> </p> <p>You can set access
// permissions using one of the following methods:</p> <ul> <li> <p>Specify a
// canned ACL with the <code>x-amz-acl</code> request header. Amazon S3 supports a
// set of predefined ACLs, known as <i>canned ACLs</i>. Each canned ACL has a
// predefined set of grantees and permissions. Specify the canned ACL name as the
// value of <code>x-amz-acl</code>. If you use this header, you cannot use other
// access control-specific headers in your request. For more information, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL">Canned
// ACL</a>.</p> </li> <li> <p>Specify access permissions explicitly with the
// <code>x-amz-grant-read</code>, <code>x-amz-grant-read-acp</code>,
// <code>x-amz-grant-write-acp</code>, and <code>x-amz-grant-full-control</code>
// headers. When using these headers, you specify explicit access permissions and
// grantees (AWS accounts or Amazon S3 groups) who will receive the permission. If
// you use these ACL-specific headers, you cannot use the <code>x-amz-acl</code>
// header to set a canned ACL. These parameters map to the set of permissions that
// Amazon S3 supports in an ACL. For more information, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html">Access
// Control List (ACL) Overview</a>.</p> <p>You specify each grantee as a type=value
// pair, where the type is one of the following:</p> <ul> <li> <p> <code>id</code>
// – if the value specified is the canonical user ID of an AWS account</p> </li>
// <li> <p> <code>uri</code> – if you are granting permissions to a predefined
// group</p> </li> <li> <p> <code>emailAddress</code> – if the value specified is
// the email address of an AWS account</p> <note> <p>Using email addresses to
// specify a grantee is only supported in the following AWS Regions: </p> <ul> <li>
// <p>US East (N. Virginia)</p> </li> <li> <p>US West (N. California)</p> </li>
// <li> <p> US West (Oregon)</p> </li> <li> <p> Asia Pacific (Singapore)</p> </li>
// <li> <p>Asia Pacific (Sydney)</p> </li> <li> <p>Asia Pacific (Tokyo)</p> </li>
// <li> <p>Europe (Ireland)</p> </li> <li> <p>South America (São Paulo)</p> </li>
// </ul> <p>For a list of all the Amazon S3 supported Regions and endpoints, see <a
// href="https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region">Regions
// and Endpoints</a> in the AWS General Reference.</p> </note> </li> </ul> <p>For
// example, the following <code>x-amz-grant-write</code> header grants create,
// overwrite, and delete objects permission to LogDelivery group predefined by
// Amazon S3 and two AWS accounts identified by their email addresses.</p> <p>
// <code>x-amz-grant-write: uri="http://acs.amazonaws.com/groups/s3/LogDelivery",
// id="111122223333", id="555566667777" </code> </p> </li> </ul> <p>You can use
// either a canned ACL or specify access permissions explicitly. You cannot do
// both.</p> <p> <b>Grantee Values</b> </p> <p>You can specify the person (grantee)
// to whom you're assigning access rights (using request elements) in the following
// ways:</p> <ul> <li> <p>By the person's ID:</p> <p> <code><Grantee
// xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
// xsi:type="CanonicalUser"><ID><>ID<></ID><DisplayName><>GranteesEmail<></DisplayName>
// </Grantee></code> </p> <p>DisplayName is optional and ignored in the request</p>
// </li> <li> <p>By URI:</p> <p> <code><Grantee
// xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
// xsi:type="Group"><URI><>http://acs.amazonaws.com/groups/global/AuthenticatedUsers<></URI></Grantee></code>
// </p> </li> <li> <p>By Email address:</p> <p> <code><Grantee
// xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
// xsi:type="AmazonCustomerByEmail"><EmailAddress><>Grantees@email.com<></EmailAddress>lt;/Grantee></code>
// </p> <p>The grantee is resolved to the CanonicalUser and, in a response to a GET
// Object acl request, appears as the CanonicalUser. </p> <note> <p>Using email
// addresses to specify a grantee is only supported in the following AWS Regions:
// </p> <ul> <li> <p>US East (N. Virginia)</p> </li> <li> <p>US West (N.
// California)</p> </li> <li> <p> US West (Oregon)</p> </li> <li> <p> Asia Pacific
// (Singapore)</p> </li> <li> <p>Asia Pacific (Sydney)</p> </li> <li> <p>Asia
// Pacific (Tokyo)</p> </li> <li> <p>Europe (Ireland)</p> </li> <li> <p>South
// America (São Paulo)</p> </li> </ul> <p>For a list of all the Amazon S3 supported
// Regions and endpoints, see <a
// href="https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region">Regions
// and Endpoints</a> in the AWS General Reference.</p> </note> </li> </ul> <p
// class="title"> <b>Related Resources</b> </p> <ul> <li> <p> <a>CreateBucket</a>
// </p> </li> <li> <p> <a>DeleteBucket</a> </p> </li> <li> <p> <a>GetObjectAcl</a>
// </p> </li> </ul>
func (c *Client) PutBucketAcl(ctx context.Context, params *PutBucketAclInput, optFns ...func(*Options)) (*PutBucketAclOutput, error) {
	stack := middleware.NewStack("PutBucketAcl", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpPutBucketAclMiddlewares(stack)
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
	addOpPutBucketAclValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketAcl(options.Region), middleware.Before)
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
			OperationName: "PutBucketAcl",
			Err:           err,
		}
	}
	out := result.(*PutBucketAclOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketAclInput struct {
	// Contains the elements that set the ACL permissions for an object per grantee.
	AccessControlPolicy *types.AccessControlPolicy
	// Allows grantee to list the objects in the bucket.
	GrantRead *string
	// Allows grantee to read the bucket ACL.
	GrantReadACP *string
	// The base64-encoded 128-bit MD5 digest of the data. This header must be used as a
	// message integrity check to verify that the request body was not corrupted in
	// transit. For more information, go to RFC 1864.
	// (http://www.ietf.org/rfc/rfc1864.txt)
	ContentMD5 *string
	// The canned ACL to apply to the bucket.
	ACL types.BucketCannedACL
	// Allows grantee to create, overwrite, and delete any object in the bucket.
	GrantWrite *string
	// The bucket to which to apply the ACL.
	Bucket *string
	// Allows grantee the read, write, read ACP, and write ACP permissions on the
	// bucket.
	GrantFullControl *string
	// Allows grantee to write the ACL for the applicable bucket.
	GrantWriteACP *string
}

type PutBucketAclOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpPutBucketAclMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpPutBucketAcl{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketAcl{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutBucketAcl(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketAcl",
	}
}
