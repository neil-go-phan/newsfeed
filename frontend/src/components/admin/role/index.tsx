import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import React, { useEffect, useState } from 'react';
import { Button } from 'react-bootstrap';
import { alertError, alertSuccess } from '@/helpers/alert';
import { ThreeDots } from 'react-loader-spinner';
import { toastifyError } from '@/helpers/toastify';
import Popup from 'reactjs-popup';
import AdminRolePagination from './pagination';
import RolesTable from './table';
import CreateRoleModal from './createRoleModal';

export const PAGE_SIZE = 10;
const FIRST_PAGE = 1;
const CREATE_ROLE_SUCCESS_MESSAGE = 'create role success';
const CREATE_ROLE_FAIL_MESSAGE = 'create role fail';
const DELETE_ROLE_SUCCESS_MESSAGE = 'delete role success';
const DELETE_ROLE_FAIL_MESSAGE = 'delete role fail';
const UPDATE_ROLE_SUCCESS_MESSAGE = 'update role success';
const UPDATE_ROLE_FAIL_MESSAGE = 'update role fail';
const GET_PAGE_ROLES_FAIL_MESSAGE = 'get page roles fail';
const COUNT_FAIL_MESSAGE = 'request count roles fail';

export default function AdminRoles() {
  const [roles, setRoles] = useState<Array<Role>>();
  const [currentPage, setCurrentPage] = useState<number>(1);
  const [total, setTotal] = useState<number>(0);
  const [isCreateRoleModalOpen, setIsCreateRoleModalOpen] =
    useState<boolean>(false);

  const handleCreate = (role: Role) => {
    requestCreateRole(role);
    handleIsCreateRoleModalClose();
  };

  const handleDelete = (id: number) => {
    requestDelete(id)
  };
  const handleUpdate = (role: Role) => {
    requestUpdate(role)
  };

  const handleIsCreateRoleModalClose = () => {
    setIsCreateRoleModalOpen(false);
  };

  const pageChangeHandler = (currentPage: number) => {
    setCurrentPage(currentPage);
    requestPageRoles(currentPage, PAGE_SIZE);
  };

  const requestDelete = async (id: number) => {
    try {
      const { data } = await axiosProtectedAPI.post('role/delete', {
        id: id,
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw DELETE_ROLE_FAIL_MESSAGE;
      }
      alertSuccess(DELETE_ROLE_SUCCESS_MESSAGE);
      requestPageRoles(currentPage, PAGE_SIZE);
      requestCountRoles();
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestUpdate = async (role: Role) => {
    try {
      const { data } = await axiosProtectedAPI.post('role/update', {
        id: role.id,
        name: role.name,
        description: role.description,
        permissions: role.permissions,
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw UPDATE_ROLE_FAIL_MESSAGE;
      }
      alertSuccess(UPDATE_ROLE_SUCCESS_MESSAGE);
      requestPageRoles(currentPage, PAGE_SIZE);
      requestCountRoles();
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestCreateRole = async (role: Role) => {
    try {
      const { data } = await axiosProtectedAPI.post('role/create', {
        id: role.id,
        name: role.name,
        description: role.description,
        permissions: role.permissions,
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw CREATE_ROLE_FAIL_MESSAGE;
      }
      alertSuccess(CREATE_ROLE_SUCCESS_MESSAGE);
      requestPageRoles(currentPage, PAGE_SIZE);
      requestCountRoles();
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestPageRoles = async (page: number, pageSize: number) => {
    try {
      const { data } = await axiosProtectedAPI.get('/role/list', {
        params: { page: page, page_size: pageSize },
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw GET_PAGE_ROLES_FAIL_MESSAGE;
      }
      setRoles(data.roles);
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestCountRoles = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('/role/count/all');
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw COUNT_FAIL_MESSAGE;
      }
      setTotal(data.total);
    } catch (error: any) {
      toastifyError(COUNT_FAIL_MESSAGE);
    }
  };

  useEffect(() => {
    requestPageRoles(FIRST_PAGE, PAGE_SIZE);
    requestCountRoles();
  }, []);

  return (
    <div className="adminRoles">
      <h1 className="adminRoles__title">Manage roles</h1>
      <div className="adminRoles__overview">
        <div className="adminRoles__overview--item">
          <p>
            Total roles: <span>{total}</span>
          </p>
        </div>
      </div>
      <div className="adminRoles__list">
        <h2 className="adminRoles__list--title">Roles list</h2>
        <div className="adminRoles__list--search d-sm-flex">
          <div className="col-sm-1"></div>
          <div className="addBtn col-sm-5">
            <Button
              variant="primary mb-2"
              onClick={() => setIsCreateRoleModalOpen(true)}
            >
              Create new role
            </Button>
            <Popup
              modal
              open={isCreateRoleModalOpen}
              onClose={handleIsCreateRoleModalClose}
            >
              <CreateRoleModal handleCreate={handleCreate} />
            </Popup>
          </div>
        </div>
        {roles ? (
          <RolesTable
            roles={roles}
            currentPage={currentPage!}
            handleDelete={handleDelete}
            handleUpdate={handleUpdate}
          />
        ) : (
          <div className="threeDotLoading">
            <ThreeDots
              height="50"
              width="50"
              radius="9"
              color="#4fa94d"
              ariaLabel="three-dots-loading"
              visible={true}
            />
          </div>
        )}
        <AdminRolePagination
          totalRows={total!}
          pageChangeHandler={pageChangeHandler}
          currentPage={currentPage}
        />
      </div>
    </div>
  );
}
