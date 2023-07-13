import React, { useState, useEffect } from 'react';
import { requestPermission, onMessageListener } from '../../firebase/firebase';
import { toast } from 'react-toastify';

function Notification() {
  const [notification, setNotification] = useState({ title: '', body: '' });
  useEffect(() => {
    requestPermission();
    const unsubscribe = onMessageListener().then((payload:any) => {
      setNotification({
        title: payload?.notification?.title,
        body: payload?.notification?.body,
      });
      toast.info(`${payload?.notification?.title}: ${payload?.notification?.body}`, {
        position: 'top-center',
        autoClose: false,
        hideProgressBar: true,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: 'light',
      });
});
    return () => {
      unsubscribe.catch();
    };
  }, []);
  return (
    <div>
      {/* <Toaster /> */}
    </div>
  );
}
export default Notification;