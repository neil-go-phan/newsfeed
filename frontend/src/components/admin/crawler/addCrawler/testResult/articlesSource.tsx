import { toDataUrl } from '@/helpers/imgUrlToData';
import axios from 'axios';
import Image from 'next/image';
import React, { useEffect, useState } from 'react';
import { Button, Form, InputGroup } from 'react-bootstrap';

type Props = {
  articlesSource: ArticlesSource;
  url: string;
};

type Field = {
  value: string;
  disable: boolean;
};

type ImageField = {
  value: string | ArrayBuffer | null;
};

const IMAGE_SIZE = 200;

const ArticlesSource: React.FC<Props> = (props: Props) => {
  const [image, setImage] = useState<string>();
  const [title, setTitle] = useState<Field>({
    value: props.articlesSource.title,
    disable: true,
  });
  const [description, setDescription] = useState<Field>({
    value: props.articlesSource.description,
    disable: true,
  });
  const [feedLink, setFeedLink] = useState<Field>({
    value: props.articlesSource.feed_link,
    disable: true,
  });
  const [link, setLink] = useState<Field>({
    value: props.articlesSource.link,
    disable: true,
  });

  useEffect(() => {
    if (props.articlesSource) {
      toDataUrl(
        props.articlesSource.image,
        (base64: string | ArrayBuffer | null) => {
          if (base64) {
            const str = base64.toString();
            setImage(str);
          }
        }
      );
    }
  }, []);

  if (props.articlesSource) {
    return (
      <div className="addCrawler__testResult--articles_source">
        <div className="title">
          <h3>Articles source</h3>
        </div>
        <div className="info">
          <div className="field">
            <label> Title </label>
            <InputGroup className="mb-3">
              <Form.Control
                placeholder="Article source title"
                type="text"
                required
                className={title.disable ? 'bg-secondary' : 'bg-light'}
                disabled={title.disable}
                value={title.value}
                onChange={(event) =>
                  setTitle({ ...title, value: event.target.value })
                }
              />
              <Button
                className="px-4"
                variant={title.disable ? 'primary' : 'success'}
                onClick={() => {
                  setTitle({ ...title, disable: !title.disable });
                }}
              >
                {title.disable ? 'Edit' : 'OK'}
              </Button>
            </InputGroup>
          </div>
          <div className="field">
            <label> Description </label>
            <InputGroup className="mb-3">
              <Form.Control
                placeholder="Article source description"
                type="text"
                required
                className={description.disable ? 'bg-secondary' : 'bg-light'}
                disabled={description.disable}
                value={description.value}
                onChange={(event) =>
                  setDescription({ ...description, value: event.target.value })
                }
              />
              <Button
                className="px-4"
                variant={description.disable ? 'primary' : 'success'}
                onClick={() => {
                  setDescription({
                    ...description,
                    disable: !description.disable,
                  });
                }}
              >
                {description.disable ? 'Edit' : 'OK'}
              </Button>
            </InputGroup>
          </div>
          <div className="field">
            <label> Feed link </label>
            <InputGroup className="mb-3">
              <Form.Control
                placeholder="Article source feed link"
                type="text"
                required
                className={feedLink.disable ? 'bg-secondary' : 'bg-light'}
                disabled={feedLink.disable}
                value={feedLink.value}
                onChange={(event) =>
                  setFeedLink({ ...feedLink, value: event.target.value })
                }
              />
              <Button
                className="px-4"
                variant={feedLink.disable ? 'primary' : 'success'}
                onClick={() => {
                  setFeedLink({ ...feedLink, disable: !feedLink.disable });
                }}
              >
                {feedLink.disable ? 'Edit' : 'OK'}
              </Button>
            </InputGroup>
          </div>
          <div className="field">
            <label> Source link </label>
            <InputGroup className="mb-3">
              <Form.Control
                placeholder="Article source link"
                type="text"
                required
                className={link.disable ? 'bg-secondary' : 'bg-light'}
                disabled={link.disable}
                value={link.value}
                onChange={(event) =>
                  setLink({ ...link, value: event.target.value })
                }
              />
              <Button
                className="px-4"
                variant={link.disable ? 'primary' : 'success'}
                onClick={() => {
                  setLink({ ...link, disable: !link.disable });
                }}
              >
                {link.disable ? 'Edit' : 'OK'}
              </Button>
            </InputGroup>
          </div>
          <div className="field">
            <label> Image </label>
            <div className="sourceLogo">
              {image ? (
                <Image
                  alt="article source logo"
                  src={image}
                  width={IMAGE_SIZE}
                  height="0"
                  style={{ height: 'auto' }}
                />
              ) : (
                <div>Not found article source logo, please add one</div>
              )}
              <input type="file" name="file" />
            </div>
          </div>
        </div>
      </div>
    );
  }
  return <div className="adminCrawler__testResult--articles_source"></div>;
};

export default ArticlesSource;
