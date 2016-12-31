// Code generated by protoc-gen-go.
// source: types.proto
// DO NOT EDIT!

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	types.proto

It has these top-level messages:
	CommonAck
	BlankMsg
	HeadBeat
	LoginReq
	LoginAck
	LogoutReq
	LogoutAck
	SetPlayerNameReq
	SetPlayerNameAck
	ChoseSideReq
	ChoseSideAck
	ScreenActive
	ActiveDetail
	ScreenDecoration
	ScreenInfo
	ScreenChangeNotify
	LoginRoomReq
	LoginRoomAck
	LeaveRoomReq
	LeaveRoomAck
	MoveActionReq
	MoveActionAck
	ChoseHeroReq
	ChoseHeroAck
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CommonAck struct {
	Status    int32  `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	TimeStamp uint32 `protobuf:"varint,2,opt,name=timeStamp" json:"timeStamp,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *CommonAck) Reset()                    { *m = CommonAck{} }
func (m *CommonAck) String() string            { return proto.CompactTextString(m) }
func (*CommonAck) ProtoMessage()               {}
func (*CommonAck) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{0} }

func (m *CommonAck) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CommonAck) GetTimeStamp() uint32 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *CommonAck) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type BlankMsg struct {
	Helper bool `protobuf:"varint,1,opt,name=helper" json:"helper,omitempty"`
}

func (m *BlankMsg) Reset()                    { *m = BlankMsg{} }
func (m *BlankMsg) String() string            { return proto.CompactTextString(m) }
func (*BlankMsg) ProtoMessage()               {}
func (*BlankMsg) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{1} }

func (m *BlankMsg) GetHelper() bool {
	if m != nil {
		return m.Helper
	}
	return false
}

