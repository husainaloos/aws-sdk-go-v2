// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gathers a complete list of on-premises servers. Connectors must be installed and
// monitoring all servers that you want to import. This call returns immediately,
// but might take additional time to retrieve all the servers.
func (c *Client) ImportServerCatalog(ctx context.Context, params *ImportServerCatalogInput, optFns ...func(*Options)) (*ImportServerCatalogOutput, error) {
	stack := middleware.NewStack("ImportServerCatalog", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpImportServerCatalogMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opImportServerCatalog(options.Region), middleware.Before)
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
			OperationName: "ImportServerCatalog",
			Err:           err,
		}
	}
	out := result.(*ImportServerCatalogOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportServerCatalogInput struct {
}

type ImportServerCatalogOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpImportServerCatalogMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpImportServerCatalog{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportServerCatalog{}, middleware.After)
}

func newServiceMetadataMiddleware_opImportServerCatalog(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "ImportServerCatalog",
	}
}
