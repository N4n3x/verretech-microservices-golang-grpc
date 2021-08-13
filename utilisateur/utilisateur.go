package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"verretech-microservices/database"
	"verretech-microservices/utilisateur/documents"
	"verretech-microservices/utilisateur/utilisateurpb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	db *database.Mongo
	utilisateurpb.UnimplementedServiceUtilisateurServer
}

func (server *server) AddUtilisateur(ctx context.Context, req *utilisateurpb.UtilisateurRequest) (*utilisateurpb.UtilisateurResponse, error) {
	mongoUtilisateur, _ := documents.FromUtilisateurPB(req.Utilisateur)
	hpw, err := HashPassword(*mongoUtilisateur.HashMotDePasse)
	mongoUtilisateur.HashMotDePasse = &hpw

	// fmt.Printf("Mongo produit %+v\n", mongoProduit)
	oid, err := mongoUtilisateur.InsertOne(*server.db.Database)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Unable to process request: %v", err))
	}
	mongoUtilisateur.ID = &oid
	var response utilisateurpb.UtilisateurResponse
	response.Utilisateur = mongoUtilisateur.ToUtilisateurPB()
	return &response, nil
}

func (server *server) GetUtilisateurs(ctx context.Context, req *utilisateurpb.UtilisateursRequest) (*utilisateurpb.UtilisateursResponse, error) {
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

func (server *server) GetUtilisateur(ctx context.Context, req *utilisateurpb.UtilisateurRequest) (*utilisateurpb.UtilisateurResponse, error) {
	var utilisateur documents.Utilisateur
	if req.Utilisateur.ID != "" {
		oid, err := primitive.ObjectIDFromHex(req.Utilisateur.ID)
		if err != nil {
			return nil, err
		}
		utilisateur.ID = &oid
	} else if req.Utilisateur.Mail != "" {
		utilisateur.Mail = &req.Utilisateur.Mail
	} else {
		err := errors.New("Utilisateur incomplet")
		return nil, err
	}
	err := utilisateur.FindOne(*server.db.Database)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Unable to process request: %v", err))
	}
	var response utilisateurpb.UtilisateurResponse
	response.Utilisateur = utilisateur.ToUtilisateurPB()
	return &response, nil
}

func (server *server) UpdateUtilisateur(ctx context.Context, req *utilisateurpb.UtilisateurRequest) (*utilisateurpb.UtilisateurResponse, error) {
	utilisateur, _ := documents.FromUtilisateurPB(req.Utilisateur)

	_, err := utilisateur.Update(*server.db.Database)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Unable to process request: %v", err))
	}
	var response utilisateurpb.UtilisateurResponse
	response.Utilisateur = utilisateur.ToUtilisateurPB()
	return &response, nil
}

func (server *server) Auth(ctx context.Context, req *utilisateurpb.UtilisateurRequest) (*utilisateurpb.AuthResponse, error) {
	response := &utilisateurpb.AuthResponse{
		State: false,
	}
	var utilisateur documents.Utilisateur
	utilisateur.Mail = &req.Utilisateur.Mail
	err := utilisateur.FindOne(*server.db.Database)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Unable to process request: %v", err))
	}
	h := *utilisateur.HashMotDePasse
	if CheckPasswordHash(req.Utilisateur.HashMotDePasse, h) {
		response.State = true
		response.Utilisateur = utilisateur.ToUtilisateurPB()
	}

	return response, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	fmt.Println("Service Utilisateur => Starting...")
	UTILISATEUR_PORT := os.Getenv("UTILISATEUR_PORT")
	if UTILISATEUR_PORT == "" {
		UTILISATEUR_PORT = "50052"
		fmt.Println("Service Produit => UTILISATEUR_PORT variable not found, 50052 used")
	}
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("Error while creating listener : %v", err)
	}
	fmt.Println("Service Utilisateur => HTTP Server OK")

	utilisateurServer := &server{
		db: database.NewMongoConnection(),
	}

	err = utilisateurServer.db.ConnectToDB("mongodb+srv://verretech:7BV5eF2zzy29LK9V@db1.3hf1n.mongodb.net", "verretech")
	if err != nil {
		log.Fatalf("Unable to connect to db : %v", err)
	}
	fmt.Println("Service Utilisateur => Database connection OK")

	gRPCServer := grpc.NewServer()
	utilisateurpb.RegisterServiceUtilisateurServer(gRPCServer, utilisateurServer)
	fmt.Println("Service Utilisateur => GRPC Server OK")
	fmt.Println("Service Utilisateur => Startup complete, listen on port ", UTILISATEUR_PORT)

	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}
