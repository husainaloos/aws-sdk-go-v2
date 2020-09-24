// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a model package that you can use to create Amazon SageMaker models or
// list on AWS Marketplace. Buyers can subscribe to model packages listed on AWS
// Marketplace to create models in Amazon SageMaker. To create a model package by
// specifying a Docker container that contains your inference code and the Amazon
// S3 location of your model artifacts, provide values for InferenceSpecification.
// To create a model from an algorithm resource that you created or subscribed to
// in AWS Marketplace, provide a value for SourceAlgorithmSpecification.
func (c *Client) CreateModelPackage(ctx context.Context, params *CreateModelPackageInput, optFns ...func(*Options)) (*CreateModelPackageOutput, error) {
	stack := middleware.NewStack("CreateModelPackage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateModelPackageMiddlewares(stack)
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
	addOpCreateModelPackageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateModelPackage(options.Region), middleware.Before)
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
			OperationName: "CreateModelPackage",
			Err:           err,
		}
	}
	out := result.(*CreateModelPackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateModelPackageInput struct {
	// Details about the algorithm that was used to create the model package.
	SourceAlgorithmSpecification *types.SourceAlgorithmSpecification
	// Specifies configurations for one or more transform jobs that Amazon SageMaker
	// runs to test the model package.
	ValidationSpecification *types.ModelPackageValidationSpecification
	// A description of the model package.
	ModelPackageDescription *string
	// Specifies details about inference jobs that can be run with models based on this
	// model package, including the following:
	//
	//     * The Amazon ECR paths of
	// containers that contain the inference code and model artifacts.
	//
	//     * The
	// instance types that the model package supports for transform jobs and real-time
	// endpoints used for inference.
	//
	//     * The input and output content formats that
	// the model package supports for inference.
	InferenceSpecification *types.InferenceSpecification
	// The name of the model package. The name must have 1 to 63 characters. Valid
	// characters are a-z, A-Z, 0-9, and - (hyphen).
	ModelPackageName *string
	// Whether to certify the model package for listing on AWS Marketplace.
	CertifyForMarketplace *bool
}

type CreateModelPackageOutput struct {
	// The Amazon Resource Name (ARN) of the new model package.
	ModelPackageArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateModelPackageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateModelPackage{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateModelPackage{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateModelPackage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateModelPackage",
	}
}
