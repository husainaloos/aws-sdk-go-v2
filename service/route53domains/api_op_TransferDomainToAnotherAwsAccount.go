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

// Transfers a domain from the current AWS account to another AWS account. Note the
// following:
//
//     * The AWS account that you're transferring the domain to must
// accept the transfer. If the other account doesn't accept the transfer within 3
// days, we cancel the transfer. See AcceptDomainTransferFromAnotherAwsAccount
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_AcceptDomainTransferFromAnotherAwsAccount.html).
//
//
// * You can cancel the transfer before the other account accepts it. See
// CancelDomainTransferToAnotherAwsAccount
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_CancelDomainTransferToAnotherAwsAccount.html).
//
//
// * The other account can reject the transfer. See
// RejectDomainTransferFromAnotherAwsAccount
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_RejectDomainTransferFromAnotherAwsAccount.html).
//
//
// <important> <p>When you transfer a domain from one AWS account to another, Route
// 53 doesn't transfer the hosted zone that is associated with the domain. DNS
// resolution isn't affected if the domain and the hosted zone are owned by
// separate accounts, so transferring the hosted zone is optional. For information
// about transferring the hosted zone to another AWS account, see <a
// href="https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/hosted-zones-migrating.html">Migrating
// a Hosted Zone to a Different AWS Account</a> in the <i>Amazon Route 53 Developer
// Guide</i>.</p> </important> <p>Use either <a
// href="https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ListOperations.html">ListOperations</a>
// or <a
// href="https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html">GetOperationDetail</a>
// to determine whether the operation succeeded. <a
// href="https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html">GetOperationDetail</a>
// provides additional information, for example, <code>Domain Transfer from Aws
// Account 111122223333 has been cancelled</code>. </p>
func (c *Client) TransferDomainToAnotherAwsAccount(ctx context.Context, params *TransferDomainToAnotherAwsAccountInput, optFns ...func(*Options)) (*TransferDomainToAnotherAwsAccountOutput, error) {
	stack := middleware.NewStack("TransferDomainToAnotherAwsAccount", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpTransferDomainToAnotherAwsAccountMiddlewares(stack)
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
	addOpTransferDomainToAnotherAwsAccountValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTransferDomainToAnotherAwsAccount(options.Region), middleware.Before)
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
			OperationName: "TransferDomainToAnotherAwsAccount",
			Err:           err,
		}
	}
	out := result.(*TransferDomainToAnotherAwsAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The TransferDomainToAnotherAwsAccount request includes the following elements.
type TransferDomainToAnotherAwsAccountInput struct {
	// The account ID of the AWS account that you want to transfer the domain to, for
	// example, 111122223333.
	AccountId *string
	// The name of the domain that you want to transfer from the current AWS account to
	// another account.
	DomainName *string
}

// The TransferDomainToAnotherAwsAccount response includes the following elements.
type TransferDomainToAnotherAwsAccountOutput struct {
	// Identifier for tracking the progress of the request. To query the operation
	// status, use GetOperationDetail
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html).
	OperationId *string
	// To finish transferring a domain to another AWS account, the account that the
	// domain is being transferred to must submit an
	// AcceptDomainTransferFromAnotherAwsAccount
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_AcceptDomainTransferFromAnotherAwsAccount.html)
	// request. The request must include the value of the Password element that was
	// returned in the TransferDomainToAnotherAwsAccount response.
	Password *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpTransferDomainToAnotherAwsAccountMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpTransferDomainToAnotherAwsAccount{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpTransferDomainToAnotherAwsAccount{}, middleware.After)
}

func newServiceMetadataMiddleware_opTransferDomainToAnotherAwsAccount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "TransferDomainToAnotherAwsAccount",
	}
}
