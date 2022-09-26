package repository

import (
	"fmt"
	"notif-engine/model"
)

func (sqlNews *mysqlNewsletterRepository) GetAllNewsletter() (result []model.GetAllNewsletter, err error) {

	query := `SELECT email FROM newsletters`

	rows, err := sqlNews.Conn.DB.Query(query)
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}

	for rows.Next() {
		var newsletter model.GetAllNewsletter
		if rows.Scan(&newsletter.Email) == nil {
			err := rows.Scan(&newsletter.Email)
			if err != nil {
				fmt.Println("error 1", err)
			}
			result = append(result, newsletter)
		}
	}
	return result, nil
}
