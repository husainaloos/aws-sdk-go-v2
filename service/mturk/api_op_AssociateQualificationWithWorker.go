// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The AssociateQualificationWithWorker operation gives a Worker a Qualification.
// AssociateQualificationWithWorker does not require that the Worker submit a
// Qualification request. It gives the Qualification directly to the Worker.  <p>
// You can only assign a Qualification of a Qualification type that you created
// (using the <code>CreateQualificationType</code> operation). </p> <note> <p>
// Note: <code>AssociateQualificationWithWorker</code> does not affect any pending
// Qualification requests for the Qualification by the Worker. If you assign a
// Qualification to a Worker, then later grant a Qualification request made by the
// Worker, the granting of the request may modify the Qualification score. To
// resolve a pending Qualification request without affecting the Qualification the
// Worker already has, reject the request with the
// <code>RejectQualificationRequest</code> operation. </p> </note>
func (c *Client) AssociateQualificationWithWorker(ctx context.Context, params *AssociateQualificationWithWorkerInput, optFns ...func(*Options)) (*AssociateQualificationWithWorkerOutput, error) {
	stack := middleware.NewStack("AssociateQualificationWithWorker", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssociateQualificationWithWorkerMiddlewares(stack)
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
	addOpAssociateQualificationWithWorkerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateQualificationWithWorker(options.Region), middleware.Before)
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
			OperationName: "AssociateQualificationWithWorker",
			Err:           err,
		}
	}
	out := result.(*AssociateQualificationWithWorkerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateQualificationWithWorkerInput struct {
	// The value of the Qualification to assign.
	IntegerValue *int32
	// The ID of the Qualification type to use for the assigned Qualification.
	QualificationTypeId *string
	// The ID of the Worker to whom the Qualification is being assigned. Worker IDs are
	// included with submitted HIT assignments and Qualification requests.
	WorkerId *string
	// Specifies whether to send a notification email message to the Worker saying that
	// the qualification was assigned to the Worker. Note: this is true by default.
	SendNotification *bool
}

type AssociateQualificationWithWorkerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssociateQualificationWithWorkerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateQualificationWithWorker{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateQualificationWithWorker{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateQualificationWithWorker(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "AssociateQualificationWithWorker",
	}
}
