// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: internal/api/synchronizer.proto

package ipb

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

const (
	Synchronizer_Synchronize_FullMethodName = "/openmatch.internal.Synchronizer/Synchronize"
)

// SynchronizerClient is the client API for Synchronizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SynchronizerClient interface {
	// Synchronize signals the caller when it is safe to run mmfs, collects the
	// mmfs' proposals, and returns the evaluated matches.
	Synchronize(ctx context.Context, opts ...grpc.CallOption) (Synchronizer_SynchronizeClient, error)
}

type synchronizerClient struct {
	cc grpc.ClientConnInterface
}

func NewSynchronizerClient(cc grpc.ClientConnInterface) SynchronizerClient {
	return &synchronizerClient{cc}
}

func (c *synchronizerClient) Synchronize(ctx context.Context, opts ...grpc.CallOption) (Synchronizer_SynchronizeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Synchronizer_ServiceDesc.Streams[0], Synchronizer_Synchronize_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &synchronizerSynchronizeClient{stream}
	return x, nil
}

type Synchronizer_SynchronizeClient interface {
	Send(*SynchronizeRequest) error
	Recv() (*SynchronizeResponse, error)
	grpc.ClientStream
}

type synchronizerSynchronizeClient struct {
	grpc.ClientStream
}

func (x *synchronizerSynchronizeClient) Send(m *SynchronizeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *synchronizerSynchronizeClient) Recv() (*SynchronizeResponse, error) {
	m := new(SynchronizeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SynchronizerServer is the server API for Synchronizer service.
// All implementations should embed UnimplementedSynchronizerServer
// for forward compatibility
type SynchronizerServer interface {
	// Synchronize signals the caller when it is safe to run mmfs, collects the
	// mmfs' proposals, and returns the evaluated matches.
	Synchronize(Synchronizer_SynchronizeServer) error
}

// UnimplementedSynchronizerServer should be embedded to have forward compatible implementations.
type UnimplementedSynchronizerServer struct {
}

func (UnimplementedSynchronizerServer) Synchronize(Synchronizer_SynchronizeServer) error {
	return status.Errorf(codes.Unimplemented, "method Synchronize not implemented")
}

// UnsafeSynchronizerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SynchronizerServer will
// result in compilation errors.
type UnsafeSynchronizerServer interface {
	mustEmbedUnimplementedSynchronizerServer()
}

func RegisterSynchronizerServer(s grpc.ServiceRegistrar, srv SynchronizerServer) {
	s.RegisterService(&Synchronizer_ServiceDesc, srv)
}

func _Synchronizer_Synchronize_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SynchronizerServer).Synchronize(&synchronizerSynchronizeServer{stream})
}

type Synchronizer_SynchronizeServer interface {
	Send(*SynchronizeResponse) error
	Recv() (*SynchronizeRequest, error)
	grpc.ServerStream
}

type synchronizerSynchronizeServer struct {
	grpc.ServerStream
}

func (x *synchronizerSynchronizeServer) Send(m *SynchronizeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *synchronizerSynchronizeServer) Recv() (*SynchronizeRequest, error) {
	m := new(SynchronizeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Synchronizer_ServiceDesc is the grpc.ServiceDesc for Synchronizer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Synchronizer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openmatch.internal.Synchronizer",
	HandlerType: (*SynchronizerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Synchronize",
			Handler:       _Synchronizer_Synchronize_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internal/api/synchronizer.proto",
}
