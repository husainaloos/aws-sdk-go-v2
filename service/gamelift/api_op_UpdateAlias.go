// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates properties for an alias. To update properties, specify the alias ID to
// be updated and provide the information to be changed. To reassign an alias to
// another fleet, provide an updated routing strategy. If successful, the updated
// alias record is returned.
//
//     * CreateAlias ()
//
//     * ListAliases ()
//
//     *
// DescribeAlias ()
//
//     * UpdateAlias ()
//
//     * DeleteAlias ()
//
//     * ResolveAlias
// ()
func (c *Client) UpdateAlias(ctx context.Context, params *UpdateAliasInput, optFns ...func(*Options)) (*UpdateAliasOutput, error) {
	stack := middleware.NewStack("UpdateAlias", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateAliasMiddlewares(stack)
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
	addOpUpdateAliasValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAlias(options.Region), middleware.Before)
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
			OperationName: "UpdateAlias",
			Err:           err,
		}
	}
	out := result.(*UpdateAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type UpdateAliasInput struct {
	// The routing configuration, including routing type and fleet target, for the
	// alias.
	RoutingStrategy *types.RoutingStrategy
	// A human-readable description of the alias.
	Description *string
	// A descriptive label that is associated with an alias. Alias names do not need to
	// be unique.
	Name *string
	// A unique identifier for the alias that you want to update. You can use either
	// the alias ID or ARN value.
	AliasId *string
}

// Represents the returned data in response to a request action.
type UpdateAliasOutput struct {
	// The updated alias resource.
	Alias *types.Alias

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateAliasMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAlias{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAlias{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateAlias(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "UpdateAlias",
	}
}
