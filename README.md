# verretech-microservices
Architecture micro services en GO pour site e-Commerce

## TODO
### Produit
* ✅ Création interface GRPC
* ✅ Connexion à la base de données
* 🛠 Conception des fonctions métiers
  * ✅ Create
  * ✅ ReadAll
  * 🛠 ReadByID
  * ✅ Update
  * 🛠 Delete
### Utilisateur
* ✅ Création interface GRPC
* ⏳ Connexion à la base de données
* ⏳ Conception des fonctions métiers
### Panier
* ✅ Création interface GRPC
* ⏳ Connexion à la base de données
* ⏳ Conception des fonctions métiers
### Commande
* ✅ Création interface GRPC
* ⏳ Connexion à la base de données
* ⏳ Conception des fonctions métiers
### Gateway
* ⏳ Référencement des routes
* ⏳ Conception middleware Auth
* ⏳ Conception middleware Log
* ⏳ Mise en place du routage
### ERP
* ⏳ Création connecteur ERP
* ⏳ Conseption taches récurentes

### Endpoints Gateway
* Produit
  * GET /produit?params=###
  * GET /produit/ID
* Utilisateur
  * POST /utilisateur
  * GET /utilisateur?params=###
  * GET /utilisateur/ID
  * PUT /utilisateur/ID
  * DELETE /utilisateur/ID
* Panier
  * POST /panier
  * GET /panier?params=###
  * GET /panier/ID
  * PUT /panier/ID
* Commande
  * POST /commande
  * GET /commande?params=###
  * GET /commande/ID
* Indicateur
  * GET /indicateur?params=###

### CMD
* Générer les interfaces du protocol buffer:
  * ```protoc --go_opt=Mproto/utilisateur=verretech-microservices/utilisateur/utilisateurpb --go_out=../ --go-grpc_out=../ ./proto/*.proto```