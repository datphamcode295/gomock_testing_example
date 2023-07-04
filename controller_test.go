package controller_test

import (
	"errors"
	"testing"

	controller "example.com/testing"
	mock_controller "example.com/testing/mocks"
	"example.com/testing/models"
	"gorm.io/gorm"

	"github.com/golang/mock/gomock"
)

func TestLogin(t *testing.T) {

	type args struct {
		email string
		pass  string
	}

	tests := []struct {
		name                string
		args                args
		want                *models.User
		wantRepositoryError error
		wantErr             error
	}{
		{
			name: "Login success",
			args: args{
				email: "test@gmail.com",
				pass:  "123456",
			},
			want: &models.User{
				ID:   1,
				Name: "John Doe",
			},
			wantRepositoryError: nil,
			wantErr:             nil,
		},
		{
			name: "Login failed",
			args: args{
				email: "test1@gmail.com",
				pass:  "123456",
			},
			want:                nil,
			wantRepositoryError: gorm.ErrRecordNotFound,
			wantErr:             errors.New("invalid email or password"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockUserRepository := mock_controller.NewMockUserRepository(mockCtrl)
			testAuthController := &controller.Auth{UserRepository: mockUserRepository}

			defer mockCtrl.Finish()

			mockUserRepository.EXPECT().GetUser(tt.args.email, tt.args.pass).Return(tt.want, tt.wantRepositoryError).Times(1)

			user, err := testAuthController.Login(tt.args.email, tt.args.pass)
			if tt.wantErr == nil {
				if err != nil {
					t.Fail()
				}

				if user.ID != tt.want.ID || user.Name != tt.want.Name {
					t.Fail()
				}
			} else {
				if err != nil && err.Error() != tt.wantErr.Error() {
					t.Fail()
				}

				if user != nil {
					t.Fail()
				}
			}
		})
	}
}
