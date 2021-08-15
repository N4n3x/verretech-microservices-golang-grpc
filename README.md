# verretech-microservices
Architecture micro services en GO pour site e-Commerce

## Déploiement
* ```git clone https://gitlab.com/alex.hern.dev/verretech-microservices.git```
* ```cd verretech-microservices```
* ```docker-compose up```

## TODO
### Produit (:50051)
* ✅ Création interface GRPC
* ✅ Connexion à la base de données
* ✅ Conception des fonctions métiers
  * ✅ Create
  * ✅ ReadAll
  * ✅ ReadByID
  * ✅ Update
  * ✅ UpdateAll
  * ✅ Delete
### Utilisateur (:50052)
* ✅ Création interface GRPC
* ✅ Connexion à la base de données
* ✅ Conception des fonctions métiers
### Panier (:50053)
* ✅ Création interface GRPC
* ✅ Connexion à la base de données
* ✅ Conception des fonctions métiers
### Commande (:50054)
* ✅ Création interface GRPC
* ✅ Connexion à la base de données
* ✅ Conception des fonctions métiers
### ERP (:50050)
* ✅ Création connecteur ERP
* ✅ Mise à jour Service Produit
* ⏳ Conseption taches récurentes
### Gateway (:10000)
* ✅ Référencement des routes
* ✅ Conception middleware Auth
* ⏳ Conception middleware Log
* ✅ Mise en place du routage
### Endpoints Gateway
* Produit
  * ✅ GET /produit?tag=###
  * ✅ GET /produit/Ref
* Utilisateur
  * ✅ GET /auth
  * ✅ POST /utilisateur
  * ✅ GET /utilisateur
  * ✅ PUT /utilisateur
* Panier
  * ✅ POST /panier
  * ✅ GET /panier
  * ✅ PUT /panier
* Commande
  * ✅ POST /commande
  * ✅ GET /commande/validation
  * ✅ GET /commande/confirmation/{idCommande}
  * ✅ GET /commande/annulation/{idCommande}

### Gestion d'erreurs ✅

### CMD
* Générer les interfaces du protocol buffer:
  * ```protoc --go_opt=Mproto/utilisateur=verretech-microservices/utilisateur/utilisateurpb --go_out=../ --go-grpc_out=../ ./proto/*.proto```