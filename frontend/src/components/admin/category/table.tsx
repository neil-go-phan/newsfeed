import React from 'react';
import { Table } from 'react-bootstrap';
import { Column, useTable } from 'react-table';
import { ThreeDots } from 'react-loader-spinner';
import Image from 'next/image';
import AdminCategoriesAction from '../category/action';
import Popup from 'reactjs-popup';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCircleQuestion } from '@fortawesome/free-solid-svg-icons';

type CategoryRow = {
  index: number;
  name: string;
  id: number;
  illustration: string;
};

type Props = {
  categories: Categories;
  currentPage: number;
  handleUpdateCategory: (id: number, oldName: string, newName: string) => void;
  handleDeleteCategory: (id: number, name: string) => void;
};

const OTHERS_CATEGORY_NAME = 'Others';
const IMAGE_SIZE_PIXEL = 100;

const CategoriseTable: React.FC<Props> = (props: Props) => {
  const columns: Column<CategoryRow>[] = React.useMemo(
    () => [
      {
        header: 'STT',
        accessor: 'index',
      },
      {
        header: 'Category',
        accessor: 'name',
        Cell: ({ row }) => (
          <>
            <p>
              {row.values.name}
              {checkIsOrphanCategory(row.values.name) ? (
                <Popup
                  trigger={() => (
                    <FontAwesomeIcon className="mx-2" icon={faCircleQuestion} />
                  )}
                  position="right center"
                  closeOnDocumentClick
                  on={['hover', 'focus']}
                >
                  <div>
                    <p>This category is not display on user interface.</p>
                    <p>
                      If you detele a category all topics belong to that deleted
                      category will automated become this category child.
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
        header: 'Illustration',
        accessor: 'illustration',
        Cell: ({ row }) => (
          <Image
            alt="category illustration image"
            src={row.values.illustration}
            width={IMAGE_SIZE_PIXEL}
            height="0"
            style={{ height: 'auto' }}
          />
        ),
      },
      {
        header: 'Action',
        accessor: 'id',
        Cell: ({ row }) => (
          <AdminCategoriesAction
            isDisabled={checkIsOrphanCategory(row.values.name)}
            name={row.values.name}
            id={row.values.id}
            illustration={row.values.illustration}
            handleDeleteCategory={props.handleDeleteCategory}
            handleUpdateCategory={props.handleUpdateCategory}
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

  const useCreateTableData = (categories: Categories | undefined) => {
    return React.useMemo(() => {
      if (!categories) {
        return [];
      }
      return categories.map((category, index) => ({
        index: index + 1 + 10 * (props.currentPage - 1),
        name: category.name,
        id: category.id,
        illustration: category.illustration,
      }));
    }, [categories]);
  };

  const data = useCreateTableData(props.categories);
  const { getTableProps, getTableBodyProps, headerGroups, prepareRow, rows } =
    useTable({
      columns,
      data,
    });
  if (props.categories.length === 0) {
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
      <div className="adminCategories__list--table">
        <Table bordered hover {...getTableProps()}>
          <thead>
            {headerGroups.map((headerGroup, index) => (
              <tr
                {...headerGroup.getHeaderGroupProps()}
                key={`categories-admin-collum-${index}`}
              >
                {headerGroup.headers.map((column) => (
                  <th
                    {...column.getHeaderProps()}
                    key={`categories-admin-collum-${column.render('header')}}`}
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
                  key={`categories-admin-tr-${i}-${row.values.name}`}
                >
                  {row.cells.map((cell, i) => {
                    return (
                      <td
                        {...cell.getCellProps()}
                        key={`categories-admin-td-${i}-${row.values.name}`}
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

export default CategoriseTable;
