import { initializeApp } from 'firebase/app';
import { getMessaging, getToken, onMessage } from 'firebase/messaging';
import firebaseConfig from './firebase.config';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';

const app = initializeApp(firebaseConfig);
const messaging = getMessaging(app);

export const requestPermission = () => {
  console.log('Requesting User Permission......');
  Notification.requestPermission().then((permission) => {
    if (permission === 'granted') {
      console.log('Notification User Permission Granted.');
      return getToken(messaging, {
        vapidKey: process.env.NEXT_PUBLIC_FIREBASE_VAPID_KEY,
      })
        .then((currentToken) => {
          if (currentToken) {
            console.log('Client Token: ', currentToken);
            requestCreateToken(currentToken)
          } else {
            console.log('Failed to generate the app registration token.');
          }
        })
        .catch((err) => {
          console.log(
            'An error occurred when requesting to receive the token.',
            err
          );
        });
    } else {
      console.log('User Permission Denied.');
    }
  });
};

const requestCreateToken = async (currentToken: string) => {
  try {
    const { data } = await axiosProtectedAPI.get(
      '/notification/create/token',
      {
        params: {
          token: currentToken,
        },
      }
    );
  } catch (error: any) {}
};

export const onMessageListener = () =>
  new Promise((resolve) => {
    onMessage(messaging, (payload) => {
      resolve(payload);
    });
  });
