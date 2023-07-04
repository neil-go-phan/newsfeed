import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import React, { useEffect, useState } from 'react';
import { Button } from 'react-bootstrap';
import { useRouter } from 'next/router';
import { alertError, alertSuccess } from '@/helpers/alert';
import { ThreeDots } from 'react-loader-spinner';
import AdmintopicPagination from './pagination';
import { toastifyError } from '@/helpers/toastify';
import Popup from 'reactjs-popup';
import TopicsTable from './table';
import CreateTopicModal from './createTopicModal';

export const PAGE_SIZE = 10;
const CREATE_TOPIC_SUCCESS_MESSAGE = 'create topic success';
const CREATE_TOPIC_FAIL_MESSAGE = 'create topic fail';
const DELETE_TOPIC_SUCCESS_MESSAGE = 'delete topic success';
const DELETE_TOPIC_FAIL_MESSAGE = 'delete topic fail';
const UPDATE_TOPIC_SUCCESS_MESSAGE = 'update topic success';
const UPDATE_TOPIC_FAIL_MESSAGE = 'update topic fail';
const GET_PAGE_TOPICS_FAIL_MESSAGE = 'get page topic fail';
const GET_CATEGORIES_NAME_FAIL_MESSAGE = 'get categories names fail';

export default function AdminTopics() {
  const [topics, setTopics] = useState<Topics>();
  const [categories, setCategories] = useState<Categories>();
  const [currentPage, setCurrentPage] = useState<number>(1);
  const [total, setTotal] = useState<number>(0);
  const [isCreateTopicModalOpen, setIsCreateTopicModalOpen] =
    useState<boolean>(false);

  const handleCreateTopic = (topicName: string, category_id: number) => {
    requestCreateTopic(topicName.trim(), category_id);
    handleIsCreateTopicModalClose();
  };

  const handleDeleteTopic = (id: number, name: string, category_id: number) => {
    requestDeleteTopic(id, name, category_id);
  };

  const handleIsCreateTopicModalClose = () => {
    setIsCreateTopicModalOpen(false);
  };

  const handleUpdateTopic = (
    id: number,
    newName: string,
    newCategoryID: number
  ) => {
    requestUpdateTopic(id, newName, newCategoryID);
  };
  // TODO: refactor paging logic
  const pageChangeHandler = (currentPage: number) => {
    setCurrentPage(currentPage);
    requestPageTopics(currentPage, PAGE_SIZE);
  };

  const requestDeleteTopic = async (
    id: number,
    name: string,
    category_id: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.post('topic/delete', {
        name: name,
        id: id,
        category_id: category_id,
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw DELETE_TOPIC_FAIL_MESSAGE;
      }
      alertSuccess(DELETE_TOPIC_SUCCESS_MESSAGE);
      requestPageTopics(currentPage, PAGE_SIZE);
      requestCountTopics();
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestUpdateTopic = async (
    id: number,
    newName: string,
    newCategoryID: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.post('topic/update', {
        id: id,
        name: newName,
        category_id: newCategoryID,
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw UPDATE_TOPIC_FAIL_MESSAGE;
      }
      alertSuccess(UPDATE_TOPIC_SUCCESS_MESSAGE);
      requestPageTopics(currentPage, PAGE_SIZE);
      requestCountTopics();
      pageChangeHandler(1);
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestListCategoriesNames = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('category/list/name');
      if (!data.success) {
        throw 'get categories names fail';
      }
      setCategories(data.categories);
    } catch (error: any) {
      alertError(GET_CATEGORIES_NAME_FAIL_MESSAGE);
    }
  };

  const requestCreateTopic = async (name: string, categoryID: number) => {
    try {
      const { data } = await axiosProtectedAPI.post('topic/create', {
        name: name,
        category_id: categoryID,
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw CREATE_TOPIC_FAIL_MESSAGE;
      }
      alertSuccess(CREATE_TOPIC_SUCCESS_MESSAGE);
      requestPageTopics(currentPage, PAGE_SIZE);
      requestCountTopics();
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestPageTopics = async (page: number, pageSize: number) => {
    try {
      const { data } = await axiosProtectedAPI.get('topic/get/page', {
        params: { page: page, page_size: pageSize },
      });
      if (!data.success) {
        throw 'get categories fail';
      }
      setTopics(data.topics);
    } catch (error: any) {
      alertError(GET_PAGE_TOPICS_FAIL_MESSAGE);
    }
  };

  const requestCountTopics = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('topic/count');
      if (!data.success) {
        throw 'count categories fail';
      }
      setTotal(data.total);
    } catch (error: any) {
      toastifyError(GET_PAGE_TOPICS_FAIL_MESSAGE);
    }
  };

  useEffect(() => {
    requestListCategoriesNames();
    requestPageTopics(1, PAGE_SIZE);
    requestCountTopics();
  }, []);

  return (
    <div className="adminTopics">
      <h1 className="adminTopics__title">Manage topics</h1>
      <div className="adminTopics__overview">
        <div className="adminTopics__overview--item">
          <p>
            Total topics: <span>{total}</span>
          </p>
        </div>
      </div>
      <div className="adminTopics__list">
        <h2 className="adminTopics__list--title">Topics list</h2>
        <div className="adminTopics__list--search d-sm-flex">
          <div className="col-sm-1"></div>
          <div className="addBtn col-sm-5">
            <Button
              variant="primary mb-2"
              onClick={() => setIsCreateTopicModalOpen(true)}
            >
              Create new topic
            </Button>
            <Popup
              modal
              open={isCreateTopicModalOpen}
              onClose={handleIsCreateTopicModalClose}
            >
              <CreateTopicModal
                categories={categories}
                handleCreateTopic={handleCreateTopic}
              />
            </Popup>
          </div>
        </div>
        {topics && categories ? (
          <TopicsTable
            categories={categories}
            currentPage={currentPage!}
            topics={topics}
            handleDeleteTopic={handleDeleteTopic}
            handleUpdateTopic={handleUpdateTopic}
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
        <AdmintopicPagination
          totalRows={total!}
          pageChangeHandler={pageChangeHandler}
          currentPage={currentPage}
        />
      </div>
    </div>
  );
}