type HeadBeat struct {
	TimeStamp int32 `protobuf:"varint,1,opt,name=time_stamp,json=timeStamp" json:"time_stamp,omitempty"`
	Code      int32 `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
}

func (m *HeadBeat) Reset()                    { *m = HeadBeat{} }
func (m *HeadBeat) String() string            { return proto.CompactTextString(m) }
func (*HeadBeat) ProtoMessage()               {}
func (*HeadBeat) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{2} }

func (m *HeadBeat) GetTimeStamp() int32 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *HeadBeat) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type LoginReq struct {
	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Uuid     string `protobuf:"bytes,2,opt,name=uuid" json:"uuid,omitempty"`
	ServerId string `protobuf:"bytes,3,opt,name=server_id,json=serverId" json:"server_id,omitempty"`
}

func (m *LoginReq) Reset()                    { *m = LoginReq{} }
func (m *LoginReq) String() string            { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()               {}
func (*LoginReq) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{3} }

func (m *LoginReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *LoginReq) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *LoginReq) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

type LoginAck struct {
	Common     *CommonAck `protobuf:"bytes,1,opt,name=common" json:"common,omitempty"`
	UserId     string     `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserName   string     `protobuf:"bytes,3,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	HeroName   string     `protobuf:"bytes,4,opt,name=hero_name,json=heroName" json:"hero_name,omitempty"`
	HeroSkin   string     `protobuf:"bytes,5,opt,name=hero_skin,json=heroSkin" json:"hero_skin,omitempty"`
	HeroSkills []string   `protobuf:"bytes,6,rep,name=hero_skills,json=heroSkills" json:"hero_skills,omitempty"`
}

func (m *LoginAck) Reset()                    { *m = LoginAck{} }
func (m *LoginAck) String() string            { return proto.CompactTextString(m) }
func (*LoginAck) ProtoMessage()               {}
func (*LoginAck) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{4} }

func (m *LoginAck) GetCommon() *CommonAck {
	if m != nil {
		return m.Common
	}
	return nil
}

func (m *LoginAck) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *LoginAck) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LoginAck) GetHeroName() string {
	if m != nil {
		return m.HeroName
	}
	return ""
}

func (m *LoginAck) GetHeroSkin() string {
	if m != nil {
		return m.HeroSkin
	}
	return ""
}

func (m *LoginAck) GetHeroSkills() []string {
	if m != nil {
		return m.HeroSkills
	}
	return nil
}

type LogoutReq struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *LogoutReq) Reset()                    { *m = LogoutReq{} }
func (m *LogoutReq) String() string            { return proto.CompactTextString(m) }
func (*LogoutReq) ProtoMessage()               {}
func (*LogoutReq) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{5} }

func (m *LogoutReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type LogoutAck struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *LogoutAck) Reset()                    { *m = LogoutAck{} }
func (m *LogoutAck) String() string            { return proto.CompactTextString(m) }
func (*LogoutAck) ProtoMessage()               {}
func (*LogoutAck) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{6} }

func (m *LogoutAck) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type SetPlayerNameReq struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *SetPlayerNameReq) Reset()                    { *m = SetPlayerNameReq{} }
func (m *SetPlayerNameReq) String() string            { return proto.CompactTextString(m) }
func (*SetPlayerNameReq) ProtoMessage()               {}
func (*SetPlayerNameReq) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{7} }

func (m *SetPlayerNameReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SetPlayerNameAck struct {
	Common *CommonAck `protobuf:"bytes,1,opt,name=common" json:"common,omitempty"`
	Name   string     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *SetPlayerNameAck) Reset()                    { *m = SetPlayerNameAck{} }
func (m *SetPlayerNameAck) String() string            { return proto.CompactTextString(m) }
func (*SetPlayerNameAck) ProtoMessage()               {}
func (*SetPlayerNameAck) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{8} }

func (m *SetPlayerNameAck) GetCommon() *CommonAck {
	if m != nil {
		return m.Common
	}
	return nil
}

func (m *SetPlayerNameAck) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ChoseSideReq struct {
	Side string `protobuf:"bytes,1,opt,name=Side" json:"Side,omitempty"`
}

func (m *ChoseSideReq) Reset()                    { *m = ChoseSideReq{} }
func (m *ChoseSideReq) String() string            { return proto.CompactTextString(m) }
func (*ChoseSideReq) ProtoMessage()               {}
func (*ChoseSideReq) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{9} }

func (m *ChoseSideReq) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

type ChoseSideAck struct {
	Common *CommonAck `protobuf:"bytes,1,opt,name=common" json:"common,omitempty"`
}

func (m *ChoseSideAck) Reset()                    { *m = ChoseSideAck{} }
func (m *ChoseSideAck) String() string            { return proto.CompactTextString(m) }
func (*ChoseSideAck) ProtoMessage()               {}
func (*ChoseSideAck) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{10} }

func (m *ChoseSideAck) GetCommon() *CommonAck {
	if m != nil {
		return m.Common
	}
	return nil
}

type ScreenActive struct {
	Id      string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	DirectX int32  `protobuf:"varint,2,opt,name=DirectX" json:"DirectX,omitempty"`
	DirectY int32  `protobuf:"varint,3,opt,name=DirectY" json:"DirectY,omitempty"`
	Speed   int32  `protobuf:"varint,4,opt,name=Speed" json:"Speed,omitempty"`
	PosX    int32  `protobuf:"varint,5,opt,name=PosX" json:"PosX,omitempty"`
	PosY    int32  `protobuf:"varint,6,opt,name=PosY" json:"PosY,omitempty"`
	Type    int32  `protobuf:"varint,7,opt,name=Type" json:"Type,omitempty"`
	Side    string `protobuf:"bytes,8,opt,name=Side" json:"Side,omitempty"`
}

func (m *ScreenActive) Reset()                    { *m = ScreenActive{} }
func (m *ScreenActive) String() string            { return proto.CompactTextString(m) }
func (*ScreenActive) ProtoMessage()               {}
func (*ScreenActive) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{11} }

func (m *ScreenActive) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ScreenActive) GetDirectX() int32 {
	if m != nil {
		return m.DirectX
	}
	return 0
}

func (m *ScreenActive) GetDirectY() int32 {
	if m != nil {
		return m.DirectY
	}
	return 0
}

func (m *ScreenActive) GetSpeed() int32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func (m *ScreenActive) GetPosX() int32 {
	if m != nil {
		return m.PosX
	}
	return 0
}

func (m *ScreenActive) GetPosY() int32 {
	if m != nil {
		return m.PosY
	}
	return 0
}

func (m *ScreenActive) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ScreenActive) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

type ActiveDetail struct {
	Id        string   `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	BelongsTo string   `protobuf:"bytes,2,opt,name=BelongsTo" json:"BelongsTo,omitempty"`
	Name      string   `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	SubType   string   `protobuf:"bytes,4,opt,name=SubType" json:"SubType,omitempty"`
	Skin      string   `protobuf:"bytes,5,opt,name=Skin" json:"Skin,omitempty"`
	Hp        int32    `protobuf:"varint,6,opt,name=Hp" json:"Hp,omitempty"`
	FullHp    int32    `protobuf:"varint,7,opt,name=FullHp" json:"FullHp,omitempty"`
	Skills    []string `protobuf:"bytes,8,rep,name=skills" json:"skills,omitempty"`
}

func (m *ActiveDetail) Reset()                    { *m = ActiveDetail{} }
func (m *ActiveDetail) String() string            { return proto.CompactTextString(m) }
func (*ActiveDetail) ProtoMessage()               {}
func (*ActiveDetail) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{12} }

func (m *ActiveDetail) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ActiveDetail) GetBelongsTo() string {
	if m != nil {
		return m.BelongsTo
	}
	return ""
}

func (m *ActiveDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ActiveDetail) GetSubType() string {
	if m != nil {
		return m.SubType
	}
	return ""
}

func (m *ActiveDetail) GetSkin() string {
	if m != nil {
		return m.Skin
	}
	return ""
}

func (m *ActiveDetail) GetHp() int32 {
	if m != nil {
		return m.Hp
	}
	return 0
}

func (m *ActiveDetail) GetFullHp() int32 {
	if m != nil {
		return m.FullHp
	}
	return 0
}

func (m *ActiveDetail) GetSkills() []string {
	if m != nil {
		return m.Skills
	}
	return nil
}

type ScreenDecoration struct {
	AssetId int32 `protobuf:"varint,1,opt,name=AssetId" json:"AssetId,omitempty"`
	X       int32 `protobuf:"varint,2,opt,name=X" json:"X,omitempty"`
	Y       int32 `protobuf:"varint,3,opt,name=Y" json:"Y,omitempty"`
}

func (m *ScreenDecoration) Reset()                    { *m = ScreenDecoration{} }
func (m *ScreenDecoration) String() string            { return proto.CompactTextString(m) }
func (*ScreenDecoration) ProtoMessage()               {}
func (*ScreenDecoration) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{13} }

func (m *ScreenDecoration) GetAssetId() int32 {
	if m != nil {
		return m.AssetId
	}
	return 0
}

func (m *ScreenDecoration) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *ScreenDecoration) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type ScreenInfo struct {
	FrameId     int64               `protobuf:"varint,1,opt,name=frame_id,json=frameId" json:"frame_id,omitempty"`
	Width       int32               `protobuf:"varint,2,opt,name=width" json:"width,omitempty"`
	Height      int32               `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	Actives     []*ScreenActive     `protobuf:"bytes,4,rep,name=actives" json:"actives,omitempty"`
	Decorations []*ScreenDecoration `protobuf:"bytes,5,rep,name=decorations" json:"decorations,omitempty"`
	Details     []*ActiveDetail     `protobuf:"bytes,6,rep,name=details" json:"details,omitempty"`
}

