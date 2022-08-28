package repository

import (
	"app/config"
	"app/domain"
	"app/usecase"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlanetaRepositoryDB struct {
	cliente *mongo.Client
	ctx     *context.Context
}

func NewPlanetaRepositoryDB(c *mongo.Client, ctx *context.Context) *PlanetaRepositoryDB {
	return &PlanetaRepositoryDB{cliente: c, ctx: ctx}
}

func (db *PlanetaRepositoryDB) CreatePlanet(req *domain.Planeta) (r interface{}, err error) {
	planetCollection := db.cliente.Database(config.GetDBName()).Collection(config.GetTable())

	regs, err := db.countPlanetByName(req.Nome)
	if err != nil {
		return nil, err
	}

	if regs > 0 {
		return nil, domain.ErrPlanetExists
	}

	result, err := planetCollection.InsertOne(*db.ctx, req)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(result.InsertedID)

	return
}

func (db *PlanetaRepositoryDB) GetAll() (interface{}, error) {
	planetCollection := db.cliente.Database(config.GetDBName()).Collection(config.GetTable())

	cur, err := planetCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println(err.Error())
	}
	defer cur.Close(*db.ctx)

	var result []domain.Planeta
	if err = cur.All(*db.ctx, &result); err != nil {
		log.Println(err.Error())
	}

	log.Println(result)
	return result, nil
}

func (db *PlanetaRepositoryDB) GetByFilter(filter string, value interface{}) (resp interface{}, err error) {
	planetCollection := db.cliente.Database(config.GetDBName()).Collection(config.GetTable())

	// No caso de consulta por id, converte para primitive.ObjectID
	if filter == "_id" {
		value, err = usecase.StringToObjectID(value.(string))
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}

	cur, err := planetCollection.Find(context.Background(), bson.D{{Key: filter, Value: value}})
	if err != nil {
		log.Println(err.Error())
	}
	defer cur.Close(*db.ctx)

	var result []domain.Planeta
	if err = cur.All(*db.ctx, &result); err != nil {
		log.Println(err.Error())
	}

	return result, nil
}

func (db *PlanetaRepositoryDB) DeletePlanetByID(id primitive.ObjectID) (err error) {
	planetCollection := db.cliente.Database(config.GetDBName()).Collection(config.GetTable())
	_, err = planetCollection.DeleteOne(*db.ctx, bson.M{"_id": id})
	return
}

func (db *PlanetaRepositoryDB) countPlanetByName(name string) (int, error) {
	planetCollection := db.cliente.Database(config.GetDBName()).Collection(config.GetTable())

	cur, err := planetCollection.Find(context.Background(), bson.D{{Key: "nome", Value: name}})
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	defer cur.Close(*db.ctx)

	var result []domain.Planeta
	if err = cur.All(*db.ctx, &result); err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return len(result), nil
}
