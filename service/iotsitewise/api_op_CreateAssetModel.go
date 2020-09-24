// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an asset model from specified property and hierarchy definitions. You
// create assets from asset models. With asset models, you can easily create assets
// of the same type that have standardized definitions. Each asset created from a
// model inherits the asset model's property and hierarchy definitions. For more
// information, see Defining Asset Models
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/define-models.html)
// in the AWS IoT SiteWise User Guide.
func (c *Client) CreateAssetModel(ctx context.Context, params *CreateAssetModelInput, optFns ...func(*Options)) (*CreateAssetModelOutput, error) {
	stack := middleware.NewStack("CreateAssetModel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateAssetModelMiddlewares(stack)
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
	addIdempotencyToken_opCreateAssetModelMiddleware(stack, options)
	addOpCreateAssetModelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAssetModel(options.Region), middleware.Before)
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
			OperationName: "CreateAssetModel",
			Err:           err,
		}
	}
	out := result.(*CreateAssetModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAssetModelInput struct {
	// A description for the asset model.
	AssetModelDescription *string
	// The property definitions of the asset model. For more information, see Asset
	// Properties
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-properties.html)
	// in the AWS IoT SiteWise User Guide. You can specify up to 200 properties per
	// asset model. For more information, see Quotas
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the
	// AWS IoT SiteWise User Guide.
	AssetModelProperties []*types.AssetModelPropertyDefinition
	// A unique case-sensitive identifier that you can provide to ensure the
	// idempotency of the request. Don't reuse this client token if a new idempotent
	// request is required.
	ClientToken *string
	// A unique, friendly name for the asset model.
	AssetModelName *string
	// A list of key-value pairs that contain metadata for the asset model. For more
	// information, see Tagging your AWS IoT SiteWise resources
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html)
	// in the AWS IoT SiteWise User Guide.
	Tags map[string]*string
	// The hierarchy definitions of the asset model. Each hierarchy specifies an asset
	// model whose assets can be children of any other assets created from this asset
	// model. For more information, see Asset Hierarchies
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-hierarchies.html)
	// in the AWS IoT SiteWise User Guide. You can specify up to 10 hierarchies per
	// asset model. For more information, see Quotas
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the
	// AWS IoT SiteWise User Guide.
	AssetModelHierarchies []*types.AssetModelHierarchyDefinition
}

type CreateAssetModelOutput struct {
	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the asset model, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:asset-model/${AssetModelId}
	AssetModelArn *string
	// The status of the asset model, which contains a state (CREATING after
	// successfully calling this operation) and any error message.
	AssetModelStatus *types.AssetModelStatus
	// The ID of the asset model. You can use this ID when you call other AWS IoT
	// SiteWise APIs.
	AssetModelId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateAssetModelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateAssetModel{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAssetModel{}, middleware.After)
}

type idempotencyToken_initializeOpCreateAssetModel struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAssetModel) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAssetModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAssetModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAssetModelInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateAssetModelMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateAssetModel{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAssetModel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "CreateAssetModel",
	}
}
