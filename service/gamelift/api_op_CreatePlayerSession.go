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

// Reserves an open player slot in an active game session. Before a player can be
// added, a game session must have an ACTIVE status, have a creation policy of
// ALLOW_ALL, and have an open player slot. To add a group of players to a game
// session, use CreatePlayerSessions (). When the player connects to the game
// server and references a player session ID, the game server contacts the Amazon
// GameLift service to validate the player reservation and accept the player. To
// create a player session, specify a game session ID, player ID, and optionally a
// string of player data. If successful, a slot is reserved in the game session for
// the player and a new PlayerSession () object is returned. Player sessions cannot
// be updated. Available in Amazon GameLift Local.
//
//     * CreatePlayerSession ()
//
//
// * CreatePlayerSessions ()
//
//     * DescribePlayerSessions ()
//
//     * Game session
// placements
//
//         * StartGameSessionPlacement ()
//
//         *
// DescribeGameSessionPlacement ()
//
//         * StopGameSessionPlacement ()
func (c *Client) CreatePlayerSession(ctx context.Context, params *CreatePlayerSessionInput, optFns ...func(*Options)) (*CreatePlayerSessionOutput, error) {
	stack := middleware.NewStack("CreatePlayerSession", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreatePlayerSessionMiddlewares(stack)
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
	addOpCreatePlayerSessionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePlayerSession(options.Region), middleware.Before)
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
			OperationName: "CreatePlayerSession",
			Err:           err,
		}
	}
	out := result.(*CreatePlayerSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type CreatePlayerSessionInput struct {
	// Developer-defined information related to a player. Amazon GameLift does not use
	// this data, so it can be formatted as needed for use in the game.
	PlayerData *string
	// A unique identifier for the game session to add a player to.
	GameSessionId *string
	// A unique identifier for a player. Player IDs are developer-defined.
	PlayerId *string
}

// Represents the returned data in response to a request action.
type CreatePlayerSessionOutput struct {
	// Object that describes the newly created player session record.
	PlayerSession *types.PlayerSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreatePlayerSessionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePlayerSession{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePlayerSession{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreatePlayerSession(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "CreatePlayerSession",
	}
}
