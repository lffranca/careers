package model

import "database/sql"

// PersonSQL PersonSQL
type PersonSQL struct {
	ID   sql.NullInt64
	Name sql.NullString
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

// WorkSQL WorkSQL
type WorkSQL struct {
	ID         sql.NullInt64
	Occupation sql.NullString
	Base       sql.NullString
}

// ConnectionsSQL ConnectionsSQL
type ConnectionsSQL struct {
	ID               sql.NullInt64
	GroupAffiliation sql.NullString
	Relatives        sql.NullString
}

// ImageSQL ImageSQL
type ImageSQL struct {
	ID  sql.NullInt64
	URL sql.NullString
}
