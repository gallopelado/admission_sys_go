package dto

type CiudadDTO struct {
	ciudId         int
	ciuDescripcion string
}

func (c *CiudadDTO) SetCiuId(ciudId int) {
	c.ciudId = ciudId
}

func (c *CiudadDTO) GetCiuId() int {
	return c.ciudId
}

func (c *CiudadDTO) SetDescripcion(ciuDescripcion string) {
	c.ciuDescripcion = ciuDescripcion
}

func (c *CiudadDTO) GetDescripcion() string {
	return c.ciuDescripcion
}
