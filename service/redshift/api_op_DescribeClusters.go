// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns properties of provisioned clusters including general cluster properties,
// cluster database properties, maintenance and backup properties, and security and
// access properties. This operation supports pagination. For more information
// about managing clusters, go to Amazon Redshift Clusters
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html) in
// the Amazon Redshift Cluster Management Guide. If you specify both tag keys and
// tag values in the same request, Amazon Redshift returns all clusters that match
// any combination of the specified keys and values. For example, if you have owner
// and environment for tag keys, and admin and test for tag values, all clusters
// that have any combination of those values are returned. If both tag keys and
// values are omitted from the request, clusters are returned regardless of whether
// they have tag keys or values associated with them.
func (c *Client) DescribeClusters(ctx context.Context, params *DescribeClustersInput, optFns ...func(*Options)) (*DescribeClustersOutput, error) {
	stack := middleware.NewStack("DescribeClusters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeClustersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClusters(options.Region), middleware.Before)
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
			OperationName: "DescribeClusters",
			Err:           err,
		}
	}
	out := result.(*DescribeClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeClustersInput struct {
	// A tag key or keys for which you want to return all matching clusters that are
	// associated with the specified key or keys. For example, suppose that you have
	// clusters that are tagged with keys called owner and environment. If you specify
	// both of these tag keys in the request, Amazon Redshift returns a response with
	// the clusters that have either or both of these tag keys associated with them.
	TagKeys []*string
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32
	// The unique identifier of a cluster whose properties you are requesting. This
	// parameter is case sensitive. The default is that all clusters defined for an
	// account are returned.
	ClusterIdentifier *string
	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeClusters () request exceed the
	// value specified in MaxRecords, AWS returns a value in the Marker field of the
	// response. You can retrieve the next set of response records by providing the
	// returned marker value in the Marker parameter and retrying the request.
	// Constraints: You can specify either the ClusterIdentifier parameter or the
	// Marker parameter, but not both.
	Marker *string
	// A tag value or values for which you want to return all matching clusters that
	// are associated with the specified tag value or values. For example, suppose that
	// you have clusters that are tagged with values called admin and test. If you
	// specify both of these tag values in the request, Amazon Redshift returns a
	// response with the clusters that have either or both of these tag values
	// associated with them.
	TagValues []*string
}

// Contains the output from the DescribeClusters () action.
type DescribeClustersOutput struct {
	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string
	// A list of Cluster objects, where each object describes one cluster.
	Clusters []*types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeClustersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeClusters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeClusters{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeClusters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeClusters",
	}
}
