package models

type EstateNumber struct {
	ID         int    `json:"id"`
	MSISDN     int64  `json:"msisdn"`
	CustomerID int    `json:"customer_id"`
	WhatsAppID string `json:"whatsapp_id"`
	WeChatID   string `json:"wechat_id"`
	Billable   string `json:"billable"`
	Status     int    `json:"status"`
	CreatedAt  string `json:"created_at"`

	// Optional fields - these fields may not be present in all cases
	UpdatedAt string `json:"updated_at,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
}