func (m *ScreenInfo) Reset()                    { *m = ScreenInfo{} }
func (m *ScreenInfo) String() string            { return proto.CompactTextString(m) }
func (*ScreenInfo) ProtoMessage()               {}
func (*ScreenInfo) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{14} }

func (m *ScreenInfo) GetFrameId() int64 {
	if m != nil {
		return m.FrameId
	}
	return 0
}

func (m *ScreenInfo) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *ScreenInfo) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ScreenInfo) GetActives() []*ScreenActive {
	if m != nil {
		return m.Actives
	}
	return nil
}

func (m *ScreenInfo) GetDecorations() []*ScreenDecoration {
	if m != nil {
		return m.Decorations
	}
	return nil
}

func (m *ScreenInfo) GetDetails() []*ActiveDetail {
	if m != nil {
		return m.Details
	}
	return nil
}

type ScreenChangeNotify struct {
	FrameId        int64               `protobuf:"varint,1,opt,name=frame_id,json=frameId" json:"frame_id,omitempty"`
	ActiveDetails  []*ActiveDetail     `protobuf:"bytes,2,rep,name=active_details,json=activeDetails" json:"active_details,omitempty"`
	AddActives     []*ScreenActive     `protobuf:"bytes,3,rep,name=add_actives,json=addActives" json:"add_actives,omitempty"`
	AddDecorations []*ScreenDecoration `protobuf:"bytes,4,rep,name=add_decorations,json=addDecorations" json:"add_decorations,omitempty"`
	DelActiveIds   []string            `protobuf:"bytes,5,rep,name=del_active_ids,json=delActiveIds" json:"del_active_ids,omitempty"`
	DelDecorations []*ScreenDecoration `protobuf:"bytes,6,rep,name=del_decorations,json=delDecorations" json:"del_decorations,omitempty"`
}

func (m *ScreenChangeNotify) Reset()                    { *m = ScreenChangeNotify{} }
func (m *ScreenChangeNotify) String() string            { return proto.CompactTextString(m) }
func (*ScreenChangeNotify) ProtoMessage()               {}
func (*ScreenChangeNotify) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{15} }

