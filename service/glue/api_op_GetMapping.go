// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates mappings.
func (c *Client) GetMapping(ctx context.Context, params *GetMappingInput, optFns ...func(*Options)) (*GetMappingOutput, error) {
	if params == nil {
		params = &GetMappingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMapping", params, optFns, addOperationGetMappingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMappingInput struct {

	// Specifies the source table.
	//
	// This member is required.
	Source *types.CatalogEntry

	// Parameters for the mapping.
	Location *types.Location

	// A list of target tables.
	Sinks []*types.CatalogEntry
}

type GetMappingOutput struct {

	// A list of mappings to the specified targets.
	//
	// This member is required.
	Mapping []*types.MappingEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetMappingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMapping{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMapping{}, middleware.After)
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
	addOpGetMappingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMapping(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetMapping(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetMapping",
	}
}
