import Swal from 'sweetalert2';

export const alertSuccess = (message:string) => {
  return Swal.fire({
    title: 'Success',
    text: message,
    icon: 'success',
    confirmButtonText: 'OK',
  });
} 
