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
	Valid(ctx context.Context, in *PanierRequest, opts ...grpc.CallOption) (*CommandeResponse, error)
	Confirm(ctx context.Context, in *CommandeRequest, opts ...grpc.CallOption) (*CommandeResponse, error)
	Cancel(ctx context.Context, in *CommandeRequest, opts ...grpc.CallOption) (*CommandeResponse, error)
	GetUserCommandes(ctx context.Context, in *ByUtilisateurRequest, opts ...grpc.CallOption) (*CommandesResponse, error)
}

type serviceCommandeClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceCommandeClient(cc grpc.ClientConnInterface) ServiceCommandeClient {
	return &serviceCommandeClient{cc}
}

func (c *serviceCommandeClient) Valid(ctx context.Context, in *PanierRequest, opts ...grpc.CallOption) (*CommandeResponse, error) {
	out := new(CommandeResponse)
	err := c.cc.Invoke(ctx, "/commande.ServiceCommande/Valid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceCommandeClient) Confirm(ctx context.Context, in *CommandeRequest, opts ...grpc.CallOption) (*CommandeResponse, error) {
	out := new(CommandeResponse)
	err := c.cc.Invoke(ctx, "/commande.ServiceCommande/Confirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceCommandeClient) Cancel(ctx context.Context, in *CommandeRequest, opts ...grpc.CallOption) (*CommandeResponse, error) {
	out := new(CommandeResponse)
	err := c.cc.Invoke(ctx, "/commande.ServiceCommande/Cancel", in, out, opts...)
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
	Valid(context.Context, *PanierRequest) (*CommandeResponse, error)
	Confirm(context.Context, *CommandeRequest) (*CommandeResponse, error)
	Cancel(context.Context, *CommandeRequest) (*CommandeResponse, error)
	GetUserCommandes(context.Context, *ByUtilisateurRequest) (*CommandesResponse, error)
	mustEmbedUnimplementedServiceCommandeServer()
}

// UnimplementedServiceCommandeServer must be embedded to have forward compatible implementations.
type UnimplementedServiceCommandeServer struct {
}

func (UnimplementedServiceCommandeServer) Valid(context.Context, *PanierRequest) (*CommandeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Valid not implemented")
}
func (UnimplementedServiceCommandeServer) Confirm(context.Context, *CommandeRequest) (*CommandeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Confirm not implemented")
}
func (UnimplementedServiceCommandeServer) Cancel(context.Context, *CommandeRequest) (*CommandeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
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

func _ServiceCommande_Valid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PanierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCommandeServer).Valid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commande.ServiceCommande/Valid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCommandeServer).Valid(ctx, req.(*PanierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceCommande_Confirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCommandeServer).Confirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commande.ServiceCommande/Confirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCommandeServer).Confirm(ctx, req.(*CommandeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceCommande_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCommandeServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commande.ServiceCommande/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCommandeServer).Cancel(ctx, req.(*CommandeRequest))
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
			MethodName: "Valid",
			Handler:    _ServiceCommande_Valid_Handler,
		},
		{
			MethodName: "Confirm",
			Handler:    _ServiceCommande_Confirm_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _ServiceCommande_Cancel_Handler,
		},
		{
			MethodName: "GetUserCommandes",
			Handler:    _ServiceCommande_GetUserCommandes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/commande.proto",
}
