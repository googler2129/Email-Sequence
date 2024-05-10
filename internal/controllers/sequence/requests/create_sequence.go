package requests

type CreateSeqCtrlReq struct {
	Name                   string          `json:"name"`
	IsOpenTrackingEnabled  bool            `json:"is_open_tracking_enabled"`
	IsClickTrackingEnabled bool            `json:"is_click_tracking_enabled"`
	SequenceSteps          []SequenceSteps `json:"sequence_steps"`
	UserID                 string          `json:"user_id"`
	UserName               string          `json:"user_name"`
}

type SequenceSteps struct {
	Subject     string `json:"subject"`
	Content     string `json:"content"`
	WaitingDays uint64 `json:"waiting_days"`
	SerialOrder uint64 `json:"serial_order"`
}
