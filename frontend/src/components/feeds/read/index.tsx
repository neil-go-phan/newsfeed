import { _ROUTES } from '@/helpers/constants';
import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react'
import ReadAllArticles from './allArticles';
import ReadArticlesBySources from './articlesSources';

function ReadFeedComponent() {
  const router = useRouter();
  const [path, setpath] = useState<string>();
  useEffect(() => {
    const path = router.asPath;
    const beforeQuestionMark = path.split('?')[0];
    setpath(beforeQuestionMark);
  }, [router.asPath]);

  const render = () => {
    switch (path) {
      case _ROUTES.FEEDS_SEARCH_WEBS:
        return <ReadAllArticles />;
      case _ROUTES.READ_FEEDS_ARTICLES_SOURCE:
        return <ReadArticlesBySources />
      default:
        return <ReadAllArticles />;
    }
  };

  return (
    <div className='readFeeds'>{render()}</div>
  )
}

export default ReadFeedComponent