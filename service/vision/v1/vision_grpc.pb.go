// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// VisionServiceClient is the client API for VisionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VisionServiceClient interface {
	// GetModelParameterSchema takes the model name and returns the parameters needed to add one to the vision registry.
	GetModelParameterSchema(ctx context.Context, in *GetModelParameterSchemaRequest, opts ...grpc.CallOption) (*GetModelParameterSchemaResponse, error)
	// GetDetectorNames returns the list of detectors in the registry.
	GetDetectorNames(ctx context.Context, in *GetDetectorNamesRequest, opts ...grpc.CallOption) (*GetDetectorNamesResponse, error)
	// AddDetector adds a new detector to the registry.
	AddDetector(ctx context.Context, in *AddDetectorRequest, opts ...grpc.CallOption) (*AddDetectorResponse, error)
	// RemoveDetector removes a detector from the registry.
	RemoveDetector(ctx context.Context, in *RemoveDetectorRequest, opts ...grpc.CallOption) (*RemoveDetectorResponse, error)
	// GetDetectionsFromCamera will return a list of detections in the next image given a camera and a detector
	GetDetectionsFromCamera(ctx context.Context, in *GetDetectionsFromCameraRequest, opts ...grpc.CallOption) (*GetDetectionsFromCameraResponse, error)
	// GetDetections will return a list of detections in the next image given the image bytes and a detector
	GetDetections(ctx context.Context, in *GetDetectionsRequest, opts ...grpc.CallOption) (*GetDetectionsResponse, error)
	// GetClassifierNames returns the list of classifiers in the registry.
	GetClassifierNames(ctx context.Context, in *GetClassifierNamesRequest, opts ...grpc.CallOption) (*GetClassifierNamesResponse, error)
	// AddClassifier adds a new classifier to the registry.
	AddClassifier(ctx context.Context, in *AddClassifierRequest, opts ...grpc.CallOption) (*AddClassifierResponse, error)
	// RemoveClassifier adds a new classifier to the registry.
	RemoveClassifier(ctx context.Context, in *RemoveClassifierRequest, opts ...grpc.CallOption) (*RemoveClassifierResponse, error)
	// GetClassificationsFromCamera will return a list of classifications in the next image given a camera and a classifier
	GetClassificationsFromCamera(ctx context.Context, in *GetClassificationsFromCameraRequest, opts ...grpc.CallOption) (*GetClassificationsFromCameraResponse, error)
	// GetClassifications will return a list of classifications in the next image given the image bytes and a classifier
	GetClassifications(ctx context.Context, in *GetClassificationsRequest, opts ...grpc.CallOption) (*GetClassificationsResponse, error)
	// GetSegmenterNames returns the list of segmenters in the registry.
	GetSegmenterNames(ctx context.Context, in *GetSegmenterNamesRequest, opts ...grpc.CallOption) (*GetSegmenterNamesResponse, error)
	// AddSegmenter adds a new segmenter to the registry.
	AddSegmenter(ctx context.Context, in *AddSegmenterRequest, opts ...grpc.CallOption) (*AddSegmenterResponse, error)
	// RemoveSegmenter removes a segmenter from the registry.
	RemoveSegmenter(ctx context.Context, in *RemoveSegmenterRequest, opts ...grpc.CallOption) (*RemoveSegmenterResponse, error)
	// GetObjectPointClouds returns all the found objects in a pointcloud from a camera of the underlying robot,
	// as well as the 3-vector center of each of the found objects.
	// A specific MIME type can be requested but may not necessarily be the same one returned.
	GetObjectPointClouds(ctx context.Context, in *GetObjectPointCloudsRequest, opts ...grpc.CallOption) (*GetObjectPointCloudsResponse, error)
}

type visionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVisionServiceClient(cc grpc.ClientConnInterface) VisionServiceClient {
	return &visionServiceClient{cc}
}

