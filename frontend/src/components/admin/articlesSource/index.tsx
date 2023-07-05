import { alertError, alertSuccess } from '@/helpers/alert';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { faFilter, faSearch } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React, { useEffect, useState } from 'react';
import { Button, Form, InputGroup } from 'react-bootstrap';
import { ThreeDots } from 'react-loader-spinner';
import ArticlesSourcesTable from './table';
import AdminArticlesSourcesPagination from './pagination';
import { toastifyError } from '@/helpers/toastify';
import FilterByTopic from './filterByTopic';
import Link from 'next/link';
import { _ROUTES } from '@/helpers/constants';

export const PAGE_SIZE = 10;
const FIRST_PAGE = 1;
const GET_PAGE_ARTICLES_SOURCES_FAIL_MESSAGE = 'request list fail';
const DELETE_ARTICLES_SOURCE_FAIL_MESSAGE = 'delete article source fail';
const DELETE_ARTICLES_SOURCE_SUCCESS_MESSAGE = 'delete source success';

function AdminArticlesSource() {
  const [articlesSources, setArticlesSources] = useState<ArticlesSourceInfoes>(
    []
  );
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [currentPage, setCurrentPage] = useState<number>(FIRST_PAGE);
  const [keyword, setKeyword] = useState<string>('');
  const [topicID, setTopicID] = useState<number>(0);
  const [total, setTotal] = useState<number>(0);
  const [found, setFound] = useState<number>(0);

  const pageChangeHandler = (newCurrentPage: number) => {
    setCurrentPage(newCurrentPage);
    setIsLoading(true);
    if (keyword === '' && topicID === 0) {
      requestListAll(newCurrentPage, PAGE_SIZE);
      setFound(total);
      return;
    }
    requestSearchNextPage(keyword, newCurrentPage, PAGE_SIZE, topicID);
    return;
  };

  const submitSearch = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    handleSearchArticleSources();
  };

  const handleChooseTopicID = (id: number) => {
    setTopicID(id);
    setKeyword('');
  };

  const handleSearchArticleSources = () => {
    setCurrentPage(FIRST_PAGE);
    setIsLoading(true);
    if (keyword === '' && topicID === 0) {
      requestListAll(FIRST_PAGE, PAGE_SIZE);
      setFound(total);
      return;
    }
    requestFirstSearch(keyword, FIRST_PAGE, PAGE_SIZE, topicID);
  };

  const handleDeleteArticlesSource = (id: number) => {
    requestDeleteArticlesSource(id);
  };

  const handleUpdate = (articlesSource: UpdateArticleSourcePayload) => {
    requestUpdate(articlesSource)
  }

  const requestUpdate = async (articlesSource: UpdateArticleSourcePayload) => {
    try {
      const { data } = await axiosProtectedAPI.post('/articles-sources/update/id', {
        id: articlesSource.id,
        title: articlesSource.title,
        description: articlesSource.description,
        image: articlesSource.image,
        topic_id: articlesSource.topic_id
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw DELETE_ARTICLES_SOURCE_FAIL_MESSAGE;
      }
      alertSuccess(DELETE_ARTICLES_SOURCE_SUCCESS_MESSAGE);
      requestListAll(currentPage, PAGE_SIZE);
      requestCountArticlesSources();
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestDeleteArticlesSource = async (id: number) => {
    try {
      const { data } = await axiosProtectedAPI.post('/articles-sources/delete/id', {
        id: id,
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw DELETE_ARTICLES_SOURCE_FAIL_MESSAGE;
      }
      alertSuccess(DELETE_ARTICLES_SOURCE_SUCCESS_MESSAGE);
      requestListAll(currentPage, PAGE_SIZE);
      requestCountArticlesSources();
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestFirstSearch = async (
    keyword: string,
    page: number,
    pageSize: number,
    topicID: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles-sources/search/filter',
        {
          params: {
            q: keyword,
            page: page,
            page_size: pageSize,
            topic_id: topicID,
          },
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw GET_PAGE_ARTICLES_SOURCES_FAIL_MESSAGE;
      }
      setArticlesSources(data.articles_sources);
      setFound(data.found);
      setIsLoading(false);
    } catch (error: any) {
      alertError(error);
      setIsLoading(false);
    }
  };

  const requestSearchNextPage = async (
    keyword: string,
    page: number,
    pageSize: number,
    topicID: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles-sources/search/filter',
        {
          params: {
            q: keyword,
            page: page,
            page_size: pageSize,
            topic_id: topicID,
          },
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw GET_PAGE_ARTICLES_SOURCES_FAIL_MESSAGE;
      }
      setArticlesSources(data.articles_sources);
      setIsLoading(false);
    } catch (error: any) {
      alertError(error);
      setIsLoading(false);
    }
  };

  const requestListAll = async (page: number, pageSize: number) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles-sources/list/all/paging',
        {
          params: { page: page, page_size: pageSize },
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw GET_PAGE_ARTICLES_SOURCES_FAIL_MESSAGE;
      }
      setArticlesSources(data.articles_sources);
      setIsLoading(false);
    } catch (error: any) {
      alertError(error);
      setIsLoading(false);
    }
  };

  const requestCountArticlesSources = async () => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles-sources/count/total'
      );
      if (!data.success) {
        throw 'count articles sources fail';
      }
      setTotal(data.total);
      setFound(data.total);
    } catch (error: any) {
      toastifyError(error);
    }
  };

  useEffect(() => {
    requestListAll(FIRST_PAGE, PAGE_SIZE);
    requestCountArticlesSources();
  }, []);

  return (
    <div className="adminArticlesSources">
      <h1 className="adminArticlesSources__title">Manage articles sources</h1>
      <div className="adminArticlesSources__overview">
        <div className="adminArticlesSources__overview--item">
          <p>
            Total articles sources: <span>{total}</span>
          </p>
        </div>
      </div>
      <div className="adminArticlesSources__list">
        <h2 className="adminArticlesSources__list--title">
          Articles sources list
        </h2>
        <div className="adminCrawler__addBtn mb-3">
          <Button type="submit" variant="primary">
            <Link href={_ROUTES.ADD_CRAWLER}>Add new articles source</Link>
          </Button>
        </div>
        <div className="adminArticlesSources__list--search mb-3">
          <div className="d-sm-flex">
            <div className=" col-sm-10 d-flex">
              <div className="icon">
                <FontAwesomeIcon icon={faFilter} />
              </div>
              <div className="filter">
                <FilterByTopic handleChooseTopicID={handleChooseTopicID} />
              </div>
            </div>
            <div className="col-sm-2 mx-3">
              <Button variant="primary" onClick={handleSearchArticleSources}>
                Search
              </Button>
            </div>
          </div>

          <div className="my-3 col-6">
            <form onSubmit={(e) => submitSearch(e)}>
              <InputGroup>
                <Form.Control
                  placeholder="Search articles sources..."
                  type="text"
                  value={keyword}
                  onChange={(event) => setKeyword(event.target.value)}
                />
                <InputGroup.Text
                  className="searchBtn"
                  onClick={handleSearchArticleSources}
                >
                  <FontAwesomeIcon icon={faSearch} fixedWidth />
                </InputGroup.Text>
              </InputGroup>
            </form>
          </div>
        </div>
        {!isLoading ? (
          <ArticlesSourcesTable
            articlesSources={articlesSources}
            currentPage={currentPage!}
            handleDeleteArticlesSource={handleDeleteArticlesSource}
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
        <AdminArticlesSourcesPagination
          totalRows={found!}
          pageChangeHandler={pageChangeHandler}
          currentPage={currentPage}
        />
      </div>
    </div>
  );
}

export default AdminArticlesSource;
