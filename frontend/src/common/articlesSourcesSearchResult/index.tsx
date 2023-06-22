import React from 'react'
import ArticleSourceCard from './articlesSourceCard';
type Props = {
  articlesSources: ArticlesSourceInfoes
};

const ArticlesSourcesSearchResult: React.FC<Props> = (props: Props) => {
  return (
    <div className='articlesSourcesSearchResult'>
      {props.articlesSources.map((articleSourceInfo) => <ArticleSourceCard key={`article source info - ${articleSourceInfo.title}`} articlesSource={articleSourceInfo}/>)}
    </div>
  )
}

export default ArticlesSourcesSearchResult