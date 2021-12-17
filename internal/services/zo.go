package services

import (
	"github.com/ArtefactGitHub/Go_P_Zo/internal/models/zo"
)

type ZoService struct {
}

func (s *ZoService) GetAll() ([]zo.Zo, error) {
	result, err := zo.FindAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ZoService) Get(id int) (*zo.Zo, error) {
	result, err := zo.Find(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ZoService) Post(z *zo.Zo) (int, error) {
	result, err := zo.Create(z)
	if err != nil {
		return -1, err
	}

	return result, nil
}

func (s *ZoService) Update(z *zo.Zo) (*zo.Zo, error) {
	err := z.Update()
	if err != nil {
		return nil, err
	}

	return z, nil
}

func (s *ZoService) Delete(id int) error {
	err := zo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
