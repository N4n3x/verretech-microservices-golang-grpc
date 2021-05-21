// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package commandepb

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

// ServiceCommandeClient is the client API for ServiceCommande service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceCommandeClient interface {
	ValidateCommande(ctx context.Context, in *CommandeRequest, opts ...grpc.CallOption) (*CommandeResponse, error)
	SubmitCommande(ctx context.Context, in *CommandeRequest, opts ...grpc.CallOption) (*CommandeResponse, error)
	GetLastCommandes(ctx context.Context, in *LastCommandesRequest, opts ...grpc.CallOption) (*CommandesResponse, error)
	GetUserCommandes(ctx context.Context, in *ByUtilisateurRequest, opts ...grpc.CallOption) (*CommandesResponse, error)
}

type serviceCommandeClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceCommandeClient(cc grpc.ClientConnInterface) ServiceCommandeClient {
	return &serviceCommandeClient{cc}
}

func (c *serviceCommandeClient) ValidateCommande(ctx context.Context, in *CommandeRequest, opts ...grpc.CallOption) (*CommandeResponse, error) {
	out := new(CommandeResponse)
	err := c.cc.Invoke(ctx, "/commande.ServiceCommande/ValidateCommande", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceCommandeClient) SubmitCommande(ctx context.Context, in *CommandeRequest, opts ...grpc.CallOption) (*CommandeResponse, error) {
	out := new(CommandeResponse)
	err := c.cc.Invoke(ctx, "/commande.ServiceCommande/SubmitCommande", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceCommandeClient) GetLastCommandes(ctx context.Context, in *LastCommandesRequest, opts ...grpc.CallOption) (*CommandesResponse, error) {
	out := new(CommandesResponse)
	err := c.cc.Invoke(ctx, "/commande.ServiceCommande/GetLastCommandes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceCommandeClient) GetUserCommandes(ctx context.Context, in *ByUtilisateurRequest, opts ...grpc.CallOption) (*CommandesResponse, error) {
	out := new(CommandesResponse)
	err := c.cc.Invoke(ctx, "/commande.ServiceCommande/GetUserCommandes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceCommandeServer is the server API for ServiceCommande service.
// All implementations must embed UnimplementedServiceCommandeServer
// for forward compatibility
type ServiceCommandeServer interface {
	ValidateCommande(context.Context, *CommandeRequest) (*CommandeResponse, error)
	SubmitCommande(context.Context, *CommandeRequest) (*CommandeResponse, error)
	GetLastCommandes(context.Context, *LastCommandesRequest) (*CommandesResponse, error)
	GetUserCommandes(context.Context, *ByUtilisateurRequest) (*CommandesResponse, error)
	mustEmbedUnimplementedServiceCommandeServer()
}

// UnimplementedServiceCommandeServer must be embedded to have forward compatible implementations.
type UnimplementedServiceCommandeServer struct {
}

func (UnimplementedServiceCommandeServer) ValidateCommande(context.Context, *CommandeRequest) (*CommandeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateCommande not implemented")
}
func (UnimplementedServiceCommandeServer) SubmitCommande(context.Context, *CommandeRequest) (*CommandeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitCommande not implemented")
}
func (UnimplementedServiceCommandeServer) GetLastCommandes(context.Context, *LastCommandesRequest) (*CommandesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastCommandes not implemented")
}
func (UnimplementedServiceCommandeServer) GetUserCommandes(context.Context, *ByUtilisateurRequest) (*CommandesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCommandes not implemented")
}
func (UnimplementedServiceCommandeServer) mustEmbedUnimplementedServiceCommandeServer() {}

// UnsafeServiceCommandeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceCommandeServer will
// result in compilation errors.
type UnsafeServiceCommandeServer interface {
	mustEmbedUnimplementedServiceCommandeServer()
}

func RegisterServiceCommandeServer(s grpc.ServiceRegistrar, srv ServiceCommandeServer) {
	s.RegisterService(&ServiceCommande_ServiceDesc, srv)
}

func _ServiceCommande_ValidateCommande_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCommandeServer).ValidateCommande(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commande.ServiceCommande/ValidateCommande",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCommandeServer).ValidateCommande(ctx, req.(*CommandeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceCommande_SubmitCommande_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCommandeServer).SubmitCommande(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commande.ServiceCommande/SubmitCommande",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCommandeServer).SubmitCommande(ctx, req.(*CommandeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceCommande_GetLastCommandes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LastCommandesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCommandeServer).GetLastCommandes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commande.ServiceCommande/GetLastCommandes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCommandeServer).GetLastCommandes(ctx, req.(*LastCommandesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceCommande_GetUserCommandes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByUtilisateurRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCommandeServer).GetUserCommandes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commande.ServiceCommande/GetUserCommandes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCommandeServer).GetUserCommandes(ctx, req.(*ByUtilisateurRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceCommande_ServiceDesc is the grpc.ServiceDesc for ServiceCommande service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceCommande_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "commande.ServiceCommande",
	HandlerType: (*ServiceCommandeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateCommande",
			Handler:    _ServiceCommande_ValidateCommande_Handler,
		},
		{
			MethodName: "SubmitCommande",
			Handler:    _ServiceCommande_SubmitCommande_Handler,
		},
		{
			MethodName: "GetLastCommandes",
			Handler:    _ServiceCommande_GetLastCommandes_Handler,
		},
		{
			MethodName: "GetUserCommandes",
			Handler:    _ServiceCommande_GetUserCommandes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/commande.proto",
}
