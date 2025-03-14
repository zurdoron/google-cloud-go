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

package appengine_test

import (
	"context"

	appengine "cloud.google.com/go/appengine/apiv1"
	"google.golang.org/api/iterator"
	appenginepb "google.golang.org/genproto/googleapis/appengine/v1"
)

func ExampleNewAuthorizedCertificatesClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := appengine.NewAuthorizedCertificatesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleAuthorizedCertificatesClient_ListAuthorizedCertificates() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := appengine.NewAuthorizedCertificatesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &appenginepb.ListAuthorizedCertificatesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/appengine/v1#ListAuthorizedCertificatesRequest.
	}
	it := c.ListAuthorizedCertificates(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleAuthorizedCertificatesClient_GetAuthorizedCertificate() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := appengine.NewAuthorizedCertificatesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &appenginepb.GetAuthorizedCertificateRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/appengine/v1#GetAuthorizedCertificateRequest.
	}
	resp, err := c.GetAuthorizedCertificate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAuthorizedCertificatesClient_CreateAuthorizedCertificate() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := appengine.NewAuthorizedCertificatesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &appenginepb.CreateAuthorizedCertificateRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/appengine/v1#CreateAuthorizedCertificateRequest.
	}
	resp, err := c.CreateAuthorizedCertificate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAuthorizedCertificatesClient_UpdateAuthorizedCertificate() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := appengine.NewAuthorizedCertificatesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &appenginepb.UpdateAuthorizedCertificateRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/appengine/v1#UpdateAuthorizedCertificateRequest.
	}
	resp, err := c.UpdateAuthorizedCertificate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAuthorizedCertificatesClient_DeleteAuthorizedCertificate() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := appengine.NewAuthorizedCertificatesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &appenginepb.DeleteAuthorizedCertificateRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/appengine/v1#DeleteAuthorizedCertificateRequest.
	}
	err = c.DeleteAuthorizedCertificate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}
