import React, { useEffect, useMemo, useState } from 'react';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Button, Form, InputGroup } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faFolder, faShapes } from '@fortawesome/free-solid-svg-icons';

type UpdateTopicsFormProperty = {
  name: string;
  categoryName: string;
};

type Props = {
  oldName: string;
  oldCategoryName: string;
  categories: Categories;
  handleUpdateTopic: (newName: string, newCategoryID:number)  => void;
};

const UpdateTopicModal: React.FC<Props> = (props: Props) => {
  const [categoryNames, setCategoryNames] = useState<Array<string>>([]);
  const schema = yup.object().shape({
    name: yup
      .string()
      .required('Please enter topic name')
      .trim('Please enter topic name'),
    categoryName: yup
      .string()
      .oneOf(categoryNames, 'category invalid')
      .required('Please enter category')
      .trim('Please enter category'),
  });

  useEffect(() => {
    const listName: Array<string> = [];
    if (props.categories) {
      props.categories.forEach((category) => listName.push(category.name));
    }
    setCategoryNames(listName);
  }, [props.categories]);

  const findCategoryByName = (categoryName: string): Category => {
    if (props.categories) {
      const category = props.categories.find(
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

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<UpdateTopicsFormProperty>({
    resolver: yupResolver(schema),
    defaultValues: useMemo(() => {
      return {
        name: props.oldName,
        categoryName: props.oldCategoryName,
      };
    }, [props]),
  });
  const onSubmit: SubmitHandler<UpdateTopicsFormProperty> = async (data) => {
    const { name, categoryName } = data;
    const category = findCategoryByName(categoryName);
    props.handleUpdateTopic(name, category.id);
  };

  return (
    <div className="adminTopics__modal">
      <form onSubmit={handleSubmit(onSubmit)}>
        <h2 className="adminTopics__modal--title">Update topic</h2>
        <div className="adminTopics__modal--line" />
        <div className="field">
          <label> Topic name </label>
          <InputGroup className="mb-3">
            <InputGroup.Text>
              <FontAwesomeIcon icon={faFolder} fixedWidth />
            </InputGroup.Text>
            <Form.Control
              {...register('name')}
              placeholder="Type topic name"
              type="text"
              required
            />
          </InputGroup>

          {errors.name && <p className="errorMessage">{errors.name.message}</p>}
        </div>

        <div className="field">
          <label> Category </label>
          <InputGroup className="mb-3">
            <InputGroup.Text>
              <FontAwesomeIcon icon={faShapes} fixedWidth />
            </InputGroup.Text>
            <Form.Select {...register('categoryName')}>
              {props.categories ? (
                props.categories.map((category) => (
                  <option
                    key={`category-create-topic-modal-option-${category.name}`}
                    value={category.name}
                  >
                    {category.name}
                  </option>
                ))
              ) : (
                <option value="not found">not found</option>
              )}
            </Form.Select>
          </InputGroup>

          {errors.categoryName && (
            <p className="errorMessage">{errors.categoryName.message}</p>
          )}
        </div>

        <Button className="w-100 px-4" variant="success" type="submit">
          Update
        </Button>
      </form>
    </div>
  );
};

export default UpdateTopicModal;
