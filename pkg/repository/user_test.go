package repository

import (
	"errors"
	"log"
	"reflect"
	"testing"

	"main/pkg/utils/models"

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

func Test_UserBlockStatus(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		stub    func(sqlmock.Sqlmock)
		want    bool
		wantErr error
	}{
		{
			name: "User is blocked",
			args: "niranjan@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^select permission from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"permission"}).AddRow(true))
			},
			want:    true,
			wantErr: nil,
		},
		{
			name: "user is not blocked",
			args: "niranjan@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^select permission from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).WithArgs().WillReturnRows(sqlmock.NewRows([]string{"permission"}).AddRow(false))
			},
			want:    false,
			wantErr: nil,
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
			result, err := u.UserBlockStatus(tt.args)
			assert.Equal(t, result, tt.want)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_SignUp(t *testing.T) {

	type args struct {
		input models.UserDetails
	}

	tests := []struct {
		name    string
		args    args
		stub    func(sqlmock.Sqlmock)
		want    models.UserResponse
		wantErr error
	}{
		{
			name: "success signup user",
			args: args{
				input: models.UserDetails{Name: "UserNew", Email: "usernew@gmail.com", Username: "usernew", Phone: "1234567890", Password: "userpass", ConfirmPassword: "userpass"},
			},

			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^INSERT INTO users (.+)$`
				mockSQL.ExpectQuery(expectedQuery).WithArgs("UserNew", "usernew@gmail.com", "userpass", "1234567890", "usernew").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone"}).AddRow(1, "UserNew", "usernew@gmail.com", "1234567890"))

			},

			want:    models.UserResponse{Id: 1, Name: "UserNew", Email: "usernew@gmail.com", Phone: "1234567890"},
			wantErr: nil,
		},

		{
			name: "error signup user",
			args: args{
				input: models.UserDetails{Name: "", Email: "", Username: "", Phone: "", Password: "", ConfirmPassword: ""},
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^INSERT INTO users (.+)$`
				mockSQL.ExpectQuery(expectedQuery).WithArgs("", "", "", "", "").
					WillReturnError(errors.New("Query 'INSERT INTO users (name, email,phone,password,username) VALUES ($1, $2, $3, $4,$5) RETURNING id, name, email, phone', arguments do not match: argument 4 expected [string - ] does not match actual [string - 12345]"))

			},

			want:    models.UserResponse{},
			wantErr: errors.New("Query 'INSERT INTO users (name, email,phone,password,username) VALUES ($1, $2, $3, $4,$5) RETURNING id, name, email, phone', arguments do not match: argument 4 expected [string - ] does not match actual [string - 12345]"),
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

			got, err := u.SignUp(tt.args.input)
			log.Println("log :", tt.name, got, tt.want)
			log.Println("log :", tt.name, tt.wantErr, err)
			assert.Equal(t, tt.wantErr, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.UserSignUp() = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_FindUserByEmail(t *testing.T) {
	tests := []struct {
		name    string
		args    models.UserLogin
		stub    func(sqlmock.Sqlmock)
		want    models.UserResponse
		wantErr error
	}{
		{
			name: "Success",
			args: models.UserLogin{
				Email:    "niranjan@gmail.com",
				Password: "niranjan2000",
			},
			stub: func(s sqlmock.Sqlmock) {
				expectedQuery := `^SELECT .* FROM users (.+)$`
				s.ExpectQuery(expectedQuery).WithArgs("niranjan@gmail.com").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "username", "email", "phone", "password"}).AddRow(1, "Niranjan", "niranjan", "niranjan@gmail.com", "8593008099", "niranjan2000"))
			},
			want: models.UserResponse{Id: 1,
				Name:     "Niranjan",
				Email:    "niranjan@gmail.com",
				Username: "niranjan",
				Phone:    "8593008099",
				Password: "niranjan2000",
			},
			wantErr: nil,
		},
		{
			name: "failure",
			args: models.UserLogin{
				Email:    "niranjan@gmail.com",
				Password: "niranjan2000",
			},
			stub: func(s sqlmock.Sqlmock) {
				expectedQuery := `^SELECT \* FROM users(.+)$`

				s.ExpectQuery(expectedQuery).WithArgs("niranjan@gmail.com").
					WillReturnError(errors.New("new error"))
			},
			want:    models.UserResponse{},
			wantErr: errors.New("error checking user details"),
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
			got, err := u.FindUserByEmail(tt.args)

			log.Println("log :", tt.name, tt.want, got)
			log.Println("log :", tt.name, tt.wantErr, err)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, err, tt.wantErr)
		})

	}

}

func Test_FindUserIDByOrderID(t *testing.T) {
	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		want    int
		wantErr error
	}{
		{
			name: "Success",
			args: 4,
			stub: func(s sqlmock.Sqlmock) {
				expectedQuery := `^SELECT user_id FROM orders (.+)$`
				s.ExpectQuery(expectedQuery).WithArgs(4).WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(1))
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "failure",
			args: 5,
			stub: func(s sqlmock.Sqlmock) {
				expectedQuery := `^SELECT user_id FROM orders(.+)$`

				s.ExpectQuery(expectedQuery).WithArgs(5).
					WillReturnError(errors.New("error checking user details"))
			},
			want:    0,
			wantErr: errors.New("error checking user details"),
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
			got, err := u.FindUserIDByOrderID(tt.args)

			log.Println("log :", tt.name, tt.want, got)
			log.Println("log :", tt.name, tt.wantErr, err)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, err, tt.wantErr)
		})

	}

}

func Test_AddAddress(t *testing.T) {

	tests := []struct {
		name string
		args struct {
			id      int
			address models.AddAddress
			result  bool
		}
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "success",
			args: struct {
				id      int
				address models.AddAddress
				result  bool
			}{
				id: 1,
				address: models.AddAddress{
					Name:      "Niranjan",
					HouseName: "thundathil",
					Street:    "mundathicode",
					City:      "thrissur",
					State:     "kerala",
					Pin:       "680601",
				},
				result: true,
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("INSERT INTO addresses").WithArgs(1, "Niranjan", "thundathil", "mundathicode", "thrissur", "kerala", "680601", true).WillReturnResult(sqlmock.NewResult(1,1))

			},
			wantErr: nil,
		},
		{
			name: "error",
			args: struct {
				id      int
				address models.AddAddress
				result  bool
			}{
				id: 1,
				address: models.AddAddress{
					Name:      "Niranjan",
					HouseName: "thundathil",
					Street:    "mundathicode",
					City:      "thrissur",
					State:     "kerala",
					Pin:       "680601",
				},
				result: true,
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("INSERT INTO addresses").WithArgs(1, "Niranjan", "thundathil", "mundathicode", "thrissur", "kerala", "680601", true).WillReturnError(errors.New("could not add address"))

			},
			wantErr: errors.New("could not add address"),
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

			err := u.AddAddress(1, tt.args.address, true)

			assert.Equal(t, tt.wantErr, err)
		})
	}

}
