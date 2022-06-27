package domain

type Seller struct {
	Id          int64  `json:"id"`
	Cid         int64  `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

type ServiceSeller interface {
	GetAll() ([]Seller, error)
	GetById(id int64) (Seller, error)
	Create(cid int64, companyName, address, telephone string) (Seller, error)
	Update(id int64, address, telephone string) (Seller, error)
	Delete(id int64) error
}

type RepositorySeller interface {
	GetAll() ([]Seller, error)
	GetById(id int64) (Seller, error)
	Create(cid int64, companyName, address, telephone string) (Seller, error)
	Update(id int64, address, telephone string) (Seller, error)
	CreatID() int64
	Delete(id int64) error
}