package tibiadata

type CharacterResponse struct {
	Characters struct {
		AccountBadges []struct {
			Description string `json:"description"`
			IconURL     string `json:"icon_url"`
			Name        string `json:"name"`
		} `json:"account_badges"`
		AccountInformation struct {
			Created      string `json:"created"`
			LoyaltyTitle string `json:"loyalty_title"`
			Position     string `json:"position"`
		} `json:"account_information"`
		Achievements []struct {
			Grade  int    `json:"grade"`
			Name   string `json:"name"`
			Secret bool   `json:"secret"`
		} `json:"achievements"`
		Character struct {
			AccountStatus     string   `json:"account_status"`
			AchievementPoints int      `json:"achievement_points"`
			Comment           string   `json:"comment"`
			DeletionDate      string   `json:"deletion_date"`
			FormerNames       []string `json:"former_names"`
			FormerWorlds      []string `json:"former_worlds"`
			Guild             struct {
				Name string `json:"name"`
				Rank string `json:"rank"`
			} `json:"guild"`
			Houses []struct {
				Houseid int    `json:"houseid"`
				Name    string `json:"name"`
				Paid    string `json:"paid"`
				Town    string `json:"town"`
			} `json:"houses"`
			LastLogin      string `json:"last_login"`
			Level          int    `json:"level"`
			MarriedTo      string `json:"married_to"`
			Name           string `json:"name"`
			Residence      string `json:"residence"`
			Sex            string `json:"sex"`
			Title          string `json:"title"`
			Traded         bool   `json:"traded"`
			UnlockedTitles int    `json:"unlocked_titles"`
			Vocation       string `json:"vocation"`
			World          string `json:"world"`
		} `json:"character"`
		Deaths []struct {
			Assists []struct {
				Name   string `json:"name"`
				Player bool   `json:"player"`
				Summon string `json:"summon"`
				Traded bool   `json:"traded"`
			} `json:"assists"`
			Killers []struct {
				Name   string `json:"name"`
				Player bool   `json:"player"`
				Summon string `json:"summon"`
				Traded bool   `json:"traded"`
			} `json:"killers"`
			Level  int    `json:"level"`
			Reason string `json:"reason"`
			Time   string `json:"time"`
		} `json:"deaths"`
		OtherCharacters []struct {
			Deleted bool   `json:"deleted"`
			Main    bool   `json:"main"`
			Name    string `json:"name"`
			Status  string `json:"status"`
			Traded  bool   `json:"traded"`
			World   string `json:"world"`
		} `json:"other_characters"`
	} `json:"characters"`
	Information struct {
		APIVersion int    `json:"api_version"`
		Timestamp  string `json:"timestamp"`
	} `json:"information"`
}

type GuildResponse struct {
	Guilds struct {
		Guild struct {
			Active           bool   `json:"active"`
			Description      string `json:"description"`
			DisbandCondition string `json:"disband_condition"`
			DisbandDate      string `json:"disband_date"`
			Founded          string `json:"founded"`
			Guildhalls       []struct {
				Name      string `json:"name"`
				PaidUntil string `json:"paid_until"`
				World     string `json:"world"`
			} `json:"guildhalls"`
			Homepage string `json:"homepage"`
			InWar    bool   `json:"in_war"`
			Invites  []struct {
				Date string `json:"date"`
				Name string `json:"name"`
			} `json:"invites"`
			LogoURL string `json:"logo_url"`
			Members []struct {
				Joined   string `json:"joined"`
				Level    int    `json:"level"`
				Name     string `json:"name"`
				Rank     string `json:"rank"`
				Status   string `json:"status"`
				Title    string `json:"title"`
				Vocation string `json:"vocation"`
			} `json:"members"`
			MembersInvited   int    `json:"members_invited"`
			MembersTotal     int    `json:"members_total"`
			Name             string `json:"name"`
			OpenApplications bool   `json:"open_applications"`
			PlayersOffline   int    `json:"players_offline"`
			PlayersOnline    int    `json:"players_online"`
			World            string `json:"world"`
		} `json:"guild"`
	} `json:"guilds"`
	Information struct {
		APIVersion int    `json:"api_version"`
		Timestamp  string `json:"timestamp"`
	} `json:"information"`
}
