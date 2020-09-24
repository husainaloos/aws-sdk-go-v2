// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a DataSource object. A DataSource references data that can be used to
// perform CreateMLModel, CreateEvaluation, or CreateBatchPrediction operations.
// <p> <code>CreateDataSourceFromS3</code> is an asynchronous operation. In
// response to <code>CreateDataSourceFromS3</code>, Amazon Machine Learning (Amazon
// ML) immediately returns and sets the <code>DataSource</code> status to
// <code>PENDING</code>. After the <code>DataSource</code> has been created and is
// ready for use, Amazon ML sets the <code>Status</code> parameter to
// <code>COMPLETED</code>. <code>DataSource</code> in the <code>COMPLETED</code> or
// <code>PENDING</code> state can be used to perform only
// <code>CreateMLModel</code>, <code>CreateEvaluation</code> or
// <code>CreateBatchPrediction</code> operations. </p> <p> If Amazon ML can't
// accept the input source, it sets the <code>Status</code> parameter to
// <code>FAILED</code> and includes an error message in the <code>Message</code>
// attribute of the <code>GetDataSource</code> operation response. </p> <p>The
// observation data used in a <code>DataSource</code> should be ready to use; that
// is, it should have a consistent structure, and missing data values should be
// kept to a minimum. The observation data must reside in one or more .csv files in
// an Amazon Simple Storage Service (Amazon S3) location, along with a schema that
// describes the data items by name and type. The same schema must be used for all
// of the data files referenced by the <code>DataSource</code>. </p> <p>After the
// <code>DataSource</code> has been created, it's ready to use in evaluations and
// batch predictions. If you plan to use the <code>DataSource</code> to train an
// <code>MLModel</code>, the <code>DataSource</code> also needs a recipe. A recipe
// describes how each input variable will be used in training an
// <code>MLModel</code>. Will the variable be included or excluded from training?
// Will the variable be manipulated; for example, will it be combined with another
// variable or will it be split apart into word combinations? The recipe provides
// answers to these questions.</p>
func (c *Client) CreateDataSourceFromS3(ctx context.Context, params *CreateDataSourceFromS3Input, optFns ...func(*Options)) (*CreateDataSourceFromS3Output, error) {
	stack := middleware.NewStack("CreateDataSourceFromS3", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateDataSourceFromS3Middlewares(stack)
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
	addOpCreateDataSourceFromS3ValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataSourceFromS3(options.Region), middleware.Before)
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
			OperationName: "CreateDataSourceFromS3",
			Err:           err,
		}
	}
	out := result.(*CreateDataSourceFromS3Output)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataSourceFromS3Input struct {
	// A user-supplied identifier that uniquely identifies the DataSource.
	DataSourceId *string
	// A user-supplied name or description of the DataSource.
	DataSourceName *string
	// The compute statistics for a DataSource. The statistics are generated from the
	// observation data referenced by a DataSource. Amazon ML uses the statistics
	// internally during MLModel training. This parameter must be set to true if the
	// DataSource needs to be used for MLModel training.
	ComputeStatistics *bool
	// The data specification of a DataSource:
	//
	//     * DataLocationS3 - The Amazon S3
	// location of the observation data.
	//
	//     * DataSchemaLocationS3 - The Amazon S3
	// location of the DataSchema.
	//
	//     * DataSchema - A JSON string representing the
	// schema. This is not required if DataSchemaUri is specified.
	//
	//     *
	// DataRearrangement - A JSON string that represents the splitting and
	// rearrangement requirements for the Datasource. Sample -
	// "{\"splitting\":{\"percentBegin\":10,\"percentEnd\":60}}"
	DataSpec *types.S3DataSpec
}

// Represents the output of a CreateDataSourceFromS3 operation, and is an
// acknowledgement that Amazon ML received the request. The CreateDataSourceFromS3
// operation is asynchronous. You can poll for updates by using the
// GetBatchPrediction operation and checking the Status parameter.
type CreateDataSourceFromS3Output struct {
	// A user-supplied ID that uniquely identifies the DataSource. This value should be
	// identical to the value of the DataSourceID in the request.
	DataSourceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateDataSourceFromS3Middlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDataSourceFromS3{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDataSourceFromS3{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDataSourceFromS3(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "CreateDataSourceFromS3",
	}
}
