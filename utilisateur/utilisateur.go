package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"verretech-microservices/database"
	"verretech-microservices/utilisateur/documents"
	"verretech-microservices/utilisateur/utilisateurpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	db *database.Mongo
	utilisateurpb.UnimplementedServiceUtilisateurServer
}

func (server *server) AddProduit(ctx context.Context, req *utilisateurpb.UtilisateurRequest) (*utilisateurpb.UtilisateurResponse, error) {
	mongoUtilisateur, _ := documents.FromUtilisateurPB(req.Utilisateur)
	// fmt.Printf("Mongo produit %+v\n", mongoProduit)
	oid, err := mongoUtilisateur.InsertOne(*server.db.Database)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Unable to process request: %v", err))
	}
	mongoUtilisateur.ID = oid
	var response utilisateurpb.UtilisateurResponse
	response.Utilisateur = mongoUtilisateur.ToUtilisateurPB()
	return &response, nil
}

func (server *server) GetAllUtilisateurs(ctx context.Context, req *utilisateurpb.UtilisateursRequest) (*utilisateurpb.UtilisateursResponse, error) {
	u, err := documents.Find(*server.db.Database)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Unable to process request: %v", err))
	}
	var utilisateurs []documents.Utilisateur
	if err = u.All(ctx, &utilisateurs); err != nil {
		log.Fatal(err)
	}
	var response utilisateurpb.UtilisateursResponse
	for _, ut := range utilisateurs {
		response.Utilisateur = append(response.Utilisateur, ut.ToUtilisateurPB())
	}
	return &response, nil
}

// func (server *server) GetProduitsByTags(ctx context.Context, req *produitpb.ProduitsByTags) (*produitpb.ProduitsResponse, error) {
// 	fmt.Printf("%+v\n", req.Tags)
// 	t := []string{}
// 	for _, v := range req.Tags {
// 		t = append(t, v)
// 	}

// 	fmt.Printf("%+v\n", t)
// 	query := bson.M{
// 		"tags": bson.M{"$all": t},
// 	}
// 	p, err := documents.FindByTags(*server.db.Database, query)
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, fmt.Sprintf("Unable to process request: %v", err))
// 	}
// 	var produits []documents.Produit
// 	if err = p.All(ctx, &produits); err != nil {
// 		log.Fatal(err)
// 	}
// 	var response produitpb.ListProduits
// 	for _, pr := range produits {
// 		response.Produits = append(response.Produits, pr.ToProduitPB())
// 	}
// 	var final produitpb.ProduitsResponse
// 	final.Listproduits = &response
// 	return &final, nil
// }

// func (server *server) GetProduitByRef(ctx context.Context, req *produitpb.ProduitByRefRequest) (*produitpb.ProduitResponse, error) {
// 	var produit documents.Produit
// 	produit.Ref = req.Ref
// 	err := produit.FindOne(*server.db.Database)
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, fmt.Sprintf("Unable to process request: %v", err))
// 	}
// 	var response produitpb.ProduitResponse
// 	response.Produit = produit.ToProduitPB()
// 	return &response, nil
// }

// func (server *server) UpdateProduit(ctx context.Context, req *produitpb.ProduitRequest) (*produitpb.ProduitResponse, error) {
// 	mongoProduit, _ := documents.FromProduitPB(req.Produit)

// 	_, err := mongoProduit.Update(*server.db.Database)
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, fmt.Sprintf("Unable to process request: %v", err))
// 	}
// 	var response produitpb.ProduitResponse
// 	response.Produit = mongoProduit.ToProduitPB()
// 	return &response, nil
// }

// func (server *server) UpdateProduits(ctx context.Context, req *produitpb.ProduitsRequest) (*produitpb.BoolResponse, error) {
// 	var p []*documents.Produit
// 	for _, e := range req.Produits {
// 		mongoProduit, _ := documents.FromProduitPB(e)
// 		p = append(p, mongoProduit)
// 	}

// 	_, err := documents.UpdateAll(*server.db.Database, p)
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, fmt.Sprintf("Unable to process request: %v", err))
// 	}

// 	response := &produitpb.BoolResponse{
// 		State: true,
// 	}

// 	return response, nil
// }

// func (server *server) DeleteProduit(ctx context.Context, req *produitpb.ProduitByRefRequest) (*produitpb.BoolResponse, error) {
// 	var produit documents.Produit
// 	produit.Ref = req.Ref
// 	_, err := produit.Delete(*server.db.Database)
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, fmt.Sprintf("Unable to process request: %v", err))
// 	}
// 	var response produitpb.BoolResponse
// 	response.State = true
// 	return &response, nil
// }

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("Error while creating listener : %v", err)
	}

	utilisateurServer := &server{
		db: database.NewMongoConnection(),
	}

	err = utilisateurServer.db.ConnectToDB("mongodb+srv://verretech:7BV5eF2zzy29LK9V@db1.3hf1n.mongodb.net", "verretech")
	if err != nil {
		log.Fatalf("Unable to connect to db : %v", err)
	}
	gRPCServer := grpc.NewServer()
	utilisateurpb.RegisterServiceUtilisateurServer(gRPCServer, utilisateurServer)

	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}
