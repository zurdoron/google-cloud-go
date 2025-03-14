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

var newServingConfigClientHook clientHook

// ServingConfigCallOptions contains the retry settings for each method of ServingConfigClient.
type ServingConfigCallOptions struct {
	CreateServingConfig []gax.CallOption
	DeleteServingConfig []gax.CallOption
	UpdateServingConfig []gax.CallOption
	GetServingConfig    []gax.CallOption
	ListServingConfigs  []gax.CallOption
	AddControl          []gax.CallOption
	RemoveControl       []gax.CallOption
	GetOperation        []gax.CallOption
	ListOperations      []gax.CallOption
}

func defaultServingConfigGRPCClientOptions() []option.ClientOption {
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

func defaultServingConfigCallOptions() *ServingConfigCallOptions {
	return &ServingConfigCallOptions{
		CreateServingConfig: []gax.CallOption{},
		DeleteServingConfig: []gax.CallOption{},
		UpdateServingConfig: []gax.CallOption{},
		GetServingConfig:    []gax.CallOption{},
		ListServingConfigs:  []gax.CallOption{},
		AddControl:          []gax.CallOption{},
		RemoveControl:       []gax.CallOption{},
		GetOperation:        []gax.CallOption{},
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

// internalServingConfigClient is an interface that defines the methods available from Retail API.
type internalServingConfigClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateServingConfig(context.Context, *retailpb.CreateServingConfigRequest, ...gax.CallOption) (*retailpb.ServingConfig, error)
	DeleteServingConfig(context.Context, *retailpb.DeleteServingConfigRequest, ...gax.CallOption) error
	UpdateServingConfig(context.Context, *retailpb.UpdateServingConfigRequest, ...gax.CallOption) (*retailpb.ServingConfig, error)
	GetServingConfig(context.Context, *retailpb.GetServingConfigRequest, ...gax.CallOption) (*retailpb.ServingConfig, error)
	ListServingConfigs(context.Context, *retailpb.ListServingConfigsRequest, ...gax.CallOption) *ServingConfigIterator
	AddControl(context.Context, *retailpb.AddControlRequest, ...gax.CallOption) (*retailpb.ServingConfig, error)
	RemoveControl(context.Context, *retailpb.RemoveControlRequest, ...gax.CallOption) (*retailpb.ServingConfig, error)
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
}

// ServingConfigClient is a client for interacting with Retail API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for modifying ServingConfig.
type ServingConfigClient struct {
	// The internal transport-dependent client.
	internalClient internalServingConfigClient

	// The call options for this service.
	CallOptions *ServingConfigCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ServingConfigClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ServingConfigClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ServingConfigClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateServingConfig creates a ServingConfig.
//
// A maximum of 100
// ServingConfigs are allowed in
// a Catalog, otherwise a
// FAILED_PRECONDITION error is returned.
func (c *ServingConfigClient) CreateServingConfig(ctx context.Context, req *retailpb.CreateServingConfigRequest, opts ...gax.CallOption) (*retailpb.ServingConfig, error) {
	return c.internalClient.CreateServingConfig(ctx, req, opts...)
}

// DeleteServingConfig deletes a ServingConfig.
//
// Returns a NotFound error if the ServingConfig does not exist.
func (c *ServingConfigClient) DeleteServingConfig(ctx context.Context, req *retailpb.DeleteServingConfigRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteServingConfig(ctx, req, opts...)
}

// UpdateServingConfig updates a ServingConfig.
func (c *ServingConfigClient) UpdateServingConfig(ctx context.Context, req *retailpb.UpdateServingConfigRequest, opts ...gax.CallOption) (*retailpb.ServingConfig, error) {
	return c.internalClient.UpdateServingConfig(ctx, req, opts...)
}

// GetServingConfig gets a ServingConfig.
//
// Returns a NotFound error if the ServingConfig does not exist.
func (c *ServingConfigClient) GetServingConfig(ctx context.Context, req *retailpb.GetServingConfigRequest, opts ...gax.CallOption) (*retailpb.ServingConfig, error) {
	return c.internalClient.GetServingConfig(ctx, req, opts...)
}

// ListServingConfigs lists all ServingConfigs linked to this catalog.
func (c *ServingConfigClient) ListServingConfigs(ctx context.Context, req *retailpb.ListServingConfigsRequest, opts ...gax.CallOption) *ServingConfigIterator {
	return c.internalClient.ListServingConfigs(ctx, req, opts...)
}

// AddControl enables a Control on the specified ServingConfig.
// The control is added in the last position of the list of controls
// it belongs to (e.g. if it’s a facet spec control it will be applied
// in the last position of servingConfig.facetSpecIds)
// Returns a ALREADY_EXISTS error if the control has already been applied.
// Returns a FAILED_PRECONDITION error if the addition could exceed maximum
// number of control allowed for that type of control.
func (c *ServingConfigClient) AddControl(ctx context.Context, req *retailpb.AddControlRequest, opts ...gax.CallOption) (*retailpb.ServingConfig, error) {
	return c.internalClient.AddControl(ctx, req, opts...)
}

// RemoveControl disables a Control on the specified ServingConfig.
// The control is removed from the ServingConfig.
// Returns a NOT_FOUND error if the Control is not enabled for the
// ServingConfig.
func (c *ServingConfigClient) RemoveControl(ctx context.Context, req *retailpb.RemoveControlRequest, opts ...gax.CallOption) (*retailpb.ServingConfig, error) {
	return c.internalClient.RemoveControl(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *ServingConfigClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *ServingConfigClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// servingConfigGRPCClient is a client for interacting with Retail API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type servingConfigGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ServingConfigClient
	CallOptions **ServingConfigCallOptions

	// The gRPC API client.
	servingConfigClient retailpb.ServingConfigServiceClient

	operationsClient longrunningpb.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewServingConfigClient creates a new serving config service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for modifying ServingConfig.
func NewServingConfigClient(ctx context.Context, opts ...option.ClientOption) (*ServingConfigClient, error) {
	clientOpts := defaultServingConfigGRPCClientOptions()
	if newServingConfigClientHook != nil {
		hookOpts, err := newServingConfigClientHook(ctx, clientHookParams{})
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
	client := ServingConfigClient{CallOptions: defaultServingConfigCallOptions()}

	c := &servingConfigGRPCClient{
		connPool:            connPool,
		disableDeadlines:    disableDeadlines,
		servingConfigClient: retailpb.NewServingConfigServiceClient(connPool),
		CallOptions:         &client.CallOptions,
		operationsClient:    longrunningpb.NewOperationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *servingConfigGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *servingConfigGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *servingConfigGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *servingConfigGRPCClient) CreateServingConfig(ctx context.Context, req *retailpb.CreateServingConfigRequest, opts ...gax.CallOption) (*retailpb.ServingConfig, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateServingConfig[0:len((*c.CallOptions).CreateServingConfig):len((*c.CallOptions).CreateServingConfig)], opts...)
	var resp *retailpb.ServingConfig
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.servingConfigClient.CreateServingConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *servingConfigGRPCClient) DeleteServingConfig(ctx context.Context, req *retailpb.DeleteServingConfigRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteServingConfig[0:len((*c.CallOptions).DeleteServingConfig):len((*c.CallOptions).DeleteServingConfig)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.servingConfigClient.DeleteServingConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *servingConfigGRPCClient) UpdateServingConfig(ctx context.Context, req *retailpb.UpdateServingConfigRequest, opts ...gax.CallOption) (*retailpb.ServingConfig, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "serving_config.name", url.QueryEscape(req.GetServingConfig().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateServingConfig[0:len((*c.CallOptions).UpdateServingConfig):len((*c.CallOptions).UpdateServingConfig)], opts...)
	var resp *retailpb.ServingConfig
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.servingConfigClient.UpdateServingConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *servingConfigGRPCClient) GetServingConfig(ctx context.Context, req *retailpb.GetServingConfigRequest, opts ...gax.CallOption) (*retailpb.ServingConfig, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetServingConfig[0:len((*c.CallOptions).GetServingConfig):len((*c.CallOptions).GetServingConfig)], opts...)
	var resp *retailpb.ServingConfig
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.servingConfigClient.GetServingConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *servingConfigGRPCClient) ListServingConfigs(ctx context.Context, req *retailpb.ListServingConfigsRequest, opts ...gax.CallOption) *ServingConfigIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListServingConfigs[0:len((*c.CallOptions).ListServingConfigs):len((*c.CallOptions).ListServingConfigs)], opts...)
	it := &ServingConfigIterator{}
	req = proto.Clone(req).(*retailpb.ListServingConfigsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*retailpb.ServingConfig, string, error) {
		resp := &retailpb.ListServingConfigsResponse{}
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
			resp, err = c.servingConfigClient.ListServingConfigs(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetServingConfigs(), resp.GetNextPageToken(), nil
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

func (c *servingConfigGRPCClient) AddControl(ctx context.Context, req *retailpb.AddControlRequest, opts ...gax.CallOption) (*retailpb.ServingConfig, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "serving_config", url.QueryEscape(req.GetServingConfig())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).AddControl[0:len((*c.CallOptions).AddControl):len((*c.CallOptions).AddControl)], opts...)
	var resp *retailpb.ServingConfig
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.servingConfigClient.AddControl(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *servingConfigGRPCClient) RemoveControl(ctx context.Context, req *retailpb.RemoveControlRequest, opts ...gax.CallOption) (*retailpb.ServingConfig, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "serving_config", url.QueryEscape(req.GetServingConfig())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).RemoveControl[0:len((*c.CallOptions).RemoveControl):len((*c.CallOptions).RemoveControl)], opts...)
	var resp *retailpb.ServingConfig
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.servingConfigClient.RemoveControl(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *servingConfigGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
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

func (c *servingConfigGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
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

// ServingConfigIterator manages a stream of *retailpb.ServingConfig.
type ServingConfigIterator struct {
	items    []*retailpb.ServingConfig
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
	InternalFetch func(pageSize int, pageToken string) (results []*retailpb.ServingConfig, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ServingConfigIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ServingConfigIterator) Next() (*retailpb.ServingConfig, error) {
	var item *retailpb.ServingConfig
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ServingConfigIterator) bufLen() int {
	return len(it.items)
}

func (it *ServingConfigIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
