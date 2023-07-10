import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { _ROUTES } from '@/helpers/constants';
import { faStar } from '@fortawesome/free-regular-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import Image from 'next/image';
import Link from 'next/link';
import React, { useEffect, useState } from 'react';
import BigLibCard from './articlesCard/bigLibCard';
import SmallLibCard from './articlesCard/smallLibCard';

const REQUEST_NEWEST_ARTILCES_FAIL_MESSAGE = 'request fail';
const PAGE = 1;
const PAGE_SIZE = 4;
const IMAGE_WIDTH = 360;
const FIRST_ARTICLE_INDEX = 0;

function LibReadlater() {
  const [articles, setArticles] = useState<Articles>([]);

  useEffect(() => {
    requestGetReadLaterArticles();
  }, []);

  const requestGetReadLaterArticles = async () => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles/get/all/followed/readlater',
        {
          params: {
            page: PAGE,
            page_size: PAGE_SIZE,
          },
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_NEWEST_ARTILCES_FAIL_MESSAGE;
      }
      setArticles(data.articles);
    } catch (error: any) {
      setArticles([]);
    }
  };

  const render = (article: Article, index: number) => {
    if (index !== FIRST_ARTICLE_INDEX) {
      return (
          <SmallLibCard article={article} />
      );
    }
  };

  return (
    <div className="library__readlater">
      <div className="title">
        <div className="text">
          <Link href={_ROUTES.FEEDS_LATER}>
            <FontAwesomeIcon icon={faStar} />
            <span>Read later</span>
          </Link>
        </div>

        <div className="viewall">
          <Link href={_ROUTES.FEEDS_LATER}>View all</Link>
        </div>
      </div>
      <div className="articlesList">
        {articles.length > 0 ? (
          <div className="foundArticles d-flex">
            <div className="left">
              <BigLibCard article={articles[FIRST_ARTICLE_INDEX]} />
            </div>
            <div className="right d-none d-md-block">
              {articles.map((article, index) => render(article, index))}
            </div>
          </div>
        ) : (
          <div className="noArticles">
            <div className="left">
              <Image
                alt="library empty images"
                src={'/images/library-empty-stars-aqua.svg'}
                width={IMAGE_WIDTH}
                height={'0'}
                style={{ height: 'auto' }}
              />
            </div>
            <div className="right">
              <div className="title">Read later</div>
              <div className="message">Star articles for reading later.</div>
              <div className="btn">
                <Link href={_ROUTES.READ_FEEDS_ALL_ARTICLES}>
                  <button>Start reading</button>
                </Link>
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

export default LibReadlater;
