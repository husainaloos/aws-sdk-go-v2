// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroups

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of AWS resource identifiers that matches tne specified query. The
// query uses the same format as a resource query in a CreateGroup or
// UpdateGroupQuery operation.
func (c *Client) SearchResources(ctx context.Context, params *SearchResourcesInput, optFns ...func(*Options)) (*SearchResourcesOutput, error) {
	stack := middleware.NewStack("SearchResources", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpSearchResourcesMiddlewares(stack)
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
	addOpSearchResourcesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSearchResources(options.Region), middleware.Before)
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
			OperationName: "SearchResources",
			Err:           err,
		}
	}
	out := result.(*SearchResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchResourcesInput struct {
	// The search query, using the same formats that are supported for resource group
	// definition. For more information, see CreateGroup ().
	ResourceQuery *types.ResourceQuery
	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that the service might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	MaxResults *int32
	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value provided by a previous call's
	// NextToken response to indicate where the output should continue from.
	NextToken *string
}

type SearchResourcesOutput struct {
	// A list of QueryError objects. Each error is an object that contains ErrorCode
	// and Message structures. Possible values for ErrorCode are
	// CLOUDFORMATION_STACK_INACTIVE and CLOUDFORMATION_STACK_NOT_EXISTING.
	QueryErrors []*types.QueryError
	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null.
	NextToken *string
	// The ARNs and resource types of resources that are members of the group that you
	// specified.
	ResourceIdentifiers []*types.ResourceIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpSearchResourcesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpSearchResources{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchResources{}, middleware.After)
}

func newServiceMetadataMiddleware_opSearchResources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resource-groups",
		OperationName: "SearchResources",
	}
}
