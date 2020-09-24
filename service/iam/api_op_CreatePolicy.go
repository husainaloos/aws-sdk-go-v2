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

// Creates a new managed policy for your AWS account. This operation creates a
// policy version with a version identifier of v1 and sets v1 as the policy's
// default version. For more information about policy versions, see Versioning for
// Managed Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html)
// in the IAM User Guide. For more information about managed policies in general,
// see Managed Policies and Inline Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
func (c *Client) CreatePolicy(ctx context.Context, params *CreatePolicyInput, optFns ...func(*Options)) (*CreatePolicyOutput, error) {
	stack := middleware.NewStack("CreatePolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreatePolicyMiddlewares(stack)
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
	addOpCreatePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePolicy(options.Region), middleware.Before)
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
			OperationName: "CreatePolicy",
			Err:           err,
		}
	}
	out := result.(*CreatePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePolicyInput struct {
	// The friendly name of the policy. IAM user, group, role, and policy names must be
	// unique within the account. Names are not distinguished by case. For example, you
	// cannot create resources named both "MyResource" and "myresource".
	PolicyName *string
	// The path for the policy. For more information about paths, see IAM Identifiers
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the
	// IAM User Guide. This parameter is optional. If it is not included, it defaults
	// to a slash (/). This parameter allows (through its regex pattern
	// (http://wikipedia.org/wiki/regex)) a string of characters consisting of either a
	// forward slash (/) by itself or a string that must begin and end with forward
	// slashes. In addition, it can contain any ASCII character from the ! (\u0021)
	// through the DEL character (\u007F), including most punctuation characters,
	// digits, and upper and lowercased letters.
	Path *string
	// The JSON policy document that you want to use as the content for the new policy.
	// You must provide policies in JSON format in IAM. However, for AWS CloudFormation
	// templates formatted in YAML, you can provide the policy in JSON or YAML format.
	// AWS CloudFormation always converts a YAML policy to JSON format before
	// submitting it to IAM. The regex pattern (http://wikipedia.org/wiki/regex) used
	// to validate this parameter is a string of characters consisting of the
	// following:
	//
	//     * Any printable ASCII character ranging from the space character
	// (\u0020) through the end of the ASCII character range
	//
	//     * The printable
	// characters in the Basic Latin and Latin-1 Supplement character set (through
	// \u00FF)
	//
	//     * The special characters tab (\u0009), line feed (\u000A), and
	// carriage return (\u000D)
	PolicyDocument *string
	// A friendly description of the policy. Typically used to store information about
	// the permissions defined in the policy. For example, "Grants access to production
	// DynamoDB tables." The policy description is immutable. After a value is
	// assigned, it cannot be changed.
	Description *string
}

// Contains the response to a successful CreatePolicy () request.
type CreatePolicyOutput struct {
	// A structure containing details about the new policy.
	Policy *types.Policy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreatePolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreatePolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreatePolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreatePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "CreatePolicy",
	}
}
