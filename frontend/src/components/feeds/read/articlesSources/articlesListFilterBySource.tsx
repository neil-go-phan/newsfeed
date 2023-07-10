import ArticleCard from '@/common/articleCard ';
import Grid from '@mui/material/Grid';
import React from 'react';

type Props = {
  articlesSource: ArticlesSourceInfo | undefined;
  articles: Articles;
};

const ArticlesListFilterBySource: React.FC<Props> = (props: Props) => {
  return (
    <div className="articlesList d-flex">
      <Grid container spacing={3}>
        {props.articles.map((article) => (
          <Grid
            item
            key={`article grid item from read feed ${article.title}`}
            xs={12}
            sm={6}
            md={4}
          >
            <ArticleCard
              key={`articles card from read feed ${article.title}`}
              articlesSource={props.articlesSource}
              article={article}
              isAdmin={false}
            />
          </Grid>
        ))}
      </Grid>
    </div>
  );
};

export default ArticlesListFilterBySource;
