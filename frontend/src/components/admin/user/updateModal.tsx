import React, { useEffect, useMemo, useState } from 'react';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Button, Form, InputGroup } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faLock } from '@fortawesome/free-solid-svg-icons';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { toastifyError } from '@/helpers/toastify';

type UpdateRoleFormProperty = {
  roleName: string;
};

type Props = {
  user: User;
  handleUpdate: (newRole: string) => void;
};

const UpdateUserRoleModal: React.FC<Props> = (props: Props) => {
  const [roleNames, setRoleNames] = useState<Array<string>>([]);
  const schema = yup.object().shape({
      roleName: yup
      .string()
      .oneOf(roleNames, 'role invalid')
      .required('Please enter role')
      .trim('Please enter role'),
  });

  useEffect(() => {
    requestRolesNames()
  }, []);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<UpdateRoleFormProperty>({
    resolver: yupResolver(schema),
    defaultValues: useMemo(() => {
      return {
        roleName: props.user.role_name
      };
    }, [props]),
  });
  const onSubmit: SubmitHandler<UpdateRoleFormProperty> = async (data) => {
    props.handleUpdate(data.roleName)
  };

  const requestRolesNames = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('/role/list/names');
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw 'fail';
      }
      setRoleNames(data.names);
    } catch (error: any) {
      toastifyError(error);
    }
  };

  return (
    <div className="adminUsers__modal">
      <form onSubmit={handleSubmit(onSubmit)}>
        <h2 className="adminUsers__modal--title">Change user roles</h2>
        <div className="adminUsers__modal--line" />
        <div className="field">
          <label> Roles </label>
          <InputGroup className="mb-3">
            <InputGroup.Text>
              <FontAwesomeIcon icon={faLock} fixedWidth />
            </InputGroup.Text>
            <Form.Select {...register('roleName')}>
              {roleNames ? (
                roleNames.map((roleName) => (
                  <option
                    key={`chage role-modal-option-${roleName}`}
                    value={roleName}
                  >
                    {roleName}
                  </option>
                ))
              ) : (
                <option value="not found">not found</option>
              )}
            </Form.Select>
          </InputGroup>

          {errors.roleName && (
            <p className="errorMessage">{errors.roleName.message}</p>
          )}
        </div>

        <Button className="w-100 px-4" variant="success" type="submit">
          Update
        </Button>
      </form>
    </div>
  );
};

export default UpdateUserRoleModal;
