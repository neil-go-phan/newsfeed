import React, { useContext, useEffect, useState } from 'react'
import { FollowedSourcesContext } from '@/common/contexts/followedSources';
import ArticleCard from '@/common/articleCard ';

type Props = {
  article: Article;
};

const BigLibCard: React.FC<Props> = (props: Props) => {
  const { followedSources } = useContext(FollowedSourcesContext);
  const [articlesSource, setArticlesSource] = useState<ArticlesSourceInfo>();
  useEffect(() => {
    if (props.article.articles_source_id) {
      setArticlesSource(getArticlesSourceByID(props.article.articles_source_id));
    }
  }, []);

  const getArticlesSourceByID = (articlesSourceID: number) => {
    const source = followedSources.find(
      (articlesSource) => articlesSource.id === articlesSourceID
    );
    return source;
  };
  return (
    <ArticleCard isAdmin={false} article={props.article} articlesSource={articlesSource}/>
  )
}

export default BigLibCard