import React, { useCallback, useRef, useState } from 'react';
import { useRouter } from 'next/router';
import * as yup from 'yup';
import { Button, Form, InputGroup } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faHeading,
  faInfo,
  faLink,
  faTag,
  faUser,
} from '@fortawesome/free-solid-svg-icons';
import { _ROUTES } from '@/helpers/constants';
import EmbedWeb from './embedWeb';
import { CRAWLER_CUSTOM_TYPE } from '.';

type Props = {
  requestCustomCrawlerTest: (crawler: Crawler) => void;
};

const ChooseHtmlClassForm: React.FC<Props> = (props: Props) =>{
  const [errorMessage, setErrorMessage] = useState({
    trigger: false,
    message: '',
  });
  const [divClass, setDivClass] = useState<string>('');
  const [titleClass, setTitleClass] = useState<string>('');
  const [descriptionClass, setDescriptionClass] = useState<string>('');
  const [linkClass, setLinkClass] = useState<string>('');
  const [authorsClass, setAuthorsClass] = useState<string>('');
  const router = useRouter();
  const { url } = router.query;

  const fieldChooseRef = useRef('');

  const handleChoose = (fieldChoosed: string) => {
    fieldChooseRef.current = fieldChoosed;
  };

  // JavaScript closure
  const handleClick = useCallback((event: Event): void => {
    const target = event.target as HTMLElement;
    const classname = target.className.trim();
    switch (fieldChooseRef.current) {
      case 'div':
        setDivClass(classname);
        break;
      case 'title':
        setTitleClass(classname);
        break;
      case 'description':
        setDescriptionClass(classname);
        break;
      case 'link':
        setLinkClass(classname);
        break;
      case 'authors':
        setAuthorsClass(classname);
        break;
      default:
        break;
    }
  }, []);

  const schema = yup.object().shape({
    article_div: yup.string().required('Article div is require'),
    article_title: yup.string().required('Article title is require'),
    article_link: yup.string().required('Article link is require'),
  });

  const handleTest = async (
    e: React.MouseEvent<HTMLButtonElement, MouseEvent>
  ) => {
    e.preventDefault();
    // validate
    try {
      await schema.validate(
        {
          article_div: divClass,
          article_title: titleClass,
          article_link: linkClass,
        },
        { abortEarly: true }
      );
      const crawler: Crawler = {
        crawl_type: CRAWLER_CUSTOM_TYPE,
        source_link: String(url),
        article_div: divClass,
        article_title: titleClass,
        article_link: linkClass,
        article_description: descriptionClass,
        article_authors: authorsClass,
        feed_link: '',
        schedule: '',
        articles_source_id: 0,
      }
      props.requestCustomCrawlerTest(crawler);
      setErrorMessage({
        trigger: false,
        message: '',
      });
    } catch (error) {
      if (error instanceof yup.ValidationError) {
        setErrorMessage({
          trigger: true,
          message: error.message,
        });
      }
    }
  };

  return (
    <div className="adminCrawler__addCrawler ">
      <div className="addCustomCrawler__embedweb">
        <div className="lineWraper">
          <div className="line"></div>
        </div>
        <div className="title">
          <h4>Site interface</h4>
        </div>
        <EmbedWeb url={String(url)} handleClick={handleClick} />
      </div>
      <div className="adminCrawler__addCrawler--add">
        <div className="adminCrawler__addCrawler--form">
          <form>
            <div className="line" />
            <h2 className="title">Input article class</h2>
            <div className="line" />
            <label> Article div <span className='colorRed'>*</span> </label>
            <InputGroup className="my-3">
              <InputGroup.Text>
                <FontAwesomeIcon icon={faTag} fixedWidth />
              </InputGroup.Text>
              <Form.Control
                placeholder="Type of choose article div classname"
                type="text"
                required
                className="bg-white"
                value={divClass}
                onChange={(event) => setDivClass(event.target.value)}
              />
              <Button
                className="px-4"
                variant="primary"
                onClick={() => handleChoose('div')}
              >
                Choose
              </Button>
            </InputGroup>

            <label> Article title <span className='colorRed'>*</span></label>
            <InputGroup className="my-3">
              <InputGroup.Text>
                <FontAwesomeIcon icon={faHeading} fixedWidth />
              </InputGroup.Text>
              <Form.Control
                placeholder="Type of choose article title classname"
                type="text"
                required
                className="bg-white"
                value={titleClass}
                onChange={(event) => setTitleClass(event.target.value)}
              />
              <Button
                className="px-4"
                variant="primary"
                onClick={() => handleChoose('title')}
              >
                Choose
              </Button>
            </InputGroup>

            <label> Article description </label>
            <InputGroup className="my-3">
              <InputGroup.Text>
                <FontAwesomeIcon icon={faInfo} fixedWidth />
              </InputGroup.Text>
              <Form.Control
                placeholder="Type of choose article description classname"
                type="text"
                required
                className="bg-white"
                value={descriptionClass}
                onChange={(event) => setDescriptionClass(event.target.value)}
              />
              <Button
                className="px-4"
                variant="primary"
                onClick={() => handleChoose('description')}
              >
                Choose
              </Button>
            </InputGroup>

            <label> Article link <span className='colorRed'>*</span></label>
            <InputGroup className="my-3">
              <InputGroup.Text>
                <FontAwesomeIcon icon={faLink} fixedWidth />
              </InputGroup.Text>
              <Form.Control
                className="bg-white"
                placeholder="Type of choose article link classname"
                type="text"
                required
                value={linkClass}
                onChange={(event) => setLinkClass(event.target.value)}
              />
              <Button
                className="px-4"
                variant="primary"
                onClick={() => handleChoose('link')}
              >
                Choose
              </Button>
            </InputGroup>

            <label> Article authors </label>
            <InputGroup className="my-3">
              <InputGroup.Text>
                <FontAwesomeIcon icon={faUser} fixedWidth />
              </InputGroup.Text>
              <Form.Control
                className="bg-white"
                placeholder="Type of choose article authors classname"
                type="text"
                required
                value={authorsClass}
                onChange={(event) => setAuthorsClass(event.target.value)}
              />
              <Button
                className="px-4"
                variant="primary"
                onClick={() => handleChoose('authors')}
              >
                Choose
              </Button>
            </InputGroup>

            {errorMessage.trigger && (
              <p className="errorMessage errorFromServer">
                {errorMessage.message}
              </p>
            )}

            <div className="btnGroup d-flex justify-content-between my-4">
              <Button
                className="px-4 w-25"
                variant="secondary"
                onClick={(e) => handleTest(e)}
                type="submit"
              >
                Test
              </Button>
            </div>
            <div className="line" />
          </form>
        </div>
      </div>
    </div>
  );
};

export default ChooseHtmlClassForm;
