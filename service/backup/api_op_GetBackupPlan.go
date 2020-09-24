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

// Returns the body of a backup plan in JSON format, in addition to plan metadata.
func (c *Client) GetBackupPlan(ctx context.Context, params *GetBackupPlanInput, optFns ...func(*Options)) (*GetBackupPlanOutput, error) {
	stack := middleware.NewStack("GetBackupPlan", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetBackupPlanMiddlewares(stack)
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
	addOpGetBackupPlanValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBackupPlan(options.Region), middleware.Before)
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
			OperationName: "GetBackupPlan",
			Err:           err,
		}
	}
	out := result.(*GetBackupPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBackupPlanInput struct {
	// Unique, randomly generated, Unicode, UTF-8 encoded strings that are at most
	// 1,024 bytes long. Version IDs cannot be edited.
	VersionId *string
	// Uniquely identifies a backup plan.
	BackupPlanId *string
}

type GetBackupPlanOutput struct {
	// A unique string that identifies the request and allows failed requests to be
	// retried without the risk of executing the operation twice.
	CreatorRequestId *string
	// Uniquely identifies a backup plan.
	BackupPlanId *string
	// Unique, randomly generated, Unicode, UTF-8 encoded strings that are at most
	// 1,024 bytes long. Version IDs cannot be edited.
	VersionId *string
	// The date and time that a backup plan is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time
	// The date and time that a backup plan is deleted, in Unix format and Coordinated
	// Universal Time (UTC). The value of DeletionDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	DeletionDate *time.Time
	// The last time a job to back up resources was executed with this backup plan. A
	// date and time, in Unix format and Coordinated Universal Time (UTC). The value of
	// LastExecutionDate is accurate to milliseconds. For example, the value
	// 1516925490.087 represents Friday, January 26, 2018 12:11:30.087 AM.
	LastExecutionDate *time.Time
	// Specifies the body of a backup plan. Includes a BackupPlanName and one or more
	// sets of Rules.
	BackupPlan *types.BackupPlan
	// An Amazon Resource Name (ARN) that uniquely identifies a backup plan; for
	// example,
	// arn:aws:backup:us-east-1:123456789012:plan:8F81F553-3A74-4A3F-B93D-B3360DC80C50.
	BackupPlanArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetBackupPlanMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetBackupPlan{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBackupPlan{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBackupPlan(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "GetBackupPlan",
	}
}
