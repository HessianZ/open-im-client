package openim

const (
	ErrCode_Success        = 0
	ErrCode_RecordNotFound = 1004
	ErrCode_TokenNotExists = 1507
)

// 1: iOS, 2: Android, 3: Windows, 4: OSX, 5: WEB, 6: Mini Program, 7: Linux, 8: AndroidPad, 9: IPad, 10: Admin
type PlatformID = int32

const (
	PlatformID_iOS      PlatformID = 1
	PlatformID_Android  PlatformID = 2
	PlatformID_Windows  PlatformID = 3
	PlatformID_OSX      PlatformID = 4
	PlatformID_Web      PlatformID = 5
	PlatformID_Mini     PlatformID = 6
	PlatformID_Linux    PlatformID = 7
	PlatformID_AndroidP PlatformID = 8
	PlatformID_IPad     PlatformID = 9
	PlatformID_Admin    PlatformID = 10
)

type MessageType = int32

const (
	MessageType_Text               MessageType = 101
	MessageType_Image              MessageType = 102
	MessageType_Audio              MessageType = 103
	MessageType_Video              MessageType = 104
	MessageType_File               MessageType = 105
	MessageType_Mention            MessageType = 106
	MessageType_Location           MessageType = 109
	MessageType_Custom             MessageType = 110
	MessageType_SystemNotification MessageType = 1400
)

type ConversationType = int32

const (
	ConversationType_Single ConversationType = 1
	ConversationType_Group  ConversationType = 3
)
