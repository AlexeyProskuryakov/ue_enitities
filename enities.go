package ue_entities

type (
	Identity string
	UserId Identity
	PostId Identity
	GroupId Identity

	Reference struct {
		From Identity          `json:"from"`
		To   Identity          `json:"to"`
		Type string            `json:"type"`
		Data map[string]string `json:"data"`
	}

	Group struct {
		Id   GroupId           `json:"id"`
		Name string            `json:"name"`
		Info map[string]string `json:"info"`

		Members []UserId `json:"members"`
		Posts   []PostId `json:"posts"`
	}

	User struct {
		Id   UserId            `json:"id"`
		Name string            `json:"name"`
		Info map[string]string `json:"info"`

		Friends       []UserId  `json:"friends"`
		Followers     []UserId  `json:"followers"`
		Subscriptions []GroupId `json:"subscriptions"`

		Posts []PostId `json:"posts"`
		Likes []PostId `json:"likes"`
	}

	Likable struct {
		Likers []Identity `json:"likers"`
	}

	Post struct {
		Likable

		Id     PostId `json:"id"`
		Author *User  `json:"author"`
		Title  string `json:"title"`
		Text   string `json:"text"`
		Url    string `json:"url"`
		Type   string `json:"type"`

		Tags []string `json:"tags"`
	}
)
