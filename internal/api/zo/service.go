package zo

type zoService struct {
	zr zoRepository
}

func (s *zoService) getAll() ([]zo, error) {
	result, err := s.zr.findall()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *zoService) get(id int) (*zo, error) {
	result, err := s.zr.find(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *zoService) post(z *zo) (int, error) {
	result, err := s.zr.create(z)
	if err != nil {
		return -1, err
	}

	return result, nil
}

func (s *zoService) update(z *zo) error {
	err := s.zr.update(z)
	if err != nil {
		return err
	}

	return nil
}

func (s *zoService) delete(id int) error {
	err := s.zr.delete(id)
	if err != nil {
		return err
	}

	return nil
}
