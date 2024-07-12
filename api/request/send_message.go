package request

type SendMessage struct {
	ChatRoomId string `json:"chatRoomId"`
	ToMemberId string `json:"toMemberId"`
	Message    string `json:"message"`
	MemberId   uint   `json:"MemberId"`
}
