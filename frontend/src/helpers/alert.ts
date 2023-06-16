import Swal from 'sweetalert2';

export const alertSuccess = (message:string) => {
  return Swal.fire({
    title: 'Success',
    text: message,
    icon: 'success',
    confirmButtonText: 'OK',
    timer: 1500,
  });
} 

export const alertError = (message:string) => {
  return Swal.fire({
    title: 'Error',
    text: message,
    icon: 'error',
    // showCloseButton: true,
    timer: 1500,
    confirmButtonText: 'OK',
  });
} 

export const alertNotFoundLink = (message:string) => {
  return Swal.fire({
    title: 'Error',
    text: message,
    icon: 'error',
    showConfirmButton: false,
    timer: 1000,
  });
} 
