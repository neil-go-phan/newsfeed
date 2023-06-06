import { NextPage } from 'next';
import { faUser } from '@fortawesome/free-regular-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faEnvelope, faLock } from '@fortawesome/free-solid-svg-icons';
import {
  Button,
  Card,
  Col,
  Container,
  Form,
  InputGroup,
  Row,
} from 'react-bootstrap';
import { useRouter } from 'next/router';
import { useState } from 'react';
import { useForm, SubmitHandler } from 'react-hook-form';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import cryptoJS from 'crypto-js';
import { _REGEX, _ROUTES } from '@/helpers/constants';
import axiosClient from '@/helpers/axiosClient';
import AuthLayout from '@/layouts/authLayout';

interface RegisterFormProperty {
  email: string;
  username: string;
  password: string;
  passwordConfirmation?: string;
}

const Register: NextPage = () => {
  const router = useRouter();

  const [errorMessage, setErrorMessage] = useState({
    trigger: false,
    message: '',
  });
  const schema = yup.object().shape({
    email: yup
      .string()
      .required('Email name must not be empty')
      .email('Email wrong format'),
    username: yup
      .string()
      .required('Username must not be empty')
      .min(8, 'Username must have 8-16 character')
      .max(16, 'Username must have 8-16 character')
      .matches(
        _REGEX.REGEX_USENAME_PASSWORD,
        'Username must not contain special character like @#$^...'
      ),
    password: yup
      .string()
      .required('Password must not be empty')
      .min(8, 'Password must have 8-16 character')
      .max(16, 'Password must have 8-16 character')
      .matches(
        _REGEX.REGEX_USENAME_PASSWORD,
        'Password must not contain special character like @#$^...'
      ),
    passwordConfirmation: yup
      .string()
      .oneOf([yup.ref('password')], 'Passwords not match'),
  });
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegisterFormProperty>({
    resolver: yupResolver(schema),
  });

  const onSubmit: SubmitHandler<RegisterFormProperty> = async (data) => {
    let { username, password, passwordConfirmation, email } = data;
    password = cryptoJS.SHA512(password).toString();
    if (passwordConfirmation) {
      passwordConfirmation = cryptoJS.SHA512(passwordConfirmation).toString();
    }
    try {
      const res = await axiosClient.post('auth/register', {
        username,
        password,
        password_confirmation: passwordConfirmation,
        email,
      });
      if (!res.data.success) {
        setErrorMessage({ trigger: true, message: res.data.message });
      } else {
        router.push(_ROUTES.LOGIN_PAGE);
        setErrorMessage({ trigger: false, message: res.data.message });
      }
    } catch (error: any) {
      setErrorMessage({
        trigger: true,
        message: error.response.data.message,
      });
    }
  };

  return (
    <AuthLayout>
      <div className="auth__register min-vh-100 d-flex flex-row align-items-center dark:bg-transparent">
        <Container>
          <Row className="justify-content-center">
            <Col md={6}>
              <Card className="mb-4 rounded-0 auth__register--form">
                <Card.Body className="p-4">
                  <h1>Register</h1>
                  <p className="text-black-50">Create your account</p>

                  <form onSubmit={handleSubmit(onSubmit)}>
                    <InputGroup className="mb-3">
                      <InputGroup.Text>
                        <FontAwesomeIcon icon={faEnvelope} fixedWidth />
                      </InputGroup.Text>
                      <Form.Control
                        {...register('email')}
                        placeholder="Type your email"
                        type="text"
                        required
                      />
                    </InputGroup>
                    {errors.email && (
                      <p className="errorMessage">{errors.email.message}</p>
                    )}
                    <InputGroup className="mb-3">
                      <InputGroup.Text>
                        <FontAwesomeIcon icon={faUser} fixedWidth />
                      </InputGroup.Text>
                      <Form.Control
                        {...register('username')}
                        placeholder="Type your username"
                        type="text"
                        required
                      />
                    </InputGroup>

                    {errors.username && (
                      <p className="errorMessage">{errors.username.message}</p>
                    )}

                    <InputGroup className="mb-3">
                      <InputGroup.Text>
                        <FontAwesomeIcon icon={faLock} fixedWidth />
                      </InputGroup.Text>
                      <Form.Control
                        {...register('password')}
                        placeholder="Type your password"
                        type="password"
                        required
                      />
                    </InputGroup>

                    {errors.password && (
                      <p className="errorMessage">{errors.password.message}</p>
                    )}

                    <InputGroup className="mb-3">
                      <InputGroup.Text>
                        <FontAwesomeIcon icon={faLock} fixedWidth />
                      </InputGroup.Text>
                      <Form.Control
                        {...register('passwordConfirmation')}
                        placeholder="Type your password again"
                        type="password"
                        required
                      />
                    </InputGroup>

                    {errors.passwordConfirmation && (
                      <p className="errorMessage">
                        {errors.passwordConfirmation.message}
                      </p>
                    )}
                    {errorMessage.trigger && (
                      <p className="errorMessage">{errorMessage.message}</p>
                    )}

                    <Button
                      type="submit"
                      className="d-block w-100"
                      variant="success"
                    >
                      Create Account
                    </Button>
                  </form>
                </Card.Body>
              </Card>
            </Col>
          </Row>
        </Container>
      </div>
    </AuthLayout>
  );
};

export default Register;
