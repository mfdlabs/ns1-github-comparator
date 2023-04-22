package account

// ActivityLog wraps an NS1 /account/activity resource
type ActivityLog struct {
	Action       string      `json:"action,omitempty"`
	ID           string      `json:"id,omitempty"`
	ResourceID   string      `json:"resource_id,omitempty"`
	Resource     interface{} `json:"resource,omitempty"` // This is an interface{} because it depends on the resource type
	ResourceType string      `json:"resource_type,omitempty"`
	Timestamp    int         `json:"timestamp,omitempty"`
	UserID       string      `json:"user_id,omitempty"`
	UserName     string      `json:"user_name,omitempty"`
	UserType     string      `json:"user_type,omitempty"`
}
