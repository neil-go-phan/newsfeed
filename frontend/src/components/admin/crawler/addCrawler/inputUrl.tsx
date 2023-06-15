import React from 'react';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Button, Form, InputGroup } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faLink } from '@fortawesome/free-solid-svg-icons';
import Link from 'next/link';
import { _ROUTES } from '@/helpers/constants';
import { useRouter } from 'next/router';

type Props = {
  // eslint-disable-next-line no-unused-vars
  handleInputUrl: (url: string) => void;
};

const InputUrl: React.FC<Props> = (props: Props) => {
  const router = useRouter()
  const schema = yup.object().shape({
    url: yup
      .string()
      .url('Enter correct url!')
      .required('Please enter website'),
  });
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<UrlFormProperty>({
    resolver: yupResolver(schema),
  });
  const onSubmit: SubmitHandler<UrlFormProperty> = async (data) => {
    let { url } = data;
    props.handleInputUrl(url);
  };
  const onCreateCusomCrawler:SubmitHandler<UrlFormProperty> = async (data) => {
    let { url } = data;
    router.push(`${_ROUTES.ADD_CUSTOM_CRAWLER}?url=${url}`)
  };
  return (
    <div className="addCrawler__inputUrl">
      <div className="addCrawler__inputUrl--line" />
      <form onSubmit={handleSubmit(onSubmit)}>
        <h2 className="addCrawler__inputUrl--title">Input url</h2>

        <label> URL </label>
        <InputGroup className="mb-3">
          <InputGroup.Text>
            <FontAwesomeIcon icon={faLink} fixedWidth />
          </InputGroup.Text>
          <Form.Control
            {...register('url')}
            placeholder="Type url"
            type="text"
            required
          />
        </InputGroup>

        {errors.url && <p className="errorMessage">{errors.url.message}</p>}
        <Button className="px-4 mx-3" variant="success" type="submit">
          Continue
        </Button>
          <Button className="px-4" variant="warning" onClick={handleSubmit(onCreateCusomCrawler)}>
            Add custom crawler
          </Button>
      </form>
      <div className="addCrawler__inputUrl--line" />
    </div>
  );
};

export default InputUrl;
