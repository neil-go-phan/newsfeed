import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import React, { useEffect, useState } from 'react';
import { ThreeDots } from 'react-loader-spinner';
import AdminArticlesPagination from './pagination';
import { alertError, alertSuccess } from '@/helpers/alert';
import { toastifyError } from '@/helpers/toastify';
import ArticlesTable from './table';
import { Button, Form, InputGroup } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faFilter, faSearch } from '@fortawesome/free-solid-svg-icons';
import DropdownFilterBySources from './dropdownFilterBySources';

export const PAGE_SIZE = 10;
const FIRST_PAGE = 1;
const DELETE_ARTICLE_FAIL_MESSAGE = 'delete article fail';
const DELETE_ARTICLE_SUCCESS_MESSAGE = 'delete success';
const GET_PAGE_ARTICLES_FAIL_MESSAGE = 'request list fail';

function AdminArticles() {
  const [articles, setArticles] = useState<Articles>([]);
  const [currentPage, setCurrentPage] = useState<number>(FIRST_PAGE);
  const [keyword, setKeyword] = useState<string>('');
  const [sourceID, setSourceID] = useState<number>(0);
  const [total, setTotal] = useState<number>(0);
  const [found, setFound] = useState<number>(0);

  const handleDeleteArticle = (id: number) => {
    requestDeleteArticle(id);
  };

  const handleChooseSourceID = (id:number) => {
    setSourceID(id)
    setKeyword('')
  }

  const pageChangeHandler = (newCurrentPage: number) => {
    setCurrentPage(newCurrentPage);
    if (keyword === '' && sourceID === 0) {
      requestListAll(newCurrentPage, PAGE_SIZE);
      setFound(total)
      return
    }
    requestSearchNextPage(keyword, newCurrentPage, PAGE_SIZE, sourceID)
    return
  };

  const submitSearch = (e:React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    handleSearchArticle();
  }

  const handleSearchArticle = () => {
    setCurrentPage(FIRST_PAGE)
    if (keyword === '' && sourceID === 0) {
      requestListAll(FIRST_PAGE, PAGE_SIZE);
      setFound(total)
      return
    }
    requestFirstSearch(keyword, FIRST_PAGE, PAGE_SIZE, sourceID)
  };

  const requestFirstSearch = async (keyword: string, page: number, pageSize: number, sourceid: number) => {
    try {
      const { data } = await axiosProtectedAPI.get('/articles/search/admin', {
        params: {q:keyword, page: page, page_size: pageSize , source_id: sourceid},
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw GET_PAGE_ARTICLES_FAIL_MESSAGE;
      }
      setArticles(data.articles);
      setFound(data.found)
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestSearchNextPage = async (keyword: string, page: number, pageSize: number, sourceid: number) => {
    try {
      const { data } = await axiosProtectedAPI.get('/articles/search/admin', {
        params: {q:keyword, page: page, page_size: pageSize , source_id: sourceid},
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw GET_PAGE_ARTICLES_FAIL_MESSAGE;
      }
      setArticles(data.articles);
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestDeleteArticle = async (id: number) => {
    try {
      const { data } = await axiosProtectedAPI.post('/articles/delete/id', {
        id: id,
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw DELETE_ARTICLE_FAIL_MESSAGE;
      }
      alertSuccess(DELETE_ARTICLE_SUCCESS_MESSAGE);
      requestListAll(currentPage, PAGE_SIZE);
      requestCountArticles();
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestListAll = async (page: number, pageSize: number) => {
    try {
      const { data } = await axiosProtectedAPI.get('/articles/list/all', {
        params: { page: page, page_size: pageSize },
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw GET_PAGE_ARTICLES_FAIL_MESSAGE;
      }
      setArticles(data.articles);
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestCountArticles = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('/articles/count/total');
      if (!data.success) {
        throw 'count articles fail';
      }
      setTotal(data.total);
      setFound(data.total);
    } catch (error: any) {
      toastifyError(error);
    }
  };

  useEffect(() => {
    requestListAll(FIRST_PAGE, PAGE_SIZE);
    requestCountArticles();
  }, []);

  return (
    <div className="adminArticles">
      <h1 className="adminArticles__title">Manage Articles</h1>
      <div className="adminArticles__overview">
        <div className="adminArticles__overview--item">
          <p>
            Total Articles: <span>{total}</span>
          </p>
        </div>
      </div>
      <div className="adminArticles__list">
        <h2 className="adminArticles__list--title">Articles list</h2>
        <div className="adminArticles__list--search d-sm-flex mb-3">
          <div className="col-sm-4">
            <form onSubmit={(e) => submitSearch(e)}>
              <InputGroup>
                <Form.Control
                  placeholder="Search article..."
                  type="text"
                  required
                  value={keyword}
                  onChange={(event) => setKeyword(event.target.value)}
                />
                <InputGroup.Text
                  className="searchBtn"
                  onClick={handleSearchArticle}
                >
                  <FontAwesomeIcon icon={faSearch} fixedWidth />
                </InputGroup.Text>
              </InputGroup>
            </form>
          </div>
          <div className=" col-sm-6 d-flex">
            <div className="icon">
            <FontAwesomeIcon icon={faFilter} />
            </div>
            <div className="filter">
              <DropdownFilterBySources handleChooseSourceID={handleChooseSourceID}/>
            </div>
          </div>
          <div className="col-sm-2">
            <Button variant='primary' onClick={handleSearchArticle}>Search</Button>
          </div>
        </div>
        {articles ? (
          <ArticlesTable
            articles={articles}
            currentPage={currentPage!}
            handleDeleteArticle={handleDeleteArticle}
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
        <AdminArticlesPagination
          totalRows={found!}
          pageChangeHandler={pageChangeHandler}
          currentPage={currentPage}
        />
      </div>
    </div>
  );
}

export default AdminArticles;
