// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Assigns the specified security profiles to the specified user.
func (c *Client) UpdateUserSecurityProfiles(ctx context.Context, params *UpdateUserSecurityProfilesInput, optFns ...func(*Options)) (*UpdateUserSecurityProfilesOutput, error) {
	stack := middleware.NewStack("UpdateUserSecurityProfiles", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateUserSecurityProfilesMiddlewares(stack)
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
	addOpUpdateUserSecurityProfilesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUserSecurityProfiles(options.Region), middleware.Before)
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
			OperationName: "UpdateUserSecurityProfiles",
			Err:           err,
		}
	}
	out := result.(*UpdateUserSecurityProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateUserSecurityProfilesInput struct {
	// The identifiers of the security profiles for the user.
	SecurityProfileIds []*string
	// The identifier of the user account.
	UserId *string
	// The identifier of the Amazon Connect instance.
	InstanceId *string
}

type UpdateUserSecurityProfilesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateUserSecurityProfilesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateUserSecurityProfiles{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateUserSecurityProfiles{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateUserSecurityProfiles(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "UpdateUserSecurityProfiles",
	}
}
