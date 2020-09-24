// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a user in a Simple AD or Microsoft AD directory. The status of a newly
// created user is "ACTIVE". New users can access Amazon WorkDocs.
func (c *Client) CreateUser(ctx context.Context, params *CreateUserInput, optFns ...func(*Options)) (*CreateUserOutput, error) {
	stack := middleware.NewStack("CreateUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateUserMiddlewares(stack)
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
	addOpCreateUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUser(options.Region), middleware.Before)
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
			OperationName: "CreateUser",
			Err:           err,
		}
	}
	out := result.(*CreateUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUserInput struct {
	// The given name of the user.
	GivenName *string
	// The surname of the user.
	Surname *string
	// The ID of the organization.
	OrganizationId *string
	// The login name of the user.
	Username *string
	// The time zone ID of the user.
	TimeZoneId *string
	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string
	// The amount of storage for the user.
	StorageRule *types.StorageRuleType
	// The email address of the user.
	EmailAddress *string
	// The password of the user.
	Password *string
}

type CreateUserOutput struct {
	// The user information.
	User *types.User

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateUser{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "CreateUser",
	}
}
