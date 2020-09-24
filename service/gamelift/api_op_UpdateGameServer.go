// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This action is part of Amazon GameLift FleetIQ with game server groups, which is
// in preview release and is subject to change. Updates information about a
// registered game server. This action is called by a game server process that is
// running on an instance in a game server group. There are three reasons to update
// game server information: (1) to change the utilization status of the game
// server, (2) to report game server health status, and (3) to change game server
// metadata. A registered game server should regularly report health and should
// update utilization status when it is supporting gameplay so that GameLift
// FleetIQ can accurately track game server availability. You can make all three
// types of updates in the same request.
//
//     * To update the game server's
// utilization status, identify the game server and game server group and specify
// the current utilization status. Use this status to identify when game servers
// are currently hosting games and when they are available to be claimed.
//
//     * To
// report health status, identify the game server and game server group and set
// health check to HEALTHY. If a game server does not report health status for a
// certain length of time, the game server is no longer considered healthy and will
// be eventually de-registered from the game server group to avoid affecting
// utilization metrics. The best practice is to report health every 60 seconds.
//
//
// * To change game server metadata, provide updated game server data and custom
// sort key values.
//
// Once a game server is successfully updated, the relevant
// statuses and timestamps are updated. Learn more GameLift FleetIQ Guide
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gsg-intro.html)
// Related operations
//
//     * RegisterGameServer ()
//
//     * ListGameServers ()
//
//     *
// ClaimGameServer ()
//
//     * DescribeGameServer ()
//
//     * UpdateGameServer ()
//
//
// * DeregisterGameServer ()
func (c *Client) UpdateGameServer(ctx context.Context, params *UpdateGameServerInput, optFns ...func(*Options)) (*UpdateGameServerOutput, error) {
	stack := middleware.NewStack("UpdateGameServer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateGameServerMiddlewares(stack)
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
	addOpUpdateGameServerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGameServer(options.Region), middleware.Before)
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
			OperationName: "UpdateGameServer",
			Err:           err,
		}
	}
	out := result.(*UpdateGameServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGameServerInput struct {
	// Indicates whether the game server is available or is currently hosting gameplay.
	UtilizationStatus types.GameServerUtilizationStatus
	// Indicates health status of the game server. An update that explicitly includes
	// this parameter updates the game server's LastHealthCheckTime time stamp.
	HealthCheck types.GameServerHealthCheck
	// The identifier for the game server to be updated.
	GameServerId *string
	// A set of custom game server properties, formatted as a single string value. This
	// data is passed to a game client or service when it requests information on a
	// game servers using DescribeGameServer () or ClaimGameServer ().
	GameServerData *string
	// An identifier for the game server group where the game server is running. Use
	// either the GameServerGroup () name or ARN value.
	GameServerGroupName *string
	// A game server tag that can be used to request sorted lists of game servers using
	// ListGameServers (). Custom sort keys are developer-defined based on how you want
	// to organize the retrieved game server information.
	CustomSortKey *string
}

type UpdateGameServerOutput struct {
	// Object that describes the newly updated game server resource.
	GameServer *types.GameServer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateGameServerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateGameServer{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateGameServer{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateGameServer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "UpdateGameServer",
	}
}
