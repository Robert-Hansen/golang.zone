package repositories

import (
	"github.com/steffen25/golang.zone/database"
	"github.com/steffen25/golang.zone/models"
	"log"
)

type TestPostRepository interface {
	GetAllPosts() ([]*models.PostTest, error)
}

type testPostRepository struct {
	*database.GormMySQLDB
}

/**
 *  Errors
 */
type NoPostsError struct{}

func (*NoPostsError) Error() string {
	return "no records were found"
}

func NewTestPostRepository(db *database.GormMySQLDB) TestPostRepository {
	return &testPostRepository{db}
}

func (pr *testPostRepository) GetAllPosts() ([]*models.PostTest, error) {
	var posts []*models.PostTest

	if err := pr.DB.Find(&posts).Error; err != nil {
		// error handling
		log.Println(err)
		return nil, err
	}

	return posts, nil
}
