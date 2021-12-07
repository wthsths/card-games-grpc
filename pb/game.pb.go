// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: game.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GameType int32

const (
	GameType_GAME_TYPE_TEXAS_HOLDEM_BONUS GameType = 0
)

// Enum value maps for GameType.
var (
	GameType_name = map[int32]string{
		0: "GAME_TYPE_TEXAS_HOLDEM_BONUS",
	}
	GameType_value = map[string]int32{
		"GAME_TYPE_TEXAS_HOLDEM_BONUS": 0,
	}
)

func (x GameType) Enum() *GameType {
	p := new(GameType)
	*p = x
	return p
}

func (x GameType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameType) Descriptor() protoreflect.EnumDescriptor {
	return file_game_proto_enumTypes[0].Descriptor()
}

func (GameType) Type() protoreflect.EnumType {
	return &file_game_proto_enumTypes[0]
}

func (x GameType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameType.Descriptor instead.
func (GameType) EnumDescriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
}

type GameStatus int32

const (
	GameStatus_GAME_STATUS_CREATED   GameStatus = 0
	GameStatus_GAME_STATUS_PLAYING   GameStatus = 1
	GameStatus_GAME_STATUS_COMPLETED GameStatus = 2
)

// Enum value maps for GameStatus.
var (
	GameStatus_name = map[int32]string{
		0: "GAME_STATUS_CREATED",
		1: "GAME_STATUS_PLAYING",
		2: "GAME_STATUS_COMPLETED",
	}
	GameStatus_value = map[string]int32{
		"GAME_STATUS_CREATED":   0,
		"GAME_STATUS_PLAYING":   1,
		"GAME_STATUS_COMPLETED": 2,
	}
)

func (x GameStatus) Enum() *GameStatus {
	p := new(GameStatus)
	*p = x
	return p
}

func (x GameStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_game_proto_enumTypes[1].Descriptor()
}

func (GameStatus) Type() protoreflect.EnumType {
	return &file_game_proto_enumTypes[1]
}

func (x GameStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameStatus.Descriptor instead.
func (GameStatus) EnumDescriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{1}
}

type PlayGameRequest_Action int32

const (
	PlayGameRequest_ACTION_DEAL  PlayGameRequest_Action = 0
	PlayGameRequest_ACTION_PLAY  PlayGameRequest_Action = 1
	PlayGameRequest_ACTION_FOLD  PlayGameRequest_Action = 2
	PlayGameRequest_ACTION_BET   PlayGameRequest_Action = 3
	PlayGameRequest_ACTION_CHECK PlayGameRequest_Action = 4
)

// Enum value maps for PlayGameRequest_Action.
var (
	PlayGameRequest_Action_name = map[int32]string{
		0: "ACTION_DEAL",
		1: "ACTION_PLAY",
		2: "ACTION_FOLD",
		3: "ACTION_BET",
		4: "ACTION_CHECK",
	}
	PlayGameRequest_Action_value = map[string]int32{
		"ACTION_DEAL":  0,
		"ACTION_PLAY":  1,
		"ACTION_FOLD":  2,
		"ACTION_BET":   3,
		"ACTION_CHECK": 4,
	}
)

func (x PlayGameRequest_Action) Enum() *PlayGameRequest_Action {
	p := new(PlayGameRequest_Action)
	*p = x
	return p
}

func (x PlayGameRequest_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayGameRequest_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_game_proto_enumTypes[2].Descriptor()
}

func (PlayGameRequest_Action) Type() protoreflect.EnumType {
	return &file_game_proto_enumTypes[2]
}

func (x PlayGameRequest_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayGameRequest_Action.Descriptor instead.
func (PlayGameRequest_Action) EnumDescriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{5, 0}
}

type Game struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   string     `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Player uint64     `protobuf:"varint,2,opt,name=player,proto3" json:"player,omitempty"`
	Type   GameType   `protobuf:"varint,3,opt,name=type,proto3,enum=GameType" json:"type,omitempty"`
	Status GameStatus `protobuf:"varint,4,opt,name=status,proto3,enum=GameStatus" json:"status,omitempty"`
	Data   []byte     `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Game) Reset() {
	*x = Game{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
}

func (x *Game) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Game) GetPlayer() uint64 {
	if x != nil {
		return x.Player
	}
	return 0
}

func (x *Game) GetType() GameType {
	if x != nil {
		return x.Type
	}
	return GameType_GAME_TYPE_TEXAS_HOLDEM_BONUS
}

func (x *Game) GetStatus() GameStatus {
	if x != nil {
		return x.Status
	}
	return GameStatus_GAME_STATUS_CREATED
}

