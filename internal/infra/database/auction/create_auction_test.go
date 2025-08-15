package auction

import (
	"context"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestAuctionAutoClose(t *testing.T) {
	os.Setenv("AUCTION_DURATION_SECONDS", "2") // fecha em 2s

	// Conectar ao Mongo (ajustar string de conexão se necessário)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to connect to Mongo: %v", err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("auction_test")
	repo := NewAuctionRepository(db)

	auction := &auction_entity.Auction{
		Id:          "test-auction-123",
		ProductName: "Produto Teste",
		Category:    "Teste",
		Description: "Leilão de teste",
		Condition:   auction_entity.New,
		Status:      auction_entity.Active,
		Timestamp:   time.Now(),
	}

	errCreate := repo.CreateAuction(context.Background(), auction)
	if errCreate != nil {
		t.Fatalf("Failed to create auction: %v", errCreate)
	}

	// Espera 3s para garantir fechamento
	time.Sleep(3 * time.Second)

	// Verifica se foi fechado
	var result AuctionEntityMongo
	errFind := repo.Collection.FindOne(context.Background(), map[string]string{"_id": auction.Id}).Decode(&result)
	if errFind != nil {
		t.Fatalf("Failed to find auction: %v", errFind)
	}
	if result.Status != auction_entity.Completed {
		t.Fatalf("Expected auction to be closed, got %v", result.Status)
	}
}
