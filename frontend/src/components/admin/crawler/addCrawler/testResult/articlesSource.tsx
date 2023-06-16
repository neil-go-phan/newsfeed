import { toDataUrl } from '@/helpers/imgUrlToData';
import Image from 'next/image';
import React, { useEffect, useMemo, useState } from 'react';
import { Button, Form, InputGroup } from 'react-bootstrap';
import { useForm, SubmitHandler } from 'react-hook-form';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';

type Props = {
  articlesSource: ArticlesSource;
  url: string;
  handleSubmit: (articlesSource: ArticlesSource) => void;
};

const IMAGE_SIZE_PIXEL = 200;
// 1mb
const IMAGE_FILE_SIZE_BYTES = 1000000; 
const DEFAULT_IMG_SIZE = 0;

const ArticlesSource: React.FC<Props> = (props: Props) => {
  const [image, setImage] = useState<string>();
  const [isTitleBtnDisable, setIsTitleBtnDisable] = useState<boolean>(true);
  const [isDescriptionBtnDisable, setIsDescriptionBtnDisable] =
    useState<boolean>(true);
  const [isFeedLinkBtnDisable, setIsFeedLinkBtnDisable] =
    useState<boolean>(true);
  const [isLinkBtnDisable, setIsLinkBtnDisable] = useState<boolean>(true);
  const schema = yup.object().shape({
    imgSize: yup
      .number()
      .required('logo image must not be empty')
      .max(IMAGE_FILE_SIZE_BYTES, 'image too big')
      .min(DEFAULT_IMG_SIZE + 1, 'please submit logo'),
    title: yup.string().required('title must not be empty'),
    description: yup.string().required('description must not be empty'),
    feed_link: yup.string().required('feed link must not be empty').url(),
    link: yup.string().required('link to website must not be empty').url(),
  });

  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors },
  } = useForm<ArticlesSourceYupValidateProp>({
    resolver: yupResolver(schema),
    defaultValues: useMemo(() => {
      return {
        title: props.articlesSource.title,
        description: props.articlesSource.description,
        feed_link: props.articlesSource.feed_link,
        link: props.articlesSource.link,
        imgSize: DEFAULT_IMG_SIZE,
      };
    }, [props]),
  });

  const onSubmit: SubmitHandler<ArticlesSourceYupValidateProp> = async (
    data
  ) => {
    let base64Img: string = '';
    if (image) {
      base64Img = image;
    }
    const articlesSource: ArticlesSource = {
      title: data.title,
      description: data.description,
      feed_link: data.feed_link,
      link: data.link,
      image: base64Img,
    };
    props.handleSubmit(articlesSource);
  };

  const photoUpload = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();
    const reader = new FileReader();
    if (e.target.files) {
      const file = e.target.files[0];
      if (reader !== undefined && file !== undefined) {
        reader.onloadend = () => {
          setImage(reader.result?.toString());
          setValue('imgSize', file.size);
        };
        reader.readAsDataURL(file);
      }
    }
  };

  useEffect(() => {
    if (props.articlesSource) {
      toDataUrl(
        props.articlesSource.image,
        (base64: string | ArrayBuffer | null) => {
          if (base64) {
            const str = base64.toString();
            const buffer = Buffer.from(str.substring(str.indexOf(',') + 1));
            setValue('imgSize', buffer.length);
            setImage(str);
          }
        }
      );
    }
  }, []);

  if (props.articlesSource) {
    return (
      <div className="info">
        <form onSubmit={handleSubmit(onSubmit)}>
          <div className="field">
            <label> Title </label>
            <InputGroup className="mb-3">
              <Form.Control
                {...register('title')}
                placeholder="Please type article source title"
                type="text"
                required
                className={isTitleBtnDisable ? 'bgDisable' : 'bg-light'}
                disabled={isTitleBtnDisable}
              />
              <Button
                className="px-4"
                variant={isTitleBtnDisable ? 'primary' : 'success'}
                onClick={() => {
                  setIsTitleBtnDisable(!isTitleBtnDisable);
                }}
              >
                {isTitleBtnDisable ? 'Edit' : 'OK'}
              </Button>
            </InputGroup>

            {errors.title && (
              <p className="errorMessage">{errors.title.message}</p>
            )}
          </div>
          <div className="field">
            <label> Description </label>
            <InputGroup className="mb-3">
              <Form.Control
                {...register('description')}
                placeholder="Please type article source description"
                type="text"
                required
                className={isDescriptionBtnDisable ? 'bgDisable' : 'bg-light'}
                disabled={isDescriptionBtnDisable}
              />
              <Button
                className="px-4"
                variant={isDescriptionBtnDisable ? 'primary' : 'success'}
                onClick={() => {
                  setIsDescriptionBtnDisable(!isDescriptionBtnDisable);
                }}
              >
                {isDescriptionBtnDisable ? 'Edit' : 'OK'}
              </Button>
            </InputGroup>

            {errors.description && (
              <p className="errorMessage">{errors.description.message}</p>
            )}
          </div>
          <div className="field">
            <label> Feed link </label>
            <InputGroup className="mb-3">
              <Form.Control
                {...register('feed_link')}
                placeholder="Please type article source feed link"
                type="text"
                required
                className={isFeedLinkBtnDisable ? 'bgDisable' : 'bg-light'}
                disabled={isFeedLinkBtnDisable}
              />
              <Button
                className="px-4"
                variant={isFeedLinkBtnDisable ? 'primary' : 'success'}
                onClick={() => {
                  setIsFeedLinkBtnDisable(!isFeedLinkBtnDisable);
                }}
              >
                {isFeedLinkBtnDisable ? 'Edit' : 'OK'}
              </Button>
            </InputGroup>

            {errors.feed_link && (
              <p className="errorMessage">{errors.feed_link.message}</p>
            )}
          </div>
          <div className="field">
            <label> Source link </label>
            <InputGroup className="mb-3">
              <Form.Control
                {...register('link')}
                placeholder="Please type article source link"
                type="text"
                required
                className={isLinkBtnDisable ? 'bgDisable' : 'bg-light'}
                disabled={isLinkBtnDisable}
              />
              <Button
                className="px-4"
                variant={isLinkBtnDisable ? 'primary' : 'success'}
                onClick={() => {
                  setIsLinkBtnDisable(!isLinkBtnDisable);
                }}
              >
                {isLinkBtnDisable ? 'Edit' : 'OK'}
              </Button>
            </InputGroup>

            {errors.link && (
              <p className="errorMessage">{errors.link.message}</p>
            )}
          </div>
          <div className="field">
            <label> Image </label>
            <div className="sourceLogo d-flex">
              <div className="col-6">
                {image ? (
                  <Image
                    alt="article source logo"
                    src={image}
                    width={IMAGE_SIZE_PIXEL}
                    height="0"
                    style={{ height: 'auto' }}
                  />
                ) : (
                  <div>Not found article source logo, please add one</div>
                )}
              </div>
              <div className="col-6">
                <input
                  type="file"
                  name="source-logo"
                  id="file"
                  accept=".jpef, .png, .jpg"
                  onChange={(event) => photoUpload(event)}
                  src={image}
                />
              </div>
            </div>
            {errors.imgSize && (
              <p className="errorMessage">{errors.imgSize.message}</p>
            )}
          </div>

          <Button className="px-4 m-3" variant="primary" type="submit">
            Create crawler
          </Button>
        </form>
      </div>
    );
  }
  return <div className="info"></div>;
};

export default ArticlesSource;
