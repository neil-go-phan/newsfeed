import ArticleCard from '@/common/articleCard ';
import { FollowedSourcesContext } from '@/common/contexts/followedSources';
import Grid from '@mui/material/Grid';
import React, { useContext, useState, useEffect } from 'react';

type Props = {
  articles: Articles;
};

const ArticlesListFromFollowedSource: React.FC<Props> = (props: Props) => {
  const [articlesSource, setArticlesSource] = useState();
  const { followedSources } = useContext(FollowedSourcesContext);
  useEffect(() => {}, [props.articles]);

  const getArticlesSourceByID = (
    articlesSourceID: number
  ): ArticlesSourceInfo | undefined => {
    const source = followedSources.find(
      (articlesSource) => articlesSource.id === articlesSourceID
    );
    return source;
  };
  return (
    <div className="articlesList d-flex">
      <Grid container spacing={3}>
        {props.articles.map((article) => {
          const articlesSource = getArticlesSourceByID(
            article.articles_source_id
          );
          return (
            <Grid
              item
              key={`article read feed card ${article.title}-${
                articlesSource!.link
              }`}
              xs={12}
              md={4}
            >
              <ArticleCard
                key={`articles from read feed ${article.title}`}
                articleSourceTitle={articlesSource?.title}
                articleSourceLink={articlesSource?.link}
                article={article}
                isAdmin={false}
              />
            </Grid>
          );
        })}
      </Grid>
    </div>
  );
};

export default ArticlesListFromFollowedSource;
