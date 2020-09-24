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

// Restores an archived copy of an object back into Amazon S3 This operation
// performs the following types of requests:
//
//     * select - Perform a select query
// on an archived object
//
//     * restore an archive - Restore an archived object
//
// To
// use this operation, you must have permissions to perform the s3:RestoreObject
// action. The bucket owner has this permission by default and can grant this
// permission to others. For more information about permissions, see Permissions
// Related to Bucket Subresource Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html) in the
// Amazon Simple Storage Service Developer Guide. Querying Archives with Select
// Requests You use a select type of request to perform SQL queries on archived
// objects. The archived objects that are being queried by the select request must
// be formatted as uncompressed comma-separated values (CSV) files. You can run
// queries and custom analytics on your archived data without having to restore
// your data to a hotter Amazon S3 tier. For an overview about select requests, see
// Querying Archived Objects
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/querying-glacier-archives.html)
// in the Amazon Simple Storage Service Developer Guide. When making a select
// request, do the following:
//
//     * Define an output location for the select
// query's output. This must be an Amazon S3 bucket in the same AWS Region as the
// bucket that contains the archive object that is being queried. The AWS account
// that initiates the job must have permissions to write to the S3 bucket. You can
// specify the storage class and encryption for the output objects stored in the
// bucket. For more information about output, see Querying Archived Objects
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/querying-glacier-archives.html)
// in the Amazon Simple Storage Service Developer Guide. For more information about
// the S3 structure in the request body, see the following:
//
//         * PutObject
// ()
//
//         * Managing Access with ACLs
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3_ACLs_UsingACLs.html) in the
// Amazon Simple Storage Service Developer Guide
//
//         * Protecting Data Using
// Server-Side Encryption
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/serv-side-encryption.html) in
// the Amazon Simple Storage Service Developer Guide
//
//     * Define the SQL
// expression for the SELECT type of restoration for your query in the request
// body's SelectParameters structure. You can use expressions like the following
// examples.
//
//         * The following expression returns all records from the
// specified object. SELECT * FROM Object
//
//         * Assuming that you are not
// using any headers for data stored in the object, you can specify columns with
// positional headers. SELECT s._1, s._2 FROM Object s WHERE s._3 > 100
//
//         *
// If you have headers and you set the fileHeaderInfo in the CSV structure in the
// request body to USE, you can specify headers in the query. (If you set the
// fileHeaderInfo field to IGNORE, the first row is skipped for the query.) You
// cannot mix ordinal positions with header column names. SELECT s.Id, s.FirstName,
// s.SSN FROM S3Object s
//
// For more information about using SQL with S3 Glacier
// Select restore, see SQL Reference for Amazon S3 Select and S3 Glacier Select
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-glacier-select-sql-reference.html)
// in the Amazon Simple Storage Service Developer Guide. When making a select
// request, you can also do the following:
//
//     * To expedite your queries, specify
// the Expedited tier. For more information about tiers, see "Restoring Archives,"
// later in this topic.
//
//     * Specify details about the data serialization format
// of both the input object that is being queried and the serialization of the
// CSV-encoded query results.
//
// The following are additional important facts about
// the select feature:
//
//     * The output results are new Amazon S3 objects. Unlike
// archive retrievals, they are stored until explicitly deleted-manually or through
// a lifecycle policy.
//
//     * You can issue more than one select request on the
// same Amazon S3 object. Amazon S3 doesn't deduplicate requests, so avoid issuing
// duplicate requests.
//
//     * Amazon S3 accepts a select request even if the object
// has already been restored. A select request doesn’t return error response
// 409.
//
// Restoring Archives Objects in the GLACIER and DEEP_ARCHIVE storage classes
// are archived. To access an archived object, you must first initiate a restore
// request. This restores a temporary copy of the archived object. In a restore
// request, you specify the number of days that you want the restored copy to
// exist. After the specified period, Amazon S3 deletes the temporary copy but the
// object remains archived in the GLACIER or DEEP_ARCHIVE storage class that object
// was restored from. To restore a specific object version, you can provide a
// version ID. If you don't provide a version ID, Amazon S3 restores the current
// version. The time it takes restore jobs to finish depends on which storage class
// the object is being restored from and which data access tier you specify. When
// restoring an archived object (or using a select request), you can specify one of
// the following data access tier options in the Tier element of the request
// body:
//
//     * Expedited - Expedited retrievals allow you to quickly access your
// data stored in the GLACIER storage class when occasional urgent requests for a
// subset of archives are required. For all but the largest archived objects (250
// MB+), data accessed using Expedited retrievals are typically made available
// within 1–5 minutes. Provisioned capacity ensures that retrieval capacity for
// Expedited retrievals is available when you need it. Expedited retrievals and
// provisioned capacity are not available for the DEEP_ARCHIVE storage class.
//
//
// * Standard - S3 Standard retrievals allow you to access any of your archived
// objects within several hours. This is the default option for the GLACIER and
// DEEP_ARCHIVE retrieval requests that do not specify the retrieval option. S3
// Standard retrievals typically complete within 3-5 hours from the GLACIER storage
// class and typically complete within 12 hours from the DEEP_ARCHIVE storage
// class.
//
//     * Bulk - Bulk retrievals are Amazon S3 Glacier’s lowest-cost
// retrieval option, enabling you to retrieve large amounts, even petabytes, of
// data inexpensively in a day. Bulk retrievals typically complete within 5-12
// hours from the GLACIER storage class and typically complete within 48 hours from
// the DEEP_ARCHIVE storage class.
//
// For more information about archive retrieval
// options and provisioned capacity for Expedited data access, see Restoring
// Archived Objects
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/restoring-objects.html) in the
// Amazon Simple Storage Service Developer Guide. You can use Amazon S3 restore
// speed upgrade to change the restore speed to a faster speed while it is in
// progress. You upgrade the speed of an in-progress restoration by issuing another
// restore request to the same object, setting a new Tier request element. When
// issuing a request to upgrade the restore tier, you must choose a tier that is
// faster than the tier that the in-progress restore is using. You must not change
// any other parameters, such as the Days request element. For more information,
// see  Upgrading the Speed of an In-Progress Restore
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/restoring-objects.html#restoring-objects-upgrade-tier.title.html)
// in the Amazon Simple Storage Service Developer Guide. To get the status of
// object restoration, you can send a HEAD request. Operations return the
// x-amz-restore header, which provides information about the restoration status,
// in the response. You can use Amazon S3 event notifications to notify you when a
// restore is initiated or completed. For more information, see Configuring Amazon
// S3 Event Notifications
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the
// Amazon Simple Storage Service Developer Guide. After restoring an archived
// object, you can update the restoration period by reissuing the request with a
// new period. Amazon S3 updates the restoration period relative to the current
// time and charges only for the request-there are no data transfer charges. You
// cannot update the restoration period when Amazon S3 is actively processing your
// current restore request for the object. If your bucket has a lifecycle
// configuration with a rule that includes an expiration action, the object
// expiration overrides the life span that you specify in a restore request. For
// example, if you restore an object copy for 10 days, but the object is scheduled
// to expire in 3 days, Amazon S3 deletes the object in 3 days. For more
// information about lifecycle configuration, see PutBucketLifecycleConfiguration
// () and Object Lifecycle Management
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in
// Amazon Simple Storage Service Developer Guide. Responses A successful operation
// returns either the 200 OK or 202 Accepted status code.
//
//     * If the object copy
// is not previously restored, then Amazon S3 returns 202 Accepted in the
// response.
//
//     * If the object copy is previously restored, Amazon S3 returns
// 200 OK in the response.
//
// Special Errors
//
//         * Code:
// RestoreAlreadyInProgress
//
//         * Cause: Object restore is already in
// progress. (This error does not apply to SELECT type requests.)
//
//         * HTTP
// Status Code: 409 Conflict
//
//         * SOAP Fault Code Prefix: Client
//
//         *
// Code: GlacierExpeditedRetrievalNotAvailable
//
//         * Cause: S3 Glacier
// expedited retrievals are currently not available. Try again later. (Returned if
// there is insufficient capacity to process the Expedited request. This error
// applies only to Expedited retrievals and not to S3 Standard or Bulk
// retrievals.)
//
//         * HTTP Status Code: 503
//
//         * SOAP Fault Code Prefix:
// N/A
//
//     <p class="title"> <b>Related Resources</b> </p> <ul> <li> <p>
// <a>PutBucketLifecycleConfiguration</a> </p> </li> <li> <p>
// <a>GetBucketNotificationConfiguration</a> </p> </li> <li> <p> <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-glacier-select-sql-reference.html">SQL
// Reference for Amazon S3 Select and S3 Glacier Select </a> in the <i>Amazon
// Simple Storage Service Developer Guide</i> </p> </li> </ul>
func (c *Client) RestoreObject(ctx context.Context, params *RestoreObjectInput, optFns ...func(*Options)) (*RestoreObjectOutput, error) {
	stack := middleware.NewStack("RestoreObject", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpRestoreObjectMiddlewares(stack)
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
	addOpRestoreObjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreObject(options.Region), middleware.Before)
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
			OperationName: "RestoreObject",
			Err:           err,
		}
	}
	out := result.(*RestoreObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreObjectInput struct {
	// VersionId used to reference a specific version of the object.
	VersionId *string
	// Container for restore job parameters.
	RestoreRequest *types.RestoreRequest
	// The bucket name or containing the object to restore. When using this API with an
	// access point, you must direct requests to the access point hostname. The access
	// point hostname takes the form
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
	// Object key for which the operation was initiated.
	Key *string
}

type RestoreObjectOutput struct {
	// Indicates the path in the provided S3 output location where Select results will
	// be restored to.
	RestoreOutputPath *string
	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpRestoreObjectMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpRestoreObject{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpRestoreObject{}, middleware.After)
}

func newServiceMetadataMiddleware_opRestoreObject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "RestoreObject",
	}
}
