import React, { useEffect, useState } from 'react';
import FeedsRecommented from './feeds';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';

const REQUEST_RECOMMENDED_SOURCES_FAIL_MESSAGE = 'request recommended articles sources fail'

function RecommendedFeeds() {
  const [articlesSources, setArticlesSources] = useState<ArticlesSourceInfoes>(
    []
  );

  useEffect(() => {
    requestRecommendedSources()
  }, [])
  
  const requestRecommendedSources = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('/articles-sources/get/most-active',);
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_RECOMMENDED_SOURCES_FAIL_MESSAGE;
      }

      setArticlesSources(data.articles_sources);
    } catch (error: any) {
      setArticlesSources([]);

    }
  };
  return (
    <div className="wrapper">
      <div className="title">Recommended feeds</div>
      {articlesSources.length > 0 ? (
        <div className="listFeeds">
          {articlesSources.map((source) => (
            <FeedsRecommented articlesSource={source} key={`feed recommended ${source.title}`}/>
          ))}
        </div>
      ) : (
        <div className="nothingNew">
          <p>Nothing to recommended</p>
        </div>
      )}
    </div>
  );
}

export default RecommendedFeeds;
