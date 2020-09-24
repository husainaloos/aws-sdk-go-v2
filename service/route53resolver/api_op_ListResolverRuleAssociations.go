// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the associations that were created between resolver rules and VPCs using
// the current AWS account.
func (c *Client) ListResolverRuleAssociations(ctx context.Context, params *ListResolverRuleAssociationsInput, optFns ...func(*Options)) (*ListResolverRuleAssociationsOutput, error) {
	stack := middleware.NewStack("ListResolverRuleAssociations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListResolverRuleAssociationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListResolverRuleAssociations(options.Region), middleware.Before)
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
			OperationName: "ListResolverRuleAssociations",
			Err:           err,
		}
	}
	out := result.(*ListResolverRuleAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResolverRuleAssociationsInput struct {
	// The maximum number of rule associations that you want to return in the response
	// to a ListResolverRuleAssociations request. If you don't specify a value for
	// MaxResults, Resolver returns up to 100 rule associations.
	MaxResults *int32
	// An optional specification to return a subset of resolver rules, such as resolver
	// rules that are associated with the same VPC ID. If you submit a second or
	// subsequent ListResolverRuleAssociations request and specify the NextToken
	// parameter, you must use the same values for Filters, if any, as in the previous
	// request.
	Filters []*types.Filter
	// For the first ListResolverRuleAssociation request, omit this value. If you have
	// more than MaxResults rule associations, you can submit another
	// ListResolverRuleAssociation request to get the next group of rule associations.
	// In the next request, specify the value of NextToken from the previous response.
	NextToken *string
}

type ListResolverRuleAssociationsOutput struct {
	// The associations that were created between resolver rules and VPCs using the
	// current AWS account, and that match the specified filters, if any.
	ResolverRuleAssociations []*types.ResolverRuleAssociation
	// The value that you specified for MaxResults in the request.
	MaxResults *int32
	// If more than MaxResults rule associations match the specified criteria, you can
	// submit another ListResolverRuleAssociation request to get the next group of
	// results. In the next request, specify the value of NextToken from the previous
	// response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListResolverRuleAssociationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListResolverRuleAssociations{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResolverRuleAssociations{}, middleware.After)
}

func newServiceMetadataMiddleware_opListResolverRuleAssociations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "ListResolverRuleAssociations",
	}
}
