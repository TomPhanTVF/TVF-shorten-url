package postgres
import(
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	models "user-service/internal/models"
	"user-service/internal/sql"
	"context"
)

// User repository
type UserRepository struct {
	db *sqlx.DB
}


// User repository constructor
func NewUserPGRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}


// Create new user
func (r *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {

	createdUser := &models.User{}
	if err := r.db.QueryRowxContext(
		ctx,
		sql.CreateUserQuery,
		user.GenID(),
		user.GetFirstName(),
		user.GetLastName(),
		user.GetEmail(),
		user.GetPassword(),
		user.GetRole(),
	).StructScan(createdUser); err != nil {
		return nil, errors.Wrap(err, "Create.QueryRowxContext")
	}

	return createdUser, nil
}

// Find by user email address
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	if err := r.db.GetContext(ctx, user, sql.FindByEmailQuery, email); err != nil {
		return nil, errors.Wrap(err, "FindByEmail.GetContext")
	}

	return user, nil
}

// Find user by uuid
func (r *UserRepository) FindById(ctx context.Context, userID string) (*models.User, error) {
	user := &models.User{}
	if err := r.db.GetContext(ctx, user, sql.FindByIDQuery, userID); err != nil {
		return nil, errors.Wrap(err, "FindById.GetContext")
	}

	return user, nil
}
