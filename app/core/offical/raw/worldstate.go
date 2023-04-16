package raw

type WorldState struct {
}

type (
	Event struct {
		Id         IdObj `json:"_id"`
		Date       *TimeObj
		Icon       string
		Priority   bool
		MobileOnly bool
		Community  bool
		Messages   []struct {
			LanguageCode string
			Message      string
		}
	}

	// Sortie 突击数据对象
	Sortie struct {
		Id         IdObj `json:"_id"`
		Activation *TimeObj
		Expiry     *TimeObj
		Reward     string
		Seed       int
		Boss       string
		Variants   []Mission
	}

	// Mission 任务数据对象
	Mission struct {
		MissionType  string `json:"missionType"`
		ModifierType string `json:"modifierType,omitempty"`
		Node         string `json:"node"`
		Tileset      string `json:"tileset,omitempty"`
	}

	SyndicateMission struct {
		Id         IdObj `json:"_id"`
		Activation *TimeObj
		Expiry     *TimeObj
		Tag        string
		Nodes      []string
		Jobs       []Job
	}

	Job struct {
		JobType       string `json:"jobType"`
		Rewards       string `json:"rewards"`
		MasterReq     int    `json:"masterReq"`
		MinEnemyLevel int    `json:"minEnemyLevel"`
		MaxEnemyLevel int    `json:"maxEnemyLevel"`
		XpAmounts     []int  `json:"xpAmounts"`
	}

	// Invasion 入侵数据对象
	Invasion struct {
		Id                  IdObj `json:"_id"`
		Faction             string
		DefenderFaction     string
		Node                string
		Count               int
		Goal                int
		LocTag              string
		Completed           bool
		ChainID             IdObj
		AttackerReward      InvasionReward
		AttackerMissionInfo InvasionMissionInfo
		DefenderReward      InvasionReward
		DefenderMissionInfo InvasionMissionInfo
		Activation          *TimeObj
	}

	InvasionReward struct {
		CountedItems []struct {
			ItemType  string
			ItemCount int
		} `json:"countedItems"`
	}

	InvasionMissionInfo struct {
		Seed    int    `json:"seed"`
		Faction string `json:"faction"`
	}

	NodeOverride struct {
		Id   IdObj `json:"_id"`
		Node string
		Hide bool
	}

	VoidTrader struct {
		Id         IdObj `json:"_id"`
		Activation *TimeObj
		Expiry     *TimeObj
		Character  string
		Node       string
	}

	PrimeVaultTrader struct {
		Id                IdObj `json:"_id"`
		Activation        *TimeObj
		Completed         bool
		InitialStartDate  *TimeObj
		Node              string
		Manifest          []Manifest
		Expiry            *TimeObj
		EvergreenManifest []Manifest
		ScheduleInfo      []struct {
			Expiry       *TimeObj
			FeaturedItem string
		}
	}

	Manifest struct {
		ItemType   string
		PrimePrice int
	}

	// VoidStorm 虚空裂隙
	VoidStorm struct {
		Id                IdObj `json:"_id"`
		Activation        *TimeObj
		Expiry            *TimeObj
		Node              string
		ActiveMissionTier string
	}

	DailyDeal struct {
		StoreItem     string
		Activation    *TimeObj
		Expiry        *TimeObj
		Discount      int
		OriginalPrice int
		SalePrice     int
		AmountTotal   int
		AmountSold    int
	}

	SeasonInfo struct {
		Activation       *TimeObj
		Expiry           *TimeObj
		AffiliationTag   string
		Season           int
		Phase            int
		ActiveChallenges []SeasonChallenge
	}

	SeasonChallenge struct {
		Id         IdObj `json:"_id"`
		Activation *TimeObj
		Expiry     *TimeObj
		Daily      bool
		Challenge  string
	}
)
