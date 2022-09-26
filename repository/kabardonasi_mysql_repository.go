package repository

import (
	"context"
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


func (s *mysqlKabarDonasiRepository) GetUserStatusNotyetByDonasiID(ctx context.Context, id string) (result []model.GetEmailUserKabarDonasi,err error){
	query := `SELECT email FROM user_kabar_donasi WHERE status = 'notyet' and donasi_id = ?`

	err = s.Conn.DB.QueryRow(query).Scan(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}