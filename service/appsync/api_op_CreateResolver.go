// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Resolver object. A resolver converts incoming requests into a format
// that a data source can understand and converts the data source's responses into
// GraphQL.
func (c *Client) CreateResolver(ctx context.Context, params *CreateResolverInput, optFns ...func(*Options)) (*CreateResolverOutput, error) {
	stack := middleware.NewStack("CreateResolver", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateResolverMiddlewares(stack)
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
	addOpCreateResolverValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResolver(options.Region), middleware.Before)
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
			OperationName: "CreateResolver",
			Err:           err,
		}
	}
	out := result.(*CreateResolverOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResolverInput struct {
	// The ID for the GraphQL API for which the resolver is being created.
	ApiId *string
	// The name of the data source for which the resolver is being created.
	DataSourceName *string
	// The name of the field to attach the resolver to.
	FieldName *string
	// The mapping template to be used for responses from the data source.
	ResponseMappingTemplate *string
	// The caching configuration for the resolver.
	CachingConfig *types.CachingConfig
	// The mapping template to be used for requests. A resolver uses a request mapping
	// template to convert a GraphQL expression into a format that a data source can
	// understand. Mapping templates are written in Apache Velocity Template Language
	// (VTL).
	RequestMappingTemplate *string
	// The SyncConfig for a resolver attached to a versioned datasource.
	SyncConfig *types.SyncConfig
	// The name of the Type.
	TypeName *string
	// The resolver type.
	//
	//     * UNIT: A UNIT resolver type. A UNIT resolver is the
	// default resolver type. A UNIT resolver enables you to execute a GraphQL query
	// against a single data source.
	//
	//     * PIPELINE: A PIPELINE resolver type. A
	// PIPELINE resolver enables you to execute a series of Function in a serial
	// manner. You can use a pipeline resolver to execute a GraphQL query against
	// multiple data sources.
	Kind types.ResolverKind
	// The PipelineConfig.
	PipelineConfig *types.PipelineConfig
}

type CreateResolverOutput struct {
	// The Resolver object.
	Resolver *types.Resolver

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateResolverMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateResolver{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateResolver{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateResolver(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "CreateResolver",
	}
}
