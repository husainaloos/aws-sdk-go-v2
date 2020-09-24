// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns domain configuration information about the specified Elasticsearch
// domains, including the domain ID, domain endpoint, and domain ARN.
func (c *Client) DescribeElasticsearchDomains(ctx context.Context, params *DescribeElasticsearchDomainsInput, optFns ...func(*Options)) (*DescribeElasticsearchDomainsOutput, error) {
	stack := middleware.NewStack("DescribeElasticsearchDomains", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeElasticsearchDomainsMiddlewares(stack)
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
	addOpDescribeElasticsearchDomainsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeElasticsearchDomains(options.Region), middleware.Before)
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
			OperationName: "DescribeElasticsearchDomains",
			Err:           err,
		}
	}
	out := result.(*DescribeElasticsearchDomainsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeElasticsearchDomains () operation.
// By default, the API returns the status of all Elasticsearch domains.
type DescribeElasticsearchDomainsInput struct {
	// The Elasticsearch domains for which you want information.
	DomainNames []*string
}

// The result of a DescribeElasticsearchDomains request. Contains the status of the
// specified domains or all domains owned by the account.
type DescribeElasticsearchDomainsOutput struct {
	// The status of the domains requested in the DescribeElasticsearchDomains request.
	DomainStatusList []*types.ElasticsearchDomainStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeElasticsearchDomainsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeElasticsearchDomains{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeElasticsearchDomains{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeElasticsearchDomains(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DescribeElasticsearchDomains",
	}
}
