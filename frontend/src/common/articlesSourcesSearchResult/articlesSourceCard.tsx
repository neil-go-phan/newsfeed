import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import Link from 'next/link';
import React, { useEffect, useState } from 'react';
import { Button } from 'react-bootstrap';
import ArticleCard from '../articleCard ';
type Props = {
  articlesSource: ArticlesSourceInfo;
};

const LOGO_SIZE = 60;
const REQUEST_NEWEST_ARTILCES_FAIL_MESSAGE = 'request newest article fail';
const FIRST_PAGE = 1;
const THREE_ARTICLES = 3;

const ArticleSourceCard: React.FC<Props> = (props: Props) => {
  const [articles, setArticles] = useState<Articles>([]);
  useEffect(() => {
    requestCreateCategory(props.articlesSource.id, FIRST_PAGE, THREE_ARTICLES);
  }, [props.articlesSource]);

  const requestCreateCategory = async (
    articlesSourceID: number,
    page: number,
    pageSize: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles/get-page-by-articles-source-id',
        {
          params: {
            page: page,
            page_size: pageSize,
            articles_source_id: articlesSourceID,
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

  const handleFollow = () => {
    console.log('follow id', props.articlesSource.id);
  };
  return (
    <div className="articlesSourcesSearchResult__item d-flex">
      <div className="info col-3">
        <div className="logo">
          <img
            alt="sources logo"
            src={props.articlesSource.image}
            style={{ height: 'auto' }}
            width={LOGO_SIZE}
          />
        </div>
        <div className="title">
          <Link target="_blank" href={props.articlesSource.link}>
            {props.articlesSource.title}
          </Link>
        </div>
        <div className="description colorGray">
          <p>{props.articlesSource.description}</p>
        </div>
        <div className="stats colorGray">
          <p>100k follower</p>
          <p>40 articles/week</p>
        </div>
        <div className="followBtn">
          <Button variant="primary mb-2" onClick={handleFollow}>
            Follow
          </Button>
        </div>
      </div>
      <div className="articles col-9">
        {articles.map((article) => (
          <div key={`article sources card ${article.title}`} className="article col-4">
            <ArticleCard
              article={article}
              isAdmin={true}
              articleSourceLink={props.articlesSource.link}
              articleSourceTitle={props.articlesSource.title}
              key={`articles from search feed ${article.title}`}
            />
          </div>
        ))}
      </div>
    </div>
  );
};

export default ArticleSourceCard;
