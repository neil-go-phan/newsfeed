import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react';
import { _ROUTES } from '@/helpers/constants';
import AddCrawler from './crawler/addCrawler';
import CrawlerComponent from './crawler';
import AddCustomCrawler from './crawler/addCrawler/addCustomCrawler';
import AdminCategories from './category';
import AdminTopics from './topic';
import AdminArticles from './article';
import AdminArticlesSource from './articlesSource';
import AdminDashboard from './dashboard';
import AdminRoles from './role';
import AdminUsers from './user';

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
      case _ROUTES.ADD_CUSTOM_CRAWLER:
        return <AddCustomCrawler />;
      case _ROUTES.ADMIN_CATEGORIES:
        return <AdminCategories />;
      case _ROUTES.ADMIN_TOPICS:
        return <AdminTopics />;
      case _ROUTES.ADMIN_ARTICLES:
        return <AdminArticles />
      case _ROUTES.ADMIN_ARTICLES_SOURCE:
        return <AdminArticlesSource />
      case _ROUTES.ADMIN_ROLE:
        return <AdminRoles />
      case _ROUTES.ADMIN_USERS:
        return <AdminUsers />
      default:
        return <AdminDashboard />;
    }
  };

  return <>{render()}</>;
}

export default AdminComponent;
