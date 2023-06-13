import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react';
import { _ROUTES } from '@/helpers/constants';
import AddCrawler from './crawler/addCrawler';
import CrawlerComponent from './crawler';

function AdminComponent() {
  const router = useRouter();
  const [path, setpath] = useState<string>();
  useEffect(() => {
    const path = router.asPath;
    const beforeQuestionMark = path.split('?')[0];
    setpath(beforeQuestionMark);
  }, [router.asPath]);
  const render = () => {
    switch (path) {
      case _ROUTES.ADMIN_CRAWLER:
        return <CrawlerComponent />;
      case _ROUTES.ADD_CRAWLER:
        return <AddCrawler />;
      default:
        return <CrawlerComponent />
    }
  };

  return <>{render()}</>;
}

export default AdminComponent;
