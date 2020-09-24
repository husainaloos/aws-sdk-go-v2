// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Rejects the transfer of a domain from another AWS account to the current AWS
// account. You initiate a transfer between AWS accounts using
// TransferDomainToAnotherAwsAccount
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html).
// <p>Use either <a
// href="https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ListOperations.html">ListOperations</a>
// or <a
// href="https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html">GetOperationDetail</a>
// to determine whether the operation succeeded. <a
// href="https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html">GetOperationDetail</a>
// provides additional information, for example, <code>Domain Transfer from Aws
// Account 111122223333 has been cancelled</code>. </p>
func (c *Client) RejectDomainTransferFromAnotherAwsAccount(ctx context.Context, params *RejectDomainTransferFromAnotherAwsAccountInput, optFns ...func(*Options)) (*RejectDomainTransferFromAnotherAwsAccountOutput, error) {
	stack := middleware.NewStack("RejectDomainTransferFromAnotherAwsAccount", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRejectDomainTransferFromAnotherAwsAccountMiddlewares(stack)
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
	addOpRejectDomainTransferFromAnotherAwsAccountValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRejectDomainTransferFromAnotherAwsAccount(options.Region), middleware.Before)
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
			OperationName: "RejectDomainTransferFromAnotherAwsAccount",
			Err:           err,
		}
	}
	out := result.(*RejectDomainTransferFromAnotherAwsAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The RejectDomainTransferFromAnotherAwsAccount request includes the following
// element.
type RejectDomainTransferFromAnotherAwsAccountInput struct {
	// The name of the domain that was specified when another AWS account submitted a
	// TransferDomainToAnotherAwsAccount
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html)
	// request.
	DomainName *string
}

// The RejectDomainTransferFromAnotherAwsAccount response includes the following
// element.
type RejectDomainTransferFromAnotherAwsAccountOutput struct {
	// The identifier that TransferDomainToAnotherAwsAccount returned to track the
	// progress of the request. Because the transfer request was rejected, the value is
	// no longer valid, and you can't use GetOperationDetail to query the operation
	// status.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRejectDomainTransferFromAnotherAwsAccountMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRejectDomainTransferFromAnotherAwsAccount{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRejectDomainTransferFromAnotherAwsAccount{}, middleware.After)
}

func newServiceMetadataMiddleware_opRejectDomainTransferFromAnotherAwsAccount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "RejectDomainTransferFromAnotherAwsAccount",
	}
}
