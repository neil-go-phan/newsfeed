import React from 'react';
import { Table } from 'react-bootstrap';
import { Column, useTable } from 'react-table';
import { ThreeDots } from 'react-loader-spinner';
import Image from 'next/image';
import AdminTopicAction from './action';
import Popup from 'reactjs-popup';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCircleQuestion } from '@fortawesome/free-solid-svg-icons';

type TopicRow = {
  index: number;
  name: string;
  id: number;
  category: Category;
};

type Props = {
  topics: Topics;
  categories: Categories;
  currentPage: number;
  handleUpdateTopic: (id: number, newName: string, newCategoryID: number) => void;
  handleDeleteTopic: (id: number, name: string, category_id: number) => void;
};

const OTHERS_CATEGORY_NAME = 'Others';
const OTHERS_TOPICS_NAME = 'Others';

const TopicsTable: React.FC<Props> = (props: Props) => {
  const columns: Column<TopicRow>[] = React.useMemo(
    () => [
      {
        header: 'STT',
        accessor: 'index',
      },
      {
        header: 'Topic',
        accessor: 'name',
        Cell: ({ row }) => (
          <>
            <p>
              {row.values.name}
              {checkIsOrphanTopic(row.values.name) ? (
                <Popup
                  trigger={() => <FontAwesomeIcon className='mx-2' icon={faCircleQuestion} />}
                  position="right center"
                  closeOnDocumentClick
                  on={['hover', 'focus']}
                >
                  <div>
                    <p>
                      This topic is not display on user interface. But article
                      sources belong to this topic can be search
                    </p>
                    <p>
                      If you detele a topic all article sources belong to that
                      deleted topic will automated become this topic child.
                    </p>
                  </div>
                </Popup>
              ) : (
                <></>
              )}
            </p>
          </>
        ),
      },
      {
        header: 'Category',
        accessor: 'category',
        Cell: ({ row }) => (
          <p>{row.values.category.name}</p>
        ),
      },
      {
        header: 'Action',
        accessor: 'id',
        Cell: ({ row }) => (
          <AdminTopicAction
            isDisabled={checkIsOrphanTopic(row.values.name)}
            topicName={row.values.name}
            id={row.values.id}
            category={row.values.category}
            categories={props.categories}
            handleDeleteTopic={props.handleDeleteTopic}
            handleUpdateTopic={props.handleUpdateTopic}
          />
        ),
      },
    ],
    []
  );

  const checkIsOrphanCategory = (categoryName: string): boolean => {
    if (categoryName === OTHERS_CATEGORY_NAME) {
      return true;
    }
    return false;
  };

  const checkIsOrphanTopic = (topicName: string): boolean => {
    if (topicName === OTHERS_TOPICS_NAME) {
      return true;
    }
    return false;
  };

  const findCategoryByID = (id: number): Category => {
    const category = props.categories.find((category) => category.id === id);
    if (category) {
      return category;
    }
    const notFoundCategory: Category = {
      id: 0,
      illustration: '',
      name: 'not found',
    };
    return notFoundCategory;
  };

  const useCreateTableData = (topics: Topics) => {
    return React.useMemo(() => {
      return topics.map((topic, index) => ({
        index: index + 1 + 10 * (props.currentPage - 1),
        name: topic.name,
        id: topic.id,
        category: findCategoryByID(topic.category_id),
      }));
    }, [topics]);
  };

  const data = useCreateTableData(props.topics);
  const { getTableProps, getTableBodyProps, headerGroups, prepareRow, rows } =
    useTable({
      columns,
      data,
    });
  if (props.topics.length === 0) {
    return (
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
    );
  }
  return (
    <>
      <div className="adminTopics__list--table">
        <Table bordered hover {...getTableProps()}>
          <thead>
            {headerGroups.map((headerGroup, index) => (
              <tr
                {...headerGroup.getHeaderGroupProps()}
                key={`topics-admin-collum-${index}`}
              >
                {headerGroup.headers.map((column) => (
                  <th
                    {...column.getHeaderProps()}
                    key={`topics-admin-collum-${column.render('header')}}`}
                  >
                    {column.render('header')}
                  </th>
                ))}
              </tr>
            ))}
          </thead>
          <tbody {...getTableBodyProps()}>
            {rows.map((row, i) => {
              prepareRow(row);
              return (
                <tr
                  {...row.getRowProps()}
                  key={`topics-admin-tr-${i}-${row.values.name}`}
                >
                  {row.cells.map((cell, i) => {
                    return (
                      <td
                        {...cell.getCellProps()}
                        key={`topics-admin-td-${i}-${row.values.name}`}
                      >
                        {cell.render('Cell')}
                      </td>
                    );
                  })}
                </tr>
              );
            })}
          </tbody>
        </Table>
      </div>
    </>
  );
};

export default TopicsTable;
