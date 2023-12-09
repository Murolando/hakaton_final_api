package postgres

import (
	"fmt"

	"github.com/Murolando/hakaton_final_api/ent"
	"github.com/jmoiron/sqlx"
)

type CardPostgres struct {
	db *sqlx.DB
}

func NewCardPostgres(db *sqlx.DB) *CardPostgres {
	return &CardPostgres{
		db: db,
	}
}

func (r *CardPostgres) AddCard(card *ent.Card, userId int) (bool, error) {
	query := fmt.Sprintf(`
	INSERT INTO "%s" (word,description)
	VALUES($1,$2)
	RETURNING id`, cardsTable)
	var id int
	row := r.db.QueryRow(query, card.Word,card.Description)
	if err:=row.Scan(&id);err!=nil{
		return false,err
	}

	query = fmt.Sprintf(`
	INSERT INTO "%s" (user_id,card_id)
	VALUES($1,$2)`, userCardsTable)
	_,err:= r.db.Exec(query, userId,id)
	if err!=nil{
		return false,err
	}
	return true, nil
}
func (r *CardPostgres) GetCard(userId int) ([]*ent.Card, error) {
	cards := make([]*ent.Card,0)
	query := fmt.Sprintf(`
	SELECT DISTINCT "cards".id,"cards".word,"cards".description
	FROM "%s" 
	JOIN "user_cards" ON user_id = $1`, cardsTable)
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next(){
		var c ent.Card
		rows.Scan(&c.Id,&c.Word,&c.Description)
		cards = append(cards, &c)
	}
	return cards, nil
}
func (r *CardPostgres) DeleteCard(cardId int) (bool, error) {
	query := fmt.Sprintf(`
	DELETE FROM "%s" 
	WHERE id = $1`,cardsTable)
	res1,err := r.db.Exec(query,cardId)
	if err!=nil{
		return false,err
	}
    count, err := res1.RowsAffected()
    if err != nil {
        return false,err
    }
	if count != 0{
		return true,nil
	}
	return false,nil
}
