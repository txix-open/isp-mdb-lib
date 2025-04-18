package consts

type ValidationStatus = string

const (
	ValidationStatusPreparingToRequest = ValidationStatus("PREPARING_TO_REQUEST")
	ValidationStatusPending            = ValidationStatus("PENDING")
	ValidationStatusProcessing         = ValidationStatus("PROCESSING")
	ValidationStatusSuccess            = ValidationStatus("SUCCESS")
	ValidationStatusFailed             = ValidationStatus("FAILED")
	ValidationStatusCanceled           = ValidationStatus("CANCELED")
	ValidationStatusExpired            = ValidationStatus("EXPIRED")
)

type ValidationSystem = string

const (
	ValidationIPEVSystem = ValidationSystem("IPEV")
	ValidationKSRDSystem = ValidationSystem("KSRD")
	ValidationEFSPSystem = ValidationSystem("EFSP")
	ValidationMDMSystem  = ValidationSystem("MDM")
	ValidationSystem4ME  = ValidationSystem("4ME")
	ValidationIsvdSystem = ValidationSystem("ISVD")
	ValidationMSHNSystem = ValidationSystem("MSHN")
)
