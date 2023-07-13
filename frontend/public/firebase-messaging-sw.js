importScripts('https://www.gstatic.com/firebasejs/9.8.4/firebase-app.js')
importScripts('https://www.gstatic.com/firebasejs/9.8.4/firebase-messaging.js')

const firebaseConfig = {
  apiKey: "AIzaSyAzfkaScszZ4Kkn6hEDoy598OR-OQui-04",
  authDomain: "news-feeds-392507.firebaseapp.com",
  projectId: "news-feeds-392507",
  storageBucket: "news-feeds-392507.appspot.com",
  messagingSenderId: "143677486946",
  appId: "1:143677486946:web:3fecb01fa5d5079c10de3b",
  measurementId: "G-W88JNN9T54"
};

firebase.initializeApp(firebaseConfig)
const messaging = firebase.messaging.isSupported() ? firebase.messaging() : null


// const messaging = firebase.messaging()

messaging.onBackgroundMessage((payload) => {
  const notificationTitle = payload.notification.title
  const notificationOptions = {
    body: payload.notification.body,
    icon: payload.notification.icon || payload.notification.image,
  }

  self.registration.showNotification(notificationTitle, notificationOptions)
})

self.addEventListener('notificationclick', (event) => {
  if (event.action) {
    clients.openWindow(event.action)
  }
  event.notification.close()
})