// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: publish.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResp) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *BaseResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type VideoBaseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PlayUrl  string `protobuf:"bytes,2,opt,name=play_url,json=playUrl,proto3" json:"play_url,omitempty"`
	CoverUrl string `protobuf:"bytes,3,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`
	Title    string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *VideoBaseInfo) Reset() {
	*x = VideoBaseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoBaseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoBaseInfo) ProtoMessage() {}

func (x *VideoBaseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoBaseInfo.ProtoReflect.Descriptor instead.
func (*VideoBaseInfo) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{1}
}

func (x *VideoBaseInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *VideoBaseInfo) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *VideoBaseInfo) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *VideoBaseInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type VideoInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId       int64          `protobuf:"varint,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	VideoBaseInfo *VideoBaseInfo `protobuf:"bytes,2,opt,name=video_base_info,json=videoBaseInfo,proto3" json:"video_base_info,omitempty"`
	LikeCount     int64          `protobuf:"varint,3,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
	CommentCount  int64          `protobuf:"varint,4,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`
	IsFavorite    bool           `protobuf:"varint,5,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite,omitempty"`
}

func (x *VideoInfo) Reset() {
	*x = VideoInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoInfo) ProtoMessage() {}

func (x *VideoInfo) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoInfo.ProtoReflect.Descriptor instead.
func (*VideoInfo) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{2}
}

func (x *VideoInfo) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *VideoInfo) GetVideoBaseInfo() *VideoBaseInfo {
	if x != nil {
		return x.VideoBaseInfo
	}
	return nil
}

func (x *VideoInfo) GetLikeCount() int64 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

func (x *VideoInfo) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *VideoInfo) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

type PublishVideoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoBaseInfo *VideoBaseInfo `protobuf:"bytes,1,opt,name=video_base_info,json=videoBaseInfo,proto3" json:"video_base_info,omitempty"`
}

func (x *PublishVideoReq) Reset() {
	*x = PublishVideoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishVideoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishVideoReq) ProtoMessage() {}

func (x *PublishVideoReq) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishVideoReq.ProtoReflect.Descriptor instead.
func (*PublishVideoReq) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{3}
}

func (x *PublishVideoReq) GetVideoBaseInfo() *VideoBaseInfo {
	if x != nil {
		return x.VideoBaseInfo
	}
	return nil
}

type PublishVideoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *PublishVideoResp) Reset() {
	*x = PublishVideoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishVideoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishVideoResp) ProtoMessage() {}

func (x *PublishVideoResp) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishVideoResp.ProtoReflect.Descriptor instead.
func (*PublishVideoResp) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{4}
}

func (x *PublishVideoResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type GetVideoListByUserIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppUserId int64 `protobuf:"varint,1,opt,name=app_user_id,json=appUserId,proto3" json:"app_user_id,omitempty"`
	UserId    int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetVideoListByUserIdReq) Reset() {
	*x = GetVideoListByUserIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoListByUserIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoListByUserIdReq) ProtoMessage() {}

func (x *GetVideoListByUserIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoListByUserIdReq.ProtoReflect.Descriptor instead.
func (*GetVideoListByUserIdReq) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{5}
}

func (x *GetVideoListByUserIdReq) GetAppUserId() int64 {
	if x != nil {
		return x.AppUserId
	}
	return 0
}

func (x *GetVideoListByUserIdReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetVideoListByUserIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp   *BaseResp    `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	VideoInfos []*VideoInfo `protobuf:"bytes,2,rep,name=video_infos,json=videoInfos,proto3" json:"video_infos,omitempty"`
}

func (x *GetVideoListByUserIdResp) Reset() {
	*x = GetVideoListByUserIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoListByUserIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoListByUserIdResp) ProtoMessage() {}

func (x *GetVideoListByUserIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoListByUserIdResp.ProtoReflect.Descriptor instead.
func (*GetVideoListByUserIdResp) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{6}
}

func (x *GetVideoListByUserIdResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *GetVideoListByUserIdResp) GetVideoInfos() []*VideoInfo {
	if x != nil {
		return x.VideoInfos
	}
	return nil
}

var File_publish_proto protoreflect.FileDescriptor

var file_publish_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x69, 0x64, 0x6c, 0x22, 0x4a, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67,
	0x22, 0x76, 0x0a, 0x0d, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c,
	0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55,
	0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0xc7, 0x01, 0x0a, 0x09, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49,
	0x64, 0x12, 0x3a, 0x0a, 0x0f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x64, 0x6c,
	0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a,
	0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x22, 0x4d, 0x0a, 0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x3a, 0x0a, 0x0f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0d, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x3e, 0x0a, 0x10, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x52, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0b,
	0x61, 0x70, 0x70, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x61, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x77, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x2a, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a,
	0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x32, 0xa2,
	0x01, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x12, 0x14, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x53,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x47, 0x65, 0x74,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_publish_proto_rawDescOnce sync.Once
	file_publish_proto_rawDescData = file_publish_proto_rawDesc
)

func file_publish_proto_rawDescGZIP() []byte {
	file_publish_proto_rawDescOnce.Do(func() {
		file_publish_proto_rawDescData = protoimpl.X.CompressGZIP(file_publish_proto_rawDescData)
	})
	return file_publish_proto_rawDescData
}

var file_publish_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_publish_proto_goTypes = []interface{}{
	(*BaseResp)(nil),                 // 0: idl.BaseResp
	(*VideoBaseInfo)(nil),            // 1: idl.VideoBaseInfo
	(*VideoInfo)(nil),                // 2: idl.VideoInfo
	(*PublishVideoReq)(nil),          // 3: idl.PublishVideoReq
	(*PublishVideoResp)(nil),         // 4: idl.PublishVideoResp
	(*GetVideoListByUserIdReq)(nil),  // 5: idl.GetVideoListByUserIdReq
	(*GetVideoListByUserIdResp)(nil), // 6: idl.GetVideoListByUserIdResp
}
var file_publish_proto_depIdxs = []int32{
	1, // 0: idl.VideoInfo.video_base_info:type_name -> idl.VideoBaseInfo
	1, // 1: idl.PublishVideoReq.video_base_info:type_name -> idl.VideoBaseInfo
	0, // 2: idl.PublishVideoResp.base_resp:type_name -> idl.BaseResp
	0, // 3: idl.GetVideoListByUserIdResp.base_resp:type_name -> idl.BaseResp
	2, // 4: idl.GetVideoListByUserIdResp.video_infos:type_name -> idl.VideoInfo
	3, // 5: idl.PublishService.PublishVideo:input_type -> idl.PublishVideoReq
	5, // 6: idl.PublishService.GetVideoListByUserId:input_type -> idl.GetVideoListByUserIdReq
	4, // 7: idl.PublishService.PublishVideo:output_type -> idl.PublishVideoResp
	6, // 8: idl.PublishService.GetVideoListByUserId:output_type -> idl.GetVideoListByUserIdResp
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_publish_proto_init() }
func file_publish_proto_init() {
	if File_publish_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_publish_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_publish_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoBaseInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_publish_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_publish_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishVideoReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_publish_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishVideoResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_publish_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoListByUserIdReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_publish_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoListByUserIdResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_publish_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_publish_proto_goTypes,
		DependencyIndexes: file_publish_proto_depIdxs,
		MessageInfos:      file_publish_proto_msgTypes,
	}.Build()
	File_publish_proto = out.File
	file_publish_proto_rawDesc = nil
	file_publish_proto_goTypes = nil
	file_publish_proto_depIdxs = nil
}