func (c *visionServiceClient) GetModelParameterSchema(ctx context.Context, in *GetModelParameterSchemaRequest, opts ...grpc.CallOption) (*GetModelParameterSchemaResponse, error) {
	out := new(GetModelParameterSchemaResponse)
	err := c.cc.Invoke(ctx, "/viam.service.vision.v1.VisionService/GetModelParameterSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visionServiceClient) GetDetectorNames(ctx context.Context, in *GetDetectorNamesRequest, opts ...grpc.CallOption) (*GetDetectorNamesResponse, error) {
	out := new(GetDetectorNamesResponse)
	err := c.cc.Invoke(ctx, "/viam.service.vision.v1.VisionService/GetDetectorNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visionServiceClient) AddDetector(ctx context.Context, in *AddDetectorRequest, opts ...grpc.CallOption) (*AddDetectorResponse, error) {
	out := new(AddDetectorResponse)
	err := c.cc.Invoke(ctx, "/viam.service.vision.v1.VisionService/AddDetector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visionServiceClient) RemoveDetector(ctx context.Context, in *RemoveDetectorRequest, opts ...grpc.CallOption) (*RemoveDetectorResponse, error) {
	out := new(RemoveDetectorResponse)
	err := c.cc.Invoke(ctx, "/viam.service.vision.v1.VisionService/RemoveDetector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visionServiceClient) GetDetectionsFromCamera(ctx context.Context, in *GetDetectionsFromCameraRequest, opts ...grpc.CallOption) (*GetDetectionsFromCameraResponse, error) {
	out := new(GetDetectionsFromCameraResponse)
	err := c.cc.Invoke(ctx, "/viam.service.vision.v1.VisionService/GetDetectionsFromCamera", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visionServiceClient) GetDetections(ctx context.Context, in *GetDetectionsRequest, opts ...grpc.CallOption) (*GetDetectionsResponse, error) {
	out := new(GetDetectionsResponse)
	err := c.cc.Invoke(ctx, "/viam.service.vision.v1.VisionService/GetDetections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visionServiceClient) GetClassifierNames(ctx context.Context, in *GetClassifierNamesRequest, opts ...grpc.CallOption) (*GetClassifierNamesResponse, error) {
	out := new(GetClassifierNamesResponse)
	err := c.cc.Invoke(ctx, "/viam.service.vision.v1.VisionService/GetClassifierNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visionServiceClient) AddClassifier(ctx context.Context, in *AddClassifierRequest, opts ...grpc.CallOption) (*AddClassifierResponse, error) {
	out := new(AddClassifierResponse)
	err := c.cc.Invoke(ctx, "/viam.service.vision.v1.VisionService/AddClassifier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visionServiceClient) RemoveClassifier(ctx context.Context, in *RemoveClassifierRequest, opts ...grpc.CallOption) (*RemoveClassifierResponse, error) {
	out := new(RemoveClassifierResponse)
	err := c.cc.Invoke(ctx, "/viam.service.vision.v1.VisionService/RemoveClassifier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visionServiceClient) GetClassificationsFromCamera(ctx context.Context, in *GetClassificationsFromCameraRequest, opts ...grpc.CallOption) (*GetClassificationsFromCameraResponse, error) {
	out := new(GetClassificationsFromCameraResponse)
	err := c.cc.Invoke(ctx, "/viam.service.vision.v1.VisionService/GetClassificationsFromCamera", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visionServiceClient) GetClassifications(ctx context.Context, in *GetClassificationsRequest, opts ...grpc.CallOption) (*GetClassificationsResponse, error) {
	out := new(GetClassificationsResponse)
	err := c.cc.Invoke(ctx, "/viam.service.vision.v1.VisionService/GetClassifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visionServiceClient) GetSegmenterNames(ctx context.Context, in *GetSegmenterNamesRequest, opts ...grpc.CallOption) (*GetSegmenterNamesResponse, error) {
	out := new(GetSegmenterNamesResponse)
	err := c.cc.Invoke(ctx, "/viam.service.vision.v1.VisionService/GetSegmenterNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visionServiceClient) AddSegmenter(ctx context.Context, in *AddSegmenterRequest, opts ...grpc.CallOption) (*AddSegmenterResponse, error) {
	out := new(AddSegmenterResponse)
	err := c.cc.Invoke(ctx, "/viam.service.vision.v1.VisionService/AddSegmenter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visionServiceClient) RemoveSegmenter(ctx context.Context, in *RemoveSegmenterRequest, opts ...grpc.CallOption) (*RemoveSegmenterResponse, error) {
	out := new(RemoveSegmenterResponse)
	err := c.cc.Invoke(ctx, "/viam.service.vision.v1.VisionService/RemoveSegmenter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visionServiceClient) GetObjectPointClouds(ctx context.Context, in *GetObjectPointCloudsRequest, opts ...grpc.CallOption) (*GetObjectPointCloudsResponse, error) {
	out := new(GetObjectPointCloudsResponse)
	err := c.cc.Invoke(ctx, "/viam.service.vision.v1.VisionService/GetObjectPointClouds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VisionServiceServer is the server API for VisionService service.
// All implementations must embed UnimplementedVisionServiceServer
// for forward compatibility
type VisionServiceServer interface {
	// GetModelParameterSchema takes the model name and returns the parameters needed to add one to the vision registry.
	GetModelParameterSchema(context.Context, *GetModelParameterSchemaRequest) (*GetModelParameterSchemaResponse, error)
	// GetDetectorNames returns the list of detectors in the registry.
	GetDetectorNames(context.Context, *GetDetectorNamesRequest) (*GetDetectorNamesResponse, error)
	// AddDetector adds a new detector to the registry.
	AddDetector(context.Context, *AddDetectorRequest) (*AddDetectorResponse, error)
	// RemoveDetector removes a detector from the registry.
	RemoveDetector(context.Context, *RemoveDetectorRequest) (*RemoveDetectorResponse, error)
	// GetDetectionsFromCamera will return a list of detections in the next image given a camera and a detector
	GetDetectionsFromCamera(context.Context, *GetDetectionsFromCameraRequest) (*GetDetectionsFromCameraResponse, error)
	// GetDetections will return a list of detections in the next image given the image bytes and a detector
	GetDetections(context.Context, *GetDetectionsRequest) (*GetDetectionsResponse, error)
	// GetClassifierNames returns the list of classifiers in the registry.
	GetClassifierNames(context.Context, *GetClassifierNamesRequest) (*GetClassifierNamesResponse, error)
	// AddClassifier adds a new classifier to the registry.
	AddClassifier(context.Context, *AddClassifierRequest) (*AddClassifierResponse, error)
	// RemoveClassifier adds a new classifier to the registry.
	RemoveClassifier(context.Context, *RemoveClassifierRequest) (*RemoveClassifierResponse, error)
	// GetClassificationsFromCamera will return a list of classifications in the next image given a camera and a classifier
	GetClassificationsFromCamera(context.Context, *GetClassificationsFromCameraRequest) (*GetClassificationsFromCameraResponse, error)
	// GetClassifications will return a list of classifications in the next image given the image bytes and a classifier
	GetClassifications(context.Context, *GetClassificationsRequest) (*GetClassificationsResponse, error)
	// GetSegmenterNames returns the list of segmenters in the registry.
	GetSegmenterNames(context.Context, *GetSegmenterNamesRequest) (*GetSegmenterNamesResponse, error)
	// AddSegmenter adds a new segmenter to the registry.
	AddSegmenter(context.Context, *AddSegmenterRequest) (*AddSegmenterResponse, error)
	// RemoveSegmenter removes a segmenter from the registry.
	RemoveSegmenter(context.Context, *RemoveSegmenterRequest) (*RemoveSegmenterResponse, error)
	// GetObjectPointClouds returns all the found objects in a pointcloud from a camera of the underlying robot,
	// as well as the 3-vector center of each of the found objects.
	// A specific MIME type can be requested but may not necessarily be the same one returned.
	GetObjectPointClouds(context.Context, *GetObjectPointCloudsRequest) (*GetObjectPointCloudsResponse, error)
	mustEmbedUnimplementedVisionServiceServer()
}

// UnimplementedVisionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVisionServiceServer struct {
}

func (UnimplementedVisionServiceServer) GetModelParameterSchema(context.Context, *GetModelParameterSchemaRequest) (*GetModelParameterSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModelParameterSchema not implemented")
}
func (UnimplementedVisionServiceServer) GetDetectorNames(context.Context, *GetDetectorNamesRequest) (*GetDetectorNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetectorNames not implemented")
}
func (UnimplementedVisionServiceServer) AddDetector(context.Context, *AddDetectorRequest) (*AddDetectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDetector not implemented")
}
func (UnimplementedVisionServiceServer) RemoveDetector(context.Context, *RemoveDetectorRequest) (*RemoveDetectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDetector not implemented")
}
func (UnimplementedVisionServiceServer) GetDetectionsFromCamera(context.Context, *GetDetectionsFromCameraRequest) (*GetDetectionsFromCameraResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetectionsFromCamera not implemented")
}
func (UnimplementedVisionServiceServer) GetDetections(context.Context, *GetDetectionsRequest) (*GetDetectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetections not implemented")
}
func (UnimplementedVisionServiceServer) GetClassifierNames(context.Context, *GetClassifierNamesRequest) (*GetClassifierNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassifierNames not implemented")
}
func (UnimplementedVisionServiceServer) AddClassifier(context.Context, *AddClassifierRequest) (*AddClassifierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddClassifier not implemented")
}
func (UnimplementedVisionServiceServer) RemoveClassifier(context.Context, *RemoveClassifierRequest) (*RemoveClassifierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveClassifier not implemented")
}
func (UnimplementedVisionServiceServer) GetClassificationsFromCamera(context.Context, *GetClassificationsFromCameraRequest) (*GetClassificationsFromCameraResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassificationsFromCamera not implemented")
}
func (UnimplementedVisionServiceServer) GetClassifications(context.Context, *GetClassificationsRequest) (*GetClassificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassifications not implemented")
}
func (UnimplementedVisionServiceServer) GetSegmenterNames(context.Context, *GetSegmenterNamesRequest) (*GetSegmenterNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSegmenterNames not implemented")
}
func (UnimplementedVisionServiceServer) AddSegmenter(context.Context, *AddSegmenterRequest) (*AddSegmenterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSegmenter not implemented")
}
func (UnimplementedVisionServiceServer) RemoveSegmenter(context.Context, *RemoveSegmenterRequest) (*RemoveSegmenterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSegmenter not implemented")
}
func (UnimplementedVisionServiceServer) GetObjectPointClouds(context.Context, *GetObjectPointCloudsRequest) (*GetObjectPointCloudsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectPointClouds not implemented")
}
func (UnimplementedVisionServiceServer) mustEmbedUnimplementedVisionServiceServer() {}

// UnsafeVisionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VisionServiceServer will
// result in compilation errors.
type UnsafeVisionServiceServer interface {
	mustEmbedUnimplementedVisionServiceServer()
}

func RegisterVisionServiceServer(s grpc.ServiceRegistrar, srv VisionServiceServer) {
	s.RegisterService(&VisionService_ServiceDesc, srv)
}

func _VisionService_GetModelParameterSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelParameterSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).GetModelParameterSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.vision.v1.VisionService/GetModelParameterSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).GetModelParameterSchema(ctx, req.(*GetModelParameterSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisionService_GetDetectorNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetectorNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).GetDetectorNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.vision.v1.VisionService/GetDetectorNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).GetDetectorNames(ctx, req.(*GetDetectorNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisionService_AddDetector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDetectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).AddDetector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.vision.v1.VisionService/AddDetector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).AddDetector(ctx, req.(*AddDetectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisionService_RemoveDetector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDetectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).RemoveDetector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.vision.v1.VisionService/RemoveDetector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).RemoveDetector(ctx, req.(*RemoveDetectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisionService_GetDetectionsFromCamera_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetectionsFromCameraRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).GetDetectionsFromCamera(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.vision.v1.VisionService/GetDetectionsFromCamera",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).GetDetectionsFromCamera(ctx, req.(*GetDetectionsFromCameraRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisionService_GetDetections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).GetDetections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.vision.v1.VisionService/GetDetections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).GetDetections(ctx, req.(*GetDetectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisionService_GetClassifierNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassifierNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).GetClassifierNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.vision.v1.VisionService/GetClassifierNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).GetClassifierNames(ctx, req.(*GetClassifierNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisionService_AddClassifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddClassifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).AddClassifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.vision.v1.VisionService/AddClassifier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).AddClassifier(ctx, req.(*AddClassifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisionService_RemoveClassifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveClassifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).RemoveClassifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.vision.v1.VisionService/RemoveClassifier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).RemoveClassifier(ctx, req.(*RemoveClassifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisionService_GetClassificationsFromCamera_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassificationsFromCameraRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).GetClassificationsFromCamera(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.vision.v1.VisionService/GetClassificationsFromCamera",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).GetClassificationsFromCamera(ctx, req.(*GetClassificationsFromCameraRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisionService_GetClassifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).GetClassifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.vision.v1.VisionService/GetClassifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).GetClassifications(ctx, req.(*GetClassificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisionService_GetSegmenterNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSegmenterNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).GetSegmenterNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.vision.v1.VisionService/GetSegmenterNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).GetSegmenterNames(ctx, req.(*GetSegmenterNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisionService_AddSegmenter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSegmenterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).AddSegmenter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.vision.v1.VisionService/AddSegmenter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).AddSegmenter(ctx, req.(*AddSegmenterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisionService_RemoveSegmenter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSegmenterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).RemoveSegmenter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.vision.v1.VisionService/RemoveSegmenter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).RemoveSegmenter(ctx, req.(*RemoveSegmenterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisionService_GetObjectPointClouds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectPointCloudsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).GetObjectPointClouds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.vision.v1.VisionService/GetObjectPointClouds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).GetObjectPointClouds(ctx, req.(*GetObjectPointCloudsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VisionService_ServiceDesc is the grpc.ServiceDesc for VisionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VisionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "viam.service.vision.v1.VisionService",
	HandlerType: (*VisionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetModelParameterSchema",
			Handler:    _VisionService_GetModelParameterSchema_Handler,
		},
		{
			MethodName: "GetDetectorNames",
			Handler:    _VisionService_GetDetectorNames_Handler,
		},
		{
			MethodName: "AddDetector",
			Handler:    _VisionService_AddDetector_Handler,
		},
		{
			MethodName: "RemoveDetector",
			Handler:    _VisionService_RemoveDetector_Handler,
		},
		{
			MethodName: "GetDetectionsFromCamera",
			Handler:    _VisionService_GetDetectionsFromCamera_Handler,
		},
		{
			MethodName: "GetDetections",
			Handler:    _VisionService_GetDetections_Handler,
		},
		{
			MethodName: "GetClassifierNames",
			Handler:    _VisionService_GetClassifierNames_Handler,
		},
		{
			MethodName: "AddClassifier",
			Handler:    _VisionService_AddClassifier_Handler,
		},
		{
			MethodName: "RemoveClassifier",
			Handler:    _VisionService_RemoveClassifier_Handler,
		},
		{
			MethodName: "GetClassificationsFromCamera",
			Handler:    _VisionService_GetClassificationsFromCamera_Handler,
		},
		{
			MethodName: "GetClassifications",
			Handler:    _VisionService_GetClassifications_Handler,
		},
		{
			MethodName: "GetSegmenterNames",
			Handler:    _VisionService_GetSegmenterNames_Handler,
		},
		{
			MethodName: "AddSegmenter",
			Handler:    _VisionService_AddSegmenter_Handler,
		},
		{
			MethodName: "RemoveSegmenter",
			Handler:    _VisionService_RemoveSegmenter_Handler,
		},
		{
			MethodName: "GetObjectPointClouds",
			Handler:    _VisionService_GetObjectPointClouds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/vision/v1/vision.proto",
}