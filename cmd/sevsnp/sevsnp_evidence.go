package sevsnp

type Attributes struct {
	Version string
	Policy string
	Family_id string
	Image_id string
	Signature_algo string
	Platform_version string
	Platform_info string
	Flags string
	Measurement string
	Host_data string
	Id_key_digest string
	Author_key_digest string
	Reported_tcb string
	Chip_id string
}
type Evidence struct {
	Scheme 	string `json:"scheme"`
	Type   	string `json:"type"`
	SubType	string `json:"sub_type"`
	Attr	Attributes `json:"attributes"`
}
