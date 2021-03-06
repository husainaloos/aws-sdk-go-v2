// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of tasks for a specified cluster. You can filter the results by
// family name, by a particular container instance, or by the desired status of the
// task with the family, containerInstance, and desiredStatus parameters. Recently
// stopped tasks might appear in the returned results. Currently, stopped tasks
// appear in the returned results for at least one hour.
func (c *Client) ListTasks(ctx context.Context, params *ListTasksInput, optFns ...func(*Options)) (*ListTasksOutput, error) {
	if params == nil {
		params = &ListTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTasks", params, optFns, addOperationListTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTasksInput struct {

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// tasks to list. If you do not specify a cluster, the default cluster is assumed.
	Cluster *string

	// The container instance ID or full ARN of the container instance with which to
	// filter the ListTasks results. Specifying a containerInstance limits the results
	// to tasks that belong to that container instance.
	ContainerInstance *string

	// The task desired status with which to filter the ListTasks results. Specifying a
	// desiredStatus of STOPPED limits the results to tasks that Amazon ECS has set the
	// desired status to STOPPED. This can be useful for debugging tasks that are not
	// starting properly or have died or finished. The default status filter is
	// RUNNING, which shows tasks that Amazon ECS has set the desired status to
	// RUNNING. Although you can filter results based on a desired status of PENDING,
	// this does not return any results. Amazon ECS never sets the desired status of a
	// task to that value (only a task's lastStatus may have a value of PENDING).
	DesiredStatus types.DesiredStatus

	// The name of the family with which to filter the ListTasks results. Specifying a
	// family limits the results to tasks that belong to that family.
	Family *string

	// The launch type for services to list.
	LaunchType types.LaunchType

	// The maximum number of task results returned by ListTasks in paginated output.
	// When this parameter is used, ListTasks only returns maxResults results in a
	// single page along with a nextToken response element. The remaining results of
	// the initial request can be seen by sending another ListTasks request with the
	// returned nextToken value. This value can be between 1 and 100. If this parameter
	// is not used, then ListTasks returns up to 100 results and a nextToken value if
	// applicable.
	MaxResults *int32

	// The nextToken value returned from a ListTasks request indicating that more
	// results are available to fulfill the request and further calls will be needed.
	// If maxResults was provided, it is possible the number of results to be fewer
	// than maxResults. This token should be treated as an opaque identifier that is
	// only used to retrieve the next items in a list and not for other programmatic
	// purposes.
	NextToken *string

	// The name of the service with which to filter the ListTasks results. Specifying a
	// serviceName limits the results to tasks that belong to that service.
	ServiceName *string

	// The startedBy value with which to filter the task results. Specifying a
	// startedBy value limits the results to tasks that were started with that value.
	StartedBy *string
}

type ListTasksOutput struct {

	// The nextToken value to include in a future ListTasks request. When the results
	// of a ListTasks request exceed maxResults, this value can be used to retrieve the
	// next page of results. This value is null when there are no more results to
	// return.
	NextToken *string

	// The list of task ARN entries for the ListTasks request.
	TaskArns []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTasks{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTasks(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListTasks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "ListTasks",
	}
}
