// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about the specified IAM user, including the user's
// creation date, path, unique ID, and ARN. If you do not specify a user name, IAM
// determines the user name implicitly based on the AWS access key ID used to sign
// the request to this API.
func (c *Client) GetUser(ctx context.Context, params *GetUserInput, optFns ...func(*Options)) (*GetUserOutput, error) {
	stack := middleware.NewStack("GetUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetUserMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetUser(options.Region), middleware.Before)
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
			OperationName: "GetUser",
			Err:           err,
		}
	}
	out := result.(*GetUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUserInput struct {
	// The name of the user to get information about. This parameter is optional. If it
	// is not included, it defaults to the user making the request. This parameter
	// allows (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	UserName *string
}

// Contains the response to a successful GetUser () request.
type GetUserOutput struct {
	// A structure containing details about the IAM user. Due to a service issue,
	// password last used data does not include password use from May 3, 2018 22:50 PDT
	// to May 23, 2018 14:08 PDT. This affects last sign-in
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_finding-unused.html)
	// dates shown in the IAM console and password last used dates in the IAM
	// credential report
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_getting-report.html),
	// and returned by this GetUser API. If users signed in during the affected time,
	// the password last used date that is returned is the date the user last signed in
	// before May 3, 2018. For users that signed in after May 23, 2018 14:08 PDT, the
	// returned password last used date is accurate. You can use password last used
	// information to identify unused credentials for deletion. For example, you might
	// delete users who did not sign in to AWS in the last 90 days. In cases like this,
	// we recommend that you adjust your evaluation window to include dates after May
	// 23, 2018. Alternatively, if your users use access keys to access AWS
	// programmatically you can refer to access key last used information because it is
	// accurate for all dates.
	User *types.User

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetUser{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GetUser",
	}
}
