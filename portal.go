package fluid

const (
	PortalTypeIonic   = "ionic"
	PortalTypeVuetify = "vuetify"
)

type MaterialColor struct {
	Color    string `json:"color" validate:"required,color"`
	Contrast string `json:"contrast,omitempty" validate:"color"`
	Shade    string `json:"shade,omitempty" validate:"color"`
	Tint     string `json:"tint,omitempty" validate:"color"`
}

type Theme struct {
	Primary   MaterialColor `json:"primary" validate:"required,dive"`
	Secondary MaterialColor `json:"secondary" validate:"required,dive"`
	Tertiary  MaterialColor `json:"tertiary" validate:"required,dive"`
	Success   MaterialColor `json:"success" validate:"required,dive"`
	Warning   MaterialColor `json:"warning" validate:"required,dive"`
	Error     MaterialColor `json:"error" validate:"required,dive"`
	Medium    MaterialColor `json:"medium" validate:"required,dive"`
	Light     MaterialColor `json:"light" validate:"required,dive"`
	Dark      MaterialColor `json:"dark" validate:"required,dive"`
}

type Portal struct {
	Name              string   `json:"name" validate:"required"`
	Type              string   `json:"type" validate:"required,portalType"`
	AccountEntityKeys []string `json:"accountEntityKeys"` // list of different account entities that can log into this platform
	LightTheme        Theme    `json:"lightTheme" validate:"required,dive"`
	DarkTheme         Theme    `json:"darkTheme" validate:"required,dive"`
}
