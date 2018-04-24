package teleBot

type GetMeResponse struct {
	Ok  bool `json:"ok"`
	Res User `json:"result"`
}

func (gmr GetMeResponse) String() string {
	return StructToString(gmr)
}

// ------------------------------------------------------------------------------------

type GetUpdatesResponse struct {
	Ok  bool     `json:"ok"`
	Res []Update `json:"result"`
}

func (udr GetUpdatesResponse) String() string {
	return StructToString(udr)
}

// ------------------------------------------------------------------------------------
