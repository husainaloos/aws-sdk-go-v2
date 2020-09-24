// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a listing of the members in a network and properties of their
// configurations.
func (c *Client) ListMembers(ctx context.Context, params *ListMembersInput, optFns ...func(*Options)) (*ListMembersOutput, error) {
	stack := middleware.NewStack("ListMembers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListMembersMiddlewares(stack)
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
	addOpListMembersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListMembers(options.Region), middleware.Before)
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
			OperationName: "ListMembers",
			Err:           err,
		}
	}
	out := result.(*ListMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMembersInput struct {
	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string
	// The unique identifier of the network for which to list members.
	NetworkId *string
	// An optional Boolean value. If provided, the request is limited either to members
	// that the current AWS account owns (true) or that other AWS accounts own (false).
	// If omitted, all members are listed.
	IsOwned *bool
	// The optional name of the member to list.
	Name *string
	// The maximum number of members to return in the request.
	MaxResults *int32
	// An optional status specifier. If provided, only members currently in this status
	// are listed.
	Status types.MemberStatus
}

type ListMembersOutput struct {
	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string
	// An array of MemberSummary objects. Each object contains details about a network
	// member.
	Members []*types.MemberSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListMembersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListMembers{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListMembers{}, middleware.After)
}

func newServiceMetadataMiddleware_opListMembers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "managedblockchain",
		OperationName: "ListMembers",
	}
}
