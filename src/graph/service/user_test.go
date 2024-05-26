package service

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMock() (*gorm.DB, *sql.DB, sqlmock.Sqlmock, error) {
	// sqlmockを初期化
	mockDB, mock, err := sqlmock.New()

	if err != nil {
		return nil, mockDB, mock, err
	}

	// DBのコネクションを初期化
	gormDB, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn:                      mockDB,
			SkipInitializeWithVersion: true,
		}),
		&gorm.Config{},
	)

	return gormDB, mockDB, mock, err
}

func TestGetAllUsers(t *testing.T) {
	gormDB, mockDB, mock, err := InitDBMock()
	assert.NoError(t, err)
	defer mockDB.Close()

	connectedDB, err := gormDB.DB()
	assert.NoError(t, err)
	defer connectedDB.Close()

	// モックデータの準備
	rows := mock.NewRows([]string{"id", "name", "email"}).
		AddRow(uuid.New().String(), "John Doe", "john@example.com").
		AddRow(uuid.New().String(), "Jane Doe", "jane@example.com")

	// クエリの期待値を設定
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).WillReturnRows(rows)

	// テスト対象の関数を呼び出し
	users, err := GetAllUsers(gormDB)

	// 結果の検証
	assert.NoError(t, err)
	assert.Len(t, users, 2)

	assert.Equal(t, "John Doe", users[0].Name)
	assert.Equal(t, "Jane Doe", users[1].Name)

	// モックが期待されたクエリを全て実行したかを検証
	assert.NoError(t, mock.ExpectationsWereMet())
}

// FIXME: テストをパスさせる
func TestGetUser(t *testing.T) {
	gormDB, mockDB, mock, err := InitDBMock()
	assert.NoError(t, err)
	defer mockDB.Close()

	connectedDB, err := gormDB.DB()
	assert.NoError(t, err)
	defer connectedDB.Close()

	id := uuid.New().String()

	// クエリの期待値を設定
	mock.ExpectQuery("SELECT").WithArgs(id).WillReturnRows(
		// モックデータの準備
		mock.NewRows([]string{"id"}).AddRow(id),
	)

	// テスト対象の関数を呼び出し
	user, err := GetUser(gormDB, id)

	// 結果の検証
	assert.NoError(t, err)
	assert.NotNil(t, user)

	assert.Equal(t, id, user.ID)

	// モックが期待されたクエリを全て実行したかを検証
	assert.NoError(t, mock.ExpectationsWereMet())
}
