package crm

// ResponseStatusTracking .
type ResponseUser struct {
	Total int    `json:"total,omitempty"`
	User  []User `json:"list,omitempty"`
}

type User struct {
	ID                     string   `json:"id,omitempty"`
	Name                   string   `json:"name,omitempty"`
	Deleted                bool     `json:"deleted,omitempty"`
	IsAdmin                bool     `json:"isAdmin,omitempty"`
	UserName               string   `json:"userName,omitempty"`
	Type                   string   `json:"type,omitempty"`
	AuthMethod             string   `json:"authMethod,omitempty"`
	APIKey                 string   `json:"apiKey,omitempty"`
	SalutationName         string   `json:"salutationName,omitempty"`
	FirstName              string   `json:"firstName,omitempty"`
	LastName               string   `json:"lastName,omitempty"`
	IsActive               bool     `json:"isActive,omitempty"`
	IsPortalUser           bool     `json:"isPortalUser,omitempty"`
	IsSuperAdmin           bool     `json:"isSuperAdmin,omitempty"`
	Title                  string   `json:"title,omitempty"`
	EmailAddress           string   `json:"emailAddress,omitempty"`
	PhoneNumber            string   `json:"phoneNumber,omitempty"`
	Gender                 string   `json:"gender,omitempty"`
	CreatedAt              string   `json:"createdAt,omitempty"`
	ModifiedAt             string   `json:"modifiedAt,omitempty"`
	LastAccess             string   `json:"lastAccess,omitempty"`
	MiddleName             string   `json:"middleName,omitempty"`
	EmailAddressIsOptedOut string   `json:"emailAddressIsOptedOut,omitempty"`
	PhoneNumberIsOptedOut  string   `json:"phoneNumberIsOptedOut,omitempty"`
	EmailAddressData       []string `json:"emailAddressData,omitempty"`
	PhoneNumberData        []string `json:"phoneNumberData,omitempty"`
	DefaultTeamID          string   `json:"defaultTeamId,omitempty"`
	DefaultTeamName        string   `json:"defaultTeamName,omitempty"`
	TeamsIds               []string `json:"teamsIds,omitempty"`
	RolesIds               []string `json:"rolesIds,omitempty"`
	PortalsIds             []string `json:"portalsIds,omitempty"`
	PortalRolesIds         []string `json:"portalRolesIds,omitempty"`
	ContactID              string   `json:"contactId,omitempty"`
	ContactName            string   `json:"contactName,omitempty"`
	AccountsIds            []string `json:"accountsIds,omitempty"`
	AvatarID               string   `json:"avatarId,omitempty"`
	AvatarName             string   `json:"avatarName,omitempty"`
	CreatedByID            string   `json:"createdById,omitempty"`
	CreatedByName          string   `json:"createdByName,omitempty"`
	DashboardTemplateID    string   `json:"dashboardTemplateId,omitempty"`
	DashboardTemplateName  string   `json:"dashboardTemplateName,omitempty"`
	UseraccountID          string   `json:"useraccountId,omitempty"`
	UseraccountName        string   `json:"useraccountName,omitempty"`
}