func (x *Game) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player uint64   `protobuf:"varint,1,opt,name=player,proto3" json:"player,omitempty"`
	Type   GameType `protobuf:"varint,2,opt,name=type,proto3,enum=GameType" json:"type,omitempty"`
}

func (x *CreateGameRequest) Reset() {
	*x = CreateGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGameRequest) ProtoMessage() {}

func (x *CreateGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGameRequest.ProtoReflect.Descriptor instead.
func (*CreateGameRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGameRequest) GetPlayer() uint64 {
	if x != nil {
		return x.Player
	}
	return 0
}

func (x *CreateGameRequest) GetType() GameType {
	if x != nil {
		return x.Type
	}
	return GameType_GAME_TYPE_TEXAS_HOLDEM_BONUS
}

type CreateGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Game *Game `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
}

func (x *CreateGameResponse) Reset() {
	*x = CreateGameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGameResponse) ProtoMessage() {}

func (x *CreateGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGameResponse.ProtoReflect.Descriptor instead.
func (*CreateGameResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{2}
}

func (x *CreateGameResponse) GetGame() *Game {
	if x != nil {
		return x.Game
	}
	return nil
}

type GetGameByUuidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetGameByUuidRequest) Reset() {
	*x = GetGameByUuidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGameByUuidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameByUuidRequest) ProtoMessage() {}

func (x *GetGameByUuidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameByUuidRequest.ProtoReflect.Descriptor instead.
func (*GetGameByUuidRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{3}
}

func (x *GetGameByUuidRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetGameByUuidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Game *Game `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
}

func (x *GetGameByUuidResponse) Reset() {
	*x = GetGameByUuidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGameByUuidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameByUuidResponse) ProtoMessage() {}

func (x *GetGameByUuidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameByUuidResponse.ProtoReflect.Descriptor instead.
func (*GetGameByUuidResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{4}
}

func (x *GetGameByUuidResponse) GetGame() *Game {
	if x != nil {
		return x.Game
	}
	return nil
}

type PlayGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action PlayGameRequest_Action   `protobuf:"varint,1,opt,name=action,proto3,enum=PlayGameRequest_Action" json:"action,omitempty"`
	Uuid   string                   `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Amount uint64                   `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Bonus  []*PlayGameRequest_Bonus `protobuf:"bytes,4,rep,name=bonus,proto3" json:"bonus,omitempty"`
}

func (x *PlayGameRequest) Reset() {
	*x = PlayGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayGameRequest) ProtoMessage() {}

func (x *PlayGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayGameRequest.ProtoReflect.Descriptor instead.
func (*PlayGameRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{5}
}

func (x *PlayGameRequest) GetAction() PlayGameRequest_Action {
	if x != nil {
		return x.Action
	}
	return PlayGameRequest_ACTION_DEAL
}

func (x *PlayGameRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *PlayGameRequest) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PlayGameRequest) GetBonus() []*PlayGameRequest_Bonus {
	if x != nil {
		return x.Bonus
	}
	return nil
}

type PlayGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Game *Game `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
}

func (x *PlayGameResponse) Reset() {
	*x = PlayGameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayGameResponse) ProtoMessage() {}

func (x *PlayGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayGameResponse.ProtoReflect.Descriptor instead.
func (*PlayGameResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{6}
}

func (x *PlayGameResponse) GetGame() *Game {
	if x != nil {
		return x.Game
	}
	return nil
}

type GetGamesByPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player uint64 `protobuf:"varint,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *GetGamesByPlayerRequest) Reset() {
	*x = GetGamesByPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGamesByPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGamesByPlayerRequest) ProtoMessage() {}

func (x *GetGamesByPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGamesByPlayerRequest.ProtoReflect.Descriptor instead.
func (*GetGamesByPlayerRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{7}
}

func (x *GetGamesByPlayerRequest) GetPlayer() uint64 {
	if x != nil {
		return x.Player
	}
	return 0
}

type GetGamesByPlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Game []*Game `protobuf:"bytes,1,rep,name=game,proto3" json:"game,omitempty"`
}

func (x *GetGamesByPlayerResponse) Reset() {
	*x = GetGamesByPlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGamesByPlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGamesByPlayerResponse) ProtoMessage() {}

func (x *GetGamesByPlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGamesByPlayerResponse.ProtoReflect.Descriptor instead.
func (*GetGamesByPlayerResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{8}
}

func (x *GetGamesByPlayerResponse) GetGame() []*Game {
	if x != nil {
		return x.Game
	}
	return nil
}

type PlayGameRequest_Bonus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index  int32  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Amount uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *PlayGameRequest_Bonus) Reset() {
	*x = PlayGameRequest_Bonus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayGameRequest_Bonus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayGameRequest_Bonus) ProtoMessage() {}

