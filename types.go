package teleBot

type User struct {
	Id           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

func (user User) String() string {
	return StructToString(user)
}

func (user User) FullName() string {
	if user.LastName == "" {
		return user.FirstName
	}
	return user.FirstName + " " + user.LastName
}

// ------------------------------------------------------------------------------------

type Message struct {
	Id   int    `json:"message_id"`
	From User   `json:"from"`
	Date int    `json:"date"`
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

func (msg Message) String() string {
	return StructToString(msg)
}

// ------------------------------------------------------------------------------------
type Chat struct {
	Id       int    `json:"id"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	Username string `json:"username"`
}

func (chat Chat) String() string {
	return StructToString(chat)
}

// ------------------------------------------------------------------------------------

type Update struct {
	Id      int     `json:"update_id"`
	Message Message `json:"message"`
	// EditedMessage Message `json:"edited_message"`
}

func (upd Update) String() string {
	return StructToString(upd)
}

// ------------------------------------------------------------------------------------

type ResponseUpdate struct {
	Ok  bool    `json:"ok"`
	Res Message `json:"result"`
}

func (upd ResponseUpdate) String() string {
	return StructToString(upd)
}
