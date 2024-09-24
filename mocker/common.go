package mocker

type RouterType string

const (
	RouterTypeFS    RouterType = "fs"
	RouterTypeProxy            = "proxy"
	RouterTypeAPI              = "api"
	RouterTypeCrud             = "crud"
	RouterTypeGroup            = "group"
)
