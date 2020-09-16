// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a replication job. The replication job schedules periodic replication
// runs to replicate your server to AWS. Each replication run creates an Amazon
// Machine Image (AMI).
func (c *Client) CreateReplicationJob(ctx context.Context, params *CreateReplicationJobInput, optFns ...func(*Options)) (*CreateReplicationJobOutput, error) {
	stack := middleware.NewStack("CreateReplicationJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateReplicationJobMiddlewares(stack)
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
	addOpCreateReplicationJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReplicationJob(options.Region), middleware.Before)

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
			OperationName: "CreateReplicationJob",
			Err:           err,
		}
	}
	out := result.(*CreateReplicationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateReplicationJobInput struct {
	// The name of the IAM role to be used by the AWS SMS.
	RoleName *string
	// The identifier of the server.
	ServerId *string
	//
	RunOnce *bool
	// The description of the replication job.
	Description *string
	// The time between consecutive replication runs, in hours.
	Frequency *int32
	// The maximum number of SMS-created AMIs to retain. The oldest will be deleted
	// once the maximum number is reached and a new AMI is created.
	NumberOfRecentAmisToKeep *int32
	// The license type to be used for the AMI created by a successful replication run.
	LicenseType types.LicenseType
	// When true, the replication job produces encrypted AMIs. See also KmsKeyId below.
	Encrypted *bool
	// KMS key ID for replication jobs that produce encrypted AMIs. Can be any of the
	// following:
	//
	//     * KMS key ID
	//
	//     * KMS key alias
	//
	//     * ARN referring to KMS
	// key ID
	//
	//     * ARN referring to KMS key alias
	//
	// If encrypted is true but a KMS key
	// id is not specified, the customer's default KMS key for EBS is used.
	KmsKeyId *string
	// The seed replication time.
	SeedReplicationTime *time.Time
}

type CreateReplicationJobOutput struct {
	// The unique identifier of the replication job.
	ReplicationJobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateReplicationJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateReplicationJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateReplicationJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateReplicationJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "CreateReplicationJob",
	}
}