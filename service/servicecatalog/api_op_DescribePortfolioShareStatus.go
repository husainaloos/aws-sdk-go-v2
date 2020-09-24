// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the status of the specified portfolio share operation. This API can only be
// called by the master account in the organization or by a delegated admin.
func (c *Client) DescribePortfolioShareStatus(ctx context.Context, params *DescribePortfolioShareStatusInput, optFns ...func(*Options)) (*DescribePortfolioShareStatusOutput, error) {
	stack := middleware.NewStack("DescribePortfolioShareStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribePortfolioShareStatusMiddlewares(stack)
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
	addOpDescribePortfolioShareStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePortfolioShareStatus(options.Region), middleware.Before)
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
			OperationName: "DescribePortfolioShareStatus",
			Err:           err,
		}
	}
	out := result.(*DescribePortfolioShareStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePortfolioShareStatusInput struct {
	// The token for the portfolio share operation. This token is returned either by
	// CreatePortfolioShare or by DeletePortfolioShare.
	PortfolioShareToken *string
}

type DescribePortfolioShareStatusOutput struct {
	// Organization node identifier. It can be either account id, organizational unit
	// id or organization id.
	OrganizationNodeValue *string
	// The portfolio identifier.
	PortfolioId *string
	// Information about the portfolio share operation.
	ShareDetails *types.ShareDetails
	// The token for the portfolio share operation. For example, share-6v24abcdefghi.
	PortfolioShareToken *string
	// Status of the portfolio share operation.
	Status types.ShareStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribePortfolioShareStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePortfolioShareStatus{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePortfolioShareStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribePortfolioShareStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "DescribePortfolioShareStatus",
	}
}
