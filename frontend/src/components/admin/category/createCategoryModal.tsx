import React, { useMemo, useState } from 'react';
import * as yup from 'yup';
import Image from 'next/image';
import { yupResolver } from '@hookform/resolvers/yup';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Button, Form, InputGroup } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faShapes } from '@fortawesome/free-solid-svg-icons';

type CreateCategoryFormProperty = {
  name: string;
  imgSize: number;
};

type Props = {
  handleCreateCategory: (
    categoryName: string,
    categoryIllustration: string
  ) => void;
};

const IMAGE_SIZE_PIXEL = 200;
// 1mb
const IMAGE_FILE_SIZE_BYTES = 1000000;
const DEFAULT_IMG_SIZE = 0;

const CreateCategoryModal: React.FC<Props> = (props: Props) => {
  const [image, setImage] = useState<string>();
  const schema = yup.object().shape({
    name: yup
      .string()
      .required('Please enter category name')
      .trim('Please enter category name'),
    imgSize: yup
      .number()
      .required('illustration image must not be empty')
      .max(IMAGE_FILE_SIZE_BYTES, 'image size must be less than 1mb')
      .min(DEFAULT_IMG_SIZE + 1, 'please submit category illustration'),
  });
  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors },
  } = useForm<CreateCategoryFormProperty>({
    resolver: yupResolver(schema),
    defaultValues: useMemo(() => {
      return {
        imgSize: DEFAULT_IMG_SIZE,
      };
    }, [props]),
  });
  const onSubmit: SubmitHandler<CreateCategoryFormProperty> = async (data) => {
    let base64Img: string = '';
    if (image) {
      base64Img = image;
    }
    props.handleCreateCategory(data.name, base64Img);
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

  return (
    <div className="adminCategories__modal">
      <form onSubmit={handleSubmit(onSubmit)}>
        <h2 className="adminCategories__modal--title">Create category</h2>
        <div className="adminCategories__modal--line" />
        <label> Category </label>
        <div className="field">
          <InputGroup className="mb-3">
            <InputGroup.Text>
              <FontAwesomeIcon icon={faShapes} fixedWidth />
            </InputGroup.Text>
            <Form.Control
              {...register('name')}
              placeholder="Type category name"
              type="text"
              required
            />
          </InputGroup>

          {errors.name && <p className="errorMessage">{errors.name.message}</p>}
        </div>
        <div className="field">
          <label> Image </label>
          <div className="illustration">
            <div className="">
              {image ? (
                <Image
                  alt="category illustration image"
                  src={image}
                  width={IMAGE_SIZE_PIXEL}
                  height="0"
                  style={{ height: 'auto' }}
                />
              ) : (
                <div>Not found category illustration image, please add one</div>
              )}
            </div>
            <div className="mt-3">
              <input
                type="file"
                name="category-illustration"
                id="file"
                accept=".jpef, .png, .jpg, .jpeg"
                onChange={(event) => photoUpload(event)}
                src={image}
              />
            </div>
          </div>
          {errors.imgSize && (
            <p className="errorMessage">{errors.imgSize.message}</p>
          )}
        </div>

        <Button className="w-100 px-4" variant="success" type="submit">
          Create
        </Button>
      </form>
    </div>
  );
};

export default CreateCategoryModal;
