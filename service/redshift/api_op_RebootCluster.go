// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Reboots a cluster. This action is taken as soon as possible. It results in a
// momentary outage to the cluster, during which the cluster status is set to
// rebooting. A cluster event is created when the reboot is completed. Any pending
// cluster modifications (see ModifyCluster ()) are applied at this reboot. For
// more information about managing clusters, go to Amazon Redshift Clusters
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html) in
// the Amazon Redshift Cluster Management Guide.
func (c *Client) RebootCluster(ctx context.Context, params *RebootClusterInput, optFns ...func(*Options)) (*RebootClusterOutput, error) {
	if params == nil {
		params = &RebootClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RebootCluster", params, optFns, addOperationRebootClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RebootClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type RebootClusterInput struct {

	// The cluster identifier.
	//
	// This member is required.
	ClusterIdentifier *string
}

type RebootClusterOutput struct {

	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRebootClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRebootCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRebootCluster{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpRebootClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRebootCluster(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRebootCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "RebootCluster",
	}
}
