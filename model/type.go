package model

// Person Person
type Person struct {
	ID          int         `json:"id"`
	Name        string      `json:"name" form:"name" binding:"required"`
	Hero        bool        `json:"hero"`
	Powerstats  Powerstats  `json:"powerstats"`
	Biography   Biography   `json:"biography"`
	Appearance  Appearance  `json:"appearance"`
	Work        Work        `json:"work"`
	Connections Connections `json:"connections"`
	Image       Image       `json:"image"`
}

// Powerstats Powerstats
type Powerstats struct {
	ID           int    `json:"id"`
	Intelligence string `json:"intelligence"`
	Strength     string `json:"strength"`
	Speed        string `json:"speed"`
	Durability   string `json:"durability"`
	Power        string `json:"power"`
	Combat       string `json:"combat"`
}

// Biography Biography
type Biography struct {
	ID              int    `json:"id"`
	FullName        string `json:"full-name"`
	AlterEgos       string `json:"alter-egos"`
	PlaceOfBirth    string `json:"place-of-birth"`
	FirstAppearance string `json:"first-appearance"`
	Publisher       string `json:"publisher"`
	Alignment       string `json:"alignment"`
}

// Appearance Appearance
type Appearance struct {
	ID        int    `json:"id"`
	Gender    string `json:"gender"`
	Race      string `json:"race"`
	Height    string `json:"height"`
	Weight    string `json:"weight"`
	EyeColor  string `json:"eye-color"`
	HairColor string `json:"hair-color"`
}

// Work Work
type Work struct {
	ID         int    `json:"id"`
	Occupation string `json:"occupation"`
	Base       string `json:"base"`
}

// Connections Connections
type Connections struct {
	ID               int    `json:"id"`
	GroupAffiliation string `json:"group-affiliation"`
	Relatives        string `json:"relatives"`
}

// Image Image
type Image struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}
