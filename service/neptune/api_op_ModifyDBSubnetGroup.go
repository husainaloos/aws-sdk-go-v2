// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies an existing DB subnet group. DB subnet groups must contain at least one
// subnet in at least two AZs in the AWS Region.
func (c *Client) ModifyDBSubnetGroup(ctx context.Context, params *ModifyDBSubnetGroupInput, optFns ...func(*Options)) (*ModifyDBSubnetGroupOutput, error) {
	stack := middleware.NewStack("ModifyDBSubnetGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyDBSubnetGroupMiddlewares(stack)
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
	addOpModifyDBSubnetGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBSubnetGroup(options.Region), middleware.Before)
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
			OperationName: "ModifyDBSubnetGroup",
			Err:           err,
		}
	}
	out := result.(*ModifyDBSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDBSubnetGroupInput struct {
	// The name for the DB subnet group. This value is stored as a lowercase string.
	// You can't modify the default subnet group. Constraints: Must match the name of
	// an existing DBSubnetGroup. Must not be default. Example: mySubnetgroup
	DBSubnetGroupName *string
	// The description for the DB subnet group.
	DBSubnetGroupDescription *string
	// The EC2 subnet IDs for the DB subnet group.
	SubnetIds []*string
}

type ModifyDBSubnetGroupOutput struct {
	// Contains the details of an Amazon Neptune DB subnet group. This data type is
	// used as a response element in the DescribeDBSubnetGroups () action.
	DBSubnetGroup *types.DBSubnetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyDBSubnetGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBSubnetGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBSubnetGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyDBSubnetGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBSubnetGroup",
	}
}
