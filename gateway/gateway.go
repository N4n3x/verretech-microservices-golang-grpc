package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"verretech-microservices/commande/commandepb"
	commandeDoc "verretech-microservices/commande/documents"
	panierDoc "verretech-microservices/panier/documents"
	"verretech-microservices/panier/panierpb"
	"verretech-microservices/produit/documents"
	"verretech-microservices/produit/produitpb"
	doc "verretech-microservices/utilisateur/documents"
	"verretech-microservices/utilisateur/utilisateurpb"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/shaj13/go-guardian/auth"
	"github.com/shaj13/go-guardian/auth/strategies/basic"
	"github.com/shaj13/go-guardian/auth/strategies/bearer"
	"github.com/shaj13/go-guardian/store"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

type AuthResponse struct {
	Token       string          `json:"token"`
	Utilisateur doc.Utilisateur `json:"utilisateur"`
}

type Values struct {
	m map[string]string
}

func (v Values) Get(key string) string {
	return v.m[key]
}

var authenticator auth.Authenticator
var cache store.Cache

//Variables environnements
var GATEWAY_PORT string
var UTILISATEUR_SERV string
var PRODUIT_SERV string
var PANIER_SERV string
var COMMANDE_SERV string

///
//UTILISATEUR
///

/// AddUtilisateur
// @return Utilisateur
func AddUtilisateur(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u utilisateurpb.Utilisateur
	_ = json.NewDecoder(r.Body).Decode(&u)
	fmt.Printf("%+v\n", &u)
	cc, err := grpc.Dial(UTILISATEUR_SERV, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Unable to connect to server : %v", err)
		json.NewEncoder(w).Encode(err)
		return
	}
	utilisateurClient := utilisateurpb.NewServiceUtilisateurClient(cc)
	b := &utilisateurpb.UtilisateurRequest{
		Utilisateur: &u,
	}
	res, err := utilisateurClient.AddUtilisateur(context.Background(), b)
	if err != nil {
		fmt.Printf("Unable to add Utilisateur: %v", err)
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(res)
	}
}

