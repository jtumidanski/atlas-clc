package constants

const (
	DeletedOrBlocked           byte = 3
	IncorrectPassword          byte = 4
	NotRegistered              byte = 5
	SystemError                byte = 6
	AlreadyLoggedIn            byte = 7
	SystemError2               byte = 8
	SystemError3               byte = 9
	TooManyConnections         byte = 10
	AgeLimit                   byte = 11
	UnableToLogOnAsMasterIp    byte = 13
	WrongGateway               byte = 14
	ProcessingRequest          byte = 15
	AccountVerificationNeeded  byte = 16
	WrongPersonalInformation   byte = 17
	AccountVerificationNeeded2 byte = 21
	LicenseAgreement           byte = 23
	MapleEuropeNotice          byte = 25
	FullClientNotice           byte = 27
)

func GetLoginFailedReason(response string) byte {
	switch response {
	case "DELETED_OR_BLOCKED":
		return DeletedOrBlocked
	case "INCORRECT_PASSWORD":
		return IncorrectPassword
	case "NOT_REGISTERED":
		return NotRegistered
	case "SYSTEM_ERROR":
		return SystemError
	case "ALREADY_LOGGED_IN":
		return AlreadyLoggedIn
	case "SYSTEM_ERROR_2":
		return SystemError2
	case "SYSTEM_ERROR_3":
		return SystemError3
	case "TOO_MANY_CONNECTIONS":
		return TooManyConnections
	case "AGE_LIMIT":
		return AgeLimit
	case "UNABLE_TO_LOG_ON_AS_MASTER_AT_IP":
		return UnableToLogOnAsMasterIp
	case "WRONG_GATEWAY":
		return WrongGateway
	case "PROCESSING_REQUEST":
		return ProcessingRequest
	case "ACCOUNT_VERIFICATION_NEEDED":
		return AccountVerificationNeeded
	case "WRONG_PERSONAL_INFORMATION":
		return WrongPersonalInformation
	case "ACCOUNT_VERIFICATION_NEEDED_2":
		return AccountVerificationNeeded2
	case "LICENSE_AGREEMENT":
		return LicenseAgreement
	case "MAPLE_EUROPE_NOTICE":
		return MapleEuropeNotice
	case "FULL_CLIENT_NOTICE":
		return FullClientNotice
	default:
		return SystemError
	}
}
