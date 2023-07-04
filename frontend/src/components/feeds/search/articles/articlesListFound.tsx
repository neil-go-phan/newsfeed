import ArticleCard from '@/common/articleCard ';
import { FollowedSourcesContext } from '@/common/contexts/followedSources';
import Grid from '@mui/material/Grid';
import React, { useContext } from 'react';

type Props = {
  articles: Articles;
};

const ArticlesListFound: React.FC<Props> = (props: Props) => {
  const {followedSources} = useContext(FollowedSourcesContext)
  const getArticlesSourceByID = (articlesSourceID: number) => {
    const source = followedSources.find(
      (articlesSource) => articlesSource.id === articlesSourceID
    );
    return source;
  };
  return (
    <div className="articlesList d-flex">
      <Grid container spacing={3}>
        {props.articles.map((article) => (
          <Grid
            item
            key={`article grid item from read feed ${article.title}`}
            xs={12}
            md={4}
          >
            <ArticleCard
              key={`articles card from read feed ${article.title}`}
              articlesSource={getArticlesSourceByID(article.articles_source_id)}
              article={article}
              isAdmin={false}
            />
          </Grid>
        ))}
      </Grid>
    </div>
  );
};

export default ArticlesListFound;
