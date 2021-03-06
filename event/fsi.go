package event

var (
	// ETFSICreateLinkEvent type for when FSI creates a link between a dataset and working directory
	ETFSICreateLinkEvent = Topic("fsi:createLinkEvent")
)

// FSICreateLinkEvent describes an FSI created link
type FSICreateLinkEvent struct {
	FSIPath  string
	Username string
	Dsname   string
}
