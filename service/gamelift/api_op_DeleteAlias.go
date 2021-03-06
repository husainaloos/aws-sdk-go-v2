// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an alias. This action removes all record of the alias. Game clients
// attempting to access a server process using the deleted alias receive an error.
// To delete an alias, specify the alias ID to be deleted.
//
//     * CreateAlias ()
//
//
// * ListAliases ()
//
//     * DescribeAlias ()
//
//     * UpdateAlias ()
//
//     *
// DeleteAlias ()
//
//     * ResolveAlias ()
func (c *Client) DeleteAlias(ctx context.Context, params *DeleteAliasInput, optFns ...func(*Options)) (*DeleteAliasOutput, error) {
	if params == nil {
		params = &DeleteAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAlias", params, optFns, addOperationDeleteAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type DeleteAliasInput struct {

	// A unique identifier of the alias that you want to delete. You can use either the
	// alias ID or ARN value.
	//
	// This member is required.
	AliasId *string
}

type DeleteAliasOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteAlias{}, middleware.After)
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
	addOpDeleteAliasValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAlias(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteAlias(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DeleteAlias",
	}
}
