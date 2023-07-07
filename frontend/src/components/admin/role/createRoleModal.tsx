import React, { useEffect, useState } from 'react';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Button, Form, InputGroup } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faAudioDescription, faLock } from '@fortawesome/free-solid-svg-icons';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { alertError } from '@/helpers/alert';
import Popup from 'reactjs-popup';

type CreateRoleFormProperty = {
  name: string;
  description: string;
};

type Props = {
  handleCreate: (role: Role) => void;
};

const GET_PERMISSIONS_FAIL_MESSAGE = 'request get permission fail';

const CreateRoleModal: React.FC<Props> = (props: Props) => {
  const [permissions, setPermissions] = useState<Array<Permission>>([]);
  const [selectedPermissionsIDs, setSelectedPermissionsIDs] = useState<
    Array<number>
  >([]);
  const schema = yup.object().shape({
    name: yup
      .string()
      .required('Please enter role name')
      .trim('Please enter role name'),
    description: yup
      .string()
      .required('Please enter role description')
      .trim('Please enter role description'),
  });

  const handleSelected = (
    e: React.MouseEvent<HTMLButtonElement, MouseEvent>
  ) => {
    const id = +e.currentTarget.value;
    const index = selectedPermissionsIDs.findIndex(
      (permissionID) => permissionID === id
    );
    if (index > -1) {
      selectedPermissionsIDs.splice(index, 1);
    } else {
      selectedPermissionsIDs.push(id);
    }
    setSelectedPermissionsIDs([...selectedPermissionsIDs]);
  };

  const handleDetectChoosePermission = (id: number): boolean => {
    const index = selectedPermissionsIDs.findIndex(
      (permissionID) => permissionID === id
    );
    if (index >= 0) {
      return true;
    }
    return false;
  };

  const getPermissions = () => {
    const selected =  permissions.filter((permission) => handleDetectChoosePermission(permission.id))
    return selected
  }

  const requestPermission = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('/permission/list');
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw GET_PERMISSIONS_FAIL_MESSAGE;
      }
      setPermissions(data.permissions);
    } catch (error: any) {
      alertError(error);
    }
  };

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<CreateRoleFormProperty>({
    resolver: yupResolver(schema),
  });
  const onSubmit: SubmitHandler<CreateRoleFormProperty> = async (data) => {
    const role: Role = {
      id: 0,
      name: data.name,
      description: data.description,
      permissions: getPermissions()
    };
    props.handleCreate(role)
  };

  useEffect(() => {
    requestPermission();
  }, []);

  return (
    <div className="adminRoles__modal">
      <form onSubmit={handleSubmit(onSubmit)}>
        <h2 className="adminRoles__modal--title">Create roles</h2>
        <div className="adminRoles__modal--line" />
        <div className="field">
          <label> Name </label>
          <InputGroup className="mb-3">
            <InputGroup.Text>
              <FontAwesomeIcon icon={faLock} fixedWidth />
            </InputGroup.Text>
            <Form.Control
              {...register('name')}
              placeholder="Type role name"
              type="text"
            />
          </InputGroup>

          {errors.name && <p className="errorMessage">{errors.name.message}</p>}
        </div>

        <div className="field">
          <label> Description </label>
          <InputGroup className="mb-3">
            <InputGroup.Text>
              <FontAwesomeIcon icon={faAudioDescription} fixedWidth />
            </InputGroup.Text>
            <Form.Control
              {...register('description')}
              placeholder="Type role description"
              type="text"
            />
          </InputGroup>

          {errors.description && (
            <p className="errorMessage">{errors.description.message}</p>
          )}
        </div>

        <div className="field">
          <label> Permission </label>
          <div className="row">
            {permissions.map((permission) => (
              <Popup  
              trigger={() => (
                  <div className="col-4">
                    <Button
                      variant={
                        handleDetectChoosePermission(permission.id)
                          ? 'primary'
                          : 'secondary'
                      }
                      value={permission.id}
                      onClick={(e) => handleSelected(e)}
                    >
                      {permission.method} {permission.entity}
                    </Button>
                  </div>
                )}
                position="bottom center"
                closeOnDocumentClick
                on={['hover', 'focus']}
              >
                <div>
                  <p>{permission.description}</p>
                </div>
              </Popup>
            ))}
          </div>
        </div>

        <Button className="w-100 px-4" variant="success" type="submit">
          Create
        </Button>
      </form>
    </div>
  );
};

export default CreateRoleModal;
