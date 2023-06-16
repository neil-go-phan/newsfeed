import AdminComponent from '@/components/admin';
import AdminLayout from '@/layouts/adminLayout';
import React from 'react';

export default function Admin() {
  return (
    <AdminLayout>
      <AdminComponent />
    </AdminLayout>
  );
}
