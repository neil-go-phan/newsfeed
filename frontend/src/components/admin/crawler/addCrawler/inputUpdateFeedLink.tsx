import React, { useEffect, useState } from 'react';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Button, Form, InputGroup } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faLink } from '@fortawesome/free-solid-svg-icons';
import { _ROUTES } from '@/helpers/constants';
import { useRouter } from 'next/router';

type Props = {
  // eslint-disable-next-line no-unused-vars
  handleSubmitNewFeedLink: (url: string) => void;
  handleTestCrawler: () => void;
  isUpdate: boolean;
  feedLink: string | undefined;
};

const InputUpdateFeedLink: React.FC<Props> = (props: Props) => {
  const router = useRouter();
  const [crawlerID, setCrawlerID] = useState<number>();
  const [sourceLink, setSourceLink] = useState<string>('');

  const schema = yup.object().shape({
    url: yup
      .string()
      .url('Enter correct url!')
      .required('Please enter website'),
  });
  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors },
  } = useForm<UrlFormProperty>({
    resolver: yupResolver(schema),
  });
  const onSubmit: SubmitHandler<UrlFormProperty> = async (data) => {
    let { url } = data;
    props.handleSubmitNewFeedLink(url);
  };

  const onUpdateCusomCrawler: SubmitHandler<UrlFormProperty> = async (data) => {
    let { url } = data;
    router.push({
      pathname: _ROUTES.ADD_CUSTOM_CRAWLER,
      query: { url: url, id: crawlerID },
    });
  };

  useEffect(() => {
    if (router.query.source_link) {
      const url = router.query.source_link as string;
      setSourceLink(url);
    }
    if (router.query.id) {
      setCrawlerID(+router.query.id);
    }
  }, []);

  useEffect(() => {
    if (props.feedLink) {
      setValue('url', props.feedLink);
    }
  }, [props.feedLink]);

  return (
    <div className="addCrawler__inputUrl">
      <div className="addCrawler__inputUrl--line" />
      <form onSubmit={handleSubmit(onSubmit)}>
        <h2 className="addCrawler__inputUrl--title"></h2>
        <label> Source link </label>
        <InputGroup className="mb-3">
          <InputGroup.Text>
            <FontAwesomeIcon icon={faLink} fixedWidth />
          </InputGroup.Text>
          <Form.Control
            placeholder="Type url"
            type="text"
            required
            disabled={true}
            value={sourceLink}
          />
        </InputGroup>
        <label> Feed link </label>
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
          Scan RSS
        </Button>
        <Button
          className="px-4 mx-3"
          variant="secondary"
          onClick={props.handleTestCrawler}
        >
          Test
        </Button>
        <Button
          className="px-4"
          variant="warning"
          onClick={handleSubmit(onUpdateCusomCrawler)}
        >
          Custom crawler
        </Button>
      </form>
      <div className="addCrawler__inputUrl--line" />
    </div>
  );
};

export default InputUpdateFeedLink;
