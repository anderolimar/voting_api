package models

import "encoding/json"

type PollSummaryRequest struct {
	PollID string `json:"pollID"`
}

type PollResponse struct {
	Response
	Poll *Poll `json:"poll,omitempty"`
}

type PollSummaryResponse struct {
	Response
	TotalVotes int   `json:"totalVotes"`
	Poll       *Poll `json:"poll"`
}

type PollRequest struct {
	Title   string       `json:"title"`
	Options []VoteOption `json:"options"`
}

type VoteRequest struct {
	CaptchaID    string `json:"captchaID"`
	CaptchaInput string `json:"captchaInput"`
	PollID       string `json:"pollID"`
	Vote         int    `json:"vote"`
	Retry        int    `json:"retry,omitempty"`
}

type VoteResponse struct {
	Response
	Poll Poll `json:"poll"`
}

type VoteMessage struct {
	ID   string     `json:"id"`
	Vote VoteOption `json:"vote"`
}

type Poll struct {
	ID      string       `json:"id"`
	Title   string       `json:"title"`
	Options []VoteOption `json:"options"`
}

func (p Poll) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

type VoteOption struct {
	Index    int    `json:"index" bson:"index"`
	Title    string `json:"title" bson:"title"`
	Quantity int    `json:"quantity,omitempty" bson:"quantity"`
}
