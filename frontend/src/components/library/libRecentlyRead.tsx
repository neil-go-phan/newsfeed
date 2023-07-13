import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { _ROUTES } from '@/helpers/constants';
import { faClock } from '@fortawesome/free-regular-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import Link from 'next/link';
import React, { useEffect, useState } from 'react';
import RecentlyReadArticle from './recentlyReadArticle';

const FIRST_PAGE = 1;
const PAGE_SIZE = 5;
const REQUEST_RECENTLY_ARTILCES_FAIL_MESSAGE =
  'request recently read articles fail';

function LibRecentlyRead() {
  const [articles, setArticles] = useState<Articles>([]);

  useEffect(() => {
    requestGetFirstPage();
  }, []);

  const requestGetFirstPage = async () => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles/get/all/recently',
        {
          params: {
            page: FIRST_PAGE,
            page_size: PAGE_SIZE,
          },
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_RECENTLY_ARTILCES_FAIL_MESSAGE;
      }

      setArticles(data.articles);
    } catch (error: any) {
      setArticles([]);
    }
  };

  return (
    <div className="library__recentlyRead">
      <div className="title">
        <div className="text">
          <Link href={_ROUTES.RECENTLY_READ_PAGE}>
            <FontAwesomeIcon icon={faClock} />
            <span>Recently read</span>
          </Link>
        </div>

        <div className="viewall">
          <Link href={_ROUTES.RECENTLY_READ_PAGE}>View all</Link>
        </div>
      </div>
      <div className="articlesList">
        {articles.length > 0 ? (
          <div className="list">
            {articles.map((article) => (
              <RecentlyReadArticle key={`recently read ${article.title}`} article={article} />
            ))}
          </div>
        ) : (
          <div className="noRead">
            <div className="message">You are not read any article recently</div>
            <div className="btn">
              <Link href={_ROUTES.READ_FEEDS_ALL_ARTICLES}>
                <button>Start reading now !</button>
              </Link>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

export default LibRecentlyRead;
