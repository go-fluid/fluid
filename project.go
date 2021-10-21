package fluid

type Project struct {
	Name        string     `json:"name" validate:"required"`
	Description string     `json:"description,omitempty"`
	Version     string     `json:"version" validate:"required"`
	Portals     []Portal   `json:"portals,omitempty" validate:"dive"`
	Entities    []Entity   `json:"entities,omitempty" validate:"dive"`
	Contracts   []Contract `json:"contracts,omitempty" validate:"dive"`
}
