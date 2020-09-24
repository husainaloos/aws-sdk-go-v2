// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an Amazon QuickSight user.
func (c *Client) UpdateUser(ctx context.Context, params *UpdateUserInput, optFns ...func(*Options)) (*UpdateUserOutput, error) {
	stack := middleware.NewStack("UpdateUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateUserMiddlewares(stack)
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
	addOpUpdateUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUser(options.Region), middleware.Before)
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
			OperationName: "UpdateUser",
			Err:           err,
		}
	}
	out := result.(*UpdateUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateUserInput struct {
	// The namespace. Currently, you should set this to default.
	Namespace *string
	// The name of the custom permissions profile that you want to assign to this user.
	// Currently, custom permissions profile names are assigned to permissions profiles
	// in the QuickSight console. You use this API to assign the named set of
	// permissions to a QuickSight user.
	CustomPermissionsName *string
	// The email address of the user that you want to update.
	Email *string
	// The ID for the AWS account that the user is in. Currently, you use the ID for
	// the AWS account that contains your Amazon QuickSight account.
	AwsAccountId *string
	// The Amazon QuickSight user name that you want to update.
	UserName *string
	// A flag that you use to indicate that you want to remove all custom permissions
	// from this user. Using this parameter resets the user to the state it was in
	// before a custom permissions profile was applied. This parameter defaults to NULL
	// and it doesn't accept any other value.
	UnapplyCustomPermissions *bool
	// The Amazon QuickSight role of the user. The user role can be one of the
	// following:
	//
	//     * READER: A user who has read-only access to dashboards.
	//
	//     *
	// AUTHOR: A user who can create data sources, datasets, analyses, and
	// dashboards.
	//
	//     * ADMIN: A user who is an author, who can also manage Amazon
	// QuickSight settings.
	Role types.UserRole
}

type UpdateUserOutput struct {
	// The Amazon QuickSight user.
	User *types.User
	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateUser{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateUser",
	}
}
