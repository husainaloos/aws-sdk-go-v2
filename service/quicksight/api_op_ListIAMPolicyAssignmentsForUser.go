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

// Lists all the IAM policy assignments, including the Amazon Resource Names (ARNs)
// for the IAM policies assigned to the specified user and group or groups that the
// user belongs to.
func (c *Client) ListIAMPolicyAssignmentsForUser(ctx context.Context, params *ListIAMPolicyAssignmentsForUserInput, optFns ...func(*Options)) (*ListIAMPolicyAssignmentsForUserOutput, error) {
	stack := middleware.NewStack("ListIAMPolicyAssignmentsForUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListIAMPolicyAssignmentsForUserMiddlewares(stack)
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
	addOpListIAMPolicyAssignmentsForUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListIAMPolicyAssignmentsForUser(options.Region), middleware.Before)
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
			OperationName: "ListIAMPolicyAssignmentsForUser",
			Err:           err,
		}
	}
	out := result.(*ListIAMPolicyAssignmentsForUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIAMPolicyAssignmentsForUserInput struct {
	// The ID of the AWS account that contains the assignments.
	AwsAccountId *string
	// The namespace of the assignment.
	Namespace *string
	// The name of the user.
	UserName *string
	// The token for the next set of results, or null if there are no more results.
	NextToken *string
	// The maximum number of results to be returned per request.
	MaxResults *int32
}

type ListIAMPolicyAssignmentsForUserOutput struct {
	// The token for the next set of results, or null if there are no more results.
	NextToken *string
	// The AWS request ID for this operation.
	RequestId *string
	// The active assignments for this user.
	ActiveAssignments []*types.ActiveIAMPolicyAssignment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListIAMPolicyAssignmentsForUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListIAMPolicyAssignmentsForUser{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListIAMPolicyAssignmentsForUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opListIAMPolicyAssignmentsForUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListIAMPolicyAssignmentsForUser",
	}
}
