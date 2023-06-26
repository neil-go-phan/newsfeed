import React from 'react';
import ArticleSourceCard from './articlesSourceCard';
type Props = {
  articlesSources: ArticlesSourceInfoes;
};

const ArticlesSourcesSearchResult: React.FC<Props> = (props: Props) => {  
  if (props.articlesSources.length !== 0) {
    return (
      <>
        {props.articlesSources.map((articleSourceInfo) => (
          <ArticleSourceCard
            key={`article source info - ${articleSourceInfo.title}`}
            articlesSource={articleSourceInfo}
          />
        ))}
      </>
    );
  }

  return (
    <div className="articlesSourcesSearchResult__notFound">Not found feeds</div>
  );
};

export default ArticlesSourcesSearchResult;
