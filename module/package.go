package module

const (
	Authority            = "github/advanced-go/agency"
	RouteName            = "observation"
	Version              = "2.2.2"
	Ver1                 = "v1"
	Ver2                 = "v2"
	ObservationAccess    = "access"
	ObservationInference = "inference"
	//ObservationAssignment = "assignment"
)

// Configuration keys used on startup for map values
const (
	PackageNameUserKey     = "user"    // type:text, empty:false
	PackageNamePasswordKey = "pswd"    // type:text, empty:false
	PackageNameRetriesKey  = "retries" // type:int, default:-1, range:0-100
)
