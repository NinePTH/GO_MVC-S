package auth

type GenerateJWTClaimsParams struct {
    Username   string
    Role       string
    PatientID  string // optional
}