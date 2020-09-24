// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the versions of a connector definition, which are containers for
// connectors. Connectors run on the Greengrass core and contain built-in
// integration with local infrastructure, device protocols, AWS, and other cloud
// services.
func (c *Client) ListConnectorDefinitionVersions(ctx context.Context, params *ListConnectorDefinitionVersionsInput, optFns ...func(*Options)) (*ListConnectorDefinitionVersionsOutput, error) {
	stack := middleware.NewStack("ListConnectorDefinitionVersions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListConnectorDefinitionVersionsMiddlewares(stack)
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
	addOpListConnectorDefinitionVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListConnectorDefinitionVersions(options.Region), middleware.Before)
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
			OperationName: "ListConnectorDefinitionVersions",
			Err:           err,
		}
	}
	out := result.(*ListConnectorDefinitionVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConnectorDefinitionVersionsInput struct {
	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string
	// The ID of the connector definition.
	ConnectorDefinitionId *string
	// The maximum number of results to be returned per request.
	MaxResults *string
}

type ListConnectorDefinitionVersionsOutput struct {
	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string
	// Information about a version.
	Versions []*types.VersionInformation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListConnectorDefinitionVersionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListConnectorDefinitionVersions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListConnectorDefinitionVersions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListConnectorDefinitionVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "ListConnectorDefinitionVersions",
	}
}
