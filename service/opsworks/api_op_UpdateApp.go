// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a specified app. Required Permissions: To use this action, an IAM user
// must have a Deploy or Manage permissions level for the stack, or an attached
// policy that explicitly grants permissions. For more information on user
// permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) UpdateApp(ctx context.Context, params *UpdateAppInput, optFns ...func(*Options)) (*UpdateAppOutput, error) {
	stack := middleware.NewStack("UpdateApp", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateAppMiddlewares(stack)
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
	addOpUpdateAppValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApp(options.Region), middleware.Before)
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
			OperationName: "UpdateApp",
			Err:           err,
		}
	}
	out := result.(*UpdateAppOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAppInput struct {
	// A Source object that specifies the app repository.
	AppSource *types.Source
	// An array of EnvironmentVariable objects that specify environment variables to be
	// associated with the app. After you deploy the app, these variables are defined
	// on the associated app server instances.For more information, see  Environment
	// Variables
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workingapps-creating.html#workingapps-creating-environment).
	// There is no specific limit on the number of environment variables. However, the
	// size of the associated data structure - which includes the variables' names,
	// values, and protected flag values - cannot exceed 20 KB. This limit should
	// accommodate most if not all use cases. Exceeding it will cause an exception with
	// the message, "Environment: is too large (maximum is 20 KB)." If you have
	// specified one or more environment variables, you cannot modify the stack's Chef
	// version.
	Environment []*types.EnvironmentVariable
	// One or more user-defined key/value pairs to be added to the stack attributes.
	Attributes map[string]*string
	// Whether SSL is enabled for the app.
	EnableSsl *bool
	// The app name.
	Name *string
	// The app type.
	Type types.AppType
	// The app's data sources.
	DataSources []*types.DataSource
	// The app ID.
	AppId *string
	// The app's virtual host settings, with multiple domains separated by commas. For
	// example: 'www.example.com, example.com'
	Domains []*string
	// A description of the app.
	Description *string
	// An SslConfiguration object with the SSL configuration.
	SslConfiguration *types.SslConfiguration
}

type UpdateAppOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateAppMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateApp{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateApp{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateApp(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "UpdateApp",
	}
}
