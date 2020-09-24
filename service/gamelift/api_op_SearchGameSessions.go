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

// Retrieves all active game sessions that match a set of search criteria and sorts
// them in a specified order. You can search or sort by the following game session
// attributes:
//
//     * gameSessionId -- A unique identifier for the game session.
// You can use either a GameSessionId or GameSessionArn value.
//
//     *
// gameSessionName -- Name assigned to a game session. This value is set when
// requesting a new game session with CreateGameSession () or updating with
// UpdateGameSession (). Game session names do not need to be unique to a game
// session.
//
//     * gameSessionProperties -- Custom data defined in a game session's
// GameProperty parameter. GameProperty values are stored as key:value pairs; the
// filter expression must indicate the key and a string to search the data values
// for. For example, to search for game sessions with custom data containing the
// key:value pair "gameMode:brawl", specify the following:
// gameSessionProperties.gameMode = "brawl". All custom data values are searched as
// strings.
//
//     * maximumSessions -- Maximum number of player sessions allowed for
// a game session. This value is set when requesting a new game session with
// CreateGameSession () or updating with UpdateGameSession ().
//
//     *
// creationTimeMillis -- Value indicating when a game session was created. It is
// expressed in Unix time as milliseconds.
//
//     * playerSessionCount -- Number of
// players currently connected to a game session. This value changes rapidly as
// players join the session or drop out.
//
//     * hasAvailablePlayerSessions --
// Boolean value indicating whether a game session has reached its maximum number
// of players. It is highly recommended that all search requests include this
// filter attribute to optimize search performance and return only sessions that
// players can join.
//
// Returned values for playerSessionCount and
// hasAvailablePlayerSessions change quickly as players join sessions and others
// drop out. Results should be considered a snapshot in time. Be sure to refresh
// search results often, and handle sessions that fill up before a player can join.
// To search or sort, specify either a fleet ID or an alias ID, and provide a
// search filter expression, a sort expression, or both. If successful, a
// collection of GameSession () objects matching the request is returned. Use the
// pagination parameters to retrieve results as a set of sequential pages. You can
// search for game sessions one fleet at a time only. To find game sessions across
// multiple fleets, you must search each fleet separately and combine the results.
// This search feature finds only game sessions that are in ACTIVE status. To
// locate games in statuses other than active, use DescribeGameSessionDetails ().
//
//
// * CreateGameSession ()
//
//     * DescribeGameSessions ()
//
//     *
// DescribeGameSessionDetails ()
//
//     * SearchGameSessions ()
//
//     *
// UpdateGameSession ()
//
//     * GetGameSessionLogUrl ()
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
func (c *Client) SearchGameSessions(ctx context.Context, params *SearchGameSessionsInput, optFns ...func(*Options)) (*SearchGameSessionsOutput, error) {
	stack := middleware.NewStack("SearchGameSessions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSearchGameSessionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opSearchGameSessions(options.Region), middleware.Before)
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
			OperationName: "SearchGameSessions",
			Err:           err,
		}
	}
	out := result.(*SearchGameSessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type SearchGameSessionsInput struct {
	// String containing the search criteria for the session search. If no filter
	// expression is included, the request returns results for all game sessions in the
	// fleet that are in ACTIVE status. A filter expression can contain one or multiple
	// conditions. Each condition consists of the following:
	//
	//     * Operand -- Name of
	// a game session attribute. Valid values are gameSessionName, gameSessionId,
	// gameSessionProperties, maximumSessions, creationTimeMillis, playerSessionCount,
	// hasAvailablePlayerSessions.
	//
	//     * Comparator -- Valid comparators are: =, <>,
	// <, >, <=, >=.
	//
	//     * Value -- Value to be searched for. Values may be numbers,
	// boolean values (true/false) or strings depending on the operand. String values
	// are case sensitive and must be enclosed in single quotes. Special characters
	// must be escaped. Boolean and string values can only be used with the comparators
	// = and <>. For example, the following filter expression searches on
	// gameSessionName: "FilterExpression": "gameSessionName = 'Matt\\'s Awesome Game
	// 1'".
	//
	// To chain multiple conditions in a single expression, use the logical
	// keywords AND, OR, and NOT and parentheses as needed. For example: x AND y AND
	// NOT z, NOT (x OR y). Session search evaluates conditions from left to right
	// using the following precedence rules:
	//
	//     * =, <>, <, >, <=, >=
	//
	//     *
	// Parentheses
	//
	//     * NOT
	//
	//     * AND
	//
	//     * OR
	//
	// For example, this filter expression
	// retrieves game sessions hosting at least ten players that have an open player
	// slot: "maximumSessions>=10 AND hasAvailablePlayerSessions=true".
	FilterExpression *string
	// Instructions on how to sort the search results. If no sort expression is
	// included, the request returns results in random order. A sort expression
	// consists of the following elements:
	//
	//     * Operand -- Name of a game session
	// attribute. Valid values are gameSessionName, gameSessionId,
	// gameSessionProperties, maximumSessions, creationTimeMillis, playerSessionCount,
	// hasAvailablePlayerSessions.
	//
	//     * Order -- Valid sort orders are ASC
	// (ascending) and DESC (descending).
	//
	// For example, this sort expression returns
	// the oldest active sessions first: "SortExpression": "creationTimeMillis ASC".
	// Results with a null value for the sort operand are returned at the end of the
	// list.
	SortExpression *string
	// Token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this action. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. The maximum number of results returned
	// is 20, even if this value is not set or is set higher than 20.
	Limit *int32
	// A unique identifier for a fleet to search for active game sessions. You can use
	// either the fleet ID or ARN value. Each request must reference either a fleet ID
	// or alias ID, but not both.
	FleetId *string
	// A unique identifier for an alias associated with the fleet to search for active
	// game sessions. You can use either the alias ID or ARN value. Each request must
	// reference either a fleet ID or alias ID, but not both.
	AliasId *string
}

// Represents the returned data in response to a request action.
type SearchGameSessionsOutput struct {
	// Token that indicates where to resume retrieving results on the next call to this
	// action. If no token is returned, these results represent the end of the list.
	NextToken *string
	// A collection of objects containing game session properties for each session
	// matching the request.
	GameSessions []*types.GameSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSearchGameSessionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSearchGameSessions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchGameSessions{}, middleware.After)
}

func newServiceMetadataMiddleware_opSearchGameSessions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "SearchGameSessions",
	}
}
