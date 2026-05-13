package mypolymarketapi

type DataAPIType int

const (
	// Misc
	DataGetLiveVolumeForEvent      DataAPIType = iota // GET Get live volume for an event
	DataGetOpenInterest                               // GET Get open interest
	DataDownloadAccountingSnapshot                    // GET Download an accounting snapshot (ZIP of CSVs)

	// Profiles
	DataGetCurrentPositions // GET Get current positions for a user
	DataGetClosedPositions  // GET Get closed positions for a user
)

var DataAPITypeMap = map[DataAPIType]string{
	// Misc
	DataGetLiveVolumeForEvent:      "/live-volume",            // GET Get live volume for an event
	DataGetOpenInterest:            "/oi",                     // GET Get open interest
	DataDownloadAccountingSnapshot: "/v1/accounting/snapshot", // GET Download an accounting snapshot (ZIP of CSVs)

	// Profiles
	DataGetCurrentPositions: "/positions",        // GET Get current positions for a user
	DataGetClosedPositions:  "/closed-positions", // GET Get closed positions for a user
}
