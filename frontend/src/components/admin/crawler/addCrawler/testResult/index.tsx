import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { toastifyError } from '@/helpers/toastify';
import React, { useEffect, useState } from 'react';
import { ThreeDots } from 'react-loader-spinner';
import ArticlesSource from './articlesSource';
import ArticleCard from '@/common/articleCard ';
import Grid from '@mui/material/Grid';

type Props = {
  url: string;
};

const ERROR_MESSAGE_WHEN_TEST_FAIL = 'Test fail';

const TestResult: React.FC<Props> = (props: Props) => {
  const [isLoading, setIsloading] = useState<boolean>(false);
  const [articles, setArticles] = useState<Articles>();
  const [articlesSource, setArticlesSource] = useState<ArticlesSource>();
  const [errorMessage, setErrorMessage] = useState<string>('');

  const requestTestCrawlerRSS = async (url: string) => {
    try {
      const res = await axiosProtectedAPI.post('crawler/test-rss', {
        source_link: url,
      });
      if (res?.data.success) {
        setArticles(res.data.articles);
        setArticlesSource(res.data.articles_source);
        setIsloading(false);
        setErrorMessage('');
      }
      if (!res?.data.success) {
        throw res;
      }
    } catch (error: any) {
      setIsloading(false);
      toastifyError(ERROR_MESSAGE_WHEN_TEST_FAIL);
      setErrorMessage(error.data.message || ERROR_MESSAGE_WHEN_TEST_FAIL);
    }
  };
  console.log(articles);
  useEffect(() => {
    if (props.url) {
      setIsloading(true);
      requestTestCrawlerRSS(props.url);
    }
  }, [props.url]);
  return (
    <div className="addCrawler__testResult">
      {isLoading ? (
        <div className="threeDotLoading">
          <ThreeDots
            height="50"
            width="50"
            radius="9"
            color="#4fa94d"
            ariaLabel="three-dots-loading"
            visible={true}
          />
        </div>
      ) : (
        <>
          <div className="addCrawler__testResult--title">
            <h1>Test result</h1>
          </div>
          <div
            className={
              errorMessage === ''
                ? 'addCrawler__testResult--error d-none'
                : 'addCrawler__testResult--error d-block'
            }
          >
            <p>
              <span>Error: </span>
              {errorMessage}
            </p>
          </div>
          <ArticlesSource url={props.url} articlesSource={articlesSource} />
          <div className="addCrawler__testResult--articles">
            {articles ? (
              <Grid container spacing={3}>
                {articles.map((article) => (
                  <Grid item key={article.title} xs={12} md={4}>
                    <ArticleCard
                      articleSourceTitle={articlesSource?.title}
                      article={article}
                      isAdmin={true}
                    />
                  </Grid>
                ))}
              </Grid>
            ) : (
              <></>
            )}
          </div>
        </>
      )}
    </div>
  );
};

export default TestResult;
