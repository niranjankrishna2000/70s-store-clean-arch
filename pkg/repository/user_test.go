package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_CheckUserAvailabilty(t *testing.T) {
	tests := []struct {
		name string
		args string
		stub func(sqlmock.Sqlmock)
		want bool
	}{
		{
			name: "User Available",
			args: "niranjan@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select count\(\*\) from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(5))

			},
			want: true,
		},
		{
			name: "user not available",
			args: "niranjan@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select count\(\*\) from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

			},

			want: false,
		},
		{
			name: "error from database",
			args: "niranjan@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select count\(\*\) from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).WithArgs("").
					WillReturnError(errors.New("text string"))

			},

			want: false,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result := u.CheckUserAvailability(tt.args)

			assert.Equal(t, tt.want, result)
		})
	}

}

func Test_UserBlockStatus(t *testing.T){
	tests:=[]struct{	
		name string
		args string
		stub func(sqlmock.Sqlmock)
		want bool
		wantErr error
	}{
		{
			name: "User is blocked",
			args: "niranjan@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock){
				expectedQuery := `^select permission from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"permission"}).AddRow(true))
			},
			want: true,
			wantErr: nil,
		},
		{
			name: "user is not blocked",
			args: "niranjan@gmail.com",
			stub: func (mockSQL sqlmock.Sqlmock)  {
				expectedQuery := `^select permission from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).WithArgs().WillReturnRows(sqlmock.NewRows([]string{"permission"}).AddRow(false))
			},
			want: false,
			wantErr: nil,
		},
	}

	for _,tt:=range tests{
		t.Run(tt.name,func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})
				tt.stub(mockSQL)
				u:=NewUserRepository(gormDB)
				result,err:=u.UserBlockStatus(tt.args)
				assert.Equal(t,result,tt.want)
				assert.Equal(t,tt.wantErr,err)
		})
	}
}