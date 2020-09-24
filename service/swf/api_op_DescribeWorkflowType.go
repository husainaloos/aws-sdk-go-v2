// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the specified workflow type. This includes
// configuration settings specified when the type was registered and other
// information such as creation date, current status, etc. Access Control You can
// use IAM policies to control this action's access to Amazon SWF resources as
// follows:
//
//     * Use a Resource element with the domain name to limit the action
// to only specified domains.
//
//     * Use an Action element to allow or deny
// permission to call this action.
//
//     * Constrain the following parameters by
// using a Condition element with the appropriate keys.
//
//         *
// workflowType.name: String constraint. The key is swf:workflowType.name.
//
//
// * workflowType.version: String constraint. The key is
// swf:workflowType.version.
//
// If the caller doesn't have sufficient permissions to
// invoke the action, or the parameter values fall outside the specified
// constraints, the action fails. The associated event attribute's cause parameter
// is set to OPERATION_NOT_PERMITTED. For details and example IAM policies, see
// Using IAM to Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) DescribeWorkflowType(ctx context.Context, params *DescribeWorkflowTypeInput, optFns ...func(*Options)) (*DescribeWorkflowTypeOutput, error) {
	stack := middleware.NewStack("DescribeWorkflowType", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpDescribeWorkflowTypeMiddlewares(stack)
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
	addOpDescribeWorkflowTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkflowType(options.Region), middleware.Before)
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
			OperationName: "DescribeWorkflowType",
			Err:           err,
		}
	}
	out := result.(*DescribeWorkflowTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorkflowTypeInput struct {
	// The name of the domain in which this workflow type is registered.
	Domain *string
	// The workflow type to describe.
	WorkflowType *types.WorkflowType
}

// Contains details about a workflow type.
type DescribeWorkflowTypeOutput struct {
	// Configuration settings of the workflow type registered through
	// RegisterWorkflowType ()
	Configuration *types.WorkflowTypeConfiguration
	// General information about the workflow type. The status of the workflow type
	// (returned in the WorkflowTypeInfo structure) can be one of the following.
	//
	//     *
	// REGISTERED – The type is registered and available. Workers supporting this type
	// should be running.
	//
	//     * DEPRECATED – The type was deprecated using
	// DeprecateWorkflowType (), but is still in use. You should keep workers
	// supporting this type running. You cannot create new workflow executions of this
	// type.
	TypeInfo *types.WorkflowTypeInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpDescribeWorkflowTypeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeWorkflowType{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeWorkflowType{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeWorkflowType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "DescribeWorkflowType",
	}
}
