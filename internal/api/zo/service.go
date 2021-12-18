package zo

type ZoService struct {
}

func (s *ZoService) GetAll() ([]Zo, error) {
	result, err := FindAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ZoService) Get(id int) (*Zo, error) {
	result, err := Find(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ZoService) Post(z *Zo) (int, error) {
	result, err := Create(z)
	if err != nil {
		return -1, err
	}

	return result, nil
}

func (s *ZoService) Update(z *Zo) (*Zo, error) {
	err := z.Update()
	if err != nil {
		return nil, err
	}

	return z, nil
}

func (s *ZoService) Delete(id int) error {
	err := Delete(id)
	if err != nil {
		return err
	}

	return nil
}
