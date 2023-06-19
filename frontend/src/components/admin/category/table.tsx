import React from 'react';
import { Table } from 'react-bootstrap';
import { Column, useTable } from 'react-table';
import { ThreeDots } from 'react-loader-spinner';
import AdminCategoriesAction from './action';

type CategoryRow = {
  index: number;
  name: string;
  id: number;
};

type Props = {
  categories: Categories;
  currentPage: number;
  handleUpdateCategory: (id: number, oldName: string, newName: string) => void;
  handleDeleteCategory: (id: number, name: string) => void;
};

const ORPHANS_CATEGORY_NAME = 'Orphans';

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
      },
      {
        header: 'Action',
        accessor: 'id',
        Cell: ({ row }) => (
          <AdminCategoriesAction
            isDisabled={checkIsOrphanCategory(row.values.name)}
            name={row.values.name}
            id={row.values.id}
            handleDeleteCategory={props.handleDeleteCategory}
            handleUpdateCategory={props.handleUpdateCategory}
          />
        ),
      },
    ],
    []
  );

  const checkIsOrphanCategory = (categoryName: string): boolean => {
    if (categoryName === ORPHANS_CATEGORY_NAME) {
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
