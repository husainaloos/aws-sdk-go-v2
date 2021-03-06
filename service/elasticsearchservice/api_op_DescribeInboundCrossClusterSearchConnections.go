// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all the inbound cross-cluster search connections for a destination domain.
func (c *Client) DescribeInboundCrossClusterSearchConnections(ctx context.Context, params *DescribeInboundCrossClusterSearchConnectionsInput, optFns ...func(*Options)) (*DescribeInboundCrossClusterSearchConnectionsOutput, error) {
	if params == nil {
		params = &DescribeInboundCrossClusterSearchConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInboundCrossClusterSearchConnections", params, optFns, addOperationDescribeInboundCrossClusterSearchConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInboundCrossClusterSearchConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeInboundCrossClusterSearchConnections
// () operation.
type DescribeInboundCrossClusterSearchConnectionsInput struct {

	// A list of filters used to match properties for inbound cross-cluster search
	// connection. Available Filter () names for this operation are:
	//
	//     *
	// cross-cluster-search-connection-id
	//
	//     * source-domain-info.domain-name
	//
	//     *
	// source-domain-info.owner-id
	//
	//     * source-domain-info.region
	//
	//     *
	// destination-domain-info.domain-name
	Filters []*types.Filter

	// Set this value to limit the number of results returned. If not specified,
	// defaults to 100.
	MaxResults *int32

	// NextToken is sent in case the earlier API call results contain the NextToken. It
	// is used for pagination.
	NextToken *string
}

// The result of a DescribeInboundCrossClusterSearchConnections () request.
// Contains the list of connections matching the filter criteria.
type DescribeInboundCrossClusterSearchConnectionsOutput struct {

	// Consists of list of InboundCrossClusterSearchConnection () matching the
	// specified filter criteria.
	CrossClusterSearchConnections []*types.InboundCrossClusterSearchConnection

	// If more results are available and NextToken is present, make the next request to
	// the same API with the received NextToken to paginate the remaining results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeInboundCrossClusterSearchConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeInboundCrossClusterSearchConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeInboundCrossClusterSearchConnections{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInboundCrossClusterSearchConnections(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeInboundCrossClusterSearchConnections(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DescribeInboundCrossClusterSearchConnections",
	}
}
