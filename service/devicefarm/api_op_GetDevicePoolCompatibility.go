// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about compatibility with a device pool.
func (c *Client) GetDevicePoolCompatibility(ctx context.Context, params *GetDevicePoolCompatibilityInput, optFns ...func(*Options)) (*GetDevicePoolCompatibilityOutput, error) {
	stack := middleware.NewStack("GetDevicePoolCompatibility", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetDevicePoolCompatibilityMiddlewares(stack)
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
	addOpGetDevicePoolCompatibilityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDevicePoolCompatibility(options.Region), middleware.Before)
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
			OperationName: "GetDevicePoolCompatibility",
			Err:           err,
		}
	}
	out := result.(*GetDevicePoolCompatibilityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to the get device pool compatibility operation.
type GetDevicePoolCompatibilityInput struct {
	// Information about the uploaded test to be run against the device pool.
	Test *types.ScheduleRunTest
	// The device pool's ARN.
	DevicePoolArn *string
	// An object that contains information about the settings for a run.
	Configuration *types.ScheduleRunConfiguration
	// The ARN of the app that is associated with the specified device pool.
	AppArn *string
	// The test type for the specified device pool. Allowed values include the
	// following:
	//
	//     * BUILTIN_FUZZ.
	//
	//     * BUILTIN_EXPLORER. For Android, an app
	// explorer that traverses an Android app, interacting with it and capturing
	// screenshots at the same time.
	//
	//     * APPIUM_JAVA_JUNIT.
	//
	//     *
	// APPIUM_JAVA_TESTNG.
	//
	//     * APPIUM_PYTHON.
	//
	//     * APPIUM_NODE.
	//
	//     *
	// APPIUM_RUBY.
	//
	//     * APPIUM_WEB_JAVA_JUNIT.
	//
	//     * APPIUM_WEB_JAVA_TESTNG.
	//
	//     *
	// APPIUM_WEB_PYTHON.
	//
	//     * APPIUM_WEB_NODE.
	//
	//     * APPIUM_WEB_RUBY.
	//
	//     *
	// CALABASH.
	//
	//     * INSTRUMENTATION.
	//
	//     * UIAUTOMATION.
	//
	//     * UIAUTOMATOR.
	//
	//
	// * XCTEST.
	//
	//     * XCTEST_UI.
	TestType types.TestType
}

// Represents the result of describe device pool compatibility request.
type GetDevicePoolCompatibilityOutput struct {
	// Information about incompatible devices.
	IncompatibleDevices []*types.DevicePoolCompatibilityResult
	// Information about compatible devices.
	CompatibleDevices []*types.DevicePoolCompatibilityResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetDevicePoolCompatibilityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetDevicePoolCompatibility{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDevicePoolCompatibility{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDevicePoolCompatibility(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "GetDevicePoolCompatibility",
	}
}
