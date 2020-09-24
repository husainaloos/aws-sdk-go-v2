// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Schedules the deletion of a customer master key (CMK). You may provide a waiting
// period, specified in days, before deletion occurs. If you do not provide a
// waiting period, the default period of 30 days is used. When this operation is
// successful, the key state of the CMK changes to PendingDeletion. Before the
// waiting period ends, you can use CancelKeyDeletion () to cancel the deletion of
// the CMK. After the waiting period ends, AWS KMS deletes the CMK and all AWS KMS
// data associated with it, including all aliases that refer to it. Deleting a CMK
// is a destructive and potentially dangerous operation. When a CMK is deleted, all
// data that was encrypted under the CMK is unrecoverable. To prevent the use of a
// CMK without deleting it, use DisableKey (). If you schedule deletion of a CMK
// from a custom key store
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html),
// when the waiting period expires, ScheduleKeyDeletion deletes the CMK from AWS
// KMS. Then AWS KMS makes a best effort to delete the key material from the
// associated AWS CloudHSM cluster. However, you might need to manually delete the
// orphaned key material
// (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html#fix-keystore-orphaned-key)
// from the cluster and its backups. You cannot perform this operation on a CMK in
// a different AWS account. For more information about scheduling a CMK for
// deletion, see Deleting Customer Master Keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html) in
// the AWS Key Management Service Developer Guide. The CMK that you use for this
// operation must be in a compatible key state. For details, see How Key State
// Affects Use of a Customer Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide.
func (c *Client) ScheduleKeyDeletion(ctx context.Context, params *ScheduleKeyDeletionInput, optFns ...func(*Options)) (*ScheduleKeyDeletionOutput, error) {
	stack := middleware.NewStack("ScheduleKeyDeletion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpScheduleKeyDeletionMiddlewares(stack)
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
	addOpScheduleKeyDeletionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opScheduleKeyDeletion(options.Region), middleware.Before)
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
			OperationName: "ScheduleKeyDeletion",
			Err:           err,
		}
	}
	out := result.(*ScheduleKeyDeletionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ScheduleKeyDeletionInput struct {
	// The unique identifier of the customer master key (CMK) to delete.  <p>Specify
	// the key ID or the Amazon Resource Name (ARN) of the CMK.</p> <p>For example:</p>
	// <ul> <li> <p>Key ID: <code>1234abcd-12ab-34cd-56ef-1234567890ab</code> </p>
	// </li> <li> <p>Key ARN:
	// <code>arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab</code>
	// </p> </li> </ul> <p>To get the key ID and key ARN for a CMK, use <a>ListKeys</a>
	// or <a>DescribeKey</a>.</p>
	KeyId *string
	// The waiting period, specified in number of days. After the waiting period ends,
	// AWS KMS deletes the customer master key (CMK). This value is optional. If you
	// include a value, it must be between 7 and 30, inclusive. If you do not include a
	// value, it defaults to 30.
	PendingWindowInDays *int32
}

type ScheduleKeyDeletionOutput struct {
	// The Amazon Resource Name (key ARN
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN))
	// of the CMK whose deletion is scheduled.
	KeyId *string
	// The date and time after which AWS KMS deletes the customer master key (CMK).
	DeletionDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpScheduleKeyDeletionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpScheduleKeyDeletion{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpScheduleKeyDeletion{}, middleware.After)
}

func newServiceMetadataMiddleware_opScheduleKeyDeletion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "ScheduleKeyDeletion",
	}
}
