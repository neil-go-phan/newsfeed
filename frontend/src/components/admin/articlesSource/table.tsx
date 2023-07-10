import React from 'react';
import { Table } from 'react-bootstrap';
import Image from 'next/image';
import AdminArticlesSourcesAction from './action';

type Props = {
  articlesSources: ArticlesSourceInfoes;
  currentPage: number;
  handleDeleteArticlesSource: (id: number) => void;
  handleUpdate: (articlesSource: UpdateArticleSourcePayload) => void;
};

const IMAGE_SIZE_PIXEL = 80;

const ArticlesSourcesTable: React.FC<Props> = (props: Props) => {
  if (props.articlesSources.length === 0) {
    return <div className="threeDotLoading">Not found articles sources</div>;
  }
  return (
    <div className="adminArticlesSources__list--table">
      <Table responsive striped bordered hover>
        <thead>
          <tr>
            <th>#</th>
            <th>Title</th>
            <th>Description</th>
            <th>Link</th>
            <th>Feed link</th>
            <th>Image</th>
            <th>Follower</th>
            <th>Topic id</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {props.articlesSources.map((articlesSource, index) => (
            <tr key={`article_source_admin_${articlesSource.title}`}>
              <td>{index + 1 + 10 * (props.currentPage - 1)}</td>
              <td>
                <p>{articlesSource.title}</p>
              </td>
              <td>
                <p>{articlesSource.description}</p>
              </td>
              <td>
                <a href={articlesSource.link} target="_blank">
                  {articlesSource.link}
                </a>
              </td>
              <td>
                <a href={articlesSource.feed_link} target="_blank">
                  {articlesSource.feed_link}
                </a>
              </td>
              <td>
                <Image
                  alt="article source logo image"
                  src={articlesSource.image}
                  width={IMAGE_SIZE_PIXEL}
                  height="0"
                  style={{ height: 'auto' }}
                />
              </td>
              <td>
                <p>{articlesSource.follower}</p>
              </td>
              <td>
                <p>{articlesSource.follower}</p>
              </td>
              <td>
                <AdminArticlesSourcesAction
                  id={articlesSource.id}
                  articlesSource={articlesSource}
                  handleDeleteArticlesSource={props.handleDeleteArticlesSource}
                  handleUpdate={props.handleUpdate}
                />
              </td>
            </tr>
          ))}
        </tbody>
      </Table>
    </div>
  );
};

export default ArticlesSourcesTable;
