package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/paolapesantez/avatweetServer/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscarPerfil - busca un perfil en la BD*/
func BuscarPerfil(ID string) (models.Usuario, error) {

	// vamos a usar una petición GET, ya que va a venir como parámetro en la URL
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("microblogging")
	col := db.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{"_id": objID}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return perfil, err
	}
	return perfil, nil
}
