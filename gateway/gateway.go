package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"verretech-microservices/produit/documents"
	"verretech-microservices/produit/produitpb"
	doc "verretech-microservices/utilisateur/documents"
	"verretech-microservices/utilisateur/utilisateurpb"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/shaj13/go-guardian/auth"
	"github.com/shaj13/go-guardian/auth/strategies/basic"
	"github.com/shaj13/go-guardian/auth/strategies/bearer"
	"github.com/shaj13/go-guardian/store"
	"google.golang.org/grpc"
)

type Token struct {
	Token string `json:"token"`
}

var authenticator auth.Authenticator
var cache store.Cache

/// AddUtilisateur
// @return Utilisateur
func AddUtilisateur(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u utilisateurpb.Utilisateur
	_ = json.NewDecoder(r.Body).Decode(&u)
	fmt.Printf("%+v\n", &u)
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to server : %v", err)
	}
	utilisateurClient := utilisateurpb.NewServiceUtilisateurClient(cc)
	b := &utilisateurpb.UtilisateurRequest{
		Utilisateur: &u,
	}
	res, err := utilisateurClient.AddUtilisateur(context.Background(), b)
	if err != nil {
		fmt.Printf("Unable to update Utilisateur: %v", err)
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(res)
	}
}

/// GetUtilisateurs
// @return []utilisateur
func GetUtilisateurs(w http.ResponseWriter, r *http.Request) {
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to server : %v", err)
	}
	utilisateurClient := utilisateurpb.NewServiceUtilisateurClient(cc)
	body := &utilisateurpb.UtilisateursRequest{}
	res, err := utilisateurClient.GetUtilisateurs(context.Background(), body)
	if err != nil {
		log.Fatalf("Unable to get Products: %v", err)
	}

	var utilisateurs []*doc.Utilisateur

	for _, ut := range res.Utilisateur {
		u, derr := doc.FromUtilisateurPB(ut)
		if derr != nil {
			log.Fatalf("Unable to get Products: %v", err)
		}
		utilisateurs = append(utilisateurs, u)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utilisateurs)
}

///UpdateUtilisateur
// @return Utilisateur (with ID)
func UpdateUtilisateur(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u utilisateurpb.Utilisateur
	_ = json.NewDecoder(r.Body).Decode(&u)
	fmt.Printf("%+v\n", &u)
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		///TODO: Gestion erreur
		log.Fatalf("Unable to connect to server : %v", err)
	}
	utilisateurClient := utilisateurpb.NewServiceUtilisateurClient(cc)
	b := &utilisateurpb.UtilisateurRequest{
		Utilisateur: &u,
	}
	res, err := utilisateurClient.UpdateUtilisateur(context.Background(), b)
	if err != nil {
		fmt.Printf("Unable to update Utilisateur: %v", err)
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(res)
	}
}

///TODO: Delete Utilisateur

/// GetProduits
// @return []produit
// @param tag (tag=tag1,tag2,tag3...) permet de filtrer les résultats par tag
func GetProduits(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to server : %v", err)
	}
	produitClient := produitpb.NewServiceProduitClient(cc)
	params := r.URL.Query().Get("tag")
	data := &produitpb.ProduitsResponse{}
	if params != "" {
		tags := strings.Split(params, ",")
		fmt.Printf("%+v\n", tags)
		b := &produitpb.ProduitsByTags{Tags: tags}
		res, err := produitClient.GetProduitsByTags(context.Background(), b)
		if err != nil {
			log.Fatalf("Unable to get Products: %v", err)
		}
		data = res
	} else {
		b := &produitpb.GetAllProduitsRequest{}
		res, err := produitClient.GetAllProduits(context.Background(), b)
		if err != nil {
			log.Fatalf("Unable to get Products: %v", err)
		}
		data = res
	}

	var produits []*documents.Produit

	for _, pr := range data.Listproduits.Produits {
		d, derr := documents.FromProduitPB(pr)
		if derr != nil {
			log.Fatalf("Unable to get Products: %v", err)
		}
		produits = append(produits, d)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(produits)

	// fmt.Fprintf(w, "Tags: %v\n", tags)
}

/// GetProduitByRef
// @return Produit
// @param Ref -> référence d'un produit
func GetProduitByRef(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ref"]

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to server : %v", err)
	}
	produitClient := produitpb.NewServiceProduitClient(cc)
	b := &produitpb.ProduitByRefRequest{
		Ref: key,
	}
	res, err := produitClient.GetProduitByRef(context.Background(), b)
	if err != nil {
		log.Fatalf("Unable to get Products: %v", err)
	}
	produit, perr := documents.FromProduitPB(res.Produit)
	if perr != nil {
		log.Fatalf("Unable to get Products: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produit)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func createToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := uuid.New().String()
	user, err := authenticator.Authenticate(r)
	if err != nil {
		fmt.Printf("Auth: %+v\n", err)
		code := http.StatusUnauthorized
		http.Error(w, http.StatusText(code), code)
		return
	}
	fmt.Printf("Auth: %+v\n", r)

	tokenStrategy := authenticator.Strategy(bearer.CachedStrategyKey)
	auth.Append(tokenStrategy, token, user, r)
	body := &Token{Token: token}
	// body := fmt.Sprintf("token: %s \n", token)
	json.NewEncoder(w).Encode(body)
}

func setupGoGuardian() {
	authenticator = auth.New()
	cache = store.NewFIFO(context.Background(), time.Minute*10)

	basicStrategy := basic.New(validateUser, cache)
	tokenStrategy := bearer.New(bearer.NoOpAuthenticate, cache)

	authenticator.EnableStrategy(basic.StrategyKey, basicStrategy)
	authenticator.EnableStrategy(bearer.CachedStrategyKey, tokenStrategy)
}

func validateUser(ctx context.Context, r *http.Request, userName, password string) (auth.Info, error) {
	///TODO: connect to Utilisateur Service
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to server : %v", err)
	}
	utilisateurClient := utilisateurpb.NewServiceUtilisateurClient(cc)
	body := &utilisateurpb.UtilisateurRequest{
		Utilisateur: &utilisateurpb.Utilisateur{
			Mail:           userName,
			HashMotDePasse: password,
		},
	}
	res, err := utilisateurClient.Auth(context.Background(), body)
	if err != nil {
		return nil, fmt.Errorf("Invalid credentials")
	}

	if res.State {
		return auth.NewDefaultUser(res.Utilisateur.Mail, res.Utilisateur.ID, nil, nil), nil
	}

	return nil, fmt.Errorf("Invalid credentials")
}

func authMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing Auth Middleware")
		user, err := authenticator.Authenticate(r)
		if err != nil {
			fmt.Printf("Auth: %+v\n", err)
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}
		log.Printf("User %s Authenticated\n", user.UserName())
		next.ServeHTTP(w, r)
	})
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(corsMiddleware)
	myRouter.HandleFunc("/auth/token", createToken).Methods("GET")
	myRouter.HandleFunc("/produit", GetProduits).Methods("GET")
	myRouter.HandleFunc("/produit/{ref}", GetProduitByRef).Methods("GET")
	myRouter.HandleFunc("/utilisateur", UpdateUtilisateur).Methods("PUT")
	myRouter.HandleFunc("/utilisateur", AddUtilisateur).Methods("POST")
	myRouter.HandleFunc("/utilisateur", authMiddleware(http.HandlerFunc(GetUtilisateurs))).Methods("GET")
	myRouter.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	setupGoGuardian()
	handleRequests()
}
