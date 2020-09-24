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

// Returns information about the specified activity type. This includes
// configuration settings provided when the type was registered and other general
// information about the type. Access Control You can use IAM policies to control
// this action's access to Amazon SWF resources as follows:
//
//     * Use a Resource
// element with the domain name to limit the action to only specified domains.
//
//
// * Use an Action element to allow or deny permission to call this action.
//
//     *
// Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
//         * activityType.name: String constraint. The key is
// swf:activityType.name.
//
//         * activityType.version: String constraint. The
// key is swf:activityType.version.
//
// If the caller doesn't have sufficient
// permissions to invoke the action, or the parameter values fall outside the
// specified constraints, the action fails. The associated event attribute's cause
// parameter is set to OPERATION_NOT_PERMITTED. For details and example IAM
// policies, see Using IAM to Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) DescribeActivityType(ctx context.Context, params *DescribeActivityTypeInput, optFns ...func(*Options)) (*DescribeActivityTypeOutput, error) {
	stack := middleware.NewStack("DescribeActivityType", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpDescribeActivityTypeMiddlewares(stack)
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
	addOpDescribeActivityTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeActivityType(options.Region), middleware.Before)
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
			OperationName: "DescribeActivityType",
			Err:           err,
		}
	}
	out := result.(*DescribeActivityTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeActivityTypeInput struct {
	// The name of the domain in which the activity type is registered.
	Domain *string
	// The activity type to get information about. Activity types are identified by the
	// name and version that were supplied when the activity was registered.
	ActivityType *types.ActivityType
}

// Detailed information about an activity type.
type DescribeActivityTypeOutput struct {
	// General information about the activity type. The status of activity type
	// (returned in the ActivityTypeInfo structure) can be one of the following.
	//
	//     *
	// REGISTERED – The type is registered and available. Workers supporting this type
	// should be running.
	//
	//     * DEPRECATED – The type was deprecated using
	// DeprecateActivityType (), but is still in use. You should keep workers
	// supporting this type running. You cannot create new tasks of this type.
	TypeInfo *types.ActivityTypeInfo
	// The configuration settings registered with the activity type.
	Configuration *types.ActivityTypeConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpDescribeActivityTypeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeActivityType{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeActivityType{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeActivityType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "DescribeActivityType",
	}
}
