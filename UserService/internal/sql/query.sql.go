package sql


const (
	CreateUserQuery = `INSERT INTO users (id, first_name, last_name, email, password, role) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id, first_name, last_name, email, password,role`

	FindByEmailQuery = `SELECT id, email, first_name, last_name, role FROM users WHERE email = $1`

	FindByIDQuery = `SELECT id, email, first_name, last_name, role  FROM users WHERE user_id = $1`
)
