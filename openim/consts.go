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
	TextMessage        MessageType = 101
	ImageMessage       MessageType = 102
	AudioMessage       MessageType = 103
	VideoMessage       MessageType = 104
	FileMessage        MessageType = 105
	MentionMessage     MessageType = 106
	LocationMessage    MessageType = 109
	CustomMessage      MessageType = 110
	SystemNotification MessageType = 1400
)

type ConversationType = int32

const (
	SingleChat ConversationType = 1
	GroupChat  ConversationType = 3
)
