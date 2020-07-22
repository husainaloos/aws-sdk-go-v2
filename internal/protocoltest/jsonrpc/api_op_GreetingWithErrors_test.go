// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc

import (
	"bytes"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/jsonrpc/types"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithytesting "github.com/awslabs/smithy-go/testing"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestClient_GreetingWithErrors_InvalidGreeting_awsAwsjson11Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectError   *types.InvalidGreeting
	}{
		// Parses simple JSON errors
		"AwsJson11InvalidGreetingError": {
			StatusCode: 400,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "__type": "InvalidGreeting",
			    "Message": "Hi"
			}`),
			ExpectError: &types.InvalidGreeting{
				Message: ptr.String("Hi"),
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for k, vs := range c.Header {
					for _, v := range vs {
						w.Header().Add(k, v)
					}
				}
				if len(c.BodyMediaType) != 0 && len(w.Header().Values("Content-Type")) == 0 {
					w.Header().Set("Content-Type", c.BodyMediaType)
				}
				if len(c.Body) != 0 {
					w.Header().Set("Content-Length", strconv.Itoa(len(c.Body)))
				}
				w.WriteHeader(c.StatusCode)
				if len(c.Body) != 0 {
					if _, err := io.Copy(w, bytes.NewReader(c.Body)); err != nil {
						t.Errorf("failed to write response body, %v", err)
					}
				}
			}))
			defer server.Close()
			client := New(Options{
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (e aws.Endpoint, err error) {
					e.URL = server.URL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient: aws.NewBuildableHTTPClient(),
				Region:     "us-west-2",
			})
			var params GreetingWithErrorsInput
			result, err := client.GreetingWithErrors(context.Background(), &params)
			if err == nil {
				t.Fatalf("expect not nil err")
			}
			if result != nil {
				t.Fatalf("expect nil result, got %v", result)
			}
			var opErr interface {
				Service() string
				Operation() string
			}
			if !errors.As(err, &opErr) {
				t.Fatalf("expect *types.InvalidGreeting operation error, got %T", err)
			}
			if e, a := client.ServiceID(), opErr.Service(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			if e, a := "GreetingWithErrors", opErr.Operation(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			var actualErr *types.InvalidGreeting
			if !errors.As(err, &actualErr) {
				t.Fatalf("expect *types.InvalidGreeting result error, got %T", err)
			}
			if err := smithytesting.CompareValues(c.ExpectError, actualErr, cmpopts.IgnoreUnexported(middleware.Metadata{})); err != nil {
				t.Errorf("expect c.ExpectError value match:\n%v", err)
			}
		})
	}
}

func TestClient_GreetingWithErrors_ComplexError_awsAwsjson11Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectError   *types.ComplexError
	}{
		// Parses a complex error with no message member
		"AwsJson11ComplexError": {
			StatusCode: 400,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "__type": "ComplexError",
			    "TopLevel": "Top level",
			    "Nested": {
			        "Fooooo": "bar"
			    }
			}`),
			ExpectError: &types.ComplexError{
				TopLevel: ptr.String("Top level"),
				Nested: &types.ComplexNestedErrorData{
					Foo: ptr.String("bar"),
				},
			},
		},
		"AwsJson11EmptyComplexError": {
			StatusCode: 400,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "__type": "ComplexError"
			}`),
			ExpectError: &types.ComplexError{},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for k, vs := range c.Header {
					for _, v := range vs {
						w.Header().Add(k, v)
					}
				}
				if len(c.BodyMediaType) != 0 && len(w.Header().Values("Content-Type")) == 0 {
					w.Header().Set("Content-Type", c.BodyMediaType)
				}
				if len(c.Body) != 0 {
					w.Header().Set("Content-Length", strconv.Itoa(len(c.Body)))
				}
				w.WriteHeader(c.StatusCode)
				if len(c.Body) != 0 {
					if _, err := io.Copy(w, bytes.NewReader(c.Body)); err != nil {
						t.Errorf("failed to write response body, %v", err)
					}
				}
			}))
			defer server.Close()
			client := New(Options{
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (e aws.Endpoint, err error) {
					e.URL = server.URL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient: aws.NewBuildableHTTPClient(),
				Region:     "us-west-2",
			})
			var params GreetingWithErrorsInput
			result, err := client.GreetingWithErrors(context.Background(), &params)
			if err == nil {
				t.Fatalf("expect not nil err")
			}
			if result != nil {
				t.Fatalf("expect nil result, got %v", result)
			}
			var opErr interface {
				Service() string
				Operation() string
			}
			if !errors.As(err, &opErr) {
				t.Fatalf("expect *types.ComplexError operation error, got %T", err)
			}
			if e, a := client.ServiceID(), opErr.Service(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			if e, a := "GreetingWithErrors", opErr.Operation(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			var actualErr *types.ComplexError
			if !errors.As(err, &actualErr) {
				t.Fatalf("expect *types.ComplexError result error, got %T", err)
			}
			if err := smithytesting.CompareValues(c.ExpectError, actualErr, cmpopts.IgnoreUnexported(middleware.Metadata{})); err != nil {
				t.Errorf("expect c.ExpectError value match:\n%v", err)
			}
		})
	}
}

func TestClient_GreetingWithErrors_FooError_awsAwsjson11Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectError   *types.FooError
	}{
		// Serializes the X-Amzn-ErrorType header. For an example service, see Amazon EKS.
		"AwsJson11FooErrorUsingXAmznErrorType": {
			StatusCode: 500,
			Header: http.Header{
				"X-Amzn-Errortype": []string{"FooError"},
			},
			ExpectError: &types.FooError{},
		},
		// Some X-Amzn-Errortype headers contain URLs. Clients need to split the URL on ':'
		// and take only the first half of the string. For example,
		// 'ValidationException:http://internal.amazon.com/coral/com.amazon.coral.validate/'
		// is to be interpreted as 'ValidationException'. For an example service see Amazon
		// Polly.
		"AwsJson11FooErrorUsingXAmznErrorTypeWithUri": {
			StatusCode: 500,
			Header: http.Header{
				"X-Amzn-Errortype": []string{"FooError:http://internal.amazon.com/coral/com.amazon.coral.validate/"},
			},
			ExpectError: &types.FooError{},
		},
		// X-Amzn-Errortype might contain a URL and a namespace. Client should extract only
		// the shape name. This is a pathalogical case that might not actually happen in
		// any deployed AWS service.
		"AwsJson11FooErrorUsingXAmznErrorTypeWithUriAndNamespace": {
			StatusCode: 500,
			Header: http.Header{
				"X-Amzn-Errortype": []string{"aws.protocoltests.restjson#FooError:http://internal.amazon.com/coral/com.amazon.coral.validate/"},
			},
			ExpectError: &types.FooError{},
		},
		// This example uses the 'code' property in the output rather than
		// X-Amzn-Errortype. Some services do this though it's preferable to send the
		// X-Amzn-Errortype. Client implementations must first check for the
		// X-Amzn-Errortype and then check for a top-level 'code' property. For example
		// service see Amazon S3 Glacier.
		"AwsJson11FooErrorUsingCode": {
			StatusCode: 500,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "code": "FooError"
			}`),
			ExpectError: &types.FooError{},
		},
		// Some services serialize errors using code, and it might contain a namespace.
		// Clients should just take the last part of the string after '#'.
		"AwsJson11FooErrorUsingCodeAndNamespace": {
			StatusCode: 500,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "code": "aws.protocoltests.restjson#FooError"
			}`),
			ExpectError: &types.FooError{},
		},
		// Some services serialize errors using code, and it might contain a namespace. It
		// also might contain a URI. Clients should just take the last part of the string
		// after '#' and before ":". This is a pathalogical case that might not occur in
		// any deployed AWS service.
		"AwsJson11FooErrorUsingCodeUriAndNamespace": {
			StatusCode: 500,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "code": "aws.protocoltests.restjson#FooError:http://internal.amazon.com/coral/com.amazon.coral.validate/"
			}`),
			ExpectError: &types.FooError{},
		},
		// Some services serialize errors using __type.
		"AwsJson11FooErrorWithDunderType": {
			StatusCode: 500,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "__type": "FooError"
			}`),
			ExpectError: &types.FooError{},
		},
		// Some services serialize errors using __type, and it might contain a namespace.
		// Clients should just take the last part of the string after '#'.
		"AwsJson11FooErrorWithDunderTypeAndNamespace": {
			StatusCode: 500,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "__type": "aws.protocoltests.restjson#FooError"
			}`),
			ExpectError: &types.FooError{},
		},
		// Some services serialize errors using __type, and it might contain a namespace.
		// It also might contain a URI. Clients should just take the last part of the
		// string after '#' and before ":". This is a pathalogical case that might not
		// occur in any deployed AWS service.
		"AwsJson11FooErrorWithDunderTypeUriAndNamespace": {
			StatusCode: 500,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "__type": "aws.protocoltests.restjson#FooError:http://internal.amazon.com/coral/com.amazon.coral.validate/"
			}`),
			ExpectError: &types.FooError{},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for k, vs := range c.Header {
					for _, v := range vs {
						w.Header().Add(k, v)
					}
				}
				if len(c.BodyMediaType) != 0 && len(w.Header().Values("Content-Type")) == 0 {
					w.Header().Set("Content-Type", c.BodyMediaType)
				}
				if len(c.Body) != 0 {
					w.Header().Set("Content-Length", strconv.Itoa(len(c.Body)))
				}
				w.WriteHeader(c.StatusCode)
				if len(c.Body) != 0 {
					if _, err := io.Copy(w, bytes.NewReader(c.Body)); err != nil {
						t.Errorf("failed to write response body, %v", err)
					}
				}
			}))
			defer server.Close()
			client := New(Options{
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (e aws.Endpoint, err error) {
					e.URL = server.URL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient: aws.NewBuildableHTTPClient(),
				Region:     "us-west-2",
			})
			var params GreetingWithErrorsInput
			result, err := client.GreetingWithErrors(context.Background(), &params)
			if err == nil {
				t.Fatalf("expect not nil err")
			}
			if result != nil {
				t.Fatalf("expect nil result, got %v", result)
			}
			var opErr interface {
				Service() string
				Operation() string
			}
			if !errors.As(err, &opErr) {
				t.Fatalf("expect *types.FooError operation error, got %T", err)
			}
			if e, a := client.ServiceID(), opErr.Service(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			if e, a := "GreetingWithErrors", opErr.Operation(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			var actualErr *types.FooError
			if !errors.As(err, &actualErr) {
				t.Fatalf("expect *types.FooError result error, got %T", err)
			}
			if err := smithytesting.CompareValues(c.ExpectError, actualErr, cmpopts.IgnoreUnexported(middleware.Metadata{})); err != nil {
				t.Errorf("expect c.ExpectError value match:\n%v", err)
			}
		})
	}
}
