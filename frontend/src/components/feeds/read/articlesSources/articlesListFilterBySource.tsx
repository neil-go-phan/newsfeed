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
            key={`article read feed card ${article.title}-${
              props.articlesSource!.link
            }`}
            xs={12}
            md={4}
          >
            <ArticleCard
              key={`articles from read feed ${article.title}`}
              articleSourceTitle={props.articlesSource?.title}
              articleSourceLink={props.articlesSource?.link}
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
