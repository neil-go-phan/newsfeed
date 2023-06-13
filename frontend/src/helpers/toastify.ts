import { toast } from 'react-toastify';
import { TOASTIFY_TIME } from './constants';

export const toastifyError = (message: string) => {
  toast.error(message, {
    position: 'top-right',
    autoClose: TOASTIFY_TIME,
    hideProgressBar: false,
    closeOnClick: true,
    pauseOnHover: true,
    draggable: true,
    progress: undefined,
    theme: 'light',
  });
};

export const toastifySuccess = (message: string) => {
  toast.success(message, {
    position: 'top-right',
    autoClose: TOASTIFY_TIME,
    hideProgressBar: false,
    closeOnClick: true,
    pauseOnHover: true,
    draggable: true,
    progress: undefined,
    theme: 'light',
  });
};
