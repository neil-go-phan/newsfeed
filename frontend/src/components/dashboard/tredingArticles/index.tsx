import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import React, { useEffect, useState } from 'react';
import TredingArticle from './tredingArticle';

const REQUEST_TREDING_ARTICLES_FAIL_MESSAGE = 'request treding article fail'

function TredingArticles() {
  const [articles, setArticles] = useState<DashboardArticles>([])

  useEffect(() => {
    requestTredingArticles()
  }, [])
  
  const requestTredingArticles = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('/articles/get/all/treding',);
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_TREDING_ARTICLES_FAIL_MESSAGE;
      }

      setArticles(data.articles);
    } catch (error: any) {
      setArticles([]);
    }
  };
  
  return (
    <div className="wrapper">
      <div className="title">Currently treding on newsfeeds</div>
      {articles.length > 0 ? (
        <div className="articles">
          {articles.map((article) => (
            <TredingArticle article={article} key={`trending article ${article.title}`}/>
          ))}
        </div>
      ) : (
        <div className="nothingNew">
          <p>Nothing treding</p>
        </div>
      )}
    </div>
  );
}

export default TredingArticles;
