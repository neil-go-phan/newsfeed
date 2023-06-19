import React from 'react';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Button, Form, InputGroup } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faShapes } from '@fortawesome/free-solid-svg-icons';

type UpdateCategoryFormProperty = {
  name: string;
};

type Props = {
  name: string;
  handleUpdateCategory: (newName: string) => void;
};

const UpdateCategoryModal: React.FC<Props> = (props: Props) => {
  const schema = yup.object().shape({
    name: yup.string()
      .required('Please enter category name')
      .trim('Please enter category name'),
  });
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<UpdateCategoryFormProperty>({
    resolver: yupResolver(schema),
  });
  const onSubmit: SubmitHandler<UpdateCategoryFormProperty> = async (data) => {
    let { name } = data;
    props.handleUpdateCategory(name);
  };
  return (
    <div className="adminCrawler__addModal">
      <form onSubmit={handleSubmit(onSubmit)}>
        <h2 className="adminCrawler__addModal--title">Input url</h2>
        <div className="adminCrawler__addModal--line" />
        <label> Category </label>
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

        {errors.name && (
          <p className="errorMessage">{errors.name.message}</p>
        )}
        <Button className="w-100 px-4" variant="success" type="submit">
          Update
        </Button>
      </form>
    </div>
  );
}

export default UpdateCategoryModal