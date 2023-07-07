import React, { useState, useEffect } from 'react';
import { Button } from 'react-bootstrap';
import { PAGE_SIZE } from '.';

type Props = {
  totalRows: number;
  currentPage: number;
  pageChangeHandler: (currentPage: number) => void;
};

const AdminUsersPagination: React.FC<Props> = (props: Props) => {
  // Calculating max number of pages
  const noOfPages = Math.ceil(props.totalRows / PAGE_SIZE);

  const [currentPage, setCurrentPage] = useState(1);

  const [canGoBack, setCanGoBack] = useState(false);
  const [canGoNext, setCanGoNext] = useState(true);

  const onNextPage = () => setCurrentPage(currentPage + 1);
  const onPrevPage = () => setCurrentPage(currentPage - 1);

  useEffect(() => {
    if (noOfPages === currentPage) {
      setCanGoNext(false);
    } else {
      setCanGoNext(true);
    }
    if (currentPage === 1) {
      setCanGoBack(false);
    } else {
      setCanGoBack(true);
    }
  }, [noOfPages, currentPage]);
  useEffect(() => {
    props.pageChangeHandler(currentPage);
  }, [currentPage]);

  useEffect(() => {
    setCurrentPage(1);
  }, [props.totalRows]);
  
  useEffect(() => {
    setCurrentPage(props.currentPage);
  }, [props.currentPage]);
  return (
    <div className="btnPaging">
      <Button
        onClick={() => onPrevPage()}
        disabled={!canGoBack}
        variant="primary"
      >
        Previous Page
      </Button>
      <Button
        onClick={() => onNextPage()}
        disabled={!canGoNext}
        variant="primary"
      >
        Next Page
      </Button>
      <p>
        Page
        <span>
          {currentPage} of {noOfPages}
        </span>
      </p>
    </div>
  );
};

export default AdminUsersPagination;