func (x *PlayGameRequest_Bonus) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayGameRequest_Bonus.ProtoReflect.Descriptor instead.
func (*PlayGameRequest_Bonus) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{5, 0}
}

func (x *PlayGameRequest_Bonus) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *PlayGameRequest_Bonus) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_game_proto protoreflect.FileDescriptor

var file_game_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a,
	0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x09, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x23, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0b, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4a, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x67,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d,
	0x65, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x22, 0x32, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x79, 0x55,
	0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x67,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x22, 0xb2, 0x02, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x62, 0x6f, 0x6e, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x52, 0x05,
	0x62, 0x6f, 0x6e, 0x75, 0x73, 0x1a, 0x35, 0x0a, 0x05, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5d, 0x0a, 0x06,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x44, 0x45, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x46, 0x4f, 0x4c, 0x44, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x45, 0x54, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x10, 0x04, 0x22, 0x2d, 0x0a, 0x10, 0x50,
	0x6c, 0x61, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x19, 0x0a, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x35, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x67, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x04,
	0x67, 0x61, 0x6d, 0x65, 0x2a, 0x2c, 0x0a, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x20, 0x0a, 0x1c, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x45,
	0x58, 0x41, 0x53, 0x5f, 0x48, 0x4f, 0x4c, 0x44, 0x45, 0x4d, 0x5f, 0x42, 0x4f, 0x4e, 0x55, 0x53,
	0x10, 0x00, 0x2a, 0x59, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x17, 0x0a, 0x13, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x41, 0x4d,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x32, 0xfe, 0x01,
	0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x42,
	0x79, 0x55, 0x75, 0x69, 0x64, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x42,
	0x79, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x47,
	0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x47, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65,
	0x73, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x47, 0x65, 0x74, 0x47,
	0x61, 0x6d, 0x65, 0x73, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x42, 0x79,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_proto_rawDescOnce sync.Once
	file_game_proto_rawDescData = file_game_proto_rawDesc
)

func file_game_proto_rawDescGZIP() []byte {
	file_game_proto_rawDescOnce.Do(func() {
		file_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_proto_rawDescData)
	})
	return file_game_proto_rawDescData
}

