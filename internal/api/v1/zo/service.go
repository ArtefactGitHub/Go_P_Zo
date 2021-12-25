package zo

type ZoService struct {
	Zr ZoRepository
}

func (s *ZoService) GetAll() ([]Zo, error) {
	result, err := s.Zr.Findall()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ZoService) Get(id int) (*Zo, error) {
	result, err := s.Zr.Find(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ZoService) Post(z *Zo) (int, error) {
	result, err := s.Zr.Create(z)
	if err != nil {
		return -1, err
	}

	return result, nil
}

func (s *ZoService) Update(z *Zo) error {
	err := s.Zr.Update(z)
	if err != nil {
		return err
	}

	return nil
}

func (s *ZoService) Delete(id int) error {
	err := s.Zr.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
