// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is documentation for AWS CloudHSM Classic. For more information, see AWS
// CloudHSM Classic FAQs (http://aws.amazon.com/cloudhsm/faqs-classic/), the AWS
// CloudHSM Classic User Guide
// (https://docs.aws.amazon.com/cloudhsm/classic/userguide/), and the AWS CloudHSM
// Classic API Reference
// (https://docs.aws.amazon.com/cloudhsm/classic/APIReference/). For information
// about the current version of AWS CloudHSM, see AWS CloudHSM
// (http://aws.amazon.com/cloudhsm/), the AWS CloudHSM User Guide
// (https://docs.aws.amazon.com/cloudhsm/latest/userguide/), and the AWS CloudHSM
// API Reference (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/). Gets
// the configuration files necessary to connect to all high availability partition
// groups the client is associated with.
func (c *Client) GetConfig(ctx context.Context, params *GetConfigInput, optFns ...func(*Options)) (*GetConfigOutput, error) {
	if params == nil {
		params = &GetConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConfig", params, optFns, addOperationGetConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConfigInput struct {

	// The ARN of the client.
	//
	// This member is required.
	ClientArn *string

	// The client version.
	//
	// This member is required.
	ClientVersion types.ClientVersion

	// A list of ARNs that identify the high-availability partition groups that are
	// associated with the client.
	//
	// This member is required.
	HapgList []*string
}

type GetConfigOutput struct {

	// The certificate file containing the server.pem files of the HSMs.
	ConfigCred *string

	// The chrystoki.conf configuration file.
	ConfigFile *string

	// The type of credentials.
	ConfigType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetConfig{}, middleware.After)
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
	addOpGetConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetConfig(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudhsm",
		OperationName: "GetConfig",
	}
}
