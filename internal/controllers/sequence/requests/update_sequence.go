package requests

type UpdateSeqCtrlReq struct {
	IsOpenTrackingEnabled  bool   `json:"is_open_tracking_enabled"`
	IsClickTrackingEnabled bool   `json:"is_click_tracking_enabled"`
	UserID                 string `json:"user_id"`
	UserName               string `json:"user_name"`
}