func (m *ScreenChangeNotify) GetFrameId() int64 {
	if m != nil {
		return m.FrameId
	}
	return 0
}

func (m *ScreenChangeNotify) GetActiveDetails() []*ActiveDetail {
	if m != nil {
		return m.ActiveDetails
	}
	return nil
}

func (m *ScreenChangeNotify) GetAddActives() []*ScreenActive {
	if m != nil {
		return m.AddActives
	}
	return nil
}

func (m *ScreenChangeNotify) GetAddDecorations() []*ScreenDecoration {
	if m != nil {
		return m.AddDecorations
	}
	return nil
}

func (m *ScreenChangeNotify) GetDelActiveIds() []string {
	if m != nil {
		return m.DelActiveIds
	}
	return nil
}

func (m *ScreenChangeNotify) GetDelDecorations() []*ScreenDecoration {
	if m != nil {
		return m.DelDecorations
	}
	return nil
}

type LoginRoomReq struct {
	RoomId string `protobuf:"bytes,1,opt,name=room_id,json=roomId" json:"room_id,omitempty"`
	PosX   int32  `protobuf:"varint,2,opt,name=pos_x,json=posX" json:"pos_x,omitempty"`
	PosY   int32  `protobuf:"varint,3,opt,name=pos_y,json=posY" json:"pos_y,omitempty"`
}

func (m *LoginRoomReq) Reset()                    { *m = LoginRoomReq{} }
func (m *LoginRoomReq) String() string            { return proto.CompactTextString(m) }
func (*LoginRoomReq) ProtoMessage()               {}
func (*LoginRoomReq) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{16} }

func (m *LoginRoomReq) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *LoginRoomReq) GetPosX() int32 {
	if m != nil {
		return m.PosX
	}
	return 0
}

func (m *LoginRoomReq) GetPosY() int32 {
	if m != nil {
		return m.PosY
	}
	return 0
}

type LoginRoomAck struct {
	Common *CommonAck  `protobuf:"bytes,1,opt,name=common" json:"common,omitempty"`
	Screen *ScreenInfo `protobuf:"bytes,2,opt,name=screen" json:"screen,omitempty"`
}

func (m *LoginRoomAck) Reset()                    { *m = LoginRoomAck{} }
func (m *LoginRoomAck) String() string            { return proto.CompactTextString(m) }
func (*LoginRoomAck) ProtoMessage()               {}
func (*LoginRoomAck) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{17} }

func (m *LoginRoomAck) GetCommon() *CommonAck {
	if m != nil {
		return m.Common
	}
	return nil
}

func (m *LoginRoomAck) GetScreen() *ScreenInfo {
	if m != nil {
		return m.Screen
	}
	return nil
}

type LeaveRoomReq struct {
	RoomId string `protobuf:"bytes,1,opt,name=room_id,json=roomId" json:"room_id,omitempty"`
}

func (m *LeaveRoomReq) Reset()                    { *m = LeaveRoomReq{} }
func (m *LeaveRoomReq) String() string            { return proto.CompactTextString(m) }
func (*LeaveRoomReq) ProtoMessage()               {}
func (*LeaveRoomReq) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{18} }

func (m *LeaveRoomReq) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

type LeaveRoomAck struct {
	Common *CommonAck `protobuf:"bytes,1,opt,name=common" json:"common,omitempty"`
}

func (m *LeaveRoomAck) Reset()                    { *m = LeaveRoomAck{} }
func (m *LeaveRoomAck) String() string            { return proto.CompactTextString(m) }
func (*LeaveRoomAck) ProtoMessage()               {}
func (*LeaveRoomAck) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{19} }

func (m *LeaveRoomAck) GetCommon() *CommonAck {
	if m != nil {
		return m.Common
	}
	return nil
}

type MoveActionReq struct {
	DirectX int32 `protobuf:"varint,1,opt,name=DirectX" json:"DirectX,omitempty"`
	DirectY int32 `protobuf:"varint,2,opt,name=DirectY" json:"DirectY,omitempty"`
	Speed   int32 `protobuf:"varint,3,opt,name=Speed" json:"Speed,omitempty"`
	PosX    int32 `protobuf:"varint,4,opt,name=PosX" json:"PosX,omitempty"`
	PosY    int32 `protobuf:"varint,5,opt,name=PosY" json:"PosY,omitempty"`
	Hp      int32 `protobuf:"varint,12,opt,name=Hp" json:"Hp,omitempty"`
	FullHp  int32 `protobuf:"varint,13,opt,name=FullHp" json:"FullHp,omitempty"`
}

