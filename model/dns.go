package model

type AddZoneInput struct {
	Name        string        `json:"name" validate:"required"`
	Kind        string        `json:"kind"`
	Masters     []interface{} `json:"masters"`
	DNSSEC      bool          `json:"dnssec"`
	NameServers []string      `json:"nameservers"`
}

type RecordIn struct {
	Type string `json:"type"`
	Name string `json:"name"`
	TTL  int    `json:"ttl"`
	Data string `json:"data"`
}
type RecordDeletIn struct {
	Type string `json:"type"`
	Name string `json:"name"`
}
type AddRecodInput struct {
	RRSets []struct {
		Name       string `json:"name" validate:"required"`
		Type       string `json:"type" validate:"required"`
		TTL        int    `json:"ttl"`
		ChangeType string `json:"changetype" validate:"required"`
		Records    []struct {
			Content  string `json:"content"`
			Disabled bool   `json:"disabled"`
		} `json:"records"`
	} `json:"rrsets" validate:"required"`
}
type DeleteRecordInput struct {
	Rrsets []struct {
		Name       string `json:"name"`
		Type       string `json:"type"`
		Changetype string `json:"changetype"`
	} `json:"rrsets"`
}