/// GetUtilisateurs
// @return []utilisateur
func GetUtilisateurs(w http.ResponseWriter, r *http.Request) {
	cc, err := grpc.Dial(UTILISATEUR_SERV, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Unable to connect to server : %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	utilisateurClient := utilisateurpb.NewServiceUtilisateurClient(cc)
	body := &utilisateurpb.UtilisateursRequest{}
	res, err := utilisateurClient.GetUtilisateurs(context.Background(), body)
	if err != nil {
		fmt.Printf("Unable to get Products: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	var utilisateurs []*doc.Utilisateur

	for _, ut := range res.Utilisateur {
		u, err := doc.FromUtilisateurPB(ut)
		if err != nil {
			fmt.Printf("Unable to get Products: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}
		utilisateurs = append(utilisateurs, u)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utilisateurs)
}

///UpdateUtilisateur
// Permission et ID exclu
// @return Utilisateur (with ID)
func UpdateUtilisateur(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := r.Context().Value("user").(Values).Get("username")

	cc, err := grpc.Dial(UTILISATEUR_SERV, grpc.WithInsecure())
	if err != nil {
		///TODO: Gestion erreur
		fmt.Printf("Unable to connect to server : %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	utilisateurClient := utilisateurpb.NewServiceUtilisateurClient(cc)

	ur := &utilisateurpb.UtilisateurRequest{
		Utilisateur: &utilisateurpb.Utilisateur{
			Mail: user,
		},
	}
	uRep, err := utilisateurClient.GetUtilisateur(context.Background(), ur)
	if err != nil {
		fmt.Printf("Erreur Utilisateur : %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	var u doc.Utilisateur
	_ = json.NewDecoder(r.Body).Decode(&u)
	// fmt.Printf("%+v\n", u)
	// utilisateur := u.ToUtilisateurPB()

	// uRep.Utilisateur.ID
	// uRep.Utilisateur.Permission
	if u.Mail != nil {
		uRep.Utilisateur.Mail = *u.Mail
	}
	if u.Nom != nil {
		uRep.Utilisateur.Nom = *u.Nom
	}
	if u.Prenom != nil {
		uRep.Utilisateur.Prenom = *u.Prenom
	}
	if u.HashMotDePasse != nil {
		pass := *u.HashMotDePasse
		hashpass, err := HashPassword(pass)
		if err != nil {

		}
		uRep.Utilisateur.HashMotDePasse = hashpass
	}
	if u.Preferences != nil {
		if u.Preferences.Localisation != nil {
			uRep.Utilisateur.Preferences.Localisation.Adresse = u.Preferences.Localisation.Adresse
			uRep.Utilisateur.Preferences.Localisation.Cp = u.Preferences.Localisation.Cp
			uRep.Utilisateur.Preferences.Localisation.Ville = u.Preferences.Localisation.Ville
		}
		if u.Preferences.PointRetrait != nil {
			uRep.Utilisateur.Preferences.PointRetrait.Nom = u.Preferences.PointRetrait.Nom
			if u.Preferences.PointRetrait.Localisation != nil {
				uRep.Utilisateur.Preferences.PointRetrait.Localisation.Adresse = u.Preferences.PointRetrait.Localisation.Adresse
				uRep.Utilisateur.Preferences.PointRetrait.Localisation.Cp = u.Preferences.PointRetrait.Localisation.Cp
				uRep.Utilisateur.Preferences.PointRetrait.Localisation.Ville = u.Preferences.PointRetrait.Localisation.Ville
			}
		}
	}

	//TODO: Hash password, Exclure autorisation

	b := &utilisateurpb.UtilisateurRequest{
		Utilisateur: uRep.Utilisateur,
	}
	res, err := utilisateurClient.UpdateUtilisateur(context.Background(), b)
	if err != nil {
		fmt.Printf("Unable to update Utilisateur: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}

func GetUtilisateur(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	u, err := infoCurrentUser(r)
	if err != nil {
		fmt.Printf("Unable to get Utilisateur: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	userInfo, err := doc.FromUtilisateurPB(u.Utilisateur)
	if err != nil {
		fmt.Printf("Unable to update Utilisateur: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userInfo)
}

///TODO: Delete Utilisateur

///
//PRODUIT
///

/// GetProduits
// @return []produit
// @param tag (tag=tag1,tag2,tag3...) permet de filtrer les r??sultats par tag
func GetProduits(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)

	cc, err := grpc.Dial(PRODUIT_SERV, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Unable to connect to server : %v", err)
		json.NewEncoder(w).Encode("Erreur grpc produit")
		return
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
			fmt.Printf("Unable to get Products: %v", err)
			json.NewEncoder(w).Encode("Erreur GetProduitsByTags")
			return
		}
		data = res
	} else {
		b := &produitpb.GetAllProduitsRequest{}
		res, err := produitClient.GetAllProduits(context.Background(), b)
		if err != nil {
			fmt.Printf("Unable to get Products: %v", err)
			json.NewEncoder(w).Encode("Erreur GetAllProduits")
			return
		}
		data = res
	}

	var produits []*documents.Produit

	for _, pr := range data.Listproduits.Produits {
		d, derr := documents.FromProduitPB(pr)
		if derr != nil {
			fmt.Printf("Unable to get Products: %v", err)
			json.NewEncoder(w).Encode("Erreur FromProduitPB")
			return
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
// @param Ref -> r??f??rence d'un produit
func GetProduitByRef(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ref"]

	cc, err := grpc.Dial(PRODUIT_SERV, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Unable to connect to server : %v", err)
		json.NewEncoder(w).Encode("Erreur grpc")
		return
	}
	produitClient := produitpb.NewServiceProduitClient(cc)
	b := &produitpb.ProduitByRefRequest{
		Ref: key,
	}
	res, err := produitClient.GetProduitByRef(context.Background(), b)
	if err != nil {
		fmt.Printf("Unable to get Products: %v", err)
		json.NewEncoder(w).Encode("Erreur GetProduitByRef")
		return
	}
	produit, perr := documents.FromProduitPB(res.Produit)
	if perr != nil {
		fmt.Printf("Unable to get Products: %v", err)
		json.NewEncoder(w).Encode("Erreur FromProduitPB")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produit)
}

///
//PANIER
///
func GetPanier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	uRep, err := infoCurrentUser(r)
	if err != nil {
		json.NewEncoder(w).Encode("Erreur Utilisateur")
		return
	}
	cc, err := grpc.Dial(PANIER_SERV, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error: %v", err)
		json.NewEncoder(w).Encode("Erreur grpc")
		return
	}
	panierClient := panierpb.NewServicePanierClient(cc)

	b := &panierpb.ByUtilisateurRequest{
		UtilisateurID: uRep.Utilisateur.ID,
	}
	res, err := panierClient.GetPanier(context.Background(), b)
	if err != nil {
		fmt.Printf("Error: %v", err)
		json.NewEncoder(w).Encode("Erreur GetPanier")
		return
	}
	panier, err := panierDoc.FromPanierPB(res.Panier)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	json.NewEncoder(w).Encode(panier)
}

func UpdatePanier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := r.Context().Value("user").(Values).Get("username")
	cu, err := grpc.Dial(UTILISATEUR_SERV, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	utilisateurClient := utilisateurpb.NewServiceUtilisateurClient(cu)
	ur := &utilisateurpb.UtilisateurRequest{
		Utilisateur: &utilisateurpb.Utilisateur{
			Mail: user,
		},
	}
	uRep, err := utilisateurClient.GetUtilisateur(context.Background(), ur)
	if err != nil {
		json.NewEncoder(w).Encode("Erreur Utilisateur")
		return
	}

	cc, err := grpc.Dial(PANIER_SERV, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	panierClient := panierpb.NewServicePanierClient(cc)
	var pa []*panierpb.Article
	err = json.NewDecoder(r.Body).Decode(&pa)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	b := &panierpb.PanierRequest{
		Panier: &panierpb.Panier{
			UtilisateurID: uRep.Utilisateur.ID,
			Article:       pa,
		},
	}
	res, err := panierClient.UpdatePanier(context.Background(), b)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	panier, err := panierDoc.FromPanierPB(res.Panier)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	json.NewEncoder(w).Encode(panier)
}

///
//Commande
///
func GetCommandes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Va chercher le panier
	uRep, err := infoCurrentUser(r)
	if err != nil {
		json.NewEncoder(w).Encode("Erreur Utilisateur")
		return
	}
	cc, err := grpc.Dial(COMMANDE_SERV, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error: %v", err)
		json.NewEncoder(w).Encode("Erreur GRPC")
		return
	}
	commandeClient := commandepb.NewServiceCommandeClient(cc)

	b := &commandepb.ByUtilisateurRequest{
		UtilisateurID: uRep.Utilisateur.ID,
	}
	res, err := commandeClient.GetUserCommandes(context.Background(), b)
	if err != nil {
		fmt.Printf("Error: %v", err)
		json.NewEncoder(w).Encode("Erreur Commande")
		return
	}
	var commandes []commandeDoc.Commande
	for _, c := range res.Commandes {
		t, err := commandeDoc.FromCommandePB(c)
		if err != nil {
			fmt.Printf("Error: %v", err)
			json.NewEncoder(w).Encode("Erreur CommandePB")
			return
		}
		commandes = append(commandes, *t)
	}
	json.NewEncoder(w).Encode(commandes)
}

func ValidCommande(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Va chercher le panier
	uRep, err := infoCurrentUser(r)
	if err != nil {
		json.NewEncoder(w).Encode("Erreur Utilisateur")
		return
	}
	cc, err := grpc.Dial(PANIER_SERV, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	panierClient := panierpb.NewServicePanierClient(cc)

	b := &panierpb.ByUtilisateurRequest{
		UtilisateurID: uRep.Utilisateur.ID,
	}
	res, err := panierClient.GetPanier(context.Background(), b)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	// le Valide
	con, err := grpc.Dial(COMMANDE_SERV, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	commandeClient := commandepb.NewServiceCommandeClient(con)
	p := &commandepb.PanierRequest{
		Panier: res.Panier,
	}
	result, err := commandeClient.Valid(context.Background(), p)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	// retourne une commande
	cmd, err := commandeDoc.FromCommandePB(result.Commande)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	json.NewEncoder(w).Encode(cmd)
}

func ConfirmCommande(w http.ResponseWriter, r *http.Request) {
	fmt.Println("00 ====>  ", r)
	vars := mux.Vars(r)
	fmt.Println("0 ====>  ", vars)
	idCmd := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	// Va chercher le panier
	uRep, err := infoCurrentUser(r)
	fmt.Println("1 ====>  ", uRep)
	if err != nil {
		json.NewEncoder(w).Encode("Erreur Utilisateur")
		return
	}
	con, err := grpc.Dial(COMMANDE_SERV, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	commandeClient := commandepb.NewServiceCommandeClient(con)
	p := &commandepb.ConfirmRequest{
		CommandeID: idCmd,
		UserID:     uRep.Utilisateur.ID,
	}
	fmt.Println("2 ====>  ", p)
	result, err := commandeClient.Confirm(context.Background(), p)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println("2 ====>  ", result)
	// // Efface le panier
	// conP, err := grpc.Dial(PANIER_SERV, grpc.WithInsecure())
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// }
	// panierClient := panierpb.NewServicePanierClient(conP)
	// panier := &panierpb.PanierRequest{
	// 	Panier: &panierpb.Panier{
	// 		UtilisateurID: uRep.Utilisateur.ID,
	// 		Article:       []*panierpb.Article{},
	// 	},
	// }
	// resultP, err := panierClient.UpdatePanier(context.Background(), panier)
	// if err != nil {
	// 	fmt.Printf("Error: %v %v", err, resultP)
	// }
	// // retourne une commande
	// fmt.Printf("3 ====>  %v", result.Commande)
	// cmd, err := commandeDoc.FromCommandePB(result.Commande)
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// }
	// json.NewEncoder(w).Encode(cmd)
}

func CancelCommande(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

///
//Utilitaire
///
func infoCurrentUser(r *http.Request) (*utilisateurpb.UtilisateurResponse, error) {
	user := r.Context().Value("user").(Values).Get("username")

	cu, err := grpc.Dial(UTILISATEUR_SERV, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	utilisateurClient := utilisateurpb.NewServiceUtilisateurClient(cu)
	ur := &utilisateurpb.UtilisateurRequest{
		Utilisateur: &utilisateurpb.Utilisateur{
			Mail: user,
		},
	}
	return utilisateurClient.GetUtilisateur(context.Background(), ur)
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
	// (*w).Header().Set("Access-Control-Allow-Headers", "*")
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
		fmt.Printf("AUTH: %+v\n", err)
		code := http.StatusUnauthorized
		http.Error(w, http.StatusText(code), code)
		return
	}
	fmt.Printf("Auth: %+v\n", user.UserName())
	tokenStrategy := authenticator.Strategy(bearer.CachedStrategyKey)
	auth.Append(tokenStrategy, token, user, r)

	cc, err := grpc.Dial(UTILISATEUR_SERV, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Unable to connect to server : %v", err)
	}
	utilisateurClient := utilisateurpb.NewServiceUtilisateurClient(cc)
	body := &utilisateurpb.UtilisateurRequest{
		Utilisateur: &utilisateurpb.Utilisateur{
			Mail: user.UserName(),
		},
	}
	res, err := utilisateurClient.GetUtilisateur(context.Background(), body)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	utilisateurInfo, errU := doc.FromUtilisateurPB(res.Utilisateur)
	if errU != nil {
		fmt.Printf("Error: %v", errU)
	}
	resp := &AuthResponse{
		Token:       token,
		Utilisateur: *utilisateurInfo,
	}
	// body := fmt.Sprintf("token: %s \n", token)
	json.NewEncoder(w).Encode(resp)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
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
	cc, err := grpc.Dial(UTILISATEUR_SERV, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Unable to connect to server : %v", err)
	}
	utilisateurClient := utilisateurpb.NewServiceUtilisateurClient(cc)
	body := &utilisateurpb.UtilisateurRequest{
		Utilisateur: &utilisateurpb.Utilisateur{
			Mail:           userName,
			HashMotDePasse: password,
		},
	}
	res, err := utilisateurClient.Auth(context.Background(), body)
	fmt.Printf("1=========>>> : %v\n", res)
	fmt.Printf("2=========>>> : %v\n", err)
	if err != nil {
		return nil, fmt.Errorf("Invalid credentials")
	}

	if res.State {
		return auth.NewDefaultUser(res.Utilisateur.Mail, res.Utilisateur.ID, nil, nil), nil
	}

	return nil, fmt.Errorf("Invalid credentials %v", res)
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

		v := Values{map[string]string{
			"username": user.UserName(),
		}}

		ctx := context.WithValue(r.Context(), "user", v)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.Use(corsMiddleware)
	myRouter.HandleFunc("/auth/token", createToken).Methods("GET")
	myRouter.HandleFunc("/produit", GetProduits).Methods("GET")
	myRouter.HandleFunc("/produit/{ref}", GetProduitByRef).Methods("GET")
	myRouter.HandleFunc("/utilisateur", authMiddleware(http.HandlerFunc(UpdateUtilisateur))).Methods("PUT")
	myRouter.HandleFunc("/utilisateur", AddUtilisateur).Methods("POST")
	myRouter.HandleFunc("/utilisateur", authMiddleware(http.HandlerFunc(GetUtilisateurs))).Methods("GET")
	myRouter.HandleFunc("/utilisateur/info", authMiddleware(http.HandlerFunc(GetUtilisateur))).Methods("GET")
	myRouter.HandleFunc("/panier", authMiddleware(http.HandlerFunc(GetPanier))).Methods("GET")
	myRouter.HandleFunc("/panier", authMiddleware(http.HandlerFunc(UpdatePanier))).Methods("POST")
	myRouter.HandleFunc("/panier", authMiddleware(http.HandlerFunc(UpdatePanier))).Methods("PUT")
	myRouter.HandleFunc("/commande", authMiddleware(http.HandlerFunc(GetCommandes))).Methods("GET")
	myRouter.HandleFunc("/commande/validation", authMiddleware(http.HandlerFunc(ValidCommande))).Methods("GET")
	myRouter.HandleFunc("/commande/confirmation/{id}", authMiddleware(http.HandlerFunc(ConfirmCommande))).Methods("GET")
	myRouter.HandleFunc("/commande/annulation", authMiddleware(http.HandlerFunc(CancelCommande))).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowCredentials: true,
	})
	///TODO: GetPointsRetrait
	myRouter.Use(loggingMiddleware)
	fmt.Println("Gateway => startup complete, listen on port ", GATEWAY_PORT)
	log.Fatal(http.ListenAndServe(":"+GATEWAY_PORT, c.Handler(myRouter)))
}

func main() {
	fmt.Println("Gateway => Starting...")
	GATEWAY_PORT = os.Getenv("GATEWAY_PORT")
	if GATEWAY_PORT == "" {
		GATEWAY_PORT = "10000"
		fmt.Println("Gateway => GATEWAY_PORT variable not found, 10000 used")
	}
	PRODUIT_SERV = os.Getenv("PRODUIT_SERV")
	if PRODUIT_SERV == "" {
		PRODUIT_SERV = "produit:50051"
		fmt.Println("Gateway => PRODUIT_SERV variable not found, produit:50051 used")
	}
	UTILISATEUR_SERV = os.Getenv("UTILISATEUR_SERV")
	if UTILISATEUR_SERV == "" {
		UTILISATEUR_SERV = "utilisateur:50052"
		fmt.Println("Gateway => UTILISATEUR_SERV variable not found, utilisateur:50052 used")
	}
	PANIER_SERV = os.Getenv("PANIER_SERV")
	if PANIER_SERV == "" {
		PANIER_SERV = "panier:50053"
		fmt.Println("Gateway => PANIER_SERV variable not found, panier:50053 used")
	}
	COMMANDE_SERV = os.Getenv("COMMANDE_SERV")
	if COMMANDE_SERV == "" {
		COMMANDE_SERV = "commande:50054"
		fmt.Println("Gateway => COMMANDE_SERV variable not found, commande:50054 used")
	}
	setupGoGuardian()
	handleRequests()
}
