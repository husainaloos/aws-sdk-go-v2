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

// Modify a setting for an Amazon Aurora global cluster. You can change one or more
// database configuration parameters by specifying these parameters and the new
// values in the request. For more information on Amazon Aurora, see  What Is
// Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. This action only applies to Aurora DB clusters.
func (c *Client) ModifyGlobalCluster(ctx context.Context, params *ModifyGlobalClusterInput, optFns ...func(*Options)) (*ModifyGlobalClusterOutput, error) {
	stack := middleware.NewStack("ModifyGlobalCluster", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyGlobalClusterMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyGlobalCluster(options.Region), middleware.Before)
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
			OperationName: "ModifyGlobalCluster",
			Err:           err,
		}
	}
	out := result.(*ModifyGlobalClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyGlobalClusterInput struct {
	// Indicates if the global database cluster has deletion protection enabled. The
	// global database cluster can't be deleted when deletion protection is enabled.
	DeletionProtection *bool
	// The DB cluster identifier for the global cluster being modified. This parameter
	// isn't case-sensitive. Constraints:
	//
	//     * Must match the identifier of an
	// existing global database cluster.
	GlobalClusterIdentifier *string
	// The new cluster identifier for the global database cluster when modifying a
	// global database cluster. This value is stored as a lowercase string.
	// Constraints:
	//
	//     * Must contain from 1 to 63 letters, numbers, or hyphens
	//
	//
	// * The first character must be a letter
	//
	//     * Can't end with a hyphen or contain
	// two consecutive hyphens
	//
	// Example: my-cluster2
	NewGlobalClusterIdentifier *string
}

type ModifyGlobalClusterOutput struct {
	// A data type representing an Aurora global database.
	GlobalCluster *types.GlobalCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyGlobalClusterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyGlobalCluster{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyGlobalCluster{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyGlobalCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyGlobalCluster",
	}
}
