package vocabulary

const (
	psPs     = "ps"
	psUptime = "uptime"
)

func (v *vocabulary) LoadPsUtil() error {

	v.NewClass("psutil")

	return nil
}
