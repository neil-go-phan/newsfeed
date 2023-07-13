import { NextPage } from 'next';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faUser } from '@fortawesome/free-regular-svg-icons';
import { faLock } from '@fortawesome/free-solid-svg-icons';
import { Button, Col, Container, Form, InputGroup, Row } from 'react-bootstrap';
import { useContext, useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import { setCookie } from 'cookies-next';
import { useForm, SubmitHandler } from 'react-hook-form';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import cryptoJS from 'crypto-js';
import { _REGEX, _ROUTES } from '@/helpers/constants';
import axiosClient from '@/helpers/axiosClient';
import { IsLoggedContext } from '@/common/contexts/isLoggedContext';
import Link from 'next/link';
import { getGoogleUrl } from '@/helpers/getGoogleUrl';
import Image from 'next/image';
import AuthLayout from '@/layouts/authLayout';
import { checkAuth } from '@/helpers/checkAuth';

const Login: NextPage = () => {
  const router = useRouter();
  const logged = useContext(IsLoggedContext);
  const [errorMessage, setErrorMessage] = useState({
    trigger: false,
    message: '',
  });
  const [auth, setAuth] = useState(false);
  useEffect(() => {
    async function checkLogIn() {
      const userChecked: boolean = await checkAuth();
      setAuth(userChecked);
    }

    checkLogIn();
  }, []);

  useEffect(() => {
    if (auth) {
      if (router.query.redirectTo) {
        router.push(_ROUTES.FEEDS_PLAN);
      } else {
        router.push(_ROUTES.DASHBOARD_PAGE);
      }
    }
  }, [auth]);

  const handleRedirect = (link: string) => {
    if (router.query.redirectTo) {
      const newLink = `${link}?redirectTo=${router.query.redirectTo}`;
      return newLink;
    }
    return link;
  };

  const schema = yup.object().shape({
    username: yup.string().required('Username or email must not be empty'),
    password: yup
      .string()
      .required('Password must not be empty')
      .min(8, 'Password must have 8-16 character')
      .max(16, 'Password must have 8-16 character')
      .matches(
        _REGEX.REGEX_USENAME_PASSWORD,
        'Password must not contain special character like @#$^...'
      ),
  });
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<LoginFormProperty>({
    resolver: yupResolver(schema),
  });

  const onSubmit: SubmitHandler<LoginFormProperty> = async (data) => {
    let { username, password } = data;
    password = cryptoJS.SHA512(password).toString();
    try {
      const res = await axiosClient.post('auth/login', {
        username,
        password,
      });
      setCookie('access_token', res.data.access_token);
      setCookie('refresh_token', res.data.refresh_token);
      setErrorMessage({
        trigger: false,
        message: res.data.message,
      });
      logged?.setIsLogged(true);
      if (router.query.redirectTo) {
        router.push(_ROUTES.FEEDS_PLAN);
      } else {
        router.push(_ROUTES.DASHBOARD_PAGE);
      }
    } catch (error: any) {
      setErrorMessage({
        trigger: true,
        message: error.response.data.message,
      });
    }
    reset({ password: '' });
  };
  if (auth) {
    return (
      <></>
    )
  }
  return (
    <AuthLayout>
      <div className="auth__login min-vh-100 d-flex flex-row align-items-center">
        <Container>
          <Row className="justify-content-center align-items-center px-3">
            <Col lg={8}>
              <Row>
                <Col md={7} className="auth__login--form border p-5">
                  <div className="">
                    <h1>Login</h1>
                    <p className="text-black-50">Login to your account</p>

                    <form onSubmit={handleSubmit(onSubmit)}>
                      <InputGroup className="mb-3">
                        <InputGroup.Text>
                          <FontAwesomeIcon icon={faUser} fixedWidth />
                        </InputGroup.Text>
                        <Form.Control
                          {...register('username')}
                          placeholder="Type your username or email"
                          type="text"
                          required
                        />
                      </InputGroup>

                      {errors.username && (
                        <p className="errorMessage">
                          {errors.username.message}
                        </p>
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
                        <p className="errorMessage">
                          {errors.password.message}
                        </p>
                      )}

                      {errorMessage.trigger && (
                        <p className="errorMessage errorFromServer">
                          {errorMessage.message}
                        </p>
                      )}
                      <Button
                        className="px-4 w-100"
                        variant="primary"
                        type="submit"
                      >
                        Login
                      </Button>
                      <div className="flex items-center my-2 before:flex-1 before:border-t before:border-gray-300 before:mt-0.5 after:flex-1 after:border-t after:border-gray-300 after:mt-0.5">
                        <p className="text-center font-semibold mx-4 mb-0">
                          OR
                        </p>
                      </div>
                      <Button
                        className="px-4 text-white font-medium text-sm leading-snug uppercase rounded shadow-md hover:shadow-lg focus:shadow-lg focus:outline-none focus:ring-0 active:shadow-lg transition duration-150 ease-in-out w-full flex justify-center items-center mb-3 w-100"
                        style={{ backgroundColor: '#3b5998' }}
                        href={getGoogleUrl(
                          handleRedirect(_ROUTES.TOKEN_REDIRECT)
                        )}
                        role="button"
                        data-mdb-ripple="true"
                        data-mdb-ripple-color="light"
                      >
                        <Image
                          src="/images/google.svg"
                          alt="google logo"
                          width={20}
                          height={20}
                        />
                        Continue with Google
                      </Button>
                    </form>
                  </div>
                </Col>
                <Col
                  md={5}
                  className="auth__login--register text-white d-flex align-items-center justify-content-center p-5"
                >
                  <div>
                    <div className="text-center">
                      <h2>Register</h2>
                      <p>Register to experience all of our great features !</p>
                      <Link href={handleRedirect(_ROUTES.REGISTER_PAGE)}>
                        <Button
                          className="btn btn-lg btn-outline-light mt-3"
                          type="button"
                        >
                          Register
                        </Button>
                      </Link>
                    </div>
                  </div>
                </Col>
              </Row>
            </Col>
          </Row>
        </Container>
      </div>
    </AuthLayout>
  );
};

export default Login;
