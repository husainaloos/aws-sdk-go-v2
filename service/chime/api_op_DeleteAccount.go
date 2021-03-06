// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified Amazon Chime account. You must suspend all users before
// deleting a Team account. You can use the BatchSuspendUser () action to do so.
// For EnterpriseLWA and EnterpriseAD accounts, you must release the claimed
// domains for your Amazon Chime account before deletion. As soon as you release
// the domain, all users under that account are suspended. Deleted accounts appear
// in your Disabled accounts list for 90 days. To restore a deleted account from
// your Disabled accounts list, you must contact AWS Support. After 90 days,
// deleted accounts are permanently removed from your Disabled accounts list.
func (c *Client) DeleteAccount(ctx context.Context, params *DeleteAccountInput, optFns ...func(*Options)) (*DeleteAccountOutput, error) {
	if params == nil {
		params = &DeleteAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAccount", params, optFns, addOperationDeleteAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAccountInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string
}

type DeleteAccountOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteAccount{}, middleware.After)
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
	addOpDeleteAccountValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAccount(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteAccount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "DeleteAccount",
	}
}
