import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { toastifyError } from '@/helpers/toastify';
import React, { useEffect, useState } from 'react';
import { ThreeDots } from 'react-loader-spinner';
import ArticlesSource from './articlesSource';
import ArticleCard from '@/common/articleCard ';
import Grid from '@mui/material/Grid';
import { CRAWLER_FEED_TYPE } from '..';
import { Table } from 'react-bootstrap';

type Props = {
  url: string;
  testType: string;
  crawler: Crawler | undefined;
  handleSubmitArticleSource: (
    articlesSource: ArticlesSource,
    topicName: string
  ) => void;
};

const ERROR_MESSAGE_WHEN_TEST_FAIL = 'Test fail';

const TestResult: React.FC<Props> = (props: Props) => {
  const [isLoading, setIsloading] = useState<boolean>(false);
  const [articles, setArticles] = useState<Articles>();
  const [articlesSource, setArticlesSource] = useState<ArticlesSource>();
  const [errorMessage, setErrorMessage] = useState<string>('');

  const requestTestCrawlerRSS = async (url: string) => {
    try {
      const res = await axiosProtectedAPI.post('crawler/test/rss', {
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

  const requestTestCustomCrawler = async (crawler: Crawler) => {
    try {
      const res = await axiosProtectedAPI.post('crawler/test/custom', {
        source_link: crawler.source_link,
        crawl_type: crawler.crawl_type,
        article_div: crawler.article_div,
        article_title: crawler.article_title,
        article_description: crawler.article_description,
        article_link: crawler.article_link,
        article_published: crawler.article_published,
        article_authors: crawler.article_authors,
      });
      setArticles(res.data.articles);
      const articlesource: ArticlesSource = res.data.articles_source;
      setArticlesSource({ ...articlesource, feed_link: articlesource.link });
      setIsloading(false);
    } catch (error: any) {
      toastifyError(ERROR_MESSAGE_WHEN_TEST_FAIL);
      setIsloading(false);
      setErrorMessage(error.data.message || ERROR_MESSAGE_WHEN_TEST_FAIL);
    }
  };

  const handleSubmitArticleSource = (
    articlesSource: ArticlesSource,
    topicName: string
  ) => {
    props.handleSubmitArticleSource(articlesSource, topicName);
  };

  useEffect(() => {
    if (props.url) {
      setArticles(undefined);
      setArticlesSource(undefined);
      setErrorMessage('');
      setIsloading(true);
      if (props.testType === CRAWLER_FEED_TYPE) {
        requestTestCrawlerRSS(props.url);
      }
    }
  }, [props.url]);

  useEffect(() => {
    if (props.crawler) {
      setArticles(undefined);
      setArticlesSource(undefined);
      setErrorMessage('');
      setIsloading(true);
      requestTestCustomCrawler(props.crawler);
    }
  }, [props.crawler]);

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
            <h2>Test result</h2>
            <div className="lineWraper">
              <div className="line"></div>
            </div>
          </div>
          <div
            className={
              errorMessage === ''
                ? 'addCrawler__testResult--error d-none'
                : 'addCrawler__testResult--error d-block'
            }
          >
            <div className="notFound">
              <p>{errorMessage}</p>
            </div>
          </div>
          <div className="addCrawler__testResult--articles_source">
            <div className="title">
              <h3>Articles source</h3>
            </div>
            {articlesSource ? (
              <ArticlesSource
                url={props.url}
                articlesSource={articlesSource}
                handleSubmit={handleSubmitArticleSource}
              />
            ) : (
              <div className="notFound">
                <p>not found article source</p>
              </div>
            )}
          </div>

          <div className="addCrawler__testResult--articles">
            <div className="title">
              <h3>Articles</h3>
            </div>
            {articles ? (
              <div className="articleFound">
                <div className="table">
                  <h3>List</h3>
                  <Table responsive striped bordered hover>
                    <thead>
                      <tr>
                        <th>#</th>
                        <th>Title</th>
                        <th>Description</th>
                        <th>Link</th>
                        <th>Published</th>
                        <th>Authors</th>
                      </tr>
                    </thead>
                    <tbody>
                      {articles.map((article, index) => (
                        <tr key={`article_crawler_test_${article.title}`}>
                          <td>{index}</td>
                          <td>{article.title}</td>
                          <td>
                            <p>{article.description}</p>
                          </td>
                          <td>
                            <a href={article.link} target="_blank">
                              {article.link}
                            </a>
                          </td>
                          <td>{article.published}</td>
                          <td>{article.authors}</td>
                        </tr>
                      ))}
                    </tbody>
                  </Table>
                </div>
                <div className="preview">
                  <h3>Preview</h3>
                  <Grid container spacing={3}>
                    {articles.map((article) => (
                      <Grid item key={article.title} xs={12} md={4}>
                        <ArticleCard
                          key={`${article.title}-card`}
                          articlesSource={articlesSource}
                          article={article}
                          isAdmin={true}
                        />
                      </Grid>
                    ))}
                  </Grid>
                </div>
              </div>
            ) : (
              <div className="notFound">
                <p>not found article</p>
              </div>
            )}
          </div>
        </>
      )}
    </div>
  );
};

export default TestResult;
