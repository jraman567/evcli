// Copyright 2023 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package sev

/**
 * The following contains a description of the EAT for SEV
 */

type SevToken struct {
	Evidence	SevEvidence	`cbor:"99989,keyasint" json:"evidence"`
	ReferenceId	string		`cbor:"99988,keyasint" json:"reference-id"`
	TrustAnchorId	string		`cbor:"99987,keyasint" json:"trust-anchor-id"`
}

type SevEvidence struct {
	Nonce			[]byte			`cbor:"99979,keyasint" json:"sev-nonce"`
	UniqueId		string			`cbor:"99978,keyasint" json:"sev-uniqueid"`
	Keys			SevKeys			`cbor:"99977,keyasint" json:"sev-keys"`
	Core			SevCoreComponents	`cbor:"99976,keyasint" json:"sev-core-components"`
	SwInfo			SevSwInfo		`cbor:"99975,keyasint" json:"sev-sw-info"`
	InstanceInfo		SevInstanceInfo		`cbor:"99973,keyasint" json:"sev-instance-info"`
}

type SevKeys struct {
	Cek	[]byte	`cbor:"99499,keyasint" json:"cek"`
	Pek	[]byte	`cbor:"99498,keyasint" json:"pek"`
	Oca	[]byte	`cbor:"99497,keyasint" json:"oca"`
	Pdh	[]byte	`cbor:"99496,keyasint" json:"pdh"`
}

type SevCoreComponents struct {
	HwModel			string	`cbor:"99399,keyasint" json:"hw-model"`
	AttestationReport	[]byte	`cbor:"99398,keyasint" json:"attestation-report"`
}

type SevSwInfo struct {
	Kernel	string	`cbor:"99349,keyasint" json:"hw-model"`
}

type SevInstanceInfo struct {
	AvailabilityDomain	string		`cbor:"99199,keyasint" json:"availabilityDomain"`
	CanonicalRegionName	string		`cbor:"99198,keyasint" json:"canonicalRegionName"`
	CompartmentId		string		`cbor:"99197,keyasint" json:"compartmentId"`
	DefinedTags		SevDefinedTags	`cbor:"99196,keyasint" json:"definedTags"`
	DisplayName		string		`cbor:"99195,keyasint" json:"displayName"`
	FaultDomain		string		`cbor:"99194,keyasint" json:"faultDomain"`
	Hostname		string		`cbor:"99193,keyasint" json:"hostname"`
	Id			string		`cbor:"99192,keyasint" json:"id"`
	Image			string		`cbor:"99191,keyasint" json:"image"`
	Metadata		SevMetadata	`cbor:"99190,keyasint" json:"metadata"`
	ociAdName		string		`cbor:"99189,keyasint" json:"ociAdName"`
	Region			string		`cbor:"99188,keyasint" json:"region"`
	RegionInfo		SevRegionInfo	`cbor:"99187,keyasint" json:"regionInfo"`
	Shape			string		`cbor:"99186,keyasint" json:"shape"`
	ShapeConfig		SevShapeConfig	`cbor:"99185,keyasint" json:"shapeConfig"`
	State			string		`cbor:"99184,keyasint" json:"state"`
	TimeCreated		float64		`cbor:"99183,keyasint" json:"timeCreated"`
}

type SevDefinedTags struct {
	Operations		SevOperations			`cbor:"99129,keyasint" json:"Operations"`
	OracleRecommendedTags	SevOracleRecommendedTags	`cbor:"99128,keyasint" json:"Oracle-Recommended-Tags"`
	OracleTags		SevOracleTags			`cbor:"99127,keyasint" json:"Oracle-Tags"`
}

type SevOperations struct {
	CreateBy		string	`cbor:"99099,keyasint" json:"CreateBy"`
	CreatedDateTime		string	`cbor:"99098,keyasint" json:"CreatedDateTime"`
}

type SevOracleRecommendedTags struct {
	ResourceOwner		string	`cbor:"99089,keyasint" json:"ResourceOwner"`
	ResourceType		string	`cbor:"99088,keyasint" json:"ResourceType"`
}

type SevOracleTags struct {
	CreatedBy		string	`cbor:"99079,keyasint" json:"CreatedBy"`
	CreatedOn		string	`cbor:"99078,keyasint" json:"CreatedOn"`
}

type SevMetadata struct {
	SshAuthorizedKeys	string	`cbor:"99069,keyasint" json:"ssh_authorized_keys"`
}

type SevRegionInfo struct {
	RealmDomainComponent	string	`cbor:"99059,keyasint" json:"realmDomainComponent"`
	RealmKey		string	`cbor:"99058,keyasint" json:"realmKey"`
	RegionIdentifier	string	`cbor:"99057,keyasint" json:"regionIdentifier"`
	RegionKey		string	`cbor:"99056,keyasint" json:"regionKey"`
}

type SevShapeConfig struct {
	BaselineOcpuUtilization		string	`cbor:"99049,keyasint" json:"baselineOcpuUtilization"`
	MaxVnicAttachments		float64	`cbor:"99048,keyasint" json:"maxVnicAttachments"`
	MemoryInGBs			float64	`cbor:"99047,keyasint" json:"memoryInGBs"`
	NetworkingBandwidthInGbps	float64	`cbor:"99046,keyasint" json:"networkingBandwidthInGbps"`
	Ocpus				float64	`cbor:"99045,keyasint" json:"ocpus"`
}
