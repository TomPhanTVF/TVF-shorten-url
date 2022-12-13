package sql

const (
	CreateURLQuery = `INSERT INTO urls (id, redirect, TVF, random, user_id) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id, redirect, TVF, random, user_id`

	FindByOwnerIDQuery = `SELECT * FROM urls WHERE user_id = $1`

	FindByIDQuery = `SELECT *  FROM urls WHERE id = $1`
)