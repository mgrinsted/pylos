package models

import "time"

// Customer represents a customer in the estate database.

type EstateCustomer struct {
	ID                  int        `json:"id" db:"id"`
	Name                *string    `json:"name" db:"name"`
	AffiliatedVendor    *string    `json:"affiliated_vendor" db:"affiliatedVendor"`
	AffiliateID         *string    `json:"affiliate_id" db:"affiliate_id"`
	TenantReferenceName *string    `json:"tenant_reference_name" db:"tenantReferenceName"`
	PortalVersion       *string    `json:"portal_version" db:"portalVersion"`
	Namespace           *string    `json:"namespace" db:"namespace"`
	SipNode             *string    `json:"sip_node" db:"sipNode"`
	Status              *int       `json:"status" db:"status"`
	Created             *time.Time `json:"created" db:"created"`
	Services            *string    `json:"services" db:"services"`
	Storage             *string    `json:"storage" db:"storage"`
	SFTP                *int       `json:"sftp" db:"sftp"`
	OnPrem1             *string    `json:"on_prem1" db:"onprem1"`
	OnPrem2             *string    `json:"on_prem2" db:"onprem2"`
	MetaURL             *string    `json:"meta_url" db:"metaURL"`
	S3Region            *string    `json:"s3_region" db:"s3region"`
	S3Key               *string    `json:"s3_key" db:"s3key"`
	S3Secret            *string    `json:"s3_secret" db:"s3secret"`
	S3Bucket            *string    `json:"s3_bucket" db:"s3bucket"`
	S3Folder            *string    `json:"s3_folder" db:"s3folder"`
	PublicKey           *string    `json:"public_key" db:"publicKey"`
	PublicKey2          *string    `json:"public_key2" db:"publicKey2"`
	SupportEmail        *string    `json:"support_email" db:"supportEmail"`
	SupportNumber       *string    `json:"support_number" db:"supportNumber"`
}

// DTO returns a Data Transfer Object representation of the EstateCustomer.
type CustomerSummaryDTO struct {
	ID            int     `json:"id" db:"id"`
	Name          string  `json:"name" db:"name"`
	AffiliateID   *string `json:"affiliate_id" db:"affiliate_id"`
	Namespace     *string `json:"namespace" db:"namespace"`
	Status        *int    `json:"status" db:"status"`
	Storage       *string `json:"storage" db:"storage"`
	SFTP          *bool   `json:"sftp" db:"sftp"`
	ActiveCount   int     `json:"active_msisdns" db:"active_msisdns"`
	InactiveCount int     `json:"inactive_msisdns" db:"inactive_msisdns"`
	TotalCount    int     `json:"total_msisdns" db:"total_msisdns"`
}
