package crm

// ResposeContact .
type ResposeContact struct {
	Total   int       `json:"total,omitempty"`
	Contact []Contact `json:"list,omitempty"`
}

// Contact .
type Contact struct {
	ID                     string   `json:"id,omitempty"`
	Name                   string   `json:"name,omitempty"`
	Deleted                bool     `json:"deleted,omitempty"`
	SalutationName         string   `json:"salutationName,omitempty"`
	FirstName              string   `json:"firstName,omitempty"`
	LastName               string   `json:"lastName,omitempty"`
	AccountID              string   `json:"accountId,omitempty"`
	Title                  string   `json:"title,omitempty"`
	Description            string   `json:"description,omitempty"`
	EmailAddress           string   `json:"emailAddress,omitempty"`
	PhoneNumber            string   `json:"phoneNumber,omitempty"`
	DoNotCall              bool     `json:"doNotCall,omitempty"`
	AddressStreet          string   `json:"addressStreet,omitempty"`
	AddressCity            string   `json:"addressCity,omitempty"`
	AddressState           string   `json:"addressState,omitempty"`
	AddressCountry         string   `json:"addressCountry,omitempty"`
	AddressPostalCode      string   `json:"addressPostalCode,omitempty"`
	AccountIsInactive      bool     `json:"accountIsInactive,omitempty"`
	AccountType            string   `json:"accountType,omitempty"`
	CreatedAt              string   `json:"createdAt,omitempty"`
	ModifiedAt             string   `json:"modifiedAt,omitempty"`
	MainContact            bool     `json:"mainContact,omitempty"`
	NfeContact             bool     `json:"nfeContact,omitempty"`
	EmailAddressIsOptedOut bool     `json:"emailAddressIsOptedOut,omitempty"`
	PhoneNumberIsOptedOut  bool     `json:"phoneNumberIsOptedOut,omitempty"`
	AccountName            string   `json:"accountName,omitempty"`
	CampaignID             string   `json:"campaignId,omitempty"`
	CreatedByID            string   `json:"createdById,omitempty"`
	CreatedByName          string   `json:"createdByName,omitempty"`
	ModifiedByID           string   `json:"modifiedById,omitempty"`
	ModifiedByName         string   `json:"modifiedByName,omitempty"`
	AssignedUserID         string   `json:"assignedUserId,omitempty"`
	AssignedUserName       string   `json:"assignedUserName,omitempty"`
	TeamsIds               []string `json:"teamsIds,omitempty"`
	UserName               string   `json:"userName,omitempty"`
	Password               string   `json:"password,omitempty"`
}