var file_game_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_game_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_game_proto_goTypes = []interface{}{
	(GameType)(0),                    // 0: GameType
	(GameStatus)(0),                  // 1: GameStatus
	(PlayGameRequest_Action)(0),      // 2: PlayGameRequest.Action
	(*Game)(nil),                     // 3: Game
	(*CreateGameRequest)(nil),        // 4: CreateGameRequest
	(*CreateGameResponse)(nil),       // 5: CreateGameResponse
	(*GetGameByUuidRequest)(nil),     // 6: GetGameByUuidRequest
	(*GetGameByUuidResponse)(nil),    // 7: GetGameByUuidResponse
	(*PlayGameRequest)(nil),          // 8: PlayGameRequest
	(*PlayGameResponse)(nil),         // 9: PlayGameResponse
	(*GetGamesByPlayerRequest)(nil),  // 10: GetGamesByPlayerRequest
	(*GetGamesByPlayerResponse)(nil), // 11: GetGamesByPlayerResponse
	(*PlayGameRequest_Bonus)(nil),    // 12: PlayGameRequest.Bonus
}
var file_game_proto_depIdxs = []int32{
	0,  // 0: Game.type:type_name -> GameType
	1,  // 1: Game.status:type_name -> GameStatus
	0,  // 2: CreateGameRequest.type:type_name -> GameType
	3,  // 3: CreateGameResponse.game:type_name -> Game
	3,  // 4: GetGameByUuidResponse.game:type_name -> Game
	2,  // 5: PlayGameRequest.action:type_name -> PlayGameRequest.Action
	12, // 6: PlayGameRequest.bonus:type_name -> PlayGameRequest.Bonus
	3,  // 7: PlayGameResponse.game:type_name -> Game
	3,  // 8: GetGamesByPlayerResponse.game:type_name -> Game
	4,  // 9: GameService.CreateGame:input_type -> CreateGameRequest
	6,  // 10: GameService.GetGameByUuid:input_type -> GetGameByUuidRequest
	8,  // 11: GameService.PlayGame:input_type -> PlayGameRequest
	10, // 12: GameService.GetGamesByPlayer:input_type -> GetGamesByPlayerRequest
	5,  // 13: GameService.CreateGame:output_type -> CreateGameResponse
	7,  // 14: GameService.GetGameByUuid:output_type -> GetGameByUuidResponse
	9,  // 15: GameService.PlayGame:output_type -> PlayGameResponse
	11, // 16: GameService.GetGamesByPlayer:output_type -> GetGamesByPlayerResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_game_proto_init() }
func file_game_proto_init() {
	if File_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Game); i {
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
		file_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGameRequest); i {
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
		file_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGameResponse); i {
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
		file_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGameByUuidRequest); i {
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
		file_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGameByUuidResponse); i {
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
		file_game_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayGameRequest); i {
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
		file_game_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayGameResponse); i {
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
		file_game_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGamesByPlayerRequest); i {
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
		file_game_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGamesByPlayerResponse); i {
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
		file_game_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayGameRequest_Bonus); i {
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
			RawDescriptor: file_game_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_proto_goTypes,
		DependencyIndexes: file_game_proto_depIdxs,
		EnumInfos:         file_game_proto_enumTypes,
		MessageInfos:      file_game_proto_msgTypes,
	}.Build()
	File_game_proto = out.File
	file_game_proto_rawDesc = nil
	file_game_proto_goTypes = nil
	file_game_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameServiceClient interface {
	CreateGame(ctx context.Context, in *CreateGameRequest, opts ...grpc.CallOption) (*CreateGameResponse, error)
	GetGameByUuid(ctx context.Context, in *GetGameByUuidRequest, opts ...grpc.CallOption) (*GetGameByUuidResponse, error)
	PlayGame(ctx context.Context, in *PlayGameRequest, opts ...grpc.CallOption) (*PlayGameResponse, error)
	GetGamesByPlayer(ctx context.Context, in *GetGamesByPlayerRequest, opts ...grpc.CallOption) (*GetGamesByPlayerResponse, error)
}

type gameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServiceClient(cc grpc.ClientConnInterface) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) CreateGame(ctx context.Context, in *CreateGameRequest, opts ...grpc.CallOption) (*CreateGameResponse, error) {
	out := new(CreateGameResponse)
	err := c.cc.Invoke(ctx, "/GameService/CreateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) GetGameByUuid(ctx context.Context, in *GetGameByUuidRequest, opts ...grpc.CallOption) (*GetGameByUuidResponse, error) {
	out := new(GetGameByUuidResponse)
	err := c.cc.Invoke(ctx, "/GameService/GetGameByUuid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) PlayGame(ctx context.Context, in *PlayGameRequest, opts ...grpc.CallOption) (*PlayGameResponse, error) {
	out := new(PlayGameResponse)
	err := c.cc.Invoke(ctx, "/GameService/PlayGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) GetGamesByPlayer(ctx context.Context, in *GetGamesByPlayerRequest, opts ...grpc.CallOption) (*GetGamesByPlayerResponse, error) {
	out := new(GetGamesByPlayerResponse)
	err := c.cc.Invoke(ctx, "/GameService/GetGamesByPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceServer is the server API for GameService service.
type GameServiceServer interface {
	CreateGame(context.Context, *CreateGameRequest) (*CreateGameResponse, error)
	GetGameByUuid(context.Context, *GetGameByUuidRequest) (*GetGameByUuidResponse, error)
	PlayGame(context.Context, *PlayGameRequest) (*PlayGameResponse, error)
	GetGamesByPlayer(context.Context, *GetGamesByPlayerRequest) (*GetGamesByPlayerResponse, error)
}

// UnimplementedGameServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (*UnimplementedGameServiceServer) CreateGame(context.Context, *CreateGameRequest) (*CreateGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGame not implemented")
}
func (*UnimplementedGameServiceServer) GetGameByUuid(context.Context, *GetGameByUuidRequest) (*GetGameByUuidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameByUuid not implemented")
}
func (*UnimplementedGameServiceServer) PlayGame(context.Context, *PlayGameRequest) (*PlayGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayGame not implemented")
}
func (*UnimplementedGameServiceServer) GetGamesByPlayer(context.Context, *GetGamesByPlayerRequest) (*GetGamesByPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGamesByPlayer not implemented")
}

func RegisterGameServiceServer(s *grpc.Server, srv GameServiceServer) {
	s.RegisterService(&_GameService_serviceDesc, srv)
}

func _GameService_CreateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).CreateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GameService/CreateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).CreateGame(ctx, req.(*CreateGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_GetGameByUuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameByUuidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetGameByUuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GameService/GetGameByUuid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetGameByUuid(ctx, req.(*GetGameByUuidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_PlayGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).PlayGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GameService/PlayGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).PlayGame(ctx, req.(*PlayGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_GetGamesByPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGamesByPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetGamesByPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GameService/GetGamesByPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetGamesByPlayer(ctx, req.(*GetGamesByPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGame",
			Handler:    _GameService_CreateGame_Handler,
		},
		{
			MethodName: "GetGameByUuid",
			Handler:    _GameService_GetGameByUuid_Handler,
		},
		{
			MethodName: "PlayGame",
			Handler:    _GameService_PlayGame_Handler,
		},
		{
			MethodName: "GetGamesByPlayer",
			Handler:    _GameService_GetGamesByPlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game.proto",
}
