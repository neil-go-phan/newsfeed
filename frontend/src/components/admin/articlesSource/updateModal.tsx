import React, { useEffect, useMemo, useState } from 'react';
import * as yup from 'yup';
import Image from 'next/image';
import { yupResolver } from '@hookform/resolvers/yup';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Button, Form, InputGroup } from 'react-bootstrap';
import FilterByTopic from './filterByTopic';

type ArticlesSourceYupValidateProp = {
  title: string;
  description: string;
  imgSize: number;
};

type Props = {
  articlesSource: ArticlesSourceInfo;
  handleUpdate: (articlesSource: UpdateArticleSourcePayload) => void;
};

const IMAGE_SIZE_PIXEL = 200;
// 1mb
const IMAGE_FILE_SIZE_BYTES = 1000000;
const DEFAULT_IMG_SIZE = 0;
const DEFAULT_TOPIC_ID = 0;

const UpdateModal: React.FC<Props> = (props: Props) => {
  const [image, setImage] = useState<string>();
  const [topicID, setTopicID] = useState<number>(0);
  const schema = yup.object().shape({
    imgSize: yup
      .number()
      .required('logo image must not be empty')
      .max(IMAGE_FILE_SIZE_BYTES, 'image size must be less than 100kb')
      .min(DEFAULT_IMG_SIZE + 1, 'please submit logo'),
    title: yup.string().required('title must not be empty'),
    description: yup.string().required('description must not be empty'),
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
        imgSize: DEFAULT_IMG_SIZE,
        topicID: DEFAULT_TOPIC_ID,
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
    const articlesSource: UpdateArticleSourcePayload = {
      id: props.articlesSource.id,
      title: data.title,
      description: data.description,
      image: base64Img,
      topic_id: topicID
    };
    props.handleUpdate(articlesSource);
  };

  const handleChooseTopicID = (id: number) => {
    setTopicID(id);
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
      if (props.articlesSource.image !== '') {
        setImage(props.articlesSource.image);
        setValue('imgSize', 100);
      }
      
    }
  }, []);

  return (
    <div className="adminArticlesSources__modal">
      <form onSubmit={handleSubmit(onSubmit)}>
      <h2 className="adminArticlesSources__modal--title">Update articles source</h2>
        <div className="adminArticlesSources__modal--line" />
        <div className="field">
          <label> Title </label>
          <InputGroup className="mb-3">
            <Form.Control
              {...register('title')}
              placeholder="Please type article source title"
              type="text"
              required
            />
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
            />
          </InputGroup>

          {errors.description && (
            <p className="errorMessage">{errors.description.message}</p>
          )}
        </div>
        <div className="field">
          <label> Feed link </label>
          <InputGroup className="mb-3">
            <Form.Control
              type="text"
              disabled={true}
              value={props.articlesSource.feed_link}
            />
          </InputGroup>
        </div>
        <div className="field">
          <label> Source link </label>
          <InputGroup className="mb-3">
            <Form.Control
              type="text"
              disabled={true}
              value={props.articlesSource.link}
            />
          </InputGroup>
        </div>

        <div className="field">
          <label> Topic </label>
          <FilterByTopic handleChooseTopicID={handleChooseTopicID} />
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
          Update articles source
        </Button>
      </form>
    </div>
  );
};

export default UpdateModal;
