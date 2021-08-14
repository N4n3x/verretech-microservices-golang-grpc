import firebase from 'firebase/app'

const firebaseConfig = {
    apiKey: "AIzaSyDD1uv8wbyz74Kd8cpzTWqDN7Hrr1-KWL4",
    authDomain: "filerougeril.firebaseapp.com",
    projectId: "filerougeril",
    storageBucket: "filerougeril.appspot.com",
    messagingSenderId: "989648905002",
    appId: "1:989648905002:web:f18b8c6cf01f1fe98cb453"
};

firebase.initializeApp(firebaseConfig);

export { firebaseConfig }