import React, { useEffect, useMemo, useState } from 'react';
import * as yup from 'yup';
import Image from 'next/image';
import { yupResolver } from '@hookform/resolvers/yup';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Button, Form, InputGroup } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faShapes } from '@fortawesome/free-solid-svg-icons';

type UpdateCategoryFormProperty = {
  name: string;
  imgSize: number;
};

type Props = {
  oldName: string;
  oldIllustration:string;
  handleUpdateCategory: (categoryName: string, categoryIllustration: string) => void;
};

const IMAGE_SIZE_PIXEL = 200;
// 1mb
const IMAGE_FILE_SIZE_BYTES = 1000000;
const DEFAULT_IMG_SIZE = 0;

const UpdateCategoryModal: React.FC<Props> = (props: Props) => {
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
  } = useForm<UpdateCategoryFormProperty>({
    resolver: yupResolver(schema),
    defaultValues: useMemo(() => {
      return {
        name: props.oldName,
        imgSize: DEFAULT_IMG_SIZE + 1,
      };
    }, [props]),
  });
  const onSubmit: SubmitHandler<UpdateCategoryFormProperty> = async (data) => {
    let base64Img: string = '';
    if (image) {
      base64Img = image;
    }
    props.handleUpdateCategory(data.name, base64Img);
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
    setImage(props.oldIllustration)
  }, [])
  

  return (
    <div className="adminCategories__modal ">
      <form onSubmit={handleSubmit(onSubmit)}>
        <h2 className="adminCategories__modal--title">Update category</h2>
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
            <div className="illustration d-flex">
              <div className="col-6">
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
              <div className="col-6">
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
          Update
        </Button>
      </form>
    </div>
  );
};

export default UpdateCategoryModal;
