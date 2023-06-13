import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import React, { useEffect, useState } from 'react';
import cryptoJS from 'crypto-js';
import { _REGEX } from '@/helpers/constants';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Button, Form, InputGroup } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faUser } from '@fortawesome/free-regular-svg-icons';
import { faLock } from '@fortawesome/free-solid-svg-icons';
import Swal from 'sweetalert2';
import { toastifyError } from '@/helpers/toastify';

type Props = {
  handleIsProfileClose: () => void;
}

const ERROR_MESSAGE_WHEN_GET_ADMIN_USERNAME = 'Error occurred while get admin username'

const AdminProfile: React.FC<Props> = (props: Props) => {
  const [errorMessage, setErrorMessage] = useState({
    trigger: false,
    message: '',
  });
  const [adminUsername, setAdminUsername] = useState()


  useEffect(() => {
    const requestGetAdminUsername = async () => {
      try {
        const { data } = await axiosProtectedAPI.get('auth/get', {
        });
        setAdminUsername(data.username);
      } catch (error) {
        toastifyError(ERROR_MESSAGE_WHEN_GET_ADMIN_USERNAME)
      }
    };
    requestGetAdminUsername();
  
  }, [])
  
  const schema = yup.object().shape({
    password: yup
      .string()
      .required('Password must not be empty')
      .min(8, 'Password must have 8-16 character')
      .max(16, 'Password must have 8-16 character')
      .matches(
        _REGEX.REGEX_USENAME_PASSWORD,
        'Password must not contain special character like @#$^...'
      ),
      // eslint-disable-next-line camelcase
    password_confirmation: yup
      .string()
      .oneOf([yup.ref('password')], 'Passwords not match'),
  });
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<ChangePasswordFormProperty>({
    resolver: yupResolver(schema),
  });
  const onSubmit: SubmitHandler<ChangePasswordFormProperty> = async (data) => {
    let { password, password_confirmation } = data;
    password = cryptoJS.SHA512(password).toString();
    // eslint-disable-next-line camelcase
    if (password_confirmation) {
      // eslint-disable-next-line camelcase
      password_confirmation = cryptoJS.SHA512(password_confirmation).toString();
    }
    try {
      const res = await axiosProtectedAPI.post(
        'auth/change-password',
        {
          username: adminUsername,
          password,
          // eslint-disable-next-line camelcase
          password_confirmation,
        }
      );
      setErrorMessage({
        trigger: false,
        message: res.data.message,
      });
      props.handleIsProfileClose();
      Swal.fire({
        title: 'Success',
        text: res.data.message,
        icon: 'success',
        confirmButtonText: 'OK',
      });
    } catch (error: any) {
      setErrorMessage({
        trigger: true,
        message: error.response.data.message,
      });
    }
    reset({ password: '' });
  };
  return (
    <div className='adminProfile'>
      <form onSubmit={handleSubmit(onSubmit)}>
        <h2 className='adminProfile__title'>Admin profile</h2>
        <div className='adminProfile__line' />
        <label> Username </label>
        <InputGroup className='mb-3'>
          <InputGroup.Text>
            <FontAwesomeIcon icon={faUser} fixedWidth />
          </InputGroup.Text>
          <Form.Control value={adminUsername} type='text' disabled />
        </InputGroup>

        <label> Password </label>
        <InputGroup className='mb-3'>
          <InputGroup.Text>
            <FontAwesomeIcon icon={faLock} fixedWidth />
          </InputGroup.Text>
          <Form.Control
            {...register('password')}
            placeholder='Type your password'
            type='password'
            required
          />
        </InputGroup>

        {errors.password && (
          <p className='errorMessage'>{errors.password.message}</p>
        )}

        <label> Confirm password </label>
        <InputGroup className='mb-3'>
          <InputGroup.Text>
            <FontAwesomeIcon icon={faLock} fixedWidth />
          </InputGroup.Text>
          <Form.Control
            {...register('password_confirmation')}
            placeholder='Confirm your password'
            type='password'
            required
          />
        </InputGroup>
        {errors.password_confirmation && (
          <p className='errorMessage'>{errors.password_confirmation.message}</p>
        )}

        {errorMessage.trigger && (
          <p className='errorMessage errorFromServer'>{errorMessage.message}</p>
        )}

        <Button
          className='w-100 px-4'
          variant='primary'
          type='submit'
        >
          Change your password
        </Button>
      </form>
    </div>
  );
};

export default AdminProfile;
