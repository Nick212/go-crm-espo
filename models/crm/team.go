package crm

// ResponseAccount .
type ResponseTeam struct {
	Total int    `json:"total,omitempty"`
	Team  []Team `json:"list,omitempty"`
}

type Team struct {
	ID         string   `json:"id,omitempty"`
	Name       string   `json:"name,omitempty"`
	Deleted    bool     `json:"deleted,omitempty"`
	CreatedAt  string   `json:"createdAt,omitempty"`
	ModifiedAt string   `json:"modifiedAt,omitempty"`
	ExternalID string   `json:"externalId,omitempty"`
	RolesIds   []string `json:"rolesIds,omitempty"`
}
