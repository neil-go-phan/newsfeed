import React, { useEffect, useState } from 'react';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { _ROUTES } from '@/helpers/constants';
import { Button } from 'react-bootstrap';
import Link from 'next/link';
import { alertError } from '@/helpers/alert';
import AdminCrawlersPagination from './pagination';
import { ThreeDots } from 'react-loader-spinner';
import CrawlersTable from './table';

export const PAGE_SIZE = 10;
const FIRST_PAGE = 1;
const ERROR_OCCRUS_WHEN_LIST_CRAWLER = 'Error occurred while get list crawler';

function CrawlerComponent() {
  const [crawlers, setCrawlers] = useState<Array<CrawlerTableRow>>([]);
  const [total, setTotal] = useState<number>(0);
  const [found, setFound] = useState<number>(0);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [currentPage, setCurrentPage] = useState<number>(FIRST_PAGE);

  const pageChangeHandler = (newCurrentPage: number) => {
    setCurrentPage(newCurrentPage);
    setIsLoading(true);
    requestListAllNextPage(newCurrentPage, PAGE_SIZE);
    return;
  };
  
  const handleEdit = () => {
    setIsLoading(true)
    requestListAllNextPage(currentPage, PAGE_SIZE);
  }

  const requestListAll = async (page: number, pageSize: number) => {
    try {
      const { data } = await axiosProtectedAPI.get('/crawler/list/all', {
        params: { page: page, page_size: pageSize },
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw ERROR_OCCRUS_WHEN_LIST_CRAWLER;
      }
      setCrawlers(data.crawlers);
      setTotal(data.found);
      setFound(data.found);
      setIsLoading(false);
    } catch (error: any) {
      alertError(error);
      setIsLoading(false);
    }
  };

  const requestListAllNextPage = async (page: number, pageSize: number) => {
    try {
      const { data } = await axiosProtectedAPI.get('/crawler/list/all', {
        params: { page: page, page_size: pageSize },
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw ERROR_OCCRUS_WHEN_LIST_CRAWLER;
      }
      setCrawlers(data.crawlers);
      setIsLoading(false);
    } catch (error: any) {
      alertError(error);
      setIsLoading(false);
    }
  };

  useEffect(() => {
    requestListAll(FIRST_PAGE, PAGE_SIZE);
  }, []);

  return (
    <div className="adminCrawler">
      <h1 className="adminCrawler__title">Manage crawlers</h1>
      <div className="adminCrawler__overview">
        <div className="adminCrawler__overview--item">
          <p>
            Total crawlers: <span>{total}</span>
          </p>
        </div>
      </div>
      <div className="adminCrawler__list">
        <h2 className="adminCrawler__list--title">Crawler list</h2>
        <div className="adminCrawler__addBtn mb-3">
          <Button type="submit" variant="primary">
            <Link href={_ROUTES.ADD_CRAWLER}>Add new crawler</Link>
          </Button>
        </div>
        {!isLoading ? (
          <CrawlersTable
            crawlers={crawlers}
            currentPage={currentPage!}
            handleEdit={handleEdit}
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
        <AdminCrawlersPagination
          totalRows={found!}
          pageChangeHandler={pageChangeHandler}
          currentPage={currentPage}
        />
      </div>
    </div>
  );
}

export default CrawlerComponent;
