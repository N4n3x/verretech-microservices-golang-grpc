# verretech-microservices
Architecture micro services en GO pour site e-Commerce

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
* ⏳ Connexion à la base de données
* ⏳ Conception des fonctions métiers
### Commande (:50054)
* ✅ Création interface GRPC
* ⏳ Connexion à la base de données
* ⏳ Conception des fonctions métiers
### ERP (:50050)
* ✅ Création connecteur ERP
* ✅ Mise à jour Service Produit
* ⏳ Conseption taches récurentes
### Gateway (:10000)
* ✅ Référencement des routes
* ✅ Conception middleware Auth
* ⏳ Conception middleware Log
* 🛠 Mise en place du routage
### Endpoints Gateway
* Produit
  * ✅ GET /produit?tag=###
  * ✅ GET /produit/Ref
* Utilisateur
  * ✅ POST /utilisateur
  * ✅ GET /utilisateur?params=###
  * ⏳ GET /utilisateur/ID
  * ⏳ PUT /utilisateur/ID
  * ⏳ DELETE /utilisateur/ID
* Panier
  * ⏳ POST /panier
  * ⏳ GET /panier?params=###
  * ⏳ GET /panier/ID
  * ⏳ PUT /panier/ID
* Commande
  * ⏳ POST /commande
  * ⏳ GET /commande?params=###
  * ⏳ GET /commande/ID
* Indicateur
  * ⏳ GET /indicateur?params=###

### CMD
* Générer les interfaces du protocol buffer:
  * ```protoc --go_opt=Mproto/utilisateur=verretech-microservices/utilisateur/utilisateurpb --go_out=../ --go-grpc_out=../ ./proto/*.proto```