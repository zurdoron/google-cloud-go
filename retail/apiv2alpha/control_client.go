// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package retail

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	retailpb "google.golang.org/genproto/googleapis/cloud/retail/v2alpha"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newControlClientHook clientHook

// ControlCallOptions contains the retry settings for each method of ControlClient.
type ControlCallOptions struct {
	CreateControl  []gax.CallOption
	DeleteControl  []gax.CallOption
	UpdateControl  []gax.CallOption
	GetControl     []gax.CallOption
	ListControls   []gax.CallOption
	GetOperation   []gax.CallOption
	ListOperations []gax.CallOption
}

func defaultControlGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("retail.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("retail.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://retail.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultControlCallOptions() *ControlCallOptions {
	return &ControlCallOptions{
		CreateControl: []gax.CallOption{},
		DeleteControl: []gax.CallOption{},
		UpdateControl: []gax.CallOption{},
		GetControl:    []gax.CallOption{},
		ListControls:  []gax.CallOption{},
		GetOperation:  []gax.CallOption{},
		ListOperations: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        300000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalControlClient is an interface that defines the methods available from Retail API.
type internalControlClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateControl(context.Context, *retailpb.CreateControlRequest, ...gax.CallOption) (*retailpb.Control, error)
	DeleteControl(context.Context, *retailpb.DeleteControlRequest, ...gax.CallOption) error
	UpdateControl(context.Context, *retailpb.UpdateControlRequest, ...gax.CallOption) (*retailpb.Control, error)
	GetControl(context.Context, *retailpb.GetControlRequest, ...gax.CallOption) (*retailpb.Control, error)
	ListControls(context.Context, *retailpb.ListControlsRequest, ...gax.CallOption) *ControlIterator
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
}

// ControlClient is a client for interacting with Retail API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for modifying Control.
type ControlClient struct {
	// The internal transport-dependent client.
	internalClient internalControlClient

	// The call options for this service.
	CallOptions *ControlCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ControlClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ControlClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ControlClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateControl creates a Control.
//
// If the Control to create already
// exists, an ALREADY_EXISTS error is returned.
func (c *ControlClient) CreateControl(ctx context.Context, req *retailpb.CreateControlRequest, opts ...gax.CallOption) (*retailpb.Control, error) {
	return c.internalClient.CreateControl(ctx, req, opts...)
}

// DeleteControl deletes a Control.
//
// If the Control to delete does not
// exist, a NOT_FOUND error is returned.
func (c *ControlClient) DeleteControl(ctx context.Context, req *retailpb.DeleteControlRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteControl(ctx, req, opts...)
}

// UpdateControl updates a Control.
//
// Control cannot be set to a different
// oneof field, if so an INVALID_ARGUMENT is returned. If the
// Control to delete does not exist, a
// NOT_FOUND error is returned.
func (c *ControlClient) UpdateControl(ctx context.Context, req *retailpb.UpdateControlRequest, opts ...gax.CallOption) (*retailpb.Control, error) {
	return c.internalClient.UpdateControl(ctx, req, opts...)
}

// GetControl gets a Control.
func (c *ControlClient) GetControl(ctx context.Context, req *retailpb.GetControlRequest, opts ...gax.CallOption) (*retailpb.Control, error) {
	return c.internalClient.GetControl(ctx, req, opts...)
}

// ListControls lists all Controls linked to this catalog.
func (c *ControlClient) ListControls(ctx context.Context, req *retailpb.ListControlsRequest, opts ...gax.CallOption) *ControlIterator {
	return c.internalClient.ListControls(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *ControlClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *ControlClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// controlGRPCClient is a client for interacting with Retail API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type controlGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ControlClient
	CallOptions **ControlCallOptions

	// The gRPC API client.
	controlClient retailpb.ControlServiceClient

	operationsClient longrunningpb.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewControlClient creates a new control service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for modifying Control.
func NewControlClient(ctx context.Context, opts ...option.ClientOption) (*ControlClient, error) {
	clientOpts := defaultControlGRPCClientOptions()
	if newControlClientHook != nil {
		hookOpts, err := newControlClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ControlClient{CallOptions: defaultControlCallOptions()}

	c := &controlGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		controlClient:    retailpb.NewControlServiceClient(connPool),
		CallOptions:      &client.CallOptions,
		operationsClient: longrunningpb.NewOperationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *controlGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *controlGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *controlGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *controlGRPCClient) CreateControl(ctx context.Context, req *retailpb.CreateControlRequest, opts ...gax.CallOption) (*retailpb.Control, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateControl[0:len((*c.CallOptions).CreateControl):len((*c.CallOptions).CreateControl)], opts...)
	var resp *retailpb.Control
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.controlClient.CreateControl(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *controlGRPCClient) DeleteControl(ctx context.Context, req *retailpb.DeleteControlRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteControl[0:len((*c.CallOptions).DeleteControl):len((*c.CallOptions).DeleteControl)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.controlClient.DeleteControl(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *controlGRPCClient) UpdateControl(ctx context.Context, req *retailpb.UpdateControlRequest, opts ...gax.CallOption) (*retailpb.Control, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "control.name", url.QueryEscape(req.GetControl().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateControl[0:len((*c.CallOptions).UpdateControl):len((*c.CallOptions).UpdateControl)], opts...)
	var resp *retailpb.Control
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.controlClient.UpdateControl(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *controlGRPCClient) GetControl(ctx context.Context, req *retailpb.GetControlRequest, opts ...gax.CallOption) (*retailpb.Control, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetControl[0:len((*c.CallOptions).GetControl):len((*c.CallOptions).GetControl)], opts...)
	var resp *retailpb.Control
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.controlClient.GetControl(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *controlGRPCClient) ListControls(ctx context.Context, req *retailpb.ListControlsRequest, opts ...gax.CallOption) *ControlIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListControls[0:len((*c.CallOptions).ListControls):len((*c.CallOptions).ListControls)], opts...)
	it := &ControlIterator{}
	req = proto.Clone(req).(*retailpb.ListControlsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*retailpb.Control, string, error) {
		resp := &retailpb.ListControlsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.controlClient.ListControls(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetControls(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *controlGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *controlGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.operationsClient.ListOperations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// ControlIterator manages a stream of *retailpb.Control.
type ControlIterator struct {
	items    []*retailpb.Control
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*retailpb.Control, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ControlIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ControlIterator) Next() (*retailpb.Control, error) {
	var item *retailpb.Control
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ControlIterator) bufLen() int {
	return len(it.items)
}

func (it *ControlIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
