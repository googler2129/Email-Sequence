package requests

type UpdateSeqSvcReq struct {
	SequenceID             uint64
	IsOpenTrackingEnabled  bool
	IsClickTrackingEnabled bool
	UserID                 string
	UserName               string
}
