package server

type ServerStatus int

const (
	StatusOK ServerStatus = iota
	StatusMaintenance
)

var statusText = map[ServerStatus]string{
	StatusOK:          "OK",
	StatusMaintenance: "Maintenance",
}

func ServerStatusText(code ServerStatus) string {
	return statusText[code]
}
