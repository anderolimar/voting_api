package repositories

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VotesRepository interface{}

func NewVotesRepository() VotesRepository {
	return &votesRepository{}
}

type votesRepository struct {
	client *mongo.Client
}

func (v votesRepository) Insert() {
	// Conectar ao MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	db := client.Database("meu_banco")        // Nome do banco de dados
	colecao := db.Collection("minha_colecao") // Nome da coleção

	// Inserir um documento
	documento := bson.D{{"nome", "João"}, {"idade", 30}, {"cidade", "São Paulo"}}
	res, err := colecao.InsertOne(context.TODO(), documento)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Documento inserido com ID: %v\n", res.InsertedID)

	// Inserir múltiplos documentos
	documentos := []interface{}{
		bson.D{{"nome", "Maria"}, {"idade", 25}, {"cidade", "Rio de Janeiro"}},
		bson.D{{"nome", "Carlos"}, {"idade", 40}, {"cidade", "Belo Horizonte"}},
	}
	resMany, err := colecao.InsertMany(context.TODO(), documentos)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("IDs dos documentos inseridos: %v\n", resMany.InsertedIDs)

	// Consultar um documento
	var resultado bson.M
	err = colecao.FindOne(context.TODO(), bson.D{{"nome", "João"}}).Decode(&resultado)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Documento encontrado:", resultado)

	// Consultar múltiplos documentos
	cursor, err := colecao.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	fmt.Println("Todos os documentos:")
	for cursor.Next(context.TODO()) {
		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			log.Fatal(err)
		}
		fmt.Println(doc)
	}
}
