import type { ReactElement } from 'react';
import { NextPageWithLayout } from '../_app';
import FeedsLayout from '@/layouts/feedLayout';
import Library from '@/components/library';

const LibraryPage: NextPageWithLayout = () => {
  return <Library />
};

LibraryPage.getLayout = function getLayout(page: ReactElement) {
  return <FeedsLayout>{page}</FeedsLayout>;
};

export default LibraryPage;
