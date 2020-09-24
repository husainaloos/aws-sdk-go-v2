// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Directory () by copying the published schema into the directory. A
// directory cannot be created without a schema. You can also quickly create a
// directory using a managed schema, called the QuickStartSchema. For more
// information, see Managed Schema
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/schemas_managed.html)
// in the Amazon Cloud Directory Developer Guide.
func (c *Client) CreateDirectory(ctx context.Context, params *CreateDirectoryInput, optFns ...func(*Options)) (*CreateDirectoryOutput, error) {
	stack := middleware.NewStack("CreateDirectory", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDirectoryMiddlewares(stack)
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
	addOpCreateDirectoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDirectory(options.Region), middleware.Before)
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
			OperationName: "CreateDirectory",
			Err:           err,
		}
	}
	out := result.(*CreateDirectoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDirectoryInput struct {
	// The Amazon Resource Name (ARN) of the published schema that will be copied into
	// the data Directory (). For more information, see arns ().
	SchemaArn *string
	// The name of the Directory (). Should be unique per account, per region.
	Name *string
}

type CreateDirectoryOutput struct {
	// The name of the Directory ().
	Name *string
	// The ARN of the published schema in the Directory (). Once a published schema is
	// copied into the directory, it has its own ARN, which is referred to applied
	// schema ARN. For more information, see arns ().
	AppliedSchemaArn *string
	// The ARN that is associated with the Directory (). For more information, see arns
	// ().
	DirectoryArn *string
	// The root object node of the created directory.
	ObjectIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDirectoryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDirectory{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDirectory{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDirectory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "CreateDirectory",
	}
}
