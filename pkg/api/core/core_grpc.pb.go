// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CoreClient is the client API for Core service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoreClient interface {
	//
	// ListKustomization lists Kustomizations from a cluster via GitOps.
	ListKustomizations(ctx context.Context, in *ListKustomizationsRequest, opts ...grpc.CallOption) (*ListKustomizationsResponse, error)
	//
	// GetKustomization gets data about a single Kustomization from a cluster.
	GetKustomization(ctx context.Context, in *GetKustomizationRequest, opts ...grpc.CallOption) (*GetKustomizationResponse, error)
	//
	// ListHelmReleases lists helm releases from a cluster.
	ListHelmReleases(ctx context.Context, in *ListHelmReleasesRequest, opts ...grpc.CallOption) (*ListHelmReleasesResponse, error)
	//
	// ListGitRepository lists git repositories objects from a cluster.
	ListGitRepositories(ctx context.Context, in *ListGitRepositoriesRequest, opts ...grpc.CallOption) (*ListGitRepositoriesResponse, error)
	//
	// ListHelmCharts lists helm chart objects from a cluster.
	ListHelmCharts(ctx context.Context, in *ListHelmChartsRequest, opts ...grpc.CallOption) (*ListHelmChartsResponse, error)
	//
	// ListHelmRepository lists helm repository objects from a cluster.
	ListHelmRepositories(ctx context.Context, in *ListHelmRepositoriesRequest, opts ...grpc.CallOption) (*ListHelmRepositoriesResponse, error)
	//
	// ListBuckets lists bucket objects from a cluster.
	ListBuckets(ctx context.Context, in *ListBucketRequest, opts ...grpc.CallOption) (*ListBucketsResponse, error)
	//
	// ListFluxRuntimeObjects lists the flux runtime deployments from a cluster.
	ListFluxRuntimeObjects(ctx context.Context, in *ListFluxRuntimeObjectsRequest, opts ...grpc.CallOption) (*ListFluxRuntimeObjectsResponse, error)
	//
	// GetReconciledObjects returns a list of objects that were created as a result a Flux automation.
	// This list is derived by looking at the Kustomization or HelmRelease specified in the request body.
	GetReconciledObjects(ctx context.Context, in *GetReconciledObjectsRequest, opts ...grpc.CallOption) (*GetReconciledObjectsResponse, error)
	//
	// GetChildObjects returns the children of a given object, specified by a GroupVersionKind.
	// Not all Kubernets objects have children. For example, a Deployment has a child ReplicaSet, but a Service has no child objects.
	GetChildObjects(ctx context.Context, in *GetChildObjectsRequest, opts ...grpc.CallOption) (*GetChildObjectsResponse, error)
	//
	// GetFluxNamespace returns with a namespace with a specific label.
	GetFluxNamespace(ctx context.Context, in *GetFluxNamespaceRequest, opts ...grpc.CallOption) (*GetFluxNamespaceResponse, error)
}

type coreClient struct {
	cc grpc.ClientConnInterface
}

func NewCoreClient(cc grpc.ClientConnInterface) CoreClient {
	return &coreClient{cc}
}

