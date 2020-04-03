package model

import "database/sql"

// PersonSQL PersonSQL
type PersonSQL struct {
	ID   sql.NullInt64
	UUID sql.NullString
	Name sql.NullString
	Hero sql.NullBool
}

// GetModelJSON GetModelJSON
func (p *PersonSQL) GetModelJSON() Person {
	return Person{
		ID:   int(p.ID.Int64),
		UUID: p.UUID.String,
		Name: p.Name.String,
		Hero: p.Hero.Bool,
	}
}

// PowerstatsSQL PowerstatsSQL
type PowerstatsSQL struct {
	ID           sql.NullInt64
	Intelligence sql.NullString
	Strength     sql.NullString
	Speed        sql.NullString
	Durability   sql.NullString
	Power        sql.NullString
	Combat       sql.NullString
}

// GetModelJSON GetModelJSON
func (p *PowerstatsSQL) GetModelJSON() Powerstats {
	return Powerstats{
		ID:           int(p.ID.Int64),
		Intelligence: p.Intelligence.String,
		Strength:     p.Strength.String,
		Speed:        p.Speed.String,
		Durability:   p.Durability.String,
		Power:        p.Power.String,
		Combat:       p.Combat.String,
	}
}

// BiographySQL BiographySQL
type BiographySQL struct {
	ID              sql.NullInt64
	FullName        sql.NullString
	AlterEgos       sql.NullString
	PlaceOfBirth    sql.NullString
	FirstAppearance sql.NullString
	Publisher       sql.NullString
	Alignment       sql.NullString
}

// GetModelJSON GetModelJSON
func (p *BiographySQL) GetModelJSON() Biography {
	return Biography{
		ID:              int(p.ID.Int64),
		FullName:        p.FullName.String,
		AlterEgos:       p.AlterEgos.String,
		PlaceOfBirth:    p.PlaceOfBirth.String,
		FirstAppearance: p.FirstAppearance.String,
		Publisher:       p.Publisher.String,
		Alignment:       p.Alignment.String,
	}
}

// AppearanceSQL AppearanceSQL
type AppearanceSQL struct {
	ID        sql.NullInt64
	Gender    sql.NullString
	Race      sql.NullString
	Height    sql.NullString
	Weight    sql.NullString
	EyeColor  sql.NullString
	HairColor sql.NullString
}

// GetModelJSON GetModelJSON
func (p *AppearanceSQL) GetModelJSON() Appearance {
	return Appearance{
		ID:        int(p.ID.Int64),
		Gender:    p.Gender.String,
		Race:      p.Race.String,
		Height:    p.Height.String,
		Weight:    p.Weight.String,
		EyeColor:  p.EyeColor.String,
		HairColor: p.HairColor.String,
	}
}

// WorkSQL WorkSQL
type WorkSQL struct {
	ID         sql.NullInt64
	Occupation sql.NullString
	Base       sql.NullString
}

// GetModelJSON GetModelJSON
func (p *WorkSQL) GetModelJSON() Work {
	return Work{
		ID:         int(p.ID.Int64),
		Occupation: p.Occupation.String,
		Base:       p.Base.String,
	}
}

// ConnectionsSQL ConnectionsSQL
type ConnectionsSQL struct {
	ID               sql.NullInt64
	GroupAffiliation sql.NullString
	Relatives        sql.NullString
}

// GetModelJSON GetModelJSON
func (p *ConnectionsSQL) GetModelJSON() Connections {
	return Connections{
		ID:               int(p.ID.Int64),
		GroupAffiliation: p.GroupAffiliation.String,
		Relatives:        p.Relatives.String,
	}
}

// ImageSQL ImageSQL
type ImageSQL struct {
	ID  sql.NullInt64
	URL sql.NullString
}

// GetModelJSON GetModelJSON
func (p *ImageSQL) GetModelJSON() Image {
	return Image{
		ID:  int(p.ID.Int64),
		URL: p.URL.String,
	}
}
