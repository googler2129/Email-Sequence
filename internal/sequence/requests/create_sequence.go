package requests

type CreateSeqSvcReq struct {
	Name                   string
	IsOpenTrackingEnabled  bool
	IsClickTrackingEnabled bool
	SequenceSteps          []SequenceStep
	UserID                 string
	UserName               string
}

type SequenceStep struct {
	Subject     string
	Content     string
	WaitingDays uint64
	SerialOrder uint64
}
