function setData(table, data, lastData = null, id = null) {

    // FIREBASE
    const firebase = require("firebase")
    require("firebase/firestore")
    var db = firebase.firestore()

    if (id == null) {
        db.collection(table)
            .add(data)
            .then(function (docRef) {
                id = docRef.id;
                data.id = id;
            })
            .catch(function (error) {
                console.error("Error writing document: ", error);
            })
            .finally(() => {
                db.collection(table).doc(data.id).update(data)
                    .then(result => generateNavMenuItems(table, data))
                    .catch(function (error) {
                        console.error("Error writing document: ", error);
                    });
            });
    } else {
        db.collection(table).doc(id).update(data)
            .then(function (result) {
                return true;
            })
            .catch(function (error) {
                console.log(error)
            })
        db.collection(table).doc(id).collection("former").doc().set(lastData)
            .then(function (result) {
                return true;
            })
            .catch(function (error) {
                console.log(error)
            })
    }
}


export { setData }