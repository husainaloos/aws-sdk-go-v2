// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns metadata associated with a restore job that is specified by a job ID.
func (c *Client) DescribeRestoreJob(ctx context.Context, params *DescribeRestoreJobInput, optFns ...func(*Options)) (*DescribeRestoreJobOutput, error) {
	stack := middleware.NewStack("DescribeRestoreJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeRestoreJobMiddlewares(stack)
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
	addOpDescribeRestoreJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRestoreJob(options.Region), middleware.Before)
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
			OperationName: "DescribeRestoreJob",
			Err:           err,
		}
	}
	out := result.(*DescribeRestoreJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRestoreJobInput struct {
	// Uniquely identifies the job that restores a recovery point.
	RestoreJobId *string
}

type DescribeRestoreJobOutput struct {
	// Specifies the IAM role ARN used to create the target recovery point; for
	// example, arn:aws:iam::123456789012:role/S3Access.
	IamRoleArn *string
	// Returns metadata associated with a restore job listed by resource type.
	ResourceType *string
	// The date and time that a restore job is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time
	// The date and time that a job to restore a recovery point is completed, in Unix
	// format and Coordinated Universal Time (UTC). The value of CompletionDate is
	// accurate to milliseconds. For example, the value 1516925490.087 represents
	// Friday, January 26, 2018 12:11:30.087 AM.
	CompletionDate *time.Time
	// Status code specifying the state of the job that is initiated by AWS Backup to
	// restore a recovery point.
	Status types.RestoreJobStatus
	// An Amazon Resource Name (ARN) that uniquely identifies a resource whose recovery
	// point is being restored. The format of the ARN depends on the resource type of
	// the backed-up resource.
	CreatedResourceArn *string
	// The size, in bytes, of the restored resource.
	BackupSizeInBytes *int64
	// An ARN that uniquely identifies a recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	RecoveryPointArn *string
	// Contains an estimated percentage that is complete of a job at the time the job
	// status was queried.
	PercentDone *string
	// Returns the account ID that owns the restore job.
	AccountId *string
	// Uniquely identifies the job that restores a recovery point.
	RestoreJobId *string
	// A message showing the status of a job to restore a recovery point.
	StatusMessage *string
	// The amount of time in minutes that a job restoring a recovery point is expected
	// to take.
	ExpectedCompletionTimeMinutes *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeRestoreJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRestoreJob{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRestoreJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeRestoreJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "DescribeRestoreJob",
	}
}