func (c *coreClient) ListKustomizations(ctx context.Context, in *ListKustomizationsRequest, opts ...grpc.CallOption) (*ListKustomizationsResponse, error) {
	out := new(ListKustomizationsResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/ListKustomizations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetKustomization(ctx context.Context, in *GetKustomizationRequest, opts ...grpc.CallOption) (*GetKustomizationResponse, error) {
	out := new(GetKustomizationResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/GetKustomization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ListHelmReleases(ctx context.Context, in *ListHelmReleasesRequest, opts ...grpc.CallOption) (*ListHelmReleasesResponse, error) {
	out := new(ListHelmReleasesResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/ListHelmReleases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ListGitRepositories(ctx context.Context, in *ListGitRepositoriesRequest, opts ...grpc.CallOption) (*ListGitRepositoriesResponse, error) {
	out := new(ListGitRepositoriesResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/ListGitRepositories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ListHelmCharts(ctx context.Context, in *ListHelmChartsRequest, opts ...grpc.CallOption) (*ListHelmChartsResponse, error) {
	out := new(ListHelmChartsResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/ListHelmCharts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ListHelmRepositories(ctx context.Context, in *ListHelmRepositoriesRequest, opts ...grpc.CallOption) (*ListHelmRepositoriesResponse, error) {
	out := new(ListHelmRepositoriesResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/ListHelmRepositories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ListBuckets(ctx context.Context, in *ListBucketRequest, opts ...grpc.CallOption) (*ListBucketsResponse, error) {
	out := new(ListBucketsResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/ListBuckets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ListFluxRuntimeObjects(ctx context.Context, in *ListFluxRuntimeObjectsRequest, opts ...grpc.CallOption) (*ListFluxRuntimeObjectsResponse, error) {
	out := new(ListFluxRuntimeObjectsResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/ListFluxRuntimeObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetReconciledObjects(ctx context.Context, in *GetReconciledObjectsRequest, opts ...grpc.CallOption) (*GetReconciledObjectsResponse, error) {
	out := new(GetReconciledObjectsResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/GetReconciledObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetChildObjects(ctx context.Context, in *GetChildObjectsRequest, opts ...grpc.CallOption) (*GetChildObjectsResponse, error) {
	out := new(GetChildObjectsResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/GetChildObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetFluxNamespace(ctx context.Context, in *GetFluxNamespaceRequest, opts ...grpc.CallOption) (*GetFluxNamespaceResponse, error) {
	out := new(GetFluxNamespaceResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/GetFluxNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreServer is the server API for Core service.
// All implementations must embed UnimplementedCoreServer
// for forward compatibility
type CoreServer interface {
	//
	// ListKustomization lists Kustomizations from a cluster via GitOps.
	ListKustomizations(context.Context, *ListKustomizationsRequest) (*ListKustomizationsResponse, error)
	//
	// GetKustomization gets data about a single Kustomization from a cluster.
	GetKustomization(context.Context, *GetKustomizationRequest) (*GetKustomizationResponse, error)
	//
	// ListHelmReleases lists helm releases from a cluster.
	ListHelmReleases(context.Context, *ListHelmReleasesRequest) (*ListHelmReleasesResponse, error)
	//
	// ListGitRepository lists git repositories objects from a cluster.
	ListGitRepositories(context.Context, *ListGitRepositoriesRequest) (*ListGitRepositoriesResponse, error)
	//
	// ListHelmCharts lists helm chart objects from a cluster.
	ListHelmCharts(context.Context, *ListHelmChartsRequest) (*ListHelmChartsResponse, error)
	//
	// ListHelmRepository lists helm repository objects from a cluster.
	ListHelmRepositories(context.Context, *ListHelmRepositoriesRequest) (*ListHelmRepositoriesResponse, error)
	//
	// ListBuckets lists bucket objects from a cluster.
	ListBuckets(context.Context, *ListBucketRequest) (*ListBucketsResponse, error)
	//
	// ListFluxRuntimeObjects lists the flux runtime deployments from a cluster.
	ListFluxRuntimeObjects(context.Context, *ListFluxRuntimeObjectsRequest) (*ListFluxRuntimeObjectsResponse, error)
	//
	// GetReconciledObjects returns a list of objects that were created as a result a Flux automation.
	// This list is derived by looking at the Kustomization or HelmRelease specified in the request body.
	GetReconciledObjects(context.Context, *GetReconciledObjectsRequest) (*GetReconciledObjectsResponse, error)
	//
	// GetChildObjects returns the children of a given object, specified by a GroupVersionKind.
	// Not all Kubernets objects have children. For example, a Deployment has a child ReplicaSet, but a Service has no child objects.
	GetChildObjects(context.Context, *GetChildObjectsRequest) (*GetChildObjectsResponse, error)
	//
	// GetFluxNamespace returns with a namespace with a specific label.
	GetFluxNamespace(context.Context, *GetFluxNamespaceRequest) (*GetFluxNamespaceResponse, error)
	mustEmbedUnimplementedCoreServer()
}

// UnimplementedCoreServer must be embedded to have forward compatible implementations.
type UnimplementedCoreServer struct {
}

func (UnimplementedCoreServer) ListKustomizations(context.Context, *ListKustomizationsRequest) (*ListKustomizationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKustomizations not implemented")
}
func (UnimplementedCoreServer) GetKustomization(context.Context, *GetKustomizationRequest) (*GetKustomizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKustomization not implemented")
}
func (UnimplementedCoreServer) ListHelmReleases(context.Context, *ListHelmReleasesRequest) (*ListHelmReleasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHelmReleases not implemented")
}
func (UnimplementedCoreServer) ListGitRepositories(context.Context, *ListGitRepositoriesRequest) (*ListGitRepositoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGitRepositories not implemented")
}
func (UnimplementedCoreServer) ListHelmCharts(context.Context, *ListHelmChartsRequest) (*ListHelmChartsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHelmCharts not implemented")
}
func (UnimplementedCoreServer) ListHelmRepositories(context.Context, *ListHelmRepositoriesRequest) (*ListHelmRepositoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHelmRepositories not implemented")
}
func (UnimplementedCoreServer) ListBuckets(context.Context, *ListBucketRequest) (*ListBucketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBuckets not implemented")
}
func (UnimplementedCoreServer) ListFluxRuntimeObjects(context.Context, *ListFluxRuntimeObjectsRequest) (*ListFluxRuntimeObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFluxRuntimeObjects not implemented")
}
func (UnimplementedCoreServer) GetReconciledObjects(context.Context, *GetReconciledObjectsRequest) (*GetReconciledObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReconciledObjects not implemented")
}
func (UnimplementedCoreServer) GetChildObjects(context.Context, *GetChildObjectsRequest) (*GetChildObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChildObjects not implemented")
}
func (UnimplementedCoreServer) GetFluxNamespace(context.Context, *GetFluxNamespaceRequest) (*GetFluxNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFluxNamespace not implemented")
}
func (UnimplementedCoreServer) mustEmbedUnimplementedCoreServer() {}

// UnsafeCoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoreServer will
// result in compilation errors.
type UnsafeCoreServer interface {
	mustEmbedUnimplementedCoreServer()
}

func RegisterCoreServer(s grpc.ServiceRegistrar, srv CoreServer) {
	s.RegisterService(&Core_ServiceDesc, srv)
}

func _Core_ListKustomizations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKustomizationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ListKustomizations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/ListKustomizations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ListKustomizations(ctx, req.(*ListKustomizationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetKustomization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKustomizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetKustomization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/GetKustomization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetKustomization(ctx, req.(*GetKustomizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ListHelmReleases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHelmReleasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ListHelmReleases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/ListHelmReleases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ListHelmReleases(ctx, req.(*ListHelmReleasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ListGitRepositories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGitRepositoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ListGitRepositories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/ListGitRepositories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ListGitRepositories(ctx, req.(*ListGitRepositoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ListHelmCharts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHelmChartsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ListHelmCharts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/ListHelmCharts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ListHelmCharts(ctx, req.(*ListHelmChartsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ListHelmRepositories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHelmRepositoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ListHelmRepositories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/ListHelmRepositories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ListHelmRepositories(ctx, req.(*ListHelmRepositoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ListBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ListBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/ListBuckets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ListBuckets(ctx, req.(*ListBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ListFluxRuntimeObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFluxRuntimeObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ListFluxRuntimeObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/ListFluxRuntimeObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ListFluxRuntimeObjects(ctx, req.(*ListFluxRuntimeObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetReconciledObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReconciledObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetReconciledObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/GetReconciledObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetReconciledObjects(ctx, req.(*GetReconciledObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetChildObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChildObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetChildObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/GetChildObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetChildObjects(ctx, req.(*GetChildObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetFluxNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFluxNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetFluxNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/GetFluxNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetFluxNamespace(ctx, req.(*GetFluxNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Core_ServiceDesc is the grpc.ServiceDesc for Core service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Core_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitops_core.v1.Core",
	HandlerType: (*CoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListKustomizations",
			Handler:    _Core_ListKustomizations_Handler,
		},
		{
			MethodName: "GetKustomization",
			Handler:    _Core_GetKustomization_Handler,
		},
		{
			MethodName: "ListHelmReleases",
			Handler:    _Core_ListHelmReleases_Handler,
		},
		{
			MethodName: "ListGitRepositories",
			Handler:    _Core_ListGitRepositories_Handler,
		},
		{
			MethodName: "ListHelmCharts",
			Handler:    _Core_ListHelmCharts_Handler,
		},
		{
			MethodName: "ListHelmRepositories",
			Handler:    _Core_ListHelmRepositories_Handler,
		},
		{
			MethodName: "ListBuckets",
			Handler:    _Core_ListBuckets_Handler,
		},
		{
			MethodName: "ListFluxRuntimeObjects",
			Handler:    _Core_ListFluxRuntimeObjects_Handler,
		},
		{
			MethodName: "GetReconciledObjects",
			Handler:    _Core_GetReconciledObjects_Handler,
		},
		{
			MethodName: "GetChildObjects",
			Handler:    _Core_GetChildObjects_Handler,
		},
		{
			MethodName: "GetFluxNamespace",
			Handler:    _Core_GetFluxNamespace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/core/core.proto",
}
