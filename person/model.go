package person

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/lffranca/careers/model"
)

func getPersonByID(db *sql.DB, id int) (*model.Person, error) {
	query := `
		select
			p.id,
			p.name,
			p.hero,
			p2.id as powerstats_id,
			p2.intelligence,
			p2.strength,
			p2.speed,
			p2.durability,
			p2.power,
			p2.combat,
			b.id as biography_id,
			b.full_name,
			b.alter_egos,
			b.place_of_birth,
			b.first_appearance,
			b.publisher,
			b.alignment,
			a.id as appearance_id,
			a.gender,
			a.race,
			a.height,
			a.weight,
			a.eye_color,
			a.hair_color,
			w.id as work_id,
			w.occupation,
			w.base,
			c.id as connections_id,
			c.group_affiliation,
			c.relatives,
			i.id as image_id,
			i.url
		from person p
		left join powerstats p2 on p.id = p2.person_id
		left join biography b on p.id = b.person_id
		left join appearance a on p.id = a.person_id
		left join work w on p.id = w.person_id
		left join connections c on p.id = c.person_id
		left join image i on p.id = i.person_id
		where p.id = $1
		limit 1
		;
	`

	rowsPerson, errRowsPerson := db.QueryContext(context.Background(), query, id)
	if errRowsPerson != nil {
		return nil, errRowsPerson
	}

	person := model.Person{}
	for rowsPerson.Next() {
		personSQL := model.PersonSQL{}
		powersSQL := model.PowerstatsSQL{}
		bioSQL := model.BiographySQL{}
		appeSQL := model.AppearanceSQL{}
		workSQL := model.WorkSQL{}
		connSQL := model.ConnectionsSQL{}
		imageSQL := model.ImageSQL{}

		if err := rowsPerson.Scan(
			&personSQL.ID,
			&personSQL.Name,
			&personSQL.Hero,

			&powersSQL.ID,
			&powersSQL.Intelligence,
			&powersSQL.Strength,
			&powersSQL.Speed,
			&powersSQL.Durability,
			&powersSQL.Power,
			&powersSQL.Combat,

			&bioSQL.ID,
			&bioSQL.FullName,
			&bioSQL.AlterEgos,
			&bioSQL.PlaceOfBirth,
			&bioSQL.FirstAppearance,
			&bioSQL.Publisher,
			&bioSQL.Alignment,

			&appeSQL.ID,
			&appeSQL.Gender,
			&appeSQL.Race,
			&appeSQL.Height,
			&appeSQL.Weight,
			&appeSQL.EyeColor,
			&appeSQL.HairColor,

			&workSQL.ID,
			&workSQL.Occupation,
			&workSQL.Base,

			&connSQL.ID,
			&connSQL.GroupAffiliation,
			&connSQL.Relatives,

			&imageSQL.ID,
			&imageSQL.URL,
		); err != nil {
			log.Println(err)
			continue
		}

		person = personSQL.GetModelJSON()
		person.Powerstats = powersSQL.GetModelJSON()
		person.Biography = bioSQL.GetModelJSON()
		person.Appearance = appeSQL.GetModelJSON()
		person.Work = workSQL.GetModelJSON()
		person.Connections = connSQL.GetModelJSON()
		person.Image = imageSQL.GetModelJSON()
	}

	if (model.Person{}) == person {
		return nil, errors.New("Invalid param id")
	}

	return &person, nil
}
