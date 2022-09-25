package repository

import "notif-engine/model"



func (sqlNews *mysqlNewsletterRepository) GetAllNewsletter() (result []model.GetAllNewsletter, err error) {

	query := `SELECT email FROM newsletter`

	err = sqlNews.Conn.DB.QueryRow(query).Scan(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
