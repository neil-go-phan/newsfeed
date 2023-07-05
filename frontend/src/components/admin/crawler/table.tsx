import React from 'react';
import { Table } from 'react-bootstrap';
import AdminCrawlerAction from './action';
import EditScheduleBtn from './editScheduleBtn';

type Props = {
  crawlers: Array<CrawlerTableRow>;
  currentPage: number;
  handleEdit: () => void;
};

const CrawlersTable: React.FC<Props> = (props: Props) => {
  if (props.crawlers.length === 0) {
    return <div className="threeDotLoading">Not found crawlers</div>;
  }
  return (
    <div className="adminCrawler__list--table">
      <Table responsive striped bordered hover>
        <thead>
          <tr>
            <th>#</th>
            <th>Source link</th>
            <th>Feed link</th>
            <th>Crawl type</th>
            <th>Schedule</th>
            <th>Articles source id</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {props.crawlers.map((crawler, index) => (
            <tr key={`crawler_admin_${crawler.source_link}`}>
              <td>{index + 1 + 10 * (props.currentPage - 1)}</td>
              <td>
                <a href={crawler.source_link} target="_blank">
                  {crawler.source_link}
                </a>
              </td>
              <td>
                <a href={crawler.feed_link} target="_blank">
                  {crawler.feed_link}
                </a>
              </td>
              <td>
                <p>{crawler.crawl_type}</p>
              </td>
              <td>
                <EditScheduleBtn
                  id={crawler.id}
                  schedule={crawler.schedule}
                  handleEdit={props.handleEdit}
                />
              </td>
              <td>
                <p>{crawler.articles_source_id}</p>
              </td>
              <td>
                <AdminCrawlerAction crawler={crawler} />
              </td>
            </tr>
          ))}
        </tbody>
      </Table>
    </div>
  );
};

export default CrawlersTable;
