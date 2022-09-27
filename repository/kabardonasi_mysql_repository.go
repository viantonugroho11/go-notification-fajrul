package repository

import (
	"context"
	"fmt"
	"notif-engine/model"
)

func(s *mysqlKabarDonasiRepository) GetAllUserByDonasiID(ctx context.Context,id string) (result []model.GetEmailUserKabarDonasi, err error) {
	query := `SELECT email FROM transaksi_donasi WHERE donasi_id = ?`

	err = s.Conn.DB.QueryRow(query,id).Scan(&result)
	// err = s.Conn.DB.QueryRow(query).Scan(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}


func (s *mysqlKabarDonasiRepository) GetUserStatusNotyetByDonasiID(id string) (result []model.GetEmailUserKabarDonasi,err error){
	query := `SELECT email FROM transaksi_donasis WHERE status = '1' and donasi_id = ?`

	rows,err := s.Conn.DB.Query(query,id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var kabarDonasi model.GetEmailUserKabarDonasi
		if rows.Scan(&kabarDonasi.Email) == nil {
			err := rows.Scan(&kabarDonasi.Email)
			if err != nil {
				fmt.Println("error 1", err)
			}
			result = append(result, kabarDonasi)
		}
	}

	return result, nil
}