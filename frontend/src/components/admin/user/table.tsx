import React from 'react';
import { Table } from 'react-bootstrap';
import DeleteBtn from './deleteBtn';
import UpdateUserRole from './updateRole';

type Props = {
  users: Array<User>;
  currentPage: number;
  handleDelete: (id: number) => void;
  handleUpdate: (id: number, newRole: string) => void;
};

const UsersTable: React.FC<Props> = (props: Props) => {
  if (props.users.length === 0) {
    return <div className="threeDotLoading">Not found users</div>;
  }
  return (
    <div className="adminUsers__list--table">
      <Table responsive striped bordered hover>
        <thead>
          <tr>
            <th>#</th>
            <th>Username</th>
            <th>Email</th>
            <th>Role</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {props.users.map((user, index) => (
            <tr key={`users_admin_${user.username}`}>
              <td>{index + 1 + 10 * (props.currentPage - 1)}</td>
              <td>
                <span>{user.username}</span>
              </td>
              <td>
                <span>{user.email}</span>
              </td>
              <td><UpdateUserRole handleUpdate={props.handleUpdate} user={user}/></td>
              <td>
                <DeleteBtn handleDelete={props.handleDelete} user={user} />
              </td>
            </tr>
          ))}
        </tbody>
      </Table>
    </div>
  );
};

export default UsersTable;
