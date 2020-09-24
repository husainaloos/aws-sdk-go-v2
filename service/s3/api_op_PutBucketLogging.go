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

// Set the logging parameters for a bucket and to specify permissions for who can
// view and modify the logging parameters. All logs are saved to buckets in the
// same AWS Region as the source bucket. To set the logging status of a bucket, you
// must be the bucket owner.  <p>The bucket owner is automatically granted
// FULL_CONTROL to all logs. You use the <code>Grantee</code> request element to
// grant access to other people. The <code>Permissions</code> request element
// specifies the kind of access the grantee has to the logs.</p> <p> <b>Grantee
// Values</b> </p> <p>You can specify the person (grantee) to whom you're assigning
// access rights (using request elements) in the following ways:</p> <ul> <li>
// <p>By the person's ID:</p> <p> <code><Grantee
// xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
// xsi:type="CanonicalUser"><ID><>ID<></ID><DisplayName><>GranteesEmail<></DisplayName>
// </Grantee></code> </p> <p>DisplayName is optional and ignored in the
// request.</p> </li> <li> <p>By Email address:</p> <p> <code> <Grantee
// xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
// xsi:type="AmazonCustomerByEmail"><EmailAddress><>Grantees@email.com<></EmailAddress></Grantee></code>
// </p> <p>The grantee is resolved to the CanonicalUser and, in a response to a GET
// Object acl request, appears as the CanonicalUser.</p> </li> <li> <p>By URI:</p>
// <p> <code><Grantee xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
// xsi:type="Group"><URI><>http://acs.amazonaws.com/groups/global/AuthenticatedUsers<></URI></Grantee></code>
// </p> </li> </ul> <p>To enable logging, you use LoggingEnabled and its children
// request elements. To disable logging, you use an empty BucketLoggingStatus
// request element:</p> <p> <code><BucketLoggingStatus
// xmlns="http://doc.s3.amazonaws.com/2006-03-01" /></code> </p> <p>For more
// information about server access logging, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerLogs.html">Server
// Access Logging</a>. </p> <p>For more information about creating a bucket, see
// <a>CreateBucket</a>. For more information about returning the logging status of
// a bucket, see <a>GetBucketLogging</a>.</p> <p>The following operations are
// related to <code>PutBucketLogging</code>:</p> <ul> <li> <p> <a>PutObject</a>
// </p> </li> <li> <p> <a>DeleteBucket</a> </p> </li> <li> <p> <a>CreateBucket</a>
// </p> </li> <li> <p> <a>GetBucketLogging</a> </p> </li> </ul>
func (c *Client) PutBucketLogging(ctx context.Context, params *PutBucketLoggingInput, optFns ...func(*Options)) (*PutBucketLoggingOutput, error) {
	stack := middleware.NewStack("PutBucketLogging", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpPutBucketLoggingMiddlewares(stack)
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
	addOpPutBucketLoggingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketLogging(options.Region), middleware.Before)
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
			OperationName: "PutBucketLogging",
			Err:           err,
		}
	}
	out := result.(*PutBucketLoggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketLoggingInput struct {
	// Container for logging status information.
	BucketLoggingStatus *types.BucketLoggingStatus
	// The MD5 hash of the PutBucketLogging request body.
	ContentMD5 *string
	// The name of the bucket for which to set the logging parameters.
	Bucket *string
}

type PutBucketLoggingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpPutBucketLoggingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpPutBucketLogging{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketLogging{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutBucketLogging(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketLogging",
	}
}
