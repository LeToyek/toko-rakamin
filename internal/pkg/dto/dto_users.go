package dto

type UserFilter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
}

type UserResponse struct {
	ID           int64  `json:"id"`
	Nama         string `json:"nama"`
	Email        string `json:"email"`
	NoTelp       string `json:"no_telp"`
	JenisKelamin string `json:"jenis_kelamin"`
	Pekerjaan    string `json:"pekerjaan"`
	IsAdmin      bool   `json:"is_admin"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

type UserRegister struct {
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=6"`
	Nama         string `json:"nama" validate:"required,min=3,max=255"`
	NoTelp       string `json:"no_telp" validate:"required,min=10"`
	JenisKelamin string `json:"jenis_kelamin"`
	Tentang      string `json:"tentang"`
	Pekerjaan    string `json:"pekerjaan"`
	IdProvinsi   string `json:"id_provinsi"`
	IdKota       string `json:"id_kota"`
	IsAdmin      bool   `json:"is_admin"`
}
