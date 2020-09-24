// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Writes multiple data records into a Kinesis data stream in a single call (also
// referred to as a PutRecords request). Use this operation to send data into the
// stream for data ingestion and processing. Each PutRecords request can support up
// to 500 records. Each record in the request can be as large as 1 MB, up to a
// limit of 5 MB for the entire request, including partition keys. Each shard can
// support writes up to 1,000 records per second, up to a maximum data write total
// of 1 MB per second. You must specify the name of the stream that captures,
// stores, and transports the data; and an array of request Records, with each
// record in the array requiring a partition key and data blob. The record size
// limit applies to the total size of the partition key and data blob. The data
// blob can be any type of data; for example, a segment from a log file,
// geographic/location data, website clickstream data, and so on. The partition key
// is used by Kinesis Data Streams as input to a hash function that maps the
// partition key and associated data to a specific shard. An MD5 hash function is
// used to map partition keys to 128-bit integer values and to map associated data
// records to shards. As a result of this hashing mechanism, all data records with
// the same partition key map to the same shard within the stream. For more
// information, see Adding Data to a Stream
// (https://docs.aws.amazon.com/kinesis/latest/dev/developing-producers-with-sdk.html#kinesis-using-sdk-java-add-data-to-stream)
// in the Amazon Kinesis Data Streams Developer Guide. Each record in the Records
// array may include an optional parameter, ExplicitHashKey, which overrides the
// partition key to shard mapping. This parameter allows a data producer to
// determine explicitly the shard where the record is stored. For more information,
// see Adding Multiple Records with PutRecords
// (https://docs.aws.amazon.com/kinesis/latest/dev/developing-producers-with-sdk.html#kinesis-using-sdk-java-putrecords)
// in the Amazon Kinesis Data Streams Developer Guide. The PutRecords response
// includes an array of response Records. Each record in the response array
// directly correlates with a record in the request array using natural ordering,
// from the top to the bottom of the request and response. The response Records
// array always includes the same number of records as the request array. The
// response Records array includes both successfully and unsuccessfully processed
// records. Kinesis Data Streams attempts to process all records in each PutRecords
// request. A single record failure does not stop the processing of subsequent
// records. A successfully processed record includes ShardId and SequenceNumber
// values. The ShardId parameter identifies the shard in the stream where the
// record is stored. The SequenceNumber parameter is an identifier assigned to the
// put record, unique to all records in the stream. An unsuccessfully processed
// record includes ErrorCode and ErrorMessage values. ErrorCode reflects the type
// of error and can be one of the following values:
// ProvisionedThroughputExceededException or InternalFailure. ErrorMessage provides
// more detailed information about the ProvisionedThroughputExceededException
// exception including the account ID, stream name, and shard ID of the record that
// was throttled. For more information about partially successful responses, see
// Adding Multiple Records with PutRecords
// (https://docs.aws.amazon.com/kinesis/latest/dev/kinesis-using-sdk-java-add-data-to-stream.html#kinesis-using-sdk-java-putrecords)
// in the Amazon Kinesis Data Streams Developer Guide. By default, data records are
// accessible for 24 hours from the time that they are added to a stream. You can
// use IncreaseStreamRetentionPeriod () or DecreaseStreamRetentionPeriod () to
// modify this retention period.
func (c *Client) PutRecords(ctx context.Context, params *PutRecordsInput, optFns ...func(*Options)) (*PutRecordsOutput, error) {
	stack := middleware.NewStack("PutRecords", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutRecordsMiddlewares(stack)
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
	addOpPutRecordsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutRecords(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

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
			OperationName: "PutRecords",
			Err:           err,
		}
	}
	out := result.(*PutRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A PutRecords request.
type PutRecordsInput struct {
	// The stream name associated with the request.
	StreamName *string
	// The records associated with the request.
	Records []*types.PutRecordsRequestEntry
}

// PutRecords results.
type PutRecordsOutput struct {
	// The encryption type used on the records. This parameter can be one of the
	// following values:
	//
	//     * NONE: Do not encrypt the records.
	//
	//     * KMS: Use
	// server-side encryption on the records using a customer-managed AWS KMS key.
	EncryptionType types.EncryptionType
	// The number of unsuccessfully processed records in a PutRecords request.
	FailedRecordCount *int32
	// An array of successfully and unsuccessfully processed record results, correlated
	// with the request by natural ordering. A record that is successfully added to a
	// stream includes SequenceNumber and ShardId in the result. A record that fails to
	// be added to a stream includes ErrorCode and ErrorMessage in the result.
	Records []*types.PutRecordsResultEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutRecordsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutRecords{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutRecords{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutRecords(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "PutRecords",
	}
}
