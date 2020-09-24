// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables server-side encryption (SSE) for the delivery stream. This operation is
// asynchronous. It returns immediately. When you invoke it, Kinesis Data Firehose
// first sets the encryption status of the stream to ENABLING, and then to ENABLED.
// The encryption status of a delivery stream is the Status property in
// DeliveryStreamEncryptionConfiguration (). If the operation fails, the encryption
// status changes to ENABLING_FAILED. You can continue to read and write data to
// your delivery stream while the encryption status is ENABLING, but the data is
// not encrypted. It can take up to 5 seconds after the encryption status changes
// to ENABLED before all records written to the delivery stream are encrypted. To
// find out whether a record or a batch of records was encrypted, check the
// response elements PutRecordOutput$Encrypted () and
// PutRecordBatchOutput$Encrypted (), respectively. To check the encryption status
// of a delivery stream, use DescribeDeliveryStream (). Even if encryption is
// currently enabled for a delivery stream, you can still invoke this operation on
// it to change the ARN of the CMK or both its type and ARN. If you invoke this
// method to change the CMK, and the old CMK is of type CUSTOMER_MANAGED_CMK,
// Kinesis Data Firehose schedules the grant it had on the old CMK for retirement.
// If the new CMK is of type CUSTOMER_MANAGED_CMK, Kinesis Data Firehose creates a
// grant that enables it to use the new CMK to encrypt and decrypt data and to
// manage the grant. If a delivery stream already has encryption enabled and then
// you invoke this operation to change the ARN of the CMK or both its type and ARN
// and you get ENABLING_FAILED, this only means that the attempt to change the CMK
// failed. In this case, encryption remains enabled with the old CMK. If the
// encryption status of your delivery stream is ENABLING_FAILED, you can invoke
// this operation again with a valid CMK. The CMK must be enabled and the key
// policy mustn't explicitly deny the permission for Kinesis Data Firehose to
// invoke KMS encrypt and decrypt operations. You can enable SSE for a delivery
// stream only if it's a delivery stream that uses DirectPut as its source. The
// StartDeliveryStreamEncryption and StopDeliveryStreamEncryption operations have a
// combined limit of 25 calls per delivery stream per 24 hours. For example, you
// reach the limit if you call StartDeliveryStreamEncryption 13 times and
// StopDeliveryStreamEncryption 12 times for the same delivery stream in a 24-hour
// period.
func (c *Client) StartDeliveryStreamEncryption(ctx context.Context, params *StartDeliveryStreamEncryptionInput, optFns ...func(*Options)) (*StartDeliveryStreamEncryptionOutput, error) {
	stack := middleware.NewStack("StartDeliveryStreamEncryption", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartDeliveryStreamEncryptionMiddlewares(stack)
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
	addOpStartDeliveryStreamEncryptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartDeliveryStreamEncryption(options.Region), middleware.Before)
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
			OperationName: "StartDeliveryStreamEncryption",
			Err:           err,
		}
	}
	out := result.(*StartDeliveryStreamEncryptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDeliveryStreamEncryptionInput struct {
	// The name of the delivery stream for which you want to enable server-side
	// encryption (SSE).
	DeliveryStreamName *string
	// Used to specify the type and Amazon Resource Name (ARN) of the KMS key needed
	// for Server-Side Encryption (SSE).
	DeliveryStreamEncryptionConfigurationInput *types.DeliveryStreamEncryptionConfigurationInput
}

type StartDeliveryStreamEncryptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartDeliveryStreamEncryptionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartDeliveryStreamEncryption{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartDeliveryStreamEncryption{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartDeliveryStreamEncryption(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "firehose",
		OperationName: "StartDeliveryStreamEncryption",
	}
}
