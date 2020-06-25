// Code generated by smithy-go-codegen DO NOT EDIT.
package query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/query/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This test case serializes XML lists for the following cases for both input and
// output:
//
//     * Normal XML lists.
//
//     * Normal XML sets.
//
//     * XML lists of
// lists.
//
//     * XML lists with @xmlName on its members
//
//     * Flattened XML
// lists.
//
//     * Flattened XML lists with @xmlName.
//
//     * Lists of structures.
func (c *Client) XmlLists(ctx context.Context, params *XmlListsInput, optFns ...func(*Options)) (*XmlListsOutput, error) {
	stack := middleware.NewStack("XmlLists", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	stack.Initialize.Add(awsmiddleware.RegisterServiceMetadata{
		Region:         options.Region,
		ServiceName:    "Query Protocol",
		ServiceID:      "queryprotocol",
		EndpointPrefix: "queryprotocol",
		OperationName:  "XmlLists",
	}, middleware.Before)
	stack.Build.Add(awsmiddleware.RequestInvocationIDMiddleware{}, middleware.After)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	stack.Deserialize.Add(awsmiddleware.AttemptClockSkewMiddleware{}, middleware.After)
	stack.Finalize.Add(retry.NewAttemptMiddleware(options.Retryer, smithyhttp.RequestCloner), middleware.After)
	stack.Finalize.Add(retry.MetricsHeaderMiddleware{}, middleware.After)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "XmlLists",
			Err:           err,
		}
	}
	out := result.(*XmlListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type XmlListsInput struct {
}

type XmlListsOutput struct {
	StringList    []*string
	StringSet     []*string
	IntegerList   []*int32
	BooleanList   []*bool
	TimestampList []*time.Time
	EnumList      []types.FooEnum
	// A list of lists of strings.
	NestedStringList   [][]*string
	RenamedListMembers []*string
	FlattenedList      []*string
	FlattenedList2     []*string
	StructureList      []*types.StructureListMember

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}
