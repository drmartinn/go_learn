package user

import "time"

type Usuario struct {
	ID        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

func (this *Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, status bool) {
	this.ID = id
	this.Nombre = nombre
	this.FechaAlta = fechaalta
	this.Status = status
}
