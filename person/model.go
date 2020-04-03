package person

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/lffranca/careers/model"
)

func createPersonPowerstats(db *sql.DB, person *model.Person) (*model.Person, error) {
	query := `
		insert into powerstats (intelligence, strength, speed, durability, power, combat, person_id)
		values (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		)
		returning id
		;
	`

	rows, errRows := db.QueryContext(
		context.Background(),
		query,
		person.Powerstats.Intelligence,
		person.Powerstats.Strength,
		person.Powerstats.Speed,
		person.Powerstats.Durability,
		person.Powerstats.Power,
		person.Powerstats.Combat,
		person.ID,
	)
	if errRows != nil {
		return nil, errRows
	}

	itemID := 0
	for rows.Next() {
		var idSQL sql.NullInt64
		if err := rows.Scan(&idSQL); err != nil {
			log.Println(err)
			continue
		}

		if idSQL.Valid {
			itemID = int(idSQL.Int64)
		}
	}

	if itemID == 0 {
		return nil, errors.New("Error save")
	}

	person.Powerstats.ID = itemID

	return person, nil
}

func createPersonBiography(db *sql.DB, person *model.Person) (*model.Person, error) {
	query := `
		insert into biography (full_name, alter_egos, place_of_birth, first_appearance, publisher, alignment, person_id)
		values (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		)
		returning id
		;
	`

	rows, errRows := db.QueryContext(
		context.Background(),
		query,
		person.Biography.FullName,
		person.Biography.AlterEgos,
		person.Biography.PlaceOfBirth,
		person.Biography.FirstAppearance,
		person.Biography.Publisher,
		person.Biography.Alignment,
		person.ID,
	)
	if errRows != nil {
		return nil, errRows
	}

	itemID := 0
	for rows.Next() {
		var idSQL sql.NullInt64
		if err := rows.Scan(&idSQL); err != nil {
			log.Println(err)
			continue
		}

		if idSQL.Valid {
			itemID = int(idSQL.Int64)
		}
	}

	if itemID == 0 {
		return nil, errors.New("Error save")
	}

	person.Biography.ID = itemID

	return person, nil
}

func createPersonAppearance(db *sql.DB, person *model.Person) (*model.Person, error) {
	query := `
		insert into appearance (gender, race, height, weight, eye_color, hair_color, person_id)
		values (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		)
		returning id
		;
	`

	rows, errRows := db.QueryContext(
		context.Background(),
		query,
		person.Appearance.Gender,
		person.Appearance.Race,
		person.Appearance.Height,
		person.Appearance.Weight,
		person.Appearance.EyeColor,
		person.Appearance.HairColor,
		person.ID,
	)
	if errRows != nil {
		return nil, errRows
	}

	itemID := 0
	for rows.Next() {
		var idSQL sql.NullInt64
		if err := rows.Scan(&idSQL); err != nil {
			log.Println(err)
			continue
		}

		if idSQL.Valid {
			itemID = int(idSQL.Int64)
		}
	}

	if itemID == 0 {
		return nil, errors.New("Error save")
	}

	person.Appearance.ID = itemID

	return person, nil
}

func createPersonWork(db *sql.DB, person *model.Person) (*model.Person, error) {
	query := `
		insert into work (occupation, base, person_id)
		values (
			$1,
			$2,
			$3
		)
		returning id
		;
	`

	rows, errRows := db.QueryContext(
		context.Background(),
		query,
		person.Work.Occupation,
		person.Work.Base,
		person.ID,
	)
	if errRows != nil {
		return nil, errRows
	}

	itemID := 0
	for rows.Next() {
		var idSQL sql.NullInt64
		if err := rows.Scan(&idSQL); err != nil {
			log.Println(err)
			continue
		}

		if idSQL.Valid {
			itemID = int(idSQL.Int64)
		}
	}

	if itemID == 0 {
		return nil, errors.New("Error save")
	}

	person.Work.ID = itemID

	return person, nil
}

func createPersonConnections(db *sql.DB, person *model.Person) (*model.Person, error) {
	query := `
		insert into connections (group_affiliation, relatives, person_id)
		values (
			$1,
			$2,
			$3
		)
		returning id
		;
	`

	rows, errRows := db.QueryContext(
		context.Background(),
		query,
		person.Connections.GroupAffiliation,
		person.Connections.Relatives,
		person.ID,
	)
	if errRows != nil {
		return nil, errRows
	}

	itemID := 0
	for rows.Next() {
		var idSQL sql.NullInt64
		if err := rows.Scan(&idSQL); err != nil {
			log.Println(err)
			continue
		}

		if idSQL.Valid {
			itemID = int(idSQL.Int64)
		}
	}

	if itemID == 0 {
		return nil, errors.New("Error save")
	}

	person.Connections.ID = itemID

	return person, nil
}

func createPersonImage(db *sql.DB, person *model.Person) (*model.Person, error) {
	query := `
		insert into image (url, person_id)
		values (
			$1,
			$2
		)
		returning id
		;
	`

	rows, errRows := db.QueryContext(
		context.Background(),
		query,
		person.Image.URL,
		person.ID,
	)
	if errRows != nil {
		return nil, errRows
	}

	itemID := 0
	for rows.Next() {
		var idSQL sql.NullInt64
		if err := rows.Scan(&idSQL); err != nil {
			log.Println(err)
			continue
		}

		if idSQL.Valid {
			itemID = int(idSQL.Int64)
		}
	}

	if itemID == 0 {
		return nil, errors.New("Error save")
	}

	person.Image.ID = itemID

	return person, nil
}

func createPerson(db *sql.DB, person *model.Person) (*model.Person, error) {
	queryPerson := `
		insert into person (uuid, name, hero)
		values (
			$1,
			$2,
			$3
		)
		returning id
		;
	`

	puuid, errNewUUID := uuid.NewUUID()
	if errNewUUID != nil {
		return nil, errNewUUID
	}

	person.UUID = puuid.String()

	rowsPerson, errRowsPerson := db.QueryContext(
		context.Background(),
		queryPerson,
		person.UUID,
		person.Name,
		person.Hero,
	)
	if errRowsPerson != nil {
		return nil, errRowsPerson
	}

	personID := 0
	for rowsPerson.Next() {
		var idSQL sql.NullInt64
		if err := rowsPerson.Scan(&idSQL); err != nil {
			log.Println(err)
			continue
		}

		if idSQL.Valid {
			personID = int(idSQL.Int64)
		}
	}

	if personID == 0 {
		return nil, errors.New("Error save")
	}

	person.ID = personID

	var err error
	person, err = createPersonPowerstats(db, person)
	if err != nil {
		return nil, err
	}

	person, err = createPersonBiography(db, person)
	if err != nil {
		return nil, err
	}

	person, err = createPersonAppearance(db, person)
	if err != nil {
		return nil, err
	}

	person, err = createPersonWork(db, person)
	if err != nil {
		return nil, err
	}

	person, err = createPersonConnections(db, person)
	if err != nil {
		return nil, err
	}

	person, err = createPersonImage(db, person)
	if err != nil {
		return nil, err
	}

	return person, nil
}

func getPersonByID(db *sql.DB, id int) (*model.Person, error) {
	query := `
		select
			p.id,
			p.uuid,
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
			&personSQL.UUID,
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
