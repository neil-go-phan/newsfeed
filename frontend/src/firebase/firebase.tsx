import { initializeApp } from 'firebase/app';
import {
  getMessaging,
  getToken,
  isSupported,
  onMessage,
} from 'firebase/messaging';
import firebaseConfig from './firebase.config';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';

const app = initializeApp(firebaseConfig);
// const messaging = getMessaging(app);
const messaging = (await isSupported()) ? getMessaging(app) : null;

export const requestPermission = () => {
  console.log('messaging', messaging)
  if (messaging) {
    Notification.requestPermission().then((permission) => {
      if (permission === 'granted') {
        return getToken(messaging, {
          vapidKey: process.env.NEXT_PUBLIC_FIREBASE_VAPID_KEY,
        })
          .then((currentToken) => {
            if (currentToken) {
              requestCreateToken(currentToken);
            } else {
            }
          })
          .catch((err) => {});
      } else {
      }
    });
  }
};

const requestCreateToken = async (currentToken: string) => {
  try {
    const { data } = await axiosProtectedAPI.get('/notification/create/token', {
      params: {
        token: currentToken,
      },
    });
  } catch (error: any) {}
};

export const onMessageListener = () =>
  new Promise((resolve) => {
    if (messaging) {
      onMessage(messaging, (payload) => {
        resolve(payload);
      });
    }
  });
