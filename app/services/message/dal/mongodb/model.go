package mongodb

type MgMessage struct {
	ThreadId    string `json:"thread_id" bson:"thread_id"`
	FromUserId  int64  `json:"from_user_id" bson:"from_user_id"`
	ToUserId    int64  `json:"to_user_id" bson:"to_user_id"`
	Contents    string `json:"contents" bson:"contents"`
	MessageUUID int64  `json:"uuid,omitempty" bson:"uuid,omitempty"`
	CreateTime  int64  `json:"create_time" bson:"create_time"`
}

type Thread struct {
	ThreadId      string `json:"thread_id" bson:"thread_id"`
	MgMessageUUID string `json:"uuid,omitempty" bson:"uuid,omitempty"`
}
