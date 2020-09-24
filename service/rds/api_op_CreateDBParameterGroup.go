// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new DB parameter group. A DB parameter group is initially created with
// the default parameters for the database engine used by the DB instance. To
// provide custom values for any of the parameters, you must modify the group after
// creating it using ModifyDBParameterGroup. Once you've created a DB parameter
// group, you need to associate it with your DB instance using ModifyDBInstance.
// When you associate a new DB parameter group with a running DB instance, you need
// to reboot the DB instance without failover for the new DB parameter group and
// associated settings to take effect. After you create a DB parameter group, you
// should wait at least 5 minutes before creating your first DB instance that uses
// that DB parameter group as the default parameter group. This allows Amazon RDS
// to fully complete the create action before the parameter group is used as the
// default for a new DB instance. This is especially important for parameters that
// are critical when creating the default database for a DB instance, such as the
// character set for the default database defined by the character_set_database
// parameter. You can use the Parameter Groups option of the Amazon RDS console
// (https://console.aws.amazon.com/rds/) or the DescribeDBParameters command to
// verify that your DB parameter group has been created or modified.
func (c *Client) CreateDBParameterGroup(ctx context.Context, params *CreateDBParameterGroupInput, optFns ...func(*Options)) (*CreateDBParameterGroupOutput, error) {
	stack := middleware.NewStack("CreateDBParameterGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateDBParameterGroupMiddlewares(stack)
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
	addOpCreateDBParameterGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBParameterGroup(options.Region), middleware.Before)
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
			OperationName: "CreateDBParameterGroup",
			Err:           err,
		}
	}
	out := result.(*CreateDBParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateDBParameterGroupInput struct {
	// The DB parameter group family name. A DB parameter group can be associated with
	// one and only one DB parameter group family, and can be applied only to a DB
	// instance running a database engine and engine version compatible with that DB
	// parameter group family. To list all of the available parameter group families,
	// use the following command: aws rds describe-db-engine-versions --query
	// "DBEngineVersions[].DBParameterGroupFamily" The output contains duplicates.
	DBParameterGroupFamily *string
	// Tags to assign to the DB parameter group.
	Tags []*types.Tag
	// The name of the DB parameter group. Constraints:
	//
	//     * Must be 1 to 255
	// letters, numbers, or hyphens.
	//
	//     * First character must be a letter
	//
	//     *
	// Can't end with a hyphen or contain two consecutive hyphens
	//
	// This value is stored
	// as a lowercase string.
	DBParameterGroupName *string
	// The description for the DB parameter group.
	Description *string
}

type CreateDBParameterGroupOutput struct {
	// Contains the details of an Amazon RDS DB parameter group. This data type is used
	// as a response element in the DescribeDBParameterGroups action.
	DBParameterGroup *types.DBParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateDBParameterGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBParameterGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBParameterGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDBParameterGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBParameterGroup",
	}
}
