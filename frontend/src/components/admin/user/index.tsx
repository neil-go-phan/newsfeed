import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import React, { useEffect, useState } from 'react';
import { Button } from 'react-bootstrap';
import { alertError, alertSuccess } from '@/helpers/alert';
import { ThreeDots } from 'react-loader-spinner';
import { toastifyError } from '@/helpers/toastify';
import Popup from 'reactjs-popup';
import AdminUsersPagination from './pagination';
import UsersTable from './table';

export const PAGE_SIZE = 10;
const FIRST_PAGE = 1;

const DELETE_USER_SUCCESS_MESSAGE = 'delete user success';
const DELETE_USER_FAIL_MESSAGE = 'delete user fail';
const UPDATE_USER_ROLE_SUCCESS_MESSAGE = 'update user success';
const UPDATE_USER_ROLE_FAIL_MESSAGE = 'update user fail';
const GET_PAGE_USERS_FAIL_MESSAGE = 'get page user fail';
const COUNT_FAIL_MESSAGE = 'request count user fail';

export default function AdminUsers() {
  const [users, setUsers] = useState<Array<User>>();
  const [currentPage, setCurrentPage] = useState<number>(1);
  const [total, setTotal] = useState<number>(0);

  const handleDelete = (id: number) => {
    requestDelete(id)
  };
  const handleUpdate = (id: number, newRole: string) => {
    requestUpdate(id, newRole)
  };

  const pageChangeHandler = (currentPage: number) => {
    setCurrentPage(currentPage);
    requestPageUsers(currentPage, PAGE_SIZE);
  };

  const requestDelete = async (id: number) => {
    try {
      const { data } = await axiosProtectedAPI.post('auth/delete', {
        id: id,
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw DELETE_USER_FAIL_MESSAGE;
      }
      alertSuccess(DELETE_USER_SUCCESS_MESSAGE);
      requestPageUsers(currentPage, PAGE_SIZE);
      requestCountUsers();
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestUpdate = async (id: number, newRole: string) => {
    try {
      const { data } = await axiosProtectedAPI.post('/auth/update/role', {
        id: id,
        new_role:newRole
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw UPDATE_USER_ROLE_FAIL_MESSAGE;
      }
      alertSuccess(UPDATE_USER_ROLE_SUCCESS_MESSAGE);
      requestPageUsers(currentPage, PAGE_SIZE);
      requestCountUsers();
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestPageUsers= async (page: number, pageSize: number) => {
    try {
      const { data } = await axiosProtectedAPI.get('/auth/list', {
        params: { page: page, page_size: pageSize },
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw GET_PAGE_USERS_FAIL_MESSAGE;
      }
      setUsers(data.users);
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestCountUsers = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('/auth/count');
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
    requestPageUsers(FIRST_PAGE, PAGE_SIZE);
    requestCountUsers();
  }, []);

  return (
    <div className="adminUsers">
      <h1 className="adminUsers__title">Manage users</h1>
      <div className="adminUsers__overview">
        <div className="adminUsers__overview--item">
          <p>
            Total users: <span>{total}</span>
          </p>
        </div>
      </div>
      <div className="adminUsers__list">
        <h2 className="adminUsers__list--title">Users list</h2>

        {users ? (
          <UsersTable
            users={users}
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
        <AdminUsersPagination
          totalRows={total!}
          pageChangeHandler={pageChangeHandler}
          currentPage={currentPage}
        />
      </div>
    </div>
  );
}
