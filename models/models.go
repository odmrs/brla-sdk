package models

type Account struct {
	Email           string  `json:"email"`
	Password        string  `json:"password"`
	ConfirmPassword string  `json:"confirmPassword"`
	Phone           string  `json:"phone"`
	TaxIdType       string  `json:"taxIdType"`
	FullName        string  `json:"fullName"`
	CPF             string  `json:"cpf"`
	BirthDate       string  `json:"birthDate"`
	CNPJ            string  `json:"cnpj,omitempty"`
	CompanyName     string  `json:"companyName,omitempty"`
	RegNum          string  `json:"regnum,omitempty"`
	StartData       string  `json:"startDate,omitempty"`
	Address         Address `json:"address"`
}

type Address struct {
	Cep        string `json:"cep"`
	City       string `json:"city"`
	State      string `json:"state"`
	Street     string `json:"street"`
	Number     string `json:"number"`
	District   string `json:"district"`
	Complement string `json:"complement"`
}

func NewAccount(email, password, confirmPassword, phone, taxIdType, fullName, cpf, birthDate string, address Address) *Account {
	return &Account{
		Email:           email,
		Password:        password,
		ConfirmPassword: confirmPassword,
		Phone:           phone,
		TaxIdType:       taxIdType,
		CPF:             cpf,
		FullName:        fullName,
		BirthDate:       birthDate,
		Address:         address,
	}
}
