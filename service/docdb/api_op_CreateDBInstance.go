// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new instance.
func (c *Client) CreateDBInstance(ctx context.Context, params *CreateDBInstanceInput, optFns ...func(*Options)) (*CreateDBInstanceOutput, error) {
	stack := middleware.NewStack("CreateDBInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateDBInstanceMiddlewares(stack)
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
	addOpCreateDBInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBInstance(options.Region), middleware.Before)
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
			OperationName: "CreateDBInstance",
			Err:           err,
		}
	}
	out := result.(*CreateDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to CreateDBInstance ().
type CreateDBInstanceInput struct {
	// A value that specifies the order in which an Amazon DocumentDB replica is
	// promoted to the primary instance after a failure of the existing primary
	// instance. Default: 1 Valid values: 0-15
	PromotionTier *int32
	// The compute and memory capacity of the instance; for example, db.r5.large.
	DBInstanceClass *string
	// Indicates that minor engine upgrades are applied automatically to the instance
	// during the maintenance window. Default: true
	AutoMinorVersionUpgrade *bool
	// The tags to be assigned to the instance. You can assign up to 10 tags to an
	// instance.
	Tags []*types.Tag
	// The name of the database engine to be used for this instance. Valid value: docdb
	Engine *string
	// The time range each week during which system maintenance can occur, in Universal
	// Coordinated Time (UTC). Format: ddd:hh24:mi-ddd:hh24:mi The default is a
	// 30-minute window selected at random from an 8-hour block of time for each AWS
	// Region, occurring on a random day of the week. Valid days: Mon, Tue, Wed, Thu,
	// Fri, Sat, Sun Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string
	// The identifier of the cluster that the instance will belong to.
	DBClusterIdentifier *string
	// The Amazon EC2 Availability Zone that the instance is created in. Default: A
	// random, system-chosen Availability Zone in the endpoint's AWS Region. Example:
	// us-east-1d Constraint: The AvailabilityZone parameter can't be specified if the
	// MultiAZ parameter is set to true. The specified Availability Zone must be in the
	// same AWS Region as the current endpoint.
	AvailabilityZone *string
	// The instance identifier. This parameter is stored as a lowercase string.
	// Constraints:
	//
	//     * Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//
	// * The first character must be a letter.
	//
	//     * Cannot end with a hyphen or
	// contain two consecutive hyphens.
	//
	// Example: mydbinstance
	DBInstanceIdentifier *string
}

type CreateDBInstanceOutput struct {
	// Detailed information about an instance.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateDBInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBInstance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDBInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBInstance",
	}
}
