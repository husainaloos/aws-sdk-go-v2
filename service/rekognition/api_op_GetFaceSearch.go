// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the face search results for Amazon Rekognition Video face search started by
// StartFaceSearch (). The search returns faces in a collection that match the
// faces of persons detected in a video. It also includes the time(s) that faces
// are matched in the video. Face search in a video is an asynchronous operation.
// You start face search by calling to StartFaceSearch () which returns a job
// identifier (JobId). When the search operation finishes, Amazon Rekognition Video
// publishes a completion status to the Amazon Simple Notification Service topic
// registered in the initial call to StartFaceSearch. To get the search results,
// first check that the status value published to the Amazon SNS topic is
// SUCCEEDED. If so, call GetFaceSearch and pass the job identifier (JobId) from
// the initial call to StartFaceSearch.  <p>For more information, see Searching
// Faces in a Collection in the Amazon Rekognition Developer Guide.</p> <p>The
// search results are retured in an array, <code>Persons</code>, of
// <a>PersonMatch</a> objects. Each<code>PersonMatch</code> element contains
// details about the matching faces in the input collection, person information
// (facial attributes, bounding boxes, and person identifer) for the matched
// person, and the time the person was matched in the video.</p> <note> <p>
// <code>GetFaceSearch</code> only returns the default facial attributes
// (<code>BoundingBox</code>, <code>Confidence</code>, <code>Landmarks</code>,
// <code>Pose</code>, and <code>Quality</code>). The other facial attributes listed
// in the <code>Face</code> object of the following response syntax are not
// returned. For more information, see FaceDetail in the Amazon Rekognition
// Developer Guide. </p> </note> <p>By default, the <code>Persons</code> array is
// sorted by the time, in milliseconds from the start of the video, persons are
// matched. You can also sort by persons by specifying <code>INDEX</code> for the
// <code>SORTBY</code> input parameter.</p>
func (c *Client) GetFaceSearch(ctx context.Context, params *GetFaceSearchInput, optFns ...func(*Options)) (*GetFaceSearchOutput, error) {
	stack := middleware.NewStack("GetFaceSearch", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetFaceSearchMiddlewares(stack)
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
	addOpGetFaceSearchValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetFaceSearch(options.Region), middleware.Before)
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
			OperationName: "GetFaceSearch",
			Err:           err,
		}
	}
	out := result.(*GetFaceSearchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFaceSearchInput struct {
	// Sort to use for grouping faces in the response. Use TIMESTAMP to group faces by
	// the time that they are recognized. Use INDEX to sort by recognized faces.
	SortBy types.FaceSearchSortBy
	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000. If you specify a value greater than 1000, a maximum of 1000
	// results is returned. The default value is 1000.
	MaxResults *int32
	// If the previous response was incomplete (because there is more search results to
	// retrieve), Amazon Rekognition Video returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of search results.
	NextToken *string
	// The job identifer for the search request. You get the job identifier from an
	// initial call to StartFaceSearch.
	JobId *string
}

type GetFaceSearchOutput struct {
	// If the job fails, StatusMessage provides a descriptive error message.
	StatusMessage *string
	// Information about a video that Amazon Rekognition analyzed. Videometadata is
	// returned in every page of paginated responses from a Amazon Rekognition Video
	// operation.
	VideoMetadata *types.VideoMetadata
	// The current status of the face search job.
	JobStatus types.VideoJobStatus
	// If the response is truncated, Amazon Rekognition Video returns this token that
	// you can use in the subsequent request to retrieve the next set of search
	// results.
	NextToken *string
	// An array of persons, PersonMatch (), in the video whose face(s) match the
	// face(s) in an Amazon Rekognition collection. It also includes time information
	// for when persons are matched in the video. You specify the input collection in
	// an initial call to StartFaceSearch. Each Persons element includes a time the
	// person was matched, face match details (FaceMatches) for matching faces in the
	// collection, and person information (Person) for the matched person.
	Persons []*types.PersonMatch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetFaceSearchMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetFaceSearch{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetFaceSearch{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetFaceSearch(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "GetFaceSearch",
	}
}
