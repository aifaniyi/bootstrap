package main

type schema struct {
	Lang     string   `json:"lang"`
	Auth     []string `json:"auth"`
	Entities []entity `json:"entities"`
}

type entity struct {
	Name        string     `json:"name"`
	Properties  []property `json:"properties"`
	Description string     `json:"description"`
	ReadBy      []string   `json:"readBy"`
	Paginate    *paginate  `json:"paginate"`
	Relations   []relation `json:"relations"`
}

type paginate struct {
	Size int `json:"size"`
}

type relation struct {
	Type   string `json:"type"`
	Entity string `json:"entity"`
}

type fieldType struct {
	Name string `json:"name"`
}

type property struct {
	Name        string    `json:"name"`
	Type        fieldType `json:"type"`
	Width       *int      `json:"width,omitempty"`
	Nullable    bool      `json:"nullable,omitempty"`
	Dto         bool      `json:"dto,omitempty"`
	Indexable   bool      `json:"indexable,omitempty"`
	Description string    `json:"description"`
	Unique      bool      `json:"unique"`
}
