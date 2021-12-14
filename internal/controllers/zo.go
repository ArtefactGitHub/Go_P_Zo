package controllers

import (
	"github.com/ArtefactGitHub/Go_P_Zo/internal/models/zo"
	"github.com/ArtefactGitHub/Go_P_Zo/internal/services"
)

type ZoController struct {
	zs services.ZoService
}

// TODO
func (c *ZoController) GetAll() ([]zo.Zo, error) {
	result, err := c.zs.GetAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ZoController) Get(id int) (*zo.Zo, error) {
	result, err := c.zs.Get(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ZoController) Post(z *zo.Zo) (int, error) {
	result, err := c.zs.Post(z)
	if err != nil {
		return -1, err
	}

	return result, nil
}

func (c *ZoController) Update(z *zo.Zo) (*zo.Zo, error) {
	result, err := c.zs.Update(z)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ZoController) Delete(z *zo.Zo) error {
	err := c.zs.Delete(z)
	if err != nil {
		return err
	}

	return nil
}
