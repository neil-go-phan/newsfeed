import { toDataUrl } from '@/helpers/imgUrlToData';
import Image from 'next/image';
import React, { useEffect, useMemo, useState } from 'react';
import { Button, Form, InputGroup } from 'react-bootstrap';
import { useForm, SubmitHandler } from 'react-hook-form';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { alertError } from '@/helpers/alert';

type ArticlesSourceYupValidateProp = {
  title: string;
  description: string;
  link: string;
  feed_link: string;
  imgSize: number;
  topicName: string;
};

type Props = {
  articlesSource: ArticlesSource;
  url: string;
  handleSubmit: (articlesSource: ArticlesSource, topicName: string) => void;
};

const IMAGE_SIZE_PIXEL = 200;
// 100kb
const IMAGE_FILE_SIZE_BYTES = 100000;
const DEFAULT_IMG_SIZE = 0;
const DEFAULT_TOPIC = 'Others';
const GET_CATEGORIES_NAME_FAIL_MESSAGE = 'get categories names fail';
const GET_TOPICS_FAIL_MESSAGE = 'get topics fail';

const ArticlesSource: React.FC<Props> = (props: Props) => {
  const [image, setImage] = useState<string>();
  const [categories, setCategories] = useState<Categories>();
  const [topics, setTopics] = useState<Topics>();
  const [categoryNames, setCategoryNames] = useState<Array<string>>([]);
  const [topicsNames, setTopicsNames] = useState<Array<string>>([]);
  const [selectedCategory, setSelectedCategory] = useState<string>('');
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
      .max(IMAGE_FILE_SIZE_BYTES, 'image size must be less than 100kb')
      .min(DEFAULT_IMG_SIZE + 1, 'please submit logo'),
    title: yup.string().required('title must not be empty'),
    description: yup.string().required('description must not be empty'),
    feed_link: yup.string().required('feed link must not be empty').url(),
    link: yup.string().required('link to website must not be empty').url(),
    topicName: yup
      .string()
      .oneOf(topicsNames, 'topic invalid')
      .required('Please select topic')
      .trim('Please select topic'),
  });

  const requestListCategoriesNames = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('category/list/name');
      if (!data.success) {
        throw GET_CATEGORIES_NAME_FAIL_MESSAGE;
      }
      setCategories(data.categories);
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestListTopics = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('topic/list');
      if (!data.success) {
        throw GET_TOPICS_FAIL_MESSAGE;
      }
      setTopics(data.topics);
    } catch (error: any) {
      alertError(error);
    }
  };

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
        topicName: DEFAULT_TOPIC,
      };
    }, [props]),
  });

  const findTopicByName = (topicName: string): Topic => {
    if (topics) {
      const topic = topics.find((topic) => topic.name === topicName);
      if (topic) {
        return topic;
      }
    }
    const notFoundTopic: Topic = {
      id: 0,
      name: 'not found',
      category_id: 0,
    };
    return notFoundTopic;
  };

  const findCategoryByName = (categoryName: string): Category => {
    if (categories) {
      const category = categories.find(
        (category) => category.name === categoryName
      );
      if (category) {
        return category;
      }
    }
    const notFoundCategory: Category = {
      id: 0,
      illustration: '',
      name: 'not found',
    };
    return notFoundCategory;
  };

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
      topic_id: findTopicByName(data.topicName).id,
    };
    props.handleSubmit(articlesSource, data.topicName);
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
    requestListCategoriesNames();
    requestListTopics();
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

  useEffect(() => {
    const listName: Array<string> = [];
    if (categories) {
      categories.forEach((category) => listName.push(category.name));
      setSelectedCategory(listName[0]);
    }
    setCategoryNames(listName);
  }, [categories]);

  useEffect(() => {
    const listName: Array<string> = [];
    const categoryID = findCategoryByName(selectedCategory).id;
    if (topics) {
      const listTopicInCategory = topics.filter(
        (topic) => topic.category_id === categoryID
      );
      listTopicInCategory.forEach((topic) => listName.push(topic.name));
    }
    setTopicsNames(listName);
  }, [selectedCategory]);

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
            <label> Category</label>
            <div className="col-6">
              <InputGroup className="mb-3">
                <Form.Select
                  onChange={(e) => setSelectedCategory(e.target.value)}
                >
                  {categoryNames ? (
                    categoryNames.map((name) => (
                      <option
                        key={`category-create-article-source-option-${name}`}
                        value={name}
                      >
                        {name}
                      </option>
                    ))
                  ) : (
                    <option value="not found">not found</option>
                  )}
                </Form.Select>
              </InputGroup>
            </div>
            <label> Topic</label>
            <div className="col-6">
              <InputGroup className="mb-3">
                <Form.Select {...register('topicName')}>
                  {topicsNames ? (
                    topicsNames.map((name) => (
                      <option
                        key={`topic-create-article-source-option-${name}`}
                        value={name}
                      >
                        {name}
                      </option>
                    ))
                  ) : (
                    <option value="not found">not found</option>
                  )}
                </Form.Select>
              </InputGroup>
            </div>
            {errors.topicName && (
              <p className="errorMessage">{errors.topicName.message}</p>
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
            Create
          </Button>
        </form>
      </div>
    );
  }
  return <div className="info"></div>;
};

export default ArticlesSource;
