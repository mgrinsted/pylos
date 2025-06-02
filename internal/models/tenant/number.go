package models

type TenantNumber struct {
	ID         int    `json:"id"`
	MSISDN     int64  `json:"msisdn"`
	WhatsAppID string `json:"whatsapp_id"`
	WeChatID   string `json:"wechat_id"`
	Billable   string `json:"billable"`
	Status     int    `json:"status"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Email      string `json:"email,omitempty"`
	RACFID     string `json:"racfid,omitempty"`
	EmployeeID int    `json:"employee_id,omitempty"`
	ECDID      string `json:"ecdid,omitempty"`
	CustomerID int    `json:"customer_id,omitempty"`
	Metadata   string `json:"metadata,omitempty"`
	Service    int    `json:"service"`
	Transcribe int    `json:"transcribe,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	DeletedAt  string `json:"deleted_at,omitempty"`
}
