package auth

import "testing"

func TestCheckEmailValidation(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{
			name:    "valid email",
			email:   "test@example.com",
			wantErr: false,
		},
		{
			name:    "missing @",
			email:   "testexample.com",
			wantErr: true,
		},
		{
			name:    "missing dot after @",
			email:   "test@examplecom",
			wantErr: true,
		},
		{
			name:    "missing @ and dot",
			email:   "testexamplecom",
			wantErr: true,
		},
		{
			name:    "dot before @ only",
			email:   "test.example@com",
			wantErr: true,
		},
		{
			name:    "valid email with subdomain",
			email:   "test@mail.example.com",
			wantErr: false,
		},
		{
			name:    "multiple @ symbols",
			email:   "test@@example.com",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckEmailValidation(tt.email)

			if (err != nil) != tt.wantErr {
				t.Errorf(
					"CheckEmailValidation(%q) error = %v, wantErr %v",
					tt.email,
					err,
					tt.wantErr,
				)
			}
		})
	}
}
