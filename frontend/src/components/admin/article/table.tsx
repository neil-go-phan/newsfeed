import React from 'react';
import { Table } from 'react-bootstrap';
import AdminArticlesAction from './action';
type Props = {
  articles: Articles;
  currentPage: number;
  handleDeleteArticle: (id: number) => void;
};

const ArticlesTable: React.FC<Props> = (props: Props) => {
  if (props.articles.length === 0) {
    return (
      <div className="threeDotLoading">
        Not found articles
      </div>
    );
  }
  return (
      <div className="adminArticles__list--table">
        <Table responsive striped bordered hover>
          <thead>
            <tr>
              <th>#</th>
              <th>Title</th>
              <th>Description</th>
              <th>Link</th>
              <th>Published</th>
              <th>Authors</th>
              <th>Source ID</th>
              <th>Action</th>
            </tr>
          </thead>
          <tbody>
            {props.articles.map((article, index) => (
              <tr key={`article_admin_${article.title}`}>
                <td>{index + 1 + 10 * (props.currentPage - 1)}</td>
                <td>
                  <p>{article.title}</p>
                </td>
                <td>
                  <p>{article.description}</p>
                </td>
                <td>
                  <a href={article.link} target="_blank">
                    {article.link}
                  </a>
                </td>
                <td>
                  <p>{article.published}</p>
                </td>
                <td>
                  <p>{article.authors}</p>
                </td>
                <td>
                  <p>{article.articles_source_id}</p>
                </td>
                <td>
                  <AdminArticlesAction
                    id={article.id}
                    article={article}
                    handleDeleteArticle={props.handleDeleteArticle}
                  />
                </td>
              </tr>
            ))}
          </tbody>
        </Table>
      </div>

  );
};

export default ArticlesTable;
