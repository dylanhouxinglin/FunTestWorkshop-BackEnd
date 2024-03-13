package data

type UpdateRankReq struct {
	CorrectCnt  int64   `json:"correctCnt" bson:"correctCnt"`
	QuestionCnt int64   `json:"questionCnt" bson:"questionCnt"`
	CorrectRate float64 `json:"correctRate" bson:"correctRate"`
}
