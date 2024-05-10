package requests

type UpdateSeqStepCtrlReq struct {
	Subject  string `json:"subject" validate:"required"`
	Content  string `json:"content" validate:"required"`
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
}
