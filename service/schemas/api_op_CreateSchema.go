// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/schemas/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a schema definition. Inactive schemas will be deleted after two years.
func (c *Client) CreateSchema(ctx context.Context, params *CreateSchemaInput, optFns ...func(*Options)) (*CreateSchemaOutput, error) {
	stack := middleware.NewStack("CreateSchema", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateSchemaMiddlewares(stack)
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
	addOpCreateSchemaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSchema(options.Region), middleware.Before)
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
			OperationName: "CreateSchema",
			Err:           err,
		}
	}
	out := result.(*CreateSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSchemaInput struct {
	// The source of the schema definition.
	Content *string
	// The name of the schema.
	SchemaName *string
	// Tags associated with the schema.
	Tags map[string]*string
	// The type of schema.
	Type types.Type
	// A description of the schema.
	Description *string
	// The name of the registry.
	RegistryName *string
}

type CreateSchemaOutput struct {
	// The name of the schema.
	SchemaName *string
	// The ARN of the schema.
	SchemaArn *string
	// The date the schema version was created.
	VersionCreatedDate *time.Time
	// The version number of the schema
	SchemaVersion *string
	// The description of the schema.
	Description *string
	// The date and time that schema was modified.
	LastModified *time.Time
	// The type of the schema.
	Type *string
	// Key-value pairs associated with a resource.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateSchemaMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateSchema{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSchema{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateSchema(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "schemas",
		OperationName: "CreateSchema",
	}
}
