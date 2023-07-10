import React from 'react';
import { Table } from 'react-bootstrap';
import Popup from 'reactjs-popup';
import PermissionColumn from './permissionColumn';
import AdminRolesAction from './action';

type Props = {
  roles: Array<Role>;
  currentPage: number;
  handleDelete: (id: number) => void;
  handleUpdate: (role: Role) => void;
};

const RolesTable: React.FC<Props> = (props: Props) => {
  if (props.roles.length === 0) {
    return <div className="threeDotLoading">Not found roles</div>;
  }
  return (
    <div className="adminRoles__list--table">
      <Table responsive striped bordered hover>
        <thead>
          <tr>
            <th>#</th>
            <th>Name</th>
            <th>Description</th>
            <th>Permissions</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {props.roles.map((role, index) => (
            <tr key={`roles_admin_${role.name}`}>
              <td>{index + 1 + 10 * (props.currentPage - 1)}</td>
              <td>
                <span>{role.name}</span>
              </td>
              <td>
                <span>{role.description}</span>
              </td>
              <td>
                <div className="row">
                  <PermissionColumn permissions={role.permissions}/>
                </div>
              </td>
              <td>
                <AdminRolesAction
                  role={role}
                  handleDelete={props.handleDelete}
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

export default RolesTable;
