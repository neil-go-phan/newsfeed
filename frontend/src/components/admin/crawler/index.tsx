import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import {_ROUTES } from '@/helpers/constants';
import { Column, usePagination, useTable } from 'react-table';
import CrawlerAction from './crawlerAction';
import { Button, Table } from 'react-bootstrap';
import Popup from 'reactjs-popup';
import UrlModal from './addCrawler/inputUrl';
import { toastifyError } from '@/helpers/toastify';
import Link from 'next/link';

type CrawlerRow = {
  index: number;
  url: string;
  article_div: string;
  article_title: string;
  article_description: string;
  article_link: string;
  next_page: string;
  next_page_type: string;
  action: boolean;
};

type Crawler = {
  url: string;
  article_div: string;
  article_title: string;
  article_description: string;
  article_link: string;
  next_page: string;
  next_page_type: string;
};

const ERROR_OCCRUS_WHEN_LIST_CRAWLER = 'Error occurred while get list crawler'

function CrawlerComponent() {
  const [crawlers, setCrawlers] = useState<Array<Crawler>>();
  const router = useRouter()
  
  // const columns: Column<CrawlerRow>[] = React.useMemo(
  //   () => [
  //     {
  //       Header: 'STT',
  //       accessor: 'index',
  //     },
  //     {
  //       Header: 'Url',
  //       accessor: 'url',
  //     },
  //     {
  //       Header: 'Article div',
  //       accessor: 'article_div',
  //     },
  //     {
  //       Header: 'Article title',
  //       accessor: 'article_title',
  //     },
  //     {
  //       Header: 'Article description',
  //       accessor: 'article_description',
  //     },
  //     {
  //       Header: 'Article link',
  //       accessor: 'article_link',
  //     },
  //     {
  //       Header: 'Next page',
  //       accessor: 'next_page',
  //     },
  //     {
  //       Header: 'Next page type',
  //       accessor: 'next_page_type',
  //     },
  //     {
  //       Header: 'Action',
  //       accessor: 'action',
  //       Cell: ({ row }) => (
  //         <CrawlerAction
  //           url={row.values.url}
  //           handleDelete={handleDelete}
  //           // handleUpdate={handleUpdate}
  //         />
  //       ),
  //     },
  //   ],
  //   // eslint-disable-next-line react-hooks/exhaustive-deps
  //   []
  // );

  // const useCreateTableData = (crawlerRow: Array<Crawler> | undefined) => {
  //   return React.useMemo(() => {
  //     if (!crawlerRow) {
  //       return [];
  //     }
  //     return crawlerRow.map((crawler, index) => ({
  //       index: index + 1,
  //       url: crawler.url,
  //       article_div: crawler.article_div,
  //       article_title: crawler.article_title,
  //       article_description: crawler.article_description,
  //       article_link: crawler.article_link,
  //       next_page: crawler.next_page,
  //       next_page_type: crawler.next_page_type,
  //       action: false,
  //     }));
  //   }, [crawlerRow]);
  // };

  // const data = useCreateTableData(crawlers);

  // const {
  //   getTableProps,
  //   getTableBodyProps,
  //   headerGroups,
  //   prepareRow,
  //   page,
  //   pageOptions,
  //   state: { pageIndex },
  //   previousPage,
  //   nextPage,
  //   canPreviousPage,
  //   canNextPage,
  // } = useTable(
  //   {
  //     columns,
  //     data,
  //     initialState: { pageIndex: 0 },
  //   },
  //   usePagination
  // );

  const handleDelete = () => {
    requestListCrawler();
  };

  const requestListCrawler = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('crawler/list');
      setCrawlers(data.config_crawlers);
    } catch (error) {
      toastifyError(ERROR_OCCRUS_WHEN_LIST_CRAWLER)
    }
  };

  useEffect(() => {
    // requestListCrawler();
  }, [router.asPath]);

  return (
    <div className="adminCrawler">
      <h2 className="adminCrawler__list--title">Crawler list</h2>
      <div className="adminCrawler__addBtn">
        <Button
          type="submit"
          variant="primary"
        >
          <Link href={_ROUTES.ADD_CRAWLER}>Add new crawler</Link>
        </Button>
      </div>
        
      {/* <div className="adminTags__list--table mt-3">
        <Table bordered hover {...getTableProps()}>
          <thead>
            {headerGroups.map((headerGroup, index) => (
              <tr
                {...headerGroup.getHeaderGroupProps()}
                key={`crawler-admin-tr-${index}`}
              >
                {headerGroup.headers.map((column, index) => (
                  <th
                    {...column.getHeaderProps()}
                    key={`crawler-admin-tr-item-${index}`}
                  >
                    {column.render('Header')}
                  </th>
                ))}
              </tr>
            ))}
          </thead>
          <tbody {...getTableBodyProps()}>
            {page.map((row) => {
              prepareRow(row);
              return (
                <tr
                  {...row.getRowProps()}
                  key={`crawler-admin-row-tr-${row.index}`}
                >
                  {row.cells.map((cell, index) => {
                    return (
                      <td
                        {...cell.getCellProps()}
                        key={`crawler-admin-row-tr-item-${index}`}
                      >
                        {cell.render('Cell')}
                      </td>
                    );
                  })}
                </tr>
              );
            })}
          </tbody>
        </Table>
        <div className="btnPaging">
          <Button
            onClick={() => previousPage()}
            disabled={!canPreviousPage}
            variant="primary"
          >
            Previous Page
          </Button>
          <Button
            onClick={() => nextPage()}
            disabled={!canNextPage}
            variant="primary"
          >
            Next Page
          </Button>
          <p>
            Page
            <span>
              {pageIndex + 1} of {pageOptions.length}
            </span>
          </p>
        </div>
      </div> */}
    </div>
  );
}

export default CrawlerComponent;
