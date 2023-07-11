import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import React, { useEffect, useState } from 'react';
import { Button } from 'react-bootstrap';
import { alertError, alertSuccess } from '@/helpers/alert';
import CategoriseTable from './table';
import { ThreeDots } from 'react-loader-spinner';
import AdminCategoryPagination from './pagination';
import { toastifyError } from '@/helpers/toastify';
import Popup from 'reactjs-popup';
import CreateCategoryModal from './createCategoryModal';

export const PAGE_SIZE = 10;
const CREATE_CATEGORY_SUCCESS_MESSAGE = 'create category success';
const CREATE_CATEGORY_FAIL_MESSAGE = 'create category fail';
const DELETE_CATEGORY_SUCCESS_MESSAGE = 'delete category success';
const DELETE_CATEGORY_FAIL_MESSAGE = 'delete category fail';
const UPDATE_CATEGORY_SUCCESS_MESSAGE = 'update category success';
const UPDATE_CATEGORY_FAIL_MESSAGE = 'update category fail';
const GET_PAGE_CATEGORIES_FAIL_MESSAGE = 'get page category fail';

export default function AdminCategories() {
  const [categories, setCategories] = useState<Categories>();
  const [currentPage, setCurrentPage] = useState<number>(1);
  const [total, setTotal] = useState<number>(0);
  const [isCreateCategoryModalOpen, setIsCreateCategoryModalOpen] =
    useState<boolean>(false);

  const handleCreateCategory = (
    categoryName: string,
    categoryIllustration: string
  ) => {
    requestCreateCategory(categoryName.trim(), categoryIllustration);
    handleIsCreateCategoryModalClose();
  };

  const handleDeleteCategory = (id: number, name: string) => {
    requestDeleteCategory(id, name);
  };

  const handleIsCreateCategoryModalClose = () => {
    setIsCreateCategoryModalOpen(false);
  };

  const handleUpdateCategory = (
    id: number,
    oldName: string,
    newName: string,
    newIllustration: string
  ) => {
    requestUpdateCategory(id, oldName, newName, newIllustration);
  };
  // TODO: refactor paging logic
  const pageChangeHandler = (currentPage: number) => {
    setCurrentPage(currentPage);
    requestPageCategories(currentPage, PAGE_SIZE);
  };

  const requestDeleteCategory = async (id: number, name: string) => {
    try {
      const { data } = await axiosProtectedAPI.post('category/delete', {
        name: name,
        id: id,
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw DELETE_CATEGORY_FAIL_MESSAGE;
      }
      alertSuccess(DELETE_CATEGORY_SUCCESS_MESSAGE);
      requestPageCategories(currentPage, PAGE_SIZE);
      requestCountCategories();
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestUpdateCategory = async (
    id: number,
    name: string,
    newName: string,
    newIllustration: string
  ) => {
    try {
      const { data } = await axiosProtectedAPI.post('category/update', {
        new_name: newName,
        new_illustration: newIllustration,
        category: {
          name: name,
          id: id,
        },
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw UPDATE_CATEGORY_FAIL_MESSAGE;
      }
      alertSuccess(UPDATE_CATEGORY_SUCCESS_MESSAGE);
      requestPageCategories(currentPage, PAGE_SIZE);
      requestCountCategories();
      pageChangeHandler(1);
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestCreateCategory = async (name: string, illustration: string) => {
    try {
      const { data } = await axiosProtectedAPI.post('category/create', {
        name: name,
        illustration: illustration,
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw CREATE_CATEGORY_FAIL_MESSAGE;
      }
      alertSuccess(CREATE_CATEGORY_SUCCESS_MESSAGE);
      requestPageCategories(currentPage, PAGE_SIZE);
      requestCountCategories();
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestPageCategories = async (page: number, pageSize: number) => {
    try {
      const { data } = await axiosProtectedAPI.get('category/get/page', {
        params: { page: page, page_size: pageSize },
      });
      if (!data.success) {
        throw 'get categories fail';
      }
      setCategories(data.categories);
    } catch (error: any) {
      alertError(GET_PAGE_CATEGORIES_FAIL_MESSAGE);
    }
  };

  const requestCountCategories = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('category/count');
      if (!data.success) {
        throw 'count categories fail';
      }
      setTotal(data.total);
    } catch (error: any) {
      toastifyError(GET_PAGE_CATEGORIES_FAIL_MESSAGE);
    }
  };

  useEffect(() => {
    requestPageCategories(1, PAGE_SIZE);
    requestCountCategories();
  }, []);

  return (
    <div className="adminCategories">
      <h1 className="adminCategories__title">Manage categories</h1>
      <div className="adminCategories__overview">
        <div className="adminCategories__overview--item">
          <p>
            Total categories: <span>{total}</span>
          </p>
        </div>
      </div>
      <div className="adminCategories__list">
        <h2 className="adminCategories__list--title">Categories list</h2>
        <div className="adminCategories__list--search d-sm-flex">
          <div className="col-sm-1"></div>
          <div className="addBtn col-sm-5">
            <Button
              variant="primary mb-2"
              onClick={() => setIsCreateCategoryModalOpen(true)}
            >
              Create new category
            </Button>
            <Popup
              modal
              open={isCreateCategoryModalOpen}
              onClose={handleIsCreateCategoryModalClose}
            >
              <CreateCategoryModal
                handleCreateCategory={handleCreateCategory}
              />
            </Popup>
          </div>
        </div>
        {categories ? (
          <CategoriseTable
            categories={categories}
            currentPage={currentPage!}
            handleDeleteCategory={handleDeleteCategory}
            handleUpdateCategory={handleUpdateCategory}
          />
        ) : (
          <div className="adminArticles__table--loading">
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
        <AdminCategoryPagination
          totalRows={total!}
          pageChangeHandler={pageChangeHandler}
          currentPage={currentPage}
        />
      </div>
    </div>
  );
}
