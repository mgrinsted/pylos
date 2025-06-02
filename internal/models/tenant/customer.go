package models

type TenantCustomer struct {
	ID                  int    `json:"id"`
	Name                string `json:"name,omitempty"`
	PortalVersion       string `json:"portal_version,omitempty"`
	Namespace           string `json:"namespace,omitempty"`
	SupportEmail        string `json:"support_email,omitempty"`
	SupportNumber       string `json:"support_number,omitempty"`
	Storage             string `json:"storage,omitempty"`
	SFTP                int    `json:"sftp,omitempty"`
	DefaultIMCRetention int    `json:"default_imc_retention,omitempty"`
	DefaultSMSRetention  int    `json:"default_sms_retention,omitempty"`
	DefaultVoiceRetention int    `json:"default_voice_retention,omitempty"`
	OnPrem1             string `json:"onprem1,omitempty"`
	OnPrem2             string `json:"onprem2,omitempty"`
	MetaURL             string `json:"meta_url,omitempty"`
	S3Region            string `json:"s3_region,omitempty"`
	S3Key               string `json:"s3_key,omitempty"`
	S3Secret            string `json:"s3_secret,omitempty"`
	S3Bucket            string `json:"s3_bucket,omitempty"`
	S3Folder            string `json:"s3_folder,omitempty"`
	PublicKey           string `json:"public_key,omitempty"`
	PublicKey2          string `json:"public_key2,omitempty"`
	// Optionnal fields - these fields may not be present in all cases
	UpdatedAt string `json:"updated_at,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
}