func (m *MoveActionReq) Reset()                    { *m = MoveActionReq{} }
func (m *MoveActionReq) String() string            { return proto.CompactTextString(m) }
func (*MoveActionReq) ProtoMessage()               {}
func (*MoveActionReq) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{20} }

func (m *MoveActionReq) GetDirectX() int32 {
	if m != nil {
		return m.DirectX
	}
	return 0
}

func (m *MoveActionReq) GetDirectY() int32 {
	if m != nil {
		return m.DirectY
	}
	return 0
}

func (m *MoveActionReq) GetSpeed() int32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func (m *MoveActionReq) GetPosX() int32 {
	if m != nil {
		return m.PosX
	}
	return 0
}

func (m *MoveActionReq) GetPosY() int32 {
	if m != nil {
		return m.PosY
	}
	return 0
}

func (m *MoveActionReq) GetHp() int32 {
	if m != nil {
		return m.Hp
	}
	return 0
}

func (m *MoveActionReq) GetFullHp() int32 {
	if m != nil {
		return m.FullHp
	}
	return 0
}

type MoveActionAck struct {
	Common *CommonAck `protobuf:"bytes,1,opt,name=common" json:"common,omitempty"`
}

func (m *MoveActionAck) Reset()                    { *m = MoveActionAck{} }
func (m *MoveActionAck) String() string            { return proto.CompactTextString(m) }
func (*MoveActionAck) ProtoMessage()               {}
func (*MoveActionAck) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{21} }

func (m *MoveActionAck) GetCommon() *CommonAck {
	if m != nil {
		return m.Common
	}
	return nil
}

type ChoseHeroReq struct {
	HeroName   string   `protobuf:"bytes,1,opt,name=hero_name,json=heroName" json:"hero_name,omitempty"`
	HeroSkin   string   `protobuf:"bytes,2,opt,name=hero_skin,json=heroSkin" json:"hero_skin,omitempty"`
	HeroSkills []string `protobuf:"bytes,3,rep,name=hero_skills,json=heroSkills" json:"hero_skills,omitempty"`
}

func (m *ChoseHeroReq) Reset()                    { *m = ChoseHeroReq{} }
func (m *ChoseHeroReq) String() string            { return proto.CompactTextString(m) }
func (*ChoseHeroReq) ProtoMessage()               {}
func (*ChoseHeroReq) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{22} }

func (m *ChoseHeroReq) GetHeroName() string {
	if m != nil {
		return m.HeroName
	}
	return ""
}

func (m *ChoseHeroReq) GetHeroSkin() string {
	if m != nil {
		return m.HeroSkin
	}
	return ""
}

func (m *ChoseHeroReq) GetHeroSkills() []string {
	if m != nil {
		return m.HeroSkills
	}
	return nil
}

type ChoseHeroAck struct {
	Common *CommonAck `protobuf:"bytes,1,opt,name=common" json:"common,omitempty"`
}

func (m *ChoseHeroAck) Reset()                    { *m = ChoseHeroAck{} }
func (m *ChoseHeroAck) String() string            { return proto.CompactTextString(m) }
func (*ChoseHeroAck) ProtoMessage()               {}
func (*ChoseHeroAck) Descriptor() ([]byte, []int) { return fileDescriptor_cspayload, []int{23} }

func (m *ChoseHeroAck) GetCommon() *CommonAck {
	if m != nil {
		return m.Common
	}
	return nil
}

func init() {
	proto.RegisterType((*CommonAck)(nil), "CommonAck")
	proto.RegisterType((*BlankMsg)(nil), "BlankMsg")
	proto.RegisterType((*HeadBeat)(nil), "HeadBeat")
	proto.RegisterType((*LoginReq)(nil), "LoginReq")
	proto.RegisterType((*LoginAck)(nil), "LoginAck")
	proto.RegisterType((*LogoutReq)(nil), "LogoutReq")
	proto.RegisterType((*LogoutAck)(nil), "LogoutAck")
	proto.RegisterType((*SetPlayerNameReq)(nil), "SetPlayerNameReq")
	proto.RegisterType((*SetPlayerNameAck)(nil), "SetPlayerNameAck")
	proto.RegisterType((*ChoseSideReq)(nil), "ChoseSideReq")
	proto.RegisterType((*ChoseSideAck)(nil), "ChoseSideAck")
	proto.RegisterType((*ScreenActive)(nil), "ScreenActive")
	proto.RegisterType((*ActiveDetail)(nil), "ActiveDetail")
	proto.RegisterType((*ScreenDecoration)(nil), "ScreenDecoration")
	proto.RegisterType((*ScreenInfo)(nil), "ScreenInfo")
	proto.RegisterType((*ScreenChangeNotify)(nil), "ScreenChangeNotify")
	proto.RegisterType((*LoginRoomReq)(nil), "LoginRoomReq")
	proto.RegisterType((*LoginRoomAck)(nil), "LoginRoomAck")
	proto.RegisterType((*LeaveRoomReq)(nil), "LeaveRoomReq")
	proto.RegisterType((*LeaveRoomAck)(nil), "LeaveRoomAck")
	proto.RegisterType((*MoveActionReq)(nil), "MoveActionReq")
	proto.RegisterType((*MoveActionAck)(nil), "MoveActionAck")
	proto.RegisterType((*ChoseHeroReq)(nil), "ChoseHeroReq")
	proto.RegisterType((*ChoseHeroAck)(nil), "ChoseHeroAck")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_cspayload) }

