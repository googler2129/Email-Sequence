package requests

type UpdateSeqStepSvcReq struct {
	SequenceID uint64
	StepID     uint64
	Subject    string
	Content    string
	UserID     string
	UserName   string
}
