// Code generated by smithy-go-codegen DO NOT EDIT.

package rdsdata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rdsdata/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Runs one or more SQL statements. This operation is deprecated. Use the
// BatchExecuteStatement or ExecuteStatement operation.
func (c *Client) ExecuteSql(ctx context.Context, params *ExecuteSqlInput, optFns ...func(*Options)) (*ExecuteSqlOutput, error) {
	stack := middleware.NewStack("ExecuteSql", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpExecuteSqlMiddlewares(stack)
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
	addOpExecuteSqlValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opExecuteSql(options.Region), middleware.Before)
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
			OperationName: "ExecuteSql",
			Err:           err,
		}
	}
	out := result.(*ExecuteSqlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request parameters represent the input of a request to run one or more SQL
// statements.
type ExecuteSqlInput struct {
	// The Amazon Resource Name (ARN) of the secret that enables access to the DB
	// cluster.
	AwsSecretStoreArn *string
	// The name of the database.
	Database *string
	// The ARN of the Aurora Serverless DB cluster.
	DbClusterOrInstanceArn *string
	// The name of the database schema.
	Schema *string
	// One or more SQL statements to run on the DB cluster. You can separate SQL
	// statements from each other with a semicolon (;). Any valid SQL statement is
	// permitted, including data definition, data manipulation, and commit statements.
	SqlStatements *string
}

// The response elements represent the output of a request to run one or more SQL
// statements.
type ExecuteSqlOutput struct {
	// The results of the SQL statement or statements.
	SqlStatementResults []*types.SqlStatementResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpExecuteSqlMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpExecuteSql{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpExecuteSql{}, middleware.After)
}

func newServiceMetadataMiddleware_opExecuteSql(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "",
		OperationName: "ExecuteSql",
	}
}
