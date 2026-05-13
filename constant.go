package mypolymarketapi

type TimeInForce string

const (
	TIME_IN_FORCE_GTC TimeInForce = "GTC" // Good-Til-Cancelled：订单持续有效，直到成交、撤销或失效
	TIME_IN_FORCE_FOK TimeInForce = "FOK" // Fill-Or-Kill：必须立即全部成交，否则整单取消
	TIME_IN_FORCE_GTD TimeInForce = "GTD" // Good-Til-Date：订单在指定到期时间前有效，到期后自动失效
	TIME_IN_FORCE_FAK TimeInForce = "FAK" // Fill-And-Kill：立即成交可撮合部分，未成交部分立即取消
)

func (t TimeInForce) String() string {
	return string(t)
}

// Sports Game Status
const (
	// NFL
	NFL_STATUS_SCHEDULED     = "Scheduled"    // Game not yet started
	NFL_STATUS_IN_PROGRESS   = "InProgress"   // Game currently playing
	NFL_STATUS_FINAL         = "Final"        // Game completed in regulation
	NFL_STATUS_F_OT          = "F/OT"         // Final after overtime
	NFL_STATUS_SUSPENDED     = "Suspended"    // Game suspended
	NFL_STATUS_POSTPONED     = "Postponed"    // Game postponed
	NFL_STATUS_DELAYED       = "Delayed"      // Game delayed
	NFL_STATUS_CANCELED      = "Canceled"     // Game canceled
	NFL_STATUS_FORFEIT       = "Forfeit"      // Game forfeited
	NFL_STATUS_NOT_NECESSARY = "NotNecessary" // Scheduled, but not needed

	// NHL
	NHL_STATUS_SCHEDULED     = "Scheduled"    // Game not yet started
	NHL_STATUS_IN_PROGRESS   = "InProgress"   // Game currently playing
	NHL_STATUS_FINAL         = "Final"        // Game completed in regulation
	NHL_STATUS_F_OT          = "F/OT"         // Final after overtime
	NHL_STATUS_F_SO          = "F/SO"         // Final after shootout
	NHL_STATUS_SUSPENDED     = "Suspended"    // Game suspended
	NHL_STATUS_POSTPONED     = "Postponed"    // Game postponed
	NHL_STATUS_DELAYED       = "Delayed"      // Game delayed
	NHL_STATUS_CANCELED      = "Canceled"     // Game canceled
	NHL_STATUS_FORFEIT       = "Forfeit"      // Game forfeited
	NHL_STATUS_NOT_NECESSARY = "NotNecessary" // Scheduled, but not needed

	// MLB
	MLB_STATUS_SCHEDULED     = "Scheduled"    // Game not yet started
	MLB_STATUS_IN_PROGRESS   = "InProgress"   // Game currently playing
	MLB_STATUS_FINAL         = "Final"        // Game completed
	MLB_STATUS_SUSPENDED     = "Suspended"    // Game suspended
	MLB_STATUS_DELAYED       = "Delayed"      // Game delayed
	MLB_STATUS_POSTPONED     = "Postponed"    // Game postponed
	MLB_STATUS_CANCELED      = "Canceled"     // Game canceled
	MLB_STATUS_FORFEIT       = "Forfeit"      // Game forfeited
	MLB_STATUS_NOT_NECESSARY = "NotNecessary" // Scheduled, but not needed

	// NBA and CBB
	NBA_CBB_STATUS_SCHEDULED     = "Scheduled"    // Game not yet started
	NBA_CBB_STATUS_IN_PROGRESS   = "InProgress"   // Game currently playing
	NBA_CBB_STATUS_FINAL         = "Final"        // Game completed
	NBA_CBB_STATUS_F_OT          = "F/OT"         // Final after overtime
	NBA_CBB_STATUS_SUSPENDED     = "Suspended"    // Game suspended
	NBA_CBB_STATUS_POSTPONED     = "Postponed"    // Game postponed
	NBA_CBB_STATUS_DELAYED       = "Delayed"      // Game delayed
	NBA_CBB_STATUS_CANCELED      = "Canceled"     // Game canceled
	NBA_CBB_STATUS_FORFEIT       = "Forfeit"      // Game forfeited
	NBA_CBB_STATUS_NOT_NECESSARY = "NotNecessary" // Scheduled, but not needed

	// CFB
	CFB_STATUS_SCHEDULED   = "Scheduled"  // Game not yet started
	CFB_STATUS_IN_PROGRESS = "InProgress" // Game currently playing
	CFB_STATUS_FINAL       = "Final"      // Game completed
	CFB_STATUS_F_OT        = "F/OT"       // Final after overtime
	CFB_STATUS_SUSPENDED   = "Suspended"  // Game suspended
	CFB_STATUS_POSTPONED   = "Postponed"  // Game postponed
	CFB_STATUS_DELAYED     = "Delayed"    // Game delayed
	CFB_STATUS_CANCELED    = "Canceled"   // Game canceled
	CFB_STATUS_FORFEIT     = "Forfeit"    // Game forfeited

	// Soccer
	SOCCER_STATUS_SCHEDULED        = "Scheduled"       // Game not yet started
	SOCCER_STATUS_IN_PROGRESS      = "InProgress"      // Game currently playing
	SOCCER_STATUS_BREAK            = "Break"           // Halftime or other break
	SOCCER_STATUS_SUSPENDED        = "Suspended"       // Game suspended
	SOCCER_STATUS_PENALTY_SHOOTOUT = "PenaltyShootout" // Penalty shootout in progress
	SOCCER_STATUS_FINAL            = "Final"           // Game completed
	SOCCER_STATUS_AWARDED          = "Awarded"         // Result awarded due to ruling/forfeit
	SOCCER_STATUS_POSTPONED        = "Postponed"       // Game postponed
	SOCCER_STATUS_CANCELED         = "Canceled"        // Game canceled

	// Esports
	ESPORTS_STATUS_NOT_STARTED = "not_started" // Match not yet started
	ESPORTS_STATUS_RUNNING     = "running"     // Match currently playing
	ESPORTS_STATUS_FINISHED    = "finished"    // Match completed
	ESPORTS_STATUS_POSTPONED   = "postponed"   // Match postponed
	ESPORTS_STATUS_CANCELED    = "canceled"    // Match canceled

	// Tennis
	TENNIS_STATUS_SCHEDULED  = "scheduled"  // Match not yet started
	TENNIS_STATUS_INPROGRESS = "inprogress" // Match currently playing
	TENNIS_STATUS_SUSPENDED  = "suspended"  // Match suspended
	TENNIS_STATUS_FINISHED   = "finished"   // Match completed
	TENNIS_STATUS_POSTPONED  = "postponed"  // Match postponed
	TENNIS_STATUS_CANCELLED  = "cancelled"  // Match canceled
)
