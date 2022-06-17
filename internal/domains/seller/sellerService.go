package seller

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Seller, error) {
	listSeller, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return listSeller, nil
}

func (s *service) GetById(id int64) (Seller, error) {
	seller, err := s.repository.GetById(id)

	if err != nil {
		return Seller{}, err
	}
	return seller, nil
}

func (s *service) CreateSeller(cid int64, companyName, address, telephone string) (Seller, error) {
	seller, err := s.repository.CreateSeller(cid, companyName, address, telephone)
	if err != nil {
		return Seller{}, err
	}
	return seller, nil
}

func (s *service) UpdateSellerAddresAndTel(id int64, address, telephone string) (Seller, error) {
	seller, err := s.repository.UpdateSellerAddresAndTel(id, address, telephone)
	if err != nil {
		return Seller{}, err
	}
	return seller, nil
}

func (s *service) DeleteSeller(id int64) error {
	err := s.repository.DeleteSeller(id)

	if err != nil {
		return err
	}
	return nil
}