var fileDescriptor_cspayload = []byte{
	// 919 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x56, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0xd6, 0x8c, 0x63, 0x7b, 0x5c, 0xb6, 0x03, 0x34, 0x08, 0x8c, 0x00, 0xb1, 0x6a, 0x56, 0x2c,
	0x27, 0x1f, 0x12, 0x4e, 0x48, 0x1c, 0xf2, 0x23, 0x14, 0x23, 0x36, 0x5a, 0xc6, 0x91, 0x88, 0xc5,
	0xc1, 0x9a, 0xf5, 0x74, 0xec, 0x51, 0xc6, 0x6e, 0x33, 0xdd, 0x0e, 0xe4, 0x91, 0x38, 0x73, 0xe1,
	0x04, 0x2f, 0xc2, 0xc3, 0x50, 0xd5, 0x3f, 0x76, 0x67, 0xd7, 0xf6, 0xfa, 0x56, 0x7f, 0xfe, 0xaa,
	0xbe, 0xaa, 0x9a, 0x6a, 0x43, 0x5b, 0x3f, 0x2e, 0x85, 0xea, 0x2f, 0x2b, 0xa9, 0x25, 0xff, 0x15,
	0x5a, 0x17, 0x72, 0x3e, 0x97, 0x8b, 0xb3, 0xc9, 0x3d, 0xfb, 0x18, 0x1a, 0x4a, 0x67, 0x7a, 0xa5,
	0x7a, 0xd1, 0xb3, 0xe8, 0x9b, 0x7a, 0xea, 0x34, 0xf6, 0x39, 0xb4, 0x74, 0x31, 0x17, 0x43, 0x9d,
	0xcd, 0x97, 0xbd, 0x18, 0x5d, 0xdd, 0x74, 0x63, 0x60, 0x3d, 0x68, 0xce, 0x85, 0x52, 0xd9, 0x54,
	0xf4, 0x6a, 0xe8, 0x6b, 0xa5, 0x5e, 0xe5, 0x1c, 0x92, 0xf3, 0x32, 0x5b, 0xdc, 0xbf, 0x54, 0x53,
	0xc2, 0x9e, 0x89, 0x72, 0x29, 0x2a, 0x83, 0x9d, 0xa4, 0x4e, 0xe3, 0xdf, 0x43, 0x72, 0x25, 0xb2,
	0xfc, 0x5c, 0x64, 0x9a, 0x7d, 0x01, 0x40, 0xb0, 0x63, 0x65, 0x12, 0xd9, 0x1a, 0x82, 0x44, 0x0c,
	0x8e, 0x26, 0x32, 0x17, 0xa6, 0x82, 0x7a, 0x6a, 0x64, 0x7e, 0x03, 0xc9, 0x4f, 0x72, 0x5a, 0x2c,
	0x52, 0xf1, 0x1b, 0xfb, 0x04, 0x9a, 0x2b, 0x25, 0xaa, 0x71, 0x91, 0x9b, 0xdf, 0xb6, 0xd2, 0x06,
	0xa9, 0x83, 0x9c, 0x7e, 0xb8, 0x5a, 0xa1, 0x35, 0x36, 0x56, 0x23, 0xb3, 0xcf, 0xa0, 0x85, 0xce,
	0x07, 0x1b, 0x6e, 0xeb, 0x4e, 0xac, 0x61, 0x90, 0xf3, 0x7f, 0x23, 0x07, 0x4b, 0x5d, 0xe1, 0xd0,
	0x98, 0x98, 0x16, 0x19, 0xd4, 0xf6, 0x09, 0xf4, 0xd7, 0x1d, 0x4b, 0x9d, 0x27, 0x4c, 0x1d, 0x3f,
	0x49, 0x8d, 0x69, 0x8c, 0x63, 0x91, 0xcd, 0x7d, 0x7b, 0x12, 0x32, 0x5c, 0xa3, 0x4e, 0xce, 0x99,
	0xa8, 0xa4, 0x75, 0x1e, 0x59, 0x27, 0x19, 0x9e, 0x38, 0xd5, 0x7d, 0xb1, 0xe8, 0xd5, 0x37, 0xce,
	0x21, 0xea, 0xec, 0x4b, 0x68, 0x7b, 0x67, 0x59, 0xaa, 0x5e, 0xe3, 0x59, 0x0d, 0xdd, 0xe0, 0xdc,
	0x68, 0xe1, 0xcf, 0xa1, 0x85, 0x04, 0xe4, 0x4a, 0xef, 0x6b, 0xcc, 0x26, 0x8a, 0x78, 0xee, 0x8c,
	0xfa, 0x1a, 0xde, 0x1f, 0x0a, 0xfd, 0xaa, 0xcc, 0x1e, 0x6d, 0xdd, 0x04, 0x89, 0x2d, 0x35, 0x55,
	0xdb, 0x48, 0x23, 0xf3, 0x1f, 0xdf, 0x88, 0x3b, 0xb4, 0x79, 0x1e, 0x2b, 0x0e, 0xb0, 0x38, 0x74,
	0x2e, 0x66, 0x52, 0x89, 0x61, 0x91, 0xfb, 0x7c, 0x24, 0xfa, 0x7c, 0x24, 0xf3, 0x93, 0x20, 0xe6,
	0xc0, 0x5c, 0xfc, 0xef, 0x08, 0x3a, 0xc3, 0x49, 0x25, 0x04, 0x5a, 0x75, 0xf1, 0x20, 0xd8, 0x31,
	0xc4, 0x03, 0x4f, 0x18, 0x25, 0xda, 0xe6, 0xcb, 0xa2, 0x12, 0x13, 0x7d, 0xeb, 0xf6, 0xcc, 0xab,
	0x1b, 0xcf, 0xc8, 0x0c, 0x72, 0xed, 0x19, 0xb1, 0x8f, 0xa0, 0x3e, 0x5c, 0x0a, 0x91, 0x9b, 0x19,
	0xd6, 0x53, 0xab, 0x50, 0xc9, 0xaf, 0xa4, 0xba, 0x35, 0xb3, 0xc3, 0x75, 0x25, 0xd9, 0xd9, 0x46,
	0x38, 0x30, 0x6f, 0x1b, 0x91, 0xed, 0x06, 0xbf, 0xc8, 0x5e, 0xd3, 0xda, 0x48, 0x5e, 0xd3, 0x4d,
	0x02, 0xba, 0xff, 0x60, 0xe9, 0xb6, 0xe8, 0x4b, 0xa1, 0xb3, 0xa2, 0x7c, 0xab, 0x74, 0xfc, 0x4c,
	0xcf, 0x45, 0x29, 0x17, 0x53, 0x75, 0x23, 0x5d, 0x33, 0x37, 0x06, 0x82, 0xbc, 0xde, 0x2c, 0xa1,
	0x91, 0x89, 0xd2, 0x70, 0xf5, 0xda, 0x64, 0xb7, 0xeb, 0xe7, 0x55, 0x53, 0xc0, 0x66, 0xf1, 0x8c,
	0x4c, 0xf9, 0xae, 0x96, 0xae, 0x74, 0x94, 0xe8, 0x93, 0xfe, 0x61, 0x55, 0x96, 0x68, 0xb3, 0xa5,
	0x3b, 0xcd, 0x9c, 0x11, 0xbb, 0x97, 0x89, 0xd9, 0x4b, 0xa7, 0xf1, 0x2b, 0xdc, 0x0f, 0xd3, 0xfa,
	0x4b, 0x31, 0x91, 0x55, 0xa6, 0x0b, 0x9c, 0x3d, 0x56, 0x70, 0xa6, 0x94, 0xd0, 0x8e, 0x08, 0x36,
	0xd5, 0xa9, 0xac, 0x03, 0x91, 0x1f, 0x41, 0x74, 0x4b, 0x9a, 0x6f, 0x7b, 0x34, 0xe2, 0xff, 0x45,
	0x00, 0x16, 0x6a, 0xb0, 0xb8, 0x93, 0xec, 0x53, 0x48, 0xee, 0x2a, 0xe4, 0xe3, 0x57, 0xb7, 0x96,
	0x36, 0x8d, 0x8e, 0x28, 0x38, 0x9a, 0xdf, 0x8b, 0x5c, 0xcf, 0x1c, 0x92, 0x55, 0xec, 0x31, 0x2a,
	0xa6, 0x33, 0xed, 0x20, 0x9d, 0xc6, 0x5e, 0x40, 0x33, 0x33, 0x1d, 0x56, 0xd8, 0x8f, 0x1a, 0xae,
	0x50, 0xb7, 0x1f, 0x2e, 0x4b, 0xea, 0xbd, 0xec, 0x14, 0xda, 0xf9, 0x9a, 0x84, 0xc2, 0x2e, 0x51,
	0xf0, 0x07, 0xfd, 0x37, 0xe9, 0xa5, 0x61, 0x14, 0xa1, 0xe7, 0x66, 0x72, 0xf6, 0x83, 0x25, 0xf4,
	0x70, 0x9e, 0xa9, 0xf7, 0xf2, 0xbf, 0x62, 0x60, 0x16, 0xea, 0x62, 0x96, 0x2d, 0xa6, 0xe2, 0x5a,
	0xea, 0xe2, 0xee, 0x71, 0x1f, 0xcd, 0x6f, 0xe1, 0xd8, 0x96, 0x36, 0xf6, 0x19, 0xe2, 0x6d, 0x19,
	0xba, 0x59, 0xa0, 0x29, 0xd6, 0x87, 0x76, 0x96, 0xe7, 0x63, 0x4f, 0xb9, 0xb6, 0x8d, 0x32, 0x60,
	0xc4, 0x99, 0x63, 0xfd, 0x1d, 0xbc, 0x47, 0xf1, 0x21, 0xf3, 0xa3, 0x5d, 0xcc, 0x8f, 0x31, 0xf2,
	0x32, 0x20, 0xff, 0x1c, 0x8e, 0x73, 0x51, 0xba, 0x5c, 0xc8, 0xc0, 0x36, 0xad, 0x95, 0x76, 0xd0,
	0x6a, 0xf1, 0x07, 0xb9, 0xc9, 0x40, 0x51, 0x61, 0x86, 0xc6, 0xce, 0x0c, 0x18, 0x19, 0x64, 0xe0,
	0x3f, 0x43, 0xc7, 0x3e, 0x05, 0x52, 0xce, 0xdd, 0xd5, 0xab, 0x50, 0x0c, 0xee, 0x19, 0xa9, 0xd8,
	0xac, 0x0f, 0xa1, 0xbe, 0x94, 0x6a, 0xfc, 0x87, 0x7f, 0x48, 0x96, 0xf4, 0x65, 0x3a, 0xe3, 0xa3,
	0xdb, 0x08, 0x32, 0x8e, 0xf8, 0x2f, 0x01, 0xe4, 0xa1, 0xd7, 0xec, 0x2b, 0xdc, 0x7e, 0x53, 0xaa,
	0x81, 0x6f, 0x9f, 0xb4, 0xfb, 0x9b, 0x4d, 0x4d, 0x9d, 0x8b, 0xbf, 0x40, 0x60, 0x91, 0x61, 0x7b,
	0xdf, 0x51, 0x2b, 0xdd, 0xb8, 0x75, 0xe0, 0xa1, 0x37, 0xee, 0xcf, 0x08, 0xba, 0x2f, 0xe5, 0x83,
	0xa0, 0xb6, 0x4a, 0xf3, 0x32, 0x06, 0x47, 0x2d, 0xda, 0x79, 0xd4, 0xe2, 0x1d, 0x47, 0xad, 0xb6,
	0xed, 0xa8, 0x1d, 0x6d, 0x39, 0x6a, 0xf5, 0xe0, 0xa8, 0xd9, 0x5b, 0xd1, 0xd9, 0x72, 0x2b, 0xba,
	0xe1, 0xad, 0xe0, 0xa7, 0x61, 0xa9, 0x87, 0x12, 0x2c, 0xdc, 0xe1, 0xbf, 0xc2, 0xf7, 0x8e, 0xe8,
	0x3d, 0x79, 0x47, 0xa3, 0x7d, 0xef, 0x68, 0xbc, 0xff, 0x1d, 0xad, 0xbd, 0xf5, 0x8e, 0x9e, 0x04,
	0xa9, 0x0e, 0x2c, 0xef, 0x75, 0xc3, 0xfc, 0xb5, 0x3a, 0xfd, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x66,
	0x70, 0xae, 0x92, 0x69, 0x09, 0x00, 0x00,
}
